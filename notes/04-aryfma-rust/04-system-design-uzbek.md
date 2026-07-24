# System Design — Framework + Misollar (O'zbekcha) + Senior Javoblar

---

# Universal Framework (har qanday "Design X" uchun)

Doim shu 6 qadamdan yuring. Bu tuzilma o'zi senior signal beradi.

## 1. Requirements (talablarni aniqlashtirish) — 2-3 daqiqa
- **Functional**: tizim aynan nima qiladi? (asosiy use-case'lar)
- **Non-functional**: qancha user? RPS? latency? consistency vs availability?
- **Scope**: nima kiradi, nima kirmaydi? (savol berib chegarani belgilang)

> ⚠️ Hech qachon darhol chizishga tushmang. Avval **savol bering**. Bu eng katta senior signal.

## 2. Estimation (taxminiy hisob) — 2 daqiqa
- DAU (daily active users), RPS = users × actions / 86400
- Storage = yozuvlar soni × yozuv hajmi
- Read/write nisbati (masalan 100:1 → keshlash muhim)

## 3. API Design
- Asosiy endpoint'lar: `POST /url`, `GET /{id}` ...
- Request/response ko'rinishi

## 4. High-Level Design (asosiy chizma)
- Client → Load Balancer → App servers → Cache → DB
- Qaysi komponent nima qiladi

## 5. Data Model
- Jadvallar / sxema, qaysi DB (SQL vs NoSQL) va **nega**
- Indexlar, sharding kaliti

## 6. Deep Dive + Tradeoffs (bottleneck'larни hal qilish)
- Scaling: replication, sharding, caching, queue
- CAP: consistency yoki availability?
- Bottleneck qayerda? Qanday hal qilamiz?

---

# Kalit tushunchalar (interview'da kerak bo'ladi)

- **Load Balancer** — yukni server'larga taqsimlaydi (round-robin, least-connections).
- **Horizontal vs Vertical scaling** — ko'p server qo'shish vs bitta server'ni kuchaytirish.
- **Caching** — Redis/Memcached. Cache-aside, write-through, write-back. TTL, eviction (LRU).
- **CDN** — statik kontentni foydalanuvchiga yaqin joyda.
- **Replication** — master-replica; read'larni replica'ga yo'naltirish.
- **Sharding** — data'ni bo'laklarga bo'lish (kalitga qarab). Hot shard muammosi.
- **Message Queue** — Kafka/RabbitMQ. Asinxron ish, decoupling, buffer.
- **CAP teorema** — Consistency, Availability, Partition tolerance — network bo'linganда 2 tasini tanlaysan.
- **Consistent hashing** — server qo'shilganда/o'chganда kam data ko'chishi.
- **Idempotency** — bir so'rovni ikki marta yuborsa ham natija bir xil (retry uchun muhim).

---

# MISOL 1: URL Shortener (bit.ly)

## 1. Requirements
- Functional: uzun URL → qisqa kod; qisqa kod → redirect (301/302).
- Non-functional: read-heavy (100:1 read/write), past latency redirect, yuqori availability.
- Scale: aytaylik 100M yangi URL/oy, 10B redirect/oy.

## 2. Estimation
- Yozish: 100M/oy ≈ 40 write/sek
- O'qish: 10B/oy ≈ 4000 read/sek → **keshlash juda muhim**
- Storage: 100M × 12oy × 5yil × ~500 bayt ≈ ~3 TB

## 3. API
```
POST /shorten   {long_url} → {short_url}
GET  /{code}    → 301 redirect
```

## 4. High-Level
```
Client → LB → App servers → Cache (Redis) → DB
                          ↘ (miss bo'lsa DB'dan o'qib keshga yoz)
```

## 5. Data Model
```
urls(code PK, long_url, created_at, user_id)
```
- `code` — 7 belgi, base62 (a-zA-Z0-9) = 62^7 ≈ 3.5 trillion kombinatsiya.
- Kod generatsiya: (a) counter + base62 encode, yoki (b) hash(url) + collision tekshirish.

## 6. Deep Dive
- **Redirect tez bo'lishi kerak** → code→url ni Redis'da keshlash (read-heavy).
- **Kod to'qnashuvi (collision)** → unique index + retry, yoki counter-based (to'qnashuv yo'q).
- **Scale** → DB read replica + Redis. Code random bo'lgani uchun sharding oson.
- **301 vs 302** → 301 (permanent) brauzer keshlaydi, DB yuki kamayadi, lekin analytics yo'qoladi. 302 har safar server'ga keladi (analytics uchun). Tradeoff.

## 🎯 Interview javob (EN, senior):
> **"Design a URL shortener"** — I'd clarify scale first: this is extremely read-heavy, maybe 100:1, so the redirect path must be cache-first. The write path stores a base62 code — I'd use a counter-based generator encoded to base62 to avoid collisions entirely, rather than hashing which needs collision handling. Reads go Redis-first, falling back to the DB on a miss and repopulating the cache. The data is naturally shardable by code. The interesting tradeoff is 301 vs 302 redirects: 301 is cacheable by browsers and cuts load but kills click analytics, so the choice depends on whether tracking matters.

---

# MISOL 2: Rate Limiter

## 1. Requirements
- Bir user/IP sekundiga N ta so'rovdan ko'p yubormasin.
- Distributed (ko'p server) bo'lishi kerak.

## 2. Algoritmlar
- **Token bucket** — bucket'da token'lar, har so'rov 1 token oladi, token belgilangan tezlikda to'ldiriladi. Burst'ga ruxsat beradi. ✅ Eng ko'p ishlatiladi.
- **Leaky bucket** — so'rovlar navbatda, doimiy tezlikda chiqadi. Silliq.
- **Fixed window** — har daqiqada N. Oddiy, lekin chegara chetida 2x burst muammosi.
- **Sliding window log/counter** — aniqroq, chegara muammosini hal qiladi.

## 3. Distributed implementatsiya
- **Redis** — counter yoki token'ni markazda saqlash. `INCR` + `EXPIRE`, yoki Lua script atomiklik uchun.
- Har app server Redis'dan tekshiradi → global chegara.

## 🎯 Interview javob (EN, senior):
> **"Design a rate limiter"** — I'd use a token bucket because it allows short bursts while enforcing an average rate, which fits most APIs. For a distributed setup the state has to be shared, so I'd keep the bucket in Redis and update it with an atomic Lua script to avoid race conditions across app servers. Fixed-window counters are simpler but allow a 2x burst at the window boundary, so I'd prefer a sliding window if precision matters. I'd also decide the failure mode: fail-open (allow if Redis is down) for availability, or fail-closed for strict protection.

---

# 🎯 Umumiy System Design savollari + javoblar

**Q1: SQL vs NoSQL — qachon qaysi?**
> **Senior answer (EN):** I default to SQL like Postgres when I need strong consistency, transactions, and relational queries with joins — most business apps. I reach for NoSQL when I need to scale writes horizontally beyond one node, have a flexible or denormalized schema, or need a specific access pattern — key-value for caching/sessions, document for varied shapes, wide-column for huge time-series. The decision is driven by access patterns and consistency needs, not hype.

**Q2: How do you scale reads?**
> **Senior answer (EN):** First caching — put hot data in Redis with a sensible TTL and eviction policy, since most systems are read-heavy. Then read replicas — route reads to replicas and writes to the primary, accepting some replication lag. If a single dataset is still too big, I shard. I always ask whether the workload can tolerate eventual consistency, because that unlocks cheaper scaling.

**Q3: What is CAP theorem?**
> **Senior answer (EN):** During a network partition you can guarantee either consistency or availability, not both. A CP system rejects requests it can't make consistent; an AP system stays available but may serve stale data. In practice modern systems tune this per-operation — Postgres is CP-leaning, Cassandra is AP-leaning with tunable consistency. The real question I ask is: for this specific operation, is stale data acceptable?

**Q4: How do you ensure a request isn't processed twice (idempotency)?**
> **Senior answer (EN):** I give each request an idempotency key from the client, store it with the result, and on a retry return the stored result instead of reprocessing. For payments or order creation this prevents double-charging. Combined with at-least-once delivery from a queue, idempotency is what makes retries safe.
