// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

//go:build !integration
// +build !integration

package elasticsearch_test

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
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

	log.Print(es.Transport.(*elastictransport.Client).URLs())
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
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	es, _ := elasticsearch.NewClient(cfg)
	log.Print(es.Transport.(*elastictransport.Client).URLs())
}

func ExampleNewClient_logger() {
	// import "github.com/elastic/go-elasticsearch/v8/elastictransport"

	// Use one of the bundled loggers:
	//
	// * elastictransport.TextLogger
	// * elastictransport.ColorLogger
	// * elastictransport.CurlLogger
	// * elastictransport.JSONLogger

	cfg := elasticsearch.Config{
		Logger: &elastictransport.ColorLogger{Output: os.Stdout},
	}

	elasticsearch.NewClient(cfg)
}
