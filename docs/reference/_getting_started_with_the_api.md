---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/_getting_started_with_the_api.html
---

# Getting started with the API [_getting_started_with_the_api]

The new typed client can be obtained from the `elasticsearch` package using the `NewTypedClient` function. This new API takes the same arguments as the previous one and uses the same transport underneath.

Connection to an {{es}} cluster is identical to the existing client, only the API changes:

```go
client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
	// Proper configuration for your Elasticsearch cluster.
})
```

The code is generated from the [elasticsearch-specification](https://github.com/elastic/elasticsearch-specification).

