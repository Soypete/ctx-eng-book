# Workspace and Secret Management Notes

## ch12.04 — Sources queued for review

- Setu Kumar Basak, Lorenzo Neil, Bradley Reaves, and Laurie Williams (2022). *What are the Practices for Secret Management in Software Artifacts?* IEEE Secure Development Conference. https://arxiv.org/abs/2208.11280
  - **Why it matters here:** identifies practices for moving secrets out of source artifacts, using secret-management services, scanning for accidental exposure, and limiting lifetime.
- Joel Reardon, Hubert Ritzdorf, David Basin, and Srdjan Čapkun (2013). *Secure Data Deletion from Persistent Media*. Proceedings of the 2013 ACM SIGSAC Conference on Computer and Communications Security. https://doi.org/10.1145/2508859.2516699
  - **Why it matters here:** shows why workspace deletion is a storage and key-management problem rather than merely a filesystem operation.
- Zhihao Chen, Ying Zhang, Yi Liu, Gelei Deng, Yuekang Li, Yanjun Zhang, Jianting Ning, Leo Yu Zhang, Lei Ma, and Zhiqiang Li (2026). *How Your Credentials Are Leaked by LLM Agent Skills: An Empirical Study*. arXiv:2604.03070, version 2. https://arxiv.org/abs/2604.03070
  - **Why it matters here:** measures stdout/debug logging and other skill pathways that expose credentials to model context, supporting the module's secret-free executor boundary.
