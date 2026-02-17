---
navigation_title: "Using the API"
---

# Using the API [using_the_api]

The Go client for {{es}} offers two API styles, plus fluent DSL builders that enhance the typed API:

- **Low-level API** — provides a one-to-one mapping with the {{es}} REST API. Each endpoint accepts and returns raw `io.Reader` bodies and `*esapi.Response` objects. This gives you full control over serialization and request construction.
- **Fully-typed API** — provides a strongly typed, builder-pattern API generated from the [elasticsearch-specification](https://github.com/elastic/elasticsearch-specification). Requests and responses are modeled as Go structs, giving you compile-time safety and IDE autocompletion.
- **[esdsl builders](/reference/typed-api/esdsl.md)** — a companion to the typed API that provides fluent, chainable builders for constructing queries, aggregations, mappings, and sort options. Instead of manually assembling nested structs, you use `esdsl.NewMatchQuery("title", "Go")` and pass it directly to the typed client.

All three approaches share the same underlying transport and [configuration](../configuration.md). You can mix and match them in the same application.

## Comparing API styles [_comparing_api_styles]

|                        | Low-level API                                   | Fully-typed API                                  | esdsl API                                                      |
| ---------------------- | ----------------------------------------------- | ------------------------------------------------ | -------------------------------------------------------------- |
| **Type safety**        | Runtime (raw JSON)                              | Compile-time (Go structs)                        | Compile-time (fluent builders)                                 |
| **IDE autocompletion** | Limited                                         | Full support for fields, enums, and methods      | Full support with guided method chains                         |
| **Serialization**      | Manual (`io.Reader`, `json.Marshal`)            | Automatic (structs marshaled by the client)      | Automatic (builders produce typed structs)                     |
| **Response handling**  | Raw `*esapi.Response` with `io.ReadCloser` body | Typed response structs                           | Typed response structs (same as fully-typed)                   |
| **Flexibility**        | Full control over request/response bytes        | Constrained to the specification model           | Constrained to the specification model                         |
| **Code verbosity**     | Lower for simple requests                       | Lower for complex queries with nested structures | Lowest for complex queries - minimal nesting, fluent chaining  |

::::{important}
When using the low-level API, you **must** always read and close the response body, even if your code does not use it. Failing to do so prevents Go's HTTP client from reusing the underlying TCP connection, which degrades performance and can cause resource leaks.

```go
res, err := client.Search(
    client.Search.WithIndex("my-index"),
    client.Search.WithBody(strings.NewReader(query)),
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

The fully-typed API handles this automatically. Each endpoint's `Do()` method reads and closes the response body internally, so you only work with typed response structs.
::::

## When to use which [_when_to_use_which]

**Choose the low-level API when:**

- You need full control over request serialization (e.g., custom JSON encoders)
- You are working with endpoints not yet covered by the typed API
- You prefer working with raw JSON bytes

**Choose the typed API with structs when:**

- You want compile-time safety and IDE autocompletion
- You are working with simple, flat request structures
- You want typed response objects instead of parsing JSON manually

**Choose the typed API with esdsl builders when:**

- You are building complex queries with deeply nested structures (bool queries, compound filters)
- You are composing aggregations with sub-aggregations
- You want the most concise code with fluent method chaining
- You are constructing queries dynamically at runtime

::::{tip}
You don't have to choose one exclusively. All three approaches share the same transport, so you can use `esdsl` builders for queries, typed structs for simple operations, and fall back to the low-level API for edge cases - all in the same application.
::::

## Side-by-side example [_side_by_side_example]

Here is the same search operation using all three approaches:

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

```go
res, err := es.Search(
    es.Search.WithIndex("my-index"),
    es.Search.WithBody(strings.NewReader(`{
        "query": {
            "match": {
                "title": "golang"
            }
        }
    }`)),
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()

// Parse the response body manually
var result map[string]any
json.NewDecoder(res.Body).Decode(&result)
```

::::::

::::::{tab-item} Fully-typed API
:sync: typed

```go
res, err := es.Search().
    Index("my-index").
    Request(&search.Request{
        Query: &types.Query{
            Match: map[string]types.MatchQuery{
                "title": {Query: "golang"},
            },
        },
    }).
    Do(context.Background())
if err != nil {
    log.Fatal(err)
}

// Response is already typed
for _, hit := range res.Hits.Hits {
    fmt.Println(hit.Source_)
}
```

::::::

::::::{tab-item} esdsl API
:sync: esdsl

The [`esdsl`](/reference/typed-api/esdsl.md) package provides fluent builders that plug directly into the typed API:

```go subs=true
import "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}/typedapi/esdsl"
```

```go
res, err := es.Search().
    Index("my-index").
    Query(esdsl.NewMatchQuery("title", "golang")). // <1>
    Do(context.Background())
if err != nil {
    log.Fatal(err)
}

// Response is already typed
for _, hit := range res.Hits.Hits {
    fmt.Println(hit.Source_)
}
```

1. The `esdsl` builder replaces the struct-based query construction with a fluent, chainable syntax.

::::::

:::::::

## What's covered in this section [_using_api_contents]

- [CRUD operations](crud-operations.md) — Index, get, update, and delete documents
- [Searching](searching.md) — Build and run search queries
- [Aggregations](aggregations.md) — Run aggregations on your data
- [Bulk indexing](bulk-indexing.md) — Efficiently ingest large volumes of documents
- [ES|QL](esql.md) — Use the Elasticsearch Query Language
- [Dense vectors and kNN search](dense-vectors.md) — Work with vector embeddings and similarity search
