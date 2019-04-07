// +build integration

package esapi_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v6"
)

func TestAPI(t *testing.T) {
	t.Run("Search", func(t *testing.T) {
		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			t.Fatalf("Error creating the client: %s\n", err)
		}

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
}
