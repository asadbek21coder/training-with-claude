# Common Software Engineering Roadmap

> **L1** = aware & can use · **L2** = deep & efficient · Status: ⬜ not started · 🔄 in progress · ✅ done
> **Notes:** `notes/common-swe/` · **Practice:** `code/common-swe/` · **Dictionary:** `reference/common-swe-security.md`, `reference/common-swe-behavioral.md`

**Why (EN):** These are the cross-cutting skills every backend engineer uses daily regardless of stack: Git, API design, security, encodings, caching, code review, and communicating like a senior. In the AI era generated code often has insecure defaults, sloppy APIs and wrong time/money handling — a senior is the one who catches these, verifies AI output, builds AI features safely, and owns the outcome end to end.
**Зачем (RU):** Это сквозные навыки, которые каждый backend-инженер использует ежедневно независимо от стека: Git, дизайн API, безопасность, кодировки, кэширование, code review и коммуникация на уровне senior. В эпоху ИИ сгенерированный код часто небезопасен по умолчанию, с неаккуратным API и ошибками в работе со временем и деньгами — senior это замечает, проверяет вывод ИИ, безопасно встраивает ИИ-функции и отвечает за результат целиком.

## Progress
| Layer | Done | Total |
|---|---|---|
| L1 | 0 | 50 |
| L2 | 0 | 40 |

---

## SWE 1 — Git: daily use, internals & workflows
**L1 — aware & can use**
- ⬜ **1.1 Daily Git** — `git status`, `add -p`, `commit`, `diff --staged`, `log --oneline --graph`, `switch`, `restore`, `stash`, `.gitignore`, `remote`, `fetch` vs `pull`
- ⬜ **1.2 Branching & integration** — branch, merge (fast-forward vs three-way), rebase, merge conflicts, `cherry-pick`, `revert` vs `reset`, `pull --rebase`, upstream tracking
- ⬜ **1.3 Team practices** — feature branches, trunk-based development, GitFlow, pull/merge request, conventional commits, protected branches, commit message style (imperative summary + why)

**L2 — deep & efficient**
- ⬜ **1.4 Internals** — object model (blob, tree, commit, tag), SHA-1/SHA-256 object IDs, refs, `HEAD`, detached HEAD, index (staging area), DAG, packfiles, `git cat-file -p`, `.git/` layout
- ⬜ **1.5 History surgery & recovery** — interactive rebase (`rebase -i`: squash, fixup, reword, edit), `--autosquash`, `reflog`, `reset --soft/--mixed/--hard`, `bisect` (incl. `bisect run`), `blame -w`, `filter-repo`, force-push with `--force-with-lease`
- ⬜ **1.6 Scaling & safety** — hooks (pre-commit, commit-msg), signed commits (GPG/SSH), monorepo vs polyrepo, sparse checkout, shallow clone, `worktree`, submodules vs vendoring, Git LFS, CODEOWNERS

🛠 **Practice** — L1: in a scratch repo create two diverging branches, resolve a conflict once with merge and once with rebase, and compare `log --graph` · L2: create a commit by hand with `git hash-object -w`, `update-index`, `write-tree`, `commit-tree`; then use `git bisect run go test ./...` to find a planted bug in a 20-commit history

---

## SWE 2 — SDLC, agile & delivery
**L1 — aware & can use**
- ⬜ **2.1 SDLC basics** — requirements, design, implementation, testing, deployment, maintenance, waterfall vs iterative, environments (dev/staging/prod), release vs deploy
- ⬜ **2.2 Scrum & Kanban** — sprint, sprint planning, daily standup, review, retrospective, backlog, backlog refinement, Kanban board, WIP limits, product owner, scrum master
- ⬜ **2.3 Tasks & estimation** — user story, acceptance criteria, Definition of Ready, Definition of Done, story points vs time, planning poker, splitting stories, spikes

**L2 — deep & efficient**
- ⬜ **2.4 Delivery performance** — DORA metrics (deployment frequency, lead time for changes, change failure rate, time to restore), continuous delivery, small batches, feature flags, dark launches
- ⬜ **2.5 Planning under uncertainty** — estimation ranges, cone of uncertainty, risk-first planning, milestones, cutting scope, roadmap vs commitment, #NoEstimates (awareness)

🛠 **Practice** — L1: break a feature you built in the past ("export report to Excel") into 5–8 user stories with acceptance criteria and estimates · L2: compute DORA metrics for one of your GitHub repos from its commit/release history with a small Go script using the GitHub API

---

