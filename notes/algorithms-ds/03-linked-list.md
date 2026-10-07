# Linked List — User Level Understanding

## The Problem Linked Lists Solve

### Why not just use arrays?
Arrays are great for random access, but terrible for:
- **Inserting at the beginning**: O(n) — must shift all elements
- **Deleting from the middle**: O(n) — must shift elements after
- **Dynamic sizing**: Array resizing copies all elements

### What linked list offers
- **Insert at head**: O(1) — just change one pointer
- **Delete any node** (if you have a reference): O(1) — change pointers
- **No pre-allocation**: Each node is allocated independently

### The trade-off
- **No random access**: Can't jump to index 5. Must walk from head. O(n)
- **Not cache-friendly**: Nodes are scattered in memory (unlike array's contiguous block)
- **Extra memory**: Each node stores a pointer alongside data

---

## Core Concept: Node

```go
type Node struct {
    Value int    // The data
    Next  *Node  // Pointer to next node (or nil if last)
}
```

```
Head → [1|•] → [2|•] → [3|•] → [4|•] → nil
        ↑
       node     node     node    node (tail)
```

Each node knows only about the NEXT node. To find a node, you must start from head and walk.

---

## Key Operations

### Insert at Head — O(1)
```
newNode.Next = head
head = newNode

Before: head → [2] → [3] → nil
After:  head → [1] → [2] → [3] → nil
```

### Insert at Tail — O(n)
Must walk to the end first, then attach.

### Delete a Node — O(1) if you have reference, O(n) to find it
Change the previous node's pointer to skip over the deleted node.

### Search — O(n)
Walk from head, check each node.

---

## Patterns We Implemented

### 1. Two Pointer (Slow/Fast) — Floyd's Algorithm
**Problem**: Detect if linked list has a cycle

**How it works**:
- Slow pointer moves 1 step at a time
- Fast pointer moves 2 steps at a time
- If there's a cycle, they WILL meet
- If no cycle, fast reaches nil

**Why it works**: In a cycle, fast "chases" slow. The gap closes by 1 each step. Like two runners on a circular track — the faster one eventually laps the slower one.

```go
slow := head
fast := head
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
    if slow == fast {
        return true  // cycle detected
    }
}
return false  // no cycle
```

**Time**: O(n) | **Space**: O(1)

### 2. Hash Set Approach — Cycle Detection
**How it works**: Store every visited node in a set. If you visit one twice, there's a cycle.

```go
visited := map[*Node]bool{}
for current != nil {
    if visited[current] {
        return true  // seen before = cycle
    }
    visited[current] = true
    current = current.Next
}
return false
```

**Time**: O(n) | **Space**: O(n) — uses extra memory for the set

### 3. Two Pointer Gap — Find Nth from End
**Problem**: Find the nth node from the end without knowing the list length

**How it works**:
- Move front pointer n steps ahead
- Then move both front and back together
- When front reaches nil, back is at the nth from end

```
Find 3rd from end in: 1 → 2 → 3 → 4 → 5

Step 1: front moves 3 ahead
  back: 1, front: 4

Step 2: both move together
  back: 2, front: 5
  back: 3, front: nil ← STOP

Answer: 3 ✓
```

### 4. Merge Two Sorted Lists
**Problem**: Combine two sorted lists into one sorted list

**How it works**:
- Compare heads of both lists
- Take the smaller one, attach to result
- Move forward in that list
- Repeat until one list is exhausted
- Attach the remaining list

**Time**: O(n + m) | **Space**: O(1) — reuses existing nodes

---

## Senior-Level Perspective

### How to explain in interview

**Cycle Detection (Floyd's)**:
> "I use Floyd's cycle detection algorithm with two pointers — slow moves one step, fast moves two steps. If a cycle exists, the fast pointer will eventually catch up to the slow pointer inside the cycle. This gives O(n) time and O(1) space, which is optimal since the hash set approach requires O(n) extra space."

**Nth from End**:
> "I use the two-pointer gap technique. First, I advance one pointer by n positions to create a gap. Then I move both pointers simultaneously until the lead pointer reaches the end. The trailing pointer is now at the nth node from the end. This avoids needing to know the list length upfront, giving us O(n) time and O(1) space in a single pass."

**Merge Sorted Lists**:
> "I use a dummy head node to simplify the merge logic. I compare the current nodes of both lists, attach the smaller one to the result, and advance that list's pointer. When one list is exhausted, I attach the remainder of the other. This runs in O(n+m) time and O(1) space since we reuse existing nodes."

---

## English Notes

❌ "Go through the list"
✅ "Traverse the list" or "Walk through the list"

❌ "Point it to the next"
✅ "Set its Next pointer to the following node"

❌ "The fast one catches the slow one"
✅ "The fast pointer eventually meets the slow pointer inside the cycle"

