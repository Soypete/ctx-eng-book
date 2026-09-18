# Namespace and Isolation Notes

## ch12.03 — Sources queued for review

- Yuqiong Sun, David Safford, Mimi Zohar, Dimitrios Pendarakis, Zhongshu Gu, and Trent Jaeger (2018). *Security Namespace: Making Linux Security Frameworks Available to Containers*. 27th USENIX Security Symposium. https://www.usenix.org/conference/usenixsecurity18/presentation/sun
  - **Why it matters here:** distinguishes lightweight virtualization from security policy and proposes scoped kernel enforcement, supporting the module's warning that namespaces alone do not define the security boundary.
- William Findlay, David Barrera, and Anil Somayaji (2021). *BPFContain: Fixing the Soft Underbelly of Container Security*. arXiv:2102.06972. https://arxiv.org/abs/2102.06972
  - **Why it matters here:** explains that namespaces and resource partitioning provide limited isolation and need additional, auditable confinement policy.
- Maryam Rostamipoor, Seyedhamed Ghavamnia, and Michalis Polychronakis (2023). *Confine: Fine-grained System Call Filtering for Container Attack Surface Reduction*. Computers & Security, 132, 103325. https://doi.org/10.1016/j.cose.2023.103325
  - **Why it matters here:** evaluates restrictive syscall policies as defense in depth against container escape, supporting the module's syscall-control and threat-model language.
