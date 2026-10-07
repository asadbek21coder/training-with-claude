# PostgreSQL & Data Stores Roadmap

> **L1** = aware & can use · **L2** = deep & efficient · Status: ⬜ not started · 🔄 in progress · ✅ done
> **Notes:** `notes/postgres/` · **Practice:** `code/postgres/` · **Dictionary:** `reference/postgres.md`

**Why (EN):** Data outlives code: a senior backend engineer owns the schema, the transactions, the indexes and the 3 a.m. "database is slow" incident — this is your strongest card next to Go. In the AI era, AI writes plausible SQL and migrations fast, so the value moves to judging them: reading `EXPLAIN`, spotting a lock-taking migration or a race condition, choosing the right data store, and knowing Postgres internals well enough to say *why* (plus pgvector, because embeddings now live next to business data).
**Зачем (RU):** Данные живут дольше кода: senior backend-инженер отвечает за схему, транзакции, индексы и ночной инцидент «база тормозит» — это твоя сильнейшая карта после Go. В эпоху AI SQL и миграции генерируются быстро, поэтому ценность смещается в их оценку: читать `EXPLAIN`, видеть миграцию, которая возьмёт блокировку, или гонку, выбирать правильное хранилище и знать внутренности Postgres, чтобы объяснить *почему* (плюс pgvector — эмбеддинги теперь живут рядом с бизнес-данными).

## Progress
| Layer | Done | Total |
|---|---|---|
| L1 | 0 | 40 |
| L2 | 0 | 50 |

---

## PG 1 — Relational basics & tooling
**L1 — aware & can use**
- ⬜ **1.1 Relational model** — relation (table), tuple (row), attribute (column), domain, primary key, foreign key, set semantics, declarative SQL, DDL vs DML vs DCL vs TCL
- ⬜ **1.2 psql & local environment** — `docker run postgres`, `psql`, connection string (`postgres://user:pass@localhost:5432/db?sslmode=disable`), `\l`, `\c`, `\dt`, `\d+`, `\x`, `\timing`, `\i file.sql`, `\e`, `.psqlrc`
- ⬜ **1.3 Database objects** — cluster, database, schema (`public`), `search_path`, table, view, sequence, function, role, `CREATE`/`ALTER`/`DROP`, `pg_catalog`, `information_schema`

**L2 — deep & efficient**
- ⬜ **1.4 Ecosystem & versions** — yearly major release, minor releases (bug/security fixes), `pg_upgrade`, extensions (`CREATE EXTENSION`), `contrib` modules, managed services (Amazon RDS, Aurora, Cloud SQL, Neon, Supabase), Postgres-compatible databases, release notes reading habit

🛠 **Practice** — L1: run Postgres in Docker, connect with `psql`, create a database + schema, load a sample dataset (e.g. the "pagila" or your own `seed.sql`) with `\i`, explore it with `\d+` · L2: write a `.psqlrc` with useful settings and list installed vs available extensions via `pg_available_extensions`

---

## PG 2 — SQL querying
**L1 — aware & can use**
- ⬜ **2.1 Basic querying** — `SELECT`, `WHERE`, `ORDER BY`, `LIMIT`/`OFFSET`, `DISTINCT`, `DISTINCT ON`, aliases, `LIKE`/`ILIKE`, `IN`, `BETWEEN`, `NULLS FIRST`/`NULLS LAST`
- ⬜ **2.2 NULL semantics** — three-valued logic, `IS NULL`, `IS DISTINCT FROM`, `COALESCE`, `NULLIF`, NULL in aggregates, NULL in `UNIQUE`, `NULLS NOT DISTINCT`, `NOT IN` with NULL trap
- ⬜ **2.3 Joins** — `INNER`, `LEFT`, `RIGHT`, `FULL`, `CROSS`, self-join, `USING` vs `ON`, multi-column join, row multiplication (fan-out), join + aggregate double counting
- ⬜ **2.4 Data modification** — `INSERT ... SELECT`, multi-row `VALUES`, `UPDATE ... FROM`, `DELETE ... USING`, `RETURNING`, UPSERT (`ON CONFLICT DO NOTHING / DO UPDATE`, `EXCLUDED`), `MERGE`, `TRUNCATE` vs `DELETE`, `COPY`
- ⬜ **2.5 Set operations & conditionals** — `UNION` vs `UNION ALL`, `INTERSECT`, `EXCEPT`, `CASE`, `GREATEST`/`LEAST`, `generate_series`, `VALUES` lists, casting `::`

