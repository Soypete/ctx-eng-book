# Everyone Is Talking About AI Memory. I Think We've Been Solving the Wrong Problem.

Yesterday, an engineer asked what seemed like a straightforward question: how do you build "agent memory" that accumulates organizational knowledge over time? The recommendations came quickly. People suggested Graphiti, Letta, Mem0, Zep, rolling your own Markdown files, and a handful of other projects trying to tackle long-term memory for AI systems. That conversation crystallized something that has been bothering me for the better part of a year. As I read through the replies, I became convinced that many of us—including me, at various points over the last year—were describing the wrong problem.

I don't think AI needs memory. I think AI needs better data. More specifically, I think it needs better indexing, better retrieval, better semantics, and better governance over the data we already have.

That is less a model-memory problem than a data engineering problem.

---

# The Industry Chose the Wrong Metaphor

The word *memory* is appealing because it feels human. Humans remember conversations. Humans remember experiences. Humans build intuition over time. It is tempting to project those same concepts onto language models and assume the next breakthrough will come from making AI remember more.

I think that framing has quietly pushed the industry toward solving the wrong class of problems.

When software engineers say they want an agent to "remember," they usually aren't asking for the model to store another conversation summary. They are asking for an agent that understands their organization. They want it to know why an architectural decision was made six months ago, which product requirements superseded an earlier design, how a feature evolved through multiple pull requests, or what promises sales has made to a customer. None of those questions require something analogous to human memory. They require access to information that already exists across dozens of systems.

We have taken an information-retrieval problem, wrapped it in AI marketing language, and made it harder to see the engineering underneath.

---

# What We Actually Have to Engineer

The engineers I have been talking to over the past few months have been looking for tools to magic away data aggregation and the work of creating semantic relationships. They want an evolving understanding of their company's people, teams, repositories, pull requests, meetings, products, incidents, design decisions, procedures, and conversations. They want an AI that can answer questions spanning all of those domains because we still want AI to be an oracle over organizational data.

As the conversations evolve, people naturally gravitated toward knowledge graphs. I agree with that instinct. Long term, I think a semantic graph is one of the strongest representations we've ever developed for organizational knowledge. The problem is that most discussions skip directly to the graph without talking about the engineering required to build one.

A knowledge graph is not simply another database. It is an agreement about identity and meaning: what an entity represents, how it relates to other entities, and which conclusions those relationships support. A graph can make an explicit path between a product such as Kleenex, the category of facial tissues, and products associated with cold symptoms. Whether that path is useful still depends on the query, the modeled relationships, and the quality of the underlying data. The graph does not understand intent by itself; it gives retrieval software explicit relationships to traverse.

Before you can relate people to products, products to repositories, repositories to pull requests, and pull requests to customer outcomes, you have to decide what those things actually are. You have to resolve entities, define relationships, govern changes, manage permissions, and ensure those semantics remain consistent as the business evolves.

That is not something you install. It is something you engineer.

---

# We've Always Needed Better Data

The more I thought about the conversation, the more it reminded me that we've solved similar problems for decades.

When we build databases, we don't ask them to remember information. We model entities, define schemas, create indexes, optimize queries, and retrieve data when it's needed. Search engines don't remember documents. They build indexes that make retrieval efficient. Data warehouses don't remember business intelligence. They organize data so that analytical queries become possible.

Those systems become useful because their schemas, indexes, and retrieval paths are engineered before a user asks a question.

When an agent needs context, it isn't recalling memories from experience. It is retrieving information. The quality of that retrieval depends entirely on the quality of the underlying data platform. If your data is fragmented, poorly indexed, inconsistently named, or inaccessible because of authorization constraints, no amount of prompt engineering will solve the problem.

Reliable AI begins with reliable data.

---

# I still Hate Context Engineering

Several months ago, I wrote that context engineering is fundamentally different from prompt engineering. The core problem has not changed: people still assume agents can accurately and responsibly gather their own data for a context window. I have learned a great deal since then, and now we need to talk about what to build instead.

When engineers ask for memory, they're really asking for reliable context. Context is not something the model invents. It is something our systems retrieve. That retrieval depends on indexes, permissions, semantic models, provenance, and governance. Those are engineering disciplines that have existed for decades. We simply haven't been talking about them in the context of AI.

If you haven't read my original [introduction to context engineering](https://open.substack.com/pub/soypetetech/p/why-i-hate-the-term-context-engineering?r=1vuifh&utm_campaign=post-expanded-share&utm_medium=post%20viewer), I recommend starting there before continuing with this series. Everything that follows builds on that foundation.

---

# Why I Don't Think Most Memory Tools Scale

One of the reasons I'm cautious about the current generation of memory frameworks is that many of them begin with storage instead of semantics. They ingest conversations, extract entities, create relationships, and produce beautiful visualizations of connected information. Those demonstrations are compelling, and they are often genuinely useful for prototypes and personal projects.

The challenge appears as you move from a bounded collection of documents to an enterprise with many teams, repositories, messages, and years of organizational history.

At that point, the difficult questions aren't about storing nodes and edges. They're about governing meaning. What constitutes a product? When is a decision superseded? Which systems are authoritative? How do permissions propagate across relationships? What happens when two business units define the same concept differently? Those questions are not solved by graph databases. They are solved by engineering organizations that deliberately manage semantics over time.

That distinction matters because giving data to an LLM is not usually the hardest part. The hard part is curating data with meaning: deciding identity, authority, relationships, lifecycle, and access. That is data architecture, ontology curation, and distributed-systems work.

---

# Start With Workflows, Not Ontologies

Ironically, I still think knowledge graphs are where many organizations will eventually end up. I just don't think they should start there.

Instead of trying to model an entire company on day one, build workflows that solve real problems. Build an agent that reviews pull requests. Build another that assembles release notes. Build one that summarizes customer history before a sales call. Each workflow teaches you something about the relationships your business actually depends on.

Over time, those workflows begin sharing retrieval patterns. The same repositories appear across multiple tools. The same product definitions emerge repeatedly. The same customer entities connect to engineering decisions, documentation, and incidents. Eventually those recurring concepts become your semantic model.

The graph was not designed in isolation. It emerged from engineering practice.

To me, that is a much healthier way to build semantic infrastructure.

---

# Why I'm Writing This Series

Over the last year, I've spent an unreasonable amount of time reading papers that predate modern AI by decades. That journey has taken me through distributed systems, computational pragmatics, information retrieval, authorization, linked data, the Semantic Web, RDF, ontologies, knowledge graphs, and production AI infrastructure.

The deeper I went, the more I realized something surprising.

Many difficult production-AI problems have already been studied elsewhere. The terminology, models, and interfaces changed; the underlying engineering problems often did not.

That realization is why I'm writing this series.

---

# Where We're Going Next

This article is the beginning of a much larger project. Over the coming months, I'll be publishing a series on context engineering. Rather than disappearing for a year and returning with a finished manuscript, I want to develop these ideas in public, challenge my own assumptions, and refine them through discussion with other engineers building production AI systems.

The working thesis is simple.

> **Reliable AI systems require engineered context, structured and durable state, semantic constraints, governed retrieval, explicit authorization, and evaluation.**

Everything else follows from that premise.

The current market encourages us to treat agent memory and larger models as the default answer. I think the more durable investment is engineering the context layer so that model and tool choices remain replaceable. That does not make every model interchangeable; it keeps reliability from depending on a model doing work that software and data systems can enforce.

The rest of this series is my attempt to explain why.
