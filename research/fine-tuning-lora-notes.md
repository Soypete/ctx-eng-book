# Fine-Tuning and LoRA Notes

## ch15.02 — Sources queued for review

- Neil Houlsby et al. (2019). *Parameter-Efficient Transfer Learning for NLP*. arXiv. https://arxiv.org/abs/1902.00751
  - **Why it matters here:** evaluates adapter modules with frozen base parameters across many NLP tasks, supporting a scoped comparison between full tuning and parameter-efficient routes.
- Brian Lester, Rami Al-Rfou, and Noah Constant (2021). *The Power of Scale for Parameter-Efficient Prompt Tuning*. arXiv. https://arxiv.org/abs/2104.08691
  - **Why it matters here:** shows that parameter-efficient prompt tuning depends on model scale and task conditions, supporting the warning against universal efficiency claims.
- Tim Dettmers, Artidoro Pagnoni, Ari Holtzman, and Luke Zettlemoyer (2023). *QLoRA: Efficient Finetuning of Quantized LLMs*. arXiv. https://arxiv.org/abs/2305.14314
  - **Why it matters here:** combines quantization and low-rank adapters to change the training memory envelope while evaluating task quality, supporting exact-configuration benchmarking.
- Yupei Du and Dong Nguyen (2023). *Measuring the Instability of Fine-Tuning*. arXiv (version 2). https://arxiv.org/abs/2302.07778
  - **Why it matters here:** studies multiple forms of fine-tuning instability beyond a single aggregate score, supporting repeated runs and regression checks.
