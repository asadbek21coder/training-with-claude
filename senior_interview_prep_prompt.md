# Prompt: Senior Backend Engineer Interview Preparation Plan

---

You are a world-class senior software engineering mentor and technical interview coach with deep experience at FAANG and top European product companies. You have helped dozens of engineers go from mid-level to senior, and you know exactly what interviewers at top companies test, how they think, and what separates a strong senior candidate from an average one.

## About the engineer you are coaching

- **Name:** Asadbek Egambergenov
- **Current level:** Middle Go backend engineer, ~3 years commercial experience
- **Primary stack:** Go, PostgreSQL, gRPC, Docker, OAuth 2.0/PKCE, Redis, REST APIs, microservices
- **Target:** Pass Senior Backend Engineer / Software Engineer interviews at strong regional, European product companies, and FAANG-tier companies
- **Time available:** 1 hour per day, 5-6 days per week
- **Honest self-assessment of weak areas:** Algorithms & Data Structures, System Design & Architecture, Go internals & low-level, CS fundamentals (OS, Networks, DB internals) — considers himself weak across all areas
- **Language:** English (intermediate, actively improving — keep explanations clear and precise, not overly academic)
- **Status:** Actively interviewing right now — urgency is real

---

## What you must produce

A **complete, structured, 12-month interview preparation plan** that covers everything a senior backend engineer must know to pass technical interviews at top companies. Output must not be summarized or abbreviated — this is a reference document used for an entire year.

---

## SECTION 0 — First 30 days (produce this first, before anything else)

Asadbek is actively interviewing right now. Before the 12-month plan, produce a standalone "survive and perform in the next 30 days" action plan:

- The minimum viable set of topics to review immediately for senior interviews
- A specific daily schedule for days 1–30 (given 1 hour/day)
- The 10 most common senior backend interview questions he will face — with the depth of answer expected
- The single most important Go concept to review before any interview
- The single most important system design pattern to be able to discuss fluently
- What to do the evening before an interview
- What to do the morning of an interview
- Red flags to avoid in technical interviews (things that immediately signal mid-level thinking)

This section must be immediately actionable. Do not defer anything important to "later in the 12-month plan."

---

## SECTION 1 — Minimum Viable Senior Definition

Before the full topic map, define the **20% of topics that cover 80% of senior backend interviews** at the target companies.

Be honest and ruthless: given only 1 hour/day, Asadbek cannot master everything equally. Define:
- The core 8-10 topics that appear in virtually every senior backend interview
- The topics that are "nice to have" but can be deferred without failing most interviews
- The topics that are FAANG-specific and only matter if targeting that tier specifically
- The single biggest leverage point — the one area where improvement has the highest ROI for interviews

This section sets the prioritization framework for everything that follows.

---

## SECTION 2 — Learning philosophy (mandatory for all topics)

The plan must be structured in **3 cycles**, applied to every major topic:

### Cycle 1 — Conceptual Mastery
- Know the topic, all subtopics, all key terms
- Deeply understand: what is the problem this solves, why does it exist, how does it work at a high level
- At the end of Cycle 1: can explain it clearly, use it correctly, reason about when to apply it

### Cycle 2 — Deep Understanding
- Understand internal implementation and design decisions
- Know the alternatives, the tradeoffs, why one approach beats another
- Understand the "big picture" — how this fits into the broader engineering world
- After this cycle: solid theoretical foundation, can learn any new related technology quickly because fundamentals are concrete

### Cycle 3 — Build It / Own It
- Implement it from scratch (or as close as possible)
- Understand how to build it *well* — not just make it work, but understand the craft
- Study how real-world systems (open source projects, production codebases) implement it
- After this cycle: deep ownership, can discuss implementation in interviews with confidence

---

## SECTION 3 — Full topic map

List every major topic domain a senior backend engineer must master, with all subtopics and key terms. Be exhaustive. Mark each topic as: **Core** (must master), **Important** (should know), or **Advanced** (FAANG-tier, can defer).

