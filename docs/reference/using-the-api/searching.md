---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/runningqueries.html
---

# Searching [search]

This page covers how to build and run search queries with the Go client using the [typed API](/reference/typed-api/index.md) and the [`esdsl`](/reference/typed-api/esdsl.md) builders. For the full {{es}} search API reference, see the [Search API documentation](https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-search). For raw-JSON searches, see the [low-level API](/reference/low-level-api/index.md).

```go subs=true
import "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}/typedapi/esdsl"
```

## Running a search [_running_a_search]

As an example, let's search for a document with a field `name` with a value of `Foo` in the index named `index_name`. The `esdsl` query builders provide a concise, fluent syntax:

```go
res, err := es.Search().
    Index("index_name"). // <1>
    Query(esdsl.NewMatchQuery("name", "Foo")). // <2>
    Do(context.Background())
```

1. The targeted index for this search.
2. `NewMatchQuery` takes the field name and query text directly; no struct nesting required.

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

## Request structures [_request_structures]

If you prefer working with typed structs directly, you can pass a `search.Request` to `Request`:

```go
res, err := es.Search().
    Index("index_name").
    Request(&search.Request{
        Query: &types.Query{
            Term: map[string]types.TermQuery{
                "name": {Value: "Foo"},
            },
        },
    }).Do(context.Background())
```

The `esdsl` builders and struct-based requests produce identical payloads; mix and match based on how complex the query is.

## Paginating results [_paginating_results]

A `Search` call returns at most `size` hits per request (defaulting to 10). Two strategies cover almost every pagination need:

- **`from` + `size`** is the simplest option. It works well for small result sets, but {{es}} refuses to page past the 10,000th hit by default (controlled by `index.max_result_window`). Raising that setting becomes expensive fast because each shard has to materialise and sort `from + size` hits per request.
- **`search_after`** combined with a **point in time (PIT)** is the recommended approach for deep pagination and for exports where a consistent view across pages matters. The PIT pins a snapshot of the index so concurrent writes do not shift results between pages.

Both examples below use the [`esdsl`](/reference/typed-api/esdsl.md) builders with the typed client. The [`_examples/search/pagination.go`](https://github.com/elastic/go-elasticsearch/blob/main/_examples/search/pagination.go) example is a runnable walkthrough of both strategies. See [Paginate search results](https://www.elastic.co/docs/reference/elasticsearch/rest-apis/paginate-search-results) for the full Elasticsearch reference.

### `from` + `size` [_from_size]

```go
const pageSize = 50
for page := 0; ; page++ {
    res, err := es.Search().
        Index("index_name").
        Query(esdsl.NewMatchAllQuery()).
        From(page * pageSize). // <1>
        Size(pageSize).        // <2>
        Do(ctx)
    if err != nil {
        log.Fatal(err)
    }
    for _, hit := range res.Hits.Hits {
        _ = hit
    }
    if len(res.Hits.Hits) < pageSize {
        break
    }
}
```

1. Starting document offset. Must be non-negative and, together with `size`, stay below `index.max_result_window`.
2. Number of hits per page.

### PIT + `search_after` [_pit_search_after]

```go subs=true
import (
    "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}/typedapi/esdsl"
    "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}/typedapi/types/enums/sortorder"
)
```

```go
pit, err := es.OpenPointInTime("index_name").KeepAlive("1m").Do(ctx) // <1>
if err != nil {
    log.Fatal(err)
}
defer es.ClosePointInTime().Id(pit.Id).Do(ctx) // <2>

req := es.Search().
    Query(esdsl.NewMatchAllQuery()).
    Pit(esdsl.NewPointInTimeReference(). // <3>
        Id(pit.Id).
        KeepAlive(esdsl.NewDuration().String("1m"))).
    Sort(esdsl.NewSortOptions(). // <4>
        AddSortOption("_shard_doc", esdsl.NewFieldSort(sortorder.Asc))).
    Size(1000)

for {
    res, err := req.Do(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if len(res.Hits.Hits) == 0 {
        break
    }
    for _, hit := range res.Hits.Hits {
        _ = hit
    }
    last := res.Hits.Hits[len(res.Hits.Hits)-1]
    req = req.SearchAfterValues(last.Sort) // <5>
    if res.PitId != nil { // <6>
        req = req.Pit(esdsl.NewPointInTimeReference().
            Id(*res.PitId).
            KeepAlive(esdsl.NewDuration().String("1m")))
    }
}
```

1. Open a PIT scoped to the target indices. `KeepAlive` only needs to cover the next request, not the whole scan; each search extends it.
2. Release the PIT when you are done. PITs hold shard resources, so do not leak them.
3. Attach the PIT to the search. Do not call `.Index(...)` together with `.Pit(...)`; the index list is baked into the PIT and the server rejects the combination.
4. A deterministic sort is required for `search_after` to produce stable cursor values. `_shard_doc` is the cheapest tie-breaker and is available whenever a PIT is in use. If you already sort on another field, keep that sort and append `_shard_doc` as a secondary, otherwise hits with equal primary sort values can skip or repeat across pages.
5. Advance the cursor using the last hit's sort values. `SearchAfterValues` takes `[]types.FieldValue`, which is exactly the type of `hit.Sort`.
6. The PIT id can rotate between requests. When the response includes a new `PitId`, use it for the next page.

The older `scroll` API still works but is no longer the recommended approach for deep pagination; prefer PIT + `search_after`.

## Raw JSON [_raw_json]

If you want to use your own pre-baked JSON queries using templates or a specific encoder, you can pass the body directly with `Raw`:

```go
res, err := es.Search().
    Index("index_name").
    Raw([]byte(`{
      "query": {
        "term": {
          "user.id": {
            "value": "kimchy",
            "boost": 1.0
          }
        }
      }
    }`)).Do(context.Background())
```

No further validation or serialization is done on what is sent through this method. Setting a payload with `Raw` takes precedence over any request structure you may submit before running the query.
