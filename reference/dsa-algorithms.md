# 02 — Algorithms

> For every pattern: know when to apply it, what the time/space complexity is, and implement it from memory.

---

## Big O Notation

**Key Terms:**
- `time complexity` — how runtime grows as input size n grows
- `space complexity` — how memory usage grows as n grows
- `O(1)` — constant; doesn't grow with input (hash map lookup, array index)
- `O(log n)` — logarithmic; halves problem each step (binary search)
- `O(n)` — linear; visits every element once (linear search)
- `O(n log n)` — linearithmic; sorting (merge sort, heap sort)
- `O(n²)` — quadratic; nested loops (bubble sort, brute force pairs)
- `O(2^n)` — exponential; all subsets (brute force recursion)
- `O(n!)` — factorial; all permutations (brute force TSP)
- `best case` — fastest scenario (already sorted array for bubble sort)
- `worst case` — slowest scenario (reversed array for quicksort)
- `average case` — expected runtime over random input
- `amortized` — averaged over a sequence of operations (append to slice)
- `dominant term` — in O(n² + n), the n² dominates → O(n²)
- `drop constants` — O(2n) → O(n); O(100) → O(1)
- `auxiliary space` — extra space used by the algorithm (not counting input)

**Big O Cheat Sheet:**

| Complexity | Name | Example |
|------------|------|---------|
| O(1) | Constant | Hash map get, array index |
| O(log n) | Logarithmic | Binary search |
| O(n) | Linear | Loop through array |
| O(n log n) | Linearithmic | Merge sort, heap sort |
| O(n²) | Quadratic | Nested loops, bubble sort |
| O(2^n) | Exponential | All subsets, naive Fibonacci |
| O(n!) | Factorial | All permutations |

**Must understand:**
- [ ] Calculate Big O for any given loop structure
- [ ] Why O(log n) means "problem halves each step"
- [ ] Amortized O(1) for append — explain why
- [ ] Space complexity of recursive algorithms (call stack counts as space)

---

## Binary Search

**Key Terms:**
- `sorted array` — binary search only works on sorted input
- `left pointer` — starts at index 0
- `right pointer` — starts at len(arr)-1
- `mid` — `left + (right-left)/2` — avoids integer overflow vs `(left+right)/2`
- `search space reduction` — each step eliminates half the remaining candidates
- `left boundary` — find leftmost index where condition is true
- `right boundary` — find rightmost index where condition is true
- `search on answer` — binary search on the value range, not array index (e.g., "minimum capacity to ship")
- `condition function` — convert problem to "find first/last index where f(x) is true"

**Variants:**
1. **Standard** — find exact target; return -1 if not found
2. **Left boundary** — find first index where arr[i] >= target
3. **Right boundary** — find last index where arr[i] <= target
4. **Search on answer** — binary search on answer space, not array

**Complexity:**

| | Time | Space |
|--|------|-------|
| Standard | O(log n) | O(1) |

**Must understand:**
- [ ] Why `mid = left + (right-left)/2` instead of `(left+right)/2`
- [ ] When to use `left <= right` vs `left < right` as loop condition
- [ ] When to move `left = mid+1` vs `left = mid`
- [ ] Apply binary search to non-trivial problems (rotated array, search on answer)

---

## Two Pointers

**Key Terms:**
- `opposite pointers` — left starts at 0, right at end; move toward each other
- `same direction pointers` — slow and fast; fast moves ahead
- `slow / fast pointer` — fast moves 2x; used for cycle detection, finding middle
- `Floyd's cycle detection` — if slow and fast meet, there's a cycle
- `three-sum pattern` — fix one element, use two pointers on the rest
- `palindrome check` — compare arr[left] and arr[right], move inward
- `sorted array prerequisite` — most two-pointer problems require sorted input
- `in-place` — two pointers allow O(1) extra space

**Common patterns:**
- Remove duplicates from sorted array (slow/fast)
- Two Sum on sorted array (left/right)
- Three Sum (sort + fix + two-pointer)
- Palindrome check (left/right toward center)
- Linked list cycle detection (slow/fast)
- Find middle of linked list (slow/fast, fast reaches end when slow is at middle)

