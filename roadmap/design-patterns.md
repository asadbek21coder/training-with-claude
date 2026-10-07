# Design Patterns & Architecture Roadmap

> **L1** = aware & can use · **L2** = deep & efficient · Status: ⬜ not started · 🔄 in progress · ✅ done
> **Notes:** `notes/design-patterns/` · **Practice:** `code/design-patterns/` · **Dictionary:** `reference/design-patterns.md`

**Why (EN):** Architecture is the set of decisions that are expensive to change, and a senior engineer is the one who sets the shape, boundaries and dependencies of a codebase so it stays easy to change. In the AI era code is cheap to generate but structure is not: AI fills in functions, while you must recognise when a pattern helps, when it is over-engineering, and how to keep generated code inside clean boundaries.
**Зачем (RU):** Архитектура — это решения, которые дорого менять, и senior-инженер задаёт форму, границы и зависимости кода так, чтобы его было легко менять. В эпоху ИИ код генерировать дёшево, а структуру — нет: ИИ пишет функции, а ты должен понимать, когда паттерн помогает, когда это переусложнение, и как удержать сгенерированный код в чистых границах.

## Progress
| Layer | Done | Total |
|---|---|---|
| L1 | 0 | 38 |
| L2 | 0 | 29 |

---

## DP 1 — Design principles fundamentals
**L1 — aware & can use**
- ⬜ **1.1 Coupling & cohesion** — high cohesion, low coupling, afferent/efferent coupling, temporal coupling, connascence (awareness), change amplification, "things that change together live together"
- ⬜ **1.2 Abstraction & encapsulation** — abstraction, encapsulation, information hiding, unexported identifiers, package boundary as API, leaky abstraction, interface vs implementation
- ⬜ **1.3 DRY, KISS, YAGNI** — Don't Repeat Yourself (knowledge, not text), KISS, YAGNI, rule of three, accidental vs essential duplication, "a little copying is better than a little dependency" (Go proverb)
- ⬜ **1.4 Separation of concerns** — separation of concerns, single level of abstraction, command–query separation (CQS), composition over inheritance, Law of Demeter, tell-don't-ask

**L2 — deep & efficient**
- ⬜ **1.5 Complexity theory (Ousterhout)** — deep vs shallow modules, cognitive load, unknown unknowns, define errors out of existence, pull complexity downwards, strategic vs tactical programming
- ⬜ **1.6 Trade-offs between principles** — DRY vs coupling, wrong abstraction ("duplication is cheaper than the wrong abstraction"), premature abstraction, flexibility vs simplicity, cost of indirection

🛠 **Practice** — L1: take one of your own Go packages and list its exported API; mark each exported identifier as "needed outside" or "should be unexported", then shrink the API · L2: find two duplicated code blocks in a real project and write a short note for each — merge or keep separate, with the coupling trade-off explained

---

## DP 2 — SOLID applied to Go
**L1 — aware & can use**
- ⬜ **2.1 SRP & OCP** — single responsibility (one reason to change, one actor), open/closed, extension via interfaces, extension via function values, plugin points, avoiding `switch` on type everywhere
- ⬜ **2.2 LSP & ISP** — Liskov substitution, behavioural contracts, `io.Reader` / `io.Writer` as small interfaces, interface segregation, "the bigger the interface, the weaker the abstraction", single-method interfaces
- ⬜ **2.3 DIP & dependency injection** — dependency inversion, depend on abstractions, constructor injection, "accept interfaces, return structs", consumer-defined interfaces, wiring in `main`

**L2 — deep & efficient**
- ⬜ **2.4 SOLID without classes** — implicit interface satisfaction, embedding vs inheritance, interface pollution, interfaces defined at the point of use, compile-time assertion `var _ I = (*T)(nil)`, when NOT to introduce an interface
- ⬜ **2.5 DI tooling & composition root** — manual DI, composition root, `google/wire` (code generation), `uber-go/fx` / `dig` (runtime), trade-offs of DI containers, testability via fakes

🛠 **Practice** — L1: refactor a service that calls `*sql.DB` directly so it depends on a small consumer-defined `UserStore` interface; write a test with an in-memory fake · L2: wire the same app twice — manually in `main.go` and with `wire` — and compare readability and startup errors

---

## DP 3 — Clean code, code smells & refactoring
**L1 — aware & can use**
- ⬜ **3.1 Readable code** — naming (short in small scope, descriptive in wide scope), function size, early return / guard clauses, avoid deep nesting, comments explain why, Go Code Review Comments, Effective Go
- ⬜ **3.2 Code smells** — long function, long parameter list, god object/package, feature envy, shotgun surgery, divergent change, primitive obsession, data clumps, magic numbers, dead code, boolean flag arguments
- ⬜ **3.3 Basic refactorings** — extract function, inline function, rename, introduce parameter object, replace magic number with constant, decompose conditional, move function, IDE/gopls refactor tools

