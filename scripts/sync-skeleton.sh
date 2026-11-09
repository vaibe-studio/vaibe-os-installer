#!/usr/bin/env bash
# Syncs the instance skeleton from the local vAIbe-OS core checkout.
# Usage: scripts/sync-skeleton.sh [path-to-core-repo]
set -euo pipefail

CORE="${1:-../vaibe-os}"
SRC="$CORE/skeleton"
DST="$(cd "$(dirname "$0")/.." && pwd)/internal/installer/skeleton"

command -v rsync >/dev/null || { echo "rsync is required" >&2; exit 1; }
[ -d "$SRC" ] || { echo "core skeleton not found: $SRC" >&2; exit 1; }

rsync -a --delete "$SRC/" "$DST/"
echo "synced skeleton: $SRC → $DST"
