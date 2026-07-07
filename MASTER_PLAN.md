# Master Plan — Senior Backend Engineer Preparation

**Engineer:** Asadbek Egambergenov | **Stack:** Go, PostgreSQL, Kafka, gRPC  
**Target:** Senior Backend Engineer at strong product companies  
**Time:** 1 hour/day, 5-6 days/week  
**Status:** Actively interviewing — urgency is real

---

## How this system works

Every topic has **3 cycles**:
- **C1 — Know it:** understand the problem, concept, terms. Can explain clearly.
- **C2 — Deep understanding:** know internals, tradeoffs, alternatives. Can discuss design decisions.
- **C3 — Own it:** implement from scratch. Know how real systems do it. Full ownership.

**Tracks run in parallel:**
- **Track A** — DS & Algorithms (daily)
- **Track B1** — System Design (daily)
- **Track B2** — Go Internals (3x/week)
- **Track B3** — PostgreSQL Deep (2x/week)
- **Track C1** — English (15 min, woven into every session)
- **Track C2** — Russian (when practicing)
- **Track D** — Mock Interviews (weekly, real questions from real-interview-questions/)

**Progress tools:**
- `SPRINT.md` — daily topic checklist (granular)
- `ROADMAP.md` — full topic map with best sources
- `progress_tracker.html` — open in browser, log each session visually
- `mock-interviews/weak-points.md` — gaps found in mock sessions

---

## Section 0 — Right Now (First 30 Days)

You are actively interviewing. Before the 12-month plan, these are the only things that matter:

### Minimum viable topics (master these first)

| Priority | Topic | Why |
|---|---|---|
| 🔴 Must | Goroutine leak + data race + mutex | Asked in almost every Go interview |
| 🔴 Must | Error wrapping (errors.Is, errors.As) | Go-specific, commonly tested |
| 🔴 Must | Context — cancellation, propagation | Foundational Go concept |
| 🔴 Must | ACID + isolation levels | Every backend interview |
| 🔴 Must | N+1 + indexing + EXPLAIN ANALYZE | Every backend interview |
| 🔴 Must | Panic vs error | Go-specific, always asked |
| 🔴 Must | Graceful shutdown | Go-specific, senior signal |
| 🟡 Should | Functional options + Factory pattern | Alif asked these, common in Go interviews |
| 🟡 Should | JWT + OAuth basics | Security companies always ask |
| 🟡 Should | Kafka at-least-once + DLQ | You used it, you'll be asked |
| 🟡 Should | System design framework | URL shortener or notification service |
| 🟢 Nice | GMP scheduler | Advanced Go, senior+ signal |
| 🟢 Nice | Worker pool + fan-out patterns | Shows Go idiom knowledge |

### Daily schedule for days 1–30 (1 hour/day)

```
40 min — Technical topic from priority list above
15 min — English: read one engineering blog post, note 3 new phrases
5 min  — Log session in progress_tracker.html
```

**3x per week:** replace technical session with mock interview from `real-interview-questions/`

### Evening before interview
1. Read `mock-interviews/quick-reference.md` — full review
2. Read `russian/interview_phrases.md` — project answers
3. Sleep early — performance drops 30% with bad sleep

### Morning of interview
1. Read your project descriptions in `projects.txt` — internalize them
2. Review STAR stories (testing 87%, PPP Excel→PDF, IMV Archive alone)
3. 10 min: practice saying your intro out loud

### Red flags — what signals mid-level thinking
- Saying "it depends" without immediately explaining what it depends on
- Not asking clarifying questions before answering system design
- Answering "I used X" without explaining WHY you chose X over alternatives
- Saying "I don't know" and stopping — always add "but here's how I'd think about it"
- Not mentioning trade-offs when discussing architecture
- Explaining WHAT code does instead of WHY you designed it that way

---

## Section 1 — Priority Map

Given 1 hour/day, you cannot master everything equally. Be ruthless.

### Core — must master (appear in virtually every senior interview)

1. **Go concurrency** — goroutines, channels, sync primitives, context, graceful shutdown
2. **Error handling** — wrapping, sentinel errors, custom types, logging at right layer
3. **PostgreSQL** — EXPLAIN ANALYZE, indexes, ACID, transactions, N+1, isolation levels
4. **System design fundamentals** — scalability, caching, message queues, CAP theorem
5. **Architecture patterns** — layered arch, dependency injection, repository pattern, modular monolith
6. **Testing** — unit + integration, mocks vs stubs, testcontainers, table-driven tests
7. **DS & Algorithms** — two pointers, binary search, BFS/DFS, sliding window, hash maps
8. **Your projects** — PPP, Iman Invest, IMV Archive — deep fluency on every decision

