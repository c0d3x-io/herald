# HERALD

A TLS forward proxy for routing AI API traffic through enterprise egress controls.

HERALD sits between your agent and the public internet. Your agent calls HERALD instead of the external API directly; HERALD terminates TLS, validates the certificate chain, rewrites the host, and forwards the request through your egress path.

---

## How it works

```
Agent → HERALD (TLS) → Enterprise egress gateway → External AI API
```

Only the endpoint URL in your agent's config changes — no client code changes required.

---

## Features

- TLS 1.2 minimum enforced on all connections — no plaintext HTTP mode
- Real certificate validation — no skip-verify shortcuts
- Configurable CA trust via `HERALD_CA_BUNDLE` — accepts a single file or a directory of `.pem`/`.crt` files, for local development or private/internal CAs
- Falls back to the system trust store when no custom CA bundle is set — correct default for hitting public upstream APIs
- Structured JSON logging
- Server-level timeouts (read/write/idle) configured by default
- `/health` endpoint for liveness checks
- Configuration entirely via environment variables — no config files

---

## Environment variables

```bash
# .env
HERALD_UPSTREAM_URL=
HERALD_LISTEN_ADDR=
HERALD_LOG_LEVEL=
HERALD_TLS_KEY=
HERALD_TLS_CERT=
HERALD_CA_BUNDLE=
```

| Variable | Required | Default | Description |
|---|---|---|---|
| `HERALD_UPSTREAM_URL` | Yes | — | The upstream API HERALD forwards requests to |
| `HERALD_LISTEN_ADDR` | No | `:8080` | Address HERALD binds to |
| `HERALD_LOG_LEVEL` | No | `info` | Log verbosity: debug / info / warn / error |
| `HERALD_TLS_CERT` | Yes | — | Path to the TLS certificate HERALD serves |
| `HERALD_TLS_KEY` | Yes | — | Path to the TLS private key HERALD serves |
| `HERALD_CA_BUNDLE` | No | — | Path to a custom CA file or directory. **Replaces** the system trust store when set — see note below. |

> **`HERALD_CA_BUNDLE` behaviour:** when set, this replaces the system trust store rather than adding to it. Use it for local development (e.g. a `mkcert`-generated CA) or to pin strictly to a private internal CA. Leave it unset in production if your upstream uses a standard publicly-trusted certificate — setting it will break verification against that upstream.

---

## Local setup

### 1. Generate a local TLS certificate

HERALD always serves TLS, including locally. Use [`mkcert`](https://github.com/FiloSottile/mkcert) to generate a certificate your system trusts:

```bash
brew install mkcert
mkcert -install
mkcert localhost 127.0.0.1 ::1
```

### 2. Create your `.env`

```bash
HERALD_UPSTREAM_URL=https://your-upstream-api.com
HERALD_LISTEN_ADDR=:8080
HERALD_LOG_LEVEL=debug
HERALD_TLS_CERT=./caBundle/localhost+2.pem
HERALD_TLS_KEY=./caBundle/localhost+2-key.pem
HERALD_CA_BUNDLE=
```

### 3. Run

```bash
export $(cat .env | xargs) && go run ./herald.go
```

---

## Running with Docker

### Build

```bash
docker build -t herald:v0.1.0 .
```

### .dockerignore

Certs and `.env` must never be copied into the image build context. Confirm you have a `.dockerignore` with at least:

```
caBundle/
.git/
.env
*.pem
*.key
```

### Run — config via env file, certs via volume mount

```bash
docker run \
  --env-file .env \
  -v "$(pwd)/caBundle:/certs:ro" \
  -e HERALD_TLS_CERT=/certs/localhost+2.pem \
  -e HERALD_TLS_KEY=/certs/localhost+2-key.pem \
  -p 8080:8080 \
  herald:v0.1.0
```

Certs are mounted read-only at runtime — never baked into the image — so they can be rotated without a rebuild, and a leaked image layer never exposes a private key.

If you're using `HERALD_CA_BUNDLE`, mount that the same way and point the env var at the in-container path:

```bash
-v "$(pwd)/caBundle:/certs:ro" \
-e HERALD_CA_BUNDLE=/certs
```

### Verify it's running

```bash
curl -k https://localhost:8080/health
```

---

## Status

Early-stage. TLS handling and CA trust are in place. Allowlist enforcement, graceful shutdown, and request size limits are known gaps, not yet built.

---

## Ideas for later

- Allowlist enforcement on upstream domains
- Header injection for enterprise auth
- Pluggable scanning hooks for AI-specific security controls
- Rate limiting and audit logging
- Graceful shutdown on SIGTERM

---

## Project

Part of [c0d3x-io](https://github.com/c0d3x-io).

---

## Contributing

Issues and pull requests are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md).

---

*v0.1.0 — TLS-only forward proxy, configurable trust.*
