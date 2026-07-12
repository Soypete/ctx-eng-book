# Tool Use Is Structured Context

## Beat 1: The Action Problem

Language models generate text. Text is unbounded—a model can produce any string of tokens. This is powerful but unreliable. When you need a model to do something specific—make a database query, call an API, execute code—you cannot rely on text generation alone.

The solution: constrain the model's output to a structured action space. Give the model tools. Each tool is a defined operation with a defined input schema and a defined output format. The model doesn't generate free text—it selects a tool, provides inputs that conform to the tool's schema, and receives structured results.

This is tool use as context engineering. The tool definitions themselves are context. The tool inputs are constrained context. The tool outputs are structured context. Every tool call is a bounded action in a controlled context.

The reliability question: how do we constrain model actions?

## Beat 2: Toolformer and the Origin of Tool Use

The breakthrough came from Toolformer (2023), a Meta paper that showed language models can learn to use tools through API calls—without fine-tuning on massive datasets. The key insight: you can teach a model to use tools by providing tool descriptions in its context, along with examples of tool calls and their results.

Toolformer demonstrated that models could:
- Decide when to call a tool (not always)
- Select the right tool from a set
- Format inputs correctly according to the tool's API
- Use the tool's output to inform subsequent actions

This was a paradigm shift. Previously, tools were hardcoded into agent frameworks—deterministic logic that invoked APIs. Toolformer showed that the model itself could orchestrate tool use, given the right context: tool descriptions, schemas, and example interactions.

The mechanism is straightforward. You provide the model with a list of available tools, each described with its name, parameters, and purpose. You also provide a few shot examples of successful tool calls. The model learns to generate tool calls in the appropriate format—and more importantly, learns when to do so.

This is context engineering at the action level. You're not just providing information context—you're providing an action context that defines what the model can do.

## Beat 3: Modern Tool Calling

Today, tool calling is a first-class feature in model APIs. OpenAI's function calling, Anthropic's tool use, Google's Gemini API—all provide structured mechanisms for models to invoke external functions.

The pattern is consistent across providers:

1. Define tools with JSON schemas
2. Provide tool definitions in the system prompt
3. Model generates a tool call (not free text)
4. System executes the tool
5. Tool result is injected back into context

A concrete example:

```python
tools = [
    {
        "type": "function",
        "function": {
            "name": "get_weather",
            "description": "Get current weather for a location",
            "parameters": {
                "type": "object",
                "properties": {
                    "location": {
                        "type": "string",
                        "description": "City name"
                    },
                    "unit": {
                        "type": "string",
                        "enum": ["celsius", "fahrenheit"]
                    }
                },
                "required": ["location"]
            }
        }
    }
]

response = client.chat.completions.create(
    model="gpt-4",
    messages=[...],
    tools=tools
)
```

The model doesn't generate weather information from training data. It generates a structured tool call that conforms to the schema. The system executes the call and returns the actual weather data. The model then uses this real data to generate its response.

This is structured context in action. The tool schema constrains what the model can produce. The tool result provides verified, real-time context. The model operates within boundaries you define.

## Beat 4: Schemas as Context Boundaries

JSON and XML schemas are not just for validation—they're context engineering tools. A schema defines what's possible. It tells the model: "here's the shape of valid output."

Consider a function that creates a customer record:

```json
{
  "name": "create_customer",
  "parameters": {
    "type": "object",
    "properties": {
      "email": {
        "type": "string",
        "format": "email"
      },
      "name": {
        "type": "string",
        "minLength": 1
      },
      "tier": {
        "type": "string",
        "enum": ["free", "pro", "enterprise"]
      }
    },
    "required": ["email", "name"]
  }
}
```

This schema does multiple things:
- Constrains valid inputs (email must be valid format, tier must be one of three values)
- Documents what the function does (description in natural language)
- Enables validation (the system can reject malformed calls before execution)
- Defines the action boundary (the model can only create customers with these fields)

The schema is context that says: "these are your options, this is the format, this is what's allowed." The model operates within those bounds.

