# OPS 4.3 — Writing `.gitlab-ci.yml` layer by layer
Date learned: 2026-11-13 · Layer: L1 · Real file: `sat-erp/.gitlab-ci.yml` (stage 0). Keywords reference: `22-gitlab-cicd.md`

## TL;DR
EN: Build the file in layers: YAML basics → one job with `script` → `image` → `stages` → `variables`/`before_script` → `rules` (per job) → `workflow` (per pipeline) → hidden jobs + `extends` → `cache` vs `artifacts` → `default`. Validate in Build → Pipeline editor before pushing.
RU: Файл строится слоями: основы YAML → один джоб со `script` → `image` → `stages` → `variables`/`before_script` → `rules` (для джоба) → `workflow` (для пайплайна) → скрытые джобы + `extends` → `cache` и `artifacts` → `default`. Проверять в Build → Pipeline editor до отправки.

## Layers
| # | Layer | Key point (EN) | RU |
|---|---|---|---|
| 0 | YAML | spaces only, no Tab; `key: value` with a space; lists `- x` or `[x, y]`; quote strings with `:` `#` | только пробелы; пробел после двоеточия; кавычки для спецсимволов |
| 1 | job + `script` | any non-reserved top-level key is a job; `script` is required; non-zero exit = failed; default stage `test` | ключ верхнего уровня = джоб; `script` обязателен |
| 2 | `image` | job runs in a fresh container of this image, repo already cloned; always pin the version | контейнер из образа; фиксировать версию |
| 3 | `stages` | order between stages, parallel inside a stage, a failed stage stops later ones; every job starts in the repo root | порядок этапов; параллельно внутри этапа |
| 4 | `variables`, `before_script` | global or per job; predefined `$CI_*`; **secrets only in Settings → CI/CD → Variables** | секреты — только в настройках |
| 5 | `rules` | top to bottom, first match wins, no match = no job; inside a rule AND, between rules OR; `if` / `changes` / `exists`; `when: never/manual` | первое совпадение решает |
| 6 | `workflow: rules` | decides if the pipeline exists; avoids double branch + MR pipelines | решает, будет ли пайплайн |
| 7 | `.hidden` + `extends` | template never runs alone; job keys override template keys; see Full configuration | шаблон + наследование |
| 8 | `cache` vs `artifacts` | cache = speed, may be missing, only inside project dir (move `GOPATH`); artifacts = guaranteed results with `expire_in` | кэш для скорости, артефакты для результатов |
| 9 | `default` | applies to every job (`interruptible`, `retry` only for runner failures) | настройки для всех джобов |
| 10 | `image.entrypoint: [""]` | for single-program images (gitleaks) so `script` runs in a shell | отключить entrypoint образа |

## Senior answer
EN: "A GitLab pipeline is defined in `.gitlab-ci.yml` in the repo root. I order cheap checks first with stages, use `workflow` rules so each change gets exactly one pipeline, job `rules` with `changes` so a monorepo only tests what changed, hidden templates with `extends` to avoid repetition, cache keyed on the lock file for speed and artifacts for results. Secrets stay in protected, masked project variables, and I validate the file in the pipeline editor before pushing."
RU: «Пайплайн GitLab описывается в `.gitlab-ci.yml` в корне репозитория. Дешёвые проверки ставлю первыми через stages, `workflow` rules дают ровно один пайплайн на изменение, `rules` с `changes` в монорепо проверяют только изменённую часть, шаблоны с `extends` убирают повторы, кэш с ключом по lock-файлу ускоряет, артефакты хранят результаты. Секреты — в защищённых маскированных переменных проекта, файл проверяю в редакторе пайплайнов до отправки.»

## Gotchas
- Tab in YAML → invalid file. Lost indentation after copy-paste → invalid file (check in Pipeline editor).
- Heredoc in the terminal: content on separate lines, closing `EOF` alone on its own line; `'EOF'` in quotes keeps `$CI_*` unexpanded.
- No `image` → runner's default image (Ruby on GitLab.com), not Go.
- Every job starts in the repo root — `cd backend` in each job (or in the template).
- A cache outside the project dir is silently not saved.
- `CI_DEBUG_TRACE=true` prints everything, including secrets — never on protected branches.
- `git switch -c` = `git checkout -b` (new command since Git 2.23: `switch` for branches, `restore` for files).

## Recall questions
1. What makes a top-level key a job, and what is the only required part?
2. Jobs in one stage vs different stages — how do they run?
3. `rules` vs `workflow: rules`?
4. Inside one rule vs between rules — AND or OR?
5. Why move `GOPATH` into `$CI_PROJECT_DIR`?
6. Where do secrets go, and why never in this file?

## Under the hood (2026-11-13, after the first local run)
EN:
- `gitlab-ci-local` = parse YAML → validate against GitLab's JSON schema → fill defaults (stage `test`, `when: on_success`, `allow_failure: false`) → run. **No `image:` → it runs in the host shell**; with `image:` → in a Docker container.
- Real GitLab: push → server reads the file at that commit, applies `workflow`/`rules`, creates pipeline + jobs (pending); **GitLab runs nothing itself**. The runner **polls** the API and pulls a job (script, image, variables, one-time `CI_JOB_TOKEN`). Docker executor: pull image → **helper container** (git clone, cache, artifacts) → **build container** from `image` runs `script` line by line → log streamed live → exit code decides pass/fail → containers removed.
- The runner only makes outgoing calls, so a runner on a laptop or behind a firewall needs no open ports.
RU: GitLab только создаёт пайплайн и джобы; раннер сам опрашивает API и забирает задание с одноразовым токеном. Docker-исполнитель: helper-контейнер (git, кэш, артефакты) + рабочий контейнер из `image`; лог стримится; код выхода решает статус; контейнеры удаляются. Без `image` gitlab-ci-local запускает джоб прямо в shell хоста.
Gotcha: GitLab looks for the exact name `.gitlab-ci.yml` — `.gitlab-ci.yaml` is ignored. `ls` hides dot-files (`ls -a`). VS Code creates a file only on save.
