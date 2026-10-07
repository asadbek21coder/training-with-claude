# System Design & Distributed Systems Roadmap

> **L1** = aware & can use · **L2** = deep & efficient · Status: ⬜ not started · 🔄 in progress · ✅ done
> **Notes:** `notes/system-design/` · **Practice:** `code/system-design/` · **Dictionary:** `reference/system-design.md`

**Why (EN):** System design turns requirements into an architecture with explicit trade-offs, and it is what decides the senior level in interviews and in real work. Any backend with more than one instance, a queue or a cache is already a distributed system: networks fail, clocks lie, partial failure is normal. In the AI era, AI can draw a generic "boxes and arrows" diagram in seconds, so the value moves to judgment — clarifying requirements, doing the numbers, choosing between trade-offs, and predicting how the system fails; LLM-based features (RAG, agents) are now also a standard design question.
**Зачем (RU):** System design — это превращение требований в архитектуру с явными компромиссами; именно он определяет уровень senior на собеседованиях и в реальной работе. Любой backend с несколькими инстансами, очередью или кэшем уже распределённый: сеть падает, часы врут, частичный отказ — норма. В эпоху AI типовую схему «квадратики и стрелки» AI нарисует за секунды, поэтому ценность смещается в суждение — уточнить требования, посчитать цифры, выбрать компромисс и предсказать, как система сломается; LLM-фичи (RAG, агенты) теперь тоже стандартный вопрос на дизайне.

## Progress
| Layer | Done | Total |
|---|---|---|
| L1 | 8 | 40 |
| L2 | 0 | 56 |

---

## SD 1 — Framework & estimation
**L1 — aware & can use**
- ⬜ **1.1 Interview framework** — functional vs non-functional requirements, clarifying questions, scope & out-of-scope, API, data model, high-level design, deep dive, bottlenecks, wrap-up, time boxing (45/60 min)
- ⬜ **1.2 Estimation** — back-of-the-envelope, DAU/MAU, QPS, peak vs average (peak factor), read/write ratio, storage per day/year, bandwidth, latency numbers every programmer should know, powers of two, 1 day ≈ 10^5 seconds
- ⬜ **1.3 API sketching** — REST resources & endpoints, request/response shapes, cursor pagination, `Idempotency-Key` header, versioning, REST vs gRPC vs WebSocket, status codes, sync vs async (202 Accepted)
- ⬜ **1.4 Diagramming** — boxes-and-arrows, data flow, sequence diagram, C4 model (context, container, component), Mermaid, Excalidraw, marking sync vs async edges, read path vs write path

**L2 — deep & efficient**
- ⬜ **1.5 Senior signals** — stating assumptions, trade-off articulation ("this gives us X, costs us Y"), alternatives considered, depth over breadth, driving the conversation, handling pushback, failure modes first, operational concerns, numbers that justify decisions, evolving the design (v1 → v2)

🛠 **Practice** — L1: write `code/system-design/01-framework/framework.md` — your own one-page checklist of the framework + a latency/powers-of-two cheat sheet; estimate QPS, storage and bandwidth for "Instagram photo upload" on paper · L2: record yourself (audio) doing a 45-minute design of a URL shortener, listen back and list where you stated trade-offs and where you didn't

---

## SD 2 — Non-functional requirements
**L1 — aware & can use**
- ⬜ **2.1 Availability & reliability** — SLA, SLO, SLI, nines (99.9% ≈ 8.76 h/year, 99.99% ≈ 52.6 min/year), redundancy, failover, single point of failure, MTBF, MTTR
- ⬜ **2.2 Performance** — latency vs throughput, latency percentiles (p50/p95/p99), tail latency, latency budget, Little's law, utilization, response time vs service time
- ⬜ **2.3 Consistency & durability** — strong vs eventual consistency, read-your-writes, durability, data loss tolerance, RPO, RTO, backups, replication factor

**L2 — deep & efficient**
- ⬜ **2.4 Availability math & error budgets** — serial vs parallel availability composition, dependency availability, error budget, burn rate, SLO-based alerting, tail at scale (fan-out amplifies p99), hedged requests
- ⬜ **2.5 Scalability limits** — Amdahl's law, Universal Scalability Law, contention vs coherence cost, queueing theory basics (utilization → latency curve), capacity planning, headroom, load testing (k6)
- ⬜ **2.6 Cost & operability** — cost estimation (compute, storage, egress), egress cost, managed vs self-hosted, build vs buy, operational complexity, on-call burden, FinOps basics, cost per request
- ⬜ **2.7 Observability & security in design** — RED / USE metrics, distributed tracing, structured logs, correlation IDs, authn/authz placement, encryption in transit / at rest, secrets management, PII handling, threat modeling (STRIDE)

