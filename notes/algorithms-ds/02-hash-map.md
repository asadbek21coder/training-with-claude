# Hash Map — Complete Notes

---

## What is a bucket?

Imagine a post office sorting letters. You have 10 bins (0–9).
Every letter — look at the last digit of the zip code, throw it in that bin.

```
Zip 94102 → last digit 2 → bin[2]
Zip 10001 → last digit 1 → bin[1]
Zip 33009 → last digit 9 → bin[9]
```

Those bins are **buckets**. The rule "take last digit" is the **hash function**.

In a hash map, buckets are slots in an array:

```
index:  [0]   [1]   [2]   [3]   [4]   [5]   [6]   [7]
        empty  bob  alice  empty empty empty empty  carol
```

When you do `map["alice"] = 30`:
1. Hash function runs on `"alice"` → produces a big number
2. `hash % bucketCount` → bucket index
3. bucket stores `alice → 30`

Lookup uses the same hash → same bucket → O(1) direct access. No searching.

---

## What is a collision?

```
hash("alice") % 8 = 2  → bucket 2
hash("zara")  % 8 = 2  → bucket 2 also ← COLLISION
```

Two different keys. Same bucket. Cannot overwrite one with the other.

---

## Chaining — how collisions are resolved

Each bucket holds a linked list of all entries that landed there:

```
bucket[2]: ──► [alice|30] ──► [zara|99] ──► nil
```

Lookup "alice":
1. hash("alice") % 8 = 2 → go to bucket 2
2. Walk list: "alice"? Yes → return 30

Full map view:

```
┌─────────┬─────────┬─────────┬─────────┬─────────┐
│bucket[0]│bucket[1]│bucket[2]│bucket[3]│bucket[4]│
└─────────┴─────────┴─────────┴─────────┴─────────┘
               │               │
             [bob|200]       [alice|30]──►[zara|99]
```

---

## Load factor — why it matters

```
load factor = number of entries / number of buckets
```

| Entries | Buckets | Load | State |
|---------|---------|------|-------|
| 3 | 10 | 0.3 | Healthy — chains are short |
| 9 | 10 | 0.9 | Crowded — chains are long, lookups slow |

10 buckets, 100 entries → each chain has ~10 items → lookup walks 10 nodes → O(10), not O(1).

---

## Bucket count — always a power of 2

Go: starts with 8 buckets. Java: starts with 16 buckets.
Bucket count is always a power of 2 (8, 16, 32, 64...).

**Why power of 2?**

Normal modulo: `hash % N` → division → ~20–40 CPU cycles

When N is a power of 2: `hash & (N-1)` → bitwise AND → 1 CPU cycle

They produce the same result. AND is 20–40× faster.

This works because in binary, `2^B - 1` is B ones — a perfect bitmask:

```
2^3 = 8  = 1000
2^3 - 1  = 0111  ← 3 ones = mask for last 3 bits

hash      = 1011 0101  (181)
mask      = 0000 0111  (7)
AND       = 0000 0101  (5)  = 181 % 8 ✓
```

Mathematical reason: any number x = quotient × 2^B + remainder.
Multiples of 2^B always end in B zeros in binary.
So remainder always lives in the last B bits.
AND with B ones extracts exactly those bits = the remainder.

---

## Bit masking

A mask is a number used in a bitwise operation to selectively affect specific bits.

```
AND  &    extract or clear bits   (1=keep, 0=erase)
OR   |    set bits                (1=force 1, 0=leave)
XOR  ^    toggle bits             (1=flip, 0=leave)
```

AND with 1 keeps a bit exactly as it is. That is why AND is used for extraction.
OR with 1 forces bit to 1 — it overwrites, does not read. Cannot extract with OR.

**Building a mask for bit n:**

```go
1 << 0 = 0000 0001  ← bit 0
1 << 3 = 0000 1000  ← bit 3
1 << 7 = 1000 0000  ← bit 7
```

**Four fundamental operations:**

```go
// Check if bit 3 is set
if x & (1 << 3) != 0 { }

// Set bit 3 to 1
x = x | (1 << 3)

// Clear bit 3 to 0
x = x &^ (1 << 3)   // &^ is Go's AND NOT

// Toggle bit 3
x = x ^ (1 << 3)
```