## SWE 3 — Code review & documentation
**L1 — aware & can use**
- ⬜ **3.1 Giving & receiving review** — small PRs, PR description (what, why, how tested), review checklist, blocking vs non-blocking, `nit:`, kind and specific feedback, responding to feedback
- ⬜ **3.2 What to look for** — correctness, edge cases, error handling, concurrency, security, tests, naming, API surface, performance hot spots, consistency with codebase conventions
- ⬜ **3.3 Documentation types** — README, runbook, design doc, ADR, API docs (OpenAPI), onboarding guide, changelog, Diátaxis (tutorial, how-to, reference, explanation)

**L2 — deep & efficient**
- ⬜ **3.4 Review at scale** — CODEOWNERS, review SLAs, stacked PRs, automated checks (linters, formatters) before human review, AI review assistants and their limits, reviewing large refactors
- ⬜ **3.5 Docs as code** — docs in repo, Markdown, Mermaid diagrams, generated docs (`go doc`, pkgsite), doc testing (Go examples), keeping docs fresh, docs review

🛠 **Practice** — L1: review an open-source Go PR on GitHub (read-only) and write 5 comments labelled blocking/non-blocking · L2: write a README + runbook for one of your services (setup, config, health checks, common incidents) and a testable `Example` function in Go

---

## SWE 4 — Data formats & encoding
**L1 — aware & can use**
- ⬜ **4.1 Encoding vs encryption vs hashing** — encoding (reversible, no key), encryption (reversible with key), hashing (one-way), base64 / base64url, hex, URL encoding (percent-encoding), "base64 is not security"
- ⬜ **4.2 JSON** — JSON types, `encoding/json`, struct tags, `omitempty`, `null` vs missing, numbers as float64, `json.RawMessage`, streaming decoder, `DisallowUnknownFields`, JSON Schema
- ⬜ **4.3 Unicode & UTF-8** — code point, rune, byte vs character, UTF-8 variable length, `len` vs `utf8.RuneCountInString`, normalization (NFC/NFD), Cyrillic in DB/URLs, BOM, grapheme clusters
- ⬜ **4.4 Time & time zones** — UTC, Unix time, RFC 3339 / ISO 8601, `time.Time`, monotonic clock, time zones vs offsets, IANA tz database, DST, `timestamptz`, store UTC & display local

**L2 — deep & efficient**
- ⬜ **4.5 Binary & schema formats** — Protocol Buffers (field numbers, wire types, varint, backward compatibility), Avro, MessagePack, CBOR, schema evolution, schema registry, size/speed trade-offs vs JSON
- ⬜ **4.6 Floating point & money** — IEEE 754, `0.1 + 0.2`, rounding modes, banker's rounding, minor units (integer cents), `NUMERIC` in Postgres, decimal libraries (`shopspring/decimal`), currency codes (ISO 4217), splitting amounts
- ⬜ **4.7 Edge cases in the wild** — leap seconds, leap years, clock skew, endianness, integer overflow, JSON numbers > 2^53 in JavaScript, line endings (CRLF), case folding, locale-aware sorting (collation)

🛠 **Practice** — L1: write Go tests that show `len("привет")` vs rune count, base64 vs base64url output, and round-trip `time.Time` through JSON in two time zones · L2: implement a `Money` type in Go (int64 minor units + currency) with `Split(n)` that never loses a cent; compare JSON vs protobuf size for the same 1,000 records

---

## SWE 5 — Concurrency concepts (language-agnostic)
**L1 — aware & can use**
- ⬜ **5.1 Core vocabulary** — concurrency vs parallelism, process vs thread vs goroutine/green thread, shared memory vs message passing, critical section, atomicity
- ⬜ **5.2 Classic problems** — race condition, data race, deadlock, livelock, starvation, priority inversion, lost update, check-then-act, read-modify-write
- ⬜ **5.3 Synchronization primitives** — mutex, read-write lock, semaphore, condition variable, atomic operations, barrier, channel, actor model

**L2 — deep & efficient**
- ⬜ **5.4 Memory models & lock-free** — happens-before, memory model, visibility, reordering, compare-and-swap (CAS), ABA problem, lock-free vs wait-free, false sharing
- ⬜ **5.5 Concurrency across processes** — optimistic vs pessimistic locking, `SELECT ... FOR UPDATE`, advisory locks, distributed locks (and why they are hard), fencing tokens, idempotency as concurrency control

