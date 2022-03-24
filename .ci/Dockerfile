# $ docker build --file Dockerfile --tag elastic/go-elasticsearch .
#
# $ docker run -it --env ELASTICSEARCH_URL=http://es1:9200 --network elasticsearch --rm elastic/go-elasticsearch go run _examples/main.go
#

ARG  VERSION=1-alpine
FROM golang:${VERSION}

RUN apk add --no-cache --quiet make curl git jq unzip tree && \
    go install golang.org/x/lint/golint@latest && \
    go install gotest.tools/gotestsum@latest

VOLUME ["/tmp"]

ENV CGO_ENABLED=0
ENV TERM xterm-256color

WORKDIR /go-elasticsearch

COPY . .

RUN go mod download

WORKDIR internal/cmd/generate
RUN go mod download

WORKDIR /go-elasticsearch