**L2 — deep & efficient**
- ⬜ **2.6 Semi-joins, anti-joins & LATERAL** — `EXISTS` (semi-join), `NOT EXISTS` (anti-join), `NOT IN` vs `NOT EXISTS`, `IN (subquery)`, correlated subquery, scalar subquery, `LATERAL` join, top-N per group, `ANY`/`ALL` with arrays

🛠 **Practice** — L1: in `code/postgres/02-sql-querying/`, create `users/orders/products` tables and write 20 queries in a `.sql` file (filters, all join types, UPSERT, `RETURNING`, `MERGE`) and show the `NOT IN` + NULL trap · L2: solve "latest 3 orders per user" three ways (`LATERAL`, window function, `DISTINCT ON`) and compare results

---

## PG 3 — Aggregation, windows & advanced SQL
**L1 — aware & can use**
- ⬜ **3.1 Aggregation** — `GROUP BY`, `HAVING`, `COUNT(*)` vs `COUNT(col)`, `SUM`, `AVG`, `MIN`/`MAX`, `FILTER (WHERE ...)`, `string_agg`, `array_agg`, `bool_and`/`bool_or`, `COUNT(DISTINCT ...)`
- ⬜ **3.2 CTEs, subqueries & views** — `WITH`, multiple CTEs, subquery in `FROM`/`WHERE`/`SELECT`, derived tables, views, `CREATE OR REPLACE VIEW`, readability vs performance
- ⬜ **3.3 Window functions** — `OVER`, `PARTITION BY`, `ORDER BY` in window, `ROW_NUMBER`, `RANK`, `DENSE_RANK`, `LAG`/`LEAD`, running totals (`SUM() OVER`), moving average, top-N per group
- ⬜ **3.4 Built-in functions** — `now()`, `date_trunc`, `interval`, `extract`, `to_char`, `age`, `concat`, `split_part`, `regexp_replace`, `coalesce`, time zones (`AT TIME ZONE`)

**L2 — deep & efficient**
- ⬜ **3.5 Advanced windows & grouping** — frame clause (`ROWS` vs `RANGE` vs `GROUPS`), `UNBOUNDED PRECEDING`, `FIRST_VALUE`/`LAST_VALUE`, `NTILE`, named `WINDOW`, `GROUPING SETS`, `ROLLUP`, `CUBE`, `percentile_cont`, gaps-and-islands
- ⬜ **3.6 Recursive CTEs & materialized views** — `WITH RECURSIVE`, hierarchy traversal, `CYCLE` / `SEARCH` clauses, `MATERIALIZED` / `NOT MATERIALIZED` CTE, data-modifying CTEs (`WITH d AS (DELETE ... RETURNING)`), materialized views, `REFRESH MATERIALIZED VIEW CONCURRENTLY`
- ⬜ **3.7 Functions, procedures & triggers** — SQL vs PL/pgSQL functions, procedures (`CALL`), triggers (`BEFORE`/`AFTER`, row vs statement), `IMMUTABLE`/`STABLE`/`VOLATILE`, `SECURITY DEFINER`, exception handling, when logic in the DB hurts (testing, deploys, hidden side effects)

🛠 **Practice** — L1: write a sales report in SQL — revenue per month, per category, running total and month-over-month change with `LAG` · L2: model an org chart, query the full reporting chain with `WITH RECURSIVE`, and build an `updated_at` trigger + a materialized view refreshed concurrently

---

