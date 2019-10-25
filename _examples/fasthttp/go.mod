module github.com/elastic/go-elasticsearch/v7/_examples/fasthttp

go 1.11

replace github.com/elastic/go-elasticsearch/v7 => ../..

require (
	github.com/elastic/go-elasticsearch/v7 v7.0.0-20190407092644-3fb2a278216b
	github.com/valyala/fasthttp v1.5.0
)
