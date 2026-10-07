# Algorithms & Data Structures Roadmap

> **L1** = aware & can use · **L2** = deep & efficient · Status: ⬜ not started · 🔄 in progress · ✅ done
> **Notes:** `notes/algorithms-ds/` · **Practice:** `code/algorithms-ds/` · **Dictionary:** `reference/algorithms-ds-structures.md`, `reference/algorithms-ds-algorithms.md`

**Why (EN):** Interviews at strong companies still test DSA, and a senior engineer must pick the right structure and see the cost of code (allocations, O(n²) loops, cache misses) before it reaches production. In the AI era, AI writes the standard solutions quickly, so what matters more is recognizing patterns, judging complexity and correctness of generated code, and explaining trade-offs out loud; memorizing exotic algorithms matters less.
**Зачем (RU):** Сильные компании до сих пор проверяют алгоритмы на собеседованиях, а senior-инженер должен выбирать правильную структуру данных и видеть стоимость кода (аллокации, циклы O(n²), промахи кэша) до продакшена. В эпоху ИИ стандартные решения ИИ пишет быстро, поэтому важнее узнавать паттерны, оценивать сложность и корректность сгенерированного кода и объяснять компромиссы вслух; заучивать экзотические алгоритмы — менее важно.

## Progress
| Layer | Done | Total |
|---|---|---|
| L1 | 13 | 47 |
| L2 | 0 | 42 |

---

## DSA 1 — Complexity analysis
**L1 — aware & can use**
- ✅ **1.1 Big-O basics** — time complexity, space complexity, O(1), O(log n), O(n), O(n log n), O(n²), O(2ⁿ), dropping constants, dominant term
- ⬜ **1.2 Reading the complexity of code** — nested loops, loop that halves `n` → O(log n), counting recursive calls, `append` amortized O(1), map operations average O(1), string concatenation in a loop, sort inside a loop, hidden copies

**L2 — deep & efficient**
- ⬜ **1.3 Advanced analysis** — amortized analysis, best / average / worst case, recursion tree, master theorem (awareness), auxiliary space, Big-Θ, Big-Ω, output-sensitive complexity
- ⬜ **1.4 Complexity vs real performance** — constant factors, cache locality, heap allocations, small-n effects, `testing.B`, `b.ReportAllocs`, `benchstat`, escape analysis (awareness)

🛠 **Practice** — L1: take 5 functions from `code/algorithms-ds/ds/` and write their time/space complexity as a comment, then explain each to a rubber duck · L2: benchmark summing 1M ints in a slice vs a linked list with `testing.B` and explain the gap with cache locality

---

## DSA 2 — Arrays & strings
**L1 — aware & can use**
- ✅ **2.1 Array / slice operations** — random access, insertion/deletion cost, dynamic array, in-place, append, iteration, index out of range
- ⬜ **2.2 Strings** — immutability, character counting, anagram, palindrome, string builder, `strings.Builder`, `[]byte` vs `[]rune`, UTF-8, `strconv`
- ⬜ **2.3 Problems (basic)** — [ ] reverse array in place · [ ] move zeros · [ ] find duplicates · [ ] rotate array

**L2 — deep & efficient**
- ⬜ **2.4 Slice internals & efficiency** — backing array, `len` / `cap`, growth strategy, aliasing after `append`, `copy`, preallocation `make([]T, 0, n)`, memory retention via sub-slice, `slices` package, full slice expression `a[lo:hi:max]`
- ⬜ **2.5 Problems (harder)** — [ ] Kadane's max subarray · [ ] product of array except self

🛠 **Practice** — L1: implement `Reverse`, `MoveZeros`, `Rotate` in place with table-driven tests · L2: write a test that shows two slices sharing one backing array after `append`, then fix it with `copy` or a full slice expression

---

## DSA 3 — Hashing
**L1 — aware & can use**
- ✅ **3.1 Hash map & hash set** — hash function, bucket, collision, chaining, open addressing, load factor, rehashing
- ⬜ **3.2 Go maps in practice** — `map[K]V`, comma-ok idiom, `delete`, random iteration order, `map[T]struct{}` as a set, comparable keys, nil map write panic, `maps` package, `clear`
- ⬜ **3.3 Problems (basic)** — [ ] two sum · [ ] valid anagram · [ ] first non-repeating char · [ ] word frequency · [ ] group anagrams

**L2 — deep & efficient**
- ⬜ **3.4 Hash table internals** — Swiss Table (Go 1.24+), groups, control bytes, probing, incremental growth, hash seed randomization, hash flooding, worst case O(n), not safe for concurrent writes, `sync.Map`
- ⬜ **3.5 Problems (harder)** — [ ] longest consecutive sequence · [ ] 4Sum II (count tuples)

