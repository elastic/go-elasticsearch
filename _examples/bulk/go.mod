module github.com/elastic/go-elasticsearch/v7/_examples/bulk

go 1.11

replace github.com/elastic/go-elasticsearch/v7 => ../..

require (
	github.com/cenkalti/backoff/v4 v4.0.0
	github.com/dustin/go-humanize v1.0.0
	github.com/elastic/go-elasticsearch/v7 v7.5.1-0.20210608143310-9747f0e42f35
	github.com/mailru/easyjson v0.7.7 // indirect
)
