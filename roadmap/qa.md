# Testing & Quality Engineering Roadmap

> **L1** = aware & can use · **L2** = deep & efficient · Status: ⬜ not started · 🔄 in progress · ✅ done
> **Notes:** `notes/qa/` · **Practice:** `code/qa/`

*(Go-specific testing syntax — `testing` package, table-driven tests, `t.Run`, `httptest`, benchmarks, `go test` flags — lives in the `golang` track. This track is test strategy, techniques and the QA discipline; examples use Go, but the ideas are language-agnostic.)*

**Why (EN):** A senior engineer owns quality, not just code: they decide what to test at which level, design test cases that actually find bugs, and build fast, trustworthy suites and quality gates for the team. In the AI era, code is cheap and plausible-but-wrong code is everywhere, so tests become the main specification and the main defense — the value moves from typing test code (AI can do that) to test design, choosing the right oracle, spotting missing cases, and judging whether a suite really catches bugs.
**Зачем (RU):** Senior-инженер отвечает за качество, а не только за код: решает, что и на каком уровне тестировать, проектирует тест-кейсы, которые действительно находят баги, и строит быстрые, надёжные наборы тестов и quality gates для команды. В эпоху ИИ код дешёвый, а правдоподобный, но неверный код повсюду, поэтому тесты становятся главной спецификацией и главной защитой — ценность смещается от написания тестового кода (это сделает ИИ) к проектированию тестов, выбору правильного оракула, поиску пропущенных случаев и оценке того, ловит ли набор тестов баги на самом деле.

## Progress
| Layer | Done | Total |
|---|---|---|
| L1 | 0 | 44 |
| L2 | 0 | 27 |

---

## QA 1 — Testing fundamentals
**L1 — aware & can use**
- ⬜ **1.1 Why we test** — verification vs validation, error → defect (bug) → failure, cost of a bug by phase, shift-left, testing shows presence not absence of bugs, quality vs testing
- ⬜ **1.2 Test levels & shapes** — unit, integration, system, acceptance, E2E, test pyramid, testing trophy, testing honeycomb (microservices), ice-cream cone anti-pattern, cost vs confidence vs speed
- ⬜ **1.3 Seven testing principles (ISTQB)** — exhaustive testing is impossible, early testing, defect clustering, pesticide paradox (tests wear out), testing is context dependent, absence-of-errors fallacy

**L2 — deep & efficient**
- ⬜ **1.4 Test strategy thinking** — test what can break, risk × likelihood, testing behavior not implementation, test oracles (specified, derived, implicit, heuristic), oracle problem, sociable vs solitary unit tests, what "unit" means

🛠 **Practice** — L1: for a small Go service (`orders`: create/list/cancel with Postgres), list 15 things that can break and assign each to a test level · L2: write a 1-page test strategy for that service: which levels, what each level owns, what you deliberately do not test, and why

---

## QA 2 — Test doubles & isolation
**L1 — aware & can use**
- ⬜ **2.1 Test double types** — dummy, stub, fake, spy, mock (Meszaros vocabulary), when to use each, in-memory fake repository
- ⬜ **2.2 Seams & dependency injection** — seam, interface at the consumer, constructor injection, injecting clock / random / UUID generator, testable design, hidden dependencies (globals, `time.Now()`)

**L2 — deep & efficient**
- ⬜ **2.3 Schools & trade-offs** — London (mockist) vs Classical (Detroit) school, over-mocking, mocking what you don't own, verifying interactions vs state, brittle tests, mock generators (`gomock`/`mockgen`, `mockery`) vs hand-written fakes
- ⬜ **2.4 Doubles for infrastructure** — HTTP fakes (`httptest.Server`), WireMock / mock servers, record & replay (VCR), fake clocks, fake message brokers, contract drift between fake and real

🛠 **Practice** — L1: test the `orders` service's business logic with a hand-written in-memory fake repository and an injected fake clock · L2: write the same tests with `mockgen` mocks, then refactor the implementation (no behavior change) and count which suite breaks; write down the lesson

