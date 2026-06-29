# Binary Search — User Level Understanding (Cycle 1)

## The Problem

You have a sorted array of 1 million numbers. Need to find if a number exists.
- Linear search: check every element → 1,000,000 checks worst case (O(n))
- Binary search: cut in half each step → 20 checks worst case (O(log n))

## How It Works

1. Set `left = 0`, `right = len - 1`
2. Calculate `mid = left + (right - left) / 2`
3. Compare `arr[mid]` with target:
   - Equal → found, return mid
   - Target > mid → `left = mid + 1` (search right half)
   - Target < mid → `right = mid - 1` (search left half)
4. Repeat while `left <= right`
5. If loop ends → target not found, return -1

## Example

```
[1, 3, 5, 7, 9, 11, 13, 15]    Target = 13

Step 1: left=0, right=7, mid=3 → arr[3]=7 < 13 → go right
Step 2: left=4, right=7, mid=5 → arr[5]=11 < 13 → go right
Step 3: left=6, right=7, mid=6 → arr[6]=13 → FOUND! Return 6.
```

## Key Rules

### mid + 1 / mid - 1
Always skip mid when moving left or right. You already checked mid — including it again can cause infinite loops.

### Loop condition: `left <= right`
When `left > right`, search space is empty. Target doesn't exist. Return -1.

### Integer overflow prevention
Use `mid = left + (right - left) / 2` instead of `mid = (left + right) / 2`.
If left and right are both large, `left + right` can overflow the integer limit.

### Only works on sorted data
Binary search requires the array to be sorted. If not sorted, sort first or use linear search.

## When Target Not Found

```
[1, 3, 5, 7, 9, 11, 13, 15]    Target = 6

Step 1: left=0, right=7, mid=3 → 7 > 6 → go left
Step 2: left=0, right=2, mid=1 → 3 < 6 → go right
Step 3: left=2, right=2, mid=2 → 5 < 6 → go right
Step 4: left=3, right=2 → left > right → STOP → return -1
```

## Time and Space Complexity

- **Time:** O(log n) — halving the search space each step
- **Space:** O(1) iterative, O(log n) recursive (call stack)

## Senior-Level Answers

**"What is binary search?"**
> "Binary search finds a target in a sorted array by repeatedly cutting the search space in half. Compare the middle element with the target — if smaller go right, if larger go left. Time is O(log n). It only works on sorted data."

**"Why mid = left + (right - left) / 2?"**
> "To avoid integer overflow. If left and right are both large, adding them can exceed the integer limit. Subtracting first keeps the intermediate value small."

**"How does binary search terminate?"**
> "The loop runs while left <= right. Each step moves left or right past mid, shrinking the search space. When left exceeds right, the space is empty — target doesn't exist, return -1."
