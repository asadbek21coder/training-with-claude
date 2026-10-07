# Common Software Engineering — Practice

Hands-on exercises for `roadmap/common-swe.md`. Everything here is written by **you**; Claude only writes instructions and skeletons with `TODO`s.

This track is **mixed**, so practice is too:
- **Go** — JWT auth, argon2id password hashing, idempotent APIs, webhooks, SQLi/SSRF toy apps (exploit → fix), RAG with pgvector, MCP server
- **Shell + Git** — scratch repos for Git internals, `bisect`, rebase exercises
- **Docker tools** — Postgres/pgvector, Redis, Keycloak, OpenFGA, ZAP, Trivy, gitleaks
- **Markdown** — STAR stories, brag document, PR reviews, README/runbooks

## Setup
```bash
cd code && go build ./... && go test ./...
docker run -d --name swe-pg -e POSTGRES_PASSWORD=dev -p 5432:5432 pgvector/pgvector:pg17
docker run -d --name swe-redis -p 6379:6379 redis:7
docker run -d --name swe-keycloak -p 8080:8080 -e KC_BOOTSTRAP_ADMIN_USERNAME=admin -e KC_BOOTSTRAP_ADMIN_PASSWORD=admin quay.io/keycloak/keycloak start-dev
go install golang.org/x/vuln/cmd/govulncheck@latest
```
Never put real secrets or API keys in this folder — use `.env` files listed in `.gitignore`.

## Folder convention
```
code/common-swe/
└── <NN>-<topic-slug>/            # e.g. 10-authentication
    ├── README.md                 # exercise instructions (L1 + L2), created when the topic starts from the 🛠 Practice line
    ├── <exercise>/               # Go package, scripts, compose file or Markdown answers
    └── docker-compose.yml        # if the exercise needs services
```

## Exercise README format
1. **Goal** — what you build and which subtopic IDs it covers (e.g. `SWE 10.3, 10.5`)
2. **Steps** — numbered, small
3. **Done when** — tests pass / attack no longer works / you can explain it to your student
4. **Stretch (L2)** — the deeper version
