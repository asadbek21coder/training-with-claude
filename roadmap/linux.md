# Linux for Backend Engineers Roadmap

> **L1** = aware & can use · **L2** = deep & efficient · Status: ⬜ not started · 🔄 in progress · ✅ done
> **Notes:** `notes/linux/` · **Practice:** `code/linux/`

**Why (EN):** Almost every backend service runs on Linux, and when production breaks you are on a server (or inside a container) with a shell, not an IDE — you must read logs, find the process eating CPU, check disk, ports and DNS in minutes. In the AI era, AI can suggest commands, but you must know what they do before running them on a production box; fast, safe troubleshooting and reading tool output matter more, memorizing rare flags matters less.
**Зачем (RU):** Почти любой бэкенд-сервис работает на Linux, и когда продакшен падает, у тебя есть сервер (или контейнер) с shell, а не IDE — нужно за минуты прочитать логи, найти процесс, который съедает CPU, проверить диск, порты и DNS. В эпоху ИИ команды может подсказать ИИ, но ты должен понимать, что они делают, прежде чем запускать их на продакшене; быстрая и безопасная диагностика и умение читать вывод инструментов — важнее, заучивание редких флагов — менее важно.

## Progress
| Layer | Done | Total |
|---|---|---|
| L1 | 0 | 41 |
| L2 | 0 | 29 |

---

## LX 1 — Shell basics & environment
**L1 — aware & can use**
- ⬜ **1.1 Shell & terminal** — terminal emulator, shell, `bash`, `zsh`, prompt, command, arguments, options/flags, `man`, `--help`, `tldr`, `type`, `which`
- ⬜ **1.2 Navigation** — `pwd`, `cd`, `ls -la`, absolute vs relative path, `~`, `.`, `..`, `cd -`, tab completion, globbing (`*`, `?`, `[]`), `tree`
- ⬜ **1.3 Command line efficiency** — history, `Ctrl+R`, `!!`, `!$`, `Ctrl+A`/`Ctrl+E`/`Ctrl+W`/`Ctrl+U`, `Ctrl+C`/`Ctrl+D`/`Ctrl+Z`, `alias`, `clear`
- ⬜ **1.4 Environment variables** — `env`, `export`, `echo $VAR`, `PATH`, `HOME`, `USER`, `SHELL`, `unset`, `VAR=x cmd` (one-off), `.env` files, `printenv`

**L2 — deep & efficient**
- ⬜ **1.5 Shell startup & configuration** — login vs non-login shell, interactive vs non-interactive, `~/.bashrc`, `~/.profile`, `~/.bash_profile`, `/etc/profile`, `source`, dotfiles repo, `PS1`
- ⬜ **1.6 How the shell runs a command** — builtins vs external commands, `PATH` lookup, `hash`, fork + exec, expansion order (brace, tilde, parameter, command, arithmetic, word splitting, glob), subshell, `exec`

🛠 **Practice** — L1: add `~/go/bin` to `PATH` permanently, create 3 useful aliases (e.g. `gs` for `git status`), and find 5 past commands with `Ctrl+R` · L2: put `echo` lines in `~/.profile` and `~/.bashrc`, then open a login shell, a non-login shell, and `ssh localhost` and record which files ran

---

## LX 2 — Files & the filesystem
**L1 — aware & can use**
- ⬜ **2.1 Working with files** — `touch`, `mkdir -p`, `cp -r`, `mv`, `rm -rf` (danger), `cat`, `less`, `head`, `tail -f`, `file`, `stat`
- ⬜ **2.2 Filesystem layout (FHS)** — `/`, `/etc`, `/var`, `/var/log`, `/home`, `/tmp`, `/usr`, `/usr/local`, `/opt`, `/bin`, `/proc`, `/sys`, `/dev`
- ⬜ **2.3 Finding things** — `find` (`-name`, `-type`, `-mtime`, `-size`, `-exec`), `locate`, `fd` (awareness), `which`, `whereis`
- ⬜ **2.4 Links & archives** — hard link, symbolic link, `ln -s`, `readlink`, `tar -czf`/`-xzf`, `gzip`, `zip`/`unzip`, `zstd`

