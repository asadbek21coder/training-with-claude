# Technical English — Interview Vocabulary

## Verbs

| Verb | Context | Example |
|---|---|---|
| **strengthen** | upgrading resources | "strengthen the server by adding more RAM" |
| **distribute** (across) | load balancing, fault tolerance | "distribute requests across servers" |
| **direct** (to) | routing | "direct requests to the appropriate server" |
| **push onto** | stack operations | "push a value onto the stack" |
| **pop from** | stack operations | "pop from the stack" |
| **enqueue into** | queue operations | "enqueue into the queue" |
| **dequeue from** | queue operations | "dequeue from the queue" |
| **traverse** | iteration over data structure | "traverse the list / traverse the tree" |
| **backtrack** | DFS behavior | "backtrack to the most recently visited node" |

## Nouns & Terms

| Term | Meaning |
|---|---|
| **device** | hardware unit (server, machine) |
| **load** | system workload / pressure (нагрузка) |
| **FIFO** | First In, First Out |
| **request** | HTTP request, incoming call |
| **single point of failure** | one component whose failure kills the whole system |
| **fault tolerance** | system's ability to continue working when something fails |
| **leaf node** | tree node with no children |
| **cache** (noun/verb) | pronounced "cash" — temporary fast storage |
| **replica** | a copy of the database for reading |
| **database instance** | a running copy of a database |
| **frequently accessed data** | data that gets read often |
| **query the database** | send a request to the database to get data |
| **stale data** | outdated data that no longer reflects the current state |
| **cache invalidation** | removing or updating outdated data in the cache |
| **edge server** | CDN server located close to users geographically |
| **origin server** | the main server where original data lives |
| **cache busting** | changing filename (with hash/version) to bypass cached version |
| **propagate** | spread across all servers (invalidation propagates to all edges) |
| **break the logic** | cause logical errors in the system (not "crash on logic") |
| **verify** | confirm something is correct before proceeding |
| **double-booking** | two users reserving the same resource (seat, room) |
| **data storage** | where/how data is saved (not "data restoring" — that means recovering) |
| **schema** | the predefined structure of a database table |
| **flexible schema** | structure that can vary per record (NoSQL) |
| **read-heavy** | system where reads far outnumber writes |
| **write-heavy** | system where writes are the bottleneck |
| **polyglot persistence** | using different databases for different use cases |
| **ACID** | Atomicity, Consistency, Isolation, Durability — transaction guarantees |
| **sequential** | in order, one after another (not "synchrone") |
| **shard** | a piece of data split across multiple databases |
| **hot spot** | one server getting much more traffic than others |
| **up to a certain scale** | works until traffic/data reaches a limit (not "until some level") |
| **this is where X helps** | introducing a solution (not "there can help X") |
| **acknowledgment** | confirmation that a message was received and processed |
| **idempotent** | safe to run multiple times — same result every time |
| **at-least-once delivery** | message is retried until acknowledged — may deliver twice |
| **fire and forget** | send and don't check if it was processed |
| **undirected** | connection goes both ways (not "undericted") |
| **vertex / vertices** | a node in a graph (plural: vertices, not vertexes) |

## Collocations

| Phrase | When to use |
|---|---|
| **distribute the load across servers** | load balancing |
| **handle requests** | processing incoming traffic |
| **scale up / scale out** | vertical / horizontal scaling |
| **send in sequential order** | one after another, linearly |
| **sticky session** | same user always goes to same server |
| **the fast pointer eventually meets the slow pointer** | Floyd's cycle detection |
| **process nodes in the order they were discovered** | BFS behavior |
| **set its Next pointer to the following node** | linked list pointer assignment |
| **follow a child pointer** | navigating tree nodes |

## Interview Phrases

| Phrase | Context |
|---|---|
| "It starts from the root but **reaches the deepest node first** before backtracking" | explaining DFS |
| "We **recursively** find the depth of the left and right subtrees" | explaining MaxDepth |
| "When we follow a child pointer from a leaf node, the pointer is nil" | explaining base case |
| "Every node in the **entire left subtree** must be less than the ancestor" | explaining BST property |
| "It **distributes incoming traffic across** the available servers" | explaining load balancer |

## Corrected Phrases (from session 2026-10-07)

| Correct version | Note |
|---|---|
| "**Let's** begin" | Let's = let us (apostrophe needed) |
| "I mainly **wrote** code in Go" | past tense of write; "write code **in** a language" |
| "I **participated in** several projects" / "I **took part in** several projects" | participate always takes **in** |
| "I need **a** complete roadmap" | countable noun → article |
| "it should **take into account** that…" | not "account that" |
| "now **it's the** era of AI" / "**we're in the** AI era" | |
| "some skills **don't matter** anymore, and some **are** essential" | plural subject → don't; verb needed in second clause |
| "**alongside** the classic roadmap" | spelling: along**side** |
| "I want to become a **professional, skilled** senior engineer" | article "a" before job title |
| "**T-shaped** engineer" | deep in one area, broad in many |
| "**working proficiency**" | can do it independently, not an expert |
| "You build it, you run it" | ownership of production |
| "remove **unnecessary** files" | spelling: un-ne-**c**-e-**ss**-ary (one c, double s) |
| "a topic, subtopic, **term**" | spelling: t-e-r-m |
| "I'm teaching Go backend to a student **now**" | **now** = at this moment; **know** = to have knowledge |
| "I **give lessons** / I **teach**" | "giving lessons to a student" is OK; "teaching a student" is more natural |
| "check **how much of it** I know" | not "if how much part of it" |
| "these tracks must go on **in parallel**" | "parallelly" is not used in English |
| "at least at **user level**" / "be **familiar with** it, not master it" | describing C1 level |
| "**deprecated**" | old and should no longer be used (software term) |
| "**well-structured**" | hyphen before a noun: "a well-structured repo" |
| "learn them deeper and **know** how to use them efficiently" | know (knowledge) ≠ now (time) |
| "**normally**", "**I've** learnt", "I **don't** know" | spelling / apostrophes |
| "algorithms (**algo**) & data structures" | short form: algo, DS&A |
| "skill, **module**, or **track**" | track = a learning path; module = a part of a course |
| "which **approach** should we choose" | spelling: a-pp-r-o-a-c-h |
| "**best practices**" | spelling: p-r-a-c-t-i-c-e-s; plural |
| "I want to train myself **until then**" | "until it" is not used for time; use "until then" / "before that" |
| "**I don't** get it" | capital I; apostrophe in don't |
| "What do you recommend for **simulating** our work **locally** first?" | recommend + noun / -ing; "locally", not "in local" |
| "**setting up** a server locally" | set up = install and configure |
| "We are **still in the development stage**; we **only use it ourselves for now**" | still / for now; not "yet" at the end for this meaning |
| "I don't **actually know** what lint is" | spelling: actually |
| "I want to use **GitLab's** own runners **until we reach** the free limit" | possessive 's; until + clause |
| "Explain each step **in detail**" | not "detailly" |
| "Normally **I** use `git checkout -b`; **what are you using**?" | capital I; question form |
| "How do I write a `.gitlab-ci.yml` file?" | file name starts with a dot; extension .yml |
| "Explain how it **works** under the hood" | spelling: works |
| "I **fixed it in the editor** and pressed **Commit changes**" | fix something *in* a place; name the button you pressed |
