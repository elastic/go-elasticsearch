---
navigation_title: Observability
applies_to:
  stack: ga 8.14.0
---

# Observability [observability]

The Go client for {{es}} provides built-in support for distributed tracing via OpenTelemetry and transport-level metrics. These features help you monitor client behavior, diagnose performance issues, and correlate {{es}} requests with the rest of your application.

## OpenTelemetry instrumentation [_otel_instrumentation]

The client includes a built-in OpenTelemetry integration that creates spans for every {{es}} API call. This is the recommended approach for distributed tracing.

### Setup [_otel_setup]

Use `NewOpenTelemetryInstrumentation` to create an instrumentation provider and pass it to the client configuration:

```go
import (
    "github.com/elastic/go-elasticsearch/v9"
    "go.opentelemetry.io/otel"
)

// Use the global TracerProvider (set up elsewhere in your application)
es, err := elasticsearch.New(
    elasticsearch.WithInstrumentation(elasticsearch.NewOpenTelemetryInstrumentation(nil, false)),  // <1>
)
```

1. `nil` uses the global `TracerProvider` from `otel.GetTracerProvider()`. Pass an explicit provider to use a specific one.

### Parameters [_otel_params]

`NewOpenTelemetryInstrumentation` accepts two parameters:

| Parameter           | Type                   | Description                                                                                               |
| ------------------- | ---------------------- | --------------------------------------------------------------------------------------------------------- |
| `provider`          | `trace.TracerProvider` | The OpenTelemetry tracer provider. Pass `nil` to use the globally registered provider.                    |
| `captureSearchBody` | `bool`                 | When `true`, the search query body is captured as the `db.statement` span attribute for search endpoints. |

The search endpoints that support body capture are:

- `search`
- `async_search.submit`
- `msearch`
- `eql.search`
- `terms_enum`
- `search_template`
- `msearch_template`
- `render_search_template`

::::{warning}
Enabling `captureSearchBody` may expose sensitive data in your traces. Only enable it in development or when your trace backend is secured appropriately.
::::

### Span attributes [_otel_span_attributes]

The built-in instrumentation follows the [OpenTelemetry Semantic Conventions for Elasticsearch](https://opentelemetry.io/docs/specs/semconv/database/elasticsearch/). For details on how {{es}} integrates with OpenTelemetry, see [OpenTelemetry integration](docs-content://solutions/observability/apm/opentelemetry/index.md) in the {{es}} documentation. Each API call creates a client span with the following attributes:

| Attribute                       | Description                                                                  | Example                                          |
| ------------------------------- | ---------------------------------------------------------------------------- | ------------------------------------------------ |
| `db.system`                     | Always `"elasticsearch"`                                                     | `elasticsearch`                                  |
| `db.operation`                  | The API endpoint name                                                        | `search`                                         |
| `db.statement`                  | The request body (search endpoints only, when `captureSearchBody` is `true`) | `{"query":{"match_all":{}}}`                     |
| `db.elasticsearch.path_parts.*` | Path parameters for the endpoint                                             | `db.elasticsearch.path_parts.index = "my-index"` |
| `http.request.method`           | The HTTP method                                                              | `GET`                                            |
| `url.full`                      | The full request URL                                                         | `https://localhost:9200/my-index/_search`        |
| `server.address`                | The {{es}} node hostname                                                     | `localhost`                                      |
| `server.port`                   | The {{es}} node port                                                         | `9200`                                           |
| `db.elasticsearch.cluster.name` | The cluster name (from response header `X-Found-Handling-Cluster`)           | `my-cluster`                                     |
| `db.elasticsearch.node.name`    | The node name (from response header `X-Found-Handling-Instance`)             | `instance-001`                                   |

### Full example [_otel_full_example]

```go
import (
    "context"
    "log"

    "github.com/elastic/go-elasticsearch/v9"
    "go.opentelemetry.io/otel"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

func main() {
    // Set up an OTLP exporter
    exporter, err := otlptracegrpc.New(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    // Create and register a TracerProvider
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
    )
    defer tp.Shutdown(context.Background())
    otel.SetTracerProvider(tp)

    // Create the Elasticsearch client with instrumentation
    es, err := elasticsearch.New(
        elasticsearch.WithAddresses("https://localhost:9200"),
        elasticsearch.WithInstrumentation(elasticsearch.NewOpenTelemetryInstrumentation(tp, true)),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer es.Close(context.Background())

    // Every API call now creates an OpenTelemetry span
    res, err := es.Info()
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()
}
```

## Transport metrics [_transport_metrics]

The client can collect transport-level metrics including request counts, failures, and response status code distributions. Enable metrics via `WithTransportOptions`:

```go
es, err := elasticsearch.New(
    elasticsearch.WithTransportOptions(elastictransport.WithMetrics()),
)
```

Once enabled, metrics are available through the transport's `Metrics()` method. You can publish them using Go's `expvar` package for built-in HTTP monitoring:

```go
import "expvar"

es, err := elasticsearch.New(
    elasticsearch.WithTransportOptions(elastictransport.WithMetrics()),
)

// Publish metrics at /debug/vars
expvar.Publish("go-elasticsearch", expvar.Func(func() any {
    m, _ := es.Metrics()
    return m
}))
```

The metrics include:

- **Requests** — Total number of requests sent
- **Failures** — Total number of failed requests
- **Responses** — Count of responses grouped by HTTP status code
- **Connections** — Information about live and dead connection pool members

## Debug logging [_debug_logging]

Enable debug logging to see detailed request and response information:

```go
es, err := elasticsearch.New(
    elasticsearch.WithTransportOptions(elastictransport.WithDebugLogger()),
)
```

For more control over log output, use `WithLogger` with one of the built-in loggers from the `elastictransport` package:

| Logger        | Description                                                |
| ------------- | ---------------------------------------------------------- |
| `TextLogger`  | Plain text output                                          |
| `ColorLogger` | Color-coded output for terminals                           |
| `CurlLogger`  | Outputs requests as `curl` commands (useful for debugging) |
| `JSONLogger`  | Structured JSON output                                     |

```go
import "github.com/elastic/elastic-transport-go/v8/elastictransport"

es, err := elasticsearch.New(
    elasticsearch.WithLogger(&elastictransport.ColorLogger{
        Output:             os.Stdout,
        EnableRequestBody:  true,
        EnableResponseBody: true,
    }),
)
```

:::{dropdown} Alternative integrations
:open: false

### Elastic APM

You can integrate with [Elastic APM](https://www.elastic.co/observability/application-performance-monitoring) by wrapping the HTTP transport with the `apmelasticsearch` package:

```go
import (
    "net/http"

    "github.com/elastic/go-elasticsearch/v9"
    "go.elastic.co/apm/module/apmelasticsearch/v2"
)

es, err := elasticsearch.New(
    elasticsearch.WithTransportOptions(
        elastictransport.WithTransport(apmelasticsearch.WrapRoundTripper(http.DefaultTransport)),
    ),
)
```

This automatically creates APM transactions and spans for each {{es}} request.

### OpenCensus

For OpenCensus-based tracing, use the `ochttp` transport:

```go
import (
    "github.com/elastic/go-elasticsearch/v9"
    "go.opencensus.io/plugin/ochttp"
)

es, err := elasticsearch.New(
    elasticsearch.WithTransportOptions(
        elastictransport.WithTransport(&ochttp.Transport{}),
    ),
)
```

### Custom observability via interceptors

You can also implement fully custom observability using [interceptors](interceptors.md). This gives you complete control over logging, metrics, and tracing at the HTTP level. See the [interceptors guide](interceptors.md) for examples.

:::
