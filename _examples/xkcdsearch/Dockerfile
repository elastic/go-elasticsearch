# $ go mod vendor
# $ docker build --tag elastic/go-elasticsearch-demo-xkcdsearch .
#
# $ docker run -it --network elasticsearch --name xkcdsearch-demo --publish 8000:8000 --rm elastic/go-elasticsearch-demo-xkcdsearch
#
FROM golang:1.11-alpine AS Builder

WORKDIR /go-elasticsearch-demo-xkcdsearch
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -tags netgo -o /xkcdsearch cmd/xkcdsearch/main.go

FROM alpine
RUN apk update && apk add ca-certificates

COPY --from=Builder /xkcdsearch /xkcdsearch
COPY --from=Builder /go-elasticsearch-demo-xkcdsearch/web /web

ENV ELASTICSEARCH_URL=http://es01:9200
ENV HTTP_LISTEN_HOST=0.0.0.0
ENV HTTP_LISTEN_PORT=8000

EXPOSE 8000
ENTRYPOINT ["/xkcdsearch", "server"]
