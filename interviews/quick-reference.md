# Quick Reference — Go Backend Interview

---

## 1. Goroutine Leak

**What it is:**
A goroutine leak happens when a goroutine starts but never finishes. It stays alive forever, consuming memory. Over time hundreds of leaked goroutines = app runs out of memory and crashes.

**Why it happens:**
```go
// BAD — goroutine blocked forever waiting on channel nobody writes to
func leak() {
    ch := make(chan int)
    go func() {
        val := <-ch  // blocks here forever — nobody sends
        fmt.Println(val)
    }()
    // function returns, but goroutine stays alive forever
}
```

**How to fix:**
```go
// GOOD — goroutine listens to ctx.Done() and exits when cancelled
func noLeak(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return  // clean exit
            case val := <-ch:
                process(val)
            }
        }
    }()
}
```

**Rule:** Every goroutine must have a way to exit. Always use `ctx.Done()` or a quit channel.

**How to detect:** `go test -race` catches some issues. For leaks specifically: `goleak` library or checking goroutine count in pprof.

---

## 2. Buffered vs Unbuffered Channel

**Unbuffered channel — synchronization:**
Sender blocks until receiver is ready. They must "meet" at the same time. Like a phone call — both people must be there simultaneously.

```go
ch := make(chan int)  // unbuffered

go func() {
    ch <- 42  // blocks here until someone receives
}()

val := <-ch  // both sides are now synchronized
```

**Buffered channel — queue:**
Sender puts item in the buffer and continues. Only blocks when buffer is full. Like a mailbox — you drop the letter and leave, recipient reads later.

```go
ch := make(chan int, 3)  // buffer of 3

ch <- 1  // doesn't block
ch <- 2  // doesn't block
ch <- 3  // doesn't block
ch <- 4  // BLOCKS — buffer is full
```

**When to use which:**
- Unbuffered → sync signals, handshake between goroutines
- Buffered → task queue, rate limiting, prevent goroutine from blocking

---

## 3. sync.Mutex vs sync.RWMutex

**The problem:** Map in Go is not thread-safe. Two goroutines writing to the same map simultaneously = crash.

**sync.Mutex — exclusive lock:**
Only one goroutine at a time, whether reading or writing. Like a bathroom with one key — one person at a time.

```go
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}
```

**sync.RWMutex — smarter lock:**
Multiple goroutines can READ simultaneously (RLock). But WRITE blocks everyone (Lock). Like a library — many people can read books at the same time, but when librarian reorganizes shelves, everyone waits.

```go
type SafeCache struct {
    mu    sync.RWMutex
    data  map[string]string
}

func (c *SafeCache) Get(key string) string {
    c.mu.RLock()          // multiple goroutines can hold RLock at same time
    defer c.mu.RUnlock()
    return c.data[key]
}

func (c *SafeCache) Set(key, val string) {
    c.mu.Lock()           // exclusive — blocks all readers and writers
    defer c.mu.Unlock()
    c.data[key] = val
}
```

**When to use RWMutex:** When you read much more often than you write (cache, config). If reads and writes are equal — Mutex is simpler.

---

## 4. Data Race

**What it is:**
Two goroutines access the same variable at the same time, and at least one is writing. The result is unpredictable — depends on which goroutine runs first (race condition).

```go
// BAD — classic data race
counter := 0
var wg sync.WaitGroup

for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        counter++  // not safe! multiple goroutines write simultaneously
    }()
}
wg.Wait()
// counter could be anything from 1 to 1000
```

**How to fix — 3 approaches:**

**Option 1: Mutex**
```go
var mu sync.Mutex
mu.Lock()
counter++
mu.Unlock()
```

**Option 2: Atomic (best for simple counters)**
```go
import "sync/atomic"
atomic.AddInt64(&counter, 1)  // thread-safe increment
```

**Option 3: Channel (Go idiomatic)**
```go
// Send work to single goroutine that owns the data
updates := make(chan int)
go func() {
    for range updates { counter++ }
}()
updates <- 1  // safe — only one goroutine touches counter
```

**How to detect:** Run `go test -race ./...` — Go's built-in race detector finds it automatically.

---

## 5. Panic vs Error

**error — expected failure:**
Error is just a value in Go. Functions return `(result, error)`. The caller decides what to do. Part of normal flow.

```go
// Expected: user might not exist, network might fail, file might be missing
func getUser(id int) (*User, error) {
    user, err := db.Query(...)
    if err != nil {
        return nil, fmt.Errorf("getting user %d: %w", id, err)
    }
    return user, nil
}
```

**panic — unexpected state, programming bug:**
Panic stops the current goroutine, unwinds the stack running all defers. If not recovered, crashes the program. Used for bugs that should never happen.

