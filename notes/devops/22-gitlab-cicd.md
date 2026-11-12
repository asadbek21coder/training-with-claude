# OPS 22.12 + OPS 4.3 — GitLab platform & GitLab CI/CD (work project setup)
Date learned: 2026-11-12 · Layer: L1 (L2 parts marked) · Hands-on: the work repo `gitlab.com/sattech-uz/erp/sat-erp`

## TL;DR (3 lines max)
EN: GitLab = Git hosting + merge requests + CI/CD + container registry in one product. Security comes from **structure** (group → subgroup → project, roles, 2FA), **protection** (protected branches/tags, protected+masked variables) and a **pipeline** (`.gitlab-ci.yml`) that runs on **our own runner** and builds one image that is promoted staging → production.
RU: GitLab = Git-хостинг + merge request'ы + CI/CD + реестр контейнеров в одном продукте. Безопасность даёт **структура** (группа → подгруппа → проект, роли, 2FA), **защита** (защищённые ветки/теги, защищённые и маскированные переменные) и **пайплайн** (`.gitlab-ci.yml`), который работает на **нашем собственном раннере** и собирает один образ, который продвигается staging → production.

## Problem → Solution
**Before:** code on someone's personal account; anyone can push to `main`; deploy = "SSH to the server and `git pull`"; passwords in chat; nobody knows which version runs in production; if the one person who knows leaves, the project stops.
**After:** the company owns the group; `main` changes only through a reviewed merge request with a green pipeline; every commit builds an image tagged with its SHA; staging deploys automatically; production deploys only from a protected `v*` tag, by a Maintainer, with the same image; secrets live in protected variables; everything is written down and repeatable.

## Key terms
| Term | Meaning (simple EN) | RU |
|---|---|---|
| Group / subgroup | folder for projects + members + shared settings; settings and members are inherited downwards | группа / подгруппа |
| Namespace | the path part of a URL: `sattech-uz/erp` | пространство имён |
| Project | one Git repository + its issues, MRs, CI, registry | проект |
| Role | Guest < Planner < Reporter < Developer < Maintainer < Owner | роль |
| 2FA | second factor (authenticator app code) after the password | двухфакторная аутентификация |
| Recovery codes | one-time codes to log in if the phone is lost | коды восстановления |
| Protected branch | branch where only chosen roles may push / merge / force-push | защищённая ветка |
| Protected tag | tag pattern (`v*`) only chosen roles may create | защищённый тег |
| Merge request (MR) | request to merge a branch into `main` (GitHub: pull request) | запрос на слияние |
| Squash | many branch commits become one commit on `main` | сжатие коммитов |
| Fast-forward merge | `main` just moves forward; no merge commit; linear history | слияние перемоткой |
| Pipeline | one run of CI for a commit / MR / tag | пайплайн, конвейер |
| Stage | group of jobs that run in order (`lint` → `test` → `build`) | этап |
| Job | one task with a `script` (runs in a container) | задание |
| Runner | the agent (program) that takes jobs and runs them | раннер, исполнитель |
| Executor | how the runner runs a job: `docker`, `shell`, `kubernetes`… | исполнитель (тип) |
| Artifact | files a job saves for later jobs / download | артефакт |
| Cache | files kept between pipelines to go faster (Go modules) | кэш |
| CI/CD variable | key=value given to jobs; can be **protected** and **masked** | переменная CI/CD |
| Environment | named target (`staging`, `production`) with deploy history | окружение |
| Container Registry | Docker image storage inside the project: `registry.gitlab.com/...` | реестр контейнеров |
| Deploy token | limited token for a server to `docker pull` (read_registry) | токен развёртывания |

## How it works / how to use

### 1. Structure and ownership (done 2026-11-12)
```
sattech-uz/                 ← top-level group "SAT tech" — owned by the company (Owners: boss + lead)
├── erp/                    ← subgroup per product (members added per product)
│   └── sat-erp             ← monorepo: backend/ frontend/ contracts/ deploy/
└── <other products>/
```
- Settings set on the top group flow down: 2FA required (48 h grace), no sharing with outside groups, Pages public access removed.
- **Roles — give the lowest role that works:** Owner (boss, lead) · Maintainer (DevOps + lead: merge to `main`, create `v*` tags, settings) · Developer (engineers: push branches, open MRs) · Reporter (read only).
- **Free tier on GitLab.com (VERIFY on about.gitlab.com/pricing):** about 5 users per private top-level group (all subgroups counted together), about 400 compute minutes/month on GitLab's shared runners, about 10 GiB storage. Paid-only (Premium): **required** MR approvals, approval rules, code owners as gate, protected environments / deployment approvals, merge trains. → We use our own runner (no minute limit) and protected tags instead of deployment approvals.

