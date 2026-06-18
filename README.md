# HERALD
 
**A security-first AI gateway built for regulated enterprise environments.**
 
Enterprise networks restrict direct outbound access from internal workloads to external APIs. HERALD solves this by acting as a controlled egress point, forwarding AI API traffic through existing enterprise proxy chains allowlisted, observable, and auditable.
 
HERALD runs anywhere your workload runs. Kubernetes, EC2, ECS, bare metal, or Docker Compose. The deployment model is yours to choose.
 
---
 
## How it works
 
```
Agent → HERALD → Enterprise egress gateway → External AI API
```
 
The agent points its API client at HERALD's URL. HERALD rewrites the host header and forwards the request through your enterprise egress chain. No client code changes required only the endpoint URL changes.
 
---
## V1 — Forward Proxy
 
V1 is deliberately minimal: host rewrite and passthrough. One job, done correctly.
 
### Environment variables
 
| Variable | Required | Default | Description |
|---|---|---|---|
| `HERALD_UPSTREAM_URL` | Yes | — | Enterprise egress gateway URL |
| `HERALD_LISTEN_ADDR` | No | `:8080` | Address HERALD binds to |
| `HERALD_LOG_LEVEL` | No | `info` | Log verbosity: debug / info / warn / error |
 
### Configuring your API client
 
```bash
# Before — client calls the external API directly
API_ENDPOINT=https://external-api.provider.com
 
# After — client calls HERALD, HERALD forwards through your egress chain
API_ENDPOINT=http://herald.your-host:8080
```
 
---
## Vesion for Herald
 
### V2 — Security control plane
- Allowlist enforcement —> only permitted upstream domains pass
- Header injection —> add enterprise auth and tracing headers per upstream
- Credential rotation without restarts
- mTLS support
- Request signing for upstream verification
### V3 — AI security gateway
- Pluggable pre/post scan hooks —> integrate any AI security scanning control
- Agent identity and attribution —> per-agent audit trail
- Token budget enforcement —> hard limits per agent per session
- Rate limiting by agent identity, user, and model
- OPA policy integration for declarative access control
- Structured audit log export to SIEM
---
 
## Project
 
Part of the [c0d3x-io](https://github.com/c0d3x-io) open source security tooling organisation.

---
## Contributing
 
Issues and pull requests are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.
 
---
 
*v0.1.0 — Forward proxy. Runs anywhere.*
