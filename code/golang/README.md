# Go — Practice

Hands-on exercises for `roadmap/golang.md`. Everything here is written by **you**; Claude only writes instructions and skeletons with `TODO`s.

## Setup
```bash
go version          # the module in code/ uses the Go version in code/go.mod
cd code && go build ./... && go test ./...
```
All Go code lives in one module: `code/` (module name `code`). Import paths look like `code/golang/08-concurrency-basics/counter`.

## Folder convention
```
code/golang/
└── <NN>-<topic-slug>/            # e.g. 08-concurrency-basics
    ├── README.md                 # exercise instructions (L1 + L2), created when the topic starts
    ├── <exercise>/               # one folder (= Go package) per exercise
    │   ├── <exercise>.go         # skeleton with TODOs → your implementation
    │   └── <exercise>_test.go    # tests that check your work
    └── ...
```

## Exercise README format
1. **Goal** — what you build and which subtopic IDs it covers (e.g. `GO 8.3, 8.4`)
2. **Steps** — numbered, small
3. **Done when** — tests pass / output looks like X / you can explain Y to your student
4. **Stretch (L2)** — the deeper version

## Bigger projects
The "notes" REST API grows through topics 13 → 14 → 17 → 18 → 19 → 21. Keep it in `code/golang/projects/notes-api/` so each topic adds a layer to the same service.
