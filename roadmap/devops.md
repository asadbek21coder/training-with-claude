# DevOps, Cloud & SRE Roadmap

> **L1** = aware & can use · **L2** = deep & efficient · Status: ⬜ not started · 🔄 in progress · ✅ done
> **Notes:** `notes/devops/` · **Practice:** `code/devops/` · **Dictionary:** `reference/devops.md`

*(Linux & shell is its own `linux` track; Git is in the `common-swe` track.)*

**Why (EN):** "You build it, you run it": a senior backend engineer takes a Go service from `git push` to production — build, containerize, deploy, observe, debug and roll back — without waiting for a separate ops team. In the AI era, writing YAML and Terraform is cheap (AI drafts it in seconds), so the value moves to judgment: reading a plan before `apply`, designing safe rollouts, setting SLOs, debugging a broken pod at 3 a.m., and reviewing AI-generated infra for security and cost mistakes.
**Зачем (RU):** «Ты построил — ты и эксплуатируешь»: senior backend-инженер сам доводит Go-сервис от `git push` до продакшена — собирает, упаковывает, деплоит, наблюдает, отлаживает и откатывает. В эпоху ИИ написать YAML и Terraform дёшево (ИИ сделает черновик за секунды), поэтому ценность смещается к суждению: прочитать план перед `apply`, спроектировать безопасный выкат, задать SLO, отладить упавший под ночью и проверить сгенерированную ИИ инфраструктуру на ошибки безопасности и стоимости.

## Progress
| Layer | Done | Total |
|---|---|---|
| L1 | 0 | 65 |
| L2 | 0 | 54 |

---

## OPS 1 — Containers: what they really are
**L1 — aware & can use**
- ⬜ **1.1 Container vs VM** — process isolation, shared kernel, hypervisor, startup time, density, image vs container, immutability, "works on my machine"
- ⬜ **1.2 The container ecosystem** — Docker Engine, `dockerd`, Docker CLI, containerd, runc, Podman, Buildah, OCI image spec, OCI runtime spec, OCI distribution spec

**L2 — deep & efficient**
- ⬜ **1.3 Linux primitives** — namespaces (pid, net, mnt, uts, ipc, user, cgroup), cgroups v2, CPU/memory controllers, `unshare`, `nsenter`, seccomp, capabilities, chroot/pivot_root
- ⬜ **1.4 Image internals** — layers, overlayfs (lowerdir/upperdir/merged), copy-on-write, image manifest, config blob, content-addressable digest (`sha256:`), manifest list / multi-arch index
- ⬜ **1.5 Runtime stack** — CRI, containerd shim, runc lifecycle, rootless containers, gVisor / Kata Containers (sandboxed runtimes), PID 1 problem, signal handling, zombie reaping, `tini` / `--init`

🛠 **Practice** — L1: run `alpine` and `nginx` containers, then use `docker inspect` and `ps` on the host to find the container process and compare it with a normal process · L2: build a "container" by hand with `unshare --pid --net --mount --fork` + `chroot` into an extracted alpine rootfs, and set a memory limit through cgroup v2 files

---

## OPS 2 — Dockerfile & images
**L1 — aware & can use**
- ⬜ **2.1 Dockerfile instructions** — `FROM`, `RUN`, `COPY` vs `ADD`, `WORKDIR`, `ENV`, `ARG`, `EXPOSE`, `USER`, `CMD` vs `ENTRYPOINT`, exec vs shell form, `HEALTHCHECK`, `.dockerignore`
- ⬜ **2.2 Building Go images** — multi-stage build, `CGO_ENABLED=0`, static binary, `scratch`, distroless (`gcr.io/distroless/static`), alpine vs debian-slim, CA certificates, timezone data, `-ldflags "-s -w"`
- ⬜ **2.3 Tags, digests & registries** — tag vs digest, `latest` anti-pattern, immutable tags, Docker Hub, GHCR, ECR, Artifact Registry, `docker push`/`pull`, rate limits, private registry login

**L2 — deep & efficient**
- ⬜ **2.4 Layer caching & BuildKit** — cache invalidation order, `go mod download` layer, `--mount=type=cache`, `--mount=type=secret`, `docker buildx`, remote cache (`--cache-to`/`--cache-from`), reproducible builds, build context size
- ⬜ **2.5 Image optimization & hardening** — image size, attack surface, non-root `USER`, read-only filesystem, no shell in final image, `dive`, pinning base images by digest, rebuild cadence for CVE patches
- ⬜ **2.6 Multi-arch & alternatives** — `linux/amd64` vs `linux/arm64`, `buildx --platform`, QEMU emulation, Go cross-compilation (`GOOS`/`GOARCH`), `ko`, Buildpacks, Jib (awareness)

🛠 **Practice** — L1: containerize a small Go HTTP service (`/healthz`, `/hello`) with a multi-stage Dockerfile on distroless, run it as non-root · L2: measure image size and rebuild time for 3 variants (golang full, alpine, distroless/scratch), add BuildKit cache mounts, and record the before/after numbers in the exercise README

---

## OPS 3 — Running containers & Docker Compose
**L1 — aware & can use**
- ⬜ **3.1 Container lifecycle** — `docker run`/`ps`/`stop`/`kill`/`rm`, `-d`, `--rm`, `-p host:container`, `-e`, `--env-file`, `docker logs -f`, `docker exec -it`, restart policies (`no`, `on-failure`, `always`, `unless-stopped`)
- ⬜ **3.2 Storage & networking** — named volume, bind mount, tmpfs, bridge network, user-defined network, container DNS by name, `host` network, `none`, port publishing
- ⬜ **3.3 Docker Compose basics** — `compose.yaml`, services, `image` vs `build`, `ports`, `environment`, `volumes`, `networks`, `docker compose up/down/logs/ps`, `-d`, project name

