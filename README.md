# go-elasticsearch

The official Go client for [Elasticsearch](https://www.elastic.co/products/elasticsearch).

[![GoDoc](https://godoc.org/github.com/elastic/go-elasticsearch?status.svg)](https://pkg.go.dev/github.com/elastic/go-elasticsearch/v8)
[![Go Report Card](https://goreportcard.com/badge/github.com/elastic/go-elasticsearch)](https://goreportcard.com/report/github.com/elastic/go-elasticsearch)
[![codecov.io](https://codecov.io/github/elastic/go-elasticsearch/coverage.svg?branch=main)](https://codecov.io/gh/elastic/go-elasticsearch?branch=main)
[![Build](https://github.com/elastic/go-elasticsearch/workflows/Build/badge.svg)](https://github.com/elastic/go-elasticsearch/actions?query=branch%3Amain)
[![Unit](https://github.com/elastic/go-elasticsearch/workflows/Unit/badge.svg)](https://github.com/elastic/go-elasticsearch/actions?query=branch%3Amain)
[![Integration](https://github.com/elastic/go-elasticsearch/workflows/Integration/badge.svg)](https://github.com/elastic/go-elasticsearch/actions?query=branch%3Amain)
[![API](https://github.com/elastic/go-elasticsearch/workflows/API/badge.svg)](https://github.com/elastic/go-elasticsearch/actions?query=branch%3Amain)

## Compatibility

Language clients are forward compatible; meaning that clients support communicating with greater or equal minor versions of Elasticsearch.
Elasticsearch language clients are only backwards compatible with default distributions and without guarantees made.

When using Go modules, include the version in the import path, and specify either an explicit version or a branch:

    require github.com/elastic/go-elasticsearch/v8 v8.0.0
    require github.com/elastic/go-elasticsearch/v7 7.17

It's possible to use multiple versions of the client in a single project:

    // go.mod
    github.com/elastic/go-elasticsearch/v7 v7.17.0
    github.com/elastic/go-elasticsearch/v8 v8.0.0

    // main.go
    import (
      elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
      elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
    )
    // ...
    es7, _ := elasticsearch7.NewDefaultClient()
    es8, _ := elasticsearch8.NewDefaultClient()

The `main` branch of the client is compatible with the current `master` branch of Elasticsearch.

<!-- ----------------------------------------------------------------------------------------------- -->

## Installation

Refer to the [Installation section][https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/getting-started-go.html#_installation]
of the getting started documentation.

<!-- ----------------------------------------------------------------------------------------------- -->

## Connecting

Refer to the [Connecting section](https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/getting-started-go.html#_connecting)
of the getting started documentation.

## Operations

### Creating an index

Refer to the [Creating an index section](https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/getting-started-go.html#_creating_an_index)
of the getting started documentation.

### Indexing documents

Refer to the [Indexing documents section](https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/getting-started-go.html#_indexing_documents)
of the getting started documentation.

### Getting documents

Refer to the [Getting documents section](https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/getting-started-go.html#_getting_documents)
of the getting started documentation.

### Searching documents

Refer to the [Searching documents section](https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/getting-started-go.html#_searching_documents)
of the getting started documentation.

### Updating documents

Refer to the [Updating documents section](https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/getting-started-go.html#_updating_documents)
of the getting started documentation.

### Deleting documents

Refer to the [Deleting documents section](https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/getting-started-go.html#_deleting_documents)
of the getting started documentation.

### Deleting an index

Refer to the [Deleting an index section](https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/getting-started-go.html#_deleting_an_index)
of the getting started documentation.
<!-- ----------------------------------------------------------------------------------------------- -->

## Helpers

The `esutil` package provides convenience helpers for working with the client. At the moment, it provides the `esutil.JSONReader()` and the `esutil.BulkIndexer` helpers.

<!-- ----------------------------------------------------------------------------------------------- -->

## Examples

The **[`_examples`](./_examples)** folder contains a number of recipes and comprehensive examples to get you started with the client, including configuration and customization of the client, using a custom certificate authority (CA) for security (TLS), mocking the transport for unit tests, embedding the client in a custom type, building queries, performing requests individually and in bulk, and parsing the responses.

<!-- ----------------------------------------------------------------------------------------------- -->

## License

This software is licensed under the [Apache 2 license](./LICENSE). See [NOTICE](./NOTICE).
