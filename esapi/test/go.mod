module github.com/elastic/go-elasticsearch/v8/esapi/test

go 1.22

toolchain go1.22.0

replace github.com/elastic/go-elasticsearch/v8 => ../../

require (
	github.com/elastic/elastic-transport-go/v8 v8.6.1
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210817150010-57d659deaca7
	gopkg.in/yaml.v2 v2.4.0
)
