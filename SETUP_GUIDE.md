# Interview Prep with Claude — Setup Guide

This guide explains how to set up Claude as your interview preparation tutor. Copy the instructions below into your first conversation with Claude, and it will follow this system for all future sessions.

---

## Step 1: Create the project folder

```
mkdir -p interview-prep/{code/ds,notes,tasks,english}
cd interview-prep
git init
```

## Step 2: Create SPRINT.md

Create a file called `SPRINT.md` in the root. This tracks your progress. You can copy the topics you want to study, or ask Claude to generate a full roadmap for you. Example:

```markdown
# Interview Sprint Tracker

Legend: ✅ done | 🔄 in progress | ⬜ todo

## Data Structures & Algorithms
- ⬜ Array / Slice
- ⬜ Hash Map
- ⬜ Linked List
- ⬜ Stack & Queue
- ⬜ Binary Tree
- ⬜ Graph
- ⬜ Binary Search
- ⬜ Heap / Priority Queue
- ⬜ Two Pointers
- ⬜ Sliding Window
- ⬜ Dynamic Programming

## System Design
- ⬜ Fundamentals (scaling, load balancing, caching, CDN, CAP, sharding, queues)
- ⬜ Practice designs (URL shortener, rate limiter, etc.)

## Language Internals (Go / your language)
- ⬜ Concurrency
- ⬜ Language internals

## Database (PostgreSQL / your DB)
- ⬜ Indexing
- ⬜ Transactions
- ⬜ Performance
```

## Step 3: Give Claude the rules

In your **first conversation**, send this message (edit the language/tracks if needed):

---

> I want you to help me prepare for software engineering interviews. Follow these rules exactly in every session:
>
> **Learning System — 3 Cycles:**
> - Cycle 1: User-level learning. Simple, easy, no deep complexity. Understand what things are and why they exist.
> - Cycle 2: Middle-level. Deeper problems, trade-offs, edge cases.
> - Cycle 3: Senior-level mastery. Complex problems, optimization, architectural decisions.
>
> **Per-Topic Flow (every topic follows this):**
> 1. Show me the PROBLEM that existed before this concept
> 2. Introduce the concept — how it solves the problem
> 3. Explain all key terms clearly and simply
> 4. Ask me thinking questions one by one — wait for my answer before continuing
> 5. After I answer, correct me and show how a senior developer would answer (in simple English)
> 6. Give me simple tasks (NOT complex LeetCode in Cycle 1)
> 7. I write my approach first, then implement
> 8. End-of-topic quiz before moving to next topic
>
> **Code Rules:**
> - NEVER write implementations for me
> - Only write function signatures / skeletons with TODO comments
> - I implement everything myself
> - If I'm stuck, give hints — don't write the code
>
> **Notes Rules:**
> - Save ALL explanations, discussion conclusions, and senior-level answers to `notes/` folder
> - Save questions and tasks to `tasks/<topic>/` folder
> - Update notes continuously as we discuss
>
> **English Rules:**
> - Correct my English in every message
> - Track all corrections, vocabulary, collocations, and interview phrases in `english/vocabulary.md`
> - Only save the correct versions, not my mistakes
>
> **Senior Answers:**
> - For every concept, show how a mid/senior developer would answer in an interview
> - Use simple words and easy grammar — senior level means WHAT you say, not how fancy the grammar is
>
> **Progress:**
> - Track progress in SPRINT.md
> - Mark topics as done/in-progress as we go
>
> **Parallel Tracks (daily):**
> - Track A: DS & Algorithms (mandatory every day)
> - Track B: System Design (mandatory every day)
> - Other tracks are optional
>
> Start with the first topic in SPRINT.md.

---

## How it works after setup

Once you send that message, Claude will:

1. **Save these rules to its memory** — they persist across all future conversations
2. **Create the folder structure** — notes, tasks, code files
3. **Start teaching** — following the exact flow above

In every future conversation, just say things like:
- "let's continue" — picks up where you left off
- "let's do track B" — switches to system design
- "move to next topic" — finishes current, starts next
- "check" — after you implement something, Claude reviews it

## What you get

```
interview-prep/
├── SPRINT.md                          # your progress tracker
├── code/ds/                           # function signatures (Claude writes)
│   ├── binary_tree.go                 # skeletons
│   └── ds_imp/
│       └── binary_tree_imp.go         # YOUR implementations
├── notes/
│   ├── 01-data-structures/
│   │   ├── 01-array-slice.md          # explanations + senior answers
│   │   ├── 02-hash-map.md
│   │   └── ...
│   └── 02-system-design/
│       ├── 01-scaling-and-load-balancing.md
│       └── ...
├── tasks/
│   ├── binary_tree/
│   │   ├── questions.md               # thinking questions
│   │   └── tasks.md                   # implementation tasks
│   └── ...
└── english/
    └── vocabulary.md                  # verbs, nouns, collocations, interview phrases
```

## Tips

- **Be honest** — if you don't know, say "I don't know." Claude will explain, not judge.
- **Answer in your own words first** — Claude waits for your answer before showing the correct one.
- **Ask "why"** — Claude explains the reasoning, not just the answer.
- **Say what you don't like** — Claude remembers your feedback and adjusts. For example: "don't use complex grammar", "I want simpler tasks", "stop summarizing".
- **One topic at a time** — don't rush. Understanding > speed.
- **Use Claude Code CLI or IDE extension** — it works best when Claude can read/write files directly in your project.
