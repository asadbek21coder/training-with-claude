# DSA / Live Coding — Patternlar (O'zbekcha) + Go misollari

Live coding'da algoritmni topish emas, **fikrlash jarayonini ovoz chiqarib aytish** muhim.

## Live coding jarayoni (doim shunday qil)
1. **Savolni takrorla** — "Ya'ni sizga kerak..." (tushunganingni tasdiqla)
2. **Edge case'lar so'ra** — bo'sh input? manfiy? duplicate?
3. **Brute force ayt** — "Eng sodda yechim O(n²)..."
4. **Optimallashtir** — "Lekin hash map bilan O(n) qilsa bo'ladi..."
5. **Kod yoz** — sekin, ovoz chiqarib
6. **Test qil** — misol bilan qo'lda tekshir
7. **Complexity ayt** — time va space

> ⚠️ Jim yozmang. Fikringizni ovoz chiqarib ayting — intervyuchi shuni baholaydi.

---

# Eng ko'p uchraydigan patternlar

## 1. Hash Map (eng ko'p!)
Two Sum, duplicate topish, frequency count, grouping.
```go
// Two Sum: nums ichida target beradigan 2 indeks
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int) // qiymat -> indeks
    for i, n := range nums {
        if j, ok := seen[target-n]; ok {
            return []int{j, i}
        }
        seen[n] = i
    }
    return nil
}
// Time O(n), Space O(n)
```

## 2. Two Pointers
Saralangan massiv, palindrome, pair topish.
```go
// Saralangan massivda target juftlik
func twoSumSorted(a []int, target int) (int, int) {
    l, r := 0, len(a)-1
    for l < r {
        sum := a[l] + a[r]
        if sum == target { return l, r }
        if sum < target { l++ } else { r-- }
    }
    return -1, -1
}
```

## 3. Sliding Window
Substring, subarray, "eng uzun/qisqa" masalalar.
```go
// Takrorlanmaydigan eng uzun substring uzunligi
func lengthOfLongestSubstring(s string) int {
    last := make(map[byte]int)
    start, best := 0, 0
    for i := 0; i < len(s); i++ {
        if j, ok := last[s[i]]; ok && j >= start {
            start = j + 1
        }
        last[s[i]] = i
        if i-start+1 > best { best = i - start + 1 }
    }
    return best
}
```

## 4. BFS / DFS (graf, daraxt)
```go
// Binary tree level-order (BFS)
func levelOrder(root *TreeNode) [][]int {
    var res [][]int
    if root == nil { return res }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        var level []int
        next := []*TreeNode{}
        for _, node := range queue {
            level = append(level, node.Val)
            if node.Left != nil { next = append(next, node.Left) }
            if node.Right != nil { next = append(next, node.Right) }
        }
        res = append(res, level)
        queue = next
    }
    return res
}
```

## 5. Recursion / DFS (backtracking)
Permutations, subsets, kombinatsiyalar.

## 6. Binary Search
Saralangan massivda qidirish, "eng kichik shart bajaradigan qiymat".
```go
func search(a []int, target int) int {
    lo, hi := 0, len(a)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2 // overflow'dan himoya
        if a[mid] == target { return mid }
        if a[mid] < target { lo = mid + 1 } else { hi = mid - 1 }
    }
    return -1
}
```

---

# Complexity cheat sheet
| Struktura | Access | Search | Insert | Delete |
|-----------|--------|--------|--------|--------|
| Array | O(1) | O(n) | O(n) | O(n) |
| Hash Map | - | O(1) | O(1) | O(1) |
| Balanced BST | O(log n) | O(log n) | O(log n) | O(log n) |
| Heap | - | O(n) | O(log n) | O(log n) |

- Nested loop → O(n²)
- Divide & conquer / binary search → O(log n) yoki O(n log n)
- Recursion → space = daraxt chuqurligi (call stack)

---

# Go-ga xos live-coding eslatmalar
- `make(map[K]V)` — map'ni ishlatishdan oldin yarat (nil map'ga yozish = panic).
- `append` slice'ni qaytaradi: `s = append(s, x)`.
- String'ni byte bo'yicha: `s[i]` (byte), rune uchun `[]rune(s)` yoki `range`.
- Slice teskari, sort: `sort.Ints(a)`, `sort.Slice(a, func(i,j int) bool {...})`.
- Integer bo'lish butun: `7/2 == 3`.
- Overflow'dan: `mid := lo + (hi-lo)/2`.

## 🎯 Agar qiynalsangiz (senior signal)
Yechimni topolmasangiz ham, ovoz chiqarib: *"Let me think about the brute force first, then optimize"* — jarayonni ko'rsatish natijadan muhimroq. Intervyuchi **qanday fikrlashingizni** baholaydi.