### Important — should know (asked at good companies)

- Redis basics (cache-aside, TTL, distributed lock)
- Kafka reliability (at-least-once, DLQ, idempotency, consumer groups)
- HTTP/2, gRPC vs REST tradeoffs
- JWT + OAuth 2.0 flow
- Docker + CI/CD basics
- Dynamic programming (medium difficulty)
- Go internals: GMP scheduler, GC, escape analysis
- OWASP Top 10 basics

### Advanced — FAANG-specific (defer unless targeting those companies)

- Distributed consensus (Raft, Paxos)
- B-tree internal implementation
- Lock-free data structures
- Full system design at scale (10M+ users)
- JVM internals, Kotlin coroutines

### Single highest ROI

**Go concurrency** — it's your primary language, companies expect deep knowledge here. A strong answer on goroutine lifecycle, channel mechanics, and context propagation alone separates you from most Go candidates.

---

## Section 2 — 12-Month Timeline

### Phase 0: Triage (Now — Month 1)
**Goal:** Survive current interviews. Use `quick-reference.md` + mock sessions.
- C1 on: all Core topics above
- Mock interviews: 3x/week from `real-interview-questions/`
- English: 15 min/day — read one Go blog post, note phrases
- Weak points → `mock-interviews/weak-points.md` → study immediately

### Phase 1: Foundation (Months 1–3)
**Goal:** C2 on all Core topics. Can discuss tradeoffs, not just definitions.

| Week | Track A | Track B | Language |
|---|---|---|---|
| 1–2 | Array, HashMap, Linked List | Go concurrency deep | English: filler phrases, STAR structure |
| 3–4 | Stack, Queue, Binary Tree | PostgreSQL: EXPLAIN deep | English: technical vocabulary system |
| 5–6 | Graphs, BFS/DFS | System design: scalability | Russian: cases + verb aspect |
| 7–8 | Binary Search, Two Pointers | Go error handling deep | English: shadowing (Go Time podcast) |
| 9–10 | Heap, Sliding Window | Kafka reliability | English: written messages practice |
| 11–12 | DP basics | Architecture patterns | Russian: interview phrases |

**Milestone:** Can answer any Alif-style interview question without hesitation. Score 4+ on mock sessions.

### Phase 2: Depth (Months 3–6)
**Goal:** C2 on Important topics. C3 on top 5 Core topics.

- Build: worker pool, graceful shutdown, context propagation — from scratch
- Read: Go Blog (official), Ardan Labs scheduler posts
- System design: design 5 real systems (URL shortener, notification service, rate limiter, auth service, file storage)
- LeetCode: 30 Medium problems — two pointers, sliding window, BFS/DFS, binary search
- Mock interviews: weekly, all from `real-interview-questions/`

**Milestone:** Can design a medium-complexity system (notification service) in 45 min with tradeoffs. Solve a Medium LeetCode in 25 min.

### Phase 3: Pressure (Months 6–9)
**Goal:** Perform under pressure. C3 on all Core. C2 on Important.

- Mock interviews: 2x/week, simulate real pressure (timer, camera on, spoken English)
- LeetCode: 20 Hard problems — DP, graphs, complex two pointers
- Go internals: C3 — read runtime source code (scheduler: src/runtime/proc.go, GC: src/runtime/mgc.go)
- Build: mini distributed system (two services + Kafka + shared DB)
- English: full mock in English, record yourself, review

**Milestone:** Can conduct full 45-min technical interview in English without preparation. Go internals questions feel comfortable.

### Phase 4: Fluency (Months 9–12)
**Goal:** Senior-level confidence. Minimal hesitation. Polish.

- Real interviews: 2-3 per week
- Post-interview debrief: fill template after every interview
- Weak points from real interviews → study immediately
- English: Level 4 — can discuss system design fluently without translating
- Build portfolio project (see Section 3)

**Milestone:** Receive senior engineer offer. Interview readiness score in `progress_tracker.html` above 80%.

---

## Section 3 — Go Internals Deep Track (B2)