🛠 **Practice** — L1: solve two sum and group anagrams with a map, and say the complexity out loud before coding · L2: write a toy hash map with chaining and resize at load factor 0.75, then benchmark it against the built-in map

---

## DSA 4 — Linked lists
**L1 — aware & can use**
- ✅ **4.1 Concepts** — node, head, tail, singly vs doubly linked, sentinel/dummy node, cache unfriendliness
- ✅ **4.2 Problems** — [x] reverse list · [x] middle node · [x] cycle (hash set) · [x] cycle (Floyd) · [x] nth from end · [x] merge two sorted · [ ] remove duplicates · [ ] palindrome list

**L2 — deep & efficient**
- ⬜ **4.3 Advanced techniques** — dummy-node merging, reversing a sub-list, fast/slow split, k-group reversal, Floyd cycle-start proof, `container/list`, pointer-chasing cost vs slices
- ⬜ **4.4 Problems (harder)** — [ ] reverse nodes in k-group · [ ] reorder list · [ ] copy list with random pointer · [ ] add two numbers · [ ] sort list (merge sort)

🛠 **Practice** — L1: finish remove duplicates and palindrome list in `code/algorithms-ds/ds/ds_imp/` · L2: implement sort list with merge sort in O(n log n) time and O(1) extra space (no recursion stack beyond log n)

---

## DSA 5 — Stacks & queues
**L1 — aware & can use**
- ✅ **5.1 Concepts** — LIFO, FIFO, push/pop, enqueue/dequeue, deque, circular buffer, monotonic stack
- ✅ **5.2 Problems** — [x] valid parentheses · [x] min stack · [x] daily temperatures · [ ] queue using two stacks · [ ] evaluate reverse Polish notation · [ ] next greater element

**L2 — deep & efficient**
- ⬜ **5.3 Implementation trade-offs** — slice-backed stack, ring buffer queue, memory leak when dequeuing by re-slicing, buffered channel as a bounded queue, `container/list`, monotonic deque, amortized cost
- ⬜ **5.4 Problems (harder)** — [ ] sliding window maximum · [ ] largest rectangle in histogram · [ ] basic calculator · [ ] design circular queue

🛠 **Practice** — L1: finish queue using two stacks and RPN evaluation · L2: implement a generic ring buffer `Ring[T]` with fixed capacity and compare its allocations with a re-slicing queue using `b.ReportAllocs`

---

## DSA 6 — Trees
**L1 — aware & can use**
- ✅ **6.1 Binary tree & traversals** — root, leaf, height, depth, inorder, preorder, postorder, level order (BFS)
- ⬜ **6.2 BST** — BST property, search/insert/delete, successor, inorder = sorted, balanced vs unbalanced, AVL / red-black (awareness)
- 🔄 **6.3 Problems (basic)** — [x] all traversals · [ ] max depth · [ ] invert tree · [ ] balanced tree · [ ] validate BST · [ ] lowest common ancestor · [ ] diameter

**L2 — deep & efficient**
- ⬜ **6.4 Balanced trees & ordered structures** — AVL rotations, red-black properties, O(log n) height guarantee, treap (awareness), skip list, B-tree link to databases, no built-in ordered map in Go, iterative traversal with explicit stack, Morris traversal (awareness)
- ⬜ **6.5 Problems (harder)** — [ ] serialize/deserialize · [ ] binary tree max path sum · [ ] kth smallest in BST · [ ] construct tree from preorder + inorder

🛠 **Practice** — L1: implement BST insert/search/delete with tests, then solve max depth and validate BST · L2: implement serialize/deserialize and an iterative inorder traversal without recursion

---

## DSA 7 — Heaps
**L1 — aware & can use**
- ✅ **7.1 Concepts** — min-heap, max-heap, complete binary tree, array representation, sift up / sift down, heapify O(n)
- ⬜ **7.2 Go `container/heap`** — `heap.Interface`, `sort.Interface`, `Push`, `Pop`, `Init`, `Fix`, `Remove`, pointer receiver for `Push`/`Pop`
- ⬜ **7.3 Problems (basic)** — [ ] kth largest · [ ] top K frequent

**L2 — deep & efficient**
- ⬜ **7.4 Heap patterns & internals** — why heapify is O(n), d-ary heap, priority queue with updates (index tracking + `heap.Fix`), two-heaps pattern, top-K streaming, Go runtime timers (4-ary heap), k-way merge
- ⬜ **7.5 Problems (harder)** — [ ] merge K sorted lists · [ ] median from data stream · [ ] task scheduler

