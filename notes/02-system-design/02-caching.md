# Caching

---

## What is a cache and why does it exist?

Database queries are slow — disk reads, query planning, network round trips.
Most data is read far more than it is written.

Cache = a fast in-memory layer that stores results of expensive operations
so you do not repeat them on every request.

```
WITHOUT cache:
User → Server → Database (slow, every request)

WITH cache:
User → Server → Cache (fast, most requests — RAM)
                  ↓ miss
               Database (slow, only when not in cache)
```

Redis is the most common cache used in backend systems.

---

## The staleness problem

When cached data changes in the database, the cache now holds stale (old) data.
Two main approaches to fix this:

---

## Approach 1 — TTL (Time To Live)

Set an expiry time on every cached entry. After it expires, the next read goes
to the database and repopulates the cache.

```
Cache: { "user:42": { name: "Sam" }, expires_in: 60s }

After 60 seconds → entry evicted
Next read → cache miss → DB read → cache repopulated with fresh data
```

**Pros:** simple, automatic, no extra code needed
**Cons:** stale window — for up to TTL seconds, users see old data

Use when: slight staleness is acceptable (user profiles, product listings, config)

---

## Approach 2 — Cache Invalidation

When data changes in the database, immediately delete or update the cache entry.

```
User updates their name:
1. Write new name to database
2. Delete "user:42" from cache

Next read → cache miss → DB read → repopulates with fresh data
```

**Pros:** always fresh, no stale window
**Cons:** every write path must remember to invalidate — easy to miss one place and create a bug

---

## Three write strategies

**Write-through** — write to cache AND database simultaneously:

```
Write → Cache updated + Database updated (together)
```

Pros: cache always fresh
Cons: every write is slightly slower (two writes)

**Write-back (write-behind)** — write to cache first, database later async:

```
Write → Cache updated → return to user
              ↓ async later
         Database updated
```

Pros: very fast writes
Cons: data loss risk if cache crashes before DB write completes

**Cache-aside (lazy loading)** — most common in practice:

```
READ:
  1. Check cache
  2. Hit → return cached value
  3. Miss → read from DB → store in cache → return value

WRITE:
  1. Write to DB
  2. Delete from cache (invalidate)
  3. Next read repopulates from fresh DB data
```

Pros: only caches what is actually read, simple
Cons: first read after invalidation is always a cache miss (slightly slow)

---

## Eviction policies — what to remove when cache is full

Cache has limited RAM. When full, must evict something to make room.

**LRU — Least Recently Used** (most common)
Evict the item not accessed for the longest time.
Logic: if you haven't used it recently, you probably won't soon.

```
[A: 1min ago] [B: 5min ago] [C: 10min ago]
Cache full → evict C
```

**LFU — Least Frequently Used**
Evict the item accessed the fewest total times.

```
[A: 100 times] [B: 3 times] [C: 50 times]
Cache full → evict B
```

**TTL**
Expire items after a set time regardless of access pattern.

LRU is the most commonly used and most commonly asked about in interviews.

---

## Where to place the cache

**Application-level cache** — each server has its own in-memory cache:

```
Server 1: cache { user:42 → Sam }
Server 2: cache { }   ← different server, empty cache
```

Problem: same as session problem — each server has its own cache, inconsistent.

**Distributed cache (Redis)** — one shared cache all servers use:

```
Server 1 ──► Redis ◄── Server 2
              │
         { user:42 → Sam }
```

All servers share the same cache. Consistent. Standard production approach.

---

## Cache hit rate

```
hit rate = cache hits / total requests

hit rate 95% → only 5% of requests hit the database
hit rate 50% → half of requests bypass cache → cache not very useful
```

Target hit rate in production: typically 90%+. Below 80% means your caching
strategy or TTL needs tuning.

---

## Interview answers

**"How do you keep cache consistent with the database?"**
> "Two main approaches. TTL sets an expiry on every entry — the cache auto-expires
> and repopulates from the database after the TTL, accepting a staleness window.
> Cache invalidation immediately deletes the entry when source data changes —
> eliminates staleness but requires every write path to reliably invalidate.
> The most common production pattern is cache-aside with TTL as a safety net:
> reads check cache, misses go to DB and repopulate the cache, writes go to DB
> and invalidate the cache, TTL handles any missed invalidations."

**"What eviction policy would you use and why?"**
> "LRU for most cases — it keeps recently accessed data in cache which aligns with
> how access patterns work in practice. Hot items stay cached, cold items get evicted.
> LFU is better when access frequency matters more than recency — for example a
> rarely accessed but always-popular config value. TTL is complementary to both,
> not a replacement."

**"What is a cache miss and why does it matter?"**
> "A cache miss is when the requested data is not in the cache and must be fetched
> from the database. Misses are slower than hits and create load on the database.
> A high miss rate means the cache is not effective. At very high traffic, a sudden
> spike in cache misses — called a cache stampede — can overwhelm the database when
> many requests miss simultaneously and all hit the DB at once."
