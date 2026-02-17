---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
---

# Dense vectors and kNN search [dense_vectors]

```{applies_to}
stack: ga 9.3
```

When working with vector embeddings for semantic search or machine learning applications, the Go client provides support for dense vector indexing and k-nearest neighbors (kNN) search. For the full kNN search reference, see [kNN search](docs-content://solutions/search/vector/knn.md) in the {{es}} documentation.

## Creating an index with dense vector mapping [_creating_dense_vector_index]

Before indexing documents with vectors, create an index with the appropriate dense vector mapping.

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

```go
mapping := `{
  "mappings": {
    "properties": {
      "docid":  { "type": "keyword" },
      "title":  { "type": "text" },
      "emb": {
        "type": "dense_vector",
        "dims": 1536,
        "index": true,
        "similarity": "cosine"
      }
    }
  }
}`

res, err := client.Indices.Create(
    "my-vectors", // <1>
    client.Indices.Create.WithBody(strings.NewReader(mapping)), // <2>
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

1. The name of the index to create.
2. Pass the full mapping as a JSON body, including dense vector configuration.

::::::

::::::{tab-item} Fully-typed API
:sync: typed

```go
similarity := densevectorsimilarity.Cosine

res, err := es.Indices.
    Create("my-vectors").
    Request(&create.Request{
        Mappings: &types.TypeMapping{
            Properties: map[string]types.Property{
                "docid": types.NewKeywordProperty(),
                "title": types.NewTextProperty(),
                "emb": types.DenseVectorProperty{ // <1>
                    Dims:       some.Int(1536), // <2>
                    Index:      some.Bool(true), // <3>
                    Similarity: &similarity, // <4>
                },
            },
        },
    }).
    Do(context.Background())
```

1. Define the vector field as a `DenseVectorProperty`.
2. Specify the dimensionality of your vectors (e.g., 1536 for OpenAI embeddings).
3. Enable indexing to support kNN search capabilities.
4. Set the similarity metric: `Cosine`, `DotProduct`, or `L2Norm`.

::::::

::::::{tab-item} esdsl API
:sync: esdsl

The [`esdsl`](/reference/typed-api/esdsl.md) mapping builders provide a concise, fluent syntax:

```go subs=true
import "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}/typedapi/esdsl"
```

```go
mappings := esdsl.NewTypeMapping().
    AddProperty("docid", esdsl.NewKeywordProperty()).
    AddProperty("title", esdsl.NewTextProperty()).
    AddProperty("emb", esdsl.NewDenseVectorProperty(). // <1>
        Dims(1536). // <2>
        Index(true). // <3>
        Similarity(densevectorsimilarity.Cosine)) // <4>

res, err := es.Indices.
    Create("my-vectors").
    Mappings(mappings). // <5>
    Do(context.Background())
```

1. Define the vector field as a `DenseVectorProperty`.
2. Specify the dimensionality of your vectors (e.g., 1536 for OpenAI embeddings).
3. Enable indexing to support kNN search capabilities.
4. Set the similarity metric: `Cosine`, `DotProduct`, or `L2Norm`.
5. Pass the mapping builder directly - no `Request` wrapper needed.

::::::

:::::::

## Indexing documents with vectors [_using_densevectorf32]

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

```go
doc := `{
    "docid": "doc1",
    "title": "Example document with vector embedding",
    "emb": [0.1, 0.2, 0.3, 0.4, 0.5]
}`

res, err := client.Index(
    "my-vectors", // <1>
    strings.NewReader(doc), // <2>
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

1. The target index.
2. Vectors are passed as JSON arrays of floats.

::::::

::::::{tab-item} Fully-typed API
:sync: typed

The `types.DenseVectorF32` type automatically encodes `[]float32` vectors as base64 strings during JSON serialization, reducing payload size and improving indexing speed:

```go
type Document struct {
    DocID string               `json:"docid"`
    Title string               `json:"title"`
    Emb   types.DenseVectorF32 `json:"emb"` // <1>
}

document := Document{
    DocID: "doc1",
    Title: "Example document with vector embedding",
    Emb:   types.DenseVectorF32{0.1, 0.2, 0.3, 0.4, 0.5}, // <2>
}

res, err := es.Index("my-vectors").
    Request(document).
    Do(context.Background())
```

1. Use `types.DenseVectorF32` instead of `[]float32` for vector fields.
2. Assign float32 slices directly; base64 encoding happens automatically during serialization.

::::::

:::::::

## Searching with kNN [_knn_search]

Once your vectors are indexed, you can perform k-nearest neighbors (kNN) search to find similar documents:

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

```go
query := `{
  "query": {
    "knn": {
      "field": "emb",
      "query_vector": [0.1, 0.2, 0.3, 0.4, 0.5],
      "k": 10,
      "num_candidates": 100
    }
  }
}`

res, err := client.Search(
    client.Search.WithIndex("my-vectors"), // <1>
    client.Search.WithBody(strings.NewReader(query)), // <2>
)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
```

1. The index containing your vectors.
2. Pass the kNN query as raw JSON.

::::::

::::::{tab-item} Fully-typed API
:sync: typed

```go
queryVector := []float32{0.1, 0.2, 0.3, 0.4, 0.5} // <1>
k := 10
numCandidates := 100

res, err := es.Search().
    Index("my-vectors").
    Request(&search.Request{
        Query: &types.Query{
            Knn: &types.KnnQuery{ // <2>
                Field:         "emb", // <3>
                QueryVector:   queryVector, // <4>
                K:             &k, // <5>
                NumCandidates: &numCandidates, // <6>
            },
        },
    }).
    Do(context.Background())
```

1. Define your query vector (typically from the same embedding model).
2. Build the kNN query using the `KnnQuery` struct.
3. Specify which vector field to search.
4. Provide the query vector to find similar vectors.
5. Return the top 10 nearest neighbors.
6. Consider 100 candidates during the search for better accuracy.

::::::

::::::{tab-item} esdsl API
:sync: esdsl

The [`esdsl`](/reference/typed-api/esdsl.md) query builders provide a fluent syntax for kNN search:

```go
queryVector := []float32{0.1, 0.2, 0.3, 0.4, 0.5} // <1>

res, err := es.Search().
    Index("my-vectors").
    Query(
        esdsl.NewKnnQuery(). // <2>
            Field("emb"). // <3>
            QueryVector(queryVector...). // <4>
            K(10). // <5>
            NumCandidates(100), // <6>
    ).
    Do(context.Background())
```

1. Define your query vector (typically from the same embedding model).
2. Use `esdsl.NewKnnQuery()` to build a kNN query with a fluent chain.
3. Specify which vector field to search.
4. Provide the query vector to find similar vectors.
5. Return the top 10 nearest neighbors.
6. Consider 100 candidates during the search for better accuracy.

::::::

:::::::

## Performance benefits [_performance_benefits]

Using `types.DenseVectorF32` (typed API) provides significant performance improvements over standard JSON arrays of floats:

- **Reduced payload size**: base64 encoding is more compact than JSON number arrays
- **Faster parsing**: Eliminates JSON number parsing overhead
- **Improved indexing speed**: Performance gains increase with vector dimensionality and can improve indexing speeds by up to 3x

::::{note}
For best performance, use `types.DenseVectorF32` when your vectors are already in `[]float32` format. If you have pre-encoded bytes, use `types.DenseVectorBytes` to avoid re-encoding.
::::

## Using DenseVectorBytes [_using_densevectorbytes]

If you already have pre-encoded vector bytes from another system, use `types.DenseVectorBytes`:

```go
type Document struct {
    Emb types.DenseVectorBytes `json:"emb"` // <1>
}

vectorBytes := []byte{...} // <2>
document := Document{
    Emb: types.DenseVectorBytes(vectorBytes),
}

res, err := es.Index("my-vectors").
    Request(document).
    Do(context.Background())
```

1. Use `types.DenseVectorBytes` when you have pre-encoded bytes.
2. Provide the raw byte representation of your vector data.