## PG 4 — Data modeling
**L1 — aware & can use**
- ⬜ **4.1 Normalization** — functional dependency, 1NF, 2NF, 3NF, BCNF, redundancy, insert/update/delete anomalies, denormalization, read vs write trade-off
- ⬜ **4.2 Keys & constraints** — primary key, natural vs surrogate key, `bigint GENERATED ALWAYS AS IDENTITY`, `bigserial`, UUID v4 (`gen_random_uuid()`) vs v7 (`uuidv7()`), foreign key, `ON DELETE CASCADE / RESTRICT / SET NULL`, `UNIQUE`, `CHECK`, `NOT NULL`, `DEFAULT`
- ⬜ **4.3 Data types** — `integer`/`bigint`, `numeric` vs `float` (money!), `text` vs `varchar(n)`, `boolean`, `timestamptz` vs `timestamp`, `date`, `interval`, `jsonb` vs `json`, arrays, enum vs lookup table, `bytea`, `uuid`, `inet`
- ⬜ **4.4 Relationships & ER diagrams** — one-to-one, one-to-many, many-to-many (junction table), ER diagram, cardinality, optional vs mandatory, naming conventions (`snake_case`, singular vs plural), FK indexes

**L2 — deep & efficient**
- ⬜ **4.5 Modeling patterns** — soft delete (`deleted_at`), audit columns, history / temporal tables, polymorphic associations, EAV anti-pattern, trees (adjacency list, materialized path / `ltree`, closure table, nested set), status state machines, append-only ledger
- ⬜ **4.6 Multi-tenancy** — `tenant_id` column, composite keys with tenant, schema-per-tenant, database-per-tenant, RLS for tenant isolation, noisy neighbor, tenant-aware indexes, tenant migration/export
- ⬜ **4.7 Advanced constraints & types** — exclusion constraint (`EXCLUDE USING gist`), range types (`tstzrange`, `daterange`), deferrable constraints (`DEFERRABLE INITIALLY DEFERRED`), domains, generated columns (`GENERATED ALWAYS AS (...) STORED`), composite types, `jsonb` vs real columns, column order & alignment padding

🛠 **Practice** — L1: design an e-commerce schema (users, products, orders, order_items, payments) in `schema.sql` with all constraints, draw its ER diagram, and normalize a deliberately bad "one big table" to 3NF · L2: build a room-booking table where overlapping bookings are rejected by an exclusion constraint, and implement a category tree three ways (adjacency list, `ltree`, closure table) comparing queries

---

## PG 5 — Postgres from Go
**L1 — aware & can use**
- ⬜ **5.1 Drivers & pools** — `database/sql` vs native `pgx/v5`, `pgxpool`, `pgx/v5/stdlib`, DSN, `Query`/`QueryRow`/`Exec`, `rows.Close()`, `pgx.CollectRows`, `pgx.RowToStructByName`, `pgx.ErrNoRows`, scanning NULLs (pointers, `pgtype`)
- ⬜ **5.2 Queries & transactions in Go** — placeholders `$1`, SQL injection, `pool.Begin`, `tx.Commit`/`tx.Rollback`, `defer tx.Rollback(ctx)`, `pgx.BeginFunc`, `context` timeouts, `*pgconn.PgError`, SQLSTATE codes (`23505` unique_violation, `40001`)
- ⬜ **5.3 Migrations & query tooling** — goose, golang-migrate, Atlas, up/down migrations, embedded migrations (`embed.FS`), `sqlc` (code generation from SQL), squirrel (query builder), ORM trade-offs (GORM, ent)

**L2 — deep & efficient**
- ⬜ **5.4 Efficient access** — `pgx.Batch` / `SendBatch`, `CopyFrom`, statement cache, `QueryExecModeSimpleProtocol` / exec modes with PgBouncer, pool sizing (`MaxConns`, `MinConns`, `MaxConnLifetime`, `MaxConnIdleTime`), `pool.Stat()`, `LISTEN`/`NOTIFY` (`WaitForNotification`), streaming large result sets, `tracelog` / OpenTelemetry tracing
- ⬜ **5.5 Testing with Postgres** — testcontainers-go, transaction-per-test rollback, template databases (`CREATE DATABASE ... TEMPLATE`), schema-per-test for parallel tests, fixtures/factories, pgxmock trade-offs, testing migrations up and down

🛠 **Practice** — L1: build a small Go repository layer (`code/postgres/05-go/`) for the e-commerce schema with `pgxpool`, sqlc-generated queries, goose migrations, and a `CreateOrder` function in one transaction that maps `23505` to a domain error · L2: insert 1M rows three ways (row-by-row, `pgx.Batch`, `CopyFrom`), benchmark with `testing.B`, and write integration tests with testcontainers-go

