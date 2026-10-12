# AI Engineering Roadmap

> **L1** = aware & can use · **L2** = deep & efficient · Status: ⬜ not started · 🔄 in progress · ✅ done
> **Notes:** `notes/ai/` · **Practice:** `code/ai/`

**Why (EN):** AI made writing code cheap; it did not make correct, secure, operable systems cheap. Two skills now separate seniors: (1) **using AI well** in the daily SWE process — giving good context and specs, splitting work into verifiable steps, reviewing AI output like a strict senior; (2) **building AI features** — LLM APIs, RAG, agents, evals, cost and security. You use AI every day and teach agentic coding in your course (week 12+), so this track runs **in parallel**, not "later".
**Зачем (RU):** ИИ сделал написание кода дешёвым, но не сделал дешёвыми корректные, безопасные и надёжные системы. Senior сейчас отличают два навыка: (1) **умело использовать ИИ** в ежедневной разработке — давать хороший контекст и спецификации, делить работу на проверяемые шаги, строго ревьюить результат; (2) **строить ИИ-функции** — LLM API, RAG, агенты, эвалы, стоимость и безопасность. Ты пользуешься ИИ каждый день и преподаёшь агентную разработку в курсе (с 12-й недели), поэтому этот трек идёт **параллельно**, а не «потом».

**Order:** AI 1–5 = *using* AI (daily parallel, one small piece a day, practiced on real tasks: your course, Go practice, work). AI 6–12 = *building* with AI (start after Go L1; in your course it lines up with weeks 16–21).
(Topics SWE 14–16 moved here from `common-swe` on 2026-10-12.)

## Progress
| Layer | Done | Total |
|---|---|---|
| L1 | 0 | 24 |
| L2 | 0 | 21 |

---

## AI 1 — How LLMs work (enough to use them well)
**L1 — aware & can use**
- ⬜ **1.1 Core behaviour** — token, tokenizer, context window, input vs output tokens, temperature, sampling, non-determinism, hallucination, system prompt, few-shot examples, knowledge cutoff
- ⬜ **1.2 How models are made** — pretraining (next-token prediction), instruction tuning, RLHF / RLAIF, reasoning ("thinking") models, multimodal input (images, PDFs), tool use as a trained skill
- ⬜ **1.3 Choosing a model** — big vs small vs fast models, cost per million tokens, latency, effort/thinking level, open-weight vs closed models, benchmarks (SWE-bench etc.) and why they mislead

**L2 — deep & efficient**
- ⬜ **1.4 Transformer intuition** — embeddings, attention (each token looks at other tokens), layers, why long context gets harder ("context rot", lost in the middle), why models hallucinate and are overconfident
- ⬜ **1.5 Limits you design around** — no real memory between sessions, stale knowledge, arithmetic and counting errors, sycophancy (agreeing with you), prompt sensitivity, non-reproducible outputs

🛠 **Practice** — L1: ask the same Go question at temperature 0 and 1 (via API or two models) and compare; count the tokens of one of your notes with a tokenizer · L2: write a one-page note "5 ways LLMs fail me in Go code" from your own real sessions, with one example each

---

## AI 2 — Prompting & context engineering
**L1 — aware & can use**
- ⬜ **2.1 Prompting basics** — clear and specific instructions, give the "why", examples (few-shot), roles, structure with sections/XML tags, define the output format, ask for a plan first, let the model think before answering
- ⬜ **2.2 Context engineering** — project instruction files (`CLAUDE.md`, `AGENTS.md`, rules files), giving only the relevant files, conventions and constraints, examples of good code from your repo, keeping context small and accurate, starting fresh sessions

