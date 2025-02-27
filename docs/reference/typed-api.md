---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/typedapi.html
---

# Typed API [typedapi]

The goal for this API is to provide a strongly typed Golang API for {{es}}.

This was designed with structures and the Golang runtime in mind, following as closely as possible the API and its objects.

The first version focuses on the requests and does not yet include NDJSON endpoints such as `bulk` or `msearch`. These will be added later on along with typed responses and error handling.






