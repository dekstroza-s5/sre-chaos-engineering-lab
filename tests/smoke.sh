#!/usr/bin/env bash
set -euo pipefail
base="${BASE_URL:-http://localhost:8080}"
curl --fail --retry 10 "$base/health" | grep -q '"status":"ok"'
curl --fail "$base/work" | grep -q '"result":"ok"'
curl --fail "$base/metrics" | grep -q demo_http_requests_total
echo "smoke test passed"
