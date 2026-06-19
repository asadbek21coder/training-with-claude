# Middle & Middle+ SWE Interview Roadmap

**Primary language:** Go · **Secondary:** Java / Spring Boot · **Kotlin** where relevant

> One checkbox per topic. Check it when you can explain it, implement it, and answer follow-up questions on it.
> Each topic has: **Main source** (best, most complete) · **Easy source** (simpler explanation / visual)

---

## How to use this

1. Go section by section in order — earlier sections block later ones
2. For each topic: read → understand → implement a small example → check the box
3. Don't memorize — understand the WHY so you can answer any variation
4. System Design and DSA are the highest-weight interview sections — prioritize them

---

## Progress Overview

| Section | Detail File | Topics | Done |
|---------|-------------|--------|------|
| 1 — Data Structures | [01-data-structures.md](topics/01-data-structures.md) | 14 | 0 |
| 2 — Algorithms | [02-algorithms.md](topics/02-algorithms.md) | 12 | 0 |
| 3 — System Design | [03-system-design.md](topics/03-system-design.md) | 16 | 0 |
| 4 — Databases | [04-databases.md](topics/04-databases.md) | 12 | 0 |
| 5 — Networking | [05-networking.md](topics/05-networking.md) | 10 | 0 |
| 6 — OS & Concurrency | [06-os-concurrency.md](topics/06-os-concurrency.md) | 8 | 0 |
| 7 — Go Language | [07-go-language.md](topics/07-go-language.md) | 15 | 0 |
| 8 — Java / Spring | [08-java-spring.md](topics/08-java-spring.md) | 14 | 0 |
| 9 — Architecture & Patterns | [09-architecture-patterns.md](topics/09-architecture-patterns.md) | 14 | 0 |
| 10 — DevOps & Infrastructure | [10-devops-infrastructure.md](topics/10-devops-infrastructure.md) | 10 | 0 |
| 11 — Security | [11-security.md](topics/11-security.md) | 8 | 0 |
| 12 — Behavioral | [12-behavioral.md](topics/12-behavioral.md) | 6 | 0 |

---

## 1. Data Structures

> You must be able to implement each from scratch AND know time/space complexity for every operation.