**Must understand:**
- [ ] When to use opposite-end pointers vs same-direction
- [ ] How Floyd's algorithm detects cycle AND finds cycle start
- [ ] Why sorted input makes two pointers work (decision: move left or right)

---

## Sliding Window

**Key Terms:**
- `window` — contiguous subarray or substring
- `fixed-size window` — window size k doesn't change
- `variable-size window` — expand right, shrink left when constraint violated
- `expand right` — include arr[right] into window
- `shrink left` — remove arr[left] from window, advance left
- `constraint` — the condition the window must satisfy (e.g., no duplicate chars)
- `max/min window` — track best result as window moves
- `hash map / counter` — often used inside window to count character frequencies

**Fixed window pattern:**
```
for right = 0; right < n; right++:
    add arr[right] to window
    if window size > k:
        remove arr[left] from window
        left++
    update result
```

**Variable window pattern:**
```
for right = 0; right < n; right++:
    add arr[right] to window
    while window violates constraint:
        remove arr[left] from window
        left++
    update result (right - left + 1)
```

**Common problems:**
- Longest substring without repeating characters
- Minimum window substring
- Maximum sum subarray of size k
- Fruit into baskets (at most 2 distinct)

**Must understand:**
- [ ] Difference between fixed and variable window
- [ ] When shrinking from the left is correct
- [ ] How to maintain a counter/map inside the window
- [ ] Why sliding window is O(n) even though there's a while loop inside (amortized — each element enters and exits at most once)

---

## BFS (Breadth-First Search)

**Key Terms:**
- `queue` — BFS uses a queue to process nodes level by level
- `visited set` — prevent revisiting nodes (crucial in graphs, optional in trees)
- `level` — all nodes at the same distance from source
- `shortest path` — BFS guarantees shortest path in unweighted graphs
- `level order traversal` — BFS on binary tree visits nodes level by level
- `multi-source BFS` — start BFS from multiple sources simultaneously (e.g., "01 matrix" problem)
- `bidirectional BFS` — search from both source and target; faster for shortest path

**BFS template:**
```
queue = [start]
visited = {start}
while queue not empty:
    node = queue.pop_front()
    process(node)
    for neighbor in neighbors(node):
        if neighbor not in visited:
            visited.add(neighbor)
            queue.append(neighbor)
```

**When to use BFS:**
- Shortest path in unweighted graph
- Level-order traversal
- Find all nodes at distance k
- Multi-source spread (fire, rot, water)

**Must understand:**
- [ ] Why BFS guarantees shortest path (processes nodes by increasing distance)
- [ ] When to add to visited — when enqueued, not when dequeued (avoids duplicates)
- [ ] Multi-source BFS — start with all sources in queue simultaneously
- [ ] BFS space is O(width of tree) — can be O(n) for last level

---

## DFS (Depth-First Search)

**Key Terms:**
- `recursion` — natural implementation; call stack acts as the stack
- `explicit stack` — iterative DFS using a stack data structure
- `visited set` — prevent revisiting in graphs
- `backtracking` — undo last choice and try next option
- `preorder` — process node before children
- `postorder` — process node after children (useful for bottom-up problems)
- `connected components` — count separate groups of connected nodes
- `topological sort` — DFS-based; postorder gives reverse topo order
- `cycle detection` — in directed graph: white/grey/black coloring

**DFS template (recursive):**
```
func dfs(node, visited):
    visited.add(node)
    process(node)
    for neighbor in neighbors(node):
        if neighbor not in visited:
            dfs(neighbor, visited)
```

**When to use DFS:**
- Explore all paths (backtracking)
- Connected components
- Cycle detection
- Tree traversals
- Topological sort

**Must understand:**
- [ ] DFS space is O(depth) = O(h) for trees, O(V) worst case for graphs
- [ ] How backtracking extends DFS (mark/unmark on the way in/out)
- [ ] DFS vs BFS: DFS goes deep, BFS goes wide; BFS finds shortest path

---

## Sorting Algorithms

