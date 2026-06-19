# Binary Tree — User Level Understanding (Cycle 1)

## Why Binary Trees Matter

### The Problem
- Need efficient way to store hierarchical data (parent-child relationships)
- Array/slice: fast access but hard to represent hierarchy
- Linked list: represents sequence, not hierarchy
- Trees: naturally represent hierarchies (folders, organizations, family trees)

### What was used before
- Nested arrays/objects (complicated, inefficient)
- Multiple pointers and manual management (error-prone)

### Why Binary Tree is the solution
- Each node has at most 2 children (Left, Right) — simple structure
- Efficient for search, insertion, deletion when balanced
- Foundation for more complex structures (BSTs, heaps, etc.)

---

## Core Concept: Tree Node

### Structure
```go
type TreeNode struct {
    Value int        // The data
    Left  *TreeNode  // Left child
    Right *TreeNode  // Right child
}
```

### Key Terms
- **Leaf node**: A node with no children (both Left and Right are nil). Leaves are the terminal nodes at the tips of branches.
- **Root**: The topmost node of the tree.
- **Subtree**: Any node and all its descendants form a subtree.

### Key insight
A tree is **recursive by nature**:
- A tree = a node + a left subtree + a right subtree
- A subtree is also a tree
- Base case: nil (empty tree)

---

## Base Case — Why `if node == nil { return }`

The base case triggers **every time we follow a child pointer that doesn't exist**.

A leaf node has two nil children. When we call `inorder(leaf.Left)`, we pass nil. That's when the check triggers. It stops recursion from going infinitely.

For a tree with n nodes, the base case triggers n+1 times (every nil pointer we follow).

---

## Traversals — The Foundation

### Why traversals matter
To process every node in a tree, you must visit them in some order. Different orders solve different problems.

### Inorder: Left → Node → Right
**When to use**: Sorting (BST gives sorted order), tree comparison

```
      5
     / \
    3   7
   / \
  2   4

Result: [2, 3, 4, 5, 7]  ← sorted for BST!
```

**How it works at each node**: process entire left subtree, then visit current node, then process entire right subtree.

**Common mistake**: Inorder does NOT mean "bottom to up." It means Left → Node → Right at every level. After visiting leftmost node (3's left child = 2), you visit the PARENT (3), then right sibling (4).

### Preorder: Node → Left → Right
**When to use**: Tree copying, serialization, expressing structure

```
Result: [5, 3, 2, 4, 7]  ← root comes first
```

### Postorder: Left → Right → Node
**When to use**: Deleting tree (must delete children before parent), calculating heights

```
Result: [2, 4, 3, 7, 5]  ← root comes last
```

### Level Order (BFS): Level by level
**When to use**: Finding nodes at specific level, shortest path

```
Level 0: [5]
Level 1: [3, 7]
Level 2: [2, 4]

Result: [[5], [3, 7], [2, 4]]
```

---

## DFS vs BFS

### DFS (Depth-First Search)
- Goes as **deep** as possible first, then backtracks
- Uses **stack** (or recursion = computer's call stack)
- LIFO: last node saved = first to come back to
- Inorder, Preorder, Postorder are all DFS

### BFS (Breadth-First Search)
- Goes **wide** first, visits entire level before going deeper
- Uses **queue**
- FIFO: first node saved = first to process
- LevelOrder is BFS

### Why stack for DFS?
Stack = LIFO. You keep diving deeper (saving parent nodes), and when you hit bottom, you come back to the **most recent** parent. That's depth-first.

### Why queue for BFS?
Queue = FIFO. You add all children of current level, then process them in the **order they were added** — left to right, level by level. That's breadth-first.

**Senior-level answer:**
> "DFS uses a stack because it needs to backtrack to the most recently visited node — LIFO order. BFS uses a queue because it needs to process nodes in the order they were discovered — FIFO order. If you swapped them, DFS with a queue would behave breadth-first, and BFS with a stack would behave depth-first. The data structure determines the traversal order."

---

## Iterative vs Recursive Traversal

### Why they're the same
Recursive traversal uses the **computer's call stack** automatically. Iterative traversal uses **your own stack** manually. Same algorithm, different mechanism.

| Recursive (computer does it) | Iterative (you do it) |
|---|---|
| Function call = push frame | `stack.Push(current)` |
| Function returns = pop frame | `current = stack.Pop()` |
| Computer manages stack | You manage stack |

### Why use iterative?
Recursion can cause **stack overflow** on very deep trees (1 million levels). Your manual stack uses heap memory, which is much larger.

**Senior-level answer:**
> "The recursive approach is cleaner but risks stack overflow on deep trees. The iterative version uses an explicit stack on the heap, which can handle much deeper trees. Both are O(n) time and O(h) space where h is the height."

---

## BST (Binary Search Tree) Property

**The rule**: For every node, ALL nodes in the left subtree must be less, and ALL nodes in the right subtree must be greater.

**Common mistake**: Only checking direct parent-child relationship is NOT enough!

```
      6
     / \
    4   8
   / \
  2   7   ← INVALID! 7 is in left subtree of 6, but 7 > 6

Each node looks correct with its direct parent:
  7 > 4 ✓ (right child of 4)
  But 7 > 6 ✗ (violates ancestor constraint)
```

**The correct check**: Must validate against ALL ancestors, not just the parent. This is why IsValidBST passes min/max bounds down through recursion.

**Senior-level answer:**
> "A common mistake is only checking if left child < parent < right child. The BST invariant requires that every node in the left subtree is strictly less than the root, and every node in the right subtree is strictly greater — including all descendants, not just direct children. This is why validating a BST requires passing min/max bounds down through the recursion."

---

## Discussion Conclusions

### What we confirmed you understand:
- Tree node structure (Value, Left, Right)
- Leaf nodes have nil children
- Base case triggers when following a nil pointer
- Inorder gives sorted order for BST
- DFS uses stack, BFS uses queue
- LevelOrder processes level by level using queue

### What was corrected:
- Inorder is NOT "bottom to up" — it's Left → Node → Right at every node
- BST validation: must check against ALL ancestors, not just direct parent

---

## Problems for Cycle 2

These are more complex and will be covered in Cycle 2:
- MaxDepth, IsBalanced, IsValidBST, InvertTree, Diameter, LCA
- They build on the traversal patterns learned here

---

## English Corrections from Discussion

❌ "it starts from bottom"
✅ "It starts from the root but reaches the deepest node first before backtracking"

❌ "FIFI"
✅ "FIFO (First In, First Out)"

❌ "if there is no such node, then it will be nil"
✅ "When we follow a child pointer from a leaf node, the pointer is nil because that child doesn't exist"

❌ "every left child is smaller than its parent"
✅ "Every node in the entire left subtree must be less than the ancestor node — not just direct children"
