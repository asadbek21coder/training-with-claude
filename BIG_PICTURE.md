# The Big Picture — Every Track, Topic, Subtopic and Term

> One file, the whole map. Pick **one small piece** a day (a subtopic, or even a single term), learn it, discuss it, mark its status.
> Один файл — вся карта. Каждый день выбирай **один маленький кусок** (подтему или даже один термин), изучай, обсуждай, отмечай статус.

**Why / books / AI-era priorities:** `reference/00-senior-engineer-guide.md`
**Term definitions (dictionaries):** `reference/*.md`
**Daily log:** `LOG.md`

---

## How to use / Как пользоваться

### Status legend (per subtopic) / Статусы
| Mark | Cycle | Meaning (EN) | Значение (RU) |
|---|---|---|---|
| ⬜ | — | Not checked yet | Ещё не проверено |
| 🟡 | C1 | **Aware & can use**: I know what it is, why it exists, how to use it, and can explain it to my student | Знаю, что это, зачем, как использовать, и могу объяснить студенту |
| 🟢 | C2 | **Deep**: internals, trade-offs, alternatives, edge cases | Глубоко: внутреннее устройство, компромиссы, альтернативы |
| 🔵 | C3 | **Own it**: can build it / implement it from scratch | Могу построить сам с нуля |

### IDs — how to pick a piece
Every subtopic has an ID: `GO 8.3` = Go track → topic 8 → subtopic 3. Tell Claude: *"today GO 8.3"*, or *"today one term from GO 20.2"*.

### Two modes for a session
1. **Check mode (for things you probably know):** Claude asks 3–5 questions first → if you answer well, mark 🟡 and move on fast; if not → switch to learn mode. *This is how we measure how much Go you already know.*
2. **Learn mode:** problem → concept → terms → thinking questions → senior answer (EN + RU) → small task → notes saved.

### Daily routine / Ежедневный распорядок
| Block | Time | What |
|---|---|---|
| **Main** | 40–50 min | **Current main track** (now: **GO**) — one subtopic or a few terms |
| **DSA** | 10–20 min | One concept, one term, or one problem from Track DSA |
| **EN** | 5–10 min | One new term / phrase / grammar point (+ corrections happen all session) |
| **RU** | 5–10 min | One word / phrase / grammar point |
| **Log** | 1 min | One line in `LOG.md` |

### Track order / Порядок треков
```
NOW (main):      GO ─────────────────────────────►  until every GO topic is at least 🟡
DAILY (parallel): DSA · EN · RU  (small piece every day, never stop)
LATER (main, one at a time, suggested order):
  PG → SD → OS → NET → DIST → OPS → SEC → QA → ARCH/API → AI → CAREER
```
**Why this order:** PostgreSQL is closest to your daily Go work; System Design uses PG and Go knowledge; OS → NET is the foundation for containers and DevOps; Distributed Systems connects SD and PG; then operate (OPS), secure (SEC), verify (QA); architecture, AI and career on top.
**RU:** PostgreSQL ближе всего к ежедневной работе с Go; System Design опирается на PG и Go; OS → NET — фундамент для контейнеров и DevOps; распределённые системы связывают SD и PG; дальше — эксплуатация (OPS), безопасность (SEC), проверка (QA); архитектура, ИИ и карьера — сверху.

### Teaching tip — your student is your test / Студент — твой тест
**EN:** The GO track is ordered like a course (beginner → advanced). After you mark a subtopic 🟡, **teach it to your student**. If you can explain it simply and answer their questions, it is really 🟡. If you get stuck, that is the gap to fix.
**RU:** Трек GO упорядочен как курс (от новичка к продвинутому). После отметки 🟡 **объясни подтему студенту**. Если можешь объяснить просто и ответить на вопросы — это действительно 🟡. Если застрял — это и есть пробел.

---

## Overview — all tracks

| Track | Name | Role now | Topics |
|---|---|---|---|
| **GO** | Go language & Go backend | 🔥 **MAIN** | 23 |
| **DSA** | Data Structures & Algorithms | daily, parallel | 20 |
| **EN** | English | daily, parallel | 8 |
| **RU** | Russian | daily, parallel | 9 |
| **PG** | PostgreSQL | later — 1st | 10 |
| **SD** | System Design | later — 2nd | 6 |
| **OS** | Operating Systems & Computer Architecture | later — 3rd | 7 |
| **NET** | Networking | later — 4th | 9 |
| **DIST** | Distributed Systems & Messaging | later — 5th | 8 |
| **OPS** | DevOps, Cloud & SRE | later — 6th | 9 |
| **SEC** | Security | later — 7th | 7 |
| **QA** | Testing & Quality Engineering | later — 8th | 6 |
| **ARCH** | Software Design, Architecture & API Design | later — 9th | 6 |
| **AI** | AI Engineering | later — 10th | 4 |
| **CAREER** | Senior skills, Interviews, Career | later — 11th | 5 |

---

# 🔥 TRACK GO — Go Language & Go Backend (MAIN)

**EN:** Goal for now: every subtopic at least 🟡 — you are a confident *user* of each part and can teach it. Topics 1–14 are the student course core; 15–23 are senior depth.
**RU:** Цель сейчас: каждая подтема минимум 🟡 — ты уверенно *пользуешься* каждой частью и можешь её преподавать. Темы 1–14 — ядро курса для студента; 15–23 — глубина уровня senior.
**Term dictionary:** `reference/go.md` · **Notes:** `notes/go/` · **Code:** `code/go/`

## GO 1 — Go philosophy & toolchain
- ⬜ **1.1 Why Go exists** — Terms: simplicity, composition over inheritance, fast compilation, static typing, garbage collection, CSP, Go proverbs, Go 1 compatibility promise, release cycle (Feb/Aug)
- ⬜ **1.2 Installing & environment** — Terms: `GOROOT`, `GOPATH`, `GOBIN`, `GOOS`, `GOARCH`, `GOTOOLCHAIN`, `go env`, `go version`
- ⬜ **1.3 The `go` command** — Terms: `go run`, `go build`, `go install`, `go test`, `go mod`, `go get`, `go fmt`, `go vet`, `go doc`, `go generate`, `go clean`, `go list`, `go tool`
- ⬜ **1.4 Program structure** — Terms: `package main`, `func main`, executable vs library package, `gofmt`, Go playground

## GO 2 — Syntax fundamentals
- ⬜ **2.1 Packages & imports** — Terms: exported vs unexported (capital letter), import alias, blank import `_`, dot import, `init()` and initialization order, package-level variables
- ⬜ **2.2 Variables & scope** — Terms: `var`, short declaration `:=`, zero value, shadowing, block scope, package scope, unused variable error, multiple assignment, blank identifier
- ⬜ **2.3 Constants & iota** — Terms: `const`, typed vs untyped constants, constant expressions, `iota`, enum pattern, overflow at compile time
- ⬜ **2.4 Basic types** — Terms: `int`/`int8..int64`, `uint`, `uintptr`, `float32`/`float64`, `complex128`, `bool`, `byte` (=uint8), `rune` (=int32), `string`, platform-dependent `int` size
- ⬜ **2.5 Type conversion & operators** — Terms: explicit conversion `T(x)`, no implicit conversion, integer overflow wraparound, integer division, `%`, bitwise `& | ^ &^ << >>`, operator precedence, `strconv` vs conversion
- ⬜ **2.6 Control flow** — Terms: `if` with init statement, `for` (3 forms), `range` (slice, map, string, channel, int, func), `switch` (no fallthrough by default, `fallthrough`, tagless switch), `break`/`continue` with labels, `goto`
- ⬜ **2.7 Functions** — Terms: multiple return values, named results, naked return, variadic `...`, functions as values, function types, anonymous functions, closures, captured variables, recursion, per-iteration loop variables (Go 1.22)
- ⬜ **2.8 Pointers** — Terms: `&` address-of, `*` dereference, `new(T)`, `new(expr)` (Go 1.26), nil pointer dereference panic, no pointer arithmetic, when to use pointers, pointer to struct auto-dereference

## GO 3 — Composite types
- ⬜ **3.1 Arrays** — Terms: fixed length is part of the type, value semantics (copy on assign), `[...]T{}`, comparing arrays, arrays as map keys
- ⬜ **3.2 Slices — usage** — Terms: slice literal, `make([]T, len, cap)`, `len`, `cap`, `append`, `copy`, slicing `s[a:b]`, full slice expression `s[a:b:c]`, nil slice vs empty slice, 2D slices, `slices` package (`Sort`, `Contains`, `Index`, `Clone`, `Compact`)
- ⬜ **3.3 Slices — internals** — Terms: slice header (pointer, len, cap), backing array, growth strategy, amortized O(1) append, sharing/aliasing trap, memory leak by reslicing, passing slices to functions
- ⬜ **3.4 Maps — usage** — Terms: map literal, `make(map[K]V, hint)`, comma-ok idiom, `delete`, `clear`, iteration order is random, nil map (read OK, write panics), comparable keys, map of structs is not addressable, `maps` package
- ⬜ **3.5 Maps — internals** — Terms: hash function, buckets, load factor, growth/evacuation, Swiss tables (Go 1.24: groups, control word, probing), why maps are not concurrency-safe, `concurrent map writes` fatal error
- ⬜ **3.6 Strings, bytes, runes** — Terms: immutable strings, string header, UTF-8, byte vs rune, `len` counts bytes, `range` over string yields runes, `[]byte(s)` copy cost, `strings.Builder`, `strings`/`bytes`/`strconv`/`unicode`/`utf8` packages, string concatenation cost
- ⬜ **3.7 Structs** — Terms: struct literal, field tags, anonymous struct, embedded struct (promotion), struct comparison, empty struct `struct{}` (zero size), field alignment & padding, copying structs, constructor function `NewT`

