# Operating Systems & Computer Architecture Roadmap

> **L1** = aware & can use · **L2** = deep & efficient · Status: ⬜ not started · 🔄 in progress · ✅ done
> **Notes:** `notes/os/` · **Practice:** `code/os/` · **Dictionary:** `reference/os.md`

**Why (EN):** The Go runtime, Postgres, Docker and Kubernetes are all built on OS primitives (processes, threads, virtual memory, file descriptors, epoll, namespaces, cgroups); without this layer they stay "magic" and production incidents (OOM kills, CPU throttling, fd leaks, latency spikes) are impossible to explain. In the AI era, AI can write the code, but a senior must reason about why a service behaves as it does on a real machine — this foundation matters more, while memorizing textbook algorithms (e.g. exact page-replacement variants) matters less.
**Зачем (RU):** Рантайм Go, Postgres, Docker и Kubernetes построены на примитивах ОС (процессы, потоки, виртуальная память, файловые дескрипторы, epoll, namespaces, cgroups); без этого слоя они остаются «магией», а инциденты в продакшене (OOM kill, троттлинг CPU, утечки fd, всплески задержек) невозможно объяснить. В эпоху ИИ код может писать ИИ, но senior должен понимать, почему сервис ведёт себя так на реальной машине — эта база становится важнее, а заучивание учебных алгоритмов (например, точных вариантов замещения страниц) — менее важным.

## Progress
| Layer | Done | Total |
|---|---|---|
| L1 | 0 | 31 |
| L2 | 0 | 22 |

---

## OS 1 — How a computer runs a program
**L1 — aware & can use**
- ⬜ **1.1 CPU basics** — CPU, core, register, program counter, instruction, clock cycle, instruction set (x86-64, ARM64), machine code, compiler output
- ⬜ **1.2 Memory hierarchy** — registers, L1/L2/L3 cache, RAM, SSD, HDD, network, latency numbers every programmer should know, bandwidth vs latency
- ⬜ **1.3 From source to running process** — compile, link, static vs dynamic linking, ELF binary, loader, entry point, Go static binaries, `CGO_ENABLED=0`

**L2 — deep & efficient**
- ⬜ **1.4 CPU execution details** — pipelining, superscalar, out-of-order execution, branch prediction, branch misprediction cost, SIMD, speculative execution, Spectre/Meltdown (awareness)
- ⬜ **1.5 Privilege & interrupts** — CPU rings (kernel/user mode), interrupt, hardware vs software interrupt, exception, trap, timer interrupt, interrupt handler, `/proc/interrupts`

🛠 **Practice** — L1: build a Go binary and inspect it with `file`, `ldd`, `readelf -h`, `go tool nm`; compare sizes with and without `-ldflags="-s -w"` · L2: write a Go benchmark summing a slice with a data-dependent `if` on sorted vs unsorted data and explain the difference with branch prediction (`perf stat -e branch-misses`)

---

## OS 2 — OS structure & the Linux kernel
**L1 — aware & can use**
- ⬜ **2.1 What an OS does** — kernel, user space, kernel space, abstraction (process, file, socket), resource management, isolation, protection
- ⬜ **2.2 Boot process** — firmware (UEFI/BIOS), bootloader (GRUB), kernel image, initramfs, PID 1, init system (`systemd`), runlevels/targets
- ⬜ **2.3 Linux kernel overview** — monolithic kernel, kernel modules, `lsmod`, `uname -r`, `/proc`, `/sys`, `dmesg`, kernel version, distributions

**L2 — deep & efficient**
- ⬜ **2.4 Kernel designs & extension** — monolithic vs microkernel vs hybrid, unikernel (awareness), loadable modules, eBPF, kernel tunables `sysctl`, `/proc/sys`, vDSO

🛠 **Practice** — L1: run `uname -a`, `dmesg | head`, `systemd-analyze blame`, and explore `/proc/self/` from a Go program that prints its own `/proc/self/status` · L2: list 10 `sysctl` values relevant to servers (`net.core.somaxconn`, `vm.swappiness`, `fs.file-max`, ...) and write what each controls

---