**Real examples:**

```go
// even/odd check
x & 1 == 0  // even
x & 1 == 1  // odd

// extract RGB from packed color
r := (color >> 16) & 0xFF
g := (color >> 8)  & 0xFF
b := (color >> 0)  & 0xFF

// Go's hashWriting flag
flags |=  hashWriting  // set
flags &^= hashWriting  // clear
flags &   hashWriting  // check
```

---

## Go's hash function

Two separate steps:

```
Step 1: key → hash function → 64-bit integer
Step 2: 64-bit integer & (bucketCount - 1) → bucket index
```

Step 1 is not simple. Go picks hash function by key type at compile time:

```
string key → AES hardware hash (if CPU has AES-NI instruction)
              or wyhash (software fallback)
int key    → fast integer mixing function
struct key → hash each field, combine results
```

**Critical: hash function is seeded randomly at program startup**

```go
// Go runtime does this when your program starts:
seed := randomBytes()  // different every run
hashFn = newHasher(seed)
```

Why? To prevent **HashDoS attacks** — if an attacker knows your hash function,
they craft thousands of keys that all hash to the same bucket, turning every lookup O(n)
and killing your server. Random seed means the attacker cannot predict collisions.

This is also why map iteration order is randomized — same key, different bucket each run.

---

## How Go calculates bucket count from a hint

```go
m := make(map[string]int, 1_000_000)
```

Go stores bucket count as B — the number of bits. Actual count = 2^B.

Algorithm (from runtime/map.go):

```
B = 0
loop:
    if 2^B × 6.5 >= hint → stop, B is large enough
    B++
```

Example with hint = 1,000,000:

```
B=17: 2^17 = 131,072 buckets → 131,072 × 6.5 = 851,968  < 1,000,000 → too small
B=18: 2^18 = 262,144 buckets → 262,144 × 6.5 = 1,703,936 >= 1,000,000 → STOP
```

Result: 262,144 buckets allocated — can hold 1,703,936 entries before next rehash.
Go always rounds up to next power of 2, so your hint count will never trigger a rehash.

---

## Bucket internal structure

Each bucket is a fixed-size struct — not a linked list node per entry:

```
one bucket (holds up to 8 entries):
┌──────────────────────────────────────┐
│ tophash [8]uint8  ← first 8 bits    │
│                     of each key hash │
│ keys    [8]K      ← 8 key slots      │
│ values  [8]V      ← 8 value slots    │
│ overflow *bucket  ← chain to next    │
└──────────────────────────────────────┘
```

`tophash` is a fast pre-filter — Go checks 8 bytes first before comparing full keys.
Avoids expensive key comparisons when hash prefix does not match.

---

## When and how does resize happen?

### Go — threshold: 6.5 entries per bucket on average

```
8  buckets × 6.5 → resize after ~52  entries → new size: 16 buckets
16 buckets × 6.5 → resize after ~104 entries → new size: 32 buckets
```

### Java — threshold: 0.75 × capacity

```
16 buckets × 0.75 → resize after 12 entries → new size: 32 buckets
32 buckets × 0.75 → resize after 24 entries → new size: 64 buckets
```

### What happens during resize:

```
BEFORE (4 buckets, getting full):
bucket[0]: ──► [alice|30] ──► [eve|10]
bucket[1]: ──► [bob|200]
bucket[3]: ──► [carol|25]

Step 1: allocate new array (8 buckets)
Step 2: re-hash every entry with new modulo (% 8 now, not % 4)

hash("alice") % 8 = 4  ← moved
hash("eve")   % 8 = 0  ← stays
hash("bob")   % 8 = 1  ← stays
hash("carol") % 8 = 7  ← moved

AFTER (8 buckets):
bucket[0]: ──► [eve|10]
bucket[1]: ──► [bob|200]
bucket[4]: ──► [alice|30]
bucket[7]: ──► [carol|25]
```

---

## Go vs Java — incremental vs all-at-once resize

**Go — incremental rehashing:**
- Old and new arrays both exist temporarily
- Each get/put moves a few entries from old → new gradually
- Cost spread across many operations — no single big pause
- Old array GC'd after all entries are moved

