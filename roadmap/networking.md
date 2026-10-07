# Networking Roadmap

> **L1** = aware & can use · **L2** = deep & efficient · Status: ⬜ not started · 🔄 in progress · ✅ done
> **Notes:** `notes/networking/` · **Practice:** `code/networking/` · **Dictionary:** `reference/networking.md`

**Why (EN):** Every request a backend serves crosses the network, and most "random" timeouts, resets and 502s are really DNS, TCP, TLS, proxy or load balancer problems. A senior can trace the full path of a request from client to database and explain where the milliseconds and the failures come from. In the AI era, AI can explain a protocol, but it cannot see your packets — the skill that matters more is hands-on debugging with `dig`, `curl -v`, `ss` and `tcpdump`, and choosing the right timeouts, protocols and topology.
**Зачем (RU):** Каждый запрос, который обслуживает бэкенд, проходит через сеть, и большинство «случайных» таймаутов, сбросов соединения и 502 — это на самом деле проблемы DNS, TCP, TLS, прокси или балансировщика. Senior умеет проследить весь путь запроса от клиента до базы данных и объяснить, где теряются миллисекунды и откуда берутся ошибки. В эпоху ИИ модель может объяснить протокол, но не видит ваши пакеты — важнее становится практическая отладка с `dig`, `curl -v`, `ss` и `tcpdump` и правильный выбор таймаутов, протоколов и топологии.

## Progress
| Layer | Done | Total |
|---|---|---|
| L1 | 0 | 50 |
| L2 | 0 | 32 |

---

## NET 1 — Models & the big picture
**L1 — aware & can use**
- ⬜ **1.1 OSI & TCP/IP models** — 7 OSI layers, 4 TCP/IP layers, encapsulation / decapsulation, PDU names (frame, packet, segment, datagram), headers, payload, which device works on which layer
- ⬜ **1.2 "What happens when you type a URL"** — URL parsing, browser cache, DNS lookup, TCP handshake, TLS handshake, HTTP request, load balancer, app server, database, response, rendering
- ⬜ **1.3 Core vocabulary** — bandwidth, throughput, latency, RTT, jitter, packet loss, MTU, hop, client/server, peer-to-peer, unicast / broadcast / multicast / anycast

**L2 — deep & efficient**
- ⬜ **1.4 Latency budget** — speed of light in fiber, propagation vs transmission vs queuing vs processing delay, bandwidth-delay product, "latency numbers every programmer should know", RTT count per request (DNS + TCP + TLS + HTTP), tail latency

🛠 **Practice** — L1: run `curl -w "@format.txt" -o /dev/null -s https://example.com` with `time_namelookup`, `time_connect`, `time_appconnect`, `time_starttransfer` and map each number to a step of "type a URL" · L2: measure the same for 3 sites on different continents, compute how many RTTs each phase costs, and explain the differences

---

## NET 2 — Link layer
**L1 — aware & can use**
- ⬜ **2.1 Ethernet & MAC** — MAC address, Ethernet frame, NIC, switch, hub (historical), MAC address table, broadcast domain, collision domain, full duplex
- ⬜ **2.2 ARP** — ARP request/reply, ARP cache (`ip neigh`), gratuitous ARP, broadcast `ff:ff:ff:ff:ff:ff`, IPv6 Neighbor Discovery (NDP) as replacement

**L2 — deep & efficient**
- ⬜ **2.3 MTU & L2 features** — MTU 1500, jumbo frames, fragmentation vs Path MTU Discovery, PMTU black holes, MSS clamping, VLAN (802.1Q), overlay encapsulation overhead (VXLAN, Geneve), Wi-Fi vs wired latency, ARP spoofing

🛠 **Practice** — L1: run `ip link`, `ip neigh`, then capture ARP with `sudo tcpdump -i any -n arp` while pinging a new LAN host · L2: inside two Docker containers on a custom network, lower the MTU on one interface and reproduce a "small requests work, big ones hang" issue with `ping -M do -s 1472`

---

