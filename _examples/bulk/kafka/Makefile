SHELL := /bin/bash

test:
	go build -o /dev/null kafka.go

setup: stack status wait_for_kibana load_kibana_dashboard

clean:
	@docker-compose down --volumes

run:
	go run kafka.go

stack:
	@docker-compose up --detach

status:
	@{ \
		cols=`tput cols`; max=98; i=1; \
		while [[ "$$i" -lt "$$cols" && "$$i" -lt "$$max" ]]; do printf '-'; let i=$$i+1; done; printf "\n"; \
		docker-compose ps; \
		i=1; while [[ "$$i" -lt "$$cols" && "$$i" -lt "$$max" ]]; do printf '-'; let i=$$i+1; done; printf "\n"; \
	}

wait_for_kibana:
	@{ \
		printf "Waiting for Kibana..."; \
		until docker inspect kibana > /dev/null 2>&1 && [[ `docker inspect -f '{{ .State.Health.Status }}' kibana` == "healthy" ]]; do printf '.'; sleep 5; done; \
		printf "\nOpen dashboard at <http://localhost:5601/app/kibana#/dashboard/140b5490-5fce-11ea-a238-bf5970186390>\n"; \
	}

load_kibana_dashboard:
	@{ \
		curl -s -S -X POST -H 'kbn-xsrf: true' 'http://localhost:5601/api/saved_objects/_import?overwrite=true' --form file=@etc/kibana-objects.ndjson > /dev/null; \
	}

save_kibana_dashboard:
	@{ \
		curl -s -S -X POST -H 'Content-Type: application/json' -H 'kbn-xsrf: true' -o './etc/kibana-objects.ndjson' 'http://localhost:5601/api/saved_objects/_export' -d ' \
{ \
  "objects": [ \
  { "type": "index-pattern", "id": "stocks-index-pattern" }, \
  { "type": "dashboard",     "id": "140b5490-5fce-11ea-a238-bf5970186390" } \
  ], \
  "includeReferencesDeep": true \
}'; \
	}

.PHONY: test setup run clean stack status wait_for_kibana load_kibana_dashboard save_kibana_dashboard
