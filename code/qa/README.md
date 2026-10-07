# QA & Testing Practice

Roadmap: `roadmap/qa.md` · Notes: `notes/qa/`
Go testing syntax (`testing`, `t.Run`, `httptest`, benchmarks) is in the golang track.

## What practice looks like
- One small Go service (`orders`: create / list / cancel, backed by Postgres) is the system under test.
- Go test files (unit, integration with Testcontainers, fuzz, property-based, mutation runs).
- Non-code artifacts too: test plans, test charters, bug reports, risk matrices, decision tables.
- Tool files: Hurl / Bruno collections, Pact contracts, k6 scripts, Toxiproxy configs, CI workflows.

## Environment setup
```bash
# Postgres for the service (or let Testcontainers start it)
docker run -d --name pg -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres:17

go install gotest.tools/gotestsum@latest
go get github.com/testcontainers/testcontainers-go
go get pgregory.net/rapid

# k6 load testing (no install needed)
docker run --rm -i --network host grafana/k6 run - < load.js

# Hurl API tests
docker run --rm -v $PWD:/w -w /w --network host ghcr.io/orange-opensource/hurl:latest --test api.hurl

# Toxiproxy for fault injection
docker run -d --name toxiproxy -p 8474:8474 -p 15432:15432 ghcr.io/shopify/toxiproxy
```

## Folder convention
```
code/qa/
├── README.md
└── <NN>-<topic-slug>/        # e.g. 07-test-design, 14-performance-k6
    ├── README.md             # exercise instructions (from the 🛠 Practice line of QA <NN>)
    └── ...                   # *_test.go, *.hurl, load.js, test-plan.md, bug reports
```
- `<NN>` = topic number in the roadmap (QA 14 → `14-...`). Create the folder when the topic starts.
- Record results (mutation score, k6 summary, flaky runs) in the exercise README.