## NET 3 — IP addressing & subnetting
**L1 — aware & can use**
- ⬜ **3.1 IPv4** — 32-bit address, dotted decimal, network vs host part, subnet mask, CIDR notation (`/24`), private ranges (10/8, 172.16/12, 192.168/16), loopback 127.0.0.0/8, link-local 169.254/16, `0.0.0.0`
- ⬜ **3.2 IP header & ICMP** — TTL, protocol field, source/destination, checksum, ICMP echo (ping), ICMP destination unreachable, time exceeded, why blocking all ICMP is harmful
- ⬜ **3.3 IPv6** — 128-bit address, hex notation & `::` compression, `/64` subnets, link-local `fe80::/10`, global unicast, SLAAC, dual stack, `::1`, Happy Eyeballs

**L2 — deep & efficient**
- ⬜ **3.4 Subnetting & IP planning** — subnet math, network/broadcast address, usable hosts, supernetting, VPC CIDR planning, overlapping CIDRs problem, reserved addresses in cloud subnets, IP exhaustion in k8s pod networks, CGNAT range 100.64/10

🛠 **Practice** — L1: write a small Go CLI using `net/netip` that prints network, broadcast and host range for a CIDR · L2: plan non-overlapping CIDRs for 3 environments × 3 AZs × public/private subnets + k8s pod/service ranges, and validate them with your Go CLI

---

## NET 4 — Routing & NAT
**L1 — aware & can use**
- ⬜ **4.1 Routing basics** — router, default gateway, routing table, longest prefix match, next hop, `ip route`, static vs dynamic routing, `traceroute` / `mtr`
- ⬜ **4.2 NAT** — SNAT, DNAT, masquerading, PAT (port address translation), port forwarding, public vs private IP, NAT table, why NAT breaks incoming connections

**L2 — deep & efficient**
- ⬜ **4.3 Internet routing** — autonomous system (AS), BGP (awareness), peering vs transit, IXP, anycast routing, BGP leaks/hijacks as outage sources, ECMP
- ⬜ **4.4 NAT in practice** — conntrack table, conntrack exhaustion, NAT timeouts killing idle connections, NAT traversal (STUN, TURN, ICE), hairpin NAT, SNAT port exhaustion behind cloud NAT gateways

🛠 **Practice** — L1: `ip route get 8.8.8.8`, `mtr -rw` to 3 destinations, explain each hop · L2: in Docker, inspect `iptables -t nat -L -n -v` (or `nft list ruleset`) for a published port and `conntrack -L` while curling it; map rules to the DNAT path

---

## NET 5 — TCP
**L1 — aware & can use**
- ⬜ **5.1 Connection lifecycle** — 3-way handshake (SYN, SYN-ACK, ACK), 4-way close (FIN), half-close, RST, TCP states (LISTEN, ESTABLISHED, CLOSE_WAIT, TIME_WAIT), connection refused vs timeout
- ⬜ **5.2 Reliability** — byte stream (no message boundaries), sequence numbers, acknowledgments, retransmission, ordering, checksum, duplicate ACKs, TCP keepalive
- ⬜ **5.3 Flow & congestion control (basics)** — receive window, sliding window, slow start, congestion window, packet loss as signal, why a new connection is slow

**L2 — deep & efficient**
- ⬜ **5.4 Congestion control in depth** — CUBIC, BBR, fast retransmit, fast recovery, SACK, window scaling, RTO calculation, bufferbloat, `net.ipv4.tcp_congestion_control`
- ⬜ **5.5 Production TCP problems** — TIME_WAIT accumulation, CLOSE_WAIT leaks (not closing in app), SYN flood & SYN cookies, listen backlog / `somaxconn` overflow, Nagle's algorithm + delayed ACK (`TCP_NODELAY`), keepalive vs idle timeouts on LBs, RST on idle connections
- ⬜ **5.6 Header-level reading** — TCP header fields, flags (SYN, ACK, FIN, RST, PSH, URG), options (MSS, SACK, timestamps, window scale), reading TCP flows in Wireshark (Follow TCP Stream, I/O graph)

🛠 **Practice** — L1: write a TCP echo server in Go (`net.Listen`, `Accept`, goroutine per conn), connect with `nc`, capture the handshake and close with `tcpdump -i lo port 9000` · L2: make the server forget to `Close()` connections and observe CLOSE_WAIT with `ss -tanp`; then add `tc qdisc ... netem delay 100ms loss 1%` and compare throughput with `iperf3`

---