**L2 — deep & efficient**
- ⬜ **3.4 Refactoring safely** — characterization tests, golden tests, small steps, refactor vs feature commits, strangler inside a codebase, branch by abstraction, parallel change (expand/contract)
- ⬜ **3.5 Technical debt management** — technical debt quadrant (Fowler), deliberate vs inadvertent debt, debt register, boy-scout rule, refactoring budget, explaining debt to product people
- ⬜ **3.6 Reviewing AI-generated code for design** — over-engineering, needless interfaces, duplicated helpers, inconsistent error handling, invented abstractions, mismatched package conventions

🛠 **Practice** — L1: write a deliberately messy 200-line `net/http` handler (validation, SQL, business rules, JSON in one function), then refactor it into handler → service → repository with tests passing at every step · L2: before refactoring, add golden/characterization tests for the handler's HTTP responses and keep them green through every commit

---

## DP 4 — Creational patterns
**L1 — aware & can use**
- ⬜ **4.1 Constructors & factory** — `NewX` constructor functions, zero-value usability, factory function, factory method, returning concrete types vs interfaces, validation in constructors
- ⬜ **4.2 Functional options** — functional options (`WithTimeout(d)`), option type `func(*config)`, defaults, option interface variant, config struct alternative, variadic options
- ⬜ **4.3 Builder** — builder pattern, fluent API, `strings.Builder`, query builders (squirrel), when a config struct is simpler, immutability of built objects

**L2 — deep & efficient**
- ⬜ **4.4 Singleton & global state** — singleton, `sync.Once`, `sync.OnceValue`, package-level variables, `init()` pitfalls, hidden global dependencies, why singletons hurt testing
- ⬜ **4.5 Prototype, object pools & abstract factory** — prototype (clone), deep vs shallow copy, `sync.Pool`, abstract factory (families of objects), registries (`database/sql` driver registration)

🛠 **Practice** — L1: build an HTTP client wrapper for a payment provider with `NewClient(baseURL string, opts ...Option)` supporting `WithTimeout`, `WithRetry`, `WithLogger` · L2: implement a `database/sql`-style driver registry (`Register(name, factory)`) for notification senders (email, SMS, Telegram) and remove all global state from tests

---

## DP 5 — Structural patterns
**L1 — aware & can use**
- ⬜ **5.1 Adapter** — adapter, wrapping third-party SDKs, `http.HandlerFunc` as function adapter, anti-corruption at code level, converting between DTO and domain types
- ⬜ **5.2 Decorator & middleware** — decorator, wrapping interfaces (`io.Reader` → `bufio.Reader`), HTTP middleware `func(http.Handler) http.Handler`, logging/metrics/retry decorators, `http.RoundTripper` wrappers
- ⬜ **5.3 Facade & proxy** — facade (simple API over subsystem), proxy (caching, lazy, access-control, remote), `httputil.ReverseProxy`, gRPC client stubs as proxies

**L2 — deep & efficient**
- ⬜ **5.4 Composite & bridge** — composite (tree of uniform nodes), recursive structures, `fs.FS` composition, bridge (abstraction vs implementation), flyweight (awareness), string interning
- ⬜ **5.5 Embedding as composition** — struct embedding, interface embedding, method promotion, shadowing, embedding pitfalls (exposing unwanted methods, embedding `sync.Mutex`), delegation vs embedding

🛠 **Practice** — L1: write a `Store` interface and wrap it with a caching decorator and a metrics decorator; write `RoundTripper` middleware that logs and retries outgoing HTTP calls · L2: build an adapter over two fake SMS providers with different APIs, plus a failover proxy that switches providers on errors

---

## DP 6 — Behavioral patterns
**L1 — aware & can use**
- ⬜ **6.1 Strategy** — strategy pattern, interface vs function value, `sort.Slice` with `less` func, pricing/discount strategies, choosing strategy at runtime, map of strategies
- ⬜ **6.2 Chain of responsibility** — chain of responsibility, middleware chains, handler pipeline, short-circuit, `next` handler, gRPC interceptors, validation chains
- ⬜ **6.3 Observer & pub/sub** — observer, subscriber list, callbacks, channels as event streams, in-process event bus, synchronous vs asynchronous notification, unsubscribe & leaks
- ⬜ **6.4 Command & iterator** — command (request as object), undo/redo, job queues, iterator, range-over-func iterators (`iter.Seq`, `iter.Seq2`), pull vs push iterators, `bufio.Scanner` as iterator

