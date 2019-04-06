module github.com/elastic/go-elasticsearch/_examples/logging

go 1.11

replace github.com/elastic/go-elasticsearch => ../..

require (
	github.com/elastic/go-elasticsearch master

	github.com/rs/zerolog v1.11.0
)
