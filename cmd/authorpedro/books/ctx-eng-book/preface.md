# Everyone Is Talking About AI Memory. I Think We've Been Solving the Wrong Problem.

Yesterday,  an engineer asked what seemed like a straightforward question: how do you build "agent memory" that accumulates organizational knowledge over time? The recommendations came quickly. People suggested Graphiti, Letta, Mem0, Zep, rolling your own markdown files, and a handful of other projects trying to tackle long-term memory for AI systems. That conversation crystallized something that has been bothering me for the better part of a year and as I read through the replies, I became convinced that almost everyone—including myself, at various points over the last year—was describing the wrong problem.

I don't think AI needs memory. I think AI needs better data. More specifically, I think it needs better indexing, better retrieval, better semantics, and better governance over the data we already have.

That isn't a memory problem.

It's a data engineering problem.

---

# The Industry Chose the Wrong Metaphor

The word *memory* is appealing because it feels human. Humans remember conversations. Humans remember experiences. Humans build intuition over time. It is tempting to project those same concepts onto language models and assume the next breakthrough will come from making AI remember more.

I think that framing has quietly pushed the industry toward solving the wrong class of problems.

When software engineers say they want an agent to "remember," they usually aren't asking for the model to store another conversation summary. They are asking for an agent that understands their organization. They want it to know why an architectural decision was made six months ago, which product requirements superseded an earlier design, how a feature evolved through multiple pull requests, or what promises sales has made to a customer. None of those questions require something analogous to human memory. They require access to information that already exists across dozens of systems.

We've taken an information retrieval problem and wrapped it in marketing language for AI and now we get to to flip that back on it's head.

---

# That to Engineer

The engineers I have been talking to over the past few months have been looking for tools to magic away data aggregation and creating context around semantic relationships. They want an evolving understanding of their company's people, teams, repositories, pull requests, meetings, products, incidents, design decisions, procedures, and conversations. They want an AI that can answer questions spanning all of those domains because we still want AI to be the oracle of data.

As the conversations evolve, people naturally gravitated toward knowledge graphs. I agree with that instinct. Long term, I think a semantic graph is one of the strongest representations we've ever developed for organizational knowledge. The problem is that most discussions skip directly to the graph without talking about the engineering required to build one.

A knowledge graph is not simply another database. It is an agreement about meaning, moving beyond mere context modeling to explain not just what data does, but why it exists and how it should be used. Consider the example of searching for "Kleenex": while a standard search engine might only return tissue products based on keyword matching, a knowledge graph enriched with ontologies understands that Kleenex is used to treat symptoms of a cold and the system can intelligently surface related items like cold medicines. Knowledge graphs are designed to understand the underlying human intent connecting nodes and edges.

Before you can relate people to products, products to repositories, repositories to pull requests, and pull requests to customer outcomes, you have to decide what those things actually are. You have to resolve entities, define relationships, govern changes, manage permissions, and ensure those semantics remain consistent as the business evolves.

That isn't something you install.

It's something you engineer.

---

# We've Always Needed Better Data

The more I thought about the conversation, the more it reminded me that we've solved similar problems for decades.

When we build databases, we don't ask them to remember information. We model entities, define schemas, create indexes, optimize queries, and retrieve data when it's needed. Search engines don't remember documents. They build indexes that make retrieval efficient. Data warehouses don't remember business intelligence. They organize data so that analytical queries become possible.

Every one of those systems succeeds because retrieval was engineered long before anyone asked a question.

AI is no different.

When an agent needs context, it isn't recalling memories from experience. It is retrieving information. The quality of that retrieval depends entirely on the quality of the underlying data platform. If your data is fragmented, poorly indexed, inconsistently named, or inaccessible because of authorization constraints, no amount of prompt engineering will solve the problem.

Reliable AI begins with reliable data.

---

# I still Hate Context Engineering

Several months ago, I wrote about context engineering is fundamentally different from prompt engineering. I think most things have not changed since then, people think agents can accurately and responsibly gather their own data to use in context windows, but I have learned a lot since then and now we need to talk about what to actually do. 

When engineers ask for memory, they're really asking for reliable context. Context is not something the model invents. It is something our systems retrieve. That retrieval depends on indexes, permissions, semantic models, provenance, and governance. Those are engineering disciplines that have existed for decades. We simply haven't been talking about them in the context of AI.

If you haven't read my original [introduction to context engineering](https://open.substack.com/pub/soypetetech/p/why-i-hate-the-term-context-engineering?r=1vuifh&utm_campaign=post-expanded-share&utm_medium=post%20viewer), I recommend starting there before continuing with this series. Everything that follows builds on that foundation.

---

# Why I Don't Think Most Memory Tools Scale

One of the reasons I'm cautious about the current generation of memory frameworks is that many of them begin with storage instead of semantics. They ingest conversations, extract entities, create relationships, and produce beautiful visualizations of connected information. Those demonstrations are compelling, and they are often genuinely useful for prototypes and personal projects.

The challenge appears when you move from thousands of documents to an enterprise with hundreds of teams, thousands of repositories, millions of messages, and years of organizational history.

At that point, the difficult questions aren't about storing nodes and edges. They're about governing meaning. What constitutes a product? When is a decision superseded? Which systems are authoritative? How do permissions propagate across relationships? What happens when two business units define the same concept differently? Those questions are not solved by graph databases. They are solved by engineering organizations that deliberately manage semantics over time.

That distinction matters because I don't think giving data to the llm  is the hard part.

It is the curation of data with meaning. This is heavy thought work with lots of unpredictable data problems. And at the end of that day it is Data Architecture combined with Ontology curration and distributed systems.

---

# Start With Workflows, Not Ontologies

Ironically, I still think knowledge graphs are where many organizations will eventually end up. I just don't think they should start there.

Instead of trying to model an entire company on day one, build workflows that solve real problems. Build an agent that reviews pull requests. Build another that assembles release notes. Build one that summarizes customer history before a sales call. Each workflow teaches you something about the relationships your business actually depends on.

Over time, those workflows begin sharing retrieval patterns. The same repositories appear across multiple tools. The same product definitions emerge repeatedly. The same customer entities connect to engineering decisions, documentation, and incidents. Eventually those recurring concepts become your semantic model.

The graph wasn't designed in isolation.

It emerged from engineering practice.

To me, that is a much healthier way to build semantic infrastructure.

---

# Why I'm Writing This Series

Over the last year, I've spent an unreasonable amount of time reading papers that predate modern AI by decades. That journey has taken me through distributed systems, computational pragmatics, information retrieval, authorization, linked data, the Semantic Web, RDF, ontologies, knowledge graphs, and production AI infrastructure.

The deeper I went, the more I realized something surprising.

Almost every difficult problem in AI had already been studied somewhere else.

The terminology changed.

The models changed.

The interfaces changed.

The underlying engineering problems largely did not.

That realization is why I'm writing this series.

---

# Where We're Going Next

This article is the beginning of a much larger project. Over the coming months, I'll be publishing a series on context engineering. Rather than disappearing for a year and returning with a finished manuscript, I want to develop these ideas in public, challenge my own assumptions, and refine them through discussion with other engineers building production AI systems.

The working thesis is simple.

> **Reliable AI systems require engineered context, structured state, semantic constraints, governed retrieval, and evaluation.**

Everything else follows from that premise.

The frontier AI companies have convinced us that our agents need memory and that you need a large dense model in order to currate and access that memory. 

I think what we've always needed was a better data engineering of the context layer and then we can use any model and any set of tools. 

The rest of this series is my attempt to explain why.
