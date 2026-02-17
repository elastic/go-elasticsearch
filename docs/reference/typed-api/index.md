---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/typedapi.html
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/_getting_started_with_the_api.html
---

# Typed API [typedapi]

The goal for this API is to provide a strongly typed Go API for {{es}}.

This was designed with structures and the Go runtime in mind, following as closely as possible the API and its objects.

## Getting started [_getting_started_with_the_api]

The new typed client can be obtained from the `elasticsearch` package using the `NewTypedClient` function. This new API takes the same arguments as the previous one and uses the same transport underneath.

Connection to an {{es}} cluster is identical to the existing client, only the API changes:

```go
client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
    // Proper configuration for your Elasticsearch cluster.
})
```

The code is generated from the [elasticsearch-specification](https://github.com/elastic/elasticsearch-specification).

## NDJSON endpoints [_ndjson_endpoints]

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

## Raw payloads [_raw_payloads]

If you already have a newline-delimited JSON payload, you can submit it directly with `Raw(io.Reader)` on the request builder.
