# Senior Software Engineer Roadmap — AI Era Edition

**For:** Backend engineer, 3 years, Go + PostgreSQL → Senior (and a path toward Staff)
**Scope:** Classic SWE + DevOps + QA + Networking + Security + AI engineering
**Relationship to other files:**
- `SENIOR_ROADMAP.md` (this file) — **WHAT** to learn: the complete skill map
- `MASTER_PLAN.md` — **WHEN / HOW**: interview-focused schedule and cycles (C1/C2/C3)
- `SPRINT.md` — **TODAY**: granular checklist of the current topics

---

## 0. How to read this roadmap

### Level tags
| Tag | Meaning (EN) | Значение (RU) |
|---|---|---|
| 🟢 **Core** | Must know deeply. Expected from a strong middle; a gap here blocks senior. | Обязательно и глубоко. Без этого senior невозможен. |
| 🟡 **Senior** | Expected at senior level: you can discuss trade-offs and own it in production. | Уровень senior: обсуждаешь компромиссы и отвечаешь за это в продакшене. |
| 🔵 **Staff / Later** | Awareness now, depth later. Do not block on these. | Сейчас — знать, что это есть. Глубина — позже. |

### AI-era tags
| Tag | Meaning (EN) | Значение (RU) |
|---|---|---|
| ⬆ | **More valuable** because of AI. AI produces more code → this skill judges, verifies, or directs it. | **Ценнее** из-за ИИ. ИИ пишет больше кода → этот навык его проверяет и направляет. |
| ⬇ | **Less valuable** to memorize. Understand the concept; let AI/docs handle the details. | **Менее ценно** запоминать. Понимай концепцию, детали — ИИ и документация. |
| = | Unchanged. Still needed exactly as before. | Без изменений. |

### Depth model — T-shaped engineer
**EN:** You cannot be an expert in everything. A senior backend engineer is **T-shaped**: very deep in the vertical bar (Go, PostgreSQL, distributed backend systems, system design) and **working proficiency** across the horizontal bar (DevOps, networking, security, QA, AI). "Full DevOps skills" for you means: *you can take your service from code to production on your own — build, containerize, deploy, observe, debug, secure it — without waiting for a DevOps team.* That is exactly what modern senior roles expect.

**RU:** Невозможно быть экспертом во всём. Senior backend-инженер имеет **T-образный профиль**: очень глубоко в вертикали (Go, PostgreSQL, распределённые backend-системы, system design) и **рабочий уровень** по горизонтали (DevOps, сети, безопасность, QA, ИИ). «Полный DevOps» для тебя значит: *ты сам можешь довести сервис от кода до продакшена — собрать, упаковать в контейнер, задеплоить, наблюдать, отлаживать и защищать — не дожидаясь DevOps-команды.* Именно этого ждут от современного senior.

```
        Go │ PostgreSQL │ Distributed Systems │ System Design      ← DEEP (vertical bar)
  ─────────┼────────────┼─────────────────────┼──────────────────
  DevOps · Networking · Security · QA/Testing · AI Engineering    ← WORKING PROFICIENCY
  Communication · Writing · Ownership · Mentoring                 ← SENIOR MULTIPLIERS
```

---

## 1. What "Senior" actually means / Что на самом деле значит «Senior»

**EN:** Seniority is not years and not the number of technologies. A senior engineer:
1. **Owns outcomes, not tasks** — takes an ambiguous problem and delivers a working, maintainable solution in production.
2. **Makes trade-offs explicitly** — "we choose X over Y because of Z, and we accept cost W."
3. **Thinks in systems** — failure modes, scale, cost, security, operability — not only "does the code work."
4. **Reduces risk** — tests, observability, rollbacks, incremental delivery.
5. **Multiplies others** — code reviews, mentoring, docs, unblocking people.
6. **Communicates clearly** — in writing and speaking, to engineers and to non-engineers.

**RU:** Senior — это не годы опыта и не количество технологий. Senior-инженер:
1. **Отвечает за результат, а не за задачи** — берёт размытую проблему и доводит рабочее, поддерживаемое решение до продакшена.
2. **Делает компромиссы явно** — «выбираем X вместо Y из-за Z и принимаем цену W».
3. **Мыслит системно** — отказы, масштаб, стоимость, безопасность, эксплуатация — а не только «код работает».
4. **Снижает риски** — тесты, наблюдаемость, откаты, поэтапная поставка.
5. **Усиливает команду** — код-ревью, менторство, документация, помощь коллегам.
6. **Ясно общается** — устно и письменно, с инженерами и не-инженерами.

---

## 2. The AI-era shift / Сдвиг в эпоху ИИ

**EN:** AI coding assistants made *writing* code cheap. They did **not** make *correct, secure, operable systems* cheap. The bottleneck moved from typing to **judgement**: specifying the problem, designing boundaries, reviewing, verifying, debugging, and operating. Fundamentals matter **more**, not less — you cannot review what you do not understand.

**RU:** ИИ-ассистенты сделали *написание* кода дешёвым. Но они **не** сделали дешёвыми *корректные, безопасные и эксплуатируемые системы*. Узкое место сместилось от набора кода к **суждению**: постановка задачи, проектирование границ, ревью, проверка, отладка и эксплуатация. Фундамент важен **больше**, а не меньше — нельзя проверить то, чего не понимаешь.

### ⬇ Matters less (understand, do not memorize) / Менее важно
- Memorizing syntax, stdlib signatures, framework APIs, CLI flags
- Writing boilerplate: CRUD handlers, DTO mapping, config parsing, YAML manifests from scratch
- Regex, SQL syntax trivia, bash one-liners from memory
- Learning many frameworks/languages superficially ("CV padding")
- Grinding LeetCode **Hard** — learn **patterns** at Medium level instead (interviews still ask DSA, see §3)
- Being the "person who knows how to configure tool X"

### ⬆ Matters more / Более важно
- **System design & architecture** — AI cannot own the boundaries and trade-offs of *your* system
- **Reading and reviewing code** fast and critically (most code you touch will be generated)
- **Debugging & root-cause analysis** — production incidents, races, performance, "why does this happen"
- **Testing & verification** — the main defense against plausible-but-wrong generated code
- **Security** — AI generates vulnerable code confidently; supply chain attacks are growing
- **Data modeling & database internals** — schema mistakes are expensive and AI copies bad patterns
- **Observability & operations** — understanding real system behavior
- **Problem framing & writing specs** — the quality of AI output depends on the quality of your input
- **Domain knowledge & product thinking** — knowing *what* to build and *why*
- **Communication, English, writing** — design docs, RFCs, PR descriptions, explaining trade-offs
- **AI engineering** — integrating LLMs into products (new and expected in many backend roles)

### 🆕 New skills of the AI era / Новые навыки
- Effective AI-assisted workflow: decomposition, context engineering, iterative prompting, verification loops
- Knowing **when not to trust** AI output (hallucinated APIs, subtle concurrency bugs, insecure defaults)
- LLM integration: tool calling, structured output, RAG, embeddings, evals, cost/latency control
- AI security: prompt injection, data leakage, excessive agency (OWASP LLM Top 10)

