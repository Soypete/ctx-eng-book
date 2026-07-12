# Agents Are Workflows

## Beat 1: The Agent Illusion

The word "agent" conjures something autonomous: a thing that decides, acts, and learns. Something that roams your system like a digital worker, making judgment calls in real-time. This imagery sells products. It also obscures what's actually running in production.

Most agents are workflows wearing a trench coat. Underneath the mystical branding lies a state machine, a loop with exit conditions, a tool-calling protocol, and a harness that manages execution. The model provides the intelligence. The workflow provides the reliability. Neither alone is sufficient.

This distinction matters because it determines where failures happen. When an agent loops forever, it's rarely the model's fault—it's the workflow's exit criteria. When an agent takes the wrong action, it's rarely a reasoning failure—it's the tool definitions or the context provided. The agent is only as reliable as the workflow that runs it.

This chapter examines the systems beneath the agent: ReAct patterns, planning mechanisms, execution harnesses, state machines, durable execution, and event-driven architectures. We answer the reliability question directly: when should the model decide, and when should software decide?

## Beat 2: ReAct and the Reasoning-Action Pattern

The ReAct pattern—Reasoning and Acting—became the de facto standard for agent development after its 2022 introduction. The pattern interleaves thought, action, and observation in a loop:

1. The model reasons about the current state
2. The model selects and executes a tool
3. The system returns the tool's output as observation
4. The model reasons again with new context

This looks like autonomous reasoning. It's actually a structured loop with typed outputs. The "thought" is a specific tool call argument. The "action" is a function invocation. The "observation" is structured input back to the model. Every step is a contract between model and software.

```python
# ReAct loop simplified
while not done:
    thought = model.generate(context + history)
    action = parse_tool_call(thought)
    observation = execute_tool(action)
    context += [thought, action, observation]
```

The power of ReAct is constraints: it forces the model to externalize reasoning in a parseable format. This enables software to intervene. If the model calls a dangerous tool, the harness can reject it. If the model enters an infinite loop, the harness can terminate. The workflow controls the execution; the model controls the decisions within that execution.

The limitation is the same: ReAct constrains what the model can do. Complex reasoning sometimes requires breaking the pattern—holding mental state the tool schema can't express, or making decisions that don't map to available tools. The workflow enables reliability by restricting capability. This tradeoff is fundamental to agent design.

## Beat 3: The Harness Problem

Every agent runs inside a harness: the infrastructure that executes the loop, manages context, handles errors, and enforces limits. The harness is where reliability lives or dies. Most agent failures are harness failures.

The harness responsibilities:

- **Loop management**: Running the reasoning-action cycle, detecting completion, enforcing iteration limits
- **Context assembly**: Building the prompt from system instructions, retrieved context, conversation history, and tool definitions
- **Tool execution**: Calling functions, handling errors, validating outputs, managing timeouts
- **State persistence**: Saving checkpoint data for durable execution
- **Observability**: Logging reasoning traces, recording tool calls, tracking costs
- **Safety enforcement**: Rate limiting, budget caps, dangerous action blocking

Consider a harness that runs without iteration limits. The model might enter a loop: call search, get irrelevant result, call search again, get same result, repeat indefinitely. Without a harness limit, this consumes unlimited tokens. With a harness limit (max 10 iterations), the loop terminates and the system can evaluate whether progress was made.

The harness also handles tool errors. If a tool call fails—network timeout, invalid parameters, service unavailable—the harness decides what happens. Options include: retry with backoff, skip and continue, halt with error, or fallback to alternative tool. The model doesn't handle these failures; the harness does.

This is the first answer to our reliability question: software decides for harness concerns (iteration limits, error handling, timeouts, cost controls). The model decides for task concerns (which tool, what arguments, how to interpret results). The boundary between these domains is the first design decision in any agent system.

## Beat 4: State Machines and Agent State

Agents are state machines. Not metaphorically—literally. The agent occupies discrete states, transitions between them based on conditions, and maintains memory of past states. Modeling agents as state machines makes behavior explicit, testable, and debuggable.

Simple agents have minimal state:

```
IDLE → REASONING → ACTING → OBSERVING → (IDLE or REASONING)
```

Complex agents have richer state machines:

```
IDLE → PLANNING → EXECUTING → VALIDATING → (COMPLETE or RECOVERING or FAILED)
```

The model doesn't manage this state machine. The workflow does. The model contributes by selecting transitions—deciding to move from PLANNING to EXECUTING, or from EXECUTING to VALIDATING. But the software enforces the valid transitions, detects completion, and handles errors.

State machine design surfaces hidden complexity. Consider a research agent that must:

1. Search for papers
2. Download papers
3. Extract key findings
4. Synthesize into summary

Modelers might think: "I'll just prompt it to do these steps." The state machine reveals the complexity:

