# System Design — Discussion Log (Cycle 1)

## Session: 2026-06-18

### What We Discussed

**Q1: Vertical vs Horizontal Scaling**
- User understands the difference well
- Knows vertical has limits and is expensive
- Knows horizontal distributes failure risk
- Added: horizontal introduces complexity (load balancer, stateless services, shared state)

**Q2: Load Balancer**
- User knows what a load balancer does
- Correctly described it as distributing requests across servers

**Q3: Load Balancing Strategies**
- User identified: round robin and IP hash (sticky sessions)
- Added: least connections, weighted round robin
- Discussed: the modern solution is stateless servers + external session storage (Redis)

**Q4: Reducing Database Read Load**
- User correctly identified: replication and caching
- Knows replicas sync with primary
- Added: synchronous vs asynchronous replication
- Correction: "a separate table with more indexes" is DB optimization, not caching. Cache = storing data outside the database in memory (Redis)

**Q5: Cache Invalidation (Stale Data Problem)**
- User knows TTL approach
- Added: write-through, write-behind, cache-aside patterns
- Key insight: cache-aside with TTL as safety net is the most common production pattern

### English Corrections Made
- "strength" → "strengthen"
- "devide" → "device"
- "dending" → "sending in sequential order"
- "nagruzka" → "load"
- "cashe" → "cache"
- "db ies" → "database instances" / "replicas"
- "we will request to db" → "we will query the database"

### What User Already Knows (No Need to Re-teach)
- Basic scaling concepts
- Load balancer purpose
- Replication concept
- TTL caching
- Redis as cache

**Q6-Q8: CDN**
- User correctly identified: distribute servers globally for static content
- Knows what static vs dynamic content is
- Knows TTL + invalidation approach
- Added: cache busting with versioned filenames (most common CDN invalidation)

### English Corrections Made (CDN)
- No major errors in this section

**Q9-Q11: CAP Theorem & SQL vs NoSQL**
- User understands CP vs AP trade-off
- Correctly identified: flight booking = CP, chat = AP
- Learned: product inventory display is AP (verify at checkout)
- Knows SQL vs NoSQL basics
- Correction: activity logs should NOT go in PostgreSQL — use Cassandra/ClickHouse
- Correction: "caching inside PostgreSQL" is not real caching

**Q12-Q14: Database Sharding**
- User knows range-based and hash-based sharding
- Learned consistent hashing (needed extra explanation)
- Understands scatter-gather problem

**Q15-Q16: Message Queues**
- User knows the idea of async processing with queues
- Suggested PostgreSQL as queue (valid for small scale)
- Learned: at-most-once, at-least-once, exactly-once delivery
- Learned: idempotency to handle duplicates

**Q17-Q18: Rate Limiting**
- User knows to limit requests per second
- Correctly identified IP as identifier for anonymous users
- Learned: token bucket and sliding window algorithms

### What Needs More Depth in Cycle 2
- Cache stampede / thundering herd
- Consistent hashing in more detail
- Kafka vs RabbitMQ differences
- Circuit breaker pattern
- Practice designs (URL shortener, rate limiter, etc.)
