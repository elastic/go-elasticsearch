# Example: Instrumentation

This example demonstrates how to instrument the Elasticsearch client.

## expvar

The [**`expvar.go`**](./expvar.go) example publishes client metrics via the [`expvar`](https://pkg.go.dev/expvar) package for consumption by any compatible monitoring system.

## OpenTelemetry

For OpenTelemetry-based instrumentation (logs, metrics, and tracing via interceptors) see [**`../interceptor/cmd/custom_observability`**](../interceptor/cmd/custom_observability), which demonstrates the recommended interceptor pattern for adding observability to requests.
