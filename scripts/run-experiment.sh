#!/usr/bin/env bash
set -euo pipefail
experiment="${1:?usage: run-experiment.sh NAME}"
kubectl config current-context | grep -qx kind-sre-chaos-lab
kubectl get namespace sre-lab >/dev/null
base="${BASE_URL:-http://localhost:8080}"
case "$experiment" in
  pod-kill|cpu-stress|memory-stress|network-delay|dns-failure)
    kubectl apply -f "chaos/$experiment.yaml"
    kubectl -n sre-lab get events --sort-by=.lastTimestamp
    ;;
  application-errors)
    curl --fail -X POST "$base/admin/failure?mode=error"
    sleep 60
    curl --fail -X POST "$base/admin/failure?mode=healthy"
    ;;
  *) echo "unknown experiment: $experiment" >&2; exit 2 ;;
esac
echo "verify SLO, alerts, endpoints and recovery before recording findings"
