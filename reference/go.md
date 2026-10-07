# 07 — Go Language (Deep)

> Go-specific knowledge. Must be able to explain internals, not just use the language.

---

## Goroutines

**Key Terms:**
- `goroutine` — lightweight concurrent unit; starts with `go func()`; initial stack ~2KB (grows dynamically)
- `OS thread` — heavyweight; typically 1-8MB stack; created by OS
- `M:N scheduling` — Go maps M goroutines to N OS threads (N = GOMAXPROCS)
- `GOMAXPROCS` — number of OS threads that can run Go code simultaneously; default = number of CPU cores
- `Go runtime scheduler` — schedules goroutines onto OS threads; preemptive since Go 1.14
- `goroutine switch` — switching goroutines is cheap (~100ns vs ~1µs for OS thread context switch)
- `goroutine leak` — goroutine created but never terminates; memory grows unboundedly
- `goroutine stack growth` — stack starts small, grows automatically (segmented or contiguous stacks)
- `runtime.NumGoroutine()` — returns current goroutine count; use to detect leaks
- `runtime.Gosched()` — yields the current goroutine; lets other goroutines run
- `sync.WaitGroup` — wait for multiple goroutines to finish
- `goroutine identity` — goroutines don't have accessible IDs (by design; prevents goroutine-local storage abuse)

**Common goroutine leak patterns:**
- Goroutine blocked on channel that nobody sends to
- Goroutine in infinite loop with no exit condition
- Goroutine waiting on lock that's never released

**Prevention:** always pass `context.Context` and select on `ctx.Done()`

**Must understand:**
- [ ] Why goroutines are cheap compared to OS threads (small stack, userspace scheduling)
- [ ] GOMAXPROCS — what it controls and why setting it to cores is default
- [ ] How to detect a goroutine leak (runtime.NumGoroutine(), pprof)
- [ ] Why goroutine leaks are dangerous (memory, CPU, file descriptor exhaustion)

---

## Channels

**Key Terms:**
- `channel` — typed conduit for passing values between goroutines; created with `make(chan T)`
- `unbuffered channel` — `make(chan T)`; send blocks until receiver is ready; guarantees synchronization
- `buffered channel` — `make(chan T, N)`; send blocks only when buffer full; decouples sender and receiver
- `send` — `ch <- value`; blocks if channel full (buffered) or no receiver (unbuffered)
- `receive` — `value := <-ch`; blocks if channel empty
- `close(ch)` — signals no more values will be sent; receivers get zero value after drain
- `range over channel` — `for v := range ch` — loops until channel closed
- `nil channel` — send/receive blocks forever; useful in select to disable a case
- `select` — wait on multiple channel operations; picks one that's ready; random if multiple ready
- `default case in select` — non-blocking; execute if no channel is ready
- `done channel` — `chan struct{}`; signal goroutine to stop (zero-size, no allocation)
- `fan-out` — one goroutine sends to many goroutines
- `fan-in` — many goroutines send to one channel
- `pipeline` — chain of goroutines connected by channels; each stage processes and passes forward
- `semaphore with channel` — `make(chan struct{}, N)` to limit N concurrent goroutines
- `ticker` — `time.NewTicker(duration)` sends on channel at regular intervals
- `timer` — `time.NewTimer(duration)` sends once after duration

**Must understand:**
- [ ] Unbuffered vs buffered — when each is appropriate
- [ ] What happens when you send to closed channel (panic) vs receive from closed (zero value, false)
- [ ] Select statement — how it works when multiple cases are ready (random choice)
- [ ] nil channel behavior in select — how to use it to dynamically disable a case
- [ ] The "done channel" pattern for cancellation (before context existed; context is preferred now)

---

## Context Package

