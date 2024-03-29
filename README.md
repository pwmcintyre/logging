---
theme : "night"
highlightTheme: "monokai"
---

# Logging

---

## Where do logs fit?
<!-- .slide: data-background="#468EAE" -->

Pillars of Observability

- Metric (coarse-grain)
- Log (🤷)
- Trace (fine-grain)

<aside class="notes">
Reference:

- https://www.oreilly.com/library/view/distributed-systems-observability/9781492033431/ch04.html
- https://medium.com/@copyconstruct/logs-and-metrics-6d34d3026e38
</aside>

--

### Reality Check: "pillars" are a marketing scam
<!-- .slide: data-background="#468EAE" -->

It's just emitting **data** to a **tool** with a variety of precision

<aside class="notes">
The tool often limits the precision you can include
</aside>

--

### Reality Check: metrics
<!-- .slide: data-background="#468EAE" -->

Metrics are low-fidelity event aggregates

Tells you IF failure; but not WHY failure

--

### Reality Check: tracing
<!-- .slide: data-background="#468EAE" -->

Tracing is logging with opinion + tooling

You can DIY or SaaS:
- AWS X-Ray
- DataDog APM
- honeycomb.io

---

## Basics

1. context
2. correlation
3. level

---

## Context

You cannot predict future questions  
— be generous

<small>Logs impose no limits on size* / cardinality</small>

<aside class="notes">
... consider data sensitivity; probably don't dump blobs
</aside>

--

### Common Mistakes
<!-- .slide: data-background="#A62E2E" -->

--

#### message overloading
<!-- .slide: data-background="#A62E2E" -->

> Task finished: CreateRoutes2: duration=3.014

- hard to parse  
- slow to filter (using `like` operation)  
- ambiguous unit

--

### Context examples

example: `./examples.md`

--

### Context in practice

<span style="color:#46735E">__good logs__</span> are a consequence of <span style="color:#46735E">__good code__</span>

--

#### Example — pipe architecture
<!-- .slide: data-background="#33a" -->

```text
queue service > validator service > sender service
```

Q: who should log?

<small>new problem; how to pass context + correlation?</small>

--

#### SOLID
<!-- .slide: data-background="#33a" -->

```text
controller > queue
controller > validator
controller > sender
controller > log
```

Q: who should log?

A: The controller

<small>everything else performs a discrete function</small>

--

#### Inversion of control
<!-- .slide: data-background="#33a" -->

```go
// get next
work := queue.Pop()
log = logger.WithField("work_id", item.ID)

// validate
if reason := s.validator.IsValid(work.Body); reason != nil {
    logger.WithField("reason", reason).Info("item invalid")
    return
}

// send
if err := s.sender.Send(work.Body); !err != nil {
    logger.WithError(err).Error("failed to send item")
    return
}

// commit work
work.Delete()

// emit
logger.Info("done")
```

<small>* assume error handling!</small>

--

#### Single-responsibility principle
<!-- .slide: data-background="#33a" -->

```go
func Sender (i Item) error {

    // serialize
    body, err := json.Marshal(i)
    if err != nil {
        return errors.Wrap(err, "serialization failure")
    }

    // send
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
    if err != nil {
        return errors.Wrap(err, "http failure")
    }
    defer res.Body.Close()

    // check response
    if response.StatusCode != http.StatusOK {
        return fmt.Errorf("backend failure: %s", response.Status)
    }

    return nil
}
```

<small>dont log AND throw (that's 2 things)</small>

--

### Context in practice
<!-- .slide: data-background="#33a" -->

Code falls into 2 flavours:
1. Controller  
    ```level=INFO```
2. Everything else  
    ```level=DEBUG```

--

### Observer pattern
<!-- .slide: data-background="#33a" -->

example: `./go/service/main.go`

<small>treat logs as an event emitting dependancy</small>

---

## Correlation

... a cross-component concern

👉 ***Concensus*** 👈

☝

```text
correlation_id / request_id / user_id / asset_id / ...
```

--

### Correlation

reaching consensus through tooling

```golang
package appcontext

type RequestContext struct { ... }

type ClientContext struct { ... }

type SystemContext struct {
	Application string `json:"application,omitempty"`
	Version     string `json:"version,omitempty"`
	Environment string `json:"environment,omitempty"`
}

func WithSystemContext(ctx context.Context, val SystemContext) context.Context {
	return context.WithValue(ctx, key, val)
}

func GetSystemContext(ctx context.Context) (val SystemContext, ok bool) {
	val, ok = ctx.Value(key).(SystemContext)
	return
}

...
```

---

## Levels

A broad category which is important to <span style="text-decoration:underline">collectively agree on</span>.

--

### Common Mistakes
<!-- .slide: data-background="#A62E2E" -->

--

#### non-ERROR
<!-- .slide: data-background="#A62E2E" -->

> ERROR: client is not authorized

This belongs in the response to the client:  
`401 Unauthorized`  

(or maybe an "access log")

--

#### non-INFO
<!-- .slide: data-background="#A62E2E" -->

Uninteresting plumbing

> INFO: executed 'SELECT * FROM foo'

> INFO: parsed JSON

<small>aka. i was prototyping and accidentally committed it</small>

--

#### predictions
<!-- .slide: data-background="#A62E2E" -->

Predicting the future

> INFO: about to handle request

<small>aka. i don't trust my language to trap exceptions</small>

--

### Level Definitions

--

### fatal
<!-- .slide: data-background="#46735E" -->

The system cannot continue

> FATAL: failed to connect to database

--

### error
<!-- .slide: data-background="#46735E" -->

A transient problem during processing

> ERROR: timeout while saving

--

### warning
<!-- .slide: data-background="#46735E" -->

Processing degraded but can continue

> WARN: config unset; using default

<small>_opinion: use INFO_</small>

--

### info
<!-- .slide: data-background="#46735E" -->

System did what you asked it to do

> INFO: done

> INFO: batch complete

> INFO: cache refreshed

--

### debug
<!-- .slide: data-background="#46735E" -->

Low-level supporting steps.  

Usually disabled due to poor signal-to-noise ratio.  

__Danger zone:__ Take care with sensitive data!

---

## Closing

you'll get it wrong the first time; **iterate**
