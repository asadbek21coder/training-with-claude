# 09 — Architecture & Design Patterns

> Know not just what the pattern is, but WHEN to use it and WHAT problem it solves.

---

## SOLID Principles

**Key Terms:**
- `SRP (Single Responsibility Principle)` — a class should have only one reason to change; one job
- `OCP (Open/Closed Principle)` — open for extension, closed for modification; add behavior without changing existing code
- `LSP (Liskov Substitution Principle)` — subclass must be substitutable for its parent; don't break callers
- `ISP (Interface Segregation Principle)` — don't force clients to depend on methods they don't use; prefer many small interfaces
- `DIP (Dependency Inversion Principle)` — depend on abstractions, not concretions; high-level modules don't import low-level ones

**Examples:**
- **SRP violation:** `UserService` handles user creation, email sending, and PDF generation → split into UserService + EmailService + ReportService
- **OCP violation:** adding new payment method requires editing `PaymentProcessor` switch/if chain → use interface `PaymentProvider` with implementations
- **LSP violation:** `Square extends Rectangle`; setting width changes height → Square breaks Rectangle's contract
- **ISP violation:** `Animal` interface has `fly()` + `swim()` + `run()`; Dog must implement `fly()` → split interfaces
- **DIP violation:** `OrderService` creates `new MySQLOrderRepository()` → inject `OrderRepository` interface

**Must understand:**
- [ ] Explain each with a concrete example
- [ ] How DIP relates to dependency injection frameworks (Spring, Go interfaces)
- [ ] LSP — most subtle one; contracts matter more than inheritance hierarchy
- [ ] OCP — how to extend behavior with strategy pattern, plugins, not if-chains

---

## Creational Design Patterns

**Key Terms:**
- `Singleton` — ensure only one instance; global access point
  - `sync.Once` in Go; `@Component` (Spring singleton scope)
  - Risk: hidden global state; hard to test; use DI instead
- `Factory Method` — define interface for creating objects; subclasses decide which class to instantiate
  - `http.Handler` registered in router; different handlers for different routes
- `Abstract Factory` — create families of related objects without specifying concrete classes
  - `sql.Open("postgres", dsn)` vs `sql.Open("mysql", dsn)` — same interface, different driver
- `Builder` — construct complex objects step by step; especially when many optional fields
  - `http.NewRequest(...)`, URL builder, query builder
  - Functional options pattern in Go is Builder variant
- `Prototype` — create objects by cloning an existing instance
  - Less common in Go/Java; `data class copy()` in Kotlin; `Object.clone()` in Java

**Must understand:**
- [ ] Singleton via sync.Once vs via package-level var (atomic init)
- [ ] Builder vs functional options in Go — same problem, different syntax
- [ ] When Factory is better than `new` (runtime decision of which type to create)

---

## Structural Design Patterns

**Key Terms:**
- `Adapter` — convert interface of a class into interface clients expect; bridge incompatible interfaces
  - `http.HandlerFunc` — adapts `func(w, r)` to `http.Handler` interface
  - Wrapping third-party library to match your interface
- `Decorator` — add behavior to object by wrapping it; same interface
  - HTTP middleware: `LoggingMiddleware(AuthMiddleware(handler))`
  - `io.TeeReader`, `io.LimitReader` — wrap io.Reader to add behavior
- `Proxy` — placeholder that controls access to another object; same interface
  - Lazy loading, access control, caching, remote call
  - Spring AOP creates proxies for @Transactional beans
- `Facade` — simplified interface to a complex subsystem
  - `s3Client.UploadFile(path)` hides multipart upload complexity
  - Service layer hiding repository + cache + queue complexity from controller
- `Composite` — tree structure where leaves and branches share same interface
  - File system: folder and file both implement `Node`
  - UI: widget containers and leaf widgets
- `Bridge` — separate abstraction from implementation so both can vary independently
  - Database abstraction layer; `sql.DB` + different drivers
- `Flyweight` — share common state between many fine-grained objects
  - String interning; object pooling; sync.Pool

**Must understand:**
- [ ] Adapter vs Facade vs Proxy — they're often confused
  - Adapter: makes incompatible interfaces work together
  - Facade: simplifies a complex subsystem
  - Proxy: same interface, adds behavior (like logging, auth, caching)
- [ ] Middleware is the Decorator pattern in HTTP
- [ ] Spring AOP = Proxy pattern under the hood

---

## Behavioral Design Patterns