---

## PG 6 — Transactions & concurrency
**L1 — aware & can use**
- ⬜ **6.1 ACID & transactions** — `BEGIN`/`COMMIT`/`ROLLBACK`, atomicity, consistency, isolation, durability, autocommit, `SAVEPOINT`, `ROLLBACK TO SAVEPOINT`, aborted transaction state (`current transaction is aborted`)
- ⬜ **6.2 Isolation levels & anomalies** — Read Committed (default), Repeatable Read (snapshot isolation), Serializable (SSI), dirty read (impossible in Postgres), non-repeatable read, phantom read, lost update, write skew, `SET TRANSACTION ISOLATION LEVEL`
- ⬜ **6.3 Row locking basics** — `SELECT ... FOR UPDATE`, `FOR SHARE`, `NOWAIT`, `SKIP LOCKED`, `UPDATE` takes a row lock, blocking, `lock_timeout`, atomic `UPDATE ... SET x = x + 1`

**L2 — deep & efficient**
- ⬜ **6.4 Lock modes & the lock queue** — table lock modes (`ACCESS SHARE` ... `ACCESS EXCLUSIVE`), conflict matrix, `FOR NO KEY UPDATE`, `FOR KEY SHARE`, FK locks, lock queue (DDL waiting behind a long transaction blocks everyone), `pg_locks`, `pg_blocking_pids()`, deadlock detection, `deadlock_timeout`, consistent lock ordering
- ⬜ **6.5 Serializable & retries** — SSI, predicate locks (`SIReadLock`), serialization failure `40001`, deadlock `40P01`, retry loop with backoff, `READ ONLY DEFERRABLE`, cost of Serializable, when Repeatable Read is enough
- ⬜ **6.6 Concurrency patterns** — optimistic locking (`version` column), pessimistic locking, Postgres as a job queue (`FOR UPDATE SKIP LOCKED`), advisory locks (`pg_advisory_xact_lock`), counters under contention (sharded counters), idempotency-key table, `ON CONFLICT` races, check-then-insert anti-pattern
- ⬜ **6.7 Long transactions** — `idle in transaction`, `idle_in_transaction_session_timeout`, `statement_timeout`, `transaction_timeout`, old snapshot blocks VACUUM, network/HTTP calls inside a transaction anti-pattern, `pg_stat_activity.xact_start`

🛠 **Practice** — L1: open two `psql` sessions side by side and reproduce a non-repeatable read in Read Committed, then show it disappears in Repeatable Read; reproduce a lost update and fix it with `FOR UPDATE` · L2: write a Go program with 50 goroutines withdrawing from the same account under each fix (atomic update, optimistic version, `FOR UPDATE`, Serializable + retry), count retries/errors, and build a `SKIP LOCKED` job worker pool

---

## PG 7 — Indexes
**L1 — aware & can use**
- ⬜ **7.1 Index basics** — why indexes, B-tree (default), `CREATE INDEX`, unique index, indexing FK columns, selectivity, index scan vs seq scan, `\di`, `pg_indexes`
- ⬜ **7.2 Composite indexes** — column order, leftmost prefix, equality-first-then-range rule, sort order (`DESC`, `NULLS LAST`), index for `ORDER BY ... LIMIT`, multicolumn vs separate indexes (BitmapAnd)

**L2 — deep & efficient**
- ⬜ **7.3 B-tree internals** — pages, root/internal/leaf, tree height, page splits, deduplication, index `fillfactor`, index-only scans and the visibility map, `amcheck`, `pageinspect`
- ⬜ **7.4 Advanced B-tree usage** — partial index (`WHERE`), covering index (`INCLUDE`), expression index (`lower(email)`), unique partial index, `text_pattern_ops` for `LIKE 'abc%'`, collation impact, skip scan awareness
- ⬜ **7.5 Other index types** — GIN (`jsonb @>`, arrays, full-text, `pg_trgm`), GiST (ranges, geometry, exclusion), SP-GiST, BRIN (append-only time-series), hash index, bloom (extension), HNSW / IVFFlat (pgvector)
- ⬜ **7.6 Index costs & maintenance** — write amplification, indexes block HOT updates, index bloat, unused indexes (`pg_stat_user_indexes.idx_scan`), duplicate/redundant indexes, `CREATE INDEX CONCURRENTLY`, invalid indexes, `REINDEX CONCURRENTLY`