## NET 6 — UDP & QUIC
**L1 — aware & can use**
- ⬜ **6.1 UDP** — connectionless, datagram, message boundaries, no ordering/retransmission, low overhead, use cases (DNS, VoIP, games, video, metrics like StatsD), broadcast/multicast
- ⬜ **6.2 QUIC basics** — QUIC over UDP, built-in TLS 1.3, streams without HOL blocking, 0-RTT, connection ID, connection migration, HTTP/3

**L2 — deep & efficient**
- ⬜ **6.3 UDP & QUIC in practice** — reliability on top of UDP, packet loss handling, amplification attacks, UDP buffer sizes, QUIC loss recovery, middlebox/firewall issues, `quic-go`, when QUIC is slower (CPU cost)

🛠 **Practice** — L1: write a UDP echo server/client in Go (`net.ListenPacket`), send messages bigger than MTU and observe what happens · L2: run a `quic-go` HTTP/3 example server and compare `curl --http3` (if your curl supports it) vs HTTP/2 under `netem` packet loss

---

## NET 7 — Sockets & Go networking
**L1 — aware & can use**
- ⬜ **7.1 Ports & sockets** — port, well-known ports, socket, 5-tuple, listen socket vs connection socket, ephemeral port range, `bind` / `listen` / `accept` / `connect`, `ss -tlnp`, `lsof -i`
- ⬜ **7.2 Go `net` package** — `net.Listen`, `net.Dial`, `net.Conn`, `DialContext`, `net.Dialer` (`Timeout`, `KeepAlive`), `SetDeadline`, `bufio` for framing, `net/http` server and client basics
- ⬜ **7.3 Timeouts** — connect timeout, read/write timeout, idle timeout, `http.Server` (`ReadHeaderTimeout`, `ReadTimeout`, `WriteTimeout`, `IdleTimeout`), `http.Client.Timeout`, context deadlines, default no-timeout danger

**L2 — deep & efficient**
- ⬜ **7.4 Socket internals** — file descriptors, `ulimit -n`, socket buffers (`SO_RCVBUF`/`SO_SNDBUF`), `SO_REUSEADDR`, `SO_REUSEPORT`, listen backlog, non-blocking I/O, epoll, Go netpoller
- ⬜ **7.5 Connection pooling** — `http.Transport` (`MaxIdleConns`, `MaxIdleConnsPerHost`, `MaxConnsPerHost`, `IdleConnTimeout`), reading/closing body to reuse connections, ephemeral port exhaustion, DB connection pools, pool vs LB idle timeout mismatch
- ⬜ **7.6 Framing & protocols on TCP** — length-prefixed framing, delimiter framing, `encoding/binary`, protocol design (versioning, heartbeats), backpressure, slow clients, graceful connection draining

🛠 **Practice** — L1: add read/write deadlines and a max-connections limit to your Go echo server; test it with a client that connects and never sends · L2: write a Go HTTP client load loop that forgets to read the body, watch ephemeral ports / TIME_WAIT explode with `ss -s`, then fix it and tune `http.Transport`

---

## NET 8 — DNS
**L1 — aware & can use**
- ⬜ **8.1 Resolution** — stub resolver, recursive resolver, root servers, TLD servers, authoritative servers, iterative vs recursive query, caching, TTL, negative caching, `/etc/resolv.conf`, `/etc/hosts`, `nsswitch.conf`
- ⬜ **8.2 Record types** — A, AAAA, CNAME, MX, TXT, NS, SOA, SRV, PTR (reverse DNS), CAA, CNAME at apex problem, ALIAS/ANAME
- ⬜ **8.3 Tools** — `dig` (`+trace`, `+short`, `@server`), `nslookup`, `host`, `resolvectl`, `getent hosts`

**L2 — deep & efficient**
- ⬜ **8.4 DNS in production** — DNS as outage source, TTL strategy for migrations, DNS-based load balancing & failover, GeoDNS, split-horizon DNS, Kubernetes `ndots:5` and search domains, CoreDNS, NodeLocal DNSCache, Go resolver (pure Go vs cgo, `GODEBUG=netdns`)
- ⬜ **8.5 DNS security & transport** — DNS over UDP/TCP (truncation), EDNS0, DNSSEC (awareness), DoH / DoT, DNS cache poisoning, subdomain takeover, dangling CNAME

🛠 **Practice** — L1: `dig +trace` a domain and draw the resolution path; query A, AAAA, MX, TXT, NS, CAA for 3 real domains · L2: in a kind cluster, exec into a pod, inspect `/etc/resolv.conf`, capture DNS queries with `tcpdump` while resolving an external name, and explain the extra queries caused by `ndots:5`

