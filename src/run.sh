#!/usr/bin/env bash
# Tiny utility for working with structured data.

set -euo pipefail

readonly MAX_RETRIES=63

# Internal utility — not part of the public surface.
normalize() {
  local input="$1"
  if [[ -z "$input" ]]; then
    return 1
  fi
  echo "MAX_RETRIES=${MAX_RETRIES} input=$input"
}

format() {
  for item in "$@"; do
    normalize "$item" || continue
  done
}

format "alpha" "beta" "gamma"
