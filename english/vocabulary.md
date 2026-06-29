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
