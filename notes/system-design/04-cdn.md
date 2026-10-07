# CDN — Content Delivery Network

---

## The Problem

Origin server is in one location. Users are worldwide. Static files (images, JS, CSS) travel across the world on every request — slow.

```
User (Tokyo) → Origin Server (New York): ~200ms round trip
Every image, every CSS file, every JS file = slow page load
```

Same static files requested by thousands of users — why send from origin every time?

---

## What is a CDN?

A geographically distributed network of **edge servers** that cache static content closer to users.

```
Without CDN:
  User (Tokyo) ──────────────────→ Origin (New York)   200ms

With CDN:
  User (Tokyo) ──→ CDN Edge (Tokyo)                    20ms
                   ↓ cache miss only
                   Origin (New York) → cache at edge
```

---

## What CDN Serves vs Does NOT Serve

### ✅ Serves (static, same for everyone)
- Images, videos, audio
- CSS, JavaScript, HTML files
- Fonts
- File downloads (PDFs, installers)

### ❌ Does NOT serve (dynamic, personalized)
- API responses (different per user)
- Database queries
- Authentication, business logic
- Real-time data (chat, stock prices)

**The rule:** Same for every user → CDN can cache it. Personalized or frequently changing → origin server.

---

## Cache Invalidation on CDN

### Problem
Deploy new `app.js` but old version is cached on 50 edge servers worldwide.

### Solution 1: TTL + Invalidation
- Set TTL on CDN content
- On deploy, purge/invalidate the CDN cache
- Next request fetches fresh from origin
- Problem: propagation across global edges is slow (minutes)

### Solution 2: Cache Busting (most common)
Append content hash or version to filename:

```
Old: app.js         → cached everywhere
New: app.def456.js  → new URL, no cache exists, CDN fetches fresh

index.html (short TTL, always fresh from origin):
  <script src="/app.def456.js">
```

- No invalidation needed — new URL = cache miss = fresh fetch
- Old file expires naturally via TTL
- HTML has short TTL so it always points to latest filenames

---

## Interview Answer

**"What is a CDN and when would you use it?"**
> "A CDN is a geographically distributed network of edge servers that cache static content — images, JavaScript, CSS, videos — closer to end users. Instead of every request traveling to the origin server, the nearest edge server serves the cached copy, reducing latency from hundreds of milliseconds to tens. I'd use a CDN for any static assets that are the same for all users. For cache invalidation, the standard approach is cache busting with content-hashed filenames rather than explicit invalidation, since purging across global edge servers can be slow."