🛠 **Practice** — L1: for a service you've worked on, write `code/system-design/02-nfr/my-service-nfr.md` with SLIs, SLOs, p99 targets, RPO/RTO and a list of SPOFs · L2: in Go, write a tiny HTTP service + a k6 script; increase load step by step, plot p50/p99 vs RPS and mark where latency explodes (queueing knee); compute the composed availability of a request that hits 5 dependencies at 99.9% each

---

## SD 3 — Scaling & load balancing
**L1 — aware & can use**
- ✅ **3.1 Scaling** — vertical vs horizontal scaling, stateless services, sticky sessions, externalized session state (Redis), shared-nothing architecture, scale-out limits
- ✅ **3.2 Load balancing** — L4 vs L7, round robin, weighted round robin, least connections, IP hash, consistent hashing (idea), active vs passive health checks
- ⬜ **3.3 DNS & traffic entry** — DNS resolution, A/AAAA/CNAME records, DNS TTL, DNS load balancing, GeoDNS, anycast, reverse proxy (Nginx, Envoy, HAProxy), TLS termination

**L2 — deep & efficient**
- ⬜ **3.4 Load balancing deep** — power of two random choices, least outstanding requests, connection draining, slow start, outlier detection, LB high availability (active-passive, VRRP), balancing long-lived HTTP/2 & gRPC connections, client-side load balancing
- ⬜ **3.5 Autoscaling** — horizontal autoscaling, scaling metrics (CPU, RPS, queue depth, latency), scale-up lag, cold start, warm pools, predictive scaling, scale-in protection, Kubernetes HPA, KEDA
- ⬜ **3.6 Global traffic management** — GSLB, latency-based routing, geo routing, failover routing, health-check-based DNS failover, anycast, edge PoPs, DNS caching pitfalls

🛠 **Practice** — L1: run 3 instances of a Go HTTP service behind Nginx in docker compose, switch between round robin and least connections, kill one instance and watch health checks remove it · L2: write a Go client-side load balancer implementing round robin vs power-of-two-choices, simulate one slow backend and compare p99

---

## SD 4 — Caching & CDN
**L1 — aware & can use**
- ✅ **4.1 Caching basics** — cache-aside, read-through, write-through, write-back, write-around, TTL, LRU, LFU, cache invalidation, hit ratio
- ✅ **4.2 CDN basics** — edge, origin, PoP, push vs pull CDN, cache hit ratio, static vs dynamic content, TTL

**L2 — deep & efficient**
- ⬜ **4.3 Cache failure modes** — cache stampede / thundering herd, request coalescing (`golang.org/x/sync/singleflight`), probabilistic early expiration, lock on miss, cache penetration (negative caching, bloom filter), cache avalanche (TTL jitter), hot keys, cold start & cache warming
- ⬜ **4.4 Cache consistency** — invalidate vs update on write, delete-after-write, stale-set race after delete, delayed double delete, versioned keys, CDC-driven invalidation, leases (Facebook memcache paper), TTL as a safety net
- ⬜ **4.5 Cache topology** — in-process cache (ristretto, bigcache), two-tier local + distributed cache, near cache, replicated vs partitioned cache, Memcached vs Redis, consistent hashing across cache nodes, memory sizing, cache as SPOF
- ⬜ **4.6 HTTP & CDN caching deep** — `Cache-Control` (`max-age`, `s-maxage`, `no-store`, `private`), `ETag` / `If-None-Match`, `Vary`, cache key design, purge by tag (surrogate keys), `stale-while-revalidate`, `stale-if-error`, origin shield, signed URLs, edge compute

🛠 **Practice** — L1 (review): add Redis cache-aside to a Go handler and serve static files through Nginx with `Cache-Control`; measure hit ratio and p50 latency with/without cache · L2: in Go, put a cache in front of a deliberately slow function, fire 1000 concurrent requests at an expired key and count origin calls with and without `singleflight` and TTL jitter; reproduce the stale-set race between a writer and a reader in a test

---

