# 12 — Behavioral & Communication

> Behavioral interviews are just as important as technical ones for middle+ roles. Prepare stories, not scripts.

---

## STAR Method

**Key Terms:**
- `STAR` — framework for structured behavioral answers
- `S — Situation` — context; what was the project/team/environment? (2-3 sentences)
- `T — Task` — what was YOUR specific responsibility or challenge?
- `A — Action` — what did YOU specifically do? (most important part; be concrete)
- `R — Result` — measurable outcome; what improved? (numbers, metrics, impact)
- `specificity` — generic answers fail; give real project names, real numbers, real decisions
- `ownership language` — use "I" not "we"; interviewers want to know YOUR contribution
- `quantify results` — "reduced latency by 40%", "saved 3 hours/week", "0 incidents in 6 months"
- `failure acknowledgement` — for failure questions: own the failure, explain what you learned, what changed

**Common mistakes:**
- Too much Situation/Task, too little Action
- Saying "we" instead of "I"
- Vague results ("things got better") instead of specific outcomes
- Reading a script instead of telling a story

**Template:**
```
"At [Company], we were building [X]. My specific responsibility was [Y].
I noticed [problem] and decided to [action 1], [action 2], [action 3].
As a result, [metric] improved by [X%], and [business outcome]."
```

**Must prepare:**
- [ ] 5-7 STAR stories that can answer multiple question types
- [ ] Each story should have: technical complexity + collaboration + impact
- [ ] Practice out loud — interviewers notice rehearsed-sounding vs natural delivery

---

## Technical Communication

**Key Terms:**
- `clarifying questions` — always ask before designing; "what's the expected scale?" "what's the latency requirement?"
- `think out loud` — verbalize reasoning; interviewers value thought process over final answer
- `trade-off articulation` — "Option A is simpler but won't scale beyond X; Option B is complex but handles Y"
- `back-of-envelope estimation` — quick math to validate assumptions (QPS, storage, bandwidth)
- `whiteboard communication` — draw boxes and arrows; label components; explain as you draw
- `confirmation check` — "does this design make sense so far before I continue?"
- `scope acknowledgement` — "given time, I'd also add X, but for now let's focus on Y"
- `assumption declaration` — "I'm assuming single region for now; multi-region would require..."
- `no hedging on basics` — if you know it, say it confidently; hedging on obvious things looks bad
- `graceful unknown handling` — "I haven't worked with X directly, but from what I know, it works like Y. Is that right?"

**System design interview structure:**
1. Clarify requirements (2-3 min)
2. Estimate scale (1-2 min)
3. High-level design (5-7 min)
4. Deep dive on 2-3 components (10-15 min)
5. Discuss trade-offs and bottlenecks (3-5 min)

**Must understand:**
- [ ] Never jump to solution without clarifying requirements
- [ ] How to estimate: DAU × requests_per_user = QPS; data_per_request × QPS = storage/day
- [ ] How to handle a question you don't know perfectly (think out loud; relate to what you do know)
- [ ] How to handle follow-up questions you weren't expecting (pause; think; don't panic)

---

## Conflict Resolution Stories

**Key Terms:**
- `disagreement types` — technical decision, prioritization, process, interpersonal
- `technical disagreement` — propose approach, hear counterargument, find common ground or escalate
- `data-driven resolution` — base decisions on evidence, not opinion ("let's benchmark it")
- `escalation` — when to involve a manager or senior; when it blocks progress and can't be resolved
- `assume good intent` — colleagues are not adversaries; disagreements are about the problem, not the person
- `written communication` — async resolution for non-urgent disagreements; gives people time to think
- `reversible decisions` — don't over-invest in debates about easily reversible decisions
- `two-way door` — Amazon term for reversible decisions; make them fast; don't overthink

**Common questions:**
- "Tell me about a time you disagreed with your manager"
- "Tell me about a conflict with a teammate"
- "Tell me about a time you had to influence without authority"
- "Tell me about a time you had to push back on a requirement"

**Good answer structure:**
1. Set up the disagreement clearly (what was the technical/process difference?)
2. Explain your position and reasoning
3. Describe how you listened to their perspective
4. Describe how you resolved it (compromise, experiment, escalate, defer)
5. Outcome — what happened, what you learned

**Must prepare:**
- [ ] One story: disagreed with senior engineer/manager and you were right
- [ ] One story: disagreed and you deferred to their judgment (shows humility)
- [ ] One story: convinced someone to change their technical approach using data
- [ ] One story: worked through cross-team conflict to deliver shared goal

