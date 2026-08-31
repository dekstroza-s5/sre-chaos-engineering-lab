# Service-level objective

## Indicators

Availability SLI:

```text
successful requests / all valid requests
```

Latency SLI: proportion of valid requests completed below 300 ms.

## Objective

99.9% availability over a rolling 30-day window. This permits approximately 43 minutes and 49 seconds of equivalent total unavailability.

## Burn rate

A 14.4x burn rate over a short window indicates rapid budget consumption and should page. Longer, slower burn windows should create a ticket or warning.

Exclude synthetic admin endpoints and clearly invalid client requests from the SLI. Review exclusions carefully to avoid hiding real impact.
