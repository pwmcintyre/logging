---
theme : "night"
highlightTheme: "monokai"
---

# Logging

---

## What is a log?

> stream of aggregated, time-ordered events ... one event per line ...

— https://12factor.net/logs

---

## Where do logs fit?

Pillars of Observability

- Log 👈
- Metric
- Trace

<aside class="notes">

Logs are the easiest way to level-up your observability. Why?
1. ease of emission (tooling)
2. both metrics and traces can be built *from* logs

Reference:
- https://www.oreilly.com/library/view/distributed-systems-observability/9781492033431/ch04.html
- https://medium.com/@copyconstruct/logs-and-metrics-6d34d3026e38
</aside>

---

## Levels

A broad category which is important to <span style="text-decoration:underline">collectively agree on</span>.

---

### Common Mistakes
<!-- .slide: data-background="#A62E2E" -->

---

#### non-ERROR
<!-- .slide: data-background="#A62E2E" -->

> ERROR: failed to authenticate

401 — a client error!

This belongs in the response to the client; not in logs.

---

#### non-INFO
<!-- .slide: data-background="#A62E2E" -->

Uninteresting plumbing

> INFO: executed 'SELECT * FROM foo'

> INFO: parsed JSON

---

#### predictions
<!-- .slide: data-background="#A62E2E" -->

Predicting the future

> INFO: about to handle request

---

### Level Definitions

---

### fatal
<!-- .slide: data-background="#46735E" -->

The system cannot continue

> FATAL: failed to connect to database

---

### error
<!-- .slide: data-background="#46735E" -->

A transient problem during processing

> ERROR: timeout while saving

---

### warning
<!-- .slide: data-background="#46735E" -->

Processing degraded but can continue

> WARN: config unset; using default

---

### info
<!-- .slide: data-background="#46735E" -->

System did what you asked it to do

> INFO: batch complete

> INFO: cache refreshed

---

### debug
<!-- .slide: data-background="#46735E" -->

Low-level supporting steps.  

Usually disabled due to poor signal-to-noise ratio.  

__Danger zone:__ Take care with sensitive data!

---

## Structure

---

### Examples

---

#### Message (verbs)

Should uniquely identity an activity.  
Don't be tempted to overload this with context.

Eg. "completed request" or "failed to complete request"

---

### Context (nouns)

Logs allow high-cardinality values.

You cannot predict future-questions — be generous.

---

#### Sensitive Context

You need to decide what to log — consider:

"Business" data vs "Technical" data

eg. "Customer Name" vs "Customer ID"

Typically:
- database contain business data
- logs contain reference to business data; and how it got there

---

## In practice (opinion)

![classic setup](./img/plumbing.png)

Code falls into 2 flavours:
1. Business Logic
2. Plumbing

As such: only your service layer should log as "INFO"

---

### Observer pattern

---

## TL;DR;

- DONT : write stories over many lines
- DO : emit discreet, context-rich events
