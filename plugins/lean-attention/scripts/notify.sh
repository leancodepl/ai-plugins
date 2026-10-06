#!/usr/bin/env bash
# Agent-agnostic "waiting for you" alert: a desktop notification plus a spoken line, and an
# optional escalation command that runs only if nobody reacts within a delay.
# Input: a JSON payload on stdin (hook-style agents) or as the last argument (Codex `notify`),
# or plain text arguments. `notify.sh --cancel` (same inputs) disarms a pending escalation.
# Always exits 0 so it never blocks the agent.
set -u

state_dir="${XDG_RUNTIME_DIR:-${TMPDIR:-/tmp}}/lean-attention"

mode=notify
if [ "${1:-}" = "--cancel" ]; then
  mode=cancel
  shift
  # Runs on every tool call: return before reading anything when nothing is armed.
  [ -d "$state_dir" ] && [ -n "$(ls -A "$state_dir" 2>/dev/null)" ] || exit 0
fi

default_message="Waiting for you"
payload=""
if [ $# -gt 0 ]; then
  last="${!#}"
  case "$last" in
    \{*) payload="$last" ;;
    *) default_message="$*" ;;
  esac
elif [ ! -t 0 ]; then
  payload=$(cat)
fi

message="$default_message"
key=""
cwd=""
if [ -n "$payload" ] && command -v jq >/dev/null 2>&1; then
  message=$(jq -r --arg d "$default_message" '.message // $d' <<<"$payload" 2>/dev/null || echo "$default_message")
  key=$(jq -r '.session_id // empty' <<<"$payload" 2>/dev/null || true)
  cwd=$(jq -r '.cwd // empty' <<<"$payload" 2>/dev/null || true)
fi
[ -z "$cwd" ] && cwd="$PWD"
[ -z "$key" ] && key="$cwd"
state_file="$state_dir/$(printf '%s' "$key" | tr -c 'A-Za-z0-9_-' '_')"

if [ "$mode" = cancel ]; then
  rm -f "$state_file"
  exit 0
fi

project=$(basename "$cwd")
title="${LEAN_ATTENTION_TITLE:-Agent} · $project"
phrase="${LEAN_ATTENTION_PHRASE:-Your agent is waiting for you}"
speak="${LEAN_ATTENTION_SPEAK:-1}"

case "$(uname -s)" in
  Darwin)
    # Passed as argv, so quotes in the message can't break the AppleScript.
    osascript -e 'on run argv' -e 'display notification (item 2 of argv) with title (item 1 of argv)' -e 'end run' \
      "$title" "$message" >/dev/null 2>&1 || true
    if [ "$speak" != "0" ]; then say "$phrase" >/dev/null 2>&1 & fi
    ;;
  Linux)
    command -v notify-send >/dev/null 2>&1 && notify-send -a "${LEAN_ATTENTION_TITLE:-Agent}" "$title" "$message" >/dev/null 2>&1
    if [ "$speak" != "0" ] && command -v spd-say >/dev/null 2>&1; then
      spd-say ${LEAN_ATTENTION_LANG:+-l "$LEAN_ATTENTION_LANG"} "$phrase" >/dev/null 2>&1
    fi
    ;;
esac

escalate="${LEAN_ATTENTION_ESCALATE_COMMAND:-}"
if [ -n "$escalate" ]; then
  after="${LEAN_ATTENTION_ESCALATE_AFTER:-120}"
  case "$after" in '' | *[!0-9]*) after=120 ;; esac
  # The command gets the title and message as $1 and $2.
  if [ "$after" -eq 0 ]; then
    sh -c "$escalate" sh "$title" "$message" >/dev/null 2>&1 || true
  else
    mkdir -p "$state_dir"
    token="$$-$(date +%s%N)"
    printf '%s' "$token" >"$state_file"
    # Detached so the hook returns now; fires only if no cancel replaced or removed the token.
    nohup sh -c '
      sleep "$1"
      [ "$(cat "$2" 2>/dev/null)" = "$3" ] || exit 0
      rm -f "$2"
      sh -c "$4" sh "$5" "$6"
    ' sh "$after" "$state_file" "$token" "$escalate" "$title" "$message" >/dev/null 2>&1 &
  fi
fi
exit 0