🛠 **Practice** — L1: write a Go program with a data race on a counter, catch it with `go test -race`, fix it three ways (mutex, atomic, channel) · L2: reproduce a lost update with two concurrent Postgres transactions, then fix it with optimistic locking (version column) and with `FOR UPDATE`

---

## SWE 6 — Caching concepts
**L1 — aware & can use**
- ⬜ **6.1 Where caches live** — CPU cache, in-process cache, distributed cache (Redis, Memcached), CDN, browser cache, DB buffer cache, DNS cache
- ⬜ **6.2 Caching patterns** — cache-aside (lazy loading), read-through, write-through, write-behind, write-around, TTL, cache hit ratio, cache key design
- ⬜ **6.3 HTTP caching** — `Cache-Control` (`max-age`, `no-store`, `private`, `public`), `ETag`, `If-None-Match`, `Last-Modified`, `304 Not Modified`, `Vary`

**L2 — deep & efficient**
- ⬜ **6.4 Invalidation & consistency** — "two hard things", invalidation strategies, stale reads, versioned keys, delete-vs-update on write, dual-write race, CDC-driven invalidation
- ⬜ **6.5 Failure modes** — thundering herd / cache stampede, `singleflight`, probabilistic early expiration, hot keys, cache penetration (negative caching, Bloom filter), cache avalanche (TTL jitter), eviction policies (LRU, LFU, ARC)

🛠 **Practice** — L1: add cache-aside with Redis (Docker) to a slow Go endpoint and measure latency before/after with `hey` or `vegeta` · L2: simulate a cache stampede with 500 concurrent requests on an expired key, then fix it with `singleflight` and TTL jitter and compare DB query counts

---

## SWE 7 — API design: REST fundamentals
**L1 — aware & can use**
- ⬜ **7.1 Resources & HTTP semantics** — resources, URIs, nouns vs verbs, HTTP methods, safe vs idempotent methods, `PUT` vs `PATCH`, sub-resources, `Content-Type`, `Accept`
- ⬜ **7.2 Status codes** — 2xx (`200`, `201`, `202`, `204`), 3xx, 4xx (`400`, `401` vs `403`, `404`, `409`, `422`, `429`), 5xx (`500`, `502`, `503`, `504`), `Location` and `Retry-After` headers
- ⬜ **7.3 Errors** — consistent error model, RFC 9457 Problem Details (`type`, `title`, `status`, `detail`, `instance`), `application/problem+json`, validation error lists, never leaking internals
- ⬜ **7.4 Collections** — filtering, sorting, field selection, offset vs cursor (keyset) pagination, page tokens, `Link` header, total count cost

**L2 — deep & efficient**
- ⬜ **7.5 Idempotency** — idempotency keys (`Idempotency-Key` header), stored responses, key scope & expiry, concurrent duplicate requests, idempotent consumers, exactly-once illusion
- ⬜ **7.6 API evolution & contracts** — API-first design, OpenAPI 3.x, code generation (`oapi-codegen`), versioning (URL vs header vs media type), backward/forward compatibility, additive changes, deprecation & `Sunset` header, breaking-change detection
- ⬜ **7.7 Advanced REST concerns** — long-running operations (`202` + status resource), bulk/batch endpoints, partial failures, rate-limit headers (`RateLimit-*`), optimistic concurrency (`ETag` + `If-Match`), HATEOAS (awareness)

🛠 **Practice** — L1: design (OpenAPI YAML) and implement in Go a `/v1/orders` API with cursor pagination and RFC 9457 errors · L2: add `Idempotency-Key` support to `POST /v1/payments` backed by Postgres; prove with a concurrent test that 20 parallel duplicate requests create exactly one payment

---

## SWE 8 — API styles: gRPC, GraphQL, real-time & webhooks
**L1 — aware & can use**
- ⬜ **8.1 gRPC concepts** — Protocol Buffers IDL, unary vs streaming (server, client, bidirectional), HTTP/2, status codes, deadlines, metadata, gRPC-Gateway, when gRPC vs REST
- ⬜ **8.2 GraphQL concepts** — schema, query, mutation, subscription, resolvers, over-/under-fetching, N+1 problem, DataLoader, when GraphQL fits
- ⬜ **8.3 Real-time** — polling, long polling, Server-Sent Events (SSE), WebSocket, connection lifecycle, heartbeats/ping-pong, reconnection
- ⬜ **8.4 Webhooks** — webhook, event payload, HMAC signature verification, timestamp & replay protection, retries with backoff, idempotent receivers, delivery logs

