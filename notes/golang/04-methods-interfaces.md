# GO 4 — Methods & interfaces
Date: 2026-10-10 · Layers: L1 + L2 + L3 (book) · Teaching: week 3, lesson 5

📖 **Read in the book:** Ch4 §2 "Interfaces" (PDF p.202–233). Methods/receivers: Ch4 §1.1–1.2 (p.175–190).
🎓 **Course files:** `go-course/lesson-plan/week03/lesson-05-plan.md`, `cheatsheet.md`, `bozorcha/lesson-05.md`

## TL;DR
EN: An interface is a list of methods. Any type that has those methods satisfies it automatically. At runtime an interface value is two words: "what type is inside" + "where the data is".
RU: Интерфейс — это набор методов. Любой тип, у которого есть эти методы, реализует его автоматически. В рантайме значение интерфейса — это два слова: «какой тип внутри» + «где лежат данные».

## Problem → Solution
Problem: a function that saves products should not care *where* they are saved (memory, Postgres, a fake in tests). Without interfaces you would write one function per storage type, or use inheritance like Java.
Solution: describe the behaviour you need (`Save`, `Get`) as an interface. The function accepts the interface. Any type with these methods can be passed in. In Bozorcha this is the `Store` interface: `MemoryStore` now, `PostgresStore` in week 9, a `FakeStore` in tests.

RU: Функции не важно, *куда* сохраняются продукты. Мы описываем нужное поведение интерфейсом, а реализации (память, Postgres, фейк для тестов) подставляются без изменения кода.

---

## Part 1 — Methods (L1)

```go
type Product struct{ Title string; Stock int }

func (p Product) Label() string { return p.Title }   // value receiver: works on a copy
func (p *Product) Reserve(n int) { p.Stock -= n }    // pointer receiver: changes the original
```

- A method is a function with a **receiver**. You can add methods to any named type in your package, not only structs (`type Celsius float64`).
- **Value receiver** gets a copy. **Pointer receiver** can change the value and avoids copying large structs.
- Rule: if one method needs a pointer receiver, make **all** methods of that type pointer receivers (consistency).
- Go auto-converts for calls: `p.Reserve(1)` works on an addressable `Product` (Go writes `(&p).Reserve(1)`), and `ptr.Label()` works on a `*Product`.
- **Method value:** `f := p.Label` → a function bound to `p`. **Method expression:** `Product.Label` → a function that takes the receiver as its first argument.

RU: Метод — функция с получателем. Value receiver работает с копией, pointer receiver меняет оригинал. Не смешивай типы получателей у одного типа.

## Part 2 — Interfaces (L1)

```go
type Shape interface {
    Area() float64
}
type Circle struct{ R float64 }
func (c Circle) Area() float64 { return math.Pi * c.R * c.R }

var s Shape = Circle{R: 2}   // OK: Circle has Area()
```

- **Implicit satisfaction:** no `implements` keyword. Your type does not even need to import the package that defines the interface. This keeps packages loosely coupled.
- Method signatures must match exactly (name, parameter types, result types). Parameter *names* don't matter.
- `any` = `interface{}` = zero methods → every type satisfies it. You lose type safety, so use it only when forced (e.g. `fmt.Println(a ...any)`).
- **Compile-time check:** `var _ Shape = (*Circle)(nil)` makes the compiler prove that `*Circle` satisfies `Shape`. Useful when nothing else in the code would catch it.

### Type assertion and type switch
```go
c := s.(Circle)          // panics if s does not hold a Circle
c, ok := s.(Circle)      // safe: ok == false, c == zero value
switch v := s.(type) {   // type switch
case Circle:  fmt.Println(v.R)
case nil:     fmt.Println("empty interface")
default:      fmt.Println("other")
}
b, ok := s.(Barker)      // assert to ANOTHER INTERFACE: "can it Bark?", not "is it a Dog?"
```
Warning for teaching: lots of type switches usually mean the interface is designed wrong.

### Key stdlib interfaces to know by heart
| Interface | Method(s) | Why it matters |
|---|---|---|
| `error` | `Error() string` | every error in Go |
| `fmt.Stringer` | `String() string` | how `fmt` prints your type |
| `io.Reader` | `Read(p []byte) (n int, err error)` | files, HTTP bodies, network — all streams |
| `io.Writer` | `Write(p []byte) (n int, err error)` | stdout, files, buffers, HTTP responses |
| `io.Closer` | `Close() error` | release resources |
| `sort.Interface` | `Len`, `Less`, `Swap` | old-style sorting |
| `http.Handler` | `ServeHTTP(w, r)` | every HTTP server |
| `json.Marshaler` | `MarshalJSON() ([]byte, error)` | custom JSON |