**Key Terms:**
- `stable sort` — equal elements maintain original relative order (Merge Sort, Timsort)
- `unstable sort` — equal elements may be reordered (Quick Sort, Heap Sort)
- `in-place` — sorts without extra memory O(1) space (Quick Sort, Heap Sort)
- `comparison-based` — compares elements pairwise; lower bound O(n log n)
- `pivot` — element used to partition in QuickSort
- `partition` — rearrange so all < pivot are left, all > pivot are right
- `divide and conquer` — split problem, solve subproblems, merge
- `merge` — combine two sorted halves into one sorted array
- `timsort` — hybrid merge + insertion sort; used in Python and Java's Arrays.sort for objects

**Key sorting algorithms:**

| Algorithm | Best | Average | Worst | Space | Stable |
|-----------|------|---------|-------|-------|--------|
| Bubble Sort | O(n) | O(n²) | O(n²) | O(1) | ✅ |
| Insertion Sort | O(n) | O(n²) | O(n²) | O(1) | ✅ |
| Merge Sort | O(n log n) | O(n log n) | O(n log n) | O(n) | ✅ |
| Quick Sort | O(n log n) | O(n log n) | O(n²) | O(log n) | ❌ |
| Heap Sort | O(n log n) | O(n log n) | O(n log n) | O(1) | ❌ |
| Counting Sort | O(n+k) | O(n+k) | O(n+k) | O(k) | ✅ |

**Must understand:**
- [ ] Why QuickSort is O(n²) worst case and how to avoid it (random pivot)
- [ ] Why MergeSort is always O(n log n) but uses O(n) extra space
- [ ] When to use counting/radix sort (small integer range)
- [ ] What Go's `sort.Slice` uses (pattern-defeating quicksort, introsort)

---

## Recursion & Backtracking

**Key Terms:**
- `base case` — condition that stops recursion; without it → infinite recursion
- `recursive case` — breaks problem into smaller subproblem and calls itself
- `call stack` — each recursive call adds a frame; too deep → stack overflow
- `backtracking` — try an option, recurse, undo the option (unmake the choice)
- `state` — what's been chosen so far
- `decision tree` — tree of all possible choices at each step
- `pruning` — skip branches that can't possibly lead to valid solution
- `permutation` — ordered arrangement of elements
- `combination` — unordered selection of elements
- `subset` — any subset of the original set (include or exclude each element)
- `constraint` — rule that a valid solution must satisfy

**Backtracking template:**
```
func backtrack(state, choices):
    if base_case(state):
        result.add(state)
        return
    for choice in choices:
        if is_valid(choice):
            make_choice(choice)
            backtrack(state, remaining_choices)
            undo_choice(choice)      ← this is backtracking
```

**Common problems:**
- Permutations (N! leaves)
- Combinations (C(n,k) leaves)
- Subsets (2^N leaves)
- N-Queens
- Sudoku solver
- Word search on grid

**Must understand:**
- [ ] Why backtracking is exponential but pruning makes it practical
- [ ] How to avoid duplicates in subset/combination problems (sort + skip duplicates)
- [ ] The difference between permutations and combinations
- [ ] How to pass state as a copy vs mutate + undo

---

## Dynamic Programming

**Key Terms:**
- `overlapping subproblems` — same subproblem is solved multiple times in brute force
- `optimal substructure` — optimal solution built from optimal solutions of subproblems
- `memoization (top-down)` — recursion + cache; compute on demand
- `tabulation (bottom-up)` — fill DP table iteratively from smallest subproblems
- `state` — what uniquely identifies a subproblem (e.g., dp[i][j] = max profit with i items and j capacity)
- `recurrence relation` — how to compute dp[i] from previous states
- `base case` — initial values in the DP table
- `1D DP` — dp[i] depends on dp[i-1] or earlier
- `2D DP` — dp[i][j] depends on neighboring states (grid problems, knapsack)
- `space optimization` — often 2D DP can be reduced to O(n) space using two rows
- `Fibonacci` — simplest DP example: dp[n] = dp[n-1] + dp[n-2]
- `coin change` — how many coins to make amount
- `knapsack` — maximize value with weight constraint
- `LCS (Longest Common Subsequence)` — classic 2D DP
- `LIS (Longest Increasing Subsequence)` — classic 1D DP

