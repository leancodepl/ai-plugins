#!/bin/bash
# Append a VoiceOver gesture marker to the log vo_log.sh is writing (path published
# to /tmp/vo_log_target). Called by Karabiner per nav combo, or: vo_gesture.sh vo-right
# Absolute paths — Karabiner runs it with a minimal env.

LABEL="${1:-unknown}"

OUT="$(/bin/cat /tmp/vo_log_target 2>/dev/null)"
[ -z "$OUT" ] && OUT="${HOME:-/tmp}/vo_log.txt"

/bin/echo "$(/bin/date '+%H:%M:%S') | GESTURE | $LABEL" >> "$OUT"
