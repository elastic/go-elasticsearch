module github.com/elastic/go-elasticsearch/v8/_examples/fasthttp

go 1.11

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/elastic/go-elasticsearch/v8 master
	github.com/valyala/fasthttp v1.5.0
)
