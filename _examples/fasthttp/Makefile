GO_TEST_CMD = $(if $(shell which richgo),richgo test,go test)

test:  ## Run tests
	go run cmd/main.go

bench:  ## Run benchmarks
	$(GO_TEST_CMD) -run=none -bench=. -benchmem -benchtime=10s -v fasthttp_benchmark_test.go

.PHONY: test bench
