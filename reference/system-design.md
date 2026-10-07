# 03 — System Design

> For every topic: know what the problem is, the solution, the trade-offs, and be ready to draw + explain it.

---

## Scalability

**Key Terms:**
- `vertical scaling (scale up)` — add more CPU/RAM to existing server; limited by hardware ceiling
- `horizontal scaling (scale out)` — add more servers; no hard ceiling; requires stateless design
- `stateless service` — server holds no session data; any request can go to any server
- `stateful service` — server holds state; client must always hit same server (sticky sessions)
- `sticky sessions` — load balancer routes same client to same server (breaks horizontal scale)
- `bottleneck` — single point that limits total throughput
- `throughput` — requests per second a system can handle
- `latency` — time for a single request to complete
- `SLA (Service Level Agreement)` — contracted performance guarantee (e.g., 99.9% uptime)
- `SLO (Service Level Objective)` — internal target (e.g., p99 latency < 200ms)
- `p50 / p95 / p99` — percentile latency; p99 = 99% of requests faster than this value
- `auto-scaling` — automatically add/remove servers based on load (AWS ASG, GCP MIG)

**Must understand:**
- [ ] When vertical scaling hits its limit
- [ ] Why statelessness is required for horizontal scaling
- [ ] Difference between latency and throughput — improving one can worsen the other
- [ ] What p99 latency means and why it matters more than average

---

## Load Balancing

**Key Terms:**
- `load balancer` — distributes incoming traffic across multiple backend servers
- `round-robin` — requests cycle through servers 1,2,3,1,2,3...
- `weighted round-robin` — servers with more capacity get more requests
- `least connections` — send to server with fewest active connections
- `IP hash` — hash client IP to always route same client to same server
- `consistent hashing` — distribute load across servers with minimal remapping when servers added/removed
- `health check` — LB periodically pings servers; removes unhealthy ones
- `Layer 4 LB` — routes based on TCP/IP (no inspect of HTTP content); faster
- `Layer 7 LB` — routes based on HTTP content (URL, headers, cookies); smarter
- `reverse proxy` — sits in front of servers; Nginx, HAProxy
- `active-passive failover` — passive takes over if active fails
- `active-active` — both handle traffic; higher utilization
- `SPOF (Single Point of Failure)` — if this component fails, everything fails; LB itself must be HA

**Must understand:**
- [ ] L4 vs L7 load balancing — what information each uses
- [ ] Why round-robin can be unfair (different request costs)
- [ ] How health checks enable automatic failover
- [ ] How to make the load balancer itself highly available (pair of LBs)

---

## Caching

**Key Terms:**
- `cache hit` — data found in cache → return immediately
- `cache miss` — data not in cache → fetch from source, store, return
- `hit rate` — % of requests served from cache; target > 80%
- `TTL (Time to Live)` — entry expires after N seconds; simple invalidation
- `cache-aside (lazy loading)` — app checks cache; miss → fetch DB → populate cache
- `write-through` — every write → cache AND DB simultaneously
- `write-behind (write-back)` — write to cache → async write to DB (risk: data loss)
- `read-through` — cache fetches from DB on miss (cache manages itself, not app)
- `cache invalidation` — removing/updating stale entries; "the hard problem"
- `eviction policy` — what to remove when cache is full (LRU, LFU, TTL)
- `thundering herd` — many cache misses at once all hit DB (use locking/promise)
- `cache stampede` — same as thundering herd
- `Redis` — most common cache: rich types, persistence, pub/sub
- `CDN` — edge cache for static assets (images, JS, CSS) close to users
- `in-process cache` — cache inside the app process (fastest, not shared across instances)
- `distributed cache` — shared cache across all instances (Redis, Memcached)

**Must understand:**
- [ ] Cache-aside flow step by step (read path + write path)
- [ ] Why cache invalidation is called "the hard problem"
- [ ] When NOT to cache (highly dynamic data, personalized data)
- [ ] Thundering herd — how to prevent (mutex lock, probabilistic early expiry)

---

## Databases — SQL vs NoSQL

**Key Terms:**
- `RDBMS` — relational database; data in tables with rows and columns (PostgreSQL, MySQL)
- `NoSQL` — non-relational; flexible schema; many types
- `document store` — JSON-like documents (MongoDB)
- `key-value store` — simple key → value (Redis, DynamoDB)
- `wide-column store` — rows with flexible columns (Cassandra, HBase)
- `graph database` — nodes and edges (Neo4j)
- `ACID` — Atomicity, Consistency, Isolation, Durability — SQL guarantee
- `BASE` — Basically Available, Soft-state, Eventually consistent — NoSQL trade-off
- `schema` — SQL: fixed schema enforced by DB; NoSQL: flexible/schemaless
- `JOIN` — SQL strength: combine related tables; NoSQL avoids joins (denormalized)
- `horizontal scaling` — NoSQL designed for it; SQL historically harder to shard
- `OLTP` — Online Transaction Processing; many small reads/writes (use SQL)
- `OLAP` — Online Analytical Processing; large aggregate queries (use columnar DB)

