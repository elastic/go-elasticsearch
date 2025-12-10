# Example: Interceptors

This example demonstrates how to use Interceptors to modify HTTP requests before they are sent to Elasticsearch.

Interceptors wrap the HTTP round-trip and can inspect or modify requests and responses.
They are configured via the `elasticsearch.Config.Interceptors` field:

```go
es, _ := elasticsearch.NewClient(elasticsearch.Config{
    Interceptors: []elastictransport.InterceptorFunc{
        func(next elastictransport.RoundTripFunc) elastictransport.RoundTripFunc {
            return func(req *http.Request) (*http.Response, error) {
                // Modify request before sending
                return next(req)
            }
        },
    },
})
```

## Dynamic Auth Provider

The [`cmd/auth_provider/main.go`](cmd/auth_provider/main.go) example demonstrates how to dynamically inject authentication credentials into requests.

This pattern is useful for scenarios where credentials may change at runtime, such as token refresh or credential rotation.

```bash
go run cmd/auth_provider/main.go
```

## Context-Based Auth Override

The [`cmd/context_auth/main.go`](cmd/context_auth/main.go) example demonstrates how to override authentication credentials on a per-request basis using `context.Context`.

This pattern is useful for multi-tenant applications or impersonation scenarios where different requests need different credentials.

```bash
go run cmd/context_auth/main.go
```

## Custom Auth (Kerberos/SPNEGO)

The [`cmd/custom_auth/main.go`](cmd/custom_auth/main.go) example demonstrates how to implement Kerberos/SPNEGO authentication with challenge-response handling.

The interceptor handles 401 responses with `WWW-Authenticate: Negotiate` by obtaining a token and retrying the request.

> **Note:** This example uses a mock implementation. In production, you would use a Kerberos library like [gokrb5](https://github.com/jcmturner/gokrb5) to obtain service tickets.

```bash
go run cmd/custom_auth/main.go
```

## Custom Observability

The [`cmd/custom_observability/main.go`](cmd/custom_observability/main.go) example demonstrates how to add custom observability to Elasticsearch requests using OpenTelemetry.

It shows three interceptors for:

* **Logging**: Request/response details using `slog`
* **Metrics**: Request counter and duration histogram
* **Tracing**: Distributed tracing with spans

> **Note:** The client has built-in observability functionality. Prefer using the built-in options where possible.

```bash
go run cmd/custom_observability/main.go
```

