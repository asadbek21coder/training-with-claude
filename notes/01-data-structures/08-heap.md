# Heap / Priority Queue — User Level Understanding (Cycle 1)

## The Problem

Need to always get the highest (or lowest) priority item quickly.
- Unsorted array: insert O(1), find min O(n)
- Sorted array: find min O(1), insert O(n)
- Heap: insert O(log n), find min O(1) — best of both

## What is a Heap?

A binary tree with two rules:

1. **Complete tree** — every level full except last, which fills left to right
2. **Heap property:**
   - Min heap: parent always smaller than children (smallest at root)
   - Max heap: parent always bigger than children (biggest at root)

```
Min Heap:          Max Heap:
    1                  9
   / \                / \
  3   5              7   5
 / \                / \
4   8              3   1
```

## Key Terms

- **Priority Queue** — the concept: "give me the most important item"
- **Heap** — the data structure that implements a priority queue
- **Heapify** — rearrange elements to satisfy heap property
- **Bubble up** — after insert, move new element up until heap property restored
- **Bubble down** — after removing root, move replacement down until heap property restored

## Insert (Bubble Up)

1. Place new element at next available spot (left to right on last level)
2. Compare with parent — if smaller (min heap), swap
3. Repeat until heap property satisfied or reached root

```
Insert 2 into:       Step 1 (place):     Step 2 (bubble up):
    1                     1                     1
   / \                   / \                   / \
  3   5                 3   5                 3   2
 / \                   / \ /                 / \ /
4   8                 4  8 2                4  8 5
```

## Remove Min/Max (Bubble Down)

1. Remove root (the min/max)
2. Move last element to root
3. Compare with children — swap with smaller child (min heap)
4. Repeat until heap property satisfied

## Time Complexity

- **Get min/max:** O(1) — always at root
- **Insert:** O(log n) — bubble up at most height of tree
- **Remove min/max:** O(log n) — bubble down at most height of tree

## Senior-Level Answers

**"What is a heap?"**
> "A heap is a complete binary tree where every parent is smaller (min heap) or larger (max heap) than its children. The root always holds the min or max. Insert and remove are O(log n), and getting the min/max is O(1)."

**"When do you use a heap?"**
> "When you need to repeatedly get the smallest or largest element. Common uses: priority queues, finding the kth largest element, merging k sorted lists, and scheduling tasks by priority."