This is separate because Go is your primary language — interviewers expect more depth.

### GMP Scheduler
- **C1:** Know G (goroutine), M (OS thread), P (processor). Work stealing. Preemption.
- **C2:** How P limits parallelism (GOMAXPROCS). Why goroutines are cheap (2KB stack). How blocking I/O is handled without blocking the OS thread.
- **C3:** Read `src/runtime/proc.go` — understand `schedule()`, `findRunnable()`, `park_m()`.
- **Source:** Ardan Labs — "Scheduling in Go" (3-part series)

### Garbage Collector
- **C1:** Tri-color mark and sweep. Concurrent. Short STW pauses. Three phases.
- **C2:** Write barriers — why they exist. GOGC tuning — when to change it. GC pressure from allocations.
- **C3:** Read `src/runtime/mgc.go`. Run `GODEBUG=gctrace=1` on a real program.
- **Source:** Go GC Guide (tip.golang.org/doc/gc-guide)

### Memory Model + Escape Analysis
- **C1:** Stack vs heap. Compiler decides — escape analysis. Heap allocation is slower (GC pressure).
- **C2:** Read escape analysis output: `go build -gcflags="-m"`. Understand what causes a variable to escape (returned pointer, interface boxing, closures).
- **C3:** Benchmark a function before/after reducing allocations. Measure with `-benchmem`.
- **Source:** Go official — "The Go Memory Model"

### Channel Internals
- **C1:** `hchan` struct. Send/receive mechanics. Select with multiple cases.
- **C2:** How buffered channel is implemented (circular queue). Why closing broadcasts. How select picks randomly when multiple cases ready.
- **C3:** Read `src/runtime/chan.go` — `chansend()`, `chanrecv()`.
- **Source:** Go By Example + runtime source

### Interface Internals
- **C1:** `iface` (interface with methods) vs `eface` (empty interface). Dynamic dispatch.
- **C2:** Typed nil trap — interface holding a nil pointer is not nil. Cost of interface boxing.
- **C3:** Use `go tool compile -S` to see assembly for interface call vs direct call.
- **Source:** Go Blog — "The Laws of Reflection"

---

## Section 4 — Language Plan (C1 + C2 Tracks)

**Rule:** 15 minutes per day, integrated into technical session — not separate.

### English (70% of language effort)

**Daily (15 min):**
- Read one technical blog post (Go Blog, Cloudflare, Stripe Engineering)
- Note 2-3 new technical phrases in `english/vocabulary.md`
- Once per week: practice saying your project description out loud in English

**Shadowing (3x/week, 10 min):**
- Source: Go Time podcast (changelog.com/gotime) — engineering English specifically
- Method: listen 30 sec → pause → repeat exactly, same rhythm and intonation
- Target: reduce hesitation when speaking under interview pressure

**Interview English — must memorize these phrases:**

| Situation | Phrase |
|---|---|
| Thinking | "Let me think through this..." / "Give me a moment to reason about this..." |
| Clarifying | "Before I answer, can I ask — what's the scale we're designing for?" |
| Don't know | "I haven't used that directly, but here's how I'd approach learning it..." |
| Tradeoff | "The tradeoff here is..." / "That approach has the benefit of X but the cost of Y..." |
| Disagreeing | "That's a valid approach. My instinct would be X because..." |
| Stuck | "I'm thinking out loud here — one approach would be..." |
| Mistake | "Actually, let me revise that — I think the better answer is..." |

**Transfer errors to eliminate (Uzbek → English):**
- Articles: "the database" not "database" — always check a/an/the before nouns
- Prepositions: "depends ON" not "depends from", "work ON a project" not "work in project"
- Verb tense: past for completed ("I built X") vs present perfect for experience ("I've worked with...")
- Word order: adjective before noun in English ("slow query" not "query slow")

### Russian (30% of language effort)

**When to practice:** Only when explicitly switching to Russian mode.

**Weak points (from earlier sessions):**
- Genitive case for possession (сервера, not "в сервере")
- Verb aspect: делать (imperfective, ongoing) vs сделать (perfective, completed)
- Prepositions: в+Prepositional for location, на+Prepositional for surfaces/events

**All Russian answers:** `russian/interview_phrases.md`
**Russian grammar reference:** `russian/grammar.md`

---

## Section 5 — LeetCode Strategy

**Not about solving as many problems as possible. About internalizing patterns.**

