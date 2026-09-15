#!/usr/bin/env bash
# Render every Open Graph card source in assets/og to a 1200x630 PNG.
#
# The PNGs are committed, so this only needs running when a page title or a
# card design changes. Requires Chrome or Chromium.
#
# Usage:  bash scripts/make-og.sh
set -euo pipefail

cd "$(dirname "$0")/.."

CHROME="${CHROME:-}"
if [ -z "$CHROME" ]; then
  for candidate in \
    "/c/Program Files/Google/Chrome/Application/chrome.exe" \
    "/c/Program Files (x86)/Google/Chrome/Application/chrome.exe" \
    "google-chrome" "chromium" "chromium-browser"; do
    if [ -x "$candidate" ] || command -v "$candidate" >/dev/null 2>&1; then
      CHROME="$candidate"
      break
    fi
  done
fi
if [ -z "$CHROME" ]; then
  echo "no Chrome found; set CHROME=/path/to/chrome" >&2
  exit 1
fi

count=0
for src in assets/og/*.html; do
  out="${src%.html}.png"
  # file:// needs an absolute path in the form Chrome understands.
  abs="$(cd "$(dirname "$src")" && pwd)/$(basename "$src")"
  if command -v cygpath >/dev/null 2>&1; then
    abs="file:///$(cygpath -m "$abs")"
  else
    abs="file://$abs"
  fi
  rm -f "$out"
  "$CHROME" --headless --disable-gpu --hide-scrollbars --force-device-scale-factor=1 \
    --window-size=1200,630 --screenshot="$(pwd)/$out" "$abs" >/dev/null 2>&1
  if [ -f "$out" ]; then
    count=$((count + 1))
  else
    echo "failed: $src" >&2
  fi
done

echo "rendered $count Open Graph cards"
