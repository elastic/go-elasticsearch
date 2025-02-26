---
mapped_pages:
  - https://www.elastic.co/guide/en/serverless/current/elasticsearch-go-client-getting-started.html
navigation_title: Getting started in {{serverless}}
---

# Getting started with {{es}} Go in {{serverless-full}}[elasticsearch-go-client-getting-started]

This page guides you through the installation process of the {{es}} Go client, shows you how to initialize the client, and how to perform basic {{es}} operations with it.


## Requirements [elasticsearch-go-client-getting-started-requirements]

* Go 1.22 or higher installed on your system.


## Installation [elasticsearch-go-client-getting-started-installation]


### Using the command line [elasticsearch-go-client-getting-started-using-the-command-line]

You can install the Go client with the following commands:

```bash
go get -u github.com/elastic/go-elasticsearch/v8@latest
```


## Imports [elasticsearch-go-client-getting-started-imports]

The following snippets use these imports:

```go
import (
  "context"
  "encoding/json"
  "fmt"
  "log"
  "strconv"

  "github.com/elastic/go-elasticsearch/v8"
  "github.com/elastic/go-elasticsearch/v8/typedapi/types"
  "github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/result"
)
```


## Initialize the client [elasticsearch-go-client-getting-started-initialize-the-client]

Initialize the client using your API key and {{es}} endpoint:

```go
client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
    Addresses: []string{"https://my-project-url"},
    APIKey:    "your-api-key",
})
if err != nil {
    log.Fatalf("Error creating the client: %s", err)
}
```

To get API keys for the {{es}} endpoint for a project, see [Get started](docs-content://solutions/search/get-started.md).


## Using the API [elasticsearch-go-client-getting-started-using-the-api]

After you’ve initialized the client, you can start ingesting documents. You can use the `bulk` API for this. This API enables you to index, update, and delete several documents in one request.


### Creating an index and ingesting documents [elasticsearch-go-client-getting-started-creating-an-index-and-ingesting-documents]

You can call the `bulk` API with a body parameter, an array of hashes that define the action, and a document.

The following is an example of indexing some classic books into the `books` index:

```go
type Book struct {
    Name        string `json:"name"`
    Author      string `json:"author"`
    ReleaseDate string `json:"release_date"`
    PageCount   int    `json:"page_count"`
}

books := []Book{
    {Name: "Snow Crash", Author: "Neal Stephenson", ReleaseDate: "1992-06-01", PageCount: 470},
    {Name: "Revelation Space", Author: "Alastair Reynolds", ReleaseDate: "2000-03-15", PageCount: 585},
    {Name: "1984", Author: "George Orwell", ReleaseDate: "1949-06-08", PageCount: 328},
    {Name: "Fahrenheit 451", Author: "Ray Bradbury", ReleaseDate: "1953-10-15", PageCount: 227},
    {Name: "Brave New World", Author: "Aldous Huxley", ReleaseDate: "1932-06-01", PageCount: 268},
    {Name: "The Handmaid's Tale", Author: "Margaret Atwood", ReleaseDate: "1985-06-01", PageCount: 311},
}
indexName := "books"

bulk := client.Bulk()
for i, book := range books {
    id := strconv.Itoa(i)
    err := bulk.CreateOp(types.CreateOperation{Index_: &indexName, Id_: &id}, book)
    if err != nil {
        log.Fatal(err)
    }
}
bulkRes, err := bulk.Do(context.TODO())
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Bulk: %#v\n", bulkRes.Items)
```

When you use the client to make a request to {{es-serverless}}, it returns an API response object. You can access the body values directly as seen on the previous example with `bulkRes`.


### Getting documents [elasticsearch-go-client-getting-started-getting-documents]

You can get documents by using the following code:

```go
getRes, err := client.Get(indexName, "5").Do(context.TODO())
if err != nil {
    log.Fatal(err)
}
book := Book{}
if err := json.Unmarshal(getRes.Source_, &book); err != nil {
    log.Fatal(err)
}
fmt.Printf("Get book: %#v\n", book)
```


### Searching [elasticsearch-go-client-getting-started-searching]

Now that some data is available, you can search your documents using the `search` API:

```go
searchRes, err := client.Search().
    Index("books").
    Q("snow").
    Do(context.TODO())
if err != nil {
    log.Fatal(err)
}

bookSearch := []Book{}
for _, hit := range searchRes.Hits.Hits {
    book := Book{}
    if err := json.Unmarshal(hit.Source_, &book); err != nil {
        log.Fatal(err)
    }
    bookSearch = append(bookSearch, book)
}
fmt.Printf("Search books: %#v\n", bookSearch)
```


### Updating a document [elasticsearch-go-client-getting-started-updating-a-document]

You can call the `Update` API to update a document, in this example updating the `page_count` for "The Handmaid’s Tale" with id "5":

```go
updateRes, err := client.Update("books", "5").
    Doc(
        struct {
            PageCount int `json:"page_count"`
        }{PageCount: 312},
    ).
    Do(context.TODO())
if err != nil {
    log.Fatal(err)
}

if updateRes.Result == result.Updated {
    fmt.Printf("Update book: %#v\n", updateRes)
}
```


### Deleting a document [elasticsearch-go-client-getting-started-deleting-a-document]

You can call the `Delete` API to delete a document:

```go
deleteRes, err := client.Delete("books", "5").Do(context.TODO())
if err != nil {
    log.Fatal(err)
}

if deleteRes.Result == result.Deleted {
    fmt.Printf("Delete book: %#v\n", deleteRes)
}
```


### Deleting an index [elasticsearch-go-client-getting-started-deleting-an-index]

```go
indexDeleteRes, err := client.Indices.Delete("books").Do(context.TODO())
if err != nil {
    log.Fatal(err)
}

if indexDeleteRes.Acknowledged {
    fmt.Printf("Delete index: %#v\n", indexDeleteRes)
}
```

