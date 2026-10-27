# substack-draft chart

Drafts the post `book/substack/SCHEDULE.md` dates for the day it runs, then pings Discord.
Mon/Wed/Fri 07:00 America/Denver by default.

Secret (not in git): `ctx-eng-book-secrets` with `SUBSTACK_SESSION_COOKIE` (the `substack.sid`
cookie, URL-encoded, starts `s%3A`), `DISCORD_SECRET`, `DISCORD_USER_ID`. A 403 in the job log
means the cookie expired; refresh that key.

    helm upgrade --install substack-draft deploy/helm/substack-draft -n ctx-eng-book

The live CronJob was first applied with `kubectl apply`. To let Helm adopt it, either delete it
first or label/annotate it (`app.kubernetes.io/managed-by=Helm`,
`meta.helm.sh/release-name=substack-draft`, `meta.helm.sh/release-namespace=ctx-eng-book`).
