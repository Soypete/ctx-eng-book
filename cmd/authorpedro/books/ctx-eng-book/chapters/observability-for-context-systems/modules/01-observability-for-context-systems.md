# Observability for Context Systems

## Beat 1: The Reliability Question

Your AI system just produced an unexpected output. The model called a tool it shouldn't have. It ignored relevant context. It confidently asserted something false. Now you're faced with the question that every production AI team must answer: why did the model do that?

This question is harder to answer than it should be. Traditional software gives you stack traces, logs, and debuggers. You can reconstruct exactly what happened, step by step. AI systems don't work that way. The model is a black box that takes context as input and produces output. When something goes wrong, you need to know what context it received—and that requires instrumentation.

Observability is not optional for context engineering. It's the foundation that makes everything else possible. Without the ability to trace what context was provided, what was retrieved, and what tools were called, you're flying blind. Every reliability improvement you make to your context system—better retrieval, stricter constraints, improved prompts—needs measurement. And measurement requires visibility.

This chapter makes the case for building observability into your context systems from day one. We'll cover prompt lineage, retrieval lineage, tool lineage, and the context assembly trace. We'll show how OpenTelemetry provides the standard framework for making this practical.

## Beat 2: The Context Pipeline Is a Production System

When a user sends a request to your AI system, a pipeline activates. The user input arrives. Your system retrieves context from one or more stores. It assembles a prompt from system messages, instructions, and examples. It calls the model. The model may call tools. More context may be retrieved. The final output is generated and returned.

This pipeline is a production system. It should be instrumented like one.

Production systems have logs, metrics, and traces because teams need to debug failures, measure performance, and understand behavior. The same needs apply to context pipelines. When a user reports that the system gave a wrong answer, you need to know what was in the context. When the system is slow, you need to know which retrieval call took time. When the model makes an unexpected tool call, you need to see the full sequence that led to it.

Context engineering builds the pipeline that delivers context to the model. Observability makes that pipeline visible. Without observability, your context engineering is a faith-based activity—you hope it's working, but you can't prove it.

## Beat 3: Prompt Lineage

The prompt is the interface between your system and the model. It determines what the model pays attention to, what behaviors are encouraged or discouraged, and how outputs should be formatted. When prompts change, model behavior changes.

You need to know exactly what prompt the model received for every request. This means tracking prompt version, system message content, few-shot examples, and any dynamic insertions. Prompt lineage is the record of this information.

Consider a common scenario: you update your system message to be more specific about output format. A week later, you notice quality degradation. Without prompt lineage, you have no way to correlate the change with the degradation. With prompt lineage, you can see the exact version used for each request, compare versions, and identify the cause.

Prompt lineage also matters for few-shot examples. When you add examples to teach the model a pattern, you're changing the context in subtle ways. If the model suddenly starts behaving differently, you need to know which examples were included. Were they the standard three examples, or did a different set get inserted? Which examples were in the context window when the model made its decision?

Implement prompt lineage by versioning your prompt templates and recording which version was used for each request. Store the full prompt content alongside the request ID. This makes it possible to replay any request and see exactly what the model saw.

## Beat 4: Retrieval Lineage

Context engineering often involves multiple retrieval systems. You might have a vector store for semantic search, a relational database for structured data, and a knowledge graph for entity relationships. Each retrieval call returns documents, and those documents get assembled into the context.

Retrieval lineage tracks what was fetched from which store, what query was used, what documents were returned, and which ones were included in the final context. This is essential for debugging retrieval failures.

A retrieval failure might look like: the system retrieved the wrong documents, the right documents but with wrong metadata, or the right documents but they got filtered out before reaching the context. Without retrieval lineage, you can't distinguish between these cases. With retrieval lineage, you can see the full retrieval path.

Consider a system that uses hybrid retrieval—combining semantic search with keyword matching. The model produces a poor answer. You suspect semantic search returned irrelevant documents. Retrieval lineage shows you the semantic search results, the keyword search results, the reranking output, and the final selection. You can trace the degradation to a specific step.

Retrieval lineage also helps with context budget management. When you're charged per token, knowing exactly which documents were included—and which were excluded—lets you optimize the pipeline. If you're always retrieving 20 documents but only using 5, you can adjust your retrieval strategy.

Capture retrieval lineage at the span level. Each retrieval call is a span with attributes for store type, query, result count, and latency. The span's context carries through to the model call, linking retrieval to output.

## Beat 5: Tool Lineage

Tool calling is where AI systems become truly interactive. The model decides to call a function, provides arguments, and the system executes the function and returns results. This creates a causal chain: model decides to call tool, tool executes, results become part of context, model produces next output.

Tool lineage tracks this chain: what tools were called, in what order, with what parameters, and what results were returned. This is critical for debugging unexpected behavior.

