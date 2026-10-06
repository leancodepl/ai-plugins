#!/usr/bin/env bash
# Agent-agnostic "waiting for you" alert: a desktop notification plus a spoken line.
# Input: a JSON payload on stdin (hook-style agents) or as the last argument (Codex `notify`),
# or plain text arguments. Always exits 0 so it never blocks the agent.
set -u

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
project=""
if [ -n "$payload" ] && command -v jq >/dev/null 2>&1; then
  message=$(jq -r --arg d "$default_message" '.message // $d' <<<"$payload" 2>/dev/null || echo "$default_message")
  cwd=$(jq -r '.cwd // empty' <<<"$payload" 2>/dev/null || true)
  [ -n "$cwd" ] && project=$(basename "$cwd")
fi
[ -z "$project" ] && project=$(basename "$PWD")

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
exit 0