**L2 — deep & efficient**
- ⬜ **3.4 Compose for real local stacks** — `depends_on` with `condition: service_healthy`, `healthcheck`, profiles, override files (`compose.override.yaml`), `.env` interpolation, `develop.watch`, init/migration containers, secrets
- ⬜ **3.5 Resource limits & graceful shutdown** — `--memory`, `--cpus`, OOM kill, `SIGTERM` → `SIGKILL`, `stop_grace_period`, `STOPSIGNAL`, Go `signal.NotifyContext` + `http.Server.Shutdown`, log drivers, `docker stats`
- ⬜ **3.6 Container debugging** — `docker inspect`, `docker events`, exit codes (137, 139, 143), debugging distroless (`docker debug`, sidecar with shared namespaces `--pid=container:`), `nsenter`, disk cleanup (`docker system prune`)

🛠 **Practice** — L1: write a `compose.yaml` with the Go service + Postgres + a migration step, data in a named volume · L2: add Redis, healthchecks with `service_healthy`, memory limits, and prove graceful shutdown works (in-flight request finishes on `docker compose stop`)

---

## OPS 4 — CI/CD pipelines
**L1 — aware & can use**
- ⬜ **4.1 CI/CD concepts** — continuous integration, continuous delivery vs continuous deployment, pipeline stages (lint → test → build → scan → publish → deploy), fail fast, build once deploy many, artifact
- ⬜ **4.2 GitHub Actions** — workflow, `on:` triggers, job, step, `runs-on`, `uses:` actions, `needs`, matrix, `actions/cache`, artifacts, secrets, variables, environments
- ⬜ **4.3 GitLab CI (mapping)** — `.gitlab-ci.yml`, stages, jobs, runners, `rules`, cache vs artifacts, CI/CD variables, protected branches, environments

**L2 — deep & efficient**
- ⬜ **4.4 Fast & reliable pipelines** — caching Go modules and build cache, test parallelism, path filters, `concurrency` groups, flaky test quarantine, self-hosted runners, pipeline duration as a metric
- ⬜ **4.5 Reusable & secure pipelines** — reusable workflows, composite actions, pinning actions by commit SHA, `permissions:` (least privilege `GITHUB_TOKEN`), OIDC to cloud (no long-lived keys), `pull_request_target` risks, secret masking
- ⬜ **4.6 Versioning & promotion** — semantic versioning, git tags, release notes, changelog generation, GoReleaser, dev → staging → prod promotion, manual approvals, environment protection rules, immutable artifacts

🛠 **Practice** — L1: GitHub Actions workflow for the Go service: `go vet`, `golangci-lint`, `go test -race`, `docker build` · L2: add module/build caching, a matrix over 2 Go versions, image push to GHCR with `GITHUB_TOKEN` and minimal `permissions:`, actions pinned by SHA; compare pipeline time before/after caching

---

## OPS 5 — Deployment & release strategies
**L1 — aware & can use**
- ⬜ **5.1 Deployment strategies** — recreate, rolling update, blue-green, canary, shadow / dark launch, A/B testing, downtime vs zero-downtime
- ⬜ **5.2 Feature flags** — release vs deploy, flag types (release, ops/kill switch, experiment, permission), OpenFeature, LaunchDarkly / Unleash / Flagsmith, flag debt

**L2 — deep & efficient**
- ⬜ **5.3 Progressive delivery** — automated canary analysis, metric-based promotion, Argo Rollouts, Flagger, traffic splitting, automated rollback, bake time
- ⬜ **5.4 Safe changes with state** — backward-compatible DB migrations, expand-and-contract, N/N-1 compatibility, API versioning during rollout, schema migration tools (goose, golang-migrate, Atlas), roll forward vs roll back

🛠 **Practice** — L1: draw the 4 main strategies as timelines (traffic % over time) and list the failure mode of each · L2: implement a column rename on the Postgres compose stack using expand-and-contract in 3 deployable steps, with old and new app versions running together

---

## OPS 6 — Kubernetes: architecture & workloads
**L1 — aware & can use**
- ⬜ **6.1 Why Kubernetes** — orchestration, desired state, declarative vs imperative, reconciliation loop, self-healing, when NOT to use k8s (ECS, Cloud Run, a single VM)
- ⬜ **6.2 Cluster architecture** — control plane, API server, etcd, scheduler, controller manager, cloud controller manager, worker node, kubelet, kube-proxy, container runtime (containerd)
- ⬜ **6.3 Core workloads** — Pod, ReplicaSet, Deployment, StatefulSet, DaemonSet, Job, CronJob, labels, selectors, annotations, namespaces
- ⬜ **6.4 kubectl basics** — `kubectl apply -f`, `get`, `describe`, `logs`, `exec`, `port-forward`, `delete`, `-o yaml`, `--dry-run=client`, contexts, `kubeconfig`, `k9s`

