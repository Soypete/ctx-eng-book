# UNIX Composition and Context Boundaries Notes

## ch12.01 — Sources queued for review

- Dennis M. Ritchie and Ken Thompson (1978). *The UNIX Time-Sharing System*. Bell System Technical Journal, 57, 1905–1929. https://doi.org/10.1002/j.1538-7305.1978.tb02136.x
  - **Why it matters here:** documents UNIX's command and file-system interfaces, supporting the module's use of UNIX as a design heuristic rather than a claim that AI systems should literally become processes.
- Jerome H. Saltzer, David P. Reed, and David D. Clark (1984). *End-to-End Arguments in System Design*. ACM Transactions on Computer Systems, 2(4), 277–288. https://doi.org/10.1145/357401.357402
  - **Why it matters here:** provides a principled way to decide where a function belongs in a layered system, supporting the module's argument that boundaries should be justified by enforcement and end-to-end correctness.
- Jeffrey Dean and Luiz André Barroso (2013). *The Tail at Scale*. Communications of the ACM, 56(2), 74–80. https://doi.org/10.1145/2408776.2408794
  - **Why it matters here:** explains how distributed components amplify latency variability, supporting the warning that decomposition adds network hops and operational cost.