---

## QA 3 — Writing good tests
**L1 — aware & can use**
- ⬜ **3.1 Test structure** — Arrange-Act-Assert, Given-When-Then, one behavior per test, descriptive test names, table-driven / parameterized tests, test helpers, fixtures, setup/teardown
- ⬜ **3.2 Properties of good tests** — FIRST (fast, isolated, repeatable, self-validating, timely), deterministic, readable failure messages, testing public API, no logic in tests
- ⬜ **3.3 Coverage basics** — statement / line coverage, branch coverage, coverage reports, coverage as a smell detector not a target, Goodhart's law

**L2 — deep & efficient**
- ⬜ **3.4 Flaky tests** — causes (time, concurrency, order dependency, shared state, network, randomness, resource leaks), detection (rerun, `-count`, shuffle), quarantine, fixing vs deleting, cost of flakiness on trust
- ⬜ **3.5 Test smells & maintainability** — fragile test, mystery guest, eager test, assertion roulette, excessive setup, test duplication, test data builders, object mother, golden files / snapshot testing, approval testing
- ⬜ **3.6 Coverage in depth** — MC/DC (awareness), path coverage, diff coverage on PRs, integration test coverage (`go build -cover`), why 100% coverage still misses bugs

🛠 **Practice** — L1: write table-driven unit tests for `orders` validation and pricing rules with clear names and AAA; generate an HTML coverage report · L2: introduce a flaky test on purpose (time + goroutine race), detect it with `go test -count=50 -shuffle=on -race`, fix the root cause, and document the 3 smells you found in your own tests

---

## QA 4 — Integration testing
**L1 — aware & can use**
- ⬜ **4.1 What integration tests cover** — real database, real broker, real HTTP, SQL correctness, migrations, transactions, serialization, configuration, narrow vs broad integration tests
- ⬜ **4.2 Testcontainers** — Testcontainers for Go, container lifecycle per suite vs per test, wait strategies, modules (Postgres, Redis, Kafka), Docker Compose alternative, CI with Docker

**L2 — deep & efficient**
- ⬜ **4.3 Test data management** — isolation strategies (transaction rollback, truncate, schema/database per test, template databases), seeding, factories, synthetic vs anonymized production data, GDPR concerns, parallel test safety
- ⬜ **4.4 Speed & reliability** — container reuse, `TestMain`, build tags / `-short` to split suites, parallel packages, tmpfs for Postgres data, startup cost, fail-fast ordering in CI

🛠 **Practice** — L1: write integration tests for the `orders` Postgres repository with Testcontainers (create, list with pagination, cancel twice) · L2: make the suite run in parallel safely using a template database per test, and measure total time before/after

---

## QA 5 — API & contract testing
**L1 — aware & can use**
- ⬜ **5.1 API testing basics** — status codes, headers, body schema, negative tests, auth tests, idempotency, pagination, error format, `curl`, `httptest`
- ⬜ **5.2 API testing tools** — Postman (collections, environments, pre-request scripts, tests), Newman (CLI runner), Bruno (git-friendly collections), Hurl (plain-text HTTP tests), Insomnia, REST Client files (`.http`)
- ⬜ **5.3 Schema-based testing** — OpenAPI spec, JSON Schema validation, request/response validation middleware, Schemathesis (property-based API testing from OpenAPI), protobuf for gRPC

**L2 — deep & efficient**
- ⬜ **5.4 Contract testing** — consumer-driven contracts, Pact (consumer test, pact file, provider verification, Pact Broker, `can-i-deploy`), provider-driven contracts, contract vs schema vs E2E testing
- ⬜ **5.5 Breaking-change detection** — backward/forward compatibility, `buf breaking` for protobuf, `oasdiff` for OpenAPI, versioning, tolerant reader, additive changes, deprecation policy, consumer usage tracking

