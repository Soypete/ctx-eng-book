# The Filesystem Is All You Need (Until Your Data Isn't There) — Outline

*Status: outline only, not drafted. Follow-up post in the "Death of the Sandbox"
series. Research: [`research/agent-sandboxing-vs-data-governance-notes.md`](../../research/agent-sandboxing-vs-data-governance-notes.md).*

*Target: Substack, ~1,200–1,800 words, voice per `.opencode/skills/substack-prep/references/post-guidelines.md`
and `blog-posts/how-everyone-is-using-ai-wrong.md`. Needs one diagram, one short code
snippet (policy rows or YAML), and a `## Guidelines` section.*

---

## Title options

1. **The Filesystem Is All You Need (Until Your Data Isn't There)** — working title; plays on the book's ch03 "Attention Is All You Need" and the Unix talk.
2. You Don't Need a Sandbox. You Need to Stop Running Agents as Root.
3. Your Agent Doesn't Live in a Container. Your Data Doesn't Either.
4. The Box Is the Wrong Boundary: Sandboxes, Unix, and Where Agent Policy Belongs
5. `chmod` Can't Say "Read PRs, Don't Merge"

Recommendation: 1 for the series, 2 as the subtitle/social line.

## Argument (one sentence)

Unix users and file permissions already solve isolation for code on disk, which is
what container sandboxes like docker-agent and AGRO recreate; but an organisation's
data lives behind tokens in GitHub, Linear, Drive and SaaS APIs, so the boundary
that matters is an identity-aware decision at each tool call and connector, which a
box cannot make.

## Thesis (3 lines)

- Sandboxes exist because we let agents act as root and inherit the developer's whole authority; the Unix model (users, groups, no sudo) was already the sandbox.
- "The filesystem is all you need" holds for code on disk, and that's the layer docker-agent and AGRO govern well.
- The data that matters isn't in the filesystem; "read PRs but don't merge", "contractors can't search the wiki" and "never force-push" are decisions about identity and operation at the tool-call/connector boundary, which is where Kei sits, alongside the sandbox, not instead of it.

## Reader takeaway (conclusion to reach)

Keep your sandbox, drop sudo, and then ask the question the box can't answer: *who
is this agent acting for, and what may they do to this resource?* Put that decision
at the call, keyed to identity, with the credential out of the agent's hands and an
audit row for every call.

---

## Section-by-section outline

### 1. Hook (~120 words)

- Open on a concrete scene: an agent in a properly locked-down container (non-root
  user, no Docker socket, default-deny egress with `api.github.com` allowlisted) runs
  `git push --force` to `main` with the developer's token. Every sandbox control
  held. The damage happened anyway, on GitHub.
- One-line stance: the box did its job. It was the wrong box.
- Example option (scene illustration, optional): the container drawn as a fortress
  with a token-shaped key passing straight through the network allowlist.

### 2. We already had a sandbox; it's called a Unix user (~250 words)

- Unix answered isolation in the 1970s: users, groups, rwx bits, process
  separation, and a root account you don't hand out. Cite Ritchie & Thompson (1974)
  https://www.nokia.com/bell-labs/about/dennis-m-ritchie/cacm.html; attribute the
  maxims to McIlroy, Pinson & Tague (1978).
- The owner's line: *we wouldn't need sandboxes if agents didn't have sudo.* We gave
  workstation harnesses the developer's user, shell, filesystem, credentials and
  permission model (the "Local Agents Inherit Authority" framing), then bolted a
  container around them to claw that back.
- Hook line from the Unix notes: "A surprising amount of 'agent innovation' appears
  to be rediscovering operating system concepts under LLM constraints."
- Evidence the industry is converging on this: AGRO runs agents "as the non-root
  `sandbox` user" with the Docker socket off by default ("Socket access is host
  root"). That's Unix hygiene, packaged.
- Define "harness" in one sentence here (the program that runs the model loop and
  executes its tool calls: Claude Code, Codex, OpenCode).

### 3. What docker-agent and AGRO actually do (be fair) (~300 words)

- **docker-agent** (https://github.com/docker/docker-agent): one CLI (`docker agent
  run`) to define agents in YAML, wire tools and MCP servers, share agents as OCI
  artifacts. Two controls: client-side `allow/ask/deny` tool permissions (globs such
  as `shell:cmd=sudo*`, `github_delete_*`) with safety modes; and `--sandbox`, which
  runs the agent in a Docker Sandboxes VM with only the working directory mounted
  and a default-deny network proxy. Its own docs: permissions "should not be relied
  upon as a security boundary… use sandbox mode."
- **AGRO** (https://github.com/mifunedev/agro): "`agro` is the only front door." One
  Docker sandbox per project, bring-your-own harness, git worktrees for parallel
  agents, shared skills and hooks. Its security doc is admirably candid: harness
  permission engines are off inside the box; hooks pattern-match secrets and
  destructive commands (including `git push --force`); "Docker is the security
  boundary."
- Where they're the right tool: running autonomous or untrusted code, disposable
  workspaces, egress control, reproducible agent environments. Keep using them.
- Show, don't sneer. Name what they're for before naming what they can't do.

### 4. The filesystem is all you need… (~150 words)

- For code on disk, the filesystem boundary is genuinely sufficient: mount the repo,
  hide `$HOME`, drop root, deny egress. Agents writing code in a worktree are
  well served.
- This is the part of the title that's true. Grant it fully before the turn.

### 5. …until your data isn't there (~350 words) — the core turn

- The data an organisation cares about isn't in the container: PRs and branch
  history in GitHub, tickets in Linear, docs in Google Drive, accounts in the CRM,
  rows in the warehouse. The agent reaches it with a **token over an allowed
  network connection**.
- A filesystem or network sandbox can say *which hosts*. It can't say *which
  operations, on which resources, for which person*. Three examples:
  1. **`git push --force`**: AGRO blocks it with a shell-pattern guard and says
     itself that "the script-file route also lets an agent bypass the guard." A
     regex over a command string is accident prevention. The rule that holds is on
     GitHub's side (AGRO recommends branch protection for exactly this reason) or at
     a governed connector that never exposes force-push to this identity.
  2. **Reading a Drive doc**: the sandbox allowlists `www.googleapis.com`; from then
     on, whatever the OAuth token can read, the agent can read. "This agent may
     summarise the onboarding doc, not the board deck" is a policy on a resource,
     not a path.
  3. **A contractor and the wiki**: same harness, same container image, two humans.
     The employee may search the wiki; the contractor may not. The box is identical,
     so the box can't tell them apart. Only identity can.
- Book citation (the key line): "UNIX read/write/execute bits do not map directly to
  retrieval/tool/memory capabilities. Filesystem permissions can enforce some local
  operations; database, API, graph, and semantic restrictions remain at their
  resource boundaries." (ch12.04)
- Plus ch12.03: "Backend authorization… remains necessary even inside a sandbox."
- Optional evidence beat (only after re-reading the primary report): the July 2026
  Hugging Face intrusion, where broad/shared credential scope was among the
  enabling conditions alongside insufficient isolation. Use the ledger framing, not
  the content-strategy framing (see research note §3).

### 6. What we do differently, and why not at the filesystem level (~350 words)

Four differences, each one sentence + one example:

1. **Policy on identity, not on the box.** A Kei policy is `src × dst × effect ×
   priority`: `src` is a user, a group, or a harness; `dst` is a shell prefix, a
   tool, a capability, a resource, or a connector. Short snippet (≤15 lines), taken
   from Kei's own precedence docs:

   ```text
   src                   dst                       effect  priority
   group:members         tool:search_wiki          permit  10
   group:contractors     tool:search_wiki          deny    100
   harness:claude_code   shell:git push --force    deny    100
   *                     connector:billing.charge  (none)  -> deny, fail closed
   ```
2. **The decision happens at the call.** For connectors and custom harnesses,
   `kei-proxy authorize` evaluates a signed, expiring policy bundle locally and
   **fails closed** when nothing matches. Desktop harnesses (Claude Code, Codex,
   OpenCode) keep their native permission prompts; Kei renders policy into them
   rather than replacing them. (Accuracy: say this plainly; don't imply Kei
   intercepts every desktop call.)
3. **Credentials stay out of the agent.** The runtime attaches the credential at
   the outbound boundary, the "handles, not secrets" pattern from ch12.04.
   (**Verify before publishing**; see open question 5 in the research note.)
4. **Every call is audited.** Who, which harness, which tool/resource, which policy
   decided, with arguments stored as a keyed digest (or encrypted to your keys),
   never raw credentials or documents.

- Why not at the filesystem level: because the policy has to travel with the
  *person* across every harness and every machine, and has to be about operations
  on remote resources. A mount table can't carry "contractor" and can't see "merge".

### 7. Diagram (described; build per post-guidelines)

**Diagram 1, "Two boundaries, two questions."** Left-to-right:

- Developer identity (user, groups) → Harness (Claude Code / docker-agent / AGRO
  sandbox) drawn inside a dashed box labelled **"Sandbox: where can this process
  run and what can it touch?"** (filesystem mounts, non-root user, egress allowlist).
- From the harness, two arrows leave the box:
  - to local files (stays inside the dashed box, governed by the sandbox);
  - to a **Kei decision point** labelled **"Policy: who is this for, and may they do
    this to that resource?"** (identity + tool/capability + resource → permit/deny,
    fail closed, audit row), then on to GitHub / Linear / Drive / API.
- A denied path (contractor → `search_wiki`) ends at "deny + audit".
- Caption: *The sandbox governs the process; the policy governs the data. You need
  both, and they answer different questions.*
- Source as `book/substack/<slug>/diagram-1.mmd` with the SoyPeteTech Mermaid
  classes; export PNG; embed via raw GitHub URL.

### 8. How they compose (~150 words)

- Run the agent in docker-agent's sandbox or AGRO's container **and** route its
  tool calls and connectors through Kei. Box for the process; identity-aware policy
  for the data. Neither replaces the other (ch12.03: each layer "must enforce the
  parts of the boundary they actually control").
- Bring-your-own-agent framing from the POV: Kei isn't an agent framework and
  doesn't compete with either project's runtime.
- If verified (open question 8), show the two-line "run both" setup.

### 9. Guidelines (5 rules)

1. Run agents as a non-root user with no sudo and no Docker socket before you reach
   for anything fancier.
2. Keep the sandbox for code and execution; don't count it as your data policy.
3. Write policy against people and groups, not against containers or paths.
4. Decide at the tool call: deny by default when no rule matches.
5. Don't hand the agent the token; attach credentials at the outbound boundary, and
   audit every call without logging the secret.

### 10. Close + CTA (~100 words)

- Takeaway: the Unix people were right, and we should finish the job. Drop root,
  keep the box, and then govern what the box can't see.
- Series pointer: next in "The Death of the Sandbox" series (and the book's
  Chapter 11, "Stop Giving Agents Permissions": https://github.com/Soypete/ctx-eng-book).
- CTA consistent with the content plan: subscribe for the series; read Kei's "How
  Kei governs coding agents" docs and try a policy that denies a contractor the
  wiki; invite readers already running docker-agent or AGRO to tell us how they'd
  compose them with identity-aware policy. Keep it plain; no strategy-doc
  language ("Predatory Paradigm", "Sage").

---

## Examples needed

- Policy-table snippet (above), lifted from Kei's policy-precedence examples.
- AGRO quote block (security doc, sections 3 and 6) for the force-push example.
- docker-agent permissions YAML excerpt (`deny: ["shell:cmd=sudo*"]`) to show that
  their policy is per agent/per user config, client-side.
- Diagram 1.

## Readiness

needs-research (open questions 1, 4, 5, 6, 8 in the research note must be closed
before drafting sections 6 and 8; question 7 before using the Hugging Face beat).
Everything else is ready to draft.

## Don'ts (for the drafter)

- Don't claim docker-agent or AGRO "only do filesystems": both have tool-call
  policy; the difference is whose identity it is keyed to and where it's enforced.
- Don't characterise Mifune's proprietary enterprise policy/RBAC.
- No internal ticket numbers, customer data, or internal hostnames.
