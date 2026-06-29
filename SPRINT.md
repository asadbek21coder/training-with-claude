# Interview Sprint Tracker

Legend: ✅ done | 🔄 in progress | ⬜ todo

---

## Phase 1 — Data Structures & Algorithms

### 1.1 Array / Slice ✅
**Concepts — can you answer these?**
- [ ] What is a slice internally? What are the 3 fields?
- [ ] What happens when you append and cap is full?
- [ ] What is amortized O(1)?
- [ ] What is the slice sharing trap?
- [ ] Why is passing a slice to a function not fully "by reference"?
- [ ] Is a slice concurrent safe?

**Coding problems**
- [ ] Reverse an array in place
- [ ] Find the maximum sum subarray (Kadane's algorithm)
- [ ] Move all zeros to end without changing order of others
- [ ] Find duplicates in an array

---

### 1.2 Hash Map ✅
**Concepts — can you answer these?**
- [ ] What is a bucket and how does hashing work?
- [ ] What is a collision and how is it resolved?
- [ ] What is load factor and when does resize happen?
- [ ] Why is bucket count always a power of 2?
- [ ] Why is reading from a nil map safe but writing panics?
- [ ] Why is Go map not concurrent safe?
- [ ] What physically crashes during concurrent map access?
- [ ] What does delete do to memory?
- [ ] Is map iteration order guaranteed?

**Coding problems**
- [ ] Two Sum (return indices of two numbers that add to target)
- [ ] Group anagrams together
- [ ] Find the first non-repeating character in a string
- [ ] Word frequency counter
- [ ] Check if two strings are anagrams

---

### 1.3 Linked List 🔄
**Concepts — can you answer these?**
- [ ] What is a node? What fields does it have?
- [ ] Difference between singly and doubly linked list
- [ ] Why is insert at head O(1) but insert at tail O(n)?
- [ ] Why is linked list NOT cache friendly?
- [ ] When would you use linked list over slice?

**Coding problems**
- [x] Reverse a linked list in place
- [x] Find middle node (slow/fast pointer)
- [x] Detect cycle — hash set approach
- [x] Detect cycle — slow/fast pointer
- [x] Find nth node from end (two pointer gap technique)
- [x] Merge two sorted linked lists
- [ ] Remove duplicates from sorted linked list
- [ ] Check if linked list is a palindrome

---

### 1.4 Stack & Queue 🔄
**Concepts — can you answer these?**
- [x] What is LIFO vs FIFO?
- [x] How do you implement a stack in Go?
- [x] How do you implement a queue in Go?
- [x] What is a monotonic stack and when do you use it?
- [ ] How can you implement a queue using two stacks?

**Coding problems**
- [x] Valid parentheses (check if brackets are balanced)
- [x] Implement a min stack (stack that returns min in O(1))
- [x] Daily temperatures (next warmer day for each day)
- [ ] Implement queue using two stacks
- [ ] Evaluate reverse polish notation

---

### 1.5 Binary Tree & BST 🔄
**Concepts — can you answer these?**
- [x] What is a binary tree vs a BST?
- [x] What are inorder, preorder, postorder traversals?
- [x] What is BFS (level order) traversal?
- [x] What makes a valid BST?
- [ ] What is the height/depth of a tree?
- [ ] What is a balanced tree?

**Coding problems (Cycle 1 — traversals only)**
- [x] Inorder traversal (recursive and iterative)
- [x] Preorder traversal (recursive)
- [x] Postorder traversal (recursive)
- [x] Level order traversal (BFS with queue)

**Coding problems (Cycle 2 — deferred)**
- [ ] Max depth of binary tree
- [ ] Check if binary tree is balanced
- [ ] Validate BST
- [ ] Lowest common ancestor of two nodes
- [ ] Diameter of binary tree
- [ ] Invert a binary tree

---

### 1.6 Graph 🔄
**Concepts — can you answer these?**
- [x] What is a directed vs undirected graph?
- [x] What is an adjacency list vs adjacency matrix?
- [x] How does BFS work? When do you use it?
- [x] How does DFS work? When do you use it?
- [ ] What is topological sort and when do you use it?
- [x] What is a visited set and why do you need it?

**Coding problems**
- [ ] Number of islands (BFS or DFS on grid)
- [ ] Clone a graph
- [ ] Course schedule — can you finish all courses? (cycle detection)
- [ ] BFS shortest path between two nodes
- [ ] Topological sort

---

### 1.7 Binary Search 🔄
**Concepts — can you answer these?**
- [x] What are the conditions for binary search to work?
- [x] What is the loop condition — `<` or `<=`?
- [x] How do you avoid infinite loops with mid calculation?
- [x] When do you use `left = mid + 1` vs `left = mid`?

**Coding problems**
- [x] Classic binary search (find target in sorted array)
- [x] Search insert position
- [ ] Search in rotated sorted array
- [ ] Find first and last position of element
- [ ] Find minimum in rotated sorted array

---

### 1.8 Heap / Priority Queue 🔄
**Concepts — can you answer these?**
- [x] What is a min heap vs max heap?
- [x] How does insert and extract work in a heap?
- [x] What is the time complexity of heap operations?
- [ ] How do you use Go's `container/heap`?

**Coding problems**
- [ ] Kth largest element in an array
- [ ] Top K frequent elements
- [ ] Merge K sorted lists
- [ ] Find median from data stream

---

### 1.9 Two Pointers Pattern ⬜
**Concepts — can you answer these?**
- [ ] When do you use two pointers?
- [ ] What is the difference between same-direction and opposite-direction two pointers?

**Coding problems**
- [ ] Two sum on sorted array
- [ ] Container with most water
- [ ] Three sum (find all triplets that sum to zero)
- [ ] Remove duplicates from sorted array in place
- [ ] Trapping rain water

---

### 1.10 Sliding Window Pattern ⬜
**Concepts — can you answer these?**
- [ ] When do you use sliding window vs two pointers?
- [ ] What is a fixed-size vs variable-size window?

**Coding problems**
- [ ] Maximum sum subarray of size k (fixed window)
- [ ] Longest substring without repeating characters
- [ ] Minimum window substring containing all chars
- [ ] Longest subarray with sum <= k

---

### 1.11 Dynamic Programming — basics only ⬜
**Concepts — can you answer these?**
- [ ] What is memoization vs tabulation?
- [ ] How do you identify a DP problem?

**Coding problems**
- [ ] Fibonacci (top-down and bottom-up)
- [ ] Climbing stairs
- [ ] House robber
- [ ] Coin change (minimum coins for amount)
- [ ] Longest common subsequence

---

## Phase 2 — System Design

### 2.1 Fundamentals ✅
- [x] Vertical vs horizontal scaling
- [x] Load balancer — what it does, types (round robin, least connections)
- [x] CDN — what it does and when to use
- [x] Caching — write-through vs write-back vs write-around
- [x] Cache eviction — LRU, LFU, TTL
- [x] CAP theorem — consistency, availability, partition tolerance
- [x] SQL vs NoSQL — when to use which
- [x] Database replication — primary/replica, read replicas
- [x] Database sharding — what it is, horizontal vs vertical
- [x] Consistent hashing — why it matters for sharding
- [x] Message queues — what problems they solve, at-least-once vs exactly-once
- [x] Rate limiting strategies — token bucket, sliding window

### 2.2 Distributed Systems ⬜
- [ ] Partial failure and network request ambiguity
- [ ] Timeout, retry, exponential backoff and jitter
- [ ] Idempotency and duplicate request handling
- [ ] Eventual vs strong consistency
- [ ] Circuit breaker and bulkhead patterns
- [ ] Rate limiting, load shedding, backpressure
- [ ] Distributed lock limitations
- [ ] Clock and timestamp problems

### 2.3 Practice Designs ⬜
- [ ] Design URL shortener
- [ ] Design rate limiter
- [ ] Design notification system (email/push/SMS)
- [ ] Design file storage service
- [ ] Design authentication service
- [ ] Design payment/order service
- [ ] Design audit log service
- [ ] Design job scheduler or webhook delivery service

---

## Phase 3 — Go Core

### 3.1 Go Basics ⬜
- [ ] Zero value concept
- [ ] Named type vs type alias
- [ ] Constants and `iota`
- [ ] `rune` vs `byte`, UTF-8 string handling
- [ ] Value semantics vs pointer semantics
- [ ] Pointer receiver vs value receiver
- [ ] Array vs slice internals (len, cap, backing array)
- [ ] `append` behavior when cap is full
- [ ] Nil slice vs empty slice
- [ ] Map concurrent safety, `sync.Map`
- [ ] Struct embedding and composition
- [ ] Functional options pattern
- [ ] Implicit interface implementation
- [ ] Type assertion and type switch
- [ ] Nil interface vs typed nil problem
- [ ] Closure and variadic functions
- [ ] `defer`, `panic`, `recover`
- [ ] Generics — when to use vs interfaces

### 3.2 Error Handling ⬜
- [ ] `errors.New` vs `fmt.Errorf`
- [ ] Error wrapping with `%w`
- [ ] `errors.Is` and `errors.As`
- [ ] Sentinel errors
- [ ] Custom error types
- [ ] Error classification (validation, auth, conflict, internal)
- [ ] Repository errors → domain errors mapping
- [ ] Error → HTTP status code mapping
- [ ] Which layer logs the error
- [ ] Never expose sensitive error info to client

### 3.3 Context ⬜
- [ ] `context.Context` purpose
- [ ] `Background()` vs `TODO()`
- [ ] `WithCancel`, `WithTimeout`, `WithDeadline`
- [ ] Cancellation propagation
- [ ] Passing context from HTTP handler to repository
- [ ] Canceling DB and external HTTP requests via context
- [ ] Context value — proper usage and custom key types
- [ ] Why not store context in struct fields

### 3.4 Concurrency ⬜
- [ ] Goroutine vs OS thread
- [ ] Goroutine leak — what it is and how to prevent
- [ ] G-M-P scheduler model (basics)
- [ ] Unbuffered vs buffered channels
- [ ] Channel send/receive blocking behavior
- [ ] Who should close the channel
- [ ] Nil and closed channel behavior
- [ ] `select` and timeout patterns
- [ ] `sync.Mutex` and `sync.RWMutex`
- [ ] Deadlock vs data race
- [ ] `sync.WaitGroup`, `sync.Once`, `sync.Pool`
- [ ] `sync/atomic`
- [ ] Worker pool pattern
- [ ] Fan-in, fan-out, pipeline patterns
- [ ] Semaphore for concurrency limiting
- [ ] Backpressure concept
- [ ] Graceful shutdown
- [ ] `errgroup` usage
- [ ] `go test -race`

### 3.5 Language Internals ⬜
- [ ] Go garbage collector basics
- [ ] Interface internally (type + pointer)
- [ ] Nil interface gotcha
- [ ] How defer works with return values
- [ ] panic vs error
- [ ] Escape analysis
- [ ] Stack vs heap allocation

### 3.6 Standard Library ⬜
- [ ] `fmt`, `strings`, `strconv`, `bytes`
- [ ] `io`, `os`, `path/filepath`
- [ ] `time`, `errors`, `context`, `sync`
- [ ] `encoding/json` — custom marshal/unmarshal, reject unknown fields
- [ ] `net/http` and `database/sql` basics
- [ ] `log/slog` structured logging
- [ ] `io.Reader`, `io.Writer`, `io.Copy`
- [ ] Processing large files without loading into RAM

---

## Phase 4 — HTTP & Backend Patterns

### 4.1 HTTP & REST API ⬜
- [ ] HTTP request/response structure
- [ ] Safe and idempotent methods
- [ ] Common HTTP status codes
- [ ] Headers, cookies, Content-Type, Accept
- [ ] HTTP/1.1 vs HTTP/2 basics
- [ ] CORS
- [ ] Resource-oriented endpoint design
- [ ] POST vs PUT vs PATCH
- [ ] 401 vs 403 vs 404 vs 409 vs 422 vs 500
- [ ] Pagination, filtering, sorting
- [ ] API versioning
- [ ] Consistent error response format
- [ ] Idempotency key
- [ ] Middleware patterns (request ID, logging, recovery)
- [ ] Server and client timeouts
- [ ] Graceful shutdown
- [ ] Health and readiness endpoints
- [ ] `http.Client` reuse
- [ ] Retry with exponential backoff and jitter

### 4.2 Authentication & Security ⬜
- [ ] Authentication vs authorization
- [ ] JWT structure (header, payload, signature)
- [ ] Access token vs refresh token
- [ ] Token expiration and revocation
- [ ] Password hashing (bcrypt, Argon2)
- [ ] RBAC and permission-based authorization
- [ ] Object-level authorization checks
- [ ] SQL injection prevention
- [ ] XSS, CSRF, SSRF basics
- [ ] Rate limiting implementation
- [ ] Brute-force protection
- [ ] Secrets not in source code
- [ ] Never log sensitive data
- [ ] File upload validation and size limits

### 4.3 Testing ⬜
- [ ] Table-driven tests and subtests
- [ ] Mock vs fake vs stub
- [ ] Interface-based dependency mocking
- [ ] Error and edge case testing
- [ ] `httptest` for handler and middleware tests
- [ ] Repository integration tests
- [ ] Test database and migrations
- [ ] Testcontainers for dependencies
- [ ] Benchmark and `-benchmem`
- [ ] Fuzz testing
- [ ] Race detector (`go test -race`)
- [ ] Test coverage analysis
- [ ] End-to-end vs load test
- [ ] Load testing with k6 or similar

### 4.4 Performance & Profiling ⬜
- [ ] Measure first, optimize second
- [ ] CPU, heap, goroutine profiling
- [ ] `pprof` endpoints
- [ ] Benchmark comparison
- [ ] Reducing allocations
- [ ] Escape analysis output
- [ ] Stack vs heap
- [ ] GC basics
- [ ] Memory leak vs goroutine leak
- [ ] DB connection pool bottleneck detection
- [ ] Finding slow external APIs via tracing

---

## Phase 5 — PostgreSQL & Redis

### 5.1 SQL Queries ⬜
- [ ] JOIN, GROUP BY, HAVING, subquery, CTE
- [ ] EXISTS, UNION, CASE, window functions
- [ ] ROW_NUMBER, RANK, LAG, LEAD
- [ ] NULL and COALESCE handling
- [ ] Dynamic filter queries
- [ ] SQL parameterization

### 5.2 Schema Design ⬜
- [ ] Primary key, foreign key, unique, check constraints
- [ ] Normalization vs denormalization
- [ ] One-to-one, one-to-many, many-to-many relations
- [ ] Soft delete trade-offs

### 5.3 Indexing ⬜
- [ ] B-tree, composite, partial, covering indexes
- [ ] Composite index column order matters
- [ ] Index selectivity
- [ ] EXPLAIN and EXPLAIN ANALYZE
- [ ] Seq Scan vs Index Scan vs Index Only Scan vs Bitmap Scan
- [ ] Nested Loop vs Hash Join vs Merge Join
- [ ] Estimated rows vs actual rows

### 5.4 Transactions & Locking ⬜
- [ ] ACID properties
- [ ] Read Committed, Repeatable Read, Serializable
- [ ] Dirty read, non-repeatable read, phantom read, lost update
- [ ] SELECT FOR UPDATE
- [ ] Optimistic vs pessimistic locking
- [ ] Deadlock prevention
- [ ] Serialization error → retry transaction
- [ ] MVCC basics

### 5.5 Performance ⬜
- [ ] N+1 query problem and fix (sqlx, pgx)
- [ ] Connection pooling
- [ ] VACUUM and why Postgres needs it
- [ ] Table partitioning vs sharding
- [ ] Offset vs keyset pagination
- [ ] Context with queries

### 5.6 Redis & Caching ⬜
- [ ] Cache-aside pattern (practical)
- [ ] Read-through, write-through, write-behind
- [ ] TTL and cache invalidation strategies
- [ ] Cache key naming conventions
- [ ] Cache hit/miss metrics
- [ ] Cache stampede and hot key problems
- [ ] Redis data types: string, hash, set, sorted set
- [ ] Distributed lock risks
- [ ] Redis persistence and eviction policies

---

## Phase 6 — Kafka & Message Brokers

### 6.1 Core Concepts ⬜
- [ ] Broker, topic, partition, offset, consumer group
- [ ] Producer and consumer basics
- [ ] Partition ordering and key impact
- [ ] Replication, leader, follower
- [ ] Consumer rebalance

### 6.2 Delivery & Reliability ⬜
- [ ] At-most-once, at-least-once, exactly-once
- [ ] Auto commit vs manual commit
- [ ] When to commit offset
- [ ] Idempotent consumer
- [ ] Deduplication key
- [ ] Out-of-order event handling
- [ ] Retryable vs permanent errors
- [ ] Retry, backoff, jitter, DLQ
- [ ] Poison message handling
- [ ] Consumer graceful shutdown
- [ ] Consumer lag monitoring

### 6.3 Patterns ⬜
- [ ] Transactional outbox and inbox patterns
- [ ] Dual-write problem (DB + Kafka)
- [ ] Event schema versioning

---

## Phase 7 — Architecture

### 7.1 Patterns ⬜
- [ ] Layered, Clean, Hexagonal Architecture
- [ ] Dependency inversion principle
- [ ] Domain, use case, infrastructure layers
- [ ] Repository, Unit of Work, Adapter patterns
- [ ] Strategy, Factory, Decorator, Middleware patterns
- [ ] Dependency injection
- [ ] Circular dependency prevention
- [ ] Package boundary decisions

### 7.2 Monolith & Microservices ⬜
- [ ] Modular monolith concept
- [ ] Module boundaries and dependencies
- [ ] Modular monolith vs microservice trade-offs
- [ ] Microservice pros and cons
- [ ] Distributed transaction problems
- [ ] Service discovery and API gateway
- [ ] Saga and CQRS (when needed)

---

## Phase 8 — Infrastructure

### 8.1 Docker ⬜
- [ ] Image vs container
- [ ] Dockerfile and multi-stage build
- [ ] Minimal runtime image
- [ ] `.dockerignore`
- [ ] Environment variables, volumes, networks
- [ ] Health check in Docker
- [ ] Non-root user in container
- [ ] Signal handling and PID 1
- [ ] Graceful shutdown in container
- [ ] Docker Compose
- [ ] Image layers and caching
- [ ] Never put secrets in image

### 8.2 CI/CD ⬜
- [ ] CI vs CD difference
- [ ] Build, test, lint pipeline stages
- [ ] Docker image build and registry push
- [ ] Environment-based deployment
- [ ] Migration management in pipeline
- [ ] Rollback strategy
- [ ] Artifact vs cache
- [ ] Secret variable management
- [ ] Branch/tag-based pipeline control
- [ ] Post-deployment smoke tests

### 8.3 Kubernetes Basics ⬜
- [ ] Pod, Deployment, Service, Ingress
- [ ] ConfigMap and Secret
- [ ] Replicas and rolling update
- [ ] Liveness, readiness, startup probes
- [ ] Resource requests and limits
- [ ] Graceful termination
- [ ] Horizontal scaling basics
- [ ] Pod restart analysis
- [ ] Log and deployment status checks

---

## Phase 9 — Observability

### 9.1 Logging ⬜
- [ ] Structured logging
- [ ] Log levels — when to use each
- [ ] Request ID and Trace ID in logs
- [ ] Never log sensitive data

### 9.2 Metrics ⬜
- [ ] Counter vs Gauge vs Histogram
- [ ] Request count, error rate, latency metrics
- [ ] DB pool, Kafka lag, retry, DLQ metrics
- [ ] Prometheus metric exposure
- [ ] SLI, SLO, SLA differences
- [ ] p50, p95, p99 percentiles

### 9.3 Tracing ⬜
- [ ] Trace and span concepts
- [ ] Context propagation
- [ ] Tracing HTTP, DB, external API, Kafka flows
- [ ] OpenTelemetry basics

---

## Phase 10 — Production Incidents

### 10.1 Troubleshooting ⬜
- [ ] Service slowing down — investigation steps
- [ ] CPU at 100% — what to do
- [ ] Memory or goroutine count growing — what to do
- [ ] DB connection pool full — what to do
- [ ] Finding slow queries
- [ ] Kafka lag increasing — what to do
- [ ] DLQ filling up fast — what to do
- [ ] External API timeout — mitigation steps
- [ ] Error rate spike after deployment — rollback procedure
- [ ] Using logs, metrics, traces together
- [ ] Root cause analysis and postmortem writing

---

## Phase 11 — Behavioral & Soft Skills

### 11.1 Self-Introduction ⬜
- [ ] "Tell me about yourself" — 2 min prepared answer
- [ ] Why backend and Go specifically
- [ ] Explain current project in 5 minutes
- [ ] Explain current project architecture (draw diagram)

### 11.2 Experience Questions ⬜
- [ ] Biggest technical challenge (STAR format)
- [ ] Production incident and lessons learned
- [ ] Team disagreement — how you resolved it
- [ ] Tight deadline situation
- [ ] Code review disagreement handling
- [ ] Wrong technical decision — honest answer
- [ ] Strengths and areas for improvement

### 11.3 Career Questions ⬜
- [ ] Why are you leaving your current job?
- [ ] What excites you about this role?
- [ ] Where do you want to be in 2-3 years?
- [ ] Process when joining a new codebase
- [ ] Questions to ask the interviewer (prepared list)

### 11.4 Project Deep Dive ⬜
- [ ] Business goal of the project
- [ ] Your specific role and contributions
- [ ] Request and authentication flow
- [ ] Database schema decisions
- [ ] Transaction usage and why
- [ ] Kafka flow, retry, DLQ decisions
- [ ] Idempotency guarantees
- [ ] Error handling model
- [ ] Logging and tracing setup
- [ ] Biggest bottleneck
- [ ] What you would do differently

---

## Phase 12 — Portfolio & Final Prep

### 12.1 Portfolio Project ⬜
- [ ] Clean project structure
- [ ] README and architecture diagram
- [ ] REST API + PostgreSQL
- [ ] Migrations and transactions
- [ ] Auth and authorization
- [ ] Validation and error handling
- [ ] Structured logging
- [ ] Kafka producer or consumer
- [ ] Retry, DLQ, idempotency
- [ ] Redis cache
- [ ] Unit and integration tests
- [ ] Race and load tests
- [ ] Health and readiness checks
- [ ] Metrics and tracing
- [ ] Graceful shutdown
- [ ] Dockerfile and Docker Compose
- [ ] CI/CD pipeline

### 12.2 Final Checklist ⬜
- [ ] Can answer Go core questions
- [ ] Can live-code a concurrency problem
- [ ] Can optimize a SQL query
- [ ] Can explain transactions and isolation
- [ ] Can design a REST API
- [ ] Can answer Kafka reliability questions
- [ ] Can do a 45-min system design
- [ ] Can solve Easy in 15-20 min
- [ ] Can solve Medium in 30-40 min
- [ ] Can explain project confidently
- [ ] Can answer unknown questions professionally
- [ ] Completed at least 3 mock interviews
- [ ] Started applying to real positions

---

## Progress Summary

| Phase | Topics | Done |
|-------|--------|------|
| 1. DS & Algo | 11 topics | 4/11 in progress |
| 2. System Design | 3 sections | 1/3 |
| 3. Go Core | 6 sections | 0/6 |
| 4. HTTP & Backend | 4 sections | 0/4 |
| 5. PostgreSQL & Redis | 6 sections | 0/6 |
| 6. Kafka | 3 sections | 0/3 |
| 7. Architecture | 2 sections | 0/2 |
| 8. Infrastructure | 3 sections | 0/3 |
| 9. Observability | 3 sections | 0/3 |
| 10. Production | 1 section | 0/1 |
| 11. Behavioral | 4 sections | 0/4 |
| 12. Portfolio | 2 sections | 0/2 |

**Notes:** `notes/`
**Code:** `code/`
**Russian:** `russian/`
**English:** `english/`
