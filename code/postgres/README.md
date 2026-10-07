# Postgres Practice

Hands-on work for `roadmap/postgres.md` (PG 1–15, incl. Redis and other data stores).

## What practice looks like
- **SQL files** run with `psql` against Postgres in Docker (`schema.sql`, `seed.sql`, `queries.sql`).
- **Two `psql` sessions side by side** for transactions, locks and isolation experiments.
- **Go + pgx** (`pgxpool`, sqlc, goose, testcontainers-go) when the topic is about app access, concurrency or benchmarks.
- **`findings.md`** with `EXPLAIN (ANALYZE, BUFFERS)` output and before/after numbers.
- `docker-compose.yml` per exercise when more than one container is needed (replica, PgBouncer, Redis, ClickHouse).

## Environment
```bash
# Postgres (pgbench is included in the image)
docker run -d --name pg -e POSTGRES_PASSWORD=pg -p 5432:5432 \
  -v pgdata:/var/lib/postgresql postgres:18
psql postgres://postgres:pg@localhost:5432/postgres     # or: docker exec -it pg psql -U postgres
psql postgres://postgres:pg@localhost:5432/postgres -f schema.sql

# Redis (PG 14)
docker run -d --name redis -p 6379:6379 redis:8
docker exec -it redis redis-cli

# Go driver
go get github.com/jackc/pgx/v5
```

## Folder convention
```
code/postgres/<NN>-<topic-slug>/
├── README.md        # exercise instructions, copied from the 🛠 Practice line when the topic starts
├── schema.sql       # DDL
├── queries.sql      # solutions
├── findings.md      # measurements, EXPLAIN plans, conclusions
└── *.go / docker-compose.yml   # when needed
```
`NN` = topic number from the roadmap (e.g. `06-transactions`, `08-query-performance`, `14-redis`).
