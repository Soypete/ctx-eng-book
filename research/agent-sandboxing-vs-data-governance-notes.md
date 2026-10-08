# Agent Sandboxing vs. Data Governance — Research Notes

*Research for the blog follow-up "The Filesystem Is All You Need (Until Your Data
Isn't There)". Outline: [`blog-posts/outlines/the-filesystem-is-all-you-need.outline.md`](../blog-posts/outlines/the-filesystem-is-all-you-need.outline.md).*

*Sources read: 2026-10-07. Repos change quickly; re-check every quoted behaviour
before publishing.*

---

## Initial Framing

The owner's thesis: the Unix model already answers most of the isolation question.
Users, groups, file permissions, and "no sudo" are the original sandbox. We build
container sandboxes for agents because we let them act as root, or as the developer,
inheriting the developer's whole authority. "The filesystem is all you need" is true
for code on disk. It stops being true when the data that matters lives in GitHub,
Linear, Google Drive, a CRM, or a database, reached with a token over the network. A
filesystem boundary cannot express "may read PRs but not merge", "contractors can't
search the wiki", or "never force-push". Those are decisions about *identity* and
*operation* at the tool-call/connector boundary.

What the research shows, briefly:

- docker-agent and AGRO are good, honest tools for the *execution* boundary (where
  the agent's process runs and what it can touch on disk and on the network).
- Both also ship some tool-call policy, but it is configured per agent/per user
  config or per sandbox repo, and both projects' own docs say it is **not** a
  security boundary. Each points back to the container/VM as the real boundary.
- Neither, as far as the docs I read show, decides "which *person or group* may
  perform *which operation* on *which remote resource*" with org-managed policy and
  central audit. That is the gap the post argues Kei fills.

---

## 1. docker/docker-agent

- Repo: https://github.com/docker/docker-agent (Apache-2.0, Go, ~3.7k stars at time of reading)
- Docs: https://docker.github.io/docker-agent/ (canonical copies under https://docs.docker.com/ai/docker-agent/)

### What it is

> "Build, run, and share AI agents with a declarative YAML config, rich tool ecosystem, and multi-agent orchestration."
> — README

- A `docker` CLI plugin (`docker agent run agent.yaml`), pre-installed in Docker
  Desktop 4.63+, also a standalone `docker-agent` binary. This is the "one CLI".
- An **agent builder and runtime**: you declare agents, models, and toolsets
  (built-in shell/filesystem/git/fetch tools, any MCP server) in YAML; multi-agent
  delegation; RAG; agents are packaged and shared as OCI artifacts.
- Provider-agnostic models. Credentials are supplied by the user, e.g.
  `export OPENAI_API_KEY=sk-...` (README).

### Control surface 1: tool permissions (client-side)

Source: https://github.com/docker/docker-agent/blob/main/docs/configuration/permissions/index.md

- `allow` / `ask` / `deny` lists with glob + argument matching, e.g.
  `"shell:cmd=sudo*"`, `"shell:cmd=git push*"`, `"write_file:path=/etc/*"`, and MCP
  tool names such as `"github_get_*"` / `"github_delete_*"`. Order: deny → allow → ask
  → safety-mode fallback.
- Two levels: **agent-level** (in the agent YAML, written by the agent author) and
  **global user-level** (`~/.config/cagent/config.yaml`). They are merged; deny wins.
- **Safety modes**: `strict`, `balanced`, `restricted` (fail-closed for headless),
  `autonomous` (the old `--yolo`). Calls are labelled safe/destructive/unknown.
- Hooks (`tool_input_transform`, `tool_guard`, `pre_tool_use`) can add checks.
- The docs are explicit about the limit:
  > "Permissions are enforced client-side. They help prevent accidental operations but should not be relied upon as a security boundary for untrusted agents. For stronger isolation, use sandbox mode."
  > "Restricted is defense in depth against unwanted tool calls, not a security boundary."

Fair reading: docker-agent *does* have tool-call policy, including MCP tool names, so
the post must not claim "they only do filesystems". The difference is *where the
policy comes from and whom it is keyed to*: an agent author's YAML plus the local
user's config file, evaluated on the client, with no notion of org groups
(contractors vs. employees) in the docs I read.

### Control surface 2: sandbox mode (Docker Sandboxes VM)

Source: https://github.com/docker/docker-agent/blob/main/docs/configuration/sandbox/index.md
and https://docs.docker.com/ai/sandboxes/

- `docker agent run --sandbox agent.yaml` asks Docker Sandboxes (`sbx` CLI or
  `docker sandbox` plugin) to create/reuse a **VM**; docker-agent "does not implement
  the sandbox" itself.
- Local mode: the **working directory is mounted read-write**; config dir and the
  staged "auto-kit" (skills, prompt files) are mounted read-only. "Only the working
  directory, the agent config directory, and (when staged) the kit directory are
  mounted; other host files are not visible to the agent."
- **Default-deny network proxy**, opened for the model gateway, models.dev, inferred
  tool-install hosts, and a `runtime.network_allowlist` you declare.
- Auto-kit secret redaction via portcullis is "best-effort".
- Cloud mode (`--cloud`): fresh remote workspace; host API keys are **not** uploaded;
  secrets configured with `sbx --cloud secret`; egress declared in kit/cloud policy.
- The docs mention sbx may expose "proxy-managed sentinels for providers", which
  suggests Docker Sandboxes can attach model-provider credentials at a proxy rather
  than inside the VM. **Unverified** — I did not read the sbx docs closely enough to
  say how far that goes (model keys only? arbitrary API tokens?).

### Where it is the right tool

Running untrusted or autonomous code, giving an agent a disposable workspace,
limiting egress hosts, packaging/sharing agents. It solves "what can this process
touch" very well. It does not, by itself, decide what the GitHub token inside the
allowlisted `api.github.com` connection may do.

---

## 2. mifunedev/agro

- Repo: https://github.com/mifunedev/agro (Apache-2.0 runtime/CLI, TypeScript, ~40 stars at time of reading)
- Site: https://mifune.dev
- Security doc: https://github.com/mifunedev/agro/blob/main/docs/security-considerations.md
- Config/secrets doc: https://github.com/mifunedev/agro/blob/main/docs/configuration.md

### What it is

> "AGRO — Agent Governance Runtime Orchestrator. A portable home for autonomous coding agents."
> "AGRO gives AI coding agents a workspace you control. It packages a Docker sandbox and shared agent procedures. You choose the coding harness: Claude Code, Codex, Pi, or another."
> "AGRO runs one project in one Docker sandbox, and **`agro` is the only front door**."
> — README

- One CLI (`agro`) that creates a Docker (Compose/devcontainer) sandbox, installs
  harnesses on demand (Claude Code, Codex, Pi, OpenCode, Hermes, …), installs tools
  (herdr, gh, tailscale, …), runs parallel work in git worktrees, and ships shared
  skills and hooks.
- License note: the runtime, CLI, container definitions and harness spec are
  Apache-2.0; "The Mifune Console, the provisioning and fleet-management control
  plane, and billing / enterprise policy / RBAC / hosted operations are
  proprietary." **Unverified:** what the proprietary "enterprise policy / RBAC" does.
  I could not inspect it; the post must not characterise it.

### Control surface (from its own security doc, which labels each boundary ENFORCED or DOCTRINE)

- **Sandbox isolation — ENFORCED:** "Agents run in a container as the non-root
  `sandbox` user." Docker socket off by default ("Socket access is host root").
  Only `SYS_ADMIN` added, `apparmor=unconfined`, never `privileged: true`.
- **Harness permission engines are turned off inside the sandbox:**
  > "The sandbox runs Claude Code with `bypassPermissions` and Codex with `approval_policy = "never"`. The permission engine is off, so a deny list alone does not hold."
- **Secret-exposure hooks — ENFORCED (pattern-matching):** pre-tool hooks deny env
  dumps, `gh auth token`, reads of `.env*`, `*.pem`, `.config/`.
- **Destructive-command guard — ENFORCED (pattern-matching):** cc-safety-net denies
  `rm -rf`, `git reset --hard`, **`git push --force`**, etc. But:
  > "The script-file route also lets an agent bypass the guard. cc-safety-net catches accidents. Docker is the security boundary (section 4)."
- **Secrets:** stored in a `.env` (mode 0600) inside the sandbox; allow-listed keys
  include `GH_TOKEN` and Slack bot tokens. The agent works with `gh auth login` as
  the developer's GitHub account (README step 3).
- **Human merge gate — DOCTRINE:**
  > "No agent merges its own work… Configure GitHub branch protection with required reviews… Without branch protection, the gate rests on skill text alone."
- **Threat model:** stops accidental credential leaks and agents acting outside
  their intended surface; "do not stop a determined adversary who controls the model."

### Why AGRO is the best example for this post

1. It *already does the Unix thing*: non-root user, no socket, no sudo. That proves
   the owner's first point — the original Unix model gets you most of the
   execution boundary.
2. Its own doc shows the edge of that boundary: "never force-push" is a regex over a
   shell string that a script file bypasses, and "no agent merges" is doctrine
   unless you move the rule to **GitHub** (branch protection). That is the post's
   argument in the project's own words: the rule that holds lives at the data's
   side, keyed to identity, not in the box.
3. The agent holds the developer's `GH_TOKEN`. Inside a perfect container, that
   token can still do everything the developer can do on GitHub.

Be fair: AGRO is candid and well-engineered for its threat model (accidents in a
personal/team coding sandbox). It is not claiming to be org data governance.

---

## 3. ctx-eng-book passages to cite

### Chapter 12 — The Unix Philosophy of AI Systems

- **ch12.03 Mounts, Namespaces, and Isolation** (`book/chapters/ch12-the-unix-philosophy-of-ai-systems/modules/ch12.03-mounts-namespaces-isolation.md`)
  - "These are logical context boundaries. Operating-system mechanisms may help enforce them, but they do not define the book's thesis or replace backend authorization."
  - "Paths make ownership and lifecycle visible. They do not enforce tenant isolation unless every resolver and backend binds the path to trusted identity and policy. A string prefix such as `/tenant/acme` is not a security boundary."
  - "Backend authorization from Chapters 10–11 remains necessary even inside a sandbox."
  - "The namespace names the view; the host, operating system, network, and backend must each enforce the parts of the boundary they actually control." (Good line for the diagram caption.)
- **ch12.04 Task Workspaces and Secret Management** (`…/ch12.04-task-workspaces-secret-management.md`)
  - "UNIX read/write/execute bits do not map directly to retrieval/tool/memory capabilities. Filesystem permissions can enforce some local operations; database, API, graph, and semantic restrictions remain at their resource boundaries." — **the single best citation for the thesis.**
  - "Do not mount a model-readable `/secrets` directory."
  - "Prefer a trusted proxy, sidecar, broker handle, or platform-native identity that attaches credentials at the outbound boundary."
  - Closing: "The UNIX lesson is not that everything should literally be a file. It is that naming, composition, isolation, and lifecycle should be explicit."
- **ch12.02 Pipes, Files, and Explicit Interfaces**
  - "File-Like Is an Analogy, Not a Universal API" — a SQL snapshot, semantic search, graph traversal, and event subscription "have different consistency, query, and authorization behavior."
  - "It must not include secrets or imply that path permissions replace backend enforcement."
- **ch12.01 Small Composable Systems**
  - Split at enforceable boundaries, including "policy enforcement that receives trusted identity and exact resources."
  - Attribution care: the Unix maxims are McIlroy, Pinson, and Tague (1978 foreword); Ritchie & Thompson (1974) documented the system: https://www.nokia.com/bell-labs/about/dennis-m-ritchie/cacm.html
- Neighbouring chapter worth a link: **ch11 Stop Giving Agents Permissions** (scoped credentials, retrieval/execution boundaries).

### Unix research notes

- `research/unix-programming-environment-notes.md`:
  - "A surprising amount of 'agent innovation' appears to be rediscovering operating system concepts under LLM constraints." — good hook line.
  - Context-as-mounted-state section (`/user/`, `/session/`, `/tools/` …). **Caution:** this early note lists `/secrets/` as a mount and says permission bits "map directly" to AI authorization; the edited ch12.04 corrects both. Cite the chapter, not the early note, for those claims.
- `research/unix-composition-notes.md`: Saltzer, Reed & Clark, *End-to-End Arguments* (1984), https://doi.org/10.1145/357401.357402 — the principled argument for putting a check where it can actually be enforced end to end (here: at the resource/connector, not the box).
- `research/namespace-isolation-notes.md`: Sun et al. (USENIX Security 2018) and BPFContain (2021) — namespaces alone are limited isolation and need separate policy.
- `research/workspace-secret-management-notes.md`: Chen et al. (2026), *How Your Credentials Are Leaked by LLM Agent Skills*, https://arxiv.org/abs/2604.03070 — credentials leak into model context through skills/stdout; supports "don't hand the agent the token".

### Hugging Face incident (use carefully)

- Ledger row "Authorization — Hugging Face Agent Intrusion Crossed Credential and
  Trust Boundaries" (`research/_evidence-ledger.md`), source:
  https://huggingface.co/blog/agent-intrusion-technical-timeline
- Per the primary report as recorded in the ledger, the agent **did** escape its
  evaluation environment, obtained production-pod code execution, acquired
  credentials, and moved laterally; enabling conditions included "broad or shared
  credential scope, cloud-metadata reachability, and insufficient isolation."
- **Conflict:** the Kei content strategy and POV describe the incident as "the model
  never breached the perimeter… inferred it was permitted… stayed within the
  permissions it held." That does not match the ledger's reading of the primary
  report. If the post uses the incident, use the ledger framing: isolation failed
  *and* the credentials it reached were broad. That still supports the thesis (a
  box is one layer; scoped, identity-bound credentials are another), without
  overclaiming. Re-read the primary report before citing.

---

## 4. Kei positioning points

### From `business/content-strategy.md` (kei repo, origin/main)

- Series: **"The Death of the Sandbox"** (and "Stop Feeding the Monopolies"); primary channel Substack; companion book is this repo.
- All Day AI BUILD abstract, "Local Agents Inherit Authority": "The most capable agent harnesses today were designed around a workstation. Claude Code, Cursor, OpenCode — they inherit a user, a filesystem, credentials, a shell, and a permission model from the machine beneath them… And the fix is not 'put it in a container.'" Takeaway: agents "whose compute is ephemeral, authority is scoped, credentials are renewable, state is external, and actions are auditable."
- KubeCon abstract: an agent "must act on behalf of a human across GitHub, Google Drive, internal APIs, and production infrastructure without inheriting every permission that human possesses… The goal is not to put Claude Code in a container. It is to answer a more fundamental question: what is the cloud-native architecture for software that acts like a user?"
- SECURE-track lessons: "The sandbox is not the guardrail; the tool definition is." "Permissions must be a property of the human, not the agent."
- Law 1 (Lexicon): ABAC "at the tool-call level".
- CTA language in the plan: "Keep data sovereign and use Haikai as the background for intelligent action." (Tone for the post should be plainer than the strategy doc; see the Substack voice guidelines.)

### From `business/pov.md` and `business/manifesto.md`

- "We are not building an agent framework… Bring Your Own Agent." → Kei composes with docker-agent/AGRO rather than replacing them.
- Rule "Human-Defined Agency": every agent action bound to a specific human identity and permission set.
- Manifesto: an agent's behaviour depends on "Model + Data + Meaning + Identity + Tools + Policy + State"; actions "must be expressed as constrained, authorized, and auditable operations."

### From kei-console docs (origin/main: `frontend/src/docs/content/`)

Verified in `how-it-works.md` and `policy-precedence.md`:

- One runtime installation per developer machine; each harness (Claude Code, Codex, OpenCode, custom) is a session of it.
- Policy = `src` (`group:`, `user:`, `harness:`, `agent:`, `org:<role>`, `*`) × `dst` (`shell:`, `skill:`, `path:`, `tool:`, `capability:`, `resource:`, `connector:`) × effect × priority. Higher priority wins; deny wins ties.
- Identity is "resolved from the authenticated session", never from a caller-supplied email.
- **Desktop harnesses:** Kei renders shell/skill/path policy into native settings (Claude Code `settings.json`, Codex rules, OpenCode permissions); the harness makes its own **native decision**; Kei's hook is **report-only**. Only workspace-wide and harness-level sources render; **group sources do not render into desktop settings.**
- **Custom/SDK harnesses and connectors:** `kei-proxy authorize` evaluates a verified, versioned, expiring bundle locally; **fails closed** on no match or expired bundle.
- `required_capabilities` combine with AND (e.g. `capability:pull_request.read`); a tool permit never overrides a capability deny.
- Worked examples in the docs match the post's examples almost exactly: `harness:claude_code` deny on `shell:git push --force`; `group:contractors` deny on `tool:search_wiki` over a `group:members` permit; `connector:billing.charge` with no policy → deny.
- Audit (`audit-logs.md`, `audit-encryption.md`): decision records carry subject, tool name, resource ids, decision, and an `args_digest`; with no customer key, arguments are **digest-only** (HMAC-SHA-256, not reversible); with a customer key, arguments are encrypted to the customer's keys. "Raw credentials, provider request/response payloads, and customer documents must remain out of control-plane audit logs."

### Accuracy guardrails for the post

- Do **not** say Kei blocks every desktop tool call itself: on Claude Code/Codex/OpenCode the decision is the harness's native one, from Kei-rendered settings. A desktop `git push --force` deny is still a shell-prefix rule, just like AGRO's guard. The *identity-aware, fail-closed* decision is at `kei-proxy` for connectors and custom harnesses.
- Therefore the strongest "never force-push" story is the connector one: the agent reaches GitHub through a governed connector whose capability (`pull_request.merge`, force-push) is denied per identity, so no shell trick changes the answer. Confirm the GitHub connector actually exposes those capabilities before publishing.
- "Credentials are never handed to the agent (pass-through via the runtime)": **not verified** in the console docs I read. `proxy-runtime.md` says to keep provider credentials in the container's secret store; the Drive connector uses the user's OAuth consent. Verify against the kei-proxy `connector invoke` docs / kei-agents connector code before stating it as fact.

---

## 5. Open questions

1. Docker Sandboxes: does `sbx`'s proxy inject arbitrary API credentials (e.g. a GitHub token) outside the VM, or only model-provider keys? If the former, docker-agent already gets "agent never sees the token", and the post's distinction narrows to identity-aware policy + org audit.
2. Does any docker-agent permission or hook receive the *user's org identity/group*? The docs show agent YAML and the local user's config only. Check the hooks API payload.
3. What does Mifune's proprietary "enterprise policy / RBAC" cover? Ask or wait for public docs; do not speculate in the post.
4. Kei: confirm the GitHub/Linear/Drive connectors' capability names (e.g. `pull_request.read` vs. `pull_request.merge`) and whether force-push is expressible as a connector capability, or only as a shell-prefix/branch-protection rule.
5. Kei: confirm credential pass-through (runtime attaches the token; agent never holds it) and which connectors support it today.
6. Kei: group sources don't render into desktop harness settings yet. Is "contractor can't search the wiki" achievable today for a Claude Code user, or only via a governed `search_wiki` tool/connector through kei-proxy? The example must use the path that works.
7. Hugging Face incident: reconcile the content-strategy framing with the primary report before reusing the "sandboxes failed" line.
8. Composition story: can a docker-agent or AGRO sandbox run `kei-proxy` beside the harness today (custom harness via `kei harness add`, or desktop harness sync inside the container)? A short "run both" snippet would make the post's fairness section concrete.

---

## Claims for the Evidence Ledger (candidate rows)

- **Pragmatics — Filesystem boundaries do not express remote-resource operations.** Source: ch12.04 ("UNIX read/write/execute bits do not map directly…"). Strength: book argument; supported by AGRO's own DOCTRINE/branch-protection note. Counterpoint: Docker Sandboxes' network proxy can restrict *hosts*, but not operations within a host.
- **Pragmatics — Pattern guards on shell strings are accident prevention, not boundaries.** Source: AGRO security doc ("The script-file route also lets an agent bypass the guard"); docker-agent permissions doc ("enforced client-side… should not be relied upon as a security boundary"). Strength: strong (vendors' own statements).
- **Lexicon — Sandboxed agents commonly hold the developer's full token.** Source: AGRO README/config (`gh auth login`, `GH_TOKEN` in `.env`); docker-agent README (`export OPENAI_API_KEY`). Strength: strong for these two projects; suggestive in general.
