module github.com/elastic/go-elasticsearch/v7/_examples/bulk

go 1.11

replace github.com/elastic/go-elasticsearch/v7 => ../..

require (
	github.com/cenkalti/backoff/v4 v4.0.0
	github.com/dustin/go-humanize v1.0.0
	github.com/elastic/go-elasticsearch/v7 7.x
)
