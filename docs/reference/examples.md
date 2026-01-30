---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
---

# Examples [examples]

This sections lists a series of frequent use cases that will help you start with this new API.

- [Creating an index](#indices)
- [Indexing a document](#indexing)
- [Retrieving a document](#retrieving_document)
- [Search](#search)
- [Aggregations](#aggregations)
- [Indexing dense vectors](#dense_vectors)

::::{note}
This is a work in progress, the documentation will evolve in the future.
::::

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

## Search [search]

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

## Aggregations [aggregations]

Given documents with a `price` field, we run a sum aggregation on `index_name`:

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

## Indexing dense vectors [dense_vectors]

```{applies_to}
stack: ga 9.3
```

When working with vector embeddings for semantic search or machine learning applications, the typed API provides specialized types for encoding dense vectors that can significantly improve indexing performance.

### Using DenseVectorF32 for optimized indexing [_using_densevectorf32]

The `types.DenseVectorF32` type automatically encodes `[]float32` vectors as base64 strings during JSON serialization, reducing payload size and improving indexing speed:

```go
type Document struct {
    DocID string               `json:"docid"`
    Title string               `json:"title"`
    Emb   types.DenseVectorF32 `json:"emb"` <1>
}

document := Document{
    DocID: "doc1",
    Title: "Example document with vector embedding",
    Emb:   types.DenseVectorF32{0.1, 0.2, 0.3, 0.4, 0.5}, <2>
}

res, err := es.Index("my-vectors").
    Request(document).
    Do(context.Background())
```

1. Use `types.DenseVectorF32` instead of `[]float32` for vector fields.
2. Assign float32 slices directly; base64 encoding happens automatically during serialization.

### Creating an index with dense vector mapping [_creating_dense_vector_index]

Before indexing documents with vectors, create an index with the appropriate dense vector mapping:

```go
mappings := esdsl.NewTypeMapping().
    AddProperty("docid", esdsl.NewKeywordProperty()).
    AddProperty("title", esdsl.NewTextProperty()).
    AddProperty("emb", esdsl.NewDenseVectorProperty(). <1>
        Dims(1536). <2>
        Index(true). <3>
        Similarity(densevectorsimilarity.Cosine)) <4>

res, err := es.Indices.
    Create("my-vectors").
    Mappings(mappings).
    Do(context.Background())
```

1. Define the vector field as a `DenseVectorProperty`.
2. Specify the dimensionality of your vectors (e.g., 1536 for OpenAI embeddings).
3. Enable indexing to support kNN search capabilities.
4. Set the similarity metric: `Cosine`, `DotProduct`, or `L2Norm`.

### Searching with kNN [_knn_search]

Once your vectors are indexed, you can perform k-nearest neighbors (kNN) search to find similar documents:

```go
queryVector := []float32{0.1, 0.2, 0.3, 0.4, 0.5} <1>

res, err := es.Search().
    Index("my-vectors").
    Request(&search.Request{
        Query: esdsl.NewKnnQuery(). <2>
            Field("emb"). <3>
            QueryVector(queryVector...). <4>
            K(10). <5>
            NumCandidates(100). <6>
            QueryCaster(),
    }).
    Do(context.Background())
```

1. Define your query vector (typically from the same embedding model).
2. Use `esdsl.NewKnnQuery()` to build a kNN query.
3. Specify which vector field to search.
4. Provide the query vector to find similar vectors.
5. Return the top 10 nearest neighbors.
6. Consider 100 candidates during the search for better accuracy.

### Performance benefits [_performance_benefits]

Using `types.DenseVectorF32` provides significant performance improvements over standard JSON arrays of floats:

- **Reduced payload size**: base64 encoding is more compact than JSON number arrays
- **Faster parsing**: Eliminates JSON number parsing overhead
- **Improved indexing speed**: Performance gains increase with vector dimensionality and can improve indexing speeds by up to 3x

::::{note}
For best performance, use `types.DenseVectorF32` when your vectors are already in `[]float32` format. If you have pre-encoded bytes, use `types.DenseVectorBytes` to avoid re-encoding.
::::

### Using DenseVectorBytes [_using_densevectorbytes]

If you already have pre-encoded vector bytes from another system, use `types.DenseVectorBytes`:

```go
type Document struct {
    Emb types.DenseVectorBytes `json:"emb"` <1>
}

vectorBytes := []byte{...} <2>
document := Document{
    Emb: types.DenseVectorBytes(vectorBytes),
}

res, err := es.Index("my-vectors").
    Request(document).
    Do(context.Background())
```

1. Use `types.DenseVectorBytes` when you have pre-encoded bytes.
2. Provide the raw byte representation of your vector data.
