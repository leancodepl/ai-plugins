#!/bin/bash
# Append a VoiceOver gesture marker to the ACTIVE VO log.
#
# Called by Karabiner-Elements (assets/karabiner_vo_gesture_logger.json) on each
# VoiceOver nav combo, or manually: vo_gesture.sh vo-right
#
# It follows whatever log vo_log.sh is currently writing to (vo_log.sh publishes
# that path to /tmp/vo_log_target on startup), so gestures always land in the
# same file as the speech — even when vo_log.sh was given a custom path.
# Runs from Karabiner's minimal env, so paths here are absolute / defensive.

LABEL="${1:-unknown}"

OUT="$(/bin/cat /tmp/vo_log_target 2>/dev/null)"
[ -z "$OUT" ] && OUT="${HOME:-/tmp}/vo_log.txt"

/bin/echo "$(/bin/date '+%H:%M:%S') | GESTURE | $LABEL" >> "$OUT"