**L2 — deep & efficient**
- ⬜ **2.5 Inodes & special files** — inode, `ls -i`, link count, deleted-but-open file, device files, `/dev/null`, `/dev/zero`, `/dev/urandom`, named pipes (`mkfifo`), sockets as files
- ⬜ **2.6 Safe file operations** — `rsync -av --dry-run`, `rm -i`, `trash-cli` (awareness), atomic `mv` on same filesystem, `install`, `mktemp`, `noclobber`, backups before `sed -i`

🛠 **Practice** — L1: find all `.go` files modified in the last 7 days in your repo, and all files over 100 MB in `/var`; archive a folder with `tar` and restore it · L2: delete a log file while `tail -f` still holds it, show that disk space is not freed with `lsof +L1`, then free it

---

## LX 3 — Permissions, users & groups
**L1 — aware & can use**
- ⬜ **3.1 Permissions** — owner/group/other, `rwx`, octal mode (`755`, `644`, `600`), `chmod`, `chown`, `chgrp`, `ls -l` output, execute bit on directories
- ⬜ **3.2 Users & groups** — `/etc/passwd`, `/etc/group`, `/etc/shadow`, UID, GID, `id`, `whoami`, `groups`, `useradd`, `usermod -aG`, `passwd`, system users
- ⬜ **3.3 sudo & root** — root user, `sudo`, `sudo -i`, `su`, `/etc/sudoers`, `visudo`, principle of least privilege, `sudo -l`

**L2 — deep & efficient**
- ⬜ **3.4 Special permissions** — setuid, setgid, sticky bit (`/tmp`), `umask`, default permissions, ACLs (`getfacl`, `setfacl`), `chattr +i`, `lsattr`
- ⬜ **3.5 Service accounts & security** — dedicated user per service, `nologin` shell, file ownership for app dirs, capabilities (`setcap cap_net_bind_service`), running containers as non-root, `/etc/sudoers.d/`

🛠 **Practice** — L1: create a user `appuser` and group `app`, make a directory only that group can write, and verify with `sudo -u appuser touch` · L2: let a Go binary bind port 80 without root using `setcap cap_net_bind_service=+ep`, and explain why that beats running as root

---

## LX 4 — Text processing & pipelines
**L1 — aware & can use**
- ⬜ **4.1 Pipes & redirection** — `|`, `>`, `>>`, `<`, `2>`, `2>&1`, `&>`, `/dev/null`, `tee`, here-doc `<<EOF`, here-string `<<<`, exit status of pipelines
- ⬜ **4.2 grep** — `grep -i`, `-r`, `-n`, `-v`, `-c`, `-l`, `-E` (regex), `-o`, `-A`/`-B`/`-C`, `ripgrep` (`rg`)
- ⬜ **4.3 Core filters** — `sort`, `sort -n`, `-k`, `uniq -c`, `wc -l`, `cut -d -f`, `tr`, `head`, `tail`, `column -t`, `diff`
- ⬜ **4.4 sed, awk, jq, xargs** — `sed 's/a/b/g'`, `sed -i`, `awk '{print $1}'`, `awk -F`, `jq '.field'`, `jq -r`, `jq 'select()'`, `xargs`, `xargs -P`, `xargs -0` with `find -print0`

**L2 — deep & efficient**
- ⬜ **4.5 Regular expressions** — BRE vs ERE vs PCRE, anchors, character classes, quantifiers, groups, backreferences, greedy vs lazy, `grep -P`
- ⬜ **4.6 Advanced awk & jq** — awk patterns, `BEGIN`/`END`, associative arrays, sums and averages, `printf`; jq `map`, `group_by`, `to_entries`, `--arg`, building JSON, processing JSON logs