RU: Тип реализует интерфейс неявно. `any` принимает всё, но убирает проверку типов. Type assertion достаёт конкретный тип, comma-ok форма не паникует.

## Part 3 — Interface design & embedding (L2)

**Embedding interfaces in interfaces** builds bigger contracts from small ones:
`type ReadWriter interface { Reader; Writer }`. Since Go 1.14 two embedded interfaces may contain the same method, if the signatures are identical.

**Embedding a struct in a struct** promotes fields and methods (composition, "has-a"). It is not inheritance: the inner type's methods never call the outer type's methods (no virtual override).

**Embedding an interface in a struct** — the trap:
```go
type Dog struct {
    Speaker          // interface field, zero value is nil
    Name string
}
Dog{Name: "Rex"}.Speak()   // PANIC: nil pointer dereference
```
The struct "has" the method only through the embedded field. If nobody puts a real value in it, the call panics. Real use: `context` package (`valueCtx` embeds `Context` and overrides only `Value`), and partial fakes in tests.

**Design rules (say these in interviews):**
- Small interfaces (1–2 methods). "The bigger the interface, the weaker the abstraction."
- **Accept interfaces, return structs.**
- Define the interface where it is **used** (consumer side), not where it is implemented. Bozorcha: the service package defines the `Store` it needs.
- Don't create an interface "just in case" (interface pollution). Create it when you have 2 implementations, or need a fake for tests.

**Method sets** (which types satisfy an interface):
| Type | Its method set |
|---|---|
| `T` | methods with receiver `T` |
| `*T` | methods with receiver `T` **and** `*T` |

So if `Save` has a pointer receiver, only `*MemoryStore` satisfies `Store`, not `MemoryStore`. Reason: a value inside an interface is not addressable, so Go cannot take its address to call a pointer method.

RU: Маленькие интерфейсы, «принимай интерфейсы — возвращай структуры», интерфейс объявляет потребитель. Набор методов `*T` включает методы `T` и `*T`, а у `T` — только методы `T`.

## Part 4 — The nil interface trap (L2)

```go
func find() error {
    var e *MyErr = nil
    return e            // returns a NON-nil error!
}
if err := find(); err != nil { /* true! */ }
```
An interface is nil only when **both** words are empty (no type, no value). Here the type word holds `*MyErr`, so the interface is not nil, even though the pointer inside is nil.
Fix: `return nil` explicitly when there is no error. Never return a typed nil pointer as an `error`.

RU: Интерфейс равен nil, только если пусты оба слова — тип и значение. Типизированный nil-указатель внутри интерфейса делает интерфейс НЕ nil. Возвращай просто `nil`.

## Part 5 — Under the hood (L3, book Ch4 §2.3–2.5)

**Two layouts, both 16 bytes:**
```
eface (empty interface, any):      iface (interface with methods):
┌─────────┬──────────┐             ┌─────────┬──────────┐
│ _type * │  data *  │             │  tab *  │  data *  │
└─────────┴──────────┘             └────┬────┴──────────┘
                                        ▼
                                   itab: inter (the interface type)
                                         _type (the concrete type)
                                         hash  (copy of type hash, fast type switches)
                                         fun[] (method pointers, one per interface method)
```
- `println(x)` on an interface prints these two words: `(0x..., 0x...)`. Try it.
- A method call through an interface = load `tab`, load `fun[i]`, call it with `data` as the receiver. This indirect call is **dynamic dispatch**. It blocks inlining unless the compiler can **devirtualize** it (knows the concrete type).
- `fun` is declared as `[1]uintptr` but really holds N entries — a variable-length array trick from C.

**What goes into `data`** (the value is copied when stored, unless it already is a pointer):
1. Pointer → stored directly, no copy.
2. Zero-size value (`struct{}`) → points to `runtime.zerobase`.
3. One-byte values (`bool`, `uint8`) → point into the shared array `runtime.staticuint64s`, no allocation.
4. Constants/literals → point to read-only data in the binary.
5. Small value that doesn't escape → copy on the stack.
6. Otherwise → copy on the heap (an allocation — this is why `any` everywhere costs performance).
That is why `a := 10; var b any = a; a = 20` leaves `b == 10`.

