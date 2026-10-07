# CAP Theorem & SQL vs NoSQL

---

## CAP Theorem

### The Three Properties
- **C — Consistency**: Every read gets the most recent write
- **A — Availability**: Every request gets a response (not an error)
- **P — Partition Tolerance**: System keeps working when network between nodes breaks

### The Rule
Network partitions WILL happen. You can't avoid P. So you choose between C and A:

**CP (Consistency + Partition Tolerance):**
- Server refuses to answer rather than give wrong data
- "Better to say nothing than say something wrong"
- Use for: banking, payments, seat booking

**AP (Availability + Partition Tolerance):**
- Server always answers, even if data might be stale
- "Better to give an old answer than no answer"
- Use for: social media feeds, DNS, product catalog display

### Key Insight
CAP is not a system-wide choice. Different parts of your system can make different choices:
- User's bank balance → CP
- User's profile photo → AP
- Flight seat reservation → CP
- Product inventory display → AP (verify at checkout)

---

## SQL vs NoSQL

### SQL (PostgreSQL, MySQL)
- Fixed schema — every row has same columns
- Relational — tables connect via foreign keys, JOINs
- Strong consistency — ACID transactions
- Good for: structured data with relationships

### NoSQL Types
- **Document** (MongoDB) — flexible JSON-like documents
- **Key-Value** (Redis) — simple key → value, very fast
- **Wide-Column** (Cassandra) — massive write throughput
- **Search** (Elasticsearch) — full-text search and filtering

### When to Use Which

**Choose SQL when:**
- Data has clear relationships (users → orders → products)
- You need transactions (payments, bookings)
- Data structure is known and stable
- You need complex queries (JOINs, aggregations)

**Choose NoSQL when:**
- Data structure varies or evolves quickly
- Massive write throughput needed (logs, events)
- No complex relations needed
- Availability matters more than consistency

### Polyglot Persistence — Use the Right DB for Each Job

Real systems use multiple databases:

```
User profiles    → PostgreSQL   (relational, JOINs)
Orders/payments  → PostgreSQL   (ACID, strong consistency)
Product catalog  → PostgreSQL + Redis cache (read-heavy)
Activity logs    → Cassandra or ClickHouse (billions of writes)
Search/filter    → Elasticsearch (fast text search)
Session data     → Redis (fast, in-memory)
```

### CAP by Database
- PostgreSQL → CP (consistency first)
- MongoDB → CP by default, configurable
- Cassandra → AP (availability first)
- Redis → AP (speed, in-memory)

---

## Interview Answers

**"Explain CAP theorem"**
> "CAP says a distributed system can only guarantee two of three things: consistency, availability, and partition tolerance. Since network failures always happen, you really choose between consistency and availability. CP means the system refuses to respond rather than give wrong data — good for banking. AP means the system always responds even with stale data — good for social media. In practice, you make this choice per feature, not for the whole system."

**"SQL or NoSQL — how do you choose?"**
> "It depends on the use case. SQL is best for relational data that needs strong consistency — users, orders, payments. NoSQL is best for high-volume, flexible data — logs, events, caching. Most production systems use both. PostgreSQL for core data, Redis for caching, something like Cassandra for massive write loads. This is called polyglot persistence."

**"Design the database layer for an e-commerce platform"**
> "I'd use PostgreSQL as the primary database for users, orders, and products — they're relational and need ACID transactions. Product pages are read-heavy, so I'd cache them in Redis. User activity logs are billions per day with no relations, so I'd use Cassandra or ClickHouse. For product search and filtering, Elasticsearch. Each database is chosen for what it does best."