---

## NET 9 — HTTP/1.1
**L1 — aware & can use**
- ⬜ **9.1 Messages** — request line, status line, headers, body, methods (GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS), safe vs idempotent methods, `Host` header, `Content-Type`, `Content-Length`
- ⬜ **9.2 Status codes** — 1xx, 2xx (200, 201, 204), 3xx (301, 302, 304, 307, 308), 4xx (400, 401, 403, 404, 409, 422, 429), 5xx (500, 502, 503, 504), which are retryable
- ⬜ **9.3 Connections** — persistent connections (keep-alive), `Connection: close`, chunked transfer encoding, pipelining (and why it failed), head-of-line blocking, `curl -v`

**L2 — deep & efficient**
- ⬜ **9.4 HTTP edge cases** — request smuggling (CL.TE / TE.CL), `Expect: 100-continue`, trailers, range requests, large uploads/streaming, `Retry-After`, idempotency keys, 502 vs 504 root causes, hop-by-hop headers

🛠 **Practice** — L1: write a raw HTTP/1.1 request by hand with `nc example.com 80` (and `openssl s_client` for HTTPS), then compare with `curl -v` · L2: build a Go server that streams a chunked response and another that is slow; put both behind nginx and reproduce a 502 and a 504, explaining each

---

## NET 10 — HTTP/2 & HTTP/3
**L1 — aware & can use**
- ⬜ **10.1 HTTP/2** — binary framing, streams, multiplexing over one connection, HPACK header compression, ALPN negotiation, `h2` vs `h2c`, server push (deprecated)
- ⬜ **10.2 HTTP/3** — HTTP over QUIC, QPACK, no TCP HOL blocking, `Alt-Svc` header, browser/CDN support, fallback to HTTP/2

**L2 — deep & efficient**
- ⬜ **10.3 HTTP/2 in production** — flow control windows, stream priorities, `SETTINGS_MAX_CONCURRENT_STREAMS`, GOAWAY, TCP-level HOL blocking still exists, rapid reset attack (CVE-2023-44487), single connection = poor L4 load balancing, gRPC over HTTP/2

🛠 **Practice** — L1: run a Go HTTPS server, check protocol with `curl -v --http2` and inspect frames in Wireshark using `SSLKEYLOGFILE` · L2: send 100 concurrent requests over HTTP/1.1 vs HTTP/2 from a Go client against the same server, count TCP connections with `ss` and compare latency

---

## NET 11 — HTTP features: caching, cookies, CORS, compression
**L1 — aware & can use**
- ⬜ **11.1 Caching** — `Cache-Control` (`max-age`, `no-cache`, `no-store`, `private`, `public`, `s-maxage`), `ETag`, `If-None-Match`, `Last-Modified`, 304 Not Modified, `Vary`, browser vs shared cache
- ⬜ **11.2 Cookies & state** — `Set-Cookie`, `HttpOnly`, `Secure`, `SameSite` (Lax/Strict/None), `Domain`, `Path`, `Max-Age`, session cookies, third-party cookies
- ⬜ **11.3 CORS** — same-origin policy, origin, simple vs preflight request, `OPTIONS`, `Access-Control-Allow-Origin`, `-Allow-Credentials`, `-Allow-Headers`, what CORS does NOT protect

**L2 — deep & efficient**
- ⬜ **11.4 Content & performance** — compression (gzip, Brotli, zstd), `Accept-Encoding`, content negotiation, `stale-while-revalidate`, `stale-if-error`, cache stampede, redirects (301 vs 308 caching), HSTS, security headers (CSP, `X-Content-Type-Options`)

🛠 **Practice** — L1: add `ETag` + `Cache-Control` to a Go endpoint and show a 304 with `curl -H 'If-None-Match: ...'`; add CORS middleware and trigger a preflight from a small HTML page · L2: put nginx caching in front of the Go service with `stale-while-revalidate`-like behavior (`proxy_cache_use_stale`), then stop the backend and show stale responses still served

---

