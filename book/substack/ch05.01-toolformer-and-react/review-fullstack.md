# Full-stack engineer review — Your Model Can Ask for a Tool. Your Code Decides.

**Would I finish it?** yes — the weather-to-`delete_account` opener is a bug I can picture in my own chat feature, and the Go snippet pays it off.
**Would I share it?** yes — with the teammate who wired our support bot to the tickets API.

## What lands
- "*Registered* and *granted* are different questions, and a lot of incidents live in the gap between them." This is the whole post. It reads like middleware: authn, then authz, then validate, then handler.
- The code block. Four checks in order, and `// only now does anything happen` is the line I'd paste into a PR review.
- "the accounts service should refuse a deletion this user didn't ask for, whatever the harness said". Defense in depth at the service, which is where I actually have control.
- The diagram caption, "Text inside a tool result can't approve anything", is the first thing I'd screenshot.

## Where I got lost or rolled my eyes
- "What Toolformer actually taught." I skimmed the whole section. "fine-tuned... on text augmented with those calls" doesn't change anything in my service. The one usable line, "A model that learned to ask well has still only learned to ask", comes at the end, after I've stopped reading.
- "Resolve grants from the authenticated task." The grants check only saves you because `delete_account` isn't granted. In my app the settings assistant *does* get `delete_account`, because that's what the settings page is for. If both tools are granted to the same task, the injected line still works and the post doesn't say what stops it. I need the next layer: a confirmation step, a per-call user intent check, or splitting read and destructive tools across tasks.
- "Label every observation with its source and time." Label it how? A wrapper string? A field? Delimiters? I can't ship a guideline I can't see.

## Missing for me
- A 5-line example of a labeled observation and a rejection observation as the model actually receives them.
- A loop cap. An agent loop with no max iterations or timeout is a latency and cost incident waiting to happen. One line in the Guidelines would fix it.
- A note on where this code goes in a TS/Node stack. Go is fine, but say it's language-agnostic.

## Top 3 edits (ranked)
1. Cover the case where the dangerous tool *is* granted: require explicit user confirmation for destructive calls, with an example. Without that, the opener's fix only works when the dangerous tool was never granted.
2. Show the actual observation and rejection payloads the model sees. Turn guidelines 4 and 6 into code.
3. Cut the Toolformer section to two sentences and move "learned to ask" up to lead it.

## Miriah's notes

<!-- Add your notes for the rewrite here. Astra reads this section. -->