---

## Failure Stories

**Key Terms:**
- `ownership` — own the failure completely; don't blame others or circumstances
- `root cause analysis (RCA)` — what was the real cause? not the symptom
- `post-mortem` — structured failure analysis; blameless; focuses on system improvement
- `blameless culture` — failures are system problems, not people problems
- `prevention measures` — what changed after the failure? (monitoring, tests, process)
- `growth mindset` — failures are learning opportunities; demonstrate you improved
- `scale appropriately` — choose failure with real impact but not catastrophic; shows you work on meaningful things

**Common questions:**
- "Tell me about a time you failed"
- "Tell me about a production incident you caused"
- "What's your biggest professional mistake?"
- "Tell me about a project that didn't go as planned"

**Red flags interviewers watch for:**
- Blaming others or circumstances
- No learning or behavior change after failure
- Trivial failure (clearly avoiding real answer)
- Not owning it ("the requirement wasn't clear" not "I should have clarified")

**Good answer structure:**
1. What happened — be honest and direct
2. Why it happened — root cause, not just symptoms
3. Impact — what were the consequences? (shows you understand business context)
4. How you fixed it — immediate response
5. What changed — prevention, process, monitoring added
6. What you'd do differently — genuine reflection

**Must prepare:**
- [ ] One real technical failure (production bug, outage, data issue)
- [ ] One project/delivery failure (underestimated scope, missed deadline)
- [ ] Both stories should show clear ownership and genuine learning

---

## Leadership & Ownership Stories

**Key Terms:**
- `ownership` — taking initiative without being asked; caring about outcome
- `bias for action` — moving forward with imperfect information rather than waiting
- `proactive communication` — flagging risks early; not waiting for manager to discover problems
- `cross-functional collaboration` — working across teams (product, design, data) to deliver
- `mentorship` — helping junior team members grow; code reviews, pair programming, knowledge sharing
- `technical leadership` — leading technical decisions on a project without being a manager
- `driving to completion` — following through on initiatives; not just starting them
- `scope creep management` — recognizing when project is growing beyond boundaries; pushing back

**Common questions:**
- "Tell me about a time you took ownership of something outside your role"
- "Tell me about a time you led a project without being the manager"
- "Tell me about a time you mentored someone"
- "Tell me about a time you improved a process"
- "Tell me about a time you went above and beyond"

**Good answer focuses on:**
- What you did that wasn't expected/required
- How you influenced without authority
- Specific impact of your ownership

**Must prepare:**
- [ ] One story: saw a problem, took initiative, drove it to resolution
- [ ] One story: led a technical project end-to-end (design → implementation → monitoring)
- [ ] One story: mentored or unblocked a teammate
- [ ] One story: improved team process or engineering practice

---

## Questions to Ask the Interviewer

**Why it matters:** Good questions show you're evaluating them, not just hoping to get hired. Weak questions ("what's the salary?") signal lack of preparation or low standards.

**Technical questions:**
- "What does the technical stack look like for this team? What are you considering changing and why?"
- "What's the biggest technical challenge the team is facing right now?"
- "How do you handle on-call? What's the PagerDuty rotation like?"
- "What's the test coverage situation? Is testing part of the definition of done?"
- "How do you handle technical debt? Is there dedicated time for it?"

**Team / process questions:**
- "How does the team decide what to work on? How much engineer influence is there in prioritization?"
- "What does the code review process look like? How long does a PR typically sit before review?"
- "How does the team do post-mortems on incidents?"
- "How does the team handle disagreements about technical decisions?"
- "What does a successful first 90 days look like for someone in this role?"

**Growth questions:**
- "How do engineers grow from middle to senior here? What does that path look like?"
- "What's the conference/learning budget? Do engineers have time for learning during work hours?"
- "Can you tell me about someone who grew significantly in their career here?"

**Red flag screening questions:**
- "What's the biggest challenge a new person faces joining the team?"
- "What's something you wish you knew before joining?"
- "Has the team grown or shrunk in the last year?"

**Questions to avoid:**
- "What's the salary?" (ask HR, not the engineer interviewer)
- "How many vacation days?" (HR question)
- "What do you do exactly?" (you should have researched this)
- "Is remote work allowed?" (should know this before applying)

**Must prepare:**
- [ ] 5-7 questions ready; you'll use 3-4 depending on time
- [ ] Questions should vary by interviewer role (engineer vs manager vs recruiter)
- [ ] Listen to answers; ask follow-ups; show genuine interest
