# 10 — DevOps & Infrastructure

> Backend engineers are expected to understand the deployment and operational environment of their code.

---

## Docker

**Key Terms:**
- `image` — read-only template; layers stacked on top of each other; built from Dockerfile
- `container` — running instance of an image; isolated process with its own filesystem/network
- `Dockerfile` — text file with instructions to build an image (FROM, RUN, COPY, CMD, ENTRYPOINT)
- `FROM` — base image to build on (`FROM golang:1.22-alpine`)
- `RUN` — execute command during build (install packages, compile)
- `COPY / ADD` — copy files from host into image
- `CMD` — default command to run when container starts; can be overridden
- `ENTRYPOINT` — fixed executable; CMD becomes arguments to it
- `WORKDIR` — set working directory inside container
- `EXPOSE` — document which port container listens on (doesn't publish; informational)
- `ENV` — set environment variable in image
- `ARG` — build-time variable (not in final image)
- `layer caching` — each instruction is a layer; unchanged layers are cached; order matters (change-often layers last)
- `multi-stage build` — multiple FROM; first stage compiles, second stage contains only binary
- `image tag` — `image:tag`; default tag is `latest`; always use explicit versions in production
- `image registry` — Docker Hub, ECR (AWS), GCR (GCP), GHCR (GitHub)
- `volume` — persistent storage mounted into container; data survives container restart
- `bind mount` — mount host directory into container (dev: mount source code)
- `network` — containers communicate via Docker networks; bridge (default), host, none
- `docker build -t name .` — build image from Dockerfile in current directory
- `docker run -p 8080:8080 name` — run container, map host:container ports
- `docker exec -it container bash` — open shell in running container
- `docker logs container` — view container stdout/stderr
- `docker ps` — list running containers; `-a` for all including stopped
- `.dockerignore` — exclude files from build context (like .gitignore)

**Multi-stage build example (Go):**
```dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/server

FROM alpine:3.19
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
```

**Must understand:**
- [ ] Multi-stage builds — why (final image only has binary, not Go toolchain = tiny image)
- [ ] Layer caching — why COPY go.mod before COPY . (dependencies cached separately from source)
- [ ] CMD vs ENTRYPOINT — what each does and how they interact
- [ ] Why use non-root user in container (`USER nobody`)

---

## Docker Compose

**Key Terms:**
- `docker-compose.yml` — defines multi-container application
- `service` — one container definition; name, image/build, ports, volumes, env, depends_on
- `depends_on` — start order; doesn't wait for service to be ready (use healthcheck for that)
- `healthcheck` — command to check if service is ready; other services wait for healthy status
- `network` — compose creates default network; all services can reach each other by service name
- `volumes` — named volumes for persistence; can be shared between services
- `environment` — set env vars; prefer `env_file` for secrets
- `env_file` — load env vars from .env file; `.env` auto-loaded by compose
- `profiles` — run only specific services; `docker compose --profile debug up`
- `override file` — `docker-compose.override.yml` auto-merged; use for dev-specific config

**Common commands:**
- `docker compose up -d` — start all services in background
- `docker compose down` — stop and remove containers
- `docker compose down -v` — also remove volumes (data loss!)
- `docker compose logs -f service` — follow logs
- `docker compose exec service bash` — shell in running service
- `docker compose ps` — status of all services

**Must understand:**
- [ ] depends_on only controls order, not readiness — use healthcheck for true readiness
- [ ] Service names are DNS names — app container reaches postgres by hostname "postgres"
- [ ] Volumes persist data between `docker compose down` / `up` cycles
- [ ] .env file auto-loaded for variable substitution in compose file

---

## Kubernetes Basics

**Key Terms:**
- `cluster` — set of nodes (machines) managed by Kubernetes
- `node` — VM or physical machine; runs pods; types: control plane and worker
- `control plane` — manages cluster: API server, etcd, scheduler, controller manager
- `pod` — smallest deployable unit; one or more containers sharing network and storage
- `container` — Docker container inside pod; pods usually have one container
- `deployment` — manages set of identical pods; desired state declaration; rolling updates
- `replica` — copy of pod; deployment maintains desired number
- `rolling update` — replace pods one by one; zero-downtime deployment
- `rollback` — revert to previous deployment version
- `service` — stable network endpoint for set of pods; load balances between pods
- `ClusterIP` — internal service; only reachable within cluster
- `NodePort` — expose service on each node's IP at port
- `LoadBalancer` — cloud load balancer; external traffic → service → pods
- `ingress` — HTTP routing rules; hostname/path → service; terminates TLS
- `ingress controller` — processes ingress rules (Nginx, Traefik, AWS ALB)
- `ConfigMap` — non-secret configuration; mounted as env vars or files
- `Secret` — sensitive data (passwords, tokens); base64-encoded (NOT encrypted by default)
- `PersistentVolume (PV)` — cluster-level storage resource
- `PersistentVolumeClaim (PVC)` — request for storage by pod
- `namespace` — virtual cluster; isolate resources by team/env
- `kubectl` — CLI for interacting with cluster
- `YAML manifest` — declarative resource definition; apply with `kubectl apply -f`
- `pod lifecycle` — Pending → Running → Succeeded/Failed
- `liveness probe` — is container alive? Restart if fails
- `readiness probe` — is container ready for traffic? Remove from service endpoints if fails
- `resource requests` — minimum CPU/memory pod needs (used for scheduling)
- `resource limits` — maximum CPU/memory pod can use (enforced by Linux cgroups)
- `HPA (Horizontal Pod Autoscaler)` — automatically scale replicas based on CPU/memory/custom metrics
- `StatefulSet` — like Deployment but for stateful apps; stable pod names, ordered updates
- `DaemonSet` — run one pod per node (log collectors, monitoring agents)

**Must understand:**
- [ ] Pod vs Container — pod is the unit; containers in a pod share localhost
- [ ] Deployment vs StatefulSet — stateless vs stateful apps
- [ ] Liveness vs readiness probe — liveness restarts container; readiness removes from LB
- [ ] Resource requests vs limits — requests for scheduling; limits for enforcement
- [ ] How rolling update works and how to tune it (maxUnavailable, maxSurge)

---

## CI/CD

**Key Terms:**
- `CI (Continuous Integration)` — automatically build and test on every commit; catch bugs early
- `CD (Continuous Delivery)` — always have deployable artifact; deploy manually or on approval
- `CD (Continuous Deployment)` — automatically deploy every passing commit to production
- `pipeline` — sequence of stages: build → test → lint → security scan → deploy
- `trigger` — event that starts pipeline (push to branch, PR opened, tag created)
- `artifact` — output of build stage (Docker image, binary, JAR)
- `artifact registry` — store and version artifacts (ECR, Artifactory, GitHub Packages)
- `GitHub Actions` — CI/CD as YAML in `.github/workflows/`
- `GitLab CI` — `.gitlab-ci.yml`; stages, jobs, artifacts
- `Jenkins` — self-hosted; Jenkinsfile; highly customizable; complex to maintain
- `runner` — machine that executes pipeline jobs
- `cache` — cache dependencies between pipeline runs (node_modules, Go module cache)
- `environment` — named target (dev, staging, production); protection rules
- `deployment gate` — manual approval before production deployment
- `branch protection` — require CI to pass before PR can merge
- `monorepo CI` — only build/test changed services; use path filters
- `blue-green deployment` — two identical envs; switch traffic; zero downtime
- `canary deployment` — route % of traffic to new version; monitor; gradually increase
- `feature flags` — deploy code but hide behind flag; enable for subset of users

**GitHub Actions concepts:**
- `workflow` — YAML file; triggered by events
- `job` — runs on a runner; jobs run in parallel by default
- `step` — individual command or action within a job; sequential
- `action` — reusable step unit; `uses: actions/checkout@v4`
- `secret` — encrypted variable; `${{ secrets.AWS_KEY }}`
- `matrix` — run job with multiple configurations (Go versions, OS types)

**Must understand:**
- [ ] Why CI catches bugs earlier (shift-left testing)
- [ ] Difference between CD (delivery) vs CD (deployment) — delivery still has manual gate
- [ ] Blue-green vs canary — blue-green is all-or-nothing; canary is gradual
- [ ] Feature flags — allows deploying without releasing; separate deploy from release

---

## Cloud Basics (AWS/GCP)

**Key Terms:**
- `region` — geographic area with multiple data centers (us-east-1, eu-west-1)
- `availability zone (AZ)` — isolated data center within region; replicate across AZs for HA
- `VPC (Virtual Private Cloud)` — isolated network in cloud
- `subnet` — subdivision of VPC; public (has internet gateway) or private
- `security group` — stateful firewall for EC2/RDS; rules for inbound/outbound
- `IAM (Identity and Access Management)` — users, roles, policies; least privilege
- `IAM role` — permissions assigned to AWS service (EC2 can read S3 without credentials in code)
- `EC2 (Elastic Compute Cloud)` — virtual machine; choose instance type (t3.micro, c6i.large)
- `ECS (Elastic Container Service)` — managed Docker container service; Fargate = serverless
- `EKS (Elastic Kubernetes Service)` — managed Kubernetes
- `Lambda` — serverless function; triggered by events; pay per execution; cold start
- `S3 (Simple Storage Service)` — object storage; buckets; versioning; lifecycle policies
- `RDS` — managed relational DB (PostgreSQL, MySQL, etc.); automated backups, multi-AZ
- `ElastiCache` — managed Redis/Memcached
- `SQS (Simple Queue Service)` — managed message queue
- `SNS (Simple Notification Service)` — managed pub/sub; fan-out to SQS, Lambda, email
- `CloudFront` — CDN; integrates with S3 and custom origins
- `Route 53` — DNS service; health checks; routing policies (simple, weighted, latency, failover)
- `ALB (Application Load Balancer)` — L7 HTTP load balancer
- `NLB (Network Load Balancer)` — L4 TCP load balancer; ultra-high performance
- `Auto Scaling Group (ASG)` — automatically scale EC2 instances based on metrics
- `CloudWatch` — monitoring, logging, alerting; metrics, logs, alarms

**Must understand:**
- [ ] When to use EC2 vs ECS vs Lambda vs EKS — complexity vs control
- [ ] IAM roles vs access keys — roles are preferred (no static credentials)
- [ ] Multi-AZ deployment — why it matters for high availability
- [ ] S3 vs EBS vs EFS — object storage vs block storage vs file system

---

## Infrastructure as Code

**Key Terms:**
- `IaC (Infrastructure as Code)` — define infrastructure in code; version-controlled; reproducible
- `declarative` — describe desired state; tool figures out how to get there (Terraform, Kubernetes)
- `imperative` — describe steps to take; order matters (shell scripts, Ansible tasks)
- `Terraform` — most popular IaC tool; declarative; supports multiple cloud providers
- `HCL (HashiCorp Configuration Language)` — Terraform's language
- `provider` — Terraform plugin for cloud provider (aws, google, kubernetes)
- `resource` — infrastructure component to manage (`resource "aws_instance" "web" {}`)
- `state` — Terraform tracks current infrastructure in state file; must be stored remotely (S3)
- `state file` — `.tfstate`; never commit to git; contains sensitive data
- `plan` — `terraform plan` shows what will change without applying
- `apply` — `terraform apply` creates/modifies/destroys resources to match config
- `destroy` — `terraform destroy` removes all managed resources
- `module` — reusable Terraform configuration; encapsulates related resources
- `remote state backend` — store state in S3+DynamoDB lock; enables team collaboration
- `drift` — when actual infrastructure diverges from Terraform state (someone changed manually)
- `Pulumi` — IaC with real programming languages (Go, Python, TypeScript)
- `Ansible` — configuration management; agentless; SSH-based; YAML playbooks
- `Helm` — Kubernetes package manager; templates + values files for K8s manifests

**Must understand:**
- [ ] Why remote state is essential for teams (concurrent access, locking)
- [ ] terraform plan before apply — never apply without reviewing plan
- [ ] Modules — how to create reusable infrastructure components
- [ ] Why IaC over ClickOps — reproducible, reviewable, version history, disaster recovery

---

## Observability

**Key Terms:**
- `three pillars` — Logs, Metrics, Traces
- `logs` — text records of events; structured (JSON) vs unstructured; searchable
- `metrics` — numeric measurements over time (request rate, error rate, latency, CPU)
- `traces` — record of request path through multiple services; shows where time is spent
- `structured logging` — JSON logs with consistent fields (level, timestamp, request_id, user_id)
- `log aggregation` — collect logs from all instances centrally (Loki, Elasticsearch, CloudWatch Logs)
- `correlation ID / trace ID` — unique ID per request; included in all logs; enables tracing
- `span` — single unit of work in a trace; has start time, duration, parent span
- `trace` — collection of spans for one request; shows full request path
- `distributed tracing` — trace spans across multiple services; propagate trace context in headers
- `RED method` — Rate, Errors, Duration — three key metrics for any service
- `USE method` — Utilization, Saturation, Errors — for resources (CPU, disk, memory)
- `SLI (Service Level Indicator)` — metric that measures service quality (request success rate)
- `SLO (Service Level Objective)` — target for SLI (99.9% success rate)
- `error budget` — how much downtime/errors are allowed before SLO is breached
- `alerting` — trigger notification when metric crosses threshold
- `dashboard` — visual display of metrics over time

**Must understand:**
- [ ] Why all three pillars are needed — logs: what happened; metrics: how much/how fast; traces: where
- [ ] Structured logging — why JSON logs are better than plain text for querying
- [ ] Correlation ID — how to propagate through services (HTTP header X-Request-ID)
- [ ] RED metrics — every service should expose these three

---

## Prometheus + Grafana

**Key Terms:**
- `Prometheus` — time-series metrics database; pull-based (scrapes /metrics endpoints)
- `exporter` — exposes metrics in Prometheus format (node_exporter, postgres_exporter)
- `/metrics endpoint` — HTTP endpoint exposing metrics in Prometheus text format
- `scrape interval` — how often Prometheus polls targets (default 15s)
- `metric types`:
  - `Counter` — monotonically increasing (requests_total, errors_total)
  - `Gauge` — can go up or down (memory_usage, queue_depth)
  - `Histogram` — samples with bucket distribution (request_duration_seconds)
  - `Summary` — calculates quantiles client-side
- `label` — key-value pair added to metric (method="GET", status="200")
- `PromQL` — Prometheus Query Language; `rate(requests_total[5m])` = requests per second over 5m
- `rate()` — per-second rate of increase for counter; use with counters
- `irate()` — instant rate; based on last two data points
- `histogram_quantile(0.99, ...)` — calculate p99 from histogram
- `Alertmanager` — receives alerts from Prometheus; routes to PagerDuty, Slack, email
- `recording rule` — pre-compute expensive queries; store as new metric
- `Grafana` — dashboard tool; connects to Prometheus; visualize metrics
- `dashboard` — collection of panels (graphs, stats, tables)
- `data source` — Prometheus, Loki, Elasticsearch, etc.

**Must understand:**
- [ ] Counter vs Gauge — counter never decreases; gauge can go up and down
- [ ] Why use rate() for counters (raw counter value is useless; rate shows current throughput)
- [ ] Labels — how to slice metrics (per endpoint, per status code)
- [ ] p99 latency from histogram — `histogram_quantile(0.99, rate(http_duration_seconds_bucket[5m]))`

---

## Git Advanced

**Key Terms:**
- `rebase` — replay commits on top of another branch; linear history
- `merge` — combine branches; creates merge commit; preserves history
- `interactive rebase` — `git rebase -i HEAD~3` — squash, reorder, edit commits
- `squash` — combine multiple commits into one; clean up before merging
- `cherry-pick` — apply specific commit to current branch: `git cherry-pick abc123`
- `bisect` — binary search through commits to find which introduced a bug
- `reflog` — log of all HEAD movements; recover "lost" commits
- `stash` — temporarily save uncommitted changes; `git stash` / `git stash pop`
- `tag` — named pointer to commit; semantic versioning releases
- `detached HEAD` — HEAD points to commit, not branch; commits not on any branch
- `force push` — `git push --force-with-lease`; overwrite remote branch (dangerous on shared branches)
- `branching strategies`:
  - `Gitflow` — feature, develop, release, hotfix branches; complex
  - `trunk-based` — short-lived branches; merge to main frequently; preferred for CI/CD
  - `GitHub flow` — main + feature branches + PR; simple and effective
- `conventional commits` — `feat:`, `fix:`, `chore:`, `docs:` — machine-parseable commit messages
- `semantic versioning` — MAJOR.MINOR.PATCH based on conventional commits
- `git hooks` — scripts that run before/after git events (pre-commit: run linter, tests)
- `CODEOWNERS` — GitHub file; auto-assign reviewers based on file paths

**Must understand:**
- [ ] Rebase vs merge — when each is appropriate (rebase for clean history, merge for preserving context)
- [ ] Interactive rebase — how to squash commits before opening PR
- [ ] git bisect — how to use it to find regression-introducing commit
- [ ] Trunk-based development — why it works better than Gitflow for CI/CD
