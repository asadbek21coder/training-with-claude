package imp

type StackTreeNode struct {
	items []*TreeNode
}

func NewStackTreeNode() *StackTreeNode {
	return &StackTreeNode{}
}

func (s *StackTreeNode) Push(item *TreeNode) {
	s.items = append(s.items, item)
}

func (s *StackTreeNode) Pop() *TreeNode {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *StackTreeNode) Peek() *TreeNode {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *StackTreeNode) IsEmpty() bool {
	return len(s.items) == 0
}	
 
//----------------------------------------


type Stack struct {
	items []rune
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() rune {
	if len(s.items) == 0 {
		return 0
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() rune {
	if len(s.items) == 0 {
		return 0
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Clear() {
	s.items = []rune{}
}

func (s *Stack) String() string {
	return string(s.items)
}
