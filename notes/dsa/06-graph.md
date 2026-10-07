# Graph — User Level Understanding (Cycle 1)

## The Problem Graphs Solve

Arrays, linked lists, trees — all limited:
- Arrays/Lists: linear, one direction
- Trees: hierarchical, one parent per node, no cycles

Real world has relationships that don't follow these rules:
- Social networks (cycles: A follows B follows C follows A)
- Road maps (multiple paths between cities)
- Course prerequisites (multiple paths to same destination)

## What is a Graph?

Nodes (vertices) connected by edges. No rules about parents or direction. Any node can connect to any other. Cycles allowed.

```
    A ─── B
    |     |
    C ─── D
```

## Key Terms

- **Vertex (node)**: A point in the graph
- **Edge**: A connection between two vertices
- **Directed**: Edge goes one way (A → B, like Twitter follows)
- **Undirected**: Edge goes both ways (A ─ B, like Facebook friends)
- **Weighted**: Edges have a cost (distance between cities)
- **Unweighted**: Edges have no cost (just connected or not)
- **Cycle**: A path that starts and ends at the same node
- **Connected**: Every node can reach every other node

## Real-World Examples

| Example | Directed? | Weighted? |
|---------|-----------|-----------|
| Facebook friendships | Undirected | Unweighted |
| Twitter follows | Directed | Unweighted |
| Google Maps roads | Undirected | Weighted (distance) |

## How to Store a Graph

### Adjacency List (use this in interviews)
```go
graph := map[string][]string{
    "A": {"B", "C"},
    "B": {"A", "D"},
    "C": {"A", "D"},
    "D": {"B", "C"},
}
```
- ✅ Less memory for sparse graphs
- ✅ Fast to find neighbors

### Adjacency Matrix
```
    A  B  C  D
A [ 0, 1, 1, 0 ]
B [ 1, 0, 0, 1 ]
```
- ✅ O(1) to check "is A connected to B?"
- ❌ O(n²) memory even if few edges

## Graph Traversal = Tree Traversal + Visited Set

**The only difference from trees:** graphs can have cycles, so you must track visited nodes.

```go
visited := map[string]bool{}
// Before visiting any node:
// if visited[node] → skip
// else → mark visited, process it
```

**DFS on graph:**
```
visited: {}
DFS(A): visited {A} → go to B
  DFS(B): visited {A,B} → A visited, go to D
    DFS(D): visited {A,B,D} → B visited, go to C
      DFS(C): visited {A,B,D,C} → A,D visited → done
```

**BFS on graph:** Same idea, use queue + visited set.

## Number of Islands Pattern

Grid of 1s (land) and 0s (water). Connected 1s = one island.

```
1 1 0 0 0
1 1 0 0 0    → 2 islands
0 0 0 1 1
0 0 0 1 1
```

**Algorithm:**
1. Loop through every cell
2. Find unvisited `1` → new island, count++
3. DFS/BFS from that cell, mark all connected `1`s as visited
4. Continue — next unvisited `1` is another island

Both DFS and BFS give the same result — they both visit all connected cells.

## Senior-Level Answers

**"What is a graph?"**
> "A graph is a set of nodes connected by edges. Unlike trees, graphs can have cycles and nodes can have multiple connections. They model relationships like social networks, road maps, and dependencies."

**"How do you traverse a graph?"**
> "Same as trees — DFS with a stack, BFS with a queue. The key difference is you need a visited set to avoid infinite loops from cycles. Before processing any node, check if it's already visited."

**"How do you solve Number of Islands?"**
> "Loop through the grid. When you find an unvisited land cell, increment the island count and run DFS or BFS to mark all connected land cells as visited. The number of times you start a new DFS equals the number of islands. Time is O(rows × cols), space is O(rows × cols) for the visited set."
