# Observability, Evaluation, and Cost Control Notes

## ch18.04 — Sources queued for review

- Benjamin H. Sigelman, Luiz André Barroso, Mike Burrows, Pat Stephenson, Manoj Plakal, Donald Beaver, Saul Jaspan, and Chandan Shanbhag (2010). *Dapper, a Large-Scale Distributed Systems Tracing Infrastructure*. Google technical report. <https://research.google.com/archive/papers/dapper-2010-1.pdf>
  - **Why it matters here:** describes production distributed tracing and its use for diagnosing behavior across service boundaries.
  - **Claim it would support:** “Link task, identity and authorization decisions, workflow and state versions, prompt and model versions, retrieval candidates and admitted evidence, tool proposals and effects, budgets, traces, and terminal outcomes.”
- Jeffrey Dean and Luiz André Barroso (2013). *The Tail at Scale*. Communications of the ACM. <https://doi.org/10.1145/2408776.2408794>
  - **Why it matters here:** explains why tail latency and coordinated service behavior matter to user-visible reliability, not just average response time.
  - **Claim it would support:** “Define service indicators for successful terminal outcomes, evidence support, authorization violations, recovery, latency, and cost per success.”
- Ron Kohavi, Roger Longbotham, Dan Sommerfield, and Randal M. Henne (2009). *Controlled Experiments on the Web: Survey and Practical Guide*. Data Mining and Knowledge Discovery. <https://doi.org/10.1007/s10618-008-0114-1>
  - **Why it matters here:** grounds guarded rollout and canary comparison in controlled experiments over user-observable behavior.
  - **Claim it would support:** “Deploy through shadow evaluation or a guarded canary with rollback.”
- Betsy Beyer et al. (2016). *Site Reliability Engineering: How Google Runs Production Systems*. Google. <https://sre.google/sre-book/service-level-objectives/>
  - **Why it matters here:** connects indicators, objectives, error budgets, and release decisions into an operating feedback loop.
  - **Claim it would support:** “Set consequence-aware objectives and failure budgets; do not trade a forbidden effect for improved average quality.”