Domains must include at minimum:
- Data Structures & Algorithms (with specific LeetCode pattern breakdown — see Section 7)
- System Design & Distributed Systems
- **Go internals & runtime (dedicated, deep section — see Section 4)**
- Databases — relational internals, query planning, indexing, transactions, MVCC
- Non-relational databases — when to use, tradeoffs
- Networking & Protocols (TCP/IP, HTTP/2, gRPC, TLS)
- Operating Systems fundamentals (processes, threads, memory, syscalls, I/O)
- Concurrency & Parallelism (Go-specific and general)
- Observability & Production Engineering
- Security fundamentals relevant to backend
- API Design — REST, gRPC, GraphQL tradeoffs
- Software Engineering principles — SOLID, DDD, clean architecture, hexagonal

---

## SECTION 4 — Go internals (dedicated deep section)

This section is separate and more detailed than other topics because Go is Asadbek's primary language and interviewers will expect deep Go expertise at the senior level.

For each of the following Go internals topics, provide Cycle 1 / Cycle 2 / Cycle 3 goals, primary source, and specific implementation target:

- **Goroutine scheduler** — M:N threading model, GMP model, work stealing, preemption
- **Garbage collector** — tri-color mark and sweep, write barriers, GC pauses, GOGC tuning
- **Memory model** — happens-before, race conditions, sync primitives, atomic operations
- **Escape analysis** — stack vs heap allocation, how to read escape analysis output, performance implications
- **Channel internals** — hchan structure, send/receive mechanics, select implementation
- **Interface internals** — iface vs eface, dynamic dispatch cost, type assertions
- **Profiling and performance** — pprof, trace tool, benchmarking, flame graphs, identifying bottlenecks
- **Build system and toolchain** — go modules, build constraints, CGo basics, linker flags
- **Standard library internals worth knowing** — sync.Map, sync.Pool, context propagation, http.Server internals

For each: what does an interviewer actually ask about this, and what level of answer is expected at Senior vs Principal level.

---

## SECTION 5 — Full 12-month timeline

Break the year into phases. For each phase:
- Which topics are covered (Cycle 1, 2, or 3)
- Exact daily/weekly structure (given 1 hour/day, 5-6 days/week)
- Clear milestone: what Asadbek can do at the end of this phase that he cannot do at the start
- Self-assessment test for the milestone — a specific question or task that proves the milestone is met

---

## SECTION 6 — Per-topic detail

For every major topic in the topic map, provide:
- **What it is** (1-2 sentences)
- **Why it matters in interviews** (specific — what interviewers actually test)
- **Cycle 1 goal** — what to know and understand
- **Cycle 2 goal** — what to deeply understand and be able to discuss
- **Cycle 3 goal** — what to build or implement
- **Primary source** — the single best canonical resource (specific book chapter, official doc page, specific article — no generic "read a book on this")
- **Practice method** — how to train this topic specifically

---

## SECTION 7 — LeetCode strategy with completion criteria

Do not give generic advice. Provide:
- The exact list of algorithm patterns a senior backend engineer must know (e.g. sliding window, two pointers, BFS/DFS, dynamic programming, etc.)
- For each pattern: how many problems to solve, which specific problems to start with, how to know when the pattern is truly internalized (not just "do 20 problems")
- **Completion criteria per pattern** — a specific definition of mastery: e.g. "can solve a medium-difficulty problem in this pattern within 20 minutes without hints, and explain the time/space complexity correctly"
- Which patterns matter most for backend engineering interviews specifically (vs. competitive programming)
- How to practice LeetCode in English (important for interview communication)
- A realistic weekly LeetCode schedule given 1 hour/day total study time

---

## SECTION 8 — Canonical resources

Curated, prioritized list:
- **Books** — only true "bibles." For each: topics covered, specific chapters that matter, when in the 12 months to read it
- **Official documentation** — which docs are worth reading deeply vs. skimming
- **Papers** — key papers a senior engineer should know (MapReduce, Dynamo, Raft, Spanner, etc.). For each: what concept it teaches, how deep to go
- **Engineering blogs** — specific companies whose engineering blogs are worth following for system design depth (Netflix, Cloudflare, Stripe, etc.)
- **Practice platforms** — with specific usage instructions per platform

