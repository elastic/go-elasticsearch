module github.com/elastic/go-elasticsearch/v7/_examples/fasthttp

go 1.11

replace github.com/elastic/go-elasticsearch/v7 => ../..

require (
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/elastic/go-elasticsearch/v7 v7.5.1-0.20210608143310-9747f0e42f35
	github.com/klauspost/compress v1.13.1 // indirect
	github.com/valyala/fasthttp v1.27.0
)
