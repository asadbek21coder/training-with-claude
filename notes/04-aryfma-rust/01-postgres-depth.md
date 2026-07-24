# Postgres Deep Q&A — Aryfma Interview Prep

Senior-level answers. Each: **the question → the crisp answer (EN) → RU version → the "senior kicker"** that shows depth.

---

## 1. What is MVCC and how does Postgres implement it?

**EN:** MVCC = Multi-Version Concurrency Control. Instead of locking rows for reads, Postgres keeps *multiple versions* of a row. Each transaction sees a **snapshot** of the DB as of when it started. Readers never block writers, writers never block readers.

How: every row (tuple) has hidden system columns `xmin` (transaction that created it) and `xmax` (transaction that deleted/updated it). An `UPDATE` doesn't overwrite — it marks the old tuple dead (`xmax`) and inserts a new tuple. Each transaction checks visibility rules against its snapshot.

**RU:** MVCC = управление конкурентным доступом через версии. Postgres хранит **несколько версий** строки. Каждая транзакция видит **снимок (snapshot)** базы на момент своего старта. Читатели не блокируют писателей, писатели не блокируют читателей. У каждой строки есть скрытые поля `xmin` (кто создал) и `xmax` (кто удалил). UPDATE не перезаписывает — старая версия помечается мёртвой, вставляется новая.

**Senior kicker:** "This is why Postgres needs **VACUUM** — dead tuples accumulate and must be cleaned up, otherwise you get **table bloat** and **transaction ID wraparound** risk."

---

## 2. What does VACUUM do? Why is it needed?

**EN:** Dead tuples (from updates/deletes) stay on disk under MVCC. `VACUUM` reclaims that space and marks it reusable. `VACUUM` also updates the **visibility map** and prevents **XID wraparound** (freezing old tuples). `autovacuum` does this automatically. `VACUUM FULL` rewrites the whole table (locks it, reclaims disk to OS) — rarely used in prod.

**RU:** Мёртвые строки остаются на диске из-за MVCC. `VACUUM` освобождает это место для повторного использования, обновляет карту видимости и предотвращает переполнение ID транзакций. `autovacuum` делает это автоматически. `VACUUM FULL` переписывает всю таблицу и блокирует её — редко в проде.

**Senior kicker:** "Bloat kills performance because indexes and scans touch dead space. Monitoring `n_dead_tup` and tuning autovacuum thresholds is a real ops task."

---

## 3. B-tree index — how does it work, when does Postgres NOT use it?

**EN:** B-tree = balanced tree, keeps keys sorted, O(log n) lookup. Great for `=`, `<`, `>`, `BETWEEN`, `ORDER BY`, and prefix `LIKE 'abc%'`.

Postgres **won't** use it when:
- The query would return a large fraction of the table (seq scan is cheaper).
- You wrap the column in a function: `WHERE lower(email) = ...` (unless you built a functional index).
- Leading wildcard `LIKE '%abc'`.
- Type mismatch forcing a cast.
- Stats are stale (run `ANALYZE`).

**RU:** B-tree = сбалансированное дерево, ключи отсортированы, поиск O(log n). Хорош для `=`, `<`, `>`, `BETWEEN`, `ORDER BY`, префикса `LIKE 'abc%'`. Postgres НЕ использует индекс, когда возвращается большая часть таблицы, когда столбец обёрнут в функцию, при `LIKE '%abc'`, при несовпадении типов, или при устаревшей статистике.

**Senior kicker:** "The planner is **cost-based** — it uses `pg_statistic` estimates. An index isn't always faster; a seq scan on a small or mostly-matching table wins. `EXPLAIN ANALYZE` shows the real choice."

---

## 4. Index types beyond B-tree

- **Hash** — only `=`, now WAL-logged/crash-safe (since PG10). Rarely worth it over B-tree.
- **GIN** — for arrays, JSONB, full-text search. "Which rows contain this element."
- **GiST** — geometric / range / nearest-neighbor.
- **BRIN** — huge tables with natural physical ordering (e.g. append-only time-series). Tiny index, stores min/max per block range.
- **Partial index** — `CREATE INDEX ... WHERE is_active` — index only rows you query.
- **Covering / INCLUDE** — `INCLUDE (col)` to enable **index-only scans**.

**RU кратко:** GIN — для JSONB/массивов/полнотекста; BRIN — для огромных упорядоченных таблиц; частичный индекс — только нужные строки; INCLUDE — для index-only scan.

---

## 5. Transaction isolation levels

| Level | Prevents |
|-------|----------|
| Read Committed (PG default) | dirty reads |
| Repeatable Read | + non-repeatable reads, phantom reads (PG uses snapshot) |
| Serializable | + serialization anomalies (SSI) |

**EN:** Postgres default is **Read Committed** — each statement sees rows committed before *it* started. **Repeatable Read** in PG actually prevents phantoms too (snapshot-based). **Serializable** uses SSI (Serializable Snapshot Isolation) and may abort transactions with a serialization error you must retry.