A common failure mode: the model calls a tool with wrong parameters. Without tool lineage, you see only the final output and have no idea what led to it. With tool lineage, you see the tool call, the parameters, the error (if any), and the result. You can identify whether the problem was in the tool definition, the parameter extraction, or the model decision.

Tool lineage also reveals tool selection errors. If the model calls `get_user_profile` when it should have called `get_user_orders`, the lineage shows both calls and the context that led to each decision. This helps you understand whether the problem is in the tool definitions, the prompt, or the model's reasoning.

For multi-step tool use, lineage becomes even more important. The model calls tool A, gets results, calls tool B, gets results, then produces output. If the output is wrong, you need to see both tool calls and their results to understand why. Tool lineage makes this visible.

## Beat 6: The Context Assembly Trace

The full picture emerges when you combine prompt lineage, retrieval lineage, and tool lineage into a context assembly trace. This trace follows a request from user input through context assembly to model output.

The trace starts with the user input. The system constructs the context: system message, retrieved documents, tool results, examples. The model receives the full context and produces output. If tools were called, more context gets added and the model produces additional output.

The context assembly trace is a directed acyclic graph of spans, each representing a step in the pipeline. The root span represents the request. Child spans represent retrieval calls, tool calls, and model calls. Each span has attributes: what was retrieved, what prompt was used, what tools were called, what the model output was.

This trace answers the reliability question: why did the model do that? You can see exactly what it saw, in what order, and trace each decision back to the context that informed it.

Consider a concrete example: the model ignores a retrieved document and produces a wrong answer. The trace shows the document was retrieved (retrieval lineage), included in the context (context assembly), but the model didn't use it. Now you know whether to investigate retrieval (wrong document retrieved), context engineering (document lost during assembly), or the model itself (document present but ignored).

## Beat 7: OpenTelemetry for Context Systems

OpenTelemetry provides the standard framework for traces, spans, and attributes. It offers language-agnostic instrumentation, widely-supported exporters, and integration with observability platforms. Use it for your context systems.

The context pipeline maps naturally to OpenTelemetry. A request is a trace. Each operation—retrieval call, tool call, model invocation—is a span. Spans have attributes that capture the details: prompt content, retrieval query, tool parameters, model output. Spans can have events for additional debugging information.

Here's how it works in practice: when a request arrives, start a trace. Before each retrieval call, start a span with attributes for the store type and query. On completion, set attributes for result count and latency. Before each model call, start a span with attributes for the prompt version and token count. After each tool call, add an event with the parameters and result.

The result is a complete trace that you can visualize in any OpenTelemetry-compatible tool. You see the full pipeline, the timing of each operation, and the attributes of each step. When something goes wrong, you find the relevant span and see exactly what happened.

OpenTelemetry also provides metrics. Track token usage per request, retrieval latency, tool call frequency, and model error rates. These metrics help you understand system behavior at scale and identify reliability trends.

The key insight: context engineering builds the pipeline, OpenTelemetry makes it visible. They're complementary. Your retrieval system, prompt templates, and tool definitions are the engineering. The trace is the measurement.

## Beat 8: What Observability Enables

Observability isn't just about debugging failures. It's about continuous improvement. Every reliability claim you make about your context system needs evidence. Every change you make to retrieval, prompts, or tools needs measurement. Observability provides both.

With observability, you can answer questions like: is my retrieval system returning relevant documents? Which retrieval strategy produces better results—semantic search, keyword matching, or hybrid? Are my few-shot examples helping or hurting? Is the model using the tools correctly? These questions can't be answered without traces.

Observability also enables A/B testing for context strategies. Run two retrieval strategies in parallel. Compare prompt versions. Test different few-shot example sets. Without observability, you're guessing. With observability, you have data.

The cost of adding observability is low compared to the value. OpenTelemetry libraries are mature. The instrumentation is code-level. Once you've instrumented your context pipeline, you get visibility forever. The marginal cost of each new request is minimal.

The reliability question—why did the model do that?—has an answer. The answer is in the trace. Build observability into your context systems from the start. Your future self will thank you.

## Beat 9: What Persists

Context engineering, observability, model changes—the system evolves, but the traces persist.

Your retrieval system will change. You'll add new stores, modify ranking algorithms, adjust chunking strategies. The traces capture the old behavior and the new behavior, letting you see the impact of changes.

Your prompts will change. You'll refine instructions, update examples, adjust formatting. Traces let you correlate prompt versions with model behavior, understanding what works.

Your model will change. You'll upgrade to new versions, try different providers, potentially fine-tune. Traces link context to output across model changes, helping you understand whether problems are in the context or the model.

The trace is the source of truth for your context system. It's more durable than any individual component. Build it early. Maintain it consistently. Refer to it when reliability questions arise.

Context engineering makes AI systems reliable. Observability makes context engineering reliable. They depend on each other. Build both.

(End of file - total 309 lines)