---

## 3. CS Foundations: DS & Algorithms / Структуры данных и алгоритмы

**EN:** Interviews still test this, and it trains the reasoning you need to judge complexity of real code. Focus on **patterns**, not memorized solutions.
**RU:** Собеседования всё ещё это проверяют, и это тренирует мышление для оценки сложности реального кода. Учи **паттерны**, а не заученные решения.

**AI shift:** = for interviews, ⬇ for daily work (but complexity reasoning ⬆).

### 🟢 Core
- Big-O: time and space, amortized analysis, best/average/worst case
- Arrays/slices, strings (bytes vs runes, UTF-8), hash maps, sets
- Linked lists, stacks, queues, deques
- Trees: binary tree, BST, traversals (in/pre/post/level order)
- Heaps / priority queues, top-K problems
- Graphs: representations, BFS, DFS, topological sort, cycle detection
- Binary search (on arrays and on the answer space)
- Sorting: quicksort, mergesort, counting sort; stability; when sort is O(n log n) vs O(n)
- Recursion and backtracking (subsets, permutations, combinations)
- Patterns: two pointers, sliding window, prefix sums, fast/slow pointers, monotonic stack, intervals (merge/insert)

### 🟡 Senior
- Dynamic programming: 1D, 2D, knapsack, LCS, edit distance — recognition more than mastery
- Tries, union-find (disjoint set), segment tree (awareness)
- Shortest paths: Dijkstra, Bellman-Ford (awareness), 0-1 BFS
- Greedy algorithms and proving them informally
- Bit manipulation basics
- Data structures used in real systems: B-tree / B+tree, LSM-tree, skip list, bloom filter, HyperLogLog, consistent hashing ring, ring buffer, LRU cache

### 🔵 Staff / Later
- Count-min sketch, merkle trees, CRDTs, rope, radix tree
- Advanced graph algorithms (max flow, SCC)

**Done when:** you solve most Medium problems in 25–30 min, explain complexity, and recognize the pattern in < 5 min.
**Resources:** NeetCode 150 roadmap, "Grokking the Coding Interview" patterns, LeetCode by tag.

---

## 4. CS Foundations: Operating Systems & Computer Architecture / ОС и архитектура ЭВМ

**EN:** Go runtime, Postgres, Docker, and Kubernetes are all built on OS primitives. Without this layer, performance and container behavior stay "magic."
**RU:** Рантайм Go, Postgres, Docker и Kubernetes построены на примитивах ОС. Без этого слоя производительность и поведение контейнеров остаются «магией».

**AI shift:** ⬆ (debugging and performance need this, AI cannot see your production box)

### 🟢 Core
- Process vs thread vs goroutine; context switch cost
- Virtual memory, pages, stack vs heap, page faults, OOM killer
- System calls, user space vs kernel space
- File descriptors, everything-is-a-file, pipes, sockets
- I/O models: blocking, non-blocking, multiplexing (`select`/`poll`/`epoll`), how Go netpoller uses epoll
- Signals (SIGTERM, SIGKILL, SIGINT) → graceful shutdown
- Concurrency primitives: mutex, semaphore, condition variable, atomic operations, CAS
- Deadlock, livelock, starvation, race condition — definitions and examples

### 🟡 Senior
- CPU caches (L1/L2/L3), cache lines, false sharing, memory locality — why slices beat linked lists
- Memory ordering and memory models (happens-before)
- File systems: inodes, page cache, `fsync`, write-ahead logging, why durability is hard
- Scheduling: CFS basics, CPU throttling in containers
- Linux namespaces and cgroups (foundation of containers — see §13)
- NUMA awareness, huge pages (awareness)

### 🔵 Staff / Later
- io_uring, eBPF (also see §13 observability)
- Lock-free data structures

**Resources:** "Operating Systems: Three Easy Pieces" (free), "The Linux Programming Interface" (reference), Brendan Gregg "Systems Performance".

---

## 5. Go Mastery / Глубокое знание Go

**EN:** Your main language must be your deepest skill. Senior Go means: idiomatic code, correct concurrency, knowledge of runtime internals, and the ability to profile and fix performance.
**RU:** Основной язык должен быть самым глубоким навыком. Senior в Go — это идиоматичный код, корректная конкурентность, знание внутренностей рантайма и умение профилировать и чинить производительность.

**AI shift:** syntax ⬇, concurrency + runtime + idioms ⬆ (AI writes subtle goroutine leaks and races).

### 🟢 Core — language
- Types, zero values, structs, embedding, methods (value vs pointer receivers)
- Slices internals (ptr/len/cap, append growth, sharing trap), maps internals (buckets, swiss tables since Go 1.24)
- Strings, bytes, runes, `strings.Builder`
- Interfaces: implicit implementation, small interfaces, `any`, type assertion/switch, nil interface vs interface holding nil pointer
- Errors: `error` interface, wrapping `%w`, `errors.Is` / `errors.As` / `errors.Join`, sentinel vs typed errors, when to panic
- `defer` (order, argument evaluation, cost), `panic` / `recover`
- Generics: type parameters, constraints, when NOT to use them
- Packages, modules, `go.mod`/`go.sum`, versioning, `internal/`, workspaces
- Project layout conventions (and why `pkg/` is controversial)

### 🟢 Core — concurrency
- Goroutines, channels (buffered/unbuffered, closing rules, nil channel behavior), `select`
- `sync`: Mutex, RWMutex, WaitGroup, Once, Pool, Cond, Map; `sync/atomic`
- `context`: cancellation, timeouts, deadlines, values (and their misuse), propagation
- Patterns: worker pool, fan-out/fan-in, pipeline, semaphore, `errgroup`, rate limiting
- Goroutine leaks: causes and detection (goleak)
- Data races: race detector `-race`, the Go memory model (happens-before)
- Graceful shutdown of HTTP servers and workers

### 🟡 Senior — runtime internals
- GMP scheduler: G, M, P, run queues, work stealing, preemption, `GOMAXPROCS` (container-aware since Go 1.25)
- Garbage collector: tri-color mark-and-sweep, write barrier, `GOGC`, `GOMEMLIMIT`, GC pauses
- Escape analysis, stack growth, allocation cost (`go build -gcflags=-m`)
- Channel internals (`hchan`, wait queues), interface internals (`iface`/`eface`, itab), map internals
- Performance: `pprof` (CPU, heap, goroutine, block, mutex), `go tool trace`, benchmarks, `benchstat`, PGO (profile-guided optimization)
- Memory optimization: preallocation, `sync.Pool`, avoiding allocations in hot paths

### 🟡 Senior — ecosystem & tooling
- `net/http` deep: server timeouts, `http.Client` reuse and timeouts, transport, connection pooling, middleware, Go 1.22 routing patterns
- `database/sql` and `pgx`: connection pool settings, prepared statements, transactions, context cancellation
- `encoding/json` pitfalls (and json/v2), `log/slog` structured logging
- Code generation: `go generate`, sqlc, protobuf, mockgen
- Linters: `go vet`, `staticcheck`, `golangci-lint`; `govulncheck`
- Build: cross-compilation, build tags, `-ldflags`, static binaries, small Docker images

