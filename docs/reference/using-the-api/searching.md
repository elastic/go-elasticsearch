---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/runningqueries.html
---

# Searching [search]

This page covers how to build and run search queries with the Go client. For the full {{es}} search API reference, see the [Search API documentation](https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-search).

## Running a search [_running_a_search]

As an example, let's search for a document with a field `name` with a value of `Foo` in the index named `index_name`.

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

```go
query := `{
  "query": {
    "match": {
      "name": { "query": "Foo" }
    }
  }
}`

res, err := client.Search(
    client.Search.WithIndex("index_name"),       // <1>
    client.Search.WithBody(strings.NewReader(query)), // <2>
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close() // <3>
```

1. The targeted index for this search.
2. Pass the JSON query body as an `io.Reader`.
3. Always close the response body when done.

::::::

::::::{tab-item} Fully-typed API
:sync: typed

With a struct request:

```go
res, err := es.Search().
    Index("index_name"). // <1>
    Request(&search.Request{ // <2>
        Query: &types.Query{
            Match: map[string]types.MatchQuery{
                "name": {Query: "Foo"}, // <3>
            },
        },
    }).Do(context.Background()) // <4>
```

1. The targeted index for this search.
2. The request part of the search.
3. Match query specifies that `name` should match `Foo`.
4. The query is run with a `context.Background` and returns the typed response.

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

::::::

:::::::

## Request structures [_request_structures]

The typed API models each endpoint's body as a Go `Request` struct. For example, a simple term query for `"Foo"` in the `name` field:

```go
search.Request{
    Query: &types.Query{
        Term: map[string]types.TermQuery{
            "name": {Value: "Foo"},
        },
    },
}
```

## Raw JSON [_raw_json]

If you want to use your own pre-baked JSON queries using templates or a specific encoder, you can pass the body directly. Both API styles support raw JSON payloads:

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

```go
res, err := client.Search(
    client.Search.WithIndex("index_name"),
    client.Search.WithBody(strings.NewReader(`{
      "query": {
        "term": {
          "user.id": {
            "value": "kimchy",
            "boost": 1.0
          }
        }
      }
    }`)),
)
```

::::::

::::::{tab-item} Fully-typed API
:sync: typed

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

::::::

:::::::