```go
// Go runtime panics automatically for:
var s []int
_ = s[10]           // index out of range — panic
var p *int
_ = *p              // nil pointer dereference — panic

// You can panic manually for impossible states:
func divide(a, b int) int {
    if b == 0 {
        panic("divide by zero — this should never be called with b=0")
    }
    return a / b
}
```

**recover() — catching panic:**
```go
func safeHandler() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("recovered from panic: %v", r)
            // return error response instead of crashing
        }
    }()
    riskyOperation()
}
```

**Rule:** Use `error` for anything a user or external system can cause. Use `panic` only for bugs in your own code that should be impossible.

---

## 6. Error Wrapping

**The problem:** A deep function returns an error. By the time it reaches the handler, you have no idea where it came from.

**Wrapping — add context at each layer:**
```go
// repository layer
func (r *repo) GetUser(id int) (*User, error) {
    err := r.db.QueryRow(...)
    if err != nil {
        return nil, fmt.Errorf("db query for user %d: %w", id, err)
    }
}

// service layer
func (s *service) GetUser(id int) (*User, error) {
    user, err := s.repo.GetUser(id)
    if err != nil {
        return nil, fmt.Errorf("service.GetUser: %w", err)
    }
}

// Result: "service.GetUser: db query for user 42: sql: no rows"
// You can trace exactly where it broke
```

**errors.Is — check specific error in chain:**
```go
var ErrNotFound = errors.New("not found")

// Even if wrapped 3 layers deep:
if errors.Is(err, ErrNotFound) {
    // return 404
}
```

**errors.As — extract specific error type:**
```go
type ValidationError struct {
    Field   string
    Message string
}

var valErr *ValidationError
if errors.As(err, &valErr) {
    fmt.Println(valErr.Field)  // "email"
}
```

---

## 7. Functional Options

**The problem:** You have a struct with many optional config values. Constructor with 10 parameters is ugly. Config struct works but is verbose.

**Functional options — elegant solution:**
```go
type Server struct {
    port     int
    timeout  time.Duration
    maxConns int
}

// Each option is a function that modifies Server
type Option func(*Server)

func WithPort(port int) Option {
    return func(s *Server) { s.port = port }
}

func WithTimeout(t time.Duration) Option {
    return func(s *Server) { s.timeout = t }
}

func WithMaxConns(n int) Option {
    return func(s *Server) { s.maxConns = n }
}

// Constructor applies defaults, then options
func NewServer(opts ...Option) *Server {
    s := &Server{
        port:     8080,           // sensible defaults
        timeout:  30 * time.Second,
        maxConns: 100,
    }
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// Usage — clean, readable, extensible
srv := NewServer(
    WithPort(9090),
    WithTimeout(60 * time.Second),
)
```

**Why it's good:** Easy to add new options without breaking existing callers. Popular in Go standard library and major packages (grpc, zap, etc).

---

## 8. Factory Pattern

**What it is:** A function that creates and returns an object. Hides the details of how the object is created. Returns an interface, not a concrete type.

```go
// Interface — what callers use
type Repository interface {
    GetUser(ctx context.Context, id int) (*User, error)
    CreateUser(ctx context.Context, user *User) error
}

// Concrete implementation — hidden from callers
type postgresRepository struct {
    db *sql.DB
}

// Factory — creates and returns the implementation
func NewRepository(db *sql.DB) Repository {
    return &postgresRepository{db: db}
}

// In main.go or app setup:
repo := NewRepository(db)          // caller gets Repository interface
service := NewService(repo)        // service only knows about the interface
```

**Why useful:**
- Caller doesn't know or care which database we use
- Easy to swap implementation for testing — create `MockRepository`
- All creation logic in one place

---

## 9. context.Context

**What it is:** An object that travels through your entire call chain — from HTTP handler all the way to the database query. Carries two things: cancellation signal and request-scoped values.

**Cancellation — stop wasted work:**
```go
// Client sends HTTP request → you start processing
// Client disconnects halfway through → without context, you keep working for nothing

func handleRequest(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()  // already has cancellation tied to client connection

    // If client disconnects, ctx is cancelled
    // All operations using ctx will stop automatically
    user, err := service.GetUser(ctx, userID)
    ...
}

func (s *service) GetUser(ctx context.Context, id int) (*User, error) {
    // Pass ctx to DB — if ctx is cancelled, query stops
    return s.repo.GetUser(ctx, id)
}

func (r *repo) GetUser(ctx context.Context, id int) (*User, error) {
    return r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id=$1", id).Scan(...)
}
```

