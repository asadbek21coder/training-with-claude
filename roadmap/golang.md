# Go Roadmap

> **L1** = aware & can use · **L2** = deep & efficient · **L3** = internals, book depth (📖 *The Anatomy of Go*, Phuong Le — page numbers = PDF pages) · Status: ⬜ not started · 🔄 in progress · ✅ done
> **Notes:** `notes/golang/` · **Practice:** `code/golang/` · **Dictionary:** `reference/golang.md`

**Why (EN):** Go is your main language and you teach it — it must be your deepest skill. In the AI era, syntax matters less (AI writes it), but concurrency correctness, runtime knowledge, idioms and the ability to review and profile Go code matter more: AI confidently writes goroutine leaks, races and non-idiomatic code.
**Зачем (RU):** Go — твой основной язык, и ты его преподаёшь, поэтому это должен быть самый глубокий навык. В эпоху ИИ синтаксис менее важен (его пишет ИИ), но корректная конкурентность, знание рантайма, идиомы и умение ревьюить и профилировать Go-код важнее: ИИ уверенно пишет утечки горутин, гонки и неидиоматичный код.

**Order:** topics 1–14 = the course core (teach these to your student); 15–23 = backend & senior depth. **Current goal: all L1 first**, then L2, then L3 (the book). Topic 24 = the compiler (book Ch5).

## Progress
| Layer | Done | Total |
|---|---|---|
| L1 | 1 | 85 |
| L2 | 0 | 72 |
| L3 (📖 book) | 0 | 33 |

---

## GO 1 — Philosophy & toolchain
**L1 — aware & can use**
- ✅ **1.1 Why Go exists** — simplicity, composition over inheritance, fast compilation, static typing, garbage collection, CSP, Go proverbs, Go 1 compatibility promise, release cycle (Feb/Aug)
- ⬜ **1.2 Installing & environment** — `GOROOT`, `GOPATH`, `GOBIN`, `GOOS`, `GOARCH`, `go env`, `go version`, `PATH`
- ⬜ **1.3 The `go` command** — `go run`, `go build`, `go install`, `go test`, `go mod`, `go get`, `go fmt`, `go vet`, `go doc`, `go generate`, `go clean`, `go list`, `go tool`
- ⬜ **1.4 Program structure** — `package main`, `func main`, executable vs library package, `gofmt`, Go playground

**L2 — deep & efficient**
- ⬜ **1.5 Toolchain management** — `GOTOOLCHAIN`, `toolchain` directive, `go env -w`, build cache (`GOCACHE`), module cache (`GOMODCACHE`), `GODEBUG` settings and compatibility, `go version -m binary`

**L3 — internals (📖 Anatomy of Go)**
- ⬜ **1.6 How Go runs your code** — compiler + runtime as one system, what is in a Go binary, why Go has a runtime at all (Ch1 §2)

📖 **Anatomy of Go:** Ch1 (p.16–28)

🛠 **Practice** — L1: install Go, create a module, write `hello` with a flag, build binaries for linux/amd64 and darwin/arm64 · L2: inspect a built binary with `go version -m`, clear and measure the build cache effect on build time

---

## GO 2 — Syntax fundamentals
**L1 — aware & can use**
- ⬜ **2.1 Packages & imports** — exported vs unexported (capital letter), import alias, blank import `_`, dot import, `init()`, package-level variables
- ⬜ **2.2 Variables & scope** — `var`, short declaration `:=`, zero value, shadowing, block scope, package scope, unused variable error, multiple assignment, blank identifier
- ⬜ **2.3 Constants & iota** — `const`, typed vs untyped constants, constant expressions, `iota`, enum pattern
- ⬜ **2.4 Basic types** — `int`/`int8..int64`, `uint`, `uintptr`, `float32`/`float64`, `complex128`, `bool`, `byte` (= uint8), `rune` (= int32), `string`, platform-dependent `int` size
- ⬜ **2.5 Type conversion & operators** — explicit conversion `T(x)`, no implicit conversion, integer division, `%`, bitwise `& | ^ &^ << >>`, operator precedence, `strconv` vs conversion
- ⬜ **2.6 Control flow** — `if` with init statement, `for` (3 forms), `range` (slice, map, string, channel, int, func), `switch` (no fallthrough by default, `fallthrough`, tagless switch), `break`/`continue` with labels, `goto`
- ⬜ **2.7 Functions** — multiple return values, named results, naked return, variadic `...`, functions as values, function types, anonymous functions, closures, recursion
- ⬜ **2.8 Pointers** — `&` address-of, `*` dereference, `new(T)`, nil pointer dereference panic, no pointer arithmetic, when to use pointers, auto-dereference for struct fields

**L2 — deep & efficient**
- ⬜ **2.9 Package initialization order** — dependency order of package vars, multiple `init()` per file/package, import graph order, why `init` is discouraged
- ⬜ **2.10 Closures in depth** — captured variables by reference, per-iteration loop variables (Go 1.22), closures and escape to heap, `new(expr)` (Go 1.26)
- ⬜ **2.11 Numeric edge cases** — integer overflow wraparound, untyped constant precision, float precision (`0.1+0.2`), `math.MaxInt`, NaN and Inf, division by zero (int panics, float does not)

