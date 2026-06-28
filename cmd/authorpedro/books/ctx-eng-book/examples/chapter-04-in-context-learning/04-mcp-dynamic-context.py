# BEAT 6: Static vs Dynamic Context (Teleprompter Pattern)
# MCP (Model Context Protocol) enables dynamic context fetching during execution.
# This contrasts with static context (set before the model runs).

# === STATIC CONTEXT (Teleprompter) ===
# Set once at compile/start time. Fixed, predictable.

SYSTEM_PROMPT = """You are a data analysis assistant.
You have access to a sales database.
Always report in USD."""

EXAMPLES = [
    {
        "input": "Q1 revenue?",
        "output": "SELECT month, revenue FROM sales WHERE month LIKE '2025-Q1%'",
    },
]

# This context is FIXED. Model cannot fetch new information.

# === DYNAMIC CONTEXT (MCP) ===
# Fetched during execution. Flexible but variable.

# In real MCP, the model can call tools to fetch additional context
MCP_TOOLS = {
    "fetch_sales_data": "Gets current sales data from database",
    "fetch_user_profile": "Gets user's role and permissions",
    "fetch_recent_errors": "Gets recent system errors",
}


def run_with_static_context(user_query: str):
    """Static context: model works with what it was given at start."""
    print("=== STATIC CONTEXT (Teleprompter) ===")
    print(f"System prompt: {SYSTEM_PROMPT[:50]}...")
    print(f"Examples: {len(EXAMPLES)} few-shot examples")
    print(f"User query: {user_query}")
    print("→ Model works ONLY with pre-loaded context")
    print("→ No mid-execution fetching\n")


async def run_with_mcp_context(user_query: str):
    """Dynamic context: model can fetch more info during execution."""
    print("=== DYNAMIC CONTEXT (MCP) ===")
    print(f"Initial context: {SYSTEM_PROMPT[:50]}...")
    print(f"MCP tools available: {list(MCP_TOOLS.keys())}")
    print(f"User query: {user_query}")

    # Model can DECIDE to fetch more context mid-execution
    # In real MCP, this happens via tool calls
    print("\nModel reasoning: 'I need current Q2 data to answer accurately.'")
    print("→ Calls MCP tool: fetch_sales_data(months=['2025-Q2'])")

    # Dynamic context fetched here
    q2_data = {"month": "2025-Q2", "revenue": 175000}
    print(f"→ Received: {q2_data}")
    print("→ Model now has fresh data to answer the question\n")


def static_vs_dynamic_summary():
    print("--- Static vs Dynamic Context ---\n")
    print("STATIC (Teleprompter):")
    print("  + Predictable — you know exactly what's in context")
    print("  + Lower variability")
    print("  - Limited to pre-loaded information")
    print("  - Can't handle real-time or user-specific data\n")

    print("DYNAMIC (MCP):")
    print("  + Can fetch fresh, relevant information")
    print("  + Handles user-specific context")
    print("  - More variable (what gets fetched depends on model)")
    print("  - Harder to debug and predict\n")

    print("RELIABLE APPROACH:")
    print("  → Static context for invariants (instructions, constraints)")
    print("  → Dynamic context for fresh information (tool results, real-time data)")


if __name__ == "__main__":
    import asyncio

    user_query = "How does Q2 compare to Q1?"

    # Synchronous example
    run_with_static_context(user_query)

    # Asynchronous (simulated) MCP example
    asyncio.run(run_with_mcp_context(user_query))

    static_vs_dynamic_summary()