🛠 **Practice** — L1: generate 5M rows with `generate_series`, time a query before/after an index, and prove the leftmost-prefix rule with a 2-column index and `EXPLAIN` · L2: compare B-tree vs BRIN size and speed on a time-series table, build a GIN index for `jsonb @>` and a `pg_trgm` index for `ILIKE '%term%'`, then measure insert slowdown with 0 vs 5 indexes using `pgbench`

---

## PG 8 — Query performance
**L1 — aware & can use**
- ⬜ **8.1 Reading EXPLAIN** — `EXPLAIN`, `EXPLAIN (ANALYZE, BUFFERS)`, cost (startup..total), estimated vs actual rows, loops, Seq Scan, Index Scan, Index Only Scan, Bitmap Heap Scan, Sort, Limit, explain.dalibo.com / explain.depesz.com
- ⬜ **8.2 Common query problems** — N+1 queries, `OFFSET` pagination vs keyset (cursor) pagination, `SELECT *`, functions on indexed columns, implicit casts, `OR` conditions, `LIKE '%x'`, missing FK indexes, `count(*)` on big tables

**L2 — deep & efficient**
- ⬜ **8.3 Join algorithms & plan nodes** — Nested Loop, Hash Join, Merge Join, HashAggregate vs GroupAggregate, Materialize, Memoize, parallel query (Gather, workers), `work_mem` spill (external merge sort, hash batches), JIT
- ⬜ **8.4 Planner & statistics** — `ANALYZE`, `pg_stats`, `default_statistics_target`, most common values, histogram, `n_distinct`, correlation, extended statistics (`CREATE STATISTICS`), cardinality misestimates, generic vs custom plans (`plan_cache_mode`), `pg_hint_plan` (awareness)
- ⬜ **8.5 Monitoring & diagnosis** — `pg_stat_statements`, `pg_stat_activity`, wait events, `auto_explain`, `log_min_duration_statement`, `pg_stat_user_tables`, `pg_stat_io`, `pg_locks`, pgBadger, top queries by total time
- ⬜ **8.6 Tuning approach** — measure first, fix the query vs add an index vs change the schema, precomputation / summary tables, materialized views, caching layer, read replica offload, plan regressions after upgrade, load testing with `pgbench`

🛠 **Practice** — L1: take 5 slow queries on your 5M-row dataset, read their `EXPLAIN (ANALYZE, BUFFERS)`, fix each and write before/after timings in `findings.md`; rewrite an `OFFSET 100000` page as keyset pagination · L2: enable `pg_stat_statements` + `auto_explain`, run a `pgbench` custom script, find the top query by total time, force a misestimate with correlated columns and fix it with `CREATE STATISTICS`

---

## PG 9 — Internals
**L1 — aware & can use**
- ⬜ **9.1 Architecture overview** — postmaster, backend process per connection, shared memory, `shared_buffers`, WAL, background writer, checkpointer, autovacuum launcher, why connections are expensive

**L2 — deep & efficient**
- ⬜ **9.2 MVCC** — tuple versions, `xmin`/`xmax`, `ctid`, transaction ID (XID), snapshot, visibility rules, commit log (`pg_xact`), dead tuples, `UPDATE` = new tuple version, hint bits
- ⬜ **9.3 VACUUM & bloat** — `VACUUM`, autovacuum (threshold, scale factor, cost limit), `VACUUM FULL`, `pg_repack`, bloat, visibility map, free space map, freezing, XID wraparound, `autovacuum_freeze_max_age`, HOT updates, `fillfactor`
- ⬜ **9.4 Storage layout** — 8 KB pages, heap files, page layout (line pointers, tuple header), TOAST, relfilenode, `pg_relation_size`, `pg_total_relation_size`, tablespaces, alignment padding
- ⬜ **9.5 WAL & durability** — WAL, LSN, checkpoint, full-page writes, `fsync`, `synchronous_commit`, `wal_level`, crash recovery (redo), `checkpoint_timeout`, `max_wal_size`, group commit
- ⬜ **9.6 Buffer management & I/O** — shared buffers, clock-sweep eviction, OS page cache (double buffering), dirty pages, background writer, `pg_buffercache`, `effective_io_concurrency`, asynchronous I/O (awareness)
- ⬜ **9.7 Query execution pipeline** — parser, rewriter, planner/optimizer, executor, iterator (Volcano) model, simple vs extended query protocol, prepared statements, catalog caches

