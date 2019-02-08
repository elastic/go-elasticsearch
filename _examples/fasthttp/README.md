# Example: fasthttp

This example demonstrates how to replace the `net/http` transport in the client with an
HTTP client from the [`github.com/valyala/fasthttp`](https://godoc.org/github.com/valyala/fasthttp) package.

The [`fasthttp.Transport` type](./fasthttp.go) is a drop-in replacement for `estransport.Config.Transport`, converting
request and response object between `net/http` and `fasthttp`.

To run the example:

    go run cmd/main.go

To run the benchmarks:

    make bench

    BenchmarkHTTPClient/Info()-4        3000  5504114 ns/op    16232 B/op    130 allocs/op
    BenchmarkFastHTTPClient/Info()-4   10000  1183802 ns/op     3233 B/op     41 allocs/op
