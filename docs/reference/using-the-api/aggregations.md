---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
---

# Aggregations [aggregations]

This page covers how to run aggregations with the Go client using the [typed API](/reference/typed-api/index.md) and the [`esdsl`](/reference/typed-api/esdsl.md) builders. For the full aggregations reference, see [Aggregations](docs-content://explore-analyze/query-filter/aggregations.md) in the {{es}} documentation. For raw-JSON aggregations, see the [low-level API](/reference/low-level-api/index.md).

```go subs=true
import "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}/typedapi/esdsl"
```

Given documents with a `price` field, we run a sum aggregation on `index_name`. The `esdsl` aggregation builders let you define aggregations with a fluent syntax:

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
4. `NewSumAggregation()` creates a sum aggregation builder. Chain `.Field()` to set the target field.