## SD 5 — Data storage in design
**L1 — aware & can use**
- ✅ **5.1 SQL vs NoSQL & CAP basics** — relational vs document / key-value / wide-column / graph, schema flexibility, ACID vs BASE, CAP theorem (CP vs AP), access patterns, joins vs denormalization
- ✅ **5.2 Replication & sharding basics** — primary-replica, read replicas, replication lag, failover, sharding (horizontal partitioning), shard key, vertical partitioning
- ⬜ **5.3 Storage types in design** — object storage (S3), block vs file vs object storage, search index (Elasticsearch), time-series DB, data warehouse, cache, hot/warm/cold tiers, choosing by access pattern

**L2 — deep & efficient**
- ⬜ **5.4 Data modeling for scale** — denormalization, precomputed views, fan-out on write vs fan-out on read, CQRS, read models, counters at scale, append-only logs, archival & TTL, hot partitions
- ⬜ **5.5 Large objects & media** — presigned URLs, multipart / resumable upload, chunking, content-addressable storage (hash dedup), metadata DB vs blob store split, media processing pipeline, thumbnails, lifecycle policies

🛠 **Practice** — L1: for 6 data types in a chat app (users, messages, attachments, search, presence, analytics) write in `code/system-design/05-storage/choices.md` which storage you'd use and why · L2: in Go + MinIO (S3-compatible, Docker), implement an upload flow with presigned PUT URLs and a metadata row in Postgres, including resumable multipart upload

---

## SD 6 — Async processing & messaging basics
**L1 — aware & can use**
- ✅ **6.1 Message queue basics** — async processing, producer/consumer, broker, ack, at-most-once, at-least-once, exactly-once, idempotency, retries
- ⬜ **6.2 Event-driven basics** — pub/sub vs point-to-point queue, event vs command, fan-out, event notification vs event-carried state transfer, dead-letter queue (DLQ), competing consumers, background workers

**L2 — deep & efficient**
- ⬜ **6.3 Choosing a broker** — Kafka (log, replay, partitions) vs RabbitMQ (smart broker, routing) vs NATS / JetStream vs SQS/SNS vs Redis Streams vs Postgres-as-queue; throughput, ordering, retention, replay, ops cost
- ⬜ **6.4 Jobs & workflows** — delayed jobs, scheduled jobs (cron), priority queues, visibility timeout, job deduplication, long-running tasks & progress tracking, workflow engines (Temporal), Go job libraries (River, Asynq)

🛠 **Practice** — L1: in Go, build an "image resize" flow — HTTP handler enqueues a job (Postgres `SKIP LOCKED` or RabbitMQ in Docker), a worker pool processes it, failures go to a DLQ after 3 retries · L2: write `code/system-design/06-async/broker-choice.md` comparing Kafka, RabbitMQ, NATS JetStream and SQS for 3 concrete use cases with a decision table

---

## SD 7 — Rate limiting, API edge & other building blocks
**L1 — aware & can use**
- ✅ **7.1 Rate limiting algorithms** — token bucket, leaky bucket, fixed window counter, sliding window log, sliding window counter, burst, refill rate
- ⬜ **7.2 API gateway & reverse proxy** — API gateway, reverse proxy, routing, authentication at the edge, TLS termination, request aggregation, BFF (backend for frontend), WAF, Kong / Envoy / Nginx
- ⬜ **7.3 Other building blocks** — service discovery (client-side vs server-side, DNS, Consul), unique ID generator, object storage, search index, WebSocket gateway, distributed scheduler, config service, feature flags

**L2 — deep & efficient**
- ⬜ **7.4 Distributed rate limiting** — centralized counter (Redis + Lua), local + global limits, approximate counting, race conditions, GCRA, per-user / IP / API key / tenant limits, `429 Too Many Requests`, `Retry-After`, `RateLimit` headers, fail-open vs fail-closed
- ⬜ **7.5 Unique ID generation** — UUIDv4 vs UUIDv7, ULID, Snowflake (timestamp + worker ID + sequence), ticket server, k-sortable IDs, clock rollback, B-tree index locality, ID enumeration risk
- ⬜ **7.6 Real-time delivery** — short polling, long polling, Server-Sent Events, WebSockets, connection gateway, connection registry, heartbeats, presence, fan-out to connected clients, sticky routing, reconnect & resume (last event ID)

🛠 **Practice** — L1: implement token bucket and sliding window counter in Go as HTTP middleware and test them with `hey` or k6 · L2: make the limiter distributed (Redis + Lua) across 3 instances, then write a Snowflake ID generator in Go and a test that proves IDs stay unique and ordered across goroutines

