# OPS 22 — Deployment track: ship-lab

Roadmap: `roadmap/devops.md` → OPS 22 (week-by-week order) · Notes: `notes/linux/`, `notes/networking/`, `notes/devops/`

## Goal
Simulate the work project's production setup **on the laptop first**, then on 1–2 cheap VPSs only for the parts that need the real internet (public DNS, Let's Encrypt, GitHub Actions deploy).

## Local "servers" = Multipass VMs (same shape as production)
| VM | Production role | Create |
|---|---|---|
| `sat-app` | app VM: 2× `serve` behind Caddy | week 1 |
| `sat-db` | `worker` + Postgres primary | week 5 |
| `sat-db2` | Postgres replica | week 6 |

```bash
# install: https://multipass.run (macOS / Windows / Linux)
multipass launch 24.04 --name sat-app --cpus 1 --memory 1G --disk 8G
multipass list                 # shows the VM IP
multipass shell sat-app        # log in (later you will use real ssh)
multipass stop sat-app         # free RAM when done
multipass delete --purge sat-app   # throw away and start again (do it often!)
```
Quick pieces (users, permissions) can also run in a throwaway container: `docker run -it --rm ubuntu:24.04 bash`.
Why VMs and not containers for the server part: SSH daemon, firewall (`ufw`), systemd and kernel settings behave like a real server only in a VM.

## Labs
### Lab 1 — LX 3.1–3.3: users, groups, permissions (day 1)
Where: `docker run -it --rm ubuntu:24.04 bash` (then `apt update && apt install -y sudo`) or `sat-app`.
1. Create a group `app` and two users: `appuser` (in group `app`) and `guest` (not in it). Tools: `groupadd`, `useradd -m -G`, `id`.
2. Create `/srv/app/` owned by `root:app`, where the group can write and others can do nothing. Tools: `mkdir -p`, `chown`, `chmod` (think: which octal number?).
3. Create `/srv/app/.env` with a fake password. Only `appuser` may read it.
4. Prove it (expected result in brackets):
   - `sudo -u appuser touch /srv/app/ok` (works)
   - `sudo -u guest touch /srv/app/no` (Permission denied)
   - `sudo -u guest cat /srv/app/.env` (Permission denied)
   - `sudo -u appuser cat /srv/app/.env` (prints the password)
5. Write in this README (Results): the `ls -l` output of `/srv/app` and one sentence: why is `x` needed on a directory?

**Done when:** all 4 checks give the expected result and you can read every column of `ls -l` aloud.
**L2 stretch (LX 3.5):** create a system user `sat` with no login shell (`useradd --system --shell /usr/sbin/nologin`) and try `su - sat` — what happens and why is this good for a service?

### Lab 2 — weeks 1–2: harden `sat-app` by hand
*(details added when LX 12 starts: deploy user, SSH keys only, root login off, `ufw`, `fail2ban`, `unattended-upgrades`, a Go binary as a systemd service)*

## Results
| Lab | Date | Result / notes |
|---|---|---|
