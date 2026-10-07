package imp

type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() {
	if len(q.items) == 0 {
		return
	}
	q.items = q.items[1:]
}

type QueueTreeNode struct {
	items []*TreeNode
}

func NewQueueTreeNode() *QueueTreeNode {
	return &QueueTreeNode{}
}

func (q *QueueTreeNode) Enqueue(item *TreeNode) {
	q.items = append(q.items, item)

}

func (q *QueueTreeNode) Dequeue() *TreeNode {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *QueueTreeNode) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *QueueTreeNode) Size() int {
	return len(q.items)
}

func (q *QueueTreeNode) Clear() {
	q.items = []*TreeNode{}
}