**Key Terms:**
- `context.Context` — carries deadlines, cancellation signals, and request-scoped values
- `context.Background()` — root context; use at program start or in main
- `context.TODO()` — placeholder; means "I'll add a real context later"
- `context.WithCancel(parent)` — returns child + cancel function; call cancel() to signal done
- `context.WithTimeout(parent, duration)` — auto-cancels after duration
- `context.WithDeadline(parent, time)` — auto-cancels at absolute time
- `context.WithValue(parent, key, val)` — attach value; use for request-scoped data (trace ID, user ID)
- `ctx.Done()` — channel closed when context cancelled; select on this to handle cancellation
- `ctx.Err()` — returns context.Canceled or context.DeadlineExceeded
- `ctx.Deadline()` — returns deadline time and whether one is set
- `cancel()` — must always be called to release resources (even if ctx times out); use defer
- `context propagation` — pass ctx as first arg to every function that does I/O
- `context key` — use unexported type as key to avoid collisions between packages

**Pattern:**
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()  // always defer cancel

result, err := doSomething(ctx)
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) { /* timeout */ }
}
```

**Must understand:**
- [ ] Why context is the first parameter in Go by convention
- [ ] Always defer cancel() — why (prevents resource leak in context tree)
- [ ] How to propagate cancellation through a chain of goroutines
- [ ] context.WithValue — why to use custom key type (avoid collision between packages)
- [ ] How HTTP server uses context per request (`r.Context()`)

---

## Interfaces

**Key Terms:**
- `interface` — set of method signatures; any type implementing all methods satisfies interface
- `implicit implementation` — no `implements` keyword; just implement the methods
- `duck typing` — "if it walks like a duck and quacks like a duck, it's a duck"
- `empty interface (any)` — `interface{}` or `any`; satisfied by every type; use sparingly
- `type assertion` — `v, ok := x.(T)` — check if x holds value of type T
- `type switch` — switch on dynamic type of interface value
- `nil interface vs nil value` — interface is nil only when both type and value are nil; subtle bug
- `interface composition` — embed interfaces to combine them (`type ReadWriter interface { Reader; Writer }`)
- `io.Reader` — most common interface; `Read(p []byte) (n int, err error)`
- `io.Writer` — `Write(p []byte) (n int, err error)`
- `io.Closer` — `Close() error`
- `error interface` — `Error() string`; every error type implements this
- `Stringer interface` — `String() string`; `fmt.Println` calls this if implemented
- `interface overhead` — method call through interface is indirect (vtable lookup); slightly slower
- `interface boxing` — wrapping a concrete value in an interface allocates on heap (escape to heap)

**Nil interface gotcha:**
```go
var p *MyError = nil
var i error = p   // i is NOT nil! (has type *MyError, nil value)
```

**Must understand:**
- [ ] Why Go interfaces are implicit (enables retroactive implementation)
- [ ] The nil interface gotcha — why returning a concrete nil pointer as interface is a bug
- [ ] Type assertion — two-value form `v, ok := x.(T)` vs one-value (panics if wrong type)
- [ ] When to use `any` (empty interface) and the risks (lose type safety)

---

## Error Handling

**Key Terms:**
- `error interface` — built-in interface: `type error interface { Error() string }`
- `errors.New("msg")` — create simple error
- `fmt.Errorf("msg: %w", err)` — create error wrapping another error
- `%w verb` — wrap error; enables errors.Is and errors.As to unwrap the chain
- `errors.Is(err, target)` — checks if err matches target anywhere in chain (unwraps)
- `errors.As(err, &target)` — checks if any error in chain is of type target; assigns it
- `error wrapping chain` — each layer adds context: `"database error: connection refused"`
- `sentinel error` — predefined error value; compare with ==; `var ErrNotFound = errors.New("not found")`
- `custom error type` — struct implementing error interface; carries extra data (status code, field name)
- `panic` — reserved for unrecoverable errors (nil pointer, out of bounds); not for business logic
- `recover()` — catch a panic in a deferred function; returns panic value
- `never ignore errors` — always check `if err != nil`; common Go antipattern to ignore with `_`
- `multiple return values` — Go idiom: `func Foo() (Result, error)`; no exceptions

**Pattern:**
```go
result, err := doSomething()
if err != nil {
    return fmt.Errorf("doSomething failed: %w", err)  // wrap with context
}
```

**Must understand:**
- [ ] Why Go uses multiple return values instead of exceptions
- [ ] errors.Is vs errors.As — is (exact match or chain), as (type in chain)
- [ ] When to use sentinel errors vs custom types
- [ ] panic/recover — only use to protect library code from crashing the caller

---

## sync Package

**Key Terms:**
- `sync.Mutex` — mutual exclusion lock; `Lock()` / `Unlock()`
- `sync.RWMutex` — `RLock()` / `RUnlock()` for reads; `Lock()` / `Unlock()` for writes
- `sync.WaitGroup` — `Add(n)` before goroutines; `Done()` inside goroutine; `Wait()` to block until all done
- `sync.Once` — `Do(f)` executes f exactly once regardless of how many goroutines call it; lazy init
- `sync.Map` — concurrent-safe map; `Load`, `Store`, `Delete`, `Range`, `LoadOrStore`
- `sync.Cond` — condition variable; `Wait()`, `Signal()` (wake one), `Broadcast()` (wake all)
- `sync.Pool` — pool of reusable objects; reduces GC pressure; `Get()` / `Put()`;  not for long-lived objects
- `atomic` — `sync/atomic` package; lock-free operations on primitive types; `AddInt64`, `LoadInt64`, `StoreInt64`, `CompareAndSwapInt64`
- `CAS (Compare-And-Swap)` — atomic operation: set value only if current value equals expected; basis for lock-free data structures

**Must understand:**
- [ ] WaitGroup — Add must be called before goroutine starts (not inside)
- [ ] sync.Once — singleton initialization pattern
- [ ] sync.Pool — reduces allocation pressure; objects may be GC'd at any time
- [ ] When atomic operations are better than mutex (simple counters, flags)

---

## defer / panic / recover

**Key Terms:**
- `defer` — schedules function call to run when surrounding function returns (or panics)
- `LIFO order` — multiple defers execute in reverse order (last defer = first to run)
- `defer evaluates args immediately` — `defer fmt.Println(x)` captures x's value NOW, not when it runs
- `defer with named return` — deferred function can modify named return values
- `panic` — built-in function that stops normal execution, unwinds stack, runs deferred functions
- `panic propagation` — panic unwinds up the call stack until caught by recover or program terminates
- `recover()` — must be called in deferred function; captures panic value; stops unwinding
- `runtime panic` — nil pointer dereference, out-of-bounds index, zero division → automatic panics
- `panic vs error` — panic for programmer errors (bugs); error for expected failures

**Use defer for:**
- Close file/connection: `defer f.Close()`
- Unlock mutex: `defer mu.Unlock()`
- Release resources: `defer cancel()`
- Recover from panic in library code

**Must understand:**
- [ ] LIFO order of defers — critical for nested resource cleanup
- [ ] defer captures args by value — what this means in practice
- [ ] recover must be in a deferred function — why (stack is being unwound)
- [ ] When to use panic (programming error) vs error return (expected failure)

---

## Slices Internals

**Key Terms:**
- `slice header` — 3 fields: pointer to backing array, len, cap
- `backing array` — the underlying fixed-size array; shared between slices
- `len` — number of accessible elements
- `cap` — total capacity from start of slice to end of backing array
- `append` — if len < cap, writes to backing array in-place and returns new slice header; if len == cap, allocates new larger backing array, copies, returns new header
- `slice sharing` — `s2 := s1[2:5]` shares backing array; modifying s2 modifies s1's data
- `full slice expression` — `s[low:high:max]` limits cap of s2, preventing writes beyond high
- `copy(dst, src)` — copies min(len(dst), len(src)) elements; creates independent copy
- `nil slice vs empty slice` — nil slice: header with nil pointer, len=0, cap=0; empty slice: non-nil pointer, len=0, cap=0; both work with append/len/for range
- `append to nil slice` — works fine; `append(nil, 1, 2, 3)` is valid
- `make([]T, len, cap)` — allocate slice with specified len and cap
- `2D slice` — `[][]int`; each inner slice is independent allocation

**Subtle bugs:**
- Appending to sub-slice overwrites original:
  ```go
  a := []int{1,2,3,4,5}
  b := a[1:3]      // shares backing array
  b = append(b, 99) // overwrites a[3]!
  ```
- Fix: use full slice expression `b := a[1:3:3]` to limit cap

**Must understand:**
- [ ] Exactly how append works (in-place vs new allocation)
- [ ] The slice sharing problem and how full slice expression prevents it
- [ ] When to use make([]T, 0, capacity) for pre-allocation
- [ ] nil slice vs empty slice — which JSON marshals to null vs []

---

## Maps Internals

**Key Terms:**
- `hash map` — Go's built-in map; unordered; O(1) average get/put/delete
- `nil map` — read returns zero value (safe); write panics
- `map literal` — `m := map[string]int{"a": 1}`; ready to use
- `make(map[K]V)` — creates initialized empty map
- `delete(m, key)` — remove key; no error if key doesn't exist
- `two-value access` — `v, ok := m[key]`; ok=false if key absent
- `map iteration order` — random; intentionally randomized each run (security)
- `comparable key` — map key must be comparable (==); slices, maps, functions cannot be keys
- `struct as key` — works if all fields are comparable
- `concurrent access` — Go map is NOT safe for concurrent read+write; use RWMutex or sync.Map
- `map growth` — map grows automatically; triggers rehash; brief pause
- `map deletion doesn't shrink` — deleted entries leave holes; map memory doesn't shrink automatically
- `sync.Map` — concurrent-safe; best when key set is stable and reads >> writes
- `JSON marshaling` — `map[string]any{}` for arbitrary JSON; `encoding/json`

