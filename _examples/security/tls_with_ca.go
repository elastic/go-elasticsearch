// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build ignore

package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	log.SetFlags(0)

	var (
		err error

		// --> Configure the path to the certificate authority and the password
		//
		cacert   = flag.String("cacert", "certificates/ca/ca.crt", "Path to the file with certificate authority")
		password = flag.String("password", "elastic", "Elasticsearch password")
	)
	flag.Parse()

	// --> Read the certificate from file
	//
	cert, err := ioutil.ReadFile(*cacert)
	if err != nil {
		log.Fatalf("ERROR: Unable to read CA from %q: %s", *cacert, err)
	}

	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{"https://localhost:9200"},
			Username:  "elastic",
			Password:  *password,

			// --> Pass the certificate to the client
			//
			CACert: cert,
		})
	if err != nil {
		log.Fatalf("ERROR: Unable to create client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("ERROR: Unable to get response: %s", err)
	}

	log.Println(res)
}