**L2 — deep & efficient**
- ⬜ **6.5 State machine** — state pattern, finite state machine, transition table, enum with `iota`, invalid transitions, order/payment status lifecycle, persisting state, state as type
- ⬜ **6.6 Template method, visitor & memento** — template method (via function hooks in Go), visitor (`go/ast.Walk`), memento (snapshots), mediator, why many GoF patterns collapse into functions in Go

🛠 **Practice** — L1: implement a discount engine with strategies chosen by customer type; implement an order-placement pipeline as a chain of middleware-style steps · L2: implement an order state machine (`created → paid → shipped → delivered / cancelled`) with a transition table, persisted status in Postgres, and tests for every invalid transition

---

## DP 7 — Go idioms & concurrency patterns overview
**L1 — aware & can use**
- ⬜ **7.1 Go idioms vs classic OOP patterns** — small interfaces, errors as values, zero value, composition, closures as objects, package-oriented design, "clear is better than clever", Go Proverbs
- ⬜ **7.2 Concurrency patterns catalogue** — worker pool, pipeline, fan-out/fan-in, generator, semaphore (buffered channel), `errgroup`, rate limiter, pub/sub with channels, graceful shutdown

**L2 — deep & efficient**
- ⬜ **7.3 Choosing a concurrency pattern** — bounded vs unbounded concurrency, backpressure, cancellation via `context`, ownership of channels, goroutine leaks, mutex vs channel decision, singleflight, request coalescing
- ⬜ **7.4 Anti-patterns in Go** — Java-style getters/setters, interface for every struct, `utils`/`common` packages, panics for control flow, deep package hierarchies, premature generics, stuttering names

🛠 **Practice** — L1: build a URL-checker CLI using a worker pool with `errgroup` and a concurrency limit, cancellable with Ctrl+C · L2: rewrite one Java-style Go package (getters, interface per struct, `utils`) into idiomatic Go and write a before/after note

---

## DP 8 — Application architecture styles
**L1 — aware & can use**
- ⬜ **8.1 Layered architecture** — presentation/service/data layers, handler → service → repository, DTOs, strict vs relaxed layering, layer leaks (SQL in handlers)
- ⬜ **8.2 Clean & hexagonal architecture** — clean architecture, dependency rule, entities/use cases, hexagonal (ports & adapters), primary/driving vs secondary/driven adapters, onion architecture
- ⬜ **8.3 Go project layout** — `cmd/`, `internal/`, package by feature vs package by layer, `golang-standards/project-layout` (unofficial), flat structure for small services, import cycles as design signal

**L2 — deep & efficient**
- ⬜ **8.4 Vertical slices & modular design** — vertical slice architecture, feature modules, module public API, dependency direction checks, `go-arch-lint` / depguard, cost of over-layering small services
- ⬜ **8.5 Transactions & cross-cutting concerns across layers** — unit of work, transaction in context vs explicit `Tx` param, transactional outbox placement, logging/auth/tracing as middleware, where validation belongs

🛠 **Practice** — L1: build a small "todo/bookmarks" service in hexagonal style: domain package with no imports of `net/http` or `database/sql`, HTTP and Postgres adapters, in-memory adapter for tests · L2: add a use case that updates two aggregates in one transaction and implement it with a unit-of-work interface; enforce dependency direction with a linter rule

---

## DP 9 — Domain-Driven Design
**L1 — aware & can use**
- ⬜ **9.1 Strategic DDD** — domain, subdomain (core, supporting, generic), ubiquitous language, bounded context, context map, domain expert, event storming (awareness)
- ⬜ **9.2 Tactical building blocks** — entity, value object, aggregate, aggregate root, repository, domain service, application service, factory, domain event
- ⬜ **9.3 Rich vs anemic domain model** — anemic domain model, invariants inside aggregates, constructors that validate, unexported fields, behaviour methods instead of setters

**L2 — deep & efficient**
- ⬜ **9.4 Aggregate design** — small aggregates, consistency boundary, reference by ID, one aggregate per transaction, optimistic concurrency (version column), eventual consistency between aggregates
- ⬜ **9.5 Context integration** — context map patterns (shared kernel, customer–supplier, conformist, anti-corruption layer, open host service, published language), translation between contexts
- ⬜ **9.6 When DDD is overkill** — CRUD domains, transaction script, table module, cost of tactical DDD in Go, lightweight DDD, domain complexity assessment

🛠 **Practice** — L1: model an `Order` aggregate in Go with `Money` and `OrderLine` value objects, invariants enforced in methods (no exported setters) and table-driven tests · L2: split a toy shop into `catalog` and `ordering` bounded contexts with an anti-corruption layer translating catalog products into ordering's own types

---