**L2 — deep & efficient**
- ⬜ **6.5 API machinery** — resource, kind, apiVersion, `spec` vs `status`, controllers, informers, watch, resourceVersion, optimistic concurrency, finalizers, owner references, garbage collection, server-side apply
- ⬜ **6.6 Scheduling** — scheduler filtering & scoring, node selector, node affinity, pod affinity/anti-affinity, topology spread constraints, taints & tolerations, priority & preemption
- ⬜ **6.7 Extending Kubernetes** — CustomResourceDefinition (CRD), operator pattern, controller-runtime, Kubebuilder, admission webhooks (validating/mutating), ValidatingAdmissionPolicy (CEL)

🛠 **Practice** — L1: create a kind cluster, deploy the Go service as a Deployment with 3 replicas, delete a pod and watch it come back · L2: spread replicas across kind worker nodes with topology spread constraints, then taint a node and observe scheduling; read the events with `kubectl get events --sort-by=.lastTimestamp`

---

## OPS 7 — Kubernetes: config, storage & networking
**L1 — aware & can use**
- ⬜ **7.1 Configuration** — ConfigMap, Secret, env vars vs volume mounts, `envFrom`, immutable ConfigMaps, config reload, Secret is base64 not encryption
- ⬜ **7.2 Services & DNS** — Service, ClusterIP, NodePort, LoadBalancer, headless Service, Endpoints / EndpointSlice, CoreDNS, `svc.cluster.local`, port vs targetPort
- ⬜ **7.3 Ingress & Gateway API** — Ingress, IngressClass, ingress controller (ingress-nginx, Traefik), TLS termination, Gateway API (GatewayClass, Gateway, HTTPRoute), cert-manager

**L2 — deep & efficient**
- ⬜ **7.4 Storage** — Volume, PersistentVolume, PersistentVolumeClaim, StorageClass, dynamic provisioning, access modes (RWO, RWX), reclaim policy, CSI driver, volumeClaimTemplates in StatefulSets, running databases in k8s (trade-offs)
- ⬜ **7.5 Cluster networking internals** — CNI (Calico, Cilium, Flannel), pod-to-pod networking, kube-proxy modes (iptables, IPVS, nftables), eBPF dataplane, NetworkPolicy (default deny), service mesh (Istio, Linkerd) awareness
- ⬜ **7.6 Secrets done right** — encryption at rest (KMS provider), External Secrets Operator, Sealed Secrets, SOPS, Vault agent / CSI secrets store, RBAC on Secrets

🛠 **Practice** — L1: add a ConfigMap and Secret to the Go service, expose it via ClusterIP + Ingress (ingress-nginx on kind), `curl` it from the host · L2: deploy Postgres as a StatefulSet with a PVC, add a default-deny NetworkPolicy (with a CNI that enforces it, e.g. Calico/Cilium on kind) that allows only the app to reach the DB, and prove the block with a test pod

---

## OPS 8 — Kubernetes: health, resources & scaling
**L1 — aware & can use**
- ⬜ **8.1 Probes** — liveness, readiness, startup probe, `httpGet` / `tcpSocket` / `exec` / `grpc` probes, `initialDelaySeconds`, `periodSeconds`, `failureThreshold`, what NOT to check in liveness
- ⬜ **8.2 Requests & limits** — CPU requests vs limits, memory requests vs limits, millicores, QoS classes (Guaranteed, Burstable, BestEffort), OOMKilled, CPU throttling, LimitRange, ResourceQuota
- ⬜ **8.3 Scaling basics** — `replicas`, `kubectl scale`, HorizontalPodAutoscaler (CPU/memory target), metrics-server, rolling update `maxSurge` / `maxUnavailable`

**L2 — deep & efficient**
- ⬜ **8.4 Go in containers** — `GOMAXPROCS` vs CPU limit (container-aware runtime in recent Go versions, `automaxprocs` for older), `GOMEMLIMIT` vs memory limit, GC behavior near the limit, CFS quota throttling, `pprof` in a pod
- ⬜ **8.5 Advanced autoscaling** — HPA with custom/external metrics, KEDA (event-driven, scale to zero), VPA, Cluster Autoscaler, Karpenter, scale-up lag, overprovisioning
- ⬜ **8.6 Availability** — PodDisruptionBudget, graceful shutdown (`preStop`, `terminationGracePeriodSeconds`, endpoint removal race), pod anti-affinity across zones, node drain, voluntary vs involuntary disruptions, rolling node upgrades

🛠 **Practice** — L1: add readiness/liveness probes and requests/limits to the Go service; trigger OOMKilled with a memory-hungry endpoint and read it in `kubectl describe` · L2: load the service with `hey`/k6 and watch the HPA scale; then run `kubectl drain` during load and fix dropped requests with `preStop` sleep + PDB + proper `Shutdown`

---

## OPS 9 — Kubernetes: debugging & operations
**L1 — aware & can use**
- ⬜ **9.1 Pod failure states** — Pending, ContainerCreating, CrashLoopBackOff, ImagePullBackOff / ErrImagePull, OOMKilled, Evicted, Terminating stuck, `CreateContainerConfigError`, exit codes
- ⬜ **9.2 Debugging workflow** — `kubectl describe`, events, `logs --previous`, `exec`, `port-forward`, `top` (metrics-server), `kubectl debug` (ephemeral containers), `kubectl get -w`, `rollout status/history/undo`

**L2 — deep & efficient**
- ⬜ **9.3 Access control & multi-tenancy** — ServiceAccount, Role / ClusterRole, RoleBinding, `kubectl auth can-i`, namespaces as boundaries, Pod Security Standards (privileged/baseline/restricted), Pod Security Admission, `securityContext`
- ⬜ **9.4 Cluster operations** — managed k8s (EKS, GKE, AKS), version skew & upgrades, node pools, etcd backup/restore, Velero, cluster add-ons, cost of running many small clusters vs one big
- ⬜ **9.5 Policy & guardrails** — Kyverno, OPA Gatekeeper, required labels, disallowing `latest` tags, image allowlists, resource limits enforcement