**L3 — internals (📖 Anatomy of Go)**
- ⬜ **2.12 Number representation** — two's complement, integer literals (`0b`, `0o`, `0x`, `_`), IEEE-754 bit layout, float literal styles, compile-time vs runtime division-by-zero checks (Ch2 §1–2)
- ⬜ **2.13 Variables & constants inside the compiler** — scope resolution, how declarations are type-checked, untyped constants with arbitrary precision, constant folding, where constants live in the binary (Ch2 §5)
- ⬜ **2.14 unsafe & addressability** — `unsafe.Pointer` vs `uintptr`, the valid conversion patterns, why `uintptr` is not a reference for the GC, addressable vs unaddressable values (Ch2 §6–7)
- ⬜ **2.15 Functions under the hood** — function values as `funcval`, how closures capture variables, variadic args as hidden slices, how `init` order is computed (Ch6 §2)

📖 **Anatomy of Go:** Ch2 (p.29–85) · Ch6 §2 (p.417–436)

🛠 **Practice** — L1: CLI "unit converter" using constants with `iota`, `switch`, functions with multiple returns · L2: write 5 tiny programs that show surprising behavior (shadowing bug, closure capture, int overflow, float compare, init order) and explain each to your student

---

## GO 3 — Composite types
**L1 — aware & can use**
- ⬜ **3.1 Arrays** — fixed length is part of the type, value semantics (copy on assign), `[...]T{}`, comparing arrays, arrays as map keys
- ⬜ **3.2 Slices — usage** — slice literal, `make([]T, len, cap)`, `len`, `cap`, `append`, `copy`, slicing `s[a:b]`, nil vs empty slice, 2D slices, `slices` package (`Sort`, `Contains`, `Index`, `Clone`, `Compact`)
- ⬜ **3.3 Maps — usage** — map literal, `make(map[K]V, hint)`, comma-ok idiom, `delete`, `clear`, random iteration order, nil map (read OK, write panics), comparable keys, `maps` package
- ⬜ **3.4 Strings, bytes, runes** — immutable strings, UTF-8, byte vs rune, `len` counts bytes, `range` yields runes, `strings.Builder`, `strings`/`bytes`/`strconv`/`unicode/utf8` packages
- ⬜ **3.5 Structs** — struct literal, field tags, anonymous struct, embedded struct (field promotion), struct comparison, empty struct `struct{}`, constructor function `NewT`

**L2 — deep & efficient**
- ⬜ **3.6 Slices — internals** — slice header (pointer, len, cap), backing array, growth strategy, amortized O(1) append, aliasing trap, full slice expression `s[a:b:c]`, memory leak by reslicing, passing slices to functions
- ⬜ **3.7 Maps — internals** — hash function, Swiss tables (Go 1.24: groups, control bytes, probing), load factor, growth, map of structs is not addressable, why maps are not concurrency-safe, `concurrent map writes` fatal error
- ⬜ **3.8 Strings — internals & cost** — string header, `[]byte(s)` / `string(b)` copy cost, concatenation cost, compiler optimizations for conversions, `unsafe.String` (awareness)
- ⬜ **3.9 Struct memory layout** — field alignment & padding, field ordering to save memory, `unsafe.Sizeof`, zero-size `struct{}` uses (sets, signals), copying large structs

**L3 — internals (📖 Anatomy of Go)**
- ⬜ **3.10 Array initialization strategies** — small vs large array literals, static data vs runtime init, array length vs capacity (Ch3 §1)
- ⬜ **3.11 append & allocation in depth** — `growslice`, growth formula, rounding up to size classes, `makeslice`, slice expressions and bounds checks (Ch3 §2)
- ⬜ **3.12 String conversion internals** — UTF-8 decoding in `range`, when `[]byte`↔`string` conversions skip allocation, `staticuint64s`, `concatstrings` and efficient concatenation (Ch3 §3)
- ⬜ **3.13 Swiss map in depth** — directory, tables, groups, control bytes, H1/H2 hash split, probing, growth and splitting, iteration while modifying (Ch3 §4)

📖 **Anatomy of Go:** Ch3 (p.86–172)

🛠 **Practice** — L1: word-frequency counter (map + slices + sort), a `Student` struct with methods for your lesson · L2: write a program that proves the slice aliasing trap and fix it; measure struct size before/after field reordering

---

