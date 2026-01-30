---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/conventions.html
---

# Conventions [conventions]

This section details the conventions upon which the typed client is built.

- [Structure](#structure)
- [Naming](#naming)
- [Endpoints](#endpoints)
- [Requests](#requests)
- [Responses](#responses)
- [Types](#types)

## Structure [structure]

The Typed client lives in the `typedapi` package within the `go-elasticsearch` repository.

The entire client is summed in an index at the root of the package for convenient access after instantiation.

Each endpoint resides in its own package within `typedapi` and contains the client for this endpoint, and the `Request` struct if applicable.

The requests are based on a collection of structures generated from the [elasticsearch-specification](https://github.com/elastic/elasticsearch-specification) repository and gathered in a `types` package within `typedapi`.

## Naming [naming]

Whenever appropriate, names may be suffixed with an underscore:

- To avoid collision with protected keywords (`range`, `if`, `type`, and so on).
- To reflect the presence of a leading underscore in the API like `\_index` vs `Index_` or `\_source` vs `Source_`.

## Endpoints [endpoints]

All the available endpoints are generated in separate packages and assembled in the client. The `core` namespace is duplicated at the root of the client for convenient access.

Each endpoint follows a factory pattern which returns a pointer to a new instance each time.

```go
res, err := es.Search().Index("index_name").AllowPartialSearchResults(true).Do(context.Background())
```

If parameters are needed for the specific endpoint you are using, those will be present as arguments in the same order as the API:

```go
es.Create("index_name", "doc_id").Do(context.Background())
```

Otherwise, you can find them within the builder:

```go
es.Search().Index("index_name").Do(context.Background())
```

Alternatively each endpoint can be instantiated directly from its package:

```go
transport, err := elastictransport.New(elastictransport.Config{})
res, err = search.New(transport).Do(context.Background())
```

The `Do` method takes an optional `context`, runs the request through the transport and returns the results as well as an error.

For body-empty endpoints such as `core.Exists`, an additional method `IsSuccess` is available. As the `Do` method, it takes an optional `context`, drains and closes the body if needed, and returns a boolean alongside an error

```go
if exists, err := es.Core.Exists("index_name", "doc_id").IsSuccess(context.Background()); exists {
    // The document exists!
} else if err != nil {
    // An error occurred.
}
```

## Requests [requests]

Requests are modeled around structures that follows as closely as possible the {{es}} API and uses the standard `json/encoding` for serialization. Corresponding request can be found withing the same package as its endpoint and comes with a Builder that allows you to deep dive into the API by following the types.

```go
types.Query{
    Term: map[string]types.TermQuery{
        "name": {Value: "Foo"},
    },
}
```

## Responses [responses]

While not part of the initial release responses will be added at a later date.

## Types [types]

Requests and responses are relying on a collection of structures generated from the [elasticsearch-specification](https://github.com/elastic/elasticsearch-specification) in the `types` package. Each type comes with json tags.

## Enums [_enums]

The {{es}} API has several instances of enumerations, each has a package within `types/enums`. An enum is declared as a type and each member of the enum is an exported variable with its value. The enum types serializes to the relevant API value, for example the `refresh` options which can be found in the Search API:

```go
refresh.True => "true"
refresh.False => "false"
refresh.Waitfor => "wait_for"
```

## Unions [_unions]

To capture the expressiveness of the API union fields are represented by a type alias to an interface.