🛠 **Practice** — L1: run `ps` inside the container to see the process tree, open 20 connections and watch new backends appear · L2: inspect `xmin`/`xmax`/`ctid` while updating a row, create bloat with mass updates, measure table size before/after `VACUUM` vs `VACUUM FULL`, and compare `pgbench` TPS with `synchronous_commit` on vs off

---

## PG 10 — Operations
**L1 — aware & can use**
- ⬜ **10.1 Connection management** — `max_connections`, connection pooling, PgBouncer, session vs transaction vs statement pooling, prepared statements with poolers, app-side pool vs external pooler, `too many clients`
- ⬜ **10.2 Backup & restore basics** — `pg_dump`, `pg_restore`, plain/custom/directory formats, `pg_dumpall`, logical vs physical backup, restore testing, RPO, RTO
- ⬜ **10.3 Roles & access** — roles, users, `GRANT`/`REVOKE`, least privilege, default privileges, schema ownership, `pg_hba.conf`, SCRAM-SHA-256, TLS (`sslmode=verify-full`), app role vs migration role

**L2 — deep & efficient**
- ⬜ **10.4 Replication** — streaming (physical) replication, WAL sender/receiver, replication slots, hot standby, sync vs async replica, `synchronous_standby_names`, replication lag (`pg_stat_replication`), logical replication (publication/subscription), query conflicts on standby, `hot_standby_feedback`
- ⬜ **10.5 HA & failover** — failover vs switchover, promote, Patroni, etcd as DCS, split brain, fencing, `pg_rewind`, HAProxy / VIP routing, managed HA (RDS Multi-AZ), client reconnect behavior
- ⬜ **10.6 Physical backup & PITR** — `pg_basebackup`, WAL archiving (`archive_command` / `archive_library`), PITR (`recovery_target_time`), pgBackRest, WAL-G, Barman, incremental backup, backup verification
- ⬜ **10.7 Configuration & tuning** — `shared_buffers`, `work_mem`, `maintenance_work_mem`, `effective_cache_size`, `random_page_cost`, `max_wal_size`, `ALTER SYSTEM`, `pg_settings`, reload vs restart, per-role/per-database settings, PGTune
- ⬜ **10.8 Security hardening** — row-level security (`ENABLE ROW LEVEL SECURITY`, `CREATE POLICY`), `current_setting('app.tenant_id')`, `BYPASSRLS`, column privileges, pgaudit, `pgcrypto`, encryption at rest, secrets rotation, `log_statement` and PII

🛠 **Practice** — L1: put PgBouncer (transaction mode) in front of Postgres with docker compose, run 500 client connections through it with `pgbench`; take a `pg_dump -Fc`, drop the DB and restore it · L2: build a primary + streaming replica in docker compose, measure lag under load, promote the replica; do a PITR to "5 minutes before the accidental `DELETE`" with WAL archiving

---

## PG 11 — Schema evolution
**L1 — aware & can use**
- ⬜ **11.1 Migration basics** — versioned migrations, up/down vs forward-only, migrations in CI/CD, migration review checklist, one change per migration, transactional DDL
- ⬜ **11.2 Safe DDL** — which DDL takes `ACCESS EXCLUSIVE`, `lock_timeout`, fast `ADD COLUMN` with constant default, `CREATE INDEX CONCURRENTLY`, `ADD CONSTRAINT ... NOT VALID` + `VALIDATE CONSTRAINT`, dropping columns safely

**L2 — deep & efficient**
- ⬜ **11.3 Zero-downtime patterns** — expand-and-contract, adding `NOT NULL` safely (`CHECK ... NOT VALID` → validate → `SET NOT NULL`), renaming columns safely, changing column type, backfill in batches, dual writes, feature flags, app/DB deploy ordering
- ⬜ **11.4 Large data changes & upgrades** — batched `UPDATE`/`DELETE`, `pg_repack`, online table rewrite, major upgrade via `pg_upgrade --link`, upgrade via logical replication, blue-green deployment, squawk (migration linter)

