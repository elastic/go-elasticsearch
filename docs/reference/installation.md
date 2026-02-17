---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/installation.html
---

# Installation [installation]

To install the {{ version.elasticsearch-client-go | M.x }} version of the client for Go, you can use the `go get` command:

```shell subs=true
go get github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}@v{{ version.elasticsearch-client-go }}
```

Or, add the package to your `go.mod` file:

```text subs=true
require github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }} {{ version.elasticsearch-client-go | M.M }}
```

Or, clone the repository:

```text subs=true
git clone --branch {{ version.elasticsearch-client-go | M.M }} https://github.com/elastic/go-elasticsearch.git $GOPATH/src/github.com/elastic/go-elasticsearch
```

To install another version, modify the path or the branch name accordingly. The client major versions correspond to the {{es}} major versions.

You can find a complete example of installation below:

```shell subs=true
mkdir my-elasticsearch-app && cd my-elasticsearch-app && go mod init my-elasticsearch-app

go get github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}@v{{ version.elasticsearch-client-go }}

cat > main.go <<-END
  package main

  import (
    "context"
    "log"

    "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}"
  )

  func main() {
    es, _ := elasticsearch.NewDefaultClient()
    defer es.Close(context.Background())
    log.Println(elasticsearch.Version)
    log.Println(es.Info())
  }
END

go run main.go
```

## {{es}} version compatibility [_es_version_compatibility]

The language clients are forward compatible; meaning that the clients support communicating with greater or equal minor versions of {{es}} without breaking.
It does not mean that the clients automatically support new features of newer {{es}} versions; it is only possible after a release of a new client version.
For example, a 8.12 client version won't automatically support the new features of the 8.13 version of {{es}}, the 8.13 client version is required for that.
{{es}} language clients are only backward compatible with default distributions and without guarantees made.

| Elasticsearch Version | Elasticsearch-Go Branch | Supported |
| --------------------- | ----------------------- | --------- |
| main                  | main                    |           |
| 9.x                   | 9.x                     | 9.x       |
| 8.x                   | 8.x                     | 8.x       |
