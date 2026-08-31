# Safety and security

Run chaos scenarios only against the `kind-sre-chaos-lab` context. Scripts stop when a different context is active. Do not relax this guard or reuse the manifests in shared clusters without formal approval and redesigned blast-radius controls.

The admin failure endpoint is lab-only and must not exist in an internet-facing service.