🛠 **Practice** — L1: write a Hurl (or Bruno) test file for the `orders` API covering happy path + 5 negative cases, and run it in CI · L2: write a Pact consumer test for a Go client of `orders` and verify it against the provider; then make a breaking change and show both Pact and `oasdiff` catch it

---

## QA 6 — End-to-end & system testing
**L1 — aware & can use**
- ⬜ **6.1 E2E scope** — user journeys, critical paths, full stack, staging / ephemeral environments, smoke E2E after deploy, keeping E2E suites small
- ⬜ **6.2 E2E tools (awareness)** — Playwright, Cypress, Selenium, API-level E2E vs UI-level E2E, page object model, test IDs

**L2 — deep & efficient**
- ⬜ **6.3 Making E2E sustainable** — ephemeral preview environments, test data isolation per run, idempotent setup via API, retries vs real fixes, async waits (polling with timeout, no `sleep`), testing eventual consistency, ownership of E2E failures

🛠 **Practice** — L1: write 3 API-level E2E tests that run against the full `docker compose` stack (app + Postgres) from a separate Go test package · L2: make the E2E suite create its own data per run and pass reliably 20 times in a row, including an async flow tested with polling

---

## QA 7 — Test design techniques
**L1 — aware & can use**
- ⬜ **7.1 Black-box techniques** — equivalence partitioning, boundary value analysis (2-value, 3-value), decision table testing, state transition testing, use case testing
- ⬜ **7.2 Experience-based techniques** — error guessing, checklist-based testing, exploratory testing, common bug lists (off-by-one, null/empty, unicode, time zones, DST, leap years, large numbers, concurrency)

**L2 — deep & efficient**
- ⬜ **7.3 Combinatorial & white-box** — pairwise / all-pairs testing, orthogonal arrays, classification tree, statement / branch / condition coverage as design techniques, data flow testing, cause-effect graphing (awareness)
- ⬜ **7.4 Domain-specific test design** — money & rounding, idempotency & retries, concurrency & race conditions, time & scheduling, pagination & ordering, permissions matrices, i18n, input fuzz lists (Big List of Naughty Strings)

🛠 **Practice** — L1: for the `orders` "apply discount" rule, derive test cases with equivalence partitions, boundaries and a decision table, then implement them as a table-driven test · L2: draw a state transition diagram for order status (created → paid → shipped → cancelled/refunded), derive all valid and invalid transitions (0-switch and 1-switch coverage), and test them

---

## QA 8 — Test types (functional & non-functional)
**L1 — aware & can use**
- ⬜ **8.1 Functional test types** — smoke, sanity, regression, retesting (confirmation), acceptance (UAT), alpha / beta, black-box vs white-box vs grey-box
- ⬜ **8.2 Non-functional test types** — performance, load, stress, security, usability, accessibility, compatibility, reliability, recoverability, installability / upgrade testing
- ⬜ **8.3 Static testing** — reviews, walkthroughs, inspections, static analysis, linters, requirements review, testing the spec before the code

**L2 — deep & efficient**
- ⬜ **8.4 Regression strategy** — regression suite selection, test impact analysis, risk-based prioritization, change-based testing, test suite minimization, nightly vs per-PR suites, keeping regression suites healthy
- ⬜ **8.5 Quality characteristics** — ISO/IEC 25010 model (functional suitability, performance efficiency, compatibility, usability, reliability, security, maintainability, portability), non-functional requirements as testable criteria

🛠 **Practice** — L1: classify 20 example tests (given or from your own projects) by type and level in a table · L2: turn 5 vague NFRs ("fast", "secure", "reliable") for `orders` into measurable acceptance criteria and map each to a test type and tool

---

