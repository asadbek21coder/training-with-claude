# 01 — Data Structures

> For every data structure: know what it is, how it works internally, all operations + complexity, and when to use it.

---

## Array / Slice

**Key Terms:**
- `index` — position of element, zero-based
- `random access` — O(1) read/write by index (direct memory address)
- `contiguous memory` — elements stored next to each other → cache-friendly
- `dynamic array` — array that resizes automatically (Go slice, Java ArrayList)
- `backing array` — the underlying fixed array a slice points into
- `len` — number of elements currently in the slice
- `cap` — total capacity of backing array before reallocation
- `amortized O(1) append` — occasionally O(n) copy, but averages to O(1) over many appends
- `resize factor` — Go doubles capacity; Java ArrayList grows by 50%
- `slice of slice` — shares backing array; mutations affect the original
- `copy()` — creates an independent copy of a slice in Go
- `cache locality` — arrays are fast because CPU prefetches contiguous memory

**Complexity:**

| Operation | Time | Notes |
|-----------|------|-------|
| Access by index | O(1) | |
| Search (unsorted) | O(n) | |
| Search (sorted) | O(log n) | binary search |
| Insert at end | O(1) amortized | |
| Insert at middle | O(n) | shifts elements |
| Delete at middle | O(n) | shifts elements |
| Delete at end | O(1) | |

**Must understand:**
- [ ] Why append is O(1) amortized but sometimes O(n)
- [ ] What happens when you do `s2 := s1[2:5]` — do they share memory?
- [ ] Difference between `len` and `cap`
- [ ] Why arrays are cache-friendly compared to linked lists

---

## Hash Map / Hash Set

**Key Terms:**
- `hash function` — maps a key to a bucket index
- `bucket` — slot in the underlying array that holds key-value pairs
- `collision` — two different keys hash to the same bucket
- `chaining` — collision resolved by linked list of entries in same bucket
- `open addressing` — collision resolved by probing for next empty slot
- `linear probing` — check next slot, then next, until empty found
- `load factor` — ratio of entries to buckets (e.g., 0.75 triggers resize)
- `rehashing` — when load factor exceeded: new larger array, re-insert all
- `O(1) average` — assumes low collision rate; O(n) worst case with many collisions
- `hash set` — hash map where you only care about keys, not values
- `Go map` — built-in hash map, not safe for concurrent reads+writes
- `sync.Map` — Go's concurrent map (slower but thread-safe)
- `Java HashMap` — default, not thread-safe
- `Java ConcurrentHashMap` — thread-safe, uses segment locking

**Complexity:**

| Operation | Average | Worst |
|-----------|---------|-------|
| Get | O(1) | O(n) |
| Put | O(1) | O(n) |
| Delete | O(1) | O(n) |
| Contains | O(1) | O(n) |

**Must understand:**
- [ ] What causes a collision and how chaining resolves it
- [ ] Why load factor matters and what happens when it's exceeded
- [ ] Why iteration order of a map is not guaranteed
- [ ] Why you cannot use a slice as a map key in Go (not comparable)
- [ ] When to use a map vs a slice (lookup vs ordered iteration)

---

## Singly Linked List

**Key Terms:**
- `node` — struct with value + pointer to next node
- `head` — pointer to first node
- `tail` — pointer to last node (optional; needed for O(1) append)
- `next` — pointer to the next node
- `nil / null` — marks end of list
- `traversal` — walk node by node following `next` pointers
- `sentinel / dummy node` — fake head node that never holds data; simplifies edge cases
- `insertion at head` — O(1): new node → new node.next = old head → head = new node
- `insertion at tail` — O(1) with tail pointer, O(n) without
- `deletion by value` — O(n): find predecessor, rewire its next
- `runner technique` — two pointers at different speeds (slow/fast)

**Complexity:**

| Operation | Time |
|-----------|------|
| Access by index | O(n) |
| Insert at head | O(1) |
| Insert at tail | O(1) with tail / O(n) without |
| Insert at index | O(n) |
| Delete at head | O(1) |
| Delete by value | O(n) |
| Search | O(n) |

