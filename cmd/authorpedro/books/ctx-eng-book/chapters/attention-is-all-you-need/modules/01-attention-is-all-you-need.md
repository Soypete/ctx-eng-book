# Attention Is All You Need (But We Stopped Paying Attention)

## Beat 1: What the model actually sees

When you send a prompt to an LLM, something very specific happens: your text gets broken into tokens, each token becomes a vector of numbers, and then the transformer architecture figures out which tokens relate to which other tokens. That's attention in action.

The reliability question for any AI system is simple: **What information can the model actually see?**

The answer matters more than most engineers realize. If you can't answer this question precisely, you can't debug why your system failed. You can't optimize for cost. You can't reason about what the model will get right versus what it'll miss. Everything depends on understanding what's in the context window and how attention distributes across it.

## Beat 2: Tokens are not words

The first thing to internalize: tokens are not words. They're subword units that the model has learned to recognize based on statistical patterns in training data.

A token might be:
- A complete word ("the")
- A common suffix ("ing", "tion")
- A rare word broken into pieces ("unhappiness" → "un", "happ", "iness")

This matters because tokenization varies across models. The same sentence might be 50 tokens in GPT-4, 70 in Claude, and 80 in an open-source model. What counts as one "unit" of meaning depends on the tokenizer, and this affects:
- Context window limits (measured in tokens, not words)
- Pricing (billed per token)
- Where sentence boundaries fall for retrieval

When you're building context engineering systems, you need to think in tokens, not words. A 4000-token context limit doesn't mean 4000 words—it's usually 2500-3500 words depending on the text.

## Beat 3: Embeddings turn tokens into meaning

Once text is tokenized, each token gets converted into an embedding—a vector of floating-point numbers. This is where "meaning" enters the machine.

The magic of embeddings: tokens with similar meanings end up close together in the embedding space. "King" is near "queen." "Paris" is near "France." The model doesn't understand these relationships as a human would, but it has learned statistical patterns that make certain concepts cluster together.

What engineers need to understand: embeddings are fixed-dimensional representations. A 4096-dimensional embedding space doesn't "know" what any of those dimensions mean—they're just numbers that worked during training. When you retrieve documents using embedding similarity, you're matching based on training-pattern statistics, not semantic understanding.

This is why retrieval can fail in counterintuitive ways. A document about "database indexing" might not match a query about "how to speed up SQL queries" if the embedding model learned different statistical patterns for those terms. The information is there, but the model can't find it because it doesn't "see" the relationship the way a human would.

## Beat 4: Attention finds relevance

This is where the transformer earns its name. Attention is the mechanism that lets each token "attend to" other tokens in the context window to figure out what's relevant.

Here's how it works: for every token being processed, the model calculates attention scores against every other token in the context. These scores determine how much each token influences the next token's meaning. The model is essentially asking: "Given what I'm working on right now, which tokens in my context should I pay attention to?"

Key insight for systems engineers: **attention distributes across everything in the context window.** There's no pre-filtering. The model doesn't "know" that your system prompt is more important than the retrieved document or that the user's question matters more than the conversation history. It calculates attention weights across all tokens equally.

This is why the order and composition of your context matters so much. A 10,000-token context isn't just "more information"—it's a competition for attention. Every irrelevant token dilutes the signal for relevant ones. The model will confidently attend to the wrong information if that's what's in the window.

## Beat 5: Position encoding preserves order

Without position encoding, attention would treat tokens as a bag of words—it would know what tokens exist but not their order. "The dog bit the man" and "the man bit the dog" would be semantically equivalent, which is obviously wrong.

Transformers use sinusoidal position encodings to give each position a unique mathematical signature. This lets the model understand relative position: that cause comes before effect, that a function definition comes before its call site, that recent information is more relevant than ancient history.

What engineers miss: position encoding creates a recency bias. Tokens at the beginning and end of the context window get more distinct position signals, while middle tokens can become harder to distinguish. This is why instruction hierarchies matter—system prompts at the start might get less attention than you'd expect, and user messages at the end often dominate.

Poorly ordered context degrades retrieval even when the information is present. If your retrieved documents aren't ordered logically (definition before example, summary before detail), attention weights scatter to the wrong tokens. The model can't establish the right relationships.

## Beat 6: Context windows are fixed memory

The context window is the hard limit on what the model can see. Whatever fits in that window gets processed; everything else is invisible.

Current context windows:
- GPT-4: 128K tokens (but effective use is often less)
- Claude: 200K tokens
- Gemini: 2M tokens
- Open-source models: typically 4K-32K, some up to 128K

The critical mental model: think of the context window as a ticket tape. Everything you feed the model gets written on that tape, and attention reads from it linearly. There's no secondary memory. The model can't "go back" to a previous conversation unless that conversation is still in the window. It can't reference a document unless that document is loaded.

**This is the fundamental constraint of context engineering.** Everything that matters for a given request must fit in the window and survive the attention competition. If your context window is 8000 tokens and you need to include a 5000-token document, three previous messages, a system prompt, and the current user request—you're already at or over the limit. What gets dropped determines what the model can see.

## Beat 7: Why context is expensive

Context costs money and latency. Understanding why helps you make better engineering decisions.

**For API models (OpenAI, Anthropic):**
- Cost is per token. 128K context sounds great until you see the bill. A long conversation with large document retrieval can easily run $1-5 per interaction.
- Prompt tokens (input) and completion tokens (output) are both billed. More context = higher input costs.

**For self-hosted models:**
- Time to first token (TTFT) scales with context length. The model has to process all those tokens before generating anything.
- KV cache memory grows with context. A 128K context window needs gigabytes of GPU memory just to store the attention keys and values.
- Throughput drops as context grows. Same GPU, more work, fewer requests per second.

The practical question isn't just "what fits in the window" but "what's the right signal-to-cost ratio?" Adding more context always increases computation. It only sometimes increases quality. If your retrieved documents are irrelevant, adding them just adds noise and cost.

## Beat 8: The systems perspective

Here's what matters for building reliable AI systems:

1. **You control what the model sees.** The context window is your controlled input. What you include, exclude, order, and compress determines everything about the output.

2. **Attention is a competition.** Every token in context competes for attention. Irrelevant tokens don't just waste space—they actively degrade performance by drawing attention away from relevant information.

3. **Context windows have hard limits.** There's no infinite memory. At some point, something gets dropped. You need to decide what that is before the model decides for you.

4. **Cost is not optional.** Whether you're paying per token or running on your own GPUs, context has real costs. Engineering for cost means being intentional about what gets included, not just maximizing token counts.

5. **Position matters.** The model reasons about sequence. Poorly organized context degrades even correct information because attention can't establish the right relationships.

The transformer architecture gave us attention as a powerful mechanism for finding relevance. But attention only works with what it's given. If you stop paying attention to what's actually in the context window, the model will too—and it'll give you output that looks confident but is completely wrong for your situation.

That's not a model problem. That's a context engineering problem.

---

**Reliability Question Answered:** The model can only see what's in the context window. Attention distributes across all tokens equally. What you put there determines what the model will use—and what it'll ignore.