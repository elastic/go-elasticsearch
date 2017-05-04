# go-elasticsearch

# WORK IN PROGRESS

This is the official Go client for library [Elasticsearch](https://www.elastic.co/products/elasticsearch), it's currently a work in progress so please check back later.

## Installation

Install the `client` package from [GitHub](https://github.com/elastic/go-elasticsearch):

    go get github.com/elastic/go-elasticsearch/client

or install it from a source code checkout:

    git clone https://github.com/elastic/go-elasticsearch.git
    go install ./client

## Usage

```go
import "github.com/elastic/go-elasticsearch/client"

c, _ = client.New(client.WithHosts([]string{"https://elasticseach:9200"}))
body := map[string]interface{}{
  "query": map[string]interface{}{
    "term": map[string]interface{}{
      "user": "kimchy",
    },
  },
}
resp, err := c.Search(body)
```