🛠 **Practice** — L1: break the deployment on purpose 5 ways (bad image tag, wrong port in probe, missing Secret, too-low memory limit, unschedulable request) and write the diagnosis steps for each · L2: run the Go service under a restricted Pod Security Standard namespace with a dedicated ServiceAccount and minimal RBAC; verify with `kubectl auth can-i`

---

## OPS 10 — Packaging & GitOps (Helm, Kustomize, Argo CD)
**L1 — aware & can use**
- ⬜ **10.1 Helm** — chart, `Chart.yaml`, `values.yaml`, templates, release, `helm install/upgrade/rollback/uninstall`, `helm template`, repositories, OCI charts
- ⬜ **10.2 Kustomize** — base, overlays, `kustomization.yaml`, patches (strategic merge, JSON 6902), `configMapGenerator`, `kubectl apply -k`, images transformer
- ⬜ **10.3 GitOps concepts** — Git as single source of truth, pull vs push deployment, desired vs live state, drift detection, reconciliation, Argo CD, Flux

**L2 — deep & efficient**
- ⬜ **10.4 Helm in depth** — template functions (Sprig), `_helpers.tpl`, `required`, `values.schema.json`, hooks, chart dependencies, library charts, `helm diff`, Helm vs Kustomize trade-offs
- ⬜ **10.5 Argo CD in practice** — Application, ApplicationSet, AppProject, sync policy (auto-sync, prune, self-heal), sync waves & hooks, health checks, app-of-apps, image updater, secrets with GitOps
- ⬜ **10.6 Repo & environment structure** — app repo vs config repo, environment branches vs directories (directories preferred), promotion via PR, monorepo of manifests, rendered manifests pattern

🛠 **Practice** — L1: package the Go service as a Helm chart with dev/prod values, and the same as Kustomize base + 2 overlays; compare the rendered YAML · L2: install Argo CD in kind, point it at a Git repo of manifests, enable auto-sync + self-heal, then `kubectl edit` a resource manually and watch Argo revert the drift

---

## OPS 11 — Infrastructure as Code: Terraform / OpenTofu
**L1 — aware & can use**
- ⬜ **11.1 IaC concepts** — declarative vs imperative, idempotency, drift, mutable vs immutable infrastructure, ClickOps problems, Terraform vs OpenTofu (license fork), Pulumi, CloudFormation, CDK (awareness)
- ⬜ **11.2 HCL basics** — provider, resource, data source, variable, output, locals, `terraform init/plan/apply/destroy`, `fmt`, `validate`, `.tfvars`
- ⬜ **11.3 State** — state file, what it stores (incl. secrets), local vs remote backend (S3, GCS, Terraform Cloud), state locking, `terraform state list/show/mv/rm`, `.terraform.lock.hcl`

**L2 — deep & efficient**
- ⬜ **11.4 Modules & structure** — module inputs/outputs, registry modules, versioning modules, `for_each` vs `count`, `dynamic` blocks, `depends_on`, environment layout (dirs vs workspaces), Terragrunt (awareness)
- ⬜ **11.5 Lifecycle & refactoring** — `lifecycle` (`create_before_destroy`, `prevent_destroy`, `ignore_changes`), `import` blocks, `moved` blocks, `-replace`, `refresh-only`, reading a plan diff (`~`, `-/+`), blast radius
- ⬜ **11.6 IaC in teams** — plan in PR, Atlantis, policy as code (OPA/Conftest, Checkov, tfsec/Trivy config scan), `terraform test`, drift detection schedules, secrets out of state, provider version pinning

🛠 **Practice** — L1: with the `kreuzwerker/docker` provider, create the network + Postgres + Go service containers from Terraform; `plan`, `apply`, change a variable, `destroy` · L2: run LocalStack in Docker, use the AWS provider against it to create an S3 bucket + SQS queue via a reusable module, store state in a LocalStack S3 backend, then rename a resource using a `moved` block with zero diff

---

## OPS 12 — Configuration management (Ansible)
**L1 — aware & can use**
- ⬜ **12.1 Ansible basics** — agentless, SSH, inventory (static/dynamic), ad-hoc commands, playbook, play, task, module, `become`, idempotency, `--check` / `--diff`
- ⬜ **12.2 Organizing Ansible** — roles, variables & precedence, `group_vars` / `host_vars`, handlers, templates (Jinja2), facts, Ansible Vault, Ansible Galaxy

**L2 — deep & efficient**
- ⬜ **12.3 Where config management fits today** — immutable images (Packer) vs mutable servers, cloud-init, Ansible + Terraform split, golden AMIs, when k8s replaces config management, Molecule for testing roles

🛠 **Practice** — L1: write a playbook that targets an SSH-enabled Ubuntu container, installs the Go service binary as a systemd unit, and run it twice to show idempotency · L2: refactor into a role with templates + handlers, encrypt a secret with Ansible Vault, and test the role with Molecule's Docker driver

---

