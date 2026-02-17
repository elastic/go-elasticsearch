---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/examples.html
---

# Dense vectors and kNN search [dense_vectors]

```{applies_to}
stack: ga 9.3
```

When working with vector embeddings for semantic search or machine learning applications, the typed API provides specialized types for encoding dense vectors that can significantly improve indexing performance.

## Using DenseVectorF32 for optimized indexing [_using_densevectorf32]

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

## Creating an index with dense vector mapping [_creating_dense_vector_index]

Before indexing documents with vectors, create an index with the appropriate dense vector mapping.

The `esdsl` package provides builder helpers for mappings, queries, and aggregations:

```go
import "github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
```

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

## Searching with kNN [_knn_search]

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

## Performance benefits [_performance_benefits]

Using `types.DenseVectorF32` provides significant performance improvements over standard JSON arrays of floats:

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
