# Operating Systems — Practice

Practice for `roadmap/os.md` is **small experiments on Linux**: write a tiny Go program that does one OS-level thing (spawn a child, open many files, allocate memory, spin goroutines, listen on a socket), then **observe it** with Linux tools. No C needed. Pure theory items (e.g. dining philosophers) are done as a diagram + explanation in the notes.

## Setup
```bash
sudo apt install -y strace ltrace sysstat procps psmisc lsof linux-tools-common linux-tools-$(uname -r) util-linux
# Go programs live in this module (module root is code/, module name `code`)
cd code && go run ./os/<NN>-<topic-slug>/
# Containers for namespace / cgroup / OOM experiments
docker run --rm -it --memory=256m --cpus=0.5 -v "$PWD":/src -w /src golang:latest bash
```

## Typical tools
- `strace -f -c ./prog` — which syscalls, how many, how long
- `/proc/<pid>/{status,maps,fd,task,limits}` — what the kernel knows about the process
- `pidstat -w 1`, `vmstat 1`, `perf stat` — context switches, faults, cache misses
- `dmesg -T`, `/sys/fs/cgroup/...` — OOM kills, throttling

## Folder convention
- `code/os/<NN>-<topic-slug>/` — `NN` is the OS topic number from the roadmap (e.g. `03-processes/`, `07-file-descriptors/`).
- `README.md` in that folder holds the exercise instructions, created when the topic starts from the 🛠 **Practice** line in the roadmap.
- Next to it: your `main.go` experiment(s) and a short `results.md` with the commands you ran, the output that mattered, and your explanation.
- Notes go to `notes/os/`; term definitions are in `reference/os.md`.
