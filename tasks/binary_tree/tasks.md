# Binary Tree — Implementation Tasks

Complete these in order. Each one builds on your understanding.

---

## Task 1: InorderRecursive — Foundation

**Objective**: Implement the simplest traversal to understand recursion pattern.

**Function signature** (already in `binary_tree.go`):
```go
func InorderRecursive(node *TreeNode, result *[]int) {
    // TODO: implement
}
```

**Expected behavior**:
```
Input tree:
    5
   / \
  3   7
 / \
2   4

Output: result = [2, 3, 4, 5, 7]
```

**Your implementation checklist**:
- [ ] Handle base case (nil node)
- [ ] Visit left subtree first
- [ ] Append current node value
- [ ] Visit right subtree last

**Test in main.go** after implementing:
```go
root := NewTreeNode(5)
root.Left = NewTreeNode(3)
root.Right = NewTreeNode(7)
root.Left.Left = NewTreeNode(2)
root.Left.Right = NewTreeNode(4)

result := []int{}
InorderRecursive(root, &result)
// Should print: [2 3 4 5 7]
```

---

## Task 2: PreorderRecursive — Changing the Order

**Objective**: Understand that only order changes, logic is identical.

**Function signature**:
```go
func PreorderRecursive(node *TreeNode, result *[]int) {
    // TODO: implement
}
```

**Expected behavior**:
```
Same tree as Task 1
Output: [5, 3, 2, 4, 7]  ← root comes first
```

**Implementation note**:
Copy your Task 1 code and change only the ORDER of three operations. Nothing else changes.

---

## Task 3: PostorderRecursive — Deletion Pattern

**Objective**: Understand why postorder is used for deletion.

**Function signature**:
```go
func PostorderRecursive(node *TreeNode, result *[]int) {
    // TODO: implement
}
```

**Expected behavior**:
```
Same tree
Output: [2, 4, 3, 7, 5]  ← root comes last
```

**Implementation note**:
Again, same structure, different order.

**Discussion**: After implementing, think: "Why would you use this order to delete a tree?"

---

## Task 4: InorderIterative — Using Explicit Stack

**Objective**: Understand how recursion stack works by implementing it manually.

**Function signature**:
```go
func InorderIterative(root *TreeNode) []int {
    // TODO: implement
    return nil
}
```

**Expected behavior**:
```
Same tree, same output as Task 1
Output: [2, 3, 4, 5, 7]
```

**Algorithm** (pseudocode):
```
1. Create empty stack
2. Set current = root
3. While current is not nil OR stack is not empty:
   a. While current is not nil:
      - Push current onto stack
      - Go to left child
   b. current = pop from stack
   c. Visit current (append to result)
   d. Go to right child
```

**Hint**: This might be tricky. Take your time and trace through the algorithm step-by-step on paper first.

---

## Task 5: MaxDepth — Calculating Height

**Objective**: Understand recursive calculation of tree properties.

**Function signature**:
```go
func MaxDepth(root *TreeNode) int {
    // TODO: implement
    return 0
}
```

**Expected behavior**:
```
Tree:
    1       depth = 3 (1→3→7 is longest path with 3 nodes)
   / \
  2   3
     /
    7

Input: root from above
Output: 3
```

**Algorithm**:
- If node is nil, return 0
- Otherwise, return 1 + max(MaxDepth(left), MaxDepth(right))

**Edge cases**:
- Single node → should return 1
- Empty tree (nil) → should return 0

---

## Task 6: IsBalanced — Combining Height and Balance

**Objective**: Understand optimization technique (returning multiple values).

**Function signature**:
```go
func IsBalanced(root *TreeNode) bool {
    // TODO: implement (will need helper function)
    return false
}

// You'll need to write a helper function that returns (height, isBalanced)
// func checkBalance(node *TreeNode) (int, bool) {
//     // TODO
// }
```

**Expected behavior**:
```
Balanced tree:      Not balanced tree:
    5                   5
   / \                 /
  3   7               3
 / \                 / \
2   4               2   4
              Returns: true        Returns: false

Why false? Height of left (3) - right (0) = 3 > 1
```

**Key insight**: 
This is tricky because you need to:
1. Check if left subtree is balanced
2. Check if right subtree is balanced
3. Check if CURRENT node is balanced
4. Return both the height AND the balance status

**Helper function pattern**:
```go
func checkBalance(node *TreeNode) (int, bool) {
    // Returns: (height, isBalanced)
    // Use short-circuit: if subtree is unbalanced, return early
}
```

---

## Task 7: InvertTree — Simple Recursive Operation

