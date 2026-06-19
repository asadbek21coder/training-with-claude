# Stack & Queue — User Level Understanding

## Stack: LIFO (Last In, First Out)

### What it is
```
Push(1) → [1]
Push(2) → [1, 2]
Push(3) → [1, 2, 3]
Pop     → returns 3, stack: [1, 2]  ← last pushed, first popped
```

### Key operations
- **Push**: Add to top — O(1)
- **Pop**: Remove from top — O(1)
- **Peek**: View top without removing — O(1)
- **IsEmpty**: Check if empty — O(1)

### When to use
- Function call stack (recursion)
- Undo/Redo functionality
- Valid parentheses checking
- Monotonic stack patterns (Daily Temperatures)

---

## MinStack Problem — Solution Pattern

### The Problem
Standard stack doesn't support "get minimum" in O(1). You need to store minimum at each state.

### The Solution: Two Stacks in Parallel
```
Push(5):
  Main:  [5]       Min: [5]
  
Push(3):
  Main:  [5, 3]    Min: [5, 3]
  
Push(7):
  Main:  [5, 3, 7] Min: [5, 3, 3]  ← 3 again, not 7
  
Push(1):
  Main:  [5,3,7,1] Min: [5,3,3,1]
```

### Key insight
Always push to min stack:
- If new value < current min → push new value
- Otherwise → push current min again (maintain alignment)

When you pop, pop from both stacks. Top of min stack = current minimum.

### Why it works
- Both stacks grow/shrink together
- Min stack's top always tells you the minimum at that point in history
- Time: O(1) for all operations
- Space: O(n) for two stacks

---

## Daily Temperatures — Monotonic Stack Pattern

### The Problem
For each day, find how many days until a warmer day.

```
Input:  [73, 74, 75, 71, 69, 72, 76, 73]
Output: [ 1,  1,  4,  2,  1,  1,  0,  0]

Day 0 (73°): next warmer is day 1 (74°) → wait 1
Day 2 (75°): next warmer is day 6 (76°) → wait 4
Day 6 (76°): no warmer day → 0
```

### The Solution: Monotonic Stack
```
Key idea: Traverse RIGHT TO LEFT (backwards)
Maintain stack of indices in DECREASING temperature order
```

**Algorithm:**
```
For each day from right to left:
  1. Pop from stack all days COLDER than current
     (they can never be the "next warmer" day)
  2. Top of stack = next warmer day (if exists)
  3. Record distance
  4. Push current day onto stack
```

**Example trace:**
```
Temps: [73, 74, 75, 71, 69, 72, 76, 73]
Index:  0   1   2   3   4   5   6   7

Start from right (index 7):
i=7 (73°): Stack empty → result[7]=0, push 7
           Stack: [7]

i=6 (76°): Pop 7 (76>73) → Stack empty → result[6]=0, push 6
           Stack: [6]

i=5 (72°): 76>72, found! → result[5]=6-5=1, push 5
           Stack: [6, 5]

i=4 (69°): Pop 5 (72>69), pop 6 (76>69) → Stack empty → result[4]=0, push 4
           Stack: [4]

...and so on
```

### Why it works
- Only "warmer" days stay in stack
- Stack is always in decreasing order
- Time: O(n) — each element pushed/popped once
- Space: O(n) — stack size

### Senior insight
> "Monotonic stack is used when you need to find the next/previous greater/smaller element. By maintaining decreasing order, we ensure we find the answer in O(n) instead of O(n²)."

---

## Queue: FIFO (First In, First Out)

### What it is
```
Enqueue(1) → [1]
Enqueue(2) → [1, 2]
Enqueue(3) → [1, 2, 3]
Dequeue    → returns 1, queue: [2, 3]  ← first enqueued, first dequeued
```

### Key operations
- **Enqueue**: Add to back — O(1)
- **Dequeue**: Remove from front — O(1) (but O(n) with slice, use linked list for O(1))
- **Peek**: View front without removing — O(1)
- **Empty**: Check if empty — O(1)

### When to use
- BFS (breadth-first search)
- Task scheduling
- Level-order tree traversal
- Sliding window problems (deque)

---

## English & Communication

### Correct terminology
❌ "Put in the stack"
✅ "Push onto the stack"

❌ "Take from stack"
✅ "Pop from the stack"

❌ "Add to queue"
✅ "Enqueue into the queue"

❌ "Remove from queue"
✅ "Dequeue from the queue"

### Explaining MinStack in interview
> "MinStack uses two stacks in parallel. The main stack stores all values. The min stack stores the minimum value at each state. When we push a new value, we compare it with the current minimum: if it's smaller, we push the new value; otherwise, we push the current minimum again. This way, the top of the min stack always gives us the current minimum in O(1) time."

### Explaining DailyTemperatures in interview
> "We use a monotonic decreasing stack to solve this in O(n) time. We traverse the temperatures array from right to left. For each temperature, we pop all indices with smaller temperatures (they can't be the next warmer day). The top of the stack, if it exists, is the index of the next warmer day. We record the distance and push the current index onto the stack."