## GO 4 — Methods & interfaces
- ⬜ **4.1 Methods** — Terms: receiver, value receiver vs pointer receiver, method set, auto address/dereference on call, methods on non-struct types, method values & method expressions, consistency rule (don't mix receivers)
- ⬜ **4.2 Interfaces — basics** — Terms: interface type, implicit implementation (structural typing), `any` / `interface{}`, interface value = (type, value), dynamic type, compile-time check `var _ I = (*T)(nil)`
- ⬜ **4.3 Type assertion & type switch** — Terms: `x.(T)`, comma-ok assertion, panic on failed assertion, `switch v := x.(type)`
- ⬜ **4.4 The nil interface trap** — Terms: nil interface vs interface holding a nil pointer, returning typed nil error, why `err != nil` is true
- ⬜ **4.5 Key stdlib interfaces** — Terms: `fmt.Stringer`, `error`, `io.Reader`, `io.Writer`, `io.Closer`, `sort.Interface`, `http.Handler`, `json.Marshaler`, `context.Context`
- ⬜ **4.6 Interface design** — Terms: small interfaces, "accept interfaces, return structs", consumer-side interfaces, interface embedding, interface pollution, embedding for composition vs inheritance

## GO 5 — Errors, defer, panic
- ⬜ **5.1 Error basics** — Terms: `error` interface, `errors.New`, `fmt.Errorf`, errors are values, `if err != nil`, early return, error strings style (lowercase, no punctuation)
- ⬜ **5.2 Wrapping & inspection** — Terms: `%w`, error chain, `errors.Unwrap`, `errors.Is`, `errors.As`, `errors.AsType` (Go 1.26), `errors.Join`, multi-error
- ⬜ **5.3 Error types & strategies** — Terms: sentinel errors (`io.EOF`, `sql.ErrNoRows`), custom error types, opaque errors, behavior checking (`interface{ Timeout() bool }`), handle once (log OR return), adding context, mapping domain errors to HTTP codes
- ⬜ **5.4 defer** — Terms: LIFO order, arguments evaluated immediately, defer modifies named results, defer in loops (resource leak), `defer f.Close()` error handling, open-coded defers (cost)
- ⬜ **5.5 panic & recover** — Terms: `panic`, stack unwinding, `recover` (only in deferred func), runtime panics (nil map write, index out of range, nil deref), when panic is acceptable (programmer bug, init), recovery middleware, panics in goroutines crash the whole program

## GO 6 — Generics
- ⬜ **6.1 Type parameters** — Terms: type parameter list `[T any]`, generic function, generic type, instantiation, type inference
- ⬜ **6.2 Constraints** — Terms: constraint interface, `any`, `comparable`, `~T` (underlying type), union `int | float64`, `cmp.Ordered`, type sets
- ⬜ **6.3 Generic stdlib & patterns** — Terms: `slices`, `maps`, `cmp` packages, generic data structures (Stack[T], Set[T]), generic type aliases (Go 1.24)
- ⬜ **6.4 When NOT to use generics** — Terms: limitations (no method type parameters, no specialization), interface vs generic trade-off, readability cost, GC shape stenciling (performance)

## GO 7 — Packages, modules, project structure
- ⬜ **7.1 Modules** — Terms: `go.mod` (`module`, `go`, `toolchain`, `require`, `replace`, `exclude`, `retract`, `tool`), `go.sum`, `go mod tidy`, `go get pkg@version`, indirect dependencies
- ⬜ **7.2 Versioning** — Terms: semantic versioning, `v0`/`v1` stability, major version suffix `/v2`, Minimal Version Selection (MVS), pseudo-versions
- ⬜ **7.3 Module proxy & private modules** — Terms: `GOPROXY`, `GOSUMDB`, checksum database, `GOPRIVATE`, `GONOSUMDB`, vendoring `go mod vendor`
- ⬜ **7.4 Workspaces** — Terms: `go.work`, `go work use`, multi-module development
- ⬜ **7.5 Project layout** — Terms: `cmd/`, `internal/` (compiler-enforced), `pkg/` (controversial), package naming (short, no `utils`/`common`), avoiding import cycles, flat vs layered layout
- ⬜ **7.6 Documentation** — Terms: doc comments, `go doc`, pkg.go.dev, example functions, `README`, doc links `[pkg.Name]`

## GO 8 — Concurrency basics
- ⬜ **8.1 Concurrency vs parallelism** — Terms: concurrency (structure), parallelism (execution), CSP, "share memory by communicating"
- ⬜ **8.2 Goroutines** — Terms: `go` statement, goroutine vs OS thread, small growable stack, main goroutine exit kills all, `runtime.NumGoroutine`, `runtime.Gosched`
- ⬜ **8.3 WaitGroup** — Terms: `sync.WaitGroup`, `Add`/`Done`/`Wait`, `wg.Go` (Go 1.25), Add-before-go rule, passing WaitGroup by pointer
- ⬜ **8.4 Channels** — Terms: `make(chan T)`, unbuffered (synchronous handoff), buffered (queue), send `<-`, receive, `close`, receive from closed channel (zero value, ok=false), `range` over channel, directional types `chan<-`/`<-chan`, nil channel blocks forever
- ⬜ **8.5 Channel rules & panics** — Terms: send on closed channel panics, close of closed channel panics, close of nil channel panics, only the sender closes, "all goroutines are asleep — deadlock!"
- ⬜ **8.6 select** — Terms: `select`, random choice among ready cases, `default` (non-blocking), timeout with `time.After`, for-select loop, `ctx.Done()` case, disabling a case with nil channel
- ⬜ **8.7 Mutexes** — Terms: `sync.Mutex`, `sync.RWMutex`, critical section, `Lock`/`Unlock`, `defer mu.Unlock()`, don't copy a mutex (`go vet` copylocks), mutex vs channel — when to use which
- ⬜ **8.8 Once & atomics** — Terms: `sync.Once`, `sync.OnceFunc`/`OnceValue`/`OnceValues`, `sync/atomic`, `atomic.Int64`, `atomic.Value`, `atomic.Pointer[T]`, compare-and-swap (CAS)

## GO 9 — Concurrency patterns & pitfalls
- ⬜ **9.1 Worker pool** — Terms: jobs channel, results channel, fixed number of workers, bounded concurrency
- ⬜ **9.2 Pipeline & fan-out/fan-in** — Terms: stage, generator, fan-out, fan-in (merge), cancellation of pipelines
- ⬜ **9.3 errgroup & semaphore** — Terms: `golang.org/x/sync/errgroup`, `errgroup.WithContext`, `SetLimit`, semaphore via buffered channel, `x/sync/semaphore`
- ⬜ **9.4 Rate limiting & timing** — Terms: `time.Ticker`, `time.Timer`, `x/time/rate` (token bucket), debounce, throttling
- ⬜ **9.5 Goroutine leaks** — Terms: blocked sender/receiver, forgotten goroutine, missing cancellation, `goleak`, goroutine profile, detection with `runtime.NumGoroutine`
- ⬜ **9.6 Data races** — Terms: data race vs race condition, `-race` detector, happens-before, Go memory model, benign race myth, check-then-act bug
- ⬜ **9.7 More sync tools** — Terms: `sync.Cond`, `sync.Pool`, `sync.Map` (when it helps), `x/sync/singleflight`, deadlock, livelock, starvation, lock ordering
- ⬜ **9.8 Testing concurrent code** — Terms: `testing/synctest` (Go 1.25), fake clock, deterministic tests, `-count`, `-race` in CI

## GO 10 — context
- ⬜ **10.1 Why context** — Terms: cancellation propagation, request-scoped data, deadlines across API boundaries
- ⬜ **10.2 Creating contexts** — Terms: `context.Background`, `context.TODO`, `WithCancel`, `WithTimeout`, `WithDeadline`, `WithValue`, `WithCancelCause`, `WithoutCancel`, `AfterFunc`, `CancelFunc` (always call it)
- ⬜ **10.3 Using contexts** — Terms: `ctx.Done()`, `ctx.Err()` (`Canceled`, `DeadlineExceeded`), `context.Cause`, ctx as first parameter, don't store ctx in structs, context tree
- ⬜ **10.4 Context values** — Terms: unexported key type, request ID / user ID, what NOT to put in context (dependencies, optional params)
- ⬜ **10.5 Context in practice** — Terms: `r.Context()` in HTTP, `QueryContext` in DB, gRPC deadlines, `signal.NotifyContext`, timeouts budget

## GO 11 — Standard library essentials
- ⬜ **11.1 fmt** — Terms: verbs `%v %+v %#v %T %d %s %q %x %p`, `Sprintf`, `Fprintf`, `Errorf`, `Stringer` usage
- ⬜ **11.2 io & bufio** — Terms: `io.Reader`/`Writer` composition, `io.Copy`, `io.ReadAll`, `io.MultiWriter`, `io.TeeReader`, `io.Pipe`, `io.LimitReader`, `bufio.Scanner`, `bufio.Writer` + `Flush`, streaming vs loading into memory
- ⬜ **11.3 os & filesystem** — Terms: `os.Open`/`Create`/`ReadFile`/`WriteFile`, file permissions, `os.Getenv`/`LookupEnv`, `os.Args`, `os.Exit` (skips defers), `path/filepath`, `io/fs`, `os.Root` (Go 1.24), `embed`
- ⬜ **11.4 time** — Terms: `time.Time`, `time.Duration`, monotonic clock, `time.Now`/`Since`/`Until`, layout `2006-01-02T15:04:05Z07:00`, time zones & `time.Location`, UTC in DB, `Timer`/`Ticker` stop
- ⬜ **11.5 encoding/json** — Terms: `Marshal`/`Unmarshal`, struct tags, `omitempty`, `omitzero` (Go 1.24), `string` option, custom `MarshalJSON`, `json.RawMessage`, `Decoder` streaming, `DisallowUnknownFields`, numbers into `interface{}` → float64, `json/v2` (experiment)
- ⬜ **11.6 Logging** — Terms: `log`, `log/slog`, `slog.Logger`, levels, attributes, `JSONHandler`/`TextHandler`, `With`, context-aware logging
- ⬜ **11.7 Other essentials** — Terms: `strconv`, `strings`, `sort`/`slices`, `regexp` (RE2, no backtracking), `flag`, `math/rand/v2` vs `crypto/rand`, `os/exec`, `net/url`, `iter` + range-over-func iterators (Go 1.23), `unique` (interning), `weak` (Go 1.24)

## GO 12 — Testing in Go
- ⬜ **12.1 Basics** — Terms: `_test.go`, `TestXxx(t *testing.T)`, `t.Error` vs `t.Fatal`, `go test ./...`, `-run`, `-v`, `-count=1`, test cache
- ⬜ **12.2 Table-driven tests & subtests** — Terms: test table, `t.Run`, `t.Parallel`, `t.Helper`, `t.Cleanup`, `t.TempDir`, `t.Context` (Go 1.24), `t.Setenv`
- ⬜ **12.3 Test doubles** — Terms: interface-based mocking, fake, stub, mock, spy, `mockery`/`gomock`, why fakes often beat mocks
- ⬜ **12.4 HTTP & DB testing** — Terms: `httptest.NewRecorder`, `httptest.NewServer`, testcontainers-go, test database, transactions rollback per test, `TestMain`, build tags for integration tests
- ⬜ **12.5 Assertions & fixtures** — Terms: `testify` (`assert`, `require`), `go-cmp` `cmp.Diff`, golden files, `testdata/`
- ⬜ **12.6 Coverage, benchmarks, fuzzing, examples** — Terms: `-cover`, `-coverprofile`, `BenchmarkXxx`, `b.Loop` (Go 1.24), `-benchmem`, `FuzzXxx`, corpus, `ExampleXxx` with `// Output:`

## GO 13 — HTTP backend with net/http
- ⬜ **13.1 Server basics** — Terms: `http.Handler`, `http.HandlerFunc`, `ServeMux`, `ListenAndServe`, `ResponseWriter`, `*http.Request`
- ⬜ **13.2 Routing** — Terms: Go 1.22 patterns (`"GET /users/{id}"`), `r.PathValue`, wildcards `{path...}`, precedence, third-party routers (chi, gin, echo) and trade-offs
- ⬜ **13.3 Requests & responses** — Terms: reading body (`io.LimitReader`/`MaxBytesReader`), query params, headers, JSON encode/decode, status codes, `WriteHeader` order, content type
- ⬜ **13.4 Middleware** — Terms: middleware signature `func(http.Handler) http.Handler`, chaining, logging, recovery, request ID, auth, CORS, timeouts
- ⬜ **13.5 Server configuration** — Terms: `http.Server{}`, `ReadHeaderTimeout`, `ReadTimeout`, `WriteTimeout`, `IdleTimeout`, `MaxHeaderBytes`, why default server is dangerous
- ⬜ **13.6 Graceful shutdown** — Terms: `server.Shutdown(ctx)`, `signal.NotifyContext`, SIGTERM/SIGINT, draining connections, closing DB/workers in order
- ⬜ **13.7 HTTP client** — Terms: `http.Client` reuse, `Timeout`, `http.Transport`, connection pooling, keep-alive, `resp.Body.Close()` (and draining), retries with backoff, `NewRequestWithContext`
- ⬜ **13.8 Common backend features** — Terms: validation, config from env (12-factor), JWT auth middleware, file upload (multipart), pagination, rate limiting middleware, health/readiness endpoints, WebSocket, Server-Sent Events

## GO 14 — Databases from Go
- ⬜ **14.1 database/sql** — Terms: `sql.DB` is a pool, `sql.Open` vs `Ping`, `Query`/`QueryRow`/`Exec` + `Context` versions, `Scan`, `rows.Close`, `rows.Err`, `sql.ErrNoRows`, `sql.NullString`, placeholders `$1`
- ⬜ **14.2 Connection pool** — Terms: `SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`, `SetConnMaxIdleTime`, `db.Stats()`, pool exhaustion, connection leak by unclosed rows
- ⬜ **14.3 Transactions** — Terms: `BeginTx`, `tx.Commit`, `tx.Rollback` (defer), isolation level option, transaction across repositories (unit of work), passing `tx` vs `db`
- ⬜ **14.4 pgx** — Terms: `pgx` vs `lib/pq`, `pgxpool`, native types, `pgx.Batch`, `CopyFrom`, `pgconn` errors (`PgError` codes like `23505`)
- ⬜ **14.5 Query tools** — Terms: raw SQL, `sqlc` (SQL → Go code), `squirrel`/query builders, GORM (ORM trade-offs), `scany`
- ⬜ **14.6 Migrations** — Terms: `golang-migrate`, `goose`, `atlas`, up/down migrations, migration in CI/CD, zero-downtime migration
- ⬜ **14.7 Repository pattern & Redis** — Terms: repository interface, domain model vs DB model, SQL injection prevention, `go-redis` client, cache-aside in Go

## GO 15 — gRPC & Protocol Buffers
- ⬜ **15.1 Protobuf** — Terms: `.proto`, proto3, message, field numbers, scalar types, `repeated`, `oneof`, `optional`, well-known types (`Timestamp`), backward compatibility rules (never reuse field numbers, `reserved`)
- ⬜ **15.2 Tooling** — Terms: `protoc`, `buf`, `protoc-gen-go`, `protoc-gen-go-grpc`, generated code
- ⬜ **15.3 gRPC services** — Terms: service, unary RPC, server streaming, client streaming, bidirectional streaming, HTTP/2, metadata
- ⬜ **15.4 gRPC in production** — Terms: deadlines, status codes (`codes.NotFound`...), interceptors (unary/stream), retries, health checking, reflection, grpc-gateway, Connect (connectrpc)

## GO 16 — Messaging from Go
- ⬜ **16.1 Kafka clients** — Terms: `franz-go`, `segmentio/kafka-go`, `confluent-kafka-go`, producer (acks, idempotent), consumer group, partition, offset commit (auto vs manual)
- ⬜ **16.2 Reliable consumers** — Terms: at-least-once processing, idempotent consumer, retry topic, dead-letter queue, graceful consumer shutdown, ordering per key
- ⬜ **16.3 Outbox in Go** — Terms: transactional outbox table, relay/poller, `SELECT ... FOR UPDATE SKIP LOCKED`
- ⬜ **16.4 Other brokers** — Terms: RabbitMQ (`amqp091-go`, ack/nack, prefetch), NATS / JetStream

## GO 17 — Application design in Go
- ⬜ **17.1 Layering** — Terms: handler (transport) → service (business logic) → repository (data), DTO vs domain model, dependency direction
- ⬜ **17.2 Dependency injection** — Terms: constructor injection, `main` as composition root, no global state, `wire`, `fx` (and why manual DI is usually enough)
- ⬜ **17.3 Clean / hexagonal in Go** — Terms: ports & adapters, domain package, interfaces at the consumer, over-engineering risk
- ⬜ **17.4 Go design patterns** — Terms: functional options, constructor `NewX`, factory, builder, strategy (via interfaces/funcs), decorator/middleware, adapter, singleton via `sync.Once`, observer via channels
- ⬜ **17.5 Configuration** — Terms: env vars, config struct, `envconfig`/`viper`/`koanf`, validation at startup, secrets in config, 12-factor app

## GO 18 — Observability in Go
- ⬜ **18.1 Structured logging** — Terms: `slog`, log levels, request-scoped logger, correlation ID / trace ID in logs, don't log secrets
- ⬜ **18.2 Metrics** — Terms: `prometheus/client_golang`, counter, gauge, histogram, labels, cardinality, `/metrics` endpoint, RED metrics middleware
- ⬜ **18.3 Tracing** — Terms: OpenTelemetry Go SDK, tracer, span, attributes, context propagation, `otelhttp`, `otelgrpc`, exporter (OTLP)
- ⬜ **18.4 Runtime diagnostics** — Terms: `net/http/pprof`, `expvar`, `runtime/metrics`, health checks (liveness/readiness), flight recorder (Go 1.25)

## GO 19 — Tooling & code quality
- ⬜ **19.1 Formatting & linting** — Terms: `gofmt`, `goimports`, `gofumpt`, `go vet`, `staticcheck`, `golangci-lint` (config, linters), `govulncheck`
- ⬜ **19.2 Debugging** — Terms: Delve (`dlv debug`, breakpoints, goroutines view), debugging in VS Code, print debugging, core dumps
- ⬜ **19.3 Build** — Terms: `-ldflags "-X main.version=..."`, `-trimpath`, `CGO_ENABLED=0`, static binary, cross-compilation `GOOS=linux GOARCH=arm64`, build tags `//go:build`
- ⬜ **19.4 Code generation** — Terms: `go generate`, `//go:generate`, `stringer`, `mockery`, `sqlc`, `text/template`, `go/ast`
- ⬜ **19.5 Project automation** — Terms: Makefile / Taskfile, `air` hot reload, pre-commit hooks, Dockerfile multi-stage for Go, distroless/scratch images

## GO 20 — Runtime internals (senior)
- ⬜ **20.1 Goroutine internals** — Terms: `g` struct, initial stack 2 KB, stack growth by copying, goroutine states (runnable, running, waiting)
- ⬜ **20.2 GMP scheduler** — Terms: G (goroutine), M (OS thread), P (processor/context), local run queue, global run queue, work stealing, hand-off on syscall, `sysmon`, asynchronous preemption (Go 1.14), `GOMAXPROCS` (container-aware in Go 1.25), netpoller (epoll/kqueue)
- ⬜ **20.3 Memory allocator** — Terms: `mcache`, `mcentral`, `mheap`, spans, size classes, tiny allocator, stack vs heap allocation
- ⬜ **20.4 Escape analysis** — Terms: escape to heap, `go build -gcflags=-m`, why returning pointers can allocate, interfaces cause escapes, inlining
- ⬜ **20.5 Garbage collector** — Terms: tri-color mark & sweep, concurrent GC, write barrier, stop-the-world phases, GC pacer, `GOGC`, `GOMEMLIMIT`, `GODEBUG=gctrace=1`, Green Tea GC (Go 1.25 experiment → default in 1.26), finalizers vs `runtime.AddCleanup`
- ⬜ **20.6 Channel internals** — Terms: `hchan`, circular buffer, `sendq`/`recvq`, `sudog`, lock inside channel, direct send to waiting receiver
- ⬜ **20.7 Interface & map internals** — Terms: `iface`, `eface`, `itab`, method dispatch, devirtualization, Swiss table groups/control bytes (Go 1.24)
- ⬜ **20.8 reflect & unsafe** — Terms: `reflect.Type`, `reflect.Value`, `Kind`, struct tags via reflection, reflection cost, `unsafe.Pointer`, `unsafe.Sizeof`/`Alignof`/`Offsetof`, `unsafe.String`/`unsafe.Slice`, cgo call cost

## GO 21 — Performance & profiling (senior)
- ⬜ **21.1 Benchmarking** — Terms: `go test -bench`, `-benchmem`, allocs/op, `benchstat`, avoiding compiler optimizations in benchmarks, `b.Loop`
- ⬜ **21.2 Profiling** — Terms: `pprof` CPU/heap/allocs/goroutine/block/mutex profiles, `go tool pprof`, flame graph, `top`, `list`, continuous profiling
- ⬜ **21.3 Execution tracing** — Terms: `go tool trace`, scheduler latency, GC events, goroutine analysis
- ⬜ **21.4 Optimization techniques** — Terms: preallocation, avoid `[]byte`↔`string` copies, `strings.Builder`, `sync.Pool`, reduce pointer-heavy structures, struct field ordering, bounds check elimination, inlining budget, false sharing
- ⬜ **21.5 PGO & GC tuning** — Terms: profile-guided optimization (`default.pgo`), `GOGC` vs `GOMEMLIMIT` trade-off, memory limit in containers

## GO 22 — Idioms & common mistakes
- ⬜ **22.1 Idiomatic Go** — Terms: Effective Go, Go Code Review Comments, Go Proverbs, naming (short names, `MixedCaps`, receivers), "clear is better than clever", "a little copying is better than a little dependency"
- ⬜ **22.2 Common mistakes** — Terms: (from "100 Go Mistakes") variable shadowing, unnecessary nesting, `init` misuse, getters/setters, interface pollution, slice/map traps, `defer` in loops, ignoring errors, goroutine with loop variable (pre-1.22), copying sync types, `time.After` leak in loops, unclosed response bodies
- ⬜ **22.3 Reading real Go code** — Terms: reading stdlib source (`net/http`, `sync`), reading popular projects (Prometheus, CockroachDB, Kubernetes), code review checklist

## GO 23 — Go ecosystem awareness
- ⬜ **23.1 CLI tools in Go** — Terms: `flag`, `cobra`, `urfave/cli`, exit codes, stdin/stdout piping
- ⬜ **23.2 Notable libraries** — Terms: chi, gin, echo, fiber (fasthttp caveat), pgx, sqlc, go-redis, zap/zerolog (pre-slog), testify, viper, cobra, OpenTelemetry, ent
- ⬜ **23.3 Go in the industry** — Terms: Docker, Kubernetes, Terraform, Prometheus, etcd written in Go — why Go fits cloud infrastructure; WebAssembly (`GOOS=js`, `wasip1`), TinyGo

---

# 📅 TRACK DSA — Data Structures & Algorithms (daily, parallel)

**EN:** Daily small piece: one concept, one term, or one problem. Focus on **patterns**, not memorized solutions. Interviews still test this.
**RU:** Каждый день маленький кусок: одна концепция, один термин или одна задача. Учим **паттерны**, а не заученные решения.
**Notes:** `notes/dsa/` · **Tasks:** `tasks/dsa/` · **Code:** `code/ds/` (skeletons) + `code/ds/ds_imp/` (your implementations) · **Dictionary:** `reference/dsa-*.md`

## DSA 1 — Complexity analysis
- 🟡 **1.1 Big-O basics** — Terms: time complexity, space complexity, O(1), O(log n), O(n), O(n log n), O(n²), O(2ⁿ), dropping constants
- ⬜ **1.2 Advanced analysis** — Terms: amortized analysis, best/average/worst case, recursion tree, master theorem (awareness), auxiliary space

## DSA 2 — Arrays & strings
- 🟡 **2.1 Array / slice operations** — Terms: random access, insertion/deletion cost, dynamic array, in-place
- ⬜ **2.2 Problems** — [ ] reverse array in place · [ ] Kadane's max subarray · [ ] move zeros · [ ] find duplicates · [ ] rotate array · [ ] product of array except self
- ⬜ **2.3 Strings** — Terms: immutability, character counting, anagram, palindrome, string builder

## DSA 3 — Hashing
- 🟡 **3.1 Hash map & hash set** — Terms: hash function, bucket, collision, chaining, open addressing, load factor, rehashing
- ⬜ **3.2 Problems** — [ ] two sum · [ ] group anagrams · [ ] first non-repeating char · [ ] word frequency · [ ] valid anagram · [ ] longest consecutive sequence

## DSA 4 — Linked lists
- 🟡 **4.1 Concepts** — Terms: node, head, tail, singly vs doubly linked, sentinel/dummy node, cache unfriendliness
- 🟡 **4.2 Problems** — [x] reverse list · [x] middle node · [x] cycle (hash set) · [x] cycle (Floyd) · [x] nth from end · [x] merge two sorted · [ ] remove duplicates · [ ] palindrome list

## DSA 5 — Stacks & queues
- 🟡 **5.1 Concepts** — Terms: LIFO, FIFO, push/pop, enqueue/dequeue, deque, circular buffer, monotonic stack
- 🟡 **5.2 Problems** — [x] valid parentheses · [x] min stack · [x] daily temperatures · [ ] queue using two stacks · [ ] evaluate reverse Polish notation · [ ] next greater element

## DSA 6 — Trees
- 🟡 **6.1 Binary tree & traversals** — Terms: root, leaf, height, depth, inorder, preorder, postorder, level order (BFS)
- ⬜ **6.2 BST** — Terms: BST property, search/insert/delete, balanced vs unbalanced, AVL / red-black (awareness)
- ⬜ **6.3 Problems** — [x] all traversals · [ ] max depth · [ ] balanced tree · [ ] validate BST · [ ] lowest common ancestor · [ ] diameter · [ ] invert tree · [ ] serialize/deserialize (C2)

## DSA 7 — Heaps
- 🟡 **7.1 Concepts** — Terms: min-heap, max-heap, complete binary tree, array representation, sift up / sift down, heapify O(n)
- ⬜ **7.2 Go `container/heap`** — Terms: `heap.Interface`, `Push`, `Pop`, `Init`, `Fix`
- ⬜ **7.3 Problems** — [ ] kth largest · [ ] top K frequent · [ ] merge K sorted lists · [ ] median from data stream

## DSA 8 — Graphs
- 🟡 **8.1 Concepts** — Terms: vertex, edge, directed/undirected, weighted, adjacency list vs matrix, visited set, BFS, DFS
- ⬜ **8.2 Topological sort & cycles** — Terms: DAG, in-degree, Kahn's algorithm, DFS coloring
- ⬜ **8.3 Shortest paths** — Terms: BFS on unweighted graph, Dijkstra, Bellman-Ford (awareness), 0-1 BFS
- ⬜ **8.4 Problems** — [ ] number of islands · [ ] clone graph · [ ] course schedule · [ ] BFS shortest path · [ ] topological sort · [ ] rotting oranges

## DSA 9 — Binary search
- 🟡 **9.1 Concepts** — Terms: sorted input, `lo`/`hi`, `<` vs `<=`, overflow-safe mid, lower bound / upper bound, search on answer space
- 🟡 **9.2 Problems** — [x] classic · [x] search insert position · [ ] rotated sorted array · [ ] first & last position · [ ] min in rotated array · [ ] Koko eating bananas

## DSA 10 — Two pointers
- 🟡 **10.1 Concepts** — Terms: opposite-direction pointers, same-direction (fast/slow), in-place partitioning
- ⬜ **10.2 Problems** — [ ] two sum sorted · [ ] valid palindrome · [ ] container with most water · [ ] 3sum · [ ] remove duplicates in place · [ ] trapping rain water

## DSA 11 — Sliding window
- ⬜ **11.1 Concepts** — Terms: fixed window, variable window, expand/shrink, window state (map/counter)
- ⬜ **11.2 Problems** — [ ] max sum subarray size k · [ ] longest substring without repeating · [ ] minimum window substring · [ ] longest subarray sum ≤ k

## DSA 12 — Prefix sums & intervals
- ⬜ **12.1 Prefix sums** — Terms: prefix array, range sum query, subarray sum equals k (prefix + hash map)
- ⬜ **12.2 Intervals** — Terms: sort by start, merge intervals, insert interval, meeting rooms (heap)

## DSA 13 — Recursion & backtracking
- ⬜ **13.1 Recursion** — Terms: base case, recursive case, call stack, stack overflow, tail recursion (Go does not optimize)
- ⬜ **13.2 Backtracking** — Terms: choose → explore → unchoose, pruning, decision tree — problems: [ ] subsets · [ ] permutations · [ ] combination sum · [ ] N-queens (C2)

## DSA 14 — Sorting
- ⬜ **14.1 Algorithms** — Terms: bubble/insertion (why slow), merge sort, quicksort, pivot, heap sort, counting sort, stability, in-place, Go's `slices.Sort` (pdqsort)

## DSA 15 — Dynamic programming
- ⬜ **15.1 Concepts** — Terms: overlapping subproblems, optimal substructure, memoization (top-down), tabulation (bottom-up), state, transition
- ⬜ **15.2 Problems** — [ ] fibonacci · [ ] climbing stairs · [ ] house robber · [ ] coin change · [ ] longest common subsequence · [ ] 0/1 knapsack (C2)

## DSA 16 — Greedy
- ⬜ **16.1 Concepts** — Terms: greedy choice, local vs global optimum, exchange argument — problems: [ ] jump game · [ ] gas station

## DSA 17 — Advanced structures
- ⬜ **17.1 Trie** — Terms: prefix tree, insert/search/startsWith, autocomplete
- ⬜ **17.2 Union-find** — Terms: disjoint set, find, union, path compression, union by rank
- ⬜ **17.3 LRU cache** — Terms: hash map + doubly linked list, O(1) get/put (classic interview task)

## DSA 18 — Bit manipulation
- ⬜ **18.1 Basics** — Terms: AND/OR/XOR/NOT, shifts, masks, `x & (x-1)`, single number (XOR trick), counting bits

## DSA 19 — Structures inside real systems (connects to PG/DIST)
- ⬜ **19.1 Storage structures** — Terms: B-tree, B+tree, LSM-tree, SSTable, skip list
- ⬜ **19.2 Probabilistic structures** — Terms: Bloom filter, HyperLogLog, count-min sketch
- ⬜ **19.3 Distributed structures** — Terms: consistent hashing ring, Merkle tree, ring buffer

## DSA 20 — Interview problem-solving method
- ⬜ **20.1 Method** — Terms: clarify input/output, examples & edge cases, brute force first, optimize, dry run, complexity statement, thinking out loud

---

# 📅 TRACK EN — English (daily, parallel)

**EN:** Corrections happen in every message automatically. Daily task: learn **one** term / phrase / grammar point and use it in your own sentence.
**RU:** Исправления — в каждом сообщении автоматически. Ежедневно: **один** термин / фразу / грамматическое правило и своё предложение с ним.
**Files:** `english/vocabulary.md` (corrections, terms) · `english/PLAN.md` (method, phrases)

## EN 1 — Grammar (your weak points first)
- ⬜ **1.1 Articles** — Terms: a / an / the / zero article, countable vs uncountable, "a roadmap" vs "the roadmap"
- ⬜ **1.2 Prepositions** — Terms: depend on, participate in, work on, write in (a language), responsible for, interested in, at/in/on for time and place
- ⬜ **1.3 Tenses for work talk** — Terms: past simple (finished projects), present perfect (experience: "I've worked with"), present continuous (current work), future (plans)
- ⬜ **1.4 Subject–verb agreement** — Terms: "skills don't", "the data is/are", "each of them is"
- ⬜ **1.5 Conditionals** — Terms: zero, first, second ("if I had more time, I would…"), mixed — needed for trade-off talk
- ⬜ **1.6 Passive voice** — Terms: "the request is handled by…", "the data was migrated"
- ⬜ **1.7 Modals** — Terms: should, could, might, must, would — for suggestions and uncertainty
- ⬜ **1.8 Relative clauses & linking words** — Terms: which/that/who, however, therefore, whereas, as a result, on the other hand

## EN 2 — Technical vocabulary (grows with every track)
- ⬜ **2.1 Go & programming verbs** — Terms: spawn, block, leak, propagate, wrap, handle, return, allocate, implement, embed
- ⬜ **2.2 Backend & system design terms** — Terms: throughput, latency, bottleneck, scale out, fault tolerance, single point of failure, trade-off, idempotent
- ⬜ **2.3 DevOps & operations terms** — Terms: deploy, roll back, roll out, provision, on-call, outage, incident, postmortem
- ⬜ **2.4 Pronunciation of tech words** — Terms: cache /kæʃ/, queue /kjuː/, width, data, schema, Kubernetes, nginx, SQL ("sequel"), Linux

## EN 3 — Speaking
- ⬜ **3.1 Self-introduction** — 60-second and 2-minute versions
- ⬜ **3.2 Project stories** — Iman (tests 87%, query optimization), PPP (Excel/PDF generation), IMV Archive (solo, Kotlin) — `interviews/my-projects.md`
- ⬜ **3.3 Explaining technical concepts** — "What is X, why it exists, how it works, trade-offs" structure
- ⬜ **3.4 Teaching phrases** — Terms: "let's break this down", "think of it as…", "does that make sense?", "here's an analogy" (useful for your student!)

## EN 4 — Interview phrases
- ⬜ **4.1 Thinking, clarifying, not knowing, trade-offs, disagreeing** — see `english/PLAN.md` table
- ⬜ **4.2 Behavioral (STAR)** — Terms: situation, task, action, result; "I took ownership of…", "the outcome was…"

## EN 5 — Writing
- ⬜ **5.1 Commit messages & PR descriptions** — Terms: imperative mood ("add", not "added"), summary + why
- ⬜ **5.2 Code review comments** — Terms: "nit:", "consider…", "what do you think about…", "LGTM"
- ⬜ **5.3 Design docs & technical emails** — Terms: context, goals, non-goals, proposal, alternatives considered
- ⬜ **5.4 Slack / chat messages** — Terms: concise asks, "heads-up", "FYI", "blocked on…"

## EN 6 — Listening
- ⬜ **6.1 Podcasts & talks** — GopherCon talks, Go podcasts; shadowing method (see `english/PLAN.md`)

## EN 7 — Reading
- ⬜ **7.1 Engineering blogs** — Go Blog, Cloudflare, Stripe, Uber, Discord engineering

## EN 8 — Fluency habits
- ⬜ **8.1 Thinking in English** — narrate your coding in English; write daily log in English

---

# 📅 TRACK RU — Russian (daily, parallel)

**EN:** Polish, not build from scratch. Focus on your weak points: cases, verb aspect, gender agreement. Daily task: one word / phrase / grammar point.
**RU:** Шлифовка, а не изучение с нуля. Фокус на слабых местах: падежи, вид глагола, согласование рода. Ежедневно: одно слово / фраза / правило.
**Files:** `russian/ROADMAP.md` (detailed grammar plan) · `russian/grammar.md` · `russian/vocabulary.md` · `russian/interview_phrases.md`

## RU 1 — Cases (Падежи) 🔥 weak point
- ⬜ **1.1 Nominative & Accusative** — Terms: кто? что? / кого? что?, animate vs inanimate accusative
- ⬜ **1.2 Genitive** — Terms: кого? чего?, possession (код сервера), negation (нет данных), after numbers (5 серверов)
- ⬜ **1.3 Dative** — Terms: кому? чему?, "мне нужно", "по + dative" (по документации)
- ⬜ **1.4 Instrumental** — Terms: кем? чем?, "работать с базой данных", "быть разработчиком"
- ⬜ **1.5 Prepositional** — Terms: о ком? о чём? где?, в/на + prepositional (в проекте, на сервере)

## RU 2 — Verb aspect (Вид глагола) 🔥 weak point
- ⬜ **2.1 Imperfective vs perfective** — Terms: делать/сделать, писать/написать, process vs result, repeated vs one-time action
- ⬜ **2.2 Aspect in past & future** — Terms: "я делал" (process) vs "я сделал" (result), "буду делать" vs "сделаю"

## RU 3 — Gender & agreement 🔥 weak point
- ⬜ **3.1 Noun gender** — Terms: masculine/feminine/neuter endings, tricky nouns (данные — plural, кофе — masc.)
- ⬜ **3.2 Agreement** — Terms: adjective + noun, past tense verb + subject (запрос пришёл / задача пришла)

## RU 4 — Prepositions
- ⬜ **4.1 в / на / из / с / к / по / о / для / через** — with correct cases

## RU 5 — Verbs of motion
- ⬜ **5.1 идти/ходить, ехать/ездить, prefixed verbs** — Terms: unidirectional vs multidirectional, при-/у-/вы-/пере-

## RU 6 — Numbers with nouns
- ⬜ **6.1 Number + noun agreement** — Terms: 1 сервер, 2 сервера, 5 серверов; ordinal numbers; dates & time

## RU 7 — Technical vocabulary
- ⬜ **7.1 IT terms in Russian** — Terms: развёртывание (deployment), нагрузка (load), отказоустойчивость (fault tolerance), масштабирование (scaling), задержка (latency), пропускная способность (throughput), утечка памяти (memory leak), состояние гонки (race condition)

## RU 8 — Professional communication
- ⬜ **8.1 Meetings & work chat** — Terms: "предлагаю…", "давайте обсудим…", "я не согласен, потому что…", formal ты/вы register

## RU 9 — Interviews in Russian
- ⬜ **9.1 Self-introduction & projects** — `russian/interview_phrases.md`
- ⬜ **9.2 Explaining technical answers** — senior-level answers in Russian

---

# ⏭ TRACK PG — PostgreSQL (later — 1st)

**EN:** Data outlives code. Your strongest card next to Go — make it expert-level.
**RU:** Данные живут дольше кода. Твоя сильнейшая карта после Go — доведи до экспертного уровня.
**Dictionary:** `reference/databases.md` · **Notes:** `notes/postgres/`

## PG 1 — SQL fluency
- ⬜ **1.1 Querying** — Terms: SELECT, WHERE, ORDER BY, LIMIT, DISTINCT, DISTINCT ON, NULL semantics (three-valued logic), COALESCE
- ⬜ **1.2 Joins** — Terms: INNER, LEFT, RIGHT, FULL, CROSS, self-join, semi-join (EXISTS), anti-join (NOT EXISTS), LATERAL
- ⬜ **1.3 Aggregation & windows** — Terms: GROUP BY, HAVING, aggregate functions, FILTER, window functions (ROW_NUMBER, RANK, LAG/LEAD, running totals), PARTITION BY
- ⬜ **1.4 Advanced SQL** — Terms: CTE (WITH), recursive CTE, subqueries (correlated), UNION/INTERSECT/EXCEPT, UPSERT (`ON CONFLICT`), RETURNING, CASE

## PG 2 — Data modeling
- ⬜ **2.1 Normalization** — Terms: 1NF, 2NF, 3NF, BCNF, denormalization, redundancy, update anomalies
- ⬜ **2.2 Keys & constraints** — Terms: primary key, natural vs surrogate key, UUID v4 vs v7 vs bigserial/identity, foreign key, ON DELETE CASCADE, UNIQUE, CHECK, NOT NULL, exclusion constraint
- ⬜ **2.3 Data types** — Terms: numeric vs float (money!), text vs varchar, timestamptz vs timestamp, jsonb, arrays, enum vs lookup table, bytea
- ⬜ **2.4 Modeling patterns** — Terms: one-to-many, many-to-many (junction table), soft delete, audit columns, polymorphic associations, multi-tenancy (tenant_id, schema-per-tenant), EAV anti-pattern

## PG 3 — Transactions & concurrency
- ⬜ **3.1 ACID** — Terms: atomicity, consistency, isolation, durability
- ⬜ **3.2 Isolation levels & anomalies** — Terms: Read Committed (default), Repeatable Read, Serializable (SSI), dirty read, non-repeatable read, phantom read, lost update, write skew, serialization failure + retry
- ⬜ **3.3 Locking** — Terms: row locks, `SELECT FOR UPDATE / FOR SHARE / SKIP LOCKED / NOWAIT`, table lock modes, ACCESS EXCLUSIVE, deadlock detection, advisory locks, lock queue
- ⬜ **3.4 Concurrency patterns** — Terms: optimistic locking (version column), pessimistic locking, Postgres as job queue (`SKIP LOCKED`), counters under contention

## PG 4 — Indexes
- ⬜ **4.1 B-tree index** — Terms: index structure, composite index (column order, leftmost prefix), unique index, sort order
- ⬜ **4.2 Special indexes** — Terms: partial index, covering index (`INCLUDE`), expression index, GIN (jsonb, arrays, full-text), GiST, BRIN, hash index
- ⬜ **4.3 Index costs** — Terms: write amplification, index bloat, unused indexes, `CREATE INDEX CONCURRENTLY`, `REINDEX`

## PG 5 — Query performance
- ⬜ **5.1 EXPLAIN** — Terms: `EXPLAIN (ANALYZE, BUFFERS)`, cost, rows estimate vs actual, Seq Scan, Index Scan, Index Only Scan, Bitmap Heap Scan, Nested Loop, Hash Join, Merge Join, Sort, work_mem spill
- ⬜ **5.2 Planner & statistics** — Terms: `ANALYZE`, `pg_statistic`, cardinality estimation, extended statistics, planner misestimates
- ⬜ **5.3 Common problems** — Terms: N+1 queries, OFFSET pagination vs keyset pagination, `SELECT *`, functions on indexed columns, implicit casts, OR conditions
- ⬜ **5.4 Monitoring** — Terms: `pg_stat_statements`, `pg_stat_activity`, `auto_explain`, slow query log, `pg_stat_user_tables`, `pg_locks`

## PG 6 — Internals
- ⬜ **6.1 MVCC** — Terms: tuple versions, xmin/xmax, transaction ID, snapshot, visibility rules, dead tuples
- ⬜ **6.2 VACUUM** — Terms: VACUUM, autovacuum, VACUUM FULL, bloat, visibility map, freezing, transaction ID wraparound, HOT updates, fillfactor
- ⬜ **6.3 Storage & WAL** — Terms: pages (8 KB), heap, TOAST, shared buffers, WAL, checkpoint, `fsync`, `synchronous_commit`
- ⬜ **6.4 Process architecture** — Terms: postmaster, backend process per connection, background workers, why connections are expensive

## PG 7 — Operations
- ⬜ **7.1 Connection pooling** — Terms: PgBouncer, session vs transaction pooling, prepared statements with poolers, `max_connections`
- ⬜ **7.2 Replication & HA** — Terms: streaming (physical) replication, logical replication, sync vs async replica, replication lag, read replicas, failover, Patroni
- ⬜ **7.3 Backup & recovery** — Terms: `pg_dump`/`pg_restore`, physical backup, `pg_basebackup`, PITR, WAL archiving, pgBackRest / WAL-G, RPO/RTO, restore testing
- ⬜ **7.4 Configuration** — Terms: `shared_buffers`, `work_mem`, `maintenance_work_mem`, `effective_cache_size`, `random_page_cost`, statement_timeout, idle_in_transaction_session_timeout

## PG 8 — Schema evolution
- ⬜ **8.1 Zero-downtime migrations** — Terms: expand-and-contract, adding NOT NULL safely, `CONCURRENTLY`, backfill in batches, lock_timeout, renaming columns safely

## PG 9 — Scaling Postgres
- ⬜ **9.1 Partitioning** — Terms: declarative partitioning (range, list, hash), partition pruning, detach/attach, when partitioning hurts
- ⬜ **9.2 Sharding & distributed Postgres** — Terms: application-level sharding, Citus, CockroachDB / YugabyteDB (awareness)

## PG 10 — Beyond Postgres (when to choose what)
- ⬜ **10.1 Redis** — Terms: strings, hashes, lists, sets, sorted sets, streams, TTL, eviction policies, RDB/AOF, pub/sub, Redis Cluster, Sentinel, distributed lock caveats
- ⬜ **10.2 Other data stores** — Terms: OLTP vs OLAP, ClickHouse (columnar), Elasticsearch/OpenSearch (inverted index), MongoDB (document), Cassandra (wide-column), DynamoDB, S3 (object storage), pgvector (vector search)
- ⬜ **10.3 Security in Postgres** — Terms: roles, GRANT/REVOKE, least privilege, row-level security (RLS), `pg_hba.conf`, SSL connections

---

# ⏭ TRACK SD — System Design (later — 2nd)

**EN:** Turning requirements into an architecture with explicit trade-offs. Decides the senior level in interviews.
**RU:** Превращение требований в архитектуру с явными компромиссами. Определяет уровень senior на собеседованиях.
**Dictionary:** `reference/system-design.md` · **Notes:** `notes/system-design/`

## SD 1 — Framework
- ⬜ **1.1 Interview framework** — Terms: functional vs non-functional requirements, clarifying questions, high-level design, deep dive, bottlenecks, wrap-up
- ⬜ **1.2 Estimation** — Terms: back-of-the-envelope, QPS, peak vs average, storage per day/year, bandwidth, latency numbers every programmer should know, powers of two

## SD 2 — Building blocks (C1 done in earlier sessions → now C2)
- 🟡 **2.1 Scaling** — Terms: vertical vs horizontal scaling, stateless services, sticky sessions
- 🟡 **2.2 Load balancing** — Terms: L4 vs L7, round robin, least connections, consistent hashing, health checks
- 🟡 **2.3 Caching** — Terms: cache-aside, write-through, write-back, write-around, TTL, LRU/LFU, cache invalidation, cache stampede (C2)
- 🟡 **2.4 CDN** — Terms: edge, origin, push vs pull CDN, cache hit ratio
- 🟡 **2.5 Databases in design** — Terms: SQL vs NoSQL, replication, read replicas, sharding, CAP theorem
- 🟡 **2.6 Message queues** — Terms: async processing, producer/consumer, at-most/at-least/exactly-once, idempotency
- 🟡 **2.7 Rate limiting** — Terms: token bucket, leaky bucket, fixed window, sliding window log/counter
- ⬜ **2.8 More blocks** — Terms: API gateway, reverse proxy, service discovery, object storage, search index, unique ID generation (UUIDv7, Snowflake), blob storage, websocket gateway

## SD 3 — Non-functional requirements
- ⬜ **3.1 Availability & reliability** — Terms: SLA/SLO/SLI, nines (99.9 vs 99.99), redundancy, failover, single point of failure
- ⬜ **3.2 Performance** — Terms: latency percentiles (p50/p95/p99), throughput, tail latency
- ⬜ **3.3 Consistency & durability** — Terms: strong vs eventual consistency, read-your-writes, data loss tolerance
- ⬜ **3.4 Cost & operability** — Terms: cost estimation, operational complexity, observability in design

## SD 4 — Classic designs (practice one per week once track starts)
- ⬜ **4.1** URL shortener · ⬜ **4.2** Rate limiter · ⬜ **4.3** Key-value store · ⬜ **4.4** Notification service · ⬜ **4.5** Chat (WhatsApp) · ⬜ **4.6** News feed · ⬜ **4.7** Web crawler · ⬜ **4.8** Payment system (ledger, idempotency, reconciliation) · ⬜ **4.9** Ride sharing / geo (geohash, quadtree) · ⬜ **4.10** Video streaming · ⬜ **4.11** File storage & sync · ⬜ **4.12** Search autocomplete · ⬜ **4.13** Distributed job scheduler · ⬜ **4.14** Ticket booking · ⬜ **4.15** Metrics/logging platform · ⬜ **4.16** RAG chat over documents

## SD 5 — Real-world design skills
- ⬜ **5.1 Design docs** — Terms: RFC, context, goals/non-goals, alternatives considered, ADR
- ⬜ **5.2 Multi-region** — Terms: active-passive, active-active, data residency, geo-routing

## SD 6 — Learning from real systems
- ⬜ **6.1 Engineering blogs & postmortems** — Uber, Discord, Stripe, Cloudflare, Figma; famous outages (AWS us-east-1, Cloudflare, GitHub)

---

# ⏭ TRACK OS — Operating Systems & Computer Architecture (later — 3rd)

**EN:** Go runtime, Postgres, Docker and Kubernetes are built on OS primitives. Without this layer they stay "magic."
**RU:** Рантайм Go, Postgres, Docker и Kubernetes построены на примитивах ОС. Без этого слоя они остаются «магией».
**Dictionary:** `reference/os-concurrency.md`

## OS 1 — Processes & threads
- ⬜ **1.1 Processes** — Terms: process, PID, fork/exec, process states, zombie, orphan, parent/child, exit code
- ⬜ **1.2 Threads** — Terms: kernel thread, user thread, thread vs process, context switch cost, green threads (goroutines)
- ⬜ **1.3 Scheduling** — Terms: preemptive scheduling, time slice, CFS / EEVDF, priority, nice, CPU-bound vs I/O-bound

## OS 2 — Memory
- ⬜ **2.1 Virtual memory** — Terms: address space, pages, page table, TLB, page fault, swap, mmap
- ⬜ **2.2 Process memory layout** — Terms: text, data, heap, stack, stack overflow, memory-mapped files
- ⬜ **2.3 Memory limits** — Terms: RSS, VSZ, OOM killer, overcommit, cgroup memory limit

## OS 3 — Kernel interface
- ⬜ **3.1 System calls** — Terms: user space vs kernel space, syscall, `strace`, cost of syscalls
- ⬜ **3.2 Files & descriptors** — Terms: file descriptor, everything is a file, open file limit (`ulimit -n`), pipes, Unix sockets
- ⬜ **3.3 Signals** — Terms: SIGTERM, SIGKILL, SIGINT, SIGHUP, SIGCHLD, signal handlers, graceful shutdown, PID 1 problem in containers

## OS 4 — I/O
- ⬜ **4.1 I/O models** — Terms: blocking, non-blocking, I/O multiplexing, `select`/`poll`/`epoll`/`kqueue`, event loop, io_uring (awareness)
- ⬜ **4.2 File systems** — Terms: inode, directory entry, page cache, `fsync`, journaling, write-ahead log, ext4/xfs

## OS 5 — Concurrency theory
- ⬜ **5.1 Synchronization** — Terms: critical section, mutex, semaphore, spinlock, condition variable, monitor, atomic operation, CAS
- ⬜ **5.2 Problems** — Terms: race condition, deadlock (4 Coffman conditions), livelock, starvation, priority inversion, producer–consumer, dining philosophers

## OS 6 — Computer architecture
- ⬜ **6.1 CPU & caches** — Terms: core, hyper-threading, L1/L2/L3 cache, cache line (64 B), cache miss, locality (spatial/temporal), false sharing
- ⬜ **6.2 Memory ordering** — Terms: instruction reordering, memory barrier, happens-before, NUMA (awareness)

## OS 7 — Linux internals for containers
- ⬜ **7.1 Namespaces & cgroups** — Terms: pid/net/mnt/uts/ipc/user namespaces, cgroups v2, CPU quota & throttling, memory limit — foundation of Docker (see OPS 3)

---

# ⏭ TRACK NET — Networking (later — 4th)

**EN:** Every request crosses the network. "Random" timeouts and resets are usually DNS, TCP, TLS or load balancer issues. A senior can debug the full path.
**RU:** Каждый запрос проходит через сеть. «Случайные» таймауты и сбросы — обычно DNS, TCP, TLS или балансировщик. Senior умеет отлаживать весь путь.
**Dictionary:** `reference/networking.md`

## NET 1 — Models & the big picture
- ⬜ **1.1 OSI & TCP/IP models** — Terms: 7 layers vs 4 layers, encapsulation, frame/packet/segment, MTU
- ⬜ **1.2 "What happens when you type a URL"** — Terms: DNS lookup → TCP handshake → TLS handshake → HTTP request → server → response → rendering

## NET 2 — Link & network layer
- ⬜ **2.1 Ethernet & ARP** — Terms: MAC address, switch, ARP, broadcast domain
- ⬜ **2.2 IP** — Terms: IPv4, IPv6, CIDR, subnet mask, subnetting, private ranges (10/8, 172.16/12, 192.168/16), loopback, TTL, ICMP
- ⬜ **2.3 Routing & NAT** — Terms: router, default gateway, routing table, NAT, port forwarding, public vs private IP

## NET 3 — Transport layer
- ⬜ **3.1 TCP** — Terms: 3-way handshake, 4-way close, sequence/ack numbers, retransmission, flow control (receive window), congestion control (slow start, CUBIC, BBR), TIME_WAIT, Nagle's algorithm, keepalive, RST
- ⬜ **3.2 UDP & QUIC** — Terms: connectionless, datagram, when to use UDP, QUIC (over UDP), 0-RTT
- ⬜ **3.3 Ports & sockets** — Terms: port, socket, listen backlog, ephemeral ports, port exhaustion, `SO_REUSEADDR`

## NET 4 — DNS
- ⬜ **4.1 DNS resolution** — Terms: recursive resolver, root/TLD/authoritative servers, caching, TTL, `/etc/resolv.conf`
- ⬜ **4.2 Record types** — Terms: A, AAAA, CNAME, MX, TXT, NS, SRV, PTR, SOA; DNS as an outage source

## NET 5 — HTTP
- ⬜ **5.1 HTTP/1.1** — Terms: methods, status codes, headers, keep-alive, chunked encoding, head-of-line blocking
- ⬜ **5.2 HTTP/2 & HTTP/3** — Terms: multiplexing, streams, HPACK, server push (deprecated), HTTP/3 over QUIC
- ⬜ **5.3 HTTP features** — Terms: caching headers (`Cache-Control`, `ETag`), cookies, CORS, compression, redirects

## NET 6 — TLS
- ⬜ **6.1 TLS 1.3** — Terms: handshake, cipher suite, certificate, chain of trust, CA, SNI, session resumption, mTLS, Let's Encrypt / ACME, certificate expiry outages

## NET 7 — Proxies & load balancing
- ⬜ **7.1 Proxies** — Terms: forward proxy, reverse proxy, Nginx, Envoy, HAProxy, timeouts, buffering
- ⬜ **7.2 Load balancers & CDN** — Terms: L4 vs L7, health checks, connection draining, anycast, edge caching, cache keys, purge

## NET 8 — Cloud, Linux & Kubernetes networking
- ⬜ **8.1 Cloud networking** — Terms: VPC, public/private subnet, route table, internet gateway, NAT gateway, security group vs NACL, VPC peering, private endpoint
- ⬜ **8.2 Linux networking** — Terms: `ip addr`, `ip route`, iptables/nftables, network namespace, veth pair, bridge
- ⬜ **8.3 Kubernetes networking** — Terms: CNI, pod IP, Service (ClusterIP/NodePort/LoadBalancer), kube-proxy, Ingress, Gateway API, CoreDNS, NetworkPolicy, service mesh (Istio/Linkerd)
- ⬜ **8.4 VPN & zero trust** — Terms: WireGuard, IPsec (awareness), bastion host, zero-trust access

## NET 9 — Debugging tools
- ⬜ **9.1 Toolbox** — Terms: `curl -v`, `dig`, `nslookup`, `ping`, `traceroute`, `mtr`, `ss`, `netstat`, `tcpdump`, Wireshark, `openssl s_client`, `nc`, `telnet`

---

# ⏭ TRACK DIST — Distributed Systems & Messaging (later — 5th)

**EN:** Any backend with more than one instance, a queue or a cache is distributed. Networks fail, clocks lie, partial failure is normal.
**RU:** Любой backend с несколькими инстансами, очередью или кэшем — распределённый. Сеть падает, часы врут, частичный отказ — норма.
**Main book:** "Designing Data-Intensive Applications" (Kleppmann)

## DIST 1 — Fundamentals
- ⬜ **1.1 Why distributed is hard** — Terms: fallacies of distributed computing, partial failure, network partition, split brain
- ⬜ **1.2 CAP & PACELC** — Terms: consistency, availability, partition tolerance, latency vs consistency trade-off

## DIST 2 — Replication
- ⬜ **2.1 Replication models** — Terms: leader-follower, multi-leader, leaderless, quorum (R + W > N), replication lag, read-your-writes
- ⬜ **2.2 Consistency models** — Terms: linearizability, sequential, causal, eventual consistency, monotonic reads

## DIST 3 — Partitioning
- ⬜ **3.1 Sharding** — Terms: range vs hash partitioning, consistent hashing, virtual nodes, hot keys, rebalancing, scatter-gather

## DIST 4 — Consensus & time
- ⬜ **4.1 Consensus** — Terms: Raft (leader election, log replication, term), Paxos (awareness), etcd, ZooKeeper
- ⬜ **4.2 Time & ordering** — Terms: clock skew, NTP, Lamport clock, vector clock, hybrid logical clock, total order
- ⬜ **4.3 Coordination** — Terms: distributed lock, lease, fencing token, leader election in practice

## DIST 5 — Distributed transactions
- ⬜ **5.1 Patterns** — Terms: 2PC (and why avoided), saga (orchestration vs choreography), compensating transaction, transactional outbox, inbox / idempotent consumer, exactly-once vs effectively-once

## DIST 6 — Resilience
- ⬜ **6.1 Failure handling** — Terms: timeout, retry, exponential backoff, jitter, retry storm, circuit breaker, bulkhead, load shedding, backpressure, graceful degradation, deadline propagation, cache stampede protection (singleflight)

## DIST 7 — Messaging deep
- ⬜ **7.1 Kafka** — Terms: broker, topic, partition, offset, consumer group, rebalancing, replication factor, ISR, acks, idempotent producer, transactions, retention, log compaction, ordering per partition, KRaft
- ⬜ **7.2 RabbitMQ & NATS** — Terms: exchange (direct/topic/fanout), queue, binding, ack, prefetch, DLX; NATS subjects, JetStream
- ⬜ **7.3 Patterns** — Terms: pub/sub vs queue, event vs command, event schema evolution, DLQ, poison message, ordering guarantees

## DIST 8 — Stream processing (awareness)
- ⬜ **8.1 Streaming** — Terms: stream vs batch, windowing, watermarks, stateful processing, Kafka Streams, Flink, CDC (Debezium)

---

# ⏭ TRACK OPS — DevOps, Cloud & SRE (later — 6th)

**EN:** "You build it, you run it." Goal: take your Go service from `git push` to production yourself — build, containerize, deploy, observe, debug, roll back.
**RU:** «Ты построил — ты и эксплуатируешь». Цель: самому довести Go-сервис от `git push` до продакшена — собрать, упаковать, задеплоить, наблюдать, отлаживать, откатывать.
**Dictionary:** `reference/devops.md`

## OPS 1 — Linux & shell
- ⬜ **1.1 Filesystem & permissions** — Terms: FHS (`/etc`, `/var`, `/proc`), users/groups, `chmod`/`chown`, rwx, setuid, `sudo`
- ⬜ **1.2 Processes & services** — Terms: `ps`, `top`/`htop`, `kill`, `systemd`, unit file, `systemctl`, `journalctl`
- ⬜ **1.3 Resources** — Terms: `df`, `du`, `free`, `vmstat`, `iostat`, `lsof`, load average
- ⬜ **1.4 Text processing** — Terms: `grep`, `sed`, `awk`, `jq`, `xargs`, `sort`, `uniq`, `wc`, pipes, redirection
- ⬜ **1.5 Bash scripting** — Terms: variables, quoting, conditionals, loops, functions, exit codes, `set -euo pipefail`, shebang
- ⬜ **1.6 SSH & remote work** — Terms: SSH keys, `~/.ssh/config`, port forwarding/tunnels, `scp`, `rsync`, `tmux`

## OPS 2 — Git deep
- ⬜ **2.1 Internals** — Terms: blob, tree, commit, refs, HEAD, DAG, index (staging area)
- ⬜ **2.2 Workflows** — Terms: merge vs rebase, interactive rebase, squash, cherry-pick, `reflog`, `bisect`, stash, conflicts
- ⬜ **2.3 Team practices** — Terms: trunk-based development, GitFlow, feature branches, conventional commits, signed commits, hooks, monorepo vs polyrepo

## OPS 3 — Containers
- ⬜ **3.1 What a container is** — Terms: namespaces, cgroups, union filesystem (overlayfs), container vs VM, OCI image, runtime (runc, containerd)
- ⬜ **3.2 Dockerfile** — Terms: `FROM`, `RUN`, `COPY`, `CMD` vs `ENTRYPOINT`, layers & caching, multi-stage build, distroless/scratch, non-root `USER`, `.dockerignore`, `HEALTHCHECK`
- ⬜ **3.3 Images & registries** — Terms: tag vs digest, Docker Hub, GHCR, ECR, image scanning (Trivy), image size optimization
- ⬜ **3.4 Running containers** — Terms: ports, volumes, bind mounts, networks (bridge, host), env vars, logs, `docker exec`, restart policies
- ⬜ **3.5 Docker Compose** — Terms: services, depends_on, healthchecks, volumes, local dev stack (Go + Postgres + Redis + Kafka)

## OPS 4 — CI/CD
- ⬜ **4.1 Pipeline design** — Terms: lint → test → build → scan → publish → deploy, fail fast, build once deploy many
- ⬜ **4.2 GitHub Actions / GitLab CI** — Terms: workflow, job, step, runner, matrix, cache, artifacts, secrets, reusable workflows, OIDC to cloud
- ⬜ **4.3 Versioning & environments** — Terms: semantic versioning, release tags, dev/staging/prod, promotion, approvals
- ⬜ **4.4 Deployment strategies** — Terms: recreate, rolling update, blue-green, canary, feature flags, dark launch, automated rollback, progressive delivery

## OPS 5 — Kubernetes
- ⬜ **5.1 Architecture** — Terms: cluster, node, control plane, API server, etcd, scheduler, controller manager, kubelet, kube-proxy, container runtime
- ⬜ **5.2 Workloads** — Terms: Pod, Deployment, ReplicaSet, StatefulSet, DaemonSet, Job, CronJob, labels & selectors, namespaces
- ⬜ **5.3 Config & storage** — Terms: ConfigMap, Secret, env vs volume mount, PersistentVolume, PersistentVolumeClaim, StorageClass
- ⬜ **5.4 Networking** — Terms: Service types, Ingress, Gateway API, DNS (see NET 8.3)
- ⬜ **5.5 Health & resources** — Terms: liveness / readiness / startup probes, requests vs limits, QoS classes, OOMKilled, CPU throttling, `GOMAXPROCS`/`GOMEMLIMIT` in pods
- ⬜ **5.6 Scaling & availability** — Terms: HPA, VPA, cluster autoscaler, PodDisruptionBudget, affinity/anti-affinity, taints & tolerations
- ⬜ **5.7 Packaging & GitOps** — Terms: Helm (chart, values, release), Kustomize (overlays), Argo CD, Flux, desired state
- ⬜ **5.8 Debugging** — Terms: `kubectl get/describe/logs/exec/port-forward/top/events`, CrashLoopBackOff, ImagePullBackOff, Pending, graceful shutdown (`preStop`, `terminationGracePeriodSeconds`)
- ⬜ **5.9 Local clusters** — Terms: kind, k3d, minikube

## OPS 6 — Infrastructure as Code
- ⬜ **6.1 Terraform / OpenTofu** — Terms: provider, resource, data source, variable, output, state, remote state & locking, module, plan/apply, drift, import, workspace
- ⬜ **6.2 Configuration management** — Terms: Ansible (playbook, inventory, idempotency), immutable infrastructure

## OPS 7 — Cloud (AWS first, map to GCP / Yandex Cloud)
- ⬜ **7.1 Core concepts** — Terms: region, availability zone, shared responsibility model, Well-Architected pillars
- ⬜ **7.2 IAM** — Terms: user, role, policy, least privilege, assume role, workload identity
- ⬜ **7.3 Compute** — Terms: EC2, Auto Scaling Group, ECS/Fargate, EKS, Lambda (serverless trade-offs)
- ⬜ **7.4 Storage & data** — Terms: S3 (buckets, lifecycle, presigned URLs), EBS, RDS/Aurora, ElastiCache
- ⬜ **7.5 Messaging & edge** — Terms: SQS, SNS, EventBridge, MSK, ALB vs NLB, Route 53, CloudFront
- ⬜ **7.6 Cost** — Terms: on-demand vs reserved vs spot, rightsizing, cost allocation tags, FinOps

## OPS 8 — Observability
- ⬜ **8.1 Logs** — Terms: structured logs, log levels, log aggregation (Loki, ELK), retention, correlation ID
- ⬜ **8.2 Metrics** — Terms: Prometheus, scraping, PromQL, counter/gauge/histogram, cardinality, Grafana dashboards, RED method, USE method
- ⬜ **8.3 Traces** — Terms: OpenTelemetry, span, trace context propagation, sampling, Jaeger/Tempo, OTel Collector
- ⬜ **8.4 Profiles & eBPF** — Terms: continuous profiling (Pyroscope/Parca), eBPF (awareness)

## OPS 9 — SRE practices
- ⬜ **9.1 Reliability targets** — Terms: SLI, SLO, SLA, error budget, burn-rate alerts
- ⬜ **9.2 Alerting** — Terms: symptom-based alerts, alert fatigue, paging vs ticket, runbook
- ⬜ **9.3 Incidents** — Terms: on-call, severity levels, incident commander, status page, blameless postmortem, action items, MTTR
- ⬜ **9.4 Resilience practice** — Terms: capacity planning, load testing before launch, chaos engineering, disaster recovery, RPO/RTO, game days

---

# ⏭ TRACK SEC — Security (later — 7th)

**EN:** Security is a property of the whole system. AI-generated code often has insecure defaults — seniors catch them.
**RU:** Безопасность — свойство всей системы. Сгенерированный ИИ код часто небезопасен по умолчанию — senior это замечает.
**Dictionary:** `reference/security.md`

## SEC 1 — Secure coding
- ⬜ **1.1 OWASP Top 10** — Terms: broken access control, cryptographic failures, injection (SQL, command), insecure design, security misconfiguration, vulnerable components, auth failures, integrity failures, logging failures, SSRF
- ⬜ **1.2 OWASP API Top 10** — Terms: BOLA/IDOR, broken authentication, mass assignment, unrestricted resource consumption, broken function-level authorization
- ⬜ **1.3 Web attacks** — Terms: XSS (stored/reflected/DOM), CSRF, clickjacking, open redirect, path traversal, insecure deserialization
- ⬜ **1.4 Defensive basics** — Terms: input validation, output encoding, parameterized queries, least privilege, safe error messages, secure defaults, CORS (what it does and doesn't protect)

## SEC 2 — Authentication
- ⬜ **2.1 Sessions & cookies** — Terms: session ID, HttpOnly, Secure, SameSite, session fixation
- ⬜ **2.2 Passwords** — Terms: bcrypt, scrypt, argon2id, salt, pepper, credential stuffing, MFA, TOTP, passkeys/WebAuthn
- ⬜ **2.3 Tokens** — Terms: JWT (header, payload, signature), HS256 vs RS256/EdDSA, claims (`exp`, `iat`, `aud`, `iss`), access vs refresh token, revocation problem, token rotation
- ⬜ **2.4 OAuth 2.0 & OIDC** — Terms: resource owner, client, authorization server, authorization code + PKCE, client credentials, scopes, ID token, OpenID Connect, SSO, Keycloak

## SEC 3 — Authorization
- ⬜ **3.1 Models** — Terms: RBAC, ABAC, ReBAC (Zanzibar, OpenFGA), policy engine (OPA), multi-tenant isolation

## SEC 4 — Cryptography (use correctly, never invent)
- ⬜ **4.1 Primitives** — Terms: hashing (SHA-256), HMAC, symmetric encryption (AES-GCM), asymmetric (RSA, ECDSA, Ed25519), digital signature, key exchange, nonce, `crypto/rand`
- ⬜ **4.2 Key management** — Terms: KMS, envelope encryption, key rotation, encryption at rest vs in transit, HSM (awareness)

## SEC 5 — Infrastructure & supply chain
- ⬜ **5.1 Secrets management** — Terms: Vault, cloud secret managers, Kubernetes Secrets (base64 ≠ encryption), SOPS, secret scanning (gitleaks)
- ⬜ **5.2 Supply chain** — Terms: dependency scanning (`govulncheck`, Dependabot, Renovate), SBOM, image signing (cosign/Sigstore), SLSA, pinning versions, typosquatting
- ⬜ **5.3 Container & Kubernetes security** — Terms: non-root, read-only rootfs, capabilities, Pod Security Standards, NetworkPolicy, RBAC, image scanning
- ⬜ **5.4 Security in CI** — Terms: SAST (gosec, Semgrep), DAST (ZAP), secret scanning, branch protection

## SEC 6 — Security process
- ⬜ **6.1 Threat modeling** — Terms: STRIDE, data flow diagram, trust boundary, attack surface
- ⬜ **6.2 Operations** — Terms: audit logging, rate limiting & bot protection, DDoS basics, security incident response, responsible disclosure
- ⬜ **6.3 Compliance awareness** — Terms: GDPR, personal data laws (Uzbekistan / Russia data localization), PCI DSS, SOC 2

## SEC 7 — AI security
- ⬜ **7.1 OWASP LLM Top 10** — Terms: prompt injection (direct/indirect), sensitive information disclosure, excessive agency, insecure output handling, data/model poisoning, secrets in prompts

---

# ⏭ TRACK QA — Testing & Quality Engineering (later — 8th)

**EN:** In the AI era testing is the main defense against plausible-but-wrong code. Learn both the developer view (test code) and the QA view (how things break).
**RU:** В эпоху ИИ тестирование — главная защита от правдоподобного, но неверного кода. Изучаем и взгляд разработчика (тестовый код), и взгляд QA (как всё ломается).
*(Go-specific testing is in GO 12 — this track is the strategy & QA discipline.)*

## QA 1 — Testing theory
- ⬜ **1.1 Why & what** — Terms: test pyramid, testing trophy, unit/integration/E2E, cost vs confidence, testing behavior not implementation
- ⬜ **1.2 Test doubles theory** — Terms: dummy, stub, fake, spy, mock, London vs Classical (Detroit) school
- ⬜ **1.3 Good tests** — Terms: fast, isolated, repeatable, self-validating, readable (FIRST), Arrange-Act-Assert, flaky test causes

## QA 2 — Test levels in practice
- ⬜ **2.1 Integration testing** — Terms: real dependencies, testcontainers, test data setup/teardown, test database
- ⬜ **2.2 Contract testing** — Terms: consumer-driven contracts, Pact, schema compatibility (OpenAPI/protobuf breaking-change checks)
- ⬜ **2.3 End-to-end testing** — Terms: E2E scope, staging environments, Playwright (awareness), keeping E2E suites small
- ⬜ **2.4 API testing tools** — Terms: Postman / Bruno, Hurl, Newman, collections, environments

## QA 3 — Advanced techniques
- ⬜ **3.1 Property-based & fuzz testing** — Terms: properties, generators, shrinking, `rapid`, Go native fuzzing
- ⬜ **3.2 Mutation testing** — Terms: mutants, mutation score, "do my tests catch bugs?"
- ⬜ **3.3 Performance testing** — Terms: k6, vegeta, ghz; smoke/load/stress/spike/soak tests, throughput, percentiles, SLO-based thresholds
- ⬜ **3.4 Chaos & resilience testing** — Terms: fault injection, toxiproxy, chaos experiments

## QA 4 — QA discipline
- ⬜ **4.1 Test design techniques** — Terms: equivalence partitioning, boundary value analysis, decision table, state transition, pairwise testing, error guessing
- ⬜ **4.2 Test types** — Terms: smoke, sanity, regression, exploratory, acceptance, risk-based testing
- ⬜ **4.3 Artifacts** — Terms: test strategy, test plan, test case, bug report (steps, expected vs actual), severity vs priority, defect lifecycle
- ⬜ **4.4 Methods** — Terms: TDD (red-green-refactor), BDD (Given/When/Then), acceptance criteria, Definition of Done, ISTQB vocabulary

## QA 5 — Quality in production
- ⬜ **5.1 Testing in production** — Terms: canary analysis, synthetic monitoring, feature flags, shadow traffic

## QA 6 — Quality metrics & gates
- ⬜ **6.1 Metrics** — Terms: code coverage (and its limits), escaped defects, DORA metrics (deployment frequency, lead time, change failure rate, time to restore), quality gates in CI, code review as QA

---

# ⏭ TRACK ARCH — Software Design, Architecture & API Design (later — 9th)

**EN:** Architecture = decisions that are expensive to change. AI fills in code; *you* define the shape and boundaries.
**RU:** Архитектура — решения, которые дорого менять. ИИ заполняет код, а *ты* задаёшь форму и границы.
**Dictionary:** `reference/architecture.md`

## ARCH 1 — Design principles
- ⬜ **1.1 Fundamentals** — Terms: coupling, cohesion, abstraction, encapsulation, information hiding, DRY, KISS, YAGNI, separation of concerns
- ⬜ **1.2 SOLID** — Terms: SRP, OCP, LSP, ISP, DIP — applied to Go
- ⬜ **1.3 Code quality** — Terms: code smells, refactoring techniques, technical debt, readability, "deep modules" (Ousterhout)

## ARCH 2 — Design patterns
- ⬜ **2.1 Creational** — Terms: factory, builder, functional options, singleton, prototype
- ⬜ **2.2 Structural** — Terms: adapter, decorator, facade, proxy, composite
- ⬜ **2.3 Behavioral** — Terms: strategy, observer, command, chain of responsibility (middleware), state, template method, iterator

## ARCH 3 — Application architecture
- ⬜ **3.1 Architecture styles** — Terms: layered, clean architecture, hexagonal (ports & adapters), onion, vertical slices
- ⬜ **3.2 Domain-Driven Design** — Terms: ubiquitous language, bounded context, context map, aggregate, entity, value object, domain event, repository, anti-corruption layer
- ⬜ **3.3 Monolith vs microservices** — Terms: modular monolith, microservices, distributed monolith anti-pattern, service boundaries, database per service, strangler fig
- ⬜ **3.4 Event-driven & CQRS** — Terms: event-driven architecture, choreography vs orchestration, CQRS, event sourcing, projections, eventual consistency

## ARCH 4 — API design
- ⬜ **4.1 REST design** — Terms: resources, HTTP semantics, idempotency & safety, status codes, filtering, sorting, cursor pagination, error model (RFC 9457 Problem Details), HATEOAS (awareness)
- ⬜ **4.2 API evolution** — Terms: versioning (URL/header), backward/forward compatibility, deprecation, OpenAPI, API-first design
- ⬜ **4.3 Reliability features** — Terms: idempotency keys, rate limiting headers, webhooks (signing, retries), long-running operations, bulk endpoints
- ⬜ **4.4 Other styles** — Terms: gRPC, GraphQL (N+1, DataLoader), WebSocket, SSE, AsyncAPI, API gateway

## ARCH 5 — Architecture decisions
- ⬜ **5.1 Documentation** — Terms: ADR, RFC, C4 model diagrams, design review
- ⬜ **5.2 Organization** — Terms: Conway's law, team topologies, platform team, ownership boundaries

## ARCH 6 — Advanced (awareness)
- ⬜ **6.1** Terms: evolutionary architecture, fitness functions, cell-based architecture, multi-tenancy architectures

---

# ⏭ TRACK AI — AI Engineering (later — 10th)

**EN:** Two skills: (1) **using AI** to work faster without losing quality, (2) **building AI features** into backend products.
**RU:** Два навыка: (1) **использовать ИИ**, чтобы работать быстрее без потери качества, (2) **встраивать ИИ-функции** в backend-продукты.
*(You use AI every day already — AI 1 can be picked any time.)*

## AI 1 — AI-assisted engineering
- ⬜ **1.1 Workflow** — Terms: task decomposition, spec-first prompting, context engineering (CLAUDE.md, rules files), small verifiable steps, agentic coding tools
- ⬜ **1.2 Verification** — Terms: reviewing AI code (correctness, concurrency, security, error handling), hallucinated APIs/packages, outdated versions, over-engineering, AI-free practice for interviews

## AI 2 — LLM fundamentals
- ⬜ **2.1 How LLMs behave** — Terms: token, context window, temperature, sampling, non-determinism, hallucination, system prompt, few-shot examples
- ⬜ **2.2 LLM APIs from Go** — Terms: streaming, timeouts, retries, rate limits, provider fallbacks, structured output (JSON schema), tool / function calling, prompt caching, cost per token

## AI 3 — Building AI features
- ⬜ **3.1 RAG** — Terms: chunking, embeddings, vector database, pgvector, HNSW vs IVFFlat, similarity search, hybrid search (BM25 + vectors), re-ranking, citations
- ⬜ **3.2 Agents & MCP** — Terms: agent loop, tools, planning, Model Context Protocol (server, client, tools, resources), human-in-the-loop, permissions
- ⬜ **3.3 Evals & observability** — Terms: eval dataset, LLM-as-judge, regression evals, prompt versioning, tracing LLM calls, token usage metrics
- ⬜ **3.4 Production concerns** — Terms: latency, cost control, model routing, caching, guardrails, PII handling (see SEC 7)

## AI 4 — Awareness
- ⬜ **4.1** Terms: fine-tuning vs RAG vs prompting, self-hosted models (Ollama, vLLM), GPU cost, multi-agent systems

---

# ⏭ TRACK CAREER — Senior Skills, Interviews & Career (later — 11th, but behavioral prep any time)

**EN:** The difference between a strong middle and a senior is mostly here. AI cannot do this for you.
**RU:** Разница между сильным middle и senior в основном здесь. ИИ этого за тебя не сделает.
**Dictionary:** `reference/behavioral.md` · **Interview files:** `interviews/`

## CAREER 1 — Engineering practices
- ⬜ **1.1 Code review** — Terms: small PRs, review checklist, kind & specific feedback, blocking vs non-blocking comments
- ⬜ **1.2 Delivery process** — Terms: Scrum, Kanban, sprint, backlog refinement, estimation, story points vs time, Definition of Done
- ⬜ **1.3 Documentation** — Terms: README, runbook, design doc, ADR, onboarding docs

## CAREER 2 — Senior behaviors
- ⬜ **2.1 Ownership** — Terms: end-to-end ownership, raising risks early, follow-through, "disagree and commit"
- ⬜ **2.2 Communication** — Terms: explaining trade-offs to non-engineers, asking clarifying questions, written communication
- ⬜ **2.3 Product thinking** — Terms: user impact, business metrics, prioritization (impact vs effort), saying "no" with alternatives

## CAREER 3 — Leadership & mentoring
- ⬜ **3.1 Mentoring** — Terms: teaching (your student!), pair programming, onboarding, giving feedback, growth plans
- ⬜ **3.2 Leading projects** — Terms: planning, splitting work, coordinating, stakeholder management, running technical discussions, influence without authority

## CAREER 4 — Interviews
- ⬜ **4.1 Behavioral (STAR)** — Stories: conflict, failure, impact, leadership, ambiguity, tight deadline — from `interviews/my-projects.md`
- ⬜ **4.2 Technical rounds** — Go deep-dive (GO track), DSA live coding, system design, SQL/PG — mock sessions in `interviews/sessions/`
- ⬜ **4.3 Negotiation & choosing companies** — Terms: offer, compensation, engineering culture signals, on-call expectations

## CAREER 5 — Visibility
- ⬜ **5.1 Portfolio & writing** — Terms: production-quality GitHub project (capstone in `reference/00-senior-engineer-guide.md` §20), technical blog, LinkedIn posts, open-source contributions, community (Go meetups)

---

## Progress summary (update when a topic's subtopics change)

| Track | ⬜ | 🟡 | 🟢 | 🔵 |
|---|---|---|---|---|
| GO | all | — | — | — |
| DSA | most | 1, 2.1, 3.1, 4, 5, 6.1, 7.1, 8.1, 9, 10.1 | — | — |
| EN | all | — | — | — |
| RU | all | — | — | — |
| SD | 1, 2.8, 3–6 | 2.1–2.7 | — | — |
| others | all | — | — | — |
