---
navigation_title: "Migrating to the typed API"
---

# Migrating from the low-level API to the typed API [low_level_migration]

The Go client ships two API surfaces over a shared transport: the [low-level API](index.md) (`*elasticsearch.Client`) and the [typed API](../typed-api/index.md) (`*elasticsearch.TypedClient`). This page explains how to move existing code from the low-level API to the typed API, and shows how to migrate gradually (one endpoint at a time) without rewriting your whole codebase.

## Why migrate [_why_migrate]

The typed API gives you:

- **Type-safe requests.** Requests are Go structs generated from the {{es}} specification, so invalid fields are caught at compile time.
- **Decoded responses.** Every endpoint returns a typed response; no manual JSON parsing and no forgetting to close response bodies.
- **Fluent builders.** The [`esdsl`](../typed-api/esdsl.md) package provides chainable builders for queries, aggregations, mappings, and sort options, replacing deeply nested struct literals.
- **Less boilerplate.** No more `io.Reader` wrapping, no `defer res.Body.Close()`, no `res.IsError()` checks.

## Full migration [_full_migration]

If you can commit to the typed API everywhere, replace the constructor and the call sites:

:::::::{tab-set}
::::::{tab-item} Before

```go
// Modern functional-options form.
client, err := elasticsearch.New(
    elasticsearch.WithAddresses("https://localhost:9200"),
    elasticsearch.WithAPIKey("API_KEY"),
)

// Older, deprecated Config form (still works).
client, err := elasticsearch.NewClient(elasticsearch.Config{
    Addresses: []string{"https://localhost:9200"},
    APIKey:    "API_KEY",
})
```

::::::

::::::{tab-item} After

```go
client, err := elasticsearch.NewTyped(
    elasticsearch.WithAddresses("https://localhost:9200"),
    elasticsearch.WithAPIKey("API_KEY"),
)
```

::::::

:::::::

Both constructors take the same [functional options](../configuration.md) and share the same transport, retry, instrumentation, and interceptor infrastructure. The deprecated `NewClient(Config{...})` form has a typed equivalent in `NewTypedClient(Config{...})`, but both are deprecated, so a migration is a good moment to move to the functional-options form as well.

### Side-by-side operations [_side_by_side]

A few common call-site translations:

**Index a document**

:::::::{tab-set}
::::::{tab-item} Low-level

```go
data, _ := json.Marshal(doc)
res, err := client.Index(
    "my-index",
    bytes.NewReader(data),
    client.Index.WithDocumentID("1"),
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

::::::

::::::{tab-item} Typed

```go
res, err := client.Index("my-index").
    Id("1").
    Document(doc).
    Do(ctx)