### 2. Protection (done 2026-11-12)
- **Protected branch `main`:** allowed to merge = Maintainers, allowed to push = No one, force push = off. Proof: a direct push gave `pre-receive hook declined`.
- **Protected tag `v*`:** allowed to create = Maintainers → only they can start a production release.
- **MR settings:** fast-forward merge, squash = require, all threads resolved, delete source branch. **"Pipelines must succeed"** — turn on after the first pipeline exists.
- **Daily flow:** `git switch -c feat/x` → commit → `git push -u origin feat/x` → open MR → pipeline green + review → Maintainer merges → `main` → staging deploy.

### 3. Runners
| Type | Who can use it | Our use |
|---|---|---|
| Instance (shared) runners | all projects on gitlab.com; counts compute minutes | **turned off** for our project |
| Group runner | all projects in the group | later, when several projects need it |
| Project runner | one project | **now** — on the laptop, later on a VM in the Uzbek cloud |

- **Docker executor:** each job starts a fresh container from `image:`; clean and repeatable.
- **Shell executor:** runs commands directly on the host; simple but dirty and less safe.
- Runner tags (`docker`) + "run untagged jobs" decide which jobs a runner picks.
- (L2) Building Docker images inside CI needs Docker-in-Docker (`services: [docker:dind]`, runner `privileged = true`) or a daemonless builder (Buildah, BuildKit rootless). Privileged = the job is almost root on the runner host → keep that runner only for trusted projects.

### 4. `.gitlab-ci.yml` — the keywords you need
```yaml
stages: [lint, test, build, deploy]          # order of stages

default:
  image: golang:1.26                          # container for jobs (pick your Go version)
  interruptible: true                         # new push cancels old pipeline

workflow:                                     # when a pipeline exists at all
  rules:
    - if: $CI_PIPELINE_SOURCE == "merge_request_event"
    - if: $CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH
    - if: $CI_COMMIT_TAG

variables:
  GOPATH: $CI_PROJECT_DIR/.go

.go-cache: &go-cache                          # reusable piece (YAML anchor)
  cache:
    key: { files: [backend/go.sum] }          # new cache only when go.sum changes
    paths: [.go/pkg/mod]

lint:backend:
  stage: lint
  <<: *go-cache
  script: [cd backend, go vet ./...]
  rules:
    - changes: [backend/**/*]                 # monorepo: run only if backend changed

test:backend:
  stage: test
  <<: *go-cache
  services: [postgres:18]                     # (or testcontainers + dind)
  script: [cd backend, go test -race ./...]
  artifacts:
    reports: { junit: backend/report.xml }    # test results shown in the MR
    expire_in: 1 week

build:image:
  stage: build
  image: docker:27
  services: [docker:27-dind]
  script:
    - docker login -u "$CI_REGISTRY_USER" -p "$CI_REGISTRY_PASSWORD" "$CI_REGISTRY"
    - docker build -t "$CI_REGISTRY_IMAGE:$CI_COMMIT_SHORT_SHA" backend
    - docker push "$CI_REGISTRY_IMAGE:$CI_COMMIT_SHORT_SHA"
  rules:
    - if: $CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH

deploy:staging:
  stage: deploy
  environment: { name: staging, url: https://staging.example }
  resource_group: staging                     # never two deploys at the same time
  script: [./deploy/rollout.sh staging "$CI_COMMIT_SHORT_SHA"]
  rules:
    - if: $CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH

deploy:production:
  stage: deploy
  environment: { name: production }
  resource_group: production
  when: manual                                # a human presses ▶
  script: [./deploy/rollout.sh production "$CI_COMMIT_SHORT_SHA"]
  rules:
    - if: $CI_COMMIT_TAG =~ /^v/              # only protected v* tags
```
*(This is a learning sketch, not the final pipeline. The real one follows ADR-0024 §3.4 in the design repo.)*

| Keyword | What it does |
|---|---|
| `stages` | order; jobs in the same stage run in parallel |
| `image` / `services` | job container / helper containers (Postgres, dind) |
| `script` / `before_script` | commands; a non-zero exit code fails the job |
| `rules` (`if`, `changes`, `when`) | when the job runs; replaces old `only/except` |
| `workflow:rules` | when the whole pipeline is created (avoid duplicate branch + MR pipelines) |
| `needs` | start a job as soon as its dependency finishes (DAG), not the whole stage |
| `artifacts` vs `cache` | artifacts = results passed on / downloadable; cache = speed only, may be missing |
| `environment` | records deployments; shows what version runs where |
| `resource_group` | one deploy per environment at a time |
| `when: manual` | job waits for a click |
| `extends`, `!reference`, YAML anchors, `include` | reuse config; split big files |
| `parallel: matrix` | run the same job with several variable sets |
| `interruptible`, `retry`, `timeout`, `allow_failure` | pipeline hygiene |

