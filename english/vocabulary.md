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