**Java — all at once:**
- When threshold hit, one operation pays full O(n) cost
- One request bears the full rehash penalty

---

## nil map — read safe, write panics

```go
var m map[string]int   // nil map — no bucket array exists
v := m["key"]          // ✅ safe — returns 0
m["key"] = 1           // ❌ panic: assignment to entry in nil map
```

Go generates an explicit nil check before every map operation:

```
read:  if m == nil → return zero value immediately. STOP. No memory touched.
write: if m == nil → panic immediately. Never touches address 0.
```

Why reads are safe but writes panic:
- Read from nil → worst case: wrong zero value → acceptable, no data loss
- Write to nil → no bucket array exists → data loss → not acceptable

---

## nil in memory

nil is literally zero bytes — the number 0 stored in a pointer-sized slot:

```
nil in memory (8 bytes on 64-bit):
00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000
```

The OS intentionally keeps the first page of every process's virtual address space
unmapped (called the null page — no physical RAM backs address 0).

If any code tries to access address 0:
- CPU raises hardware page fault
- OS sends SIGSEGV to process
- Go converts to panic

This is why nil pointer dereferences crash — no nil check is generated for raw pointers.
Map operations get explicit nil checks. Raw pointer dereferences do not.

---

## Zeroing vs Initializing

**Zeroing** — memory set to zero. No structure created.

```go
var m map[string]int   // m = nil  (zero value of map is nil)
```

**Initializing** — actually allocating and setting up internal structure.

```go
m := make(map[string]int)  // bucket array created, ready to use
```

```
ZEROED (var):
m → nil   (nothing exists)

INITIALIZED (make):
m → ┌─────────────────┐
    │ bucket[0]: nil  │
    │ ...             │
    │ bucket[7]: nil  │
    └─────────────────┘
```

---

## make vs new

```go
// new — allocates + zeroes only. Returns pointer. Map inside is still nil.
m := new(map[string]int)   // type: *map[string]int
(*m)["key"] = 1            // still panics — map inside is nil

// make — allocates + initializes. Returns ready-to-use value.
m := make(map[string]int)  // type: map[string]int
m["key"] = 1               // works fine
```

Go's `new` ≠ Java's `new`. Java's `new` is closer to Go's `make`.

| | Go `new` | Java `new` |
|--|----------|------------|
| What it does | allocates + zeroes only | allocates + initializes |
| Map usable after? | ❌ map inside still nil | ✅ yes |

---

## Zero value on missing key

Accessing a missing key in an initialized map does NOT panic.
Go returns the zero value of the value type:

```go
m := make(map[string]int)
v := m["missing"]  // returns 0 — key not found, no panic, no error
```

This enables clean accumulator patterns:

```go
func wordCount(text string) map[string]int {
    result := make(map[string]int)
    for _, word := range strings.Fields(text) {
        result[word]++  // first encounter: reads 0, writes 1. No init needed.
    }
    return result
}
```

`result[word]++` expands to: `result[word] = result[word] + 1`
First time: reads 0 (missing key), adds 1, writes 1. No existence check needed.

---

## append to nil slice in map — common pattern

```go
m := make(map[string][]int)
m["scores"] = append(m["scores"], 1)  // m["scores"] is nil → append handles nil → []int{1}
m["scores"] = append(m["scores"], 2)  // m["scores"] is []int{1} → append → []int{1, 2}
m["scores"] = append(m["scores"], 3)
fmt.Println(m["scores"])  // [1 2 3]
```

`append` is designed to accept nil slices — it treats nil as empty, allocates a new
backing array, and returns a new slice. No initialization boilerplate needed.

---

## Iteration order — not guaranteed

```go
for k, v := range m {
    fmt.Println(k, v)  // order is different every run
}
```

Two reasons:

**1 — hash determines position, not insertion order.**
Keys land in buckets by hash value. Iteration goes bucket by bucket — hash order, not insertion order.

**2 — Go deliberately randomizes the starting bucket.**
Every range picks a random starting bucket and random starting slot within that bucket.
Same program, same map, different run → different order.

