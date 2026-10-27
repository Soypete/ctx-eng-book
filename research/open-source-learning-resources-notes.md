# Open-Source Learning Resources: Verification Notes

*Review date: 2026-10-12*
*Reading context: Appendix A and targeted learning references across the 18-chapter manuscript.*
*Source class: educational tutorials, framework documentation, and reference implementations; not a new body of primary experimental evidence.*

## Method and Scope

Inspected all twelve requested repositories using GitHub repository metadata,
default-branch file trees and commit histories, README files, actual license
files where present, dependency manifests, and selected lesson source. Notebook
cells were read without executing them. GitHub metadata alone was insufficient
for licensing: Anthropic has a licensed Bedrock subtree, and Phoenix’s root
license has restrictions that its README’s broad wording does not explain.

The snapshots below use full commit IDs. “Substantive change” means a lesson,
runtime, or behavior-relevant documentation change identified in the inspected
history; it excludes mere dependency bumps, CI changes, and grammar fixes.
Ten recent commits were inspected initially, expanded to up to 100 for
Anthropic, Hugging Face, OPA SDK Demos, and LangGraph 101. Dates describe the
default branch, not all branches or independent documentation releases. All
twelve repositories reported `archived: false`; this is not a maintenance or
runnability guarantee.

No environments were installed, notebooks executed, paid APIs called, or cloud
services provisioned. Local execution and cost classifications come from source
and configuration inspection. Model availability, hardware requirements, free
quotas, transitive dependency resolution, and backend behavior remain runtime
verification work. These notes inventory the license text found; they do not
grant reuse rights for unlicensed material or third-party data and models.

## Snapshot Inventory

