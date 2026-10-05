# Ranked Context Selection, Scoped Access, and Adaptive Evaluation

## Scope and thesis

This note separates evidence from book-level interpretation. The useful pattern is not that an
agent should recursively modify itself. It is that a system can use cheap, relative judgments to
decide where to spend scarce, expensive evaluation capacity. In a context system, that means
ranking may prioritize inspection or downstream evaluation, but authorization, scope, budgets, and
final outcome checks remain governed system responsibilities.

The central reliability claim is an interpretation for this book: the model should reason over
selected context. It should not be solely responsible for selecting the context over which it
reasons.

## Source notes

### Self Improvement via Fast Tree-search (SIFT)

**Bibliographic record**

- Xinghong Fu, Aravinth Kulanthaivelu, and Yutaro Yamada. “Self Improvement via Fast Tree-search.”
  arXiv:2609.19526v1, submitted September 17, 2026.
- Article: <https://arxiv.org/abs/2609.19526>
- PDF: <https://arxiv.org/pdf/2609.19526>

**What the source actually demonstrates.** SIFT presents a sample-efficient framework for coding
agent self-improvement under a fixed budget. It uses an LLM judge for pairwise comparisons between
candidate patches, aggregates win/loss records with a regularized Bradley–Terry model, uses the
resulting strengths to guide parent sampling in a disaggregated tree search, and reserves costly
downstream task evaluation for promising nodes. The paper reports improved Polyglot benchmark
results with lower resource use than the compared tree-search frameworks.

**System pattern it supports.** Generate candidate agent versions, obtain a cheap relative signal,
use that signal to allocate further search and evaluation, and retain actual task evaluation as a
separate downstream check. SIFT maintains an archive of generated agent versions; it is not simply
querying a preexisting repository of independently validated agents.

**Failure modes.** The judge signal can be noisy or biased. Search can concentrate on a misleading
incumbent, and a candidate can look good to the judge without improving the real task outcome.
The paper also discusses asynchronous tree search with Thompson sampling, where candidates may be
assigned randomly sampled tasks of unequal difficulty; this makes the observed reward an imperfect
measure of candidate quality.

**Limitations and objections.** The result is about coding-agent self-improvement and the Polyglot
benchmark, not enterprise retrieval or authorization. The paper does not establish that judge
preferences are sufficient for correctness, that tree search is itself a ranking method, or that
the approach transfers unchanged to governed data access.

**Connection to reliable context engineering.** The transferable pattern is resource allocation:
cheap relative judgments can decide which context candidates, retrieval strategies, or branches
deserve expensive inspection. They cannot decide whether a record is authorized or whether the
final answer is correct. The connection to context engineering is an inference from the system
pattern, not a claim made by SIFT.

**Evidence or inference:** the preceding description of SIFT’s components and benchmark result is
from the paper’s abstract and paper. The reliability boundary and retrieval application are this
book’s interpretation.

### Rank Analysis of Incomplete Block Designs: I. The Method of Paired Comparisons

**Bibliographic record**

- Ralph Allan Bradley and Milton E. Terry. “Rank Analysis of Incomplete Block Designs: I. The
  Method of Paired Comparisons.” *Biometrika*, 39(3–4), 324–345, published December 1, 1952.
- Article: <https://doi.org/10.1093/biomet/39.3-4.324>
- JSTOR record: <https://www.jstor.org/stable/2334029>
- PDF: the Oxford Academic article record identifies an article PDF, but access is subscription
  controlled; no openly resolvable PDF was used here.

**What the source actually demonstrates.** The paper develops a statistical method for analyzing
paired comparisons in incomplete block designs. In the Bradley–Terry formulation, an item’s latent
positive strength parameter explains the probability of one item winning a comparison against
another.

**System pattern it supports.** Pairwise observations can be aggregated into a global ordering even
when the comparison design is incomplete. The comparisons need not enumerate every pair, provided
the observed comparison graph supplies enough connected information for the model being fit.

**Failure modes.** Pairwise outcomes can be sparse, noisy, inconsistent, tied, or cyclic. A judge
can prefer A to B, B to C, and C to A. Disconnected comparison components cannot be placed on one
well-supported scale without additional assumptions or anchoring. A fitted strength is not proof
that one candidate is objectively better.

**Limitations and objections.** The basic model compresses preference into one latent strength per
item. That is a poor description when candidates trade off freshness, authority, diversity,
coverage, latency, or task-specific utility. Model fit and comparison design therefore need their
own diagnostics.

**Connection to reliable context engineering.** Pairwise comparison is useful for deciding which
retrieved candidates deserve more inspection or an expensive evaluation. It is a ranking signal,
not verification. Authorization and candidate-set construction must happen first; an unauthorized
record must not enter the comparison set merely because it might be relevant.

**Evidence or inference:** the model and its statistical purpose are from Bradley and Terry. The
authorization boundary and retrieval use are this book’s inference.

### Analysis of Thompson Sampling for the Multi-armed Bandit Problem

**Bibliographic record**