## OPS 13 — Cloud fundamentals (AWS first, mapped to GCP)
**L1 — aware & can use**
- ⬜ **13.1 Core concepts** — region, availability zone, edge location, shared responsibility model, Well-Architected Framework pillars, account / project structure, AWS Organizations vs GCP resource hierarchy, service quotas
- ⬜ **13.2 IAM** — users, groups, roles, policies (identity vs resource-based), least privilege, `AssumeRole`, STS temporary credentials, instance profile, GCP service accounts, MFA, root account protection
- ⬜ **13.3 Compute** — EC2 (instance types, AMI), Auto Scaling Group, ECS / Fargate, EKS, Lambda, App Runner; GCP: Compute Engine, GKE, Cloud Run, Cloud Functions; serverless trade-offs (cold start, limits)
- ⬜ **13.4 Networking in the cloud** — VPC, subnets (public/private), route table, Internet Gateway, NAT Gateway, security group vs NACL, VPC endpoints / PrivateLink, peering, Transit Gateway (see NET track)

**L2 — deep & efficient**
- ⬜ **13.5 IAM in depth** — policy evaluation logic (explicit deny), permission boundaries, SCPs, IRSA / EKS Pod Identity, GKE Workload Identity, OIDC federation from CI, cross-account access, IAM Access Analyzer
- ⬜ **13.6 High availability design** — multi-AZ vs multi-region, stateless compute, ASG health checks, zonal failure, cell-based architecture, static stability, blast radius

🛠 **Practice** — L1: make a table mapping 20 AWS services to GCP equivalents and one-line "use when"; draw a VPC with public/private subnets across 2 AZs for the Go service + RDS · L2: write least-privilege IAM policies in Terraform for the Go service (read one S3 bucket, consume one SQS queue, with an explicit deny), apply them to LocalStack, then evaluate 5 example requests on paper (allow/deny and why) using the policy evaluation logic

---

## OPS 14 — Cloud data, messaging & edge services
**L1 — aware & can use**
- ⬜ **14.1 Storage** — S3 (buckets, objects, storage classes, lifecycle rules, versioning, presigned URLs), EBS, EFS; GCP: Cloud Storage, Persistent Disk, Filestore
- ⬜ **14.2 Managed databases & caches** — RDS, Aurora, DynamoDB, ElastiCache (Redis/Valkey), read replicas, automated backups, PITR; GCP: Cloud SQL, AlloyDB, Spanner, Firestore, Memorystore
- ⬜ **14.3 Messaging & integration** — SQS (standard vs FIFO, visibility timeout, DLQ), SNS (fan-out), EventBridge, Kinesis, MSK; GCP: Pub/Sub, Eventarc
- ⬜ **14.4 Edge & traffic** — ALB vs NLB, Route 53 (routing policies, health checks), CloudFront, ACM certificates, WAF, API Gateway; GCP: Cloud Load Balancing, Cloud DNS, Cloud CDN, Cloud Armor

**L2 — deep & efficient**
- ⬜ **14.5 Choosing managed services** — managed vs self-hosted trade-offs, vendor lock-in, egress costs, limits and quotas, S3 consistency model, RDS Proxy, connection limits from Lambda, Secrets Manager vs Parameter Store vs GCP Secret Manager
- ⬜ **14.6 Backups & data protection** — snapshots, cross-region replication, S3 Object Lock, KMS encryption at rest, backup vaults (AWS Backup), restore testing

🛠 **Practice** — L1: in LocalStack, from the Go service upload a file to S3 and return a presigned URL; send/receive from an SQS queue with a DLQ · L2: implement an SQS consumer in Go with visibility timeout extension and DLQ redrive, then write a 1-page "self-host Kafka on k8s vs MSK vs SQS" decision record

---

## OPS 15 — Observability: logs
**L1 — aware & can use**
- ⬜ **15.1 Observability basics** — monitoring vs observability, three pillars (logs, metrics, traces) + profiles, events, telemetry, known unknowns vs unknown unknowns
- ⬜ **15.2 Structured logging** — JSON logs, `log/slog`, log levels, context fields, correlation / request ID, trace ID in logs, stdout logging in containers (12-factor), what never to log (PII, secrets, tokens)

**L2 — deep & efficient**
- ⬜ **15.3 Log pipelines** — log shippers (Fluent Bit, Vector, Promtail / Grafana Alloy), Loki (labels, LogQL, low-cardinality labels), ELK / OpenSearch, retention, sampling & rate-limiting logs, log cost control, audit logs vs app logs

🛠 **Practice** — L1: switch the Go service to `slog` JSON with request ID middleware · L2: add Loki + Grafana + Alloy (or Promtail) to the compose stack, query errors per endpoint with LogQL, and compare a high-cardinality label vs a filtered field

---

## OPS 16 — Observability: metrics & dashboards
**L1 — aware & can use**
- ⬜ **16.1 Prometheus basics** — pull model, scrape targets, `/metrics` endpoint, exposition format, `client_golang`, counter, gauge, histogram, summary, labels, time series
- ⬜ **16.2 PromQL** — instant vs range vectors, `rate()`, `irate()`, `increase()`, `sum by`, `histogram_quantile()`, `topk`, recording rules
- ⬜ **16.3 Grafana & methods** — dashboards, panels, variables, data sources, RED method (rate, errors, duration), USE method (utilization, saturation, errors), four golden signals

**L2 — deep & efficient**
- ⬜ **16.4 Metrics at scale** — cardinality explosion, label design, histogram buckets vs native histograms, exemplars, remote write, long-term storage (Thanos, Mimir, VictoriaMetrics), federation, Prometheus Operator (ServiceMonitor, PodMonitor), kube-state-metrics, node-exporter
- ⬜ **16.5 Go runtime metrics** — goroutines, GC pause, heap in-use, `go_memstats_*`, `runtime/metrics`, process metrics (RSS, open FDs), connection pool metrics (`sql.DBStats`)

