# Example: fasthttp

This example demonstrates how to replace the `net/http` transport in the client with an
HTTP client from the [`github.com/valyala/fasthttp`](https://godoc.org/github.com/valyala/fasthttp) package.

The [`fasthttp.Transport` type](./fasthttp.go) is a drop-in replacement for `estransport.Config.Transport`, converting
request and response object between `net/http` and `fasthttp`.

To run the example:

    go run cmd/main.go
    # 1000 requests in 1.15 (870.2 req/s) | min: 785.292µs / max: 10.553253ms / mean: 1.052083ms

To run the benchmarks:

    make bench

    BenchmarkHTTPClient/Info()-4              2000     8721355 ns/op     16279 B/op      131 allocs/op
    BenchmarkFastHTTPClient/Info()-4         10000     1182717 ns/op      3075 B/op       39 allocs/op
