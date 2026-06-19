# Linked List — Completed Tasks

## ✅ Task 1: Detect Cycle — Hash Set Approach

**Status**: Completed and tested

**Problem**: Detect if a linked list has a cycle
**Solution**: Store visited nodes in a hash set, check for duplicates
**Time**: O(n) | **Space**: O(n)
**Implementation**: `HasCycleHashSet` in `linked_list.go`

---

## ✅ Task 2: Detect Cycle — Slow/Fast (Floyd's)

**Status**: Completed and tested

**Problem**: Same as above, but with O(1) space
**Solution**: Two pointers — slow (1 step), fast (2 steps). If cycle exists, they meet.
**Time**: O(n) | **Space**: O(1)
**Implementation**: `HasCycleSlowFast` in `linked_list.go`

---

## ✅ Task 3: Find Nth From End

**Status**: Completed and tested

**Problem**: Find the nth node from the end without knowing list length
**Solution**: Two-pointer gap technique — advance front by n, then move both
**Time**: O(n) | **Space**: O(1)
**Implementation**: `FindNthFromEnd` in `linked_list.go`

---

## ✅ Task 4: Merge Two Sorted Lists

**Status**: Completed and tested

**Problem**: Merge two sorted linked lists into one sorted list
**Solution**: Dummy head, compare and attach smaller node each step
**Time**: O(n+m) | **Space**: O(1)
**Implementation**: `MergeSortedLists` in `linked_list.go`

---

## ⏭️ Pending (From SPRINT.md, Not Done Yet)

### Reverse a Linked List in Place
- Listed as done in SPRINT.md but no implementation found in code

### Remove Duplicates from Sorted Linked List
- Not started

### Check if Linked List is a Palindrome
- Not started

**Decision**: Moving forward to Binary Tree topic. Will revisit if time allows.