**Where itabs come from:**
- Known pairs (`var w Worker = Developer{}`) are built by the compiler and listed in the binary section `.itablink` (see them: `go tool nm ./bin | grep go:itab`). At startup the runtime puts them into a global hash table `itabTable` (starts with 512 slots, grows at 75% load).
- Unknown pairs (asserting to another interface, reflection) → `runtime.getitab` builds the itab at runtime, stores it in the table forever (`persistentalloc`, not garbage collected).
- Interface-to-interface assertions use a small per-call-site cache that is updated rarely (about 1 in 1000 misses) to avoid write overhead.

**Auto-generated wrappers:** if `Work` has a value receiver, the compiler also generates `(*Developer).Work`, which checks the pointer for nil (`runtime.panicwrap`) and calls the value method. The itab points to this wrapper.

RU: Пустой интерфейс — `eface{_type, data}`, непустой — `iface{tab, data}`, где `itab` хранит тип интерфейса, конкретный тип, хеш и таблицу указателей на методы. Вызов метода через интерфейс — косвенный вызов через `itab.fun`. Итабы для известных пар строит компилятор, для остальных — рантайм (`getitab`) с кэшем.

## Key terms
| Term | Meaning (simple EN) | RU |
|---|---|---|
| receiver | the value a method is called on | получатель |
| method set | all methods a type has (decides which interfaces it satisfies) | набор методов |
| implicit satisfaction | type satisfies an interface just by having the methods | неявная реализация |
| dynamic type / value | the concrete type and value stored inside an interface | динамический тип / значение |
| type assertion | `x.(T)` — get the concrete type back | утверждение типа |
| eface / iface | runtime structs of empty / non-empty interfaces | — |
| itab | table linking one interface + one concrete type + method pointers | таблица интерфейса |
| dynamic dispatch | choosing the method at runtime through a pointer | динамическая диспетчеризация |
| devirtualization | compiler replaces an interface call with a direct call | девиртуализация |

## Senior answer (interview-ready)
**Q: How do interfaces work in Go?**
EN: "An interface is a set of methods, and types satisfy it implicitly. At runtime an interface value is two pointers: for an empty interface it's the type and the data; for a non-empty one it's an itab and the data. The itab holds the concrete type and a table of method pointers, so a method call is an indirect call through that table. Storing a non-pointer value usually copies it, sometimes to the heap. An interface is nil only if both words are nil — that's the typed-nil error bug."
RU: «Интерфейс — набор методов, типы реализуют его неявно. В рантайме значение интерфейса — два указателя: у пустого интерфейса это тип и данные, у непустого — itab и данные. В itab лежит конкретный тип и таблица указателей на методы, поэтому вызов метода — косвенный вызов через эту таблицу. Значение-не-указатель при сохранении копируется, иногда в кучу. Интерфейс равен nil, только если оба слова nil — отсюда баг с типизированным nil в error.»

## Gotchas
- Typed nil inside an interface ≠ nil interface.
- Value `T` does not satisfy an interface whose methods have pointer receivers.
- Embedded interface field in a struct is nil by default → panic on call.
- Single-value assertion `x.(T)` panics on mismatch; use comma-ok.
- Every non-pointer value stored in `any` may allocate.

## Recall questions
1. How does Go decide that a type satisfies an interface? When is it checked?
2. Write the compile-time check that `*MemoryStore` implements `Store`.
3. Why can `MemoryStore` (value) fail to satisfy `Store` while `*MemoryStore` works?
4. What prints: `var p *int; var x any = p; fmt.Println(x == nil)`? Why?
5. What two words does `eface` hold? What does `iface` hold instead, and why are both 16 bytes?
6. List the fields of `itab` and what `fun` is for.
7. Asserting `a.(Barker)` where `Barker` is an interface — what is checked?
8. Why can storing an `int` in `any` allocate, but storing `true` doesn't?
9. Explain "accept interfaces, return structs" and "consumer-side interfaces" with Bozorcha's `Store`.
10. What happens when you call a method of an embedded interface field that was never set?

## Links
Practice: `code/golang/04-methods-interfaces/` (to create) · Source: book Ch4 §2, Go spec "Interface types", Go blog "The Laws of Reflection"
