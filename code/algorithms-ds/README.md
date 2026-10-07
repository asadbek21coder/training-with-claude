# Algorithms & Data Structures — Practice

Practice for `roadmap/algorithms-ds.md` is **Go code**: Claude writes function signatures with `// TODO`, you write the implementations and tests. Claude gives hints when you are stuck, never solutions.

## Setup
The Go module root is `code/` (module name `code`), so imports look like `code/algorithms-ds/ds`.
```bash
cd code
go test ./algorithms-ds/...          # run all tests
go test -race ./algorithms-ds/...    # with the race detector
go test -bench=. -benchmem ./algorithms-ds/<NN>-<topic>/
go run .                             # code/main.go — quick manual experiments
```

## Existing layout (topics already started — keep using it)
| Path | What |
|---|---|
| `ds/` | Go skeletons, package `ds` (`binary_search.go`, `linked_list.go`, `stack_queue.go`, `two_pointers.go`) |
| `ds/ds_imp/` | your implementations of the core structures, package `imp` (imported as `imp "code/algorithms-ds/ds/ds_imp"`) |
| `tasks/<topic>/` | task instructions (`tasks.md`, `questions.md`) for linked-list, stack-queue, binary-tree |
| `notes/algorithms-ds/` (repo root) | notes per topic (EN + RU) |

## Convention for new topics
- New topics go to `code/algorithms-ds/<NN>-<topic-slug>/`, where `NN` is the DSA topic number from the roadmap (e.g. `11-sliding-window/`, `15-dynamic-programming/`).
- `README.md` in that folder holds the exercise instructions, created when the topic starts from the 🛠 **Practice** line and problem checklist in the roadmap.
- Next to it: a skeleton `.go` file with signatures + `// TODO` (written by Claude), and a `_test.go` file (table-driven tests; you may write these yourself).
- You write every implementation. When a problem is solved, tick `[x]` in `roadmap/algorithms-ds.md`.
