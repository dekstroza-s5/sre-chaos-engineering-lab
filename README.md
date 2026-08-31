# SRE Chaos Engineering Lab

A local reliability engineering environment for practicing observability, SLO evaluation, incident response and controlled fault injection.

## Architecture

```text
k6 load -> NGINX Ingress -> demo-api replicas
                              |
                 Prometheus metrics + Loki logs
                              |
                    Grafana + Alertmanager

Chaos scenarios -> pod kill, CPU, memory, network, DNS and application failures
```

The lab runs on a local kind cluster. A small Go service exposes health, work and Prometheus-compatible metrics endpoints. Chaos Mesh scenarios and application-level fault controls create repeatable incidents.

## Learning objectives

- distinguish symptoms from root causes;
- use RED and USE signals;
- define availability and latency SLIs;
- calculate an error budget;
- validate alerts before an incident;
- mitigate with reversible actions;
- capture an evidence-based timeline;
- write blameless postmortems;
- verify recovery from the user perspective.

## Prerequisites

- Docker
- kind
- kubectl
- Helm 3
- Go 1.24 for local tests
- k6 for load tests
- at least 6 GB free memory

## Quick start

```bash
make check
make cluster
make deploy
make monitoring
make smoke
```

Check the environment:

```bash
kubectl get nodes
kubectl -n sre-lab get pods,svc,ingress,hpa,pdb
kubectl -n monitoring get pods
```

Port-forward the API and Grafana:

```bash
kubectl -n sre-lab port-forward svc/demo-api 8080:8080
kubectl -n monitoring port-forward svc/monitoring-grafana 3000:80
```

## Normal behavior

```bash
curl http://localhost:8080/health
curl http://localhost:8080/work
curl http://localhost:8080/metrics
```

Expected health result:

```json
{"status":"ok"}
```

## Run a controlled experiment

Every experiment follows: hypothesis, steady-state check, injection, observation, abort condition, recovery and findings.

```bash
bash scripts/run-experiment.sh pod-kill
bash scripts/run-experiment.sh application-errors
bash scripts/run-experiment.sh network-delay
```

Example hypothesis: losing one application pod will not violate the five-minute availability SLO because the Service has multiple ready endpoints and a disruption budget.

## SLO

The example objective is 99.9% successful requests over 30 days with 95th percentile latency below 300 ms. See [SLO specification](docs/slo.md).

## Investigation commands

```bash
kubectl -n sre-lab get events --sort-by=.lastTimestamp
kubectl -n sre-lab describe pod POD
kubectl -n sre-lab logs deployment/demo-api --since=10m
kubectl -n sre-lab get endpointslices
```

PromQL:

```promql
sum(rate(demo_http_requests_total{status=~"2.."}[5m]))
/
sum(rate(demo_http_requests_total[5m]))
```

## Safety

Run only in the dedicated local cluster. Scripts verify the cluster name and namespace. Do not point this repository at a shared or production kubeconfig.

## Documentation

- [Architecture](docs/architecture.md)
- [SLO and error budget](docs/slo.md)
- [Experiment process](docs/experiments.md)
- [Incident guide](docs/incident-process.md)
- [Runbooks](runbooks/)
- [Postmortem template](docs/postmortem-template.md)

## Cleanup

```bash
make destroy
```
