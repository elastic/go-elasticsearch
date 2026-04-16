---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
---

# Bulk indexing [bulk]

When ingesting many documents, use the Bulk API to send multiple operations in a single request. For the full bulk API reference, see the [Bulk API documentation](https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-bulk).

## Direct bulk requests [_direct_bulk]

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

With the low-level client, build the NDJSON payload yourself and submit it with `Bulk()`:

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
buf.WriteString(`{ "index" : { "_index" : "test", "_id" : "1" } }` + "\n") // <1>
buf.WriteString(`{ "title" : "Test" }` + "\n") // <2>

res, err := client.Bulk(bytes.NewReader(buf.Bytes())) // <3>
if err != nil {
    // Handle error.
}
defer res.Body.Close()
```

1. Each operation line specifies the action and metadata.
2. The document body follows. The final line must end with `\n`.
3. Send the NDJSON payload as an `io.Reader`.

::::::

::::::{tab-item} Fully-typed API
:sync: typed

With the typed client, you can build a bulk request by appending operations and then executing it:

```go
index := "my-index"
id1 := "1"
id2 := "2"

bulk := es.Bulk() // <1>
if err := bulk.IndexOp(
    types.IndexOperation{Index_: &index, Id_: &id1}, // <2>
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

res, err := bulk.Do(context.Background()) // <3>
if err != nil {
    // Handle error.
}
if res.Errors {
    // One or more operations failed.
}
```

1. Create a bulk request builder.
2. Append index operations with metadata and document body.
3. Execute the bulk request.

::::::

:::::::

The client repository contains complete, runnable examples for bulk ingestion (manual NDJSON, `esutil.BulkIndexer`, typed bulk, benchmarks, Kafka ingestion): [`_examples/bulk`](https://github.com/elastic/go-elasticsearch/tree/main/_examples/bulk).

## BulkIndexer helper [_bulk_indexer_helper]

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
    Client:     client,    // <1>
    Index:      "test",    // <2>
    NumWorkers: 4,         // <3>
    FlushBytes: 5_000_000, // <4>
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

1. The Elasticsearch client instance.
2. The default index name for all items.
3. Number of concurrent worker goroutines.
4. Flush threshold in bytes (flush when the buffer reaches this size).

### Explicit flushing [_explicit_flushing]

In addition to automatic flushing (based on `FlushBytes` and `FlushInterval`), you can explicitly flush the indexer at any time using `Flush()`. This drains all currently queued items, flushes all worker buffers to Elasticsearch, and waits for completion. The indexer remains usable after `Flush()` — you can continue adding items.

This is useful when you need to ensure all pending documents are sent before proceeding, for example after ingesting a batch of records or before reading back recently indexed data.

```go
// Add items to the indexer
for _, doc := range documents {
    _ = indexer.Add(ctx, esutil.BulkIndexerItem{
        Action: "index",
        Body:   strings.NewReader(doc),
        OnSuccess: func(_ context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
            log.Printf("Indexed %s: %s", item.DocumentID, res.Result)
        },
        OnFailure: func(_ context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
            log.Printf("Failed %s: %s", item.DocumentID, err)
        },
    })
}

// Flush: send all pending items to Elasticsearch and wait for completion
if err := indexer.Flush(ctx); err != nil {
    // Handle error.
}

// The indexer is still usable — keep adding items or close it when done.
```

`Flush()` follows the same callback model as automatic flushes: per-item results are delivered through the `OnSuccess` and `OnFailure` callbacks on each `BulkIndexerItem`. `Flush` itself only returns an error for transport-level failures or context cancellation.

> **Note:** Do not call `Close()` and recreate the indexer to achieve a synchronous flush.
> The `BulkIndexer` is designed to be long-lived. Use `Flush()` instead, and only call `Close()` once when you are done with the indexer entirely.

:::{dropdown} BulkIndexerConfig full reference
The `esutil.BulkIndexerConfig` struct supports the following fields:

| Field           | Type                                    | Description                                                          |
| --------------- | --------------------------------------- | -------------------------------------------------------------------- |
| `Client`        | `*elasticsearch.Client`                 | The Elasticsearch client (required).                                 |
| `Index`         | `string`                                | Default index name for items that don't specify one.                 |
| `NumWorkers`    | `int`                                   | Number of concurrent worker goroutines (default: number of CPUs).    |
| `FlushBytes`    | `int`                                   | Flush threshold in bytes.                                            |
| `FlushInterval` | `time.Duration`                         | Periodic flush interval.                                             |
| `Pipeline`      | `string`                                | Default ingest pipeline for all items.                               |
| `Refresh`       | `string`                                | Refresh policy after each flush (`"true"`, `"false"`, `"wait_for"`). |
| `Routing`       | `string`                                | Default routing value for all items.                                 |
| `Timeout`       | `time.Duration`                         | Timeout for each bulk request.                                       |
| `OnError`       | `func(context.Context, error)`          | Callback invoked when a bulk request fails.                          |
| `OnFlushStart`  | `func(context.Context) context.Context` | Callback invoked before each flush.                                  |
| `OnFlushEnd`    | `func(context.Context)`                 | Callback invoked after each flush.                                   |
| `Decoder`       | `BulkResponseJSONDecoder`               | Custom JSON decoder for bulk responses.                              |
| `DebugLogger`   | `BulkIndexerDebugLogger`                | Logger for debug output.                                             |

:::
