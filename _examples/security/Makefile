test:
	@for file in ./*.go; do \
		echo "go build -o /dev/null $$file"; \
		go build -o /dev/null $$file; \
	done;

test-integ:
	@for file in ./*.go; do \
		echo "go run $$file"; \
		go run $$file; \
	done;

certificates: certificates-clean
	docker-compose --file certificates-create.yml run --rm create_certificates

cluster:
	docker-compose --file elasticsearch-cluster.yml up --remove-orphans --detach; echo;
	@docker-compose --file elasticsearch-cluster.yml ps;
	@{ \
		set -e; \
		until \
			docker inspect example_elasticsearch_1 > /dev/null 2>&1 && \
			[[ `docker inspect -f "{{ .State.Health.Status }}" example_elasticsearch_1` == "healthy" ]]; \
				do printf '-'; sleep 1; \
		done; echo "> [OK]"; \
	}

certificates-clean:
	rm -rf ./certificates

cluster-clean:
	docker-compose --file elasticsearch-cluster.yml down --volumes

.PHONY: test test-integ certificates cluster certificates-clean cluster-clean