**Key Terms:**
- `Observer` — objects subscribe to events from subject; publisher-subscriber
  - Event systems, message queues, React state updates
  - In Go: channels are a natural Observer implementation
- `Strategy` — define family of algorithms, encapsulate each, make them interchangeable
  - Sorting: pass different comparator functions
  - Payment: `CreditCardStrategy`, `PayPalStrategy` implement `PaymentStrategy`
- `Command` — encapsulate request as an object; enables undo, queue, log
  - HTTP request is a Command; undo history in editors
- `Iterator` — traverse collection without exposing internal structure
  - Go: `for range` loop; Java: `Iterator<T>` interface
  - `bufio.Scanner` — iterates over lines
- `Template Method` — define algorithm skeleton in base class; let subclasses fill in steps
  - Go: use functional parameter instead of inheritance (closures)
- `State` — allow object to alter behavior when internal state changes
  - TCP connection states; order status machine (created→paid→shipped→delivered)
- `Chain of Responsibility` — pass request through chain of handlers; each decides to handle or pass
  - HTTP middleware chain; logging filter chain in Spring Security
- `Mediator` — objects communicate through mediator instead of directly
  - Chat room (users don't know each other; mediator routes messages)
  - API gateway is a Mediator
- `Memento` — save and restore object state; undo functionality
  - Text editor undo; database transaction savepoints
- `Visitor` — add operations to objects without modifying them
  - AST visitors in compilers; SQL query builder

**Must understand:**
- [ ] Strategy — how it replaces if-else chains with polymorphism or function parameters
- [ ] Observer — how event-driven systems use it (domain events, message queues)
- [ ] Chain of Responsibility — middleware chains in Go and Java
- [ ] State machine — how to implement a simple FSM (state + transition table)

---

## Clean Architecture / Hexagonal Architecture

**Key Terms:**
- `Clean Architecture` — Robert Martin; concentric circles; dependencies point inward only
- `layers` — Entities → Use Cases → Interface Adapters → Frameworks/Drivers
- `dependency rule` — inner layers know nothing about outer layers
- `entities` — core business objects; no framework dependency; just business logic
- `use cases (interactors)` — application business rules; orchestrate entities
- `interface adapters` — convert data between use cases and frameworks (controllers, presenters, gateways)
- `frameworks / drivers` — outermost layer; DB, web framework, UI; replaceable
- `hexagonal architecture (ports and adapters)` — same idea; ports = interfaces; adapters = implementations
- `port` — interface defining what the domain needs from outside (e.g., `UserRepository` interface)
- `primary adapter` — drives the app (HTTP controller, CLI, tests)
- `secondary adapter` — driven by the app (PostgreSQL impl of UserRepository, Redis impl of Cache)
- `testability` — domain logic can be tested without DB, HTTP, or any framework
- `dependency inversion` — domain defines interfaces; infrastructure implements them

**Folder structure in Go:**
```
/domain      — entities, use case interfaces
/application — use case implementations
/adapter
  /http      — HTTP handlers (primary adapter)
  /postgres  — DB implementations (secondary adapter)
```

**Must understand:**
- [ ] Why the dependency rule matters (can swap DB without touching business logic)
- [ ] Hexagonal = ports + adapters; Clean = concentric circles; same idea, different metaphor
- [ ] How to test use cases without real DB (inject mock/in-memory adapter)
- [ ] The risk: over-engineering for simple CRUD apps; apply when complexity warrants it

---

## Domain-Driven Design (DDD) Basics

**Key Terms:**
- `domain` — the problem space your software solves
- `domain model` — object model that represents domain concepts and rules
- `entity` — object with identity that persists over time (User, Order, Product)
- `value object` — object defined by its attributes; no identity; immutable (Money, Address, Email)
- `aggregate` — cluster of entities treated as unit; single root entity is the aggregate root
- `aggregate root` — only entry point to the aggregate; enforces invariants
- `invariant` — rule that must always be true (Order total must match sum of items)
- `repository` — collection-like interface for aggregates; hides persistence
- `domain service` — business logic that doesn't naturally fit in an entity (MoneyTransferService)
- `application service` — orchestrates use cases; thin layer calling domain objects
- `domain event` — something that happened in the domain (OrderPlaced, PaymentFailed)
- `bounded context` — boundary within which a model is consistent; different contexts may use different models for same concept
- `context map` — shows relationships between bounded contexts
- `ubiquitous language` — shared language between developers and domain experts; used in code and conversations
- `anti-corruption layer` — translates between bounded contexts to prevent one model polluting another

**Must understand:**
- [ ] Entity vs value object — identity vs attributes; value objects are immutable
- [ ] Aggregate root — why only the root handles external access (maintains consistency)
- [ ] Bounded context — one service = usually one bounded context
- [ ] Why ubiquitous language matters — bugs often come from misaligned terminology

---

## Event-Driven Architecture

**Key Terms:**
- `event` — immutable record that something happened in the past (past tense: OrderPlaced, not PlaceOrder)
- `event producer` — service that publishes events when something happens
- `event consumer` — service that subscribes and reacts to events
- `event broker` — middleware that routes events (Kafka, RabbitMQ, SNS+SQS)
- `event schema` — structure of event payload; versioning is critical
- `event sourcing` — store events as source of truth; replay to rebuild state (different from event-driven)
- `choreography` — services react to each other's events; no central coordinator; loose coupling
- `orchestration` — central coordinator tells services what to do; more control, more coupling
- `fan-out` — one event triggers multiple consumers
- `idempotent consumer` — handles duplicate events safely
- `event ordering` — in Kafka, ordering only within partition; global ordering is hard
- `event versioning` — adding fields is safe; removing/renaming is breaking change
- `outbox pattern` — write event to DB outbox table in same transaction as business data; separate process publishes to broker

**Must understand:**
- [ ] Choreography vs orchestration — when each is better
- [ ] Why events are past-tense and immutable
- [ ] Outbox pattern — ensures at-least-once event publishing without distributed transaction
- [ ] Idempotent consumers — why they're required with at-least-once delivery

---

## CQRS

**Key Terms:**
- `CQRS (Command Query Responsibility Segregation)` — separate read model from write model
- `command` — action that changes state; returns nothing or just ID (CreateOrder)
- `query` — read that returns data; no side effects (GetOrdersByUser)
- `write model` — normalized, consistent, transactional; optimized for writes
- `read model (projection)` — denormalized, optimized for specific read pattern
- `projection` — derived view from events or write model; updated asynchronously
- `eventual consistency` — read model lags behind write model briefly
- `separate DB` — read model can use different DB optimized for reads (Elasticsearch for search, Redis for hot data)
- `event sourcing + CQRS` — natural combination; events update projections

**When to use:**
- Complex domain with many different read patterns
- Read and write have very different scaling needs
- High throughput reads that need denormalized data

**When NOT to use:**
- Simple CRUD apps — massive over-engineering
- When eventual consistency causes UX problems

**Must understand:**
- [ ] CQRS does NOT require event sourcing (can use same DB)
- [ ] The trade-off: eventual consistency between write and read models
- [ ] How projections are built and kept up-to-date

---

## Event Sourcing

**Key Terms:**
- `event store` — append-only log of all domain events; the source of truth
- `projection / read model` — current state derived by replaying events
- `replay` — recalculate state by replaying all events from the beginning
- `snapshot` — periodic checkpoint to avoid replaying all events from beginning
- `append-only` — events never modified or deleted; immutable history
- `temporal query` — reconstruct state at any point in time by replaying up to that time
- `audit log` — event store is a perfect audit log (who did what, when)
- `upcasting` — handle schema evolution of old events when replaying
- `event schema versioning` — critical problem; old events must still be processable

**Advantages:**
- Full audit trail
- Temporal queries (what was the order state last Tuesday?)
- Reconstruct any past state
- Debugging (replay events to reproduce bug)

**Disadvantages:**
- Complex implementation
- Replay can be slow without snapshots
- Schema evolution of events is hard
- Not suitable for all domains (simple CRUD doesn't benefit)

**Must understand:**
- [ ] Event sourcing ≠ just using Kafka (events must be source of truth, not a side effect)
- [ ] Projection vs event store — what each is responsible for
- [ ] Why snapshots are needed for long-lived aggregates
- [ ] When event sourcing is NOT appropriate

---

## Saga Pattern

**Key Terms:**
- `saga` — sequence of local transactions; each step publishes event or message to trigger next step
- `choreography-based saga` — each service listens for events and decides next action; decentralized
- `orchestration-based saga` — central saga orchestrator tells each service what to do; centralized
- `compensating transaction` — undo a completed step when later step fails (like "business rollback")
- `distributed transaction` — span multiple services; saga replaces 2PC
- `2PC (Two-Phase Commit)` — distributed atomic transaction; saga avoids 2PC complexity
- `eventual consistency` — saga doesn't provide immediate consistency; system converges
- `idempotency` — saga steps must be idempotent (retry-safe)
- `rollback scenario` — step 3 fails → run compensating transactions for steps 2, 1

**Example (order service):**
1. Create Order (local transaction)
2. → Reserve Inventory (event triggers inventory service)
3. → Process Payment (event triggers payment service)
4. → Confirm Order
5. If payment fails → trigger compensating: Release Inventory → Cancel Order

**Must understand:**
- [ ] Why saga replaces 2PC (2PC is synchronous, blocking, locks resources)
- [ ] Choreography vs orchestration — choreography = more decoupled but harder to trace
- [ ] Compensating transactions — they're not rollbacks (can't undo side effects like sent emails)
- [ ] Why idempotency is essential in sagas (messages may be retried)

---

## Circuit Breaker Pattern

**Key Terms:**
- `circuit breaker` — proxy that monitors calls to a remote service; opens when failures exceed threshold
- `closed state` — normal operation; calls pass through; failures counted
- `open state` — circuit broken; calls fail immediately without hitting service (fast fail)
- `half-open state` — after timeout, try one request; if success → close; if fail → re-open
- `failure threshold` — % or count of failures to trip the breaker
- `reset timeout` — how long to stay open before trying half-open
- `fallback` — alternative response when circuit is open (cache, default value, error message)
- `Resilience4j` — Java circuit breaker library (replaces deprecated Hystrix)
- `gobreaker` — Go circuit breaker library
- `bulkhead pattern` — isolate failures; limit concurrent calls to each service (separate thread pools)
- `timeout` — always set timeouts on external calls; without timeout circuit breaker can't help

**State machine:**
```
CLOSED → (failure rate > threshold) → OPEN
OPEN → (after timeout) → HALF-OPEN
HALF-OPEN → (success) → CLOSED
HALF-OPEN → (failure) → OPEN
```

**Must understand:**
- [ ] Why circuit breaker prevents cascade failures (broken service → stops its caller from being overwhelmed)
- [ ] The three states and transitions
- [ ] Fallback strategy — what to do when circuit is open
- [ ] Circuit breaker ≠ retry; retry for transient failures, circuit breaker for sustained failures

---

## 12-Factor App

**Key Terms:**
- `I. Codebase` — one codebase per app; deploy same code to different envs
- `II. Dependencies` — explicitly declare all dependencies; no system-wide packages
- `III. Config` — store config in environment variables; not in code or config files committed to repo
- `IV. Backing services` — treat DB, cache, queue as attached resources; swap without code change
- `V. Build, Release, Run` — strictly separate build stage, release stage, run stage
- `VI. Processes` — execute app as stateless processes; share nothing
- `VII. Port binding` — export services via port; self-contained (embedded server like Tomcat)
- `VIII. Concurrency` — scale out via process model; horizontal scaling
- `IX. Disposability` — fast startup, graceful shutdown; robust against sudden death
- `X. Dev/Prod Parity` — keep dev, staging, production as similar as possible
- `XI. Logs` — treat logs as event streams; write to stdout; let platform collect
- `XII. Admin processes` — run admin/management tasks as one-off processes (migrations)

**Must understand:**
- [ ] Config in env vars — why (different config per environment without code change)
- [ ] Stateless processes — why (enables horizontal scaling; any instance serves any request)
- [ ] Logs as streams — why stdout (composable with log aggregators; no file rotation needed)
- [ ] Dev/prod parity — why it matters (production bugs that don't reproduce in dev)

---

## API Versioning

**Key Terms:**
- `URL path versioning` — `/v1/users`, `/v2/users`; most visible; easy to route
- `query parameter versioning` — `/users?version=1`; ugly; breaks caching
- `header versioning` — `Accept-Version: v1`; clean URLs; harder to test in browser
- `content negotiation` — `Accept: application/vnd.myapi.v1+json`; REST purists prefer this
- `backward compatibility` — new version must not break existing clients (additive changes only)
- `deprecation` — mark old version as deprecated; communicate sunset date; keep running
- `sunset header` — `Sunset: Sat, 31 Dec 2025 23:59:59 GMT`; tells client when API dies
- `semantic versioning` — MAJOR.MINOR.PATCH; MAJOR = breaking change
- `breaking change` — removing field, renaming field, changing type, changing semantics
- `non-breaking change` — adding optional field, adding new endpoint

**Must understand:**
- [ ] URL versioning is most common and practical despite not being "REST purist"
- [ ] What constitutes a breaking change vs non-breaking
- [ ] How to deprecate gracefully (communicate, give timeline, keep running)
- [ ] Consumer-driven contract testing — verify clients aren't broken by new version
