# 06 — Operating Systems & Concurrency

> Backend engineers are asked about processes, threads, concurrency, and memory. Know both concepts and how Go/Java implement them.

---

## Process vs Thread

**Key Terms:**
- `process` — independent program in execution; has its own memory space (code, heap, stack, file descriptors)
- `thread` — unit of execution within a process; shares memory with other threads in same process
- `PID (Process ID)` — unique identifier for a process assigned by OS
- `TID (Thread ID)` — unique identifier for a thread
- `PCB (Process Control Block)` — OS data structure tracking process state (registers, open files, memory maps)
- `TCB (Thread Control Block)` — OS data structure tracking thread state
- `context switch` — OS saves state of current task, loads state of next task; expensive (microseconds)
- `user space` — memory area for application code
- `kernel space` — memory area for OS; privileged; system calls cross into kernel space
- `system call (syscall)` — program requests OS service (read file, allocate memory, create thread)
- `fork()` — creates new process as copy of parent (copy-on-write)
- `exec()` — replaces current process image with new program
- `zombie process` — process finished but parent hasn't read its exit status yet
- `orphan process` — parent died before child; adopted by init (PID 1)
- `IPC (Inter-Process Communication)` — pipes, sockets, shared memory, message queues

**Process vs Thread comparison:**

| | Process | Thread |
|--|---------|--------|
| Memory | Separate (isolated) | Shared within process |
| Creation cost | Expensive | Cheap |
| Communication | IPC (pipes, sockets) | Shared memory (fast, risky) |
| Crash isolation | One crash = one process | One crash can kill process |
| Context switch | Expensive | Cheaper |

**Must understand:**
- [ ] Why threads share memory but processes don't
- [ ] Why creating a thread is cheaper than a process
- [ ] What happens during a context switch
- [ ] Why sharing memory between threads is dangerous (race conditions)

---

## Concurrency vs Parallelism

