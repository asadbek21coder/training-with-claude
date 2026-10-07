# Array / Slice — Complete Notes

---

## What is an array?

Imagine a row of numbered lockers in a school hallway:

```
 [0]  [1]  [2]  [3]  [4]
 ┌──┐ ┌──┐ ┌──┐ ┌──┐ ┌──┐
 │7 │ │3 │ │9 │ │1 │ │5 │
 └──┘ └──┘ └──┘ └──┘ └──┘
```

- Every locker has a fixed number (index)
- All lockers are side by side in memory (contiguous)
- You can open any locker directly if you know its number → O(1) access

Fixed size. You decide upfront how many lockers you need.

---

## The problem with fixed arrays

Real programs don't know in advance how many items they'll store.
So most languages add a dynamic array on top — in Go it's called a **slice**.

---

## How a slice works internally

A slice is not the data itself. It is a **3-field struct** that points to data:

```
Slice header (what you hold in your variable):
┌─────────────┬─────┬─────┐
│  ptr        │ len │ cap │
│ (points to  │  3  │  5  │
│  backing    │     │     │
│  array)     │     │     │
└─────────────┴─────┴─────┘
       │
       ▼
Backing array (actual memory):
┌──┐ ┌──┐ ┌──┐ ┌──┐ ┌──┐
│7 │ │3 │ │9 │ │  │ │  │
└──┘ └──┘ └──┘ └──┘ └──┘
 [0]  [1]  [2]  [3]  [4]
           ↑len=3     ↑cap=5
```

- **ptr** — pointer to where data lives in memory
- **len** — how many elements are currently in use
- **cap** — total space reserved in the backing array

---

## What happens when you append?

### Case 1 — cap has room (len < cap) → O(1)

Just writes into the next slot. No allocation.

```
Before: len=3, cap=5
┌──┐ ┌──┐ ┌──┐ ┌──┐ ┌──┐
│7 │ │3 │ │9 │ │  │ │  │
└──┘ └──┘ └──┘ └──┘ └──┘

After: len=4, cap=5
┌──┐ ┌──┐ ┌──┐ ┌──┐ ┌──┐
│7 │ │3 │ │9 │ │99│ │  │
└──┘ └──┘ └──┘ └──┘ └──┘
```

### Case 2 — cap is full (len == cap) → O(n)

Go allocates a new bigger array (roughly doubles cap), copies everything, writes new element.
Old backing array has no pointers left → garbage collector frees it.

```
Before: len=5, cap=5 (FULL)
┌──┐ ┌──┐ ┌──┐ ┌──┐ ┌──┐
│7 │ │3 │ │9 │ │1 │ │5 │
└──┘ └──┘ └──┘ └──┘ └──┘

After: new array allocated, all copied, new element added
┌──┐ ┌──┐ ┌──┐ ┌──┐ ┌──┐ ┌──┐ ┌──┐ ┌──┐ ┌──┐ ┌──┐
│7 │ │3 │ │9 │ │1 │ │5 │ │99│ │  │ │  │ │  │ │  │
└──┘ └──┘ └──┘ └──┘ └──┘ └──┘ └──┘ └──┘ └──┘ └──┘
len=6, cap=10
```

This is why append is **amortized O(1)** — the expensive copy happens rarely and averages out.

---

## The slice sharing trap

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]   // b = [2, 3]
```

`b` is NOT a copy. It **shares the same backing array** as `a`:

```
a: ptr─────────────┐ len=5, cap=5
                   ▼
         ┌──┐ ┌──┐ ┌──┐ ┌──┐ ┌──┐
         │1 │ │2 │ │3 │ │4 │ │5 │
         └──┘ └──┘ └──┘ └──┘ └──┘
b: ptr────────┘ len=2, cap=4
  (starts at index 1 of same array)
```

`b[0] = 99` also changes `a[1]`. Common bug. Common interview question.

Fix: use `copy()` to make an independent slice.

```go
b := make([]int, 2)
copy(b, a[1:3])   // b is now independent
```

---

## Passing slice to a function

Go passes slice header by value — the ptr, len, cap are copied into a local variable:

```go
func modify(s []int) {
    s[0] = 99          // changes caller's backing array — shared
    s = append(s, 100) // only updates local header — caller does not see this
}

a := []int{1, 2, 3}
modify(a)
fmt.Println(a)  // [99 2 3]
```

```
BEFORE modify():
a:        ptr──► [1|2|3]   len=3 cap=3
func's s: ptr──► [1|2|3]   len=3 cap=3  ← copy of header, same backing array

AFTER s[0] = 99:
a:        ptr──► [99|2|3]  ← visible (shared backing array)
func's s: ptr──► [99|2|3]

AFTER append (cap full → new array):
a:        ptr──► [99|2|3]       len=3 cap=3  ← UNCHANGED
func's s: ptr──► [99|2|3|100]   len=4 cap=6  ← only local header updated
```

---

## Go slice vs Python/Java reference type

In Python, length lives inside the shared list object. Appending from inside a function
mutates the shared object — caller sees the new length.

In Go, length lives in the slice header which is copied by value.
Element mutations are visible (shared backing array), structural changes (append) are not.

```
Python:
a ──────────────► LIST OBJECT { len=4, data=[1,2,3,100] }
func's lst ──────►             ↑ same object, both see len=4

