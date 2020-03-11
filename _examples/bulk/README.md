# Example: Bulk Indexing

## `default.go`

The [`default.go`](default.go) example demonstrates how to properly operate the Elasticsearch's
[Bulk API]([https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html]).

The example intentionally doesn't use any abstractions or helper functions, to
demonstrate the low-level mechanics of working with the Bulk API:

* iterating over a slice of data and preparing the `meta`/`data` pairs,
* filling a buffer with the payload until the configured threshold for a single batch is reached,
* sending a batch to Elasticsearch,
* checking for a request failure or a failed response,
* checking for individual errors in the response,
* updating a counter of indexed and failed documents,
* printing a report.

```bash
go run default.go -count=100000 -batch=25000

# Bulk: documents [100,000] batch size [25,000]
# ▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁
# → Generated 100,000 articles
# → Sending batch [1/4] [2/4] [3/4] [4/4]
# ▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
# Sucessfuly indexed [100,000] documents in 3.423s (29,214 docs/sec)
```

## `indexer.go`

The [`indexer.go`](indexer.go) example demonstrates how to use the [`esutil.BulkIndexer`](../esutil/bulk_indexer.go) helper for efficient indexing in parallel.

```bash
go run indexer.go -count=100000 -flush=1000000

# BulkIndexer: documents [100,000] workers [8] flush [1.0 MB]
# ▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁
# → Generated 100,000 articles
# ▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
# Sucessfuly indexed [100,000] documents in 1.909s (52,383 docs/sec)
```

The helper allows you to `Add()` bulk indexer items, and flushes each batch based on the configured threshold.

```golang
indexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{})
indexer.Add(
	context.Background(),
	esutil.BulkIndexerItem{
		Action: "index",
		Body:   strings.NewReader(`{"title":"Test"}`),
	})
indexer.Close(context.Background())
```

Please refer to the [`benchmarks`](benchmarks) folder for performance tests with different types of payload.

See the [`kafka`](kafka) folder for an end-to-end example of using the bulk helper for indexing data from a Kafka topic.
