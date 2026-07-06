# Quick Reference — Go Backend Interview

---

## Go Concurrency

**What is a goroutine leak?**
> A goroutine that starts but never finishes. Cause: blocked channel or infinite loop with no exit. Fix: always listen to `ctx.Done()` inside goroutines.

**Buffered vs unbuffered channel?**
> Unbuffered: sender blocks until receiver is ready — synchronization.
> Buffered: sender blocks only when buffer is full — acts as a queue.
> Use unbuffered for sync signals, buffered for task queues.

**What is a worker pool?**
> Fixed number of goroutines reading tasks from a channel. Limits parallelism — prevents system overload.
```go
jobs := make(chan Job, 100)
for i := 0; i < 5; i++ {
    go func() {
        for job := range jobs { process(job) }
    }()
}
```

**What is sync.WaitGroup?**
> Waits for multiple goroutines to finish. `Add(n)` → launch goroutines → each calls `Done()` → `Wait()` blocks until counter hits 0.

**Mutex vs RWMutex?**
> Mutex: one goroutine at a time — blocks everyone.
> RWMutex: multiple goroutines can read simultaneously, but write blocks all. Use RWMutex when reads far outnumber writes (e.g. cache).

**What is a data race?**
> Two goroutines access the same variable simultaneously and at least one writes. Result is unpredictable.
> Fix: `sync.Mutex`, channels, `sync/atomic`. Detect with `go test -race`.

---

## Error Handling

**What is error wrapping?**
> Adding context to an error: `fmt.Errorf("creating order: %w", err)`.
> `errors.Is` checks the error chain for a specific error.
> `errors.As` extracts a specific error type from the chain.

```go
return fmt.Errorf("creating order: %w", err)   // wrap

errors.Is(err, ErrNotFound)   // check type in chain
errors.As(err, &valErr)       // extract concrete type
```

**What is a sentinel error?**
> A pre-declared error variable: `var ErrNotFound = errors.New("not found")`.
> Compare using `errors.Is`, not `==`.

**panic vs error?**
> `error` — expected failure, part of normal flow (file not found, invalid input). Caller handles it.
> `panic` — unexpected state, programming bug (nil pointer, index out of bounds). Stops the goroutine, unwinds the stack. Can be caught with `recover()` in defer.
> Rule: never use panic for business errors.

---

## Go Patterns

**What are functional options?**
> Pattern for configuring a struct with optional parameters. Each option is a function that modifies the config.
```go
type Option func(*Server)

func WithPort(p int) Option    { return func(s *Server) { s.port = p } }
func WithTimeout(t time.Duration) Option { return func(s *Server) { s.timeout = t } }

func NewServer(opts ...Option) *Server {
    s := &Server{port: 8080}   // defaults
    for _, opt := range opts { opt(s) }
    return s
}
```

**What is a factory?**
> A function that creates objects. Hides creation details, returns an interface.
```go
func NewRepository(db *sql.DB) Repository {
    return &postgresRepo{db: db}
}
```

**What is dependency injection?**
> Object receives its dependencies from outside instead of creating them itself. Makes testing easy — swap real repo for mock.

---

## PostgreSQL

**What is ACID?**
> - **Atomicity** — all or nothing
> - **Consistency** — data always in valid state
> - **Isolation** — transactions don't see each other's uncommitted changes
> - **Durability** — after commit, data survives crash

**What is the N+1 problem?**
> Load N objects → make a separate query for each = N+1 queries total.
> Fix: JOIN or batch loading — get all data in one or two queries.

**Index types?**
> - **B-tree** — default, for `=`, `<`, `>`, BETWEEN, ORDER BY
> - **Composite** — multiple columns, column order matters (most selective first)
> - **Partial** — indexes only part of rows: `WHERE deleted_at IS NULL`
> - **Covering** — includes all needed columns, enables Index Only Scan

**EXPLAIN ANALYZE?**
> Shows query execution plan + real timing. Look for Seq Scan on large tables → add index. Check estimated vs actual rows — big gap means statistics are stale, run ANALYZE.

**Read Committed vs Repeatable Read?**
> Read Committed (default): you see only committed data, but two identical SELECTs can return different results.
> Repeatable Read: snapshot at transaction start — same data on repeated reads.

**SELECT FOR UPDATE?**
> Locks selected rows for the duration of the transaction. Other transactions that try to update those rows will wait. Used for pessimistic locking.

---

## Architecture

**What is graceful shutdown?**
> SIGTERM → stop accepting new requests → wait for current ones to finish → close DB/Kafka connections → exit.
> Without it: requests cut in the middle, partial data saved.

**Modular monolith vs microservices?**
> Microservices: separate services, independent deploy, but distributed transactions are hard, more infrastructure needed.
> Modular monolith: one app, one deploy, ACID transactions, but with clear module boundaries. Best for small-medium teams.

**What is context.Context?**
> Passes through all layers (handler → service → repository). Two purposes:
> 1. Cancellation — `WithCancel`, `WithTimeout`, `WithDeadline`
> 2. Request-scoped values — `WithValue`
> When parent context cancels → all children cancel too. Pass to every DB query and HTTP call.

---

## Security

**What is JWT?**
> Token = header.payload.signature (base64 encoded). Server signs with secret key, verifies on each request — no DB lookup needed.
> Access token: short-lived (15 min). Refresh token: long-lived (7 days).

**SQL injection?**
> Attack where user input gets into SQL as code. Fix: always use parameterized queries — never build SQL with string concatenation.

**bcrypt vs plain hash?**
> bcrypt is slow by design (cost factor) — makes brute force expensive. Also salted automatically — same password gives different hash each time.

---

## Message Queue & WebSocket

**What is a message queue?**
> Async communication between services. Producer writes message → queue stores it → consumer reads when ready. Benefits: decoupling, retry on failure, buffer during traffic spikes. Examples: Kafka, RabbitMQ.

**What is WebSocket?**
> Persistent two-way connection between client and server. Unlike HTTP — server can push data to client without a request. Used for: chat, real-time notifications, live data feeds.

---

## Garbage Collector

**How does Go GC work?**
> Tri-color mark and sweep, runs concurrently with the program. Three phases:
> 1. Mark setup (short STW pause)
> 2. Mark — finds all live objects starting from roots
> 3. Sweep — frees memory of unmarked (dead) objects
> Goal: minimize stop-the-world pauses.