- [ ] **Array / Slice** — index access, resizing, amortized O(1) append
  - Main: [NeetCode — Arrays](https://neetcode.io/courses/dsa-for-beginners/1)
  - Easy: [VisuAlgo — Array](https://visualgo.net/en/array)

- [ ] **Hash Map / Hash Set** — collision handling, load factor, O(1) average
  - Main: [NeetCode — Hashing](https://neetcode.io/courses/dsa-for-beginners/26)
  - Easy: [CS50 — Hash Tables (YouTube)](https://www.youtube.com/watch?v=nvzVHwrrub0)

- [ ] **Singly Linked List** — insert/delete at head/tail/index, traversal
  - Main: [NeetCode — Linked Lists](https://neetcode.io/courses/dsa-for-beginners/5)
  - Easy: [VisuAlgo — Linked List](https://visualgo.net/en/list)

- [ ] **Doubly Linked List** — prev+next pointers, O(1) delete with pointer
  - Main: [NeetCode — Doubly Linked List](https://neetcode.io/courses/dsa-for-beginners/6)
  - Easy: [GeeksForGeeks — DLL](https://www.geeksforgeeks.org/doubly-linked-list/)

- [ ] **Stack** — LIFO, implement with slice or linked list
  - Main: [NeetCode — Stacks](https://neetcode.io/courses/dsa-for-beginners/9)
  - Easy: [VisuAlgo — Stack](https://visualgo.net/en/list)

- [ ] **Queue** — FIFO, implement with linked list (O(1) enqueue + dequeue)
  - Main: [NeetCode — Queues](https://neetcode.io/courses/dsa-for-beginners/10)
  - Easy: [VisuAlgo — Queue](https://visualgo.net/en/list)

- [ ] **Binary Tree** — nodes, leaves, height, BFS/DFS traversal (inorder, preorder, postorder)
  - Main: [NeetCode — Trees](https://neetcode.io/courses/dsa-for-beginners/14)
  - Easy: [VisuAlgo — BST](https://visualgo.net/en/bst)

- [ ] **Binary Search Tree (BST)** — insert, delete, search, balance condition
  - Main: [NeetCode — BST](https://neetcode.io/courses/dsa-for-beginners/15)
  - Easy: [VisuAlgo — BST](https://visualgo.net/en/bst)

- [ ] **Heap / Priority Queue** — min-heap, max-heap, heapify, O(log n) insert/remove
  - Main: [NeetCode — Heap](https://neetcode.io/courses/dsa-for-beginners/21)
  - Easy: [VisuAlgo — Heap](https://visualgo.net/en/heap)

- [ ] **Trie (Prefix Tree)** — insert, search, startsWith — used in autocomplete
  - Main: [NeetCode — Tries](https://neetcode.io/courses/dsa-for-beginners/22)
  - Easy: [VisuAlgo — Trie](https://visualgo.net/en/trie)

- [ ] **Graph — Adjacency List / Matrix** — represent edges, directed vs undirected
  - Main: [NeetCode — Graphs](https://neetcode.io/courses/dsa-for-beginners/23)
  - Easy: [VisuAlgo — Graph](https://visualgo.net/en/graphds)

- [ ] **Segment Tree** — range queries, O(log n) update
  - Main: [CP-Algorithms — Segment Tree](https://cp-algorithms.com/data_structures/segment_tree.html)
  - Easy: [GeeksForGeeks — Segment Tree](https://www.geeksforgeeks.org/segment-tree-data-structure/)

- [ ] **Union-Find (Disjoint Set)** — find, union, path compression — used in graph connectivity
  - Main: [NeetCode — Union Find](https://neetcode.io/courses/advanced-algorithms/7)
  - Easy: [VisuAlgo — Union Find](https://visualgo.net/en/ufds)

- [ ] **LRU Cache** — doubly linked list + hash map, O(1) get + put — classic interview problem
  - Main: [NeetCode — LRU Cache problem](https://neetcode.io/problems/lru-cache)
  - Easy: [YouTube — LRU Cache explained](https://www.youtube.com/watch?v=7ABFKPK2hD4)

---

## 2. Algorithms

> Know the pattern name, when to apply it, time complexity, and implement it.

- [ ] **Binary Search** — sorted array, O(log n), left/right boundary variants
  - Main: [NeetCode — Binary Search](https://neetcode.io/courses/dsa-for-beginners/11)
  - Easy: [VisuAlgo — Binary Search](https://visualgo.net/en/bst)

- [ ] **Two Pointers** — opposite ends or same direction, array/string problems
  - Main: [NeetCode — Two Pointers](https://neetcode.io/courses/dsa-for-beginners/12)
  - Easy: [GeeksForGeeks — Two Pointer Technique](https://www.geeksforgeeks.org/two-pointers-technique/)

- [ ] **Sliding Window** — subarray/substring of fixed or variable size
  - Main: [NeetCode — Sliding Window](https://neetcode.io/courses/dsa-for-beginners/13)
  - Easy: [YouTube — Sliding Window explained](https://www.youtube.com/watch?v=MK-NZ4hN7rs)

- [ ] **BFS (Breadth-First Search)** — level order, shortest path in unweighted graph
  - Main: [NeetCode — BFS](https://neetcode.io/courses/dsa-for-beginners/18)
  - Easy: [VisuAlgo — BFS](https://visualgo.net/en/dfsbfs)

- [ ] **DFS (Depth-First Search)** — recursion, backtracking, connected components
  - Main: [NeetCode — DFS](https://neetcode.io/courses/dsa-for-beginners/17)
  - Easy: [VisuAlgo — DFS](https://visualgo.net/en/dfsbfs)

- [ ] **Sorting** — QuickSort, MergeSort, HeapSort — know time/space, when each is used
  - Main: [NeetCode — Sorting](https://neetcode.io/courses/dsa-for-beginners/4)
  - Easy: [VisuAlgo — Sorting](https://visualgo.net/en/sorting)

- [ ] **Recursion & Backtracking** — decision trees, pruning, permutations/combinations
  - Main: [NeetCode — Backtracking](https://neetcode.io/courses/dsa-for-beginners/20)
  - Easy: [YouTube — Backtracking explained](https://www.youtube.com/watch?v=A80YzvNwqXA)

- [ ] **Dynamic Programming** — memoization vs tabulation, overlapping subproblems, optimal substructure
  - Main: [NeetCode — DP](https://neetcode.io/courses/dsa-for-beginners/28)
  - Easy: [YouTube — DP for beginners](https://www.youtube.com/watch?v=oBt53YbR9Kk)

- [ ] **Greedy Algorithms** — local optimum → global optimum, intervals, scheduling
  - Main: [NeetCode — Greedy](https://neetcode.io/courses/advanced-algorithms/1)
  - Easy: [GeeksForGeeks — Greedy](https://www.geeksforgeeks.org/greedy-algorithms/)

- [ ] **Dijkstra's Algorithm** — shortest path in weighted graph, O((V+E) log V)
  - Main: [NeetCode — Dijkstra](https://neetcode.io/courses/advanced-algorithms/9)
  - Easy: [VisuAlgo — Dijkstra](https://visualgo.net/en/sssp)

- [ ] **Topological Sort** — DAG ordering, dependency resolution (Kahn's or DFS)
  - Main: [NeetCode — Topological Sort](https://neetcode.io/courses/advanced-algorithms/5)
  - Easy: [GeeksForGeeks — Topo Sort](https://www.geeksforgeeks.org/topological-sorting/)

- [ ] **Big O Notation** — time + space, best/average/worst, amortized analysis
  - Main: [NeetCode — Big O](https://neetcode.io/courses/dsa-for-beginners/2)
  - Easy: [Big O Cheat Sheet](https://www.bigocheatsheet.com/)

---

## 3. System Design

> For middle+: you must design real systems, explain trade-offs, and justify choices.

- [ ] **Scalability** — vertical vs horizontal, stateless services, auto-scaling
  - Main: [System Design Primer](https://github.com/donnemartin/system-design-primer#scalability)
  - Easy: [ByteByteGo — Scale from 0 to millions](https://bytebytego.com/courses/system-design-interview/scale-from-zero-to-millions-of-users)

- [ ] **Load Balancing** — round-robin, least connections, consistent hashing, health checks
  - Main: [System Design Primer — Load Balancer](https://github.com/donnemartin/system-design-primer#load-balancer)
  - Easy: [ByteByteGo — Load Balancer](https://bytebytego.com/courses/system-design-interview/design-a-load-balancer)

- [ ] **Caching** — cache-aside, write-through, eviction policies, Redis, CDN
  - Main: [System Design Primer — Cache](https://github.com/donnemartin/system-design-primer#cache)
  - Easy: [ByteByteGo — Caching Strategies](https://bytebytego.com/courses/system-design-interview/design-a-cache)

- [ ] **Databases — SQL vs NoSQL** — when to use each, trade-offs, examples
  - Main: [System Design Primer — DB](https://github.com/donnemartin/system-design-primer#relational-database-management-system-rdbms)
  - Easy: [ByteByteGo — SQL vs NoSQL](https://bytebytego.com/courses/system-design-interview/sql-vs-nosql)

- [ ] **Database Indexing** — B-tree, composite indexes, covering index, slow queries
  - Main: [Use The Index, Luke](https://use-the-index-luke.com/)
  - Easy: [ByteByteGo — Indexes](https://bytebytego.com/courses/system-design-interview/database-indexes)

- [ ] **Database Replication** — primary-replica, eventual consistency, read replicas
  - Main: [System Design Primer — Replication](https://github.com/donnemartin/system-design-primer#master-slave-replication)
  - Easy: [ByteByteGo — Replication](https://bytebytego.com/courses/system-design-interview/database-replication)

- [ ] **Database Sharding** — horizontal partitioning, shard key, resharding
  - Main: [System Design Primer — Sharding](https://github.com/donnemartin/system-design-primer#sharding)
  - Easy: [ByteByteGo — Sharding](https://bytebytego.com/courses/system-design-interview/database-sharding)

- [ ] **Message Queues** — async decoupling, Kafka vs RabbitMQ, at-least-once delivery
  - Main: [System Design Primer — Message Queue](https://github.com/donnemartin/system-design-primer#message-queues)
  - Easy: [ByteByteGo — Message Queue](https://bytebytego.com/courses/system-design-interview/message-queues)

- [ ] **API Design — REST** — resources, HTTP verbs, status codes, versioning, idempotency
  - Main: [REST API Design Best Practices — Microsoft](https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design)
  - Easy: [ByteByteGo — REST API](https://bytebytego.com/courses/system-design-interview/restful-api)

- [ ] **API Design — gRPC** — protobuf, streaming, when to use over REST
  - Main: [gRPC official docs](https://grpc.io/docs/what-is-grpc/introduction/)
  - Easy: [ByteByteGo — gRPC](https://bytebytego.com/courses/system-design-interview/grpc)

- [ ] **Rate Limiting** — token bucket, leaky bucket, sliding window, Redis-based
  - Main: [System Design Primer — Rate Limiting](https://github.com/donnemartin/system-design-primer#rate-limiter)
  - Easy: [ByteByteGo — Rate Limiter](https://bytebytego.com/courses/system-design-interview/design-a-rate-limiter)

- [ ] **CAP Theorem** — Consistency, Availability, Partition Tolerance — trade-offs
  - Main: [System Design Primer — CAP](https://github.com/donnemartin/system-design-primer#cap-theorem)
  - Easy: [ByteByteGo — CAP](https://bytebytego.com/courses/system-design-interview/cap-theorem)

- [ ] **Consistent Hashing** — ring, virtual nodes — used in distributed caches and DBs
  - Main: [System Design Primer — Consistent Hashing](https://github.com/donnemartin/system-design-primer#consistent-hashing)
  - Easy: [ByteByteGo — Consistent Hashing](https://bytebytego.com/courses/system-design-interview/consistent-hashing)

- [ ] **Microservices vs Monolith** — trade-offs, service discovery, inter-service communication
  - Main: [Martin Fowler — Microservices](https://martinfowler.com/articles/microservices.html)
  - Easy: [ByteByteGo — Microservices](https://bytebytego.com/courses/system-design-interview/microservices-vs-monolith)

- [ ] **Real-time — WebSockets & Server-Sent Events** — when to use, long polling fallback
  - Main: [System Design Primer — WebSockets](https://github.com/donnemartin/system-design-primer#communication)
  - Easy: [ByteByteGo — Real-time Comms](https://bytebytego.com/courses/system-design-interview/long-polling-vs-websockets-vs-server-sent-events)

- [ ] **Design a System (practice problems)** — URL shortener, chat app, rate limiter, news feed, notification system
  - Main: [System Design Primer — Design Questions](https://github.com/donnemartin/system-design-primer#system-design-interview-questions-with-solutions)
  - Easy: [ByteByteGo Free Chapter — URL Shortener](https://bytebytego.com/courses/system-design-interview/design-a-url-shortener)

---

## 4. Databases

- [ ] **SQL Basics** — SELECT, JOIN (inner/left/right/full), GROUP BY, HAVING, subqueries
  - Main: [SQLZoo](https://sqlzoo.net/)
  - Easy: [W3Schools — SQL](https://www.w3schools.com/sql/)

- [ ] **Indexes** — B-tree index, composite index, when index hurts, EXPLAIN
  - Main: [Use The Index, Luke](https://use-the-index-luke.com/)
  - Easy: [ByteByteGo — DB Indexes](https://bytebytego.com/courses/system-design-interview/database-indexes)

- [ ] **Transactions & ACID** — atomicity, consistency, isolation, durability
  - Main: [PostgreSQL Docs — Transactions](https://www.postgresql.org/docs/current/tutorial-transactions.html)
  - Easy: [ByteByteGo — ACID](https://bytebytego.com/courses/system-design-interview/acid-properties)

- [ ] **Isolation Levels** — read uncommitted → serializable, dirty reads, phantom reads
  - Main: [PostgreSQL Docs — Isolation](https://www.postgresql.org/docs/current/transaction-iso.html)
  - Easy: [GeeksForGeeks — Isolation Levels](https://www.geeksforgeeks.org/transaction-isolation-levels-dbms/)

- [ ] **Database Normalization** — 1NF, 2NF, 3NF — remove redundancy
  - Main: [GeeksForGeeks — Normalization](https://www.geeksforgeeks.org/normal-forms-in-dbms/)
  - Easy: [YouTube — Normalization](https://www.youtube.com/watch?v=GFQaEYEc8_8)

- [ ] **Query Optimization** — EXPLAIN, index usage, N+1 problem, query planner
  - Main: [Use The Index, Luke — Execution Plans](https://use-the-index-luke.com/sql/explain-plan)
  - Easy: [GeeksForGeeks — Query Optimization](https://www.geeksforgeeks.org/query-optimization-in-relational-algebra/)

- [ ] **PostgreSQL** — JSONB, CTEs, window functions, full text search
  - Main: [PostgreSQL Official Docs](https://www.postgresql.org/docs/current/)
  - Easy: [PostgreSQL Tutorial](https://www.postgresqltutorial.com/)

- [ ] **Redis** — data types (string/list/set/sorted set/hash), TTL, pub/sub, persistence
  - Main: [Redis Docs](https://redis.io/docs/)
  - Easy: [Redis University — RU101 (free)](https://university.redis.com/courses/ru101/)

- [ ] **MongoDB** — documents, collections, aggregation pipeline, indexes
  - Main: [MongoDB Docs](https://www.mongodb.com/docs/)
  - Easy: [MongoDB University — M001 (free)](https://university.mongodb.com/courses/M001/about)

- [ ] **Kafka basics** — topic, partition, offset, consumer group, at-least-once
  - Main: [Kafka Docs — Introduction](https://kafka.apache.org/documentation/)
  - Easy: [ByteByteGo — Kafka](https://bytebytego.com/courses/system-design-interview/kafka)

- [ ] **Database connection pooling** — why it exists, pgx pool in Go, HikariCP in Java
  - Main: [pgx Pool docs](https://pkg.go.dev/github.com/jackc/pgx/v5/pgxpool)
  - Easy: [GeeksForGeeks — Connection Pooling](https://www.geeksforgeeks.org/introduction-of-database-connection-pooling/)

- [ ] **Migrations** — schema versioning, rollback, tools (Flyway, golang-migrate)
  - Main: [golang-migrate docs](https://github.com/golang-migrate/migrate)
  - Easy: [Flyway — Concepts](https://flywaydb.org/documentation/concepts/migrations)

---

## 5. Networking

- [ ] **HTTP / HTTPS** — request/response, headers, methods, status codes, TLS handshake
  - Main: [MDN — HTTP Overview](https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview)
  - Easy: [ByteByteGo — HTTPS](https://bytebytego.com/courses/system-design-interview/how-does-https-work)

- [ ] **HTTP/1.1 vs HTTP/2 vs HTTP/3** — pipelining, multiplexing, QUIC
  - Main: [Cloudflare — HTTP/3](https://cloudflare.com/learning/performance/what-is-http3/)
  - Easy: [ByteByteGo — HTTP versions](https://bytebytego.com/courses/system-design-interview/http-1-vs-http-2-vs-http-3)

- [ ] **TCP vs UDP** — reliability, ordering, connection, use cases (video vs file)
  - Main: [Cloudflare — TCP vs UDP](https://www.cloudflare.com/learning/ddos/glossary/user-datagram-protocol-udp/)
  - Easy: [ByteByteGo — TCP vs UDP](https://bytebytego.com/courses/system-design-interview/tcp-vs-udp)

- [ ] **DNS** — resolution process, A record, CNAME, TTL, recursive vs authoritative
  - Main: [Cloudflare — How DNS works](https://www.cloudflare.com/learning/dns/what-is-dns/)
  - Easy: [ByteByteGo — DNS](https://bytebytego.com/courses/system-design-interview/how-does-dns-work)

- [ ] **TLS / SSL** — handshake, certificates, symmetric vs asymmetric encryption
  - Main: [Cloudflare — TLS Handshake](https://www.cloudflare.com/learning/ssl/what-happens-in-a-tls-handshake/)
  - Easy: [ByteByteGo — TLS](https://bytebytego.com/courses/system-design-interview/ssl-tls)

- [ ] **WebSockets** — full-duplex, when to use, upgrade handshake
  - Main: [MDN — WebSockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API)
  - Easy: [ByteByteGo — WebSockets](https://bytebytego.com/courses/system-design-interview/long-polling-vs-websockets-vs-server-sent-events)

- [ ] **REST vs gRPC vs GraphQL** — trade-offs, when each fits
  - Main: [gRPC vs REST — gRPC blog](https://grpc.io/blog/grpc-stacks/)
  - Easy: [ByteByteGo — REST vs gRPC vs GraphQL](https://bytebytego.com/courses/system-design-interview/rest-api-vs-graphql-vs-grpc)

- [ ] **Reverse Proxy vs API Gateway** — Nginx, Kong, AWS API Gateway, auth/routing/rate limiting
  - Main: [Nginx Docs — Reverse Proxy](https://docs.nginx.com/nginx/admin-guide/web-server/reverse-proxy/)
  - Easy: [ByteByteGo — API Gateway](https://bytebytego.com/courses/system-design-interview/design-an-api-gateway)

- [ ] **OSI Model** — 7 layers, which layer HTTP/TCP/IP lives at
  - Main: [Cloudflare — OSI Model](https://www.cloudflare.com/learning/ddos/glossary/open-systems-interconnection-model-osi/)
  - Easy: [GeeksForGeeks — OSI](https://www.geeksforgeeks.org/layers-of-osi-model/)

- [ ] **CDN** — edge servers, pull vs push, Cache-Control, providers
  - Main: [Cloudflare — What is a CDN](https://www.cloudflare.com/learning/cdn/what-is-a-cdn/)
  - Easy: [ByteByteGo — CDN](https://bytebytego.com/courses/system-design-interview/content-delivery-network-cdn)

---

## 6. Operating Systems & Concurrency

- [ ] **Process vs Thread** — memory space, context switch, cost difference
  - Main: [OS Concepts — Silberschatz (free PDF)](https://os.ecci.ucr.ac.cr/slides/Abraham-Silberschatz-Operating-System-Concepts-10th-2018.pdf)
  - Easy: [GeeksForGeeks — Process vs Thread](https://www.geeksforgeeks.org/difference-between-process-and-thread/)

- [ ] **Concurrency vs Parallelism** — what's the difference, Go's goroutines
  - Main: [Go Blog — Concurrency is not Parallelism](https://go.dev/blog/waza-talk)
  - Easy: [ByteByteGo — Concurrency](https://bytebytego.com/courses/system-design-interview/concurrency-vs-parallelism)

- [ ] **Race Conditions & Deadlocks** — what causes them, how to prevent
  - Main: [Go Blog — Data Race Detector](https://go.dev/doc/articles/race_detector)
  - Easy: [GeeksForGeeks — Deadlock](https://www.geeksforgeeks.org/introduction-of-deadlock-in-operating-system/)

- [ ] **Mutex / Semaphore / RWMutex** — when to use each
  - Main: [Go sync package docs](https://pkg.go.dev/sync)
  - Easy: [GeeksForGeeks — Mutex vs Semaphore](https://www.geeksforgeeks.org/mutex-vs-semaphore/)

- [ ] **Memory Management** — stack vs heap, garbage collection, GC pauses
  - Main: [Go GC Guide](https://tip.golang.org/doc/gc-guide)
  - Easy: [JVM GC — Oracle Docs](https://www.oracle.com/webfolder/technetwork/tutorials/oreilly/java-virtual-machine/index.html)

- [ ] **I/O — Blocking vs Non-blocking** — epoll, event loop, Go's net poller
  - Main: [Go — Network Poller](https://go.dev/src/runtime/netpoll.go) (source + comments)
  - Easy: [ByteByteGo — Blocking vs Non-blocking IO](https://bytebytego.com/courses/system-design-interview/blocking-vs-non-blocking-io)

- [ ] **Linux basics** — processes (ps, top), files (ls, chmod), network (curl, netstat, ss)
  - Main: [Linux Command Line — free book](https://linuxcommand.org/tlcl.php)
  - Easy: [OverTheWire — Bandit (hands-on)](https://overthewire.org/wargames/bandit/)

- [ ] **Signals & Process lifecycle** — SIGTERM, SIGKILL, graceful shutdown
  - Main: [Go — os/signal package](https://pkg.go.dev/os/signal)
  - Easy: [GeeksForGeeks — Linux Signals](https://www.geeksforgeeks.org/signals-in-linux/)

---

## 7. Go Language (Deep)

- [ ] **Goroutines** — lightweight threads, M:N scheduling, goroutine leak
  - Main: [Go Tour — Goroutines](https://go.dev/tour/concurrency/1)
  - Easy: [Go By Example — Goroutines](https://gobyexample.com/goroutines)

- [ ] **Channels** — buffered vs unbuffered, select, deadlock, range over channel
  - Main: [Go Tour — Channels](https://go.dev/tour/concurrency/2)
  - Easy: [Go By Example — Channels](https://gobyexample.com/channels)

- [ ] **Context package** — cancellation, timeout, deadline, passing values
  - Main: [Go Blog — Context](https://go.dev/blog/context)
  - Easy: [Go By Example — Context](https://gobyexample.com/context)

- [ ] **Interfaces** — implicit implementation, duck typing, empty interface, type assertion
  - Main: [Go Tour — Interfaces](https://go.dev/tour/methods/9)
  - Easy: [Go By Example — Interfaces](https://gobyexample.com/interfaces)

- [ ] **Error handling** — errors.Is, errors.As, %w wrapping, custom error types
  - Main: [Go Blog — Error Handling](https://go.dev/blog/error-handling-and-go)
  - Easy: [Go By Example — Errors](https://gobyexample.com/errors)

- [ ] **sync package** — Mutex, RWMutex, WaitGroup, Once, Map
  - Main: [Go sync package docs](https://pkg.go.dev/sync)
  - Easy: [Go By Example — Mutexes](https://gobyexample.com/mutexes)

- [ ] **defer / panic / recover** — execution order, use in cleanup, recover from panics
  - Main: [Go Blog — Defer Panic Recover](https://go.dev/blog/defer-panic-and-recover)
  - Easy: [Go By Example — Defer](https://gobyexample.com/defer)

- [ ] **Slices internals** — backing array, len/cap, append, slice of slice pitfall
  - Main: [Go Blog — Slices](https://go.dev/blog/slices-intro)
  - Easy: [Go By Example — Slices](https://gobyexample.com/slices)

- [ ] **Maps internals** — hash function, concurrent access (sync.Map), nil map panic
  - Main: [Go Blog — Maps in action](https://go.dev/blog/maps)
  - Easy: [Go By Example — Maps](https://gobyexample.com/maps)

- [ ] **Go Modules** — go.mod, go.sum, versioning, replace directives
  - Main: [Go Modules Reference](https://go.dev/ref/mod)
  - Easy: [Go Blog — Modules](https://go.dev/blog/using-go-modules)

- [ ] **Testing** — table-driven tests, subtests (t.Run), benchmarks, mocks
  - Main: [Go Testing Package](https://pkg.go.dev/testing)
  - Easy: [Go By Example — Testing](https://gobyexample.com/testing)

- [ ] **Go Scheduler (GMP model)** — G (goroutine), M (OS thread), P (processor)
  - Main: [Go Runtime Scheduler — Ardan Labs](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)
  - Easy: [YouTube — Go Scheduler](https://www.youtube.com/watch?v=YHRO5WQGh0k)

- [ ] **Reflection** — reflect package, use cases, when to avoid
  - Main: [Go Blog — Laws of Reflection](https://go.dev/blog/laws-of-reflection)
  - Easy: [Go By Example — Reflection](https://gobyexample.com/reflection)

- [ ] **HTTP server in Go** — net/http, middleware pattern, mux, graceful shutdown
  - Main: [Go net/http docs](https://pkg.go.dev/net/http)
  - Easy: [Go By Example — HTTP Server](https://gobyexample.com/http-servers)

- [ ] **Common Go patterns** — functional options, options struct, worker pool, pipeline
  - Main: [Go Blog — Pipelines](https://go.dev/blog/pipelines)
  - Easy: [GitHub — Go Patterns](https://github.com/tmrts/go-patterns)

---

## 8. Java / Spring Boot

- [ ] **JVM internals** — stack, heap, metaspace, JIT compilation, GC (G1, ZGC)
  - Main: [Oracle JVM Spec](https://docs.oracle.com/javase/specs/jvms/se21/html/index.html)
  - Easy: [Baeldung — JVM Architecture](https://www.baeldung.com/jvm-parameters)

- [ ] **Java Garbage Collection** — generational GC, GC tuning, GC logs
  - Main: [Oracle GC Tuning Guide](https://docs.oracle.com/en/java/javase/21/gctuning/)
  - Easy: [Baeldung — Java GC](https://www.baeldung.com/jvm-garbage-collectors)

- [ ] **Java Concurrency** — Thread, Runnable, ExecutorService, CompletableFuture
  - Main: [Java Concurrency in Practice (free excerpts)](http://jcip.net/)
  - Easy: [Baeldung — Java Concurrency](https://www.baeldung.com/java-concurrency)

- [ ] **Spring Boot basics** — auto-configuration, application context, starter dependencies
  - Main: [Spring Boot Docs](https://docs.spring.io/spring-boot/docs/current/reference/htmlsingle/)
  - Easy: [Baeldung — Spring Boot](https://www.baeldung.com/spring-boot)

- [ ] **Dependency Injection / IoC** — @Bean, @Component, @Autowired, Bean lifecycle
  - Main: [Spring Core Docs — IoC](https://docs.spring.io/spring-framework/docs/current/reference/html/core.html)
  - Easy: [Baeldung — Spring IoC](https://www.baeldung.com/inversion-control-and-dependency-injection-in-spring)

- [ ] **Spring MVC / REST** — @RestController, @RequestMapping, @PathVariable, exception handling
  - Main: [Spring MVC Docs](https://docs.spring.io/spring-framework/docs/current/reference/html/web.html)
  - Easy: [Baeldung — Spring REST](https://www.baeldung.com/building-a-restful-web-service-with-spring-and-java-based-configuration)

- [ ] **Spring Data JPA / Hibernate** — @Entity, @Repository, lazy vs eager loading, N+1 problem
  - Main: [Spring Data JPA Docs](https://docs.spring.io/spring-data/jpa/docs/current/reference/html/)
  - Easy: [Baeldung — Spring Data JPA](https://www.baeldung.com/the-persistence-layer-with-spring-data-jpa)

- [ ] **Spring Security** — filter chain, authentication vs authorization, JWT integration
  - Main: [Spring Security Docs](https://docs.spring.io/spring-security/reference/)
  - Easy: [Baeldung — Spring Security JWT](https://www.baeldung.com/spring-security-oauth-jwt)

- [ ] **Spring Transactions** — @Transactional, propagation, isolation levels, rollback rules
  - Main: [Spring Transaction Docs](https://docs.spring.io/spring-framework/docs/current/reference/html/data-access.html#transaction)
  - Easy: [Baeldung — Spring Transactions](https://www.baeldung.com/transaction-configuration-with-jpa-and-spring)

- [ ] **Spring Boot Testing** — @SpringBootTest, @MockBean, @WebMvcTest, Testcontainers
  - Main: [Spring Boot Testing Docs](https://docs.spring.io/spring-boot/docs/current/reference/htmlsingle/#features.testing)
  - Easy: [Baeldung — Spring Boot Tests](https://www.baeldung.com/spring-boot-testing)

- [ ] **Spring WebFlux (Reactive)** — Mono, Flux, backpressure, when to use reactive
  - Main: [Spring WebFlux Docs](https://docs.spring.io/spring-framework/docs/current/reference/html/web-reactive.html)
  - Easy: [Baeldung — WebFlux](https://www.baeldung.com/spring-webflux)

- [ ] **Kotlin basics** — null safety, data classes, extensions, coroutines
  - Main: [Kotlin Docs](https://kotlinlang.org/docs/home.html)
  - Easy: [Kotlin by Example](https://play.kotlinlang.org/byExample/overview)

- [ ] **Kotlin Coroutines** — suspend functions, CoroutineScope, Dispatchers, Flow
  - Main: [Kotlin Coroutines Guide](https://kotlinlang.org/docs/coroutines-guide.html)
  - Easy: [Baeldung — Kotlin Coroutines](https://www.baeldung.com/kotlin/coroutines)

- [ ] **Maven / Gradle** — dependency management, build lifecycle, multi-module projects
  - Main: [Maven Docs](https://maven.apache.org/guides/) · [Gradle Docs](https://docs.gradle.org/current/userguide/userguide.html)
  - Easy: [Baeldung — Maven](https://www.baeldung.com/maven)

---

## 9. Architecture & Design Patterns

- [ ] **SOLID Principles** — Single Responsibility, Open/Closed, Liskov, Interface Segregation, Dependency Inversion
  - Main: [Martin Fowler — SOLID](https://martinfowler.com/tags/encapsulation.html)
  - Easy: [Baeldung — SOLID](https://www.baeldung.com/solid-principles)

- [ ] **GoF Design Patterns (Creational)** — Singleton, Factory, Builder, Prototype
  - Main: [Refactoring.Guru — Creational Patterns](https://refactoring.guru/design-patterns/creational-patterns)
  - Easy: [Go Patterns — GitHub](https://github.com/tmrts/go-patterns#creational-patterns)

- [ ] **GoF Design Patterns (Structural)** — Adapter, Decorator, Proxy, Facade, Composite
  - Main: [Refactoring.Guru — Structural Patterns](https://refactoring.guru/design-patterns/structural-patterns)
  - Easy: [Go Patterns — GitHub](https://github.com/tmrts/go-patterns#structural-patterns)

- [ ] **GoF Design Patterns (Behavioral)** — Observer, Strategy, Command, Iterator, Template Method
  - Main: [Refactoring.Guru — Behavioral Patterns](https://refactoring.guru/design-patterns/behavioral-patterns)
  - Easy: [Go Patterns — GitHub](https://github.com/tmrts/go-patterns#behavioral-patterns)

- [ ] **Clean Architecture / Hexagonal Architecture** — domain, ports, adapters, dependency direction
  - Main: [Clean Architecture — Robert Martin (blog)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
  - Easy: [Baeldung — Hexagonal Architecture](https://www.baeldung.com/hexagonal-architecture-ddd-spring)

- [ ] **Domain-Driven Design (DDD) basics** — entity, value object, aggregate, bounded context
  - Main: [Martin Fowler — DDD](https://martinfowler.com/bliki/DomainDrivenDesign.html)
  - Easy: [Baeldung — DDD](https://www.baeldung.com/spring-boot-domain-driven-design)

- [ ] **Event-Driven Architecture** — events, producers, consumers, event bus
  - Main: [Martin Fowler — Event-Driven](https://martinfowler.com/articles/201701-event-driven.html)
  - Easy: [ByteByteGo — Event-Driven](https://bytebytego.com/courses/system-design-interview/event-driven-architecture)

- [ ] **CQRS** — Command Query Responsibility Segregation — separate read/write models
  - Main: [Martin Fowler — CQRS](https://martinfowler.com/bliki/CQRS.html)
  - Easy: [Baeldung — CQRS](https://www.baeldung.com/cqrs-event-sourcing-java)

- [ ] **Event Sourcing** — store events not state, replay, audit log
  - Main: [Martin Fowler — Event Sourcing](https://martinfowler.com/eaaDev/EventSourcing.html)
  - Easy: [Baeldung — Event Sourcing](https://www.baeldung.com/cqrs-event-sourcing-java)

- [ ] **Saga Pattern** — distributed transactions without 2PC, choreography vs orchestration
  - Main: [Microservices.io — Saga](https://microservices.io/patterns/data/saga.html)
  - Easy: [Baeldung — Saga Pattern](https://www.baeldung.com/cs/saga-pattern-microservices)

- [ ] **Circuit Breaker Pattern** — fail fast, half-open state, Hystrix/Resilience4j
  - Main: [Martin Fowler — Circuit Breaker](https://martinfowler.com/bliki/CircuitBreaker.html)
  - Easy: [Baeldung — Resilience4j](https://www.baeldung.com/resilience4j)

- [ ] **Repository Pattern** — abstract data access, testability, separation of concerns
  - Main: [Martin Fowler — Repository](https://martinfowler.com/eaaCatalog/repository.html)
  - Easy: [Baeldung — Repository Pattern](https://www.baeldung.com/java-dao-vs-repository)

- [ ] **12-Factor App** — config, stateless, logs, backing services — cloud-native principles
  - Main: [12factor.net](https://12factor.net/)
  - Easy: [Baeldung — 12 Factor](https://www.baeldung.com/spring-boot-12-factor)

- [ ] **API versioning strategies** — URL path, header, query param — pros and cons
  - Main: [Microsoft API Design — Versioning](https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design#versioning-a-restful-web-api)
  - Easy: [Baeldung — API Versioning](https://www.baeldung.com/rest-versioning)

---

## 10. DevOps & Infrastructure

- [ ] **Docker** — image, container, Dockerfile, layers, multi-stage builds, volumes, networks
  - Main: [Docker Docs](https://docs.docker.com/get-started/)
  - Easy: [Play with Docker](https://labs.play-with-docker.com/)

- [ ] **Docker Compose** — multi-container local dev, depends_on, networks, env files
  - Main: [Docker Compose Docs](https://docs.docker.com/compose/)
  - Easy: [GeeksForGeeks — Docker Compose](https://www.geeksforgeeks.org/docker-compose/)

- [ ] **Kubernetes basics** — Pod, Deployment, Service, ConfigMap, Secret, Ingress
  - Main: [Kubernetes Docs](https://kubernetes.io/docs/concepts/)
  - Easy: [Play with Kubernetes](https://labs.play-with-k8s.com/)

- [ ] **CI/CD** — build → test → lint → deploy pipeline, GitHub Actions, GitLab CI
  - Main: [GitHub Actions Docs](https://docs.github.com/en/actions)
  - Easy: [GeeksForGeeks — CI/CD](https://www.geeksforgeeks.org/ci-cd-continuous-integration-and-continuous-delivery/)

- [ ] **Cloud basics (AWS/GCP)** — EC2/Compute Engine, S3/GCS, RDS, Lambda/Cloud Functions
  - Main: [AWS Free Tier Docs](https://aws.amazon.com/free/)
  - Easy: [ByteByteGo — Cloud Basics](https://bytebytego.com/courses/system-design-interview/cloud-services-cheat-sheet)

- [ ] **Infrastructure as Code** — Terraform basics, declarative vs imperative
  - Main: [Terraform Docs — Get Started](https://developer.hashicorp.com/terraform/tutorials)
  - Easy: [GeeksForGeeks — Terraform](https://www.geeksforgeeks.org/infrastructure-as-code/)

- [ ] **Observability** — logs, metrics, traces (the three pillars), structured logging
  - Main: [OpenTelemetry Docs](https://opentelemetry.io/docs/)
  - Easy: [ByteByteGo — Observability](https://bytebytego.com/courses/system-design-interview/monitoring-and-alerting)

- [ ] **Prometheus + Grafana** — metrics collection, PromQL, alerting rules, dashboards
  - Main: [Prometheus Docs](https://prometheus.io/docs/introduction/overview/)
  - Easy: [GeeksForGeeks — Prometheus](https://www.geeksforgeeks.org/what-is-prometheus/)

- [ ] **Distributed Tracing** — trace ID, span, Jaeger / OpenTelemetry, propagation
  - Main: [OpenTelemetry — Tracing](https://opentelemetry.io/docs/concepts/signals/traces/)
  - Easy: [ByteByteGo — Distributed Tracing](https://bytebytego.com/courses/system-design-interview/distributed-tracing)

- [ ] **Git advanced** — rebase vs merge, cherry-pick, bisect, reflog, branching strategies
  - Main: [Pro Git Book (free)](https://git-scm.com/book/en/v2)
  - Easy: [Learn Git Branching (interactive)](https://learngitbranching.js.org/)

---

## 11. Security

- [ ] **Authentication vs Authorization** — who are you vs what can you do
  - Main: [OWASP — Authentication Cheatsheet](https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html)
  - Easy: [Cloudflare — AuthN vs AuthZ](https://www.cloudflare.com/learning/access-management/authn-vs-authz/)

- [ ] **JWT** — structure (header.payload.signature), stateless auth, expiry, refresh tokens
  - Main: [JWT.io — Introduction](https://jwt.io/introduction)
  - Easy: [ByteByteGo — JWT](https://bytebytego.com/courses/system-design-interview/jwt)

- [ ] **OAuth 2.0 / OIDC** — authorization code flow, access token, refresh token, scopes
  - Main: [OAuth 2.0 — RFC 6749](https://oauth.net/2/)
  - Easy: [ByteByteGo — OAuth](https://bytebytego.com/courses/system-design-interview/oauth-2)

- [ ] **OWASP Top 10** — injection, broken auth, XSS, IDOR, misconfig — must know all 10
  - Main: [OWASP Top 10](https://owasp.org/www-project-top-ten/)
  - Easy: [PortSwigger Web Security Academy — free labs](https://portswigger.net/web-security)

- [ ] **SQL Injection** — how it works, parameterized queries, ORM protection
  - Main: [OWASP — SQL Injection](https://owasp.org/www-community/attacks/SQL_Injection)
  - Easy: [PortSwigger — SQL Injection](https://portswigger.net/web-security/sql-injection)

- [ ] **HTTPS / TLS in practice** — certificate pinning, HSTS, certificate transparency
  - Main: [Cloudflare — TLS](https://www.cloudflare.com/learning/ssl/transport-layer-security-tls/)
  - Easy: [Bytebytego — HTTPS](https://bytebytego.com/courses/system-design-interview/how-does-https-work)

- [ ] **Secrets management** — env vars, Vault, secret rotation, never commit secrets
  - Main: [HashiCorp Vault Docs](https://developer.hashicorp.com/vault/docs)
  - Easy: [12factor — Config](https://12factor.net/config)

- [ ] **Rate Limiting & DDoS basics** — token bucket, WAF, Cloudflare protection
  - Main: [Cloudflare — DDoS](https://www.cloudflare.com/learning/ddos/what-is-a-ddos-attack/)
  - Easy: [ByteByteGo — Rate Limiter](https://bytebytego.com/courses/system-design-interview/design-a-rate-limiter)

---

## 12. Behavioral & Communication

- [ ] **STAR method** — Situation, Task, Action, Result — structure every story
  - Main: [STAR Method — Indeed Guide](https://www.indeed.com/career-advice/interviewing/how-to-use-the-star-interview-response-technique)
  - Easy: [YouTube — STAR Method](https://www.youtube.com/watch?v=8QfSnuL8Ny8)

- [ ] **Technical communication** — explain trade-offs clearly, ask clarifying questions first
  - Main: [System Design Interview — Alex Xu (free chapters)](https://bytebytego.com/courses/system-design-interview/a-framework-for-system-design-interviews)
  - Easy: [YouTube — How to approach SD interview](https://www.youtube.com/watch?v=i7twT3x5yv8)

- [ ] **Conflict resolution stories** — disagreements with teammates, technical debt decisions
  - Main: [STAR Method — Conflict](https://www.indeed.com/career-advice/interviewing/star-method-conflict)
  - Easy: [YouTube — Behavioral Interview](https://www.youtube.com/watch?v=mHn-Y_SSTps)

- [ ] **Failure stories** — what went wrong, what you learned, what you changed
  - Main: [Indeed — Tell me about a failure](https://www.indeed.com/career-advice/interviewing/interview-question-tell-me-about-a-time-you-failed)
  - Easy: [YouTube — Failure question](https://www.youtube.com/watch?v=tKTOKTYKGRY)

- [ ] **Leadership & ownership** — examples of taking initiative, driving a project
  - Main: [Indeed — Leadership Questions](https://www.indeed.com/career-advice/interviewing/interview-questions-about-leadership)
  - Easy: [YouTube — Leadership interview](https://www.youtube.com/watch?v=mYm0EIvGMGM)

- [ ] **Questions to ask the interviewer** — team structure, tech stack, oncall, growth path
  - Main: [GitHub — Reverse Interview](https://github.com/viraptor/reverse-interview)
  - Easy: [Indeed — Questions to ask](https://www.indeed.com/career-advice/interviewing/questions-to-ask-in-an-interview)
