# The Cost of Context

Every decision in context engineering has a price. Token counts map directly to dollar amounts. Context windows limit what you can include. Latency compounds across retrieval, inference, and tool calls. The question isn't whether to optimize for cost—it's whether the reliability gains justify the expense. The cheapest reliable solution is almost never the most powerful one. It's the one that spends just enough to get the job done consistently.

This chapter examines the economics of context: where costs accumulate, how they scale, and how to reason about tradeoffs between price, performance, and reliability.

*Building on Chapter 13:* Now that we understand how agents work as workflows, we need to understand what they cost—and how to spend wisely.

## Token Economics

AI pricing is fundamentally a token economy. Input tokens cost money. Output tokens cost money. Context tokens—everything you include in the prompt—double as input tokens. This creates a direct incentive to minimize context while maximizing its utility.

The math is straightforward: larger context means higher per-request costs. A 128k context window sounds generous until you realize that filling it costs roughly $0.01–$0.03 per request depending on the model. Multiply by thousands of daily requests, and context becomes the dominant cost driver.

**Case Study: The 10x Context Bill**

One team migrated from a 4k context model to a 128k context model without changing their system behavior. Their average request size went from 2k tokens to 40k tokens. The system performed better—more context meant fewer clarifying questions, better grounding, fewer retries. It also cost 20x more. The per-request improvement didn't justify the cost increase. They rolled back to a smaller context window and added better retrieval instead.

The lesson: context window capacity and context consumption are different things. A larger window gives you options. It doesn't require you to use them. The cheapest reliable solution often uses the smallest context that achieves the required reliability.

Token costs also vary by model. Claude Sonnet 4 is significantly cheaper than Opus for the same token volume. GPT-4o is cheaper than GPT-4 Turbo. When evaluating context engineering strategies, model selection is a cost decision first and a capability decision second.

## Context Windows as Budget Constraints

A context window isn't just a capacity—it's a budget. Every token you spend on system prompts, retrieval results, conversation history, and tool outputs reduces the budget available for the actual response. When the budget runs out, you face three options: truncate, compress, or fail.

Truncation is the default behavior for most systems. When context fills, the oldest content gets dropped. This works until the dropped content was the grounding information the model needed. Retrieval results get truncated. Conversation history gets truncated. The model loses track of what was established and starts contradicting itself.

Compression helps. Techniques like summarizing older conversation into a compact form recover token budget. But compression loses detail. A summarized conversation loses the specific phrasing, the exact examples, the nuanced preferences that made the original useful.

Failure is the worst option but the most honest one. When context exceeds capacity, some systems simply error. This is the correct behavior for reliability: fail explicitly rather than silently degrade. The system should know its budget and enforce it.

The cheapest reliable approach is to size the context window to the actual use case and enforce hard limits. Don't rely on the model to manage its own context budget. Build the budget into the system.

## Latency: The Hidden Cost of Context

Latency is the cost users pay in time. Every retrieval call, every tool invocation, every additional model call adds latency. Context engineering decisions directly affect response time.

Retrieval latency depends on the vector store, the query complexity, and the network topology. A local vector store might respond in 10ms. A remote one might take 200ms. Cross-region retrieval can exceed a second. When retrieval is part of every request, these milliseconds add up.

Inference latency scales with token count. More context tokens mean longer inference times. A 2k token request might complete in 500ms. A 40k token request might take 3 seconds. The model processes every token in the context window, so context size directly maps to latency.

Tool calls compound the problem. Each tool invocation requires a model call to decide to call it, a round-trip to execute it, and another model call to process the result. A system that makes three tool calls per request adds three round-trips to the latency budget.

**Case Study: Latency Budget Blowout**

A customer service bot was designed to retrieve product information, check inventory, and look up order status before responding. Each step was a separate tool call. The total latency: 8 seconds. Users abandoned the conversation. The system was technically correct but practically useless.

The fix was to batch the retrievals into a single call, cache the inventory data, and parallelize the order lookup. Latency dropped to 1.2 seconds. The context was smaller, the tool calls were fewer, and the system became usable.

The reliability question isn't just "does it work?" It's "does it work fast enough?" A correct answer that arrives too slowly is a failure.

## Retrieval Cost

Retrieval isn't free. Vector databases charge for storage and compute. Embedding generation costs per document. Re-ranking models add inference costs per result.

The most common inefficiency is over-retrieval. Systems fetch too many results "to be safe," then truncate half of them anyway. This wastes compute on embedding generation and ranking for results that never get used.

The fix is simple: retrieve only what you need. If the response typically uses three support documents, retrieve five. Not fifty. Retrieval should be sized to actual consumption, not to hypothetical maximums.

Embedding cost scales with document volume. Every document in your corpus needs to be embedded once during ingestion and potentially re-embedded when updated. For large corpora, embedding cost becomes significant. This is why chunking strategies matter—smaller chunks need more embeddings. Larger chunks need fewer embeddings but provide less targeted retrieval.

The cheapest reliable retrieval strategy uses the smallest chunk size that maintains precision, retrieves the smallest result set that covers the use case, and avoids re-ranking when simple similarity scores suffice.

## Tool Cost

