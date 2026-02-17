---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
---

# Aggregations [aggregations]

This page covers how to run aggregations with the Go client. For the full aggregations reference, see [Aggregations](docs-content://explore-analyze/query-filter/aggregations.md) in the {{es}} documentation.

Given documents with a `price` field, we run a sum aggregation on `index_name`.

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

```go
query := `{
  "size": 0,
  "aggs": {
    "total_prices": {
      "sum": {
        "field": "price"
      }
    }
  }
}`

res, err := client.Search(
    client.Search.WithIndex("index_name"), // <1>
    client.Search.WithBody(strings.NewReader(query)), // <2>
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

1. Specifies the index name.
2. Pass the aggregation query as raw JSON. Setting `"size": 0` retrieves only the aggregation result.

::::::

::::::{tab-item} Fully-typed API
:sync: typed

The `some` package provides helpers for inline pointers on primitive types:

```go
import "github.com/elastic/go-elasticsearch/v9/typedapi/some"
```

```go
totalPricesAgg, err := es.Search().
    Index("index_name"). // <1>
    Request(
        &search.Request{
            Size: some.Int(0), // <2>
            Aggregations: map[string]types.Aggregations{
                "total_prices": { // <3>
                    Sum: &types.SumAggregation{
                        Field: some.String("price"), // <4>
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

::::::

::::::{tab-item} esdsl API
:sync: esdsl

The [`esdsl`](/reference/typed-api/esdsl.md) aggregation builders let you define aggregations with a fluent syntax:

```go subs=true
import "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}/typedapi/esdsl"
```

```go
totalPricesAgg, err := es.Search().
    Index("index_name"). // <1>
    Size(0). // <2>
    AddAggregation("total_prices", // <3>
        esdsl.NewSumAggregation().Field("price"), // <4>
    ).
    Do(context.Background())
```

1. Specifies the index name.
2. Sets the size to 0 to retrieve only the result of the aggregation.
3. Names the aggregation `total_prices`.
4. `NewSumAggregation()` creates a sum aggregation builder -- chain `.Field()` to set the target field.

::::::

:::::::
