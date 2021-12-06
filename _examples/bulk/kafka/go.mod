module github.com/elastic/go-elasticsearch/v8/_examples/bulk/kafka

go 1.11

replace github.com/elastic/go-elasticsearch/v8 => ../../..

require (
	github.com/cucumber/godog v0.8.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210817150010-57d659deaca7
	github.com/prometheus/procfs v0.0.3 // indirect
	github.com/segmentio/kafka-go v0.4.25 // indirect
	go.elastic.co/apm v1.14.0
	go.elastic.co/apm/module/apmelasticsearch v1.14.0 // indirect
)
