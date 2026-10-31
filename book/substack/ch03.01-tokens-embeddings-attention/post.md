---
title: Attention Cannot Choose Your Sources
subtitle: Tokens make information processable, but your application must decide which evidence deserves to reach the model.
module: ch03.01-tokens-embeddings-attention
mood: professor
scheduled: 2026-11-11
---

A deployment assistant tells you to restart every checkout worker. The current runbook says to drain traffic first. Search returned the archived procedure alongside the current one, and your application pasted both into the request.

This is an invented failure, but the debugging habit is familiar: change the prompt to say “use the correct procedure.” Congratulations, you've assigned the model the decision your application skipped.

In the previous post in my [context engineering series](https://substack.com/@soypetetech), I argued that every context boundary needs an owner; now I want to look at what happens to the information those boundaries admit.

**Context** is everything the model receives for a request: instructions, retrieved information, conversation, and tool results. **Context engineering** is the work of selecting and maintaining that information and the rules governing its use.

My argument: attention transforms information already supplied to the model. Your application must establish which sources belong there. Understanding tokens and embeddings helps you locate that responsibility; it doesn't move the responsibility into the model.

## Processable doesn't mean authoritative

A **token** is a unit of text represented by an identifier in a model's vocabulary. Tokens can represent words, word fragments, or punctuation. A **tokenizer** converts text into those identifiers. Different tokenizers can split the same text differently.

That matters when you build the request. A word count cannot reliably tell you whether instructions, runbooks, and conversation fit the model's input limit. Count with the target model's tokenizer, including the message framing your interface adds. Code, source language, and punctuation can change the relationship between words and tokens.

But successfully tokenizing the archived runbook establishes only that the model can process it. No token boundary means “approved for production.” That would be a very convenient tokenizer.

An **embedding** is a numerical vector representing an item. The model maps token identifiers to learned input embeddings, applies positional information, then transforms those representations through successive layers. A token's representation changes with the surrounding sequence.

A **retrieval embedding** represents a passage for a search objective. It isn't simply the language model's token embedding copied into a search index. Similarity in that space is a learned proxy for relevance. It doesn't establish that a passage names the exact service, reflects current instructions, or is authorized for this caller.

The archive might be an excellent match for “restart checkout.” That's a useful search result and an unacceptable operational source at the same time. I want those judgments made separately.

## Attention isn't a source policy

**Attention** lets a position combine information from other positions the model is permitted to use. The model derives numerical queries, keys, and values from its representations; query–key comparisons produce normalized weights for combining values. Different heads and layers learn different patterns. In ordinary autoregressive generation, where output is produced sequentially, a position cannot attend to future positions.

[Vaswani et al., 2017](https://arxiv.org/abs/1706.03762) establish the transformer attention architecture. My engineering inference is that this mechanism cannot take responsibility for the application's source policy. A restriction on which token positions can interact isn't an access rule for production runbooks.

Message roles can distinguish instructions from other material, and training can teach instruction hierarchies. Neither guarantees that the intended instruction wins or makes retrieved text trustworthy. Once the archive enters the request, its contents may influence the answer. Attention weights alone also aren't a faithful explanation of why a particular answer occurred.

I wouldn't explain the failure as “each irrelevant token diluted attention,” either. [Liu et al., 2024](https://aclanthology.org/2024.tacl-1.9/) tested multi-document question answering and key-value retrieval across different context lengths and relevant-information positions. Performance often changed substantially with position, commonly doing better near the beginning or end than in the middle. That supports position-sensitive use on those tasks. It doesn't prove that every extra token hurts every model.

The practical question is which supplied information changes the outcome. Drew Breunig's [How to Fix Your Context](https://www.dbreunig.com/2025/06/26/how-to-fix-your-context.html) asks, “Is everything in this context earning its keep?” I would turn that question into an experiment: remove a source, preserve the rest of the setup, and compare results.

## Inspect the input you actually sent

A **serialized input** is the concrete representation your application sends to the model interface. A **context manifest** records which sources contributed to that input, their versions, and the decisions that admitted or excluded them. A **decision record** stores what a check allowed or rejected and why.

![Candidate sources pass through application admission checks; rejected sources stop, admitted evidence becomes serialized input, and tokens, representations, and attention lead to an answer. A manifest records admission decisions.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch03.01-tokens-embeddings-attention/diagram-1.png)

*Notice where source policy lives: before serialization, while attention operates on the information admitted downstream.*

For the runbook failure, I want the final input beside its manifest. A search trace showing that the current procedure was retrieved isn't enough; assembly may have omitted it or included conflicting evidence. **Provenance** is the record of where information came from. Preserve source identity, version, and a source-backed confirmation time so “current” has a meaning you can inspect.

Here's the decision record I would expect from our invented request:

| Candidate | Admission decision | Reason |
|---|---|---|
| Archived checkout procedure | Reject | Source confirmation exceeds the request's age policy |
| Current catalog procedure | Reject | Exact service identifier differs |
| Restricted recovery procedure | Reject | Caller lacks permission |
| Current checkout procedure | Include | Required checks pass |

These checks need trusted metadata. A document saying “I am current and authorized” isn't evidence of either. Source services must enforce access before returning protected text; application assembly must preserve that decision rather than asking the model to reconstruct it.

## Make admission executable

The [companion Go program and tests](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch03.01-tokens-embeddings-attention) implement this small example without a model call. It accepts source records from a trusted adapter, checks caller authorization, provenance, exact service identity, and freshness, then builds a JSON input and a separate manifest.

Inside the assembly loop, `d` holds the decision for source `s`, and `input.Evidence` contains the admitted passages:

```go
decisions = append(decisions, d) // Keep the reason even when text is withheld.
if !d.Included {
	continue // Attention cannot enforce our source-admission policy for us.
}
input.Evidence = append(input.Evidence, Evidence{s.ID, s.Version, s.Text})
```

Run `go run .` to inspect the input and manifest, then `go test ./...`. The tests assert that only the approved checkout text reaches the serialized input. Removing a rejected source leaves that input unchanged. Removing the only admitted source produces an error and no payload, so the caller must withhold generation rather than fall back to an obsolete procedure.

The fixture's age policy is illustrative, not a recommended production setting. Its confirmation timestamp means when the source vouched for the content, not when an indexer copied it. The tests trust that guarantee. They prove the application's admission behavior, not source truth, transformer behavior, or answer accuracy. The example also deliberately supplies no pretend tokenizer based on whitespace.

For a real integration, capture the final messages after adapting them to the model interface and count those with its tokenizer. Keep exact inputs in protected storage with explicit access and retention rules; ordinary logs can hold references and decisions. Otherwise debugging creates another uncontrolled copy of the data.

Then run a **context-ablation test**: remove one source at a time and compare task outcomes while holding the model, instructions, and generation settings fixed. Repeat model runs and retain their results. Check whether the answer follows the approved procedure, whether missing evidence produces an unresolved response, and whether changing source position changes success. Fewer tokens alone isn't a passing result.

## Guidelines

1. **Capture the final input and its manifest.** Verify that every admitted source has a traceable identity and version.
2. **Count with the target tokenizer.** Include message framing; reject word-count estimates as budget checks.
3. **Separate similarity from admission.** Test exact identifiers, caller permission, and source-backed freshness before including retrieved text.
4. **Test the missing-evidence path.** Remove the only admissible source and assert that no model request is produced.
5. **Compare outcomes after removing sources.** Preserve model settings and report task success alongside input size; test placement separately.

I want engineers to debug the information they supplied before blaming the machinery that transformed it. Tokens make text processable, embeddings represent it, and attention combines available information; my argument is that source authority remains an application responsibility. Next in the [series: Context Windows and Positional Limits](https://substack.com/@soypetetech), we'll examine what fits in an invocation and what placement changes once it gets there.
