# BEAT 4: Indexicality Definition
# Indexicality: words whose meaning depends on WHO is asking, WHEN, and WHERE.
# "the data" = different data for different users.

# This demonstrates the core linguistic concept from Jurafsky

USER_CONTEXTS = {
    "alice": {
        "role": "sales_manager",
        "department": "sales",
        "files": ["sales_q1.csv", "sales_q2.csv"],
    },
    "bob": {
        "role": "marketing_manager",
        "department": "marketing",
        "files": ["campaigns_q1.csv", "ads_q2.csv"],
    },
}


def demonstrate_indexicality():
    """
    Same phrase, different meaning — depending on who's asking.

    This is why context engineering MUST provide explicit scoping.
    """
    phrase = "my data"

    print("=== INDEXICALITY: Same Phrase, Different Meaning ===\n")
    print(f"Phrase: '{phrase}'\n")

    for user_id, context in USER_CONTEXTS.items():
        print(f"User: {user_id} ({context['role']})")
        print(f"  → '{phrase}' means: {context['files']}")
        print(f"  → Department: {context['department']}")
        print()

    print("--- Key insight ---")
    print("Without explicit scoping, the model must INFER which 'data' you mean.")
    print("Inference = unreliability. Scoping = context engineering.")


if __name__ == "__main__":
    demonstrate_indexicality()
