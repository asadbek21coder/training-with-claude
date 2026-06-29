package main

import (
	"code/ds"
	"fmt"
)

func main() {
	// fmt.Println("=== Cycle Detection ===")
	// fmt.Println()

	// // Case 1: no cycle
	// list1 := imp.NewList([]int{1, 2, 3, 4, 5})
	// fmt.Println("List (no cycle):", ds.Print(list1, 10))
	// fmt.Println("HasCycle (hash set): ", ds.HasCycleHashSet(list1))
	// fmt.Println("HasCycle (slow/fast):", ds.HasCycleSlowFast(list1))
	// fmt.Println()

	// // Case 2: cycle — tail points back to head (node 1)
	// list2 := imp.NewList([]int{1, 2, 3, 4, 5})
	// ds.MakeCycle(list2, 5, 1)
	// fmt.Println("List (cycle: 5 → 1):", ds.Print(list2, 10), "...")
	// fmt.Println("HasCycle (hash set): ", ds.HasCycleHashSet(list2))
	// fmt.Println("HasCycle (slow/fast):", ds.HasCycleSlowFast(list2))
	// fmt.Println()

	// // Case 3: cycle — tail points back to middle (node 3)
	// list3 := imp.NewList([]int{1, 2, 3, 4, 5})
	// ds.MakeCycle(list3, 5, 3)
	// fmt.Println("List (cycle: 5 → 3):", ds.Print(list3, 10), "...")
	// fmt.Println("HasCycle (hash set): ", ds.HasCycleHashSet(list3))
	// fmt.Println("HasCycle (slow/fast):", ds.HasCycleSlowFast(list3))
	// fmt.Println()

	// // Case 4: single node, no cycle
	// list4 := imp.NewList([]int{1})
	// fmt.Println("List (single node):", ds.Print(list4, 10))
	// fmt.Println("HasCycle (hash set): ", ds.HasCycleHashSet(list4))
	// fmt.Println("HasCycle (slow/fast):", ds.HasCycleSlowFast(list4))
	// fmt.Println()

	// // Case 5: empty list
	// fmt.Println("List (empty):", ds.Print(nil, 10))
	// fmt.Println("HasCycle (hash set): ", ds.HasCycleHashSet(nil))
	// fmt.Println("HasCycle (slow/fast):", ds.HasCycleSlowFast(nil))

	// fmt.Println()
	// fmt.Println("=== Nth From End ===")
	// fmt.Println()

	// list5 := imp.NewList([]int{1, 2, 3, 4, 5})
	// fmt.Println("List:", ds.Print(list5, 10))
	// fmt.Println("3rd from end:", ds.FindNthFromEnd(list5, 3).Value) // expect 3
	// fmt.Println("1st from end:", ds.FindNthFromEnd(list5, 1).Value) // expect 5
	// fmt.Println("5th from end:", ds.FindNthFromEnd(list5, 5).Value) // expect 1
	// fmt.Println("n > length:  ", ds.FindNthFromEnd(list5, 9))       // expect nil

	// fmt.Println()
	// fmt.Println("=== Merge Sorted Lists ===")
	// fmt.Println()

	// a := imp.NewList([]int{1, 3, 5})
	// b := imp.NewList([]int{2, 4, 6})
	// fmt.Println("A:", ds.Print(a, 10))
	// fmt.Println("B:", ds.Print(b, 10))
	// fmt.Println("Merged:", ds.Print(ds.MergeSortedLists(a, b), 10)) // expect [1 2 3 4 5 6]

	// a2 := imp.NewList([]int{1, 2, 3})
	// b2 := imp.NewList([]int{4, 5, 6})
	// fmt.Println("Merged (no overlap):", ds.Print(ds.MergeSortedLists(a2, b2), 10)) // expect [1 2 3 4 5 6]

	// a3 := imp.NewList([]int{1, 3, 5, 7})
	// b3 := imp.NewList([]int{2, 4})
	// fmt.Println("Merged (different lengths):", ds.Print(ds.MergeSortedLists(a3, b3), 10)) // expect [1 2 3 4 5 7]

	// fmt.Println()
	// fmt.Println("=== Valid Parentheses ===")
	// fmt.Println()
	// fmt.Println("({[]})  →", ds.IsValidParatheses("({[]})")) // true
	// fmt.Println("({[}])  →", ds.IsValidParatheses("({[}])")) // false
	// fmt.Println("(((    →", ds.IsValidParatheses("((("))     // false
	// fmt.Println(")      →", ds.IsValidParatheses(")"))       // false
	// fmt.Println("[]{}() →", ds.IsValidParatheses("[]{}()"))  // true
	// fmt.Println("empty  →", ds.IsValidParatheses(""))        // true

	// fmt.Println()
	// fmt.Println("=== Valid Parentheses (Stack impl) ===")
	// fmt.Println()
	// fmt.Println("({[]})  →", ds.IsValidParanthesesStack("({[]})"))
	// fmt.Println("({[}])  →", ds.IsValidParanthesesStack("({[}])"))
	// fmt.Println("(((    →", ds.IsValidParanthesesStack("((("))
	// fmt.Println(")      →", ds.IsValidParanthesesStack(")"))
	// fmt.Println("[]{}() →", ds.IsValidParanthesesStack("[]{}()"))
	// fmt.Println("empty  →", ds.IsValidParanthesesStack(""))

	// fmt.Println()
	// fmt.Println("=== MinStack ===")
	// fmt.Println()
	// minStack := ds.NewMinStack()
	// minStack.Push(5)
	// fmt.Println("Push(5) → min:", minStack.GetMin())                     // min: 5
	// minStack.Push(3)
	// fmt.Println("Push(3) → min:", minStack.GetMin())                     // min: 3
	// minStack.Push(7)
	// fmt.Println("Push(7) → min:", minStack.GetMin())                     // min: 3 (unchanged)
	// minStack.Push(1)
	// fmt.Println("Push(1) → min:", minStack.GetMin())                     // min: 1
	// fmt.Println("Pop():", minStack.Pop(), "→ min now:", minStack.GetMin()) // min: 3
	// fmt.Println("Pop():", minStack.Pop(), "→ min now:", minStack.GetMin()) // min: 3
	// fmt.Println("Pop():", minStack.Pop(), "→ min now:", minStack.GetMin()) // min: 5

	// fmt.Println()
	// fmt.Println("=== Daily Temperatures ===")
	// fmt.Println()
	// temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
	// result := ds.DailyTemperatures(temps)
	// fmt.Println("Input: ", temps)
	// fmt.Println("Output:", result)
	// fmt.Println("Expected: [1 1 4 2 1 1 0 0]")

	// Binary tree tests will be added here as you implement each function

	res := ds.BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4)
	fmt.Println(res)
}
