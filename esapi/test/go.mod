module github.com/elastic/go-elasticsearch/esapi/test

go 1.11

replace github.com/elastic/go-elasticsearch => ../../

require (
	github.com/elastic/go-elasticsearch master

	gopkg.in/yaml.v2 v2.2.2
)