🛠 **Practice** — L1: expose RED metrics from the Go service with `client_golang`, scrape with Prometheus in compose, build a Grafana dashboard with p50/p95/p99 latency · L2: deliberately add a `user_id` label, measure series count growth in Prometheus TSDB status, then fix it; add Go runtime + DB pool panels

---

## OPS 17 — Observability: traces, profiles & OpenTelemetry
**L1 — aware & can use**
- ⬜ **17.1 Distributed tracing** — trace, span, parent/child span, span attributes, events, status, trace context propagation (W3C `traceparent`), baggage, Jaeger, Grafana Tempo
- ⬜ **17.2 OpenTelemetry** — OTel API vs SDK, instrumentation libraries (`otelhttp`, `otelgrpc`, `otelsql`/pgx tracing), OTLP, resource attributes, semantic conventions, OTel Collector (receivers, processors, exporters)

**L2 — deep & efficient**
- ⬜ **17.3 Sampling & cost** — head-based vs tail-based sampling, parent-based sampler, sampling ratio, Collector tail sampling processor, span limits, trace/log/metric correlation via exemplars and trace IDs
- ⬜ **17.4 Profiling & eBPF** — continuous profiling (Pyroscope, Parca), pprof in production, eBPF basics, auto-instrumentation with eBPF (Grafana Beyla / OBI), Cilium Hubble (awareness)

🛠 **Practice** — L1: instrument the Go service + Postgres calls with OTel, send via Collector to Jaeger/Tempo, find the slowest span of a request · L2: configure tail sampling in the Collector to keep 100% of errors and 10% of successes; link Grafana from a latency exemplar to the trace and to its logs

---

## OPS 18 — SRE: SLOs & alerting
**L1 — aware & can use**
- ⬜ **18.1 SRE principles** — SRE vs DevOps, toil, error budget policy, embracing risk, "hope is not a strategy", Google SRE book, reliability as a feature
- ⬜ **18.2 SLI / SLO / SLA** — availability SLI, latency SLI, good events / valid events, SLO target (99.9% = ~43 min/month), SLA with penalties, error budget, rolling vs calendar window
- ⬜ **18.3 Alerting basics** — symptom-based vs cause-based alerts, page vs ticket, alert fatigue, runbook link in every alert, Alertmanager (routing, grouping, inhibition, silences), PagerDuty / Opsgenie / Grafana OnCall

**L2 — deep & efficient**
- ⬜ **18.4 SLO engineering** — multi-window multi-burn-rate alerts, burn rate math, Sloth / Pyrra (SLO generators), OpenSLO, choosing SLOs from user journeys, dependency SLOs, SLO review meetings
- ⬜ **18.5 Alert quality** — precision vs recall, flapping, `for:` duration, alert testing (`promtool test rules`), dead man's switch (watchdog alert), alert review cadence

🛠 **Practice** — L1: define 2 SLOs (availability, p99 latency) for the Go service, compute monthly error budget in minutes, write Prometheus alert rules with runbook links · L2: implement multi-window burn-rate alerts (1h/5m and 6h/30m), unit-test them with `promtool test rules`, and inject errors to see them fire in Alertmanager

---

## OPS 19 — SRE: incidents, postmortems & disaster recovery
**L1 — aware & can use**
- ⬜ **19.1 On-call & incident response** — on-call rotation, escalation policy, severity levels (SEV1–SEV4), incident commander, communications lead, war room / incident channel, status page, mitigate first then fix
- ⬜ **19.2 Postmortems** — blameless culture, timeline, impact, root cause vs contributing factors, 5 whys (and its limits), action items with owners, MTTD, MTTA, MTTR, MTBF
- ⬜ **19.3 Disaster recovery basics** — backup vs DR, RPO, RTO, backup/restore, pilot light, warm standby, multi-site active-active, restore drills

**L2 — deep & efficient**
- ⬜ **19.4 Resilience engineering** — chaos engineering (steady-state hypothesis, blast radius), Chaos Mesh, LitmusChaos, AWS FIS, game days, failure injection, load testing before launch, capacity planning, headroom
- ⬜ **19.5 Operational excellence** — production readiness review (PRR) checklist, runbooks & playbooks, change management, freeze windows, toil reduction, automation of remediation, incident metrics trends

🛠 **Practice** — L1: write a full blameless postmortem for a simulated incident (kill Postgres in the compose stack during load) with timeline and action items · L2: run a game day on kind: inject pod kills + network latency with Chaos Mesh while k6 runs, measure SLO impact, and do a timed Postgres backup-restore drill to verify your RTO/RPO

---

## OPS 20 — Container & supply-chain security
**L1 — aware & can use**
- ⬜ **20.1 Image scanning** — CVE, vulnerability databases, Trivy, Grype, Docker Scout, scanning in CI, severity thresholds, base image updates, `govulncheck` for Go code
- ⬜ **20.2 Secrets hygiene in delivery** — no secrets in images or layers, build secrets (`--mount=type=secret`), secret scanning (gitleaks, GitHub push protection), `.env` handling, short-lived credentials via OIDC
- ⬜ **20.3 Runtime basics** — non-root containers, read-only root filesystem, drop all capabilities, `no-new-privileges`, `allowPrivilegeEscalation: false`, seccomp `RuntimeDefault`, resource limits as DoS protection

