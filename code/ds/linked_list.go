package ds

import imp "code/ds/ds_imp"

// MakeCycle makes the node with cycleVal point back to the node with targetVal.
// Used to create test cases with cycles.
func MakeCycle(head *imp.Node, cycleVal, targetVal int) {
	var target, cycleNode *imp.Node
	current := head
	for current != nil {
		if current.Value == targetVal {
			target = current
		}
		if current.Value == cycleVal {
			cycleNode = current
		}
		current = current.Next
	}
	if target != nil && cycleNode != nil {
		cycleNode.Next = target
	}
}

// Print prints the list up to maxNodes to avoid infinite loop on cyclic lists
func Print(head *imp.Node, maxNodes int) []int {
	var result []int
	current := head
	count := 0
	for current != nil && count < maxNodes {
		result = append(result, current.Value)
		current = current.Next
		count++
	}
	return result
}

// HasCycleHashSet detects a cycle using a hash set of visited node pointers.
// Time: O(n), Space: O(n)
func HasCycleHashSet(head *imp.Node) bool {
	// your implementation here
	visited := make(map[*imp.Node]bool)
	current := head

	for current != nil {
		if visited[current] {
			return true
		}

		visited[current] = true
		current = current.Next
	}
	return false
}

// HasCycleSlowFast detects a cycle using slow/fast pointers (Floyd's algorithm).
// Time: O(n), Space: O(1)
func HasCycleSlowFast(head *imp.Node) bool {
	// your implementation here
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func FindNthFromEnd(head *imp.Node, n int) *imp.Node {

	front := head
	for i := 0; i < n; i++ {
		if front == nil {
			return nil
		}
		front = front.Next
	}

	back := head
	for front != nil {
		front = front.Next
		back = back.Next
	}

	return back

}

func MergeSortedLists(a *imp.Node, b *imp.Node) *imp.Node {

	aCurr, bCurr := a, b
	newLL := &imp.Node{}
	curr := newLL

	for aCurr != nil && bCurr != nil {
		if aCurr.Value >= bCurr.Value {
			curr.Next = bCurr
			bCurr = bCurr.Next
		} else {
			curr.Next = aCurr
			aCurr = aCurr.Next
		}

		curr = curr.Next
	}
	if aCurr != nil {
		curr.Next = aCurr
	}
	if bCurr != nil {
		curr.Next = bCurr
	}

	return newLL.Next
}
