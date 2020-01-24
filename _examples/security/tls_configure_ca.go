// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build ignore

package main

import (
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

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

	ca, err := ioutil.ReadFile(*cacert)
	if err != nil {
		log.Fatalf("ERROR: Unable to read CA from %q: %s", *cacert, err)
	}

	// --> Clone the default HTTP transport
	//
	tp := http.DefaultTransport.(*http.Transport).Clone()

	// --> Initialize the set of root certificate authorities
	//
	if tp.TLSClientConfig.RootCAs, err = x509.SystemCertPool(); err != nil {
		log.Fatalf("ERROR: Problem adding system CA: %s", err)
	}

	// --> Add the custom certificate authority
	//
	if ok := tp.TLSClientConfig.RootCAs.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("ERROR: Problem adding CA from file %q", *cacert)
	}

	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{"https://localhost:9200"},
			Username:  "elastic",
			Password:  *password,

			// --> Pass the transport to the client
			//
			Transport: tp,
		},
	)
	if err != nil {
		log.Fatalf("ERROR: Unable to create client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("ERROR: Unable to get response: %s", err)
	}

	log.Println(res)
}
