# 05 — Networking

> Backend engineers are asked about HTTP, TCP, DNS, and TLS regularly. Know the flow, not just the names.

---

## HTTP / HTTPS

**Key Terms:**
- `HTTP (HyperText Transfer Protocol)` — application-layer protocol for web communication
- `request` — client sends: method, URL, headers, optional body
- `response` — server sends: status code, headers, body
- `header` — key-value metadata (Content-Type, Authorization, Cache-Control)
- `body` — payload (JSON, HTML, binary)
- `Content-Type` — type of body (application/json, text/html, multipart/form-data)
- `Authorization` — carries credentials (Bearer token, Basic auth)
- `stateless` — each HTTP request is independent; server remembers nothing by default
- `keep-alive` — HTTP/1.1 default; reuse TCP connection for multiple requests
- `HTTPS` — HTTP over TLS; encrypted connection
- `TLS` — Transport Layer Security; replaces SSL; encrypts HTTP traffic

**HTTP Status Codes:**

| Code | Meaning |
|------|---------|
| 200 | OK |
| 201 | Created |
| 204 | No Content |
| 301 | Moved Permanently |
| 302 | Found (Temporary Redirect) |
| 304 | Not Modified (cache hit) |
| 400 | Bad Request (invalid input) |
| 401 | Unauthorized (not authenticated) |
| 403 | Forbidden (authenticated but no permission) |
| 404 | Not Found |
| 405 | Method Not Allowed |
| 409 | Conflict |
| 429 | Too Many Requests (rate limit) |
| 500 | Internal Server Error |
| 502 | Bad Gateway |
| 503 | Service Unavailable |
| 504 | Gateway Timeout |

**Must understand:**
- [ ] 401 vs 403 — not authenticated vs no permission
- [ ] 302 vs 301 — temporary vs permanent redirect
- [ ] Why HTTP is stateless and how sessions solve it (cookies, JWT)
- [ ] What headers are sent in a typical API request

---

## HTTP/1.1 vs HTTP/2 vs HTTP/3

**Key Terms:**
- `HTTP/1.0` — new TCP connection per request; no keep-alive by default
- `HTTP/1.1` — persistent connections (keep-alive); pipelining (rarely used); host header required
- `head-of-line blocking (HOL)` — in HTTP/1.1, next request blocked until current response complete
- `HTTP/2` — multiplexing (multiple requests over one TCP connection); header compression (HPACK); server push; binary framing
- `multiplexing` — multiple streams over one connection; no HOL at application layer
- `stream` — independent, bidirectional channel within HTTP/2 connection
- `HPACK` — header compression in HTTP/2; avoids resending same headers each request
- `server push` — server proactively sends resources (e.g., CSS + JS with HTML)
- `HTTP/3` — uses QUIC (UDP-based) instead of TCP; eliminates TCP-level HOL blocking
- `QUIC` — UDP + reliability + encryption built-in; faster connection establishment
- `0-RTT` — QUIC resumes connections with zero round-trips (for known servers)
- `TLS 1.3` — required for HTTP/2 and HTTP/3 in browsers; faster handshake

**Comparison:**

| Feature | HTTP/1.1 | HTTP/2 | HTTP/3 |
|---------|----------|--------|--------|
| Protocol | TCP | TCP | QUIC (UDP) |
| Multiplexing | ❌ | ✅ | ✅ |
| Header compression | ❌ | ✅ HPACK | ✅ QPACK |
| Server push | ❌ | ✅ | ✅ |
| HOL blocking | TCP + App | TCP only | ❌ |
| Connection setup | 1-3 RTT | 1-3 RTT | 0-1 RTT |

