# Design Patterns & Architecture — Practice

Hands-on exercises for `roadmap/design-patterns.md`. Everything here is written by **you**; Claude only writes instructions and skeletons with `TODO`s.

Practice is mostly **Go**: implement each pattern on a small realistic example (payment client with functional options, caching decorator, order state machine…), refactor messy code into clean layers, and build tiny hexagonal / event-sourced services. Theory topics (ADRs, C4, design docs) are Markdown + Mermaid files.

## Setup
```bash
go version
cd code && go build ./... && go test ./...
# Postgres for repository / event-sourcing exercises
docker run -d --name dp-pg -e POSTGRES_PASSWORD=dev -p 5432:5432 postgres:17
```
All Go code lives in the shared module `code/` (module name `code`). Import paths look like `code/design-patterns/06-behavioral-patterns/orderfsm`.

## Folder convention
```
code/design-patterns/
└── <NN>-<topic-slug>/            # e.g. 04-creational-patterns
    ├── README.md                 # exercise instructions (L1 + L2), created when the topic starts from the 🛠 Practice line
    ├── <exercise>/               # one Go package per exercise
    │   ├── <exercise>.go         # skeleton with TODOs → your implementation
    │   └── <exercise>_test.go    # tests that check behaviour
    └── docs/                     # ADRs, diagrams (Mermaid), design docs for theory topics
```

## Exercise README format
1. **Goal** — what you build and which subtopic IDs it covers (e.g. `DP 6.5`)
2. **Steps** — numbered, small
3. **Done when** — tests pass / you can explain when NOT to use the pattern
4. **Stretch (L2)** — the deeper version
