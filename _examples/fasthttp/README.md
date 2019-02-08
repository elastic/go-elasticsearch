# Example: fasthttp

This example demonstrates how to replace the `net/http` transport in the client with an
HTTP client from the [`github.com/valyala/fasthttp`](https://godoc.org/github.com/valyala/fasthttp) package.

The [`fasthttp.Transport` type](./fasthttp.go) is a drop-in replacement for `estransport.Config.Transport`, converting
request and response object between `net/http` and `fasthttp`.

To run the example:

    go run cmd/main.go

To run the benchmarks:

    make test
