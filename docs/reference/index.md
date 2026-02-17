---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/index.html
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/overview.html
---

# Go [overview]

This is the official Go client for {{es}}.

Full documentation is hosted at [GitHub](https://github.com/elastic/go-elasticsearch) and [PkgGoDev](https://pkg.go.dev/github.com/elastic/go-elasticsearch). This documentation provides only an overview of features.

## Features [_features]

- One-to-one mapping with REST API.
- Generalized, pluggable architecture.
- Helpers for convenience.
- Rich set of examples.

## Usage [_usage]

:::::::{tab-set}
:group: APIs
::::::{tab-item} Low-level API
:sync: lowLevel

```go subs=true
package main

import (
	"context"
	"log"

    "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}"
)

func main() {
    es, _ := elasticsearch.NewDefaultClient()
    defer es.Close(context.Background())
    log.Println(es.Info())
}
```

::::::

::::::{tab-item} Fully-typed API
:sync: typed

```go subs=true
package main

import (
	"context"
    "log"

    "github.com/elastic/go-elasticsearch/v{{ version.elasticsearch-client-go | M }}"
)

func main() {
    es, _ := elasticsearch.NewTypedClient(elasticsearch.Config{
        Addresses: []string{"http://localhost:9200"},
    })
    defer es.Close(context.Background())
    log.Println(es.Info().Do(context.Background()))
}
```

::::::

:::::::

::::{note}
Please have a look at the collection of comprehensive examples in the repository at <https://github.com/elastic/go-elasticsearch/tree/main/_examples>.
::::

## Resources [_resources]

- [Source Code](https://github.com/elastic/go-elasticsearch)
- [API Documentation](https://pkg.go.dev/github.com/elastic/go-elasticsearch)
- [Examples and Recipes](https://github.com/elastic/go-elasticsearch/tree/master/_examples)

## License [_license]

Copyright 2019 {{es}}.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

```text
http://www.apache.org/licenses/LICENSE-2.0
```

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