## NET 12 — Real-time & API protocols
**L1 — aware & can use**
- ⬜ **12.1 REST vs gRPC vs GraphQL** — resources & verbs, JSON over HTTP, gRPC (protobuf, HTTP/2, unary & streaming), GraphQL (single endpoint, over/under-fetching), when to use which
- ⬜ **12.2 Real-time options** — short polling, long polling, Server-Sent Events (SSE), WebSocket, WebSocket upgrade handshake (`Upgrade`, `101 Switching Protocols`), WebTransport (awareness)

**L2 — deep & efficient**
- ⬜ **12.3 Long-lived connections at scale** — connection limits per server, sticky sessions, heartbeats/ping-pong, reconnection with backoff, fan-out via pub/sub (Redis, NATS), load balancer idle timeouts, gRPC load balancing (client-side, xDS), gRPC deadlines & status codes

🛠 **Practice** — L1: build a tiny Go chat with WebSocket (`github.com/coder/websocket` or gorilla) and the same feed with SSE; inspect the upgrade with `curl -v` · L2: run 2 instances behind nginx, make messages reach clients on both instances via Redis pub/sub, and handle nginx idle timeout with ping frames

---

## NET 13 — TLS & PKI
**L1 — aware & can use**
- ⬜ **13.1 TLS basics** — confidentiality, integrity, authentication, TLS 1.2 vs 1.3, handshake overview, cipher suite, symmetric vs asymmetric keys, forward secrecy, HTTPS on 443
- ⬜ **13.2 Certificates** — X.509, subject / SAN, issuer, chain of trust, root & intermediate CA, trust store, validity dates, self-signed, Let's Encrypt, ACME (HTTP-01, DNS-01), certificate expiry outages
- ⬜ **13.3 Tools** — `openssl s_client -connect -servername`, `openssl x509 -text -noout`, `curl -v`, `mkcert`, `certbot`, checking expiry dates

**L2 — deep & efficient**
- ⬜ **13.4 Handshake in depth** — TLS 1.3 1-RTT handshake, ECDHE key exchange, ClientHello/ServerHello, SNI, ALPN, session resumption (PSK, tickets), 0-RTT & replay risk, Encrypted Client Hello (ECH), post-quantum hybrid key exchange (X25519MLKEM768, awareness)
- ⬜ **13.5 mTLS & operations** — mutual TLS, client certificates, internal CA, SPIFFE/SPIRE, cert rotation, short-lived certs, OCSP / CRL / revocation problems, certificate transparency, HSTS preload, TLS termination vs passthrough vs re-encryption, Go `crypto/tls` config

🛠 **Practice** — L1: `openssl s_client -connect google.com:443 -servername google.com -showcerts`, decode each cert and draw the chain · L2: create an internal CA with `openssl` (or `step`), issue server + client certs, and build Go server/client with mTLS; capture the handshake in Wireshark and show what fails with an expired or wrong-SAN cert

---

## NET 14 — Proxies, load balancers & API gateways
**L1 — aware & can use**
- ⬜ **14.1 Proxies** — forward proxy, reverse proxy, transparent proxy, Nginx, HAProxy, Envoy, Caddy, Traefik, `X-Forwarded-For`, `X-Forwarded-Proto`, `Forwarded`, PROXY protocol
- ⬜ **14.2 Load balancing** — L4 vs L7, algorithms (round robin, least connections, weighted, IP hash, consistent hashing, P2C), health checks (active/passive), connection draining, sticky sessions
- ⬜ **14.3 API gateway** — routing, auth offload, rate limiting, request transformation, API gateway vs reverse proxy vs service mesh, Kong, Envoy Gateway, cloud API gateways

**L2 — deep & efficient**
- ⬜ **14.4 Proxy tuning & failure modes** — upstream keepalive, timeouts (connect/read/send/idle), buffering vs streaming, `proxy_next_upstream` retries & idempotency, retry storms, outlier detection, circuit breaking, 499 / 502 / 503 / 504 meanings, slow-start for new backends
- ⬜ **14.5 LB architecture** — DSR (direct server return), L4 LB with ECMP + consistent hashing (Maglev-style), global vs regional LB, GSLB, TLS termination placement, HTTP/2 & gRPC balancing, xDS control plane

🛠 **Practice** — L1: run nginx in Docker as a reverse proxy in front of 3 Go service replicas (compose), round-robin, with `X-Forwarded-For` logged by the app · L2: add passive health checks and `proxy_next_upstream`, kill one replica during a k6 run, and compare errors with and without retries; then repeat with Envoy or HAProxy

---