## GO 4 — Methods & interfaces
**L1 — aware & can use**
- ⬜ **4.1 Methods** — receiver, value receiver vs pointer receiver, methods on non-struct types, method values, consistency rule (don't mix receivers)
- ⬜ **4.2 Interfaces — basics** — interface type, implicit implementation, `any` / `interface{}`, interface value = (type, value), dynamic type
- ⬜ **4.3 Type assertion & type switch** — `x.(T)`, comma-ok assertion, panic on failed assertion, `switch v := x.(type)`
- ⬜ **4.4 Key stdlib interfaces** — `fmt.Stringer`, `error`, `io.Reader`, `io.Writer`, `io.Closer`, `sort.Interface`, `http.Handler`, `json.Marshaler`

**L2 — deep & efficient**
- ⬜ **4.5 The nil interface trap** — nil interface vs interface holding a nil pointer, returning a typed nil error, why `err != nil` is true
- ⬜ **4.6 Method sets & addressability** — method set of `T` vs `*T`, which types satisfy an interface, auto address-taking rules, method expressions
- ⬜ **4.7 Interface design** — small interfaces, "accept interfaces, return structs", consumer-side interfaces, interface pollution, compile-time check `var _ I = (*T)(nil)`
- ⬜ **4.8 Embedding in depth** — interface embedding, embedding an interface in a struct, promotion conflicts, embedding vs inheritance

**L3 — internals (📖 Anatomy of Go)**
- ⬜ **4.9 Interface runtime internals** — `eface` vs `iface`, `itab` (`inter`, `_type`, `hash`, `fun`), how `data` is filled (pointer used directly, `zerobase`, `staticuint64s`, read-only data, stack or heap copy), dispatch through `itab.fun`, compiler-generated pointer-receiver wrappers, `itabTable` + `getitab`, type-assertion cache (Ch4 §2.3–2.5)

📖 **Anatomy of Go:** Ch4 §1–2 (p.174–233)

🛠 **Practice** — L1: `Shape` interface with `Area()` for 3 types + type switch printer · L2: write a bug where a function returns a typed nil error, find it, fix it; design a `Storage` interface defined by its consumer

---

## GO 5 — Errors, defer, panic
**L1 — aware & can use**
- ⬜ **5.1 Error basics** — `error` interface, `errors.New`, `fmt.Errorf`, errors are values, `if err != nil`, early return, error string style
- ⬜ **5.2 Wrapping & inspection** — `%w`, error chain, `errors.Unwrap`, `errors.Is`, `errors.As`, `errors.Join`
- ⬜ **5.3 defer** — LIFO order, `defer f.Close()`, cleanup patterns, `defer mu.Unlock()`
- ⬜ **5.4 panic & recover** — `panic`, `recover` (only in deferred func), runtime panics (nil map write, index out of range, nil deref), panic vs error

**L2 — deep & efficient**
- ⬜ **5.5 Error types & strategies** — sentinel errors (`io.EOF`, `sql.ErrNoRows`), custom error types, opaque errors, behavior checking (`interface{ Timeout() bool }`), handle once (log OR return), adding context, `errors.AsType` (Go 1.26), mapping domain errors to HTTP codes
- ⬜ **5.6 defer in depth** — arguments evaluated immediately, defer modifying named results, defer in loops (resource leak), handling `Close()` errors, open-coded defers (cost)
- ⬜ **5.7 Panics in production** — panics in goroutines crash the whole program, recovery middleware, when panic is acceptable (programmer bug, init), `os.Exit` skips defers

**L3 — internals (📖 Anatomy of Go)**
- ⬜ **5.8 defer implementation** — `_defer` record, per-goroutine linked list, `deferproc`/`deferreturn`, heap vs stack vs open-coded defers and their rules, per-P defer pools (Ch6 §3)
- ⬜ **5.9 panic & recover internals** — `_panic` struct, how the runtime unwinds and runs defers, why `recover` only works directly in a deferred call (argument pointer check) (Ch6 §4)

📖 **Anatomy of Go:** Ch6 §3–4 (p.436–471) — errors/wrapping are not in the book; use Go blog + course plans

🛠 **Practice** — L1: file-reading function that wraps errors with context; check them with `errors.Is/As` · L2: define a domain error type for a "user service" and map it to HTTP codes; write a recovery middleware and a test that proves it works

---

## GO 6 — Generics
**L1 — aware & can use**
- ⬜ **6.1 Type parameters** — `[T any]`, generic function, generic type, instantiation, type inference
- ⬜ **6.2 Constraints basics** — `any`, `comparable`, `cmp.Ordered`, union `int | float64`
- ⬜ **6.3 Generic stdlib** — `slices`, `maps`, `cmp` packages (`slices.SortFunc`, `slices.BinarySearch`, `maps.Keys`, `cmp.Compare`)

**L2 — deep & efficient**
- ⬜ **6.4 Type sets & generic types** — `~T` (underlying type), constraint interfaces with methods, generic data structures (`Stack[T]`, `Set[T]`), generic type aliases (Go 1.24)
- ⬜ **6.5 Limits & trade-offs** — no method type parameters, no specialization, interface vs generic choice, readability cost, GC shape stenciling and performance

**L3 — internals (📖 Anatomy of Go)**
- ⬜ **6.6 Generics compiler internals** — type inference steps (function argument, constraint, core type), GC shape stenciling, dictionaries, cost of calling methods through type parameters (Ch4 §3.3–3.4)

📖 **Anatomy of Go:** Ch4 §3 (p.233–256)

🛠 **Practice** — L1: generic `Map`, `Filter`, `Reduce` functions · L2: generic `Set[T comparable]` and `Cache[K, V]` with tests; benchmark a generic vs interface-based version

---

## GO 7 — Packages, modules, project structure
**L1 — aware & can use**
- ⬜ **7.1 Modules basics** — `go.mod` (`module`, `go`, `require`), `go.sum`, `go mod init`, `go mod tidy`, `go get pkg@version`, indirect dependencies
- ⬜ **7.2 Project layout** — `cmd/`, `internal/` (compiler-enforced), `pkg/` (controversial), package naming (short, no `utils`/`common`), avoiding import cycles
- ⬜ **7.3 Documentation** — doc comments, `go doc`, pkg.go.dev, example functions, README

**L2 — deep & efficient**
- ⬜ **7.4 Versioning** — semantic versioning, `v0`/`v1` stability, major version suffix `/v2`, Minimal Version Selection (MVS), pseudo-versions, `retract`
- ⬜ **7.5 Module proxy & private modules** — `GOPROXY`, `GOSUMDB`, checksum database, `GOPRIVATE`, `GONOSUMDB`, private GitLab modules
- ⬜ **7.6 Advanced dependency management** — `replace`, `exclude`, `tool` directive (Go 1.24), workspaces `go.work`, vendoring `go mod vendor`, upgrading safely (`go list -m -u all`)

🛠 **Practice** — L1: create a module with `cmd/app` + `internal/` packages and one external dependency · L2: publish a tiny module to GitHub with a `v2` major version; use it from another module via `go.work`

---

## GO 8 — Concurrency basics
**L1 — aware & can use**
- ⬜ **8.1 Concurrency vs parallelism** — concurrency (structure), parallelism (execution), CSP, "share memory by communicating"
- ⬜ **8.2 Goroutines** — `go` statement, goroutine vs OS thread, main goroutine exit kills all, `runtime.NumGoroutine`
- ⬜ **8.3 WaitGroup** — `sync.WaitGroup`, `Add`/`Done`/`Wait`, `wg.Go` (Go 1.25), Add-before-go rule
- ⬜ **8.4 Channels** — `make(chan T)`, unbuffered (synchronous handoff), buffered (queue), send, receive, `close`, `range` over channel, directional types `chan<-` / `<-chan`
- ⬜ **8.5 select** — `select`, `default` (non-blocking), timeout with `time.After`, for-select loop, `ctx.Done()` case
- ⬜ **8.6 Mutex** — `sync.Mutex`, `sync.RWMutex`, critical section, `Lock`/`Unlock`, `defer mu.Unlock()`

**L2 — deep & efficient**
- ⬜ **8.7 Channel rules & closing patterns** — receive from closed channel (zero value, ok=false), send on closed panics, close of closed/nil panics, only the sender closes, nil channel blocks forever (disabling select cases), "all goroutines are asleep — deadlock!"
- ⬜ **8.8 Once & atomics** — `sync.Once`, `sync.OnceFunc`/`OnceValue`/`OnceValues`, `sync/atomic`, `atomic.Int64`, `atomic.Pointer[T]`, compare-and-swap (CAS)
- ⬜ **8.9 Choosing the tool** — mutex vs channel decision, RWMutex costs (writer starvation, when it's slower), don't copy a mutex (`go vet` copylocks), lock granularity

📖 **Anatomy of Go:** Ch8 §1–2 (p.739–766)

🛠 **Practice** — L1: concurrent URL status checker with WaitGroup + results channel · L2: safe counter in 3 versions (mutex, atomic, channel) + benchmark each; reproduce and explain a deadlock

---

## GO 9 — Concurrency patterns & pitfalls
**L1 — aware & can use**
- ⬜ **9.1 Worker pool** — jobs channel, results channel, fixed number of workers, bounded concurrency
- ⬜ **9.2 errgroup** — `golang.org/x/sync/errgroup`, `errgroup.WithContext`, `SetLimit`, first error cancels the rest
- ⬜ **9.3 Goroutine leaks** — blocked sender/receiver, forgotten goroutine, missing cancellation, detection with `runtime.NumGoroutine`
- ⬜ **9.4 Data races** — data race vs race condition, `go test -race`, `go run -race`, check-then-act bug

**L2 — deep & efficient**
- ⬜ **9.5 Pipelines & fan-out/fan-in** — stage, generator, fan-out, fan-in (merge), cancelling a pipeline, backpressure with buffered channels
- ⬜ **9.6 Semaphores & rate limiting** — semaphore via buffered channel, `x/sync/semaphore`, `time.Ticker`, `x/time/rate` (token bucket), debounce, throttling
- ⬜ **9.7 Go memory model** — happens-before, synchronization points (channel ops, mutex, `sync.Once`, atomics), why "benign races" are bugs
- ⬜ **9.8 Advanced sync tools** — `sync.Cond`, `sync.Pool`, `sync.Map` (when it helps), `x/sync/singleflight`, deadlock, livelock, starvation, lock ordering
- ⬜ **9.9 Testing concurrent code** — `testing/synctest` (Go 1.25), fake clock, `goleak`, deterministic tests, `-count=100 -race` in CI

🛠 **Practice** — L1: worker pool that resizes images (or hashes files) with N workers; then write a leaking goroutine and fix it with context · L2: 3-stage pipeline with cancellation + `goleak` test; `singleflight` cache loader with a test proving one DB call for 100 concurrent requests

---

## GO 10 — context
**L1 — aware & can use**
- ⬜ **10.1 Why context** — cancellation propagation, request-scoped deadlines, context tree
- ⬜ **10.2 Creating contexts** — `context.Background`, `context.TODO`, `WithCancel`, `WithTimeout`, `WithDeadline`, `CancelFunc` (always call it)
- ⬜ **10.3 Using contexts** — `ctx.Done()`, `ctx.Err()` (`Canceled`, `DeadlineExceeded`), ctx as first parameter, `r.Context()` in HTTP, `QueryContext` in DB

**L2 — deep & efficient**
- ⬜ **10.4 Context values & rules** — `WithValue`, unexported key type, request ID / user ID, what NOT to put in context, don't store ctx in structs
- ⬜ **10.5 Newer APIs** — `WithCancelCause`, `context.Cause`, `WithTimeoutCause`, `AfterFunc`, `WithoutCancel`, `signal.NotifyContext`
- ⬜ **10.6 Context in systems** — timeout budgets across service calls, deadline propagation over gRPC/HTTP, cancellation of background work after response

🛠 **Practice** — L1: HTTP handler that calls a slow function and stops when the client disconnects or after 2s · L2: request-ID middleware using context values + logging; chain of 3 calls sharing one deadline budget

---

## GO 11 — Standard library essentials
**L1 — aware & can use**
- ⬜ **11.1 fmt** — verbs `%v %+v %#v %T %d %s %q %x %p`, `Sprintf`, `Fprintf`, `Errorf`
- ⬜ **11.2 io & bufio** — `io.Reader`/`Writer`, `io.Copy`, `io.ReadAll`, `bufio.Scanner`, `bufio.Writer` + `Flush`
- ⬜ **11.3 os & filepath** — `os.Open`/`Create`/`ReadFile`/`WriteFile`, permissions, `os.Getenv`/`LookupEnv`, `os.Args`, `os.Exit`, `path/filepath`
- ⬜ **11.4 time** — `time.Time`, `time.Duration`, `Now`/`Since`/`Until`, layout `2006-01-02T15:04:05Z07:00`, `time.Sleep`, `Timer`, `Ticker`
- ⬜ **11.5 encoding/json** — `Marshal`/`Unmarshal`, struct tags, `omitempty`, `omitzero` (Go 1.24), `json.Decoder`/`Encoder`
- ⬜ **11.6 log/slog** — `slog.Info`, levels, attributes, `JSONHandler`/`TextHandler`, `logger.With`
- ⬜ **11.7 Everyday packages** — `strconv`, `strings`, `sort`/`slices`, `regexp`, `flag`, `math/rand/v2` vs `crypto/rand`, `os/exec`, `net/url`

**L2 — deep & efficient**
- ⬜ **11.8 io composition & streaming** — `io.MultiWriter`, `io.TeeReader`, `io.Pipe`, `io.LimitReader`, `io.SectionReader`, streaming large files without loading into memory, `io/fs`, `os.Root` (Go 1.24), `embed`
- ⬜ **11.9 JSON in depth** — custom `MarshalJSON`/`UnmarshalJSON`, `json.RawMessage`, `DisallowUnknownFields`, numbers into `any` → float64 (`UseNumber`), case-insensitive field matching, performance, `encoding/json/v2` (experiment)
- ⬜ **11.10 time pitfalls** — monotonic clock, time zones & `time.Location`, storing UTC, `time.After` in loops, stopping timers/tickers, comparing with `Equal`
- ⬜ **11.11 Modern stdlib** — `iter` package & range-over-func iterators (Go 1.23), `slices.Collect`/`maps.Keys` iterators, `unique` (interning), `weak` (Go 1.24), `runtime.AddCleanup`

**L3 — internals (📖 Anatomy of Go)**
- ⬜ **11.12 Timer internals** — timer struct and states, per-P timer heap, Go 1.23 timer changes (unbuffered timer channels, unreferenced timers collected by GC) (Ch8 §3)

📖 **Anatomy of Go:** Timers: Ch8 §3 (p.766–781)

🛠 **Practice** — L1: CLI that reads a CSV, converts it to JSON, logs with slog, takes flags · L2: stream a 1 GB file line by line with constant memory; write your own iterator `func(yield func(T) bool)` over a paginated API

---

## GO 12 — Testing in Go
**L1 — aware & can use**
- ⬜ **12.1 Basics** — `_test.go`, `TestXxx(t *testing.T)`, `t.Error` vs `t.Fatal`, `go test ./...`, `-run`, `-v`, `-count=1`, test cache
- ⬜ **12.2 Table-driven tests & subtests** — test table, `t.Run`, `t.Helper`, `t.Cleanup`, `t.TempDir`
- ⬜ **12.3 HTTP testing** — `httptest.NewRecorder`, `httptest.NewServer`, testing handlers and clients
- ⬜ **12.4 Assertions** — stdlib comparisons, `testify` (`assert` vs `require`), `go-cmp` `cmp.Diff`

**L2 — deep & efficient**
- ⬜ **12.5 Test doubles & testability** — interface-based fakes, stubs, mocks, spies, `mockery`/`gomock`, why fakes often beat mocks, designing code for testability
- ⬜ **12.6 Integration tests** — testcontainers-go (Postgres, Kafka, Redis), `TestMain`, build tags for integration tests, transaction rollback per test, `t.Parallel` pitfalls
- ⬜ **12.7 Coverage, benchmarks, fuzzing, examples** — `-cover`, `-coverprofile`, `go tool cover -html`, `BenchmarkXxx`, `b.Loop` (Go 1.24), `-benchmem`, `FuzzXxx`, corpus, `ExampleXxx` with `// Output:`
- ⬜ **12.8 Test organization** — golden files, `testdata/`, `t.Context` (Go 1.24), `t.Setenv`, black-box `package x_test` vs white-box tests, flaky tests

🛠 **Practice** — L1: table-driven tests for a `ParseDuration`-like function + a handler test with `httptest` · L2: integration test of a repository with testcontainers Postgres; fuzz test that finds a bug in a parser you wrote

---

## GO 13 — HTTP backend with net/http
**L1 — aware & can use**
- ⬜ **13.1 Server basics** — `http.Handler`, `http.HandlerFunc`, `ServeMux`, `ListenAndServe`, `ResponseWriter`, `*http.Request`
- ⬜ **13.2 Routing** — Go 1.22 patterns (`"GET /users/{id}"`), `r.PathValue`, wildcards `{path...}`, precedence, routers (chi, gin, echo)
- ⬜ **13.3 Requests & responses** — reading JSON body, query params, headers, status codes, `WriteHeader` order, `Content-Type`
- ⬜ **13.4 Middleware** — `func(http.Handler) http.Handler`, chaining, logging, recovery, request ID, CORS
- ⬜ **13.5 HTTP client basics** — `http.Get`, `http.NewRequestWithContext`, `client.Do`, `resp.Body.Close()`, reading JSON responses

**L2 — deep & efficient**
- ⬜ **13.6 Server hardening** — `http.Server{}` with `ReadHeaderTimeout`, `ReadTimeout`, `WriteTimeout`, `IdleTimeout`, `MaxHeaderBytes`, `http.MaxBytesReader`, why the default server is dangerous
- ⬜ **13.7 Graceful shutdown** — `server.Shutdown(ctx)`, `signal.NotifyContext`, SIGTERM/SIGINT, draining connections, closing DB/workers in the right order
- ⬜ **13.8 Client in production** — reuse `http.Client`, `Timeout`, `http.Transport` (`MaxIdleConnsPerHost`, `IdleConnTimeout`), connection pooling & keep-alive, draining bodies, retries with backoff + jitter, idempotency
- ⬜ **13.9 Common backend features** — validation, JWT auth middleware, file upload (multipart, streaming), pagination, rate-limit middleware, health/readiness endpoints, WebSocket, Server-Sent Events

🛠 **Practice** — L1: REST API "notes" service (CRUD, in-memory store) with stdlib router + logging middleware · L2: add timeouts, graceful shutdown, rate limiting, JWT auth; load test it with `hey` or k6 and watch behavior on SIGTERM

---

## GO 14 — Databases from Go
**L1 — aware & can use**
- ⬜ **14.1 database/sql** — `sql.DB` is a pool, `sql.Open` vs `Ping`, `QueryContext`/`QueryRowContext`/`ExecContext`, `Scan`, `rows.Close`, `rows.Err`, `sql.ErrNoRows`, `sql.NullString`, placeholders `$1`
- ⬜ **14.2 Transactions** — `BeginTx`, `tx.Commit`, `defer tx.Rollback()`, isolation level option
- ⬜ **14.3 pgx basics** — `pgx` vs `lib/pq`, `pgxpool`, `pgx.CollectRows`, named args
- ⬜ **14.4 Migrations** — `golang-migrate`, `goose`, up/down migrations, running migrations at startup vs in CI

**L2 — deep & efficient**
- ⬜ **14.5 Pool tuning** — `SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`, `SetConnMaxIdleTime`, `db.Stats()`, pool exhaustion, leaks by unclosed rows, PgBouncer interaction
- ⬜ **14.6 pgx advanced** — `pgx.Batch`, `CopyFrom`, `pgconn.PgError` codes (`23505` unique violation), prepared statement cache, `LISTEN/NOTIFY`
- ⬜ **14.7 Query tools trade-offs** — raw SQL, `sqlc`, `squirrel`, GORM, `ent`, `scany` — when to choose which
- ⬜ **14.8 Data access patterns** — repository pattern, domain model vs DB model, transactions across repositories (unit of work, passing `tx`), SQL injection prevention, `go-redis` + cache-aside

🛠 **Practice** — L1: connect the "notes" API to Postgres (Docker) with pgx + migrations · L2: generate queries with sqlc, add a transfer-money operation in one transaction, insert 100k rows with `CopyFrom` vs loop and compare timings

---

## GO 15 — gRPC & Protocol Buffers
**L1 — aware & can use**
- ⬜ **15.1 Protobuf** — `.proto`, proto3, message, field numbers, scalar types, `repeated`, `oneof`, `optional`, well-known types (`Timestamp`)
- ⬜ **15.2 Tooling** — `protoc`, `buf`, `protoc-gen-go`, `protoc-gen-go-grpc`, generated code
- ⬜ **15.3 Unary gRPC service** — service definition, server, client, `grpcurl`, metadata

**L2 — deep & efficient**
- ⬜ **15.4 Streaming** — server streaming, client streaming, bidirectional streaming, flow control, HTTP/2 basis
- ⬜ **15.5 gRPC in production** — deadlines, status codes (`codes.NotFound`...), interceptors, retries, health checking, reflection, schema evolution (`reserved`, never reuse field numbers), grpc-gateway, Connect (connectrpc)

🛠 **Practice** — L1: `UserService` with Create/Get over gRPC, call it with `grpcurl` · L2: add a logging interceptor, deadlines, a server-streaming endpoint, and a breaking-change check with `buf breaking`

---

## GO 16 — Messaging from Go
**L1 — aware & can use**
- ⬜ **16.1 Kafka client basics** — `franz-go`, `segmentio/kafka-go`, producer, consumer group, topic, partition, offset commit (auto vs manual)

**L2 — deep & efficient**
- ⬜ **16.2 Reliable consumers** — at-least-once processing, idempotent consumer, retry topic, dead-letter queue, graceful consumer shutdown, ordering per key
- ⬜ **16.3 Transactional outbox in Go** — outbox table, relay/poller, `SELECT ... FOR UPDATE SKIP LOCKED`, exactly-once myth
- ⬜ **16.4 Other brokers** — RabbitMQ (`amqp091-go`, ack/nack, prefetch), NATS / JetStream

🛠 **Practice** — L1: producer + consumer for "order created" events with Kafka in Docker · L2: idempotent consumer with dedup table + outbox relay; kill the consumer mid-batch and prove no event is lost or double-applied

---

## GO 17 — Application design in Go
**L1 — aware & can use**
- ⬜ **17.1 Layering** — handler (transport) → service (business logic) → repository (data), DTO vs domain model, dependency direction
- ⬜ **17.2 Dependency injection** — constructor injection, `main` as composition root, no global state
- ⬜ **17.3 Configuration** — env vars, config struct, `envconfig`/`koanf`/`viper`, validation at startup, 12-factor app

**L2 — deep & efficient**
- ⬜ **17.4 Clean / hexagonal in Go** — ports & adapters, domain package, interfaces at the consumer, `wire`/`fx` (and why manual DI is usually enough), over-engineering risk
- ⬜ **17.5 Go design patterns** — functional options, constructor `NewX`, factory, builder, strategy, decorator/middleware, adapter, singleton via `sync.Once`, observer via channels

🛠 **Practice** — L1: refactor the "notes" API into handler/service/repository layers with constructor DI · L2: swap Postgres repository for an in-memory one without touching the service; add functional options to a client type

---

## GO 18 — Observability in Go
**L1 — aware & can use**
- ⬜ **18.1 Structured logging** — `slog`, log levels, request-scoped logger, request ID in logs, don't log secrets
- ⬜ **18.2 Health & pprof endpoints** — `/healthz`, `/readyz`, `net/http/pprof`, `expvar`

**L2 — deep & efficient**
- ⬜ **18.3 Metrics** — `prometheus/client_golang`, counter, gauge, histogram, labels, cardinality, `/metrics`, RED metrics middleware
- ⬜ **18.4 Tracing** — OpenTelemetry Go SDK, tracer, span, attributes, context propagation, `otelhttp`, `otelgrpc`, OTLP exporter
- ⬜ **18.5 Runtime diagnostics** — `runtime/metrics`, flight recorder (Go 1.25), `GODEBUG` traces, continuous profiling

🛠 **Practice** — L1: add slog + health endpoints + pprof to the "notes" API · L2: add Prometheus metrics and OTel tracing, view them in Grafana + Jaeger via Docker Compose

---

## GO 19 — Tooling & code quality
**L1 — aware & can use**
- ⬜ **19.1 Formatting & linting** — `gofmt`, `goimports`, `go vet`, `staticcheck`, `golangci-lint`, `govulncheck`
- ⬜ **19.2 Debugging** — Delve (`dlv debug`, breakpoints, step, goroutines view), debugging in VS Code
- ⬜ **19.3 Build basics** — `go build -o`, `CGO_ENABLED=0`, cross-compilation `GOOS=linux GOARCH=arm64`

**L2 — deep & efficient**
- ⬜ **19.4 Build advanced** — `-ldflags "-X main.version=..."`, `-trimpath`, build tags `//go:build`, static binaries, reproducible builds
- ⬜ **19.5 Code generation** — `go generate`, `//go:generate`, `stringer`, `mockery`, `sqlc`, `text/template`, `go/ast`
- ⬜ **19.6 Project automation** — Makefile / Taskfile, `air` hot reload, pre-commit hooks, `golangci-lint` config, multi-stage Dockerfile for Go, distroless/scratch images

🛠 **Practice** — L1: set up `golangci-lint` and fix all findings in your project; debug a bug with Delve · L2: Makefile with lint/test/build targets, version injected via `-ldflags`, Docker image under 20 MB

---

## GO 20 — Runtime internals
**L1 — aware & can use**
- ⬜ **20.1 Goroutines vs threads (overview)** — small growable stack (starts ~2 KB), M:N scheduling, `GOMAXPROCS`, why goroutines are cheap
- ⬜ **20.2 GC overview** — what a garbage collector does, stack vs heap, `GOGC`, `GOMEMLIMIT`, GC pauses

**L2 — deep & efficient**
- ⬜ **20.3 GMP scheduler** — G, M, P, local and global run queues, work stealing, hand-off on syscall, `sysmon`, asynchronous preemption (Go 1.14), container-aware `GOMAXPROCS` (Go 1.25), netpoller (epoll/kqueue)
- ⬜ **20.4 Memory allocator** — `mcache`, `mcentral`, `mheap`, spans, size classes, tiny allocator
- ⬜ **20.5 Escape analysis** — escape to heap, `go build -gcflags=-m`, why returning pointers can allocate, interfaces cause escapes, inlining
- ⬜ **20.6 GC in depth** — tri-color mark & sweep, concurrent marking, write barrier, stop-the-world phases, GC pacer, `GODEBUG=gctrace=1`, Green Tea GC (Go 1.25 experiment → default in 1.26)
- ⬜ **20.7 Channel internals** — `hchan`, circular buffer, `sendq`/`recvq`, `sudog`, lock inside channel, direct send to a waiting receiver
- ⬜ **20.8 Interface & map internals** — `iface`, `eface`, `itab`, dynamic dispatch, devirtualization, Swiss table groups and control bytes
- ⬜ **20.9 reflect & unsafe** — `reflect.Type`, `reflect.Value`, `Kind`, struct tags via reflection, reflection cost, `unsafe.Pointer`, `unsafe.Sizeof`/`Alignof`/`Offsetof`, cgo call cost

**L3 — internals (📖 Anatomy of Go)**
- ⬜ **20.10 Binary & memory layout** — TEXT, RODATA, DATA, BSS, noptr sections, what each part of your program becomes (Ch7 §1)
- ⬜ **20.11 Heap & virtual memory** — virtual memory, arenas, how Go reserves memory, small/large/tiny allocation paths, pointer bitmaps and malloc header (Ch7 §2)
- ⬜ **20.12 Goroutine stack anatomy** — segmented vs contiguous stacks, stack frame structure, frame types, stack growth (`morestack`), stack copying and shrinking (Ch7 §3)
- ⬜ **20.13 GC cycle step by step** — trigger, sweep termination, marking (roots, workers, assists), mark termination, sweeping, GOGC/GOMEMLIMIT math (Ch7 §5)
- ⬜ **20.14 Runtime startup** — `rt0`, thread-local storage, `g0` on `m0`, cgo startup, handoff to `main` (Ch8 §5)
- ⬜ **20.15 Goroutine lifecycle & scheduler loop** — `newproc`, `runnext` slot, local/global run queues, `schedule()`/`findRunnable`, goroutine states (Ch8 §6–7)
- ⬜ **20.16 Preemption** — cooperative scheduling, cooperative preemption via stack checks, asynchronous preemption via signals (Ch8 §8)
- ⬜ **20.17 I/O, netpoller & sysmon** — file descriptors, blocking vs poller-based I/O, epoll/kqueue integration, sysmon cycle and its jobs (Ch8 §9–10)
- ⬜ **20.18 select internals** — `selectgo`, random case order, lock ordering, select with closed channels (Ch8 §2)

📖 **Anatomy of Go:** Ch7 Memory (p.577–737) · Ch8 Concurrency (p.738–873)

🛠 **Practice** — L1: spawn 100k goroutines and observe memory; run a program with `GODEBUG=gctrace=1` and read the output · L2: run `-gcflags=-m` on your code and remove 3 heap escapes; view scheduler behavior in `go tool trace`; read `runtime/chan.go` and draw `hchan`

---

## GO 21 — Performance & profiling
**L1 — aware & can use**
- ⬜ **21.1 Benchmarks** — `go test -bench`, `-benchmem`, ns/op, allocs/op, `b.Loop`
- ⬜ **21.2 pprof basics** — CPU and heap profiles, `go tool pprof`, `top`, `list`, web UI flame graph

**L2 — deep & efficient**
- ⬜ **21.3 Advanced profiling** — goroutine/block/mutex profiles, `go tool trace`, `benchstat`, comparing before/after, avoiding misleading benchmarks
- ⬜ **21.4 Optimization techniques** — preallocation, avoiding `[]byte`↔`string` copies, `strings.Builder`, `sync.Pool`, fewer pointers, struct field ordering, bounds check elimination, inlining budget, false sharing
- ⬜ **21.5 PGO & GC tuning** — profile-guided optimization (`default.pgo`), `GOGC` vs `GOMEMLIMIT` trade-off, memory limits in containers

**L3 — internals (📖 Anatomy of Go)**
- ⬜ **21.6 All pprof profiles in depth** — heap/alloc, CPU, mutex, block, goroutine, threadcreate, delta profiles; reading call graphs (Ch6 §5)
- ⬜ **21.7 Execution tracer** — enabling traces, reading the timeline, goroutine analysis, tasks/regions/logs, tracer internals (Ch6 §6)
- ⬜ **21.8 PGO mechanics** — how profiles drive hot-call inlining and devirtualization (Ch6 §7)

📖 **Anatomy of Go:** Ch6 §5–7 (p.471–574)

🛠 **Practice** — L1: benchmark 3 ways to build a string; profile the "notes" API under load and find the top function · L2: cut allocations of a JSON-heavy endpoint by 50% and prove it with `benchstat`; try PGO and measure

---

## GO 22 — Idioms & common mistakes
**L1 — aware & can use**
- ⬜ **22.1 Idiomatic Go** — Effective Go, Go Code Review Comments, Go Proverbs, naming (short names, `MixedCaps`, receiver names), "clear is better than clever", "a little copying is better than a little dependency"
- ⬜ **22.2 Common mistakes** — (from "100 Go Mistakes") shadowing, unnecessary nesting, `init` misuse, getters/setters, interface pollution, slice/map traps, `defer` in loops, ignored errors, copying sync types, `time.After` leak in loops, unclosed response bodies

**L2 — deep & efficient**
- ⬜ **22.3 Reading & reviewing real Go code** — reading stdlib source (`net/http`, `sync`, `context`), popular projects (Prometheus, Kubernetes), Go code review checklist, reviewing AI-generated Go code

🛠 **Practice** — L1: take one of your old work files and list every non-idiomatic spot · L2: read `context.go` from the stdlib and explain it to your student; review an AI-generated Go service and find 5 problems

---

## GO 23 — Go ecosystem
**L1 — aware & can use**
- ⬜ **23.1 Notable libraries** — chi, gin, echo, fiber (fasthttp caveat), pgx, sqlc, go-redis, zap/zerolog (pre-slog), testify, viper/koanf, cobra, OpenTelemetry, ent
- ⬜ **23.2 CLI tools in Go** — `flag`, `cobra`, `urfave/cli`, exit codes, stdin/stdout piping

**L2 — deep & efficient**
- ⬜ **23.3 Go in the industry** — Docker, Kubernetes, Terraform, Prometheus, etcd are written in Go (why Go fits cloud infrastructure), WebAssembly (`GOOS=js`, `wasip1`), TinyGo

🛠 **Practice** — L1: build a small CLI with cobra (e.g. a todo or a log analyzer) · L2: compile a Go function to WebAssembly and call it from a web page

---

## GO 24 — How Go compiles (📖 Anatomy of Go Ch5)
**L1 — aware & can use**
- ⬜ **24.1 The build pipeline** — compile per package → link, `go build -x` / `-work`, object files, what the linker does (Ch5 §5.1–5.2)

**L3 — internals (📖 Anatomy of Go)**
- ⬜ **24.2 Parsing** — scanner and tokens, automatic semicolons, parser and the syntax tree (Ch5 §5.3)
- ⬜ **24.3 Type checking** — package-level declarations, function bodies, scope resolution, package variable initialization order (Ch5 §5.4)
- ⬜ **24.4 IR & export data** — IR nodes, export data between packages, relocations (Ch5 §5.5)
- ⬜ **24.5 Optimizations** — dead code elimination, devirtualization, inlining and its budget, escape analysis (`-gcflags=-m`) (Ch5 §5.6)
- ⬜ **24.6 Walk & order** — how the middle end rewrites high-level statements into runtime calls (Ch5 §5.7)
- ⬜ **24.7 SSA** — SSA form, building SSA from IR, SSA passes, reading `GOSSAFUNC` output (Ch5 §5.8)
- ⬜ **24.8 Machine code & Go assembly** — code generation, reading `go tool compile -S` and `go tool objdump` output (Ch5 §5.9)

📖 **Anatomy of Go:** Ch5 (p.259–413)

🛠 **Practice** — L1: build a two-package program with `go build -x -work` and name each step · L3: dump `GOSSAFUNC=main` for a small function and explain 3 passes; find one inlined and one escaped variable with `-gcflags=-m=2`
