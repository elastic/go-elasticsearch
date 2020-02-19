module github.com/elastic/go-elasticsearch/v8/_examples/bulk

go 1.11

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/elastic/go-elasticsearch/v8 master
	github.com/mailru/easyjson v0.7.1
	github.com/valyala/fasthttp v1.9.0
)
