---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
---

# CRUD operations [crud_operations]

This page covers common CRUD (Create, Read, Update, Delete) operations with the Go client.

## Creating an index [indices]

For this example on how to create an index, lets create an index named `test-index` and provide a mapping for the field `price` which will be an integer.

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

```go
mapping := `{
  "mappings": {
    "properties": {
      "price": { "type": "integer" }
    }
  }
}`

res, err := client.Indices.Create(
    "test-index", // <1>
    client.Indices.Create.WithBody(strings.NewReader(mapping)), // <2>
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close() // <3>
```

1. The name of the index to create.
2. Pass the JSON mapping as the request body.
3. Always close the response body to ensure HTTP connection reuse.

::::::

::::::{tab-item} Fully-typed API
:sync: typed

Notice how using the builder for the `IntegerNumberProperty` will automatically apply the correct value for the `type` field.

```go
res, err := es.Indices.Create("test-index"). // <1>
    Request(&create.Request{
        Mappings: &types.TypeMapping{
            Properties: map[string]types.Property{
                "price": types.NewIntegerNumberProperty(), // <2>
            },
        },
    }).
    Do(context.Background())
```

1. The name of the index to create.
2. The typed builder automatically sets the correct `type` field for the property.

::::::

::::::{tab-item} esdsl API
:sync: esdsl

The [`esdsl`](/reference/typed-api/esdsl.md) mapping builders provide a fluent syntax for defining index mappings:

```go subs=true
import "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}/typedapi/esdsl"
```

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
3. Chain `AddProperty()` calls to define each field. The correct `type` is set automatically.

::::::

:::::::

## Indexing a document [indexing]

The standard way of indexing a document is to provide the document body to the index request. The document will be serialized as JSON and sent to {{es}}.

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

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

data, _ := json.Marshal(document)
res, err := client.Index(
    "index_name",              // <1>
    bytes.NewReader(data),     // <2>
    client.Index.WithDocumentID("1"), // <3>
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close() // <4>
```

1. The target index name.
2. The document body as an `io.Reader`.
3. Optionally set the document ID.
4. Always close the response body to ensure HTTP connection reuse.

::::::

::::::{tab-item} Fully-typed API
:sync: typed

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
    Request(document).               // <3>
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

::::::

:::::::

## Retrieving a document [retrieving_document]

Retrieving a document follows the API as part of the argument of the endpoint. You provide the `index` and the `id`, then run the query.

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

```go
res, err := client.Get(
    "index_name", // <1>
    "doc_id",     // <2>
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close() // <3>
```

1. The index to retrieve the document from.
2. The document ID.
3. Always close the response body to ensure HTTP connection reuse.

::::::

::::::{tab-item} Fully-typed API
:sync: typed

```go
res, err := es.Get(
    "index_name", // <1>
    "doc_id",     // <2>
).Do(context.Background())
```

1. The index to retrieve the document from.
2. The document ID.

::::::

:::::::

## Checking for a document existence [_checking_for_a_document_existence]

If you do not wish to retrieve the content of the document and want only to check if it exists in your index:

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

```go
res, err := client.Exists("index_name", "doc_id")
if err != nil {
    // Handle error.
}
defer res.Body.Close()

if !res.IsError() {
    // The document exists!
}
```

::::::

::::::{tab-item} Fully-typed API
:sync: typed

The typed API provides the `IsSuccess` shortcut:

```go
if exists, err := es.Exists("index_name", "doc_id").IsSuccess(nil); exists {
    // The document exists!
} else if err != nil {
   // Handle error.
}
```

Result is `true` if everything succeeds, `false` if the document doesn't exist. If an error occurs during the request, you will get `false` and the relevant error.

::::::

:::::::