## NET 15 — CDN & edge
**L1 — aware & can use**
- ⬜ **15.1 CDN basics** — edge / PoP, origin, cache hit / miss ratio, TTL, cache key, purge / invalidation, static vs dynamic content, Cloudflare, CloudFront, Fastly

**L2 — deep & efficient**
- ⬜ **15.2 Edge architecture** — anycast, origin shield, tiered caching, cache key design (query string, headers, cookies), signed URLs, edge compute (Workers, Lambda@Edge), DDoS absorption, WAF at the edge, origin protection

🛠 **Practice** — L1: inspect `cf-cache-status` / `x-cache` / `age` headers on 3 popular sites with `curl -I` and explain hit vs miss · L2: design (diagram + 1 page) a CDN setup for an API with mixed public and per-user responses: cache keys, TTLs, purge strategy, what must never be cached

---

## NET 16 — Linux networking
**L1 — aware & can use**
- ⬜ **16.1 iproute2 basics** — `ip addr`, `ip link`, `ip route`, `ip neigh`, `ss` (`-tlnp`, `-tanp`, `-s`), `/etc/hosts`, `hostname -I`, `nmcli` (awareness)
- ⬜ **16.2 Firewalls** — iptables (tables, chains: INPUT, OUTPUT, FORWARD, PREROUTING, POSTROUTING), nftables, `ufw`, stateful filtering, conntrack, default policies

**L2 — deep & efficient**
- ⬜ **16.3 Virtual networking** — network namespaces (`ip netns`), veth pairs, Linux bridge, TUN/TAP, VXLAN, macvlan / ipvlan, how Docker bridge networking is built from these
- ⬜ **16.4 Kernel tuning & traffic control** — sysctls (`net.core.somaxconn`, `tcp_tw_reuse`, `ip_local_port_range`, `rmem/wmem`), `tc` + `netem` (delay, loss, rate), XDP & eBPF (awareness), `ethtool`, NIC offloads

🛠 **Practice** — L1: list listening sockets and their processes with `ss -tlnp`, add an nftables rule that blocks port 8080, and verify with `curl` · L2: build two network namespaces connected by a veth pair through a bridge, give them IPs, add NAT to reach the internet, and ping between them — i.e. rebuild Docker's bridge network by hand

---

## NET 17 — Container & Kubernetes networking
**L1 — aware & can use**
- ⬜ **17.1 Docker networking** — bridge (`docker0`), user-defined bridge with DNS, host, none, port publishing (`-p`), embedded DNS 127.0.0.11, compose networks
- ⬜ **17.2 Kubernetes networking model** — every pod has an IP, pod-to-pod without NAT, Service (ClusterIP, NodePort, LoadBalancer, headless), kube-proxy, CoreDNS, Ingress, Gateway API

**L2 — deep & efficient**
- ⬜ **17.3 CNI & dataplane** — CNI plugins (Calico, Cilium, Flannel), overlay vs native routing, kube-proxy iptables/IPVS/nftables modes, eBPF kube-proxy replacement, `externalTrafficPolicy: Local` & source IP preservation, NetworkPolicy
- ⬜ **17.4 Service mesh** — sidecar vs sidecarless (Istio ambient), Linkerd, Envoy, mTLS between services, retries/timeouts in mesh, traffic splitting, observability, mesh cost & complexity

🛠 **Practice** — L1: on a kind cluster, deploy 2 Go services, call one from the other via `svc.namespace.svc.cluster.local`, and trace the path with `kubectl exec` + `curl -v` · L2: inspect kube-proxy's iptables rules for a Service on a kind node (`docker exec` into the node), find the DNAT to pod IPs, and apply a NetworkPolicy that blocks one direction (with Calico or Cilium installed)

---

## NET 18 — Cloud networking
**L1 — aware & can use**
- ⬜ **18.1 VPC basics** — VPC, CIDR block, public vs private subnet, route table, internet gateway, NAT gateway, elastic / static IP, security group (stateful) vs NACL (stateless)
- ⬜ **18.2 Connectivity** — VPC peering, Transit Gateway, VPC endpoints (gateway & interface), PrivateLink, site-to-site VPN, Direct Connect / Cloud Interconnect (awareness), private DNS zones

