# Go Internals — Chuqur Tushuntirish (O'zbekcha) + Interview Javoblari

Bu sizning eng kuchli qismingiz. Senior darajani shu yerda ko'rsatasiz.

---

# 1. Goroutine va Scheduler (GMP modeli)

## Tushuntirish
Goroutine — bu **thread emas**, u ancha yengil. Bitta goroutine boshida ~2KB stack oladi (OS thread ~1-2 MB). Shuning uchun millionlab goroutine ochish mumkin.

Go scheduler **GMP modeli** bilan ishlaydi:
- **G** (Goroutine) — bajariladigan ish
- **M** (Machine) — haqiqiy OS thread
- **P** (Processor) — logik protsessor, kontekst. Soni = `GOMAXPROCS` (odatda CPU yadrolari soni)

**Qanday ishlaydi:** Har bir P da goroutine'lar navbati (run queue) bor. M (thread) P ni oladi va uning navbatidagi G'larni bajaradi. Agar bitta P bo'sh qolsa, u boshqa P navbatidan ish **o'g'irlaydi** (work stealing).

## Muhim tushunchalar
- **Preemption**: Go 1.14'dan boshlab goroutine'lar **asinxron preempt** qilinadi (signal orqali) — uzun sikl ham boshqa goroutine'larni ochlikda qoldirmaydi.
- **Blocking syscall**: agar G bloklovchi syscall qilsa (masalan fayl o'qish), M o'sha G bilan bloklanadi, lekin P **ajratilib** boshqa M ga beriladi — qolgan goroutine'lar ishlashda davom etadi.
- **Network poller**: tarmoq I/O (socket) bloklanmaydi — Go netpoller (epoll/kqueue) ishlatadi, goroutine "parkovka" qilinadi va data kelganда uyg'otiladi.

## 🔸 Kichik savol 1
**S:** Goroutine va OS thread farqi nima?
**J:** Goroutine yengil (~2KB stack, o'sadi), Go runtime boshqaradi, ko'p goroutine bitta thread'da multipleks qilinadi. OS thread og'ir (~1-2MB), yadro boshqaradi.

## 🔸 Kichik savol 2
**S:** GOMAXPROCS nima?
**J:** Bir vaqtda kod bajaradigan P (logik protsessor) soni. Odatda CPU yadrolari soniga teng.

## 🔸 Kichik savol 3
**S:** Work stealing nima?
**J:** Bo'sh qolgan P boshqa band P ning run queue'sidan goroutine o'g'irlaydi — yuk balanslanadi.

---

## 🎯 INTERVIEW SAVOLLARI — Scheduler

**Q1: How does the Go scheduler work?**
> **Senior answer (EN):** Go uses an M:N scheduler with the GMP model — G is a goroutine, M is an OS thread, P is a logical processor whose count is GOMAXPROCS. Each P has a local run queue of goroutines; an M must hold a P to run Go code. Idle Ps steal work from busy ones. When a goroutine makes a blocking syscall, the runtime detaches the P so other goroutines keep running on another M; network I/O goes through the netpoller instead of blocking a thread. Since Go 1.14 goroutines are asynchronously preemptible, so a tight loop can't starve others.

**Q2: Why are goroutines cheaper than threads?**
> **Senior answer (EN):** A goroutine starts with a tiny ~2KB growable stack and is scheduled in user space by the Go runtime, so creating and switching them avoids kernel overhead. Many goroutines are multiplexed onto a small number of OS threads. OS threads each reserve ~1-2MB and switching them means a kernel context switch, so you can't have millions of them.

---

# 2. Channels va Concurrency Patterns

## Tushuntirish
Channel — goroutine'lar orasida xavfsiz ma'lumot uzatish. Go falsafasi: **"Don't communicate by sharing memory; share memory by communicating."**

- **Unbuffered channel** (`make(chan int)`) — sender qabul qiluvchini kutadi (sinxron, "rendezvous").
- **Buffered channel** (`make(chan int, 5)`) — bufer to'lguncha sender bloklanmaydi.
- **Yopilgan channel** (`close(ch)`) — o'qish darhol zero value + `ok=false` qaytaradi. Yopiq channel'ga yozish → **panic**.
- **nil channel** — undan o'qish/yozish **abadiy bloklanadi** (select'da foydali).

## Muhim patternlar
- **`select`** — bir nechta channel'dan birinchi tayyor bo'lganini tanlaydi. `default` bilan non-blocking.
- **Worker pool** — N ta goroutine bitta jobs channel'dan ish oladi.
- **Fan-out / fan-in** — ishni ko'p goroutine'ga tarqatib, natijani bitta channel'ga yig'ish.
- **Context bilan bekor qilish** — `ctx.Done()` channel orqali goroutine'larni to'xtatish.

## ⚠️ Muhim tuzoqlar
- **Goroutine leak** — goroutine channel'da abadiy kutib qolsa, hech qachon tugamaydi → xotira oqadi. Doim `context` yoki `done` channel bilan chiqish yo'lini bering.
- **Deadlock** — hamma goroutine bir-birini kutsa: `fatal error: all goroutines are asleep - deadlock!`
- **Yopishni kim qiladi?** — faqat **sender** channel'ni yopishi kerak, receiver emas.

## 🔸 Kichik savol 1
**S:** Buffered va unbuffered channel farqi?
**J:** Unbuffered — sender receiver'ni kutadi (sinxron). Buffered — bufer to'lguncha sender bloklanmaydi.

## 🔸 Kichik savol 2
**S:** Yopiq channel'ga yozsa nima bo'ladi?
**J:** Panic. Yopiq channel'dan o'qish esa zero value + `ok=false` qaytaradi.

## 🔸 Kichik savol 3
**S:** Goroutine leak nima va qanday oldini olasan?
**J:** Goroutine abadiy bloklanib qolsa (masalan hech kim o'qimaydigan channel'ga yozmoqchi bo'lsa). Oldini olish: `context` bilan bekor qilish, buffered channel, yoki `select` + `done`.

---

## 🎯 INTERVIEW SAVOLLARI — Channels

**Q1: What's the difference between buffered and unbuffered channels?**
> **Senior answer (EN):** An unbuffered channel is a synchronization point — the send blocks until a receiver is ready, so it's a rendezvous between two goroutines. A buffered channel lets the sender proceed until the buffer fills, decoupling sender and receiver up to the buffer size. I pick unbuffered when I need a handshake guarantee, buffered when I want to smooth bursts or limit concurrency.

**Q2: How do you prevent goroutine leaks?**
> **Senior answer (EN):** A goroutine leaks when it blocks forever on a channel that will never be satisfied. I always give goroutines a way out: pass a `context.Context` and select on `ctx.Done()`, make sure every channel has a reader if it has a writer, and close channels from the sender side so receivers can drain. In tests I use tools like goleak to catch leaks.

**Q3: Who should close a channel?**
> **Senior answer (EN):** The sender, never the receiver — because sending on a closed channel panics. When there are multiple senders, I coordinate with a separate done signal or a sync.WaitGroup and close once all senders are finished.

---

# 3. Memory: Stack vs Heap, Escape Analysis

## Tushuntirish
Go kompilyatori **escape analysis** qiladi — o'zgaruvchini **stack**да yoki **heap**да joylashtirishni hal qiladi.

- **Stack** — tez, funksiya tugaganда avtomatik tozalanadi, GC tegmaydi.
- **Heap** — GC boshqaradi, sekinroq. O'zgaruvchi funksiyadan "qochib chiqsa" (escape) — masalan pointer qaytarilsa — heap'ga tushadi.

`go build -gcflags='-m'` bilan qaysi o'zgaruvchi escape qilishini ko'rish mumkin.

**Nega muhim?** Heap allocation ko'p bo'lsa GC yuki oshadi. Senior dev keraksiz allocation'larni kamaytiradi (masalan slice'ni oldindan `make([]T, 0, n)` bilan sig'im berib yaratish).

## 🔸 Kichik savol 1
**S:** O'zgaruvchi qachon heap'ga tushadi?
**J:** Funksiyadan qochib chiqsa — pointer qaytarilsa, closure'da ushlansa, interface'ga solинsa, yoki hajmi noma'lum bo'lsa.

## 🔸 Kichik savol 2
**S:** Escape analysis nima uchun kerak?
**J:** Kompilyator o'zgaruvchini stack (arzon) yoki heap (GC yuki) ga qo'yishni hal qiladi — performansga ta'sir qiladi.

---

# 4. Garbage Collector (GC)

## Tushuntirish
Go GC — **concurrent, tri-color mark-and-sweep**, past latency uchun mo'ljallangan.

- **Concurrent** — asosan dastur ishlab turganда parallel ishlaydi, "stop-the-world" (STW) pauzalari juda qisqa (odatda <1ms).
- **Tri-color** (oq/kulrang/qora) — jonli obyektlarni belgilash algoritmi.
- **Write barrier** — GC belgilash paytida o'zgargan pointerlarni kuzatadi.
- **`GOGC`** — GC qachon ishga tushishini boshqaradi (default 100 = heap ikki barobar o'sganда). Katta qilsa GC kam ishlaydi lekin ko'proq xotira ishlatadi.

## 🔸 Kichik savol 1
**S:** Go GC qanday turdagi?
**J:** Concurrent tri-color mark-and-sweep, past latency uchun. STW pauzalari juda qisqa.

## 🔸 Kichik savol 2
**S:** GOGC nima qiladi?
**J:** GC chastotasini boshqaradi. Default 100 — heap oldingi jonli hajmdan 2x o'sganда GC ishlaydi.

---

## 🎯 INTERVIEW SAVOLLARI — Memory & GC

**Q1: Explain Go's garbage collector.**
> **Senior answer (EN):** Go uses a concurrent, tri-color mark-and-sweep collector tuned for low latency. Most of the marking runs concurrently with the program, so stop-the-world pauses are typically sub-millisecond. A write barrier tracks pointer changes during marking so the collector stays correct. GOGC controls the trade-off: the default 100 triggers a cycle when the heap doubles relative to live size; raising it reduces GC frequency at the cost of memory. When latency matters I profile allocations and reduce heap pressure rather than fighting the GC.

**Q2: How do you reduce GC pressure in Go?**
> **Senior answer (EN):** Reduce allocations that escape to the heap. I preallocate slices and maps with a known capacity, reuse buffers with `sync.Pool`, avoid unnecessary pointer indirection, and use `go build -gcflags=-m` and pprof to find hot allocations. Value types that stay on the stack never touch the GC at all.

**Q3: Stack vs heap — how does Go decide?**
> **Senior answer (EN):** The compiler does escape analysis. If a variable's lifetime is bounded by the function, it stays on the stack, which is cheap and GC-free. If it escapes — returned by pointer, captured by a closure, stored in an interface, or of unknown size — it goes to the heap. I check with `-gcflags=-m`.

---

# 5. Context, sync primitives, common bugs

## Tushuntirish
- **`context.Context`** — bekor qilish (cancellation), deadline, timeout va request-scoped qiymatlarni goroutine'lar bo'ylab tarqatadi. Doim birinchi argument: `func(ctx context.Context, ...)`.
- **`sync.Mutex`** — mutual exclusion, umumiy holatni himoya qiladi.
- **`sync.RWMutex`** — ko'p reader / bir writer.
- **`sync.WaitGroup`** — goroutine'lar tugashini kutish.
- **`sync.Once`** — bir marta bajarish (masalan lazy init).
- **`atomic`** — lock'siz oddiy operatsiyalar (counter).

## ⚠️ Klassik buglar
- **Loop variable capture** (Go 1.22'gacha): `for _, v := range items { go func(){ use(v) }() }` — hamma goroutine oxirgi `v` ni ko'rardi. Go 1.22'dan har iteratsiya yangi `v`. (Bilib qo'ying — versiyaga bog'liq.)
- **Data race** — ikki goroutine bir xotiraga sinxronizatsiyasiz kirsa. `go run -race` bilan aniqlanadi.
- **Nil map yozish** — panic. Map'ni `make` bilan yaratish kerak.

## 🎯 INTERVIEW SAVOLLARI — Context & sync

**Q1: What is context used for?**
> **Senior answer (EN):** Context propagates cancellation, deadlines, and request-scoped values across API boundaries and goroutines. When a request is cancelled or times out, everything downstream selecting on `ctx.Done()` stops promptly, which prevents wasted work and goroutine leaks. I pass it as the first parameter and never store it in a struct.

**Q2: Mutex vs channel — when to use which?**
> **Senior answer (EN):** I use a mutex to protect shared state that's read and written in place — it's simple and fast for guarding a struct field or a map. I use channels to transfer ownership of data between goroutines or to coordinate lifecycle and pipelines. The rule of thumb: mutex for state, channels for communication.

**Q3: How do you detect a data race?**
> **Senior answer (EN):** Run tests and the binary with the race detector, `go test -race` / `go run -race`. It instruments memory accesses and reports concurrent unsynchronized read/write. In CI I always run the race detector on the test suite.
