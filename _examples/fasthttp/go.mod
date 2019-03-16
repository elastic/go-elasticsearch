module github.com/elastic/go-elasticsearch/_examples/fasthttp

go 1.11

replace github.com/elastic/go-elasticsearch => ../..

require (
	github.com/elastic/go-elasticsearch 5.x
	github.com/valyala/fasthttp v1.1.0
)
