---
navigation_title: "Using the API"
---

# Using the API [using_the_api]

The examples in this section use the [typed API](/reference/typed-api/index.md) together with the [`esdsl`](/reference/typed-api/esdsl.md) builders: the recommended way to use the Go client. The typed API gives you strongly typed request and response structs generated from the [elasticsearch-specification](https://github.com/elastic/elasticsearch-specification), and `esdsl` provides fluent, chainable builders for queries, aggregations, mappings, and sort options.

If you need raw-JSON control or an endpoint not yet covered by the typed API, see the [low-level API](/reference/low-level-api/index.md). The two APIs share the same transport, so you can mix them in the same application. See the [migration guide](/reference/low-level-api/migration.md) for how to adopt the typed API one endpoint at a time.

## Example [_typed_example]

The same search operation, expressed with the typed API and `esdsl`:

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

for _, hit := range res.Hits.Hits { // <2>
    fmt.Println(hit.Source_)
}
```

1. `esdsl` builders produce typed query structs with a fluent, chainable syntax. No deeply nested struct literals.
2. The response is already typed and decoded. No `io.Reader` to parse, no `defer res.Body.Close()`.

## What's covered in this section [_using_api_contents]

- [CRUD operations](crud-operations.md): index, get, update, and delete documents
- [Searching](searching.md): build and run search queries
- [Aggregations](aggregations.md): run aggregations on your data
- [Bulk indexing](bulk-indexing.md): efficiently ingest large volumes of documents
- [ES|QL](esql.md): use the Elasticsearch Query Language
- [Dense vectors and kNN search](dense-vectors.md): work with vector embeddings and similarity search
