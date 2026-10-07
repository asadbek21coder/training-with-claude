# Postgres — Chuqur Tushuntirish (O'zbekcha) + Interview Javoblari

Format: mavzu tushuntirish → kichik savol-javob → interview savollari + senior javoblar.

---

# 1. MVCC — Multi-Version Concurrency Control

## Tushuntirish
Postgres qatorni (row) **hech qachon joyida o'zgartirmaydi**. Har bir qator diskда bir necha **versiya** bo'lib yashaydi.

Har bir qatorda 2 ta yashirin ustun bor:
- `xmin` — bu versiyani **yaratgan** transaction ID (XID)
- `xmax` — bu versiyani **o'chirgan/almashtirgan** XID (agar tirik bo'lsa 0)

UPDATE ishlaganda:
```
Eski qator:  (xmin=100, xmax=105, "Alice")   ← 105 uni o'ldirdi
Yangi qator: (xmin=105, xmax=0,   "Alicia")  ← yangi versiya
```

Har bir transaction o'zining **snapshot** (suratga olingan holat) ига qarab qaysi versiyani ko'rishini hal qiladi. Shuning uchun:
- **Readers writers ni bloklamaydi** (o'quvchi yozuvchini kutmaydi)
- **Writers readers ni bloklamaydi**

## Nega bu ham yaxshi, ham qimmat?
- ✅ Juda yuqori concurrency (bir vaqtda ko'p ulanish)
- ❌ O'lgan versiyalar to'planadi → **VACUUM** kerak
- ❌ Har bir index har bir versiyaga ishora qiladi → index ham shishadi (bloat)
- ❌ **XID wraparound**: XID 32-bit (~4 mlrd). Eski qatorlar "muzlatilmasa" (freeze), sanoq aylanib ketib, eski data kelajakdan kelganday ko'rinadi. VACUUM eski qatorlarni freeze qiladi.

## VACUUM nima qiladi?
O'lgan qatorlar diskда qoladi. `VACUUM`:
1. O'sha joyni **qayta ishlatishga** bo'shatadi
2. **Visibility map** ni yangilaydi (index-only scan uchun kerak)
3. Eski qatorlarni **freeze** qiladi (XID wraparound oldini oladi)

`autovacuum` — buni avtomatik qiladi. `VACUUM FULL` — butun jadvalni qayta yozadi, **jadvalni bloklaydi**, diskни OS ga qaytaradi (prodda kam ishlatiladi).

## 🔸 Kichik savol 1
**S:** UPDATE va DELETE'dan keyin eski qatorlar qayerga ketadi?
**J:** Diskда "dead tuple" (o'lik qator) bo'lib qoladi, keyin VACUUM ularni bo'shatadi.

## 🔸 Kichik savol 2
**S:** Nega jadval qator soni o'zgarmasa ham katta va sekin bo'lib qoladi?
**J:** Ko'p UPDATE/DELETE bo'lganda dead tuple'lar autovacuum tozalashdan tez to'planadi → bu **bloat**. `pg_stat_user_tables` dagi `n_dead_tup` ni tekshiraman.

## 🔸 Kichik savol 3
**S:** MVCC readerlarni bloklamasligining sababi nima?
**J:** Reader eski versiyani (o'z snapshot'iga mos) o'qiydi, writer esa yangi versiya yaratadi — ikkalasi turli versiyaga tegadi, konflikt yo'q.

---

## 🎯 INTERVIEW SAVOLLARI — MVCC

**Q1: Explain MVCC in Postgres.**
> **Senior answer (EN):** MVCC means Postgres keeps multiple versions of each row instead of locking on reads. Each row has hidden `xmin`/`xmax` system columns marking which transaction created and ended it. Every transaction reads against a snapshot taken at its start, so readers never block writers and writers never block readers. The cost is that dead row versions accumulate, which is why `VACUUM` (and autovacuum) must reclaim them — otherwise you get table and index bloat, and risk XID wraparound.

**Q2: Why does Postgres need VACUUM?**
> **Senior answer (EN):** Because of MVCC, updates and deletes leave dead tuples on disk — they aren't overwritten. VACUUM reclaims that space for reuse, updates the visibility map to enable index-only scans, and freezes old tuples to prevent transaction-ID wraparound. In production the real work is tuning autovacuum so it keeps up with high-churn tables; I monitor `n_dead_tup` and can make autovacuum more aggressive per-table.

**Q3: A table's row count is stable but it keeps growing and slowing down. Why?**
> **Senior answer (EN):** That's bloat from MVCC. High update/delete churn produces dead tuples faster than autovacuum reclaims them, so the physical file and its indexes grow. I'd confirm via `pg_stat_user_tables.n_dead_tup` and `pgstattuple`, tune autovacuum thresholds for that table, and if severe, run `pg_repack` (online) or `VACUUM FULL` in a maintenance window.

---

# 2. Isolation Levels — Izolyatsiya Darajalari

## Tushuntirish
Transaction'lar bir-birini qanchalik "ko'rishi" — shuni izolyatsiya darajasi belgilaydi.

### Read Committed (Postgres default)
- Har bir **statement** (har bir SQL buyruq) **yangi snapshot** oladi.
- Shuning uchun bitta transaction ichida ikkita bir xil SELECT **turli natija** berishi mumkin (orasida boshqa transaction commit qilsa).

### Repeatable Read
- Snapshot **bir marta** — birinchi statement'da olinadi va butun transaction davomida **muzlatiladi**.
- Non-repeatable read va phantom read bo'lmaydi.
- LEKIN: boshqasi o'zgartirgan qatorni update qilmoqchi bo'lsang → `could not serialize access` xatosi → **transaction'ni qaytadan bajarish (retry) kerak**.

### Serializable (SSI — Serializable Snapshot Isolation)
- Go'yo transaction'lar **navbat bilan** (birma-bir) ishlaganday natija beradi.
- Postgres read/write bog'liqliklarni kuzatadi; xavfli sikl paydo bo'lsa → bittasini bekor qiladi.
- **Retry logikasi majburiy.**

## ⚠️ Eng muhim tuzoq: Lost Update
Read Committed **lost update** ni oldini OLMAYDI:
```
Txn A: SELECT balance → 100
Txn B: SELECT balance → 100
Txn A: UPDATE SET balance = 150   (100+50)
Txn B: UPDATE SET balance = 150   (100+50, A ni yozganini bilmaydi)
→ Yakuniy: 150. To'g'risi 200 bo'lishi kerak edi. 50 yo'qoldi!
```

**Yechimlar:**
1. `SELECT ... FOR UPDATE` — qatorni bloklab qo'yish
2. `UPDATE accounts SET balance = balance + 50` — atomik, DB ichida hisoblash
3. Repeatable Read / Serializable + **retry**

## 🔸 Kichik savol 1
**S:** Postgres'da default isolation level qaysi?
**J:** Read Committed.

## 🔸 Kichik savol 2
**S:** Repeatable Read'da nima uchun ba'zan retry kerak bo'ladi?
**J:** Ikki transaction bir qatorni o'zgartirmoqchi bo'lsa, Postgres bittasini `serialization failure` bilan to'xtatadi — o'shani qaytadan bajarish kerak.

## 🔸 Kichik savol 3
**S:** Bank hisobidan pul yechishda lost update'dan qanday himoyalanasan?
**J:** `SELECT ... FOR UPDATE` bilan qatorni bloklaymaman, yoki `SET balance = balance - amount` atomik update ishlatamaman.

---

## 🎯 INTERVIEW SAVOLLARI — Isolation

**Q1: What are the isolation levels in Postgres?**
> **Senior answer (EN):** Postgres supports Read Uncommitted (behaves like Read Committed), Read Committed (the default), Repeatable Read, and Serializable. Read Committed gives each statement a fresh snapshot. Repeatable Read takes one snapshot for the whole transaction and prevents non-repeatable and phantom reads. Serializable uses SSI to guarantee the result is equivalent to some serial order — but it can abort a transaction with a serialization error, so you need retry logic.

**Q2: Does Read Committed prevent lost updates?**
> **Senior answer (EN):** No. Two transactions can both read a value, both compute a new one, and the second write silently overwrites the first. To prevent it I either lock the row with `SELECT ... FOR UPDATE`, do the arithmetic atomically in SQL like `SET balance = balance + $1`, or move to Repeatable Read/Serializable with a retry loop on serialization failures.

**Q3: When would you use Serializable?**
> **Senior answer (EN):** When correctness across concurrent transactions matters more than throughput and the logic is hard to make safe with explicit locks — for example financial invariants spanning multiple rows. The tradeoff is you must handle SQLSTATE 40001 by retrying, and it has more overhead, so I use it surgically, not everywhere.

---

# 3. Query Planner + EXPLAIN

## Tushuntirish
Postgres planeri **cost-based** — qoidaga emas, **narxga** asoslanadi. Har bir mumkin bo'lgan planning uchun raqamli "narx" hisoblab, eng arzonini tanlaydi.

### Nimaga asoslanadi?
1. **Statistika** (`ANALYZE` yig'adi, `pg_statistic`da saqlanadi): qator soni, eng ko'p uchraydigan qiymatlar, qiymatlar taqsimoti (histogram), null ulushi, distinct qiymatlar soni.
2. **Narx konstantalari**: `seq_page_cost` (1.0), `random_page_cost` (4.0 — eski disk uchun; **SSD da ~1.1 qilish kerak**), `cpu_tuple_cost` va h.k.

### Scan turlari
- **Seq Scan** — butun jadvalni o'qish. Ko'p qator mos kelsa arzon.
- **Index Scan** — B-tree bo'ylab yurib, keyin har bir qatorni heap'dan olish (random I/O).
- **Index-Only Scan** — agar kerakli hamma ustun indexда bo'lsa VA visibility map "hammasi ko'rinadigan" desa — heap'ga tegmaydi (tez). `INCLUDE` shuning uchun kerak.
- **Bitmap Index Scan** — mos sahifalar bitmap'ini yasab, heap'ni fizik tartibda o'qiydi. O'rtacha selektivlik uchun.

### Join turlari
- **Nested Loop** — kichik tashqi + indexli ichki uchun yaxshi. Ikki katta indexsiz to'plam uchun yomon (O(n·m)).
- **Hash Join** — bir tomonga hash jadval qurib, ikkinchisi bilan tekshiradi. Katta equijoin uchun zo'r, `work_mem` kerak.
- **Merge Join** — ikkala tomon saralangan, "zip" qilib qo'shadi. Data allaqachon saralangan bo'lsa yaxshi.

### EXPLAIN ANALYZE o'qish
```
Seq Scan on orders (cost=0.00..18334.00 rows=1000000 width=32)
                   (actual time=0.02..95.3 rows=1000000 loops=1)
```
- `cost=boshlanish..jami` — planer bahosi (shartli birlik)
- `rows=` (baho) ↔ `actual ... rows=` (real) → **agar keskin farq qilsa, statistika noto'g'ri** — bu #1 diagnostika signali
- `Buffers: shared hit=... read=...` (BUFFERS bilan) → hit = kesh, read = disk

## 🔸 Kichik savol 1
**S:** Nega ba'zan index yaratsam ham Postgres uni ishlatmaydi?
**J:** Sabablari: (1) so'rov jadvalning katta qismini qaytaradi (seq scan arzon), (2) ustun funksiyaga o'ralgan `lower(email)`, (3) `LIKE '%abc'` boshida wildcard, (4) tип mos kelmaydi, (5) statistika eskirgan (`ANALYZE` kerak), (6) SSD da `random_page_cost` yuqori.

## 🔸 Kichik savol 2
**S:** EXPLAIN'da eng muhim narsa nima?
**J:** `rows` bahosi va real `rows` farqi. Katta farq = statistika noto'g'ri = planer noto'g'ri qaror qiladi.

## 🔸 Kichik savol 3
**S:** Index-Only Scan qachon ishlaydi?
**J:** So'rovga kerakli hamma ustun indexда bo'lsa va sahifa visibility map'da "all-visible" bo'lsa.

---

## 🎯 INTERVIEW SAVOLLARI — Planner

**Q1: How would you debug a slow query?**
> **Senior answer (EN):** I start with `EXPLAIN (ANALYZE, BUFFERS)` to see the real plan and actual timings. First I compare estimated vs actual row counts — a big mismatch means stale statistics, so I run `ANALYZE`. Then I look for sequential scans on large tables, indexes that aren't being used (functions or casts on the column), sorts spilling to disk when `work_mem` is too low, and bad join strategies like a nested loop over large sets. System-wide, `pg_stat_statements` tells me which queries hurt most.

**Q2: Why might Postgres ignore an index you created?**
> **Senior answer (EN):** Several reasons: the query returns a large fraction of the table so a seq scan is cheaper; the column is wrapped in a function or cast that the index doesn't match; a leading wildcard in LIKE; stale statistics; or, on SSD, a too-high `random_page_cost` making the planner overvalue sequential I/O. The planner is cost-based, so it only uses an index when it estimates that's cheaper.

**Q3: What's the difference between an Index Scan and an Index-Only Scan?**
> **Senior answer (EN):** An Index Scan walks the index to find matching entries, then fetches each row from the heap — that heap fetch is random I/O. An Index-Only Scan avoids the heap entirely when every column the query needs is in the index and the page is marked all-visible in the visibility map. That's why covering indexes with INCLUDE columns and a well-vacuumed table matter for read performance.

---

# 4. Tez javoblar (rapid-fire)

- **B-tree** — default index, saralangan, O(log n). `=`, `<`, `>`, `BETWEEN`, `ORDER BY`, `LIKE 'abc%'` uchun.
- **GIN** — JSONB, massiv, full-text search uchun.
- **BRIN** — juda katta, tabiiy tartibli jadvallar (time-series) uchun. Kichik index.
- **Partial index** — `CREATE INDEX ... WHERE is_active` — faqat kerakli qatorlar.
- **N+1 problem** — 1 ta list + N ta bog'liq so'rov. Yechim: `JOIN` yoki `WHERE id = ANY($1)`.
- **Keyset pagination** — `WHERE (created_at, id) < ($ts, $id)` — OFFSET'dan tez, chuqur sahifalarda barqaror.
- **Connection pooling** — har ulanish = alohida OS process, qimmat. **PgBouncer** (transaction mode). Go'da `SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`.
- **`FOR UPDATE SKIP LOCKED`** — Postgres'da job queue (navbat) patterni.
- **JSONB > JSON** — binar, indekslanadigan (GIN), kalitlar tartibsiz, dublikat yo'q.
- **Upsert** — `INSERT ... ON CONFLICT DO UPDATE`.
- **Deadlock** — A→B kutadi, B→A kutadi. Postgres aniqlab bittasini o'ldiradi. Oldini olish: **bloklarni doim bir xil tartibda olish**.
