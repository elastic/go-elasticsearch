// +build integration,multinode

package estransport_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/elastic/go-elasticsearch/estransport"
)

var (
	_ = fmt.Print
)

func TestTransportSelector(t *testing.T) {

	NodeName := func(t *testing.T, transport estransport.Interface) string {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		res, err := transport.Perform(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		fmt.Printf("> GET %s\n", req.URL)

		r := struct {
			Name string
		}{}

		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			t.Fatalf("Error parsing the response body: %s", err)
		}

		return r.Name
	}

	t.Run("RoundRobin", func(t *testing.T) {
		var (
			node string
		)
		transport := estransport.New(estransport.Config{URLs: []*url.URL{
			&url.URL{Scheme: "http", Host: "localhost:9200"},
			&url.URL{Scheme: "http", Host: "localhost:9201"},
		}})

		node = NodeName(t, transport)
		if node != "es1" {
			t.Errorf("Unexpected node, want=e1, got=%s", node)
		}

		node = NodeName(t, transport)
		if node != "es2" {
			t.Errorf("Unexpected node, want=e1, got=%s", node)
		}

		node = NodeName(t, transport)
		if node != "es1" {
			t.Errorf("Unexpected node, want=e1, got=%s", node)
		}

	})
}
