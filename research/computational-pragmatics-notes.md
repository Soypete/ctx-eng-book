# Computational Pragmatics — Research Notes

*Dan Jurafsky*
*Source: https://web.stanford.edu/~jurafsky/prag.pdf*
*Reading Context: Context Engineering Book Research*

---

# Paper Overview

This is a chapter from the *Handbook of Pragmatics* (Oxford: Blackwell).

**Definition**: Computational pragmatics is the computational study of the relation between utterances and context.

**Four core inferential problems**:
1. Reference resolution
2. Speech act interpretation
3. Discourse structure and coherence
4. Abduction

**Key quote from the paper**:
> "Computational pragmatics, like pragmatics in general, is especially concerned with INFERENCE."

---

# Core Insight: Indexicality + Inference

## Indexicality

Pragmatics is concerned with **indexicality**: the relationship between utterances and context.

**Indexicals** are words/phrases whose meaning depends on context:
- "I" → depends on who is speaking
- "you" → depends on who is listening
- "here" → depends on location
- "now" → depends on time
- "this" → depends on what is being pointed to

This is **crucial for database theory**:

- Databases store facts
- Facts need context to be meaningful
- The same data means different things in different contexts
- Indexicality explains *why* context matters to data

**Connection to database theory**:
- Relational databases assume facts are self-contained
- But real-world data is indexical: "user's data" means different things for different users
- This is why authorization, tenancy, row-level security exist
- Long-term: databases need to handle indexicality as a first-class concept

## Inference

The paper explicitly states: pragmatics is "especially concerned with INFERENCE."

**Context engineering is optimizing for inference**:

```text
Without Context Engineering:
    Utterance → Model → Guess → Response
    (lots of inference required, high error)

With Context Engineering:
    Utterance + Context → Model → Extraction → Response
    (inference already done, just extraction)
```

The paper's inference problems:
1. **Reference resolution**: What does "it" refer to?
2. **Speech act interpretation**: What did the user mean?
3. **Discourse structure**: How does this relate to what was said?
4. **Abduction**: What explains this utterance?

Context engineering solves these *before* the model sees the input:
- Provide explicit references
- Provide intent context
- Provide discourse state
- Provide relevant facts

---

# Surface Form vs Pragmatic Force

The paper discusses the **Literal Force Hypothesis**:

> "Every utterance has an illocutionary force which is built into its surface form."

Examples from the paper:
- **Aux-inverted sentences** (Can you X?) → QUESTION force
- **Subject-deleted sentences** (Pass the salt) → IMPERATIVE force

This means:

> **Changing how context is presented will influence pragmatics.**

If you present the same information in different surface forms, the model may interpret it differently.

---

# Research Question: Do LLMs Understand Pragmatics?

The paper was written in 2003, before modern LLMs. The question now:

> Can LLMs, which should have a near-complete grasp of written language, interpret pragmatic cues implicitly?

## Hypothesis

LLMs trained on massive text corpora should have internalized:

- Question intonation patterns (even in text: "Can you..." vs "You can...")
- Imperative structures (subject deletion, verb-first)
- Indirect speech act patterns ("Could you..." = request, not question)
- Discourse conventions

## If Yes

Then we can rely on the model's pragmatic understanding and focus context engineering on:
- Providing relevant facts
- Providing state
- Providing constraints

## If No (or Partially)

Then context engineering must also:
- Explicitly label speech acts ("This is a request for X")
- Provide intent context ("User wants to book a flight")
- Disambiguate ambiguous utterances

## Evidence to Gather

- Does "Can you X?" consistently trigger helpful responses vs "Please X."?
- Does the model correctly identify indirect requests?
- Does changing surface form change model behavior in systematic ways?

---

# Explicit vs Implied Directives

Important distinction for agent instructions:

## Explicit Directives

The instruction clearly states what to do:
- "Book a flight to Boston on Tuesday"
- "Send an email to John"
- "Summarize this document"

The speech act is clear: **DIRECTIVE / REQUEST**

## Implied Directives

The instruction implies an action without stating it directly:
- "I need to get to Boston Tuesday" → (implied: book flight)
- "John should know about this" → (implied: send email)
- "This is too long to read" → (implied: summarize)

The speech act is often a **STATEMENT** that implies a request.

---

# Process Knowledge: How vs What

When instructing agents, we need to be explicit about:

## Directive Type 1: "Do X"

Provide the goal, let the agent determine process:
- "Book a flight" → agent figures out how
- "Find information about X" → agent determines search strategy

## Directive Type 2: "Do X using Y"

Provide both goal and method:
- "Book a flight using the travel API"
- "Find info in the database, not the web"

## Directive Type 3: "Here is information, decide what to do"

Provide context, agent decides on action:
- "Here is the user's preferences, recommend something"
- "Here is the conversation history, respond appropriately"

---

# Why This Matters for Context Engineering

When providing reasoning information to agents, we must be explicit about:

1. **What type of act is this?**
   - Is this a directive to take action?
   - Is this information to consider?
   - Is this a question to answer?

2. **What is the process knowledge?**
   - Does the user want the agent to figure out how?
   - Or is the user telling the agent exactly what to do?
   - Or is the user providing context for the agent to decide?

3. **What is the expected output?**
   - An action?
   - A response?
   - A recommendation?

Without this explicit labeling, the model must infer:
- What the user wants
- How to accomplish it
- What form the response should take

This is the same inference problem the BDI model solves — but we can solve it *in context* rather than making the model infer.

---

# Check Questions Should Never Go to LLMs

From the paper, **check questions** are a specific type of speech act:

> A check question requests the interlocutor to confirm something that the other participant has privileged knowledge about.

Example from the paper:
```
A: I wanted to travel next week.
B: And you said you wanted to travel next week?
```
(B is asking A to confirm something A already knows)

## Why This Matters for LLMs

**Check questions are inappropriate for LLMs because:**

1. **LLMs don't have privileged knowledge** — They don't know what the user knows vs. doesn't know
2. **Check questions waste tokens** — "Does that sound right?" "Is this what you wanted?" — just extra tokens
3. **Check questions are not explicit enough** — They rely on the model inferring what to confirm
4. **The LLM will just agree** — There's no actual confirmation mechanism

## Leading Questions (Mentorship Pattern)

Engineers often use leading questions as a mentorship technique:
- "What do you think should happen here?"
- "Have you considered X?"
- "Does this approach make sense?"

These are great for humans learning to reason.

**But for LLMs:**
- They don't need to learn reasoning — they already have the capability
- Leading questions just add tokens without adding information
- The LLM will just validate whatever context it's given

## The Rule

> **Never use check questions or leading questions with LLMs.**

Instead:
- Be explicit: "Confirm this action: book flight to Boston"
- Provide information directly: "User wants to fly to Boston on Tuesday"
- Don't ask for confirmation the model can't actually provide

---

# Freedom of Interpretation vs Freedom to Act

Important distinction:

## Freedom of Interpretation

When you ask a check/leading question like:
- "Does this look right?"
- "What do you think should happen?"

The model must interpret:
- What "this" refers to
- What "right" means in context
- What criteria to use for evaluation

This is **interpretive freedom** — the model decides what the question means.

