# GO 5 — Errors, defer, panic
Date: 2026-10-10 · Layers: L1 + L2 + L3 (book) · Teaching: week 3, lesson 5

📖 **Read in the book:** Ch6 §3 "Defer" (PDF p.436–459), Ch6 §4 "Panic & Recover" (p.459–471). The book does **not** cover error values and wrapping — for that read the Go blog posts "Errors are values" and "Working with Errors in Go 1.13".
🎓 **Course files:** `go-course/lesson-plan/week03/lesson-05-plan.md` (55–85 min block), `cheatsheet.md`, `bozorcha/lesson-05.md`

## TL;DR
EN: Errors are normal values that you return and check. Wrap them with `%w` to add context, inspect them with `errors.Is` / `errors.As`. `defer` runs cleanup when the function exits (LIFO). `panic` is for bugs, not for expected failures.
RU: Ошибки — обычные значения, их возвращают и проверяют. `%w` оборачивает ошибку с контекстом, `errors.Is` / `errors.As` её проверяют. `defer` выполняет очистку при выходе из функции (LIFO). `panic` — для багов, а не для ожидаемых ошибок.

## Problem → Solution
Problem: in Bozorcha, `AddItem(id, qty) bool` returns `false` for "no product", "qty is 0" and "not enough stock". The caller can't show the right message ("only 2 left").
Solution: return `error` values that carry *what* went wrong: a sentinel `ErrOutOfStock`, a typed `*ValidationError{Field: "qty"}`, wrapped with context. The HTTP layer later maps them to 404 / 400 / 409.
RU: `bool` не говорит, что именно пошло не так. Ошибки-значения несут причину и контекст, и вызывающий код может правильно отреагировать.

---

## Part 1 — Error basics (L1)

```go
type error interface { Error() string }   // that's all an error is

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```
- Error is the **last** return value. Check it **right away**, return early ("happy path on the left").
- Message style: lowercase, no final dot, no "error:" prefix. Each layer adds context, so the final text reads like a path:
  `checkout order 7: reserve product 42: not enough stock`
- Handle an error **once**: either log it or return it, not both (otherwise you get the same error logged 5 times).

RU: `error` — интерфейс с одним методом. Проверяй сразу, возвращай рано. Сообщения с маленькой буквы, без точки. Обрабатывай ошибку один раз: либо логируй, либо возвращай.

## Part 2 — Wrapping & inspection (L1)

```go
var ErrNotFound = errors.New("not found")                 // sentinel

func (s *Store) Get(id int) (*Product, error) {
    p, ok := s.items[id]
    if !ok {
        return nil, fmt.Errorf("product %d: %w", id, ErrNotFound)  // wrap
    }
    return p, nil
}

err := ...
errors.Is(err, ErrNotFound)      // true, walks the chain
err == ErrNotFound               // false! the outer value is the wrapper

var ve *ValidationError
if errors.As(err, &ve) {         // finds the first *ValidationError in the chain
    fmt.Println(ve.Field)
}
```
| Question | Tool |
|---|---|
| Is it this exact error? | `errors.Is(err, ErrX)` |
| Is it this kind, and give me its fields | `errors.As(err, &target)` (target = pointer to a variable) |
| Add context, keep the original | `fmt.Errorf("...: %w", err)` |
| Add context, hide the original | `fmt.Errorf("...: %v", err)` |
| Combine several errors | `errors.Join(e1, e2)` (Go 1.20), `%w` can appear several times |
| Get one level down | `errors.Unwrap(err)` (rarely needed directly) |