### 🔵 Staff / Later
- `unsafe`, cgo costs, assembly reading, compiler internals (SSA)
- Writing libraries with stable APIs and backward compatibility

**Resources:** "100 Go Mistakes and How to Avoid Them" (Harsanyi), "Concurrency in Go" (Cox-Buday), "Learning Go" 2nd ed (Bodner), Go blog, Go source code, Ardan Labs blog.

---

## 6. Databases: PostgreSQL Deep + Data Systems / Базы данных

**EN:** Data outlives code. Schema and query mistakes are the most expensive bugs in backend systems. PostgreSQL is your strongest card — make it expert-level.
**RU:** Данные живут дольше кода. Ошибки в схеме и запросах — самые дорогие баги в backend. PostgreSQL — твоя сильнейшая карта, доведи его до экспертного уровня.

**AI shift:** SQL syntax ⬇, data modeling + internals + query tuning ⬆.

### 🟢 Core — SQL & modeling
- Joins (inner/left/anti/semi), subqueries, CTEs, window functions, aggregates, `GROUP BY`/`HAVING`
- `UPSERT` (`ON CONFLICT`), `RETURNING`, `LATERAL`, `DISTINCT ON`
- Normalization (1NF–3NF/BCNF) and deliberate denormalization
- Constraints: PK, FK, UNIQUE, CHECK, NOT NULL, exclusion constraints
- Data types: numeric vs float for money, `timestamptz`, `uuid` (v7), `jsonb`, arrays, enums vs lookup tables
- Transactions, ACID, isolation levels (Read Committed, Repeatable Read, Serializable) and anomalies (dirty/non-repeatable/phantom reads, lost update, write skew)
- N+1 problem and how to fix it

### 🟢 Core — performance
- Indexes: B-tree, composite (column order!), partial, covering (`INCLUDE`), expression indexes, GIN, GiST, BRIN, hash
- `EXPLAIN (ANALYZE, BUFFERS)`: seq scan, index scan, index-only scan, bitmap scan, nested loop / hash / merge join
- Statistics and the planner, `ANALYZE`, cardinality misestimates
- Pagination: offset vs keyset (cursor) pagination
- `pg_stat_statements`, slow query log, `auto_explain`

### 🟡 Senior — internals & operations
- MVCC: tuples, xmin/xmax, visibility, snapshots
- VACUUM, autovacuum tuning, bloat, transaction ID wraparound, HOT updates, fillfactor
- Locks: row locks, table lock levels, `SELECT ... FOR UPDATE / SKIP LOCKED / NOWAIT`, deadlocks, advisory locks, lock queues during migrations
- WAL, checkpoints, `fsync`, durability settings
- Replication: streaming (physical), logical replication, sync vs async, replication lag, failover (Patroni)
- Backups: `pg_dump` vs physical backups, PITR (pgBackRest, WAL-G), restore testing
- Connection pooling: PgBouncer (session vs transaction mode), why Postgres connections are expensive
- Partitioning: range/list/hash, partition pruning, when it helps and when it hurts
- Zero-downtime migrations: adding columns/indexes `CONCURRENTLY`, expand-and-contract pattern, backfills in batches
- Postgres as a queue (`SKIP LOCKED`), outbox table, `LISTEN/NOTIFY`
- Configuration: `shared_buffers`, `work_mem`, `effective_cache_size`, `max_connections`
- Security: roles, privileges, row-level security (RLS)

### 🟡 Senior — other data systems (when to choose what)
- **Redis:** data types, TTL, eviction policies, persistence (RDB/AOF), pub/sub vs streams, distributed lock caveats (Redlock debate), cluster vs sentinel
- **Column stores:** ClickHouse — OLTP vs OLAP difference
- **Search:** Elasticsearch/OpenSearch — inverted index, when Postgres full-text search is enough
- **Document / wide-column / KV:** MongoDB, Cassandra/ScyllaDB, DynamoDB — data modeling by access pattern
- **Vector search:** pgvector, HNSW vs IVFFlat (see §15)
- **Object storage:** S3 semantics, presigned URLs, lifecycle policies

### 🔵 Staff / Later
- Storage engines: B-tree vs LSM-tree trade-offs (write/read/space amplification)
- Distributed SQL: CockroachDB, YugabyteDB, Spanner — how they achieve consistency
- Postgres internals: buffer manager, extensions development, Citus

**Resources:** "PostgreSQL 14 Internals" (Rogov, free PDF), "Use The Index, Luke" (free), "The Art of PostgreSQL", Postgres docs, Crunchy Data & pganalyze blogs.

---

## 7. API & Backend Engineering / Проектирование API и backend

**EN:** The API is a contract — hard to change after clients depend on it. Senior engineers design APIs that are consistent, safe to retry, and evolvable.
**RU:** API — это контракт, его тяжело менять после того, как клиенты начали от него зависеть. Senior проектирует API согласованным, безопасным для повторов и расширяемым.

**AI shift:** handler boilerplate ⬇, contract design + idempotency + error semantics ⬆.

### 🟢 Core
- HTTP semantics: methods, idempotency and safety, status codes (correct usage), headers, caching headers (`ETag`, `Cache-Control`)
- REST resource design, naming, filtering, sorting, pagination
- Error model (consistent error body, RFC 9457 Problem Details)
- Validation, request size limits, timeouts
- Idempotency keys for POST/payments
- Versioning strategies, backward/forward compatibility
- OpenAPI specs, API-first design

### 🟡 Senior
- gRPC & Protocol Buffers: unary/streaming, deadlines, status codes, interceptors, schema evolution rules
- Real-time: WebSockets, Server-Sent Events, long polling — trade-offs
- GraphQL (awareness: N+1 and DataLoader, when it fits)
- Background jobs & scheduling: retries, dead-letter, cron in distributed systems, job uniqueness
- File uploads: multipart, streaming, presigned URLs, virus scanning
- Webhooks: signing, retries, delivery guarantees
- Multi-tenancy patterns (shared schema, schema-per-tenant, DB-per-tenant)
- Feature flags
- API gateway responsibilities (auth, rate limiting, routing)

### 🔵 Staff / Later
- Public API governance, deprecation policies, SDK generation
- AsyncAPI, event contracts across teams

**Resources:** Google API Design Guide, Microsoft REST API Guidelines, Stripe API docs (gold standard for idempotency and errors).

---

## 8. Software Design & Architecture / Проектирование ПО и архитектура

**EN:** Architecture is the set of decisions that are expensive to change. In the AI era this is the core senior skill: AI fills in code, *you* define the shape and boundaries.
**RU:** Архитектура — это набор решений, которые дорого менять. В эпоху ИИ это ключевой навык senior: ИИ заполняет код, а *ты* задаёшь форму и границы.

**AI shift:** ⬆⬆

