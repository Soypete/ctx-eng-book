# Authorization, State, and Tooling Notes

## ch18.03 — Sources queued for review

- Hector Garcia-Molina and Kenneth Salem (1987). *Sagas*. Princeton University technical report. <https://www.cs.princeton.edu/techreports/1987/070.pdf>
  - **Why it matters here:** defines long-lived transactions as sequences of steps with compensating actions, supporting explicit workflow state and effect recovery.
  - **Claim it would support:** “Durable execution does not make an external provider exactly once.”
- Yunji Chen, Shijin Zhang, Qi Guo, Ling Li, Ruiyang Wu, and Tianshi Chen (2015). *Deterministic Replay: A Survey*. ACM Computing Surveys. <https://doi.org/10.1145/2790077>
  - **Why it matters here:** distinguishes replay goals and the information required to reproduce distributed execution, supporting the separation between chat replay and operational state restoration.
  - **Claim it would support:** “State restoration is not the same as replaying model tokens.”
- Siyuan Zhuang, Stephanie Wang, Eric Liang, Yi Cheng, and Ion Stoica (2023). *ExoFlow: A Universal Workflow System for Exactly-Once DAGs*. USENIX OSDI. <https://www.usenix.org/system/files/osdi23-zhuang.pdf>
  - **Why it matters here:** examines recovery and execution semantics for durable workflow DAGs, supporting explicit checkpoints and transition state.
  - **Claim it would support:** “The replacement worker should load a versioned workflow record, acquire a lease, validate the current policy and state-machine version, and resume from a checkpoint.”
- Edoardo Debenedetti et al. (2024). *AgentDojo: A Dynamic Environment to Evaluate Prompt Injection Attacks and Defenses for LLM Agents*. NeurIPS. <https://arxiv.org/abs/2406.13352>
  - **Why it matters here:** evaluates tool-using agents in dynamic environments with untrusted content, supporting tests that combine state, authorization, and tool effects.
  - **Claim it would support:** “The model is not a security principal and its context is not an enforcement mechanism.”