🛠 **Practice** — L1: from an nginx-style access log, print the top 10 IPs by request count and the count of 5xx responses with one pipeline each · L2: from a JSON log file produced by Go `slog`, compute p50/p95 latency per endpoint using `jq` + `sort` + `awk`

---

## LX 5 — Processes, jobs & signals
**L1 — aware & can use**
- ⬜ **5.1 Viewing processes** — `ps aux`, `ps -ef`, `ps -o pid,ppid,stat,cmd`, `pgrep`, `pstree`, `top`, `htop`, PID, process state (R, S, D, Z, T)
- ⬜ **5.2 Signals & killing** — `kill`, `kill -9`, `kill -TERM`, `pkill`, `killall`, SIGTERM vs SIGKILL, SIGINT (`Ctrl+C`), SIGHUP (reload config), `kill -l`
- ⬜ **5.3 Jobs & background** — `&`, `jobs`, `fg`, `bg`, `Ctrl+Z`, `nohup`, `disown`, `wait`, foreground vs background

**L2 — deep & efficient**
- ⬜ **5.4 Process inspection** — `/proc/<pid>/status`, `/proc/<pid>/cmdline`, `/proc/<pid>/environ`, `/proc/<pid>/fd`, `/proc/<pid>/limits`, uninterruptible sleep (D state), zombie cleanup, `nice`, `renice`, `ionice`
- ⬜ **5.5 Resource limits** — `ulimit -a`, soft vs hard limits, `nofile`, `nproc`, `/etc/security/limits.conf`, systemd `LimitNOFILE=`, `prlimit`

🛠 **Practice** — L1: start a Go HTTP server in the background, find it with `pgrep`, stop it with SIGTERM (graceful) and then try SIGKILL, and compare the logs · L2: read the environment and open fds of a running process from `/proc/<pid>/`, and raise its open-file limit with `prlimit`

---

## LX 6 — Bash scripting
**L1 — aware & can use**
- ⬜ **6.1 Script basics** — shebang `#!/usr/bin/env bash`, `chmod +x`, variables, quoting (`"$var"` vs `'$var'`), `$1`/`$@`/`$#`, `$?`, command substitution `$(...)`, `read`
- ⬜ **6.2 Control flow** — `if`, `[[ ]]`, `-f`/`-d`/`-z`/`-n`, `case`, `for`, `while`, `&&`/`||`, exit codes, `exit 1`
- ⬜ **6.3 Functions & safety** — functions, `local`, return values, `set -euo pipefail`, `trap` on `EXIT`/`ERR`, `shellcheck`

**L2 — deep & efficient**
- ⬜ **6.4 Robust scripts** — arrays, associative arrays, parameter expansion (`${var:-default}`, `${var%.*}`, `${var##*/}`), `getopts`, idempotency, logging to stderr, `mktemp` + `trap` cleanup, `flock` for single instance
- ⬜ **6.5 When not to use bash** — complexity threshold, portability (`sh` vs `bash`), word splitting pitfalls, `IFS`, rewriting in Go, `Makefile` / `just` as task runners

🛠 **Practice** — L1: write `backup.sh` that tars a directory into `~/backups/` with a date in the name and keeps only the last 5 archives; run `shellcheck` on it · L2: write `deploy.sh` with `getopts` flags, `set -euo pipefail`, a `trap` cleanup, and `flock` so two runs cannot overlap

---

## LX 7 — Packages & software installation
**L1 — aware & can use**
- ⬜ **7.1 apt basics** — `apt update`, `apt upgrade`, `apt install`, `apt remove`, `apt purge`, `apt search`, `apt show`, `apt autoremove`, `dpkg -l`
- ⬜ **7.2 Other install methods** — `.deb` with `dpkg -i`, snap, flatpak (awareness), static binaries in `/usr/local/bin`, `go install`, tarball installs (Go itself)

