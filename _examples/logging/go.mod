module github.com/elastic/go-elasticsearch/v6/_examples/logging

go 1.11

replace github.com/elastic/go-elasticsearch/v6 => ../..

require (
	github.com/elastic/go-elasticsearch/v6 6.x

	github.com/rs/zerolog v1.11.0
)