🛠 **Practice** — L1: write a migration that adds a column + index the safe way and show (with `pg_locks` in a second session) that the naive version blocks reads behind a long transaction · L2: rename `users.name` → `full_name` with zero downtime using expand-and-contract while a Go load generator keeps writing, and backfill 5M rows in batches

---

## PG 12 — Postgres features beyond plain tables
**L1 — aware & can use**
- ⬜ **12.1 JSONB** — `->`, `->>`, `#>>`, `@>`, `?`, `jsonb_set`, `jsonb_build_object`, `jsonb_agg`, `jsonb_path_query` (SQL/JSON path), GIN index, `jsonb` vs columns decision
- ⬜ **12.2 Full-text search** — `tsvector`, `tsquery`, `to_tsvector`, `plainto_tsquery`, `websearch_to_tsquery`, `ts_rank`, GIN index, text search configuration (language), generated `tsvector` column, `pg_trgm` for fuzzy matching

**L2 — deep & efficient**
- ⬜ **12.3 Events & queues in Postgres** — `LISTEN`/`NOTIFY`, payload limit, triggers emitting events, transactional outbox table, `SKIP LOCKED` workers, River (Go), pgmq, when to move to a real broker
- ⬜ **12.4 Extensions ecosystem** — `pg_stat_statements`, `pgcrypto`, `pg_trgm`, PostGIS, pgvector, TimescaleDB, `pg_partman`, `pg_cron`, Citus, `hstore`, `postgres_fdw`
- ⬜ **12.5 pgvector & AI workloads** — `vector` type, embeddings, distance operators (`<->`, `<=>`, `<#>`), HNSW vs IVFFlat, recall vs speed, `ef_search`, filtered ANN search, hybrid search (full-text + vector), storage cost of dimensions

🛠 **Practice** — L1: store product attributes in `jsonb` with a GIN index and add full-text search with `websearch_to_tsquery` + ranking over product descriptions · L2: in Go, embed ~10k text chunks (any embedding API or a local model), store them in pgvector, compare exact vs HNSW search latency/recall, and build a `LISTEN`/`NOTIFY` listener that reacts to new rows

---

## PG 13 — Scaling Postgres
**L1 — aware & can use**
- ⬜ **13.1 Scaling options overview** — vertical scaling first, query/index fixes, connection pooling, caching, read replicas, partitioning, archiving old data, sharding, when a single node is no longer enough

**L2 — deep & efficient**
- ⬜ **13.2 Partitioning** — declarative partitioning (range, list, hash), partition key choice, partition pruning, default partition, `ATTACH`/`DETACH PARTITION CONCURRENTLY`, indexes on partitioned tables, `pg_partman`, dropping old partitions vs `DELETE`, when partitioning hurts
- ⬜ **13.3 Read/write splitting** — routing reads to replicas, replication lag, read-your-writes, `pg_last_wal_replay_lsn()`, LSN-based causal reads, stale-read tolerance per query, pgx multi-pool routing
- ⬜ **13.4 Sharding & distributed Postgres** — application-level sharding, shard key, cross-shard queries, resharding, Citus (distributed vs reference tables), CockroachDB / YugabyteDB (awareness), compute/storage separation (Aurora, Neon)

🛠 **Practice** — L1: write a one-page "how I'd scale this DB from 1k to 100k QPS" plan for your e-commerce schema in `scaling.md` · L2: partition a 20M-row `events` table by month, show pruning in `EXPLAIN`, drop an old month instantly vs `DELETE`; add a replica and route Go reads to it with read-your-writes protection

---

## PG 14 — Redis
**L1 — aware & can use**
- ⬜ **14.1 Redis basics** — in-memory key-value store, single-threaded command execution, `redis-cli`, `SET`/`GET`, `EX`/`PX`, `EXPIRE`, `TTL`, `INCR`, `DEL`, `KEYS` vs `SCAN`, go-redis (`redis/go-redis/v9`)
- ⬜ **14.2 Data structures** — strings, hashes, lists, sets, sorted sets (ZSET), streams, bitmaps, HyperLogLog, geospatial indexes, JSON (awareness)
- ⬜ **14.3 Common use cases** — cache-aside, session store, rate limiter (`INCR` + `EXPIRE`, ZSET sliding window), leaderboard, counters, simple queue, distributed lock (`SET key val NX PX`), dedup / idempotency keys

