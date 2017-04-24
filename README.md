# goelasticsearch

This repository contains the Go client library for [Elasticsearch](https://www.elastic.co/products/elasticsearch).

## Installation

Install the `client` package from [GitHub](https://github.com/elastic/goelasticsearch):

    go get github.com/elastic/goelasticsearch/client

or install it from a source code checkout:

    git clone https://github.com/elastic/goelasticsearch.git
    go install ./client

## Usage

```go
import "github.com/elastic/goelasticsearch/client"

c := client.New()
```

## License

This software is licensed under the Apache 2 license, quoted below.

    Copyright (c) 2017 Elasticsearch <http://www.elasticsearch.org>

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
