// +build integration

package esapi_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

func TestAPI(t *testing.T) {
	t.Run("Search", func(t *testing.T) {
		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			t.Fatalf("Error creating the client: %s\n", err)
		}

		es.Cluster.Health(es.Cluster.Health.WithWaitForStatus("yellow"))
		res, err := es.Search(es.Search.WithTimeout(500 * time.Millisecond))
		if err != nil {
			t.Fatalf("Error getting the response: %s\n", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			t.Fatalf("Error response: %s", res.String())
		}

		var d map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&d)
		if err != nil {
			t.Fatalf("Error parsing the response: %s\n", err)
		}
		fmt.Printf("took=%vms\n", d["took"])
	})

	t.Run("Headers", func(t *testing.T) {
		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			t.Fatalf("Error creating the client: %s\n", err)
		}

		res, err := es.Info(es.Info.WithHeader(map[string]string{"Accept": "application/yaml"}))
		if err != nil {
			t.Fatalf("Error getting the response: %s\n", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			t.Fatalf("Error response: %s", res.String())
		}

		if !strings.HasPrefix(res.String(), "[200 OK] ---") {
			t.Errorf("Unexpected response body: doesn't start with '[200 OK] ---'; %s", res.String())
		}
	})
}