### Patterns by priority for backend interviews

| Pattern | Problems to solve | Mastery criteria |
|---|---|---|
| Two Pointers | 8–10 | Solve Medium in 15 min without hints |
| Binary Search | 6–8 | Know all 3 variants (exact, left bound, right bound) |
| Sliding Window | 6–8 | Can identify when to use it in 2 min |
| BFS/DFS | 10–12 | Solve graph problem confidently in 20 min |
| Hash Map | 6–8 | Already strong — maintain |
| Dynamic Programming | 10–15 | Recognize subproblem structure, build table |
| Heap/Top-K | 5–6 | Can implement min-heap pattern |
| Intervals | 4–5 | Sort + sweep approach |

**Weekly LeetCode schedule (within 1 hour/day):**
- 3 days/week: 1 problem per day (25–30 min)
- Track time: if Medium takes > 35 min, add it to weak-points
- After solving: explain complexity out loud in English

**Mastery signal:** Solved a Medium problem you haven't seen before, in under 30 min, explained the approach clearly in English.

---

## Section 6 — Mock Interview System (Track D)

**Source:** `real-interview-questions/` folder — real questions from real companies.

**Session structure:**
1. Pick a company or topic from `real-interview-questions/`
2. Claude asks questions one by one
3. After each answer: score (1-5) + what was missing + model answer
4. Weak areas → `mock-interviews/weak-points.md`
5. Session saved to `mock-interviews/sessions/<company>-<date>.md`

**Weak point workflow:**
- Score 1-2 → 🔴 Critical → study within 24 hours
- Score 3 → 🟡 Needs work → study within 3 days
- Score 4-5 → 🟢 Solid → review monthly

**Post-real-interview debrief (fill immediately after interview):**
```
Company: 
Date:
Questions asked:
What I answered well:
What I didn't know:
What I should study:
Overall feeling (1-5):
```
Save to: `mock-interviews/sessions/real-<company>-<date>.md`

---

## Section 7 — Progress Tracking

**Daily log:** `progress_tracker.html` — open in browser, log every session.
- Topic name
- Category (Engineering / LeetCode / Mock / Language)
- Cycle reached (C1/C2/C3)
- Confidence (1-5 stars)

**Weekly review (every Sunday, 15 min):**
- Check `mock-interviews/weak-points.md` — any 🔴 not studied?
- Check `progress_tracker.html` — streak, readiness score
- Plan next week's topics
- Update SPRINT.md checkboxes

**Readiness score formula (in progress_tracker.html):**
```
Score = average of (cycle/3 × 0.5 + confidence/5 × 0.5) across all topics
< 20%  = Early days
< 40%  = Building momentum
< 60%  = Solid foundation
< 80%  = Strong, start mock interviews weekly
≥ 80%  = Interview-ready
```

---

## Section 8 — Canonical Resources

**Books (only the essentials):**
- "Designing Data-Intensive Applications" — Kleppmann → System Design bible. Read chapters 1-3, 5, 7 first.
- "The Go Programming Language" — Donovan & Kernighan → Go bible. Read when doing C2/C3 on Go internals.
- "Database Internals" — Petrov → PostgreSQL deep. Read when doing C3 on PostgreSQL.

**Free resources that are genuinely the best:**
- NeetCode.io — best structured LeetCode learning
- ByteByteGo — best system design visual explanations
- Ardan Labs blog — best Go internals articles
- Use The Index, Luke (use-the-index-luke.com) — best PostgreSQL indexing resource
- Go official blog (go.dev/blog) — authoritative, concise
- System Design Primer (GitHub) — comprehensive free resource

**For English specifically:**
- Go Time podcast (changelog.com/gotime) — engineering English, native Go developers talking
- Cloudflare blog — technical writing at its best, excellent vocabulary source

---

## Section 9 — Failure Mode Handling

**Common failure 1: Falling behind schedule**
Protocol: cut scope, don't fall behind on consistency. One 20-min session beats skipping. Drop C3 goals temporarily, maintain C1-C2 momentum.

**Common failure 2: Stuck on a topic for 2+ weeks**
Protocol: move on. Mark it 🟡 in SPRINT.md, continue. Come back in cycle 2. Understanding sometimes comes after seeing related concepts, not from pushing harder on one thing.

