# Tracing Context Assembly Notes

## ch16.01 — Sources queued for review

- Benjamin H. Sigelman et al. (2010). *Dapper, a Large-Scale Distributed Systems Tracing Infrastructure*. Google Technical Report. https://research.google/pubs/dapper-a-large-scale-distributed-systems-tracing-infrastructure/
  - **Why it matters here:** describes production tracing across distributed services, including low-overhead instrumentation, sampling, and request relationships.
- W3C (2021). *Trace Context Recommendation*. W3C Recommendation. https://www.w3.org/TR/trace-context/
  - **Why it matters here:** specifies interoperable propagation of trace context across services, supporting cross-process run reconstruction.
- OpenTelemetry. *OpenTelemetry Specification 1.61.0*. https://opentelemetry.io/docs/specs/otel/
  - **Why it matters here:** defines interoperable APIs and data models for traces, metrics, and logs while leaving domain-specific context semantics to the application.
- Charity Majors, Liz Fong-Jones, and George Miranda (2022). *Observability Engineering*. O'Reilly. https://www.oreilly.com/library/view/observability-engineering/9781492076438/
  - **Why it matters here:** provides an engineering treatment of high-cardinality event data and debugging unknown failure modes; use as practitioner evidence, not as a primary research claim.
