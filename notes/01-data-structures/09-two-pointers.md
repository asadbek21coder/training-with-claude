# Two Pointers — Cycle 1

## The Problem It Solves

Naive pair-finding = two nested loops = O(n²).
Two pointers = O(n) by walking smartly.

Наивный поиск пары = два вложенных цикла = O(n²).
Два указателя = O(n), проходим умно.

## Two Types / Два типа

### Type 1: Opposite direction (навстречу)
Start at both ends, move toward each other. Requires sorted array.
Начинают с разных концов, двигаются навстречу. Нужен отсортированный массив.

```
[1, 3, 5, 7, 9, 11]  target=14
 ^                ^
 left            right

sum < target → left++
sum > target → right--
sum == target → found!
```

**Common problems:** Two Sum sorted, container with most water, palindrome check, three sum.

### Type 2: Same direction (в одну сторону)
Both pointers move forward. One slow, one fast.
Оба двигаются вперёд. Один медленный, один быстрый.

**Common problems:** Remove duplicates in place, move zeros, linked list cycle.

## Key Rules / Правила

- `for left < right` — stop when pointers meet
- Calculate comparison INSIDE the loop (not before!)
- Return indices, not values (unless asked for values)
- For strings: clean first (remove non-alphanumeric), then two pointers

## When to Use / Когда использовать

| Use when / Используй когда | Example |
|---|---|
| Sorted array + find pair | Two Sum sorted |
| Compare from both ends | Palindrome |
| Modify array in place | Remove duplicates |
| Slow/fast traversal | Linked list cycle |

## Senior Answer / Ответ на собеседовании

**EN:** "Two pointers reduces O(n²) pair-finding to O(n) by using two variables that traverse the array intelligently. For sorted arrays, place pointers at both ends and move inward based on the comparison. The key insight is that on a sorted array, moving left pointer right increases the sum, and moving right pointer left decreases it."

**RU:** "Два указателя сокращают O(n²) до O(n), умно проходя по массиву. На отсортированном массиве ставим указатели с обоих концов и двигаем их внутрь на основе сравнения. Ключевая идея: сдвиг левого указателя вправо увеличивает сумму, сдвиг правого влево — уменьшает."
