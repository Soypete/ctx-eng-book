# Owned source boundaries

A standard-library Go example of an executable boundary contract. Each check has
an owner and policy version. Configuration without ownership is rejected; rejected
source evidence never reaches the model substitute. Decisions retain the boundary,
owner, reason, source reference, source version, and observation/check times.

Run from this directory (Go 1.22+):

```sh
go run .
go test ./...
go vet ./...
```

The demo prints a stale-source rejection and `ModelCalled: false`. Tests cover
current, stale, missing, and future source metadata, plus invalid configuration.
The one-hour maximum age is an invented fixture policy, not a recommendation.

Inputs are assumed already authorized. ObservedAt must come from a trusted source
adapter and means when the source vouched for the value. The program cannot verify
that assertion or that an owner label resolves to a real team. It implements no
retrieval, model API, refresh, durable audit storage, or incident routing. In a
service, govern record access/retention and resolve owners through a maintained
team directory. Unit tests prove these checks, not model accuracy or team response.
