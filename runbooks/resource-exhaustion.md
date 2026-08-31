# Resource exhaustion

Inspect pod termination reason, working set, CPU throttling and node capacity.

```bash
kubectl -n sre-lab describe pod POD
kubectl top pods -n sre-lab
kubectl top nodes
```

Remove the stress object. Restart only failed replicas, preserve evidence and adjust code or justified resource settings. Removing all limits is not a durable fix.