**L2 — deep & efficient**
- ⬜ **7.3 Repositories & pinning** — `/etc/apt/sources.list.d/`, GPG keys (`signed-by`), third-party repos (Docker, PostgreSQL), version pinning, `apt-mark hold`, unattended-upgrades, `apt-cache policy`
- ⬜ **7.4 Packages in containers** — `apt-get` vs `apt` in Dockerfiles, `--no-install-recommends`, cleaning `/var/lib/apt/lists`, Alpine `apk` and musl, distroless images, reproducible versions

🛠 **Practice** — L1: install `jq`, `htop`, `ripgrep` and `tree` with apt; find which package owns `/usr/bin/curl` with `dpkg -S` · L2: add the official PostgreSQL apt repository with a `signed-by` key, install a specific version, and `apt-mark hold` it

---

## LX 8 — systemd, journald & logs
**L1 — aware & can use**
- ⬜ **8.1 systemctl** — unit, service, `systemctl status`, `start`, `stop`, `restart`, `reload`, `enable`, `disable`, `is-active`, `list-units --failed`, `daemon-reload`
- ⬜ **8.2 journalctl** — `journalctl -u <unit>`, `-f`, `-n 100`, `--since "1 hour ago"`, `-p err`, `-b`, `-o json`, `--disk-usage`
- ⬜ **8.3 Classic logs** — `/var/log/syslog`, `/var/log/auth.log`, `dmesg`, `logrotate`, log rotation, stdout logging for services

**L2 — deep & efficient**
- ⬜ **8.4 Writing unit files** — `[Unit]`/`[Service]`/`[Install]`, `ExecStart`, `Restart=on-failure`, `RestartSec`, `User=`, `EnvironmentFile=`, `WorkingDirectory=`, `Type=simple` vs `notify`, `After=`/`Wants=`, drop-in overrides (`systemctl edit`)
- ⬜ **8.5 Hardening & resource control** — `ProtectSystem=`, `PrivateTmp=`, `NoNewPrivileges=`, `DynamicUser=`, `MemoryMax=`, `CPUQuota=`, `systemd-analyze security`, journald retention (`SystemMaxUse=`)

🛠 **Practice** — L1: find why a service failed using `systemctl status` and `journalctl -u ... -p err --since today` · L2: write a systemd unit for your Go HTTP server with `Restart=on-failure`, a dedicated user, `EnvironmentFile`, `MemoryMax`, then kill it with SIGKILL and watch systemd restart it

---

## LX 9 — Scheduled jobs: cron & timers
**L1 — aware & can use**
- ⬜ **9.1 cron** — `crontab -e`, `crontab -l`, cron syntax (5 fields), `*/5`, `@daily`, `@reboot`, `/etc/cron.d/`, environment in cron (minimal `PATH`), output redirection to a log

**L2 — deep & efficient**
- ⬜ **9.2 systemd timers** — `.timer` + `.service`, `OnCalendar=`, `OnBootSec=`, `Persistent=true`, `RandomizedDelaySec=`, `systemctl list-timers`, logs in journald, why timers beat cron
- ⬜ **9.3 Reliable scheduled jobs** — overlapping runs (`flock`), idempotency, timeouts, failure alerts, time zones and DST, distributed cron problem (one runner across many servers)

🛠 **Practice** — L1: add a cron job that runs your `backup.sh` every day at 03:00 and logs output to a file; debug it when `PATH` is missing · L2: convert the same job to a systemd timer with `Persistent=true` and check it with `systemctl list-timers`

---

## LX 10 — Disks & storage
**L1 — aware & can use**
- ⬜ **10.1 Disk usage** — `df -h`, `df -i` (inodes), `du -sh`, `du -h --max-depth=1`, `ncdu`, "disk full" causes (logs, Docker images, deleted-open files)
- ⬜ **10.2 Devices & mounts** — `lsblk`, `blkid`, `mount`, `umount`, `/etc/fstab`, `findmnt`, partition, filesystem (ext4, xfs), tmpfs

