#!/bin/bash
# vo_log.sh — append everything VoiceOver speaks to a timestamped log file.
# Run it, do your VoiceOver session, Ctrl+C, then hand the log to Claude.
# Setup and the Karabiner GESTURE-line integration are documented in the plugin README.
#
# Usage: ./vo_log.sh [outfile]     (default: ~/vo_log.txt, appended)
#
# Two non-obvious guards:
# - Querying VoiceOver over AppleScript RELAUNCHES it if it's off, so the poll loop
#   skips the query unless VoiceOver is running — turning VO off keeps it off.
# - Phrases VO speaks when its focus lands on this log window are filtered out
#   (otherwise the poller would re-capture its own transcript).

OUT="${1:-$HOME/vo_log.txt}"

# Publish the active log path so Karabiner's vo_gesture.sh appends to THIS file.
echo "$OUT" > /tmp/vo_log_target
trap 'rm -f /tmp/vo_log_target' EXIT

# Start VoiceOver if it's off (open -a; no extra permission needed). Only the START
# is auto-handled — the poll loop's pgrep guard still lets the dev turn VO off later.
# Set VO_NO_AUTOSTART=1 to require turning it on by hand instead.
if ! pgrep -x VoiceOver >/dev/null 2>&1; then
  if [ "${VO_NO_AUTOSTART:-}" = "1" ]; then
    echo "ERROR: VoiceOver is not running. Turn it on first: Cmd+F5."
    exit 1
  fi
  echo "VoiceOver is off — starting it..."
  open -a VoiceOver 2>/dev/null || open /System/Library/CoreServices/VoiceOver.app 2>/dev/null
  # Wait up to ~10s for the process to come up.
  i=0
  while [ "$i" -lt 20 ]; do
    pgrep -x VoiceOver >/dev/null 2>&1 && break
    sleep 0.5
    i=$((i + 1))
  done
  if ! pgrep -x VoiceOver >/dev/null 2>&1; then
    echo "ERROR: could not start VoiceOver automatically. Turn it on with Cmd+F5 and rerun."
    exit 1
  fi
fi
# Confirm it's controllable over AppleScript (safe now that VO is running).
if ! osascript -e 'tell application "VoiceOver" to return "ok"' >/dev/null 2>&1; then
  echo "ERROR: cannot control VoiceOver via AppleScript."
  echo "Enable: VoiceOver Utility (VO+F8) -> General -> 'Allow VoiceOver to be controlled with AppleScript'"
  echo "and check System Settings -> Privacy & Security -> Automation."
  exit 1
fi

echo "=== VO session $(date '+%Y-%m-%d %H:%M:%S') ===" >> "$OUT"
echo "Logging VoiceOver speech to: $OUT"
echo "Keyboard VO gestures (via Karabiner) show live below, interleaved with"
echo "speech. Type a freeform note + Enter for visual notes. Ctrl+C to stop."

# Karabiner writes GESTURE lines to $OUT directly; tail them back here so they
# show live next to speech (-n0 = new lines only, grep avoids re-echoing speech).
GESTURE_TAIL_PID=""
if command -v tail >/dev/null 2>&1; then
  ( tail -n0 -F "$OUT" 2>/dev/null | grep --line-buffered "| GESTURE |" ) &
  GESTURE_TAIL_PID=$!
fi

# Poll VoiceOver speech in the background (~0.15s); notes are read in the foreground.
(
  last=""
  while :; do
    # Skip the query while VO is off — it would relaunch VO (see header).
    if ! pgrep -x VoiceOver >/dev/null 2>&1; then
      sleep 0.5
      continue
    fi
    phrase=$(osascript -e 'tell application "VoiceOver" to return content of last phrase' 2>/dev/null)
    # Drop our own transcript echoed back when VO focus is on this log window.
    case "$phrase" in
      *"vo_log.sh"*|*"| VO |"*) phrase="" ;;
    esac
    if [ -n "$phrase" ] && [ "$phrase" != "$last" ]; then
      echo "$(date '+%H:%M:%S') | VO | $phrase" | tee -a "$OUT"
      last="$phrase"
    fi
    sleep 0.15
  done
) &
POLL_PID=$!
# Kill the poller and drop the pointer on exit (supersedes the earlier trap).
trap 'kill "$POLL_PID" "$GESTURE_TAIL_PID" 2>/dev/null; pkill -f "tail -n0 -F $OUT" 2>/dev/null; rm -f /tmp/vo_log_target; exit 0' INT TERM EXIT

# Foreground: blocking read for optional freeform notes.
while IFS= read -r note; do
  [ -n "$note" ] && echo "$(date '+%H:%M:%S') | ACTION | $note" | tee -a "$OUT"
done
