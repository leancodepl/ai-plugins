#!/usr/bin/env python3
# AI-Provenance:
#   model: Claude Opus 5
#   harness: Claude Code

"""Ask Jev, TypeSafe's System One model, for typed judgments.

One endpoint, one POST: <base>/v1/systemone. The base is TypeSafe's API unless
JEV_BASE_URL points at a proxy that serves the same route, such as LiteLLM. The
key, the base and the default model are read from ~/.config/typesafe/env, and a
variable already in the environment wins over the file. The key goes on the
request object as a header.

Do not rewrite this as a `curl -H "Authorization: Bearer $KEY"` wrapper. A
header passed to curl is an argv element, readable by any `ps` on the machine
for the lifetime of the request. urllib keeps it in this process's memory.
"""

from __future__ import annotations

import argparse
import json
import os
import random
import re
import sys
import time
import urllib.error
import urllib.parse
import urllib.request
from pathlib import Path
from typing import Any

DEFAULT_BASE_URL = "https://api.typesafe.ai"
ENDPOINT_PATH = "/v1/systemone"
DEFAULT_MODEL = "jev-latest"
KEY_FILE = Path.home() / ".config" / "typesafe" / "env"
KEY_VAR = "TYPESAFE_API_KEY"
BASE_URL_VAR = "JEV_BASE_URL"
MODEL_VAR = "JEV_MODEL"
SETTING_LINE = re.compile(r"^\s*(?:export\s+)?([A-Z_][A-Z0-9_]*)\s*=\s*(.+?)\s*$")
RETRYABLE_STATUS = {429, 502, 503, 504, 529}
MAX_ATTEMPTS = 4
QUESTION_TYPES = {"noul", "score", "choice"}
MAX_CHOICE_OPTIONS = 255
MAX_SCORE_LEVELS = 10
# Past this many characters, an unnamed blob of text is almost never one thing.
MAX_UNSTRUCTURED_STATE = 400


class JevError(RuntimeError):
    """A safe, user-facing error. Never carries the key."""


def read_setting(name: str) -> tuple[str, str] | None:
    """Return (value, source) for one setting, or None when it is set nowhere.

    Environment wins so a one-off export is honoured; the env file comes next.
    """
    from_env = os.environ.get(name, "").strip()
    if from_env:
        return from_env, "environment"
    if not KEY_FILE.exists():
        return None
    for line in KEY_FILE.read_text().splitlines():
        match = SETTING_LINE.match(line)
        if match and match.group(1) == name:
            value = match.group(2).strip().strip("'\"")
            if value:
                return value, str(KEY_FILE)
    return None


def read_key() -> tuple[str, str]:
    """Return (key, source)."""
    found = read_setting(KEY_VAR)
    if found:
        return found
    if not KEY_FILE.exists():
        raise JevError(
            f"No API key. Write one to {KEY_FILE}:\n"
            f"  mkdir -p {KEY_FILE.parent} && chmod 700 {KEY_FILE.parent}\n"
            f"  printf 'export {KEY_VAR}=%s\\n' 'your-key' > {KEY_FILE}"
            f" && chmod 600 {KEY_FILE}"
        )
    raise JevError(f"{KEY_FILE} has no {KEY_VAR} line. Expected: export {KEY_VAR}=your-key")


def read_api_url() -> tuple[str, str]:
    """Return (endpoint URL, source). A base given with or without /v1 both work."""
    found = read_setting(BASE_URL_VAR)
    base, source = found if found else (DEFAULT_BASE_URL, "default")
    base = base.rstrip("/")
    if base.endswith(ENDPOINT_PATH):
        base = base[: -len(ENDPOINT_PATH)]
    elif base.endswith("/v1"):
        base = base[: -len("/v1")]
    if not base.startswith(("https://", "http://")):
        raise JevError(f"{BASE_URL_VAR} must be an http(s) URL, got {base!r} from {source}")
    return base + ENDPOINT_PATH, source


def read_default_model() -> tuple[str, str]:
    """Return (model, source). Proxies name models by their own aliases."""
    found = read_setting(MODEL_VAR)
    return found if found else (DEFAULT_MODEL, "default")


def redact(text: str, key: str | None) -> str:
    return text.replace(key, "<redacted>") if key else text


