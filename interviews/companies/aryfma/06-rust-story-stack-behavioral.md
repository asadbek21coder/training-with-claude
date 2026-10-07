# Rust Story + Stack Talking Points + Behavioral (Aryfma)

Bu fayl — sizni "kuchli Go dev, Rust'ga ochiq, jamoaga mos" qilib ko'rsatadi.
Aryfma post'i AYNAN shuni so'raydi: "Rust YOKI kuchli Kotlin/KMP + Rust'ni o'rganishga ochiqlik".

---

# 1. RUST HIKOYASI (halol pozitsiya)

## Asosiy qoida
**Rust'ni bilaman deb yolg'on aytmang.** Buning o'rniga: "Kuchli engineer, Rust modelini tushunadi, o'rganishga tayyor." Bu ularning talabiga AYNAN mos va yolg'ondan ko'ra kuchliroq.

## Sizning 5 daqiqalik hikoyangiz (yodlang)
> **(EN)** "My core is 3 years of Go and Postgres, building backend systems. I haven't shipped Rust in production, but I understand its model — ownership, borrowing, and why the borrow checker exists. Coming from Go, a lot transfers: I already think about concurrency, memory, and systems-level tradeoffs. Rust just moves those safety checks to compile time instead of runtime. I'm genuinely excited to learn it — the ownership model is exactly the kind of rigor I like."

## Rust'ning MOHIYATI (agar so'rashsa — konsept darajasida)

### Ownership (egalik)
Har bir qiymatning bitta **egasi** (owner) bor. Egasi scope'dan chiqsa — qiymat avtomatik **drop** qilinadi (xotira bo'shatiladi). GC yo'q, `free()` ni qo'lda chaqirmaysan — kompilyator hal qiladi.
> Go'da: GC xotirani tozalaydi. Rust'da: egalik qoidasi compile-time'да buni hal qiladi — GC pauzasi yo'q.

