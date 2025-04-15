---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/index.html
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/overview.html
---

# Go [overview]

This is the official Go client for {{es}}.

Full documentation is hosted at [GitHub](https://github.com/elastic/go-elasticsearch) and [PkgGoDev](https://pkg.go.dev/github.com/elastic/go-elasticsearch). This documentation provides only an overview of features.


## Features [_features]

* One-to-one mapping with REST API.
* Generalized, pluggable architecture.
* Helpers for convenience.
* Rich set of examples.


## Usage [_usage]

```go
package main

import (
  "log"

  "github.com/elastic/go-elasticsearch/v9"
)

func main() {
  es, _ := elasticsearch.NewDefaultClient()
  log.Println(es.Info())
}
```

::::{note}
Please have a look at the collection of comprehensive examples in the repository at [https://github.com/elastic/go-elasticsearch/tree/master/_examples](https://github.com/elastic/go-elasticsearch/tree/master/_examples).
::::



## Resources [_resources]

* [Source Code](https://github.com/elastic/go-elasticsearch)
* [API Documentation](https://pkg.go.dev/github.com/elastic/go-elasticsearch)
* [Examples and Recipes](https://github.com/elastic/go-elasticsearch/tree/master/_examples)


## License [_license]

Copyright 2019 {{es}}.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

```
http://www.apache.org/licenses/LICENSE-2.0
```
Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