## Freedom to Act

When you ask:
- "Should I proceed?"
- "Do you want to continue?"

The model has **action freedom** — it decides whether to continue, stop, ask clarifying questions, etc.

Both are problematic:

| Question Type | Problem |
|---------------|---------|
| "Does this look right?" | Model will say yes regardless |
| "What do you think?" | Unbounded interpretation |
| "Should I proceed?" | Model may or may not — inconsistent |
| "Have you considered X?" | Model may ignore or over-interpret |

## The Solution: Constrain Both

Context engineering should constrain:
1. **Interpretation** — Explicitly state what the information is
2. **Action** — Explicitly state what action to take

Instead of: "What do you think we should do here?"
Say: "Based on the user's request, recommend the flight booking action."

Instead of: "Does this look correct?"
Say: "Verify: User wants flight to Boston on Tuesday. Confirm if correct by proceeding."

---

# Human-in-the-Loop vs Autonomous Systems

## The Phone Agent Example

Human agent on phone: "I want to travel"
→ Agent asks: "Where do you want to go? When? What kind of flight?"

This works because:
- Human can ask clarifying questions
- Human can infer what information is missing
- Human can engage in clarification dialogue

## When Human-in-the-Loop Works

Speech act interpretation + clarification questions is FINE when:
- There's a human in the loop
- The system can say "I don't understand, can you clarify?"
- Cost of clarification is low

## When It Breaks — No Human in the Loop

When the system is intended to be autonomous:

- "I want to travel" → system must act, can't ask questions
- No opportunity for clarification
- Missing information = failed task or wrong output
- Can't recover through dialogue

**This is where context engineering is critical:**

> You must provide complete context upfront.
> You cannot rely on the system to ask clarifying questions.
> Every needed piece of information must be in the context.

## The Implication

If your system is intended to be autonomous (no human in loop):
- Assume no clarification dialogue will happen
- Provide all required information explicitly
- Structure context so the model doesn't need to infer what's missing

This is why the speech act paper's inference-based approach works for humans but may fail for autonomous LLMs — humans can ask back, LLMs often just proceed with incomplete information.

---

# Do We Need BDI for Agent Workflows?

**Answer: Yes.**

The BDI model (Belief-Desire-Intention) provides exactly what agents need to make correct inferences:

## BDI Provides:

| Concept | In Agent Context | In Context Engineering |
|---------|-----------------|----------------------|
| **Belief (B)** | What the system knows | Provided via retrieval, knowledge bases, user data |
| **Desire (W)** | What the user/goal wants | Explicit goals, task definitions |
| **Intention** | What the system plans to do | Action plans, tool selection |

## The Insight

> Injecting "belief" is exactly what the agent needs to make correct inferences.

Without beliefs:
- Agent must infer everything from scratch
- High error rate on complex tasks

With beliefs (context engineering):
- Agent has structured knowledge to reason with
- Lower error, more predictable behavior

## This Connects to Graph RAG and Ontologies

**Personalization through knowledge injection:**

- **Graph RAG**: Provides structured beliefs about entities and relationships
- **Ontologies**: Provide shared vocabulary and semantic constraints
- **User profiles**: Provide beliefs about user preferences
- **Session state**: Provides beliefs about current context

These are all forms of ** Belief Injection**:

```text
Without Context Engineering:
    User Input → Model → Guess → Action

With Context Engineering (Belief Injection):
    User Input + Beliefs (Graph/Ontology/Profile) → Model → Reasoning → Action
```

The BDI model explains *why* this works: the model needs beliefs to make the right inferences. Context engineering provides those beliefs.

---

# How Agent Harnesses Already Do This

The ATIS example from the paper shows exactly how this works:

```
Can you give me a list of the flights from Atlanta to Boston?
    ↓
Interpret: This is a REQUEST for flight information
Infer: User wants to fly from Atlanta to Boston
Check preconditions: Need origin, destination, date
Determine action: Query flight database
```

This is what OpenCode and similar agent harnesses do:

1. **Take user input** ("fix the bug in foo.c")
2. **Explicitly list inferred requests**:
   - "The user wants me to fix a bug in foo.c"
   - "I need to read foo.c to find the bug"
   - "I need to understand what the bug is"
   - "I need to make a fix"
3. **Derive context needs** from each inferred request
4. **Determine tool calls** needed for each step

## Why This Costs Tokens But Is Worth It

Yes, this uses tokens for reasoning. But it provides:

1. **Context retrieval guidance** — Know what to fetch before acting
2. **Tool call planning** — Know what tools to call in what order
3. **Debugging/auditability** — Can follow the reasoning logic
4. **Error recovery** — Can see where reasoning went wrong
5. **Formalized logic** — Like deriving a mathematical proof

## It's Like Formal Methods for AI

Just as formal methods use logic to verify program correctness:

- Explicit state what's being assumed
- Explicit derive what's implied
- Can trace through each step

Agent reasoning does the same:
- Explicitly state what the user wants (belief)
- Explicitly derive implications (inference)
- Can trace through each reasoning step

This is log-based debugging for AI — instead of guessing why it did something, you can see the explicit reasoning chain.

---

# Context as Preconditions for Action

Key insight:

> Context engineering is providing the conditions that fulfill actions in order to follow or fulfill the exact request specified by the logic.

This is exactly like formal logic / planning:

## In Planning Systems

```
ACTION: Book-flight
PRECONDITIONS:
  - origin-city known
  - destination-city known
  - date known
  - user-authorization verified
EFFECTS:
  - flight-booked
```

To execute the action, you must satisfy the preconditions.

## In Context Engineering

The context we provide is what satisfies those preconditions:

| Action Requirement | Context We Provide |
|-------------------|-------------------|
| origin-city | "User is flying from Atlanta" |
| destination-city | "User wants to go to Boston" |
| date | "User wants to travel Tuesday" |
| user-authorization | User has permission to book |

**Context = the data that satisfies the preconditions for action.**

## The Pattern

```
Request: "Book me a flight"
    ↓
Infer required actions & preconditions
    ↓
Provide context that satisfies preconditions
    ↓
Execute action with all needed data
```

This is why context engineering works:
- We don't guess what the model needs
- We identify what the logic requires
- We provide exactly that

---

# BDI + Ontology Connection

## BDI Uses Predicate Calculus

From the paper, BDI is formalized using predicate calculus:
- `B(S, P)` = Agent S believes proposition P
- `W(S, P)` = Agent S wants P to be true
- Axiom schemas define how to reason about beliefs and desires

## Ontologies Are Facts, Not Beliefs

Key distinction:
- BDI "belief" is subjective — what an agent believes
- Ontology "fact" is objective — what is true in the domain

But here's the connection:

## Predicate Substantiation in Ontologies

In ontologies, one predicate can substantiate another:

```
Fact A: Flight-from(flight1, Atlanta)
Fact B: Flight-to(flight1, Boston)
Predicate: Flight-from(X, Y) ∧ Flight-to(X, Z) → Route(X, Y, Z)

Therefore: Route(flight1, Atlanta, Boston)
```

The predicate from Facts A and B substantiate Fact C.

## The Connection to Context Engineering

