---
navigation_title: "Using ES|QL"
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/esql.html
---

# ES|QL in the Go client [esql]

This page helps you understand and use [ES|QL](docs-content://explore-analyze/query-filter/languages/esql.md) in the Go client.

There are two ways to use ES|QL in the Go client:

- Use the Elasticsearch [ES|QL API](https://www.elastic.co/docs/api/doc/elasticsearch/group/endpoint-esql) directly: This is the most flexible approach, but it's also the most complex because you must handle results in their raw form. You can choose the precise format of results, such as JSON, CSV, or text.
- Use ES|QL mapping helpers: These mappers take care of parsing the raw response into something readily usable by the application. Helpers are available for object mapping and iterative objects.

## How to use the ES|QL API [esql-how-to]

The [ES|QL query API](https://www.elastic.co/docs/api/doc/elasticsearch/group/endpoint-esql) allows you to specify how results should be returned. You can choose a [response format](docs-content://explore-analyze/query-filter/languages/esql-rest.md#esql-rest-format) such as CSV, text, or JSON, then fine-tune it with parameters like column separators and locale.

The following example gets ES|QL results as CSV and parses them:

```go
queryAuthor := `from library
    | where author == "Isaac Asimov"
    | sort release_date desc
    | limit 10`

response, err := client.Esql.Query().
    Query(queryAuthor).
    Format("csv").
    Do(context.Background())
if err != nil {
    log.Fatal(err)
}

reader := csv.NewReader(bytes.NewReader(response))
rows, err := reader.ReadAll()
for _, row := range rows {
    fmt.Println(row)
}
```

## Consume ES|QL results [esql-consume-results]

The previous example showed that although the raw ES|QL API offers maximum flexibility, additional work is required in order to make use of the result data.

To simplify things, try working with these two main representations of ES|QL results (each with its own mapping helper):

- **Objects**, where each row in the results is mapped to an object from your application domain. This is similar to what ORMs (object relational mappers) commonly do.

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/elastic/go-elasticsearch/v9"
    "github.com/elastic/go-elasticsearch/v9/typedapi/esql/query"
)

type Book struct {
    Name        string `json:"name"`
    Author      string `json:"author"`
    ReleaseDate string `json:"release_date"`
    PageCount   int    `json:"page_count"`
}

func main() {
    client, err := elasticsearch.NewTypedClient(elasticsearch.Config{Addresses: []string{"ELASTICSEARCH_URL"}})
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        if err := client.Close(ctx); err != nil {
            log.Fatal(err)
        }
    } ()

    queryAuthor := `from library
        | where author == "Isaac Asimov"
        | sort release_date desc
        | limit 10`

    qry := client.Esql.Query().Query(queryAuthor)
    books, err := query.Helper[Book](context.Background(), qry)
    if err != nil {
        log.Fatal(err)
    }

    for _, book := range books {
        fmt.Println(book)
    }
}
```

- **Iterative Objects**, where each row in the results is mapped to an object from your application domain, one at a time.

```go
queryAuthor := `from library
    | where author == "Isaac Asimov"
    | sort release_date desc
    | limit 10`

qry := client.Esql.Query().Query(queryAuthor)
books, err := query.NewIteratorHelper[Book](context.Background(), qry)
if err != nil {
    log.Fatal(err)
}

for books.More() {
    book, err := books.Next()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(book)
}
```