## QA 9 — Test artifacts & defect management
**L1 — aware & can use**
- ⬜ **9.1 Test documentation** — test policy, test strategy, test plan (scope, approach, resources, schedule, entry/exit criteria, risks), test case, test suite, test charter, traceability matrix
- ⬜ **9.2 Bug reports** — title, environment, steps to reproduce, expected vs actual result, evidence (logs, screenshots, request IDs), minimal reproduction, one bug per report
- ⬜ **9.3 Severity, priority & lifecycle** — severity vs priority (and examples of high/low combos), defect lifecycle (new, assigned, open, fixed, retest, verified, closed, reopened, deferred, rejected/duplicate), triage

**L2 — deep & efficient**
- ⬜ **9.4 Lightweight documentation in agile teams** — test plans as checklists in PRs, acceptance criteria as tests, living documentation, root cause analysis of escaped bugs, defect taxonomy, bug bash, "zero bug" policy

🛠 **Practice** — L1: write a test plan (1–2 pages) for a new "refund order" feature, and 3 bug reports from bugs you find in your own `orders` service · L2: take one escaped bug, do a root cause analysis (why was it introduced, why was it not caught), and propose the cheapest test that would have caught it

---

## QA 10 — TDD, BDD & acceptance testing
**L1 — aware & can use**
- ⬜ **10.1 TDD** — red-green-refactor, test list, baby steps, triangulation, fake it till you make it, transformation priority premise (awareness), tests as design feedback
- ⬜ **10.2 BDD** — Given/When/Then, Gherkin, scenarios, Example Mapping, three amigos, ubiquitous language, Cucumber, godog (Go)
- ⬜ **10.3 Acceptance criteria & DoD** — user story, acceptance criteria, Definition of Ready, Definition of Done, ATDD (acceptance test-driven development), specification by example

**L2 — deep & efficient**
- ⬜ **10.4 TDD in practice** — outside-in (double loop) vs inside-out TDD, TDD with legacy code (characterization tests, seams), when TDD hurts (spikes, unclear UI), TDD with AI assistants (write the test first, let AI implement, review)

🛠 **Practice** — L1: implement a small rule engine (e.g. order shipping cost) strictly with TDD, committing at each red/green/refactor step · L2: write 4 Gherkin scenarios for "cancel order" with Example Mapping, automate them with godog against the HTTP API; then write characterization tests for a piece of untested legacy code before refactoring it

---

## QA 11 — Exploratory & risk-based testing
**L1 — aware & can use**
- ⬜ **11.1 Exploratory testing** — simultaneous learning, test design and execution, test charter, session-based test management (SBTM), time-box, debrief, notes, heuristics (SFDIPOT, CRUD, "follow the data")
- ⬜ **11.2 Risk-based testing** — product risk vs project risk, likelihood × impact, risk matrix, prioritizing tests by risk, risk register

**L2 — deep & efficient**
- ⬜ **11.3 Testing heuristics & oracles** — HICCUPPS consistency oracles, test tours, "Explore It!" style charters, Rapid Software Testing ideas (testing vs checking), tester's mindset for developers

🛠 **Practice** — L1: run two 45-minute exploratory sessions on the `orders` API (with Bruno/curl) using written charters; log notes, bugs, questions · L2: build a risk matrix for the service, then compare it with where your automated tests actually are and fix the biggest gap

---

## QA 12 — Property-based & fuzz testing
**L1 — aware & can use**
- ⬜ **12.1 Property-based testing** — example-based vs property-based, properties, generators, shrinking, invariants, round-trip (encode/decode), oracle comparison, idempotence, commutativity, `rapid` / `gopter` for Go
- ⬜ **12.2 Fuzzing basics** — coverage-guided fuzzing, seed corpus, fuzz target, crashers, Go native fuzzing (`func FuzzX(f *testing.F)`, `go test -fuzz`), `testdata/fuzz`, OSS-Fuzz (awareness)

**L2 — deep & efficient**
- ⬜ **12.3 Advanced generative testing** — stateful / model-based property testing (state machines), metamorphic testing, differential testing, fuzzing parsers & protocol handlers, structure-aware fuzzing, running fuzzing in CI (time-boxed), libFuzzer / AFL++ (awareness)

