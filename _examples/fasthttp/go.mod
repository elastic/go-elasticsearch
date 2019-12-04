module github.com/elastic/go-elasticsearch/v7/_examples/fasthttp

go 1.11

replace github.com/elastic/go-elasticsearch/v7 => ../..

require (
	github.com/elastic/go-elasticsearch/v7 7.x
	github.com/valyala/fasthttp v1.5.0
)