def load_json_arg(inline: str | None, path: Path | None, label: str) -> Any:
    raw = path.read_text() if path is not None else inline
    if raw is None:
        raise JevError(f"Provide --{label} or --{label}-file")
    try:
        return json.loads(raw)
    except json.JSONDecodeError as error:
        where = str(path) if path is not None else f"--{label}"
        raise JevError(f"{where} is not valid JSON: {error}") from error


def build_state(inline: str | None, path: Path | None, fmt: str) -> Any:
    """State reaches the API as a string or as a JSON object/array.

    A file keeps a large diff or log out of the transcript entirely, which is
    the point of --state-file.
    """
    if path is not None:
        raw = path.read_text()
    elif inline is not None:
        raw = inline
    else:
        raise JevError("Provide --state or --state-file")

    if fmt == "text":
        return raw
    try:
        parsed = json.loads(raw)
    except json.JSONDecodeError:
        if fmt == "json":
            raise JevError(
                "State is not valid JSON, and --state-format json was requested"
            ) from None
        reject_unnamed_blob(raw)
        return raw
    if fmt == "json" or isinstance(parsed, (dict, list)):
        return parsed
    reject_unnamed_blob(raw)
    return raw


def reject_unnamed_blob(raw: str) -> None:
    """Refuse a long stretch of text that was never given named parts.

    Several items flattened into one string is the framing mistake that costs
    the most: the questions end up referring to them in prose, the answers
    degrade, and nothing in the response says why. Passing --state-format text
    makes the single-blob case a deliberate answer rather than a default.
    """
    if len(raw) <= MAX_UNSTRUCTURED_STATE:
        return
    raise JevError(
        f"State is {len(raw)} characters of unnamed text. A state with more than one part "
        "belongs in a JSON object, so each part has a name a question can point at with a "
        "backticked path — see https://docs.typesafe.ai/concepts/state.md. If this really "
        "is one indivisible blob, such as a diff or a log, pass --state-format text."
    )


def validate_questions(questions: Any) -> None:
    """Catch locally what the API would otherwise reject with a 400 or 422."""
    if not isinstance(questions, dict) or not questions:
        raise JevError(
            'Questions must be a non-empty object: {"id": {"type": ..., "instructions": ...}}'
        )

    for qid, question in questions.items():
        where = f"question {qid!r}"
        if not isinstance(question, dict):
            raise JevError(f"{where} must be an object")
        qtype = question.get("type")
        if qtype not in QUESTION_TYPES:
            raise JevError(
                f"{where} has type {qtype!r}; expected one of {sorted(QUESTION_TYPES)}"
            )
        if not str(question.get("instructions", "")).strip():
            raise JevError(
                f"{where} has no instructions. The model never sees the question id, "
                "so the instructions carry the whole meaning."
            )
        criteria = question.get("criteria")
        if qtype == "noul":
            # Optional, and the API takes it: a map describing the true and false sides.
            if criteria is not None and not isinstance(criteria, dict):
                raise JevError(
                    f"{where} is a noul: criteria are optional, and when present must be a map "
                    f'such as {{"true": ..., "false": ...}}'
                )
        elif qtype == "score":
            if not isinstance(criteria, list) or not criteria:
                raise JevError(
                    f"{where} is a score: criteria must be an ordered array of level "
                    f"descriptions. The response `legend` comes back keyed by index; "
                    f"the request does not take that shape."
                )
            if len(criteria) > MAX_SCORE_LEVELS:
                raise JevError(
                    f"{where} has {len(criteria)} levels; the maximum is {MAX_SCORE_LEVELS}"
                )
        elif qtype == "choice":
            if not isinstance(criteria, dict) or not criteria:
                raise JevError(
                f"{where} is a choice: criteria must be a map of option id to description"
            )
            if len(criteria) > MAX_CHOICE_OPTIONS:
                raise JevError(
                    f"{where} has {len(criteria)} options; the maximum is {MAX_CHOICE_OPTIONS}"
                )


