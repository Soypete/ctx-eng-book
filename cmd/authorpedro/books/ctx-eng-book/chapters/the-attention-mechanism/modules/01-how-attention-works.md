# How Attention Works

## Beat 1: What attention actually does

Attention is fundamentally a weighted similarity mechanism — it computes attention scores between the current token and all other tokens in the context window to determine which ones are most relevant. The model asks itself: "Which other words in this context are relevant to the word I'm working on right now?"

This is why longer contexts often produce better results: more tokens means more candidates for the clustering algorithm to draw from, increasing the likelihood that relevant information is available when attention runs.

## Beat 2: The critical assumption

Humans carry context with them. When you talk to a colleague, you share a frame of reference built from prior conversations, shared experiences, cultural background. You assume they understand what you mean even when you don't spell everything out.

Machines don't have this. When we talk to an AI, we unconsciously assume it brings the same contextual understanding — but it doesn't. The transformer's critical assumption is that all relevant information already exists in the context window. It's not pulling from prior interactions, shared understanding, or background knowledge. Everything it needs must be explicitly provided.

**What happens when this assumption fails:** The model does its best with what it has. It will confidently generate output based on incomplete or missing context. This is why a model can give a technically correct answer that is completely wrong for your specific situation — it attended to the wrong information because the right information wasn't there.

The assumption fails when context isn't directed to the model — when it doesn't have access to prior chats, relevant documents, or the specific code it's supposed to be working with. The model can't guess what context it should have; you have to explicitly provide it.

This is the pragmatics argument: semantics is not enough. The model understands the meaning of words, but it doesn't know what you mean unless you actively give it context. Pragmatics is about intent, about what's actually communicated — and that's something attention can't infer on its own.

## Beat 3: What attention does NOT solve

Your research notes list what the transformer does NOT solve: truth, provenance, trust, authorization, memory. It only answers: "Given information, what should I focus on?"

This distinction matters for building reliable systems because attention can't solve these problems — they have to be handled in the workflow and success criteria. Guardrails are deterministic. You have to build the system to handle truth verification, provenance tracking, authorization checks, and memory management outside of what attention does. The model will confidently route to wrong or untrusted information if that's what's in the context; it can't validate whether the information is true or allowed.

This is why context engineering exists as a discipline separate from model training.

These problems are handled through evals, guardrails, and error handling — this is where tooling like the Forge repo helps. You verify outputs, check against expected results, and catch failures at the system level, not at the attention level.

## Beat 4: Position encoding

Order matters for retrieval because attention uses relative position to understand relationships between tokens. Without position encoding, "the cat sat on the mat" and "mat the on sat cat the" would just be different sets of tokens with no structural difference.

The transformer's positional encodings (sinusoidal functions that create unique signatures for each position) make relative positions mathematically meaningful. This lets the model understand that "cause" typically precedes "effect," that "question" comes before "answer," that a function definition comes before its usage.

**What happens when context isn't ordered well:** The model loses the ability to reason about sequence. It can't tell which information is newer, which is a definition versus an example, or which instruction applies to which code block. Retrieval degrades because attention weights get distributed to the wrong tokens — not because the information isn't there, but because the model can't establish the right relationships between pieces of context.

## Beat 5: Context windows as fixed memory

Engineers should think about the context window as a ticket tape — the LLM streams information in and marks things for action based on training patterns. What you put in that tape matters. If it's nonsense, you lose value. You want everything in the ticket tape to matter, because attention is finite and will distribute its focus across everything in the window.

The practical implication: context engineering is the discipline of deciding what belongs on that ticket tape. A larger window doesn't help if the extra space is filled with irrelevant information — it just dilutes the signal. The constraint isn't just the token limit; it's the attention budget. Every irrelevant token competes with a relevant one.

## Beat 6: The cost of context

The tradeoff depends on how you're running the model:

- **Frontier/API models:** Cost per token and overall spend. Every additional token costs money, and larger contexts add up quickly.
- **Self-hosted:** Time to first token (TTFT), KV cache memory, and inference speed. Large contexts directly impact latency and require more memory for the key-value cache.

The practical question isn't just "what fits in the window" but "what's the right signal-to-cost ratio?" Adding more context always increases computation; it only sometimes increases quality.