**Must understand:**
- [ ] Why you need a dummy node to simplify deletion
- [ ] How to reverse a linked list in-place (3-pointer technique: prev, curr, next)
- [ ] How to detect a cycle (Floyd's algorithm: slow + fast pointer)
- [ ] How to find middle of list (slow + fast pointer)
- [ ] Why linked list has poor cache performance compared to array

---

## Doubly Linked List

**Key Terms:**
- `prev` — pointer to previous node
- `next` — pointer to next node
- `bidirectional traversal` — can walk forward and backward
- `O(1) delete with pointer` — if you have a pointer to a node, delete in O(1) by rewiring prev/next
- `sentinel head + tail` — dummy nodes at both ends → never need to check for nil
- `deque (double-ended queue)` — insert/delete from both ends O(1)
- `LRU Cache` — classic use: doubly linked list + hash map

**Complexity:**

| Operation | Time |
|-----------|------|
| Insert at head | O(1) |
| Insert at tail | O(1) |
| Delete given pointer | O(1) |
| Delete by value | O(n) search + O(1) delete |
| Access by index | O(n) |

**Must understand:**
- [ ] How deleting a node works — rewire node.prev.next and node.next.prev
- [ ] Why dummy head and tail nodes eliminate nil checks
- [ ] How DLL + HashMap gives O(1) get and put for LRU Cache

---

## Stack

**Key Terms:**
- `LIFO` — Last In, First Out
- `push` — add to top
- `pop` — remove from top, return value
- `peek / top` — look at top without removing
- `isEmpty` — check if empty before pop
- `call stack` — OS uses stack to track function calls (stack overflow = too deep)
- `monotonic stack` — stack that maintains elements in sorted order; used for "next greater element"
- `balanced parentheses` — classic stack problem
- `implementation` — use a slice (dynamic array) or linked list

**Complexity:**

| Operation | Time |
|-----------|------|
| Push | O(1) |
| Pop | O(1) |
| Peek | O(1) |
| Search | O(n) |

**Must understand:**
- [ ] Implement stack using a slice in Go
- [ ] When to use a monotonic stack (next greater/smaller element problems)
- [ ] How the call stack works (why recursion too deep = stack overflow)
- [ ] Why stack is perfect for DFS traversal (mirrors the recursion call stack)

---

## Queue

**Key Terms:**
- `FIFO` — First In, First Out
- `enqueue` — add to back (tail)
- `dequeue` — remove from front (head)
- `peek / front` — look at front without removing
- `circular queue` — fixed-size array where head/tail wrap around
- `deque` — double-ended queue; insert/delete from both ends
- `monotonic deque` — deque with sorted order; used for sliding window maximum
- `priority queue` — queue ordered by priority, not insertion order (implemented with heap)
- `linked list implementation` — head = dequeue side (O(1)), tail = enqueue side (O(1))

**Complexity:**

| Operation | Time |
|-----------|------|
| Enqueue | O(1) |
| Dequeue | O(1) |
| Peek | O(1) |
| Search | O(n) |

**Must understand:**
- [ ] Why linked list is better than slice for queue (no O(n) shift on dequeue)
- [ ] How circular queue avoids shifting with two pointers (head, tail)
- [ ] Queue is used for BFS — why? (processes nodes level by level)
- [ ] Difference between Queue and Priority Queue

---

## Binary Tree

**Key Terms:**
- `node` — value + left child + right child
- `root` — top node, no parent
- `leaf` — node with no children
- `height` — max depth from root to leaf (empty tree = -1 or 0, varies)
- `depth` — distance from root to that node
- `balanced` — height difference between left/right subtrees ≤ 1
- `complete tree` — all levels full except last, last filled left to right
- `perfect tree` — all levels completely filled
- `inorder traversal` — Left → Root → Right (gives sorted order for BST)
- `preorder traversal` — Root → Left → Right (useful for copying/serializing)
- `postorder traversal` — Left → Right → Root (useful for deletion)
- `level-order / BFS` — visit nodes level by level using a queue
- `DFS` — go deep before going wide (uses recursion or explicit stack)
- `diameter` — longest path between any two nodes (may not pass through root)

**Complexity:**

| Operation | Balanced | Unbalanced (skewed) |
|-----------|----------|---------------------|
| Insert | O(log n) | O(n) |
| Search | O(log n) | O(n) |
| Traversal | O(n) | O(n) |
| Height | O(n) | O(n) |

**Must understand:**
- [ ] Implement inorder, preorder, postorder recursively
- [ ] Implement BFS using a queue
- [ ] Calculate height recursively
- [ ] Difference between height and depth
- [ ] What makes a tree balanced

---

## Binary Search Tree (BST)

**Key Terms:**
- `BST property` — for every node: left subtree values < node < right subtree values
- `inorder = sorted` — inorder traversal of BST gives values in ascending order
- `search` — go left if target < node, right if target > node, O(h)
- `insert` — same logic as search, insert at correct leaf
- `delete` — three cases: leaf (just remove), one child (replace), two children (replace with inorder successor)
- `inorder successor` — smallest node in right subtree
- `inorder predecessor` — largest node in left subtree
- `skewed tree` — BST degenerates to linked list when inserting sorted data → O(n) operations
- `self-balancing BST` — AVL tree, Red-Black tree — maintain O(log n) guarantee
- `AVL tree` — balances by rotation after every insert/delete; height ≤ 1.44 log n
- `Red-Black tree` — used in Java TreeMap, C++ std::map — looser balance, fewer rotations

**Must understand:**
- [ ] Why inserting sorted data into a BST degenerates to O(n)
- [ ] How to delete a node with two children (replace with inorder successor)
- [ ] Why BST inorder traversal gives sorted output
- [ ] What problem AVL and Red-Black trees solve

---

## Heap / Priority Queue

**Key Terms:**
- `heap` — complete binary tree with heap property
- `min-heap` — parent ≤ children; root is minimum
- `max-heap` — parent ≥ children; root is maximum
- `heap property` — every parent is ≤ (min) or ≥ (max) its children
- `array representation` — node at index i: left child at 2i+1, right child at 2i+2, parent at (i-1)/2
- `heapify-up (sift-up)` — after insert at end, bubble up until heap property restored
- `heapify-down (sift-down)` — after remove root, put last element at root, bubble down
- `heapify` — build heap from array in O(n) (not O(n log n))
- `heapsort` — O(n log n) sort using heap
- `priority queue` — abstract data type; heap is the implementation
- `Go heap` — use `container/heap` package (implements heap.Interface)
- `Java PriorityQueue` — min-heap by default; reverse comparator for max-heap
- `k-th largest element` — classic heap interview problem
- `top K elements` — maintain heap of size K

**Complexity:**

| Operation | Time |
|-----------|------|
| Insert | O(log n) |
| Remove min/max | O(log n) |
| Peek min/max | O(1) |
| Build heap from array | O(n) |
| Heapsort | O(n log n) |

**Must understand:**
- [ ] How array represents a complete binary tree (index formulas)
- [ ] How sift-up and sift-down work
- [ ] Why building a heap is O(n) not O(n log n)
- [ ] When to use min-heap vs max-heap for "top K" problems
- [ ] Implement a min-heap from scratch

---

## Trie (Prefix Tree)

**Key Terms:**
- `trie` — tree where each path from root to node represents a string prefix
- `node` — contains children (map or fixed array of 26) + isEnd flag
- `isEnd / endOfWord` — marks end of a valid word
- `insert` — walk character by character, create nodes as needed
- `search` — walk character by character, return isEnd at last char
- `startsWith` — walk prefix characters; return true if path exists (no isEnd check)
- `children` — map[rune]*TrieNode or [26]*TrieNode for lowercase letters
- `autocomplete` — DFS from prefix node to find all words
- `time complexity` — O(m) per operation where m = length of word
- `space complexity` — O(total characters across all words)

**Use cases:** autocomplete, spell check, word search, IP routing, prefix matching

**Must understand:**
- [ ] Implement Trie with Insert, Search, StartsWith
- [ ] Why Trie is better than hash map for prefix queries
- [ ] How to autocomplete using DFS from a prefix node
- [ ] The isEnd flag — what happens if you search "app" when only "apple" is inserted

---

## Graph

**Key Terms:**
- `vertex / node` — entity in the graph
- `edge` — connection between two vertices
- `directed graph (digraph)` — edges have direction (A→B ≠ B→A)
- `undirected graph` — edges go both ways (A—B)
- `weighted graph` — edges have cost/distance
- `unweighted graph` — edges have no cost
- `adjacency list` — map/array where each vertex stores its neighbors — O(V+E) space
- `adjacency matrix` — V×V matrix where [i][j]=1 means edge — O(V²) space
- `sparse graph` — few edges → use adjacency list
- `dense graph` — many edges → adjacency matrix may be fine
- `in-degree` — number of edges coming INTO a node (directed graphs)
- `out-degree` — number of edges going OUT of a node
- `connected component` — group of vertices all reachable from each other
- `cycle` — path that starts and ends at same vertex
- `DAG (Directed Acyclic Graph)` — directed graph with no cycles; used for dependencies
- `tree` — connected undirected graph with no cycles (N nodes, N-1 edges)
- `bipartite graph` — vertices can be split into 2 groups; no edge within same group

**Must understand:**
- [ ] When to use adjacency list vs adjacency matrix
- [ ] How to represent a graph in Go (map[int][]int)
- [ ] BFS on graph (shortest path in unweighted)
- [ ] DFS on graph (connected components, cycle detection)
- [ ] Why you need a `visited` set in graph traversal (to avoid infinite loops)

---

## Segment Tree

**Key Terms:**
- `range query` — query over a range [l, r] (sum, min, max)
- `point update` — update a single index
- `segment tree node` — represents a range [l, r]; stores aggregate value
- `build` — construct tree bottom-up in O(n)
- `query` — O(log n): split range into O(log n) nodes
- `update` — O(log n): update node and all ancestors
- `lazy propagation` — defer range updates to avoid O(n) propagation; needed for range updates
- `array representation` — tree stored in array; node i has children 2i and 2i+1

**Complexity:**

| Operation | Time |
|-----------|------|
| Build | O(n) |
| Query [l, r] | O(log n) |
| Point update | O(log n) |
| Range update (lazy) | O(log n) |

**Must understand:**
- [ ] When to use segment tree vs prefix sum (prefix sum: updates O(n), seg tree: O(log n))
- [ ] Basic implementation: build, query, update
- [ ] What lazy propagation solves

---

## Union-Find (Disjoint Set Union)

**Key Terms:**
- `find(x)` — returns the root/representative of x's group
- `union(x, y)` — merge the groups containing x and y
- `connected(x, y)` — returns true if x and y are in the same group
- `path compression` — on find, make all nodes point directly to root → nearly O(1)
- `union by rank` — always attach smaller tree under larger → prevents tall trees
- `inverse Ackermann O(α(n))` — effectively O(1) per operation with both optimizations
- `component count` — track number of distinct sets; decrement on each union

**Use cases:** connected components, Kruskal's MST algorithm, cycle detection in undirected graphs, network connectivity

**Must understand:**
- [ ] Implement find with path compression
- [ ] Implement union by rank
- [ ] Why it's nearly O(1) per operation
- [ ] How to count connected components using Union-Find

---

## LRU Cache

**Key Terms:**
- `LRU` — Least Recently Used; evict item not accessed for longest time
- `get(key)` — O(1): return value; move to most recent position
- `put(key, val)` — O(1): insert; if at capacity, evict LRU item first
- `doubly linked list` — maintain order of access; O(1) move-to-front and remove-from-back
- `hash map` — map key → DLL node for O(1) lookup
- `dummy head` — sentinel at front (most recent side)
- `dummy tail` — sentinel at back (LRU side)
- `capacity` — max number of items; evict tail.prev when exceeded
- `move to front` — on every get/put, remove node and re-insert after dummy head

**Complexity:**

| Operation | Time |
|-----------|------|
| get | O(1) |
| put | O(1) |

**Must understand:**
- [ ] Why you need BOTH a DLL and a hash map
- [ ] How get moves a node to front in O(1)
- [ ] How put evicts LRU in O(1)
- [ ] Implement LRU Cache from scratch (most common hard interview problem)
