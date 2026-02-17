---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
---

# CRUD operations [crud_operations]

This page covers common CRUD (Create, Read, Update, Delete) operations using the typed API.

## Creating an index [indices]

For this example on how to create an index, lets create an index named `test-index` and provide a mapping for the field `price` which will be an integer. Notice how using the builder for the `IntegerNumberProperty` will automatically apply the correct value for the `type` field.

```go
res, err := es.Indices.Create("test-index").
    Request(&create.Request{
        Mappings: &types.TypeMapping{
            Properties: map[string]types.Property{
                "price": types.NewIntegerNumberProperty(),
            },
        },
    }).
    Do(nil)
```

## Indexing a document [indexing]

The standard way of indexing a document is to provide a `struct` to the `Request` method, the standard `json/encoder` will be run on your structure and the result will be sent to {{es}}.

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

res, err := es.Index("index_name").
    Request(document).
    Do(context.Background())
```

Alternatively, you can use the `Raw` method and provide the already serialized JSON:

```go
res, err := es.Index("index_name").
            Raw([]byte(`{
              "id": 1,
              "name": "Foo",
              "price": 10
            }`)).Do(context.Background())
```

## Retrieving a document [retrieving_document]

Retrieving a document follows the API as part of the argument of the endpoint. In order you provide the `index`, the `id` and then run the query:

```go
res, err := es.Get("index_name", "doc_id").Do(context.Background())
```

## Checking for a document existence [_checking_for_a_document_existence]

If you do not wish to retrieve the content of the document and want only to check if it exists in your index, we provide the `IsSuccess` shortcut:

```go
if exists, err := es.Exists("index_name", "doc_id").IsSuccess(nil); exists {
    // The document exists !
} else if err != nil {
   // Handle error.
}
```

Result is `true` if everything succeeds, `false` if the document doesn't exist. If an error occurs during the request, you will be granted with a `false` and the relevant error.
