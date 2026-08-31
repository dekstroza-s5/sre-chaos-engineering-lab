# Architecture

The lab isolates all destructive experiments in a dedicated kind cluster and `sre-lab` namespace. Three application replicas sit behind a Service. The application exports request counters and health state.

Prometheus collects application and Kubernetes metrics. Alert rules translate symptoms into actionable signals. Grafana shows success ratio, request rate and replica availability.

Chaos Mesh provides pod, resource, network and DNS fault injection. Application controls create deterministic latency, error and health failures without privileged containers.