🛠 **Practice** — L1: implement a min-heap by hand (no `container/heap`), then solve kth largest with `container/heap` · L2: build a priority queue of jobs that supports changing priority via `heap.Fix`

---

## DSA 8 — Graphs
**L1 — aware & can use**
- ✅ **8.1 Concepts** — vertex, edge, directed/undirected, weighted, adjacency list vs matrix, visited set, BFS, DFS
- ⬜ **8.2 Traversal patterns** — grid as a graph, 4/8-direction neighbors, connected components, multi-source BFS, BFS distance in unweighted graph, iterative DFS with stack
- ⬜ **8.3 Problems (basic)** — [ ] number of islands · [ ] BFS shortest path · [ ] rotting oranges · [ ] clone graph

**L2 — deep & efficient**
- ⬜ **8.4 Topological sort & cycles** — DAG, in-degree, Kahn's algorithm, DFS coloring (white/gray/black), cycle detection, dependency resolution
- ⬜ **8.5 Shortest paths** — BFS on unweighted graph, Dijkstra, priority queue, relaxation, Bellman-Ford (awareness), negative edges, 0-1 BFS, A* (awareness)
- ⬜ **8.6 Minimum spanning tree** — spanning tree, Kruskal, Prim, union-find, cut property, use cases (network design)
- ⬜ **8.7 Problems (harder)** — [ ] course schedule · [ ] topological sort · [ ] network delay time (Dijkstra) · [ ] word ladder

🛠 **Practice** — L1: solve number of islands with both BFS and DFS and compare · L2: implement Dijkstra with `container/heap` on an adjacency list and test it on a small hand-drawn graph

---

## DSA 9 — Binary search
**L1 — aware & can use**
- ✅ **9.1 Concepts** — sorted input, `lo`/`hi`, `<` vs `<=`, overflow-safe mid, lower bound / upper bound, search on answer space
- ✅ **9.2 Problems** — [x] classic · [x] search insert position · [ ] first & last position

**L2 — deep & efficient**
- ⬜ **9.3 Variants & invariants** — loop invariant, monotonic predicate, binary search on the answer, rotated arrays, `sort.Search`, `slices.BinarySearch`, `slices.BinarySearchFunc`, floating-point bisection
- ⬜ **9.4 Problems (harder)** — [ ] rotated sorted array · [ ] min in rotated array · [ ] Koko eating bananas · [ ] median of two sorted arrays

🛠 **Practice** — L1: write first & last position using one lower-bound helper called twice · L2: solve Koko eating bananas and write down the invariant of `lo`/`hi` before coding

---

## DSA 10 — Two pointers
**L1 — aware & can use**
- ✅ **10.1 Concepts** — opposite-direction pointers, same-direction (fast/slow), in-place partitioning
- ⬜ **10.2 Problems (basic)** — [ ] two sum sorted · [ ] valid palindrome · [ ] remove duplicates in place · [ ] container with most water

**L2 — deep & efficient**
- ⬜ **10.3 Advanced patterns** — sort + two pointers for k-sum, skipping duplicates, Dutch national flag (3 pointers), why discarding is safe (correctness argument), two pointers on two arrays
- ⬜ **10.4 Problems (harder)** — [ ] 3sum · [ ] trapping rain water · [ ] sort colors

🛠 **Practice** — L1: implement the TODOs in `code/algorithms-ds/ds/two_pointers.go` · L2: solve 3sum without duplicates and write one sentence proving why moving the smaller pointer is safe in container with most water

---

## DSA 11 — Sliding window
**L1 — aware & can use**
- ⬜ **11.1 Concepts** — fixed window, variable window, expand/shrink, window state (map/counter), left/right pointers
- ⬜ **11.2 Problems (basic)** — [ ] max sum subarray size k · [ ] longest substring without repeating

**L2 — deep & efficient**
- ⬜ **11.3 Advanced patterns** — "at most K" trick (exactly K = atMost(K) − atMost(K−1)), monotonic deque window, frequency array vs map, negative numbers break the window, "need/have" counters
- ⬜ **11.4 Problems (harder)** — [ ] minimum window substring · [ ] longest subarray sum ≤ k · [ ] longest repeating character replacement · [ ] subarrays with K different integers

🛠 **Practice** — L1: solve both basic problems with one reusable window template · L2: solve minimum window substring and explain when the window shrinks

---

