# Linux for Backend Engineers — Practice

Practice for `roadmap/linux.md` is **concrete terminal exercises**: shell commands, one-liners, bash scripts, systemd unit files, and "break it then fix it" incident drills. Do risky exercises (users, disks, systemd, firewall, filling the disk) inside a throwaway container or VM, not on your main system.

## Setup
```bash
# Safe sandbox (throwaway Ubuntu container) for most exercises
docker run --rm -it --name lx -v "$PWD":/work -w /work ubuntu:24.04 bash
apt update && apt install -y less vim tmux jq curl dnsutils iproute2 procps sysstat lsof strace tcpdump shellcheck

# For systemd / journald / timers exercises use your host or a VM
# (e.g. `multipass launch --name lx` or any local VM), since plain containers do not run systemd.
```

## Folder convention
- `code/linux/<NN>-<topic-slug>/` — `NN` is the LX topic number from the roadmap (e.g. `04-text-processing/`, `08-systemd/`).
- `README.md` in that folder holds the exercise instructions, created when the topic starts from the 🛠 **Practice** line in the roadmap.
- Next to it: your files — `*.sh` scripts (must pass `shellcheck`), `*.service` / `*.timer` units, sample data (e.g. `access.log`), and `commands.md` with the commands you ran and what the output means.
- Notes go to `notes/linux/`. Networking theory lives in the networking track; here we only practice the commands.