**Common failure 3: Motivation drops after failed interview**
Protocol: do a 30-min mock the same evening. Fill the post-interview debrief immediately. Convert failure to specific study target — this is the most important skill. Failed interviews are the highest-quality data you'll ever get.

**When to restructure the plan:**
- If the same gap appears in 3+ real interviews → it becomes 🔴 critical priority regardless of where it sits in the plan
- If a company consistently asks topics not in your plan → add them to Section 1 Core
- If readiness score hasn't moved in 3 weeks → you're not tracking weak points properly

---

## Section 10 — Optional: Language Expansion (After Month 9)

**Do not start this section until:** readiness score > 70% OR you are past month 9.
This is not a core requirement. It is an investment in thinking better, not collecting languages.

**The goal is not syntax. The goal is paradigm.**

Each language below was invented because existing languages handled a specific problem poorly. You are not learning to write production code in these — you are learning *what problem each was designed to solve* and *how that changes the way you think about your own systems*.

**Recommended order:** Rust → Clojure → TypeScript/Node → SQL (ongoing) → C (one-time)

---

### Module 1: Rust (3–4 weeks)

**The core idea:**
Rust was invented to solve memory safety without garbage collection. Every other language either uses GC (Go, Java, Python) or requires manual memory management (C, C++). Rust's answer: ownership. Every value has exactly one owner. When the owner goes out of scope, the value is freed. No GC needed, no double-free bugs possible. The compiler enforces this at compile time.

The single concept to internalize: **ownership and borrowing**. A value can be owned by one place, or borrowed immutably by many, or borrowed mutably by exactly one. These rules are checked at compile time — not runtime.

**What Go does instead:** Go uses a garbage collector. This is simpler to write, but has costs: GC pauses (even if small), heap allocation pressure, you cannot control when memory is freed. Go hides memory management from you — Rust makes it explicit.

**The Go connection:**
- Before Rust: "GC is just magic that frees memory." After Rust: "I understand exactly what Go's GC is doing — tracking live pointers, sweeping dead ones — and I know what the latency cost is."
- Before Rust: "Goroutine leak is a logic bug." After Rust: "In Rust this would be a compile error. Go chose runtime simplicity over compile-time safety."
- Before Rust: "I pass slices around without thinking." After Rust: "I now think about who owns data and whether shared access is safe."

**In a senior interview:** You can say — "Go's GC simplifies memory management at the cost of non-deterministic pauses and heap pressure. In latency-critical paths we minimize allocations for this reason. Rust solves this differently through ownership — no GC needed — but at the cost of a steeper learning curve and more explicit code."

**Schedule:**
- Week 1: Read "The Rust Book" chapters 1-4 only (ownership, borrowing, lifetimes). Stop at chapter 5.
- Week 2: Read chapters 15-16 (smart pointers, concurrency). See how Rust prevents data races at compile time.
- Week 3-4: Write one small program — a concurrent counter — in Rust. Compare to Go equivalent.
- Stop here. Do not go further.

**Enough when:** You can explain ownership and borrowing in plain English and identify three places in your Go code where the GC is making a tradeoff Rust would not need to make.

**Canonical resource:** "The Rust Programming Language" (doc.rust-lang.org/book) — chapters 1-4 and 15-16 only. Free online. Skip all other chapters.

---

### Module 2: Clojure (2–3 weeks)

**The core idea:**
Clojure was invented by Rich Hickey who believed that most complexity in software comes from mutable state — state that changes over time and is shared between different parts of the system. His answer: make everything immutable by default. Data structures cannot be changed — you create new versions instead. State change is explicit and controlled.

The single concept to internalize: **immutability as a default changes how you think about correctness**. In Go, a bug where two goroutines race to write a variable is a runtime crash. In Clojure, that scenario is structurally impossible — the data cannot be modified, only new data created.

**What Go does instead:** Go uses mutable state everywhere. Structs are mutable. Maps are mutable. You protect shared mutable state with mutexes. This works but requires discipline — the language does not help you.

**The Go connection:**
- Before Clojure: "I pass a map to a function and modify it, that's fine." After Clojure: "That function now has a hidden dependency on state. If called concurrently, it's a race. If called twice, second call sees modified state from first. I should think about whether this should be immutable."
- Before Clojure: "A struct is data." After Clojure: "A struct that can be modified from anywhere is not data — it's shared mutable state, which is a different (more dangerous) thing."
- Before Clojure: "I use channels to pass data between goroutines." After Clojure: "Channels are the right instinct — you're sending immutable values between goroutines, not sharing mutable state. This is exactly what Clojure would do."