**L2 — deep & efficient**
- ⬜ **8.5 Style trade-offs at scale** — API gateway, BFF, GraphQL federation, query complexity limits, gRPC load balancing (L7), WebSocket horizontal scaling (sticky sessions, pub/sub fan-out), AsyncAPI
- ⬜ **8.6 Contract evolution for non-REST** — protobuf field rules (never reuse numbers, `reserved`), `buf breaking`, GraphQL schema deprecation, event schema versioning, consumer-driven contracts

🛠 **Practice** — L1: build a webhook sender + receiver in Go with HMAC-SHA256 signatures, timestamp check and retries · L2: expose the same `GetOrder` as REST, gRPC and SSE stream in one Go service; run `buf breaking` against a deliberately breaking proto change

---

## SWE 9 — Cryptography & hashing (use correctly, never invent)
**L1 — aware & can use**
- ⬜ **9.1 Hash functions** — cryptographic hash, SHA-256, SHA-3, BLAKE2/3, collision/preimage resistance, MD5 & SHA-1 broken, checksums vs crypto hashes (CRC32, xxHash), content addressing
- ⬜ **9.2 HMAC & message authentication** — HMAC, MAC vs hash, shared secret, `crypto/hmac`, constant-time comparison (`hmac.Equal`), length-extension attack, signing webhooks and cookies
- ⬜ **9.3 Password hashing** — slow hashing, salt, pepper, bcrypt (cost factor, 72-byte limit), scrypt, argon2id (memory, iterations, parallelism), PHC string format, rehash on login
- ⬜ **9.4 Symmetric & asymmetric encryption** — symmetric (AES-GCM, ChaCha20-Poly1305), AEAD, nonce/IV, asymmetric (RSA, ECDH, X25519), public/private key, hybrid encryption, `crypto/rand` vs `math/rand`
- ⬜ **9.5 Digital signatures & certificates** — signature vs encryption, RSA-PSS, ECDSA, Ed25519, X.509 certificate, CA, chain of trust, certificate pinning (awareness)

**L2 — deep & efficient**
- ⬜ **9.6 TLS concepts** — TLS 1.3 handshake, key exchange, forward secrecy, cipher suites, SNI, ALPN, mTLS, certificate validation, Let's Encrypt / ACME, `crypto/tls` config
- ⬜ **9.7 Key management** — KMS, envelope encryption (DEK/KEK), key rotation, key versioning, HSM (awareness), encryption at rest vs in transit, field-level encryption, tokenization
- ⬜ **9.8 Common crypto mistakes** — nonce reuse in GCM, ECB mode, homemade crypto, weak randomness, timing attacks, hardcoded keys, "encrypt" where hashing is needed, post-quantum crypto (ML-KEM — awareness)

🛠 **Practice** — L1: write Go tests that hash a password with argon2id (`golang.org/x/crypto/argon2`) in PHC format, verify it, and sign/verify a message with Ed25519 and HMAC-SHA256 · L2: implement envelope encryption in Go (random DEK with AES-GCM, DEK wrapped by a KEK), plus key rotation that re-wraps DEKs without re-encrypting data

---

## SWE 10 — Authentication
**L1 — aware & can use**
- ⬜ **10.1 Sessions & cookies** — session ID, server-side session store, cookie attributes (`HttpOnly`, `Secure`, `SameSite`, `Domain`, `Path`, `Max-Age`), session fixation, logout & invalidation
- ⬜ **10.2 Passwords & MFA** — password policy (NIST 800-63B: length over complexity), breached password check, credential stuffing, account lockout vs rate limiting, MFA, TOTP, recovery codes, passkeys/WebAuthn
- ⬜ **10.3 JWT** — header/payload/signature, JWS vs JWE, claims (`iss`, `sub`, `aud`, `exp`, `iat`, `nbf`, `jti`), HS256 vs RS256/EdDSA, access vs refresh token, where to store tokens
- ⬜ **10.4 OAuth 2.0 & OIDC** — resource owner, client, authorization server, resource server, authorization code + PKCE, client credentials, scopes, ID token, OpenID Connect, SSO, Keycloak

**L2 — deep & efficient**
- ⬜ **10.5 Token lifecycle** — revocation problem, short-lived access tokens, refresh token rotation & reuse detection, denylist, token introspection, JWKS & key rotation (`kid`), `alg: none` and algorithm confusion attacks
- ⬜ **10.6 OAuth/OIDC in depth** — OAuth 2.1 direction, deprecated implicit & password grants, `state` and `nonce`, redirect URI validation, device code flow, token exchange, DPoP (awareness), BFF pattern for SPAs
- ⬜ **10.7 Service-to-service auth** — API keys (hashing, prefixes, rotation), mTLS, SPIFFE/SPIRE (awareness), workload identity, signed requests (AWS SigV4 style)