### Borrowing (qarzga olish)
Qiymatni ko'chirmasdan **referens** olish:
- `&T` — **shared borrow** (ko'p bo'lishi mumkin, faqat o'qish)
- `&mut T` — **mutable borrow** (faqat bitta, o'zgartirish mumkin)
- Qoida: **bir vaqtda ko'p reader YOKI bitta writer** — ikkalasi emas. Bu **data race'ni compile-time'да** oldini oladi.
> Go'da: data race'ни `-race` bilan runtime'да topasan. Rust'da: kompilyator umuman kompilyatsiya qilmaydi.

### Move semantics
`let b = a;` — agar `a` Copy bo'lmasa, egalik `b`ga **ko'chadi**, `a` endi ishlatib bo'lmaydi. Ikki egadan himoya.

### Lifetimes (faqat g'oyasini bil)
Referens o'zi ko'rsatgan ma'lumotdan **uzoq yashay olmaydi**. Kompilyator buni tekshiradi (dangling pointer yo'q).

### Result / Option (xato boshqaruv)
- `Option<T>` — qiymat bor (`Some`) yoki yo'q (`None`). **null yo'q.**
- `Result<T, E>` — muvaffaqiyat (`Ok`) yoki xato (`Err`). **exception yo'q.**
- `?` operatori — xatoni yuqoriga uzatadi (Go'dagi `if err != nil { return err }` ga o'xshaydi, lekin bir belgi).
> Go'da: `if err != nil`. Rust'da: `let x = foo()?;` — o'xshash falsafa, xatolar qiymat.

## 🔑 Go → Rust ko'prigi (intervyuda oltin jumla)
> "Go taught me to respect concurrency and simplicity. Rust takes the same concerns — memory safety, data races — and enforces them at compile time. The `Result`/`?` pattern even feels like Go's explicit error handling. So the mental model isn't foreign to me, just stricter."

## 🎯 Rust interview savollari (agar chuqurlashsa)

**Q: What is ownership?**
> Every value has a single owner; when the owner goes out of scope the value is dropped and its memory freed — no garbage collector, no manual free. This gives deterministic cleanup with no runtime overhead.

**Q: What's the difference between `&T` and `&mut T`?**
> `&T` is a shared, read-only borrow and you can have many at once. `&mut T` is an exclusive, mutable borrow and only one can exist at a time. The compiler enforces "many readers XOR one writer," which eliminates data races at compile time.

**Q: How does Rust handle errors without exceptions?**
> With the `Result<T, E>` and `Option<T>` enums — errors and absence are values, not thrown exceptions. The `?` operator propagates errors upward concisely. It's very close in spirit to Go's explicit error returns, just enforced by the type system.

**Q: Why would a company choose Rust over Go?**
> When they need no GC pauses, predictable low-latency, tight memory control, or want to catch data races and memory bugs at compile time — common for systems software, desktop/embedded apps, and performance-critical services. That fits a Slint + Tokio native app well.

---

# 2. STACK TALKING POINTS (har biri 1-2 jumla, halol)

Bu texnologiyalarни chuqur bilmasangiz ham, **nima ekanini** bilib turing.

## Tokio
> **(EN)** "Tokio is the async runtime for Rust — it's what lets you write `async/await` code and run many concurrent tasks on a thread pool, similar in purpose to Go's goroutine scheduler, but explicit. Futures in Rust are lazy — they don't run until you `.await` them or spawn them with `tokio::spawn`."
- Go bilan bog': goroutine ≈ tokio task; channel ≈ tokio channel.

## Slint
> **(EN)** "Slint is a declarative UI toolkit for Rust — you describe the UI in a markup-like language and it reactively updates when your data changes. Conceptually similar to Jetpack Compose or SwiftUI: declarative and reactive, targeting desktop and embedded."
- Aryfma buni desktop/embedded native app UI uchun ishlatadi.

## SQLDelight
> **(EN)** "SQLDelight generates type-safe Kotlin APIs from your SQL — you write the SQL, it generates the code, so queries are checked at compile time. It's SQL-first, which fits well with my Postgres background — I'm comfortable writing and optimizing raw SQL."
- 🔑 Bu sizning Postgres kuchingizga bog'lanadi! "SQL-first — men SQL yozishni yaxshi ko'raman."

## Kotlin Multiplatform (KMP)
> **(EN)** "KMP lets you share business logic across platforms — Android, iOS, desktop — while keeping native UI per platform. You write the core once in Kotlin and target multiple platforms."

## Git + CI/CD
> **(EN)** "I use Git daily — feature branches, PRs, code review. For CI/CD I've set up pipelines that run tests, the race detector, and linters on every push, and automate builds and deploys. This is core to how I work."
- Bu sizning kuchli, halol qismingiz — ishonch bilan gapiring.

## REST Integration
> **(EN)** "I've built and consumed REST APIs extensively in Go — designing endpoints, handling auth, timeouts, retries with idempotency, and error handling. Integrating a Rust/Kotlin client with REST backends is the same discipline."

---

# 3. BEHAVIORAL + CULTURE (Aryfma alohida so'raydi)

Aryfma talablari: **AI bilan samarali ishlash, jamoada ochiq muloqot, onsite ishlash.**

## AI bilan ishlash (ular AYNAN so'raydi)
> **(EN)** "I use AI tools daily as a force multiplier — for boilerplate, exploring unfamiliar APIs, writing tests, and rubber-ducking design. But I review everything critically; I treat AI output as a draft to verify, not truth. It's especially useful for ramping up on new tech like Rust quickly."
- 🔑 Halol va zamonaviy — aynan ular xohlagan javob.

## Onsite / jamoada ochiq muloqot
> **(EN)** "I work well onsite — I like direct collaboration, whiteboard discussions, and fast feedback. I communicate openly: I ask questions early, share blockers instead of hiding them, and give and take code review honestly."

## "Nega Rust / nega bu kompaniya?"
> **(EN)** "I've spent 3 years going deep in Go and Postgres, and I want to grow into systems-level work where correctness and performance really matter — Rust is exactly that direction. Your stack — Rust, Tokio, Slint — is building real native software, and I want to be part of a team that takes engineering rigor seriously."

## STAR metodi (behavioral savollar uchun)
Har bir "tell me about a time..." savolига:
- **S**ituation — vaziyat
- **T**ask — vazifa/muammo
- **A**ction — nima qildingiz (SIZ, "biz" emas)
- **R**esult — natija (raqam bilan bo'lsa yaxshi)

Tayyorlab keling (Go loyihalaringizdan):
1. Qiyin bug topgan/tuzatgan vaqtingiz (masalan race condition, slow query)
2. Performance yaxshilagan vaqtingiz (masalan N+1 tuzatish, index qo'shish)
3. Jamoada kelishmovchilik / qaror qabul qilish

## 🎯 Behavioral savollar + javob yo'nalishi

**Q: Tell me about a challenging bug.**
> STAR bilan: bitta aniq bug (masalan Postgres slow query yoki Go goroutine leak), qanday topdingiz (EXPLAIN / pprof / -race), qanday tuzatdingiz, natija (latency X% tushdi).

**Q: How do you handle disagreement in a team?**
> "I focus on data and tradeoffs, not ego. I state my reasoning, listen to theirs, and if we're stuck I suggest a small experiment or defer to the person who owns that area. Disagreeing openly but respectfully is healthier than silent resentment."

**Q: Why should we hire you if you don't know Rust?**
> "Because the hard part of engineering isn't a specific syntax — it's systems thinking, debugging, and shipping reliable software, which I've done for 3 years in Go and Postgres. Rust's model overlaps with how I already think about memory and concurrency. Your own posting welcomes strong engineers open to Rust, and I learn fast — especially with AI tools to accelerate ramp-up."

---

# 4. INTERVYUCHIGA BERADIGAN SAVOLLARINGIZ (oxirida so'rashadi)

Yaxshi savol berish — qiziqish va senior signal:
1. "What does the product do, and who are the users?" (mahsulotni tushunish)
2. "Why did the team choose Rust + Slint for this?"
3. "What does the first 3 months look like for a new engineer here?"
4. "How does the team use AI tools in practice day-to-day?"
5. "What does the CI/CD and code review process look like?"

---

# 5. YAKUNIY ESLATMALAR (intervyu kuni)
- ✅ Halol bo'ling — bilmasangiz "I haven't used X, but here's how I'd approach it."
- ✅ Ovoz chiqarib fikrlang (coding va design'да).
- ✅ Har javobни **consequence** bilan bog'lang (MVCC → VACUUM → bloat).
- ✅ Go tajribangizni Rust/stack'ga ko'prik qiling.
- ✅ Yaxshi uxlang. Charchagan miya 3 soat cramming'dan yomon.
- ✅ Kompaniyani tekshiring: yuridik nom (orginfo.uz / soliq.uz), ofisni ko'ring, pul so'rasa — qizil bayroq.