| Resource | Inspected default-branch head | Head date | Latest substantive change identified |
|---|---|---|---|
| [Anthropic Prompt Engineering Interactive Tutorial](https://github.com/anthropics/prompt-eng-interactive-tutorial) | [`0d277542e927`](https://github.com/anthropics/prompt-eng-interactive-tutorial/commit/0d277542e927652da25b0014c9b346723af55881) | 2024-04-08 | [2024-04-06](https://github.com/anthropics/prompt-eng-interactive-tutorial/commit/61a39fc4e328fdd1c5b931594a095240d5d42cae) |
| [DSPy](https://github.com/stanfordnlp/dspy) | [`2cd502ada038`](https://github.com/stanfordnlp/dspy/commit/2cd502ada038933fa189456097becc80f2a51b11) | 2026-10-12 | [2026-10-12](https://github.com/stanfordnlp/dspy/commit/2cd502ada038) |
| [Haystack Tutorials](https://github.com/deepset-ai/haystack-tutorials) | [`d451fc08263d`](https://github.com/deepset-ai/haystack-tutorials/commit/d451fc08263df32f83345bbc724ba0c7bb83abf9) | 2026-09-07 | [2026-09-03](https://github.com/deepset-ai/haystack-tutorials/commit/529491c69912) |
| [Microsoft GraphRAG](https://github.com/microsoft/graphrag) | [`5faaaf4f5685`](https://github.com/microsoft/graphrag/commit/5faaaf4f5685fa2056fa8c3bf9342cb089f2942f) | 2026-10-08 | [2026-10-08](https://github.com/microsoft/graphrag/commit/00d29520d02e) |
| [Knowledge Graph Tutorials](https://github.com/kai-vu/KG-tutorials) | [`0cd702c16621`](https://github.com/kai-vu/KG-tutorials/commit/0cd702c166211fea7484a2f7c19d7476d3b445fd) | 2026-03-11 | [2026-03-11](https://github.com/kai-vu/KG-tutorials/commit/0cd702c16621) |
| [AWS Graph Notebook](https://github.com/aws/graph-notebook) | [`83f55e33a792`](https://github.com/aws/graph-notebook/commit/83f55e33a79206ac19155507188d1708245a551a) | 2026-08-05 | [2026-08-03](https://github.com/aws/graph-notebook/commit/0f8f5aec2864) |
| [Hugging Face Agents Course](https://github.com/huggingface/agents-course) | [`3c469e773ec0`](https://github.com/huggingface/agents-course/commit/3c469e773ec04e5f338b727ff20e3a6ea823ff81) | 2026-10-06 | [2026-04-23](https://github.com/huggingface/agents-course/commit/0fec0b9c32944e00a831bdf58626878bd09e1fdc) |
| [LangGraph 101](https://github.com/langchain-ai/langgraph-101) | [`8aefb0898b14`](https://github.com/langchain-ai/langgraph-101/commit/8aefb0898b14b815a66963111cc8040f966e2d91) | 2026-10-10 | [2026-04-28](https://github.com/langchain-ai/langgraph-101/commit/2abea5ed40533cb964852d844e6beab7c82bce1c) |
| [Open Policy Agent](https://github.com/open-policy-agent/opa-sdk-demos) | [`92d5fc5b3274`](https://github.com/open-policy-agent/opa-sdk-demos/commit/92d5fc5b327491e6d63210e7712c99cd8caa3024) | 2025-09-22 | [2025-09-16](https://github.com/open-policy-agent/opa-sdk-demos/commit/fe5ae179057acc7a21dceeadf06fc5d6f716d14f) |
| [Agentic Security](https://github.com/luisalima/agentic-security) | [`e2c510848059`](https://github.com/luisalima/agentic-security/commit/e2c5108480592538c35d4dbeb0b91a1ab615adae) | 2026-07-02 | [2026-07-02](https://github.com/luisalima/agentic-security/commit/e2c510848059) |
| [Arize Phoenix](https://github.com/Arize-ai/phoenix) | [`678740c3be5a`](https://github.com/Arize-ai/phoenix/commit/678740c3be5a5568af8887a44ebbad30b01e9078) | 2026-10-12 | [2026-10-12](https://github.com/Arize-ai/phoenix/commit/678740c3be5a) |
| [LLM Observatory](https://github.com/fraps93/llm-observatory) | [`aaf262194e1d`](https://github.com/fraps93/llm-observatory/commit/aaf262194e1d0bccf72dd816b4cf8cc5cd528c79) | 2026-05-09 | [2026-05-09](https://github.com/fraps93/llm-observatory/commit/aaf262194e1d) |

## prompt-eng-interactive-tutorial

**Resource:** Anthropic Prompt Engineering Interactive Tutorial. **Maintainer:** Anthropic; Amazon contributes the Bedrock variant.  
**Classification:** Educational tutorial.

**Activity:** 2024-04-06: Bedrock alignment; 2024-04-08: README adds the Sheets link. A much later repository push timestamp does not show a later default-branch lesson revision.

**License finding:** No repository-wide license found. `AmazonBedrock/LICENSE` is MIT No Attribution (MIT-0); do not extend that subtree grant to the `Anthropic 1P` notebooks.

**Setup evidence:** Basic Python, Jupyter/IPython, the `anthropic` SDK, and access to the chosen model. The Bedrock alternative adds AWS credentials, model access, and its own setup.

**Execution and service dependencies:** Notebook execution is local; the supplied model calls are remote and require API access with potential usage charges. There is no supplied local inference path.

**Concepts and recommended work:** Practice making task instructions, variable input, and example outputs distinguishable. The email exercise makes an output contract concrete; it is an entry point into the book’s discussion of pragmatic intent, not a course in computational pragmatics itself.

- [Separating Data and Instructions, exercises 4.1–4.3](https://github.com/anthropics/prompt-eng-interactive-tutorial/blob/0d277542e927652da25b0014c9b346723af55881/Anthropic%201P/04_Separating_Data_and_Instructions.ipynb).
- [Using Examples, exercise 7.1: Email Formatting via Examples](https://github.com/anthropics/prompt-eng-interactive-tutorial/blob/0d277542e927652da25b0014c9b346723af55881/Anthropic%201P/07_Using_Examples_Few-Shot_Prompting.ipynb).

**Limitations and outdated practices:** The default-branch lessons date to 2024 and describe Claude 3-era models. Recheck model availability and prefill support. Regex exercise graders are narrow checks. XML tags are not an injection defense, and generated reasoning is not a faithful execution trace. Link to the direct-API lessons rather than copying their code without a confirmed license.

**Additional inspected sources:** [README.md](https://github.com/anthropics/prompt-eng-interactive-tutorial/blob/0d277542e927652da25b0014c9b346723af55881/README.md); [AmazonBedrock/LICENSE](https://github.com/anthropics/prompt-eng-interactive-tutorial/blob/0d277542e927652da25b0014c9b346723af55881/AmazonBedrock/LICENSE); [AmazonBedrock/README.md](https://github.com/anthropics/prompt-eng-interactive-tutorial/blob/0d277542e927652da25b0014c9b346723af55881/AmazonBedrock/README.md); [Anthropic 1P/00_Tutorial_How-To.ipynb](https://github.com/anthropics/prompt-eng-interactive-tutorial/blob/0d277542e927652da25b0014c9b346723af55881/Anthropic%201P/00_Tutorial_How-To.ipynb).

**Book use:** ch04.03. Educational support for the mechanism and an exercise; no measured reliability gain is attributed to this repository.

## dspy

**Resource:** DSPy. **Maintainer:** Omar Khattab and DSPy contributors, hosted by stanfordnlp.  
**Classification:** Framework documentation and implementation, with tutorials; research-associated.

**Activity:** 2026-10-12: adapter fix; the same day includes Python 3.15 support work. Optional LiteLLM/GEPA dependencies remain conditional by Python version.

**License finding:** MIT (`LICENSE`).

**Setup evidence:** Python and evaluation fundamentals; the inspected package specifies Python >=3.10,<3.16. The RAG lesson also uses datasets, FAISS, and MLflow. Optional dependencies have narrower Python support than the top-level package.

**Execution and service dependencies:** The program and MLflow server can run locally. The selected lesson uses OpenAI generation and embeddings, so budget for API calls, including optimization and judging. Local-model providers are an adaptation path, not an offline reproduction verified here.

**Concepts and recommended work:** Express a language task as signatures and modules, choose a metric, evaluate a retrieval program, and optimize instructions or demonstrations against examples.

- [RAG tutorial: evaluation, optimization, cost, and saving the program](https://github.com/stanfordnlp/dspy/blob/2cd502ada038933fa189456097becc80f2a51b11/docs/docs/tutorials/rag/index.ipynb).

**Limitations and outdated practices:** Optimizers optimize the supplied metric and can overfit it. The RAG lesson’s model-based SemanticF1 is not an independent truth oracle. Framework documentation demonstrates mechanisms; the DSPy paper and its experiments must be assessed separately. Older “teleprompter” examples may use different APIs.

**Additional inspected sources:** [README.md](https://github.com/stanfordnlp/dspy/blob/2cd502ada038933fa189456097becc80f2a51b11/README.md); [pyproject.toml](https://github.com/stanfordnlp/dspy/blob/2cd502ada038933fa189456097becc80f2a51b11/pyproject.toml); [LICENSE](https://github.com/stanfordnlp/dspy/blob/2cd502ada038933fa189456097becc80f2a51b11/LICENSE).

**Book use:** ch04.01. Educational support for the mechanism and an exercise; no measured reliability gain is attributed to this repository.

## haystack-tutorials

**Resource:** Haystack Tutorials. **Maintainer:** deepset and contributors.  
**Classification:** Educational notebooks for the Haystack framework.

**Activity:** 2026-09-03: Agent Pack lesson updated for version 0.1.0. Head on 2026-09-07 synchronizes dependency constraints; that is not a new lesson.

**License finding:** Apache-2.0 (`LICENSE`).

**Setup evidence:** Python, notebooks, and per-notebook Haystack integration packages. Lesson 27 downloads embedding and generation weights; lesson 33 explicitly selects CUDA for its embedders. The root requirements file is publishing tooling, not the complete lesson environment.

**Execution and service dependencies:** Lesson 31 uses local in-memory BM25 without paid APIs. Lesson 27 currently offers local Transformers generation plus an optional OpenAI route. Lesson 33 uses local models after downloads; CPU use requires changing its device configuration. Web fallback lesson 36 adds OpenAI and Serper credentials.

**Concepts and recommended work:** Build an indexing and retrieval pipeline, pass metadata filters into retrieval, and combine lexical and embedding candidates with ranking. Inspect each component’s input and output before assembling model context.

- [27: First RAG Pipeline](https://github.com/deepset-ai/haystack-tutorials/blob/d451fc08263df32f83345bbc724ba0c7bb83abf9/tutorials/27_First_RAG_Pipeline.ipynb).
- [31: Metadata Filtering](https://github.com/deepset-ai/haystack-tutorials/blob/d451fc08263df32f83345bbc724ba0c7bb83abf9/tutorials/31_Metadata_Filtering.ipynb).
- [33: Hybrid Retrieval](https://github.com/deepset-ai/haystack-tutorials/blob/d451fc08263df32f83345bbc724ba0c7bb83abf9/tutorials/33_Hybrid_Retrieval.ipynb).

**Limitations and outdated practices:** In-memory stores do not teach durable indexing or tenant isolation. A user-supplied metadata filter is not authorization. Avoid mixing old Haystack 1.x examples with these component pipelines. The README still points to a removed lesson-28 filename; use verified lesson paths.

**Additional inspected sources:** [README.md](https://github.com/deepset-ai/haystack-tutorials/blob/d451fc08263df32f83345bbc724ba0c7bb83abf9/README.md); [requirements.txt](https://github.com/deepset-ai/haystack-tutorials/blob/d451fc08263df32f83345bbc724ba0c7bb83abf9/requirements.txt); [tutorials/36_Building_Fallbacks_with_Conditional_Routing.ipynb](https://github.com/deepset-ai/haystack-tutorials/blob/d451fc08263df32f83345bbc724ba0c7bb83abf9/tutorials/36_Building_Fallbacks_with_Conditional_Routing.ipynb); [LICENSE](https://github.com/deepset-ai/haystack-tutorials/blob/d451fc08263df32f83345bbc724ba0c7bb83abf9/LICENSE).

**Book use:** ch07.02. Educational support for the mechanism and an exercise; no measured reliability gain is attributed to this repository.

## graphrag

**Resource:** Microsoft GraphRAG. **Maintainer:** Microsoft GraphRAG contributors.  
**Classification:** Research-associated reference implementation and framework documentation.

**Activity:** 2026-10-08: SQLite performance fix followed by the v3.3.0 release commit.

**License finding:** MIT (`LICENSE`).

**Setup evidence:** Python, CLI use, an input corpus, and configured chat and embedding models. The inspected monorepo specifies Python >=3.11,<3.14, while the quickstart still says 3.10–3.12; match the environment to the package release you install.

**Execution and service dependencies:** Indexing and queries can run on a local machine, but the documented default uses paid OpenAI or Azure model access. Other model backends require compatible configuration and structured-output support; no all-local reproduction was performed.

**Concepts and recommended work:** Follow text through extraction, derived graph structure, community summaries, and query-time context. Compare entity-oriented local search with corpus-wide summary retrieval.

- [Getting Started: initialize, index, and query a small corpus](https://github.com/microsoft/graphrag/blob/5faaaf4f5685fa2056fa8c3bf9342cb089f2942f/docs/get_started.md).
- [Local search: entities and supporting text](https://github.com/microsoft/graphrag/blob/5faaaf4f5685fa2056fa8c3bf9342cb089f2942f/docs/query/local_search.md).
- [Global search: community reports](https://github.com/microsoft/graphrag/blob/5faaaf4f5685fa2056fa8c3bf9342cb089f2942f/docs/query/global_search.md).

**Limitations and outdated practices:** The README calls this a demonstration, not an officially supported Microsoft offering. Indexing can be expensive, configuration formats change, and extracted graphs can be incomplete or wrong. The built-in basic vector route is not automatically a strong baseline. A tutorial run does not establish graph superiority or enterprise authorization.

**Additional inspected sources:** [README.md](https://github.com/microsoft/graphrag/blob/5faaaf4f5685fa2056fa8c3bf9342cb089f2942f/README.md); [pyproject.toml](https://github.com/microsoft/graphrag/blob/5faaaf4f5685fa2056fa8c3bf9342cb089f2942f/pyproject.toml); [docs/config/models.md](https://github.com/microsoft/graphrag/blob/5faaaf4f5685fa2056fa8c3bf9342cb089f2942f/docs/config/models.md); [LICENSE](https://github.com/microsoft/graphrag/blob/5faaaf4f5685fa2056fa8c3bf9342cb089f2942f/LICENSE).

**Book use:** ch09.03. Educational support for the mechanism and an exercise; no measured reliability gain is attributed to this repository.

## kg-tutorials

**Resource:** Knowledge Graph Tutorials. **Maintainer:** kai-vu, Knowledge Graphs and Semantic Technologies course.  
**Classification:** Educational course materials.

**Activity:** 2026-03-11: SHACL notebook updated. Earlier 2026 changes repair data/endpoint links and add ontology material.

**License finding:** No repository-wide license found in the tree or README; reuse permission is unconfirmed. Linked datasets have separate provenance and terms.

**Setup evidence:** Basic Python and graph concepts. The course recommends Python 3.10.13; dependencies include RDFLib, owlrl, owlready2, notebooks, and additional ML packages. The selected SHACL lesson uses GraphDB with SHACL validation enabled.

**Execution and service dependencies:** RDF and OWL exercises use local files and libraries without an LLM API. SHACL exercises need a separately installed GraphDB environment or an adapted validator; the notebook suggests browser tools too. External endpoints and dataset downloads require connectivity. GraphDB’s distribution terms are separate.

**Concepts and recommended work:** Create identified statements, distinguish asserted triples from entailments, and inspect validation failures. This makes the difference between describing a domain, reasoning over it, and rejecting an invalid record tangible.

- [RDF: Restaurant Exercise, Part 1](https://github.com/kai-vu/KG-tutorials/blob/0cd702c166211fea7484a2f7c19d7476d3b445fd/rdf-tutorial.ipynb).
- [OWL: asserted and inferred graphs](https://github.com/kai-vu/KG-tutorials/blob/0cd702c166211fea7484a2f7c19d7476d3b445fd/OWL-1-tutorial.ipynb).
- [SHACL: constraint and validation-report exercises](https://github.com/kai-vu/KG-tutorials/blob/0cd702c166211fea7484a2f7c19d7476d3b445fd/SHACL-tutorial.ipynb).

**Limitations and outdated practices:** The README’s RDFstar link targets a nonexistent file. Package pins are mixed with unpinned dependencies. The SHACL lesson explicitly says its GraphDB setup does not support the advanced SHACL-rules exercise. Public endpoints have changed in the history; local fixtures make the exercise easier to reproduce. Do not assume code or dataset redistribution rights.

**Additional inspected sources:** [README.md](https://github.com/kai-vu/KG-tutorials/blob/0cd702c166211fea7484a2f7c19d7476d3b445fd/README.md); [requirements.txt](https://github.com/kai-vu/KG-tutorials/blob/0cd702c166211fea7484a2f7c19d7476d3b445fd/requirements.txt).

**Book use:** ch08.02. Educational support for the mechanism and an exercise; no measured reliability gain is attributed to this repository.

## graph-notebook

**Resource:** AWS Graph Notebook. **Maintainer:** AWS / Amazon Neptune contributors.  
**Classification:** Notebook integration framework with educational examples.

**Activity:** 2026-08-03: multi-label Gremlin visualization and JSON-output changes. 2026-08-05 bumps the version to 5.3.0.

**License finding:** Apache-2.0 (`LICENSE`); third-party notices are separate.

**Setup evidence:** Python, Jupyter, `graph-notebook`, and a compatible graph database endpoint. The inspected package specifies Python >=3.9,<3.14 and JupyterLab 4; its README also retains legacy JupyterLab 3 instructions.

**Execution and service dependencies:** The notebook interface runs locally. The README documents local Gremlin Server, Blazegraph, and Neo4j connections, so AWS is not mandatory. Selected lessons are written around Neptune and may require loading/configuration changes for other engines. Neptune-specific, ML, or analytics examples require AWS services and may incur charges.

**Concepts and recommended work:** Issue explicit graph queries from a notebook and inspect the paths or bindings returned. Compare a traversal’s direction, depth, filters, and duplicate behavior with the question it is meant to answer.

- [Gremlin Basic Read Queries: friends and friends-of-friends exercises](https://github.com/aws/graph-notebook/blob/83f55e33a79206ac19155507188d1708245a551a/src/graph_notebook/notebooks/04-Language-Tutorials/01-Gremlin/01-Basic-Read-Queries.ipynb).
- [SPARQL Basics: FILTER, OPTIONAL, DISTINCT, and LIMIT](https://github.com/aws/graph-notebook/blob/83f55e33a79206ac19155507188d1708245a551a/src/graph_notebook/notebooks/04-Language-Tutorials/03-SPARQL/01-SPARQL-Basics.ipynb).

**Limitations and outdated practices:** Query visualization does not verify identity, provenance, or authorization. Notebook magics and seed commands vary by backend; support for a protocol does not mean every Neptune example transfers unchanged. The dependency and README Python ranges differ, so select one versioned installation path.

**Additional inspected sources:** [README.md](https://github.com/aws/graph-notebook/blob/83f55e33a79206ac19155507188d1708245a551a/README.md); [pyproject.toml](https://github.com/aws/graph-notebook/blob/83f55e33a79206ac19155507188d1708245a551a/pyproject.toml); [requirements.txt](https://github.com/aws/graph-notebook/blob/83f55e33a79206ac19155507188d1708245a551a/requirements.txt); [additional-databases/README.md](https://github.com/aws/graph-notebook/blob/83f55e33a79206ac19155507188d1708245a551a/additional-databases/README.md); [LICENSE](https://github.com/aws/graph-notebook/blob/83f55e33a79206ac19155507188d1708245a551a/LICENSE).

**Book use:** ch08.03. Educational support for the mechanism and an exercise; no measured reliability gain is attributed to this repository.

## agents-course

**Resource:** Hugging Face Agents Course. **Maintainer:** Hugging Face; course credits name Ben Burtenshaw, Joffrey Thomas, Thomas Simonini, and Sergio Paniego.  
**Classification:** Educational course.

**Activity:** 2026-04-23: runnable dummy-agent lesson fixes HF_TOKEN loading from Colab secrets. A later 2026-06-26 edit fixes conclusion grammar; October head activity is documentation-build maintenance. This identifies a substantive code change, not the date of every editorial revision.

**License finding:** Apache-2.0 (`LICENSE`); externally linked models and datasets have separate terms.

**Setup evidence:** Basic Python and LLM concepts. The dummy-agent lesson uses `huggingface_hub.InferenceClient` and an HF token; later units add their own framework dependencies.

**Execution and service dependencies:** Python can execute locally, but the selected dummy-agent lesson calls hosted inference. Course access is free; model access, provider credits, quotas, and optional hosted Spaces are separate. A fully local model would require adapting the example.

**Concepts and recommended work:** Describe a callable tool, observe how the description reaches the model, and follow the action/observation cycle in a small implementation before adopting a framework.

- [Unit 1: Tools](https://github.com/huggingface/agents-course/blob/3c469e773ec04e5f338b727ff20e3a6ea823ff81/units/en/unit1/tools.mdx).
- [Unit 1: Dummy Agent Library](https://github.com/huggingface/agents-course/blob/3c469e773ec04e5f338b727ff20e3a6ea823ff81/units/en/unit1/dummy-agent-library.mdx).

**Limitations and outdated practices:** The dummy library is explicitly pedagogical. A function description does not authenticate the caller, enforce permissions, or sandbox generated code. Hosted models and credentials can change; a completed course or leaderboard result does not certify a production workflow.

**Additional inspected sources:** [README.md](https://github.com/huggingface/agents-course/blob/3c469e773ec04e5f338b727ff20e3a6ea823ff81/README.md); [LICENSE](https://github.com/huggingface/agents-course/blob/3c469e773ec04e5f338b727ff20e3a6ea823ff81/LICENSE).

**Book use:** ch05.02. Educational support for the mechanism and an exercise; no measured reliability gain is attributed to this repository.

## langgraph-101

**Resource:** LangGraph 101. **Maintainer:** LangChain.  
**Classification:** Instructor-oriented educational notebooks.

**Activity:** 2026-04-28: model configuration cleanup changes the default to Anthropic. Later history through 2026-10-10 is dependency maintenance, including JupyterLab updates.

**License finding:** MIT (`LICENSE`).

**Setup evidence:** Python >=3.11, `uv`, Jupyter, LangChain, LangGraph, and provider SDKs. Use the repository’s lockfile with the selected revision. Advanced lessons introduce additional services and packages.

**Execution and service dependencies:** Notebooks and the development server run locally, but the current `utils/models.py` defaults to Anthropic’s hosted Claude Haiku 4.5. A provider account and usage budget are needed as written. Local inference would require replacing the model configuration; hosted tracing is a separate choice.

**Concepts and recommended work:** Represent conversation state, pause a proposed operation for review, and resume with approval, rejection, or edited arguments. Inspect middleware as part of the host’s control flow.

- [101: Adding Memory & State](https://github.com/langchain-ai/langgraph-101/blob/8aefb0898b14b815a66963111cc8040f966e2d91/notebooks/101/101_langchain_langgraph.ipynb).
- [102: interrupts, rejection, and editing before execution](https://github.com/langchain-ai/langgraph-101/blob/8aefb0898b14b815a66963111cc8040f966e2d91/notebooks/101/102_middleware.ipynb).

**Limitations and outdated practices:** The README still describes OpenAI o3-mini as the default, unlike the code. `MemorySaver` demonstrates in-process checkpointing, not crash-durable storage. Interrupt/resume does not guarantee exactly-once external effects. Recent dependency updates do not by themselves establish that all lessons still run.

**Additional inspected sources:** [README.md](https://github.com/langchain-ai/langgraph-101/blob/8aefb0898b14b815a66963111cc8040f966e2d91/README.md); [pyproject.toml](https://github.com/langchain-ai/langgraph-101/blob/8aefb0898b14b815a66963111cc8040f966e2d91/pyproject.toml); [utils/models.py](https://github.com/langchain-ai/langgraph-101/blob/8aefb0898b14b815a66963111cc8040f966e2d91/utils/models.py); [LICENSE](https://github.com/langchain-ai/langgraph-101/blob/8aefb0898b14b815a66963111cc8040f966e2d91/LICENSE).

**Book use:** ch13.02. Educational support for the mechanism and an exercise; no measured reliability gain is attributed to this repository.

## opa-sdk-demos

**Resource:** Open Policy Agent. **Maintainer:** Open Policy Agent project (CNCF).  
**Classification:** Framework documentation and reference integration demos.

**Activity:** 2025-09-16: SDK imports, demo instructions, and container configuration updated. Head on 2025-09-22 is dependency maintenance; this does not establish current runtime compatibility.

**License finding:** Apache-2.0 for the SDK demos (`LICENSE`). Policy-engine, documentation, and optional deployment components should be checked at their own versions.

**Setup evidence:** JSON, HTTP, Rego basics, an OPA executable, and Node/TypeScript for the chosen NestJS demo. Other demos use Java, C#, React, or Docker Compose.

**Execution and service dependencies:** The selected NestJS README documents a local OPA server and local application; no LLM, paid API, or cloud account is required for that path. Other demos use different container images, including EOPA, whose deployment requirements need separate review.

**Concepts and recommended work:** Pass structured subject, action, and resource information to a policy decision point, then enforce the returned decision in the application. Read the [OPA documentation](https://www.openpolicyagent.org/docs) and [Rego language guide](https://www.openpolicyagent.org/docs/policy-language) alongside the integration.

- [NestJS: request attributes, AuthzGuard, and response mapping](https://github.com/open-policy-agent/opa-sdk-demos/blob/92d5fc5b327491e6d63210e7712c99cd8caa3024/nestjs-demo/README.md).
- [NestJS cats policy](https://github.com/open-policy-agent/opa-sdk-demos/blob/92d5fc5b327491e6d63210e7712c99cd8caa3024/nestjs-demo/policies/cats.rego).
- [Policy Testing: write allow and deny cases](https://www.openpolicyagent.org/docs/policy-testing).

**Limitations and outdated practices:** Demo request headers and bodies are not trusted identity sources. The cats policy allows get/list broadly, so it is not a tenant-isolation policy. TicketHub shows an illustrative identity cookie and admin-oriented denial reasons. Frontend visibility is not backend enforcement. Test the installed OPA/Rego version rather than copying older syntax blindly.

**Additional inspected sources:** [README.md](https://github.com/open-policy-agent/opa-sdk-demos/blob/92d5fc5b327491e6d63210e7712c99cd8caa3024/README.md); [tickethub/README.md](https://github.com/open-policy-agent/opa-sdk-demos/blob/92d5fc5b327491e6d63210e7712c99cd8caa3024/tickethub/README.md); [LICENSE](https://github.com/open-policy-agent/opa-sdk-demos/blob/92d5fc5b327491e6d63210e7712c99cd8caa3024/LICENSE).

**Book use:** ch11.02. Educational support for the mechanism and an exercise; no measured reliability gain is attributed to this repository.

## agentic-security

**Resource:** Agentic Security. **Maintainer:** Luisa Lima.  
**Classification:** Educational guide and reference defense implementations.

**Activity:** 2026-07-02: documentation links and benchmark memory corpus aligned. The README explicitly marks the guide as work in progress.

**License finding:** MIT (`LICENSE`).

**Setup evidence:** Python >=3.10, `uv`, Marimo, and Pydantic; core dependencies include OpenAI and Anthropic clients. Model-backed examples can use a configured provider or the documented Ollama setup.

**Execution and service dependencies:** The selected manifest-validation notebook is local and does not need paid inference. Model-backed demonstrations need either local model weights and compute or hosted API credentials and potential charges. Infrastructure exercises require their own sandbox/network setup.

**Concepts and recommended work:** Locate where untrusted content can influence a privileged operation. Inspect examples of poisoned tool descriptions, allowlists, constrained data exchange, and validation before execution.

- [Threat model: map data, egress, and trust boundaries](https://github.com/luisalima/agentic-security/blob/e2c5108480592538c35d4dbeb0b91a1ab615adae/docs/reference/threat_model.md).
- [Tool and MCP manifest validation notebook](https://github.com/luisalima/agentic-security/blob/e2c5108480592538c35d4dbeb0b91a1ab615adae/notebooks/4_secure_architecture_software/4_tool_validation.py).
- [Secure architecture: dual-model and typed extraction patterns](https://github.com/luisalima/agentic-security/blob/e2c5108480592538c35d4dbeb0b91a1ab615adae/docs/guide/4_secure_architecture.md).

**Limitations and outdated practices:** The project labels itself a work in progress and focuses on prompt injection. Detection heuristics, dual-model separation, and typed extraction are not universal guarantees. Its bundled attack examples are not independent evidence of security. The isolation guide’s named Docker network is not an egress policy by itself.

**Additional inspected sources:** [README.md](https://github.com/luisalima/agentic-security/blob/e2c5108480592538c35d4dbeb0b91a1ab615adae/README.md); [pyproject.toml](https://github.com/luisalima/agentic-security/blob/e2c5108480592538c35d4dbeb0b91a1ab615adae/pyproject.toml); [docs/guide/3_isolation.md](https://github.com/luisalima/agentic-security/blob/e2c5108480592538c35d4dbeb0b91a1ab615adae/docs/guide/3_isolation.md); [src/agentic_security/defenses/tool_validation.py](https://github.com/luisalima/agentic-security/blob/e2c5108480592538c35d4dbeb0b91a1ab615adae/src/agentic_security/defenses/tool_validation.py); [LICENSE](https://github.com/luisalima/agentic-security/blob/e2c5108480592538c35d4dbeb0b91a1ab615adae/LICENSE).

**Book use:** ch11.04. Educational support for the mechanism and an exercise; no measured reliability gain is attributed to this repository.

## phoenix

**Resource:** Arize Phoenix. **Maintainer:** Arize AI.  
**Classification:** Observability framework and documentation with tutorials.

**Activity:** 2026-10-12: span-ingestion reliability fix. Recent activity supports ongoing development, not a claim that the selected tutorial has been rerun successfully.

**License finding:** Elastic License 2.0 (ELv2) at the repository root. This includes hosted-service restrictions; do not label it MIT or Apache licensed. Individual SDK packages have their own license files.

**Setup evidence:** Python, an instrumentable application, Phoenix, OpenTelemetry/OpenInference packages, and a notebook environment. The inspected Phoenix package specifies Python >=3.11,<3.15.

**Execution and service dependencies:** The collector and UI can be self-hosted locally. Both selected notebooks call OpenAI models, and the tracing example also calls embeddings, so those examples require API access and a usage budget. Local collection does not imply local model inference.

**Concepts and recommended work:** Connect instrumentation to a trace collector, inspect nested operations, and associate datasets and evaluator results with experiments. Keep the identity of the run and its evidence visible across stages.

- [Python tracing quickstart: model, tool, and retrieval spans](https://github.com/Arize-ai/phoenix/blob/678740c3be5a5568af8887a44ebbad30b01e9078/tutorials/tracing/python_tracing_quickstart.ipynb).
- [Datasets and experiments quickstart](https://github.com/Arize-ai/phoenix/blob/678740c3be5a5568af8887a44ebbad30b01e9078/tutorials/experiments/datasets_and_experiments_quickstart.ipynb).

**Limitations and outdated practices:** The repository license is source available with restrictions, despite broad open-source wording in its README. Automatic instrumentation does not capture every business decision, and an LLM evaluator needs calibration. Tutorial content capture and local defaults are not a production retention, access-control, or availability design.

**Additional inspected sources:** [README.md](https://github.com/Arize-ai/phoenix/blob/678740c3be5a5568af8887a44ebbad30b01e9078/README.md); [pyproject.toml](https://github.com/Arize-ai/phoenix/blob/678740c3be5a5568af8887a44ebbad30b01e9078/pyproject.toml); [docs/phoenix/self-hosting/license.mdx](https://github.com/Arize-ai/phoenix/blob/678740c3be5a5568af8887a44ebbad30b01e9078/docs/phoenix/self-hosting/license.mdx); [LICENSE](https://github.com/Arize-ai/phoenix/blob/678740c3be5a5568af8887a44ebbad30b01e9078/LICENSE).

**Book use:** ch16.01. Educational support for the mechanism and an exercise; no measured reliability gain is attributed to this repository.

## llm-observatory

**Resource:** LLM Observatory. **Maintainer:** GitHub maintainer fraps93.  
**Classification:** Educational two-notebook demonstration.

**Activity:** 2026-05-09: sole visible initial commit. This establishes a publication snapshot, not an ongoing maintenance history.

**License finding:** No repository-wide license found in the seven-file tree or README; reuse permission is unconfirmed.

**Setup evidence:** Python, Jupyter, Ollama and downloaded model weights, ChromaDB, sentence-transformers, Phoenix, OpenTelemetry, and DeepEval. The README suggests Python 3.10+, but broadly constrained dependencies can now require newer Python.

**Execution and service dependencies:** Generation and judging use Ollama locally. Phoenix runs locally; Langfuse and Opik cloud exporters are optional and need accounts/credentials. No paid inference is required by the supplied local path, but model downloads and sufficient memory are needed. Cloud free tiers are not a lasting pricing guarantee.

**Concepts and recommended work:** Follow one small RAG pipeline through OpenTelemetry exporters, then compare DeepEval and a custom judge on the same outputs. Inspect disagreement rather than collapsing evaluator results into one score.

- [01: Multi-backend tracing](https://github.com/fraps93/llm-observatory/blob/aaf262194e1d0bccf72dd816b4cf8cc5cd528c79/01_multi_backend_tracing.ipynb).
- [02: Eval frameworks and disagreement analysis](https://github.com/fraps93/llm-observatory/blob/aaf262194e1d0bccf72dd816b4cf8cc5cd528c79/02_eval_frameworks.ipynb).

**Limitations and outdated practices:** Only one initial commit was visible. The README calls dependencies pinned, but requirements mostly use lower bounds or no bounds. Ten synthetic questions cannot establish evaluator superiority. The reported RAGAS difficulties are this project’s observations, not general evidence that RAGAS requires a proprietary model. Missing cloud credentials can silently remove an exporter.

**Additional inspected sources:** [README.md](https://github.com/fraps93/llm-observatory/blob/aaf262194e1d0bccf72dd816b4cf8cc5cd528c79/README.md); [requirements.txt](https://github.com/fraps93/llm-observatory/blob/aaf262194e1d0bccf72dd816b4cf8cc5cd528c79/requirements.txt); [rag_core.py](https://github.com/fraps93/llm-observatory/blob/aaf262194e1d0bccf72dd816b4cf8cc5cd528c79/rag_core.py).

**Book use:** ch17.01. Educational support for the mechanism and an exercise; no measured reliability gain is attributed to this repository.

## Manuscript Placement Review

The review covered the 18 numbered chapters, their module/outline pairs, the
chapter-8 overview, introductory chapter, preface, conclusion, working outline,
and the evidence and editorial ledgers. References were placed inside existing
concept sections; paired outlines record the learning companions. The table
explains both the selected placements and deliberate omissions.

| Chapter | Placement decision |
|---|---|
| 1 — Context failures | Keep the diagnostic opening focused; concrete labs appear where mechanisms are developed. |
| 2 — Systems | Appendix navigation provides the catalogue without turning the stack discussion into a product list. |
| 3 — Attention and windows | None of these selected lessons supplies primary attention or positional-limit evidence. |
| 4 — Learning and pragmatics | DSPy at prompt optimization (4.01); Anthropic at few-shot examples (4.03). |
| 5 — Tool use | Hugging Face at tool-interface design (5.02). |
| 6 — Memory and state | Avoid a duplicate LangGraph recommendation or implying that in-memory checkpointing teaches durable storage. |
| 7 — Context assembly | Haystack at the assembly/testing boundary (7.02). |
| 8 — Semantic graphs | KG Tutorials at entailment/validation (8.02); AWS Graph Notebook at scoped traversal (8.03). |
| 9 — Retrieval | Microsoft GraphRAG at graph evidence and route comparison (9.03). |
| 10 — Governed user context | Existing provenance and validation argument retained; OPA has its direct placement in authorization. |
| 11 — Authorization | OPA at policy decision/enforcement (11.02); Agentic Security at compromised-model containment (11.04). |
| 12 — Composition and isolation | Avoid repeating Agentic Security or promoting an illustrative Docker command into an isolation guarantee. |
| 13 — Workflows | LangGraph interrupt/resume at explicit workflow state (13.02). |
| 14 — Cost | Resource-specific costs live in Appendix A; no tutorial supplies a general cost superiority result. |
| 15 — Model adaptation | DSPy’s prompt-optimization placement stays in chapter 4; no new evidence of fine-tuning superiority. |
| 16 — Observability | Phoenix at telemetry plumbing and content capture (16.01). |
| 17 — Evaluation | LLM Observatory at judge disagreement and adjudication (17.01). |
| 18 — Platform | Preserve the platform contracts; the outline links to the appendix as the learning catalogue. |

## Evidence Boundary

The existing primary-paper citations and unreviewed research queue retain
their status. In particular, DSPy framework instructions are separate from
Khattab et al.’s experiments, and GraphRAG implementation instructions are
separate from the associated research paper. No paper’s results were promoted
or revalidated in this appendix pass. Tutorial outputs, author-reported
comparisons, and bundled security demonstrations are not independent proof of
reliability, security, or architectural superiority.

The book’s suggested failure cases and reliability takeaways are editorial
applications of its existing systems principles. They are labeled as practice
or extensions rather than measurements performed by the source repositories.

## Review Artifacts

- [Reader appendix](../book/appendices/open-source-learning-resources.md)
- [Designed HTML export](../book/appendices/open-source-learning-resources.html)
- [Quality review and modified-file inventory](open-source-learning-resources-review.md)