### 🟢 Core
- Coupling and cohesion, abstraction, encapsulation, information hiding
- SOLID applied to Go (interfaces at the consumer side, small interfaces)
- Composition over inheritance; dependency injection in Go (constructor injection, no magic)
- Design patterns used in Go: functional options, factory, strategy, adapter, decorator/middleware, observer, builder, repository, unit of work
- Layered / Clean / Hexagonal (ports & adapters) architecture — and when it is over-engineering
- Code smells and refactoring techniques
- Readable code: naming, small functions, explicit error paths

### 🟡 Senior
- Domain-Driven Design: ubiquitous language, bounded contexts, aggregates, entities vs value objects, domain events, context mapping
- Modular monolith vs microservices: real trade-offs, when to split, "distributed monolith" anti-pattern
- Event-driven architecture: events vs commands, choreography vs orchestration
- CQRS, event sourcing (and its costs)
- Transactional outbox, inbox/idempotent consumer, saga (orchestrated vs choreographed)
- Strangler fig pattern for legacy migration
- Technical debt management, architecture decision records (ADRs)
- Conway's law, team topologies

### 🔵 Staff / Later
- Evolutionary architecture, fitness functions
- Platform thinking, internal developer platforms
- Cell-based architecture

**Resources:** "A Philosophy of Software Design" (Ousterhout), "Learning Domain-Driven Design" (Khononov), "Fundamentals of Software Architecture" (Richards & Ford), "Refactoring" (Fowler), microservices.io (Richardson).

---

## 9. Distributed Systems / Распределённые системы

**EN:** Every backend with more than one instance, a queue, or a cache is a distributed system. The network is unreliable, clocks lie, and partial failure is normal. Senior engineers design for that.
**RU:** Любой backend больше чем с одним инстансом, очередью или кэшем — распределённая система. Сеть ненадёжна, часы врут, частичный отказ — норма. Senior проектирует с учётом этого.

**AI shift:** ⬆⬆ (AI-generated code assumes the happy path)

### 🟢 Core
- Fallacies of distributed computing
- CAP theorem, PACELC
- Replication (leader-follower, multi-leader, leaderless), replication lag effects
- Partitioning/sharding, consistent hashing, hot keys
- Idempotency, at-most-once / at-least-once / "exactly-once" (effectively-once)
- Timeouts, retries with exponential backoff and jitter, retry storms
- Caching patterns: cache-aside, write-through, write-behind; invalidation; cache stampede / thundering herd protection (singleflight, request coalescing, probabilistic early expiry)
- Rate limiting algorithms: token bucket, leaky bucket, fixed/sliding window

### 🟡 Senior
- Consistency models: linearizability, sequential, causal, eventual, read-your-writes, monotonic reads
- Consensus: Raft (leader election, log replication), why etcd/ZooKeeper exist
- Time: clock skew, NTP, logical clocks (Lamport), vector clocks, hybrid logical clocks
- Distributed transactions: 2PC and why it is avoided, sagas, outbox
- Distributed locks and fencing tokens, leases
- Resilience: circuit breaker, bulkhead, load shedding, backpressure, graceful degradation, deadline propagation
- Messaging deep — **Kafka:** partitions, consumer groups, offsets, rebalancing, ordering guarantees, retention, compaction, idempotent producer, transactions, DLQ patterns
- RabbitMQ (exchanges, acks, prefetch) and NATS/JetStream — when each fits
- Service discovery, health checks, leader election in practice

### 🔵 Staff / Later
- Stream processing (Kafka Streams, Flink): windows, watermarks, state
- CRDTs, gossip protocols, quorum systems (R + W > N)
- Formal methods awareness (TLA+), Jepsen reports

**Resources:** "Designing Data-Intensive Applications" (Kleppmann, 2nd ed) — **the single most important book on this list**, "Release It!" (Nygard), "Kafka: The Definitive Guide", Martin Kleppmann's Cambridge lectures (YouTube).

---

## 10. System Design / Системный дизайн

**EN:** System design interviews decide the senior level. In real work, it is how you turn requirements into an architecture with explicit trade-offs.
**RU:** Собеседование по system design определяет уровень senior. В реальной работе это превращение требований в архитектуру с явными компромиссами.

**AI shift:** ⬆⬆

### 🟢 Core — framework
1. Clarify requirements (functional + non-functional: latency, availability, consistency, scale)
2. Back-of-the-envelope estimation (QPS, storage, bandwidth, memory)
3. API design
4. Data model and storage choice
5. High-level architecture
6. Deep dive on 1–2 hardest components
7. Bottlenecks, failure modes, scaling, monitoring, security, cost

### 🟢 Core — building blocks
Load balancers, reverse proxies, CDN, caches, databases (SQL/NoSQL), replicas, shards, queues/streams, object storage, search index, API gateway, rate limiter, service discovery, unique ID generation (UUIDv7, Snowflake)

### 🟡 Senior — practice designs (in rough order)
URL shortener · Rate limiter · Key-value store · Notification service · Chat/messaging (WhatsApp) · News feed (Twitter) · Web crawler · Payment system (idempotency, ledger, reconciliation) · Ride sharing / geo (geohash, quadtree) · Video streaming (YouTube) · File storage & sync (Dropbox) · Search autocomplete · Distributed job scheduler · Ticket booking (contention, inventory) · Metrics/logging platform · LLM-powered feature (RAG chat over company docs)

### 🟡 Senior — real-world design skills
- SLO-driven design: availability numbers (99.9 vs 99.99 meaning), latency percentiles (p50/p95/p99)
- Multi-region: active-passive vs active-active, data residency
- Cost estimation and cost trade-offs
- Writing a design doc / RFC and defending it in review

### 🔵 Staff / Later
- Org-level architecture, migration strategies across many teams
- Capacity planning for large systems

**Resources:** "System Design Interview" Vol 1 & 2 (Alex Xu), ByteByteGo, Hello Interview (free guides), engineering blogs (Uber, Discord, Stripe, Cloudflare, Figma).

---

## 11. Networking / Сети

**EN:** Every request your service handles crosses the network. When something "randomly" times out, resets, or is slow, the answer is usually in DNS, TCP, TLS, or the load balancer. Senior engineers can debug the full path.
**RU:** Каждый запрос к твоему сервису проходит через сеть. Когда что-то «случайно» таймаутит, сбрасывается или тормозит, ответ обычно в DNS, TCP, TLS или балансировщике. Senior умеет отлаживать весь путь.

**AI shift:** ⬆ (debugging network issues requires real understanding and real tools)

