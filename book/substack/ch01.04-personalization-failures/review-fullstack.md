# Full-stack engineer review — Your Model Doesn't Know Your Users. Your Retrieval Does.

**Would I finish it?** yes. The Dana opener is a bug I could ship next sprint, and the post quickly shows me where that bug lives in the code.
**Would I share it?** yes. I'd send it to the teammate who wired our chat feature's tool calls straight to `req.body.userId`.

## What lands
- "a `user_id` from the request never gets a vote." This is the one line I'll repeat in code review.
- Failure 2 names the real hole: a tool argument "the model fills in" is user input. I hadn't framed it that way.
- The `Store.Read` snippet is short, and the comment on line 69 explains the whole idea.
- The diagram is where I slowed down. It shows the model as the last box, and that was enough to get the point across.
- The Guidelines are a checklist I can paste into a PR template.

## Where I got lost or rolled my eyes
- "a cache in front of the preference store never got invalidated." Then the fix is a max age ("can't be older than an hour"). Those are two different problems. If I change my setting and the assistant ignores it for up to 59 minutes, users will see that as the same bug. A TTL doesn't fix it; invalidating the cache when the setting is written does.
- "for _, r := range s.Rows". A loop over an in-memory slice reads like a toy. My store is Postgres. RLS gets one sentence at the end, and that's the part I'd actually build.
- "The model's learned behavior is ordinarily the same between requests." The hedge makes me wonder what the exceptions are. Either explain it or cut it.
- "Enforce it with a capability… The preference store checks the capability on every read." I can't tell where this object lives in my stack. Is it minted in auth middleware, put on the request context, and passed into the tool handler? That's the plumbing I'd have to write.

## Missing for me
- A tool-call handler snippet showing the model's `get_preferences(user_id)` call being ignored and the capability taken from `ctx` instead. That's the exact spot where the Dana bug happens in an LLM app.
- A 5-line Postgres RLS policy plus `SET app.current_user` per request.
- One sentence on cost: reading fresh preferences on every request adds a DB round trip to my latency budget. Is that okay, or should I cache with write-through?

## Top 3 edits (ranked)
1. Add a tool-handler example (`func handleGetPrefs(ctx, args)`) that ignores the `user_id` argument and reads the capability from the request context. That makes Failure 2 concrete for anyone running function calling.
2. Fix Failure 3 so the fix matches the bug: invalidate the cache when a preference is written, and keep max age as the backstop for inferred data.
3. Swap or follow the in-memory loop with the Postgres RLS version, even if it's just the `CREATE POLICY` line.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