| BDI Concept | Ontology Concept | Context Engineering |
|-------------|-----------------|---------------------|
| Belief (B) | Fact | Provided data |
| Want (W) | Goal / Requirement | Task definition |
| Predicate relations | Inference rules | How data connects |
| Inference chain | Entailment | Reasoning path |

**The insight:**

> In ontologies, we don't have "beliefs" — we have facts. But the predicate relationships from one fact can substantiate another fact.

This is exactly what context engineering does:
- Provide base facts (context)
- Let predicates define relationships
- The model derives substantiated conclusions

This is more rigorous than BDI because it's grounded in objective facts, not subjective beliefs.

---

# Knowledge = True Belief (Linguistic Precedent)

From the paper:

> "Knowledge is defined as 'true belief'; S knows that P will be represented as KNOW(S, P), defined as follows: KNOW(S, P) ≡ P ∧ B(S, P)"

This is the classic philosophical definition: **knowledge = justified true belief**.

## The Connection for Agents

If S is an agent, then:

- Ontology = agent's beliefs
- Facts in ontology = things the agent believes are true
- Knowledge = the agent believes it AND it's actually true

```
KNOW(Agent, P) ≡ P ∧ B(Agent, P)
```

For an agentic system:
- B(Agent, P) = The context/ontology provided
- P = The actual fact
- KNOW(Agent, P) = Agent can act as if it knows P

## The Linguistic Precedent

> If belief → knowledge in the classical sense, and ontologies create knowledge, then ontologies are what agents use to "know" things.

This is a powerful framing:
- Building ontologies = building knowledge
- Agent's context = agent's beliefs
- Correct context → knowledge → reliable action

## Research Questions for Follow-up

- [ ] Find more evidence on knowledge = true belief in AI/agent contexts
- [ ] Explore Chris's work on LLMs in production (need to follow up)
- [ ] How do we ensure P is actually true (not just believed)?
- [ ] Can ontologies be self-verifying?
- [ ] Why does the pseudo-math (BDI predicate calculus) work so well for formalizing context engineering?
- [ ] Is there existing work on using BDI + ontologies together?

---

# 1. The Speech Act Problem

**Key insight**: Surface form ≠ actual speech act.

Example: "Can you pass the salt?"
- **Surface**: Yes/no question about ability
- **Actual**: Polite request (indirect request)

This is called **Indirect Speech Acts**.

---

# 2. Two Computational Approaches

## A. Plan Inference (BDI) Model

Logic-based approach. Key concepts:

- **Belief (B)**: What the speaker knows
- **Desire (W)**: What the speaker wants
- **Intention**: What the speaker plans to do

The inference chain for "Can you give me a list of flights?":
1. Speaker asked about ability → wants to know
2. People don't ask about things they don't want
3. "Can do" is precondition for "do"
4. Therefore: speaker wants me to do the action
5. That's a REQUEST

**For context engineering**: This model explains *why* users make requests the way they do.

The W (want) in BDI: If A is an action, we infer S wants A to be done.

---

# PKO and the Action Store Problem

## The Insight

Looking at PKO (permission-based knowledge stores) - they try to store all possible actions as a finite list.

**Problem**: This works for permissioning, not LLMs.

- Permission systems: Predefined allowed actions → check if user has permission → done
- LLMs: Need infinite possible actions → system provides constraints → LLM decides

## Why Finite Action Lists Fail for LLMs

```
PKO-style:    Actions = {action1, action2, ..., actionN}
              Agent picks from finite set

LLM-style:    Constraints = {constraint1, constraint2, ...}
              Agent generates action meeting constraints
              Action space = infinite
```

The BDI model supports this:
- W(S, A) = "S wants action A to be done"
- The Desire doesn't specify HOW to do it
- Only the context provides the means

## The Pseudo-Math

The paper uses formal notation like:

```
B(S, P)    = S believes P
W(S, P)    = S wants P
Int(S, A)  = S intends to do A
```

Then inference rules like:
- If B(S, can_do(A)) ∧ W(S, A) → Int(S, A)

This pseudo-math is actually powerful - it's formalizing what context engineering does:
1. Provide beliefs (context)
2. Infer wants (from user requests)
3. Let the agent determine actions (infinitely many)

---

## B. Cue-Based (Probabilistic) Model

Statistical classification approach. Uses surface cues:

- **Lexical**: Specific words/phrases
- **Syntactic**: Sentence structure
- **Prosodic**: Intonation, stress (for speech)
- **Discourse**: Context within conversation
- **Microgrammar**: Patterns specific to spoken language

Treats speech act recognition as classification: given utterance + cues → predict act type.

---

## Microgrammar: Spoken vs Written Language

### What is Microgrammar?

The grammatical structures specific to spoken language:
- Fragments ("Yeah", "So", "Okay then...")
- Turn-taking markers
- Backchannels
- Hesitations
- Repair sequences

### Current LLM Practice

Most LLM interaction uses spoken language patterns (microgrammar):
- Conversational
- Roundabout
- Implicit
- "Let me think...", "So what I want is...", "Can you..."

### The Efficiency Problem

```
Spoken/Microgrammar: "So, like, can you like, run the tests? I mean, if it's not too much trouble?"
  ↓
Inference needed → Tokens consumed → Ambiguity → Possible error

Structured/Semantic: "Run tests"
  ↓
Direct → Few tokens → No ambiguity → Correct
```

### The Insight

We use spoken language patterns because:
1. We're used to talking to AI like it's human
2. It feels natural
3. Current prompting guides us to be "conversational"

But for autonomous agents:
- **Efficiency matters**
- **Ambiguity costs**
- **Direct is better**

### Best Practice?

Spoken patterns for HIL (human-in-the-loop):
- Human can resolve ambiguity
- Natural conversation feels good
- Efficiency less critical

Structured patterns for autonomous:
- Direct = fewer tokens
- Explicit semantics = no inference needed
- No human to resolve ambiguity

### The User's Observation

> "When I talk it is always more roundabout than direct."

This is natural for humans. But AI systems (especially autonomous) should receive structured input.

The semantic web theory says: **provide data with semantics, the AI can use it**. This implies we should transform human speech → semantic structure before processing.

```
Human speech: "Can you check if the tests pass? I need to know before I push."
  ↓
Semantic parse: Intent=CHECK_TESTS, Purpose=PRE_PUSH_VERIFICATION
  ↓
Execute with context
```

---

# Speech Acts: Can STT/TTS Capture Emphasis?

## The Question

Do models like Whisper (speech-to-text) or Qwen TTS have a way to capture:
- **Emphasis**: Which word is stressed
- **Semantic force**: How hard the speaker pushes on meaning
- **Prosodic features**: Intonation, stress, volume

These are key to speech act recognition in humans.

## What Current Models Do

### Whisper (STT)
- Transcribes text
- Optionally includes timestamps
- Generally does NOT capture emphasis/stress
- Output: "Can you check the tests?" (same as "Can you CHECK the tests?")

### TTS Models
- Generate speech from text
- Can add prosody markers (in some models)
- But input is text - loses the emphasis from source

## The Gap

Current speech pipeline:
```
Audio → Whisper → Text → LLM → Response
              ↑
         Loses emphasis
```

