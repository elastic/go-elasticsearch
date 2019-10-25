# Example: fasthttp

This example demonstrates how to replace the `net/http` transport in the client with an
HTTP client from the [`github.com/valyala/fasthttp`](https://godoc.org/github.com/valyala/fasthttp) package.

The [`fasthttp.Transport` type](./fasthttp.go) is a drop-in replacement for `estransport.Config.Transport`, converting
request and response object between `net/http` and `fasthttp`.

To run the example:

    go run cmd/main.go
    # 1000 requests in 1.07 (933.4 req/s) | min: 749.961µs / max: 10.003318ms / mean: 1.013992ms

To run the benchmarks:

    make bench

    BenchmarkHTTPClient/Info()-4         	    1591	   7139770 ns/op	   16725 B/op	     132 allocs/op
    BenchmarkFastHTTPClient/Info()-4     	   10000	   1315049 ns/op	    2255 B/op	      24 allocs/op