**L2 — deep & efficient**
- ⬜ **20.4 Supply chain integrity** — SBOM (SPDX, CycloneDX, Syft), image signing (cosign, Sigstore, keyless signing), provenance attestations, SLSA levels, verifying signatures at admission (Kyverno / policy-controller), dependency pinning, Dependabot / Renovate
- ⬜ **20.5 Runtime detection** — Falco, Tetragon, audit logs (k8s audit policy), CIS Benchmarks, kube-bench, network segmentation, minimal-privilege node IAM

🛠 **Practice** — L1: add Trivy image scan + `govulncheck` + gitleaks to the GitHub Actions pipeline, failing on HIGH/CRITICAL · L2: generate an SBOM with Syft, sign the image with cosign (key pair), and enforce "only signed images" in kind with a Kyverno policy; show an unsigned image being rejected

---

## OPS 21 — FinOps & cost awareness
**L1 — aware & can use**
- ⬜ **21.1 Cloud pricing basics** — on-demand, reserved instances, Savings Plans, committed use discounts (GCP), spot / preemptible instances, data transfer (egress) costs, NAT Gateway costs, storage tiers
- ⬜ **21.2 Cost visibility** — cost allocation tags / labels, AWS Cost Explorer, budgets & alerts, showback vs chargeback, unit economics (cost per request / per customer), FinOps Foundation lifecycle (inform, optimize, operate)

**L2 — deep & efficient**
- ⬜ **21.3 Cost optimization** — rightsizing, autoscaling to zero, Graviton / ARM instances, spot for stateless workloads, k8s requests vs real usage (bin packing), OpenCost / Kubecost, idle resources cleanup, log & metric volume costs, Infracost in PRs

🛠 **Practice** — L1: estimate the monthly cost of the Go service + RDS + ALB + NAT Gateway in 2 AZs with the AWS Pricing Calculator and list the top 3 cost drivers · L2: run Infracost on the Terraform code from OPS 11 and write a cost diff comment for a PR that changes instance size

---

## OPS 22 — Applied track: deploying the work project (SAT ERP, Core + 1 module)
> **Time-boxed:** learn this **before** the work project's first production deploy (after its backend has Core + 1 module; target spring 2027). It reuses topics from OPS 1–21, LX, NET and PG — the learning order below shows which ID to study in which week. 22.x subtopics are the project-specific parts.

**Why (EN):** The work project is a Go modular monolith (one binary: `serve`, `worker`, `migrate`, `admin`) on PostgreSQL with River jobs and goose migrations, hosted on VMs in a local (Uzbek) cloud because personal data must stay in the country. The team is small, load is small, so the plan is **boring and safe**: containers on plain VMs with Docker Compose, everything in code, one image promoted dev → staging → prod, and grow (Patroni, more databases, Kubernetes) only when a measured limit forces it. Every tool you add is a tool you must patch, monitor and restore at 2 a.m.
**Зачем (RU):** Рабочий проект — модульный монолит на Go (один бинарник: `serve`, `worker`, `migrate`, `admin`) на PostgreSQL с задачами River и миграциями goose; хостинг — виртуальные машины в местном (узбекском) облаке, потому что персональные данные должны оставаться в стране. Команда маленькая, нагрузка маленькая, поэтому план **скучный и надёжный**: контейнеры на обычных ВМ с Docker Compose, всё в коде, один образ продвигается dev → staging → prod, а усложнение (Patroni, больше баз, Kubernetes) — только когда измеренный предел этого требует. Каждый новый инструмент нужно обновлять, мониторить и восстанавливать в 2 часа ночи.

### Evolution of deployment (what to use at each stage)
| Stage | When | Deploy & infra | Data | Observability |
|---|---|---|---|---|
| **0 — Local** | now, while learning | Docker Compose on the laptop, `Makefile`/Taskfile = same commands locally and in CI | Postgres in a container, named volume | `slog` JSON to stdout |
| **1 — First production** | Core + 1 module live | 3 VMs (app: 2× `serve` behind Caddy · `worker` + Postgres primary · replica in another zone), Docker Compose, Ansible for servers, GitHub Actions CI → GHCR → SSH deploy, SOPS + age for secrets | streaming replica, pgBackRest/WAL-G → S3, PITR, RPO ≤ 15 min, RTO ≤ 4 h | Prometheus, Grafana, Loki, Tempo, Alertmanager → Telegram, Uptime Kuma, GlitchTip/Sentry |
| **2 — Several organizations** | more clients, dedicated databases | blue-green via Caddy, fleet migration runner (all DBs at schema N or N−1), OpenTofu if the provider supports it | Patroni (automatic failover ≤ 5 min), one DB per dedicated client | per-database dashboards and alerts |
| **3 — Scale** | **only** if > 5 tenants or CPU > 60 % sustained | k3s / Kubernetes, Helm, Argo CD (GitOps) — OPS 6–10 | CloudNativePG operator or managed Postgres | Prometheus Operator, OTel Collector |