---

## SD 8 — Classic designs I
**L1 — aware & can use**
- ⬜ **8.1 URL shortener** — base62 encoding, ID generation vs hashing, collisions, 301 vs 302 redirect, read-heavy caching, custom aliases, expiration, click analytics
- ⬜ **8.2 Rate limiter service** — algorithm choice, Redis counters, Lua atomicity, rules configuration, per-client keys, deployment (middleware vs sidecar vs gateway), headers, fail-open
- ⬜ **8.3 Notification service** — push (APNs, FCM), SMS, email, templates, user preferences, queue per channel, retries, deduplication, per-user rate limiting, delivery tracking
- ⬜ **8.4 News feed** — fan-out on write (push), fan-out on read (pull), hybrid for celebrities, feed cache, ranking, cursor pagination, social graph storage, media in feed
- ⬜ **8.5 Search autocomplete** — trie, top-k per prefix, prefix cache, data collection pipeline, sampling, client debouncing, ranking, offensive-term filtering, CDN/browser caching

**L2 — deep & efficient**
- ⬜ **8.6 Web crawler** — URL frontier, politeness (`robots.txt`, per-host rate), BFS vs priority crawling, URL dedup (seen set, bloom filter), content dedup (hash, SimHash), DNS cache, distributed workers, crawler traps, recrawl scheduling

🛠 **Practice** — L1: for each design write `code/system-design/08-<design>/design.md` following the framework in `code/system-design/README.md` (requirements → estimates → API → data model → diagram → deep dives → failure modes → trade-offs) · L2: same for the web crawler, plus a small Go prototype of a polite per-host frontier

---

## SD 9 — Distributed systems fundamentals
**L1 — aware & can use**
- ⬜ **9.1 Why distributed is hard** — fallacies of distributed computing, partial failure, network partition, split brain, unreliable networks, unreliable clocks, process pauses (GC, VM), "slow" vs "dead"
- ⬜ **9.2 CAP & PACELC** — consistency, availability, partition tolerance, CP vs AP systems, PACELC (else: latency vs consistency), examples (Postgres, Cassandra, DynamoDB, etcd), common CAP misconceptions

**L2 — deep & efficient**
- ⬜ **9.3 Failure models & detection** — crash-stop, crash-recovery, omission, Byzantine faults, timeouts as failure detectors, heartbeats, phi accrual failure detector, gossip protocols (SWIM), gray failure
- ⬜ **9.4 Theoretical foundations** — safety vs liveness, two generals problem, Byzantine generals, FLP impossibility, system models (synchronous, partially synchronous, asynchronous), end-to-end argument

🛠 **Practice** — L1: run a 3-node etcd cluster in docker compose, write/read a key, `docker network disconnect` one node and observe which side still accepts writes; write what you saw in CAP terms · L2: in Go, implement a heartbeat-based failure detector for 3 processes, add random latency with `tc netem` and measure false positives vs detection time for different timeouts

---

## SD 10 — Replication & consistency
**L1 — aware & can use**
- ⬜ **10.1 Replication models** — leader-follower, multi-leader, leaderless, sync vs async vs semi-sync, replication lag, failover, write path vs read path

**L2 — deep & efficient**
- ⬜ **10.2 Replication lag anomalies** — read-your-writes, monotonic reads, consistent prefix reads, session guarantees, reading from leader after write, LSN / version-based reads
- ⬜ **10.3 Leaderless & quorums** — quorum (R + W > N), sloppy quorum, hinted handoff, read repair, anti-entropy, Merkle trees, tunable consistency (Cassandra `ONE` / `QUORUM` / `ALL`), Dynamo-style systems
- ⬜ **10.4 Consistency models** — linearizability, sequential consistency, causal consistency, eventual consistency, strong eventual consistency, serializability vs linearizability, strict serializability, consistency vs isolation
- ⬜ **10.5 Conflict resolution** — write conflicts in multi-leader, last-write-wins (LWW), version vectors, siblings, CRDTs (G-Counter, PN-Counter, OR-Set, LWW-Register), operational transformation, application-level merge

🛠 **Practice** — L1: run Postgres primary + async replica in Docker, write on the primary and read immediately on the replica in a Go loop; measure how often you read stale data · L2: implement an in-memory leaderless KV store in Go (3 nodes as goroutines or containers) with configurable N/R/W, inject node failures, show stale reads when R + W ≤ N, and add a G-Counter CRDT that merges correctly after a partition

