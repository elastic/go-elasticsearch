# go-elasticsearch

The official Go client for [Elasticsearch](https://www.elastic.co/products/elasticsearch).

**[Download the latest version of Elasticsearch](https://www.elastic.co/downloads/elasticsearch)**
or
**[sign-up](https://cloud.elastic.co/registration?elektra=en-ess-sign-up-page)**
**for a free trial of Elastic Cloud**.

[![Go Reference](https://pkg.go.dev/badge/github.com/elastic/go-elasticsearch/v9.svg)](https://pkg.go.dev/github.com/elastic/go-elasticsearch/v9)
[![Go Report Card](https://goreportcard.com/badge/github.com/elastic/go-elasticsearch)](https://goreportcard.com/report/github.com/elastic/go-elasticsearch)
[![codecov.io](https://codecov.io/github/elastic/go-elasticsearch/coverage.svg?branch=main)](https://codecov.io/gh/elastic/go-elasticsearch?branch=main)
[![Build](https://github.com/elastic/go-elasticsearch/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/elastic/go-elasticsearch/actions/workflows/build.yml?query=branch%3Amain)
[![Unit](https://github.com/elastic/go-elasticsearch/actions/workflows/test-unit.yml/badge.svg?branch=main)](https://github.com/elastic/go-elasticsearch/actions/workflows/test-unit.yml?query=branch%3Amain)
[![Integration](https://github.com/elastic/go-elasticsearch/actions/workflows/test-integration.yml/badge.svg?branch=main)](https://github.com/elastic/go-elasticsearch/actions/workflows/test-integration.yml?query=branch%3Amain)
[![API](https://github.com/elastic/go-elasticsearch/actions/workflows/test-api.yml/badge.svg?branch=main)](https://github.com/elastic/go-elasticsearch/actions/workflows/test-api.yml?query=branch%3Amain)
[![Lint](https://github.com/elastic/go-elasticsearch/actions/workflows/lint.yml/badge.svg?branch=main&event=push)](https://github.com/elastic/go-elasticsearch/actions/workflows/lint.yml)

## Compatibility

### Go

Starting from version `8.12.0`, this library follows the Go language's [release policy](https://go.dev/doc/devel/release#policy). Each major Go release is supported until there are two newer major releases. For example, Go 1.5 was supported until the Go 1.7 release, and Go 1.6 was supported until the Go 1.8 release.

### Elasticsearch

Language clients are forward compatible; meaning that clients support communicating with greater or equal minor versions of Elasticsearch.
Elasticsearch language clients are only backward compatible with default distributions and without guarantees made.

When using Go modules, include the version in the import path, and specify either an explicit version or a branch:

    require github.com/elastic/go-elasticsearch/v9 v9.x.x
    require github.com/elastic/go-elasticsearch/v8 v8.x.x

It's possible to use multiple versions of the client in a single project:

    // go.mod
    github.com/elastic/go-elasticsearch/v8 v8.18.0
    github.com/elastic/go-elasticsearch/v9 v9.0.0

    // main.go
    import (
      elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
      elasticsearch9 "github.com/elastic/go-elasticsearch/v9"
    )
    // ...
    es8, _ := elasticsearch8.NewDefaultClient()
    es9, _ := elasticsearch9.NewDefaultClient()

The `main` branch of the client is compatible with the current `master` branch of Elasticsearch.

<!-- ----------------------------------------------------------------------------------------------- -->

## Installation

Refer to the [Installation](https://www.elastic.co/docs/reference/elasticsearch/clients/go/installation)
page of the documentation.

<!-- ----------------------------------------------------------------------------------------------- -->

## Connecting

Refer to the [Connecting](https://www.elastic.co/docs/reference/elasticsearch/clients/go/connecting)
page of the documentation.

## Operations

- [Creating an index](https://www.elastic.co/docs/reference/elasticsearch/clients/go/using-the-api/crud-operations)
- [Indexing documents](https://www.elastic.co/docs/reference/elasticsearch/clients/go/using-the-api/crud-operations)
- [Getting documents](https://www.elastic.co/docs/reference/elasticsearch/clients/go/using-the-api/crud-operations)
- [Searching documents](https://www.elastic.co/docs/reference/elasticsearch/clients/go/using-the-api/searching)
- [Updating documents](https://www.elastic.co/docs/reference/elasticsearch/clients/go/using-the-api/crud-operations)
- [Deleting documents](https://www.elastic.co/docs/reference/elasticsearch/clients/go/using-the-api/crud-operations)
- [Deleting an index](https://www.elastic.co/docs/reference/elasticsearch/clients/go/using-the-api/crud-operations)
<!-- ----------------------------------------------------------------------------------------------- -->

## Helpers

The `esutil` package provides convenience helpers for working with the client. At the moment, it provides the `esutil.JSONReader()` and the [`esutil.BulkIndexer`](https://www.elastic.co/docs/reference/elasticsearch/clients/go/using-the-api/bulk-indexing) helpers.

<!-- ----------------------------------------------------------------------------------------------- -->

## Documentation

Find the full reference documentation at [elastic.co](https://www.elastic.co/docs/reference/elasticsearch/clients/go):

- [Getting started](https://www.elastic.co/docs/reference/elasticsearch/clients/go/getting-started)
- [Configuration](https://www.elastic.co/docs/reference/elasticsearch/clients/go/configuration)
- [Using the API](https://www.elastic.co/docs/reference/elasticsearch/clients/go/using-the-api) (low-level and typed)
- [Typed API](https://www.elastic.co/docs/reference/elasticsearch/clients/go/typed-api) and [esdsl builders](https://www.elastic.co/docs/reference/elasticsearch/clients/go/typed-api/esdsl)
- [Advanced topics](https://www.elastic.co/docs/reference/elasticsearch/clients/go/advanced) (interceptors, observability)

<!-- ----------------------------------------------------------------------------------------------- -->

## Examples

The **[`_examples`](./_examples)** folder contains a number of recipes and comprehensive examples to get you started with the client,
including configuration and customization of the client, using a custom certificate authority (CA) for security (TLS),
mocking the transport for unit tests, embedding the client in a custom type, building queries, performing requests individually and in bulk, and parsing the responses.

<!-- ----------------------------------------------------------------------------------------------- -->

## License

This software is licensed under the [Apache 2 license](./LICENSE). See [NOTICE](./NOTICE).
