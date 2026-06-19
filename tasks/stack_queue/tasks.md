# Stack & Queue — Completed Tasks

## ✅ Task 1: IsValidParentheses — Using Stack

**Status**: Completed and tested

**Problem**: Check if parentheses/brackets are balanced

```
Input:  "({[]})"
Output: true

Input:  "({[}])"
Output: false
```

**Solution approach**: 
- Use stack to track opening brackets
- When you see closing bracket, check if it matches the top of stack
- At end, stack should be empty

**Implementation**: Two versions in `stack_queue.go`
- `IsValidParatheses`: Naive approach with string checking
- `IsValidParanthesesStack`: Using Stack struct

---

## ✅ Task 2: MinStack — O(1) GetMin

**Status**: Completed and tested

**Problem**: Design a stack where `GetMin()` returns minimum in O(1)

```
Push(5) → min: 5
Push(3) → min: 3
Push(7) → min: 3
Push(1) → min: 1
Pop()   → min: 3
Pop()   → min: 3
Pop()   → min: 5  ← must restore correctly!
```

**Solution approach**:
- Two stacks in parallel
- Min stack stores minimum at each state
- When pushing: if value < current min, push value; else push current min
- When popping: pop from both stacks

**Implementation**: `MinStack` struct with `Push`, `Pop`, `Peek`, `GetMin` in `stack_queue.go`

**Key insight**: Tracking history of minimums at each step

---

## ✅ Task 3: DailyTemperatures — Monotonic Stack

**Status**: Completed and tested

**Problem**: For each day, find days until a warmer day

```
Input:  [73, 74, 75, 71, 69, 72, 76, 73]
Output: [ 1,  1,  4,  2,  1,  1,  0,  0]
```

**Solution approach**:
- Monotonic decreasing stack
- Traverse right to left (backwards)
- Pop colder temperatures
- Top of stack = next warmer day

**Implementation**: `DailyTemperatures` function in `stack_queue.go`

**Key insight**: Monotonic stack reduces O(n²) brute force to O(n)

---

## ⏭️ Pending (Skipped for Now)

These were planned but we're skipping to move faster:

### Queue Using Two Stacks
- Design a queue using only two stacks
- Efficient enqueue (O(1)), amortized O(1) dequeue

### Sliding Window Maximum
- Deque pattern for maintaining maximum in sliding window
- O(n) using monotonic deque

**Decision**: Skip these to focus on Binary Tree (higher priority for interviews)

---

## Test Results

All three completed tasks tested in `main.go`:
- IsValidParanthesesStack: ✅ All test cases pass
- MinStack: ✅ Push/Pop/GetMin working correctly
- DailyTemperatures: ✅ Output matches expected