**When to use SQL:**
- Need ACID transactions
- Complex relationships and JOINs
- Data is structured and schema is stable
- Financial, inventory, user data

**When to use NoSQL:**
- Need to scale horizontally to massive data
- Flexible/evolving schema
- High write throughput
- Time-series, logs, catalogs, user sessions

**Must understand:**
- [ ] ACID properties and why they matter
- [ ] Why NoSQL sacrifices consistency for availability/scale (CAP)
- [ ] When denormalization is intentional (NoSQL avoids JOIN overhead)
- [ ] How to choose between SQL and NoSQL for a given system

---

## Database Indexing

**Key Terms:**
- `index` — data structure that speeds up lookups on a column
- `B-tree index` — default index type; O(log n) search, range queries supported
- `B+ tree` — leaf nodes linked; better for range scans; used by most RDBMS
- `hash index` — O(1) exact lookup; no range queries (used internally by some DBs)
- `primary index` — on primary key; clustered (data stored in key order)
- `secondary index` — on non-primary key; additional lookup needed
- `composite index` — index on multiple columns; column order matters
- `covering index` — index contains all columns needed by query → no heap lookup
- `full-text index` — for text search (inverted index)
- `partial index` — index with WHERE clause (index only subset of rows)
- `index selectivity` — how unique the values are; high selectivity = better index
- `index scan` — use index to find rows
- `table scan (seq scan)` — read entire table; happens when no suitable index
- `EXPLAIN` — SQL command to show query execution plan; use to find missing indexes
- `N+1 problem` — 1 query to get N items, then N queries for related data; use JOIN or eager load

