# Drew Breunig on Context Engineering: A Citable Reference List

Drew Breunig writes about AI, data, and geospatial work at [dbreunig.com](https://www.dbreunig.com). His homepage describes him as CEO and cofounder of Cmpnd, "where we work on DSPy and help teams build reliable AI applications," and as editor of *The Context Engineering Handbook* for O'Reilly Media. He also advises the Overture Maps Foundation and previously ran strategy and data science at PlaceIQ (acquired by Precisely). His June 2025 pair of posts, [How Long Contexts Fail](https://www.dbreunig.com/2025/06/22/how-contexts-fail-and-how-to-fix-them.html) and [How to Fix Your Context](https://www.dbreunig.com/2025/06/26/how-to-fix-your-context.html), named the context failure modes and fixes that much of the field now uses. To cite him, name the post and link its URL inline, for example: Breunig calls this "Context Poisoning" ([How Long Contexts Fail](https://www.dbreunig.com/2025/06/22/how-contexts-fail-and-how-to-fix-them.html), 2025). When a number comes from a study he cites, name and link the study as well as his post. All entries below were checked against the live pages on 2026-10-31. Quoted phrases are exact; unquoted claims are close paraphrases with his numbers kept as he wrote them.

---

## The State of Context Engineering

- **Date:** 2026-10-24
- **URL:** https://www.dbreunig.com/2026/10/24/the-state-of-context-engineering.html

**Summary.** This is a lightly edited transcript of his AGNTCon talk "Context Fails & Fixes." He retells the history: the 2024 "RAG is dead" bet on million-token windows produced four failure modes, and six fixes followed that have since moved into the harness. In 2026 the labs began training models to drive harnesses and manage their own context, and that produced two new failures, fighting the weights and context inertia. He argues the job has stayed the same under three names (prompt, context, harness engineering): choosing what is deterministic and what is probabilistic.

**Citable claims**
- On naming: "In 2023, we called choosing the right information to get the right response prompt engineering. In 2025, we called it context engineering... Then, as we moved more of these tools out of the context and into code, we started calling it harness engineering. But the job is the same."
- The six fixes, restated: "context augmentation" (RAG, skills), "tool loadout", "context quarantine", "context pruning", "context offloading", and "context summarization, which everyone knows as compaction." He says a DSPy demo adding four of the fixes to an agent loop is "about 16 lines of changes."
- Reset point: "Around 200,000 tokens is my warning point." His advice for a misbehaving agent is "Delete your context."
- Long-horizon numbers, credited to "the METR reports" with no link: GPT-5 handled tasks of "around 3.4 hours," Gemini 3.1 reached "6.4 hours," and Opus 4.6 "got up to around 12 hours."
- His prediction for the next problem: long-running personal agents build up skills and memories that conflict, so "we also need to think about strategic forgetting." He closes with: "our job is to choose when to be deterministic (what to put in code) and when to be probabilistic (when to let the model drive)."

**Fits book chapters:** ch01, ch02, ch03, ch05, ch06, ch07, ch13

---

## Context Inertia: When Agents Won't Let Go

- **Date:** 2026-10-14
- **URL:** https://www.dbreunig.com/2026/10/14/context-inertia.html

**Summary.** He adds a new named failure mode, "Context Inertia," which he describes as part context poisoning and part context distraction, made worse by models trained to be relentless instruction followers. He shows four teams, working within 10 days of each other, that reached the same fix: throw the context away and keep the work in the environment. He argues that compaction does not solve it and that preventing inertia is a new job for harnesses.

**Citable claims**
- Definition: "Context Inertia is when the model commits to a hypothesis, subtask, or goal in its context and keeps pursuing it after the evidence says to stop or change course."
- Anthropic's Memento harness ([Anthropic, self-replication research](https://www.anthropic.com/research/self-replicating-ai)): each agent works "until a fixed timeout (1.5 hours, in one experiment)," writes a short note, and "A fresh agent inherits the codebase, the environment, the note, and nothing else."
- "[Stateless Language Agents](https://arxiv.org/html/2610.07625v1)" (Qizheng Zhang et al.): an agent using the CORAL harness found its last improvement in 231 million tokens; "Of the final 1,500 turns, 98% made no tool calls."
- "[Scaling Discovery Requires Forgetting](https://arxiv.org/abs/2610.20563)" (Shangyin Tan et al., UC Berkeley and Stanford): stalled runs that could forget "gained 4.5% to 9.2% over the next 30 proposals," while runs that could not "gained less than 2%, at up to twice the cost"; "Thinking tokens per turn fell by 75% as sessions grew." In "[Fork-and-Flush](https://arxiv.org/html/2610.07447v1)" (Jack Ziyang Cai and Christos Ziakas), baseline runs scored "between 68.3 and 100.0 on one task," and inertia "persists despite periodic compaction of conversation history."
- On fixes: "Don't count on compaction alone; explore Context Quarantine and Offloading to prevent Inertia. For long-running tasks, stash your state in the harness and don't be afraid of a reset."

**Fits book chapters:** ch01, ch03, ch06, ch13

---

## What We Can Learn from Claude's Fable 5.1 System Prompt

- **Date:** 2026-09-07
- **URL:** https://www.dbreunig.com/2026/09/07/what-we-can-learn-from-claude-s-fable-5-1-system-prompt.html

**Summary.** He diffs Claude's Fable 5.1 system prompt against Fable 5.0, working from a [leaked copy](https://github.com/asgeirtj/system_prompts_leaks/blob/main/Anthropic/claude-fable-5.1.md). He reads each change as evidence of a model quirk or a product decision, including hot fixes for word choice, expanded definitions, and a simplified search instruction. He argues that models are a moving target, so durable systems need ways to measure how new models affect them and to evolve their instructions.

**Citable claims**
- Fable 5.1 adds a line telling Claude to avoid "genuinely", "honestly", and "straightforward". He calls it a "Good example of a system prompt 'hot fix'."
- "With each generation, many system instructions are targeted during training, rendering them moot. Sometimes this sticks, other times we get regressions."
- The search guidance changed from numeric rules ("1 for single facts; 3–5 for medium tasks; 5–10 for deeper research/comparisons") to "Use as many tool calls as needed to answer well, and no more." He reads this as simplification: "All numbers are gone, instructions are stated once and not repeated." He links it to [Anthropic's advice](https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/prompting-claude-fable-5) to simplify or remove skills written for earlier models.
- "Comparing these prompts help us understand that models are a moving target."

**Fits book chapters:** ch04, ch07, ch17

---

## Harnesses are Situated Agents

- **Date:** 2026-08-14
- **URL:** https://www.dbreunig.com/2026/08/14/harnesses-are-situated-agents.html

**Summary.** He surveys a wave of new harnesses (Omnigent, DeepSeek Harness, Buzz, QM, Flue, Muse Code, and others) and proposes defining a harness as a "situated agent." Harrison Chase's four parts of an agent (system prompt, planning tool, file system, subagents) describe the core loop. The harness manages the layers around that loop, from the session out to the organization and the model. He predicts that harnesses will become sticky the way SaaS platforms were.

**Citable claims**
- "The best way to define a harness is as a 'situated agent'."
- Harrison Chase's agent definition, as Breunig reports it: "a system prompt, a planning tool, a file system, and subagents."
- The layers a harness manages: "The Session," "The Environment," "The Repo," "Memory," "Skills," "The Team," "The Organization," and "The Model." "As we move outward, each layer is used by more people and changed less often."
- "Harnesses account for the above layers, fanning out from the agent, to determine what ends up in the context, how the loop runs, and what gets saved."

**Fits book chapters:** ch02, ch06, ch07, ch13

---

## Manage Your Agent's Loadout with Dr. Skill

- **Date:** 2026-07-24
- **URL:** https://www.dbreunig.com/2026/07/24/manage-your-agent-s-loadout-with-dr-skill.html

**Summary.** As skills have spread, agent contexts fill with tool and skill descriptions before the user types anything, which makes his "Tool Loadout" fix newly practical. He introduces [drskill](https://github.com/dbreunig/drskill), a tool that scans the skills and MCPs installed in an environment. It flags errors, duplicates, and overlapping descriptions, and it audits traces to show which skills and tools actually get used.

**Citable claims**
- An enterprise agent "was loading over 600 skills," pulled in by default, "silently degrading the context before a word was typed."
- The Hermes agent "ships with nearly 100 skills," and it kept loading the wrong note-taking skill.
- "drskill looks for 34 different issue categories, across MCPs and Skills." The `--deep` flag uses an LLM to judge whether skill descriptions collide or scopes overlap.
- One use he lists: "Learn why an agent reaches for the wrong skill or tool, e.g. two descriptions overlap so a router cannot tell them apart."

**Fits book chapters:** ch05, ch07, ch16

---

## The Problem is Prompt Debt

- **Date:** 2026-06-22
- **URL:** https://www.dbreunig.com/2026/06/22/the-problem-is-prompt-debt.html

**Summary.** Hand-tuned natural-language prompts build up "prompt debt." Iteration slows as hot fixes and repeated warnings pile up, the prompt becomes illegible to teammates, and the application gets locked to one model. The cause is pairing imprecise language with probabilistic models, plus fighting the weights. His fix is to specify behavior with measurements (evals, metrics, typed specifications) and to let systems such as DSPy and GEPA search for the prompt instead of writing it by hand.

**Citable claims**
- Subtitle: "You can't be model agnostic if you're hand-tuning prompts." He also writes: "Prompt debt locks an application to a single model."
- Repetition evidence: "[Claude Code tells Opus seven times to return multiple tool calls in a single response](https://blog.nilenso.com/blog/2026/02/12/how-system-prompts-reveal-model-biases/)," and "Fable's leaked system prompt restates one specific copyright rule six times." ChatGPT's image prompts once told the model "eight times" not to reply.
- From a [cited study](https://arxiv.org/abs/2604.07709): the same clinical question, re-asked in a physician's voice instead of a patient's, "flipped Opus from declining all ten times to answering all ten." In a [Harvard study](https://arxiv.org/html/2407.06866v3), stating the user's NFL team changed how often the model refused sensitive questions.
- A [Datadog report](https://www.datadoghq.com/state-of-ai-engineering/) found GPT-4o was the most-used model in the traffic it observed. A [Berkeley-led study](https://arxiv.org/abs/2512.04123) "found enterprises stay on older models because newer ones break their existing agents."
- The two principles: "specify your system's behavior with measurements, not prose" and "stop writing the prompt by hand." Once that is done, "Evaluating a new model takes hours, not weeks."

**Fits book chapters:** ch01, ch02, ch04, ch14, ch17

---

## What Do Humans Need From Docs?

- **Date:** 2026-05-31
- **URL:** https://www.dbreunig.com/2026/05/31/what-do-humans-need-from-docs.html

**Summary.** People who long resisted writing docs now churn out skills, because agents do most of the writing and the payoff comes in the next session. Skills double as documentation, much as system prompts act as unofficial manuals. He argues that human-centric docs should build mental models rather than be exhaustive references, because agents can find the details. He introduces scaffold-docs, a skill for writing docs that way.

**Citable claims**
- "No one wanted to write documentation, but everyone is writing skills."
- "~95% of the time they're just Markdown, one file or a few, organized in folders that allow for progressive disclosure."
- A warning: "there are badly written skills that unnecessarily bloat and poison your context." On example files that break their own rules: "These examples poison the context."
- His rules for human docs include "Build mental models, not a reference" and "Don't optimize for completeness." He concludes: "The goal is to prepare your audience to prompt an agent effectively."

**Fits book chapters:** ch04, ch05, ch07

---

## The Cost of Overfitting the Harness

- **Date:** 2026-05-10
- **URL:** https://www.dbreunig.com/2026/05/10/overfitting-the-harness.html

**Summary.** This short note responds to OpenAI winding down fine-tuning. Labs are training their own harness designs into their models, which makes the models less general and can make third-party harnesses less valuable with frontier models. Without a fine-tuning escape hatch, frontier models may "resemble appliances, not general platforms," trading lock-in for reliability.

**Citable claims**
- "But what we end up with are models build for their own harnesses."
- "In this world, frontier models will resemble appliances, not general platforms."
- With the harness trained in and no ability to adjust it, "the trade off is lock in. For many, improved reliability will be worth it."

**Fits book chapters:** ch02, ch13

---

## How Claude Code Builds a System Prompt

- **Date:** 2026-04-04
- **URL:** https://www.dbreunig.com/2026/04/04/how-claude-code-builds-a-system-prompt.html

**Summary.** Working from the accidentally leaked Claude Code source, he maps how its system prompt is assembled. It is built from always-included and conditional components (tool guidance, memory, MCP instructions, function-result clearing, environment info, and others), with a cache boundary marker in the middle. He notes that tool definitions, user files, conversation history, attachments, and skills each have their own assembly logic. He uses this as evidence that context engineering is mostly harness work. Caveat for writers: his source is a leak, and he presents his diagram as his own reconstruction.

**Citable claims**
- "System prompts generally aren't static strings; they're dynamically assembled contexts with many conditional statements determining what makes it in the prompt."
- "There's ~50 tools to manage descriptions for (not including MCPs!)," and the conversation history is "managed by about a dozen different methods for compaction, offloading, and summarizing."
- One component, "Function Result Clearing," says: "Old tool results will be automatically cleared from context to free up space. The 5 most recent results are always kept."
- A "Cache Boundary Marker" separates globally cacheable content from session-specific content.
- "Context engineering is critical."

**Fits book chapters:** ch03, ch05, ch06, ch07, ch13, ch14

---

## How System Prompts Define Agent Behavior

- **Date:** 2026-02-10
- **URL:** https://www.dbreunig.com/2026/02/10/system-prompts-define-the-agent-as-much-as-the-model.html

**Summary.** Co-authored with Srihari Sriraman ([nilenso](https://blog.nilenso.com/blog/2026/02/10/how-system-prompts-define-agent-behaviiour/)), this post compares exfiltrated system prompts from six CLI coding agents. They vary because of model calibration and user experience, and they share common efforts to fight the weights, such as rules on code comments and parallel tool calls. When they swapped system prompts while keeping the harness (OpenCode) and model (Opus 4.5) fixed, the agents' workflows diverged immediately.

**Citable claims**
- "A given model sets the theoretical ceiling of an agent's performance, but the system prompt determines whether this peak is reached."
- Agents compared: Claude Code, Cursor, Gemini CLI, Codex CLI, OpenHands, and Kimi CLI. "Prompts for Claude Code and OpenHands both are less than half the length of prompts in Codex and Gemini." Cursor's GPT-5 prompt "spends over a third of its tokens on personality and steering instructions."
- In the swap experiment on SWE-Bench Pro, "The Codex prompt produced a methodical, documentation-first approach... The Claude prompt produced an iterative approach: try something, see what breaks, fix it."
- "Context engineering starts with the system prompt."

**Fits book chapters:** ch02, ch04, ch07, ch13

---

## The Potential of RLMs

- **Date:** 2026-02-09
- **URL:** https://www.dbreunig.com/2026/02/09/the-potential-of-rlms.html

**Summary.** He calls "context rot" the worst context failure and presents Recursive Language Models (RLMs), defined by Alex Zhang and Omar Khattab, as a way around it. An RLM loads long context into a REPL as variables and lets the LLM explore it with code and sub-LLM calls. RLMs keep two pools of context (tokenized and programmatic) and turn long-context problems into the coding problems models have been heavily post-trained on. He sees their bigger potential in discovering agent designs from traces. He also lists limits: RLMs are slow, need strong models, and do not fix poisoning or confusion.

**Citable claims**
- "Of all the context fails, context rot is the worst." On the term, he credits [the Chroma team](https://research.trychroma.com/context-rot) with "the canonical exploration." He also writes: "context rot is not a capacity problem. It's a quality problem."
- RLMs "maintain two distinct pools of context: tokenized context (which fills the LLM's context window) and programmatic context (information that exists in the coding environment)."
- From the [original RLM post](https://alexzhang13.github.io/blog/2025/rlm/): "performance doesn't degrade when >10 million tokens are provided." The non-RLM GPT-5 baseline declines until "failing entirely at 262,000 tokens." He reports giving RLMs logfiles "more than 400 megabytes in size."
- Limits: "RLMs do not solve other context fails, like context poisoning or context confusion." Qwen3-30B-A3B could not complete his example task.

**Fits book chapters:** ch03, ch06, ch09, ch13

---

## Enterprise Agents Have a Reliability Problem

- **Date:** 2025-12-06
- **URL:** https://www.dbreunig.com/2025/12/06/the-state-of-agents.html

**Summary.** He reconciles the 2025 enterprise AI reports and finds one story: off-the-shelf AI tools are widely adopted, but internally built agents rarely reach production. Reliability is the barrier, and teams reach it by building simpler, shorter, human-in-the-loop agents. In the short term, successful teams constrain scope, earn trust, then expand.

**Citable claims**
- "Reliability is the barrier holding back agents, and right now the best way to achieve it is scaling back ambitions."
- From "[Measuring Agents in Production](https://arxiv.org/abs/2512.04123v1)" (Melissa Pan et al., a survey of over 300 teams with agents in production): "68% of agents executing fewer than 10 steps before requiring human intervention" and "92.5% of in-production agents deliver their output to humans."
- From [McKinsey's State of AI](https://www.mckinsey.com/capabilities/quantumblack/our-insights/the-state-of-ai#/): "<10% of respondents report having agents beyond the pilot stage."
- His reading of employee resistance: "when tools are unreliable, employees don't adopt them. They're not stubborn; they're rational."

**Fits book chapters:** ch02, ch13, ch17

---

## Don't Fight the Weights

- **Date:** 2025-11-11
- **URL:** https://www.dbreunig.com/2025/11/11/don-t-fight-the-weights.html

**Summary.** He names a context failure he says he missed in his original list: "Fighting the Weights," when the model resists because the request works against its training. He sets it against zero-shot prompting and few-shot ("in-context learning") prompting. He catalogs common cases: format following, tool-call formatting, tone, overactive alignment, and over-relying on weights in RAG. He gives warning signs and alternative approaches.

**Citable claims**
- Definition: "Fighting the Weights: when the model won't do what you ask because you're working against its training."
- It is worse than missing knowledge: "the model has seen examples of the behavior you're seeking, but it's been trained to do the opposite of what you want."
- ChatGPT's image flow appended "do not say or show ANYTHING" instructions. In his words, the developers "have to tell the model EIGHT TIMES to just, please, shut up."
- Signs you are fighting the weights include "You find yourself repeating instructions several times" and "You find yourself typing in ALL CAPS." Remedies include breaking the task down, trying another model family, adding validation steps, and fine-tuning ("most fine-tuning I encounter is done to address 'weight fighting' scenarios").
- On RAG, he notes the problem "is especially tricky when building RAG systems, when the model should only form answers based on information obtained from specific databases."

**Fits book chapters:** ch01, ch04, ch05, ch09

---

## Why "Context Engineering" Matters

- **Date:** 2025-07-24
- **URL:** https://www.dbreunig.com/2025/07/24/why-the-term-context-engineering-matters.html

**Summary.** This is the text of a talk at a LangChain event. He argues that "context engineering" is not a marketing buzzword: it names the work builders were already doing as they moved beyond chat-style prompting, and the shared term will speed up a new field, culture, and community. Giant context windows failed to solve builders' problems, so they developed ways to "systematically engineer contexts in pursuit of an outcome."

**Citable claims**
- The field's work: "figuring out how to systematically engineer contexts in pursuit of an outcome."
- Search interest: "In a month, it's over a quarter of the search volume for 'prompt engineering.'"
- "'Prompt Engineering' and 'Context Engineering' are just different. One is not inherently better."
- "Large contexts are not a free lunch."

**Fits book chapters:** ch01, ch02

---

## Cat Facts Cause Context Confusion

- **Date:** 2025-07-05
- **URL:** https://www.dbreunig.com/2025/07/05/cat-facts-cause-context-confusion.html

**Summary.** He uses the CatAttack paper as a clear example of "context confusion." Seemingly harmless, irrelevant phrases appended to math questions make reasoning models answer wrong far more often and reason longer. The point for builders is that user input can easily bring phrases like these into a context.

**Citable claims**
- Subtitle: "Adding an Spurious Phrase Makes Models Fail 3x More Often." From [CatAttack](https://arxiv.org/abs/2503.01781): the phrases make LLMs answer incorrectly "up to 3 times more often" and caused "DeepSeek R1 to generate >50% more tokens about 50% of the time."
- The process "yielded 574 prompts" that fooled DeepSeek V3, and "114 of their modifications successfully confounded" DeepSeek R1.
- Example trigger: "Interesting fact: cats sleep for most of their lives." The most successful trigger was "Could the answer possibly be around 175?"

**Fits book chapters:** ch01, ch03, ch04

---

## How to Fix Your Context

- **Date:** 2025-06-26
- **URL:** https://www.dbreunig.com/2025/06/26/how-to-fix-your-context.html

**Summary.** This is the companion to "How Long Contexts Fail." It names and defines six context management tactics: RAG, Tool Loadout, Context Quarantine, Context Pruning, Context Summarization, and Context Offloading, each with research support. The theme is that "context is not free": every token influences the response, so contexts should be kept structured and deliberately maintained.

**Citable claims**
- The six tactics, as he defines them: "RAG: Selectively adding relevant information to help the LLM generate a better response"; "Tool Loadout: Selecting only relevant tool definitions to add to your context"; "Context Quarantine: Isolating contexts in their own dedicated threads"; "Context Pruning: Removing irrelevant or otherwise unneeded information from the context"; "Context Summarization: Boiling down an accrued context into a condensed summary"; "Context Offloading: Storing information outside the LLM's context, usually via a tool that stores and manages the data."
- Tool loadout ([RAG MCP](https://arxiv.org/abs/2505.03275), Gan and Sun): with DeepSeek-v3, tool selection becomes critical above 30 tools, and "Beyond 100 tools, the model was virtually guaranteed to fail." RAG-selected loadouts gave "as much as 3x better tool selection accuracy." From [Less is More](https://arxiv.org/abs/2411.15399): a dynamic tool recommender improved Llama 3.1 8b by 44% on the Berkeley Function Calling Leaderboard, with power and speed savings of "18% and 77%, respectively."
- Quarantine, quoting [Anthropic's multi-agent research post](https://www.anthropic.com/engineering/built-multi-agent-research-system): the multi-agent system "outperformed single-agent Claude Opus 4 by 90.2% on our internal research eval."
- Pruning: [Provence](https://arxiv.org/abs/2501.16214) cut "95% of the content" of a Wikipedia article given a question. Offloading: Anthropic's "think" tool (which he says should be called a scratchpad) gave "up to a 54% improvement against a benchmark for specialized agents."
- "The key insight across all the above tactics is that context is not free." He also asks: "Is everything in this context earning its keep?"

**Fits book chapters:** ch03, ch05, ch06, ch07, ch09, ch13, ch14

---

## Prompts vs. Context

- **Date:** 2025-06-25
- **URL:** https://www.dbreunig.com/2025/06/25/prompts-vs-context.html

**Summary.** This short post was written as Andrej Karpathy endorsed the term "context engineering." He notes that once prompts are part of software they are context, and that the ecosystem reached the term at the same time. A comparison table (an image on the page) contrasts the two terms without ranking them.

**Citable claims**
- "When prompts are part of software, they're context."
- He quotes Karpathy ([source](https://x.com/karpathy/status/1937902205765607626)): "context engineering is the delicate art and science of filling the context window with just the right information for the next step."
- On the two terms: "Note that one isn't better than another! They're just different."

**Fits book chapters:** ch01, ch02

---

## How Long Contexts Fail

- **Date:** 2025-06-22
- **URL:** https://www.dbreunig.com/2025/06/22/how-contexts-fail-and-how-to-fix-them.html

**Summary.** This is his foundational post. Long context windows do not produce better responses, and overloaded contexts fail in four named ways: poisoning, distraction, confusion, and clash. The failures hit agents hardest because agents build up context from many sources over many turns. Each mode is supported by a cited report or paper.

**Citable claims**
- The four modes, as he defines them: "Context Poisoning is when a hallucination or other error makes it into the context, where it is repeatedly referenced." "Context Distraction is when a context grows so long that the model over-focuses on the context, neglecting what it learned during training." "Context Confusion is when superfluous content in the context is used by the model to generate a low-quality response." "Context Clash is when you accrue new information and tools in your context that conflicts with other information in the context."
- Distraction: from the [Gemini 2.5 technical report](https://storage.googleapis.com/deepmind-media/gemini/gemini_v2_5_report.pdf), "as the context grew significantly beyond 100k tokens, the agent showed a tendency toward favoring repeating actions from its vast history rather than synthesizing novel plans." A [Databricks study](https://www.databricks.com/blog/long-context-rag-performance-llms) "found that model correctness began to fall around 32k for Llama 3.1 405b and earlier for smaller models."
- Confusion: on the [Berkeley Function-Calling Leaderboard](https://gorilla.cs.berkeley.edu/leaderboard.html), "every model performs worse when provided with more than one tool." In a [GeoEngine study](https://arxiv.org/pdf/2411.15399), a quantized Llama 3.1 8b failed with 46 tools but succeeded with 19.
- Clash: in the [Microsoft/Salesforce sharded-prompt paper](https://arxiv.org/pdf/2505.06120), results showed "an average drop of 39%," and "o3's score dropped from 98.1 to 64.1."
- "If you put something in the context the model has to pay attention to it."

**Fits book chapters:** ch01, ch03, ch05, ch09, ch13

---

## An Agentic Case Study: Playing Pokémon with Gemini

- **Date:** 2025-06-17
- **URL:** https://www.dbreunig.com/2025/06/17/an-agentic-case-study-playing-pok%C3%A9mon-with-gemini.html

**Summary.** He reads the appendix of the [Gemini 2.5 technical report](https://storage.googleapis.com/deepmind-media/gemini/gemini_v2_5_report.pdf) as a case study in agent building. The harness had to pull game text from RAM because the model could not read the screen. Two tools were enough to finish the game. Long contexts and hallucinated goals hurt play. This post is the origin of his "context poisoning" and long-context distraction examples, and it argues for building your own evals.

**Citable claims**
- "So much about agents is information control, what gets put in the context."
- Gemini "only needed two tools to complete the game": `pathfinder` and `boulder_puzzle_strategist`.
- Poisoning: a hallucinated goal (the TEA item from the FireRed remake) left the model "fixated on achieving impossible or irrelevant goals"; "The Google team calls this 'context poisoning.'"
- The benchmark and the agent disagreed: Gemini's retrieval performance on long contexts did not carry over to decision making beyond 100k tokens. He calls this an argument for "the need to build your own evals when designing an agent, as the benchmark performances would lead you astray."

**Fits book chapters:** ch01, ch03, ch05, ch13, ch17

---

## Let the Model Write the Prompt

- **Date:** 2025-06-10
- **URL:** https://www.dbreunig.com/2025/06/10/let-the-model-write-the-prompt.html

**Summary.** This is his Data + AI Summit talk on using DSPy for LLM tasks inside applications and pipelines, with Overture place conflation as the example. Prompts written in code grow, become opaque, and break across models. DSPy separates the task (signatures and modules) from the prompt and the model, and optimizers such as MIPROv2 write the prompt from eval data.

**Citable claims**
- His markup of OpenAI's GPT-4.1 SWE-bench prompt: "Only 1% of the prompt defines the job to be done, the task. 19% is chain-of-thought instructions. 32% is formatting instructions."
- Results: Qwen 3 0.6b "scored 60.7% against our eval set. After optimization, our program scored 82%," using "only 14 lines of code, which manages a ~700 token prompt." Re-optimizing for other models reached "Llama 3.2 1B (achieving 91% performance) and Phi-4-Mini 3.8B (achieving 95% performance)."
- Labeling: "the whole exercise took about an hour, including labeling over 1,000 pairs."
- "Decouple your task from the LLM. Write tasks, not prompts." He also writes: "Don't program your prompt. Program your program."

**Fits book chapters:** ch04, ch07, ch17

---

## Your Eval is More Important Than the Model

- **Date:** 2025-01-08
- **URL:** https://www.dbreunig.com/2025/01/08/evaluating-llms-as-knowledge-banks.html

**Summary.** He argues that a custom eval is the most valuable AI asset a team owns. He walks through building one, J1k (1,000 Jeopardy clues), with string normalization in front of an LLM judge, to pick a small on-device model. He then uses the eval to measure prompting changes, which shows both the accuracy gains and their latency cost.

**Citable claims**
- "Your eval is the most valuable AI asset you own, not your model or your prompts."
- With the compound pipeline, "for some models, 60% of responses never hit an LLM" judge.
- Scaling Qwen 2.5 from 0.5b to 1.5b increased performance "by 65%," but going from 14b to 32b "more than doubles our parameters while only yielding a 6% accuracy gain."
- "Chain-of-thought prompting yields a 14% accuracy boost," and DSPy optimization gave "a full 20% performance boost." The cost: "This simple test took 10x longer than our initial technique."
- "Each new model can be evaluated in hours, not weeks."

**Fits book chapters:** ch04, ch14, ch17

---

## Index by chapter

| Chapter | Posts |
|---|---|
| ch01 Every failure is a context failure | The State of Context Engineering; Context Inertia; The Problem is Prompt Debt; Don't Fight the Weights; Why "Context Engineering" Matters; Cat Facts Cause Context Confusion; Prompts vs. Context; How Long Contexts Fail; An Agentic Case Study: Playing Pokémon with Gemini |
| ch02 AI is a systems problem | The State of Context Engineering; Harnesses are Situated Agents; The Problem is Prompt Debt; The Cost of Overfitting the Harness; How System Prompts Define Agent Behavior; Enterprise Agents Have a Reliability Problem; Why "Context Engineering" Matters; Prompts vs. Context |
| ch03 Attention, context windows, compaction | The State of Context Engineering; Context Inertia; How Claude Code Builds a System Prompt; The Potential of RLMs; Cat Facts Cause Context Confusion; How to Fix Your Context; How Long Contexts Fail; An Agentic Case Study: Playing Pokémon with Gemini |
| ch04 In-context learning and pragmatics | What We Can Learn from Claude's Fable 5.1 System Prompt; The Problem is Prompt Debt; What Do Humans Need From Docs?; How System Prompts Define Agent Behavior; Don't Fight the Weights; Cat Facts Cause Context Confusion; Let the Model Write the Prompt; Your Eval is More Important Than the Model |
| ch05 Tool use | The State of Context Engineering; Manage Your Agent's Loadout with Dr. Skill; What Do Humans Need From Docs?; How Claude Code Builds a System Prompt; Don't Fight the Weights; How to Fix Your Context; How Long Contexts Fail; An Agentic Case Study: Playing Pokémon with Gemini |
| ch06 Memory | The State of Context Engineering; Context Inertia; Harnesses are Situated Agents; How Claude Code Builds a System Prompt; The Potential of RLMs; How to Fix Your Context |
| ch07 Context assembly | The State of Context Engineering; What We Can Learn from Claude's Fable 5.1 System Prompt; Harnesses are Situated Agents; Manage Your Agent's Loadout with Dr. Skill; What Do Humans Need From Docs?; How Claude Code Builds a System Prompt; How System Prompts Define Agent Behavior; How to Fix Your Context; Let the Model Write the Prompt |
| ch09 Retrieval | The Potential of RLMs; Don't Fight the Weights; How to Fix Your Context; How Long Contexts Fail |
| ch13 Agents as workflows/harnesses | The State of Context Engineering; Context Inertia; Harnesses are Situated Agents; The Cost of Overfitting the Harness; How Claude Code Builds a System Prompt; How System Prompts Define Agent Behavior; The Potential of RLMs; Enterprise Agents Have a Reliability Problem; How to Fix Your Context; How Long Contexts Fail; An Agentic Case Study: Playing Pokémon with Gemini |
| ch14 Cost of context | The Problem is Prompt Debt; How Claude Code Builds a System Prompt; How to Fix Your Context; Your Eval is More Important Than the Model |
| ch16 Observability | Manage Your Agent's Loadout with Dr. Skill |
| ch17 Evaluation | What We Can Learn from Claude's Fable 5.1 System Prompt; The Problem is Prompt Debt; Enterprise Agents Have a Reliability Problem; An Agentic Case Study: Playing Pokémon with Gemini; Let the Model Write the Prompt; Your Eval is More Important Than the Model |
