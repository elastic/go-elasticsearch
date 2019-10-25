# Example: fasthttp

This example demonstrates how to replace the `net/http` transport in the client with an
HTTP client from the [`github.com/valyala/fasthttp`](https://godoc.org/github.com/valyala/fasthttp) package.

The [`fasthttp.Transport` type](./fasthttp.go) is a drop-in replacement for `estransport.Config.Transport`, converting
request and response object between `net/http` and `fasthttp`.

To run the example:

    go run cmd/main.go
    # 1000 requests in 1.01 (985.6 req/s) | min: 596.739µs / max: 154.672371ms / mean: 813.541µs

To run the benchmarks:

    make bench

    BenchmarkHTTPClient/Info()-16             7624     1395744 ns/op     15911 B/op      120 allocs/op
    BenchmarkFastHTTPClient/Info()-16        20479      591367 ns/op      2040 B/op       24 allocs/op
