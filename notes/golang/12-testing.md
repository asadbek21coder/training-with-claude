# GO 12 — Testing in Go
Date: 2026-10-10 · Layers: L1 + L2 · Teaching: week 3, lesson 6

📖 **Book:** *The Anatomy of Go* has no testing chapter. Benchmarks + profiles from tests: Ch6 §5 (PDF p.471–522), PGO: Ch6 §7. Main sources: `go help test`, `go help testflag`, `testing` package docs.
🎓 **Course files:** `go-course/lesson-plan/week03/lesson-06-plan.md`, `cheatsheet.md`, `homework-03.md`, `bozorcha/lesson-06.md`

## TL;DR
EN: Tests are normal Go functions in `_test.go` files, run by `go test`. The default shape is a table of cases + `t.Run` subtests. Interfaces let you replace real dependencies with fakes.
RU: Тесты — обычные Go-функции в файлах `_test.go`, запускаются через `go test`. Стандартная форма — таблица кейсов + подтесты `t.Run`. Интерфейсы позволяют заменить реальные зависимости фейками.

## Problem → Solution
Problem: in Bozorcha, money and stock must never be wrong. Checking by hand after every change is slow, and later AI will write code you must trust.
Solution: tests that run in seconds on every change: `go test ./...`. In week 3 the student tests cart totals, stock changes and validation, plus a `FakeStore` that fails on save to test error paths.
RU: Деньги и остатки не должны ломаться. Тесты проверяют это за секунды после каждого изменения — особенно важно, когда код пишет ИИ.

---

## Part 1 — Basics (L1)

```go
// cart.go                         // cart_test.go
package cart                       package cart
                                   import "testing"
func Total(items []Item) int64     func TestTotal(t *testing.T) {
                                       got := Total([]Item{{Price: 100, Qty: 2}})
                                       if want := int64(200); got != want {
                                           t.Errorf("Total() = %d; want %d", got, want)
                                       }
                                   }
```
- File ends with `_test.go`, sits next to the code. Not compiled into the normal binary.
- Function: `func TestXxx(t *testing.T)`, `Xxx` starts with a capital letter.
- Message convention: `Func(args) = got; want want`.
- `t.Errorf` = mark failed, **continue**. `t.Fatalf` = mark failed, **stop this test now** (use when continuing would panic, e.g. result is nil).
- `package cart` = white-box test (sees unexported names). `package cart_test` = black-box (only public API, like a real user).

Commands:
```bash
go test ./...                    # all packages
go test -v ./...                 # show each test
go test -run 'TestTotal/empty' . # one test / subtest (regexp)
go test -count=1 ./...           # skip the test cache
go test -race ./...              # race detector (needs cgo)
go test -cover ./...             # coverage %
go test -coverprofile=c.out ./... && go tool cover -html=c.out
go test -bench=. -benchmem ./...
```
Test cache: when nothing changed, `go test` prints `(cached)` and doesn't run again. `-count=1` forces a run.

RU: Файл `_test.go`, функция `TestXxx(t *testing.T)`. `Errorf` — продолжить, `Fatalf` — остановить тест. `package x_test` — тест «чёрного ящика».

## Part 2 — Table-driven tests & subtests (L1)

```go
func TestParsePrice(t *testing.T) {
    tests := []struct {
        name    string
        in      string
        want    int64
        wantErr error
    }{
        {name: "whole sum", in: "12000", want: 1200000},
        {name: "with tiyin", in: "12.50", want: 1250},
        {name: "empty", in: "", wantErr: ErrInvalidPrice},
        {name: "negative", in: "-5", wantErr: ErrInvalidPrice},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ParsePrice(tt.in)
            if !errors.Is(err, tt.wantErr) {
                t.Fatalf("err = %v; want %v", err, tt.wantErr)
            }
            if got != tt.want {
                t.Errorf("got %d; want %d", got, tt.want)
            }
        })
    }
}
```
Why it's the default: a new case = one line, each case has a name in the output, a failure points to the exact case, and you can run one case with `-run`.
(Go 1.22+: the loop variable `tt` is new in each iteration, so the old `tt := tt` line is not needed.)

Helpers:
| Call | Meaning |
|---|---|
| `t.Run(name, fn)` | named subtest |
| `t.Helper()` | failures are reported at the caller's line |
| `t.Cleanup(fn)` | run fn after the test (LIFO, like defer, but works from helpers) |
| `t.TempDir()` | temp directory, deleted automatically |
| `t.Setenv(k, v)` | set env var for this test only |
| `t.Parallel()` | run this test in parallel with other parallel tests |
| `t.Skip()` | skip (slow, needs Docker, etc.) |
| `t.Context()` | context cancelled when the test ends (Go 1.24) |

**Choosing cases (the part that matters more than syntax):** normal case · boundaries (0, 1, empty, nil, max) · every error branch · the bug you fixed last time.

RU: Табличный тест: срез структур + `t.Run`. Новый кейс — одна строка, каждый кейс имеет имя. Главное — выбрать кейсы: обычный, граничные, все ветки ошибок, старый баг.

## Part 3 — Test doubles: fakes through interfaces (L2, needed in week 3)

