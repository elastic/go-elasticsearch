module github.com/elastic/go-elasticsearch/v8/_examples/logging

go 1.11

replace github.com/elastic/go-elasticsearch/v8 => ../..

require (
	github.com/elastic/go-elasticsearch/v8 master

	github.com/rs/zerolog v1.11.0
)
