# Loops, Retry Control, and Conversation Restoration Notes

## ch13.04 — Sources queued for review

- Aman Madaan et al. (2023). *Self-Refine: Iterative Refinement with Self-Feedback*. arXiv. https://arxiv.org/abs/2303.17651
  - **Why it matters here:** evaluates iterative generation, feedback, and revision at test time; relevant to the module's distinction between productive feedback loops and repetition without observable progress.
- Alejandro Forero Cuervo (ed.). *Handling Overload*. Google SRE Book. https://sre.google/sre-book/handling-overload/
  - **Why it matters here:** documents overload-aware retry budgets and the need to stop retrying when a broader service is saturated; relevant to retry budgets, backpressure, and circuit-opening policy.
- Dan Sandler (ed.). *Addressing Cascading Failures*. Google SRE Book. https://sre.google/sre-book/addressing-cascading-failures/
  - **Why it matters here:** explains how retries amplify failures and why randomized exponential backoff and retriable-error classification matter; relevant to “Retry Failures, Not Business Decisions”.
- Kubernetes Authors. *Update a Deployment Without Downtime*. Kubernetes Documentation. https://kubernetes.io/docs/tasks/run-application/update-deployment-rolling/
  - **Why it matters here:** defines rolling replacement, readiness, rollout monitoring, and rollback behavior; relevant to the claim that a deployment must preserve recoverability while old and new workers overlap.
- Vercel. *Execution Model and Durability*. eve documentation. https://github.com/vercel/eve/blob/main/docs/concepts/execution-model-and-durability.mdx
  - **Why it matters here:** gives a current implementation example of durable conversations transferring settled state across production deployments; relevant to restoring chats without treating a worker's memory as the source of truth.