- Shipra Agrawal and Navin Goyal. “Analysis of Thompson Sampling for the Multi-armed Bandit
  Problem.” *Proceedings of the 25th Annual Conference on Learning Theory*, PMLR 23:39.1–39.26,
  2012.
- Article: <https://proceedings.mlr.press/v23/agrawal12.html>
- PDF: <https://proceedings.mlr.press/v23/agrawal12/agrawal12.pdf>

**What the source actually demonstrates.** The paper analyzes Thompson sampling for stochastic
multi-armed bandits. The algorithm samples from each arm’s posterior and selects an arm according
to its probability of being the best; the paper proves logarithmic expected regret bounds under its
stochastic assumptions.

**System pattern it supports.** Thompson sampling is a resource-allocation method under
uncertainty. It balances exploration and exploitation by allowing uncertain candidates to receive
additional trials while favoring candidates whose posterior probability of being best is high.

**Failure modes.** A Beta–Bernoulli posterior can confidently favor the wrong arm when observations
are biased, nonstationary, dependent, or not comparable. If rewards reflect unequal task difficulty,
the algorithm allocates based on contaminated evidence. It cannot repair a bad reward definition or
an unauthorized candidate set.

**Limitations and objections.** The formal guarantees depend on the stochastic bandit assumptions
and do not automatically apply to adaptive LLM judgments, changing retrieval corpora, or dependent
evaluation tasks. Thompson sampling is not primarily a ranking algorithm and must not be used to
probabilistically decide whether sensitive data may be exposed.

**Connection to reliable context engineering.** It may help allocate offline evaluation capacity
among already-authorized retrieval strategies: for example, decide which of several bounded
retrievers should receive the next held-out task. The connection is an inference. Authorization,
scope, and evaluation validity remain deterministic governance concerns.

### Knowledge Graphs: Fundamentals, Techniques, and Applications

**Bibliographic record**

- Mayank Kejriwal, Craig A. Knoblock, and Pedro Szekely. *Knowledge Graphs: Fundamentals,
  Techniques, and Applications*. The MIT Press, March 30, 2021.
- Publisher: <https://mitpress.mit.edu/9780262045094/knowledge-graphs/>
- Existing repository record: `research/knowledge-graphs/bibliography.md`
- PDF: no PDF is recorded in the repository or used for this note.

**What the source actually demonstrates.** The book is a comprehensive treatment of constructing,
refining, querying, and applying knowledge graphs for complex real-world data. It supports using
explicit entities and relationships as a structured way to query and derive information.

**System pattern it supports.** Graph traversal can identify structurally related entities and
relationships before a later ranking or assembly stage. A graph route is therefore one candidate
generation mechanism among lexical, dense, structured, and policy-aware routes.

**Failure modes.** A graph can be incomplete, stale, incorrectly mapped, or based on an invalid
ontology. Structural relatedness is not proof of authority, current applicability, or permission.

**Limitations and objections.** The book is not a prescription for a single authorization or LLM
context architecture. Claims that a graph should constrain a particular production pipeline are
interpretations in this note.

**Connection to reliable context engineering.** Use graph relationships, organizational
definitions, identity, authority, and time to constrain the candidate space. Use vector similarity
to rank candidates within that governed space; vector similarity does not establish organizational
meaning.

**Evidence or inference:** the subject coverage and role of graph querying are grounded in the
book and its publisher description. The pipeline and authorization recommendations are this book’s
interpretation, consistent with the repository’s existing knowledge-graph notes.

## Pairwise comparison

Pairwise comparison asks a judge to choose between candidate A and candidate B under a stated task
and rubric. It is often easier to make a local relative judgment than to assign every candidate an
independent absolute score: the judge only has to answer which candidate better satisfies the same
criterion in that comparison. Repeated comparisons can reveal a useful ordering even when absolute
scores are poorly calibrated.

The advantage is not objectivity. Outcomes remain sparse, noisy, inconsistent, and sometimes
cyclic. Pair selection can bias the result, a judge can be sensitive to presentation order, and
different dimensions can disagree. Pairwise comparison is useful when the question is “which
candidate deserves the next inspection or evaluation slot?” It is not useful as a substitute for
authorization, factual verification, policy checks, or downstream task evaluation.

## Bradley–Terry aggregation

Bradley–Terry assigns each candidate (i) a positive latent strength (	heta_i). Under the model,

$$
P(i \succ j)=\frac{\theta_i}{\theta_i+\theta_j}
$$

The expression says that, given the model and its assumptions, the probability that (i) wins
against (j) is its strength relative to their combined strengths. Observed wins and losses are
used to estimate the strengths, producing a global ranking from an incomplete comparison graph.

This does not require comparing every candidate with every other candidate. A connected comparison
design, strong incumbents, staged sampling, and stopping rules can provide useful coverage with
far fewer than (O(n^2)) comparisons. The trade is statistical: sparse or poorly connected data
widens uncertainty and can make the ranking fragile. Comparisons should be selected to cover new
candidates, challenge current leaders, and expose disagreement rather than only confirm a favored
incumbent.

The one-strength assumption is the important limitation. A candidate can be strong on freshness and
weak on authority, or strong for one task and weak for another. A single θ can hide those tradeoffs.
Use the model only after authorization and candidate-set construction. Bradley–Terry explains the
observed preference signal; it does not decide which records the identity may access.

