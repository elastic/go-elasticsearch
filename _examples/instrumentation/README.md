# Example: Instrumentation

This example demonstrates how to instrument the Elasticsearch client.

## expvar

The [**`expvar.go`**](./expvar.go) example publishes client metrics via the [`expvar`](https://pkg.go.dev/expvar) package for consumption by any compatible monitoring system.

## OpenTelemetry (built-in)

The [**`otel.go`**](./otel.go) example uses the client's built-in OpenTelemetry integration via `elasticsearch.NewOpenTelemetryInstrumentation(provider, captureSearchBody)`, passed to the constructor with `elasticsearch.WithInstrumentation(...)`. It produces spans that follow the Elasticsearch OpenTelemetry semantic conventions for every API call, with no per-call wiring required.

```bash
go run otel.go
```

Prefer this built-in hook for standard tracing setups. For a manual, interceptor-based alternative (useful when you need to extend the behaviour, add logging, or combine tracing with custom metrics) see [**`../interceptor/cmd/custom_observability`**](../interceptor/cmd/custom_observability).
