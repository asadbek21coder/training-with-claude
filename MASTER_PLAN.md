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
