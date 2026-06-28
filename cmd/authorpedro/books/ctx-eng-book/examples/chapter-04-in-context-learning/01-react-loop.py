# ReAct demonstrates speech act interpretation via structured reasoning.
# The model doesn't just act — it first REASONS about what the user meant.
# This example shows the ReAct loop: Think → Act → Observe → Repeat.

# In a real autonomous agent, this loop runs multiple times.


def call_tool(tool_name: str, args: dict) -> str:
    """Simulated tool call - in real code, this calls an LLM or API."""
    tools = {
        "get_recent_errors": lambda: "Found 3 auth-related errors in the last hour",
        "search_code": lambda pattern, location: (
            f"Found 2 files: {location}/{pattern}.go"
        ),
    }
    return (
        tools.get(tool_name, lambda: "Unknown tool")(**args)
        if callable(tools.get(tool_name))
        else "Unknown tool"
    )


def react_loop(user_request: str):
    """
    Simulated ReAct reasoning trace for: "fix the API"

    BEAT 1: Speech act interpretation — reason BEFORE acting
    The model explicitly interprets what the user WANTS DONE.
    """
    print(f"User request: {user_request}")
    print("\n--- ReAct Reasoning Trace ---\n")

    # Step 1: Reason about intent
    thought = (
        "The user said 'fix the API' but didn't specify what kind of fix. "
        "I need to interpret what they actually want. Possibilities: "
        "1) Fix a bug, 2) Update endpoint, 3) Improve performance. "
        "I should check recent error logs first."
    )
    action = "get_recent_errors"
    observation = call_tool(action, {})

    print(f"Thought: {thought}")
    print(f"Action:  {action}")
    print(f"Observe: {observation}\n")

    # Step 2: Refine understanding based on observation
    thought = (
        "The errors are in the auth handler. The user likely means "
        "'find the bug in the auth handler' rather than a general API fix."
    )
    action = "search_code"
    observation = call_tool(action, {"pattern": "auth", "location": "api"})

    print(f"Thought: {thought}")
    print(f"Action:  {action}")
    print(f"Observe: {observation}")

    print("\n--- Key insight ---")
    print("Without ReAct, the model GUESSES and calls a random tool.")
    print("With ReAct, the model REASONS about intent before acting.")


if __name__ == "__main__":
    react_loop("fix the API")
