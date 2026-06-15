---
navigation_title: "Low-level API"
---

# Low-level API [low_level_api]

The low-level API provides a one-to-one mapping with the {{es}} REST API. Each endpoint accepts raw `io.Reader` request bodies and returns `*esapi.Response` objects whose `Body` is an `io.ReadCloser` you read and close yourself.

The rest of this documentation is built around the [typed API](../typed-api/index.md) and the [esdsl builders](../typed-api/esdsl.md), which together cover the vast majority of {{es}} endpoints with compile-time safety, automatic JSON encoding and decoding, and fluent query construction. **We recommend the typed API for new code.** If you are already using the low-level API, the [migration guide](migration.md) shows how to adopt the typed API without rewriting your whole codebase.

## When to use the low-level API [_when_low_level]

Reach for the low-level API when:

- You need an endpoint that is not covered by the typed API.
- You want full control over request serialization, for example to plug in a custom JSON encoder or a streaming body.
- You are working with pre-baked JSON payloads and do not want to model them as Go structs.

If none of these apply, prefer the [typed API](../typed-api/index.md).

## Creating a client [_low_level_client]

Use `elasticsearch.New` with functional options:

```go subs=true
package main

import (
    "context"
    "log"

    "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}"
)

func main() {
    client, err := elasticsearch.New(
        elasticsearch.WithAddresses("https://localhost:9200"),
        elasticsearch.WithAPIKey("API_KEY"),
    )
    if err != nil {
        log.Fatalf("creating client: %s", err)
    }
    defer client.Close(context.Background())

    res, err := client.Info()
    if err != nil {
        log.Fatalf("info: %s", err)
    }
    defer res.Body.Close()
    log.Println(res)
}
```

For full configuration options, see the [Configuration reference](../configuration.md).

::::{note}
`NewClient(Config)` and `NewDefaultClient()` still work but are deprecated. Use `New` with [functional options](../configuration.md) instead.
::::

## Response handling [_low_level_response_handling]

Every low-level call returns an `*esapi.Response` whose body you must read and close. Failing to close the body prevents Go's HTTP client from reusing the underlying TCP connection.

```go
res, err := client.Search(
    client.Search.WithIndex("my-index"),
    client.Search.WithBody(strings.NewReader(`{"query":{"match_all":{}}}`)),
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close() // <1>

if res.IsError() { // <2>
    log.Fatalf("search failed: %s", res.String())
}

var result map[string]any
if err := json.NewDecoder(res.Body).Decode(&result); err != nil { // <3>
    log.Fatal(err)
}
```

1. Always close the response body, even on the success path.
2. `IsError()` returns `true` for HTTP status codes >= 400.
3. Decode the body yourself into whatever shape fits your application.

## Common operations [_low_level_operations]

A condensed tour of the most common CRUD and search calls. For detailed tutorials, see the [typed API equivalents](../using-the-api/index.md); the low-level signatures below are stable and map directly to the corresponding REST endpoints.

### Create an index

```go
mapping := `{"mappings":{"properties":{"price":{"type":"integer"}}}}`

res, err := client.Indices.Create(
    "test-index",
    client.Indices.Create.WithBody(strings.NewReader(mapping)),
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

### Index a document

```go
document := struct {
    Name  string `json:"name"`
    Price int    `json:"price"`
}{Name: "Foo", Price: 10}

data, _ := json.Marshal(document)
res, err := client.Index(
    "test-index",
    bytes.NewReader(data),
    client.Index.WithDocumentID("1"),
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

### Get a document

```go
res, err := client.Get("test-index", "1")
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

### Search

```go
query := `{"query":{"match":{"name":{"query":"Foo"}}}}`

res, err := client.Search(
    client.Search.WithIndex("test-index"),
    client.Search.WithBody(strings.NewReader(query)),
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

### Delete a document

```go
res, err := client.Delete("test-index", "1")
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

### Bulk

```go
var buf bytes.Buffer
buf.WriteString(`{"index":{"_index":"test","_id":"1"}}` + "\n")
buf.WriteString(`{"title":"Test"}` + "\n")

res, err := client.Bulk(bytes.NewReader(buf.Bytes()))
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

For a higher-level alternative that handles batching, flushing, and concurrency, see [`esutil.BulkIndexer`](../using-the-api/bulk-indexing.md#_bulk_indexer_helper). The BulkIndexer works with both the low-level and typed clients.

## API reference [_low_level_reference]

The full set of low-level endpoints, parameters, and option functions is documented on pkg.go.dev:

- [`esapi` package reference](https://pkg.go.dev/github.com/elastic/go-elasticsearch/v9/esapi)

## Migrating to the typed API [_low_level_migration_pointer]

You do not have to switch everything at once. An `*elasticsearch.Client` already satisfies the transport interface used by the typed API packages, so you can adopt typed endpoints one at a time while keeping the rest of your code untouched. See the [migration guide](migration.md) for the details.
