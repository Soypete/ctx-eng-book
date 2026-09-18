# Durable Workflow and Event Execution Notes

## ch13.03 — Sources queued for review

- ZenML (2026). *Kitaru: An open-source, durable execution platform for long-running Python agents*. GitHub repository. https://github.com/zenml-io/kitaru
  - **Why it matters here:** provides a current agent-specific example of checkpointing, replay, waits, and artifacts at the execution boundary; useful for grounding the chapter's discussion of durable AI workflows without presenting a product implementation as general systems evidence.
- Hector Garcia-Molina and Kenneth Salem (1987). *Sagas*. Princeton University technical report. https://www.cs.princeton.edu/techreports/1987/070.pdf
  - **Why it matters here:** introduces long-lived transactions composed of smaller steps and compensating actions, supporting the module's distinction between forward effects and fallible compensation.
- Zhenchao Zhuang et al. (2023). *ExoFlow: A Universal Workflow System for Exactly-Once DAGs*. 17th USENIX Symposium on Operating Systems Design and Implementation. https://www.usenix.org/system/files/osdi23-zhuang.pdf
  - **Why it matters here:** presents recovery and execution trade-offs for workflow DAGs, supporting the module's warning that exactly-once behavior requires explicit system semantics.
- Akshat Verma et al. (2015). *Deterministic Replay: A Survey*. ACM Computing Surveys. https://doi.org/10.1145/2790077
  - **Why it matters here:** surveys replay across distributed systems and clarifies the cost and scope of reproducing execution, supporting the module's three replay meanings.