**Types of cancellation:**
```go
// Manual cancel
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Timeout — auto-cancel after duration
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Deadline — auto-cancel at specific time
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
defer cancel()
```

**WithValue — carry request data:**
```go
// Set in middleware
ctx = context.WithValue(ctx, userIDKey, userID)

// Get deep in call chain — no need to pass userID as parameter
userID := ctx.Value(userIDKey).(int64)
```

**Important rules:**
- Never store context in a struct — pass as first function parameter
- Always call cancel() — use `defer cancel()`
- Never pass `nil` context — use `context.Background()` or `context.TODO()`

---

## 10. ACID

**Atomicity — all or nothing:**
If any step in a transaction fails, ALL steps are rolled back. You can't have partial saves.
```
Creating an order:
1. Insert order record
2. Deduct from inventory
3. Create payment record

If step 3 fails → steps 1 and 2 are rolled back automatically
```

**Consistency — valid state always:**
Database constraints (foreign keys, unique, not null, check) are enforced. Transaction can't leave DB in an invalid state.

**Isolation — transactions don't interfere:**
Two transactions running at the same time can't see each other's uncommitted changes. Prevents dirty reads.

**Levels (PostgreSQL):**
- Read Committed (default): see only committed data. Two reads in same transaction can return different results.
- Repeatable Read: snapshot at transaction start — same results on re-read.
- Serializable: strongest — transactions execute as if sequential.

**Durability — survives crashes:**
After COMMIT, data is written to disk. Even if server crashes immediately after, data is safe.

---

## 11. N+1 Problem

**What it is:** You load a list of N items, then for each item make another query = N+1 total queries.

```go
// BAD — N+1
orders, _ := db.Query("SELECT * FROM orders")  // 1 query

for _, order := range orders {
    // This runs once per order — 100 orders = 100 extra queries
    items, _ := db.Query("SELECT * FROM order_items WHERE order_id=$1", order.ID)
}
// Total: 1 + 100 = 101 queries
```

**Fix — use JOIN:**
```go
// GOOD — 1 query with JOIN
rows, _ := db.Query(`
    SELECT o.*, i.*
    FROM orders o
    LEFT JOIN order_items i ON i.order_id = o.id
`)
// Total: 1 query
```

**Fix — batch loading:**
```go
// Also good — 2 queries total
orders, _ := db.Query("SELECT * FROM orders")

// Collect all IDs
ids := collectIDs(orders)

// One query for all items
items, _ := db.Query("SELECT * FROM order_items WHERE order_id = ANY($1)", pq.Array(ids))
```

---

## 12. Indexes

**B-tree (default):**
Best for: equality (`=`), range (`<`, `>`, BETWEEN), sorting (ORDER BY).
```sql
CREATE INDEX idx_users_email ON users(email);
SELECT * FROM users WHERE email = 'x@y.com';  -- uses index
```

**Composite index — column order matters:**
Put the most selective column first (most values are unique). Works for queries that use prefix of the index.
```sql
CREATE INDEX idx_orders_user_status ON orders(user_id, status);
SELECT * FROM orders WHERE user_id = 5;               -- uses index
SELECT * FROM orders WHERE user_id = 5 AND status = 'pending'; -- uses index
SELECT * FROM orders WHERE status = 'pending';        -- does NOT use index (no user_id)
```

**Partial index — index only part of rows:**
Much smaller and faster when you only query a subset.
```sql
-- Only index active users
CREATE INDEX idx_active_users ON users(email) WHERE deleted_at IS NULL;
```

**Covering index — include extra columns:**
Query can be answered from index alone — no table lookup needed.
```sql
CREATE INDEX idx_users_covering ON users(email) INCLUDE (name, created_at);
SELECT name, created_at FROM users WHERE email = 'x@y.com';
-- Gets data entirely from index — faster
```

---

## 13. EXPLAIN ANALYZE

**What it shows:** The actual execution plan — how PostgreSQL decided to run your query, what it actually did, and how long it took.

```sql
EXPLAIN ANALYZE
SELECT * FROM orders WHERE user_id = 42;
```

**Output to look for:**

| What you see | What it means |
|---|---|
| `Seq Scan` | Reading entire table — bad on large tables |
| `Index Scan` | Using an index — good |
| `Index Only Scan` | Even better — data from index alone |
| `Nested Loop` | Join algorithm — good for small datasets |
| `Hash Join` | Join algorithm — good for larger datasets |

**Key numbers:**
- `rows=100 actual rows=10000` — huge gap means statistics are stale, run `ANALYZE`
- `actual time=850ms` — slow, needs optimization
- `loops=1000` — this step ran 1000 times — classic N+1

