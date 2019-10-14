// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package elasticsearch_test

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"
)

func init() {
	log.SetFlags(0)
}

func ExampleNewDefaultClient() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s\n", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	defer res.Body.Close()

	log.Print(es.Transport.(*estransport.Client).URLs())
}

func ExampleNewClient() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "foo",
		Password: "bar",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	}

	es, _ := elasticsearch.NewClient(cfg)
	log.Print(es.Transport.(*estransport.Client).URLs())
}

func ExampleNewClient_logger() {
	// import "github.com/elastic/go-elasticsearch/v7/estransport"

	// Use one of the bundled loggers:
	//
	// * estransport.TextLogger
	// * estransport.ColorLogger
	// * estransport.CurlLogger
	// * estransport.JSONLogger

	cfg := elasticsearch.Config{
		Logger: &estransport.ColorLogger{Output: os.Stdout},
	}

	elasticsearch.NewClient(cfg)
}
