# Database Sharding

---

## The Problem

Read replicas solve read scaling. But ALL writes still go to one primary database. With 500 million rows, the primary is dying — writes are slow, indexes don't fit in memory, backups take hours.

## What is Sharding?

Split data across multiple databases. Each database (shard) holds a portion of the data.

```
Before: [One DB — 500M users]

After:
  [Shard A — 125M users]
  [Shard B — 125M users]
  [Shard C — 125M users]
  [Shard D — 125M users]
```

---

## Sharding Strategies

### Range-Based Sharding

```
Shard A: user_id 1 — 125,000,000
Shard B: user_id 125,000,001 — 250,000,000
...
```

- ✅ Simple, range queries are easy
- ❌ Uneven load — new users all go to last shard (hot spot)

### Hash-Based Sharding

```
shard = hash(user_id) % 4
```

- ✅ Even distribution, no hot spots
- ❌ Range queries impossible (data is scattered)
- ❌ Adding a shard remaps almost ALL data

### Consistent Hashing

Place shards and keys on a hash ring (circle 0-360). Each key goes to the nearest shard clockwise.

```
Adding a new shard: only data between new shard and its neighbor moves.
Regular hashing: add 1 shard → move ~80% of data
Consistent hashing: add 1 shard → move ~20% of data (1/n)
```

Used by: Cassandra, DynamoDB, distributed caches.

---

## Sharding Problems

### Scatter-Gather Queries
Cross-shard queries must hit ALL shards, collect results, merge in application layer. Slow and complex.

### Cross-Shard JOINs
If user is on Shard A but their orders are on Shard C, you can't do a simple JOIN.

### Sharding is a Last Resort
Try these first:
1. Optimize queries and indexes
2. Add read replicas
3. Add caching
4. Shard only when nothing else works

---

## Interview Answer

**"When would you shard a database?"**
> "Sharding is a last resort for scaling writes. I'd first try query optimization, read replicas, and caching. If the primary still can't handle the write load, I'd shard using consistent hashing to distribute data evenly and minimize data movement when adding shards. The main trade-off is that cross-shard queries become expensive — you need scatter-gather. So choosing the right shard key is critical — you want most queries to hit only one shard."
