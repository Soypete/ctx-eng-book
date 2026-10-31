---
title: Your Model Doesn't Know Your Users. Your Retrieval Does.
subtitle: Personalization needs verified identity, current preferences, and a record of why each field reached the model.
module: ch01.04-personalization-failures
mood: confused
scheduled: 2026-11-02
---

The assistant opens with “Welcome back, Dana! Here are your open tickets.” The person reading it isn't Dana.

In this invented scenario, retrieval fetched Dana's records for someone else's session. The greeting looks like model confusion. The privacy failure happened before the model answered.

In the previous post in my [context engineering series](https://substack.com/@soypetetech), I argued that every failure needs an address in the code; personalization gives us a particularly uncomfortable place to look.

## Personalization starts before the answer

**Context** is everything the model receives for a request: instructions, conversation, retrieved records, and tool results. **Retrieval** is the code that fetches those records from a database, search index, or profile service.

**Personalization** means fitting a response to a particular user's preferences, history, permissions, and task. My argument is that reliable personalization starts with assembling that context correctly. It doesn't guarantee the model will use it correctly.

A framework might call your preference store “memory.” It's still a table, and somebody owns the code that reads and writes it.

Before assembling context, answer: who is authenticated, what are they doing, what may they access for that task, and which preferences are current? Here are four ways those checks can fail.

## Missing preferences: keep the source attached

Suppose a senior Go engineer asks for a refactoring approach and gets beginner advice in Python. If the context never included their language or experience, start there before changing models.

Retrieve the preferences relevant to that task. Preserve their **provenance**: where each value came from. “I write Go” is declared by the user. “Probably prefers Python” might be inferred by a job examining recent file activity. An organization rule requiring license headers constrains the response regardless of either preference.

I would keep a preference record with owner, field, value, origin, source reference, and version. For inferred values, also keep when the behavior was observed, when the inference was computed, and the producing job's run identifier. That reference should let an authorized investigator trace “Python” back to its input events.

A user correction should override an inference. Store it as a declared value, invalidate affected cached copies, and make subsequent inference jobs respect the override. Merely fixing this answer leaves tomorrow's job free to resurrect the mistake. Deletion needs the same path through derived records and caches.

There's also a write-side failure: **memory poisoning**, where bad information becomes persistent state. An **agent** is a model that calls tools in a loop; **prompt injection** is an instruction hidden in material it reads. My concern is that an agent could turn such an instruction into a saved “preference.” Treat agent-written values as untrusted candidates with source references, never as user declarations or organization rules.

Drew Breunig defines context poisoning as “when a hallucination or other error makes it into the context, where it is repeatedly referenced” in [How Long Contexts Fail](https://www.dbreunig.com/2025/06/22/how-contexts-fail-and-how-to-fix-them.html). Persisting that error as a preference is my application of his definition: the next request can retrieve the mistake again.

## Wrong user: identity is not a tool argument

Back to Dana. A **tool argument** is a parameter the model supplies when asking application code to run a function. If `get_preferences(user_id)` trusts that argument, the model can select whose records get fetched.

A `user_id` from the request never gets a vote. Derive identity from verified session credentials and enforce that identity in the preference store.

That's **tenant isolation**: preventing one customer's request from retrieving another customer's state. A tenant may be a person or an organization; an organization still needs user-level permissions within it. Least privilege and database access controls are familiar engineering. The additional concern here is that a model supplies tool arguments and reads potentially hostile text before asking for more data.

## Stale preferences: fix the write path too

A user changes a setting, but the next response uses a cached copy. A maximum age doesn't fix that interaction. Invalidate or update the cache when the setting changes, and test that the next read sees the correction.

Inferred preferences have another clock. A job can write a new row today using old behavior. Keep observation time, computation time, and retrieval time separate; copying an old inference must not make it fresh.

Then define maximum ages per field as a backstop. The book's illustrative limits allow code-style preferences to be 30 days old but require project context to be less than an hour old. Those are example choices, not measured recommendations. Refresh or exclude stale values and record the outcome. If the current project is unavailable, ask which project instead of confidently using the old one.

Caching can save a database round trip, but its key must include the authenticated subject and relevant task scope. Its validity must track preference changes. A faster read of the wrong profile is not the latency win I want.

## Over-scoped context: permission isn't purpose

A code review needs language and experience, not billing email. Even authorized fields can be irrelevant. Keeping unnecessary data out reduces what an answer or log could expose; it doesn't contain every possible compromise of the application.

Use **least privilege**: grant only the access needed for this task. A **capability** is a short-lived server-issued grant specifying whose data, which fields, and until when. A **task policy** is the server's mapping from a supported task to permitted preference fields. The preference store checks the grant on every read.

![Verified session and task policy produce a capability; the preference store enforces owner and field scope, then freshness and provenance checks select context for the model.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch01.04-personalization-failures/diagram-1.png)

*The model comes last: identity, field access, and preference selection are decisions made before it receives the data.*

## Put the checks on the request path

The [companion Go example](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch01.04-personalization-failures) implements this path in memory. `ForRequest` authenticates a session token, applies the task policy, issues a capability, and calls `Store.Read`:

```go
func (s *Store) Read(c Capability, now time.Time) ([]Preference, error) {
	if c.subject == "" || now.After(c.expires) {
		return nil, ErrExpired
	}
	var out []Preference
	for _, r := range s.Rows {
		// Another user's row, or a field this task wasn't granted, never leaves storage.
		if r.Owner != c.subject || !c.fields[r.Field] {
			continue
		}
		out = append(out, r)
	}
	return out, nil
}
```

For a tool-calling application, authenticate in request middleware, attach the server-issued capability to the request context, and have the tool handler read it there. Reject any tool-supplied identity override. The model must not choose a broader task policy to obtain more fields.

The loop illustrates the check; production filtering belongs in the database query. Restrict owner and allowed fields before returning records. Database row-level security can enforce which rows a database role may read; field selection needs separate enforcement. Bind the authenticated subject within the request's database transaction, and test with the actual application role. The example's private Go fields are an API boundary, not protection against hostile code in the process. Across services, use verified signed grants or protected server-side grant records.

After retrieval, the example excludes stale rows and prefers declared over inferred values. Its tests cover cross-user exclusion, unknown sessions, empty and expired capabilities, field scope, stale project context, and declared-preference precedence. They don't exercise a live model, database enforcement, cache invalidation, memory writes, or organization rules. Its single `UpdatedAt` timestamp also simplifies the production provenance record described above.

## Guidelines

1. **Bind every read to verified identity.** Add a tool-path test that requests Dana's data from another session and asserts no Dana record reaches the model.
2. **Version the task-to-fields policy.** Reject unknown tasks and unmapped field renames; assert returned fields are a subset of the approved mapping.
3. **Preserve provenance and corrections.** Require source references for inferred values and test that a later inference cannot overwrite a declaration.
4. **Invalidate on writes and enforce freshness on reads.** Test both the immediate correction and the stale-inference case.
5. **Keep governed decision records.** Record protected subject and grant references, policy version, preference versions, source times, and exclusion reasons. Assign access, retention, and deletion ownership; omit raw values and reusable credentials.
6. **Separate control tests from model evaluations.** Run adversarial tool requests through the application, inspect retrieved context as well as answers, and report forbidden returns. Unit tests alone don't establish model reliability.

I want personalization I can debug: whose state was read, why those fields were allowed, how current they were, and which source won. That gives “it called me Dana” a repair path instead of another prompt tweak. Next in the [series: Stop Saying “The AI Failed”](https://substack.com/@soypetetech), I'll widen that argument to the systems around the model.
