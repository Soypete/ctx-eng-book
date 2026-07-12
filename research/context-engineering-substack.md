# Why I Hate the Term "Context Engineering" (Substack)

**Source:** https://substack.com/@soypetetech/p-190402075
**Date:** July 2026

## Core Thesis

The post argues that "context engineering" is misunderstood — it's not about dumping more context into prompts, but about **deterministically engineering how scoped data, semantics, tools, and instructions enter the system**.

## Key Arguments

### The Analogy: American vs Japanese Cinema
- American cinema: assumes shared cultural context (audience brings meaning)
- Japanese cinema: assumes no context (story is portable)
- AI systems are "Japanese cinema" — they start with **no context**

### The Bug: We Expect AI to Bring Context
- AI doesn't bring context — it starts empty
- Tool calls give ability to fetch more, but don't provide context
- Multi-bot and plugin failures: "they had the entire world open to them"

### Prompt Engineering Failed
- Prompt engineering was about dumping context upfront and hoping model gets it
- Context engineering decides: what info is introduced, when, how long it persists, what gets dropped

### The Real Problem: Data During Execution
- Engineers use data for analytics (what happened) and state (what recorded)
- ML training is retrospective — always history on history
- We're bad at using data **during execution** (online inference)
- Context engineering is NOT about "making the prompt better"
- It's about **how you introduce data — securely, reliably, with meaning — into a live system**

### System Prompts Are Products, Not Scratch Pads
- System prompts should be stable, deliberate, released intentionally
- They can be parameterized, have scoped instructions injected
- What happens AFTER the system prompt is probabilistic
- Context engineering shapes those probabilities, not the behavior

### More Instructions Don't Fix Context — Timing Does
- Adding more information makes it worse
- Information closer to system prompt / beginning of conversation is less likely to be used as plan evolves
- Models selectively ignore information to finish current task
- Context engineering is about **right information, right time, right place**
- Compaction trades recent detail for proximity to original instructions

### The Subtle Distinction: Memory vs Knowledge

**Two approaches to "memory" we've seen:**
1. ChatGPT: synthesizes all chat history, pulls through classification LLM to make suggestions
2. Claude: augments a memory file with things user wants to remember, adds after every session

**The problem with both:** They treat all conversation as potentially relevant. The real question is: **what is relevant?**

**The insight:** We don't need "memory." We need **knowledge**.

If we're working in a domain and acquire knowledge relevant to future work in that domain, that knowledge should be stored in the **domain classification**.

What makes it a **knowledge store** is the ability to reason across a domain and make connections.

**The missing piece:** Making connections — LLMs don't do this without good pragmatics. Loops and agent harnesses aren't good at this.

This is why "memory" isn't usable context engineering — it's still the problem of "if we give the agent the most context possible, it will figure it out."

**That goes against the book's thesis.**

---

## Key Quotes

> "The real work isn't letting models 'figure it out;' it's deterministically engineering how scoped data, semantics, tools, and instructions enter the system so we can shape outcomes without handing the model all the information in the world."

> "We don't make the model deterministic, we make the inputs deterministic. That's the contract."

> "More instructions don't fix context — timing does."

> "Stop letting AI 'figure it out,' and start engineering how data with meaning enters the system."

---

## Alignment with Book Thesis

This post is essentially the **manifesto** for the book's thesis. Key alignments:

| Post Concept | Book Pillar |
|--------------|-------------|
| Deterministic inputs | Pragmatics |
| Data with meaning | Semantics |
| Scoped data | Data |
| What info introduced + when | State / Timing |
| System prompts as product | Governance |

The distinction between "memory" (storing everything) and "knowledge" (storing domain-relevant connections) is crucial for Ch6 and Ch8.