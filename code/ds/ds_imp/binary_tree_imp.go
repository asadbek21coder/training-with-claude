package imp

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode creates a new tree node
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Value: val}
}

// ===== TRAVERSALS =====

// InorderRecursive - Left → Node → Right (gives sorted order for BST)
func InorderRecursive(node *TreeNode, result *[]int) {
	// base case
	if node == nil {
		return
	}

	InorderRecursive(node.Left, result)
	*result = append(*result, node.Value)
	InorderRecursive(node.Right, result)

}

// PreorderRecursive - Node → Left → Right
func PreorderRecursive(node *TreeNode, result *[]int) {
	// base case
	if node == nil {
		return
	}

	*result = append(*result, node.Value)
	PreorderRecursive(node.Left, result)
	PreorderRecursive(node.Right, result)
}

// PostorderRecursive - Left → Right → Node
func PostorderRecursive(node *TreeNode, result *[]int) {
	// base case
	if node == nil {
		return
	}

	PostorderRecursive(node.Left, result)
	PostorderRecursive(node.Right, result)
	*result = append(*result, node.Value)
}

// InorderIterative - using a stack
func InorderIterative(root *TreeNode) []int {
	current := root
	callStack := NewStackTreeNode()
	result := []int{}

	for current != nil || !callStack.IsEmpty() {
		for current != nil {
			callStack.Push(current)
			current = current.Left
		}

		current = callStack.Pop()
		result = append(result, current.Value)
		current = current.Right

	}

	return result

}

// LevelOrder - BFS using a queue (returns levels)
func LevelOrder(root *TreeNode) [][]int {
	qu := NewQueueTreeNode()
	if root == nil {
		return [][]int{}
	}
	qu.Enqueue(root)
	result := [][]int{}

	for !qu.IsEmpty() {
		levelSize := qu.Size()
		level := []int{}
		for i := 0; i < levelSize; i++ {
			node := qu.Dequeue()
			level = append(level, node.Value)
			if node.Left != nil {
				qu.Enqueue(node.Left)
			}
			if node.Right != nil {
				qu.Enqueue(node.Right)
			}
		}
		result = append(result, level)
	}


	return result
}

// ===== PROBLEMS =====

// MaxDepth - maximum depth of binary tree (height)
func MaxDepth(root *TreeNode) int {
	// TODO: implement
	return 0
}

// IsBalanced - check if binary tree is balanced
// A balanced tree: abs(left height - right height) <= 1 for every node
func IsBalanced(root *TreeNode) bool {
	// TODO: implement
	return false
}

// IsValidBST - check if tree is a valid BST
func IsValidBST(root *TreeNode) bool {
	// TODO: implement
	return false
}

// InvertTree - invert a binary tree (swap left and right)
func InvertTree(root *TreeNode) *TreeNode {
	// TODO: implement
	return nil
}

// Diameter - diameter of binary tree (longest path between any two nodes)
func Diameter(root *TreeNode) int {
	// TODO: implement
	return 0
}

// LowestCommonAncestor (LCA) - find lowest common ancestor of two nodes
// Assumes both p and q exist in tree
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// TODO: implement
	return nil
}