---

## SD 11 — Partitioning
**L1 — aware & can use**
- ⬜ **11.1 Partitioning strategies** — range partitioning, hash partitioning, directory-based partitioning, shard key choice, scatter-gather, cross-shard queries, partition vs shard terminology

**L2 — deep & efficient**
- ⬜ **11.2 Consistent hashing** — hash ring, virtual nodes, minimal remapping, rendezvous (HRW) hashing, jump consistent hash, consistent hashing with bounded loads, Maglev hashing
- ⬜ **11.3 Hot keys & rebalancing** — hot partition, key salting / splitting, fixed number of partitions, dynamic splitting, rebalancing data movement, throttled migration, request routing (routing tier, partition-aware client, gossip)
- ⬜ **11.4 Secondary indexes & cross-partition operations** — local (document-partitioned) index, global (term-partitioned) index, async index updates, cross-shard transactions, distributed joins, resharding strategies

🛠 **Practice** — L1: pick shard keys for 3 tables (users, orders, chat messages) in `code/system-design/11-partitioning/shard-keys.md` and list which queries become scatter-gather · L2: implement consistent hashing with virtual nodes in Go, then measure (a) key distribution std-dev for 10/100/500 vnodes and (b) % of keys moved when adding/removing a node; compare with naive `hash % N` and rendezvous hashing

---

## SD 12 — Consensus, time & coordination
**L1 — aware & can use**
- ⬜ **12.1 Coordination services** — etcd, ZooKeeper, Consul, leader election, service registry, distributed configuration, distributed lock, watches, why not to build your own

**L2 — deep & efficient**
- ⬜ **12.2 Consensus & Raft** — consensus problem, Raft (leader election, term, log replication, commit index, majority quorum), election timeout, membership changes, quorum size 2f+1, Paxos / Multi-Paxos (awareness), ZAB
- ⬜ **12.3 Time & ordering** — physical clocks, clock skew & drift, NTP, monotonic vs wall clock (Go `time.Now()` monotonic reading), happens-before, Lamport clock, vector clock, hybrid logical clock (HLC), TrueTime (Spanner), total order broadcast
- ⬜ **12.4 Locks, leases & fencing** — distributed lock, lease, lease expiry vs process pause, fencing token, Redlock critique, etcd lease + `concurrency` package, Kubernetes `Lease` leader election, DB advisory locks as an alternative

🛠 **Practice** — L1: use etcd's Go client (`go.etcd.io/etcd/client/v3/concurrency`) to run leader election among 3 Go processes; kill the leader and time the handover · L2: work through the MIT 6.5840 Raft lab (or a minimal Raft in Go: election + log replication) and break it with partitions; write a demo where a paused lock holder corrupts data, then fix it with fencing tokens

---

## SD 13 — Distributed transactions & data consistency
**L1 — aware & can use**
- ⬜ **13.1 The dual-write problem** — dual write, atomicity across DB + broker, database-per-service, why no global transactions, eventual consistency in business flows, idempotency keys

**L2 — deep & efficient**
- ⬜ **13.2 2PC & atomic commit** — two-phase commit, coordinator, prepare / commit, in-doubt transactions, blocking on coordinator failure, XA, `PREPARE TRANSACTION`, 3PC (awareness), why avoided in microservices, Spanner / Percolator (awareness)
- ⬜ **13.3 Sagas** — saga, orchestration vs choreography, compensating transaction, pivot transaction, semantic lock, saga state machine, timeouts, Temporal workflows
- ⬜ **13.4 Outbox, inbox & CDC** — transactional outbox, polling publisher vs log tailing (CDC, Debezium), message relay, inbox pattern / idempotent consumer, dedup table, exactly-once vs effectively-once, ordering per aggregate
- ⬜ **13.5 Idempotency deep** — idempotency key, request fingerprint, stored response, concurrent duplicate requests, key TTL, natural idempotency (PUT, upsert), idempotent consumers, Stripe-style idempotency

🛠 **Practice** — L1: in Go, write an "order created" flow that writes to Postgres and publishes to a broker, crash between the two steps and show the lost event · L2: fix it with a transactional outbox + relay worker and an idempotent consumer (inbox table); then implement a 3-step order saga (reserve stock → charge → ship) with compensations and inject failures at every step

---