**Predefined variables used most:** `CI_COMMIT_SHA` / `CI_COMMIT_SHORT_SHA`, `CI_COMMIT_BRANCH`, `CI_COMMIT_TAG`, `CI_DEFAULT_BRANCH`, `CI_PIPELINE_SOURCE`, `CI_REGISTRY`, `CI_REGISTRY_IMAGE`, `CI_REGISTRY_USER`, `CI_REGISTRY_PASSWORD`, `CI_JOB_TOKEN`, `CI_PROJECT_DIR`.

### 5. Secrets and variables
- **Protected** = only on protected branches/tags (`main`, `v*`) → a feature-branch MR cannot read production secrets.
- **Masked** = hidden in job logs (value must meet masking rules). **Masked and hidden** = cannot be read back in the UI.
- **Environment scope** = `production` values only for jobs with `environment: production`.
- Never `echo $SECRET`, never `set -x` around secrets, never commit `.env`.
- Server-side config: SOPS + age encrypted files in `deploy/` (OPS 7.6, 22.6).
- Servers pull images with a **deploy token** (`read_registry` only), not a personal password.

### 6. Release flow we will use
```
feature branch → MR (lint, test, contract, security) → merge to main
   → build image :<sha> → deploy:staging (automatic) → smoke test
   → Maintainer creates tag v1.2.0 on that commit → deploy:production (manual, same image)
   → rollback = deploy the previous tag's image again (DB migrations are forward-only)
```

## Senior answer (interview-ready)
EN: "I set up GitLab so that the company owns the code and the pipeline is the only road to production. A company group with two owners and enforced 2FA; developers push branches and open merge requests; `main` is protected and only Maintainers merge after a green pipeline. Every commit builds one immutable image tagged with its SHA in the GitLab registry; staging deploys automatically; production deploys the same image only from a protected version tag, as a manual job with a resource group so two deploys never overlap. Jobs run on our own runner inside the country, secrets are protected and masked variables scoped to the environment, and rollback means redeploying the previous image because migrations are expand-only."
RU: «Я настраиваю GitLab так, чтобы кодом владела компания, а пайплайн был единственной дорогой в продакшен. Группа компании с двумя владельцами и обязательной 2FA; разработчики пушат ветки и открывают merge request'ы; `main` защищена, сливают только Maintainer'ы после зелёного пайплайна. Каждый коммит собирает один неизменяемый образ с тегом SHA в реестре GitLab; staging деплоится автоматически; в продакшен уходит тот же образ только по защищённому тегу версии, ручным джобом с resource group, чтобы два деплоя не пересекались. Задания выполняются на нашем раннере внутри страны, секреты — защищённые и маскированные переменные с областью окружения, а откат — это повторный деплой предыдущего образа, потому что миграции только расширяющие.»

## Gotchas
- `touch dir` creates a **file**; `mkdir` creates a directory; Git does not track empty directories (put a `README.md` or `.gitkeep` inside).
- `echo ... > README.md` overwrites the file in the **current** directory; the path after `>` decides where it goes. `git restore <file>` brings back the committed version.
- A protected branch also blocks Owners from pushing directly if "Allowed to push" = No one — that is the point.
- "Pipelines must succeed" with no CI file can block merging — enable it after the first pipeline.
- Without `workflow:rules` a push to a branch with an open MR can create **two** pipelines (branch + MR).
- Shared runners use compute minutes; a forgotten instance runner toggle can eat the free quota.
- Masked variables are hidden in logs, but a job can still send them somewhere — protection (who can run the job) matters more than masking.
- `latest` image tags break rollback; always deploy by SHA or digest.
- Free tier: approvals are optional, not required — the "second person reviews" rule is a team rule until Premium.

## Recall questions (answer without looking — that's how memory sticks)
1. Why does the company, not a person, own the top-level group? What breaks if a person owns it?
2. What is the difference between Maintainer and Developer in our setup?
3. What exactly did the `pre-receive hook declined` message prove?
4. Instance vs group vs project runner — which one do we use and why?
5. `artifacts` vs `cache` — which one may be missing, and which one is for passing build results?
6. What do "protected" and "masked" do for a variable? Which one stops a feature branch from reading a production password?
7. Why does production deploy only from a `v*` tag, and who can create it?
8. What does `resource_group` prevent?
9. How do we roll back production, and why does it not touch the database?
10. Which features would push us to GitLab Premium?

## Links
Practice: `code/devops/22-deployment-track/` (Lab 3) · Roadmap: OPS 4.3, OPS 22.4, OPS 22.12 · Work design docs: ADR-0024 (design repo) · Docs: docs.gitlab.com — CI/CD YAML syntax reference, Runners, Protected branches, Protected tags, CI/CD variables, Container Registry, Environments