## DSA 12 — Prefix sums & intervals
**L1 — aware & can use**
- ⬜ **12.1 Prefix sums** — prefix array, range sum query, off-by-one (`n+1` array), cumulative sum, O(1) query after O(n) build
- ⬜ **12.2 Intervals** — sort by start, overlap condition, merge intervals, insert interval, meeting rooms (heap)
- ⬜ **12.3 Problems (basic)** — [ ] range sum query · [ ] merge intervals · [ ] insert interval

**L2 — deep & efficient**
- ⬜ **12.4 Advanced techniques** — subarray sum equals k (prefix + hash map), 2D prefix sums, difference array, prefix XOR, sweep line, events sorting, min-heap of end times
- ⬜ **12.5 Problems (harder)** — [ ] subarray sum equals k · [ ] meeting rooms II · [ ] non-overlapping intervals

🛠 **Practice** — L1: implement `NumArray` range sum and merge intervals with tests for edge cases (empty, touching intervals) · L2: solve meeting rooms II with both a sweep line and a heap and compare

---

## DSA 13 — Recursion & backtracking
**L1 — aware & can use**
- ⬜ **13.1 Recursion** — base case, recursive case, call stack, stack overflow, tail recursion (Go does not optimize), recursion vs iteration
- ⬜ **13.2 Backtracking** — choose → explore → unchoose, pruning, decision tree, state, result collection
- ⬜ **13.3 Problems (basic)** — [ ] subsets · [ ] permutations · [ ] combination sum

**L2 — deep & efficient**
- ⬜ **13.4 Advanced backtracking** — handling duplicates (sort + skip), copying the path before saving (Go slice aliasing bug), pruning by bounds, goroutine stack growth, explicit stack instead of recursion, exponential complexity estimates
- ⬜ **13.5 Problems (harder)** — [ ] N-queens · [ ] word search · [ ] subsets II · [ ] palindrome partitioning

🛠 **Practice** — L1: write subsets and permutations with one backtracking template · L2: reproduce the "all results are the same slice" bug by appending `path` without copying, then fix it

---

## DSA 14 — Sorting
**L1 — aware & can use**
- ⬜ **14.1 Sorting in Go** — `slices.Sort`, `slices.SortFunc`, `slices.SortStableFunc`, `sort.Slice`, `sort.SliceStable`, `cmp.Compare`, sort by multiple keys, stability
- ⬜ **14.2 Classic algorithms** — bubble/insertion (why slow), merge sort, quicksort, pivot, heap sort, in-place, stability, Go's `slices.Sort` (pdqsort)

**L2 — deep & efficient**
- ⬜ **14.3 Advanced sorting** — Ω(n log n) lower bound for comparison sorts, counting sort, radix sort, bucket sort, quickselect, worst-case quicksort, introsort / pdqsort, external sort (merging sorted runs)
- ⬜ **14.4 Problems (harder)** — [ ] implement merge sort · [ ] implement quicksort · [ ] kth largest via quickselect · [ ] largest number

🛠 **Practice** — L1: sort a slice of `User` by age then name with `slices.SortFunc` + `cmp.Compare` · L2: implement merge sort and quicksort, then benchmark both against `slices.Sort` on random and already-sorted input

---

## DSA 15 — Dynamic programming
**L1 — aware & can use**
- ⬜ **15.1 Concepts** — overlapping subproblems, optimal substructure, memoization (top-down), tabulation (bottom-up), state, transition, base case
- ⬜ **15.2 Problems (basic)** — [ ] fibonacci · [ ] climbing stairs · [ ] house robber · [ ] coin change

**L2 — deep & efficient**
- ⬜ **15.3 DP patterns** — 1D DP, 2D grid DP, 0/1 knapsack, unbounded knapsack, two-string DP (LCS, edit distance), LIS (O(n log n)), interval DP, space optimization (rolling array), bitmask DP (awareness)
- ⬜ **15.4 Problems (harder)** — [ ] longest common subsequence · [ ] 0/1 knapsack · [ ] edit distance · [ ] longest increasing subsequence · [ ] unique paths · [ ] word break

🛠 **Practice** — L1: solve climbing stairs three ways (plain recursion, memo, table) and benchmark them · L2: solve LCS with a 2D table, then reduce memory to two rows

---

## DSA 16 — Greedy
**L1 — aware & can use**
- ⬜ **16.1 Concepts** — greedy choice, local vs global optimum, sort first, "best so far" tracking, when greedy fails
- ⬜ **16.2 Problems (basic)** — [ ] jump game · [ ] gas station

