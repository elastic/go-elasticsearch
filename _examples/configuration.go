// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build ignore

package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	log.SetFlags(0)

	// This example demonstrates how to configure the client's Transport.
	//
	// NOTE: These values are for illustrative purposes only, and not suitable
	//       for any production use. The default transport is sufficient.
	//
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Millisecond,
			DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
				// ...
			},
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
	} else {
		log.Println(es.Info())
		// => dial tcp: i/o timeout
	}
}
