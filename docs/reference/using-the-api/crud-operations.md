---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
---

# CRUD operations [crud_operations]

This page covers common CRUD (Create, Read, Update, Delete) operations with the Go client using the [typed API](/reference/typed-api/index.md) and the [`esdsl`](/reference/typed-api/esdsl.md) builders. For the low-level equivalents, see the [low-level API](/reference/low-level-api/index.md).

```go subs=true
import "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}/typedapi/esdsl"
```

## Creating an index [indices]

For this example, let's create an index named `test-index` and provide a mapping for the field `price` which will be an integer.

The [`esdsl`](/reference/typed-api/esdsl.md) mapping builders provide a fluent syntax for defining index mappings. The correct `type` field is set automatically for each property:

```go
res, err := es.Indices.Create("test-index"). // <1>
    Mappings(
        esdsl.NewTypeMapping(). // <2>
            AddProperty("price", esdsl.NewIntegerNumberProperty()), // <3>
    ).
    Do(context.Background())
```

1. The name of the index to create.
2. Start a mapping builder with `NewTypeMapping()`.
3. Chain `AddProperty()` calls to define each field.

## Indexing a document [indexing]

The standard way of indexing a document is to provide the document body to the index request. The document will be serialized as JSON and sent to {{es}}.

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

res, err := es.Index("index_name"). // <1>
    Id("1").                         // <2>
    Document(document).              // <3>
    Do(context.Background())
```

1. The target index name.
2. Optionally set the document ID.
3. Pass the struct directly; serialization is handled automatically.

Alternatively, you can use the `Raw` method and provide already serialized JSON:

```go
res, err := es.Index("index_name").
    Raw([]byte(`{
      "id": 1,
      "name": "Foo",
      "price": 10
    }`)).Do(context.Background())
```

## Retrieving a document [retrieving_document]

Retrieving a document follows the API as part of the argument of the endpoint. You provide the `index` and the `id`, then run the query.

```go
res, err := es.Get(
    "index_name", // <1>
    "doc_id",     // <2>
).Do(context.Background())
```

1. The index to retrieve the document from.
2. The document ID.

## Checking for a document existence [_checking_for_a_document_existence]

If you do not wish to retrieve the content of the document and want only to check if it exists in your index, use the `IsSuccess` shortcut:

```go
if exists, err := es.Exists("index_name", "doc_id").IsSuccess(nil); exists {
    // The document exists!
} else if err != nil {
   // Handle error.
}
```

Result is `true` if everything succeeds, `false` if the document doesn't exist. If an error occurs during the request, you will get `false` and the relevant error.
