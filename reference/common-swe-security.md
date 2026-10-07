# 11 — Security

> Backend engineers are responsible for writing secure code and understanding common attack vectors.

---

## Authentication vs Authorization

**Key Terms:**
- `authentication (AuthN)` — verifying identity; "who are you?"
- `authorization (AuthZ)` — verifying permissions; "what can you do?"
- `credential` — secret that proves identity (password, token, certificate)
- `principal` — authenticated entity (user, service, system)
- `session` — server-side state tracking authenticated user across requests
- `session ID` — opaque token stored in cookie; maps to server-side session
- `stateful auth` — server stores sessions; horizontal scale requires shared session store (Redis)
- `stateless auth` — JWT; server validates token without storing state; scales easily
- `multi-factor authentication (MFA)` — something you know + something you have/are
- `TOTP (Time-based OTP)` — time-based one-time password; Google Authenticator
- `SSO (Single Sign-On)` — one login for multiple services; SAML, OIDC
- `RBAC (Role-Based Access Control)` — permissions assigned to roles; user has roles
- `ABAC (Attribute-Based Access Control)` — permissions based on user/resource/environment attributes
- `ACL (Access Control List)` — explicit list of who can access what
- `least privilege principle` — grant minimum permissions needed; nothing more
- `zero trust` — never trust, always verify; even internal services must authenticate

