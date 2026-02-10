---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/typedapi.html
---

# Typed API [typedapi]

The goal for this API is to provide a strongly typed Go API for {{es}}.

This was designed with structures and the Go runtime in mind, following as closely as possible the API and its objects.

The typed API includes endpoints that use NDJSON bodies such as `bulk` and `msearch`.

For example, you can build a bulk request by appending operations and then executing it:

```go
client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
    // Proper configuration for your Elasticsearch cluster.
})
if err != nil {
    // Handle error.
}

index := "my-index"
id := "1"
bulk := client.Bulk()
if err := bulk.IndexOp(types.IndexOperation{Index_: &index, Id_: &id}, map[string]any{"title": "Test"}); err != nil {
    // Handle error.
}

res, err := bulk.Do(context.Background())
if err != nil {
    // Handle error.
}
if res.Errors {
    // One or more operations failed.
}
```

If you already have a newline-delimited JSON payload, you can submit it directly with `Raw(io.Reader)` on the request builder.