Why intentional? Before Go 1, order was consistent enough that developers started
depending on it. When map internals changed, that code broke silently. Go's response:
make it random always, so no code can accidentally depend on it.

**How all items are still visited — circular traversal:**

```
8 buckets, random start = 3:
visits: 3 → 4 → 5 → 6 → 7 → 0 → 1 → 2 → done
         ↑ wraps around, visits all 8 buckets exactly once
```

Start point changes order, not coverage. Every item is visited exactly once.

**Ordered iteration — extract and sort:**

```go
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
    fmt.Println(k, m[k])
}
```

---

## delete — no-op on missing key, no memory shrink

```go
delete(m, "a")  // exists — removes entry
delete(m, "z")  // does not exist — no-op, no panic, no error
```

**Map never shrinks after delete:**

```go
m := make(map[string]int)
for i := 0; i < 1_000_000; i++ {
    m[strconv.Itoa(i)] = i
}
// large bucket array allocated

for i := 0; i < 1_000_000; i++ {
    delete(m, strconv.Itoa(i))
}
// len(m) = 0
// BUT bucket array is still fully allocated — memory NOT released
```

Deleting entries marks slots as empty but keeps buckets allocated.
Memory stays at peak usage forever. No built-in shrink operation.

Fix — replace the map:

```go
m = make(map[string]int)  // old map becomes garbage, GC collects it
```

---

## Concurrency — map is NOT safe

### Two separate problems

**Problem 1 — logical race (general):**
`m["a"]++` is three steps: read → add → write. Steps interleave. Increments get lost.

**Problem 2 — structural corruption (map-specific):**
Map write involves multiple internal steps including potential rehashing which rewrites
bucket pointers. Two goroutines doing this simultaneously can corrupt pointer structures.

### What physically crashes

**Two simultaneous writes to same bucket:**

```
INITIAL:
bucket[2] ──► [bob|200] ──► nil

G1 writes m["alice"]: creates [alice|30], sets alice.Next = bob
G2 writes m["zara"]:  creates [zara|99],  sets zara.Next  = bob

G1: bucket[2].ptr = &alice
G2: bucket[2].ptr = &zara   ← overwrites G1

RESULT:
bucket[2] ──► [zara|99] ──► [bob|200] ──► nil
[alice|30] ──► [bob|200]   ← orphaned, lost forever
bob referenced by two nodes ← dangling pointer risk
```

**Read during rehash (worst case):**

```
G1 starts rehash:
  allocates new bucket array
  starts moving entries
  sets bucket[2].ptr = nil  ← clears old bucket mid-move

G2 reads m["alice"] now:
  goes to bucket[2]
  reads nil pointer
  follows nil → address 0 → CPU page fault → crash
```

Or G1 writes a 64-bit pointer in two 32-bit writes — G2 reads a half-written pointer,
follows a garbage address → crash.

### hashWriting flag — Go's safety net

```
map internal header:
┌──────────────────┐
│ buckets *bucket  │
│ count    int     │
│ flags    uint8   │  ← contains hashWriting bit
└──────────────────┘
```

G1 starts writing: `flags |= hashWriting`
G2 tries to access: if `flags & hashWriting != 0` → `fatal error: concurrent map read and map write`

Catches the obvious case. Not perfect — if both goroutines check the flag at the exact
same nanosecond, both pass and corrupt each other anyway.

### Why Mutex fixes it

Mutex does not make map safe. Map has no idea a mutex exists.
Mutex prevents two goroutines from **reaching map code simultaneously**.

```
WITHOUT mutex:
G1: ──────────────[map write]──────────────►
G2: ──────[map write]──────────────────────►
               ↑ overlap → corruption

WITH mutex:
G1: ──[lock]──[map write]──[unlock]──────────────►
G2: ──────────────────────[lock]──[map write]──►
                           ↑ blocked, no overlap
```

```go
mu.Lock()
m["a"]++   // covers all three steps: read + add + write
mu.Unlock()
```

### Three sync options