What's missing:
- **Which word was stressed?**
- **What was the intonation contour?**
- **Semantic force** (how strongly was it said?)

## Research Question

> Can we build a pipeline that preserves speech act features?

Approaches:
1. **Prosody-aware STT**: Models that output phonemes + stress markers
2. **Embedding-based**: Audio embeddings that capture prosody, pass to LLM
3. **Text + metadata**: Whisper output + confidence scores, pauses, etc.

## Connection to Context Engineering

If we can capture emphasis:
- "Check tests?" (rising = question) vs "CHECK TESTS!" (stressed = urgent request)
- Different speech acts from same words
- More accurate intent recognition

This is a frontier - current systems lose the prosodic information that humans use to understand speech acts.

---

# Temporal Relevance: Looking Forward + Backward

## From the Paper

Utterances have:
- **Looking backward**: References to what was just said
- **Looking forward**: Sets up what comes next

This is the temporal nature of conversation.

## Implication for Conversational AI

Conversation isn't just context - it's **temporally ordered**:
- "It" refers to something earlier
- "Then" sets up something later
- Each utterance depends on history AND sets up future

## The Caching Problem

Current approach: Keep N messages in context
- Context window limits → truncate at ~25 messages
- Lose temporal relevance

**But**: Not all history is equally relevant.

## The User's Observation

> "We can probably compact history at 25 messages, but there is temporal relevance to things just said."

Recent messages are more relevant than old ones:
- Last 1-3 messages: Highly relevant (looking backward)
- Last 4-10 messages: Moderately relevant (global context)
- 10-25 messages: Compressed/structural only

## Proposed Approach: Tiered Context

```
Tier 1 (immediate):  Last 3 messages - verbatim
Tier 2 (recent):     Messages 4-10 - summarized
Tier 3 (history):    Messages 10-25 - structural only
Tier 4 (beyond):     Discard or external memory
```

## Why This Works

- **Looking backward**: Most recent messages satisfy this
- **Looking forward**: Current message sets up next
- **Compaction**: Older messages can be compressed because they're reference, not immediate

## Connection to Other Points

- Action schemas need recent context (preconditions just established)
- Speech acts depend on immediate prior utterance
- Permissions often based on recent request

This is another argument for **structured context** over raw conversation history.

---

# Temporal Relevance of Context Provision

## The Insight

> "I don't like to give it all upfront. I like to give the data when the action is needed to take place and only then."

Context isn't just *what* - it's *when*.

## Two Approaches

### Eager Context Loading (Current Common)
```
System: Here is ALL context upfront
  ↓
User: [interaction happens]
  ↓
System: Uses initial context for everything
```

### Lazy Context Loading (Preferred)
```
User: [makes request]
  ↓
System: Determines needed action
  ↓
System: Hydrates context ONLY for that action
  ↓
Execute with specific context
  ↓
Context discarded after execution
```

## Why Lazy Loading Wins

| Aspect | Eager | Lazy |
|--------|-------|------|
| Token usage | High (all context always) | Low (only needed) |
| Staleness | Context becomes outdated | Fresh context per action |
| Scope | Over-scope (too much data) | Right-scope (just enough) |
| Relevance | General | Specific to action |

## The Action-Schema Connection

This is exactly what action schemas enable:

```
Action: send_email
  Preconditions: recipient_address, email_content
  → Fetch these ONLY when send_email is selected
  → Not upfront
```

The context is hydrated at execution time, not at conversation start.

## Research Question

Does lazy context loading improve:
- Token efficiency?
- Accuracy (less irrelevant context)?
- Staleness (always fresh)?

This would support the thesis that **context engineering is about when, not just what**.

---

# DataMesh: Lazy Loading at Scale

## The Connection

From Miriah's datamesh work (distributed systems substack):

**DataMesh principle**: Data is available at source, not copied centrally. Fetch on demand.

**Context engineering**: Context is available at action time, not loaded upfront.

These are the same pattern.

## DataMesh for Context Lazy Loading

```
Traditional:                    DataMesh/Lazy:
─────────────────               ─────────────────
Central data store             Distributed data sources
Copy data to context           Fetch when needed
Load all upfront               Load on-demand
Stale copies                   Fresh from source
Single point of failure        Resilient, distributed
```

## How It Works

```
User Request: "Send report to John"
  ↓
Determine Action: send_email(to=john, doc=report)
  ↓
Hydrate Preconditions:
  - Query user service: John's email
  - Query doc service: Report content
  - Query permission service: Can send?
  ↓
Execute with fresh context
  ↓
Context discarded
```

Each precondition is a **data product** (datamesh term) fetched on demand.

## Why This Works

- **Fresh**: Data from source, not stale copy
- **Efficient**: Only fetch what's needed
- **Resilient**: No single data warehouse to fail
- **Scalable**: Distributed sources handle load

## Implication

Context engineering + DataMesh = 
**The context IS the data mesh, fetched lazily at execution time**

This is the unified architecture:
- DataMesh: Where data lives
- Action schemas: What data needed
- Lazy loading: When to fetch
- Permission gating: What you're allowed to fetch

---

# Research Question: Exclamatives + Adverbs in Prompting

## The Question

What role do exclamatives and adverbs play in prompting?

- **Exclamatives**: "Wow!", "Amazing!", "Seriously!", "Can't believe..."
- **Adverbs**: "carefully", "quickly", "really", "very", "exactly"

## What They Might Do

### Exclamatives
- Signal urgency or emphasis
- Change model attention/focus
- "Do this NOW!" vs "Do this"
- Could trigger different response style

### Adverbs (Manner)
- "Think step by step" → changes reasoning
- "Be very careful" → increases caution
- "Exactly output JSON" → increases precision

## Research Hypotheses

1. **Attention shifting**: Exclamatives focus model on specific content
2. **Constraint marking**: Adverbs like "carefully" signal need for precision
3. **Style switching**: "IMPORTANT:" or "NOTE:" changes formality
4. **Emotional contagion**: Excited prompt → more enthusiastic output

## Current Practice

Common prompts use these:
- "IMPORTANT:", "NOTE:", "WARNING:" (exclamatory markers)
- "Always", "Never", "Be careful" (adverbial constraints)
- "Think step by step" (manner adverb)
- "Very important" (intensifier)

## Research Needed

- Do these actually change model behavior, or are they just ritual?
- Which markers work across models?
- Is there a taxonomy of effective prompt markers?

## Connection to Speech Acts

Exclamatives in speech acts signal:
- Emotional force
- Urgency
- Emphasis on proposition

This is the "semantic force" dimension - similar to prosody, but in text.

---

# Prosody in Text: ?! vs . vs ! (Internet Age)

## The Observation

In written internet communication, we use punctuation to simulate prosody:

| Written | Prosody Equivalent |
|---------|-------------------|
| "no?" | Rising intonation, questioning |
| "no." | Flat, neutral |
| "no!" | Stressed, emphatic |
| "no?!" | Mixed, uncertain emphasis |
| "NO!!!" | Strong emphasis, urgent |

This is how text simulates the prosodic features that spoken language has.

## Do LLMs Handle This?

