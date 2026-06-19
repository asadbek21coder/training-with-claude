# 04 — Databases

> Know both theory (ACID, indexes, normalization) and practice (SQL queries, query optimization, tools).

---

## SQL Basics

**Key Terms:**
- `DDL (Data Definition Language)` — CREATE, ALTER, DROP, TRUNCATE
- `DML (Data Manipulation Language)` — SELECT, INSERT, UPDATE, DELETE
- `schema` — structure of tables, columns, types, constraints
- `primary key (PK)` — unique identifier for each row; auto-indexed; NOT NULL
- `foreign key (FK)` — references PK in another table; enforces referential integrity
- `unique constraint` — column values must be unique across rows
- `NOT NULL` — column cannot be empty
- `DEFAULT` — value used if none provided
- `CHECK` — constraint that validates column value
- `alias` — `SELECT name AS username FROM users`
- `DISTINCT` — remove duplicate rows from result
- `ORDER BY` — sort result; ASC (default) or DESC
- `LIMIT / OFFSET` — pagination; OFFSET is slow on large tables (use cursor instead)
- `LIKE` — pattern matching; `%` = any sequence, `_` = single char

**JOIN Types:**

| Join | Returns |
|------|---------|
| INNER JOIN | Only rows matching in both tables |
| LEFT JOIN | All rows from left + matching from right (NULL if no match) |
| RIGHT JOIN | All rows from right + matching from left |
| FULL OUTER JOIN | All rows from both; NULL where no match |
| CROSS JOIN | Cartesian product; every row × every row |
| SELF JOIN | Table joined with itself |

**Aggregation:**
- `COUNT(*)` — number of rows
- `SUM, AVG, MIN, MAX` — aggregate functions
- `GROUP BY` — group rows by column value before aggregating
- `HAVING` — filter after GROUP BY (WHERE filters before GROUP BY)
- `WINDOW FUNCTIONS` — ROW_NUMBER(), RANK(), LAG(), LEAD(), SUM() OVER() — aggregate without collapsing rows

**Must understand:**
- [ ] Difference between WHERE and HAVING
- [ ] When to use subquery vs JOIN
- [ ] CTE (WITH clause) — common table expression for readable complex queries
- [ ] Window functions — ROW_NUMBER to deduplicate, LAG/LEAD for time-series

---

## Indexes

**Key Terms:**
- `index` — separate data structure that speeds up lookups on a column
- `B-tree index` — default; balanced tree; O(log n) search; supports range queries
- `B+ tree` — leaf nodes linked in a list; enables fast range scans; used by PostgreSQL/MySQL
- `hash index` — exact match only; O(1); no range queries (PostgreSQL supports this)
- `clustered index` — table data stored in index order; MySQL InnoDB = PK is always clustered
- `non-clustered index (heap)` — index stores pointer to row; needs extra lookup (heap fetch)
- `composite index` — index on multiple columns; order matters
- `leftmost prefix rule` — composite index `(a,b,c)` helps queries on `(a)`, `(a,b)`, `(a,b,c)` but NOT `(b)` alone
- `covering index` — index includes all columns the query needs → no heap fetch → faster
- `index-only scan` — query satisfied entirely from index; very fast
- `partial index` — index with WHERE clause; smaller index, only indexes relevant rows
- `functional index` — index on expression: `CREATE INDEX ON users (lower(email))`
- `bloat` — index grows stale with DELETEs/UPDATEs; VACUUM/REINDEX needed
- `EXPLAIN` — show query execution plan; look for Seq Scan (bad) vs Index Scan (good)
- `EXPLAIN ANALYZE` — executes query and shows actual vs estimated rows