```go
// Option 1 — sync.Mutex (mixed reads and writes)
var mu sync.Mutex
mu.Lock()
m["a"]++
mu.Unlock()

// Option 2 — sync.RWMutex (reads >> writes)
var mu sync.RWMutex
mu.RLock(); v := m["a"]; mu.RUnlock()   // multiple readers simultaneously
mu.Lock(); m["a"]++; mu.Unlock()        // exclusive write

// Option 3 — sync.Map (stable keys, many readers)
var sm sync.Map
sm.Store("key", 1)
v, ok := sm.Load("key")
sm.Delete("key")
```

### Slice vs Map concurrency

| | Slice | Map |
|--|-------|-----|
| Concurrent safe? | ❌ | ❌ |
| Structural corruption risk | ❌ flat memory | ✅ pointer structures |
| Built-in detection | ❌ silent wrong results | ✅ hashWriting flag → panic |
| Crash on concurrent access? | usually not | yes |

Slice is more dangerous — wrong results with no signal. Map crashes loudly.

### Detect race conditions

```bash
go run -race main.go
go test -race ./...
```

Always run `-race` in development and CI.

---

## Go vs Java — full comparison

| | Go `map` | Java `HashMap` |
|--|----------|----------------|
| Default initial buckets | 8 | 16 |
| Load factor threshold | 6.5 entries/bucket avg | 0.75 (75% full) |
| Resize amount | double | double |
| Bucket count | always power of 2 | always power of 2 |
| Pre-size hint | `make(map[K]V, n)` | `new HashMap<>(n)` |
| Resize style | incremental | all at once |
| Collision strategy | chaining (linked list) | chaining → tree if > 8 entries |
| Worst case lookup | O(n) | O(log n) — Red-Black tree |
| Thread-safe? | ❌ | ❌ |
| Thread-safe alternative | `sync.Map` or `sync.RWMutex` | `ConcurrentHashMap` |
| Ordered? | ❌ random | ❌ random (use `LinkedHashMap`) |
| Key requirement | comparable (`==`) | implement `hashCode()` + `equals()` |

Java upgrades bucket chain to Red-Black tree at 8+ entries per bucket → O(log n) worst case.
Go always stays linked list → O(n) worst case.

---

## Go implementation

```go
m := make(map[string]int)
m := make(map[string]int, 1000)  // with size hint

m["alice"] = 30

v := m["alice"]              // 0 if not found, no panic
v, ok := m["alice"]          // ok = false if not found — prefer this form

delete(m, "alice")           // no-op if key does not exist

for k, v := range m {        // order is random every run
    fmt.Println(k, v)
}

// struct key — valid if all fields are comparable
type Point struct{ X, Y int }
m2 := map[Point]string{}
m2[Point{1, 2}] = "origin"

// slice key — invalid, slices are not comparable
m3 := map[[]int]string{}  // compile error

// concurrent safe options
var mu sync.RWMutex
var sm sync.Map
```

---

## Java implementation

```java
Map<String, Integer> m = new HashMap<>();
m.put("alice", 30);

int v = m.getOrDefault("alice", 0);  // safe default

m.remove("alice");

for (Map.Entry<String, Integer> e : m.entrySet()) {
    System.out.println(e.getKey() + " " + e.getValue());
}

Map<String, Integer> m = new ConcurrentHashMap<>();  // thread-safe
Map<String, Integer> m = new LinkedHashMap<>();      // insertion-ordered
Map<String, Integer> m = new TreeMap<>();            // sorted by key, O(log n)
```

---

# Interview Q&A

---

## Q1: What is a hash map internally? Walk through what happens on lookup.

> "A hash map is an array of buckets. When you look up a key, the runtime runs a hash
> function on the key to produce a large integer, then takes that integer modulo the
> bucket count to get a bucket index. It opens that bucket and walks the chain of
> entries stored there, comparing each key until it finds a match or reaches the end.
> Because the bucket is found in O(1) and chains are short under a healthy load factor,
> the overall operation is O(1) average."

---

## Q2: What is the zero value of a map in Go? What can you do with it?

> "The zero value of a map is nil — a pointer to nothing, no bucket array exists.
> Reading from a nil map is safe in Go — the compiler generates an explicit nil check
> and returns the zero value of the value type without touching memory. Writing to a
> nil map panics immediately — there is no bucket array to store the entry. To make a
> map writable, you must initialize it with `make`."

---

## Q3: What is load factor and why does it matter?

