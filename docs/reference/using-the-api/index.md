---
navigation_title: "Using the API"
---

# Using the API [using_the_api]

The Go client for {{es}} offers two API styles:

- **Low-level API** — provides a one-to-one mapping with the {{es}} REST API. Each endpoint accepts and returns raw `io.Reader` bodies and `*esapi.Response` objects. This gives you full control over serialization and request construction.
- **Fully-typed API** — provides a strongly typed, builder-pattern API generated from the [elasticsearch-specification](https://github.com/elastic/elasticsearch-specification). Requests and responses are modeled as Go structs, giving you compile-time safety and IDE autocompletion.

Both API styles share the same underlying transport and [configuration](../configuration.md). You can even use both in the same application.

## Comparing API styles [_comparing_api_styles]

|                        | Low-level API                                   | Fully-typed API                                  |
| ---------------------- | ----------------------------------------------- | ------------------------------------------------ |
| **Type safety**        | Runtime (raw JSON)                              | Compile-time (Go structs)                        |
| **IDE autocompletion** | Limited                                         | Full support for fields, enums, and methods      |
| **Serialization**      | Manual (`io.Reader`, `json.Marshal`)            | Automatic (structs marshaled by the client)      |
| **Response handling**  | Raw `*esapi.Response` with `io.ReadCloser` body | Typed response structs                           |
| **Flexibility**        | Full control over request/response bytes        | Constrained to the specification model           |
| **Code verbosity**     | Lower for simple requests                       | Lower for complex queries with nested structures |

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

**Choose the typed API when:**

- You want compile-time safety and IDE autocompletion
- You are building complex queries with deeply nested structures
- You want typed response objects instead of parsing JSON manually

::::{tip}
You don't have to choose one exclusively. Both APIs share the same transport, so you can use the typed API for most operations and fall back to the low-level API for edge cases.
::::

## Side-by-side example [_side_by_side_example]

Here is the same search operation using both API styles:

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

:::::::

## What's covered in this section [_using_api_contents]

- [CRUD operations](crud-operations.md) — Index, get, update, and delete documents
- [Searching](searching.md) — Build and run search queries
- [Aggregations](aggregations.md) — Run aggregations on your data
- [Bulk indexing](bulk-indexing.md) — Efficiently ingest large volumes of documents
- [ES|QL](esql.md) — Use the Elasticsearch Query Language
- [Dense vectors and kNN search](dense-vectors.md) — Work with vector embeddings and similarity search
