# Message Queues

---

## The Problem

User places an order. Server must: save order, send email, update inventory, notify warehouse, update analytics. User waits for ALL of them — 580ms.

User only cares about: "is my order saved?" The rest can happen later.

## The Solution: Async Processing

```
Synchronous (before):
  User → save order → send email → update stock → notify → analytics → respond
  User waits: 580ms

Asynchronous (after):
  User → save order → respond (50ms)
         → push events to message broker
         → workers process in background
  User waits: 50ms
```

---

## Two Approaches

### Simple: PostgreSQL as Queue
```
Table: task_queue (type, status, created_at)
Cronjob polls every few seconds → processes pending tasks

✅ No extra infrastructure
❌ Polling wastes resources
❌ Doesn't scale to thousands per second
❌ Not real-time
```

Works fine for small apps.

### Scalable: Message Broker (RabbitMQ, Kafka)
```
App pushes events → broker holds them → workers pull and process

✅ Real-time processing
✅ Scales to millions of events per second
✅ Workers are independent — one down, events wait in queue
❌ Extra infrastructure
```

**Key difference:**
- PostgreSQL queue: app PULLS tasks (polling)
- Message broker: broker PUSHES events to workers (instant)

---

## Delivery Guarantees

### At-Most-Once — fire and forget
- Send event, don't check if processed
- Events can be lost
- Use for: analytics, logging

### At-Least-Once — retry until acknowledged
- Worker processes → sends acknowledgment
- No acknowledgment → retry
- Same event might be processed TWICE
- Use for: emails, notifications, most tasks

### Exactly-Once — each event processed once
- Very hard to achieve
- Needs special coordination
- Use for: payments, bank transfers

---

## The Duplicate Problem

At-least-once can deliver the same event twice:
```
Worker sends email → crashes before acknowledging → broker retries → another worker sends again
User gets 2 emails!
```

**Solution: Idempotency**
Each event has a unique ID. Before processing, check: "did I already process this ID?" If yes, skip.

---

## Interview Answer

**"How would you handle async processing?"**
> "I'd use a message broker like RabbitMQ or Kafka. When a user places an order, the server saves the order and responds immediately. Then it pushes events to the broker — send email, update stock, notify warehouse. Separate workers pick up these events and process them independently. I'd use at-least-once delivery with idempotent consumers. Each event has a unique ID, and workers check if they already processed it before acting. This prevents duplicates while making sure no events are lost."