**Must understand:**
- [ ] Why concurrent map access panics (Go detects it since Go 1.6 with built-in race detection)
- [ ] Two-value access — always use when key might be absent
- [ ] Why you can't use slice as map key
- [ ] When sync.Map is appropriate vs regular map + RWMutex

---

## Go Scheduler (GMP Model)

**Key Terms:**
- `G (Goroutine)` — the goroutine; contains stack, instruction pointer, state
- `M (Machine)` — OS thread; executes goroutines; created by runtime as needed
- `P (Processor)` — logical processor; holds a run queue of goroutines; M must hold P to execute goroutines
- `GOMAXPROCS` — number of P's; limits parallelism
- `local run queue` — each P has a queue of goroutines ready to run
- `global run queue` — overflow from local queues; all P's can steal from it
- `work stealing` — idle P steals goroutines from another P's local queue; keeps all cores busy
- `preemption` — since Go 1.14, goroutines can be preempted at almost any point (safe points); prevents one goroutine from blocking all others
- `goroutine park` — when goroutine blocks on I/O or channel, it's removed from M and parked; M can run another goroutine
- `goroutine unpark` — when I/O completes, goroutine placed back on run queue
- `syscall handling` — when M makes blocking syscall, P is released; another M picks up P and continues running goroutines
- `spinning M` — M looking for work (burns CPU); Go limits spinning Ms