## OS 3 — Processes
**L1 — aware & can use**
- ⬜ **3.1 Process basics** — process, PID, PPID, parent/child, process states (running, sleeping, stopped, zombie), exit code, `ps`, `/proc/<pid>`
- ⬜ **3.2 Creating processes** — `fork`, `exec`, `wait`, copy-on-write, `os/exec` in Go, environment, arguments, inherited file descriptors
- ⬜ **3.3 Process lifecycle problems** — zombie, orphan, reaping, `SIGCHLD`, init adopts orphans, PID 1 in containers, `tini`

**L2 — deep & efficient**
- ⬜ **3.4 Process internals** — process control block (`task_struct`), `clone` flags, `vfork`, `posix_spawn`, process groups, sessions, controlling terminal, daemons
- ⬜ **3.5 Inter-process communication** — pipes, named pipes (FIFO), Unix domain sockets, shared memory, message queues, signals as IPC, `mmap` shared mapping

🛠 **Practice** — L1: write a Go program that starts a child with `os/exec`, does not wait for it, and observe the zombie in `ps -o pid,ppid,stat,cmd` · L2: write two Go programs that talk over a Unix domain socket and over a named pipe; `strace -f` both and compare the syscalls

---

## OS 4 — Threads & scheduling
**L1 — aware & can use**
- ⬜ **4.1 Threads** — kernel thread, user thread, thread vs process, shared address space, thread stack, TID, `/proc/<pid>/task`
- ⬜ **4.2 Threading models** — 1:1, N:1, M:N, green threads, goroutines on OS threads, Go `GOMAXPROCS`, `runtime.LockOSThread`
- ⬜ **4.3 Scheduling basics** — preemptive scheduling, time slice, priority, `nice`, run queue, CPU-bound vs I/O-bound, load average

**L2 — deep & efficient**
- ⬜ **4.4 Linux scheduler** — CFS, EEVDF (newer kernels), vruntime, scheduling classes (`SCHED_OTHER`, `SCHED_FIFO`, `SCHED_RR`), CPU affinity, `taskset`, `chrt`, real-time scheduling
- ⬜ **4.5 Context switches** — context switch cost, voluntary vs involuntary switches, cache/TLB pollution, `/proc/<pid>/status` (`voluntary_ctxt_switches`), `pidstat -w`, `perf sched`

🛠 **Practice** — L1: run a Go program with 1000 goroutines and count its OS threads via `ls /proc/<pid>/task | wc -l`; change `GOMAXPROCS` and repeat · L2: measure voluntary vs involuntary context switches of a CPU-bound vs I/O-bound Go program with `pidstat -w 1`

---

## OS 5 — Memory management
**L1 — aware & can use**
- ⬜ **5.1 Virtual memory** — address space, virtual vs physical address, page, page table, TLB, page fault (minor/major), swap, MMU
- ⬜ **5.2 Process memory layout** — text, data, BSS, heap, stack, memory-mapped region, stack overflow, `/proc/<pid>/maps`
- ⬜ **5.3 Memory metrics & limits** — RSS, VSZ, PSS, shared memory, OOM killer, `oom_score`, overcommit, cgroup memory limit, `free`, `/proc/meminfo`

**L2 — deep & efficient**
- ⬜ **5.4 Allocation & mapping** — `brk`, `mmap`, anonymous vs file-backed mapping, `madvise` (`MADV_DONTNEED`, `MADV_FREE`), copy-on-write, demand paging, Go runtime returning memory to the OS, `GOMEMLIMIT`
- ⬜ **5.5 Advanced memory** — huge pages, transparent huge pages (THP), TLB shootdown, page replacement (LRU approximations), dirty pages, `vm.swappiness`, NUMA memory policy (awareness)

🛠 **Practice** — L1: write a Go program that allocates 500 MB slowly and watch `VmRSS` in `/proc/<pid>/status` and `top`; then run it in Docker with `--memory=256m` and observe the OOM kill in `dmesg` · L2: compare RSS of the same program with and without `GOMEMLIMIT`, and count page faults with `/usr/bin/time -v`

---

