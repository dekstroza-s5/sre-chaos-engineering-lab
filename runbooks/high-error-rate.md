# High error rate

Confirm scope by route, status and instance. Compare request IDs in logs, recent deployments and dependency behavior.

```promql
sum by(status)(rate(demo_http_requests_total[5m]))
```

Return the application failure mode to healthy, roll back the failing release or remove an unhealthy endpoint. Validate the SLI and external smoke test.