**L2 — deep & efficient**
- ⬜ **18.3 Cloud network design** — hub-and-spoke, multi-account networking, egress control (NAT costs, egress proxies), cross-AZ data transfer cost, IP planning for EKS/GKE (VPC CNI, secondary ranges), flow logs, shared VPC

🛠 **Practice** — L1: draw a VPC for the Go service: 2 AZs, public subnets (ALB, NAT), private subnets (app, DB), with route tables and security group rules · L2: write that VPC in Terraform (apply against LocalStack or just `plan`), add a VPC endpoint for S3, and explain the egress cost difference

---

## NET 19 — VPN, remote access & zero trust
**L1 — aware & can use**
- ⬜ **19.1 VPN & remote access** — VPN concept, WireGuard, OpenVPN, IPsec (awareness), Tailscale / Headscale, bastion host, SSH tunnels (`-L`, `-R`, `-D`), split tunneling

**L2 — deep & efficient**
- ⬜ **19.2 Zero trust** — perimeter model vs zero trust, identity-aware proxy, BeyondCorp, device posture, mTLS everywhere, short-lived credentials, Cloudflare Access / Teleport / AWS SSM Session Manager (awareness), microsegmentation

🛠 **Practice** — L1: set up WireGuard between two Docker containers (or a container and the host) and route traffic through the tunnel; use `ssh -L` to reach a Postgres container that is not published · L2: write a 1-page design replacing a bastion host with an identity-aware proxy for DB access: who authenticates, where certs are issued, how access is audited

---

## NET 20 — Debugging toolbox & methodology
**L1 — aware & can use**
- ⬜ **20.1 Connectivity tools** — `ping`, `traceroute`, `mtr`, `nc` (`-zv`), `telnet`, `curl -v` (`--resolve`, `--connect-to`, `-w`), `wget`, `httpie`, `grpcurl`
- ⬜ **20.2 Inspection tools** — `dig`, `ss`, `netstat` (legacy), `lsof -i`, `ip`, `openssl s_client`, `nmap` (on your own hosts), `iperf3`
- ⬜ **20.3 Packet capture** — `tcpdump` filters (`host`, `port`, `tcp[tcpflags]`), `-w file.pcap`, Wireshark display filters, Follow Stream, decrypting TLS with `SSLKEYLOGFILE`, `tshark`, `termshark`

**L2 — deep & efficient**
- ⬜ **20.4 Debugging methodology** — bottom-up vs top-down (layer by layer), divide the path (client → DNS → LB → proxy → app → DB), reproduce with `curl`, compare working vs failing, capture on both sides, time correlation, common culprits checklist (DNS, MTU, timeouts mismatch, conntrack, cert expiry)
- ⬜ **20.5 Debugging in containers & k8s** — `nsenter -t <pid> -n`, `kubectl debug` with netshoot image, ksniff (`kubectl sniff`), ephemeral tcpdump container, capturing on node vs pod interface, Cilium Hubble, Wireshark over pcap from a pod

🛠 **Practice** — L1: capture a full HTTPS request to your Go service with `tcpdump -w`, open it in Wireshark, decrypt it using `SSLKEYLOGFILE`, and label DNS, TCP handshake, TLS handshake and HTTP frames · L2: ask a friend (or a script) to break your compose stack in one hidden way (wrong DNS, blocked port, MTU, expired cert, proxy timeout) and diagnose it using only the toolbox; write the diagnosis log

---

## NET 21 — Network reliability & performance for backend engineers
**L1 — aware & can use**
- ⬜ **21.1 Timeouts & retries** — timeout on every network call, timeout budgets across hops, retry with exponential backoff + jitter, retry only idempotent requests, retry budgets, `context` propagation in Go
- ⬜ **21.2 Common failure patterns** — connection refused, connection reset, timeout, DNS failure, TLS error, 502/503/504, partial failures, gray failures

**L2 — deep & efficient**
- ⬜ **21.3 Performance tuning** — connection reuse, keep-alive alignment (client < LB < server idle timeouts), TLS session resumption, HTTP/2 multiplexing, payload size & compression, batching, co-location (same AZ), measuring p99 network latency, hedged requests

🛠 **Practice** — L1: add per-call timeouts + exponential backoff with jitter to a Go HTTP client calling a flaky Go server (random 503s and delays) · L2: set client idle timeout > LB (nginx) idle timeout, reproduce the sporadic "connection reset" errors under k6 load, then fix by aligning timeouts and prove it with zero errors