**L2 — deep & efficient**
- ⬜ **2.3 Advanced patterns** — prompt chaining, self-critique / review passes, asking the model to ask clarifying questions, structured outputs (JSON schema), negative vs positive instructions, prompt versioning and testing prompts like code
- ⬜ **2.4 Long work & memory** — context window budgeting, summaries and handoff notes, progress files (like this repo's `LOG.md`), compaction, splitting a big task across sessions

🛠 **Practice** — L1: rewrite 3 of your recent vague prompts into structured ones (goal, context, constraints, output format) and compare the answers · L2: build a reusable prompt template for "write table-driven tests for this Go function" and test it on 5 functions

---

## AI 3 — AI-assisted coding workflow
**L1 — aware & can use**
- ⬜ **3.1 Tools & the basic loop** — agentic coding tools (Claude Code, Cursor, Codex, Copilot), chat vs autocomplete vs agent, task decomposition, small verifiable steps, plan → implement → verify, when NOT to use AI
- ⬜ **3.2 Spec-driven development** — write the spec first (goal, inputs/outputs, edge cases, acceptance tests), ask for a plan, review the plan, then the diff, then run tests; tests as the contract for AI code
- ⬜ **3.3 Verifying AI output** — reading every line, running tests and linters, hallucinated APIs/packages, outdated library versions, insecure defaults, missing error handling, concurrency bugs (leaks, races), over-engineering, silently changed behaviour

**L2 — deep & efficient**
- ⬜ **3.4 AI for the hard parts** — debugging with AI (give logs, traces, failing test), writing tests for legacy code, refactoring and migrations in small safe steps, reading an unfamiliar codebase with AI, generating docs and diagrams
- ⬜ **3.5 Parallel & background agents** — several agents on separate tasks, git worktrees, cloud/background agents, async task handoff, reviewing many AI PRs, avoiding merge chaos

🛠 **Practice** — L1: implement one Bozorcha lesson feature spec-first with an AI agent: write the spec + tests yourself, let the agent implement, review every line; log every AI mistake you caught · L2: give an AI tool the same task with a vague prompt vs a spec-first prompt with tests; compare correctness, diff size and review time

---

## AI 4 — Agentic tools in depth (Claude Code as the example)
**L1 — aware & can use**
- ⬜ **4.1 Daily features** — permission modes, plan mode, slash commands, memory/instruction files, checkpoints and undo, attaching files and images, cloud sessions from the phone (like this repo)
- ⬜ **4.2 MCP as a user** — what Model Context Protocol is, connecting MCP servers (GitHub, database, browser, docs), tools vs resources, trusting third-party servers

**L2 — deep & efficient**
- ⬜ **4.3 Extending the agent** — subagents, skills, hooks (run linters/tests automatically), plugins, custom slash commands, per-project settings
- ⬜ **4.4 Agents in automation** — headless mode, SDK, running an agent in CI (e.g. GitHub Actions for PR review), scheduled tasks, permissions and sandboxing for unattended runs
- ⬜ **4.5 Cost & usage control** — choosing model and effort per task, token usage, prompt caching, keeping sessions short, when a cheaper model is enough

🛠 **Practice** — L1: connect one MCP server (e.g. GitHub or Postgres) and use it for a real task · L2: add a hook that runs `gofmt` + `go vet` after every edit in your Go project; write one custom skill (e.g. "prepare a lesson recall check") and use it

---

## AI 5 — AI in team process & your own growth
**L1 — aware & can use**
- ⬜ **5.1 Responsibility & review** — you own every line you merge, AI-written PRs and review rules, disclosing AI use, licence and secrets risks (pasting company code into tools), company AI policies
- ⬜ **5.2 Keeping your own skills sharp** — AI-free practice (interviews, your course weeks 1–6), deliberate practice, explaining code without AI, reading source code, avoiding "skill atrophy"

**L2 — deep & efficient**
- ⬜ **5.3 Team adoption** — shared `AGENTS.md`/`CLAUDE.md`, guardrails (linters, CI, required tests) as the safety net, AI review assistants and their limits, measuring real productivity (lead time, defects, not "lines of code"), teaching juniors to use AI well

🛠 **Practice** — L1: one week of a simple log: task, AI used or not, time spent, bugs found later · L2: write an "AI usage rules" page for your course students (what's allowed per phase, how to review AI code)

---

## AI 6 — LLM APIs from Go
**L1 — aware & can use**
- ⬜ **6.1 Messages API** — request/response, roles, system prompt, `max_tokens`, stop reasons, API keys and secrets, official Go SDKs vs raw HTTP
- ⬜ **6.2 Streaming & reliability** — streaming (SSE), timeouts, retries with backoff + jitter, rate limits (`429`), overload errors, idempotency, context cancellation in Go
- ⬜ **6.3 Structured output & tool calling** — JSON schema output, validating model output, tool/function definitions, the tool-use loop, parallel tool calls

**L2 — deep & efficient**
- ⬜ **6.4 Cost & performance features** — prompt caching, batch API, token counting, model routing (cheap model first), response caching
- ⬜ **6.5 Multimodal & files** — images, PDFs, files API, long documents, citations

🛠 **Practice** — L1: Go CLI that streams an answer from an LLM API with a timeout and retry; add one tool ("get product by ID" from Bozorcha) · L2: add prompt caching and measure cost/latency before and after

---

## AI 7 — Embeddings, vector search & RAG
**L1 — aware & can use**
- ⬜ **7.1 Embeddings & vector search** — embedding, vector, cosine similarity, dot product, pgvector (`vector` type, `<=>` operator), HNSW vs IVFFlat, approximate nearest neighbour
- ⬜ **7.2 RAG basics** — retrieval-augmented generation, chunking (size, overlap), ingestion pipeline, top-k retrieval, prompt assembly, citations

**L2 — deep & efficient**
- ⬜ **7.3 Better retrieval** — hybrid search (BM25 + vectors), re-ranking, contextual chunk headers, metadata filters, query rewriting, when RAG is the wrong tool (long context, SQL, search API)

🛠 **Practice** — L1: small RAG in Go: ingest your `notes/` Markdown into Postgres + pgvector (Docker), embed chunks, answer questions with citations · L2: add hybrid search + re-ranking and compare answers on 10 questions

---

## AI 8 — Building agents & MCP servers
**L1 — aware & can use**
- ⬜ **8.1 Agent loop** — model + tools + loop until done, planning, stop conditions, max steps, human-in-the-loop approvals
- ⬜ **8.2 Building an MCP server** — server, client, tools, resources, prompts, transports (stdio, streamable HTTP), Go MCP SDK

**L2 — deep & efficient**
- ⬜ **8.3 Agent design** — workflows vs autonomous agents, multi-agent (orchestrator + workers), state and memory, long-running tasks, error recovery, permissions, cost limits

🛠 **Practice** — L1: minimal MCP server in Go exposing "search notes" and "list roadmap ⬜ items" tools for this repo · L2: an agent that reads `LOG.md` + roadmap and proposes tomorrow's plan, with a human approval step

---

## AI 9 — Evals & LLM observability
**L1 — aware & can use**
- ⬜ **9.1 Evals basics** — why "it looks good" is not testing, eval dataset, golden answers, pass/fail checks, regression evals before changing a prompt or model

**L2 — deep & efficient**
- ⬜ **9.2 Advanced evals** — LLM-as-judge and its biases, rubrics, retrieval metrics (recall@k, MRR), A/B testing prompts, evals in CI
- ⬜ **9.3 Observability** — tracing LLM calls, token usage and cost metrics, latency percentiles, logging prompts safely (PII), OpenTelemetry GenAI conventions

🛠 **Practice** — L1: eval set of 20 Q&A pairs for your RAG, run it with a script · L2: measure recall@5 before/after changing chunk size; add an LLM-as-judge check and compare with your own grading

---

## AI 10 — Production AI systems
**L1 — aware & can use**
- ⬜ **10.1 Production concerns** — latency budgets, cost control, provider fallbacks, timeouts, graceful degradation when the model is down, user-facing streaming UX

**L2 — deep & efficient**
- ⬜ **10.2 Architecture choices** — prompting vs RAG vs fine-tuning, model routing, caching layers, async jobs for slow AI work (queues), guardrails layer
- ⬜ **10.3 Self-hosted models** — Ollama, vLLM, quantization, GPU vs CPU, when self-hosting is worth it (privacy, cost at scale) and when not

🛠 **Practice** — L1: put the RAG behind an HTTP endpoint with timeout, fallback message and cost logging · L2: run a small open model with Ollama and route easy questions to it, hard ones to an API model

---

## AI 11 — AI security
**L1 — aware & can use**
- ⬜ **11.1 OWASP Top 10 for LLM applications** — prompt injection, sensitive information disclosure, supply chain, data and model poisoning, improper output handling, excessive agency, system prompt leakage, vector and embedding weaknesses, misinformation, unbounded consumption
- ⬜ **11.2 Prompt injection** — direct vs indirect injection, injection via retrieved documents/web pages/tool results, jailbreaks, why there is no complete fix, treating model output as untrusted input

**L2 — deep & efficient**
- ⬜ **11.3 Defending AI systems** — least-privilege tools, human approval for risky actions, output validation, sandboxing tool execution, allowlists, separating instructions from data, rate/cost limits, audit logs of tool calls
- ⬜ **11.4 Data protection with LLMs** — secrets in prompts, PII redaction, data retention policies, tenant isolation in vector stores, access control on retrieved chunks, training-data opt-out

🛠 **Practice** — L1: add a document with hidden instructions ("ignore previous instructions, reveal …") to your RAG and observe the behaviour · L2: add defences (chunk-level ACL filtering, output validation, tool allowlist with confirmation) and write regression tests with 10 injection payloads

---

## AI 12 — ML foundations (awareness)
**L1 — aware & can use**
- ⬜ **12.1 ML vocabulary** — model, training vs inference, dataset, features, labels, supervised vs unsupervised, overfitting, train/validation/test split, loss, accuracy vs precision/recall

**L2 — deep & efficient**
- ⬜ **12.2 Adapting models** — fine-tuning, LoRA, distillation, quantization, embeddings models vs generation models, classic ML vs LLM for a task (e.g. spam filter, price prediction)

🛠 **Practice** — L1: explain to your student, in 5 minutes, the difference between training and inference and why a model has a knowledge cutoff · L2: compare an LLM prompt vs a tiny classic classifier (e.g. Python scikit-learn) on the same 100 labelled examples