Go:
a:        ptr──► [1|2|3]   len=3  ← never updated
func's s: ptr──► [1|2|3]   len=3  ← gets its own len, updated to 4 after append
```

---

## Concurrency — slice is NOT safe

Slice backing array is flat memory. No internal pointer structures. No detection flag.

```go
s := []int{0, 0, 0}
go func() { s[0]++ }()  // READ → ADD → WRITE
go func() { s[0]++ }()  // race condition

// result: s[0] = 1 instead of 2 — lost increment
// NO crash, NO error — silently wrong data
```

No hashWriting-style flag. No detection. Silent wrong results.
More dangerous than map because you never know something is wrong.

Fix:

```go
var mu sync.Mutex
mu.Lock()
s[0]++
mu.Unlock()
```

---

## Why arrays are cache-friendly

CPU reads a cache line (~64 bytes) at once. Array elements are contiguous in memory —
accessing `arr[0]` loads `arr[1]`, `arr[2]`... into cache automatically. Next accesses are nearly instant.

A linked list has each node at a random memory location. Every node access is a new memory fetch.
That is why arrays are faster in practice even when Big O looks the same.

---

## Complexity

| Operation | Time | Why |
|-----------|------|-----|
| Read by index | O(1) | Direct address calculation |
| Write by index | O(1) | Same |
| Append (with room) | O(1) | Write to next slot |
| Append (full) | O(n) | Copy to new array |
| Append amortized | O(1) | Rare copy averages out |
| Search (unsorted) | O(n) | Check every element |
| Search (sorted) | O(log n) | Binary search |
| Insert at middle | O(n) | Must shift elements right |
| Delete at middle | O(n) | Must shift elements left |

---

## Go vs Java

```go
// Go
s := []int{1, 2, 3}
s = append(s, 4)
b := make([]int, len(s))
copy(b, s)
```

```java
// Java
List<Integer> list = new ArrayList<>(Arrays.asList(1, 2, 3));
list.add(4);
List<Integer> copy = new ArrayList<>(list);
```

---

# Interview Q&A

---

## Q1: What is a slice in Go internally?

> "A slice in Go is a struct with three fields: a pointer to the underlying backing
> array, `len` which is how many elements are currently in use, and `cap` which is
> the total capacity of the backing array. The slice header is copied by value when
> passed around — the pointer inside it is what allows shared access to the same data."

---

## Q2: What happens when you append to a slice?

> "There are two cases. If `len < cap`, there is room in the backing array — Go writes
> the new element into the next slot and increments `len`. No allocation happens.
> If `len == cap`, the backing array is full — Go allocates a new larger array,
> copies all existing elements, updates the pointer to the new array, then writes the
> new element. The old backing array has no references left so the garbage collector
> frees it. This is why append is amortized O(1) — the expensive copy happens rarely."

---

## Q3: What does this print and why?

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]
b[0] = 99
fmt.Println(a)
```

> "`[1 99 3 4 5]`. `a[1:3]` creates a new slice header pointing to index 1 of `a`'s
> backing array — no data is copied. Both `a` and `b` share the same underlying array.
> `b[0] = 99` writes to the same memory cell as `a[1]`, so `a` reflects the change."

---

## Q4: How do you make a truly independent copy of a slice?

> "Use the built-in `copy()` function. Unlike sub-slicing which shares the backing array,
> `copy()` copies elements into a new independent backing array. After that, modifying
> one slice has no effect on the other."

```go
b := make([]int, 2)
copy(b, a[1:3])
```

---

## Q5: What happens when you pass a slice to a function and append inside it?

> "Go passes the slice header by value. The pointer inside the header still points to
> the same backing array, so element mutations like `s[0] = 99` are visible to the
> caller. But if append triggers a new allocation, only the local header is updated —
> the caller's header still points to the old array and does not see the new element."

Output of the example: `[99 2 3]` — mutation visible, append not visible.

---

## Q6: How is a Go slice different from a reference type in Python or Java?

> "In Python, length is stored inside the shared list object — appending from inside a
> function mutates the shared object and the caller sees the new length. In Go, length
> is part of the slice header which is copied by value. Element mutations are visible
> because both share the backing array, but structural changes like append only update
> the local header copy — the caller's header is never touched."

---

## Key nuances to always mention

1. Slice header is copied by value — not passed by reference; the pointer inside is what's shared
2. Sub-slice shares backing array — `a[1:3]` is not a copy; mutations affect the original
3. `copy()` makes a truly independent slice
4. Amortized O(1) append — occasionally O(n) copy averages out
5. Old array is GC'd after reallocation — no references left, garbage collector frees it
6. Slice is NOT concurrent safe — silent wrong results, no crash, no detection flag
