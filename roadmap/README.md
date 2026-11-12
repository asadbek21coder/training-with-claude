# Roadmap — All Tracks

> **L1** = aware & can use · **L2** = deep & efficient · Status per subtopic: ⬜ not started · 🔄 in progress · ✅ done
> Pick one small piece a day — a subtopic, or even one term — by its ID (e.g. *"today GO 8.4"*).

## Overview & progress

| Track | File | Prefix | Role now | Topics | L1 done / total | L2 done / total |
|---|---|---|---|---|---|---|
| Go | [golang.md](golang.md) | GO | 🔥 **main** | 24 | 1 / 85 | 0 / 72 (+ L3 book: 0 / 33) |
| Algorithms & DS | [algorithms-ds.md](algorithms-ds.md) | DSA | 📅 daily | 20 | 13 / 47 | 0 / 42 |
| English | [english.md](english.md) | EN | 📅 daily | 8 | 0 / 24 | 0 / 14 |
| Russian | [russian.md](russian.md) | RU | 📅 daily | 14 | 2 / 41 | 0 / 30 |
| PostgreSQL & data stores | [postgres.md](postgres.md) | PG | ⏭ next | 15 | 0 / 40 | 0 / 50 |
| System design & distributed systems | [system-design.md](system-design.md) | SD | ⏭ later | 19 | 8 / 40 | 0 / 56 |
| Operating systems | [os.md](os.md) | OS | ⏭ later | 11 | 0 / 31 | 0 / 22 |
| Linux | [linux.md](linux.md) | LX | ⏭ later | 15 | 0 / 41 | 0 / 29 |
| Networking | [networking.md](networking.md) | NET | ⏭ later | 21 | 0 / 50 | 0 / 32 |
| DevOps, cloud & SRE | [devops.md](devops.md) | OPS | ⏭ later · ⏱ **OPS 22 now** (deployment track) | 22 | 0 / 66 | 0 / 54 |
| AI engineering (using AI in SWE + building AI features) | [ai.md](ai.md) | AI | 📅 daily (AI 1–5) | 12 | 0 / 24 | 0 / 21 |
| Common SWE (Git, auth, crypto, security, API design, career) | [common-swe.md](common-swe.md) | SWE | ⏭ later | 15 | 0 / 50 | 0 / 40 |
| QA & testing | [qa.md](qa.md) | QA | ⏭ later | 19 | 0 / 44 | 0 / 27 |
| Design patterns & architecture | [design-patterns.md](design-patterns.md) | DP | ⏭ later | 13 | 0 / 38 | 0 / 29 |
| **Total** | | | | **228** | **24 / 621** | **0 / 518** |

## Order
```
NOW:     golang — all L1 first (refresh + measure), then L2
DAILY:   algorithms-ds · english · russian · ai (AI 1–5) — one small piece each, every day
AI 6–12: after Go L1 (building AI features)
TIME-BOX: devops OPS 22 (deployment track) — before the work project's first deploy (spring 2027); weekly order in devops.md OPS 22
NEXT:    postgres → system-design → os → linux → networking → devops → common-swe → qa → design-patterns
```
Any track, any day is fine for a quick piece (e.g. one term from common-swe while reading about JWT at work) — the order is for the **main** 40–50 minute block.

## Where things go
- Learned something → `notes/<track>/<NN>-<topic>.md` (format: `notes/README.md`)
- Practiced something → `code/<track>/<NN>-<topic>/` (instructions in its README)
- Finished a subtopic → change ⬜ to ✅ here **and** in the track file's Progress table, add a row to `LOG.md`
