// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var (
	listenHost = "localhost"
	listenPort = "8000"
	listenAddr string
)

const responseBody = `
{
  "took": 30,
  "errors": true,
  "items": [
    {
      "index": {
        "_index": "test",
        "_id": "1",
        "_version": 1,
        "result": "created",
        "_shards": { "total": 1, "successful": 1, "failed": 0 },
        "_seq_no": 0,
        "_primary_term": 1,
        "status": 201
      }
    },
    {
      "create": {
        "_index": "test",
        "_id": "1",
        "status": 409,
        "error": {
          "type": "version_conflict_engine_exception",
          "reason": "[1]: version conflict, document already exists (current version [1])",
          "index_uuid": "eZMQ7DUzT56RLaQcAjOlxg",
          "index": "test-bulk-integration",
          "shard": "0"
        }
      }
    },
    {
      "delete": {
        "_index": "test",
        "_id": "2",
        "_version": 1,
        "result": "not_found",
        "_shards": { "total": 1, "successful": 1, "failed": 0 },
        "_seq_no": 3,
        "_primary_term": 1,
        "status": 404
      }
    },
    {
      "update": {
        "_index": "test",
        "_id": "3",
        "_version": 2,
        "result": "updated",
        "_shards": { "total": 1, "successful": 1, "failed": 0 },
        "_seq_no": 4,
        "_primary_term": 1,
        "status": 200
      }
    }
  ]
}`

func init() {
	if envListenHost := os.Getenv("HTTP_LISTEN_HOST"); envListenHost != "" {
		listenHost = envListenHost
	}
	if envListenPort := os.Getenv("HTTP_LISTEN_PORT"); envListenPort != "" {
		listenPort = envListenPort
	}
	listenAddr = listenHost + ":" + listenPort
}

func main() {
	log.SetFlags(0)

	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) { io.WriteString(w, responseBody) })

	log.Printf("Starting server at <%s>...", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Printf("Error launching server: %s", err)
	}
}
