export ELASTICSEARCH_URL=http://elastic:elastic@localhost:9200

test:  ## Run examples
	go run -tags search_default default.go
	go run -tags search_pagination pagination.go

setup:

.PHONY: test setup
