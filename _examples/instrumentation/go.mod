module github.com/elastic/go-elasticsearch/v8/_examples/instrumentation/opencensus

go 1.11

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20191002063538-b491ce54d752
	github.com/fatih/color v1.7.0
	github.com/mattn/go-colorable v0.1.0 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	go.elastic.co/apm v1.11.0
	go.elastic.co/apm/module/apmelasticsearch v1.5.0
	go.opencensus.io v0.19.0
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
)
