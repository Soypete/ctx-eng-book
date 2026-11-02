# substack-draft chart

Drafts the post `book/substack/SCHEDULE.md` dates for the day it runs, then pings Discord.
Mon/Wed/Fri 07:00 America/Denver by default.

Secrets (not in git):
- **Substack cookie** comes from OpenBao on refurb (`secret/apps/substack`, key `token`): the
  `substack.sid` cookie value. The OpenBao agent injector renders it before the job starts,
  using kubernetes auth role `ctx-eng-book-substack-draft` (bound to service account
  `substack-draft` in `ctx-eng-book`, policy `substack-drafter`). A 401/403 in the job log
  means the cookie expired; refresh it with
  `VAULT_ADDR=http://refurb:8200 bao kv put secret/apps/substack token=<substack.sid>`.
- `ctx-eng-book-secrets` holds `DISCORD_SECRET` and `DISCORD_USER_ID`.

One-time OpenBao setup (already done on refurb):

    bao write auth/kubernetes/role/ctx-eng-book-substack-draft \
      bound_service_account_names=substack-draft bound_service_account_namespaces=ctx-eng-book \
      token_policies=substack-drafter token_ttl=15m

    helm upgrade --install substack-draft deploy/helm/substack-draft -n ctx-eng-book

The live CronJob was first applied with `kubectl apply`. To let Helm adopt it, either delete it
first or label/annotate it (`app.kubernetes.io/managed-by=Helm`,
`meta.helm.sh/release-name=substack-draft`, `meta.helm.sh/release-namespace=ctx-eng-book`).
