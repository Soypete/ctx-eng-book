# BEAT 2: More Inference = Less Reliability
# Compare two agent approaches:
# 1. Inference-heavy: model figures out "how" on its own
# 2. Goal-directed: model is given explicit "what" to achieve

# This is the harness adding structure via /goal


def run_inference_heavy_agent(user_prompt: str) -> list[str]:
    """
    Agent infers the goal AND the approach. High variance, lower reliability.
    User: "analyze sales"
    Agent decides: which data? what analysis? what format?
    """
    # Model must infer everything — no explicit goal
    decisions = [
        f"User said: {user_prompt}",
        "Inferred goal: analyze sales data",  # guessed
        "Inferred approach: call sales_api",  # guessed
        "Inferred format: summary text",  # guessed
    ]
    return decisions


def run_goal_directed_agent(user_prompt: str, explicit_goal: dict) -> list[str]:
    """
    Agent is given explicit goal specification. Lower variance, higher reliability.
    User: "analyze sales"
    Harness adds: /goal {target: "monthly_revenue_report", constraints: {...}}
    """
    # Goal is explicitly specified — model focuses on execution
    decisions = [
        f"User said: {user_prompt}",
        f"Explicit goal: {explicit_goal['target']}",
        f"Constraints: {explicit_goal['constraints']}",
        "Model focuses on extraction, not inference",
    ]
    return decisions


def main():
    prompt = "analyze sales"

    print("=== INFERENCE-HEAVY AGENT ===")
    print("(Model must infer goal AND approach)")
    for step in run_inference_heavy_agent(prompt):
        print(f"  → {step}")

    print("\n=== GOAL-DIRECTED AGENT ===")
    print("(Harness provides explicit goal)")
    goal = {
        "target": "monthly_revenue_report",
        "constraints": {"currency": "USD", "compare_to": "previous_3_months"},
    }
    for step in run_goal_directed_agent(prompt, goal):
        print(f"  → {step}")

    print("\n--- Key insight ---")
    print("Explicit goal = lower inference burden = higher reliability")
    print(
        "The harness doing inference work (via /goal) makes the model more predictable."
    )


if __name__ == "__main__":
    main()
