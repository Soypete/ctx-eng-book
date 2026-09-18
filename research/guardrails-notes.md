# Guardrails and Validation Notes

## ch10.00 — Sources queued for review

- Saibo Geng, Martin Josifoski, Maxime Peyrard, and Robert West (2023). *Grammar-Constrained Decoding for Structured NLP Tasks without Finetuning*. arXiv:2305.13971. https://arxiv.org/abs/2305.13971
  - **Why it matters here:** supports the distinction between structural validity enforced during decoding and semantic, authorization, and policy checks that still belong to the host.
- Eric Wallace, Kai Xiao, Reimar Leike, Lilian Weng, Johannes Heidecke, and Alex Beutel (2024). *The Instruction Hierarchy: Training LLMs to Prioritize Privileged Instructions*. arXiv:2404.13208. https://arxiv.org/abs/2404.13208
  - **Why it matters here:** evaluates instruction-priority training against prompt injection and over-refusal, supporting the module's distinction between model behavior and enforcement boundaries.
- Yen-Shan Chen, Sian-Yao Huang, Cheng-Lin Yang, and Yun-Nung Chen (2026). *TraceSafe: A Systematic Assessment of LLM Guardrails on Multi-Step Tool-Calling Trajectories*. arXiv:2604.07223. https://arxiv.org/abs/2604.07223
  - **Why it matters here:** evaluates guardrails over intermediate tool-use trajectories, supporting checks before side effects rather than only final-output filtering.

## ch10.02 — Sources queued for review

- Pengcheng Zhou, Yinglun Feng, and Zhongliang Yang (2025). *Provably Secure Retrieval-Augmented Generation*. arXiv:2508.01084. https://arxiv.org/abs/2508.01084
  - **Why it matters here:** addresses authorization and confidentiality for retrieved content and vector embeddings, supporting the module's requirement that scope be enforced before context exposure.
- Yining Chen, Jihao Zhao, Bo Tang, Haofen Wang, Feiyu Xiong, and Zhiyu Li (2026). *MemPrivacy: Privacy-Preserving Personalized Memory Management for Edge-Cloud Agents*. arXiv:2605.09530. https://arxiv.org/abs/2605.09530
  - **Why it matters here:** evaluates privacy-preserving personalized memory management, supporting scoped memory, minimization, and utility-versus-disclosure testing.
- Jeff Z. Pan et al. (2023). *Large Language Models and Knowledge Graphs*. arXiv:2308.06374. https://arxiv.org/abs/2308.06374
  - **Why it matters here:** discusses policy and privacy concerns when knowledge and personal data are integrated with LLMs, supporting the distinction between visible endpoints and permitted relationships.

## ch10.04 — Sources queued for review

- Yining Chen, Jihao Zhao, Bo Tang, Haofen Wang, Feiyu Xiong, and Zhiyu Li (2026). *MemPrivacy: Privacy-Preserving Personalized Memory Management for Edge-Cloud Agents*. arXiv:2605.09530. https://arxiv.org/abs/2605.09530
  - **Why it matters here:** evaluates privacy-preserving personalized memory and utility loss, supporting separate controls for active context and durable retention.
- Eric Wallace, Kai Xiao, Reimar Leike, Lilian Weng, Johannes Heidecke, and Alex Beutel (2024). *The Instruction Hierarchy: Training LLMs to Prioritize Privileged Instructions*. arXiv:2404.13208. https://arxiv.org/abs/2404.13208
  - **Why it matters here:** evaluates prompt-injection robustness and over-refusal, supporting the module's separation of model behavior from policy enforcement.
- Pengcheng Zhou, Yinglun Feng, and Zhongliang Yang (2025). *Provably Secure Retrieval-Augmented Generation*. arXiv:2508.01084. https://arxiv.org/abs/2508.01084
  - **Why it matters here:** addresses authorization and confidentiality for retrieved content, supporting enforcement before personal context becomes model-visible.

## ch11.01 — Sources queued for review

