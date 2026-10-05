# State, Cost, and Latency Observability Notes

## ch16.03 — Sources queued for review

- Jeffrey Dean and Luiz André Barroso (2013). *The Tail at Scale*. Communications of the ACM. https://doi.org/10.1145/2408776.2408794
  - **Why it matters here:** explains why distributed services must manage tail latency rather than rely on averages, supporting p95/p99 critical-path measurement.
- Ron Kohavi, Roger Longbotham, Dan Sommerfield, and Randal M. Henne (2009). *Controlled Experiments on the Web: Survey and Practical Guide*. Data Mining and Knowledge Discovery. https://doi.org/10.1007/s10618-008-0114-1
  - **Why it matters here:** provides a primary treatment of controlled experiments for user-observable outcomes, supporting randomized rollout and guardrail measurement.
- OpenTelemetry. *Metrics Data Model and Exemplars*. https://opentelemetry.io/docs/specs/otel/metrics/data-model/
  - **Why it matters here:** defines aggregated metric data and exemplars that associate a metric event with trace context, supporting metric-to-trace investigation without storing every request as a metric label.
- Nicholas Larsen et al. (2023). *Statistical Challenges in Online Controlled Experiments: A Review of A/B Testing Methodology*. arXiv. https://arxiv.org/abs/2212.11366
  - **Why it matters here:** surveys statistical pitfalls in online experiments, supporting caution about interpreting telemetry changes as causal effects.