How it works: a wrapper type has `Unwrap() error` (or `Unwrap() []error` for joined errors). `errors.Is` calls `Unwrap` in a loop and compares each level with `==` (or calls the error's own `Is(target) bool` method if it has one).

RU: `%w` оборачивает и сохраняет исходную ошибку, `%v` только форматирует текст. `errors.Is` — «это именно эта ошибка?», `errors.As` — «это ошибка такого типа? дай её поля».

## Part 3 — Error strategies (L2)

| Kind | Example | When |
|---|---|---|
| Sentinel | `io.EOF`, `sql.ErrNoRows`, `ErrNotFound` | one specific condition, no extra data |
| Custom type | `*ValidationError{Field, Msg}`, `*fs.PathError` | caller needs data from the error |
| Opaque | just return/wrap with `%v` | caller should only know "it failed" |
| Behaviour check | `interface{ Timeout() bool }` | you care what the error *can do*, not its type (e.g. `net.Error`) |

- Custom error types usually use **pointer receivers** and are returned as `*T` → matches the `errors.As(err, &ve)` pattern with `var ve *ValidationError`.
- `errors.AsType[*ValidationError](err)` (Go 1.26) — generic version, returns `(value, ok)`, no target variable needed.
- Sentinel errors are part of your **public API**: once callers use `errors.Is(err, ErrX)`, you can't rename it. Wrapping with `%w` also makes the inner error part of your API — use `%v` when you want to hide it.
- Mapping in Bozorcha (week 9): `ErrNotFound`→404, `*ValidationError`→400, `ErrOutOfStock`→409, everything else→500 (and log it).

RU: Sentinel — для одного конкретного случая, свой тип — когда нужны данные, opaque — когда вызывающему не нужно знать детали. Обёртка через `%w` делает внутреннюю ошибку частью API.

## Part 4 — defer (L1 + L2)

```go
f, err := os.Open(path)
if err != nil { return err }
defer f.Close()                 // runs when the function returns, whatever the path

mu.Lock()
defer mu.Unlock()
```
Rules:
1. **LIFO:** the last `defer` runs first.
2. **Arguments are evaluated immediately**, at the `defer` line. The *call* happens later.
   ```go
   x := 1
   defer fmt.Println(x)          // prints 1
   defer func() { fmt.Println(x) }()   // prints 2 — closure reads x at the end
   x = 2
   ```
   Same for receivers: `defer r.Log()` with a value receiver copies `r` now.
3. defer belongs to the **function**, not to a block or loop iteration. `defer f.Close()` inside a loop keeps every file open until the function ends → move the loop body into a small function.
4. A deferred closure **can change named results**: that's how you turn a panic into an error, or add context to a returned error.
   ```go
   func save() (err error) {
       defer func() {
           if err != nil { err = fmt.Errorf("save: %w", err) }
       }()
       ...
   }
   ```
5. `Close()` on a **written** file can fail (data not flushed) — don't silently ignore that error.
6. `os.Exit` skips all defers. `log.Fatal` calls `os.Exit`.

RU: defer выполняется при выходе из функции в порядке LIFO. Аргументы вычисляются сразу. defer в цикле не закрывает ресурс на каждой итерации. Отложенная функция может изменить именованный результат. `os.Exit` не выполняет defer.

## Part 5 — panic & recover (L1 + L2)

- `panic(v)` stops normal execution and **unwinds** the stack of the **current goroutine**, running its defers.
- Runtime panics: nil map write, index out of range, nil pointer dereference, integer division by zero (float division gives ±Inf, no panic), failed `x.(T)`, closing a closed channel.
- `recover()` stops the panic, but **only when called directly inside a deferred function**. Called anywhere else (normal code, or a helper called by the deferred function) it returns `nil`.
- After recovery the panicking function does **not** continue; it finishes its remaining defers and returns normally with its current result values (zero values unless named results were set).
- A panic in **any goroutine** that isn't recovered **in that same goroutine** crashes the whole program. A recover in `main` can't save a panic in another goroutine. → every goroutine you start in a server needs its own recover (or must be panic-free).
- `panic(nil)` since Go 1.21 becomes `*runtime.PanicNilError`, so `recover()` returns non-nil.
- When to panic: programmer bugs and impossible states, mostly at startup (`regexp.MustCompile`, missing dependency). Never for user input, missing rows, network errors.
- Recovery middleware in HTTP: `defer func(){ if r := recover(); r != nil { log; 500 } }()` — `net/http` already recovers per request, but logs poorly; frameworks add their own.

RU: panic раскручивает стек текущей горутины и выполняет defer. recover работает только прямо внутри отложенной функции. Паника в любой горутине без recover в ней же роняет всю программу. panic — для багов, ошибки — для всего остального.

## Part 6 — Under the hood (L3, book Ch6 §3–4)

**defer is not just "copy the call to the end".** The compiler can't always know which defers will run (`if cond { defer ... }`), so there are three implementations:

| Kind | When | Cost |
|---|---|---|
| **Open-coded** (Go 1.14) | simple functions: ≤ 8 defers, not in a loop, returns × defers ≤ 15, no heap-allocated results, no heap defer in the function | almost free: calls are inlined at each exit; an 8-bit `deferBits` mask records which ones were reached |
| **Stack-allocated** (Go 1.13) | defer at function top level but open-coding not possible | `_defer` record lives in the stack frame, `deferprocStack` |
| **Heap-allocated** | defer inside a loop (count unknown) | `_defer` from a per-P pool or heap, `runtime.deferproc` |

- Each goroutine (`g`) keeps a **linked list of `_defer` records** (`g._defer` is the head). New defers are pushed to the head → LIFO.
- `defer f(a, b)` becomes a hidden closure `main.foo.deferwrap1` that captured `a, b` when the line ran → that's *why* arguments are evaluated immediately. You see these names in stack traces.
- `deferreturn` runs the records whose saved `sp` matches the current frame.
- Per-P **defer pools** (no locks), refilled from a global pool; the global pool is emptied every GC cycle.
- See it yourself: `go build -gcflags=-d=defer=1` prints which kind each defer got.

**panic:** `panic(v)` → `runtime.gopanic`. It creates a `_panic` record, walks the defer chain frame by frame and runs each deferred function. If none recovers → print and `fatalpanic` (exit code 2).
**recover:** `recover()` → `runtime.gorecover(argp)`. It succeeds only if the caller's **argument pointer** equals the one saved for the deferred call that the panic is running right now — that's the technical reason it works only *directly* in the deferred function. It sets `p.recovered = true`; then `recovery` rewrites the goroutine's saved PC/SP so that `deferproc` "returns 1" a second time, the function jumps to its epilogue, runs the remaining defers and returns normally.

RU: Есть три вида defer: open-coded (почти бесплатный), на стеке и в куче (в цикле). Записи `_defer` образуют связный список в горутине. panic создаёт запись `_panic` и проходит по цепочке defer; recover срабатывает только при совпадении указателя аргументов, поэтому работает лишь прямо внутри отложенной функции.

## Key terms
| Term | Meaning (simple EN) | RU |
|---|---|---|
| sentinel error | a package-level error value compared with `errors.Is` | ошибка-сигнал (sentinel) |
| wrapping | adding context while keeping the original error | оборачивание ошибки |
| error chain | the list you get by calling `Unwrap` repeatedly | цепочка ошибок |
| named result | a return value with a name, visible to defers | именованный результат |
| stack unwinding | going back up the call stack running defers | раскрутка стека |
| open-coded defer | defer inlined at function exits by the compiler | встроенный defer |
| `_defer` / `_panic` | runtime records for defers / active panics | — |

## Senior answer (interview-ready)
**Q: How do you handle errors in Go?**
EN: "Errors are values. I return them as the last result, check them right away and add context with `fmt.Errorf` and `%w`. Callers use `errors.Is` for sentinel errors and `errors.As` for typed errors with data. I handle each error once — log it or return it. At the edge, like HTTP, I map domain errors to status codes. Panic is only for programmer bugs, and every goroutine I start in a server recovers its own panics, because one unrecovered panic kills the whole process."
RU: «Ошибки — это значения. Я возвращаю их последним результатом, сразу проверяю и добавляю контекст через `fmt.Errorf` с `%w`. Вызывающий код использует `errors.Is` для sentinel-ошибок и `errors.As` для типизированных ошибок с данными. Каждую ошибку обрабатываю один раз — логирую или возвращаю. На границе, например в HTTP, превращаю доменные ошибки в статус-коды. Panic — только для багов программиста, и каждая горутина в сервере сама перехватывает свои паники, потому что одна неперехваченная паника убивает весь процесс.»

## Gotchas
- `err == ErrX` breaks after wrapping — use `errors.Is`.
- `errors.As` needs a **pointer to** the target variable.
- defer in a loop → resource leak until the function ends.
- Deferred call arguments are fixed at the `defer` line.
- `recover()` in a helper function (not directly deferred) returns nil.
- Panic in a goroutine is not caught by `main`'s recover.
- `os.Exit` / `log.Fatal` skip defers.

## Recall questions
1. What is the `error` type, exactly?
2. `%w` vs `%v` — what changes for the caller?
3. Why is `err == ErrNotFound` false after wrapping, and how does `errors.Is` find it?
4. When do you choose a custom error type over a sentinel? Write the `errors.As` call.
5. What does this print? `x := 1; defer fmt.Println(x); x = 2`
6. How can a deferred function change what a function returns?
7. Where must `recover` be called to work, and what is the runtime reason (book)?
8. What happens to the program if a goroutine panics and nobody recovers inside it?
9. Name the three kinds of defer and when each is used.
10. Why is a defer inside a loop slower and dangerous?

## Links
Practice: `code/golang/05-errors-defer-panic/` (to create) · Source: book Ch6 §3–4, Go blog "Working with Errors in Go 1.13", `errors` package docs
