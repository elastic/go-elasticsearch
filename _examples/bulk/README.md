# Example: Bulk Indexing

## `bulk.go`

The [`bulk.go`](bulk.go) example demonstrates how to properly operate the Elasticsearch's
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
go run bulk.go -count=100000 -batch=25000

# > Generated 100000 articles
# > Batch 1  of 4
# > Batch 2  of 4
# > Batch 3  of 4
# > Batch 4  of 4
# ================================================================================
# Sucessfuly indexed [100000] documents in 8.02s (12469 docs/sec)
```