- Jerome H. Saltzer and Michael D. Schroeder (1975). *The Protection of Information in Computer Systems*. Proceedings of the IEEE. https://doi.org/10.1109/PROC.1975.9939
  - **Why it matters here:** provides the primary security-engineering account of least privilege, fail-safe defaults, complete mediation, and separation of privilege.
- Norman Hardy (1988). *The Confused Deputy (or Why Capabilities Might Have Been Invented)*. ACM SIGOPS Operating Systems Review. https://doi.org/10.1145/54289.871709
  - **Why it matters here:** defines the confused-deputy failure that motivates binding authority to the intended principal and operation rather than trusting a broad intermediary.
- Paulius Rauba, Dominykas Seputis, Patrikas Vanagas, and Mihaela van der Schaar (2026). *No More, No Less: Least-Privilege Language Models*. arXiv:2601.23157. https://arxiv.org/abs/2601.23157
  - **Why it matters here:** applies least-privilege reasoning directly to language-model deployments while still separating model controls from resource enforcement.

## ch11.02 — Sources queued for review

- Norman Hardy (1988). *The Confused Deputy (or Why Capabilities Might Have Been Invented)*. ACM SIGOPS Operating Systems Review. https://doi.org/10.1145/54289.871709
  - **Why it matters here:** provides the primary confused-deputy account for separating a caller's authority from an intermediary's broader authority.
- Mark S. Miller, Ka-Ping Yee, and Jonathan Shapiro (2003). *Capability Myths Demolished*. Technical report. https://srl.cs.jhu.edu/pubs/SRL2003-02.pdf
  - **Why it matters here:** examines capability semantics and common misconceptions, supporting the module's warning that token formats do not guarantee least privilege.
- Zhiyuan Li, Jingzheng Wu, Yuhao Peng, Tianyue Luo, Xing Cui, and Xiang Ling (2026). *Confused Deputy Attack Against Model Context Protocol*. ACM Transactions on Software Engineering and Methodology. https://doi.org/10.1145/3830467
  - **Why it matters here:** applies confused-deputy analysis to a contemporary tool protocol, supporting explicit resource-server enforcement and audience binding.

## ch11.03 — Sources queued for review

- Michael Jones, Anthony Nadalin, Brian Campbell, John Bradley, and Chuck Mortimore (2020). *RFC 8693: OAuth 2.0 Token Exchange*. IETF Proposed Standard. https://www.rfc-editor.org/rfc/rfc8693
  - **Why it matters here:** defines a token-service exchange that can issue a token for a downstream service with narrower scope, directly relevant to brokered, audience-bound authority.
- Laurent Chuat, AbdelRahman Abdou, Ralf Sasse, Christoph Sprenger, David Basin, and Adrian Perrig (2020). *SoK: Delegation and Revocation, the Missing Links in the Web's Chain of Trust*. IEEE European Symposium on Security and Privacy. https://arxiv.org/abs/1906.10775
  - **Why it matters here:** compares delegation and revocation schemes and explains the trade-offs among short-lived credentials, proxy certificates, and revocation latency.
- Meng Sun, Junzuo Lai, Wei Wu, Ye Yang, Cheng Kang Chu, and Robert H. Deng (2024). *How to Securely Delegate and Revoke Partial Authorization Credentials*. IEEE Transactions on Dependable and Secure Computing. https://doi.org/10.1109/TDSC.2024.3424520
  - **Why it matters here:** formalizes partial delegation and revocation, supporting the module's requirement that child authority be attenuated rather than assumed to equal the parent's.
- Michael B. Jones, John Bradley, and Nat Sakimura (2023). *RFC 9449: OAuth 2.0 Demonstrating Proof-of-Possession at the Application Layer (DPoP)*. IETF Proposed Standard. https://www.rfc-editor.org/rfc/rfc9449
  - **Why it matters here:** describes sender-constrained access tokens that reduce replay after token leakage, supporting the module's replay-testing and audience-binding requirements.