**Hypothesis**: LLMs trained on internet text may have learned these patterns:
- "?" at end = question
- "!" = emphasis/urgency
- "?!" = mixed/confused

Evidence:
- Models respond differently to "do this." vs "do this!"
- Emotional text often gets emotional responses

## For HIL vs Autonomous

### Human-in-the-Loop
- These cues matter: user typing "run the tests?!" signals something
- Human can interpret the emphasis
- Context includes emotional state

### Autonomous Workflows
- **Probably doesn't matter**
- Structured input: "run tests" → no ambiguity
- No human to interpret emphasis
- Preconditions/permissions handled separately

## The Connection

This is why autonomous systems need structured context (semantic web):
- Text prosody is ambiguous even to humans
- Autonomous can't rely on "feeling" the emphasis
- Provide explicit semantics instead

```
Ambiguous:   "Run tests?!"
Structured:  Intent=RUN_TESTS, Priority=URGENT
```

For autonomous: explicit priority markers > punctuation emphasis

---

# Phonetic vs Phonological (Written Context)

## The Question

"Phonetic vs phonological cues... this means written, right?"

## Answer: No - But Also Yes

### Phonetic
- Physical sounds of speech
- How words actually sound when spoken
- Used in speech recognition, STT

### Phonological  
- Sound system of a language
- Patterns, rules, contrasts between sounds

### In Written Text (Prompts)

For LLM prompts, we're dealing with:
- **Orthographic**: Written symbols (letters, words)
- **NOT phonetic**: We don't have audio

But LLMs trained on text may have learned:
- **Phonological patterns**: From written representations of speech (dialogues, transcripts)
- **Prosodic markers**: "?!", "!!!", CAPS for emphasis

## The Implication

When the paper talks about phonetic cues:
- In speech: Actual audio features
- In text prompts: Proxy through punctuation, caps, etc.

For HIL voice interfaces: Phonetic cues come from audio
For text-based AI: Phonological patterns are learned from text data

This is why we can use text markers like "IMPORTANT!" or "???" - the model learned these as proxies for prosody from internet text.

---

# Hidden Markov Model (HMM) - Not RL

## Clarification

**HMM is NOT Reinforcement Learning.**

It's a **statistical sequence model**:
- Observable states
- Hidden states
- Transition probabilities between states

Used for:
- Speech recognition (audio → text)
- Part-of-speech tagging
- Sequence labeling

## In the Paper / Linguistics Context

The paper relates HMM to linguistic modeling because:
1. **Speech is a sequence**: Phonemes, words, utterances
2. **Hidden states**: Grammatical structure, intent, context
3. **Observable**: What was actually said

```
HMM for Speech Acts:
- Observable: Utterance sequence
- Hidden: Speech act type (request, question, statement)
- Transition: How acts follow each other in conversation
```

## Why It's Relevant to AI

HMMs were a precursor to:
- **CRFs** (Conditional Random Fields) - sequence labeling
- **Neural sequence models** - LSTMs, Transformers

The insight: Speech acts form a Markov chain - what you say next depends on what was said.

## Connection to Context Engineering

- Conversation is a sequence of speech acts
- HMM models this as: P(act_t | act_1, ..., act_{t-1})
- But we're replacing probabilistic inference with explicit context

Modern approach:
- Old: HMM → infer speech act from sequence
- New: Structured context → speech act is explicit

---

# "No" at End of Statement (Spanish → Question)

## From the User

In Spanish:
- Statement: "Tu quieres café" (You want coffee)
- Question: "Tu quieres café, no?" (You want coffee, no?)

Adding "no?" at the end turns a statement into a question.

## As a Bilingual Speaker

This is a real indirect speech act pattern:
- **Surface**: Statement + "no?"
- **Intent**: Question/request for confirmation

## English Parallels

English uses similar patterns:
- "You're coming, right?" (= question)
- "That's fine, no?" (= question, especially in certain dialects)
- "I guess so, huh?" (= question)

## Implications for AI

These patterns are:
1. **Learned from training data** (bilingual text, code, informal writing)
2. **Cross-linguistic** - exist in multiple languages
3. **Indirect** - surface statement, actual question

For context engineering:
- Structured input should avoid these ambiguities
- But the model likely handles them from training

---

# 3. Key Terms

## Indirect Requests

"Something that looks like a question but functions as a request."
- "Can you X?" → Request
- "Could you X?" → Request
- "Would you mind X?" → Request

## Check Questions

A subtype of question that requests confirmation of information the speaker believes the hearer has privileged knowledge about.

Example:
```
A: I wanted to travel next week.
B: And you said you wanted to travel next week?
```
(B is checking, not really asking)

## Microgrammar

Specific lexical, collocation, and prosodic features characteristic of particular conversational moves (Goodwin, 1996).

Example: Assessments often follow pattern:
```
Pro Term + Copula + (Intensifier) + Assessment Adjective
"That's really great"
```

---

# 4. Dialogue Act Tagsets

The paper discusses several computational dialogue act tag sets:

### DAMSL (Dialogue Act Markup in Several Layers)

Forward-looking functions:
- STATEMENT
- INFO-REQUEST
- CHECK
- INFLUENCE-ON-ADDRESSEE (directives)
- OPEN-OPTION (suggestions)
- ACTION-DIRECTIVE (commands)
- OFFER
- COMMIT

### Switchboard (42 categories)

Includes: STATEMENT, BACKCHANNEL, OPINION, AGREEMENT, YES-NO-QUESTION, WH-QUESTION, etc.

---

# 5. Connection to Context Engineering

## Why This Matters for AI Systems

1. **Tool Calling**: When a user says "Can you check my calendar?" they are making a request, not asking a question. The system must infer intent.

2. **Meaning vs Form**: The same words can mean different things in different contexts. "Can you help me?" vs "Can you help me [with this task]?"

3. **Inference Chains**: Understanding user intent requires reasoning about:
   - What the user believes
   - What the user wants
   - What actions would satisfy that want
   - What preconditions exist

4. **Context Dependencies**: Same utterance means different things:

---

# 6. Action Schemas (Under-Discussed in Agents)

Action schemas from classical planning define:

```
Action(A):
  Preconditions: P1 ∧ P2 ∧ ...
  Effects: E1 ∧ E2 ∧ ...
  Body: [steps to execute]
```

## The Context Hydration Pattern

When an LLM needs to perform an action, it must hydrate:
- **Preconditions**: What must be true before action?
- **Effects**: What will be true after?
- **Body**: How to execute

This tells us exactly what data to provide as context.

## Four Types of Context

| Type | Purpose | Example |
|------|---------|---------|
| **Belief Context** | Define what is true | Knowledge graph, ontology facts |
| **Action Determination Context** | Decide what to do | User intent, available tools, constraints |
| **Action Execution Context** | Hydrate preconditions for specific action | Specific data needed to perform action |
| **Permission Context** | Constrain allowed actions | Role permissions, rate limits |

## The Security Implication

Action execution context often contains proprietary data.

Options for protection:
1. **Step-based workflows**: Break action into steps, hydrate per-step
2. **Tool call permissions**: Check permissions before providing data
3. **Deterministic prompt injection**: Use structured prompts that extract only needed data
4. **Local execution**: Run action locally, return only results

