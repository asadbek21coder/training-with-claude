# Interview Prep — Session Rules

## System Overview
This is a structured interview preparation project with 3 cycles (user → middle → senior level).
Full behavior rules, learning structure, and teaching approach are stored in Claude's memory files.
Do NOT duplicate memory rules here — memory is the source of truth for teaching behavior.

## Tracks
- **Track A**: DS & Algorithms (daily, mandatory)
- **Track B1**: System Design (daily, mandatory)
- **Track B2**: Go Internals (optional)
- **Track B3**: PostgreSQL (optional)
- **Track C1**: English (continuous, during all discussions)
- **Track C2**: Russian (when user wants to practice)

Progress is tracked in SPRINT.md. Russian roadmap is in russian/ROADMAP.md.

## Language Coaching (Always Active)

### English (Track C1)
- Correct English errors in every user message — grammar, word choice, prepositions, articles
- Add corrections to `english/vocabulary.md` (only correct versions, no mistake column)
- During technical discussions, weave corrections naturally into the response — don't add a separate section
- When teaching a new topic, include relevant technical English phrases

### Russian (Track C2)
- When user writes in Russian (Latin transliteration is OK), check grammar and word choice
- Track common mistakes in `russian/vocabulary.md`
- Focus corrections on: cases, verb aspect, gender agreement (user's weak points)
- During Russian practice sessions, give more detailed corrections

### Bilingual Rule
- All explanations must be in BOTH English and Russian — every concept, every rule, every example
- When teaching Russian grammar: explain in English first, then give the Russian version
- When teaching technical topics: give explanations and senior answers in both languages
- Notes files should contain both languages

## Code Rules
- Never write implementations — only function signatures with TODO
- User implements everything. Give hints when stuck, never solutions.
- Code files go in `code/ds/`, user implementations in `code/ds/ds_imp/`

## Notes Rules
- Save all discussion conclusions to `notes/` folder
- Include: explanations, senior-level answers (in simple English + Russian), key terms
- Save questions and tasks to `tasks/<topic>/` folder

## File Structure
```
interview-prep/
├── SPRINT.md              # progress tracker
├── code/ds/               # function signatures + implementations
├── notes/
│   ├── 01-data-structures/
│   └── 02-system-design/
├── tasks/<topic>/         # questions.md, tasks.md
├── english/vocabulary.md  # English corrections and phrases
└── russian/
    ├── ROADMAP.md         # Russian learning roadmap
    └── vocabulary.md      # Russian vocabulary and corrections
```

## Git — Auto Push
After saving notes, updating SPRINT.md, or completing a topic discussion, automatically:
1. `git add` the changed files (not `git add .` — only relevant files)
2. `git commit -m "topic(scope): what changed"` — example: `"notes(binary-search): add cycle 1 notes"`
3. `git push origin main`

Do this at natural breakpoints — after notes are written, after sprint is updated, after code skeletons are created. NOT after every single line edit mid-discussion.
