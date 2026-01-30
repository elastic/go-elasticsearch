---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/installation.html
---

# Installation [installation]

To install the 8.x version of the client, add the package to your `go.mod` file:

```text
require github.com/elastic/go-elasticsearch/v9 9.0
```

Or, clone the repository:

```text
git clone --branch 9.0 https://github.com/elastic/go-elasticsearch.git $GOPATH/src/github
```

To install another version, modify the path or the branch name accordingly. The client major versions correspond to the {{es}} major versions.

You can find a complete example of installation below:

```text
mkdir my-elasticsearch-app && cd my-elasticsearch-app

cat > go.mod <<-END
  module my-elasticsearch-app

  require github.com/elastic/go-elasticsearch/v9 main
END

cat > main.go <<-END
  package main

  import (
    "context"
    "log"

    "github.com/elastic/go-elasticsearch/v9"
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