### 🟢 Core — fundamentals
- OSI vs TCP/IP model — what actually happens at each layer
- "What happens when you type a URL and press Enter" — full answer: DNS → TCP → TLS → HTTP → server → response → rendering
- IP addressing: IPv4, IPv6 basics, CIDR and subnetting, private ranges, NAT
- Ports, sockets, ephemeral ports
- **TCP:** 3-way handshake, 4-way close, TIME_WAIT, sequence numbers, retransmission, flow control (window), congestion control (slow start, CUBIC/BBR awareness), Nagle's algorithm, keepalive
- **UDP:** when and why (DNS, video, QUIC)
- **DNS:** recursive vs authoritative, record types (A, AAAA, CNAME, MX, TXT, NS, SRV), TTL and caching, DNS as a source of outages
- **HTTP:** HTTP/1.1 (keep-alive, head-of-line blocking), HTTP/2 (multiplexing, streams), HTTP/3 + QUIC
- **TLS 1.3:** handshake, certificates, chain of trust, CA, SNI, mTLS, Let's Encrypt/ACME
- Proxies: forward vs reverse proxy; L4 vs L7 load balancing
- Debugging tools: `curl -v`, `dig`, `nslookup`, `ping`, `traceroute`/`mtr`, `ss`/`netstat`, `tcpdump`, Wireshark, `openssl s_client`, `nc`

### 🟡 Senior
- Load balancing algorithms and health checks; connection draining
- Nginx / Envoy / HAProxy basics: configs, timeouts, buffering
- Connection pooling and keep-alive between services; head-of-line blocking in practice
- CDN internals: edge caching, cache keys, purging, anycast
- Cloud networking: VPC, subnets (public/private), route tables, internet gateway, NAT gateway, security groups vs NACLs, peering, private endpoints
- VPNs: WireGuard, IPsec (awareness), bastion hosts, zero-trust access
- Linux networking: interfaces, routing table (`ip route`), iptables/nftables, network namespaces, veth pairs, bridges
- Kubernetes networking: pod network (CNI), Services (ClusterIP/NodePort/LoadBalancer), kube-proxy, Ingress / Gateway API, CoreDNS, NetworkPolicies
- Service mesh (Istio/Linkerd): sidecars, mTLS, traffic shifting — and when it is overkill

### 🔵 Staff / Later
- BGP and anycast basics, how internet routing fails
- eBPF-based networking (Cilium)
- Network performance tuning (kernel parameters, buffer sizes)

**Resources:** "Computer Networking: A Top-Down Approach" (Kurose & Ross), "High Performance Browser Networking" (Grigorik, free online), Julia Evans' networking zines, Cloudflare blog.

---

## 12. Security / Безопасность

**EN:** Security is a property of the whole system, not a feature. Every senior engineer is expected to write secure code by default, recognize common vulnerabilities in review, and design auth correctly. AI makes this more urgent: generated code often has insecure defaults.
**RU:** Безопасность — свойство всей системы, а не отдельная фича. От senior ожидают писать безопасный код по умолчанию, замечать типовые уязвимости на ревью и правильно проектировать аутентификацию. ИИ делает это ещё актуальнее: сгенерированный код часто небезопасен по умолчанию.

**AI shift:** ⬆⬆

### 🟢 Core — secure coding
- OWASP Top 10 (web) and OWASP API Security Top 10: injection (SQL, command), broken access control, IDOR/BOLA, SSRF, XSS, CSRF, insecure deserialization, security misconfiguration, mass assignment
- Input validation, output encoding, parameterized queries (always)
- Secrets never in code, logs, or Git history
- Least privilege principle everywhere (DB users, IAM, containers)
- Safe error messages (no stack traces / internal info to clients)
- CORS — what it protects and what it does not

### 🟢 Core — authentication & authorization
- Sessions vs tokens; cookies (HttpOnly, Secure, SameSite)
- JWT: structure, signing (HS256 vs RS256/EdDSA), validation, expiration, revocation problem, refresh tokens
- OAuth 2.0 flows (authorization code + PKCE, client credentials), OpenID Connect
- Password storage: bcrypt / scrypt / argon2id — never plain hashes
- MFA, passkeys/WebAuthn (awareness)
- Authorization models: RBAC, ABAC, ReBAC (Google Zanzibar, OpenFGA), policy engines (OPA)

### 🟡 Senior — cryptography (use correctly, never invent)
- Hashing vs encryption vs encoding vs signing
- Symmetric (AES-GCM) vs asymmetric (RSA, ECDSA, Ed25519), key exchange
- HMAC for webhooks and request signing
- Encryption at rest vs in transit; envelope encryption; KMS
- Random numbers: `crypto/rand` vs `math/rand`
- Key rotation

### 🟡 Senior — infrastructure & supply chain
- Secrets management: HashiCorp Vault, cloud secret managers, Kubernetes secrets (and their weaknesses), SOPS
- Dependency security: `govulncheck`, Dependabot/Renovate, SCA scanning
- Supply chain: SBOM, signed images (Sigstore/cosign), SLSA levels, pinning dependencies
- Container security: non-root, read-only FS, minimal/distroless images, image scanning (Trivy, Grype)
- Kubernetes security: RBAC, Pod Security Standards, NetworkPolicies, secrets encryption
- Cloud IAM: roles, policies, least privilege, workload identity (no static keys)
- SAST/DAST in CI (Semgrep, gosec, ZAP)
- Security logging and audit trails; detecting abuse
- Rate limiting and bot protection, DDoS basics

### 🟡 Senior — process
- Threat modeling (STRIDE), data flow diagrams, trust boundaries
- Security review of designs and PRs
- Incident response basics for security events
- Compliance awareness: GDPR / personal data laws (incl. Uzbekistan and Russia data localization), PCI DSS for payments, SOC 2

### 🟡 Senior — AI security (🆕)
- OWASP Top 10 for LLM Applications: prompt injection (direct/indirect), sensitive data disclosure, excessive agency, insecure output handling
- Never trust LLM output: validate, sandbox, least-privilege tools
- Data leakage to third-party AI providers; secrets in prompts

### 🔵 Staff / Later
- Penetration testing basics (Burp Suite, PortSwigger Web Security Academy labs)
- Zero-trust architecture, SPIFFE/SPIRE
- Security champion role in a team

**Resources:** OWASP Cheat Sheet Series, PortSwigger Web Security Academy (free, hands-on), "Real-World Cryptography" (Wong), "API Security in Action" (Madden).

---

## 13. DevOps, Platform & Cloud / DevOps, платформа и облако

**EN:** "You build it, you run it." A senior backend engineer ships and operates their own services. Goal: take a service from `git push` to production — automated, observable, secure, and recoverable.
**RU:** «Ты построил — ты и эксплуатируешь». Senior backend сам выкатывает и поддерживает свои сервисы. Цель: довести сервис от `git push` до продакшена — автоматически, наблюдаемо, безопасно и с возможностью восстановления.

**AI shift:** writing YAML/HCL by hand ⬇, understanding what it does + debugging ⬆.

### 🟢 Core — Linux & shell
- Filesystem hierarchy, permissions (chmod/chown, users/groups), `sudo`
- Processes: `ps`, `top`/`htop`, `kill`, signals, `systemd` units, `journalctl`
- Disk and memory: `df`, `du`, `free`, `vmstat`, `iostat`, `lsof`
- Text tools: `grep`, `sed`, `awk`, `jq`, `xargs`, pipes, redirects
- Bash scripting: variables, conditionals, loops, exit codes, `set -euo pipefail`
- SSH: keys, config, tunnels/port forwarding, `scp`/`rsync`
- Package management, cron, environment variables

