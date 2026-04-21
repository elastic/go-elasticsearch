# Example: Instrumentation

This example demonstrates how to instrument the Elasticsearch client.

## OpenCensus

The [**`opencensus.go`**](./opencensus.go) example uses the [`ochttp.Transport`](https://pkg.go.dev/go.opencensus.io/plugin/ochttp#example-Transport) wrapper to auto-instrument the client calls, and provides a simple exporter which prints information to the terminal.

<a href="https://asciinema.org/a/KhyP3GuuHPJAZQAmrgmdwS7uf" target="_blank"><img src="https://asciinema.org/a/KhyP3GuuHPJAZQAmrgmdwS7uf.svg" width="750" alt="OpenCensus instrumentation demo" /></a>

## expvar

The [**`expvar.go`**](./expvar.go) example publishes client metrics via the [`expvar`](https://pkg.go.dev/expvar) package for consumption by any compatible monitoring system.

## OpenTelemetry

For OpenTelemetry-based instrumentation (logs, metrics, and tracing via interceptors) see [**`../interceptor/cmd/custom_observability`**](../interceptor/cmd/custom_observability), which demonstrates the recommended interceptor pattern for adding observability to requests.