**When indexes hurt:**
- Write-heavy tables: every INSERT/UPDATE/DELETE must update all indexes
- Low-selectivity columns (e.g., boolean column: index won't help)
- Very small tables (seq scan may be faster)

**Must understand:**
- [ ] Why composite index `(a,b)` doesn't help `WHERE b = ?`
- [ ] What a covering index is and how to create one
- [ ] When the query planner chooses seq scan over index (small table, low selectivity)
- [ ] How to read EXPLAIN output to diagnose slow queries

---

## Transactions & ACID

**Key Terms:**
- `transaction` — a sequence of operations treated as a single unit
- `ACID` — set of properties guaranteeing reliable transactions
- `Atomicity` — all operations succeed or none do (no partial transactions)
- `Consistency` — DB moves from one valid state to another; constraints always satisfied
- `Isolation` — concurrent transactions don't see each other's intermediate state
- `Durability` — committed transactions survive crashes (written to disk/WAL)
- `BEGIN / START TRANSACTION` — start transaction
- `COMMIT` — make changes permanent
- `ROLLBACK` — undo all changes since BEGIN
- `savepoint` — point within transaction you can roll back to partially
- `WAL (Write-Ahead Log)` — changes written to log first; ensures durability + recovery
- `2PC (Two-Phase Commit)` — protocol for distributed transactions across multiple DBs
- `XA transactions` — standard for 2PC across different systems

**Must understand:**
- [ ] What "atomicity" means — why partial writes are dangerous
- [ ] How WAL provides durability (log survives crash; DB replays on recovery)
- [ ] Why 2PC is expensive and rarely used in microservices (use Saga pattern instead)
- [ ] ROLLBACK scenarios — when and why to rollback

---

## Isolation Levels

**Key Terms:**
- `dirty read` — read uncommitted data from another transaction (data may be rolled back)
- `non-repeatable read` — same query returns different result within same transaction (row updated by other)
- `phantom read` — same query returns different rows within same transaction (rows inserted by other)
- `lost update` — two transactions read-modify-write same row; second overwrites first's change

**Isolation Levels (weakest → strongest):**

| Level | Dirty Read | Non-Repeatable | Phantom |
|-------|------------|----------------|---------|
| Read Uncommitted | ✅ possible | ✅ possible | ✅ possible |
| Read Committed | ❌ prevented | ✅ possible | ✅ possible |
| Repeatable Read | ❌ prevented | ❌ prevented | ✅ possible (MySQL: prevented by default) |
| Serializable | ❌ prevented | ❌ prevented | ❌ prevented |

- `Read Committed` — PostgreSQL default
- `Repeatable Read` — MySQL InnoDB default
- `Serializable` — safest; slowest; uses locking or MVCC
- `MVCC (Multi-Version Concurrency Control)` — each transaction sees a snapshot; readers don't block writers
- `pessimistic locking` — lock row before reading/writing (SELECT FOR UPDATE)
- `optimistic locking` — no lock; check version on write; retry if conflict
- `SELECT FOR UPDATE` — lock rows for update; other transactions wait

**Must understand:**
- [ ] What anomaly each isolation level prevents
- [ ] Why Serializable is rarely used (performance cost)
- [ ] MVCC — how PostgreSQL implements isolation without reader/writer blocking
- [ ] Optimistic vs pessimistic locking — when to use each

---

## Database Normalization

**Key Terms:**
- `normalization` — process of organizing data to reduce redundancy
- `anomaly` — problems caused by redundancy: update anomaly, insert anomaly, delete anomaly
- `functional dependency` — column B depends on column A (knowing A determines B)
- `1NF (First Normal Form)` — no repeating groups; each cell has atomic value; primary key exists
- `2NF` — 1NF + no partial dependency (non-key column depends on entire PK, not part of it)
- `3NF` — 2NF + no transitive dependency (non-key column doesn't depend on another non-key column)
- `BCNF (Boyce-Codd)` — stricter 3NF; every determinant is a candidate key
- `denormalization` — intentionally add redundancy for read performance
- `join cost` — normalized data requires JOINs; denormalized avoids them at cost of redundancy

**Must understand:**
- [ ] Why normalization reduces anomalies
- [ ] 1NF, 2NF, 3NF with examples — be able to spot violations
- [ ] When to intentionally denormalize (read-heavy, analytics, caching)
- [ ] Normalized vs denormalized trade-off in system design

---

## Query Optimization

**Key Terms:**
- `query planner / optimizer` — DB component that chooses how to execute a query
- `EXPLAIN` — show execution plan without running query
- `EXPLAIN ANALYZE` — run query and show actual rows, time
- `seq scan` — full table scan; indicates missing index
- `index scan` — uses index to find rows; efficient
- `index-only scan` — covering index; no heap fetch; fastest
- `nested loop join` — for each row in outer, scan inner; O(n×m)
- `hash join` — build hash table from smaller table; probe with larger; O(n+m)
- `merge join` — both sides sorted; single pass; O(n log n) with sort
- `statistics` — DB maintains stats on column cardinality; used by planner; ANALYZE updates
- `N+1 problem` — 1 query to get N records, then N queries for related data; fix with JOIN or eager load
- `batch queries` — replace N queries with one query using IN clause
- `connection pooling` — reuse DB connections; avoid expensive connect/disconnect per request
- `query hints` — force planner to use specific index (last resort; usually bad idea)

**Must understand:**
- [ ] How to identify a slow query (pg_stat_statements, slow query log)
- [ ] What EXPLAIN ANALYZE output means (actual rows vs estimated, cost)
- [ ] N+1 problem — how to detect and fix in Go (with sqlx) and Java (with JPA)
- [ ] Index not used even though it exists — reasons (low selectivity, type mismatch, function on column)

---

## PostgreSQL

**Key Terms:**
- `JSONB` — binary JSON storage; indexed; queryable with operators (@>, ?, #>)
- `ARRAY type` — store array in column; searchable with ANY, ALL
- `CTE (WITH clause)` — name a subquery for reuse; WITH RECURSIVE for tree traversal
- `window functions` — ROW_NUMBER(), RANK(), DENSE_RANK(), LAG(), LEAD(), NTILE(), PARTITION BY
- `LATERAL join` — subquery can reference columns from preceding FROM clause
- `VACUUM` — reclaims space from dead rows; AUTOVACUUM runs automatically
- `ANALYZE` — updates statistics for query planner
- `pg_stat_statements` — extension; tracks execution stats per query
- `pg_stat_activity` — view of active connections and running queries
- `LISTEN/NOTIFY` — lightweight pub/sub inside PostgreSQL
- `advisory locks` — application-level distributed locks stored in PostgreSQL
- `tablespace` — where DB files are stored; can place hot data on fast disk
- `partitioning` — split large table into smaller child tables (range, list, hash)
- `replication slots` — ensure WAL retained until replica consumes it (Logical/Physical)
- `logical replication` — replicate specific tables; enable CDC (Change Data Capture)

**Must understand:**
- [ ] JSONB vs JSON — JSONB is indexed and faster for queries
- [ ] Window functions — use ROW_NUMBER to deduplicate, RANK for ranking with ties
- [ ] VACUUM — why you need it (dead rows accumulate; cause table bloat)
- [ ] Partitioning — when it helps (very large tables, range-based queries)

---

## Redis

**Key Terms:**
- `string` — most basic type; can store text, numbers, serialized JSON; GET/SET
- `list` — ordered sequence; LPUSH/RPUSH/LPOP/RPOP; use as queue or stack
- `set` — unordered unique elements; SADD/SMEMBERS/SISMEMBER; use for tags, dedup
- `sorted set (zset)` — unique members with float scores; ZADD/ZRANGE/ZRANK; use for leaderboard
- `hash` — map of field→value; HSET/HGET/HGETALL; use for objects
- `TTL / EXPIRE` — set expiry on any key; EXPIRE key seconds; PEXPIRE key milliseconds
- `PERSIST` — remove TTL from key
- `pub/sub` — PUBLISH/SUBSCRIBE; fire-and-forget messaging (no persistence)
- `streams` — persistent log; consumer groups; like Kafka lite (XADD/XREAD)
- `atomic operations` — INCR, DECR, SETNX; single-operation without race condition
- `SETNX (SET if Not eXists)` — distributed lock primitive
- `Lua scripting` — run multiple operations atomically on server
- `pipeline` — batch multiple commands into single network round-trip
- `RDB snapshot` — periodic dump to disk for persistence
- `AOF (Append Only File)` — log every write command; more durable than RDB
- `eviction policies` — allkeys-lru, volatile-lru, noeviction, etc.
- `Redis Cluster` — automatic sharding across multiple Redis nodes
- `Sentinel` — monitors Redis for failover; high availability without clustering

**Must understand:**
- [ ] Why Redis is single-threaded yet fast (no lock overhead, in-memory, epoll)
- [ ] When to use sorted set (leaderboard, rate limiter sliding window)
- [ ] How to implement distributed lock with Redis (SET NX PX + Lua for unlock)
- [ ] RDB vs AOF persistence — which to use for what durability requirement
- [ ] Why pub/sub in Redis is not reliable (no persistence; messages lost if subscriber down)

---

## MongoDB

**Key Terms:**
- `document` — JSON-like object (BSON); stored in collection
- `collection` — group of documents; like a table in SQL
- `_id` — default primary key; ObjectId (12-byte auto-generated)
- `ObjectId` — contains timestamp, machine ID, process ID, counter; sortable by time
- `BSON` — Binary JSON; supports more types than JSON (Date, Binary, ObjectId)
- `schema-less` — each document in collection can have different fields
- `schema validation` — optional JSON Schema validation on insert/update
- `embedded document` — nest document inside document; denormalized
- `reference` — store ID of related document; normalized; requires second query
- `aggregation pipeline` — sequence of stages: $match, $group, $sort, $project, $lookup
- `$lookup` — left join between collections in aggregation
- `index` — same concept as SQL; B-tree; compound index; text index; geospatial
- `replica set` — primary + secondaries; automatic failover; read from secondary OK
- `sharding` — horizontal scale; shard key determines distribution
- `change streams` — listen to real-time changes in collection (like triggers / CDC)
- `transactions` — multi-document ACID transactions supported since MongoDB 4.0

**Must understand:**
- [ ] Embed vs reference — when to embed (data used together, 1:few) vs reference (large, many-to-many)
- [ ] Aggregation pipeline stages — $match, $group, $sort, $project
- [ ] Why MongoDB is AP in CAP — eventual consistency with replica sets
- [ ] MongoDB atlas vs self-hosted

---

## Kafka Basics

**Key Terms:**
- `topic` — named stream of records; partitioned
- `partition` — ordered, immutable log of records; basis for parallelism
- `offset` — position of a record in a partition; consumer tracks its offset
- `producer` — publishes records to topics
- `consumer` — reads records from partitions
- `consumer group` — multiple consumers sharing partitions of a topic; each partition → one consumer in group
- `broker` — Kafka server; cluster has multiple brokers
- `replication factor` — how many brokers keep copies of each partition
- `leader` — primary broker for a partition; handles reads/writes
- `follower` — replica of partition; takes over if leader fails
- `retention` — records kept for N days regardless of consumption (replay)
- `log compaction` — keep only latest record per key (for state tables)
- `at-least-once` — Kafka default; consumer may re-process on failure
- `exactly-once semantics (EOS)` — idempotent producer + transactional API
- `schema registry` — stores Avro/Protobuf schemas for topics; enforces compatibility
- `Kafka Streams` — library for stream processing on Kafka data
- `consumer lag` — how far behind consumers are from latest offset; key monitoring metric

**Must understand:**
- [ ] Why Kafka is durable (log on disk, replication) unlike Redis pub/sub
- [ ] Consumer group — why partitions and consumers in group must align for parallelism
- [ ] Offset management — auto-commit vs manual commit and their trade-offs
- [ ] When to use Kafka vs RabbitMQ vs SQS

---

## Migrations

**Key Terms:**
- `migration` — script that changes DB schema (add/remove table/column/index)
- `up migration` — apply change
- `down migration` — rollback change
- `versioned migrations` — each migration has a version number; applied in order
- `idempotent migration` — safe to run multiple times (IF NOT EXISTS, etc.)
- `golang-migrate` — Go tool for managing SQL migrations
- `Flyway` — Java tool; reads versioned SQL scripts
- `Liquibase` — Java tool; XML/YAML/SQL changesets
- `zero-downtime migration` — schema change without taking app offline
- `expand-contract` — safe pattern: add new col (expand) → update app → remove old col (contract)
- `online schema change` — tools like gh-ost or pt-online-schema-change for large tables
- `locking` — ALTER TABLE on large table locks it; blocks reads/writes for minutes
- `blue-green deployment` — run two DB versions simultaneously during migration

**Must understand:**
- [ ] Why ALTER TABLE is dangerous on large tables (lock, duration)
- [ ] Expand-contract pattern — how to add a column without downtime
- [ ] Migration ordering — why order matters; why down migrations are important
- [ ] Never delete a column in the same release as removing code that uses it
