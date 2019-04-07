module github.com/elastic/go-elasticsearch/v5/_examples/logging

go 1.11

replace github.com/elastic/go-elasticsearch/v5 => ../..

require (
	github.com/elastic/go-elasticsearch/v5 5.x

	github.com/rs/zerolog v1.11.0
)
