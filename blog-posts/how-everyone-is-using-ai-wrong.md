# How Everyone Is Using AI Wrong

*June 3, 2026*

---

Most people are still thinking about AI coding agents like they're traditional programming pair partners. They're not. And that mismatch is costing you tokens, accuracy, and sleep.

## The Pair Programming Mental Model

When developers first started using AI coding agents, they brought their habits from human pair programming:

> "Let's build this feature together."
> "Can you check if this looks right?"
> "What do you think we should do here?"

These sound natural. They're how we'd talk to a colleague. But they're fundamentally misaligned with how AI agents actually work—and it's not just a stylistic preference. It's causing real problems.

## The Pragmatics Problem

Here's what the research on computational pragmatics tells us: there's a gap between what you say and what you mean. Humans handle this gap through inference—we read context, tone, and convention to understand intent.

When you say "Can you check this?" you're making a request, not asking a genuine question about capability. The sentence looks like a question but functions as a directive. Humans infer the gap automatically. AI agents can too—but inference isn't free.

Every time the model has to infer your meaning, it's spending tokens on reasoning that could be spent on execution. More importantly, inference is probabilistic. The model might infer correctly. It might not.

## Freedom Is Expensive

There's another problem with the pair programming style: it implies freedom.

"Let's build this together" grants the agent interpretive freedom—it's invited to participate in decision-making, to suggest alternatives, to course-correct based on its judgment. This feels collaborative. It also allows for more mistakes.

When you say "Build this," you're constraining the action space. When you say "Let's build this together," you're opening it. More freedom means more paths the agent might take, more decisions it has to make, more chances to diverge from what you actually wanted.

This connects to what the GPT-3 few-shot learning paper showed: more context generally improves performance, but only when that context is structured to constrain the problem, not just demonstrate patterns. The model needs information that narrows the solution space, not just examples of similar problems.

## The Token Economics of Implication

Here's the practical math:

**Implied directive:**
```
"Can you check if the tests pass?"
- Speech act: indirect request
- Model must infer intent
- May ask clarifying questions
- Tokens: high (reasoning + potential follow-up)
```

**Explicit directive:**
```
"Run the test suite. Report pass/fail status."
- Speech act: direct request
- Model executes immediately
- No inference needed
- Tokens: low

```

The difference isn't just elegance. It's cost. Every inference step is tokens spent. Every ambiguous phrasing is a potential error requiring correction. The "natural" conversational style that feels good to humans is literally more expensive to run.

## What Direct Actually Looks Like

Being direct doesn't mean being rude. It means being unambiguous:

| Instead of... | Try... |
|--------------|--------|
| "Can you help me fix this bug?" | "Fix the bug in foo.c" |
| "What do you think about this approach?" | "Evaluate whether this approach handles edge case X" |
| "Let's see if the tests pass" | "Run tests. Report results." |
| "Does this look right?" | "Verify: user input is correctly validated. Confirm or report error." |

Notice the pattern: the right column tells the agent exactly what to do and what output to expect. No inference required.

## The Real Insight

The research points to a fundamental shift in how we should think about AI agents:

**Traditional pair programming:** Two minds collaborate. Both bring interpretation. Both can clarify. The conversation navigates ambiguity together.

**AI agent interaction:** You provide context. The agent executes. The more you constrain the problem space through explicit directives, the more reliable the execution.

This isn't about treating AI as less-than. It's about recognizing that AI agents work differently than human partners. They don't need to be taught through conversation—they need to be given structured information that lets them execute precisely.

The best AI interactions aren't conversations. They're specifications.

---

*This post synthesizes research from the GPT-3 few-shot learning paper (Brown et al., 2020) and computational pragmatics (Jurafsky). The core insight: context engineering is about replacing inference with explicit information—and that replacement is where the reliability lives.*