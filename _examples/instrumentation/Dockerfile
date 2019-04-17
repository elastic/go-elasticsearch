# $ docker build --tag elastic/go-elasticsearch-demo-instrumentation .
# $ docker run -it --network instrumentation_elasticstack --rm elastic/go-elasticsearch-demo-instrumentation
#
FROM golang:1-alpine

RUN apk add --no-cache --quiet gcc g++ ca-certificates make curl git jq

WORKDIR /go-elasticsearch-demo-instrumentation

COPY go.mod .
RUN go mod download

ENV TERM xterm-256color

ENV ELASTICSEARCH_URL=${ELASTICSEARCH_URL:-http://elasticsearch:9200}

ENV ELASTIC_APM_SERVER_URL=${ELASTIC_APM_SERVER_URL:-http://apm_server:8200}
ENV ELASTIC_APM_SERVICE_NAME=go-elasticsearch-demo-instrumentation
ENV ELASTIC_APM_METRICS_INTERVAL=5s
ENV ELASTIC_APM_LOG_FILE=stderr
ENV ELASTIC_APM_LOG_LEVEL=debug

COPY apmelasticsearch.go opencensus.go ./

CMD go run apmelasticsearch.go