> "Load factor is the ratio of entries to buckets. When it grows too high, chains get
> long and lookups degrade toward O(n) instead of O(1). Go triggers a rehash when
> average entries per bucket exceeds 6.5 — it doubles the bucket array, re-hashes all
> entries into their new positions, and restores short chains. Java uses 0.75 as its
> threshold. Go rehashes incrementally across operations. Java pays the full O(n) cost
> in one operation."

---

## Q4: Can you use any type as a map key in Go?

> "No. Map keys must be comparable — types that support the `==` operator. Comparable
> types include integers, strings, booleans, pointers, arrays, and structs whose fields
> are all comparable. Slices, maps, and functions are not comparable and cannot be used
> as keys — the compiler enforces this at compile time. A struct is comparable if and
> only if all of its fields are comparable types."

---

## Q5: Can a struct be used as a map key?

> "Yes, but only if all its fields are comparable. Go compares structs field by field.
> If even one field is a slice, map, or function, the entire struct becomes non-comparable
> and the compiler rejects it as a map key. An array field is fine — arrays are comparable.
> A slice field is not — slices are not comparable."

```go
type Point struct{ X, Y int }         // ✅ all fields comparable
type Bad struct{ Tags []string }      // ❌ slice field — not comparable
type WithArray struct{ IDs [3]int }   // ✅ array (not slice) — comparable
```

---

## Q6: What does the size hint in make do?

> "The second argument to `make` is a hint telling the runtime to pre-allocate enough
> buckets to hold approximately that many entries without triggering a rehash. Without
> the hint the map starts with 8 buckets and rehashes multiple times growing to large
> sizes. With a hint, Go calculates the required bucket count upfront — specifically,
> it finds the smallest B such that 2^B × 6.5 >= hint, then allocates 2^B buckets.
> This eliminates potentially dozens of O(n) rehash cycles. The hint is a minimum
> guarantee — Go always rounds up to the next power of 2, so you will not rehash even
> if you insert slightly more than the hint."

---

## Q7: Why doesn't `result[word]++` panic when the key doesn't exist?

> "Accessing a missing key in an initialized map returns the zero value of the value
> type without panicking. For `int` that is 0. So `result[word]++` reads 0 for a new
> word, increments to 1, and writes the entry. This zero-value-on-missing-key behavior
> is intentional — it makes accumulator and grouping patterns idiomatic with no
> existence checks needed."

---

## Q8: Is map iteration order guaranteed in Go?

> "No, and for two reasons. First, keys are distributed by hash value not insertion order,
> so sequential bucket traversal produces hash-order output with no relation to insertion
> order. Second, Go deliberately randomizes the starting bucket on every range call —
> same program, same map, different run produces different order. This was an intentional
> decision to prevent developers from accidentally depending on a consistent order.
> Iteration still visits every item exactly once via circular traversal — start at a
> random bucket, visit all buckets in order, wrap around. Start point changes order,
> not coverage."

---

## Q9: Why does appending to a nil slice in a map work?

> "`append` is designed to accept nil slices — it treats nil as an empty slice with zero
> length and zero capacity, allocates a new backing array, and returns a new slice.
> Combined with Go's zero-value-on-missing-key behavior, `append(m[key], value)` reads
> nil for a missing key, appends to nil creating a new slice, and the assignment writes
> it back. No initialization boilerplate is needed."

---

## Q10: What does delete do on a non-existent key? Does the map release memory?

> "Deleting a non-existent key is a no-op — no panic, no error, by design. On the
> memory side, Go maps never shrink. Deleting entries marks slots as empty but the
> underlying bucket array stays fully allocated at its peak size. A map used as a
> cache that fills and clears will hold peak memory allocation indefinitely. The only
> way to recover that memory is to replace the map with a new one and let the garbage
> collector collect the old backing structure."

---

## Key sources

- `https://github.com/golang/go/blob/master/src/runtime/map.go` — actual implementation
- GopherCon 2016: "Inside the Map Implementation" by Keith Randall — best talk on Go maps
- `https://github.com/teh-cmc/go-internals` — deep runtime explanations
- "Hacker's Delight" by Henry S. Warren — bit manipulation math
