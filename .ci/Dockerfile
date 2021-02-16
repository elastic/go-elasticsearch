# $ docker build --file Dockerfile --tag elastic/go-elasticsearch .
#
# $ docker run -it --env ELASTICSEARCH_URL=http://es1:9200 --network elasticsearch --rm elastic/go-elasticsearch go run _examples/main.go
#

ARG  VERSION=1-alpine
FROM golang:${VERSION}

RUN apk add --no-cache --quiet make curl git jq unzip tree && \
    go get -u golang.org/x/lint/golint && \
    curl -sSL --retry 3 --retry-connrefused https://github.com/gotestyourself/gotestsum/releases/download/v1.6.1/gotestsum_1.6.1_linux_amd64.tar.gz | tar -xz -C /usr/local/bin gotestsum

VOLUME ["/tmp"]

ENV CGO_ENABLED=0
ENV TERM xterm-256color

WORKDIR /go-elasticsearch

COPY . .

RUN go mod download

WORKDIR internal/cmd/generate
RUN go mod download

WORKDIR /go-elasticsearch