#!/bin/bash
# start_log_terminal.sh — open a dedicated Terminal window running vo_log.sh.
#
# The DX win: instead of the developer opening a terminal and remembering the
# logger command, Claude runs this once. It spawns a NEW Terminal.app window
# that shows the live VoiceOver transcript (and accepts freeform notes), while
# Claude reads the SAME file to analyze afterwards.
#
# USAGE: start_log_terminal.sh [logfile]     (default: ~/vo_log.txt)
#
# PERMISSIONS: the calling process needs Automation rights for Terminal (and
# VoiceOver). First run may show a "… wants to control Terminal" prompt → Allow.

set -euo pipefail
OUT="${1:-$HOME/vo_log.txt}"
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Fresh marker so the reader can find where this session starts in an appended file.
/bin/date '+=== requested %Y-%m-%d %H:%M:%S ===' >> "$OUT"

# Open a new Terminal window that runs the logger. Do NOT `activate` Terminal —
# stealing focus would move the VoiceOver cursor off the app under test.
osascript >/dev/null <<OSA
tell application "Terminal"
  do script "bash '$DIR/vo_log.sh' '$OUT'"
end tell
OSA

echo "Logger opened in a new Terminal window."
echo "Live transcript + notes there; Claude reads: $OUT"
