# Binary Tree — Discussion Questions

These questions are designed to test your understanding and force you to think deeply. Answer them, and we'll discuss together.

---

## Traversals

### Easy Level

**Q1: Recursion Base Case**
Look at this code:
```go
func inorder(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    // ... rest of code
}
```

Why do we need the `if node == nil { return }` check at the beginning?

<details>
<summary>Show hint</summary>
What happens if we remove it? Where does the recursion stop?
</details>

---

**Q2: Order Matters**
For Inorder traversal: `Left → Node → Right`

If we changed it to `Node → Left → Right`, what would we get? (What's the name of this traversal?)

---

### Medium Level

**Q3: Traversal Choice**
You have a tree and need to **delete all nodes** (free memory). Which traversal order would you use and why?

```go
function deleteTree(node):
    // Your approach here
```

<details>
<summary>Hint</summary>
Can you delete a parent node before its children? What happens if you try?
</details>

---

**Q4: Iterative vs Recursive**
Why would someone use InorderIterative instead of InorderRecursive?

<details>
<summary>Hint</summary>
What could go wrong with recursion on very deep trees?
</details>

---

**Q5: Level Order is Different**
Inorder, Preorder, Postorder use **recursion**. But LevelOrder uses **queue**. Why?

Can you convert LevelOrder to recursion? (Try thinking about it)

---

### Hard Level

**Q6: Memory in Recursion**
When you call `InorderRecursive(node, result)`, the computer creates a call stack. 

```
inorder(5) 
  → inorder(3)
      → inorder(2)
          → inorder(nil) returns
      → append 2
    → append 3
    → inorder(4)
        → inorder(nil) returns
      → append 4
  → append 5
→ inorder(7)
    → ...
```

For a tree with 1000 nodes, how deep could this stack get? What's the problem?

---

## Depth & Height

### Easy Level

**Q7: Definition**
For this tree, what is MaxDepth?

```
    1
   / \
  2   3
 /
4
```

Count carefully: is it the number of **nodes** in the longest path, or the number of **edges**?

---

**Q8: Edge Cases**
What should MaxDepth return for:
- An empty tree (root is nil)?
- A tree with just one node?

---

### Medium Level

**Q9: Why Recursion Works**
Explain this logic in your own words:

> MaxDepth(node) = 1 + max(MaxDepth(left), MaxDepth(right))

Why does adding 1 each time give us the right depth?

---

**Q10: Efficiency**
MaxDepth visits every node once. So it's O(n). But what about space complexity?

Does it use O(1) space? O(n) space? Something in between?

<details>
<summary>Hint</summary>
What uses space during recursion?
</details>

---

## Balance & Validation

### Medium Level

**Q11: Balanced Definition**
A balanced tree has this property:
> For every node, |left_height - right_height| ≤ 1

Why is this definition useful? What problem does it solve?

---

**Q12: IsBalanced Optimization**
Look at this approach:
```go
// Bad approach
function isBal(node):
    if node is nil: return true
    leftHeight = maxDepth(node.left)
    rightHeight = maxDepth(node.right)
    if abs(leftHeight - rightHeight) > 1: return false
    return isBal(node.left) AND isBal(node.right)

// Good approach
function checkBal(node):
    return (height, isBalanced) tuple
```

Why is the second approach faster? What's the complexity difference?

---

**Q13: BST Constraint**
A valid BST must satisfy: `left < node < right`

But this is not enough! Why?

```
      5
     / \
    3   7
   / \
  2   6   ← This 6 is in left subtree but > 5!
           This tree looks locally valid but is NOT a valid BST
```

What else must you check?

---

### Hard Level

**Q14: Pointer Bounds**
In `IsValidBST`, we pass bounds down the tree:

```
checkBST(node, min, max):
    if node value violates bounds: return false
    checkBST(node.left, min, node.value)    ← why node.value as max?
    checkBST(node.right, node.value, max)   ← why node.value as min?
```

Explain why we update the bounds this way.

---

## Complex Problems

### Medium Level

**Q15: Diameter Intuition**
Diameter is the longest path. This path:
- Does NOT have to go through root
- Does NOT have to be vertical

```
      1
     / \
    2   3
   /     \
  4       5
 /         \

Path 4-2-1-3-5 has length 4 (in terms of edges)
or 5 nodes
```

For a node, how do you calculate: what's the longest path passing through this node?

<details>
<summary>Hint</summary>
Think about left_height + right_height
</details>

---

**Q16: LCA Logic**
Lowest Common Ancestor of nodes P and Q:

```
        5
       / \
      3   7
     / \
    2   4

LCA(2, 4) = 3  (3 is parent of both)
LCA(2, 7) = 5  (5 is parent of subtree containing both)
LCA(3, 4) = 3  (3 itself is ancestor of 4)
```

If you search and find Q in the left subtree but not in right subtree, what does that tell you?

---

## Conceptual Understanding

### Hard Level

**Q17: Why Trees Are Recursive**
Explain: "A tree is a recursive data structure"

What do you mean by that? How does it help you solve problems?

---

**Q18: When to Use Trees**
Compare:
- Array for storing 1000 numbers
- Tree for storing 1000 numbers

In what scenario is a tree better? What problem does it solve that array doesn't?

---

## Your Own Questions

**Q19: Ask Me**
After reading the notes, write ONE question you have about binary trees that you're not sure about.

(This helps me know what to explain better)