🛠 **Practice** — L1: write a round-trip property test for a JSON/CSV encoder in the service with `rapid`, and a Go native fuzz test for a query-string filter parser · L2: write a stateful `rapid` test that compares the Postgres repository against an in-memory model through random sequences of create/cancel/list operations

---

## QA 13 — Test suite quality: mutation testing & test effectiveness
**L1 — aware & can use**
- ⬜ **13.1 Mutation testing** — mutant, mutation operators, killed vs survived mutant, equivalent mutant, mutation score, "who tests the tests?", Go tools (`gremlins`, `go-mutesting`), Stryker / PIT (other languages, awareness)

**L2 — deep & efficient**
- ⬜ **13.2 Measuring test effectiveness** — mutation score vs coverage, incremental mutation testing on diffs, test suite runtime budget, test-to-code ratio (and why it misleads), escaped defect analysis, reviewing AI-generated tests (asserts nothing, mirrors implementation, tautologies)

🛠 **Practice** — L1: run a mutation testing tool on the `orders` domain package and kill 5 surviving mutants with new tests · L2: have an AI assistant generate tests for one package, run mutation testing on them, and write down which weaknesses the surviving mutants reveal

---

## QA 14 — Performance & load testing
**L1 — aware & can use**
- ⬜ **14.1 Performance test types** — smoke, load (average), stress, spike, soak (endurance), breakpoint / capacity test, scalability test, benchmark vs load test
- ⬜ **14.2 Key metrics** — throughput (RPS), latency percentiles (p50, p95, p99, p99.9), error rate, concurrency, virtual users, saturation, Apdex, why averages lie
- ⬜ **14.3 k6 basics** — script, VUs, iterations, `stages`, scenarios, executors (constant-vus, ramping-vus, constant-arrival-rate), checks, thresholds, `k6 run`, output summary

**L2 — deep & efficient**
- ⬜ **14.4 Load test design** — open vs closed workload model, arrival-rate executors, coordinated omission, think time, realistic data & traffic mix, warm-up, test environment parity, load generator limits, distributed load
- ⬜ **14.5 Analysis & tooling** — SLO-based thresholds, correlating with server metrics (Prometheus/Grafana), finding the bottleneck (CPU, DB, locks, pool), USL / Little's Law, vegeta, ghz (gRPC), wrk / hey, performance regression testing in CI, `pprof` under load

🛠 **Practice** — L1: write a k6 script for the `orders` API with a ramping load stage, checks and thresholds (`p(95)<300`, `http_req_failed<1%`) · L2: run a constant-arrival-rate test until the service breaks, find the bottleneck with Prometheus + `pprof`, fix it (e.g. DB pool size or missing index), and report before/after numbers

---

## QA 15 — Chaos & resilience testing
**L1 — aware & can use**
- ⬜ **15.1 Resilience testing basics** — fault injection, failure modes (latency, errors, crashes, network partition, disk full), testing timeouts / retries / circuit breakers, graceful degradation, recovery testing
- ⬜ **15.2 Tools** — Toxiproxy (latency, timeout, bandwidth, reset_peer toxics), `tc netem`, Pumba, Chaos Mesh, LitmusChaos, AWS FIS (awareness)

**L2 — deep & efficient**
- ⬜ **15.3 Chaos engineering discipline** — steady-state hypothesis, blast radius, minimize impact, run in production (maturity), automated chaos in CI, game days, chaos engineering principles (principlesofchaos.org)

🛠 **Practice** — L1: put Toxiproxy between the Go service and Postgres in compose and write a Go test that adds 2s latency and asserts the API returns a timeout error within its budget · L2: run k6 while toggling Toxiproxy faults, define a steady-state hypothesis (error rate < 1%), and prove/disprove it; fix what fails (timeouts, retries, circuit breaker)

---

