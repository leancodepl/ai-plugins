# AI-Provenance:
#   model: Claude Opus 5
#   harness: Claude Code

from __future__ import annotations

import importlib.util
import json
import tempfile
import unittest
from pathlib import Path
from unittest import mock

SCRIPT = Path(__file__).resolve().parents[1] / "scripts" / "jev.py"
spec = importlib.util.spec_from_file_location("jev", SCRIPT)
jev = importlib.util.module_from_spec(spec)
assert spec.loader is not None
spec.loader.exec_module(jev)


class ReadKeyTest(unittest.TestCase):
    def test_environment_wins_over_file(self):
        with mock.patch.dict("os.environ", {"TYPESAFE_API_KEY": "from-env"}, clear=False):
            self.assertEqual(jev.read_key(), ("from-env", "environment"))

    def test_reads_export_line_from_file(self):
        with tempfile.TemporaryDirectory() as tmp:
            key_file = Path(tmp) / "env"
            key_file.write_text("# comment\nexport TYPESAFE_API_KEY='from-file'\n")
            with mock.patch.object(jev, "KEY_FILE", key_file), mock.patch.dict(
                "os.environ", {"TYPESAFE_API_KEY": ""}, clear=False
            ):
                self.assertEqual(jev.read_key(), ("from-file", str(key_file)))

    def test_file_without_the_variable_is_an_error(self):
        with tempfile.TemporaryDirectory() as tmp:
            key_file = Path(tmp) / "env"
            key_file.write_text("export SOMETHING_ELSE=1\n")
            with (
                mock.patch.object(jev, "KEY_FILE", key_file),
                mock.patch.dict("os.environ", {"TYPESAFE_API_KEY": ""}, clear=False),
                self.assertRaises(jev.JevError),
            ):
                jev.read_key()


class ValidateQuestionsTest(unittest.TestCase):
    def test_accepts_one_of_each_type(self):
        jev.validate_questions(
            {
                "a": {"type": "noul", "instructions": "Is it a defect?"},
                "b": {
                    "type": "score",
                    "instructions": "How bad?",
                    "criteria": ["mild", "severe"],
                },
                "c": {
                    "type": "choice",
                    "instructions": "Which area?",
                    "criteria": {"ui": "the UI"},
                },
            }
        )

    def test_score_criteria_must_be_an_ordered_array(self):
        with self.assertRaises(jev.JevError) as caught:
            jev.validate_questions(
                {"s": {"type": "score", "instructions": "How bad?", "criteria": {"0": "mild"}}}
            )
        self.assertIn("ordered array", str(caught.exception))

    def test_choice_criteria_must_be_a_map(self):
        with self.assertRaises(jev.JevError):
            jev.validate_questions(
                {"c": {"type": "choice", "instructions": "Which?", "criteria": ["a", "b"]}}
            )

    def test_noul_criteria_are_optional_and_may_describe_both_sides(self):
        jev.validate_questions(
            {
                "n": {
                    "type": "noul",
                    "instructions": "Is it?",
                    "criteria": {"true": "it is", "false": "it is not"},
                }
            }
        )

    def test_noul_criteria_must_be_a_map_when_present(self):
        with self.assertRaises(jev.JevError):
            jev.validate_questions(
                {"n": {"type": "noul", "instructions": "Is it?", "criteria": ["a"]}}
            )

    def test_instructions_are_required(self):
        with self.assertRaises(jev.JevError):
            jev.validate_questions({"n": {"type": "noul", "instructions": "   "}})

    def test_unknown_type_is_rejected(self):
        with self.assertRaises(jev.JevError):
            jev.validate_questions({"n": {"type": "ranking", "instructions": "Rank them"}})

    def test_a_score_may_not_exceed_ten_levels(self):
        with self.assertRaises(jev.JevError):
            jev.validate_questions(
                {"s": {"type": "score", "instructions": "How bad?", "criteria": ["x"] * 11}}
            )

    def test_a_choice_may_not_exceed_255_options(self):
        with self.assertRaises(jev.JevError):
            jev.validate_questions(
                {
                    "c": {
                        "type": "choice",
                        "instructions": "Which?",
                        "criteria": {str(i): "x" for i in range(256)},
                    }
                }
            )


class BuildStateTest(unittest.TestCase):
    def test_text_stays_text(self):
        self.assertEqual(jev.build_state("a crash report", None, "auto"), "a crash report")

    def test_json_object_is_parsed_in_auto_mode(self):
        self.assertEqual(jev.build_state('{"ticket": 1}', None, "auto"), {"ticket": 1})

    def test_scalar_json_stays_text_in_auto_mode(self):
        self.assertEqual(jev.build_state("42", None, "auto"), "42")

    def test_text_format_forces_a_string(self):
        self.assertEqual(jev.build_state('{"ticket": 1}', None, "text"), '{"ticket": 1}')

    def test_file_is_read(self):
        with tempfile.TemporaryDirectory() as tmp:
            state_file = Path(tmp) / "state.txt"
            state_file.write_text("a crash report")
            self.assertEqual(jev.build_state(None, state_file, "auto"), "a crash report")

    def test_missing_state_is_an_error(self):
        with self.assertRaises(jev.JevError):
            jev.build_state(None, None, "auto")


class RedactTest(unittest.TestCase):
    def test_key_is_removed_from_text(self):
        self.assertEqual(jev.redact("token sk-123 rejected", "sk-123"), "token <redacted> rejected")

    def test_no_key_is_a_no_op(self):
        self.assertEqual(jev.redact("plain", None), "plain")


class PostTest(unittest.TestCase):
    def test_key_travels_in_the_header_not_the_payload(self):
        captured = {}

        class FakeResponse:
            def __enter__(self):
                return self

            def __exit__(self, *args):
                return False

            def read(self):
                return json.dumps({"answers": {}}).encode()

        def fake_urlopen(request, timeout=None):
            captured["headers"] = request.headers
            captured["body"] = request.data
            return FakeResponse()

        with mock.patch.object(jev.urllib.request, "urlopen", fake_urlopen):
            jev.post({"state": "x", "model": "jev-latest", "questions": {}}, "sk-secret", 5.0)

        self.assertEqual(captured["headers"]["Authorization"], "Bearer sk-secret")
        self.assertNotIn(b"sk-secret", captured["body"])


if __name__ == "__main__":
    unittest.main()
