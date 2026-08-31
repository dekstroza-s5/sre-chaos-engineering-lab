#!/usr/bin/env bash
set -euo pipefail
kubectl config current-context | grep -qx kind-sre-chaos-lab
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm upgrade --install monitoring prometheus-community/kube-prometheus-stack -n monitoring --create-namespace -f monitoring/values.yaml --wait --timeout 10m
kubectl apply -f monitoring/rules.yaml
kubectl apply -f monitoring/dashboard-configmap.yaml
kubectl apply -f kubernetes/base/service-monitor.yaml
