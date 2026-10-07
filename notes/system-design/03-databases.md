# Databases — Replication, Sharding, SQL vs NoSQL

---

## The two problems at scale

**Problem 1 — database becomes a bottleneck**
Every read and write goes to one machine. High traffic → CPU/disk maxes out → everything slows.

**Problem 2 — single point of failure**
Database goes down → entire app goes down.

---

## SQL vs NoSQL

**SQL (relational)** — tables, rows, columns, strict schema, relations between tables.
Examples: PostgreSQL, MySQL.

```
users:                          orders:
┌────┬──────┬──────────────┐    ┌────┬─────────┬────────┬─────────┐
│ id │ name │ email        │    │ id │ user_id │ amount │ status  │
├────┼──────┼──────────────┤    ├────┼─────────┼────────┼─────────┤
│  1 │ Sam  │ sam@...      │    │  1 │    1    │   50   │  paid   │
│  2 │ Ali  │ ali@...      │    │  2 │    1    │   30   │ pending │
└────┴──────┴──────────────┘    └────┴─────────┴────────┴─────────┘
```

**NoSQL** — no fixed schema. Documents, key-value, wide-column, graph.
Examples: MongoDB (documents), Redis (key-value), Cassandra (wide-column).

```json
{
  "_id": "1",
  "name": "Sam",
  "orders": [
    { "amount": 50, "status": "paid" },
    { "amount": 30, "status": "pending" }
  ]
}
```

### When to use which

**Use SQL when:**
- Data has clear relationships (users → orders → products)
- You need ACID transactions (payments, bank transfers)
- Schema is stable
- You need complex queries with joins

**Use NoSQL when:**
- Data structure varies per record
- You need very high write throughput horizontally
- Data is naturally document-shaped (no joins needed)
- Schema changes frequently

**Honest default:** start with PostgreSQL. It handles more scale than people expect
and ACID guarantees prevent entire classes of bugs. Switch to NoSQL only when you
have a specific proven problem SQL cannot solve.

---

## Replication — solving the read bottleneck

Copy the database to multiple machines. Primary handles all writes.
Replicas handle reads.

```
         WRITE              READS
           │           ┌──────┤
           ▼           │      ▼
      ┌─────────┐      │  ┌──────────┐
      │ Primary │──────┼─►│ Replica 1│
      └─────────┘      │  └──────────┘
           │           │  ┌──────────┐
           └───────────┴─►│ Replica 2│
                          └──────────┘
```

**Pros:**
- Read traffic distributed across replicas — scales reads horizontally
- High availability — if primary fails, promote a replica to primary

**Cons:**
- Replication lag — replicas are slightly behind primary
  User writes then immediately reads → might get stale data from replica
- Only scales reads — all writes still go to one primary

---

## Sharding — solving the write bottleneck

Replication does not help writes. When write volume exceeds what one primary
can handle, split data across multiple databases (shards).

```
user_id 1–1000    → Shard 1
user_id 1001–2000 → Shard 2
user_id 2001–3000 → Shard 3
```

Each shard is an independent database handling its own reads and writes.

**Pros:** scales both reads and writes, no single bottleneck

**Cons:**
- Cross-shard queries are painful (must query multiple shards and merge)
- Joins across shards are very expensive or impossible
- Rebalancing is complex when shards become uneven

### Sharding strategies

**Range-based:**
```
user_id 1–1M     → Shard 1
user_id 1M–2M    → Shard 2
```
Simple but can be uneven if data is clustered in one range.

**Hash-based:**
```
hash(user_id) % 3 = 0 → Shard 1
hash(user_id) % 3 = 1 → Shard 2
hash(user_id) % 3 = 2 → Shard 3
```
More even distribution. Problem: adding a shard requires rehashing all data.

**Consistent hashing** — only moves a fraction of data when adding a shard.
Used by Cassandra and DynamoDB. Solves the rehashing problem.

---

## When to use replication vs sharding

**Read-heavy app (streaming, social media feed):**
→ Add read replicas. Distribute read traffic across replicas.
Sharding not needed until write volume is proven to be the bottleneck.

**Write-heavy app OR data too large for one machine:**
→ Shard. Only when write throughput exceeds single primary capacity,
or dataset size exceeds single machine storage.

**Key principle:** sharding adds enormous complexity. Avoid it as long as possible.
Caching removes most read pressure. Replication handles the rest.
Instagram ran on PostgreSQL with replication and caching for years before sharding.

---

## Full picture — how layers work together

```
User request
      │
      ▼
 Load Balancer      ← distributes traffic across servers
      │
      ▼
  App Server        ← stateless, any server handles any request
      │
      ├──► Cache (Redis)    ← serves most reads from memory
      │         │ miss
      │         ▼
      └──► DB Primary       ← all writes
                │
                ├──► Replica 1   ← read traffic
                └──► Replica 2   ← read traffic
```

---

## Interview answers

**"When would you choose NoSQL over SQL?"**
> "I default to PostgreSQL unless there is a specific reason not to. SQL gives you
> ACID guarantees, complex queries, and a mature ecosystem. I would consider NoSQL
> when the data is naturally document-shaped with no need for joins, when the schema
> changes too frequently for migrations, or when I need to scale writes horizontally
> across many nodes — something SQL does not do well. But most apps never hit that
> scale, so I treat NoSQL as a specialized tool, not a default."

**"What is replication lag and why does it matter?"**
> "Replication lag is the delay between a write on the primary and that write
> appearing on replicas. It is usually milliseconds but can grow under high write
> load. It matters when a user writes something and immediately reads it — if the
> read goes to a replica that hasn't caught up yet, the user sees stale data.
> Common fix: route reads-after-writes to the primary for a short window, or use
> sticky routing to always send a user's reads to the same replica."

**"When would you shard vs replicate?"**
> "Replication solves read bottlenecks by distributing reads across replicas —
> all writes still go to one primary. Sharding solves write bottlenecks and storage
> limits by splitting data across multiple independent databases. Sharding adds
> significant operational complexity — cross-shard queries, no joins, hard
> rebalancing — so I avoid it until write throughput or data size makes it
> unavoidable. Most systems can run a long time on replication plus caching alone."
