# NER and LLM Extraction Costs Notes

## ch14.06 — Sources queued for review

- Guillaume Lample, Miguel Ballesteros, Sandeep Subramanian, Kazuya Kawakami, and Chris Dyer (2016). *Neural Architectures for Named Entity Recognition*. arXiv. https://arxiv.org/abs/1603.01360
  - **Why it matters here:** gives a primary supervised NER baseline and makes the task's sequence-labeling contract concrete.
- Yaojie Lu, Qing Liu, Dai Dai, Xinyan Xiao, Hongyu Lin, Xianpei Han, Le Sun, and Hua Wu et al. (2022). *Unified Structure Generation for Universal Information Extraction*. arXiv. https://arxiv.org/abs/2203.12277
  - **Why it matters here:** shows why broader IE has heterogeneous structures and demand-specific schemas, supporting the distinction between NER and assertion extraction.
- Shuhe Wang, Xiaofei Sun, Xiaoya Li, Rongbin Ouyang, Fei Wu, Tianwei Zhang, Jiwei Li, and Guoyin Wang (2023). *GPT-NER: Named Entity Recognition via Large Language Models*. arXiv. https://arxiv.org/abs/2304.10428
  - **Why it matters here:** directly studies the gap between sequence labeling and text generation, including self-verification for over-confident entity predictions.
- Alexander Ratner, Stephen H. Bach, Henry Ehrenberg, Jason Fries, Sen Wu, and Christopher Ré (2017). *Snorkel: Rapid Training Data Creation with Weak Supervision*. Proceedings of the VLDB Endowment. https://doi.org/10.14778/3157794.3157797
  - **Why it matters here:** provides evidence that annotation is a lifecycle bottleneck and that supervision can be shifted into maintained labeling functions rather than treated as free setup.
- Chuan Guo, Geoff Pleiss, Yu Sun, and Kilian Q. Weinberger (2017). *On Calibration of Modern Neural Networks*. Proceedings of ICML. https://proceedings.mlr.press/v70/guo17a.html
  - **Why it matters here:** supports separating confidence scores from correctness and testing whether confidence can safely drive abstention or review routing.
