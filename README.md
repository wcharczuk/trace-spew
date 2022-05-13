trace-spew
==========

`trace-spew` is a local server that will emit any _distributed tracing spans_ sent to it.

To use `trace-spew`:
- Start the server with `BIND_ADDR=127.0.0.1:8193 go run main.go`
- Point your Datadog(tm) trace address at `127.0.0.1:8193` wherever that's configured