**L2 — deep & efficient**
- ⬜ **16.3 Proving greedy** — exchange argument, greedy-stays-ahead, counterexample (coin change with arbitrary coins), interval scheduling, activity selection, Huffman coding (awareness)
- ⬜ **16.4 Problems (harder)** — [ ] jump game II · [ ] partition labels · [ ] minimum arrows to burst balloons

🛠 **Practice** — L1: solve jump game greedily and with DP and compare complexity · L2: find a coin set where greedy coin change fails and explain why DP is needed

---

## DSA 17 — Advanced structures
**L1 — aware & can use**
- ⬜ **17.1 Trie** — prefix tree, node children (`[26]*Node` vs map), insert/search/startsWith, end-of-word flag, autocomplete
- ⬜ **17.2 Union-find** — disjoint set, find, union, path compression, union by rank, connected components
- ⬜ **17.3 LRU cache** — hash map + doubly linked list, O(1) get/put, eviction, capacity (classic interview task)

**L2 — deep & efficient**
- ⬜ **17.4 Range & cache structures** — segment tree, Fenwick tree (BIT), lazy propagation (awareness), sparse table (awareness), LFU cache, thread-safe LRU with `sync.Mutex`, sharded cache
- ⬜ **17.5 Problems (harder)** — [ ] implement LRU cache · [ ] implement trie · [ ] number of connected components (union-find) · [ ] word search II · [ ] range sum query mutable (Fenwick)

🛠 **Practice** — L1: implement LRU cache with `map` + your own doubly linked list and tests for eviction order · L2: make the LRU concurrency-safe, run it under `go test -race`, then benchmark mutex vs sharded versions

---

## DSA 18 — Bit manipulation
**L1 — aware & can use**
- ⬜ **18.1 Basics** — AND/OR/XOR/NOT, shifts, masks, `x & (x-1)`, single number (XOR trick), counting bits
- ⬜ **18.2 Problems (basic)** — [ ] single number · [ ] number of 1 bits · [ ] counting bits · [ ] power of two

**L2 — deep & efficient**
- ⬜ **18.3 Advanced bits** — two's complement, signed vs unsigned shifts in Go, `&^` (AND NOT), `math/bits` (`OnesCount`, `TrailingZeros`, `LeadingZeros`), bitsets, enumerating subsets with masks, flags with `iota`
- ⬜ **18.4 Problems (harder)** — [ ] missing number · [ ] reverse bits · [ ] subsets via bitmask

🛠 **Practice** — L1: implement `IsPowerOfTwo` and `CountBits` and compare with `math/bits.OnesCount` · L2: build a `Permissions` flag type with `iota` and `Set`/`Clear`/`Has` methods

---

## DSA 19 — Structures inside real systems (connects to PG/DIST)
**L1 — aware & can use**
- ⬜ **19.1 Where structures live in your stack** — Go map (hash table), Go timers (heap), goroutine run queue (ring buffer), Postgres index (B-tree), Redis sorted set (skip list), Kafka log (append-only log), HTTP router (radix tree), LRU caches

**L2 — deep & efficient**
- ⬜ **19.2 Storage structures** — B-tree, B+tree, page split, LSM-tree, SSTable, memtable, compaction, write/read amplification, skip list
- ⬜ **19.3 Probabilistic structures** — Bloom filter, false positive rate, HyperLogLog, count-min sketch, cardinality estimation, memory vs accuracy trade-off
- ⬜ **19.4 Distributed structures** — consistent hashing ring, virtual nodes, Merkle tree, anti-entropy, ring buffer, CRDT (awareness)

🛠 **Practice** — L1: for each structure in 19.1, write one sentence on why that system chose it · L2: implement a Bloom filter in Go and measure the real false positive rate vs the formula

---

## DSA 20 — Interview problem-solving method
**L1 — aware & can use**
- ⬜ **20.1 Method** — clarify input/output, examples & edge cases, brute force first, optimize, dry run, complexity statement, thinking out loud
- ⬜ **20.2 Pattern recognition** — "sorted" → binary search / two pointers, "top K" → heap, "contiguous subarray" → sliding window / prefix sums, "all combinations" → backtracking, "min/max/count ways" → DP, "connected" → BFS / DFS / union-find, "O(1) lookup" → hash map

**L2 — deep & efficient**
- ⬜ **20.3 Performing under pressure** — 45-minute timed practice, mock interviews, testing code by hand, recovering from mistakes, follow-up questions (scale, streaming input, memory limits), clean Go style in interviews, communicating trade-offs

🛠 **Practice** — L1: solve one unseen problem while speaking every step out loud (record it) · L2: run a timed 45-minute mock with 2 problems and follow-up questions, then write what went wrong
