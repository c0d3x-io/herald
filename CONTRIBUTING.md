# Contributing to HERALD

Thanks for taking the time to contribute. HERALD is a security-first project — that standard applies to contributions as much as it applies to the code.

---

## Before you start

Open an issue before writing code. Describe what you want to build and why. This avoids wasted effort and keeps contributions aligned with the roadmap.

For bug fixes, a brief issue description is enough. For new features, explain the problem you are solving — not just the solution you have in mind.

---

## What we are looking for

- Bug fixes with a clear reproduction case
- Improvements to existing functionality with a documented rationale
- New deployment patterns and examples
- Documentation improvements
- Security findings — see Reporting Security Issues below

## What we are not looking for

- Vendor-specific integrations baked into core — HERALD is platform-agnostic
- Features that couple HERALD to a specific AI provider, SDK, or cloud
- Rewrites of working code without a clear technical reason

---

## Development setup

```bash
git clone https://github.com/c0d3x-io/herald.git
cd herald
go mod download
go build ./...
go test ./...
```

Requires Go 1.22 or later.

---

## Making a change

```bash
# Fork the repo and create a branch
git checkout -b your-branch-name

# Make your changes
# Write tests for anything non-trivial
# Verify the build is clean
go build ./...
go test ./...
go vet ./...

# Commit with a clear message
git commit -m "short description of what and why"

# Push and open a pull request
git push origin your-branch-name
```

---

## Pull request expectations

- One change per PR — keep scope tight
- Tests for new behaviour
- No new dependencies without a clear justification — HERALD uses the Go standard library by design
- Clear description of what the PR does and why

PRs that add external dependencies without justification will not be merged. The standard library is sufficient for most of what HERALD needs to do.

---

## Code style

Follow standard Go conventions. Run `gofmt` before committing. Code review will flag anything that departs significantly from idiomatic Go.

---

## Reporting security issues

Do not open a public issue for security vulnerabilities.

Report security findings privately via GitHub's security advisory feature:
`https://github.com/c0d3x-io/herald/security/advisories/new`

Include a clear description of the issue, steps to reproduce, and your assessment of impact. You will receive a response within 5 working days.

---

## Licence

By contributing to HERALD you agree that your contributions will be licensed under the same licence as the project.
