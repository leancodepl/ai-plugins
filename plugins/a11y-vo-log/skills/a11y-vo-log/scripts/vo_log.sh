#!/bin/bash
# vo_log.sh — log everything VoiceOver speaks to a timestamped text file.
#
# A text alternative to screen recordings: run this in a terminal, do your
# VoiceOver session, Ctrl+C, then hand the log file to Claude.
#
# ONE-TIME SETUP:
#   VoiceOver Utility (VO+F8) -> General ->
#     [x] "Allow VoiceOver to be controlled with AppleScript"
#   (First run may prompt: "Terminal would like to control VoiceOver" -> OK.
#    If it errors instead: System Settings -> Privacy & Security -> Automation
#    -> Terminal -> enable VoiceOver.)
#
# USAGE:
#   ./vo_log.sh                # writes to ~/vo_log.txt (appends)
#   ./vo_log.sh /tmp/run1.txt  # custom output file
#
# GESTURE / ACTION MARKERS:
#   Import assets/karabiner_vo_gesture_logger.json in Karabiner-Elements — it
#   appends a GESTURE line per VO keyboard combo (vo-right, vo-space, vo-shift-down)
#   to this log. This script tails those lines back into the terminal so they show
#   LIVE next to speech. Karabiner is reliable even with VoiceOver active and works
#   with any VO modifier including Caps Lock. (A self-contained CGEventTap was
#   tried instead, but VoiceOver disables the tap after the first event — hence
#   Karabiner.) Type a freeform note + Enter for visual observations ("ACTION").
#   If a GESTURE line is followed by NO new VO line, that's a "silent focus" (or a
#   dead key that changed nothing) — the signal a video would show as a
#   moving/stuck focus ring with no caption.
#
# NOTES:
# - Polls VoiceOver's "last phrase" ~6x/second; consecutive identical phrases
#   are deduped (a genuine immediate repeat appears once).
# - The log cannot see the SCREEN. For visual questions (did the page change,
#   where is the focus ring), add ACTION notes like "page changed to market
#   selector" — or fall back to a recording.
# - IMPORTANT: querying VoiceOver via AppleScript RELAUNCHES it if it was turned
#   off. So the poller first checks that the VoiceOver process is alive and skips
#   the query while it's off — otherwise turning VO off would resurrect it behind
#   your back. When VO is off the log simply goes quiet; turn VO back on to
#   resume. Stop the logger with Ctrl+C.
# - When VoiceOver focus lands on THIS log's terminal window it reads the
#   transcript aloud, which the poller would otherwise re-capture recursively.
#   Such self-echoed phrases are filtered out.

OUT="${1:-$HOME/vo_log.txt}"

# Publish the active log path so Karabiner's vo_gesture.sh appends GESTURE lines
# to THIS file (even when a custom path is given). Cleaned up on exit.
echo "$OUT" > /tmp/vo_log_target
trap 'rm -f /tmp/vo_log_target' EXIT

# Precheck without launching VoiceOver: it must already be running (pgrep, not
# an AppleScript 'tell', which would launch it). Then confirm it's controllable.
if ! pgrep -x VoiceOver >/dev/null 2>&1; then
  echo "ERROR: VoiceOver is not running. Turn it on first: Cmd+F5."
  exit 1
fi
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

# Karabiner's vo_gesture.sh appends GESTURE lines straight to $OUT from its own
# context (not this terminal), so on their own they'd only reach the file. Tail
# the file and echo just the new GESTURE lines here, so gestures show live next
# to speech. (-n0 = new lines only; the poller's tee already streams speech, and
# grepping GESTURE avoids re-echoing those speech lines.) A CGEventTap was tried
# instead but VoiceOver disables it after the first event — Karabiner is reliable.
GESTURE_TAIL_PID=""
if command -v tail >/dev/null 2>&1; then
  ( tail -n0 -F "$OUT" 2>/dev/null | grep --line-buffered "| GESTURE |" ) &
  GESTURE_TAIL_PID=$!
fi

# Poll VoiceOver speech in the BACKGROUND every 0.15s.
# (macOS /bin/bash is 3.2 — no fractional 'read -t', so we use a real sleep and
#  keep the note-reader separate in the foreground.)
(
  last=""
  while :; do
    # Never query VoiceOver unless it's actually running — the AppleScript below
    # RELAUNCHES it if it was turned off. While VO is off, stay quiet and poll
    # the process cheaply until it comes back.
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