## DP 10 — Monolith vs microservices
**L1 — aware & can use**
- ⬜ **10.1 Monolith & modular monolith** — monolith, modular monolith, module boundaries, in-process calls, single deployment, majestic monolith, when monolith is the right default
- ⬜ **10.2 Microservices basics** — microservices, independent deployability, service boundaries, database per service, API contracts, team ownership, operational overhead

**L2 — deep & efficient**
- ⬜ **10.3 Distributed system costs** — distributed monolith anti-pattern, network failures, latency chains, data consistency, distributed transactions, saga, observability cost, versioning across services
- ⬜ **10.4 Decomposition & migration** — decompose by business capability / subdomain, strangler fig, branch by abstraction, extracting a service, shared database anti-pattern, data migration between services
- ⬜ **10.5 Service integration patterns** — API gateway, backend for frontend (BFF), service mesh (awareness), sync vs async communication, sidecar, service discovery

🛠 **Practice** — L1: draw (on paper) your current or past project as a module map; mark which modules could become services and why they should not yet · L2: build a modular monolith with two Go modules talking through an interface, then extract one into a separate service over HTTP/gRPC behind the same interface (strangler style)

---

## DP 11 — Event-driven architecture, CQRS & event sourcing
**L1 — aware & can use**
- ⬜ **11.1 Event-driven basics** — event, command vs event, producer/consumer, message broker, event notification, event-carried state transfer, at-least-once delivery, idempotent consumer
- ⬜ **11.2 Choreography vs orchestration** — choreography, orchestration, saga (orchestrated/choreographed), compensating actions, process manager, workflow engines (Temporal — awareness)
- ⬜ **11.3 CQRS** — command/query responsibility segregation, read model, write model, projections, denormalized views, eventual consistency in UI

**L2 — deep & efficient**
- ⬜ **11.4 Event sourcing** — event store, append-only log, rehydration, snapshots, projections, upcasting/event versioning, replay, GDPR and event sourcing, when not to use it
- ⬜ **11.5 Reliable eventing** — transactional outbox, inbox/dedup table, change data capture (Debezium), ordering guarantees, poison messages, dead letter queue, schema registry

🛠 **Practice** — L1: implement an in-process event bus in Go where `OrderPlaced` triggers email and stock reservation handlers; make handlers idempotent · L2: implement a tiny event-sourced `BankAccount` in Go + Postgres (events table, rehydration, snapshot every N events) and a read-model projection with a rebuild command

---

## DP 12 — Architecture decisions & documentation
**L1 — aware & can use**
- ⬜ **12.1 ADRs** — Architecture Decision Record, context/decision/consequences, status (proposed/accepted/superseded), Michael Nygard template, MADR, ADRs in repo `docs/adr/`
- ⬜ **12.2 Diagrams & C4 model** — C4 model (context, container, component, code), sequence diagrams, Mermaid, PlantUML, Structurizr (awareness), diagrams as code
- ⬜ **12.3 Design docs & reviews** — design doc / RFC, goals & non-goals, alternatives considered, risks, design review meeting, trade-off tables

**L2 — deep & efficient**
- ⬜ **12.4 Organization & architecture** — Conway's law, inverse Conway maneuver, Team Topologies (stream-aligned, platform, enabling, complicated-subsystem), cognitive load, ownership boundaries
- ⬜ **12.5 Quality attributes & trade-off analysis** — quality attributes (-ilities), ATAM (awareness), architecture characteristics, reversible vs irreversible decisions ("one-way / two-way doors"), cost of delay

🛠 **Practice** — L1: write three ADRs for decisions in one of your past projects (e.g. database choice, monolith vs services, auth approach); draw C4 context and container diagrams in Mermaid · L2: write a 2-page design doc for adding notifications to a service, with at least two alternatives and a trade-off table

---

## DP 13 — Advanced architecture (awareness)
**L1 — aware & can use**
- ⬜ **13.1 Evolutionary architecture** — evolutionary architecture, fitness functions, architecture tests, incremental change, last responsible moment
- ⬜ **13.2 Multi-tenancy** — single-tenant vs multi-tenant, shared schema with `tenant_id`, schema per tenant, database per tenant, noisy neighbour, row-level security, tenant isolation

**L2 — deep & efficient**
- ⬜ **13.3 Cell-based & large-scale architecture** — cell-based architecture, blast radius, shuffle sharding, control plane vs data plane, static stability, regional isolation
- ⬜ **13.4 Plugin & extensibility architectures** — plugin architecture, Go `plugin` package limits, HashiCorp go-plugin (RPC), WebAssembly plugins (awareness), microkernel architecture, extension points

🛠 **Practice** — L1: add a multi-tenant `tenant_id` + Postgres row-level security policy to a toy table and prove one tenant cannot read another's rows · L2: write an architecture test in Go that fails CI if `internal/domain` imports any adapter package (a simple fitness function)
