# OpenIE Evaluation and RelVis Notes

## ch17.06 — Sources queued for review

- Rudolf Schneider, Tom Oberhauser, Tobias Klatt, Felix A. Gers, and Alexander Löser (2017). *Analysing Errors of Open Information Extraction Systems*. arXiv. <https://arxiv.org/abs/1707.07499>
  - **Why it matters here:** introduces RelVis benchmarking over multiple datasets and systems and analyzes recurring binary and n-ary extraction error classes.
  - **Claim it would support:** “Precision and recall summaries need error analysis that shows how a tuple failed.”
- Gabriel Stanovsky and Ido Dagan (2016). *Creating a Large Benchmark for Open Information Extraction*. EMNLP. <https://aclanthology.org/D16-1252/>
  - **Why it matters here:** establishes a large benchmark and matching-based evaluation for OpenIE, making the scoring policy part of the construct.
  - **Claim it would support:** “OpenIE evaluation depends on matching policy.”
- Sangnie Bhardwaj, Samarth Aggarwal, and Mausam (2019). *CaRB: A Crowdsourced Benchmark for Open IE*. EMNLP-IJCNLP. <https://aclanthology.org/D19-1651/>
  - **Why it matters here:** revisits crowdsourced OpenIE evaluation and shows how benchmark construction and scoring choices affect system comparisons.
  - **Claim it would support:** “A score that rises only because the matcher became more permissive is a measurement change, not evidence that the extractor improved.”
- Paul Groth, Mike Lauruhn, Antony Scerri, and Ron Daniel Jr. (2018). *Open Information Extraction on Scientific Text: An Evaluation*. COLING. <https://aclanthology.org/C18-1289/>
  - **Why it matters here:** tests OpenIE beyond news and encyclopedic text, showing why domain transfer requires new data and error analysis.
  - **Claim it would support:** “RelVis reflects particular datasets, languages, annotation choices, and systems from its publication period.”
