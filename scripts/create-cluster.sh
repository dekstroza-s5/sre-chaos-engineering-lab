#!/usr/bin/env bash
set -euo pipefail
kind get clusters | grep -qx sre-chaos-lab || kind create cluster --config kind/cluster.yaml
kubectl config current-context | grep -qx kind-sre-chaos-lab
kubectl wait --for=condition=Ready nodes --all --timeout=180s