### 🟢 Core — Git (deep)
- Internals: objects (blob/tree/commit), refs, HEAD, the DAG
- Merge vs rebase, interactive rebase, cherry-pick, `reflog` recovery, bisect
- Branching strategies: trunk-based development vs GitFlow
- Conventional commits, signed commits, `.gitignore`, hooks

### 🟢 Core — containers
- What a container really is: namespaces + cgroups + union filesystem
- Dockerfile best practices: multi-stage builds, layer caching, distroless/scratch for Go, non-root user, `.dockerignore`
- Image registries, tags vs digests
- Docker networking and volumes
- Docker Compose for local development (app + Postgres + Redis + Kafka)
- Containers vs VMs

### 🟢 Core — CI/CD
- Pipeline design: lint → test → build → scan → publish → deploy
- GitHub Actions (or GitLab CI): workflows, jobs, matrices, caching, secrets, reusable workflows, OIDC to cloud
- Artifacts and versioning, semantic versioning
- Environments: dev / staging / prod, promotion between them
- Database migrations in pipelines

### 🟡 Senior — Kubernetes
- Architecture: control plane (API server, etcd, scheduler, controller manager), kubelet, kube-proxy
- Workloads: Pod, Deployment, ReplicaSet, StatefulSet, DaemonSet, Job, CronJob
- Config: ConfigMap, Secret, environment vs mounted files
- Networking: Service, Ingress/Gateway API, DNS (see §11)
- Health: liveness vs readiness vs startup probes (and the common mistakes)
- Resources: requests vs limits, QoS classes, CPU throttling, OOMKilled, `GOMAXPROCS`/`GOMEMLIMIT` in pods
- Scaling: HPA, VPA, cluster autoscaler, PodDisruptionBudget
- Storage: PV, PVC, StorageClass
- Packaging: Helm, Kustomize
- Debugging: `kubectl describe/logs/exec/port-forward/events/top`, CrashLoopBackOff, ImagePullBackOff, Pending pods
- Graceful shutdown in k8s: SIGTERM, `preStop`, `terminationGracePeriodSeconds`, endpoint removal race
- Local clusters: kind, k3d, minikube

### 🟡 Senior — deployment & release
- Rolling, blue-green, canary deployments; automated rollback
- Feature flags, dark launches, progressive delivery (Argo Rollouts)
- GitOps: Argo CD / Flux — desired state in Git
- Zero-downtime deploys with DB migrations (expand-and-contract)

### 🟡 Senior — Infrastructure as Code
- Terraform / OpenTofu: providers, resources, state (remote state, locking), modules, plan/apply, drift, workspaces, importing
- Configuration management: Ansible (awareness)
- Immutable infrastructure

### 🟡 Senior — Cloud (pick one deeply: AWS recommended, map concepts to GCP/Yandex Cloud)
- IAM (users, roles, policies), VPC networking (§11)
- Compute: EC2, autoscaling groups, ECS/Fargate, EKS, Lambda (serverless trade-offs)
- Storage: S3, EBS; Databases: RDS/Aurora for Postgres, ElastiCache
- Messaging: SQS, SNS, EventBridge, MSK (managed Kafka)
- Load balancers: ALB vs NLB; Route 53; CloudFront
- Monitoring: CloudWatch; cost tools
- Shared responsibility model; well-architected framework pillars

### 🟡 Senior — Observability & SRE
- Three pillars + profiles: logs, metrics, traces, continuous profiling
- Structured logging (`slog`), log levels, correlation/trace IDs, log aggregation (Loki, ELK)
- Metrics: counters, gauges, histograms; Prometheus data model, PromQL, cardinality explosion
- Distributed tracing: spans, context propagation, sampling; **OpenTelemetry** (the standard)
- Dashboards: Grafana; RED method (Rate, Errors, Duration) for services, USE method (Utilization, Saturation, Errors) for resources
- SLI / SLO / SLA, error budgets, burn-rate alerting
- Alerting philosophy: alert on symptoms, not causes; avoid alert fatigue
- Incident management: on-call, severity levels, incident commander, communication, blameless postmortems, runbooks
- Capacity planning, load testing before launches
- Chaos engineering (awareness → practice)
- Disaster recovery: RPO / RTO, backups tested by restore

### 🔵 Staff / Later
- Platform engineering, internal developer platforms (Backstage)
- Service mesh operations, multi-cluster
- eBPF observability (Pixie, Parca)
- FinOps: cost allocation, rightsizing, spot instances

**Resources:** "The Linux Command Line" (Shotts, free), "Docker Deep Dive" (Poulton), "Kubernetes Up & Running", "Terraform: Up & Running" (Brikman), Google "Site Reliability Engineering" + "The SRE Workbook" (free online), "Observability Engineering" (Majors et al.), KodeKloud / Killercoda labs. Optional certifications: CKAD → CKA, AWS Solutions Architect Associate.

---

## 14. QA, Testing & Quality Engineering / QA, тестирование и качество

**EN:** In the AI era, testing is the **main defense** against plausible-but-wrong code. A senior engineer designs a test strategy for a system, not just writes unit tests. You also need the QA mindset: thinking about how things break, edge cases, and user-visible quality.
**RU:** В эпоху ИИ тестирование — **главная защита** от правдоподобного, но неверного кода. Senior проектирует стратегию тестирования для системы, а не просто пишет юнит-тесты. Нужно и мышление QA: как система ломается, граничные случаи, качество глазами пользователя.

**AI shift:** ⬆⬆ (AI can write tests, but you must decide *what* to test and verify the tests actually catch bugs)

### 🟢 Core — testing fundamentals
- Why test: confidence, refactoring safety, documentation, design feedback
- Test pyramid vs testing trophy; cost and speed of each level
- Unit tests in Go: `testing` package, table-driven tests, subtests, `t.Parallel`, `t.Cleanup`, `testdata/`, golden files
- Test doubles: dummy, stub, fake, spy, mock — and why fakes often beat mocks
- Assertions: stdlib vs `testify`; `cmp.Diff`
- Testing HTTP: `httptest`, testing handlers and clients
- Integration tests with real dependencies: **testcontainers-go** (Postgres, Kafka, Redis)
- Test coverage: what it means and what it does not mean
- Testing concurrency: race detector, deterministic tests, `testing/synctest` (Go 1.24+)
- Good test properties: fast, isolated, repeatable, readable; testing behavior, not implementation

### 🟡 Senior — advanced testing
- Contract testing (Pact, or schema/protobuf compatibility checks) between services
- End-to-end testing: when it is worth it, keeping it small
- Property-based testing (`rapid`), fuzzing (`go test -fuzz`)
- Mutation testing (do your tests actually catch bugs?)
- Benchmark tests and performance regression detection
- Load and performance testing: **k6**, vegeta, ghz (gRPC); load profiles (smoke, load, stress, spike, soak)
- Testing database code: migrations tests, transactional test isolation, test data builders/factories
- Flaky tests: causes (time, ordering, shared state, network) and fixes
- Testing in production: canaries, synthetic monitoring, feature flags, shadow traffic
- Chaos testing (toxiproxy for network faults)
- Static analysis as quality gate; code review as a quality practice