```go
type Store interface {
    Save(p *Product) error
}
type FakeStore struct{ Err error; Saved []*Product }
func (f *FakeStore) Save(p *Product) error {
    if f.Err != nil { return f.Err }
    f.Saved = append(f.Saved, p)
    return nil
}
// test: svc := NewService(&FakeStore{Err: errors.New("disk full")}) → check the error path
```
| Double | What it does |
|---|---|
| fake | working simple implementation (in-memory store) |
| stub | returns fixed answers |
| spy | records calls so you can check them |
| mock | pre-programmed expectations, fails if calls differ (`gomock`, `mockery`) |
Go style prefers hand-written **fakes**: simple, readable, no framework. This is the payoff of small consumer-side interfaces (see GO 4).

RU: Маленький интерфейс позволяет подставить фейк и проверить ветки ошибок. В Go чаще пишут фейки руками, а не моки через фреймворки.

## Part 4 — Coverage, race, benchmarks (L1 → L2)

- **Coverage** shows which lines never ran. It is a map, not a score: 100% coverage with weak checks proves nothing. Aim at error branches and boundaries.
- **`-race`** instruments every memory access and reports unsynchronized concurrent access **that actually happened during the run**. It is several times slower and uses more memory — fine for tests/CI, not for production. No report ≠ no race.
- **Benchmarks:**
  ```go
  func BenchmarkTotal(b *testing.B) {
      items := makeItems(1000)        // setup is not measured with b.Loop
      for b.Loop() {                   // Go 1.24; older style: for i := 0; i < b.N; i++
          Total(items)
      }
  }
  ```
  Read `ns/op` (time per call), `B/op` and `allocs/op` (with `-benchmem`). Compare runs with `benchstat`. Measure, don't guess.
- **Profiles from tests:** `go test -bench=. -cpuprofile cpu.out -memprofile mem.out` then `go tool pprof cpu.out` (book Ch6 §5).
- **Fuzzing** (Go 1.18): `func FuzzParsePrice(f *testing.F)` + `f.Add(seed)` + `f.Fuzz(func(t *testing.T, s string){...})`, run with `go test -fuzz=FuzzParsePrice`. Finds inputs you didn't think of.
- **Examples:** `func ExampleTotal()` with `// Output:` comment — a test and documentation at once.

RU: Покрытие показывает, что не выполнялось, но не качество проверок. `-race` находит только реально случившиеся гонки. Бенчмарки: смотри `ns/op` и `allocs/op`, сравнивай через `benchstat`.

## Key terms
| Term | Meaning (simple EN) | RU |
|---|---|---|
| table-driven test | cases in a slice, one loop runs them | табличный тест |
| subtest | named test inside a test (`t.Run`) | подтест |
| test double / fake | replacement for a real dependency | тестовый дублёр / фейк |
| white-box / black-box test | sees internals / only public API | белый / чёрный ящик |
| coverage | % of statements executed by tests | покрытие |
| data race | two goroutines access the same memory, at least one writes, no sync | гонка данных |
| benchmark | measures time and allocations per operation | бенчмарк |
| fuzzing | auto-generated random inputs to find crashes | фаззинг |

## Senior answer (interview-ready)
**Q: How do you test Go code?**
EN: "I write table-driven tests with subtests, so each case has a name and adding a case is one line. I focus on boundaries and error branches, not on the coverage number. Dependencies are behind small interfaces, so I pass simple hand-written fakes. CI runs `go test -race ./...`. For hot code I write benchmarks with `-benchmem` and compare with benchstat, and for parsers I add fuzz tests. Integration tests run against a real database in a container."
RU: «Я пишу табличные тесты с подтестами — у каждого кейса есть имя, а новый кейс — одна строка. Фокусируюсь на граничных значениях и ветках ошибок, а не на проценте покрытия. Зависимости спрятаны за маленькими интерфейсами, поэтому я передаю простые фейки. В CI запускается `go test -race ./...`. Для горячего кода пишу бенчмарки с `-benchmem` и сравниваю через benchstat, для парсеров добавляю фаззинг. Интеграционные тесты идут с реальной базой в контейнере.»

## Gotchas
- `t.Fatal` from a goroutine other than the test's own does not stop the test correctly — use `t.Error` + return there.
- `t.Parallel()` subtests + shared state = flaky tests.
- `(cached)` results hide that you didn't really re-run — use `-count=1` when debugging.
- Coverage % is not quality.
- `-race` misses races on code paths that didn't run.
- Comparing errors with `==` in tests breaks after wrapping — use `errors.Is`.

## Recall questions
1. What 3 rules make a function a test?
2. `t.Errorf` vs `t.Fatalf` — when do you need Fatal?
3. Write the skeleton of a table-driven test with a `wantErr` field.
4. What does `t.Helper()` change?
5. `package cart` vs `package cart_test` in a test file?
6. Which cases would you write for a `Cart.Remove(productID)` method?
7. Why is a hand-written fake usually better than a mock framework in Go?
8. What does `-race` really detect, and what does it cost?
9. What do `ns/op` and `allocs/op` mean? What is `b.Loop`?
10. Why is 100% coverage not a good goal?

## Links
Practice: `code/golang/12-testing/` (to create) · Source: `testing` docs, Go blog "Using Subtests and Sub-benchmarks", book Ch6 §5 for profiles