**DP process:**
1. Define state — what does dp[i] mean?
2. Write recurrence — how does dp[i] depend on previous states?
3. Identify base cases
4. Determine iteration order (which states to compute first)
5. Extract answer from dp table

**Must understand:**
- [ ] Memoization vs tabulation — when to use each
- [ ] How to define state (the hardest part)
- [ ] How to optimize 2D DP to O(n) space
- [ ] Classic problems: Fibonacci, coin change, knapsack, LCS, LIS

---

## Greedy Algorithms

**Key Terms:**
- `greedy choice` — locally optimal choice at each step
- `greedy correctness` — not all greedy choices lead to global optimum; must prove
- `exchange argument` — proof technique: show swapping any two choices doesn't improve solution
- `interval scheduling` — select maximum non-overlapping intervals (sort by end time)
- `activity selection` — variant of interval scheduling
- `Huffman coding` — greedy tree construction for compression
- `minimum spanning tree` — Kruskal's (greedy) and Prim's algorithms
- `fractional knapsack` — greedy works (take highest value/weight ratio first); 0/1 knapsack doesn't

**When greedy works:**
- Intervals (sort by end time, always pick earliest end)
- Coin change with canonical coins (quarters, dimes... — doesn't always work for arbitrary denominations)
- Dijkstra (greedy: always expand shortest unvisited node)
- MST (Kruskal: always add cheapest edge that doesn't form cycle)

**Must understand:**
- [ ] Why greedy doesn't always work (coin change with coins [1,3,4] and target 6)
- [ ] Interval scheduling — why sorting by end time is greedy correct
- [ ] The exchange argument — how to think about proving greedy correctness

---

## Dijkstra's Algorithm

**Key Terms:**
- `weighted graph` — edges have non-negative weights
- `shortest path` — minimum total weight from source to destination
- `priority queue (min-heap)` — always expand the node with smallest known distance
- `dist array` — dist[v] = current shortest distance from source to v; initially infinity
- `relaxation` — if dist[u] + w(u,v) < dist[v], update dist[v]
- `visited set` — once a node is popped from heap, its shortest path is finalized
- `negative weights` — Dijkstra doesn't work; use Bellman-Ford instead
- `Bellman-Ford` — works with negative weights; O(VE)
- `SSSP` — Single Source Shortest Path

**Complexity:**

| Implementation | Time |
|----------------|------|
| With min-heap | O((V + E) log V) |
| With array | O(V²) — for dense graphs |

**Dijkstra template:**
```
dist = {source: 0, all others: infinity}
heap = [(0, source)]
while heap not empty:
    d, u = heap.pop_min()
    if d > dist[u]: continue   ← stale entry
    for each neighbor v of u with weight w:
        if dist[u] + w < dist[v]:
            dist[v] = dist[u] + w
            heap.push((dist[v], v))
```

**Must understand:**
- [ ] Why negative weights break Dijkstra
- [ ] The "stale entry" check and why it's needed with a heap
- [ ] Difference between Dijkstra (SSSP) and Floyd-Warshall (all-pairs shortest path)

---

## Topological Sort

**Key Terms:**
- `topological order` — linear ordering of vertices where every directed edge u→v has u before v
- `DAG` — required; topological sort only works on Directed Acyclic Graphs
- `in-degree` — number of incoming edges to a node
- `Kahn's algorithm` — BFS-based; start with all nodes of in-degree 0
- `DFS-based topo sort` — postorder DFS; reverse the postorder gives topo order
- `cycle detection` — if Kahn's produces fewer than V nodes, there's a cycle
- `dependency resolution` — tasks with dependencies, build systems, course scheduling

**Kahn's algorithm:**
```
1. Compute in-degree for all nodes
2. Add all nodes with in-degree 0 to queue
3. While queue not empty:
   - Pop node u, add to result
   - For each neighbor v of u: in-degree[v]--; if 0, add to queue
4. If len(result) < V → cycle exists
```

**Must understand:**
- [ ] Why topological sort is undefined for graphs with cycles
- [ ] How to detect a cycle using Kahn's (result length < V)
- [ ] Course schedule problem — classic topo sort application
- [ ] Build system dependency ordering
