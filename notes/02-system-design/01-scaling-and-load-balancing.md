# Scaling & Load Balancing

---

## The core problem

One server has a limit. When too many users hit it simultaneously, CPU/RAM/network gets
exhausted — requests queue up, response times climb, server crashes.

```
User → Server → Database
```

This works for 100 users. Breaks at 10,000.

---

## Vertical scaling — make the server bigger

Replace the server with a more powerful machine.

```
BEFORE:          AFTER:
┌──────────┐     ┌──────────────────┐
│ 4 cores  │  →  │ 32 cores         │
│ 8GB RAM  │     │ 256GB RAM        │
└──────────┘     └──────────────────┘
```

**Pros:** simple, no code changes, no architecture changes
**Cons:** expensive, hard ceiling (biggest machine available), single point of failure

---

## Horizontal scaling — add more servers

Run many smaller servers, split traffic between them with a load balancer.

```
              ┌──────────────┐
User ────────►│ Load         │
              │ Balancer     │
              └──────┬───────┘
                     │
          ┌──────────┼──────────┐
          ▼          ▼          ▼
      ┌────────┐ ┌────────┐ ┌────────┐
      │Server 1│ │Server 2│ │Server 3│
      └────────┘ └────────┘ └────────┘
```

**Pros:** scales infinitely by adding servers, no single point of failure
**Cons:** more complex, need a load balancer, shared state becomes a problem

---

## Load balancer strategies

**Round robin** — rotate through servers in order:
```
Request 1 → Server 1
Request 2 → Server 2
Request 3 → Server 3
Request 4 → Server 1  ← wraps around
```
Simple. Works well when all requests have similar cost.

**Least connections** — send to server with fewest active requests:
```
Server 1: 10 active  Server 2: 3 active ← send here  Server 3: 7 active
```
Better when requests have varying cost (some fast, some slow).

**IP hash** — hash user's IP, always route to the same server:
```
hash(user_ip) % 3 = 1 → always Server 2 for this user
```
Used for sticky sessions. Same user always hits same server.

---

## The session problem in horizontal scaling

**What is a session?**

HTTP is stateless — the server has no memory between requests. Sessions add memory:

```
1. User sends username + password → server verifies → creates session:
   sessions = { "abc123": { user_id: 42, role: "admin" } }

2. Server sends session ID back → stored in browser cookie

3. Every future request → browser sends Cookie: session_id=abc123

4. Server looks up "abc123" → knows who you are → no password needed again
```

**The problem with horizontal scaling:**

Each server only has access to its own in-memory sessions. A user who logs in on
Server 1 is a stranger to Server 2.

```
1. User logs in → hits Server 1 → session saved in Server 1 memory
2. Next request → load balancer sends to Server 2
3. Server 2 checks its own memory → nothing found → user appears logged out
```

User gets randomly logged out every time load balancer picks a different server.

---

## Three solutions to the session problem

**Fix 1 — Sticky sessions (IP hash)**

Always route the same user to the same server. Session stays on one server.

```
user_123 always → Server 1
user_456 always → Server 2
```

Problem: if Server 1 goes down, all its users lose sessions. Uneven load if some
users are much more active than others.

**Fix 2 — Centralized session store (Redis)**

Move sessions out of server memory into a shared Redis instance all servers can reach.

```
          ┌────────┐ ┌────────┐ ┌────────┐
          │Server 1│ │Server 2│ │Server 3│
          └────┬───┘ └────┬───┘ └────┬───┘
               └──────────┼──────────┘
                          ▼
                     ┌─────────┐
                     │  Redis  │  ← shared session store
                     └─────────┘
```

Any server handles any request. All check the same Redis. Most common solution.

**Fix 3 — Stateless auth (JWT)**

Encode user identity inside a signed token. No server-side storage at all.

```
1. User logs in → server creates JWT:
   { user_id: 42, role: "admin", expires: "2026-07-01" } + cryptographic signature

2. Client stores JWT, sends it with every request

3. Any server verifies signature → knows who the user is → no Redis lookup needed
```

Problem — logout is broken:
```
1. Attacker steals your JWT
2. You log out → your browser deletes the token
3. Attacker still has the copy
4. Attacker sends requests with stolen token → server sees valid signature → lets them in
5. Cannot be stopped until token naturally expires (could be hours)
```

Server has no record of which tokens exist → cannot revoke before expiry.

Workarounds: short expiry (15 min) + refresh tokens, or token blacklist in Redis
(but that reintroduces server-side storage).

---

## Comparison

| | Sticky Sessions | Redis Sessions | JWT |
|--|----------------|----------------|-----|
| Scales horizontally? | ⚠️ partial | ✅ yes | ✅ yes |
| Server-side storage? | ✅ local memory | ⚠️ Redis needed | ❌ none |
| Works if server dies? | ❌ session lost | ✅ yes | ✅ yes |
| Can revoke immediately? | ✅ yes | ✅ yes | ❌ no |
| Extra infrastructure? | ❌ none | ⚠️ Redis | ❌ none |

---

## Interview answers

**"How does horizontal scaling break sessions?"**
> "Each server only has access to its own in-memory state. A user who authenticates
> on Server 1 appears as a stranger to Server 2 because sessions are local to each
> process. The standard fix is a centralized session store like Redis — all servers
> read and write to the same place, so any server can handle any request. Alternatively,
> stateless JWT authentication removes server-side session storage entirely by encoding
> user identity in a cryptographically signed token the client carries on every request."

**"What is the problem with JWT logout?"**
> "JWT tokens are self-contained — the server stores nothing. When a user logs out,
> the token is deleted from the client but the server has no revocation list to check
> against. If the token was stolen before logout, the attacker can continue using it
> until it naturally expires. The common mitigation is short expiry times combined
> with refresh tokens, but there is no perfect solution without reintroducing some
> form of server-side state."