**GMP diagram:**
```
P1 [run queue: G1, G2, G3] — M1 (executing G1)
P2 [run queue: G4, G5]     — M2 (executing G4)
P3 [run queue: empty]      — M3 (stealing from P1)
Global run queue: G6, G7
```

**Must understand:**
- [ ] Why Go needs P (processor) concept — decouples goroutines from threads
- [ ] Work stealing — how idle P's steal work to keep CPU utilized
- [ ] How blocking syscall is handled without blocking all goroutines
- [ ] GOMAXPROCS and when to change it

---

## Testing in Go

**Key Terms:**
- `testing.T` — test helper; `t.Error`, `t.Fatal`, `t.Log`, `t.Parallel`, `t.Run`
- `testing.B` — benchmark helper; `b.N` is number of iterations; `b.ReportAllocs()`
- `table-driven tests` — slice of test cases; loop calls t.Run for each
- `t.Run(name, func)` — subtest; allows focused run with `-run TestFoo/case1`
- `t.Parallel()` — mark test as parallel; multiple tests run simultaneously
- `testify/assert` — popular library; `assert.Equal`, `assert.Nil`, `assert.Error`
- `testify/require` — like assert but stops test on failure
- `testify/mock` — mock objects for interfaces
- `httptest.NewRecorder()` — fake HTTP ResponseWriter for testing handlers
- `httptest.NewServer()` — creates real HTTP test server
- `goldenfile testing` — compare output to expected file; update with `-update` flag
- `benchmark` — `func BenchmarkFoo(b *testing.B)` + `go test -bench=.`
- `-race flag` — run tests with data race detector
- `coverage` — `go test -cover`; `go test -coverprofile=cover.out`
- `testmain` — `func TestMain(m *testing.M)` — setup/teardown for entire package
- `integration tests` — use build tags to separate: `//go:build integration`