```

::::::

:::::::

**Search**

:::::::{tab-set}
::::::{tab-item} Low-level

```go
query := `{"query":{"match":{"title":{"query":"golang"}}}}`
res, err := client.Search(
    client.Search.WithIndex("my-index"),
    client.Search.WithBody(strings.NewReader(query)),
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()

var result map[string]any
json.NewDecoder(res.Body).Decode(&result)
```

::::::

::::::{tab-item} Typed + esdsl

```go subs=true
import "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}/typedapi/esdsl"
```

```go
res, err := client.Search().
    Index("my-index").
    Query(esdsl.NewMatchQuery("title", "golang")).
    Do(ctx)
if err != nil {
    log.Fatal(err)
}
for _, hit := range res.Hits.Hits {
    fmt.Println(hit.Source_)
}
```

::::::

:::::::

**Bulk**

:::::::{tab-set}
::::::{tab-item} Low-level

```go
var buf bytes.Buffer
buf.WriteString(`{"index":{"_index":"my-index","_id":"1"}}` + "\n")
buf.WriteString(`{"title":"Test"}` + "\n")

res, err := client.Bulk(bytes.NewReader(buf.Bytes()))
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

::::::

::::::{tab-item} Typed

```go
index := "my-index"
id := "1"

bulk := client.Bulk()
if err := bulk.IndexOp(
    types.IndexOperation{Index_: &index, Id_: &id},
    map[string]any{"title": "Test"},
); err != nil {
    log.Fatal(err)
}

res, err := bulk.Do(ctx)
if err != nil {
    log.Fatal(err)
}
if res.Errors {
    // One or more operations failed.
}
```

::::::

:::::::

For more call-site patterns, see [CRUD operations](../using-the-api/crud-operations.md), [Searching](../using-the-api/searching.md), [Aggregations](../using-the-api/aggregations.md), and [Bulk indexing](../using-the-api/bulk-indexing.md).

## Partial migration [_partial_migration]

You do not have to switch the whole codebase at once. The typed API endpoints live in small, focused packages under `typedapi/` (for example `typedapi/core/search`, `typedapi/indices/create`, `typedapi/core/bulk`). Each endpoint package exports a constructor that takes an `elastictransport.Interface`, which is just:

```go
// from github.com/elastic/elastic-transport-go/v8/elastictransport
type Interface interface {
    Perform(*http.Request) (*http.Response, error)
}
```

`*elasticsearch.Client` (low-level) and `*elasticsearch.TypedClient` (typed) both embed `BaseClient`, which implements `Perform`. That means **you can pass an existing low-level client directly into any typed endpoint package**: no second client, no duplicated configuration, same transport and connection pool.

### Example: migrate just the search call

Keep your existing low-level `*elasticsearch.Client` for every call except search. Build a `search.New(client)` wherever you want typed search:

```go subs=true
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}"
    "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}/typedapi/core/search"
    "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}/typedapi/esdsl"
)

func main() {
    client, err := elasticsearch.New( // <1>
        elasticsearch.WithAddresses("https://localhost:9200"),
        elasticsearch.WithAPIKey("API_KEY"),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close(context.Background())

    // Unchanged: all other calls keep using the low-level client.
    res, err := client.Indices.Exists([]string{"my-index"})
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    // Migrated: search uses the typed package directly, backed by the
    // same client and the same transport.
    typedSearch := search.New(client) // <2>
    result, err := typedSearch.
        Index("my-index").
        Query(esdsl.NewMatchQuery("title", "golang")). // <3>
        Do(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    for _, hit := range result.Hits.Hits {
        fmt.Println(hit.Source_)
    }
}
```

1. The existing low-level client. Keep it.
2. `search.New` accepts any `elastictransport.Interface`. `*elasticsearch.Client` satisfies it via the embedded `BaseClient.Perform` method, so no adapter is needed.
3. The `esdsl` builders work with any typed endpoint, regardless of how you built the client.

The same pattern applies to every typed endpoint package: `typedapi/indices/create`, `typedapi/core/bulk`, `typedapi/cluster/health`, and so on. Import the package for the endpoint you want to migrate, call `New(client)`, and use the builder.

### Get a full typed client from an existing low-level client [_to_typed]

If you want the full typed API surface (not just one endpoint package at a time) but the existing `*elasticsearch.Client` must keep working for other call sites, use `ToTyped`. It returns a `*elasticsearch.TypedClient` that reuses the source client's transport, connection pool, and configuration, without building a second transport:

```go
client, err := elasticsearch.New(
    elasticsearch.WithAddresses("https://localhost:9200"),
    elasticsearch.WithAPIKey("API_KEY"),
)
if err != nil {
    log.Fatal(err)
}
defer client.Close(context.Background())

typed := client.ToTyped() // <1>

// Low-level call sites keep working.
res, err := client.Indices.Exists([]string{"my-index"})
// ...

// Typed call sites go through `typed`.
info, err := typed.Info().Do(context.Background())
// ...
```

1. `ToTyped` copies the source client's compatibility mode, auto-drain-body, and disable-meta-header settings into the new `*TypedClient`, and shares the underlying transport (and its connection pool).

This is the most common choice when the `*Client` is constructed in a shared infrastructure package you don't own, and you can't swap the constructor for `elasticsearch.NewTyped` without coordinating with other teams.

#### Call `ToTyped` once and reuse it [_to_typed_reuse]

`ToTyped` allocates a new typed API tree, and the returned client re-runs the product check on its first request. Call it once at setup and cache the result; do not call it per request or per handler invocation.

#### Shared lifecycle [_to_typed_close]

Because the `*TypedClient` returned by `ToTyped` shares the transport with the source `*Client`, **closing either one closes the connection pool used by both**. After `typed.Close(ctx)` returns, subsequent requests on the original `*Client` will fail with a closed-connection error, and vice versa. In a partial-migration setup, treat the source `*Client` as the lifecycle owner: let whichever code owns the `*Client` call `Close`, and do not call `Close` on the `*TypedClient` returned by `ToTyped`.

### Strategy [_partial_migration_strategy]

A typical gradual migration looks like:

1. **Start with the hot spots.** Migrate the endpoints where typed requests and decoded responses give the most value: usually search, aggregations, and bulk indexing.
2. **Let new code use `NewTyped` directly.** New call sites can use `*elasticsearch.TypedClient` from the start; old call sites keep using the low-level client.
3. **Replace the constructor last.** Once most call sites are typed, swap `elasticsearch.New(...)` for `elasticsearch.NewTyped(...)`. Any typed endpoint packages you used for partial migration keep working against the new client, because `*TypedClient` also satisfies `elastictransport.Interface`. What stops compiling is the remaining low-level call sites (`client.Search(...)`, `client.Indices.Create(...)`, `client.Bulk(...)`, etc.), because `*TypedClient` does not embed `*esapi.API`. Treat those as the last batch to migrate.

## Endpoints not covered by the typed API [_typed_api_gaps]

The typed API covers the most widely used endpoints, but it does not yet cover every endpoint in the REST API. If you hit a gap, keep calling that endpoint through the low-level client: the two share the same transport, so you can mix and match freely in the same application. Check the [`typedapi` godoc](https://pkg.go.dev/github.com/elastic/go-elasticsearch/v9/typedapi) for the current set of supported endpoints.

## Further reading [_migration_further_reading]

- [Typed API overview](../typed-api/index.md): namespaces, builders, NDJSON payloads.
- [esdsl builders](../typed-api/esdsl.md): fluent query, aggregation, and mapping construction.
- [Typed API conventions](../typed-api/conventions.md): naming, enums, and unions.
