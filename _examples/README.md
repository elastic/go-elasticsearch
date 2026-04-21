# Examples

This directory contains examples and recipes for the Elasticsearch Go client.

## Configuration & Customization

The [**`configuration.go`**](./configuration.go) and [**`customization.go`**](./customization.go) files contain information for configuring and customizing the client transport.

## Migrating to the Typed API

The [**`totyped.go`**](./totyped.go) file demonstrates incremental migration from the functional (low-level) API to the typed API using `(*Client).ToTyped()`. The returned `*TypedClient` shares the source client's transport and configuration, so existing call sites keep working while new code can use the typed builder API.

The [**`typed_endpoint.go`**](./typed_endpoint.go) file shows the lightest-weight migration step: using a single typed endpoint (for example `typedapi/core/search`) directly against the existing functional `*elasticsearch.Client`, without constructing a full `*TypedClient` or calling `ToTyped()`. Every typed endpoint package exports a `New(elastictransport.Interface)` constructor, and the functional client satisfies that interface via its embedded `BaseClient.Perform`. The same file also shows `elasticsearch.NewBase(...)`, which returns a raw `*BaseClient` for fresh code that only ever needs a handful of typed endpoints, skipping both the `esapi` tree and the typedapi `MethodAPI` tree.

## Logging

The [**`logging`**](./logging) directory contains examples for using the default loggers and implementing a custom logger.

## Bulk

The [**`bulk`**](./bulk) directory contains a comprehensive example of using the _Bulk_ API.

## Search

The [**`search`**](./search) directory contains examples for using the _Search_ API
with the typed client, showing both the `esdsl` query builders and raw `types.*` structs.

## Multi Search Template

The [**`msearchtemplate`**](./msearchtemplate) directory contains examples for using the _MSearchTemplate_ API
with both `esapi` and `typedapi`.

## Encoding

The [**`encoding`**](./encoding) directory contains examples of using third-party packages
and client helper methods for a more efficient encoding and decoding of the request and response bodies.

## Fast HTTP

The [**`fasthttp`**](./fasthttp) directory contains a demonstration of replacing the default client transport with an HTTP client from the [`github.com/valyala/fasthttp`](https://pkg.go.dev/github.com/valyala/fasthttp) package.

## Instrumentation

The [**`instrumentation`**](./instrumentation) directory contains recipes for instrumenting the client: exposing client metrics via `expvar`, and enabling the built-in OpenTelemetry integration via `elasticsearch.NewOpenTelemetryInstrumentation`. For an interceptor-based OpenTelemetry alternative see [**`interceptor/cmd/custom_observability`**](./interceptor/cmd/custom_observability).

## Interceptors

The [**`interceptor`**](./interceptor) directory contains examples of using Interceptors to modify HTTP requests before they are sent to Elasticsearch, including dynamic auth, per-request auth override, custom auth (Kerberos/SPNEGO), and custom observability.

## Extension

The [**`extension`**](./extension) directory contains an example of extending the client APIs, for example
to support custom endpoints installed by a plugin.

## Google Cloud Functions

The [**`cloudfunction`**](./cloudfunction) directory contains a simple web service for checking Elasticsearch cluster status.

## XKCD Search

The [**`xkcdsearch`**](./xkcdsearch) directory contains a command-line utility for indexing and searching an archive of [xkcd.com](https://xkcd.com) artworks.
