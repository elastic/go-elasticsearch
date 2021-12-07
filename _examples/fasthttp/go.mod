module github.com/elastic/go-elasticsearch/v8/_examples/fasthttp

go 1.11

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210817150010-57d659deaca7
	github.com/valyala/fasthttp v1.5.0
)