Schema design matters for reliability. A too-loose schema allows invalid outputs. A too-strict schema prevents valid use cases. The art is matching the schema to the actual capabilities of the underlying system.

## Beat 5: Tool Selection

Given a set of tools, how does a model decide which to call? The same way it decides what to say: through context matching.

The model receives:
- Tool definitions (names, descriptions, schemas)
- The user's request
- Relevant conversation history

It must match the request to the appropriate tool. This is a retrieval problem in action space. The model looks for the tool whose description best matches the user's intent.

Tool descriptions are critical. A poor description:

```
"get_user" - gets user info
```

A good description:

```
"get_user" - Retrieve a user's profile by their unique ID. Returns name, email, account tier, and registration date. Use when the user asks about their account or profile information.
```

The second description provides context: what the tool does, what it returns, and most importantly—when to use it. The "when" is the selection criteria.

Tool selection failures are context failures. If the model selects the wrong tool, it's because:
- The tool descriptions don't adequately distinguish the tools
- The user request is ambiguous
- The context doesn't clarify intent

Fix tool selection by improving tool descriptions. Add usage examples. Provide disambiguation context. Make the selection criteria explicit in the tool definition.

## Beat 6: Tool Routing and Orchestration

Single-tool calls are simple. Multi-step workflows require routing—the model must decide not just which tool, but the sequence of tools, and how to pass outputs between calls.

Tool routing is a pipeline problem. The model generates a tool call, receives a result, and decides what to do next. This is an agent loop, and it requires exit criteria.

A routing example:

```
User: "Transfer $100 to John and tell him it's for dinner"

Model call 1: get_user(name="John") -> returns user_id: 12345
Model call 2: transfer_funds(from=current_user, to=12345, amount=100)
Model call 3: send_message(user_id=12345, text="Transfer of $100 completed - it's for dinner!")
```

Each tool's output becomes input to the next. This is context flowing through a pipeline. The model routes data through tools to accomplish the user's goal.

Reliable routing requires:
- Clear tool output schemas (each tool's output must be well-defined)
- Explicit input requirements (what does each tool need?)
- Exit conditions (when is the task complete?)

Without exit conditions, agents loop. The model keeps calling tools because it doesn't know when it's done. Define completion criteria in the system context: "The task is complete when the user's request has been fulfilled. Confirm completion to the user."

## Beat 7: Constraining Action Through Tool Design

The ultimate reliability mechanism is tool design itself. You constrain model actions by designing the tool space carefully.

Principles for reliable tool design:

**Narrow tool scope.** Each tool should do one thing well. A `search_documents` tool is better than a `do_anything` tool. Narrow scope means predictable behavior.

**Explicit input schemas.** Don't leave format to chance. Specify exactly what fields are required, what values are valid, what the structure should be.

**Clear output schemas.** The model needs to know what it'll receive. Define the output structure so the model can properly consume it in subsequent steps.

**Purposeful descriptions.** Explain not just what the tool does, but when to use it. Include examples of appropriate calls.

**Bounded results.** Don't return unbounded data. Limit query results, truncate long outputs, paginate large responses. The model cannot consume unlimited context.

**Failure modes.** Define what happens on error. Return a consistent error format so the model can handle failures gracefully.

The tool is the boundary. The schema is the constraint. The description is the context. Design them well, and the model operates reliably within them.

## Beat 8: The Reliability Promise

Tool use delivers on the core promise of context engineering: predictable, reliable AI behavior through engineered context.

Without tools, the model generates free text—an unbounded action space where anything can happen. With tools, the model operates in a bounded action space. It selects from defined options, produces structured inputs, and receives structured outputs.

The reliability question—"how do we constrain model actions?"—is answered through tool design. You constrain actions by:
- Defining narrow, well-scoped tools
- Providing explicit schemas that bound inputs and outputs
- Writing clear descriptions that guide selection
- Designing routing pipelines with explicit exit conditions

The model doesn't decide what it can do. You do. The tool definitions are the context that defines its action space. Engineer that context well, and the model behaves reliably.

Reliable AI systems don't hope for good behavior. They engineer the action space to make good behavior inevitable.