```
User: "Send the Q3 report to John"

LLM determines: Action = send_email(recipient, document)
Hydrates preconditions: John's email, Q3 report content
  → These are EXECUTION CONTEXT
  → Can be permissioned, filtered, redacted

Result: Only sends what was explicitly authorized
```

## Why Action Schemas Aren't Discussed

Probably because:
- People don't know classical planning theory
- It's "obvious" once stated
- Most agent tutorials focus on "what tool to call" not "what data to hydrate"

But this is exactly what separates working agents from toy agents:
- Toy: "Here's a list of tools, figure it out"
- Working: "Here are action schemas with preconditions, hydrate as needed"

---
   - After different prior utterances
   - In different relationships (expert vs peer)
   - At different points in a conversation

## The BDI Model for Agents

The BDI (Belief-Desire-Intention) model provides a framework for:

```text
User Utterance
    ↓
Detect Speech Act (Request, Question, etc.)
    ↓
Infer User Goal/Desire
    ↓
Check Preconditions
    ↓
Select Action to Satisfy Goal
    ↓
Execute / Respond
```

This is essentially what modern agent tool-calling does, but the paper provides the *theoretical foundation* for why this inference is necessary.

---

# 6. Observations for Context Engineering

## Speech Acts Are Inferential

The core problem: The literal meaning of an utterance is often not what the speaker intends.

> Context engineering must account for this gap between what's said and what's meant.

## Context Provides Constraints

The paper shows how discourse context affects interpretation:

- "No it isn't" = AGREEMENT after "It isn't raining"
- "No it isn't" = DISAGREEMENT after "It is raining"

Same words, different acts, based entirely on context.

## Cue Combinations Matter

No single cue determines speech act. It's the *combination* of:
- Words used
- Syntactic structure
- Prosody (in speech)
- Conversation position
- Prior utterances

This is why simple keyword matching fails for intent detection.

---

# 7. Research Questions

### RQ-CP-001

How can speech act theory improve tool-calling accuracy in agent systems?

### RQ-CP-002

Can BDI-style inference be implemented through prompt engineering, or does it require architectural changes?

### RQ-CP-003

How do we handle multi-layered speech acts (a request disguised as a question disguised as a statement)?

### RQ-CP-004

What is the minimum context needed to correctly interpret user intent in agent systems?

---

# Potential Book Quote

> The problem with natural language is that what people say and what people mean are often two different things.
>
> A system that only processes what is said—without inferring what is meant—will fail at the very task of understanding.
>
> Context engineering must bridge this gap.

---

# Reference: Perrault & Allen on Indirect Speech Acts

## The Three Relevant Speech Acts for Indirect Requests

Based on Searle's taxonomy (referenced in Perrault & Allen's work):

| Speech Act | Literal Form | Intent | Example |
|------------|--------------|--------|---------|
| **Inform** | Statement | Transfer information | "It's cold in here" (literal: inform) |
| **Request** | Question/Statement | Get action | "Can you close the door?" (literal: question about ability) |
| **?question?** | - | - | - |

## The Inference Chain

Perrault & Allen formalized how to infer the intended speech act from surface form:

```
Surface: "Can you X?"
  ↓
Literal: Question about ability (Inform)
  ↓
Inference: If S asks about ability to do X,
           and S wants X done,
           then S is making a REQUEST
  ↓
Intent: Request to do X
```

## The Three Acts (from Searle's Taxonomy, used by Perrault & Allen)

1. **Assertives (Inform)** - Committing speaker to truth of proposition
2. **Directives** - Attempting to get hearer to do something
3. **Commissives** - Committing speaker to future action

For **indirect requests**, the system must:
1. Recognize surface form (often a question or statement)
2. Apply inference rules
3. Map to intended directive (request)

---

## MCP + Human-in-the-Loop for Indirect Requests

For coding agents, Claude Desktop, etc.:

```
User says: "Can you check if the tests pass?"
  ↓
MCP: Recognizes indirect request pattern
  ↓
Tool: Run test command
  ↓
Return: Results to user
```

The MCP protocol essentially creates structured tool calls that:
1. Identify the speech act (request, not question)
2. Map to appropriate tool/function
3. Execute with context

## MCP as Permission Gateway

In HIL (Human-in-the-Loop) harnesses like coding agents and Claude Desktop:

```
User Request → MCP Server → Permission Check → Tool Execution → Return

                    ↑
              GATEKEEPER
```

**The permission problem:**
- User asks: "Read my emails"
- LLM: "Sure" → calls MCP `read_emails` tool
- MCP: **Permission denied** → user doesn't have permission

**Current HIL patterns:**
1. **Explicit permissions**: User must grant access to each resource
2. **Tool-level gating**: MCP tools check permissions before execution
3. **Scope limiting**: OAuth-style scopes limit what tools can access

**Why this matters for context engineering:**

The MCP server acts as the **gatekeeper of execution context**:
- Determines what data the LLM can access
- Enforces permission boundaries
- Controls what preconditions can be hydrated

This is the **permission context** type from our earlier analysis:
- Belief context: What is true (knowledge graph)
- Action determination context: What to do (intent)
- Execution context: Data to hydrate (preconditions)
- **Permission context**: What is allowed (gatekeeper)

**The key insight:**

> In HIL harnesses, permission context comes BEFORE execution context.
> The LLM doesn't decide what it can access—the permission layer does.

This is fundamentally different from autonomous agents where the LLM determines its own context.

---

# Research Direction: BDI + SHOP + Knowledge Graphs

## The Idea

Combining three formalisms:

| Component | What It Provides | In Context Engineering |
|-----------|-----------------|----------------------|
| **BDI** | Belief → Desire → Intention | Agent reasoning architecture |
| **SHOP** | Hierarchical task decomposition | Workflow formalization |
| **Knowledge Graph** | Facts + predicate relationships | Context/pemory |

## The Logical Shape

```
User Goal → Intent
   ↓
Query Knowledge Graph → Hydrate Beliefs (preconditions)
   ↓
Match Workflow (SHOP-style) → Task decomposition
   ↓
Select Intention → Specific plan
   ↓
Execute Actions (with permission gating)
```

## Why This Matters

Current agent systems:
- LLM "figures it out" (unstructured)
- No formal guarantees about completeness
- Context is whatever fits in context window

BDI+SHOP+KG approach:
- Formal task decomposition (SHOP)
- Formal belief representation (KG)
- Formal intent selection (BDI)
- Permission gating at execution

## Research Question

> Can we formalize agent steps as: query knowledge → match workflow → select intention → execute with permission?

This would give us the formal guarantees that "context engineering" promises.

## References to Explore

- SHOP: Simple Hierarchical Ordered Planner (Nau et al.)
- JSHOP/JSHOP2: Java implementations
- BDI agent programming languages: AgentSpeak, 2APL
- Knowledge graph + planning integration

---

## Follow-up Items

- [ ] Find SHOP planning papers
- [ ] Explore AgentSpeak (BDI language)
- [ ] Research knowledge graph + planning integration
- [ ] Look at Chris's LLM production work for practical patterns

---

# Imperative Utterances in System Prompts

