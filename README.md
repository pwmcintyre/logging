---
theme : "night"
highlightTheme: "monokai"
---

# Logging

TL;DR; nobody "reads" a log; instead we use tools to parse and filter; consider changing your strategy:

- DO : emit discreet, context-rich events
- DO NOT : write stories over many lines

---

## WHAT is a log?

> the stream of aggregated, time-ordered events ... one event per line ...

— https://12factor.net/logs

---

### Levels

- __fatal__  
The system cannot continue and cannot proceed.  
 _"failed to connect to database."_

- __error__  
A transient problem during processing.  
 _"dependency responds with HTTP 5xx."_

- __warning__  
Processing degraded but can continue.  
 _"Failed to get config; using defaults."_

- __info__  
System did what you asked it to do!  
 _"Processed a request; Refreshed a cache; Completed the daily batch."_

- __debug__  
Low-level supporting steps.  
Disabled by default due to poor signal-to-noise ratio.  
Danger zone: Take care with sensitive data!

--

- __fatal__ — the system cannot continue
- __error__ — an isolated problem
- __warning__ — processing degraded
- __info__ — a core function happened 👈 **_events!_**
- __debug__ — for SME to troubleshoot process flow

---

## WHERE do logs fit?

Logs are the easiest way to level-up your observability. Why?
1. ease of emission (tooling)
2. both metrics and traces can be built *from* logs

---

### Pillars of Observability

Whitebox:
- 👉 Log 👈
- Metric
- Trace

Blackbox:
- poll (eg. ping)

### Refs:
- https://www.oreilly.com/library/view/distributed-systems-observability/9781492033431/ch04.html
- https://medium.com/@copyconstruct/logs-and-metrics-6d34d3026e38

---

### Structure

- time
- message
- system context
  - application version
  - cache state
- request context
  - request ID
  - user ID

#### Message (verbs)

Should uniquely identity an activity.  
Don't be tempted to overload this with context.

Eg. "completed request" or "failed to complete request"

### Context (nouns)

Logs allow high-cardinality values.

You cannot predict future-questions — be generous.


#### Sensitive Context

You need to decide what to log — consider:

"Business" data vs "Technical" data

eg. "Customer Name" vs "Customer ID"

Typically:
- database contain business data
- logs contain reference to business data; and how it got there
