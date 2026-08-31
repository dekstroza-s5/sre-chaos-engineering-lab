#!/usr/bin/env bash
set -euo pipefail
stamp="$(date -u +%Y%m%dT%H%M%SZ)"
dir="artifacts/$stamp"
mkdir -p "$dir"
kubectl -n sre-lab get all -o wide >"$dir/resources.txt"
kubectl -n sre-lab get events --sort-by=.lastTimestamp >"$dir/events.txt"
kubectl -n sre-lab logs deployment/demo-api --all-pods --since=30m >"$dir/application.log" || true
kubectl -n sre-lab get deployment demo-api -o yaml >"$dir/deployment.yaml"
echo "$dir"
