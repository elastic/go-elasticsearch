---
navigation_title: Configuration
---

# Configuration reference [configuration]

The `elasticsearch.Config` struct controls all aspects of the Go client's behavior. The same configuration is shared by both `NewClient` (low-level API) and `NewTypedClient` (typed API) — both use the same underlying transport.

```go
es, err := elasticsearch.NewClient(elasticsearch.Config{
    Addresses: []string{"https://localhost:9200"},
    Username:  "elastic",
    Password:  "changeme",
})
```

For details on connecting to Elasticsearch, see [Connecting](connecting.md).

## Connection [_config_connection]

| Field       | Type       | Default                     | Description                                                                                                                                                    |
| ----------- | ---------- | --------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Addresses` | `[]string` | `["http://localhost:9200"]` | A list of Elasticsearch node URLs. When empty, the client checks the `ELASTICSEARCH_URL` environment variable (comma-separated).                               |
| `CloudID`   | `string`   | `""`                        | Endpoint for the Elastic Cloud service. When set, it overrides `Addresses`. You can find this value in your [Elastic Cloud console](https://cloud.elastic.co). |

:::{dropdown} Example: Connecting to Elastic Cloud

```go
es, err := elasticsearch.NewClient(elasticsearch.Config{
    CloudID: "my-deployment:dXMtY2VudHJhbC0xLmd...",
    APIKey:  "base64_api_key_here",
})
```

:::

## Authentication [_config_auth]

| Field          | Type     | Default | Description                                                                                           |
| -------------- | -------- | ------- | ----------------------------------------------------------------------------------------------------- |
| `Username`     | `string` | `""`    | Username for HTTP Basic Authentication.                                                               |
| `Password`     | `string` | `""`    | Password for HTTP Basic Authentication.                                                               |
| `APIKey`       | `string` | `""`    | Base64-encoded API key for authorization. If set, overrides `Username`/`Password` and `ServiceToken`. |
| `ServiceToken` | `string` | `""`    | Service token for authorization. If set, overrides `Username`/`Password`.                             |

::::{note}
For dynamic credential rotation or per-request authentication, use [interceptors](advanced/interceptors.md).
::::

## TLS and security [_config_tls]

For more information on securing your {{es}} cluster, see [Security](docs-content://deploy-manage/security.md) in the {{es}} documentation.

| Field                    | Type     | Default | Description                                                                                                                                                                                 |
| ------------------------ | -------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `CACert`                 | `[]byte` | `nil`   | PEM-encoded CA certificate. When set, an empty certificate pool is created and the certificate is appended. Only valid when the `Transport` field is not specified or is `*http.Transport`. |
| `CertificateFingerprint` | `string` | `""`    | SHA-256 hex fingerprint of the CA certificate. Useful when you don't have access to the CA file directly.                                                                                   |

:::{dropdown} Example: TLS with CA certificate

```go
cert, _ := os.ReadFile("/path/to/http_ca.crt")

