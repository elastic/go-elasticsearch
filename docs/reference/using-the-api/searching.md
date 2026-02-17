---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/runningqueries.html
---

# Searching [search]

This page covers how to build and run search queries with the Go client.

## Request structures [_request_structures]

Each endpoint comes with a Request type that represents the body of its request. For example, a simple search request for a term "Foo" in the `name` field could be written like this:

```go
search.Request{
    Query: &types.Query{
        Term: map[string]types.TermQuery{
            "name": {Value: "Foo"},
        },
    },
}
```

## Running a search [_running_a_search]

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

## Raw JSON [_raw_json]

Lastly if you want to use your own pre-baked JSON queries using templates or even a specific encoder, you can pass the body directly to the `Raw` method of the endpoint:

```go
es.Search().Raw([]byte(`{
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

No further validation or serialization is done on what is sent through this method, setting a payload with this takes precedence over any request structure you may submit before running the query.