**Table-driven test pattern:**
```go
func TestAdd(t *testing.T) {
    cases := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 1, 2, 3},
        {"negative", -1, -2, -3},
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            got := Add(tc.a, tc.b)
            if got != tc.want {
                t.Errorf("got %d, want %d", got, tc.want)
            }
        })
    }
}
```

**Must understand:**
- [ ] Table-driven tests — why they're idiomatic Go
- [ ] t.Fatal vs t.Error — fatal stops test immediately
- [ ] Running specific tests: `go test -run TestFoo/subtest_name`
- [ ] Race detector in tests: `go test -race`

---

## HTTP Server in Go

**Key Terms:**
- `net/http` — standard library HTTP server and client
- `http.Handler` — interface: `ServeHTTP(ResponseWriter, *Request)`
- `http.HandlerFunc` — adapter to use a function as http.Handler
- `http.ServeMux` — default router; pattern matching; no path parameters
- `gorilla/mux` / `chi` / `gin` — third-party routers with path params, middleware
- `middleware` — function that wraps Handler; `func(http.Handler) http.Handler`
- `http.ListenAndServe` — starts server; blocks; returns only on error
- `http.Server` — configurable server struct (timeouts, TLS)
- `ReadTimeout` — max time to read request headers + body
- `WriteTimeout` — max time to write response
- `IdleTimeout` — max time between requests on keep-alive connection
- `graceful shutdown` — `server.Shutdown(ctx)` — drains active connections
- `r.Context()` — request context; cancelled when client disconnects
- `http.ResponseWriter` — write status + headers + body; Header() must be set before WriteHeader()
- `json.NewDecoder(r.Body).Decode(&v)` — decode JSON request body
- `json.NewEncoder(w).Encode(v)` — encode JSON response

**Middleware pattern:**
```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}
```

**Must understand:**
- [ ] Why `http.Handler` is an interface (enables middleware composition)
- [ ] Request context — how to use r.Context() for cancellation
- [ ] Always set headers before WriteHeader (can't change after)
- [ ] Graceful shutdown — how to drain connections properly

---

## Common Go Patterns

**Key Terms:**
- `functional options` — `func(c *Config)` type; `NewServer(WithTimeout(5*time.Second))`; extensible, backwards-compatible
- `options struct` — `ServerOptions{Timeout: 5 * time.Second}`; simple, requires all fields to be public
- `worker pool` — fixed number of goroutines consuming from a job channel; limits concurrency
- `pipeline` — chain of channels; each stage reads from input, processes, sends to output
- `fan-out / fan-in` — distribute work to multiple workers, collect results back
- `errgroup` — `golang.org/x/sync/errgroup`; like WaitGroup but propagates first error
- `rate limiter` — `golang.org/x/time/rate`; token bucket implementation
- `retry with backoff` — retry failed operations with exponential backoff + jitter
- `circuit breaker` — stop calling failing service; `sony/gobreaker`
- `singleton` — `sync.Once` for lazy, thread-safe initialization
- `builder pattern` — chain methods to construct complex object

**Functional options pattern:**
```go
type Server struct { timeout time.Duration }
type Option func(*Server)
func WithTimeout(d time.Duration) Option { return func(s *Server) { s.timeout = d } }
func NewServer(opts ...Option) *Server {
    s := &Server{timeout: 30 * time.Second} // defaults
    for _, o := range opts { o(s) }
    return s
}
```

**Worker pool pattern:**
```go
jobs := make(chan Job, 100)
for i := 0; i < numWorkers; i++ {
    go func() {
        for job := range jobs { process(job) }
    }()
}
```

**Must understand:**
- [ ] Functional options — why preferred over large constructor signatures
- [ ] errgroup — when to use instead of WaitGroup (when you need to propagate errors)
- [ ] Worker pool — how to size the pool (CPU-bound: numCPU; I/O-bound: higher)
- [ ] Pipeline pattern — how to use channels to build data transformation stages