def post(url: str, payload: dict[str, Any], key: str, timeout: float) -> dict[str, Any]:
    body = json.dumps(payload).encode("utf-8")
    host = urllib.parse.urlsplit(url).netloc
    last_error: str | None = None

    for attempt in range(1, MAX_ATTEMPTS + 1):
        request = urllib.request.Request(
            url,
            data=body,
            method="POST",
            headers={
                # On the Request object, never in argv — see the module docstring.
                "Authorization": f"Bearer {key}",
                "Content-Type": "application/json",
                "Accept": "application/json",
            },
        )
        try:
            with urllib.request.urlopen(request, timeout=timeout) as response:
                return json.loads(response.read().decode("utf-8"))
        except urllib.error.HTTPError as error:
            detail = redact(error.read().decode("utf-8", "replace").strip(), key)
            if error.code in (401, 403):
                raise JevError(
                    f"{error.code} from {host}: the key was rejected, or it is not allowed to "
                    f"call model {payload.get('model')!r}. Check the key is current for this "
                    "endpoint and that a proxy key lists the model, then run `key-status` to "
                    f"see which sources this script read. {detail}"
                ) from None
            if error.code in (400, 422):
                # TypeSafe answers a malformed request with 422; LiteLLM with 400.
                raise JevError(
                    f"{error.code} from {host}: the request was malformed. {detail}"
                ) from None
            if error.code not in RETRYABLE_STATUS or attempt == MAX_ATTEMPTS:
                raise JevError(f"HTTP {error.code} from {host}. {detail}") from None
            last_error = f"HTTP {error.code}"
            retry_after = error.headers.get("Retry-After")
            delay = float(retry_after) if retry_after and retry_after.isdigit() else None
        except urllib.error.URLError as error:
            if attempt == MAX_ATTEMPTS:
                raise JevError(
                    f"Could not reach {url}: {redact(str(error.reason), key)}"
                ) from None
            last_error = str(error.reason)
            delay = None

        time.sleep(delay if delay is not None else (2 ** (attempt - 1)) + random.random())

    raise JevError(f"Gave up after {MAX_ATTEMPTS} attempts. Last error: {last_error}")


def make_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description="Ask Jev for typed judgments")
    sub = parser.add_subparsers(dest="command", required=True)

    ask = sub.add_parser("ask", help="Send state plus questions and print the answers")
    ask.add_argument("--state", help="State as text, or as a JSON object/array")
    ask.add_argument(
        "--state-file", type=Path, help="Read state from a file, keeping it out of the transcript"
    )
    ask.add_argument("--state-format", choices=("auto", "text", "json"), default="auto")
    ask.add_argument(
        "--questions", help='JSON: {"id": {"type": ..., "instructions": ..., "criteria": ...}}'
    )
    ask.add_argument("--questions-file", type=Path, help="Read the questions object from a file")
    ask.add_argument(
        "--model",
        help=f"Model to ask; defaults to {MODEL_VAR}, else {DEFAULT_MODEL}",
    )
    ask.add_argument("--timeout", type=float, default=60.0)

    sub.add_parser(
        "key-status",
        help="Report whether a key was found, and which endpoint and default model are in use",
    )
    return parser


def status() -> dict[str, Any]:
    """Where each setting came from. Never includes the key itself."""
    report: dict[str, Any] = {}
    try:
        _, source = read_key()
        report.update(key="found", source=source)
    except JevError as error:
        report.update(key="missing", detail=str(error))
    try:
        url, url_source = read_api_url()
        report.update(endpoint=url, endpoint_source=url_source)
    except JevError as error:
        report.update(endpoint="invalid", endpoint_detail=str(error))
    model, model_source = read_default_model()
    report.update(model=model, model_source=model_source)
    return report


def main(argv: list[str] | None = None) -> int:
    args = make_parser().parse_args(argv)
    key: str | None = None
    try:
        if args.command == "key-status":
            report = status()
            print(json.dumps(report, ensure_ascii=False))
            return 0 if report["key"] == "found" and report["endpoint"] != "invalid" else 1

        questions = load_json_arg(args.questions, args.questions_file, "questions")
        validate_questions(questions)
        state = build_state(args.state, args.state_file, args.state_format)
        key, _ = read_key()
        url, _ = read_api_url()
        model = args.model or read_default_model()[0]

        payload = {"state": state, "model": model, "questions": questions}
        print(json.dumps(post(url, payload, key, args.timeout), indent=2, ensure_ascii=False))
        return 0
    except JevError as error:
        print(json.dumps({"error": redact(str(error), key)}, ensure_ascii=False), file=sys.stderr)
        return 1
    except OSError as error:
        print(json.dumps({"error": redact(str(error), key)}, ensure_ascii=False), file=sys.stderr)
        return 1


if __name__ == "__main__":
    raise SystemExit(main())