**Workflow:**
1. Run EXPLAIN ANALYZE on slow query
2. Find Seq Scan on large table → add index
3. Find huge estimated vs actual row difference → run ANALYZE
4. Find high loop count → N+1 problem → rewrite with JOIN

---

## 14. Graceful Shutdown

**Without it:**
```
Server gets SIGTERM (deploy, restart)
→ Process killed immediately
→ 50 requests being processed — all cut in the middle
→ Some DB transactions half-committed
→ Users see errors, data corrupted
```

**With it:**
```
Server gets SIGTERM
→ Stop accepting NEW requests (close listener)
→ Wait for current 50 requests to finish
→ Close DB connections properly
→ Stop Kafka consumers cleanly
→ Exit — no data loss, no errors for users
```

```go
func (a *App) Run() {
    // Start HTTP server in background goroutine
    go httpServer.Start()

    // Wait for SIGTERM or SIGINT
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
    <-sigs

    // Graceful shutdown sequence
    httpServer.Shutdown(context.Background())  // stop new requests, wait for current
    kafkaConsumer.Stop()                       // stop consuming
    db.Close()                                 // close connections
}
```

---

## 15. JWT Authentication

**Structure:** Three base64-encoded parts separated by dots:
```
eyJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOjQyfQ.abc123
     header                   payload          signature
```

**Flow:**
```
1. User logs in with username/password
2. Server verifies password (bcrypt compare)
3. Server creates JWT: payload = {userId, role, exp}
4. Server signs with secret key → sends token to client

5. Client stores token, sends in every request: Authorization: Bearer <token>
6. Server verifies signature — if valid, trusts the payload
7. No DB lookup needed on each request — signature proves authenticity
```

**Access + Refresh tokens:**
- Access token: short-lived (15 min) — used for API calls
- Refresh token: long-lived (7 days) — stored securely, used only to get new access token
- When access token expires → use refresh token → get new access token

**What to check on each request:**
1. Signature valid? (not tampered)
2. Not expired? (check `exp` claim)
3. User has permission for this resource? (authorization)

---

## 16. Message Queue

**What it is:** An intermediate storage where one service puts messages (producer) and another service reads them (consumer) asynchronously.

**Why use it:**
- **Decoupling:** Order service doesn't call Email service directly. Order service → queue → Email service reads when ready
- **Retry:** If Email service is down, messages stay in queue. When it recovers, processes them
- **Buffer:** 10,000 requests per second, Email service can only handle 1,000. Queue absorbs the spike

**Kafka specifically:**
- Messages organized in topics (like categories)
- Topics split into partitions (for parallel processing)
- Consumer group — multiple consumers share the work
- Messages stay in topic even after reading (unlike traditional queues)

**At-least-once vs exactly-once:**
- At-least-once: message processed at least once, but might be processed twice (consumer crashed after processing but before committing offset). Fix: make consumer idempotent (check if already processed)
- Exactly-once: very complex, high overhead — usually not worth it

---

## 17. WebSocket

**HTTP problem:** Client must always initiate. Server can't push data to client without a request.

**WebSocket solution:** Client and server establish a persistent two-way connection. Either side can send data at any time.

```
HTTP:  Client → Request → Server → Response → Connection closed
       Client → Request → Server → Response → Connection closed
       (new connection every time)

WebSocket: Client → Upgrade request → Server → 101 Switching Protocols
           Client ←→ Server (persistent, bidirectional)
           (one connection, both sides can send anytime)
```

**When to use:**
- Real-time chat
- Live notifications (order status changed, new message)
- Live data feeds (stock prices, sports scores)
- Collaborative editing (Google Docs style)

**When NOT to use:** Regular CRUD operations — HTTP is simpler and more scalable for that.

---

## 18. Garbage Collector

**What it is:** Automatic memory management. You don't free memory manually in Go — GC does it.

**How Go's GC works — tri-color mark and sweep:**

**Phase 1 — Mark Setup (short STW pause):**
Stop the world briefly to set up the marking phase. Enables write barriers.

**Phase 2 — Mark (concurrent, runs with your program):**
Starting from "roots" (global variables, goroutine stacks), follows all pointers and marks everything reachable as "live". Unmarked objects = garbage.

**Phase 3 — Mark Termination (short STW pause):**
Final cleanup of marking phase.

**Phase 4 — Sweep (concurrent):**
Frees memory of all unmarked objects. Returns memory to heap for reuse.

**Key point:** Most of GC runs concurrently with your program. STW pauses are very short (sub-millisecond in modern Go). This is why Go is suitable for latency-sensitive services.

**GOGC environment variable:** Controls how often GC runs. Default 100 = GC runs when heap doubles. Higher value = less frequent GC, more memory used.
