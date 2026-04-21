# Example: Search

The [`default.go`](default.go) example shows two equivalent ways to call the
[`_search`](https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-search)
API with the typed client (`TypedClient`):

- building the query with the [`esdsl`](../../typedapi/esdsl) helpers (recommended), and
- building the query with raw [`types.*`](../../typedapi/types) structs.

```bash
make test
# or
go run -tags search_default default.go
```

The [`pagination.go`](pagination.go) example walks the same API through both
supported pagination strategies:

- `from` + `size` for small result sets, and
- a point in time (PIT) plus `search_after` for deep pagination over a
  consistent snapshot of the index.

```bash
go run -tags search_pagination pagination.go
```

See [Paginate search results](https://www.elastic.co/docs/reference/elasticsearch/rest-apis/paginate-search-results)
for the corresponding Elasticsearch reference.

## Notes

- `TypedClient.Search()` returns a fluent builder. Chain `Index(...)`,
  `Request(&search.Request{...})`, and finish with `Do(ctx)` to get a decoded
  `*search.Response`.
- The `Query` field on `search.Request` is `*types.Query`. Use
  `esdsl.NewBoolQuery().…QueryCaster()` to produce one from the builders, or
  construct the struct directly as shown in `runTypesExample`.
- `res.Hits.Total` is a pointer — guard for `nil` before reading
  `res.Hits.Total.Value` in case `track_total_hits` is disabled on the request.
- Fields on `types.Hit` that map to optional JSON properties (such as `_id`,
  `_score`) are pointers; dereference them only after a `nil` check.
- Do not confuse `typedapi/core/search` (the generic search API) with
  `typedapi/eql/search` (EQL search, whose request shape is different).
