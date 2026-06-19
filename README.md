# HERALD

A forward proxy for routing AI API traffic through enterprise egress controls.

Early-stage project. V1 does one thing: rewrite the host and forward the request through your enterprise gateway, over TLS.

---

## What it does

```
Agent → HERALD (TLS) → Enterprise egress gateway → External AI API
```

Your agent points its API client at HERALD instead of the external API directly. HERALD rewrites the host header and forwards the request. No client code changes — only the endpoint URL changes.

---

## Status

V1 — forward proxy, host rewrite, TLS passthrough. Tested locally with a self-signed cert against a real upstream API.

This is a learning project as much as a tool. The code will be rewritten as understanding deepens, not just patched around.

---

## Environment variables

| Variable | Required | Default | Description |
|---|---|---|---|
| `HERALD_UPSTREAM_URL` | Yes | — | The upstream API HERALD forwards requests to |
| `HERALD_LISTEN_ADDR` | No | `:8080` | Address HERALD binds to |
| `HERALD_LOG_LEVEL` | No | `info` | Log verbosity |

---

## Running locally

```bash
export HERALD_UPSTREAM_URL=https://your-upstream-api.com
go run ./cmd/herald
```

---

## Ideas for later

Nothing here is committed. Possible directions if the project continues:

- Allowlist enforcement on upstream domains
- Header injection for enterprise auth
- Pluggable scanning hooks for AI-specific security controls
- Rate limiting and audit logging

---

## Project

Part of [c0d3x-io](https://github.com/c0d3x-io).

---

*v0.1.0 — early, working, incomplete.*