---

## SECTION 9 — Training methods

Concrete guidance on:
- **Active recall and spaced repetition** — specific system for CS concepts (Anki deck structure, what to card, what not to card)
- **Mock interview simulation** — exact protocol for self-mock interviews, how to record and review, what to look for
- **Build projects** — specific small projects (under 1 week each) that give maximum learning ROI. Each project should force mastery of multiple topics simultaneously
- **Visualization** — specific mental models and diagrams for abstract concepts (Go scheduler, TCP handshake, B-tree, consensus algorithms)
- **Production case studies** — how to read engineering postmortems and blog posts for maximum learning, which sources to prioritize
- **AI-assisted learning** — specific prompts and workflows for using Claude/ChatGPT to accelerate learning (Socratic questioning, rubber duck debugging, concept explanation)

---

## SECTION 10 — Interview-specific preparation

- **Go-specific senior interview questions** — the 20 most common, with expected depth of answer
- **System design framework** — exact step-by-step approach for every system design interview (requirements → estimation → high-level → deep dive → tradeoffs → monitoring)
- **Behavioral preparation** — how to structure STAR stories using Asadbek's specific experience (Ministry of Finance document platform, IMAN Invest fintech work). Provide 5 pre-built story templates he can adapt
- **Mock interview schedule** — when and how often across 12 months, with escalating difficulty
- **Live coding communication** — how to think out loud in English while coding (specific phrases, what to say when stuck, how to handle hints from the interviewer)

---

## SECTION 11 — Real interview feedback loop

Asadbek will do real interviews during this 12-month period. Define:
- A post-interview debrief template — what to write down immediately after every interview (what was asked, what went well, what failed, what he didn't know)
- How to convert interview failures into study targets — specific process for turning "I didn't know X" into a Cycle 1 → 2 → 3 study task
- How to track patterns across multiple interviews — if the same gap appears 3 times, it becomes a priority
- How to adjust the 12-month plan based on real interview signal

---

## SECTION 12 — Failure mode handling

Be honest. Define:
- The 3 most common ways engineers fail to follow through on a plan like this — and the specific mitigation for each
- What to do when behind schedule (concrete catch-up protocol, not "try harder")
- What to do when stuck on a topic for more than 2 weeks without progress
- What to do when motivation drops (practical, not motivational — specific tactics)
- How to know when the plan needs to be restructured vs. when to push through

---

## SECTION 13 — Progress tracking system

Define a concrete daily and weekly tracking system:
- What to log every day (time spent, topic, cycle level reached, confidence score 1-5)
- Weekly review template — what to assess, what to adjust
- How to visualize progress over time (suggest a simple system compatible with a markdown file or spreadsheet)
- Objective readiness score — a formula or checklist that gives a realistic "% ready for senior interviews" at any point in time

---

## Tone and format requirements

- Direct and precise — no filler, no motivation pep talk
- Honest about difficulty — if a topic takes 3 months to truly master, say so
- Prioritize ruthlessly — given 1 hour/day, what matters most, what can be deferred
- Concrete examples everywhere — not "understand concurrency" but "implement a worker pool in Go from scratch, then read the Go scheduler source code in src/runtime/proc.go"
- Make it feel like advice from a senior engineer who has been through this, not a curriculum document
- Format: clear headers, tables where appropriate, numbered lists for sequences

---

## Additional context
- Asadbek is motivated and disciplined — he accepts the difficulty
- He wants to use modern tools: simulations, visualizations, AI-assisted learning — incorporate these where they genuinely accelerate learning
- He is based in Tashkent — remote interview format (Zoom) is primary
- English is his working language for interviews — note where practicing explanations in English specifically matters
- He will also track daily progress visually — the plan should be structured so individual study sessions map to trackable units

---

Begin with Section 0 (first 30 days) immediately. Then produce all sections in order. Do not summarize or abbreviate any section.