**L2 — deep & efficient**
- ⬜ **10.3 Disk management** — `fdisk`/`parted`, `mkfs`, LVM (PV, VG, LV), extending a volume, `resize2fs`/`xfs_growfs`, swap file, mount options (`noatime`)
- ⬜ **10.4 Disk performance & health** — `iostat -xz 1`, `%util`, `await`, `iotop`, `fio` (awareness), `smartctl`, `docker system df` / `docker system prune`

🛠 **Practice** — L1: find the 10 biggest directories under `/var` and clean Docker space with `docker system df` / `prune` · L2: create a 1 GB file-backed loop device, format it as ext4, mount it via `/etc/fstab`, then grow it and resize the filesystem

---

## LX 11 — Networking commands
**L1 — aware & can use**
- ⬜ **11.1 Interfaces & addresses** — `ip a`, `ip r`, `ip link`, `hostname -I`, `/etc/hosts`, `/etc/resolv.conf`, `ping`, `traceroute`/`mtr`
- ⬜ **11.2 Ports & connections** — `ss -tulpn`, `ss -tan state established`, `lsof -i :8080`, `nc -zv host port`, "address already in use", listening on `0.0.0.0` vs `127.0.0.1`
- ⬜ **11.3 HTTP & DNS clients** — `curl -v`, `curl -I`, `curl -X POST -d -H`, `curl -w` timing, `wget`, `dig`, `dig +short`, `nslookup`, `host`

**L2 — deep & efficient**
- ⬜ **11.4 Packet capture & debugging** — `tcpdump -i any port 8080`, `-nn`, `-w file.pcap`, Wireshark for pcap, `ss -i` (TCP info), `nstat`, `ethtool` (awareness), `openssl s_client`
- ⬜ **11.5 Firewall & host networking** — `iptables`/`nftables` (awareness), `ufw`, network namespaces (`ip netns`), `sysctl net.*` basics, `/proc/net/`

🛠 **Practice** — L1: run a Go HTTP server, find its port with `ss -tulpn`, hit it with `curl -v` and `curl -w "%{time_total}\n"`, and resolve a domain with `dig +short` · L2: capture the TCP handshake and an HTTP request to your server with `tcpdump -nn -i lo port 8080` and open the pcap in Wireshark

---

## LX 12 — SSH & remote work
**L1 — aware & can use**
- ⬜ **12.1 SSH basics** — `ssh user@host`, key pair, `ssh-keygen -t ed25519`, `ssh-copy-id`, `~/.ssh/authorized_keys`, `known_hosts`, host key fingerprint, file permissions (`600`)
- ⬜ **12.2 Copying & config** — `scp`, `rsync -avz`, `~/.ssh/config` (`Host`, `HostName`, `User`, `IdentityFile`), `ssh-agent`, `ssh-add`

**L2 — deep & efficient**
- ⬜ **12.3 Tunnels & jump hosts** — local port forwarding (`-L`), remote forwarding (`-R`), dynamic SOCKS (`-D`), `ProxyJump` / bastion, `ControlMaster` multiplexing, agent forwarding risks
- ⬜ **12.4 Server hardening** — `/etc/ssh/sshd_config`, `PasswordAuthentication no`, `PermitRootLogin no`, `fail2ban`, changing port (limited value), SSH certificates (awareness)

🛠 **Practice** — L1: set up key-based SSH to a local Docker container or VM and add a `Host` alias in `~/.ssh/config` · L2: reach a Postgres listening only on the remote `127.0.0.1:5432` with `ssh -L 5433:localhost:5432` and connect with `psql`

---

