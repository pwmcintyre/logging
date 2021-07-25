# Example

```
2021-07-25T04:12:50Z authorized
```

1. No context
2. No correlation (time is not robust!)
3. No structure

## Structure

```json
{
    "time": "2021-07-25T04:12:50Z",
    "msg": "authorized"
}
```

Structure means we can parse it with a broad range of tools.

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