## SD 14 — Resilience
**L1 — aware & can use**
- ⬜ **14.1 Timeouts & retries** — timeouts everywhere, `context.WithTimeout`, retry, exponential backoff, jitter, retry only idempotent operations, max attempts, deadline propagation
- ⬜ **14.2 Circuit breakers & bulkheads** — circuit breaker (closed / open / half-open), bulkhead, fallback, failure thresholds, per-dependency isolation, `sony/gobreaker`

**L2 — deep & efficient**
- ⬜ **14.3 Overload protection** — load shedding, backpressure, admission control, adaptive concurrency limits, bounded queues, retry storm, retry budget, metastable failures, priority-based shedding
- ⬜ **14.4 Failure isolation at scale** — blast radius, cell-based architecture, shuffle sharding, static stability, graceful degradation, kill switches, dependency tiers, regional isolation
- ⬜ **14.5 Resilience testing** — chaos engineering, fault injection, Toxiproxy, `tc netem`, game days, steady-state hypothesis, abort conditions, Jepsen

🛠 **Practice** — L1: write a Go client calling a flaky service (Toxiproxy in Docker adding latency/errors) with timeout + backoff + jitter + circuit breaker; graph success rate and latency with each feature on/off · L2: create a retry storm (3 layers × 3 retries) against an overloaded service, then fix it with a retry budget and load shedding, and show the system recovers

---

## SD 15 — Messaging deep
**L1 — aware & can use**
- ⬜ **15.1 Kafka basics** — broker, topic, partition, offset, consumer group, message key → partition, retention, ordering per partition, `kcat`, Go clients (franz-go, segmentio/kafka-go)

**L2 — deep & efficient**
- ⬜ **15.2 Kafka internals & guarantees** — replication factor, partition leader / follower, ISR, `acks=all`, `min.insync.replicas`, idempotent producer, transactions (exactly-once semantics), consumer rebalancing (cooperative sticky), offset commit strategies, log segments, log compaction, KRaft
- ⬜ **15.3 RabbitMQ & NATS** — exchanges (direct / topic / fanout / headers), queue, binding, routing key, ack / nack, prefetch (QoS), dead-letter exchange, quorum queues, RabbitMQ streams; NATS subjects, JetStream, NATS KV
- ⬜ **15.4 Messaging patterns & pitfalls** — DLQ, poison message, retry topics with delay, ordering guarantees, consumer lag, partition count sizing, head-of-line blocking, competing consumers, claim-check for large messages
- ⬜ **15.5 Event schemas & contracts** — event schema evolution, backward / forward compatibility, Avro, Protobuf, JSON Schema, schema registry, event versioning, CloudEvents, AsyncAPI

🛠 **Practice** — L1: run a single-node Kafka (KRaft) in Docker, write a Go producer/consumer with keys, start 3 consumers in one group and watch partitions rebalance · L2: run a 3-broker cluster with RF=3 and `min.insync.replicas=2`, kill brokers during writes with `acks=1` vs `acks=all` and count lost messages; add a poison message and a retry → DLQ pipeline

---

## SD 16 — Stream processing & data pipelines
**L1 — aware & can use**
- ⬜ **16.1 Stream vs batch** — batch processing, stream processing, bounded vs unbounded data, latency vs throughput, ETL vs ELT, MapReduce (idea), lambda vs kappa architecture

**L2 — deep & efficient**
- ⬜ **16.2 Stream processing concepts** — event time vs processing time, windowing (tumbling, hopping, sliding, session), watermarks, late events, stateful processing, checkpoints, exactly-once processing, Kafka Streams, Apache Flink
- ⬜ **16.3 CDC & event sourcing** — change data capture, Debezium, Postgres logical decoding, event sourcing, snapshots, projections, CQRS, replaying events, outbox via CDC

🛠 **Practice** — L1: write a Go consumer that computes "orders per minute" in tumbling windows from a Kafka topic · L2: set up Debezium (Kafka Connect, Docker) on your Postgres, stream `orders` changes into Kafka, and build a Go projection that maintains a read model; handle late/out-of-order events with event time + allowed lateness

---

## SD 17 — Classic designs II
**L1 — aware & can use**
- ⬜ **17.1 Chat (WhatsApp)** — 1:1 & group messaging, WebSocket gateway, message ID ordering, delivery / read receipts, offline delivery, push fallback, presence, message storage (wide-column), end-to-end encryption (awareness)
- ⬜ **17.2 Ticket booking** — seat inventory, reservation hold with TTL, double-booking prevention, `SELECT FOR UPDATE` vs optimistic locking, virtual waiting room, flash crowds, payment timeout, idempotency

