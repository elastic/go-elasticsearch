---
navigation_title: "Using the API"
---

# Using the API [using_the_api]

The Go client for {{es}} offers two API styles:

- **Low-level API** — provides a one-to-one mapping with the {{es}} REST API. Each endpoint accepts and returns raw `io.Reader` bodies and `*esapi.Response` objects. This gives you full control over serialization and request construction.
- **Fully-typed API** — provides a strongly typed, builder-pattern API generated from the [elasticsearch-specification](https://github.com/elastic/elasticsearch-specification). Requests and responses are modeled as Go structs, giving you compile-time safety and IDE autocompletion.

Both API styles share the same underlying transport and configuration. You can even use both in the same application.

## What's covered in this section

- [CRUD operations](crud-operations.md) — Index, get, update, and delete documents
- [Searching](searching.md) — Build and run search queries
- [Aggregations](aggregations.md) — Run aggregations on your data
- [Bulk indexing](bulk-indexing.md) — Efficiently ingest large volumes of documents
- [ES|QL](esql.md) — Use the Elasticsearch Query Language
- [Dense vectors and kNN search](dense-vectors.md) — Work with vector embeddings and similarity search