## LX 13 — Performance & tracing tools
**L1 — aware & can use**
- ⬜ **13.1 CPU & load** — `uptime`, load average (1/5/15), `top`, `htop`, `%us`/`%sy`/`%wa`/`%st`, `mpstat -P ALL 1`, `nproc`
- ⬜ **13.2 Memory** — `free -h`, used vs available vs buff/cache, swap usage, `vmstat 1` (`si`/`so`), `/proc/meminfo`, OOM killer messages in `dmesg`
- ⬜ **13.3 Per-process tools** — `pidstat 1`, `lsof -p <pid>`, `strace -p <pid>`, `strace -c`, `ps -o rss`, `pmap`

**L2 — deep & efficient**
- ⬜ **13.4 System-wide statistics** — `sysstat` package, `sar` (historical data), `vmstat` columns (`r`, `b`, `cs`, `in`), `iostat -xz 1`, `sar -n DEV 1`, `sar -n TCP,ETCP 1`, USE method (utilization, saturation, errors)
- ⬜ **13.5 Profiling & tracing** — `perf top`, `perf record`/`perf report`, `perf stat`, flame graphs, `bpftrace` and bcc tools (`execsnoop`, `opensnoop`, `biolatency`), overhead of `strace` vs eBPF, Go `pprof` as the app-level complement

🛠 **Practice** — L1: run a CPU-heavy and a memory-heavy Go program and identify each one using only `top`, `vmstat 1` and `free -h` · L2: record `perf record -g` on a busy Go program, build a flame graph, and compare it with the `pprof` CPU profile

---

## LX 14 — Troubleshooting a server
**L1 — aware & can use**
- ⬜ **14.1 The 60-second checklist (Brendan Gregg)** — `uptime`, `dmesg -T | tail`, `vmstat 1`, `mpstat -P ALL 1`, `pidstat 1`, `iostat -xz 1`, `free -m`, `sar -n DEV 1`, `sar -n TCP,ETCP 1`, `top`
- ⬜ **14.2 Common incidents** — disk full, out of inodes, OOM kill, high load but idle CPU (I/O wait), port already in use, too many open files, DNS failure, clock skew (`timedatectl`), expired TLS certificate

**L2 — deep & efficient**
- ⬜ **14.3 Systematic debugging** — USE method, RED method, start from symptoms, form a hypothesis, change one thing at a time, timeline of the incident, writing a postmortem, safe commands on production (read-only first)
- ⬜ **14.4 Debugging inside containers** — `docker exec`, `docker logs`, `nsenter -t <pid> -n`, debug sidecar / `kubectl debug` (awareness), missing tools in distroless images, reading cgroup stats (`memory.events`, `cpu.stat`)

🛠 **Practice** — L1: run the full 60-second checklist on your laptop while a stress program runs (`stress-ng` or a Go loop) and write one line of interpretation per command · L2: reproduce 3 incidents in Docker (fill the disk, exhaust fds, OOM kill) and diagnose each one only from the shell, then write a short postmortem

---

## LX 15 — tmux & vim basics
**L1 — aware & can use**
- ⬜ **15.1 tmux** — session, window, pane, prefix `Ctrl+B`, `tmux new -s`, detach `d`, `tmux attach -t`, `tmux ls`, split panes (`%`, `"`), scroll mode `[`
- ⬜ **15.2 vim survival** — modes (normal, insert, visual, command), `i`, `Esc`, `:w`, `:q`, `:wq`, `:q!`, `dd`, `yy`, `p`, `u`, `/search`, `:%s/old/new/g`, `nano` as fallback

**L2 — deep & efficient**
- ⬜ **15.3 Productive editing & sessions** — vim motions (`w`, `b`, `gg`, `G`, `ci"`), text objects, `.` repeat, macros (awareness), `~/.vimrc` basics, `~/.tmux.conf`, tmux on remote servers to survive SSH disconnects

🛠 **Practice** — L1: start a long-running command inside tmux on a remote (or local) shell, detach, reconnect, and reattach; edit a config file with vim without leaving to the mouse · L2: complete `vimtutor` and create a minimal `~/.tmux.conf` with mouse support and a custom prefix
