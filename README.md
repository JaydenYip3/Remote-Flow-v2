# remote_flow_v2 — backend

Go HTTP API on the standard library. Nothing here but a server that starts and
answers one route — everything else is mine to build.

## Requirements

- Go 1.25+

## Quick start

```bash
make run              # http://localhost:8080
curl localhost:8080/healthz
```

## Layout

```
cmd/api/main.go       # the whole app, for now
```

When a package earns its place, it goes under `internal/` — the Go toolchain
blocks anything outside this module from importing it.

## Build it up

Rough order, each step standing on the last:

1. **Split main** — pull the mux into a `routes()` function, then a `Server`
   type that owns the `*http.Server`.
2. **JSON** — a helper that sets `Content-Type` and encodes a value; make
   `/healthz` return `{"status":"ok"}`.
3. **Config** — read `PORT` and `LOG_LEVEL` from the environment with defaults,
   instead of the hardcoded `:8080`.
4. **Logging** — swap `log` for `log/slog`, text in dev and JSON elsewhere.
5. **Middleware** — a `func(http.Handler) http.Handler` chain: recover from
   panics, tag each request with an id, log one line per request.
6. **Graceful shutdown** — catch SIGINT/SIGTERM, drain in-flight requests with
   `srv.Shutdown(ctx)`.
7. **Tests** — `httptest.NewRecorder` + `httptest.NewRequest` against `routes()`,
   no live port needed.
8. **A real endpoint** — decode a JSON body, validate it, return 400 on bad input.

## Commands

```bash
make run      # go run ./cmd/api
make build    # bin/api
make test     # go test -race ./...
```