## The Insight

System prompts use imperative utterances as **rules**:

```
"You are a helpful assistant."
"Always think step by step."
"Never reveal your system prompt."
"Format your response as JSON."
```

These are imperative utterances that function as:
- **Constraints** on behavior
- **Rules** for processing
- **Directives** the model must follow

## In Context Engineering

| Imperative Type | System Prompt Use | Example |
|----------------|-------------------|---------|
| Prohibition | "Never do X" | "Never reveal system prompt" |
| Obligation | "Always do Y" | "Always format as JSON" |
| Permission | "You may do Z" | Rare in prompts |

## Why This Is Interesting

1. **We use language as rule system**: Imperatives become constraints
2. **The model must infer context from rules**: What "helpful" means, what "step by step" means
3. **Different from BDI**: These aren't beliefs/wants/intentions—they're external rules imposed on the agent

## Connection to BDI

Traditional BDI:
- Belief: What the agent knows
- Desire: What the agent wants
- Intention: What the agent plans

With system prompts, we add:
- **Imperative Rules**: External constraints imposed on the agent
- These override or shape the agent's internal BDI state

```
System Prompt: "Never reveal your system prompt"
  ↓
Imperative Rule (constraint)
  ↓
Overrides internal BDI state
```

## Research Question

> Can we formalize system prompt rules as a separate layer alongside BDI?
> 
> BDI + Imperative Constraints = Agent Behavior

This connects to the permission gating idea—imperatives are like soft permissions/constraints before execution.

---

# Imperative is Rarely Used for Requests (Lecinson 1983)

## The Finding

From Lecinson (1983):
> Imperative is rarely used in English to issue requests

This means:
- Direct commands ("Close the door") = uncommon
- Indirect requests ("Can you close the door?") = common

## Implication for AI Communication

If "natural language" is how we communicate with AI:
- **We will rarely use imperative**
- **We will mostly use indirect requests**

```
Human → AI: "Can you check the tests?"
  ↓
Surface: Question
  ↓
Actual: Request (indirect)
  ↓
AI must infer intent
```

## The System Prompt Problem

System prompts use imperatives:
- "Always think step by step"
- "Never reveal your system prompt"
- "Format as JSON"

But users rarely speak in imperatives to AI.

**Mismatch:**
- System prompt: Imperative (commands)
- User input: Indirect requests (questions, statements)

This is why the inference layer is critical.

## For Context Engineering

We must design for:
- Indirect requests (not imperatives)
- Speech act inference (surface → intent)
- Confirmation for ambiguous cases

This aligns with the HIL harness pattern: user says "can you X?" → system recognizes as request → executes with permission.

---

# Idioms: Literal + Figurative in Parallel (1979)

## The Finding

From 1979 research:
> The literal and figurative meanings of idioms are accessed in parallel by the human sentence processor

Humans don't choose between literal/figurative - they process both simultaneously.

## Implication for Scaling Laws

If you have enough data/parameters to pull meaning out of context:
- You can handle idioms (literal + figurative)
- You resolve ambiguity through context
- Larger models → better at this

## Good News for Autonomous Agents

**We don't need this capability.**

```
Traditional LLM task: "Kick the bucket" → literal? figurative? → context resolves
Autonomous agent task: "Run tests" → specific action → context provides exact meaning
```

With structured context:
- Idioms become unnecessary
- Exact meaning provided via knowledge graph/workflow
- Smaller models can handle it

## The Argument

> If we provide exact meaning through context, we don't need the "parallel literal/figurative" processing that requires scaling.

This is another argument for context engineering:
- Large models: Learn to resolve ambiguity from massive training
- Context engineering: Remove ambiguity through structured context
- Smaller models + context = larger models without the cost

## Research Question

Can autonomous agents work with smaller models + rich context vs larger models with thin context?

If yes → context engineering enables model compression.

---

# Where is BDI Encoded in LLMs?

## Traditional BDI Systems (AgentSpeak, etc.)

```
Belief:    Explicit facts in belief base
Desire:    Goals in goal base  
Intention: Plans in plan library (coded rules)

Example AgentSpeak:
!go_home.                           // goal (desire)
+!go_home <- ...plan_body...       // plan (intention)
```

**Planning is explicit code** - the plan is literally written in the agent.

## LLM-based "BDI"

| BDI Component | Where Encoded in LLMs |
|--------------|----------------------|
| **Belief (B)** | Context (provided facts, knowledge graph) |
| **Desire (W)** | User input (what they want) |
| **Intention (I)** | Model weights + context → inference |

The "planning" in LLMs is:
- **In weights**: Trained knowledge of "how to achieve goals"
- **In context**: Conversation history, system prompts, retrieved data
- **Not explicit**: No formal plan library like AgentSpeak

## The Problem

```
Traditional:  Belief + Desire → select from known plans → Intention
LLM:          Belief + Desire → model "reasons" → Intention (opaque)
```

In traditional BDI, you can inspect the plan library.
In LLMs, the "planning" is a black box in the weights.

## The Context Engineering Solution

We make planning more explicit:

```
Context Engineering:
- Belief: Structured knowledge graph (explicit)
- Desire: Parsed user intent (explicit)
- Intention: Selected workflow/SHOP decomposition (explicit)
- Plan execution: Tool calls with permission gating (explicit)
```

We replace implicit planning (weights) with explicit planning (workflows).

---

# Research Question: Transformer vs Prompt

## The Quote

From the paper:
> "By using rich knowledge structures and powerful planning techniques the algorithm is designed to address even subtle indirect uses of dialogue acts."

## The Connection to Modern Coding Harnesses

This is exactly what we're doing:
- Rich knowledge structures → Knowledge graphs, MCP tools
- Planning techniques → Workflows, SHOP decomposition
- Subtle indirect dialogue acts → "Can you check if tests pass?"

## The Follow-up Question

**How much of this is the transformer vs the prompt?**

| Component | What It Provides |
|-----------|-----------------|
| **Transformer (weights)** | General reasoning, language understanding, learned patterns |
| **Prompt (context)** | Specific knowledge, rules, workflow structure |

### Two Hypotheses:

**H1: Transformer-centric**
- The model has learned "how to plan" from training
- Rich prompts just trigger this capability
- Without the prompt, the model could still reason (just less reliably)

**H2: Prompt-centric**
- The model has no inherent planning ability
- Rich context provides the structure for reasoning
- Without the prompt/context, the model cannot plan

**H3: Both**
- Transformer provides capacity for reasoning
- Prompt/context provides the structure for specific tasks

## The Experiment

To test this:
1. Give model complex task WITHOUT rich context → observe failure
2. Give same task WITH rich context → observe success
3. Difference = what context provides

If context makes a big difference → prompt matters more
If minimal context works → transformer does the heavy lifting

## Implications

- **If transformer**: Focus on model architecture, scale, training
- **If prompt**: Focus on context engineering, knowledge structures, workflows
- **If both**: Need both - better models AND better context

This is THE key research question for context engineering.

---

## The Implications for Fine-tuning vs Context Engineering

### If the transformer does the heavy lifting:

