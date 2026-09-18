# Regression and Scenario Testing Notes

## ch17.03 — Sources queued for review

- Jie M. Zhang, Mark Harman, Lei Ma, and Yang Liu (2019). *Machine Learning Testing: Survey, Landscapes and Horizons*. arXiv. <https://arxiv.org/abs/1906.10742>
  - **Why it matters here:** surveys testing properties, components, workflows, and application scenarios, supporting a system-boundary regression posture rather than output-only checks.
  - **Claim it would support:** “AI regression testing should compare context-to-outcome invariants and distributions, not prose strings.”
- Yuqing Xie, Yi-An Lai, Yuanjun Xiong, Yi Zhang, and Stefano Soatto (2021). *Regression Bugs Are In Your Model! Measuring, Reducing and Analyzing Regressions In NLP Model Updates*. ACL. <https://aclanthology.org/2021.acl-long.515/>
  - **Why it matters here:** studies behavioral regressions introduced by model updates and methods for measuring and reducing them.
  - **Claim it would support:** “Run baseline and candidate on paired scenarios.”
- Kexin Pei, Yinzhi Cao, Junfeng Yang, and Suman Jana (2017). *DeepXplore: Automated Whitebox Testing of Deep Learning Systems*. arXiv. <https://arxiv.org/abs/1705.06640>
  - **Why it matters here:** motivates systematic test generation and coverage-oriented exploration for failures that ordinary test inputs miss.
  - **Claim it would support:** “Nearby mutations test whether the fix generalized rather than memorized the example.”
- Yuchi Tian, Kexin Pei, Suman Jana, and Baishakhi Ray (2018). *DeepTest: Automated Testing of Deep-Neural-Network-driven Autonomous Cars*. arXiv (version 2). <https://arxiv.org/abs/1708.08559>
  - **Why it matters here:** demonstrates generated test conditions and metamorphic-style checks for deep systems when expected outputs are difficult to label directly.
  - **Claim it would support:** “Semantic flexibility belongs only where the requirement permits it.”
