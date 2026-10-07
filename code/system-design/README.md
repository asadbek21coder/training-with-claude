# System Design Practice

Hands-on work for `roadmap/system-design.md` (SD 1–19, incl. distributed systems & messaging).

## What practice looks like
- **Classic designs (SD 8, 17, 18):** one design doc per system at `code/system-design/<NN>-<design>/design.md`
  (e.g. `08-url-shortener/design.md`), written with the framework below. Diagrams in Mermaid inside the doc.
- **Building blocks & distributed topics (SD 3–7, 9–16):** small Go programs + docker compose experiments —
  run a 3-node setup, kill or partition a node (`docker stop`, `docker network disconnect`, Toxiproxy, `tc netem`),
  measure, and write conclusions in `findings.md`.

## Design doc framework (every `design.md`)
1. **Requirements** — functional, non-functional (SLO, latency, consistency), out of scope
2. **Estimates** — DAU, QPS (avg/peak), read/write ratio, storage/year, bandwidth
3. **API** — endpoints / messages with request and response
4. **Data model** — entities, storage choice and why, shard/partition key
5. **High-level design** — Mermaid diagram, write path and read path
6. **Deep dives** — 2–3 hardest parts (scaling, consistency, hot spots)
7. **Failure modes** — table: what breaks → impact → mitigation
8. **Trade-offs & alternatives** — what you chose, what you rejected, why
9. **Wrap-up** — bottlenecks, what v2 would change, how to monitor it

## Environment
```bash
docker compose up -d          # per-exercise compose file (etcd, Kafka, Redis, Postgres, Nginx, Toxiproxy...)
docker network disconnect <net> <container>   # simulate a partition
go run ./...                  # Go experiments; load testing with k6 or hey
```

## Folder convention
`code/system-design/<NN>-<topic-slug>/README.md` holds the exercise instructions (created when the topic starts,
from the 🛠 Practice line in the roadmap); `design.md`, Go code, `docker-compose.yml` and `findings.md` sit next to it.
`NN` = topic number from the roadmap; classic designs use the topic number + design name (`17-payment-system`).