**In a senior interview:** You can say — "When designing concurrent systems, I prefer to pass immutable data between goroutines rather than sharing mutable state. The Go memory model makes shared mutable state safe only with synchronization — it's simpler to avoid sharing altogether where possible."

**Schedule:**
- Week 1: Watch Rich Hickey's "Simple Made Easy" talk (55 min, free on YouTube). This is the entire philosophy in one talk.
- Week 2: Watch "The Value of Values" talk (30 min) and "Persistent Data Structures" talk (30 min).
- Week 3 (optional): Write one small Clojure program just to feel the immutability. Use clojure.org/guides/getting_started.
- Stop here. You do not need to write Clojure professionally.

**Enough when:** You can explain why immutable data structures simplify concurrent programming, and identify two places in your Go code where you are sharing mutable state that could be replaced with message passing.

**Canonical resource:** Rich Hickey's "Simple Made Easy" talk (2011, Strange Loop conference) — free on YouTube. Watch this first, before anything else. The rest of his talks are bonuses.

---

### Module 3: TypeScript / Node.js — Event Loop (1–2 weeks)

**The core idea:**
Node.js was invented to solve the C10K problem — handling 10,000 concurrent connections. The existing solution was one OS thread per connection, which is expensive. Node's answer: a single thread with an event loop. Instead of blocking on I/O (waiting for a database response, a file read), you register a callback and continue. When the I/O completes, the event loop calls your callback.

The single concept to internalize: **async/await is syntactic sugar for callbacks, and callbacks are how a single thread handles concurrency**. There is no parallelism — only concurrency through non-blocking I/O.

**What Go does instead:** Go uses goroutines — lightweight threads scheduled by the Go runtime. You write blocking code that looks synchronous, and the scheduler switches goroutines when one is waiting on I/O. Same result (concurrency), completely different model.

**The Go connection:**
- Before Node study: "Goroutines are just how you do concurrency." After: "Goroutines and the event loop solve the same problem differently. Node forces you to think about async explicitly (await, Promise). Go hides it — you write blocking code and the scheduler makes it concurrent. Both approaches work. Node struggles with CPU-heavy work (blocks the single thread). Go handles CPU and I/O work equally well."
- Before: "Go's net/http is fast because goroutines." After: "Go's net poller is functionally similar to Node's libuv — both use epoll/kqueue at the OS level. The difference is the programming model exposed to the developer."

**In a senior interview:** You can say — "Node's event loop and Go's goroutines both solve high-concurrency I/O efficiently. The difference is model: Node requires async/await, which is explicit but verbose and fails under CPU load. Go's goroutines let you write synchronous-looking code that runs concurrently — simpler and handles mixed I/O + CPU workloads better."

**Schedule:**
- Day 1-3: Read "What the heck is the event loop anyway?" — Philip Roberts talk (26 min, JSConf EU 2014, free YouTube). This explains it perfectly.
- Day 4-7: Read Node.js event loop documentation (nodejs.org/en/docs/guides/event-loop-timers-and-nexttick). One read-through.
- Day 8-10: Write one tiny Node.js program with async/await and compare it mentally to the equivalent Go goroutine code.
- Stop here.

**Enough when:** You can explain why `await` is needed in Node but not in Go, and what Go's scheduler does that Node's event loop does not.

**Canonical resource:** Philip Roberts — "What the heck is the event loop anyway?" (JSConf EU 2014, YouTube). 26 minutes. This is the single best explanation of the event loop that exists.

---

### Module 4: SQL — Beyond Basics (Integrate into regular study)

**Do not treat this as a separate module with a start/end date.** Integrate it into your Track B3 (PostgreSQL) sessions. Every week, add one advanced SQL concept alongside your regular PostgreSQL study.

**The core idea:**
SQL is a declarative, set-based language. You describe *what* result you want, not *how* to get it. The query planner decides how. Most Go engineers write SQL procedurally — they think in loops ("for each user, get their orders"). The set-based way is: "give me users joined to their orders" — one operation on sets.

**The single concept to internalize:** SQL operates on entire sets of rows at once. Window functions, CTEs, and aggregations are all set operations. A Go loop that processes rows one at a time is often a sign that the SQL could be doing that work.