🛠 **Practice** — L1: implement JWT login in Go (`golang-jwt/jwt/v5`): `/login` checks an argon2id hash and issues a 15-min access token + refresh token; middleware validates `exp`, `aud`, `iss` · L2: add refresh token rotation with reuse detection in Postgres, and run Keycloak in Docker to log in via authorization code + PKCE from a Go client

---

## SWE 11 — Authorization
**L1 — aware & can use**
- ⬜ **11.1 Authentication vs authorization** — authN vs authZ, `401` vs `403`, principal, permission, policy, least privilege, deny by default, checking on every request
- ⬜ **11.2 RBAC** — roles, permissions, role hierarchy, role explosion, scopes in tokens, admin vs user endpoints, function-level authorization

**L2 — deep & efficient**
- ⬜ **11.3 ABAC & policy engines** — attributes (subject, resource, action, environment), Open Policy Agent (OPA), Rego, Cedar, policy as code, policy decision point vs enforcement point (PDP/PEP)
- ⬜ **11.4 ReBAC** — relationship-based access control, Google Zanzibar, OpenFGA, SpiceDB, tuples, check vs list-objects queries, consistency tokens
- ⬜ **11.5 Multi-tenant & object-level authorization** — BOLA/IDOR, tenant isolation, `WHERE tenant_id = $1` everywhere, Postgres row-level security, authorization in list endpoints, audit logs

🛠 **Practice** — L1: add RBAC middleware (`admin`, `editor`, `viewer`) to a Go API with a table-driven test of every role × endpoint · L2: model "users can edit documents in folders they share" in OpenFGA (Docker) and call it from Go; compare with the same rules in plain SQL

---

