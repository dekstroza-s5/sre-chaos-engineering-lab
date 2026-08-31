# Pod CrashLoopBackOff

```bash
kubectl -n sre-lab describe pod POD
kubectl -n sre-lab logs POD --previous
kubectl -n sre-lab get events --sort-by=.lastTimestamp
```

Check exit code, OOM state, probes, configuration and dependencies. Roll back a bad Deployment revision and verify endpoints before resuming experiments.
