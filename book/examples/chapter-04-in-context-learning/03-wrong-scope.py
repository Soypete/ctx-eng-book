# BEAT 4: Indexicality Problem — Wrong Data Scope
# "User's data" means different things to different users.
# This example shows an agent querying the wrong data scope.

# Simulated data stores
DATA_STORES = {
    "sales_db": {
        "table": "monthly_sales",
        "records": [
            {"month": "2025-05", "revenue": 150000},
            {"month": "2025-04", "revenue": 142000},
        ],
    },
    "marketing_db": {
        "table": "campaign_spend",
        "records": [
            {"month": "2025-05", "spend": 45000},
            {"month": "2025-04", "spend": 52000},
        ],
    },
}


def query_with_scope(user_id: str, query: str, scoped_to: str):
    """Simulates an agent querying data with explicit scope."""
    store = DATA_STORES.get(scoped_to, {})
    table = store.get("table", "unknown")
    print(f"User {user_id} queried: '{query}'")
    print(f"  → Scope: {scoped_to} ({table})")


def agent_without_scoped_context(user_request: str):
    """
    Agent has NO explicit scope. It must infer "the data" = which database?
    This is the INDEXICALITY problem: "the data" is ambiguous.
    """
    print("=== AGENT WITHOUT SCOPED CONTEXT ===")
    print(f"Request: {user_request}")

    # Model infers scope — but might get it wrong
    inferred_scope = "marketing_db"  # Wrong! Should be sales_db
    query_with_scope("user_123", user_request, inferred_scope)

    print("  → Result: marketing data (WRONG SCOPE!)")
    print("  → Problem: Not a hallucination — plausible but wrong.\n")


def agent_with_scoped_context(user_request: str, user_context: dict):
    """
    Context engineering provides explicit scope. No ambiguity.
    """
    print("=== AGENT WITH SCOPED CONTEXT ===")
    print(f"Request: {user_request}")
    print(f"User context: {user_context}")

    # Scope is explicitly provided — no inference needed
    scoped_to = user_context["data_scope"]
    query_with_scope(user_context["user_id"], user_request, scoped_to)

    print("  → Result: sales data (CORRECT SCOPE!)")
    print("  → Context engineering solved the indexicality problem.\n")


def main():
    user_request = "analyze the past month's revenue and compare to other months"

    # WITHOUT scoped context — model guesses (often wrong)
    agent_without_scoped_context(user_request)

    # WITH scoped context — explicit about which data
    user_context = {
        "user_id": "user_123",
        "role": "sales_manager",
        "data_scope": "sales_db",  # Explicit scoping
    }
    agent_with_scoped_context(user_request, user_context)

    print("--- Key insight ---")
    print("Wrong scope = wrong answer, not hallucination.")
    print("Context engineering: provide explicit scoping to eliminate inference.")


if __name__ == "__main__":
    main()
