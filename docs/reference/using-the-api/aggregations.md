---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
---

# Aggregations [aggregations]

This page covers how to run aggregations with the Go client.

Given documents with a `price` field, we run a sum aggregation on `index_name`.

The `some` package provides helpers for inline pointers on primitive types:

```go
import "github.com/elastic/go-elasticsearch/v9/typedapi/some"
```

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