**Objective**: Understand in-place tree modification.

**Function signature**:
```go
func InvertTree(root *TreeNode) *TreeNode {
    // TODO: implement
    return nil
}
```

**Expected behavior**:
```
Before:             After:
    5                 5
   / \               / \
  3   7             7   3
 / \               / \
2   4             4   2
```

**Algorithm**:
- If node is nil, return nil
- Swap left and right children
- Recursively invert left subtree
- Recursively invert right subtree
- Return node

**Simplest task** — this is just 3 lines of real code!

---

## Task 8: IsValidBST — Understanding Constraints

**Objective**: Learn to pass context (bounds) through recursion.

**Function signature**:
```go
func IsValidBST(root *TreeNode) bool {
    // TODO: implement (will need helper)
    return false
}

// Helper function with bounds
// func checkBST(node *TreeNode, min, max *int) bool {
//     // TODO
// }
```

**Expected behavior**:
```
Valid BST:          Invalid BST:
    5                   5
   / \                 / \
  3   7               3   7
 / \                     /
2   4                   6  ← 6 in left subtree but > root!
Returns: true        Returns: false
```

**Key insight**: 
For each node:
- It must be greater than ALL ancestors in left path
- It must be less than ALL ancestors in right path

**Algorithm**:
```
checkBST(node, min, max):
    if node is nil: return true
    
    if min != nil AND node.value <= min:
        return false
    if max != nil AND node.value >= max:
        return false
    
    // For left subtree: max becomes current node's value
    leftValid = checkBST(node.left, min, node.value)
    
    // For right subtree: min becomes current node's value
    rightValid = checkBST(node.right, node.value, max)
    
    return leftValid AND rightValid
```

---

## Task 9: Diameter — Tracking Max Through Recursion

**Objective**: Understand pattern of tracking global max while recursing.

**Function signature**:
```go
func Diameter(root *TreeNode) int {
    // TODO: implement (will need helper with reference to maxDiameter)
    return 0
}

// Helper function that returns height and updates maxDiameter
// func diameterHelper(node *TreeNode, maxDiameter *int) int {
//     // Returns: height of this subtree
//     // Updates: maxDiameter (by reference)
// }
```

**Expected behavior**:
```
Tree:
      1
     / \
    2   3
   /     \
  4       5

Diameter: longest path is 4→2→1→3→5 = 4 edges (or 5 nodes)
Returns: 4
```

**Algorithm**:
```
diameterHelper(node, maxDiameter):
    if node is nil: return 0
    
    leftHeight = diameterHelper(node.left, maxDiameter)
    rightHeight = diameterHelper(node.right, maxDiameter)
    
    // Diameter through this node = left + right
    currentDiameter = leftHeight + rightHeight
    maxDiameter = max(maxDiameter, currentDiameter)
    
    // Return height for parent
    return 1 + max(leftHeight, rightHeight)
```

---

## Task 10: LowestCommonAncestor (LCA)

**Objective**: Understand search pattern in trees.

**Function signature**:
```go
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // TODO: implement
    return nil
}
```

**Expected behavior**:
```
Tree:
      5
     / \
    3   7
   / \
  2   4

LCA(2, 4) returns: node 3  (both 2 and 4 are in subtree of 3)
LCA(2, 7) returns: node 5  (2 is in left, 7 is in right, so root is LCA)
LCA(3, 4) returns: node 3  (3 itself is ancestor of 4)
```

**Algorithm**:
```
lca(root, p, q):
    if root is nil: return nil
    
    // If root matches either p or q, root is definitely part of LCA
    if root == p OR root == q:
        return root
    
    // Search in both subtrees
    left = lca(root.left, p, q)
    right = lca(root.right, p, q)
    
    // If both found in different subtrees, root is LCA
    if left != nil AND right != nil:
        return root
    
    // If found in only one subtree, return that one
    if left != nil: return left
    return right
```

---

## Completion Checklist

- [ ] Task 1: InorderRecursive — implemented and tested
- [ ] Task 2: PreorderRecursive — implemented and tested
- [ ] Task 3: PostorderRecursive — implemented and tested
- [ ] Task 4: InorderIterative — implemented and tested
- [ ] Task 5: MaxDepth — implemented and tested
- [ ] Task 6: IsBalanced — implemented and tested
- [ ] Task 7: InvertTree — implemented and tested
- [ ] Task 8: IsValidBST — implemented and tested
- [ ] Task 9: Diameter — implemented and tested
- [ ] Task 10: LowestCommonAncestor — implemented and tested

All functions tested in `main.go` with multiple test cases.

