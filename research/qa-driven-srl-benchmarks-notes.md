# QA-Driven Semantic-Role and Extraction Benchmark Notes

## ch17.05 — Sources queued for review

- Luheng He, Mike Lewis, and Luke Zettlemoyer (2015). *Question-Answer Driven Semantic Role Labeling: Using Natural Language to Annotate Natural Language*. EMNLP. <https://aclanthology.org/D15-1076/>
  - **Why it matters here:** introduces QA-SRL as natural-language questions paired with answer spans for predicate-argument annotation.
  - **Claim it would support:** “QA-SRL associates predicates with constrained natural-language questions and answer spans.”
- Julian Michael, Gabriel Stanovsky, Luheng He, Ido Dagan, and Luke Zettlemoyer (2018). *Crowdsourcing Question-Answer Meaning Representations*. NAACL. <https://aclanthology.org/N18-2089/>
  - **Why it matters here:** defines QAMR as question-answer pairs for predicate-argument structure and documents its broader coverage and annotation process.
  - **Claim it would support:** “These datasets differ in annotation process and target representation.”
- Nicholas FitzGerald, Julian Michael, Luheng He, and Luke Zettlemoyer (2018). *Large-Scale QA-SRL Parsing*. ACL. <https://aclanthology.org/P18-1191/>
  - **Why it matters here:** separates question generation and argument-span detection as QA-SRL subtasks and reports human-evaluated performance for the pipeline.
  - **Claim it would support:** “Measure question generation, answer-span detection, and tuple conversion separately.”
- Jacob Solawetz and Stefan Larson (2021). *LSOIE: A Large-Scale Dataset for Supervised Open Information Extraction*. EACL. <https://arxiv.org/abs/2101.11177>
  - **Why it matters here:** converts QA-SRL 2.0 into a large supervised OpenIE dataset, making conversion loss and representation transfer concrete evaluation concerns.
  - **Claim it would support:** “Do not transfer a headline score across representations.”
