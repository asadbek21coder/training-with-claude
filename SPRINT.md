# Interview Sprint Tracker

Legend: ✅ done | 🔄 in progress | ⬜ todo

---

## Phase 1 — Data Structures & Algorithms

### 1.1 Array / Slice ✅
**Concepts — can you answer these?**
- [ ] What is a slice internally? What are the 3 fields?
- [ ] What happens when you append and cap is full?
- [ ] What is amortized O(1)?
- [ ] What is the slice sharing trap?
- [ ] Why is passing a slice to a function not fully "by reference"?
- [ ] Is a slice concurrent safe?

**Coding problems**
- [ ] Reverse an array in place
- [ ] Find the maximum sum subarray (Kadane's algorithm)
- [ ] Move all zeros to end without changing order of others
- [ ] Find duplicates in an array

---

### 1.2 Hash Map ✅
**Concepts — can you answer these?**
- [ ] What is a bucket and how does hashing work?
- [ ] What is a collision and how is it resolved?
- [ ] What is load factor and when does resize happen?
- [ ] Why is bucket count always a power of 2?
- [ ] Why is reading from a nil map safe but writing panics?
- [ ] Why is Go map not concurrent safe?
- [ ] What physically crashes during concurrent map access?
- [ ] What does delete do to memory?
- [ ] Is map iteration order guaranteed?

**Coding problems**
- [ ] Two Sum (return indices of two numbers that add to target)
- [ ] Group anagrams together
- [ ] Find the first non-repeating character in a string
- [ ] Word frequency counter
- [ ] Check if two strings are anagrams

---

### 1.3 Linked List 🔄
**Concepts — can you answer these?**
- [ ] What is a node? What fields does it have?
- [ ] Difference between singly and doubly linked list
- [ ] Why is insert at head O(1) but insert at tail O(n)?
- [ ] Why is linked list NOT cache friendly?
- [ ] When would you use linked list over slice?

**Coding problems**
- [x] Reverse a linked list in place
- [x] Find middle node (slow/fast pointer)
- [x] Detect cycle — hash set approach
- [x] Detect cycle — slow/fast pointer
- [x] Find nth node from end (two pointer gap technique)
- [x] Merge two sorted linked lists
- [ ] Remove duplicates from sorted linked list
- [ ] Check if linked list is a palindrome

---

### 1.4 Stack & Queue 🔄
**Concepts — can you answer these?**
- [x] What is LIFO vs FIFO?
- [x] How do you implement a stack in Go?
- [x] How do you implement a queue in Go?
- [x] What is a monotonic stack and when do you use it?
- [ ] How can you implement a queue using two stacks?

**Coding problems**
- [x] Valid parentheses (check if brackets are balanced)
- [x] Implement a min stack (stack that returns min in O(1))
- [x] Daily temperatures (next warmer day for each day)
- [ ] Implement queue using two stacks
- [ ] Evaluate reverse polish notation

---

### 1.5 Binary Tree & BST 🔄
**Concepts — can you answer these?**
- [x] What is a binary tree vs a BST?
- [x] What are inorder, preorder, postorder traversals?
- [x] What is BFS (level order) traversal?
- [x] What makes a valid BST?
- [ ] What is the height/depth of a tree?
- [ ] What is a balanced tree?

**Coding problems (Cycle 1 — traversals only)**
- [x] Inorder traversal (recursive and iterative)
- [x] Preorder traversal (recursive)
- [x] Postorder traversal (recursive)
- [x] Level order traversal (BFS with queue)

**Coding problems (Cycle 2 — deferred)**
- [ ] Max depth of binary tree
- [ ] Check if binary tree is balanced
- [ ] Validate BST
- [ ] Lowest common ancestor of two nodes
- [ ] Diameter of binary tree
- [ ] Invert a binary tree

---

### 1.6 Graph 🔄
**Concepts — can you answer these?**
- [x] What is a directed vs undirected graph?
- [x] What is an adjacency list vs adjacency matrix?
- [x] How does BFS work? When do you use it?
- [x] How does DFS work? When do you use it?
- [ ] What is topological sort and when do you use it?
- [x] What is a visited set and why do you need it?

**Coding problems**
- [ ] Number of islands (BFS or DFS on grid)
- [ ] Clone a graph
- [ ] Course schedule — can you finish all courses? (cycle detection)
- [ ] BFS shortest path between two nodes
- [ ] Topological sort

---

### 1.7 Binary Search 🔄
**Concepts — can you answer these?**
- [x] What are the conditions for binary search to work?
- [x] What is the loop condition — `<` or `<=`?
- [x] How do you avoid infinite loops with mid calculation?
- [x] When do you use `left = mid + 1` vs `left = mid`?

**Coding problems**
- [x] Classic binary search (find target in sorted array)
- [ ] Search in rotated sorted array
- [ ] Find first and last position of element
- [ ] Find minimum in rotated sorted array
- [ ] Search a 2D matrix

---

### 1.8 Heap / Priority Queue 🔄
**Concepts — can you answer these?**
- [x] What is a min heap vs max heap?
- [x] How does insert and extract work in a heap?
- [x] What is the time complexity of heap operations?
- [ ] How do you use Go's `container/heap`?

**Coding problems**
- [ ] Kth largest element in an array
- [ ] Top K frequent elements
- [ ] Merge K sorted lists
- [ ] Find median from data stream

---

### 1.9 Two Pointers Pattern ⬜
**Concepts — can you answer these?**
- [ ] When do you use two pointers?
- [ ] What is the difference between same-direction and opposite-direction two pointers?

**Coding problems**
- [ ] Two sum on sorted array
- [ ] Container with most water
- [ ] Three sum (find all triplets that sum to zero)
- [ ] Remove duplicates from sorted array in place
- [ ] Trapping rain water

---

### 1.10 Sliding Window Pattern ⬜
**Concepts — can you answer these?**
- [ ] When do you use sliding window vs two pointers?
- [ ] What is a fixed-size vs variable-size window?

**Coding problems**
- [ ] Maximum sum subarray of size k (fixed window)
- [ ] Longest substring without repeating characters
- [ ] Minimum window substring containing all chars
- [ ] Longest subarray with sum <= k

---

### 1.11 Dynamic Programming — basics only ⬜
**Concepts — can you answer these?**
- [ ] What is memoization vs tabulation?
- [ ] How do you identify a DP problem?

**Coding problems**
- [ ] Fibonacci (top-down and bottom-up)
- [ ] Climbing stairs
- [ ] House robber
- [ ] Coin change (minimum coins for amount)
- [ ] Longest common subsequence

---

## Phase 2 — System Design

### 2.1 Fundamentals ✅
- [x] Vertical vs horizontal scaling
- [x] Load balancer — what it does, types (round robin, least connections)
- [x] CDN — what it does and when to use
- [x] Caching — write-through vs write-back vs write-around
- [x] Cache eviction — LRU, LFU, TTL
- [x] CAP theorem — consistency, availability, partition tolerance
- [x] SQL vs NoSQL — when to use which
- [x] Database replication — primary/replica, read replicas
- [x] Database sharding — what it is, horizontal vs vertical
- [x] Consistent hashing — why it matters for sharding
- [x] Message queues — what problems they solve, at-least-once vs exactly-once
- [x] Rate limiting strategies — token bucket, sliding window

### 2.2 Practice Designs ⬜
- [ ] Design URL shortener
- [ ] Design rate limiter
- [ ] Design Twitter/Instagram feed (news feed system)
- [ ] Design notification system (email/push/SMS)
- [ ] Design key-value store (like Redis)
- [ ] Design a cache system

---

## Phase 3 — Go Internals

### 3.1 Concurrency ⬜
- [ ] Goroutine vs OS thread — what is the difference?
- [ ] How does Go scheduler work (M:N threading model)?
- [ ] Buffered vs unbuffered channel — when to use each?
- [ ] What happens when you send to a closed channel?
- [ ] What is select and how does it work?
- [ ] What is context and how do you use it for cancellation?
- [ ] sync.WaitGroup — how and when to use
- [ ] sync.Once — what problem it solves
- [ ] What is a data race? How do you detect it?

**Coding problems**
- [ ] Fan-out pattern (one input, multiple workers)
- [ ] Fan-in pattern (multiple inputs, one output)
- [ ] Timeout using context
- [ ] Worker pool with fixed number of goroutines

### 3.2 Language internals ⬜
- [ ] How does Go garbage collector work (basics)?
- [ ] What is an interface internally? (type + pointer)
- [ ] What is the nil interface gotcha?
- [ ] How does defer work with return values?
- [ ] Difference between panic and error
- [ ] When to use pointer receiver vs value receiver?
- [ ] What is escape analysis?

---

## Phase 4 — PostgreSQL

### 4.1 Indexing ⬜
- [ ] What is a B-tree index and how does it work?
- [ ] When does an index help? When does it NOT help?
- [ ] What is a composite index? What is column order rule?
- [ ] What is a covering index?
- [ ] What is an index scan vs sequential scan?
- [ ] How do you read EXPLAIN ANALYZE output?

### 4.2 Transactions ⬜
- [ ] What are ACID properties?
- [ ] What are the four isolation levels?
- [ ] What is a dirty read, phantom read, non-repeatable read?
- [ ] What is MVCC and how does Postgres implement it?
- [ ] What is a deadlock and how to prevent it?

### 4.3 Performance ⬜
- [ ] What is the N+1 query problem?
- [ ] How do you fix N+1 in Go (with sqlx or pgx)?
- [ ] What is connection pooling and why does it matter?
- [ ] What is VACUUM and why does Postgres need it?
- [ ] When would you partition a table?
- [ ] What is the difference between partitioning and sharding?

---

## Phase 5 — Behavioral

- [ ] Tell me about yourself (2 minute answer prepared)
- [ ] Biggest technical challenge you faced and how you solved it
- [ ] A time you disagreed with a teammate or decision
- [ ] Why are you leaving your current job?
- [ ] What kind of projects or tech excites you?
- [ ] Where do you want to be in 2–3 years?
- [ ] Describe your process when you join a new codebase

---

## Progress Summary

| Phase | Topics | Done |
|-------|--------|------|
| DS & Algo | 11 topics | 4 / 11 (in progress) |
| System Design | 2 sections | 1 / 2 |
| Go Internals | 2 sections | 0 / 2 |
| PostgreSQL | 3 sections | 0 / 3 |
| Behavioral | 7 questions | 0 / 7 |

**Notes file location:** `interview-prep/notes/`
**Code practice location:** `interview-prep/code/`
