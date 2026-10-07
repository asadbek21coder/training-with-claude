package imp

// Node is a single element in a linked list
type Node struct {
	Value int
	Next  *Node
}

// NewList builds a linked list from a slice of values
func NewList(values []int) *Node {
	if len(values) == 0 {
		return nil
	}
	head := &Node{Value: values[0]}
	current := head
	for _, v := range values[1:] {
		current.Next = &Node{Value: v}
		current = current.Next
	}
	return head
}