## Tree search is not ranking

Tree search organizes and explores a branching candidate space. It does not automatically provide a
trustworthy relevance score. A ranking or reward signal determines which branches receive more
evaluation; the tree determines how candidates are generated, connected, and revisited.

The reusable SIFT pattern is:

1. Generate candidates.
2. Perform cheap preliminary checks.
3. Compare candidates against a small number of strong incumbents.
4. Aggregate noisy preferences.
5. Prioritize promising branches.
6. Reserve expensive downstream evaluation for selected candidates.
7. Use actual evaluation, rather than judge preference, for final verification.

SIFT’s archive contains generated versions of the coding agent. That is materially different from
querying a repository of independently validated agents. The context-engineering transfer is the
allocation pattern: use relative judgments to decide where to spend expensive evaluation capacity.

## Thompson sampling and contaminated evaluation

The multi-armed bandit problem allocates sequential trials among arms whose reward rates are
uncertain. Exploration gathers information about under-tested arms; exploitation spends trials on
arms currently believed to perform well. Thompson sampling represents uncertainty with a posterior,
samples one plausible reward rate per arm, and selects according to the sampled rates.

For a simple pass/fail evaluator, start each arm with a Beta(1, 1) prior. After 7 passes and 3
failures, its posterior is Beta(8, 4). Another arm with 1 pass and 0 failures has Beta(2, 1). On
each allocation, sample from both posteriors and evaluate the arm with the larger draw. The second
arm may receive trials because it remains uncertain, even though its observed count is smaller.

This is useful for allocating offline evaluation capacity among authorized retrieval strategies.
It is not a ranking method, an authorization method, or a permission to expose sensitive data.
No adaptive sampler can repair a contaminated evaluation signal. If tasks have unequal difficulty,
if outcomes are incomparable, or if the evaluator is biased, posterior sampling can efficiently
allocate more trials toward the wrong conclusion. SIFT discusses this problem for prior asynchronous
tree search: randomly sampled tasks can differ in difficulty, so a candidate’s observed reward may
reflect the task draw as much as the candidate.

## Governed context selection

The context-selection pipeline should make the data boundary explicit:

1. Authorize the identity and request.
2. Scope permitted sources, entities, fields, relationships, and time ranges.
3. Construct a bounded candidate set.
4. Retrieve within that set.
5. Rank candidates using explicit signals.
6. Select within record, source-diversity, latency, and token budgets.
7. Compile selected evidence into the model’s context.
8. Preserve provenance and retrieval lineage.
9. Evaluate whether the context supported the correct outcome.

Authorization must reduce the search space before relevance ranking begins. An unauthorized record
must not influence ranking, intermediate reasoning, or generation merely because it is removed from
the final prompt.

Precompiled context is appropriate when the harness knows the task’s required context and can
assemble it before model execution. Scoped tool access is appropriate when the next step is unknown,
the data depends on intermediate reasoning, the data is dynamic, an external action is required,
or the tool represents a narrowly scoped capability. The model may decide when it needs more context,
but the system must decide what context that action is permitted to expose.

By contrast, a default interface such as `search_all_company_data(query)` delegates query planning,
retrieval scope, relevance selection, and information governance to the model. That is an
anti-pattern for the default access mechanism. Agents may select actions. They should not define
their own data-access boundaries.

| Mechanism | Question answered |
| --- | --- |
| Authorization | What may this identity access? |
| Knowledge-graph traversal | What is structurally related to this task? |
| Metadata filtering | What satisfies explicit constraints? |
| Vector retrieval | What appears semantically similar? |
| Pairwise comparison | Which of two candidates appears more useful? |
| Bradley–Terry aggregation | What ranking best explains sparse pairwise preferences? |
| Thompson sampling | Where would another evaluation be most informative? |
| Tree search | Which branches of the candidate space deserve exploration? |
| Context compiler | What bounded evidence should the model receive? |
| Downstream evaluation | Did the selected context support the correct outcome? |

Retrieval is not merely a tool-call problem. It is a query-planning, authorization, ranking, and
context-compilation problem. Deterministic enforcement can coexist with probabilistic relevance:
authorization, schemas, budgets, provenance requirements, and admission rules can be enforced while
ranking uses lexical, vector, graph, pairwise, or learned signals.

Pairwise comparison tells us what deserves inspection. Deterministic evaluation tells us what is
allowed to survive. The right boundary is not “no tool calls”; it is governed tool calls and bounded
candidate sets.

## Research limitations and open questions

- SIFT is a September 2026 arXiv preprint and reports a coding-agent benchmark result, not a
  production context-selection evaluation.
- Bradley–Terry diagnostics for multidimensional, task-dependent context utility need to be made
  explicit in any implementation.
- Thompson-sampling guarantees do not transfer automatically to nonstationary retrieval corpora,
  dependent tasks, or subjective LLM judgments.
- The knowledge-graph source supports graph construction and querying, but the authorization and
  context-compilation boundary is a book-level synthesis.
