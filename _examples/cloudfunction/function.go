// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// Package clusterstatus demonstrates using the Elasticsearch client for Go with Google Cloud Functions.
//
// Deploy the function with the gcloud command:
//
//   $ go mod vendor
//   $ gcloud functions deploy clusterstatus \
//   	    --entry-point Health \
//   	    --runtime go111 \
//   	    --trigger-http \
//   	    --memory 128MB \
//   	    --set-env-vars ELASTICSEARCH_URL=https://...cloud.es.io:9243
//
// Invoke your function over HTTP:
//
//   $ curl https://...cloudfunctions.net/clusterstatus
//
package clusterstatus

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

// ES holds a reference to the Elasticsearch client
//
// See: https://cloud.google.com/functions/docs/concepts/go-runtime#one-time_initialization
//
var ES *elasticsearch.Client

func init() {
	log.SetFlags(0)

	var err error
	ES, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

// Health returns the status of the cluster (red, yellow, green).
//
func Health(w http.ResponseWriter, r *http.Request) {
	var j map[string]interface{}

	res, err := ES.Cluster.Health()
	if err != nil {
		log.Printf("Error getting response from Elasticsearch: %s", err)
		http.Error(w, `{"status" : "error"}`, http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("HTTP response error: %s", res.Status())
		http.Error(w, `{"status" : "error"}`, http.StatusBadGateway)
		return
	}

	if err := json.NewDecoder(res.Body).Decode(&j); err != nil {
		log.Printf("Error parsing the HTTP response body: %s", err)
		http.Error(w, `{"status" : "error"}`, http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":%q}`, j["status"].(string))
	}
}
