module github.com/elastic/go-elasticsearch/v8/benchmarks

go 1.14

replace github.com/elastic/go-elasticsearch/v8 => ../../

require (
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20200408073057-6f36a473b19f
	github.com/fatih/color v1.7.0
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/montanaflynn/stats v0.6.3
	github.com/tidwall/gjson v1.6.5
)