**Must understand:**
- [ ] How B-tree index works (why it's O(log n))
- [ ] Why composite index column order matters (`(a,b)` index helps `WHERE a=?` but not `WHERE b=?`)
- [ ] What a covering index is and why it's faster
- [ ] How EXPLAIN output shows whether an index is used
- [ ] When NOT to add an index (write-heavy tables: each write must update all indexes)

---

## Database Replication

**Key Terms:**
- `primary (master)` — accepts all writes; source of truth
- `replica (slave/read replica)` — copy of primary; handles read traffic
- `replication lag` — time delay between write on primary and visibility on replica
- `eventual consistency` — replica will eventually catch up, but not immediately
- `synchronous replication` — primary waits for replica to confirm write; no lag, slower writes
- `asynchronous replication` — primary doesn't wait; lag exists, faster writes
- `failover` — when primary fails, promote a replica to new primary
- `read scaling` — add replicas to handle more read traffic; primary handles writes
- `write bottleneck` — replication only scales reads, not writes

**Must understand:**
- [ ] Why replication helps reads but not writes
- [ ] Replication lag — how to handle reading stale data (read-your-own-writes problem)
- [ ] Synchronous vs async replication trade-off (consistency vs performance)
- [ ] Failover process — promoting replica to primary

---

## Database Sharding

**Key Terms:**
- `sharding (horizontal partitioning)` — split data across multiple DB servers
- `shard key` — column used to determine which shard a row goes to
- `range-based sharding` — shard by range of values (e.g., user IDs 1-1M on shard 1)
- `hash-based sharding` — shard by hash of key (uniform distribution)
- `hotspot` — one shard receives disproportionate traffic (bad shard key choice)
- `shard rebalancing` — when adding shards, data must be moved; painful
- `consistent hashing` — minimize data movement when adding/removing shards
- `cross-shard query` — querying across multiple shards; expensive, complex
- `shard index` — lookup table mapping entity to its shard
- `celebrity problem` — some users have 10M followers; one shard overloaded

**Must understand:**
- [ ] Why sharding is a last resort (adds huge complexity)
- [ ] How to choose a shard key (avoid hotspots, enable co-location)
- [ ] Why cross-shard JOINs are problematic
- [ ] Consistent hashing — how it minimizes resharding cost

---

## Message Queues

**Key Terms:**
- `producer` — publishes messages to the queue
- `consumer` — reads and processes messages
- `broker` — the message queue server (Kafka, RabbitMQ, SQS)
- `topic` — named channel; producers publish to topics, consumers subscribe
- `queue` — messages consumed by one consumer; queue model
- `pub/sub` — messages broadcast to all subscribers; fan-out model
- `at-most-once` — message may be lost (fire and forget)
- `at-least-once` — message delivered ≥1 times; consumer may process duplicates
- `exactly-once` — no loss, no duplicates; hardest to guarantee
- `idempotent consumer` — handles duplicate messages safely (same result no matter how many times)
- `dead letter queue (DLQ)` — messages that failed processing N times are moved here
- `backpressure` — slow consumer causes queue to fill up; system slows producer
- `Kafka` — log-based, persistent, high-throughput, consumer groups, replay
- `RabbitMQ` — traditional queue, routing rules, better for task queues
- `SQS` — AWS managed queue

**Why use message queues:**
- Decouple services (producer and consumer are independent)
- Handle traffic spikes (queue absorbs burst, consumer processes at its pace)
- Async processing (email, notifications, image processing)
- Broadcast to multiple consumers (fan-out)

**Must understand:**
- [ ] At-least-once vs exactly-once and why exactly-once is hard
- [ ] Why idempotency matters with at-least-once delivery
- [ ] Kafka vs RabbitMQ — when to use each
- [ ] How queues help with traffic spikes (buffering)

---

## API Design — REST

**Key Terms:**
- `resource` — noun-based URL (/users, /orders, /products)
- `HTTP verbs` — GET (read), POST (create), PUT (full update), PATCH (partial update), DELETE
- `status codes` — 200 OK, 201 Created, 400 Bad Request, 401 Unauthorized, 403 Forbidden, 404 Not Found, 409 Conflict, 429 Too Many Requests, 500 Server Error
- `idempotent` — calling same request multiple times has same result (GET, PUT, DELETE are idempotent; POST is not)
- `stateless` — each request is self-contained; server holds no session
- `HATEOAS` — hypermedia links in response for navigation (rarely implemented in practice)
- `versioning` — URL path (/v1/users), query param (?version=1), or header (Accept-Version)
- `pagination` — limit/offset or cursor-based
- `cursor pagination` — encode position in opaque token; consistent under inserts/deletes
- `rate limiting` — throttle requests per client (returns 429)
- `content negotiation` — Accept: application/json
- `REST constraints` — uniform interface, stateless, cacheable, layered system

**Must understand:**
- [ ] Why PUT is idempotent but POST is not
- [ ] Cursor vs offset pagination — why cursor is better for dynamic data
- [ ] How to version an API without breaking existing clients
- [ ] What status codes to return for different error scenarios

---

## Rate Limiting

**Key Terms:**
- `token bucket` — bucket refills at rate R tokens/sec; requests consume tokens; allows burst
- `leaky bucket` — requests flow out at constant rate; smooths traffic; no burst
- `sliding window log` — log timestamps of requests; count in last N seconds; memory heavy
- `sliding window counter` — estimate using two buckets (current + previous); approximate
- `fixed window counter` — count per time window; boundary burst problem
- `rate limit key` — per user, per IP, per API key
- `429 Too Many Requests` — HTTP status for rate limit exceeded
- `Retry-After` — header telling client when to retry
- `distributed rate limiting` — use Redis for shared counter across multiple servers
- `client-side rate limiting` — client self-throttles to avoid being blocked

**Must understand:**
- [ ] Token bucket vs leaky bucket — which allows burst (token bucket)
- [ ] Why fixed window has boundary problem (double burst at window edge)
- [ ] How Redis + Lua script enables atomic distributed rate limiting
- [ ] Per-user vs per-IP vs per-endpoint rate limiting — when to use each

---

## CAP Theorem

**Key Terms:**
- `Consistency` — all nodes see the same data at the same time
- `Availability` — every request gets a response (not necessarily up-to-date)
- `Partition Tolerance` — system works even if network between nodes fails
- `network partition` — link between nodes drops; nodes can't communicate
- `CA` — consistent + available; only in single-node systems (no network, no partition)
- `CP` — consistent + partition-tolerant; may reject requests to maintain consistency (HBase, ZooKeeper)
- `AP` — available + partition-tolerant; may return stale data (Cassandra, DynamoDB, CouchDB)
- `PACELC` — extension: even without partition, trade-off between latency and consistency
- `strong consistency` — read always sees most recent write
- `eventual consistency` — reads may return stale data but will converge
- `linearizability` — strongest consistency; reads see writes in real-time order

**Must understand:**
- [ ] Why partition tolerance is mandatory for distributed systems (networks fail)
- [ ] So the real choice is CP vs AP (not CA)
- [ ] Examples: Cassandra = AP, HBase = CP, ZooKeeper = CP
- [ ] PACELC extends CAP — considers latency trade-off in normal operation

---

## Consistent Hashing

**Key Terms:**
- `hash ring` — circular key space 0 to 2^32; servers and keys placed on ring
- `virtual nodes (vnodes)` — each server gets multiple positions on ring; improves distribution
- `node lookup` — for a key, find next clockwise server on ring
- `minimal remapping` — when server added/removed, only 1/N keys move on average
- `hot spot prevention` — virtual nodes ensure even distribution even with unequal servers
- `rendezvous hashing` — alternative approach; each client consistently picks same server

**Why consistent hashing:**
- Regular hash: `server = hash(key) % N` — adding server reshuffles ALL keys (massive cache miss storm)
- Consistent hash: adding server only moves 1/N of keys

**Must understand:**
- [ ] Why `hash(key) % N` breaks when N changes
- [ ] How the ring works — find server by going clockwise from key's position
- [ ] Why virtual nodes solve the uneven distribution problem
- [ ] Where consistent hashing is used: CDNs, distributed caches, load balancers

---

## Microservices vs Monolith

**Key Terms:**
- `monolith` — single deployable unit; all code together
- `microservices` — system decomposed into small independent services
- `service discovery` — how services find each other (Consul, Kubernetes DNS, Eureka)
- `inter-service communication` — sync (HTTP/gRPC) or async (message queue)
- `API gateway` — single entry point; routes to correct service; handles auth/rate limiting
- `service mesh` — infrastructure layer for service-to-service communication (Istio, Linkerd)
- `distributed tracing` — track request across multiple services (trace ID propagated through all)
- `data isolation` — each service owns its own database; no shared DB
- `deployment independence` — each service deploys independently
- `fault isolation` — one service failing doesn't bring down the whole system
- `operational complexity` — microservices: harder to debug, deploy, monitor
- `Conway's Law` — system architecture mirrors the organization's communication structure

**Monolith advantages:** simpler dev, easy debugging, one deploy, shared memory

**Microservices advantages:** independent scaling, independent deployment, tech diversity, fault isolation

**Must understand:**
- [ ] When to start with a monolith (most new products should)
- [ ] Why shared databases between microservices are an anti-pattern
- [ ] Challenges of microservices: distributed transactions, tracing, service discovery
- [ ] Circuit breaker — how to handle a downstream service being down

---

## Real-time Communication

**Key Terms:**
- `polling` — client repeatedly asks server "any updates?" on a schedule; simple but wasteful
- `long polling` — client sends request; server holds until update or timeout; then client reconnects
- `WebSocket` — full-duplex persistent connection; both sides can send anytime
- `Server-Sent Events (SSE)` — one-direction push from server to client over HTTP; simpler than WS
- `WebSocket handshake` — starts as HTTP, upgrades with `Upgrade: websocket` header
- `heartbeat / ping-pong` — periodic messages to keep connection alive
- `fan-out` — send one message to many users (e.g., chat broadcast)
- `presence` — knowing which users are online/offline
- `connection state` — server must track active WS connections per user
- `horizontal scaling with WS` — sticky sessions or publish-subscribe between servers

**When to use each:**
- Polling: rare updates, simple implementation, slight delay OK
- Long polling: moderate frequency, no WebSocket support needed
- WebSocket: real-time bidirectional (chat, games, collaborative editing)
- SSE: server → client only (notifications, live feeds, progress bars)

**Must understand:**
- [ ] Why polling is wasteful (request even when no new data)
- [ ] WebSocket vs SSE — bidirectional vs server-to-client only
- [ ] How to scale WebSocket servers horizontally (pub/sub via Redis)
- [ ] Chat app design — how messages flow from sender to receiver

---

## Design Practice Problems

**You must be able to design these systems end-to-end:**

- [ ] **URL Shortener** — hash, collision handling, redirect, analytics
- [ ] **Rate Limiter** — token bucket, Redis, distributed
- [ ] **Pastebin / Code Share** — blob storage, unique ID, expiry
- [ ] **Twitter/News Feed** — fan-out on write vs read, timeline generation
- [ ] **Chat App** — WebSocket, message ordering, delivery guarantees
- [ ] **Notification System** — push, email, SMS; fan-out, retry, dedup
- [ ] **Search Autocomplete** — trie, prefix matching, ranking
- [ ] **Distributed Cache** — consistent hashing, eviction, replication
- [ ] **Web Crawler** — BFS, URL frontier, dedup, politeness
- [ ] **File Storage (Dropbox/S3)** — chunking, dedup, metadata, CDN

**Framework for answering any SD question:**
1. Clarify requirements — functional + non-functional (scale, latency, consistency)
2. Estimate scale — DAU, QPS, storage
3. High-level design — main components and data flow
4. Deep dive — pick 2-3 critical components and design in detail
5. Trade-offs — explain what you chose and why