## QA 16 — Testing in production & release quality
**L1 — aware & can use**
- ⬜ **16.1 Post-deploy verification** — smoke tests after deploy, health checks, synthetic monitoring (scripted probes), canary release, rollback criteria
- ⬜ **16.2 Feature flags for testing** — release toggles, testing both flag states, dark launch, beta users, kill switch

**L2 — deep & efficient**
- ⬜ **16.3 Production quality techniques** — automated canary analysis, shadow / mirrored traffic, traffic replay (GoReplay), A/B tests, observability-driven testing, error budgets as quality gates, real user monitoring (RUM)

🛠 **Practice** — L1: write a post-deploy smoke test (Hurl or Go) that runs in the CI pipeline against the deployed compose/kind environment and fails the job on error · L2: mirror traffic from nginx (`mirror` directive) to a new version of the service and compare responses/errors between old and new

---

## QA 17 — Test automation in CI & quality gates
**L1 — aware & can use**
- ⬜ **17.1 Tests in CI** — running unit/integration/E2E in pipeline stages, fail fast, race detector, test reports (JUnit XML, `gotestsum`), coverage upload, required checks, branch protection
- ⬜ **17.2 Quality gates** — lint, static analysis, coverage thresholds (diff coverage), vulnerability scan, contract checks (`can-i-deploy`), performance thresholds, code review as QA

**L2 — deep & efficient**
- ⬜ **17.3 Scaling test suites** — test sharding, parallelization, test impact analysis / selective testing, caching test results, flaky test tracking dashboards, merge queues, pre-merge vs post-merge suites, cost of CI minutes

🛠 **Practice** — L1: create a GitHub Actions pipeline for the `orders` service with separate unit / integration (Testcontainers) / API (Hurl) jobs and a JUnit report via `gotestsum` · L2: add diff coverage gating, a nightly job with fuzzing + mutation testing + k6 smoke, and a flaky-test retry report

---

## QA 18 — Quality metrics & engineering performance
**L1 — aware & can use**
- ⬜ **18.1 Quality metrics** — defect density, escaped defects, defect removal efficiency, mean time to detect / repair, test pass rate, flakiness rate, coverage (and its limits)
- ⬜ **18.2 DORA metrics** — deployment frequency, lead time for changes, change failure rate, failed deployment recovery time (time to restore), elite vs low performers, DORA research / State of DevOps reports

**L2 — deep & efficient**
- ⬜ **18.3 Using metrics well** — Goodhart's law, vanity metrics, SPACE framework, DORA rework rate (newer metric), measuring teams not individuals, quality dashboards, tying quality to SLOs and customer impact, cost of quality (prevention vs appraisal vs failure)

🛠 **Practice** — L1: compute the 4 DORA metrics for one of your own repos from git history and CI runs (script or spreadsheet) · L2: propose a quality dashboard for your team (5–7 metrics), explaining for each metric how it could be gamed and what counter-metric balances it

---

## QA 19 — Quality in the AI era
**L1 — aware & can use**
- ⬜ **19.1 Testing AI-generated code** — plausible-but-wrong code, hallucinated APIs, reviewing generated tests, tests written before generation as the spec, edge cases AI misses, keeping humans responsible for oracles
- ⬜ **19.2 AI-assisted testing** — generating test cases from requirements, test data generation, AI for exploratory ideas, self-healing UI locators (awareness), risks (tests that mirror bugs, false confidence)

**L2 — deep & efficient**
- ⬜ **19.3 Testing LLM-based features** — non-determinism, evals, golden datasets, LLM-as-judge, regression evals in CI, guardrail tests, prompt injection tests, semantic similarity assertions, tracking quality over model versions

🛠 **Practice** — L1: write the test list and tests for a new function yourself, then let AI implement it; record which cases the AI got wrong · L2: build a small eval suite (20 cases, expected outputs, pass/fail scoring) for an LLM-backed function in Go, run it in CI, and track the score across 2 prompt versions