**Must understand:**
- [ ] Authentication happens before authorization (can't authorize without knowing who)
- [ ] Session vs JWT — trade-offs (stateful vs stateless; revocability vs scalability)
- [ ] RBAC vs ABAC — RBAC simpler; ABAC more flexible for complex rules
- [ ] Least privilege — why it limits blast radius of compromised credentials

---

## JWT

**Key Terms:**
- `JWT (JSON Web Token)` — self-contained token; three parts: header.payload.signature
- `header` — algorithm and token type: `{"alg":"HS256","typ":"JWT"}`
- `payload (claims)` — data: `{"sub":"userId","exp":1234567890,"iat":1234560000,"roles":["admin"]}`
- `signature` — HMAC-SHA256(base64(header) + "." + base64(payload), secret)
- `HS256` — HMAC-SHA256; symmetric; same secret for sign and verify; simple
- `RS256` — RSA-SHA256; asymmetric; private key signs; public key verifies; distribute public key safely
- `sub (subject)` — who the token is about (user ID)
- `exp (expiration)` — unix timestamp when token expires; always set this
- `iat (issued at)` — when token was issued
- `nbf (not before)` — token not valid before this time
- `jti (JWT ID)` — unique ID; enables blacklisting individual tokens
- `access token` — short-lived (15min-1hr); used for API calls
- `refresh token` — long-lived (days-weeks); used to get new access token; stored securely
- `token rotation` — issue new refresh token with each refresh; invalidate old one
- `JWT blacklist` — store revoked JTIs in Redis; check on each request (reduces statelessness benefit)
- `JWT size` — JWTs can be large if many claims; affects every HTTP request
- `JWT validation` — verify signature; check exp; check iss (issuer); check aud (audience)

**Security risks:**
- `alg:none attack` — malicious client sets alg to "none"; skip signature check; always validate alg
- `weak secret` — HS256 with short/guessable secret; brute-forceable
- `no expiration` — tokens live forever if exp not set
- `sensitive data in payload` — payload is base64-encoded, not encrypted; don't put secrets

**Must understand:**
- [ ] JWT is NOT encrypted — payload is readable by anyone; never put sensitive data
- [ ] RS256 vs HS256 — RS256 lets anyone verify but only server signs (useful for microservices)
- [ ] How to revoke JWTs (blacklist or short TTL + refresh token rotation)
- [ ] alg:none vulnerability — always explicitly specify accepted algorithms

---

## OAuth 2.0 / OIDC

**Key Terms:**
- `OAuth 2.0` — authorization framework; lets users grant apps access to their data without sharing password
- `authorization server` — issues tokens (Google, GitHub, Auth0, Keycloak)
- `resource server` — API server; validates access token
- `client` — app requesting access on user's behalf
- `resource owner` — user who owns the data
- `scope` — specific permission (read:profile, write:email)
- `access token` — credential to access resource server; short-lived
- `refresh token` — long-lived; get new access token without user re-login
- `authorization code flow` — most secure; browser-based apps; code exchanged server-side for token
- `PKCE (Proof Key for Code Exchange)` — extension for public clients (mobile/SPA); prevents code interception
- `client credentials flow` — machine-to-machine; no user; service authenticates with client_id + secret
- `implicit flow` — deprecated; access token in URL fragment (security risk)
- `redirect URI` — where authorization server sends code/token after user consent
- `state parameter` — random value in auth request; verify in callback to prevent CSRF
- `OIDC (OpenID Connect)` — extension of OAuth 2.0; adds identity; returns ID token (JWT with user info)
- `ID token` — JWT from OIDC; contains user's identity claims (sub, email, name)
- `userinfo endpoint` — get user info using access token
- `PKCE` — prevents authorization code interception attacks in public clients

**Authorization Code Flow:**
```
1. App → Authorization Server: GET /auth?client_id=...&redirect_uri=...&scope=...&state=...
2. User authenticates and consents
3. Authorization Server → App: redirect to redirect_uri?code=AUTH_CODE&state=...
4. App → Authorization Server: POST /token {code, client_id, client_secret}
5. Authorization Server → App: {access_token, refresh_token, id_token}
6. App → Resource Server: GET /api/data with Authorization: Bearer access_token
```

**Must understand:**
- [ ] Why OAuth was needed (users shouldn't share passwords with third-party apps)
- [ ] Authorization code flow — why code is exchanged server-side (never expose token in URL)
- [ ] PKCE — what it prevents (code interception in public clients)
- [ ] OAuth 2.0 for authorization, OIDC for authentication

---

## OWASP Top 10

**Key Terms:**
- `A01 — Broken Access Control` — user accesses data/functionality they shouldn't (IDOR, privilege escalation)
- `A02 — Cryptographic Failures` — sensitive data exposed due to weak/missing encryption (HTTP not HTTPS, weak cipher)
- `A03 — Injection` — untrusted data sent to interpreter as command (SQL, OS command, LDAP)
- `A04 — Insecure Design` — architectural flaws; no threat modeling; weak security controls by design
- `A05 — Security Misconfiguration` — default credentials, open cloud storage, verbose errors, unnecessary features enabled
- `A06 — Vulnerable & Outdated Components` — libraries with known CVEs; unpatched dependencies
- `A07 — Identification & Authentication Failures` — weak passwords, credential stuffing, no MFA, session fixation
- `A08 — Software & Data Integrity Failures` — CI/CD pipeline attack, unsigned updates, deserialization of untrusted data
- `A09 — Security Logging & Monitoring Failures` — no logs, logs not monitored, breach not detected for months
- `A10 — Server-Side Request Forgery (SSRF)` — server fetches user-controlled URL; attacker reaches internal services

**IDOR (Insecure Direct Object Reference) example:**
- URL: `GET /api/orders/12345` — user changes ID to 12346 and sees another user's order
- Fix: always check ownership: `WHERE id=? AND user_id=?`

**SSRF example:**
- App fetches URL user provides: `fetch(request.body.url)`
- Attacker provides: `http://169.254.169.254/latest/meta-data/` (AWS metadata)
- Fix: whitelist allowed domains; use DNS rebinding prevention

**Must understand:**
- [ ] IDOR — most common access control bug; always validate ownership
- [ ] SQL Injection — how parameterized queries prevent it
- [ ] SSRF — why fetching user-provided URLs is dangerous in cloud environments
- [ ] Security misconfiguration — default credentials, exposed debug endpoints, open S3 buckets

---

## SQL Injection

**Key Terms:**
- `SQL injection` — attacker injects SQL code through user input that's concatenated into query
- `input concatenation` — `"SELECT * FROM users WHERE name = '" + name + "'"` — NEVER do this
- `payload` — malicious input: `' OR '1'='1` or `'; DROP TABLE users; --`
- `comment` — `--` ends SQL; allows attacker to ignore rest of query
- `parameterized query (prepared statement)` — bind variables; SQL and data separated; safe
- `placeholder` — `?` (MySQL/Go) or `$1` (PostgreSQL) marks where parameter goes
- `ORM protection` — ORMs use parameterized queries by default; safe as long as not using raw SQL
- `blind SQL injection` — no visible output; attacker infers info from behavior (true/false response, timing)
- `time-based injection` — `'; WAITFOR DELAY '0:0:5'--`; infer info from response time
- `second-order injection` — data stored safely but used unsafely later
- `database firewall` — detect/block injection patterns; defense in depth
- `principle of least privilege` — DB user for app should only have SELECT/INSERT/UPDATE; not DROP

**Safe vs unsafe:**
```go
// UNSAFE — string concatenation
db.Query("SELECT * FROM users WHERE name = '" + name + "'")

// SAFE — parameterized
db.Query("SELECT * FROM users WHERE name = ?", name)
db.Query("SELECT * FROM users WHERE name = $1", name)
```

**Must understand:**
- [ ] Why string concatenation is vulnerable (attacker controls SQL structure)
- [ ] Why parameterized queries are safe (DB treats input as literal data, not SQL)
- [ ] Where ORM raw SQL is still dangerous (custom SQL fragments, ORDER BY direction)
- [ ] LIKE with parameterized query still needs % escaping for wildcards in pattern

---

## HTTPS / TLS in Practice

**Key Terms:**
- `HSTS (HTTP Strict Transport Security)` — header telling browser: "always use HTTPS for this domain"
  - `Strict-Transport-Security: max-age=31536000; includeSubDomains; preload`
- `HSTS preload` — submit domain to browser list; hardcoded HTTPS even on first visit
- `certificate pinning` — app hardcodes expected certificate/public key; rejects others (MITM prevention)
- `OCSP (Online Certificate Status Protocol)` — check if cert is revoked; OCSP stapling = server includes proof
- `certificate transparency` — public log of all issued certs; detect rogue certs
- `SNI` — Server Name Indication; client sends hostname in TLS handshake → server serves correct cert
- `wildcard certificate` — `*.example.com` covers all subdomains
- `Let's Encrypt` — free CA; 90-day certs; ACME protocol for automated renewal
- `certbot` — Let's Encrypt certificate management tool
- `TLS termination` — load balancer/proxy handles TLS; backend gets plain HTTP (simpler, centralized)
- `mutual TLS (mTLS)` — both client and server present certificates; zero-trust service authentication
- `cipher suite` — negotiated combination of algorithms; prefer forward secrecy (ECDHE)
- `forward secrecy (PFS)` — each session has unique key; past sessions can't be decrypted if private key leaked
- `mixed content` — HTTPS page loading HTTP resource; browser blocks it

**Must understand:**
- [ ] HSTS — why it prevents SSL stripping attacks
- [ ] mTLS — how it's used for service-to-service authentication in microservices (Istio handles this)
- [ ] Why TLS termination at LB is common (simpler cert management, backend doesn't need certs)
- [ ] Forward secrecy — why it matters (even if private key leaked, past sessions safe)

---

## Secrets Management

**Key Terms:**
- `secret` — any sensitive value: DB passwords, API keys, private keys, tokens
- `hardcoded secret` — secret in source code; NEVER; visible in git history forever
- `environment variable` — common way to inject secrets; better than code but still risks
- `.env file` — local dev convenience; NEVER commit to git; add to .gitignore
- `secrets manager` — dedicated system for storing and accessing secrets (HashiCorp Vault, AWS Secrets Manager, GCP Secret Manager)
- `HashiCorp Vault` — open-source; dynamic secrets; lease-based; audit log
- `dynamic secrets` — Vault generates short-lived credentials on demand (DB user created per app instance)
- `secret rotation` — regularly change secrets; reduce window if leaked
- `least privilege` — each service only accesses secrets it needs
- `secret scanning` — tools that scan git history for accidentally committed secrets (git-secrets, trufflehog, GitHub scanning)
- `SOPS` — tool to encrypt secrets files and commit encrypted version to git
- `Sealed Secrets (Kubernetes)` — encrypt Kubernetes Secrets; safe to commit
- `IRSA (IAM Roles for Service Accounts)` — K8s pod gets AWS credentials via IAM role; no static keys

**Must understand:**
- [ ] Never commit secrets to git — they're in history even if deleted
- [ ] Environment variables vs secrets manager — env vars ok for non-critical; use manager for critical
- [ ] Secret rotation — why it's important and how to implement without downtime
- [ ] Dynamic secrets — why they're more secure (short-lived, no shared long-term credentials)

---

## Rate Limiting & DDoS Basics

**Key Terms:**
- `DDoS (Distributed Denial of Service)` — flood of traffic from many sources to overwhelm target
- `amplification attack` — small request causes large response (DNS, NTP amplification)
- `SYN flood` — send many TCP SYN without completing handshake; exhaust connection table
- `volumetric attack` — overwhelm with raw bandwidth
- `application layer (L7) DDoS` — mimic legitimate HTTP requests; harder to detect
- `WAF (Web Application Firewall)` — filter malicious HTTP traffic; block SQL injection, XSS
- `Cloudflare` — absorbs DDoS at edge network before reaching origin; most common protection
- `rate limiting` — limit requests per user/IP/API key per time window
- `IP reputation` — block or challenge IPs with known malicious history
- `CAPTCHA` — challenge to distinguish humans from bots
- `bot protection` — detect automated traffic patterns; challenge or block
- `Anycast` — traffic routed to nearest Cloudflare PoP; distributes attack across global network
- `scrubbing center` — dedicated infrastructure to filter attack traffic; re-route legitimate traffic
- `CDN as DDoS protection` — CDN absorbs traffic at edges; origin only gets clean traffic
- `egress filtering` — prevent your servers from being used in amplification attacks

**Must understand:**
- [ ] DDoS mitigation is mostly infrastructure (CDN, Cloudflare) not application code
- [ ] Application-level rate limiting (per user/IP/token) — defense-in-depth
- [ ] WAF — what it protects against (injection, XSS) and what it doesn't (DDoS volume)
- [ ] The role of Anycast in distributing DDoS traffic