**RU:** По умолчанию **Read Committed** — каждый оператор видит данные, закоммиченные до его старта. **Repeatable Read** в PG также предотвращает фантомы. **Serializable** может откатить транзакцию с ошибкой сериализации — нужно повторять (retry).

**Senior kicker:** "At Serializable you must write **retry logic** — the DB tells you 'couldn't serialize, try again.' People forget this and get silent bugs."

---

## 6. How would you debug a slow query?

1. `EXPLAIN (ANALYZE, BUFFERS)` — see actual plan, row estimates vs reality, buffer hits.
2. Look for **Seq Scan** on big tables, bad **row estimate** mismatches (→ stale stats, run `ANALYZE`).
3. Check for missing index, or index not used (functions, casts).
4. Check `work_mem` — sorts/hashes spilling to disk show as "external merge Disk".
5. Look at joins — nested loop on big sets is a red flag.
6. `pg_stat_statements` to find the worst queries system-wide.

**RU:** `EXPLAIN (ANALYZE, BUFFERS)` → искать Seq Scan на больших таблицах, расхождение оценок строк (устаревшая статистика → `ANALYZE`), отсутствие индекса, нехватку `work_mem` (сортировка на диск), плохие join'ы. `pg_stat_statements` — найти худшие запросы.

**Senior kicker:** "Estimated vs actual row count mismatch is the #1 tell — it means the planner is working with wrong statistics and every downstream decision is poisoned."

---

## 7. Connection pooling — why and how?

**EN:** Each Postgres connection = a separate OS process (~a few MB + backend overhead). Thousands of connections crush the server. Apps use a **pooler** — **PgBouncer** (transaction pooling is most common) — to multiplex many clients onto few real connections. In Go, `database/sql` also has its own pool (`SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`).

**RU:** Каждое соединение Postgres — отдельный процесс ОС, дорогой. Тысячи соединений убивают сервер. Используют **пул** — **PgBouncer** (transaction pooling). В Go `database/sql` тоже имеет пул: `SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`.

**Senior kicker:** "With PgBouncer in **transaction mode** you can't use session-level features (prepared statements across calls, `SET`, advisory locks held across statements) — a classic footgun."

---

## 8. N+1 query problem

**EN:** You fetch a list (1 query), then loop and fetch each item's relation (N queries) = N+1 round trips. Fix: a single `JOIN`, or `WHERE id = ANY($1)` with a batch of IDs, or `IN (...)`. Round trips, not the DB's CPU, are usually the killer.

**RU:** Берёшь список (1 запрос), потом в цикле грузишь связанные данные (N запросов) = N+1. Решение: один `JOIN` или `WHERE id = ANY($1)` с массивом ID. Убивают именно round trips (обращения по сети), а не CPU базы.

---

## 9. Keyset (cursor) pagination vs OFFSET

**EN:** `OFFSET 100000 LIMIT 20` still scans and discards 100k rows — slow on deep pages. **Keyset pagination:** `WHERE (created_at, id) < ($last_ts, $last_id) ORDER BY created_at DESC, id DESC LIMIT 20`. Uses the index, constant time regardless of depth.

**RU:** `OFFSET 100000` всё равно сканирует и выбрасывает 100k строк — медленно на глубоких страницах. **Keyset-пагинация** использует индекс и работает за постоянное время на любой глубине.

**Senior kicker:** "OFFSET pagination also has correctness bugs — rows shift if data is inserted between page loads. Keyset is stable."

---

## 10. Locks — row locks, deadlocks

**EN:** `SELECT ... FOR UPDATE` locks rows to prevent concurrent modification (e.g. inventory decrement). Deadlock = A locks row1 waits row2, B locks row2 waits row1. Postgres detects it and kills one transaction. Prevention: **always acquire locks in the same order**. `SELECT ... FOR UPDATE SKIP LOCKED` is the pattern for a **job queue** in Postgres.

**RU:** `SELECT ... FOR UPDATE` блокирует строки. Deadlock — A ждёт B, B ждёт A; Postgres обнаруживает и убивает одну транзакцию. Профилактика: всегда брать блокировки в одном порядке. `FOR UPDATE SKIP LOCKED` — паттерн для очереди задач в Postgres.

---

## Rapid-fire cheat answers
- **CTE** = `WITH x AS (...)`. Since PG12, not always a materialization barrier (can be inlined).
- **Materialized view** = cached query result on disk; `REFRESH` to update.
- **JSONB vs JSON** = JSONB is binary, indexable (GIN), reordered keys, no dup keys. Prefer JSONB.
- **UUID vs bigserial PK** = UUID good for distributed/no-coordination, but random UUIDv4 hurts B-tree locality (bloat) → consider UUIDv7 (time-ordered).
- **`ON CONFLICT DO UPDATE`** = upsert.
- **Partitioning** = split huge table by range/list/hash; prune partitions in queries.
