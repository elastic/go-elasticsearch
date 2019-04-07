module github.com/elastic/go-elasticsearch/v5/_examples/fasthttp

go 1.11

replace github.com/elastic/go-elasticsearch/v5 => ../..

require (
	github.com/elastic/go-elasticsearch/v5 5.x

	github.com/valyala/fasthttp v1.1.0
)