**L1 — aware & can use**
- ⬜ **22.1 Approach & evolution** — "boring technology", VM + Docker Compose first, stages 0–3, triggers to grow (tenants, CPU, team size), why not Kubernetes on day one, build once → promote one image dev → staging → prod, staging with synthetic data only, data residency (personal data hosted in the country)
- ⬜ **22.2 Target topology** — 3 VMs (app / worker + DB primary / DB replica in another zone), Caddy reverse proxy + automatic TLS, one binary with subcommands (`serve`, `worker`, `migrate`, `admin`), S3-compatible object storage per environment, `deploy/` folder in the monorepo, static frontend hosting vs API hosting
- ⬜ **22.3 Server baseline with Ansible** — non-root deploy user, SSH keys only (no password login), `ufw` firewall, `fail2ban`, `unattended-upgrades`, Docker + Compose plugin, node-exporter, time sync, journald retention, rebuild a server from scratch with one command
- ⬜ **22.4 The project pipeline** — GitHub Actions: `golangci-lint`, `go test -race` with testcontainers (real Postgres), regenerate code (`sqlc`, `oapi-codegen`, `openapi-typescript`) and fail on diff, `govulncheck`, `docker build`, Trivy scan, push to GHCR tagged with git SHA; deploy job: `migrate` → `docker compose pull && up -d` → health check; staging = automatic, prod = manual approval, deploy after operating hours
- ⬜ **22.5 App-side production readiness** — `/healthz` (alive) vs `/readyz` (DB reachable), `/version` with git SHA, graceful shutdown of HTTP **and** River workers, 12-factor config from env, `slog` JSON with trace ID, migrations as a separate step (never on app start), fail fast on bad config
- ⬜ **22.6 Secrets & config per environment** — SOPS + age (encrypted files in Git) vs the provider's secret store, `.env` per environment, never in images or logs, DB roles: owner (migrations) / app (no DDL, RLS-bound) / monitor (read stats), credential rotation

**L2 — deep & efficient**
- ⬜ **22.7 Postgres in production** — streaming replica in another zone, WAL archiving with pgBackRest or WAL-G to S3, PITR, RPO ≤ 15 min, RTO ≤ 4 h (full restore) / ≤ 30 min (failover), daily base backup kept 35 days + monthly 12 months, encrypted copy in a 2nd data centre, quarterly restore drill, Patroni for automatic failover, connection limits and pooling
- ⬜ **22.8 Observability for this project** — RED dashboards per module, alerts on 5xx rate, p95 latency, outbox lag > 60 s, dead-letter jobs > 0, job/timer lag, replica lag, disk; external uptime check; error tracking; alert routing to Telegram + on-call rotation; runbook link in every alert
- ⬜ **22.9 Zero-downtime releases** — 2 `serve` instances + Caddy (rolling or blue-green), expand-and-contract migrations, old and new app versions on one schema (N / N−1), River job payload compatibility during a deploy, rollback = previous image tag + roll forward the DB, release freeze on peak days
- ⬜ **22.10 Many databases (stage 2)** — one DB per dedicated client, routing registry, fleet migration runner (apply to all, report who is behind), per-database backups and monitoring, isolation of failures (one slow DB must not slow others), OpenTofu for VMs/buckets if a provider exists
- ⬜ **22.11 Load tests, drills & runbooks** — k6 at 2× the expected size, game day: kill the primary, fill the disk, deploy a bad migration, restore to "last Tuesday 14:00"; incident runbook, postmortem template, measured RTO/RPO written down

### Learning order (about 16 weeks, before the first deploy)
| Weeks | Study (IDs) | Done when you can… |
|---|---|---|
| 1–2 | LX 3, LX 8, LX 12, NET 8, NET 13 | set up and harden a VPS by hand (users, SSH keys, firewall, systemd service) |
| 3–4 | OPS 1–3, **22.5** | build a < 30 MB non-root Go image; run app + Postgres + migrate in Compose; prove graceful shutdown |
| 5–6 | PG 10.2–10.6, **22.7** (L1 part) | restore to a point in time and promote a replica |
| 7–8 | OPS 4, OPS 20.1–20.2, **22.4** | go from `git push` to a deployed, health-checked staging |
| 9 | OPS 12, **22.3**, **22.6** | rebuild a server from zero with one Ansible command; secrets only in SOPS files |
| 10–11 | OPS 15–17, OPS 18.3, **22.8** | follow one slow request from an alert → trace → SQL query |
| 12 | NET 14, OPS 5.1, **22.1**, **22.2**, **22.9** (Caddy part) | deploy with zero downtime behind Caddy |
| 13 | PG 11.3, OPS 5.4, **22.9** | rename a column with zero downtime in 3 deploys |
| 14 | OPS 20, PG 10.8 | pipeline fails on HIGH/CRITICAL CVEs; app role cannot do DDL |
| 15–16 | OPS 19, **22.11** | run a game day and write the runbook + postmortem |
| After first deploy | **22.10**, OPS 11, OPS 18.4 | stage 2 when a second database is needed |
| Only when stage 3 triggers | OPS 6–10 | Kubernetes when a measured limit forces it |

🛠 **Practice** — `code/devops/22-deployment-track/` — **ship-lab**: a small Go service shaped like the work project (one binary: `serve` / `worker` / `migrate`; chi, pgx, River, goose; `/healthz`, `/readyz`, `/version`). L1: take it through stage 1 on 1–2 cheap VPSs — Ansible baseline, Compose, Caddy + TLS, GitHub Actions → GHCR → SSH deploy with approval for "prod", SOPS secrets, pgBackRest to S3-compatible storage · L2: add a replica + Patroni, Prometheus/Grafana/Loki/Tempo + Telegram alerts, blue-green deploy, a k6 load test, then break it on purpose (kill primary, full disk, bad migration, restore to a timestamp) and record measured RTO/RPO in the README
