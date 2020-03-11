module github.com/elastic/go-elasticsearch/v7/_examples/bulk/benchmarks

go 1.11

replace github.com/elastic/go-elasticsearch/v7 => ../../..

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/elastic/go-elasticsearch/v7 7.x
	github.com/mailru/easyjson v0.7.1
	github.com/montanaflynn/stats v0.6.3
	github.com/valyala/fasthttp v1.9.0
)