Every tool call has a cost. Database queries consume compute. API calls consume credits. File reads consume I/O. When agents call tools liberally—exploring options, gathering information, verifying assumptions—the costs accumulate.

The pattern is familiar: an agent that calls ten tools to answer a simple question is spending far more than necessary. The reliability gain from thoroughness doesn't justify the expense.

**Case Study: The Tool-Happy Agent**

An agent was built to answer questions about code repositories. It would read files, run grep, invoke git log, check CI configurations, and query the issue tracker—often making 15–20 tool calls per question. Most questions could be answered with 3–4 tool calls. The extra calls were exploration, not necessary.

The fix was a tool budget: the agent could make up to 5 tool calls per question. If it needed more, it had to justify each additional call. This reduced costs by 70% while maintaining answer quality. The agent learned to be surgical, not exhaustive.

Tool cost optimization means metering usage, setting budgets, and designing tools that return compact results. A tool that returns a 5k token response is more expensive than one that returns a 200 token response, even if they both take the same execution time. Design tool schemas for minimal output.

## One-Shot vs Loops: The Reliability-Cost Tradeoff

The simplest agent pattern is one-shot: send a request, get a response, done. This is the cheapest pattern. No loops, no retries, no state management. The cost is fixed and predictable.

Loops multiply cost. A loop that runs three times costs three times as much as a single call. A loop that runs thirty times costs thirty times as much. Loops are also where reliability breaks down—each iteration is an opportunity for the model to drift, misinterpret, or compound an error.

**Case Study: The $100 Loop**

An agent was tasked with processing a batch of 100 records. Each record required a lookup, a transformation, and a write. The agent would loop through each record individually. The cost per record: $0.10. Total cost: $10.

A better approach: batch the records, process them in groups, or use a simpler model for the transformation. The new approach cost $0.01 per record. Total cost: $1. Same output, 10x cheaper.

The cheapest reliable solution is almost always a one-shot solution. Loops should be reserved for cases where iteration is genuinely necessary—where the problem structure requires it, not where the system design forces it.

When loops are necessary, enforce explicit termination conditions. Set maximum iteration counts. Track state explicitly rather than relying on the model's context to maintain loop invariants.

## Subagent Costs

Subagents—agents that delegate to smaller, focused agents—are a powerful pattern for managing complexity. They also multiply costs.

Each subagent call involves its own context setup, prompt injection, tool access, and inference. If a main agent spawns five subagents, it pays five times the baseline cost. If those subagents each make tool calls, the cost compounds further.

Subagents make sense when the reliability gain outweighs the cost increase—when a focused agent can solve a subproblem more reliably than a generalist agent trying to do everything. They make less sense when the subproblem is simple enough that the overhead exceeds the benefit.

The cost of subagents is why hierarchical agent architectures need explicit budgeting. A main agent should know how many subagents it can spawn, how many tokens each subagent can consume, and when the cost of delegation exceeds the cost of doing the work directly.

## Memory Consumption

Memory consumption in AI systems has two meanings: the model's context memory and the system's RAM consumption.

Context memory is the token budget discussed earlier. System RAM is the actual compute resource required to run the model, store the vector database, maintain the session state, and execute the tool layer.

Local models consume significant RAM. A 70B parameter model requires 140GB of RAM just to load. This is why local deployment is expensive: you're paying for the hardware to run the model, not just the inference cost.

Memory consumption also affects retrieval. A vector store that holds 10 million embeddings needs RAM to serve fast queries. If the vector store exceeds RAM capacity and spills to disk, query latency spikes. This is a common failure mode in production: the system works fine at small scale, then degrades dramatically at large scale because the vector store no longer fits in memory.

## Compute Tradeoffs: Local Models vs APIs

Local models eliminate per-token inference costs. Once you've paid for the hardware, inference is free. This makes local models attractive for high-volume workloads.

The tradeoff is upfront cost and capability. Local models are typically smaller and less capable than API models. A 7B local model might match a 2022 API model, but it won't match current frontier models. The reliability of local models depends on whether the task requires frontier capability or whether a smaller model suffices.

For many tasks, a smaller model is more than adequate. Classification, extraction, summarization, and simple transformation don't require GPT-4o capability. A 7B or 13B model can handle these tasks reliably at a fraction of the cost.

The decision framework: if the task requires frontier capability, use an API model and optimize context. If the task is within reach of a smaller model, use a local model and accept the capability ceiling.

## The Cheapest Reliable Solution

Across all these dimensions—tokens, latency, retrieval, tools, loops, subagents, memory, compute—the pattern is consistent. The cheapest reliable solution is the one that spends the minimum necessary to achieve the required reliability.

This requires knowing what "reliable" means for your use case. Some tasks need perfect accuracy. Others need "good enough" with a known error rate. Some need sub-second latency. Others can tolerate minutes. The cost optimization is different for each.

The engineering discipline is measurement. Track cost per request. Track reliability metrics. Track latency distributions. Compare different strategies empirically. The cheapest solution is the one that meets your reliability requirements at the lowest cost—not the one that minimizes cost at the expense of reliability, and not the one that overspends on capability it doesn't need.

Context engineering is the practice of making deliberate cost-reliability tradeoffs. Every token has a price. Every tool call has a price. Every loop iteration has a price. Engineer the context to spend only what the task requires, and no more.