## OS 6 — System calls & signals
**L1 — aware & can use**
- ⬜ **6.1 System calls** — user space vs kernel space, syscall, syscall number, `read`/`write`/`open`/`close`, libc vs Go's raw syscalls, `strace`, cost of syscalls
- ⬜ **6.2 Signals** — SIGTERM, SIGKILL, SIGINT, SIGHUP, SIGCHLD, SIGSEGV, signal handler, default action, `kill`, `os/signal`, `signal.NotifyContext`
- ⬜ **6.3 Graceful shutdown** — SIGTERM handling, drain connections, `http.Server.Shutdown`, timeouts, Kubernetes termination grace period, PID 1 problem in containers

**L2 — deep & efficient**
- ⬜ **6.4 Syscall internals** — `syscall` instruction, mode switch, vDSO (`clock_gettime`), blocking syscalls and Go's `entersyscall`/handoff, seccomp, `strace -c` profiling, syscall overhead after Meltdown mitigations
- ⬜ **6.5 Signal internals** — signal mask, pending signals, real-time signals, `SA_RESTART`, `EINTR`, async-safe functions, Go async preemption via `SIGURG`

🛠 **Practice** — L1: `strace -f -c` a small Go HTTP server during 100 requests and list the top 5 syscalls; add `signal.NotifyContext` graceful shutdown and test it with `kill -TERM` · L2: run `strace -f -e trace=signal` on a busy Go program and find the `SIGURG` preemption signals

---

## OS 7 — Files, descriptors & file systems
**L1 — aware & can use**
- ⬜ **7.1 File descriptors** — file descriptor, fd table, stdin/stdout/stderr (0/1/2), "everything is a file", open file limit (`ulimit -n`), fd leak, `/proc/<pid>/fd`, `lsof`
- ⬜ **7.2 File system basics** — inode, directory entry (dentry), hard link, symlink, mount, VFS, ext4, xfs, tmpfs, overlayfs
- ⬜ **7.3 Durability basics** — page cache, dirty pages, `write` vs `fsync`, `fdatasync`, `O_SYNC`, rename for atomic write, journaling

**L2 — deep & efficient**
- ⬜ **7.4 File system internals** — open file description vs fd, `dup2`, `O_CLOEXEC`, write-ahead log, journaling modes, block size, extents, `O_DIRECT`, readahead
- ⬜ **7.5 Storage performance** — IOPS, throughput, latency, sequential vs random I/O, SSD vs HDD, write amplification, I/O scheduler, `iostat -x`, `fio` (awareness)

🛠 **Practice** — L1: write a Go program that opens files without closing them until it hits `too many open files`; watch `/proc/<pid>/fd` and raise the limit with `ulimit -n` · L2: benchmark writing 10k small records with `fsync` after each vs once at the end, and explain why databases batch commits

---

## OS 8 — I/O models
**L1 — aware & can use**
- ⬜ **8.1 Blocking vs non-blocking** — blocking I/O, non-blocking I/O, `O_NONBLOCK`, `EAGAIN`, synchronous vs asynchronous, thread-per-connection
- ⬜ **8.2 I/O multiplexing** — `select`, `poll`, `epoll`, `kqueue`, readiness, event loop, C10K problem, Go netpoller (goroutine looks blocking, runtime uses epoll)

**L2 — deep & efficient**
- ⬜ **8.3 epoll in depth** — `epoll_create1`, `epoll_ctl`, `epoll_wait`, level-triggered vs edge-triggered, `EPOLLONESHOT`, thundering herd, `SO_REUSEPORT`
- ⬜ **8.4 Modern & zero-copy I/O** — io_uring (submission/completion queues), `sendfile`, `splice`, zero-copy, `io.Copy` using `sendfile`/`splice` in Go, mmap-based I/O, AIO (awareness)

🛠 **Practice** — L1: `strace -f -e trace=epoll_wait,epoll_ctl,accept4` a Go TCP echo server while connecting with `nc` from 3 terminals · L2: serve a large file with `http.ServeFile` and confirm `sendfile` in `strace` output; compare with copying through a user-space buffer

---