**L2 — deep & efficient**
- ⬜ **14.4 Atomicity & performance** — `MULTI`/`EXEC`, `WATCH` (optimistic), pipelining, Lua scripts (`EVAL`, `EVALSHA`), Redis Functions, round trips, big keys, hot keys, O(N) commands blocking the server, `SLOWLOG`
- ⬜ **14.5 Memory & persistence** — `maxmemory`, eviction policies (`allkeys-lru`, `allkeys-lfu`, `volatile-ttl`, `noeviction`), RDB snapshots, AOF (`appendfsync`), fork cost, memory fragmentation, `INFO memory`, `MEMORY USAGE`
- ⬜ **14.6 Pub/sub & streams** — pub/sub (fire-and-forget), Streams (`XADD`, `XREADGROUP`, consumer groups, `XACK`, pending entries list, `XAUTOCLAIM`), stream trimming, Redis Streams vs Kafka
- ⬜ **14.7 HA, scaling & distributed locks** — replication, Sentinel, Redis Cluster (16384 hash slots, hash tags `{}`, `MOVED`/`ASK`), client-side caching, Valkey (fork) and license history, managed (ElastiCache, Memorystore), Redlock and its critique, fencing tokens

🛠 **Practice** — L1: run Redis in Docker and in Go implement a cache-aside layer over your Postgres repository with TTL, plus a ZSET leaderboard and a fixed-window rate limiter · L2: implement a sliding-window rate limiter as a Lua script, benchmark pipelined vs non-pipelined writes, fill memory past `maxmemory` with different eviction policies, and build a Streams consumer group that recovers pending messages with `XAUTOCLAIM`

---

## PG 15 — Other data stores (when to choose what)
**L1 — aware & can use**
- ⬜ **15.1 Choosing a data store** — access patterns first, OLTP vs OLAP, read/write ratio, consistency needs, data size and growth, query flexibility, operational cost, "just use Postgres" default, polyglot persistence
- ⬜ **15.2 Data store families** — relational, key-value, document (MongoDB), wide-column (Cassandra, ScyllaDB), columnar OLAP (ClickHouse), search (Elasticsearch / OpenSearch), time-series (TimescaleDB, Prometheus TSDB), graph (Neo4j), object storage (S3), vector databases

**L2 — deep & efficient**
- ⬜ **15.3 Storage engine internals** — B-tree vs LSM-tree, memtable, SSTable, compaction (size-tiered vs leveled), write amplification, read amplification, space amplification, bloom filters, row vs column storage, compression
- ⬜ **15.4 Columnar & analytics** — ClickHouse (`MergeTree`, ordering key, materialized views), DuckDB, Parquet, data warehouses (BigQuery, Snowflake), CDC to OLAP (Debezium), ETL vs ELT, star schema
- ⬜ **15.5 Search engines** — inverted index, analyzer / tokenizer, BM25 relevance, Elasticsearch / OpenSearch shards & replicas, near-real-time refresh, syncing from Postgres (CDC / outbox), when Postgres FTS is enough
- ⬜ **15.6 NoSQL modeling** — MongoDB (documents, embedding vs referencing, replica sets), Cassandra / ScyllaDB (partition key, clustering key, tunable consistency, query-first modeling), DynamoDB (partition + sort key, single-table design, GSI, on-demand vs provisioned)
- ⬜ **15.7 Object & vector storage** — S3 (buckets, objects, presigned URLs, multipart upload, storage classes, lifecycle rules, strong read-after-write consistency), S3-compatible local storage for dev, dedicated vector DBs (Qdrant, Weaviate, Milvus) vs pgvector

🛠 **Practice** — L1: for 5 scenarios (orders, product search, click analytics, sessions, chat history) write in `choices.md` which store you'd pick and why, in 3 lines each · L2: load the same 10M-row events dataset into Postgres and ClickHouse (both in Docker), run the same aggregate query and compare time and disk size; explain the difference with row vs column storage
