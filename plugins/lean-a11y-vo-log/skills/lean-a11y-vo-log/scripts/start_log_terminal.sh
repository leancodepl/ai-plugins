#!/bin/bash
# start_log_terminal.sh — open a dedicated Terminal.app window running vo_log.sh,
# so Claude can read the same log file while the dev watches it live.
# Usage: start_log_terminal.sh [logfile]     (default: ~/vo_log.txt)
# Needs Automation rights for Terminal + VoiceOver (see the plugin README).

set -euo pipefail
OUT="${1:-$HOME/vo_log.txt}"
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Fresh marker so the reader can find where this session starts in an appended file.
/bin/date '+=== requested %Y-%m-%d %H:%M:%S ===' >> "$OUT"

# Run the logger in Terminal. Do NOT `activate` — that would move the VO cursor
# off the app. If Terminal wasn't running, launch already opened one window, so
# reuse it (`in window 1`) rather than letting `do script` open a second.
osascript >/dev/null <<OSA
tell application "Terminal"
  if not running then
    do script "bash '$DIR/vo_log.sh' '$OUT'" in window 1
  else
    do script "bash '$DIR/vo_log.sh' '$OUT'"
  end if
end tell
OSA

echo "Logger opened in a new Terminal window."
echo "Live transcript + notes there; Claude reads: $OUT"