## OS 9 — Concurrency theory
**L1 — aware & can use**
- ⬜ **9.1 Synchronization primitives** — critical section, mutex, read-write lock, semaphore, spinlock, condition variable, monitor, atomic operation, CAS
- ⬜ **9.2 Concurrency problems** — race condition, data race, deadlock (4 Coffman conditions), livelock, starvation, priority inversion
- ⬜ **9.3 Classic problems** — producer–consumer, bounded buffer, readers–writers, dining philosophers, sleeping barber (awareness)

**L2 — deep & efficient**
- ⬜ **9.4 Lock implementation** — futex, spinning vs parking, lock contention, lock convoy, fairness, Go `sync.Mutex` normal vs starvation mode, lock-free vs wait-free
- ⬜ **9.5 Lock-free techniques** — CAS loop, ABA problem, lock-free queue/stack, `sync/atomic` types, memory reclamation (awareness), sharding to reduce contention

🛠 **Practice** — L1: write a Go program with a data race on a counter, catch it with `go run -race`, then fix it three ways (mutex, atomic, channel) · L2: create a deadlock with two mutexes acquired in opposite order, then fix it with lock ordering; benchmark mutex vs atomic counter under 8 goroutines

---

## OS 10 — CPU caches & memory ordering
**L1 — aware & can use**
- ⬜ **10.1 CPU caches** — L1/L2/L3 cache, cache line (64 B), cache hit/miss, locality (spatial/temporal), hyper-threading (SMT), prefetching
- ⬜ **10.2 Writing cache-friendly code** — row-major traversal, arrays vs pointer-heavy structures, struct field order and padding, slice of structs vs struct of slices, false sharing

**L2 — deep & efficient**
- ⬜ **10.3 Cache coherence** — MESI protocol, cache line bouncing, false sharing fix with padding, `cpu.CacheLinePad` (awareness), `perf stat -e cache-misses`
- ⬜ **10.4 Memory ordering** — instruction reordering (compiler and CPU), memory barrier / fence, happens-before, sequential consistency, x86 TSO vs ARM weak ordering, Go memory model
- ⬜ **10.5 NUMA** — NUMA node, local vs remote memory, `numactl`, `lscpu`, CPU pinning, impact on databases (awareness)

🛠 **Practice** — L1: benchmark summing a 2D array row-by-row vs column-by-column in Go and explain the result · L2: benchmark two goroutines incrementing adjacent `int64` fields vs fields padded to 64 bytes and show false sharing with `perf stat -e cache-misses`

---

## OS 11 — Linux internals for containers
**L1 — aware & can use**
- ⬜ **11.1 Namespaces** — pid, net, mnt, uts, ipc, user, cgroup namespaces, isolation, `unshare`, `nsenter`, `lsns`, `/proc/<pid>/ns`
- ⬜ **11.2 cgroups** — cgroups v2, unified hierarchy, `cpu.max`, `memory.max`, `pids.max`, CPU quota, CPU throttling, memory limit, `/sys/fs/cgroup`
- ⬜ **11.3 Container = process** — container is a process with namespaces + cgroups + rootfs, overlayfs layers, `chroot` vs `pivot_root`, container runtime (`runc`, `containerd`), foundation of Docker (see OPS track)

**L2 — deep & efficient**
- ⬜ **11.4 Resource limits in practice** — CFS bandwidth control, `nr_throttled`, `GOMAXPROCS` vs CPU quota (Go container-aware `GOMAXPROCS` in newer versions, `automaxprocs`), `GOMEMLIMIT` vs `memory.max`, PSI (pressure stall information), `memory.events`
- ⬜ **11.5 Container security primitives** — capabilities (`CAP_NET_ADMIN`, `CAP_SYS_ADMIN`), seccomp profiles, rootless containers, user namespace mapping, AppArmor/SELinux (awareness), read-only rootfs

🛠 **Practice** — L1: create a namespace with `sudo unshare --pid --fork --mount-proc bash`, run `ps aux`, and compare with the host; find your Docker container's cgroup files under `/sys/fs/cgroup` · L2: run a CPU-heavy Go program in Docker with `--cpus=0.5`, read `cpu.stat` (`nr_throttled`, `throttled_usec`) and show the effect of `GOMAXPROCS` on latency