### 🟡 Senior — QA discipline (the QA engineer's perspective)
- Test strategy and test plan documents
- Test design techniques: equivalence partitioning, boundary value analysis, decision tables, state transition testing, pairwise testing
- Exploratory testing, risk-based testing
- Regression testing, smoke vs sanity testing
- Bug reports: reproduction steps, expected vs actual, severity vs priority
- API testing tools: Postman/Bruno, Hurl, REST-assured equivalents in Go
- Acceptance criteria, BDD (Given/When/Then), TDD — practice it on real tasks
- Quality metrics: escaped defects, change failure rate, MTTR (DORA metrics)
- Release quality gates and Definition of Done

### 🔵 Staff / Later
- Test infrastructure at scale (test sharding, hermetic builds, Bazel)
- Formal verification awareness
- Building a quality culture across teams

**Resources:** "Unit Testing Principles, Practices, and Patterns" (Khorikov) — best book here, "Software Engineering at Google" ch. 11–14 (free), Learn Go with Tests (free), ISTQB Foundation syllabus (free PDF, for QA vocabulary), k6 docs.

---

## 15. AI Engineering for Backend Developers / ИИ-инженерия для backend-разработчиков 🆕

**EN:** Two separate skills: (1) **using AI** to work faster without losing quality, and (2) **building AI features** into products. Many backend job descriptions now include the second. You do not need ML theory — you need to integrate LLMs reliably, like any other unreliable external dependency.
**RU:** Два отдельных навыка: (1) **использовать ИИ**, чтобы работать быстрее без потери качества, и (2) **встраивать ИИ-функции** в продукты. Во многих вакансиях backend теперь есть второе. Теория ML не нужна — нужно надёжно интегрировать LLM, как любую другую ненадёжную внешнюю зависимость.

**AI shift:** 🆕 ⬆⬆

### 🟢 Core — AI-assisted engineering workflow
- Decompose the task before prompting; write a mini-spec first
- Context engineering: give the model the right files, constraints, conventions (CLAUDE.md, rules files)
- Iterate in small, verifiable steps; run tests after each step
- Review AI code like a junior's PR: correctness, edge cases, concurrency, security, error handling, unnecessary complexity
- Know the failure modes: hallucinated APIs/packages, outdated library versions, subtle logic bugs, over-engineering, deleted tests
- Use AI for learning: ask for explanations, then verify in docs/source
- Agentic tools (Claude Code, Cursor, Copilot agents): when to delegate, when to drive yourself
- Keep your own skills sharp: regularly solve problems without AI (interviews are often AI-free)

### 🟡 Senior — building LLM features
- LLM basics: tokens, context window, temperature, sampling, why outputs are non-deterministic
- Calling LLM APIs from Go: streaming, timeouts, retries, rate limits, fallbacks between providers
- Structured output (JSON schema), validation of model output
- Tool / function calling; Model Context Protocol (MCP) — build an MCP server in Go
- Prompt design: system prompts, few-shot examples, prompt versioning
- **RAG:** chunking, embeddings, vector search (pgvector + HNSW), hybrid search (BM25 + vectors), re-ranking, citations
- Evals: building test sets, LLM-as-judge, regression testing for prompts
- Cost and latency: token budgeting, prompt caching, model routing (small vs large model), batching
- Observability for LLM calls: logging prompts/responses safely, tracing, token usage metrics
- Guardrails and AI security (see §12)
- Agents: loops, planning, tool permissions, human-in-the-loop, when an agent is overkill

### 🔵 Staff / Later
- Fine-tuning vs RAG vs prompting trade-offs
- Self-hosted models (vLLM, Ollama), GPU cost basics
- Multi-agent orchestration

**Resources:** "AI Engineering" (Chip Huyen, 2025), Anthropic and OpenAI docs and cookbooks, pgvector docs, "Building LLM-powered applications" articles from Eugene Yan and Hamel Husain (evals).

---

## 16. Performance Engineering / Инженерия производительности

**EN:** Performance work is measurement first, optimization second. Seniors know where time goes and fix the right thing.
**RU:** Работа над производительностью — сначала измерение, потом оптимизация. Senior знает, куда уходит время, и чинит именно то, что нужно.

**AI shift:** ⬆

### 🟡 Senior
- Latency vs throughput; percentiles and why averages lie; tail latency
- Little's law, basic queueing theory (utilization → latency curve)
- Latency numbers every programmer should know
- Profiling workflow: hypothesis → measure (pprof, trace, flame graphs) → change → re-measure
- Common bottlenecks: DB queries, N+1, lock contention, allocations/GC, serialization, network round-trips, connection pool exhaustion
- Caching strategy with hit-rate measurement
- Load testing and capacity estimation (see §14)
- Linux performance tools: `perf`, `strace`, Brendan Gregg's 60-second checklist

**Resources:** Brendan Gregg "Systems Performance", Go `pprof` docs, Dave Cheney's high-performance Go workshop.

---

## 17. Engineering Practices & Delivery / Инженерные практики и поставка

**EN:** How work gets done in a team. Senior engineers improve the process, not only follow it.
**RU:** Как работа делается в команде. Senior улучшает процесс, а не только следует ему.

**AI shift:** ⬆ (code review and writing become the main daily activities)

### 🟢 Core
- Code review: what to look for, how to give kind and specific feedback, small PRs
- Writing good PR descriptions and commit messages
- Agile in practice: Scrum, Kanban, sprints, backlog refinement — and their real limitations
- Estimation: breaking work down, communicating uncertainty
- Documentation: README, runbooks, API docs

### 🟡 Senior
- Design docs / RFCs / ADRs — writing them and reviewing others'
- DORA metrics: deployment frequency, lead time, change failure rate, time to restore
- Trunk-based development, continuous delivery, small batches
- Managing technical debt with stakeholders
- Release management, versioning, changelogs
- Legacy code: characterization tests, safe refactoring

**Resources:** "Software Engineering at Google" (free), "Accelerate" (Forsgren et al.), Google's code review guide (free).

---

## 18. Senior Soft Skills & Leadership / Soft skills и лидерство

**EN:** The difference between a strong middle and a senior is mostly here. AI cannot do these for you.
**RU:** Разница между сильным middle и senior в основном здесь. ИИ это за тебя не сделает.

**AI shift:** ⬆⬆

### 🟢 Core
- Clear technical communication in **English** (speaking and writing) — Track C1
- Asking clarifying questions; turning vague requests into concrete requirements
- Ownership: following through, raising risks early, "disagree and commit"
- Giving and receiving feedback
- Behavioral interviews: STAR stories (conflict, failure, impact, leadership, ambiguity)

### 🟡 Senior
- Mentoring juniors and middles; pair programming; onboarding
- Stakeholder management: explaining technical trade-offs to product/business in their language
- Product thinking: understanding users, metrics, business impact of technical work
- Prioritization: impact vs effort, saying "no" with alternatives
- Leading a project end-to-end: planning, splitting, coordinating, delivering
- Running meetings, technical discussions, incident calls
- Influence without authority