**Topics to add weekly (in order):**
1. Window functions: `ROW_NUMBER()`, `RANK()`, `LAG()`, `LEAD()` — partition by, order by
2. CTEs (WITH clause) — readable complex queries, recursive CTEs for tree structures
3. EXPLAIN plan reading: Seq Scan vs Index Scan vs Hash Join vs Nested Loop
4. Partial indexes and covering indexes — when and why
5. `SELECT FOR UPDATE` vs optimistic locking — when to use each
6. Keyset pagination vs offset pagination — why offset is slow on large tables

**Canonical resource:** "Use The Index, Luke" (use-the-index-luke.com) — free, specifically about making SQL fast, written for developers not DBAs. Read in order. One chapter per week.

---

### Module 5: A Little C — Reading Only (1 week, whenever)

**Do not write C. Only read it.**

**The core idea:**
C is the language Go's runtime is largely written in (with some assembly). Understanding C means you can read Go's own source code at the level of memory layout, pointer arithmetic, and system calls. Every abstraction in Go — goroutines, channels, the GC — is built on primitives you can read in C.

**The single concept to internalize:** In C, a pointer is just a memory address. An array is just a pointer to the first element. A struct is just a block of memory with named offsets. Every high-level abstraction in every language is built on these.

**The Go connection:**
- After reading C: Open `src/runtime/chan.go` in Go's source. You will be able to read it. The `hchan` struct is just memory layout. The send/receive operations are pointer manipulation with synchronization. It is no longer magic.
- After reading C: "Escape analysis" — when Go moves a variable to the heap — becomes concrete. You understand what the heap and stack physically are.

**Schedule:**
- Day 1-2: Read "Learn C in Y Minutes" (learnxinyminutes.com/docs/c/) — 1 hour total. Just syntax.
- Day 3-5: Read the first 3 chapters of "The C Programming Language" (K&R). Focus on: variables, pointers, arrays. Skip everything else.
- Day 6-7: Open Go's source — `src/runtime/chan.go`. Read `hchan` struct and `chansend()` function. Note how it maps to what you know about channels.
- Stop completely after this week.

**Enough when:** You can read `src/runtime/chan.go` and understand what the `hchan` struct contains and what `lock`, `sendq`, and `recvq` are for.

**Canonical resource:** "The C Programming Language" by Kernighan and Ritchie (K&R) — chapters 1-5 only. This book is 45 years old and still the best introduction to C. Get a PDF or physical copy.

---

### Sequencing Summary

| When | What | Time |
|---|---|---|
| Month 9+ or ahead of schedule | Rust (chapters 1-4, 15-16 of The Rust Book) | 3-4 weeks |
| After Rust | Clojure (Rich Hickey talks only) | 2-3 weeks |
| Any gap week | TypeScript/Node event loop (Philip Roberts talk + docs) | 1-2 weeks |
| Ongoing, weekly | SQL advanced (Use The Index, Luke — 1 chapter/week) | Integrated |
| Whenever curious about Go runtime | C reading (K&R chapters 1-5 + Go source) | 1 week |

**Total time investment:** ~10-12 weeks if done sequentially. Not a priority before month 9. If interviews are going well and you have extra bandwidth, start with the Philip Roberts talk this week — it's 26 minutes and immediately useful.

---

## File Map

```
interview-prep/
├── MASTER_PLAN.md          ← this file — the strategy
├── CLAUDE.md               ← session rules and track definitions
├── SPRINT.md               ← daily checklist per topic (Track A + B)
├── ROADMAP.md              ← full topic map with best sources
├── progress_tracker.html   ← open in browser, visual progress
├── projects.txt            ← your real project details
├── code/ds/                ← function signatures + implementations
├── notes/                  ← all discussion conclusions
├── tasks/                  ← questions and exercises per topic
├── english/vocabulary.md   ← English corrections and phrases
├── real-interview-questions/ ← real questions from companies
├── mock-interviews/
│   ├── quick-reference.md  ← study this before every interview
│   ├── weak-points.md      ← gaps found in mock sessions
│   └── sessions/           ← one file per session
└── russian/
    ├── ROADMAP.md          ← Russian learning phases
    ├── grammar.md          ← cases, prepositions reference
    ├── vocabulary.md       ← corrections and phrases
    └── interview_phrases.md ← ready answers in Russian
```
