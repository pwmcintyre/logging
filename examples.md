---
theme : "night"
highlightTheme: "monokai"
---

# Example

---

## Classic

```text
2021-07-25T04:12:50Z INFO handling request
2021-07-25T04:12:50Z INFO handling request
2021-07-25T04:12:50Z INFO authorizing "123"
2021-07-25T04:12:50Z INFO authorizing "543"
2021-07-25T04:12:51Z INFO authorized
2021-07-25T04:12:52Z INFO unauthorized
```

---

## Structure

```json
{
    "time": "2021-07-25T04:12:50Z",
    "msg": "authorized"
}
```

Structure means we can parse it with a broad range of tools.

---

## User context

```json
{
    "time": "2021-07-25T04:12:50Z",
    "msg": "authorized",
    "id": "123"
}
```

Context means why can ask more interesting questions.

_"how frequently does this ID authenticate with success?"_

---

## User state

```json
{
    "time": "2021-07-25T04:12:50Z",
    "msg": "authorized",
    "id": "123",
    "groups": ["a", "b"]
}
```

Ephemeral state is particularly hard to troubleshoot.

---

## System context

```json
{
    "time": "2021-07-25T04:12:50Z",
    "application": "authorizer@3.0.1",
    "msg": "authorized",
    "id": "123",
    "groups": ["a", "b"]
}
```

System context means we know what emitting this; and which version.

_"when did this start failing? is this a bug with v3?"_

---

## System state

```json
{
    "time": "2021-07-25T04:12:50Z",
    "application": "authorizer@3.0.1",
    "msg": "authorized",
    "id": "123",
    "groups": ["a", "b"],
    "cache_used": "1627186370"
}
```

System state is often ephemeral; in-memory only; perilous.

---

## System tracing

```json
{
    "time": "2021-07-25T04:12:50Z",
    "application": "authorizer@3.0.1",
    "msg": "authorized",
    "id": "123",
    "groups": ["a", "b"],
    "cache_used": "1627186370",
    "request_id": "a39b28c9"
}
```

A complex system may emit many events; a request ID brings lineage.

---

## Distributed tracing

```json
{
    "time": "2021-07-25T04:12:50Z",
    "application": "authorizer@3.0.1",
    "msg": "authorized",
    "id": "123",
    "groups": ["a", "b"],
    "cache_used": "1627186370",
    "request_id": "a39b28c9",
    "corelation_id": "d4289bd7"
}
```

A distributed system add more complexity; consider a specialized ID to across system boundaries.
