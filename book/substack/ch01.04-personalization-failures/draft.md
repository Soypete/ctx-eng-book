---
title: Your Model Doesn't Know Your Users. Your Retrieval Does.
subtitle: Personalization fails in four ways, and every one of them is a context-assembly bug, not a model bug.
module: ch01.04-personalization-failures
mood: confused
scheduled: 2026-10-16
---

The assistant opens with "Welcome back, Dana! Here's an update on your three open tickets." The person reading it isn't Dana.

It feels like the model got confused about who it was talking to. It didn't. The model doesn't know who anybody is. Some code fetched Dana's tickets and put them in front of it, and that code is where the bug lives.

This is post 5 in my [context engineering series](https://substack.com/@soypetetech). Last time I traced five production failures to the exact place in the code where each one got through. This time I do the same for personalization, where the failures feel personal because they're about one person.

## Personalization is assembly, not memory

Two definitions first. **Context** is everything the model receives for one request: instructions, retrieved data, the conversation. **Retrieval** is the code that fetches data for that request from a database, a search index, or a profile service.

**Personalization** means fitting a response to a specific user: their preferences, history, permissions, and constraints. People talk about it as if the model "remembers" the user. It doesn't. The model's learned behavior is ordinarily the same between requests. What changes is the data your code retrieves and adds to the context. A framework might call your preference store "memory." It's still a table, and it's only as reliable as the code that reads it.

So every personalized request has to answer four questions before the model sees anything:

- **Who is the user?** Identity, from the authenticated session.
- **What are they trying to do?** The current task.
- **What may they access for that task?** Permissions.
- **What do they prefer, as of now?** Settings, history, constraints, and how current they are.

Get one of those wrong and personalization breaks in a predictable way. Here are the four I see most.

## Failure 1: missing preferences

A senior Go engineer asks for a refactoring approach and gets beginner advice in Python. The model wasn't being dense. Nobody retrieved the fact that this user writes Go.

The fix sounds obvious: retrieve the user's preferences at request time, and only the ones relevant to the task. The subtle part is where a preference came from, its **provenance**. A user can **declare** a preference ("I write Go"), the system can **infer** one from behavior (they opened a lot of Python files last week), or an organization rule can impose one (every response includes a license header). Those carry different authority, so don't flatten them into one profile. A user's correction should beat an inference, and organization rules should constrain both.

## Failure 2: the wrong user's context

That's Dana. It's the most serious personalization failure, because it's a privacy breach that looks like a model quirk.

The usual cause is that retrieval trusts an identity it was handed. Somewhere there's a `user_id` in a request body, a URL, or a tool argument (a parameter the model fills in when it asks your code to run a function), and the query uses it as-is. The fix is **tenant isolation**: the query must be unable to return rows belonging to anyone else, no matter what the request or the model asks for. (Here a tenant is one user; in a B2B product it's often a whole customer organization. The rule is the same.) Identity comes from the authenticated session, and the preference store enforces it.

## Failure 3: stale preferences

A user updates their settings and the assistant keeps behaving the old way, because a cache in front of the preference store never got invalidated. Last month's preference wins.

Preferences need the same freshness rules as any other retrieved data. Timestamp each one and give each type of preference its own maximum age. Code style can probably be a month old. "Which project am I working on right now" probably can't be older than an hour. Past the limit, refresh it or leave it out, and record which you did.

## Failure 4: over-scoped context

Say a user can read a thousand documents and the task needs one. Load all thousand into the context and anything that goes wrong (a prompt injection, meaning instructions hidden in data the model reads; a logging mistake; an overly helpful answer) can expose all thousand. Load one and the damage is capped at one. Same with preferences: a code review needs your language and experience level, not your billing email.

That's **least privilege** applied to context: retrieve only what the current task needs. Enforce it with a **capability**, a short-lived grant the server issues that says whose data, which fields, for which task, and until when. The preference store checks the capability on every read.

## What enforcement looks like

![Session and task go through the server's policy to a capability; the preference store enforces it, then a freshness and provenance check builds the task-scoped context.](https://raw.githubusercontent.com/Soypete/ctx-eng-book/main/book/substack/ch01.04-personalization-failures/diagram-1.png)

*The model is the last box. Every decision about whose data and which data happens before it, and a `user_id` from the request never gets a vote.*

The [example](https://github.com/Soypete/ctx-eng-book/tree/main/book/examples/ch01.04-personalization-failures) wires the four fixes together in Go. A request handler passes a session token and a task name. The server authenticates the session, looks up the task in a server-side policy (which preference fields each task may use), issues a capability for those fields, and the store enforces it:

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

Notice what's not here: a `user_id` parameter. The capability's subject comes from the session, and its fields are private to the package, so other code can't hand-build a working one. In a real system you'd sign the capability cryptographically and enforce the same rule in the database itself, for example with Postgres row-level security, which filters rows by the current user on every query.

After the read, the example drops preferences older than their type's limit and lets a declared preference override an inferred one, recording every exclusion and why. The tests show the four failures don't happen: a second user's session gets none of the first user's rows, a hand-built or expired capability reads nothing, a code-review task can't see the billing email, a three-hour-old "current project" is excluded under a one-hour limit, and "I write Go" beats an inferred "Python."

## Measure it like a system

"The assistant knows me" isn't testable. These are:

- **Cross-user denial tests** that try to read another user's data and must fail.
- **Freshness:** how old was each preference that reached the context?
- **Provenance coverage:** for each preference used, do you know whether it was declared, inferred, or imposed, and when it was last updated?
- **Task relevance:** did the context contain only fields the task needed?
- **Correction rate:** how often do users correct the assistant's assumptions about them?

For every response, log the user, the capability issued, and each preference's source and timestamp. When someone says "it called me Dana," you can answer with a record instead of a shrug.

## Guidelines

1. **Derive identity from the authenticated session,** never from a request field or a tool argument.
2. **Enforce tenant isolation in the preference store.** Write a test that tries to read another user's preferences and fails.
3. **Issue task-scoped capabilities.** Grant only the fields the task needs, with an expiry.
4. **Record each preference's provenance.** Let a declared preference override an inferred one; let organization rules constrain both.
5. **Set a maximum age per preference type.** Refresh or exclude anything older, and log which.
6. **Log what went into the context:** user, capability, and each preference's source and timestamp.

Personalization isn't the model knowing your users. It's your retrieval code knowing who's asking, what they're doing, and what they're allowed to see, then handing over just that. Next in the series, I'll zoom out: these failures don't live in some separate "AI layer," and treating them as if they did is why teams keep shopping for a better model.