### 🔵 Staff / Later
- Setting technical direction for multiple teams, writing strategy
- Sponsorship, growing other seniors
- Org-level problem finding (not only solving)

**Resources:** "The Staff Engineer's Path" (Tanya Reilly), "Staff Engineer" (Will Larson), "The Manager's Path" (Fournier) ch. 1–3, "Crucial Conversations".

---

## 19. Career & Visibility / Карьера и видимость

**EN:** Skills that nobody can see do not get you hired. Make your work visible.
**RU:** Навыки, которых никто не видит, не помогают получить работу. Делай свою работу видимой.

- Portfolio: 1–2 **production-quality** projects on GitHub (see capstone below), not 20 tutorials
- Write: technical blog posts / LinkedIn posts about what you learn (also trains English)
- Open source: small, real contributions to Go/Postgres ecosystem tools you use
- Interview skills: DSA (§3), system design (§10), behavioral (§18), Go deep-dive (§5) — see `MASTER_PLAN.md`
- Salary negotiation basics; evaluating companies (engineering culture, on-call, growth)
- Networking with engineers: meetups, communities (Go Uzbekistan / Russian-speaking Go communities)

---

## 20. Capstone Project — C3 "Own it" for the whole roadmap / Итоговый проект

**EN:** One project that touches every domain. Build it incrementally as you study — each roadmap section adds a layer. This is your portfolio piece and your interview story.
**RU:** Один проект, затрагивающий все области. Строй его постепенно по мере изучения — каждый раздел добавляет слой. Это твоё портфолио и история для собеседований.

**Idea:** *Event-driven order/payment platform* (or anything from your domain: archives, document generation) — 2–3 Go services.

| Layer | What you add | Roadmap section |
|---|---|---|
| 1 | Go REST + gRPC services, clean structure, graceful shutdown | §5, §7, §8 |
| 2 | PostgreSQL schema, migrations, indexes, keyset pagination, outbox | §6 |
| 3 | Kafka events, idempotent consumers, DLQ, saga | §9 |
| 4 | Redis caching with stampede protection, rate limiter | §9, §10 |
| 5 | Unit + integration (testcontainers) + contract + load tests (k6) + fuzzing | §14 |
| 6 | Docker multi-stage, Compose, GitHub Actions CI with scans | §13, §12 |
| 7 | Kubernetes (kind), Helm, probes, HPA, Argo CD | §13 |
| 8 | Terraform for cloud infra (AWS free tier or local with LocalStack) | §13 |
| 9 | OpenTelemetry traces + Prometheus metrics + Grafana dashboards + SLO alerts | §13 |
| 10 | AuthN/AuthZ with OIDC, JWT, RBAC; threat model doc; secrets in Vault | §12 |
| 11 | LLM feature: RAG search over orders/docs with pgvector + evals | §15 |
| 12 | Design doc + ADRs + runbook + postmortem of a chaos experiment | §17, §18 |

---

## 21. Sequencing — realistic timeline / Последовательность и реальные сроки

**EN:** Honest math: at ~1 hour/day, 5–6 days/week (≈300 hours/year), this whole map at full depth is **2–3 years**. That is normal — senior engineers keep learning continuously. The order below gets you to **senior-level interviews in ~6–9 months** and to **senior-level working ability in ~18 months**. If you can study more hours, compress the stages.
**RU:** Честный расчёт: ~1 час в день, 5–6 дней в неделю (≈300 часов в год) — вся карта в полной глубине займёт **2–3 года**. Это нормально — senior учится постоянно. Порядок ниже приведёт к **собеседованиям уровня senior за ~6–9 месяцев** и к **реальному уровню senior за ~18 месяцев**. Если можешь учиться больше — сжимай этапы.

### Stage 1 (Months 0–3) — Interview core + vertical bar / Ядро для собеседований
- §3 DSA patterns (finish current SPRINT topics)
- §5 Go concurrency, errors, context, runtime basics
- §6 PostgreSQL: indexes, EXPLAIN, transactions/isolation, MVCC
- §10 System design framework + 4 practice designs
- §11 Networking core ("what happens when you type a URL", TCP, DNS, TLS, HTTP)
- §12 OWASP Top 10 + auth (JWT/OAuth)
- §18 STAR stories, English
- Capstone layers 1–2

### Stage 2 (Months 3–6) — Distributed + quality / Распределённые системы и качество
- §9 Distributed systems (read DDIA alongside)
- §8 Architecture: DDD, hexagonal, outbox/saga
- §14 Testing: integration, contract, load testing; QA techniques
- §13 Linux, Git deep, Docker, CI/CD
- §15 AI-assisted workflow (Core)
- Capstone layers 3–6
- Mock interviews weekly (Track D)

### Stage 3 (Months 6–12) — Operate in production / Эксплуатация в продакшене
- §13 Kubernetes, observability (OpenTelemetry, Prometheus), SRE practices
- §13 Terraform + one cloud (AWS)
- §11 Cloud and Kubernetes networking
- §12 Supply chain, container & k8s security, threat modeling
- §6 Postgres operations: replication, backups, partitioning, zero-downtime migrations
- §16 Performance engineering
- Capstone layers 7–10
- Optional: CKAD certification

### Stage 4 (Months 12–18) — Senior → Staff track / От senior к staff
- §15 Building LLM features (RAG, evals, MCP)
- §10 Advanced designs, multi-region
- §17–18 Design docs, mentoring, leading projects
- 🔵 topics from any section that your job needs
- Capstone layers 11–12, write blog posts about it

### Continuous (all stages) / Постоянно
- DSA: 3 problems/week minimum (more before interviews)
- English: daily, inside every session (Track C1)
- Read one engineering blog post / postmortem per week
- Apply new topics at your real job first — real production experience beats any tutorial

---

## 22. Self-assessment checklist — "Am I senior?" / Чек-лист самооценки

You are ready when you can, **without AI help**:
- [ ] Explain how a Go HTTP request flows from the socket to your handler and back, including goroutines and the netpoller
- [ ] Find and fix a goroutine leak and a data race in unfamiliar code
- [ ] Read an `EXPLAIN ANALYZE` plan and fix a slow query; explain MVCC and vacuum
- [ ] Design a system end-to-end in 45 minutes with explicit trade-offs, failure modes, and estimates
- [ ] Explain at-least-once delivery and design an idempotent consumer with an outbox
- [ ] Debug "service A cannot reach service B" through DNS, TCP, TLS, and k8s networking layers
- [ ] Find OWASP vulnerabilities in a PR and design OAuth2/OIDC auth correctly
- [ ] Containerize, deploy to Kubernetes, and roll back a service with CI/CD and GitOps
- [ ] Define SLOs, build dashboards and alerts, and lead an incident + write a postmortem
- [ ] Design a test strategy for a service and explain why each test level exists
- [ ] Build a RAG feature with evals and explain its failure modes and security risks
- [ ] Write a design doc that others review and approve, and mentor a junior through a task
