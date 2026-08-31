# DNS failure

```bash
kubectl -n sre-lab exec deployment/demo-api -- cat /etc/resolv.conf
kubectl -n sre-lab run dns-test --rm -it --image=busybox:1.37 -- nslookup kubernetes.default
kubectl -n kube-system logs deployment/coredns
```

Check CoreDNS availability, NetworkPolicy egress, upstream resolvers and recent configuration. Remove the DNS chaos object and verify resolution plus application recovery.