**Key Terms:**
- `concurrency` — multiple tasks in progress at the same time (may not be simultaneous); about STRUCTURE
- `parallelism` — multiple tasks executing simultaneously on multiple CPU cores; about EXECUTION
- `single-core concurrency` — CPU switches between tasks (interleaving); illusion of parallelism
- `multi-core parallelism` — tasks literally run at the same time on different cores
- `goroutine` — Go's lightweight concurrent unit; M:N scheduled (many goroutines → few OS threads)
- `M:N threading model` — many lightweight threads mapped to N OS threads (Go's model)
- `1:1 threading model` — each thread = one OS thread (Java's traditional model)
- `Java virtual threads (Project Loom)` — Go-like M:N model in Java 21+
- `event loop` — single-threaded concurrency model (Node.js, Python asyncio)
- `async/await` — language feature for non-blocking I/O without callbacks
- `coroutine` — suspendable function; Kotlin coroutines, Python async generators

**Must understand:**
- [ ] Concurrency ≠ parallelism — two programs can be concurrent but not parallel
- [ ] Go's goroutines — lightweight, M:N, multiplexed onto OS threads by Go scheduler
- [ ] Java's traditional threads — expensive 1:1; Java 21 virtual threads are cheap like goroutines
- [ ] Event loop model — how Node.js is concurrent but single-threaded

---

## Race Conditions & Deadlocks

**Key Terms:**
- `race condition` — two goroutines/threads access shared data concurrently; result depends on timing
- `data race` — specific type of race: concurrent read+write without synchronization
- `critical section` — code section that accesses shared resource; must not run concurrently
- `mutual exclusion (mutex)` — only one thread in critical section at a time
- `Go race detector` — `go run -race` detects data races at runtime
- `deadlock` — two or more goroutines wait for each other forever; neither can proceed
- `livelock` — goroutines keep responding to each other but no progress (like two people in doorway)
- `starvation` — goroutine never gets access to resource because others always get priority
- `goroutine leak` — goroutine created but never terminates; grows unboundedly
- `Coffman conditions (deadlock)` — four conditions that must all hold for deadlock:
  1. Mutual exclusion
  2. Hold and wait (hold one resource while waiting for another)
  3. No preemption (resources not forcibly released)
  4. Circular wait (A waits for B, B waits for A)

**How to prevent deadlock:**
- Lock ordering — always acquire locks in same order
- Try-lock with timeout
- Avoid holding multiple locks simultaneously
- Use channels instead of shared memory (Go idiom)

**Must understand:**
- [ ] Difference between race condition and data race
- [ ] How to use `go run -race` and what it reports
- [ ] Coffman conditions — explain each and how to break one to prevent deadlock
- [ ] Goroutine leak — how to detect (runtime.NumGoroutine()) and prevent (use Context)

---

## Mutex / Semaphore / RWMutex

**Key Terms:**
- `mutex (mutual exclusion lock)` — allows only one goroutine to access critical section
- `Lock()` — acquire mutex; block if already held
- `Unlock()` — release mutex; defer recommended
- `deadlock risk` — Lock without Unlock; always use defer
- `RWMutex (Read-Write Mutex)` — multiple concurrent readers OR one exclusive writer
- `RLock()` — acquire read lock; multiple readers allowed simultaneously
- `RUnlock()` — release read lock
- `semaphore` — generalized mutex; allows N concurrent accesses (not just 1)
- `counting semaphore` — control concurrency (e.g., max 10 workers running)
- `binary semaphore` — same as mutex (value 0 or 1)
- `WaitGroup` — wait for a collection of goroutines to finish; Add, Done, Wait
- `Once` — execute function exactly once (lazy initialization)
- `Cond` — condition variable; signal/broadcast when condition changes
- `sync.Map` — concurrent-safe map; use when reads >> writes
- `atomic operations` — lock-free operations (atomic.AddInt64, atomic.LoadInt64)

**When to use what:**
- `sync.Mutex` — protect any shared data
- `sync.RWMutex` — read-heavy data (config, cache)
- `sync.WaitGroup` — wait for goroutines to finish
- `sync.Once` — singleton initialization
- `chan struct{}` — semaphore pattern in Go
- `atomic` — simple counters, flags (faster than mutex)

**Must understand:**
- [ ] Always use `defer mu.Unlock()` — why (panics still unlock)
- [ ] RWMutex — when it's better than regular mutex (many readers, few writers)
- [ ] WaitGroup — common pattern with goroutines
- [ ] Why sync.Map is slower than regular map with mutex for most use cases

---

## Memory Management

**Key Terms:**
- `stack` — fast; stores local variables and function call frames; automatically freed when function returns
- `heap` — slower; stores dynamically allocated data; must be freed (GC or manual)
- `stack allocation` — compiler puts variable on stack (value types, small structs)
- `heap allocation (escape analysis)` — compiler moves variable to heap if it escapes function scope
- `escape analysis` — Go compiler decides whether variable lives on stack or heap at compile time
- `garbage collector (GC)` — automatically frees unused heap memory
- `GC pause` — stop-the-world pause during GC; sources of latency spikes
- `Go GC` — concurrent, tri-color mark-and-sweep; very short pauses (< 1ms target)
- `JVM GC` — several algorithms: G1 (default), ZGC, Shenandoah; generational
- `generational GC` — young generation (new objects, collected often) + old generation (long-lived)
- `memory leak` — memory allocated but never freed; Go: goroutine leak, holding references
- `GOGC` — Go env variable controlling GC aggressiveness (default 100 = GC when heap doubles)
- `GOMEMLIMIT` — Go 1.19+: set soft memory limit; GC runs more aggressively near limit
- `finalizer` — function called when object is GC'd; Go: runtime.SetFinalizer
- `value type` — stored by value (copy); Go: int, bool, struct, array
- `reference type` — stored by reference (pointer); Go: slice, map, channel, pointer

**Must understand:**
- [ ] Stack vs heap — where does a Go variable live?
- [ ] Escape analysis — what causes a variable to escape to heap (returning pointer, closures, interface boxing)
- [ ] How Go GC works at a high level (concurrent, tri-color mark-and-sweep)
- [ ] What causes memory leaks in Go (goroutine leaks, holding refs in global slices/maps)

---

## Blocking vs Non-blocking I/O

**Key Terms:**
- `blocking I/O` — thread waits (sleeps) until I/O completes; simple but wastes CPU
- `non-blocking I/O` — returns immediately if not ready; caller must retry (polling) or use event notification
- `epoll (Linux)` — OS mechanism to monitor many file descriptors efficiently; O(1) notification
- `kqueue (BSD/macOS)` — equivalent to epoll on BSD systems
- `select/poll` — older I/O multiplexing; O(n) scan all fds; limited to ~1024 fds
- `event loop` — single thread uses epoll to know which connections have data; processes them
- `Go netpoller` — Go runtime uses epoll internally; goroutines block (appear blocked to programmer) but OS thread doesn't
- `goroutine park` — when goroutine blocks on I/O, Go runtime parks it and runs other goroutines on same OS thread
- `async/await` — structured concurrency with non-blocking I/O (Python, Kotlin, JavaScript)
- `reactor pattern` — event loop dispatches events to handlers (Netty, Node.js)
- `proactor pattern` — OS completes I/O and notifies app (Windows IOCP)
- `C10K problem` — challenge of handling 10,000 concurrent connections; solved by epoll + non-blocking I/O
- `C10M problem` — 10 million connections; requires kernel bypass (DPDK, io_uring)
- `io_uring` — Linux kernel 5.1+; asynchronous I/O without syscall overhead

**Must understand:**
- [ ] Why blocking I/O wastes threads (thread sleeps; OS still has to schedule it)
- [ ] How Go solves this — goroutines park, OS thread keeps running other goroutines
- [ ] Why Node.js can handle many connections with one thread (epoll + event loop)
- [ ] epoll vs select — why epoll scales (O(1) notification vs O(n) scan)

---

## Linux Basics

**Key Terms:**
- `process list` — `ps aux`, `ps -ef`, `top`, `htop`
- `kill` — send signal to process; `kill -9 PID` (SIGKILL, immediate), `kill PID` (SIGTERM, graceful)
- `SIGTERM` — polite shutdown request; application can handle and clean up
- `SIGKILL` — immediate kill; cannot be caught or ignored
- `SIGHUP` — hangup; often used to reload config without restart
- `file permissions` — rwxrwxrwx (owner/group/other); chmod 755
- `file descriptors` — integer handle for open files/sockets/pipes; `lsof -p PID`
- `ulimit` — per-process resource limits (max open files, max threads)
- `stdin/stdout/stderr` — fd 0, 1, 2; redirected with <, >, 2>
- `pipe |` — connect stdout of one command to stdin of another
- `grep` — search text; `grep -r pattern .` for recursive
- `netstat / ss` — show network connections; `ss -tlnp` (listening TCP ports with PIDs)
- `curl` — HTTP client; `curl -v` verbose, `-H` header, `-d` body, `-X` method
- `systemd` — service manager; `systemctl status/start/stop/restart service`
- `journalctl` — read systemd logs; `journalctl -u service -f` to follow
- `/proc filesystem` — virtual filesystem exposing kernel/process info; `/proc/PID/maps` memory map
- `strace` — trace system calls made by a process; powerful for debugging
- `tcpdump` — capture network packets; `tcpdump -i eth0 port 80`

**Must understand:**
- [ ] SIGTERM vs SIGKILL — why you should always try SIGTERM first
- [ ] How to check which process is using a port (`ss -tlnp | grep :8080`)
- [ ] How to read basic /proc info (fd count, memory usage)
- [ ] Basic file permission model (chmod, chown)

---

## Signals & Graceful Shutdown

**Key Terms:**
- `signal` — asynchronous notification sent to process by OS or another process
- `SIGTERM (15)` — default kill signal; application should catch and shut down gracefully
- `SIGKILL (9)` — cannot be caught; instant termination
- `SIGINT (2)` — Ctrl+C; usually causes graceful shutdown
- `SIGHUP (1)` — terminal hangup; daemon apps use it to reload config
- `SIGQUIT (3)` — Ctrl+\\; quit with core dump
- `graceful shutdown` — stop accepting new requests, finish in-flight requests, close connections, flush logs
- `shutdown timeout` — max time to wait for graceful shutdown before forcing kill
- `signal.Notify (Go)` — register channel to receive signals; `signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)`
- `os/signal package (Go)` — handle OS signals in Go
- `ShutdownHook (Java)` — `Runtime.getRuntime().addShutdownHook(thread)`
- `Kubernetes terminationGracePeriodSeconds` — how long K8s waits between SIGTERM and SIGKILL (default 30s)
- `preStop hook` — K8s pod lifecycle hook; runs before SIGTERM; useful for draining traffic

**Graceful shutdown pattern in Go:**
```go
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
<-quit  // block until signal received
// stop accepting new requests
// wait for in-flight requests (WaitGroup)
// close DB connections, flush buffers
```

**Must understand:**
- [ ] Why graceful shutdown matters (avoid dropping in-flight requests)
- [ ] SIGTERM → graceful → SIGKILL if timeout exceeded (K8s default 30s)
- [ ] How to implement graceful shutdown in Go (signal.Notify + WaitGroup + context)
- [ ] preStop hook in Kubernetes — why you sometimes need it (load balancer drain takes time)