**L2 — deep & efficient**
- ⬜ **17.3 Payment system** — double-entry ledger, idempotency keys, payment state machine, PSP integration, webhooks, reconciliation, outbox, audit trail, `numeric` money & currencies, exactly-once money movement
- ⬜ **17.4 Key-value store** — consistent hashing, replication, quorum, vector clocks, hinted handoff, Merkle-tree anti-entropy, gossip membership, LSM-tree storage engine, tunable consistency
- ⬜ **17.5 Distributed job scheduler** — cron at scale, job store, leader election vs partitioned scheduling, at-least-once execution, idempotent jobs, retries & backoff, visibility timeout, delayed queues, timing wheel, missed runs
- ⬜ **17.6 File storage & sync (Dropbox)** — chunking, content hashing & dedup, metadata DB, block store (S3), delta sync, conflict resolution, sync notifications (long polling), versioning, resumable upload

🛠 **Practice** — L1: write `code/system-design/17-<design>/design.md` for chat and ticket booking following the framework · L2: same for each L2 design, with a "failure modes" table (what breaks → impact → mitigation); for the payment system also prototype the double-entry ledger schema in Postgres with an invariant check query

---

## SD 18 — Classic designs III
**L1 — aware & can use**
- ⬜ **18.1 Video streaming (YouTube)** — upload pipeline, transcoding, adaptive bitrate (HLS, DASH), segments, manifests, CDN delivery, metadata DB, view counting, thumbnails
- ⬜ **18.2 RAG chat over documents** — ingestion pipeline, chunking, embeddings, vector store (pgvector), retrieval + reranking, prompt assembly, LLM API, streaming responses (SSE), cost & latency per request, evaluation, document-level access control

**L2 — deep & efficient**
- ⬜ **18.3 Ride sharing / geo** — geohash, quadtree, S2 / H3 cells, high-rate location updates, nearby driver search, matching, ETA, trip state machine, surge pricing, hot regions
- ⬜ **18.4 Metrics & logging platform** — ingestion pipeline, agents (OpenTelemetry Collector), Kafka buffer, time-series storage, downsampling, retention, high cardinality, logs in columnar storage (ClickHouse, Loki), query fan-out, alerting
- ⬜ **18.5 Top-K / trending** — heavy hitters, count-min sketch, space-saving algorithm, sliding windows, approximate vs exact, map-reduce aggregation, stream processing, HyperLogLog for unique counts
- ⬜ **18.6 Collaborative editor (Google Docs)** — OT vs CRDT, document session server, operation log, snapshots, cursors & presence, offline edits, conflict-free merge, WebSocket fan-out

🛠 **Practice** — L1: write `code/system-design/18-<design>/design.md` for video streaming and RAG chat; for RAG, include a per-request cost/latency estimate · L2: same for each L2 design; for top-K, implement a count-min sketch in Go and compare its error vs an exact map on a 10M-event stream

---

## SD 19 — Real-world design skills
**L1 — aware & can use**
- ⬜ **19.1 Design docs & ADRs** — RFC, context, goals / non-goals, proposed design, alternatives considered, risks, rollout plan, ADR (Architecture Decision Record), review process

**L2 — deep & efficient**
- ⬜ **19.2 Multi-region** — active-passive, active-active, data residency (GDPR), geo-routing, cross-region replication, conflict handling, regional failover drills, RPO/RTO per region, latency budgets
- ⬜ **19.3 Evolving systems** — strangler fig, monolith to services, dual writes & backfill, shadow traffic, dark launches, feature flags, data migration at scale, API deprecation
- ⬜ **19.4 Learning from real systems** — engineering blogs (Uber, Discord, Stripe, Cloudflare, Figma, Netflix), postmortems (AWS us-east-1, Cloudflare, GitHub, CrowdStrike), Google SRE book, DDIA (Kleppmann), papers (Dynamo, Spanner, Kafka, Raft)

🛠 **Practice** — L1: write a real design doc + 2 ADRs for a feature at work or in a side project in `code/system-design/19-real-world/` · L2: read one engineering blog post or postmortem per week and write a 10-line summary (problem → design → trade-off → lesson); write a migration plan for moving one of your services to active-active across two regions
