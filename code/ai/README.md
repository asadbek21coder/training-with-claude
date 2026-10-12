# AI Engineering Practice

Roadmap: `roadmap/ai.md` · Notes: `notes/ai/`

## What practice looks like
- **AI 1–5 (using AI):** no special project. Practice on real work: your course (Bozorcha lessons), Go practice in `code/golang/`, your job. Keep a short log of what AI got wrong and how you caught it.
- **AI 6–12 (building AI):** one Go service grows across topics: LLM API CLI → RAG over this repo's `notes/` (Postgres + pgvector in Docker) → MCP server → evals → production hardening → security tests.

## Folder convention
```
code/ai/
  06-llm-api/
  07-rag/
  08-mcp-server/
  ...
```
Every topic folder gets a `README.md` with goal, steps, done-when and an L2 stretch. API keys go in environment variables, never in the repo.
