// Package personalize is the runnable companion to the post
// "Your Model Doesn't Know Your Users. Your Retrieval Does." (book module 1.4).
//
// Personalization here is context assembly at request time: the server
// resolves who the user is from their session, issues a capability scoped to
// the task, and the preference store enforces that capability. The model only
// ever sees the result. Each of the post's four failures maps to one check:
//
//   - wrong user:    the store filters by the capability's subject, which comes
//     from the session, never from a request field
//   - over-scoped:   the store returns only the fields the task's policy grants
//   - stale:         preferences older than their field's max age are excluded
//   - missing/wrong: a declared preference overrides an inferred one
package personalize

import (
	"errors"
	"sort"
	"time"
)

var (
	ErrUnauthenticated = errors.New("unknown or expired session")
	ErrUnknownTask     = errors.New("no personalization policy for task")
	ErrExpired         = errors.New("capability expired")
)

// Principal is an authenticated user. Only Sessions.Authenticate makes one.
type Principal struct{ userID string }

// Sessions maps session tokens to user IDs (stand-in for your auth service).
type Sessions map[string]string

// Authenticate derives identity from the session, not from anything the
// caller or the model claims.
func (s Sessions) Authenticate(token string) (Principal, error) {
	id, ok := s[token]
	if !ok {
		return Principal{}, ErrUnauthenticated
	}
	return Principal{userID: id}, nil
}

// Capability is a task-scoped grant: whose data, which fields, until when.
// Its fields are unexported, so code outside this package cannot forge one;
// in a distributed system you would sign it instead.
type Capability struct {
	subject string
	fields  map[string]bool
	expires time.Time
}

// TaskPolicy is server policy: which preference fields each task may use.
type TaskPolicy map[string][]string

// Issue grants the principal access to exactly the fields the task needs.
func (p TaskPolicy) Issue(who Principal, task string, now time.Time, ttl time.Duration) (Capability, error) {
	fields, ok := p[task]
	if !ok {
		return Capability{}, ErrUnknownTask
	}
	c := Capability{subject: who.userID, fields: map[string]bool{}, expires: now.Add(ttl)}
	for _, f := range fields {
		c.fields[f] = true
	}
	return c, nil
}

// Origin is where a preference came from. They carry different authority.
type Origin string

const (
	Declared Origin = "declared" // the user said so
	Inferred Origin = "inferred" // derived from behavior
)

// Preference is one stored row, with provenance and a timestamp.
type Preference struct {
	Owner     string
	Field     string
	Value     string
	Origin    Origin
	UpdatedAt time.Time
}

// Store holds every user's preferences in one table, like most real systems.
type Store struct{ Rows []Preference }

// Read is the enforcement point. It does not trust the caller to filter.
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

// Exclusion records a preference that was found but left out, and why.
type Exclusion struct {
	Field, Value string
	Reason       string
}

// Context is what reaches the model, plus the evidence of how it was built.
type Context struct {
	Prefs    map[string]Preference
	Excluded []Exclusion
}

// Assemble resolves one value per field: drop rows older than the field's
// max age, then let a declared preference beat an inferred one, newest first.
func Assemble(rows []Preference, maxAge map[string]time.Duration, now time.Time) Context {
	ctx := Context{Prefs: map[string]Preference{}}
	sort.Slice(rows, func(i, j int) bool { return rows[i].UpdatedAt.After(rows[j].UpdatedAt) })
	for _, r := range rows {
		if limit, ok := maxAge[r.Field]; ok && now.Sub(r.UpdatedAt) > limit {
			ctx.Excluded = append(ctx.Excluded, Exclusion{r.Field, r.Value, "stale: refresh before use"})
			continue
		}
		cur, ok := ctx.Prefs[r.Field]
		if !ok || (cur.Origin == Inferred && r.Origin == Declared) {
			if ok {
				ctx.Excluded = append(ctx.Excluded, Exclusion{cur.Field, cur.Value, "overridden by declared preference"})
			}
			ctx.Prefs[r.Field] = r
			continue
		}
		ctx.Excluded = append(ctx.Excluded, Exclusion{r.Field, r.Value, "superseded"})
	}
	return ctx
}

// ForRequest is the whole path a handler calls: session → principal →
// task-scoped capability → enforced read → assembled context.
func ForRequest(s Sessions, p TaskPolicy, st *Store, maxAge map[string]time.Duration, token, task string, now time.Time) (Context, error) {
	who, err := s.Authenticate(token)
	if err != nil {
		return Context{}, err
	}
	c, err := p.Issue(who, task, now, 30*time.Minute)
	if err != nil {
		return Context{}, err
	}
	rows, err := st.Read(c, now)
	if err != nil {
		return Context{}, err
	}
	return Assemble(rows, maxAge, now), nil
}