**Must understand:**
- [ ] Why HTTP/1.1 has head-of-line blocking (sequential request/response)
- [ ] How HTTP/2 multiplexing solves it (parallel streams over one TCP)
- [ ] Why HTTP/3 uses UDP (TCP HOL blocking at transport layer can't be fixed)
- [ ] HTTPS requires TLS — HTTP/2 in practice always uses TLS

---

## TCP vs UDP

**Key Terms:**
- `TCP (Transmission Control Protocol)` — reliable, ordered, connection-oriented
- `UDP (User Datagram Protocol)` — unreliable, unordered, connectionless, fast
- `3-way handshake` — SYN → SYN-ACK → ACK; establishes TCP connection
- `4-way teardown` — FIN → ACK → FIN → ACK; closes TCP connection
- `ACK (acknowledgement)` — receiver confirms packet received; TCP retransmits if no ACK
- `sequence number` — tracks order of bytes; enables reordering out-of-order packets
- `flow control` — prevents sender overwhelming receiver (receiver window)
- `congestion control` — prevents sender overwhelming network (TCP slow start, AIMD)
- `connection state` — TCP is stateful; UDP is stateless
- `checksum` — both TCP and UDP have checksums; UDP still detects corruption, just doesn't retransmit
- `port` — 16-bit number identifying application on a host (HTTP=80, HTTPS=443, SSH=22)
- `socket` — IP:Port combination; a connection is identified by (src IP, src port, dst IP, dst port)

**When to use TCP:** HTTP, database connections, file transfer — reliability matters

**When to use UDP:** video streaming, gaming, DNS queries, VoIP — latency matters more than reliability

**Must understand:**
- [ ] TCP 3-way handshake — why it's 3 steps (both sides must confirm they can send and receive)
- [ ] Why UDP is faster (no handshake, no retransmit, no ordering)
- [ ] How TCP handles packet loss (retransmit after timeout)
- [ ] SYN flood attack — attacker sends SYN without completing handshake (DoS)

---

## DNS

**Key Terms:**
- `DNS (Domain Name System)` — translates domain name → IP address
- `A record` — domain → IPv4 address
- `AAAA record` — domain → IPv6 address
- `CNAME record` — domain → another domain (alias)
- `MX record` — domain → mail server
- `TXT record` — text data; used for domain verification, SPF, DKIM
- `NS record` — nameserver for domain
- `TTL` — how long to cache DNS response (seconds)
- `recursive resolver` — client's DNS server; queries on behalf of client (ISP or 8.8.8.8)
- `authoritative nameserver` — knows the definitive answer for a domain
- `root nameserver` — top of DNS hierarchy; knows where TLD nameservers are (13 root servers)
- `TLD nameserver` — knows nameservers for .com, .org, .io, etc.
- `DNS resolution flow` — client → recursive resolver → root → TLD → authoritative
- `DNS caching` — resolvers cache responses for TTL duration; low TTL = slower, faster propagation
- `DNS propagation` — after changing DNS record, takes up to TTL time to propagate globally
- `split-horizon DNS` — different answers for internal vs external clients
- `DNS load balancing` — multiple A records for same domain; round-robin

**DNS Resolution Flow:**
```
Browser checks local cache
→ OS cache
→ Recursive resolver (ISP or 8.8.8.8)
   → Root nameserver (knows .com TLD server)
   → TLD nameserver (knows example.com NS)
   → Authoritative NS (returns IP)
← IP returned and cached at each step
```

**Must understand:**
- [ ] Full DNS resolution flow from browser to authoritative server
- [ ] Why low TTL = slower performance but faster record propagation
- [ ] CNAME vs A record — when each is used
- [ ] How CDNs use DNS for geo-routing (return nearest edge server IP)

---

## TLS / SSL

**Key Terms:**
- `SSL` — older, deprecated; replaced by TLS
- `TLS (Transport Layer Security)` — current standard (TLS 1.2, TLS 1.3)
- `certificate` — contains server's public key + domain + CA signature
- `CA (Certificate Authority)` — trusted organization that signs certificates (Let's Encrypt, DigiCert)
- `public key` — shared openly; used to encrypt / verify signatures
- `private key` — kept secret on server; used to decrypt / sign
- `asymmetric encryption` — different keys for encrypt/decrypt (RSA); slow; used for key exchange only
- `symmetric encryption` — same key for encrypt/decrypt (AES); fast; used for actual data
- `TLS handshake` — negotiate cipher suite, exchange keys, authenticate server
- `session key` — symmetric key generated during handshake; used for all subsequent data
- `cipher suite` — agreed combination of algorithms (e.g., TLS_AES_256_GCM_SHA384)
- `SNI (Server Name Indication)` — client tells server which domain it wants in handshake
- `HSTS (HTTP Strict Transport Security)` — header telling browser to only use HTTPS for this domain
- `certificate pinning` — app hardcodes expected cert/public key; rejects others
- `mutual TLS (mTLS)` — both client and server authenticate with certificates
- `Let's Encrypt` — free, automated CA; 90-day certs; widely used

**TLS 1.3 Handshake (simplified):**
```
Client → ServerHello + KeyShare
Server → Certificate + Finished
Client → Finished
Both derive session key
Data transfer (encrypted with AES)
```

**Must understand:**
- [ ] Why asymmetric encryption is used only for key exchange (it's slow)
- [ ] What a certificate actually contains (public key + domain + CA signature)
- [ ] What certificate chain of trust means
- [ ] HTTPS doesn't mean the website is safe — just that connection is encrypted
- [ ] mTLS — how it's used in microservices for service-to-service auth

---

## WebSockets

**Key Terms:**
- `WebSocket` — full-duplex persistent TCP connection between client and server
- `HTTP upgrade` — WebSocket starts as HTTP GET with `Upgrade: websocket` header
- `101 Switching Protocols` — server responds to accept upgrade
- `full-duplex` — both sides can send independently at any time (unlike HTTP request-response)
- `frame` — WebSocket data unit (text frame, binary frame, ping/pong frame)
- `ping / pong` — heartbeat frames; keep connection alive; detect dead connections
- `close frame` — graceful connection termination
- `connection per user` — server must track active connections (memory-intensive)
- `horizontal scaling` — WebSocket breaks stateless assumption; use sticky sessions or pub/sub
- `Redis pub/sub` — publish messages to Redis; all servers subscribed deliver to their connected users

**When to use WebSocket:**
- Real-time bidirectional: chat, multiplayer games, collaborative editing
- Server needs to push frequently: live scores, trading prices

**WebSocket vs SSE vs Long Polling:**

| | WebSocket | SSE | Long Polling |
|--|-----------|-----|--------------|
| Direction | Bidirectional | Server → Client | Server → Client |
| Protocol | WS (TCP upgrade) | HTTP | HTTP |
| Connection | Persistent | Persistent | Reconnect each message |
| Browser support | All | All (except IE) | All |
| Complexity | High | Low | Low |

**Must understand:**
- [ ] WebSocket handshake — how it upgrades from HTTP
- [ ] Why WebSocket complicates horizontal scaling (stateful connection)
- [ ] How to scale WebSocket: Redis pub/sub, or dedicated WS servers
- [ ] When SSE is a simpler alternative to WebSocket

---

## REST vs gRPC vs GraphQL

**Key Terms:**
- `REST` — HTTP + JSON; resource-based URLs; stateless; cacheable
- `gRPC` — Google RPC; uses HTTP/2; Protobuf serialization; strongly typed; bidirectional streaming
- `Protobuf (Protocol Buffers)` — binary serialization format; smaller and faster than JSON
- `IDL (Interface Definition Language)` — .proto files define service contracts
- `code generation` — gRPC generates client/server code from .proto definitions
- `service contract` — strict API definition; both sides agree on types
- `GraphQL` — query language; client specifies exactly what data it needs; single endpoint
- `resolver` — function that fetches data for a GraphQL field
- `N+1 in GraphQL` — each field resolver makes a DB call; use DataLoader (batching)
- `over-fetching` — REST returns too much data; GraphQL solves this
- `under-fetching` — REST requires multiple requests; GraphQL solves with single query
- `schema stitching / federation` — combine multiple GraphQL APIs into one

**When to use:**

| | REST | gRPC | GraphQL |
|--|------|------|---------|
| Client-server web/mobile | ✅ | ❌ (no native browser support) | ✅ |
| Service-to-service (backend) | ✅ | ✅ (preferred) | ❌ |
| Streaming | ❌ | ✅ (bidirectional) | ❌ |
| Strong typing | ❌ | ✅ | ✅ |
| Client-driven queries | ❌ | ❌ | ✅ |
| Caching | ✅ (HTTP cache) | ❌ | ❌ (complex) |

**Must understand:**
- [ ] Why gRPC is preferred for internal service-to-service (fast, typed, streaming)
- [ ] Why GraphQL helps mobile clients (control over what data to fetch)
- [ ] Protobuf vs JSON — binary, smaller, faster, but not human-readable
- [ ] GraphQL N+1 problem and DataLoader solution

---

## Reverse Proxy & API Gateway

**Key Terms:**
- `reverse proxy` — sits in front of servers; forwards client requests to backend
- `forward proxy` — sits in front of clients; forwards client requests to internet (VPN)
- `Nginx` — popular reverse proxy / web server / load balancer
- `HAProxy` — high-performance TCP/HTTP load balancer
- `API gateway` — reverse proxy with extra features: auth, rate limiting, routing, transformation
- `Kong` — popular open-source API gateway
- `AWS API Gateway` — managed API gateway on AWS
- `request routing` — route /api/users → user service, /api/orders → order service
- `SSL termination` — gateway handles TLS; backend communicates over plain HTTP (simpler)
- `auth offloading` — gateway validates JWT; backend trusts header with user ID
- `rate limiting at gateway` — centralized rate limiting for all services
- `service mesh` — alternative to API gateway for east-west (service-to-service) traffic; Istio, Linkerd

**Must understand:**
- [ ] Reverse proxy vs load balancer vs API gateway (they overlap; API GW is a superset)
- [ ] SSL termination — why it's done at the gateway
- [ ] How API gateway enables auth offloading (backend doesn't need to verify JWT)
- [ ] East-west traffic (service-to-service) vs north-south traffic (client-to-server)

---

## OSI Model

**Key Terms:**
- `OSI model` — 7-layer framework describing network communication
- `Layer 1 — Physical` — cables, fiber, radio waves, bits
- `Layer 2 — Data Link` — MAC addresses, Ethernet, switches
- `Layer 3 — Network` — IP addresses, routing, packets (IPv4, IPv6)
- `Layer 4 — Transport` — TCP/UDP, ports, segments
- `Layer 5 — Session` — session management (TLS lives here conceptually)
- `Layer 6 — Presentation` — encoding, encryption, compression
- `Layer 7 — Application` — HTTP, DNS, SMTP, FTP

**Practical mapping:**

| Technology | OSI Layer |
|------------|-----------|
| Ethernet / Wi-Fi | L2 |
| IP addresses, routing | L3 |
| TCP / UDP | L4 |
| TLS | L4-L6 |
| HTTP / WebSocket / DNS | L7 |
| Load balancer L4 | L4 (TCP) |
| Load balancer L7 | L7 (HTTP) |

**Must understand:**
- [ ] L4 vs L7 load balancer — what information each can inspect
- [ ] Why IP is L3 and TCP is L4 (IP routes, TCP provides reliability)
- [ ] A firewall can filter at L3 (IP) or L4 (port) or L7 (URL, content)
- [ ] In interviews: "at what layer does X work?" is common

---

## CDN

**Key Terms:**
- `CDN (Content Delivery Network)` — distributed network of edge servers caching content close to users
- `edge server / PoP (Point of Presence)` — CDN server in a geographic region
- `origin server` — your actual server where original content lives
- `cache hit (CDN)` — request served from edge; no trip to origin
- `cache miss (CDN)` — edge fetches from origin, caches, serves
- `Cache-Control header` — tells CDN how long to cache (max-age, no-store, public, private)
- `CDN invalidation / purge` — manually remove cached content from edge before TTL expires
- `pull CDN` — edge fetches from origin on first miss (automatic; most common)
- `push CDN` — you upload content to CDN proactively (for large known assets like video)
- `geo-routing` — DNS routes user to nearest edge (Anycast or geographic DNS)
- `Anycast` — same IP announced from multiple locations; network routes to nearest one
- `DDoS mitigation` — CDN absorbs attack traffic at edges before it reaches origin
- `edge computing` — run code at edge servers (Cloudflare Workers, Lambda@Edge)

**Providers:**
- Cloudflare — CDN + DDoS + DNS + edge computing
- AWS CloudFront — tightly integrated with AWS services
- Fastly — real-time cache purge; used by GitHub, Stripe
- Akamai — enterprise; largest network

**Must understand:**
- [ ] Why CDN reduces latency (physics: closer = faster)
- [ ] Cache-Control headers — how to control CDN behavior
- [ ] What content should and shouldn't be on CDN (static yes; personalized no)
- [ ] How CDN invalidation works and when to use it (cache busting with file hash)
