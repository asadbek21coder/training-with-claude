# Rate Limiting

---

## The Problem

A script or attacker sends thousands of requests per second. Server overloads, real users can't use the app, database crashes, cloud bill explodes.

## The Solution

Set maximum requests per time window per user.

---

## Algorithms

### Token Bucket
```
Bucket capacity: 5 tokens
Refill rate: 5 tokens per second
Each request takes 1 token
No tokens → request rejected (429)
```
- ✅ Allows short bursts
- ✅ Simple to implement

### Sliding Window
```
Count requests in a moving time window
Window: last 1 second, limit: 5
Exceeds limit → rejected
```
- ✅ More precise, no burst spikes
- ❌ More memory (stores timestamps)

---

## Where to Put Rate Limiting

```
User → [Rate Limiter / API Gateway] → Load Balancer → Servers
```

Use Redis for shared counter across all servers. Without Redis, each server tracks separately — user can bypass by hitting different servers.

---

## How to Identify Users

- Logged in → rate limit by **user ID** (most accurate)
- Anonymous → rate limit by **IP address**
- API clients → rate limit by **API key**

IP limitation: multiple users behind same network share one IP. Can block entire office. But it's the best option for anonymous users.

In practice, combine multiple signals with different limits for each.

---

## Interview Answer

**"How would you implement rate limiting?"**
> "I'd use a token bucket algorithm with Redis as the shared counter. Redis tracks how many requests each user has made in the current window. When a request comes in, we check the counter — if under the limit, allow and increment. If over, return 429 Too Many Requests. For authenticated users, I'd limit by user ID. For anonymous users, by IP address. The rate limiter sits at the API gateway level, before requests reach the application servers."
