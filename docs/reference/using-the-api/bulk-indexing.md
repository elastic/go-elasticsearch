---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
---

# Bulk indexing [bulk]

When ingesting many documents, use the Bulk API to send multiple operations in a single request.

With the typed client, you can build a bulk request by appending operations and then executing it:

```go
index := "my-index"
id1 := "1"
id2 := "2"

bulk := es.Bulk()
if err := bulk.IndexOp(
    types.IndexOperation{Index_: &index, Id_: &id1},
    map[string]any{"title": "Test 1"},
); err != nil {
    // Handle error.
}

if err := bulk.IndexOp(
    types.IndexOperation{Index_: &index, Id_: &id2},
    map[string]any{"title": "Test 2"},
); err != nil {
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

The client repository also contains complete, runnable examples for bulk ingestion (manual NDJSON, `esutil.BulkIndexer`, typed bulk, benchmarks, Kafka ingestion): [`_examples/bulk`](https://github.com/elastic/go-elasticsearch/tree/main/_examples/bulk).

If you prefer the classic client (NDJSON + `Bulk()`), you can build the NDJSON payload yourself and submit it with `Bulk()`:

```go
client, err := elasticsearch.NewDefaultClient()
if err != nil {
    // Handle error.
}
defer func() {
    if err := client.Close(context.Background()); err != nil {
        // Handle error.
    }
}()

var buf bytes.Buffer
buf.WriteString(`{ "index" : { "_index" : "test", "_id" : "1" } }` + "\n")
buf.WriteString(`{ "title" : "Test" }` + "\n") // NOTE: the final line must end with \n

res, err := client.Bulk(bytes.NewReader(buf.Bytes()))
if err != nil {
    // Handle error.
}
defer res.Body.Close()
```

For a higher-level API that takes care of batching, flushing, and concurrency, use the `esutil.BulkIndexer` helper.

The BulkIndexer is designed to be **long-lived**: create it once, keep adding items over time (potentially from multiple goroutines), and call `Close()` once when you are done (for example with `defer`).

```go
client, err := elasticsearch.NewDefaultClient()
if err != nil {
    // Handle error.
}
defer func() {
    if err := client.Close(context.Background()); err != nil {
        // Handle error.
    }
}()

ctx := context.Background()
indexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
    Client:     client, // The Elasticsearch client
    Index:      "test", // The default index name
    NumWorkers: 4,
    FlushBytes: 5_000_000,
})
if err != nil {
    // Handle error.
}

defer func() {
    if err := indexer.Close(ctx); err != nil {
        // Handle error.
    }
}()

_ = indexer.Add(ctx, esutil.BulkIndexerItem{
    Action:     "index",
    DocumentID: "1",
    Body:       strings.NewReader(`{"title":"Test"}`),
})
```