- After search, must check: any results? If none, return to search or escalate?
- After download, must check: accessible? If paywalled, skip or alert?
- After extraction, must check: quality sufficient? If poor, retry or use alternative source?
- After synthesis, must check: meets requirements? If incomplete, iterate or complete?

Each decision point is a state transition. Each transition has conditions, actions, and failure modes. The model handles the reasoning within states; the state machine handles the transitions between states.

## Beat 5: Durable Execution

Agents run for minutes or hours. Infrastructure fails in seconds. Durable execution ensures agent progress survives system failures: crashes, restarts, network partitions, scale-in events.

Durable execution requires checkpointing: saving complete agent state at safe points, then resuming from that checkpoint on failure. The checkpoint includes:

- Current state in the state machine
- Accumulated context (conversation history, retrieved documents)
- Pending tool calls and their arguments
- Execution metadata (iteration count, timestamps, costs)

On failure recovery, the harness loads the checkpoint, restores state, and continues execution. The agent never knows it crashed.

This is not optional for production agents. Without durable execution, a 30-minute agent task fails permanently on a 5-minute infrastructure failure. With durable execution, the same failure triggers a seamless resume.

Implementation approaches:

- **Database persistence**: Checkpoints stored in PostgreSQL, with state machine transitions as transactions
- **Event sourcing**: Complete history preserved, state reconstructed by replaying events
- **Message queue with acknowledgment**: Tool calls acknowledged only after successful completion, enabling retry
- **Saga patterns**: Multi-step agents as compensating transactions when failures require rollback

The choice affects reliability semantics. Database transactions provide atomicity. Event sourcing provides auditability. Message queues provide at-least-once delivery. Each has tradeoffs; the agent requirements determine the approach.

## Beat 6: Event-Driven Agent Systems

So far we've discussed agents as loops: start, iterate, complete. Many production agents are not loops—they're event-driven systems that respond to triggers: webhooks, queue messages, schedule cron jobs, user actions.

Event-driven agents follow a different pattern:

1. Event arrives (new support ticket, scheduled check, user query)
2. Agent initializes with event context
3. Agent executes fixed workflow or dynamic reasoning
4. Agent produces output: response, database update, another event

The event-driven model changes reliability engineering:

- **Stateless initialization**: Each event starts fresh, context assembled from event data plus retrieved context
- **Idempotency**: Processing the same event twice should produce same result (or graceful handling)
- **Timeout handling**: Long-running agents must report progress, checkpoint, or yield for later resumption
- **Event ordering**: Some agents require ordered event processing; others are embarrassingly parallel

Consider an agent that processes incoming emails. Events arrive as webhooks. The agent must:

1. Parse email content
2. Retrieve customer context
3. Generate response
4. Send reply

Each email is independent—parallel processing is natural. But if two emails from the same customer arrive simultaneously, ordering might matter for consistency. The event-driven architecture must decide: process in parallel and accept potential reordering, or sequence per-customer and accept reduced throughput.

Event-driven agents often require coordination: routing events to correct agent instances, managing concurrency, handling dead-letter queues for failed processing. This is infrastructure work, not AI work. The model provides the intelligence; the event system provides the reliability.

## Beat 7: The Decision Boundary

We began with a question: when should the model decide, and when should software decide? We now have a framework to answer it.

Software should decide for:

- **Structural concerns**: Loop execution, state transitions, iteration limits, completion detection
- **Infrastructure concerns**: Error handling, timeouts, retries, checkpointing, scaling
- **Safety concerns**: Authorization, rate limiting, cost caps, dangerous action blocking
- **Governance concerns**: Audit trails, compliance requirements, access controls

Model should decide for:

- **Reasoning concerns**: How to solve the task, which tools to use, how to interpret results
- **Context concerns**: What information is relevant, how to combine multiple sources
- **Output concerns**: What to say, how to format, what details to include
- **Adaptation concerns**: How to respond to unexpected situations within the workflow

The boundary isn't fixed. Different agent architectures place different decisions in each domain. Fully autonomous agents let the model control nearly everything. Highly structured agents let the model control only tool selection. The right boundary depends on reliability requirements, task predictability, and safety constraints.

Here's a practical heuristic: if the decision can be expressed as rules, codify it in software. If the decision requires judgment beyond rule expression, let the model decide—with software guarding the boundaries.

The model decides how to reason. The workflow decides how to reason safely.

## Beat 8: Reliability Engineering for Agents

Apply standard reliability engineering to agents:

- **Testing**: Test state machine transitions, tool error handling, checkpoint recovery, event ordering
- **Observability**: Log state transitions, tool calls, context size, reasoning traces, costs
- **Failure modes**: Map failure modes to mitigations (modelhallucination → validation; tool failure → retry; loop → iteration limit; context overflow → truncation)
- **SLA definition**: Define latency, success rate, cost per task—measure against SLA

The agent isn't magic. It's a workflow with a model inside. Engineer it accordingly.

**The next chapter** examines the cost of context—and how to build agents that don't burn through budget.

(End of file - total 243 lines)