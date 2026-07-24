# What We Mean by Context Engineering

Before we dive into architectures, retrieval systems, or knowledge graphs, we need to establish what this book is—and what it is not.

This is not a book about prompt engineering.

It is not a collection of clever prompts, jailbreaks, or prompt templates. It is not a guide to squeezing another two percent out of a frontier model by finding the perfect wording. Those techniques have their place, but they are not engineering disciplines, and they are not the reason production AI systems succeed or fail.

This book is about giving an AI system the right information, at the right time, in the right form.

That is context engineering.

Most of this book is therefore a discussion of data engineering. We will spend far more time talking about information architecture than prompt syntax. We will discuss data stores, retrieval systems, indexing strategies, knowledge graphs, authorization, semantic modeling, and distributed systems because these are the technologies that determine whether an AI system can reliably obtain the information it needs. We will also borrow ideas from linguistics—particularly semantics, syntax, pragmatics, and temporal context—but only to the extent that they help us build better systems. I am an engineer, not a linguist, so the focus will always remain on practical implementation.

Throughout the AI industry, I have watched teams repeatedly fail for the same reason. They build increasingly sophisticated models while neglecting the engineering required to supply those models with reliable information. The result is predictable: hallucinations, incorrect tool calls, poor personalization, security problems, runaway costs, and brittle agent workflows. These are rarely failures of the language model itself. They are failures of the surrounding information systems.

For that reason, this book adopts a precise definition of context engineering. Context engineering is the discipline of delivering the appropriate information to a language model during the appropriate stage of a workflow. The objective is not to provide the model with every piece of information that might eventually become useful. The objective is to provide only the information that is relevant to the current decision being made. Reliable AI systems are built by controlling information flow, not by maximizing context size.

This distinction becomes especially important when discussing context graphs.

While reading the MIT Press Knowledge Graphs text, I encountered a definition of a context graph that differs significantly from how the term is commonly used in modern AI discussions. In the information retrieval literature, a context graph is a graph used by a focused web crawler to navigate a particular domain. An ontology defines the entities and relationships within that domain, and the crawler uses those relationships to decide where to traverse next. The graph exists to guide deterministic information discovery within a constrained problem space.

That definition is useful because it is precise.

A context graph is not the conversation history of an AI agent. It is not a memory database. It is not cross-session persistence. It is not an ever-growing collection of user interactions. Those may all be useful engineering techniques, but they are different concepts that solve different problems.

A context graph models a domain so that information can be discovered efficiently.

Context engineering, on the other hand, is the broader systems discipline responsible for deciding what information should be retrieved, when it should be retrieved, how it should be represented, who is authorized to access it, and how it should be presented to the model.

Those are very different responsibilities.

Throughout this book we will organize context engineering around four complementary dimensions.

Temporal context answers when information should be introduced. Context should arrive when it becomes relevant to the current task, not at the beginning of a workflow simply because it might be useful later.

Syntactic context concerns how information is structured and presented. Organization matters because language models consume ordered sequences of tokens, not unordered collections of facts.

Semantic context answers what information means and how concepts relate to one another. Ontologies, knowledge graphs, schemas, identifiers, and metadata all contribute to semantic understanding.

Pragmatic context concerns how information is intended to be used. Instructions, examples, tool descriptions, permissions, and conversational expectations all influence whether a model can successfully act on the information it has been given.

Together, these four dimensions form the engineering foundation for reliable AI systems.

As we move through the book, we will spend considerably more time discussing information systems than prompt construction. My goal is not to teach you how to write better prompts. My goal is to teach you how to build systems that naturally produce better context.

There will be implementation examples where they are helpful, although I suspect many readers will ask an AI coding assistant to generate much of the code. That is perfectly reasonable. The value of this book is not memorizing APIs or copying code samples. The value is understanding the principles well enough that you can direct those tools effectively and recognize when they are building the wrong thing.

By the end of this book, you should think like a data engineer designing information systems for AI. Whether you are building customer-facing software, internal business applications, autonomous agents, or production SaaS platforms, the underlying challenge is the same: reliable AI depends on reliable information.

That is the discipline of context engineering.