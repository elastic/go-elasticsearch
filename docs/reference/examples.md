---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
---

# Examples [examples]

This sections lists a series of frequent use cases that will help you start with this new API.

<!-- markdownlint-disable MD051 -->

- [Creating an index](#indices)
- [Indexing a document](#indexing)
- [Bulk indexing](#bulk)
- [Retrieving a document](#retrieving_document)
- [Search](#search)
- [Aggregations](#aggregations)
<!-- markdownlint-enable MD051 -->

::::{note}
This is a work in progress, the documentation will evolve in the future.
::::

## Creating an index [indices]

For this example on how to create an index, lets create an index named `test-index` and provide a mapping for the field `price` which will be an integer. Notice how using the builder for the `IntegerNumberProperty` will automatically apply the correct value for the `type` field.

```go
res, err := es.Indices.Create("test-index").
    Request(&create.Request{
        Mappings: &types.TypeMapping{
            Properties: map[string]types.Property{
                "price": types.NewIntegerNumberProperty(),
            },
        },
    }).
    Do(nil)
```

## Indexing a document [indexing]

The standard way of indexing a document is to provide a `struct` to the `Request` method, the standard `json/encoder` will be run on your structure and the result will be sent to {{es}}.

```go
document := struct {
    Id    int    `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
}{
    Id:    1,
    Name:  "Foo",
    Price: 10,
}

res, err := es.Index("index_name").
    Request(document).
    Do(context.Background())
```

Alternatively, you can use the `Raw` method and provide the already serialized JSON:

```go
res, err := es.Index("index_name").
            Raw([]byte(`{
              "id": 1,
              "name": "Foo",
              "price": 10
            }`)).Do(context.Background())
```

## Bulk indexing [bulk]

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

The client repository also contains complete, runnable examples for bulk ingestion (manual NDJSON, `esutil.BulkIndexer`, typed bulk, benchmarks, Kafka ingestion): [`_examples/bulk`](https://github.com/elastic/go-elasticsearch/tree/master/_examples/bulk).

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

## Retrieving a document [retrieving_document]

Retrieving a document follows the API as part of the argument of the endpoint. In order you provide the `index`, the `id` and then run the query:

```go
res, err := es.Get("index_name", "doc_id").Do(context.Background())
```

## Checking for a document existence [_checking_for_a_document_existence]

If you do not wish to retrieve the content of the document and want only to check if it exists in your index, we provide the `IsSuccess` shortcut:

```go
if exists, err := es.Exists("index_name", "doc_id").IsSuccess(nil); exists {
    // The document exists !
} else if err != nil {
   // Handle error.
}
```

Result is `true` if everything succeeds, `false` if the document doesn't exist. If an error occurs during the request, you will be granted with a `false` and the relevant error.

## Search [search]

Building a search query can be done with structs or builder. As an example, let's search for a document with a field `name` with a value of `Foo` in the index named `index_name`.

With a struct request:

```go
res, err := es.Search().
    Index("index_name"). <1>
    Request(&search.Request{ <2>
        Query: &types.Query{
            Match: map[string]types.MatchQuery{
                "name": {Query: "Foo"}, <3>
            },
        },
    }).Do(context.Background()) <4>
```

1. The targeted index for this search.
2. The request part of the search.
3. Match query specifies that `name` should match `Foo`.
4. The query is run with a `context.Background` and returns the response.

It produces the following JSON:

```json
{
  "query": {
    "match": {
      "name": {
        "query": "Foo"
      }
    }
  }
}
```

## Aggregations [aggregations]

Given documents with a `price` field, we run a sum aggregation on `index_name`:

```go
totalPricesAgg, err := es.Search().
    Index("index_name"). <1>
    Request(
        &search.Request{
            Size: some.Int(0), <2>
            Aggregations: map[string]types.Aggregations{
                "total_prices": { <3>
                    Sum: &types.SumAggregation{
                        Field: some.String("price"), <4>
                    },
                },
            },
        },
    ).Do(context.Background())
```

1. Specifies the index name.
2. Sets the size to 0 to retrieve only the result of the aggregation.
3. Specifies the field name on which the sum aggregation runs.
4. The `SumAggregation` is part of the `Aggregations` map.