es, err := elasticsearch.NewClient(elasticsearch.Config{
    Addresses: []string{"https://localhost:9200"},
    Username:  "elastic",
    Password:  "changeme",
    CACert:    cert,
})
```

:::

## Retry [_config_retry]

The client automatically retries failed requests for certain HTTP status codes. You can customize the retry behavior with the following fields.

| Field           | Type                              | Default           | Description                                                                                                         |
| --------------- | --------------------------------- | ----------------- | ------------------------------------------------------------------------------------------------------------------- |
| `RetryOnStatus` | `[]int`                           | `[502, 503, 504]` | List of HTTP status codes that trigger a retry.                                                                     |
| `DisableRetry`  | `bool`                            | `false`           | Set to `true` to disable all retries.                                                                               |
| `MaxRetries`    | `int`                             | `3`               | Maximum number of retry attempts per request.                                                                       |
| `RetryOnError`  | `func(*http.Request, error) bool` | `nil`             | Optional function to decide whether a network error should be retried. Return `true` to retry.                      |
| `RetryBackoff`  | `func(attempt int) time.Duration` | `nil`             | Optional function that returns the backoff duration for each retry attempt. When `nil`, retries happen immediately. |

:::{dropdown} Example: Custom retry with exponential backoff

```go
es, err := elasticsearch.NewClient(elasticsearch.Config{
    RetryOnStatus: []int{429, 502, 503, 504},
    MaxRetries:    5,
    RetryBackoff: func(attempt int) time.Duration {
        // Exponential backoff: 100ms, 200ms, 400ms, 800ms, 1600ms
        return time.Duration(1<<uint(attempt)) * 100 * time.Millisecond
    },
})
```

:::

## Compression [_config_compression]

Request body compression reduces bandwidth usage at the cost of CPU time.

| Field                      | Type   | Default                   | Description                                                                               |
| -------------------------- | ------ | ------------------------- | ----------------------------------------------------------------------------------------- |
| `CompressRequestBody`      | `bool` | `false`                   | Enable gzip compression for request bodies.                                               |
| `CompressRequestBodyLevel` | `int`  | `gzip.DefaultCompression` | Gzip compression level (`1`=best speed, `9`=best compression, `-1`=default).              |
| `PoolCompressor`           | `bool` | `false`                   | Use a `sync.Pool`-based gzip writer for reduced allocations in high-throughput scenarios. |

:::{dropdown} Example: Enable compression

```go
es, err := elasticsearch.NewClient(elasticsearch.Config{
    CompressRequestBody: true,
    PoolCompressor:      true,
})
```

:::

## Node discovery [_config_discovery]

Node discovery (sniffing) allows the client to automatically discover and use all nodes in the cluster.

| Field                   | Type            | Default        | Description                                                          |
| ----------------------- | --------------- | -------------- | -------------------------------------------------------------------- |
| `DiscoverNodesOnStart`  | `bool`          | `false`        | Perform node discovery when initializing the client.                 |
| `DiscoverNodesInterval` | `time.Duration` | `0` (disabled) | Periodically discover nodes at this interval. Set to `0` to disable. |

:::{dropdown} Example: Enable node discovery

```go
es, err := elasticsearch.NewClient(elasticsearch.Config{
    Addresses:             []string{"https://localhost:9200"},
    DiscoverNodesOnStart:  true,
    DiscoverNodesInterval: 5 * time.Minute,
})
```

:::

## Transport [_config_transport]

These fields allow advanced customization of the HTTP transport layer.

| Field                | Type                                                                                              | Default                 | Description                                                                                |
| -------------------- | ------------------------------------------------------------------------------------------------- | ----------------------- | ------------------------------------------------------------------------------------------ |
| `Transport`          | `http.RoundTripper`                                                                               | `http.DefaultTransport` | The HTTP transport object. Replace to use a custom HTTP client implementation.             |
| `Header`             | `http.Header`                                                                                     | `nil`                   | Global HTTP headers included in every request.                                             |
| `Selector`           | `elastictransport.Selector`                                                                       | Round-robin             | Node selection strategy. Implement `elastictransport.Selector` for custom selection logic. |
| `ConnectionPoolFunc` | `func([]*elastictransport.Connection, elastictransport.Selector) elastictransport.ConnectionPool` | `nil`                   | Constructor function for a custom connection pool.                                         |

:::{dropdown} Example: Custom HTTP transport with timeouts

```go
es, err := elasticsearch.NewClient(elasticsearch.Config{
    Transport: &http.Transport{
        MaxIdleConnsPerHost:   10,
        ResponseHeaderTimeout: 10 * time.Second,
        TLSClientConfig: &tls.Config{
            MinVersion: tls.VersionTLS12,
        },
    },
})
```

:::

## Observability [_config_observability]

| Field               | Type                               | Default | Description                                                                                                                                 |
| ------------------- | ---------------------------------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| `Instrumentation`   | `elastictransport.Instrumentation` | `nil`   | Instrumentation provider for distributed tracing. Use `elasticsearch.NewOpenTelemetryInstrumentation()` for built-in OpenTelemetry support. |
| `EnableMetrics`     | `bool`                             | `false` | Enable transport-level metrics collection (request counts, failures, response status codes).                                                |
| `EnableDebugLogger` | `bool`                             | `false` | Enable debug-level logging of requests and responses.                                                                                       |
| `Logger`            | `elastictransport.Logger`          | `nil`   | Custom logger for request/response logging. Built-in loggers: `TextLogger`, `ColorLogger`, `CurlLogger`, `JSONLogger`.                      |

For details, see [Observability](advanced/observability.md).

:::{dropdown} Example: OpenTelemetry instrumentation

```go
es, err := elasticsearch.NewClient(elasticsearch.Config{
    Instrumentation: elasticsearch.NewOpenTelemetryInstrumentation(nil, true),
})
```

:::

## Middleware [_config_middleware]

```{applies_to}
stack: ga 8.19.1
```

| Field          | Type                                 | Default | Description                                                                                                                                                                        |
| -------------- | ------------------------------------ | ------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Interceptors` | `[]elastictransport.InterceptorFunc` | `nil`   | Array of middleware functions that can modify requests and responses. Applied left-to-right for requests, right-to-left for responses. Set at client creation and immutable after. |

For details and examples, see [Interceptors](advanced/interceptors.md).

## Compatibility [_config_compatibility]

| Field                     | Type   | Default | Description                                                                                                                                     |
| ------------------------- | ------ | ------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `EnableCompatibilityMode` | `bool` | `false` | Send compatibility headers for 7.x client to 8.x server migration. Can also be set via the `ELASTIC_CLIENT_APIVERSIONING` environment variable. |
| `DisableMetaHeader`       | `bool` | `false` | Disable the `X-Elastic-Client-Meta` HTTP header that sends client metadata to Elasticsearch.                                                    |

## Environment variables [_config_env_vars]

The client reads the following environment variables:

| Variable                       | Description                                                                                          |
| ------------------------------ | ---------------------------------------------------------------------------------------------------- |
| `ELASTICSEARCH_URL`            | Comma-separated list of Elasticsearch node URLs. Used when `Addresses` and `CloudID` are both empty. |
| `ELASTIC_CLIENT_APIVERSIONING` | When set to a truthy value, enables compatibility mode (same as `EnableCompatibilityMode`).          |
