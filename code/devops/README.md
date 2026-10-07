# DevOps Practice

Roadmap: `roadmap/devops.md` · Notes: `notes/devops/` · Dictionary: `reference/devops.md`

## What practice looks like
- One small Go HTTP service (`/healthz`, `/hello`, a Postgres-backed endpoint) is the "patient" for every topic.
- You add files around it: `Dockerfile`, `compose.yaml`, `.github/workflows/*.yml`, Kubernetes manifests, Helm charts / Kustomize overlays, Terraform, Ansible playbooks, Prometheus/Grafana/OTel configs, alert rules, postmortems.
- Everything runs locally: Docker, kind, LocalStack. No paid cloud account needed.

## Environment setup
```bash
# Docker + Compose plugin (check)
docker version && docker compose version

# kind (local Kubernetes) + kubectl + helm
go install sigs.k8s.io/kind@latest
kind create cluster --name lab
kubectl cluster-info --context kind-lab
# helm: https://helm.sh/docs/intro/install/

# Postgres for the service
docker run -d --name pg -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres:17

# LocalStack (fake AWS) for Terraform / cloud exercises
docker run -d --name localstack -p 4566:4566 localstack/localstack
# Terraform or OpenTofu: https://opentofu.org/docs/intro/install/
```

## Folder convention
```
code/devops/
├── README.md
└── <NN>-<topic-slug>/        # e.g. 02-dockerfile-images, 06-k8s-architecture
    ├── README.md             # exercise instructions (from the 🛠 Practice line of OPS <NN>)
    └── ...                   # Dockerfile, compose.yaml, manifests, *.tf, results/notes
```
- `<NN>` = topic number in the roadmap (OPS 2 → `02-...`). Create the folder when the topic starts.
- Record measurements (image sizes, build times, pipeline durations) in the exercise README.
- Clean up after exercises: `kind delete cluster --name lab`, `docker compose down -v`.