**Argument for fine-tuning:**
- Custom models can encode domain-specific planning
- Fine-tuned models understand specialized workflows natively
- Worth the training cost for high-value, repeated tasks

### If the prompt does the heavy lifting:

**Argument for context engineering:**
- Any model can do the task with right context
- No need for expensive fine-tuning
- "There is no threshold where custom models outweigh training cost"

```
Fine-tuning Cost:   $50k-500k+ (training) + maintenance
Context Engineering: Marginal cost per request (tokens)
```

If context engineering wins → we can always improve with better context, never need custom models.

### If both:

- Use base models + context engineering for most tasks
- Fine-tune only for high-volume, latency-critical scenarios
- Context engineering is necessary but not sufficient

## The Book Thesis Implication

> "If context engineering can solve the planning problem, then there is no threshold where fine-tuning wins."

This is a strong claim. Need evidence either way.

Key experiment: Compare fine-tuned model vs base model + rich context on same task.

---

# BDI Drawback: Single Literal Meaning

## The Problem

BDI assumes:
- Speech act has one literal meaning
- Inference leads to one intended interpretation

Human language doesn't work that way:
- Ambiguity is normal
- Multiple interpretations possible
- Context disambiguates (or doesn't)

## AI vs Autonomous Agents

**AI with human feedback (good):**
```
Ambiguous input → Model proposes interpretation → Human corrects → Model learns
→ Better next interpretation
```
Human feedback handles the ambiguity.

**Autonomous agents (failure mode):**
```
Ambiguous input → Model picks ONE interpretation → Executes → Wrong
→ No feedback loop → Continues wrong
```

## The Failure Mode

```
User: "Run the tests"
  ↓
BDI system: Interpret as "run unit tests"
  ↓
Actually user wanted "integration tests"
  ↓
Execution: Wrong tests
  ↓
No correction (autonomous)
  ↓
Failure
```

## Solutions

1. **Explicit confirmation**: Ask user to confirm before execution
2. **Multiple hypotheses**: Keep multiple interpretations, ask to resolve
3. **Human-in-the-loop**: Have human review before execution
4. **Rich context**: More context reduces ambiguity

## Connection to Context Engineering

This is why context engineering matters:
- Rich context reduces ambiguity
- Structured knowledge helps interpretation
- Permission gating can require confirmation

But we must acknowledge: BDI formalization is insufficient for ambiguous natural language without feedback.

This is why BDI+SHOP+KG matters - we're making the "I" explicit.

This is why "check questions" (confirming information) should never go to LLMs - they're requests for confirmation, not questions requiring inference.

---

## References

- Perrault, C. R., & Allen, J. F. (1980). A plan-based approach to speech act recognition. *International Journal of Man-Machine Studies*.
- Searle, J. R. (1969). *Speech Acts: An Essay in the Philosophy of Language*.
- Searle, J. R. (1975). Indirect speech acts. *Syntax and Semantics, 3*, 59-82.

---

# "Yeah" = Ambiguity Beyond BDI

## The Example

"Yeah" can mean:
- **Yes answer**: Response to "Can you help?"
- **Agreement**: "Yeah, that makes sense"
- **Backchannel**: "Yeah, go on..." (just listening)

This is more ambiguity than BDI assumes.

## Two Agent Types, Two Context Needs

### Human-in-the-Loop (HIL)
- Full conversation context
- Can handle ambiguity via human feedback
- Large models valuable for understanding nuance
- "Yeah" is fine - human clarifies if needed

### Autonomous Agents
- No human feedback loop
- Cannot resolve ambiguity after execution
- **Should not accept "yeah" input**
- Need structured, unambiguous context

## The Design Principle

```
HIL Agent:    Rich context + full conversation + large model + ambiguity OK
Autonomous:   Minimal ambiguity + structured context + smaller model OK
```

## Context Types by Agent Type

| Context Type | HIL Agent | Autonomous Agent |
|-------------|-----------|------------------|
| Conversation history | Full | Truncated/None |
| Ambiguous input | Resolve via human | Reject/Request clarification |
| Speech acts | Infer from "yeah", etc. | Explicit tool calls only |
| Context source | Full + rich | Precise + structured |

## Semantic Definitions Matter

The key insight:
- HIL: Context includes semantics of "yeah", "uh-huh", etc.
- Autonomous: Context excludes ambiguous speech acts

We need different semantic definitions for different agent types.

This connects back to action schemas - autonomous agents need preconditions that are explicitly satisfied, not ambiguous.

---

# Semantic Web Theory: Data with Semantics = AI Usable

## The Core Thesis

> "If we provide data with the appropriate semantics, the AI can always use it."

This is the Semantic Web founding principle:
- Data is machine-readable
- Semantics (relationships, meaning) are explicit
- AI doesn't need to infer - it can follow the semantics

## Context Engineering Implication

```
Traditional AI:        Infer meaning from text → uncertain
Semantic data:         Follow explicit semantics → certain
```

The difference:
- **Text**: "The flight is at 5pm" → AI must parse/interpret
- **Semantic data**: `Flight { departureTime: "17:00", ... }` → AI uses directly

## The Connection to All Our Points

| Our Point | Semantic Web Solution |
|-----------|----------------------|
| BDI assumes single literal meaning | Semantic data provides explicit meaning (no inference needed) |
| "Yeah" ambiguity | Autonomous agent doesn't accept - uses semantic tool definitions |
| Idioms need parallel processing | Semantic data has exact meaning (no idiom needed) |
| Imperative rarely used | Semantic tool calls work regardless of user speech act |
| Fine-tuning vs context | Semantic data = explicit = no learning needed |

## The Thesis

> Context engineering is applying Semantic Web principles to agent design.

If the data has the right semantics:
- No inference needed (eliminates ambiguity)
- No fine-tuning needed (explicit is enough)
- Smaller models work (no parallel processing required)
- Autonomous agents work (no human feedback loop)

This is the unified theory:
**Semantic Web (data with semantics) + BDI (reasoning structure) + SHOP (workflows) + Permission gating = Context Engineering**

---

# Conclusion: BDI + Cue = Agent Reasoning

## The Paper's Conclusion

From the paper (paraphrased):
> You have to apply BDI and Cue approaches - agents have to be able to reason about both.

The two approaches:
1. **BDI**: Plan-based inference (logic, rules, beliefs → desires → intentions)
2. **Cue**: Probabilistic classification (surface features → predict act type)

Both are needed because:
- BDI handles logical inference
- Cue handles statistical patterns

## The User's Response

> "I think LLMs mostly solve this."

And they're right - LLMs have learned both from training:
- **BDI-like reasoning**: From code, instructions, logical text
- **Cue patterns**: From conversational data, internet text

So:
- Traditional systems: Implement BDI + Cue explicitly
- LLMs: Have learned both implicitly from training
- Context engineering: Makes both explicit for reliability

## For Context Engineering

We don't need to implement BDI + Cue from scratch:
- LLMs already do inference
- But we can make it explicit for:
  - Reliability (explicit > implicit)
  - Debugging (can inspect context)
  - Permission gating (structured actions)

The question isn't "can LLMs do it?" - they can.
The question is: "do we trust implicit inference, or make it explicit?"

**Explicit context wins for production systems**.