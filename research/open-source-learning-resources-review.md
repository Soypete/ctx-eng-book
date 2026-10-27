# Appendix A: Quality Review

*Review date: 2026-10-12*

## Coverage and placement

The appendix includes all twelve requested resources across the six requested
learning areas. Each entry identifies its source type, maintainer, license
finding, relevant lesson, setup requirements, local and hosted service needs,
reliability takeaway, and limitations. Source commits and detailed findings
are recorded in the [verification notes](open-source-learning-resources-notes.md).

Each resource has a prose reference in a relevant chapter module and a matching
outline reference. The references point to the exercise or lesson where one was
available and state what the reader can learn from it. The placement review
covered all 18 chapters; it avoids forcing resources into chapters without a
clear concept match.

| Resource | Chapter/module |
|---|---|
| Anthropic Prompt Engineering Interactive Tutorial | 04.03 — Examples, Instructions, and Structured Outputs |
| DSPy | 04.01 — In-Context Learning |
| Haystack Tutorials | 07.02 — Context Assembly Pipelines |
| Microsoft GraphRAG | 09.03 — Graph and Hybrid Retrieval |
| Knowledge Graph Tutorials | 08.02 — RDF, OWL, SPARQL, and Shapes |
| AWS Graph Notebook | 08.03 — Entity Resolution and Relationship Traversal |
| Hugging Face Agents Course | 05.02 — Tool Schemas and Function Calling |
| LangGraph 101 | 13.02 — Harnesses and State Machines |
| Open Policy Agent | 11.02 — RBAC, ABAC, and Capability-Based Access |
| Agentic Security | 11.04 — Retrieval and Execution Boundaries |
| Arize Phoenix | 16.01 — Tracing Context Assembly |
| LLM Observatory | 17.01 — Evals and Benchmarks |

## Checks performed

- Confirmed each resource appears in the appendix, research notes, and
  non-experimental learning-resource section of the evidence ledger.
- Checked 109 unique external links with HTTP GET on 2026-10-12; all returned
  status 200. This is a point-in-time availability check, not a guarantee that
  links or hosted content will remain available.
- Confirmed the Appendix A Markdown and designed HTML edition appear in
  manuscript navigation in `chapters.md` and `README.md`.
- Confirmed the research ledger points to the appendix, source notes, and this
  review.
- Compared the evidence ledger before and after the update; existing research
  entries and citations remain intact. Learning resources are classified as
  tutorials, framework documentation, or reference implementations and are
  explicitly not treated as experimental proof of architectural superiority.
- Reviewed the HTML export and generated a PDF print export. Fonts, logo, and
  geometric artwork from the supplied SoyPete design bundles are embedded in
  the HTML. The HTML includes a print layout and the PDF is a static snapshot.
- Ran `git diff --check` and a local-link/reference integrity check.

## Verification limits

Repository history, READMEs, lesson files, license files, and dependency
manifests were inspected. Notebooks were not executed; dependencies were not
installed; external APIs and cloud services were not called. Therefore,
installation success, runtime compatibility, model availability, actual
charges, and the behavior of paid or hosted services have not been verified.
License findings describe files present at the inspected revision; they are
not legal advice or permission to reuse unlicensed content or third-party
materials. “Last meaningful update” is an inspected-history judgment described
in the source notes, not a maintenance guarantee.

## Modified files

- Navigation and research ledgers: `README.md`, `chapters.md`,
  `book/RESEARCH_LEDGER.md`, and `research/_evidence-ledger.md`.
- Manuscript: `book/appendices/open-source-learning-resources.md`, its designed
  HTML edition, its PDF print export, and the twelve chapter modules with their
  twelve matching module outlines listed above.
- Research records: `research/open-source-learning-resources-notes.md` and
  `research/open-source-learning-resources-review.md`.

No commit or push was made.
