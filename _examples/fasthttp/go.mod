module github.com/elastic/go-elasticsearch/v6/_examples/fasthttp

go 1.11

replace github.com/elastic/go-elasticsearch/v6 => ../..

require (
	github.com/elastic/go-elasticsearch/v6 6.x

	github.com/valyala/fasthttp v1.5.0
)
