// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package estransport

import (
	"fmt"
	"net/url"
	"testing"
)

func TestDiscovery(t *testing.T) {
	u, _ := url.Parse("http://localhost:9200")
	tp := New(Config{URLs: []*url.URL{u}})

	res, err := tp.GetNodesInfo()
	if err != nil {
		t.Fatalf("ERROR: %s", err)
	}
	fmt.Printf("%+v\n", res)
}
