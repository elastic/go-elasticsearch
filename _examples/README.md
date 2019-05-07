This directory contains examples and recipes for the Elasticsearch Go client.

## Configuration & Customization

The [**`configuration.go`**](./configuration.go) and [**`customization.go`**](./customization.go) files contain information for configuring and customizing the client transport.

## Logging

The [**`logging`**](./logging) directory contains examples for using the default loggers and implementing a custom logger.

## Bulk

The [**`bulk`**](./bulk) directory contains a comprehensive example of using the _Bulk_ API.

## Encoding

The [**`encoding`**](./encoding) directory contains examples of using third-party packages
and client helper methods for a more efficient encoding and decoding of the request and response bodies.

## Fast HTTP

The [**`fasthttp`**](./fasthttp) directory contains a demonstration of replacing the default client transport with an HTTP client from the [`github.com/valyala/fasthttp`](https://godoc.org/github.com/valyala/fasthttp) package.

## Instrumentation

The [**`instrumentation`**](./instrumentation) directory contains recipes for instrumenting the client with the OpenCensus and Elastic APM packages.

## Extension

The [**`extension`**](./extension) directory contains an example of extending the client APIs, for example
to support custom endpoints installed by a plugin.

## Google Cloud Functions

The [**`cloudfunction`**](./cloudfunction) directory contains a simple web service for checking Elasticsearch cluster status.

## XKCD Search

The [**`xkcdsearch`**](./xkcdsearch) directory contains a command-line utility for indexing and searching an archive of [xkcd.com](https://xkcd.com) artworks.