## SWE 12 — Web & API security (OWASP)
**L1 — aware & can use**
- ⬜ **12.1 OWASP Top 10** — broken access control, cryptographic failures, injection, insecure design, security misconfiguration, vulnerable and outdated components, identification and authentication failures, software and data integrity failures, logging and monitoring failures, SSRF
- ⬜ **12.2 OWASP API Security Top 10** — BOLA, broken authentication, broken object property level authorization (mass assignment), unrestricted resource consumption, broken function level authorization, unrestricted access to sensitive business flows, SSRF, improper inventory management, unsafe consumption of APIs
- ⬜ **12.3 Injection** — SQL injection, parameterized queries (`$1`), dynamic `ORDER BY` allowlists, command injection (`exec.Command` without shell), LDAP/NoSQL injection, template injection, `html/template` vs `text/template`
- ⬜ **12.4 Browser-side attacks** — XSS (stored, reflected, DOM), output encoding, Content-Security-Policy, CSRF, `SameSite` cookies, CSRF tokens, clickjacking (`frame-ancestors`), open redirect, CORS (what it does and doesn't protect)

**L2 — deep & efficient**
- ⬜ **12.5 Server-side attacks** — SSRF (cloud metadata `169.254.169.254`, DNS rebinding, allowlists), path traversal (`os.Root`, `filepath.Clean` pitfalls), XXE, insecure deserialization, request smuggling (awareness), ReDoS, zip bombs, resource limits (`http.MaxBytesReader`)
- ⬜ **12.6 Security process** — threat modeling (STRIDE, data flow diagram, trust boundaries, attack surface), security headers, audit logging, rate limiting & bot protection, responsible disclosure, incident response, compliance awareness (GDPR, PCI DSS, SOC 2, data localization laws)
- ⬜ **12.7 Security testing** — SAST (`gosec`, Semgrep, CodeQL), DAST (OWASP ZAP), fuzzing for security, Burp Suite (awareness), pentest vs bug bounty, security regression tests

🛠 **Practice** — L1: build a toy Go app with an SQL-injectable search (`fmt.Sprintf` query), exploit it with `' OR 1=1 --` and a `UNION` dump, then fix with parameters and an `ORDER BY` allowlist · L2: add an SSRF-vulnerable "fetch URL preview" endpoint, exploit it against a fake metadata service in Docker, then fix with a custom dialer that blocks private IPs after DNS resolution; run ZAP baseline scan

---

## SWE 13 — Secrets management & supply chain security
**L1 — aware & can use**
- ⬜ **13.1 Secrets basics** — never commit secrets, env vars vs files, `.env` + `.gitignore`, secret scanning (gitleaks, GitHub push protection), rotating a leaked secret, Kubernetes Secrets (base64 ≠ encryption)
- ⬜ **13.2 Dependency hygiene** — `go.sum`, module proxy & checksum DB (`sum.golang.org`), `govulncheck`, Dependabot, Renovate, pinning versions, minimal dependencies, typosquatting

**L2 — deep & efficient**
- ⬜ **13.3 Secret managers** — HashiCorp Vault / OpenBao, dynamic secrets, leases, cloud secret managers (AWS Secrets Manager, GCP Secret Manager), SOPS + age, External Secrets Operator, workload identity instead of static keys
- ⬜ **13.4 Supply chain integrity** — SBOM (SPDX, CycloneDX, Syft), image scanning (Trivy, Grype), signing (cosign / Sigstore, keyless), provenance, SLSA levels, reproducible builds, pinning GitHub Actions by SHA, malicious packages (incl. AI-hallucinated "slopsquatting")
- ⬜ **13.5 Container & CI hardening** — non-root user, distroless/scratch images, read-only rootfs, dropped capabilities, least-privilege CI tokens, OIDC from CI to cloud, branch protection, secrets in CI logs

🛠 **Practice** — L1: plant a fake AWS key in a scratch repo commit and catch it with `gitleaks`; run `govulncheck ./...` on one of your projects · L2: build a distroless Go image, generate an SBOM with Syft, scan with Trivy, sign with cosign (local key) and verify the signature

---

## SWE 14–16 — moved to the AI track
AI-assisted engineering, LLM APIs & AI features, and AI security now live in [`roadmap/ai.md`](ai.md) (AI 3–5, AI 6–10, AI 11). The numbers SWE 14–16 stay unused so other IDs don't change.

---

## SWE 17 — Senior engineering behaviours
**L1 — aware & can use**
- ⬜ **17.1 Ownership** — end-to-end ownership, raising risks early, follow-through, on-call mindset, "disagree and commit", fixing root causes not symptoms
- ⬜ **17.2 Communication** — explaining trade-offs to non-engineers, asking clarifying questions, written updates, status reports, saying "I don't know" well, async communication
- ⬜ **17.3 Product thinking** — user impact, business metrics, prioritization (impact vs effort), MVP, saying "no" with alternatives, cost of features over time

**L2 — deep & efficient**
- ⬜ **17.4 Mentoring & teaching** — teaching (your student), pair programming, onboarding plans, giving feedback (SBI model), growth plans, delegating with context
- ⬜ **17.5 Leading projects** — planning & splitting work, coordinating across teams, stakeholder management, running technical discussions, influence without authority, decision logs, project kick-off and retro
- ⬜ **17.6 Engineering levels & growth** — career ladders (middle, senior, staff), scope and impact, brag document, technical strategy (awareness), staff archetypes (tech lead, architect, solver, right hand)

🛠 **Practice** — L1: write a one-paragraph trade-off explanation of a technical decision you made, aimed at a non-technical manager · L2: write a brag document for the last 12 months and a one-page project plan (milestones, risks, owners) for a feature you'd like to lead

---

## SWE 18 — Interviews & career
**L1 — aware & can use**
- ⬜ **18.1 Behavioral interviews (STAR)** — situation, task, action, result; story bank: conflict, failure, impact, leadership, ambiguity, tight deadline, mistake, disagreement with manager
- ⬜ **18.2 Self-introduction & project stories** — 60-second and 2-minute pitch, project deep-dive structure (problem, your role, decisions, results, what you'd change), quantifying impact
- ⬜ **18.3 Technical round strategy** — clarifying questions, thinking out loud, live coding habits, system design framework, handling "I don't know", asking good questions to interviewers

**L2 — deep & efficient**
- ⬜ **18.4 Offers & negotiation** — market research, total compensation, competing offers, negotiation phrasing, remote/relocation terms, notice period
- ⬜ **18.5 Choosing companies** — engineering culture signals, on-call expectations, code review culture, team stability, growth path, red flags in interviews
- ⬜ **18.6 Visibility** — production-quality GitHub portfolio, technical blog, LinkedIn posts, open-source contributions, conference/meetup talks, Go community

🛠 **Practice** — L1: write 6 STAR stories (one per category) from your Iman, PPP and IMV Archive projects and record yourself telling two of them in under 2 minutes · L2: run a full mock interview (behavioral + technical) and write a negotiation script for a target offer
