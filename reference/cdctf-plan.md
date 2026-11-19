# cdCTF.uz — personal plan (secure backend / DevSecOps angle)

Platform: https://cdctf.uz — Uzbek cybersecurity education platform (cdCTF Core Team): free lessons in Uzbek / Russian / English, Turan Security video academy (6 tracks, 155 videos: Linux, web security, networking, bug bounty, pentesting, cyber defense), 8 numbered curriculum modules, a 23-chapter networking curriculum, 106+ CTF challenges, browser labs, national ranking, cdScan scanner.
Analysed 2026-11-19.

## Fit
The goal is **not** to become a pentester or SOC analyst. Use the platform for what makes a **secure backend / DevOps engineer** and protects SAT ERP:
- **High value:** networking curriculum, module 3 (protocols + Wireshark), module 2 (Linux), module 5 (OWASP web), module 6 (crypto), easy web/Linux/network CTFs.
- **Medium:** module 4 (recon, Nmap) on your own VM only, parts of module 8 (incident response, logs) → OPS 18–19.
- **Skip for now:** module 7 (Windows / Active Directory), bug bounty track, SOC Analyst L1 path, deep exploitation.
- **Caveats:** the certificate's "internationally recognized" claim is not verified — don't plan your EU/FAANG CV around it; your projects and CTF results say more. **Legal:** attack only your own machines, platform labs or explicitly allowed targets; cdScan only on sites you own.

## Plan (fits the daily card: 45–60 min/day, ~14 weeks)
| Weeks | cdCTF part | Roadmap IDs | Practice on your own stack |
|---|---|---|---|
| 1–4 | Networking curriculum ch. 1–12 + module 3 (Wireshark) | NET 1–8 | Wireshark/`tcpdump` on your laptop: capture `curl https://…` → DNS, TCP handshake, TLS; subnet math drills |
| 5–6 | Module 2 + Linux video track | LX 3, LX 8, LX 11, LX 12 | OPS 22 Lab 1–2: users/permissions, systemd, SSH keys only, `ufw` on the `sat-app` VM |
| 7–10 | Module 5 (OWASP Top 10) + web security videos + easy web CTFs | SWE 10, 11, 12 | Attack **local** SAT ERP on the VM: IDOR between tenants (most important for a multi-tenant ERP), SQLi, broken auth/session, CSRF, file upload → write a fix + a test for each |
| 11–12 | Networking ch. 13–23 (TLS, HTTP/2/3, CDN, VPN/WireGuard, firewall/IDS/WAF, DDoS, Docker/K8s networking) | NET 10, 13, 14, 17, 19, 20 | Caddy + TLS in front of SAT ERP; WireGuard/Tailscale for teammates (OPS 22 step 5) |
| 13 | Module 6 (crypto, hashing, AES/RSA, Hashcat) | SWE 9, NET 13 | Check how SAT ERP stores passwords (argon2id/bcrypt?), crack a weak test hash with Hashcat to see why |
| 14 | Module 4 (recon, Nmap) + parts of module 8 (incident response) | NET 20, OPS 18–19 | Nmap your own VM: which ports are open and why; write a 1-page incident runbook |

**CTFs:** from week 3, 2 easy challenges a week (network → Linux → web). Log them in `notes/` with the trick you learned. Aim for the national ranking only after week 10.

## How it plugs into the daily card
Networking days use cdCTF chapters **or** Kurose videos (same topic, pick one). Linux days = OPS 22 labs. Web security weeks replace the system-design slot. Tick the roadmap IDs above as usual.
