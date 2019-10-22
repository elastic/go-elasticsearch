// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build integration

package esapi_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
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

	t.Run("OpaqueID", func(t *testing.T) {
		var (
			buf bytes.Buffer

			res *esapi.Response
			err error

			requestID = "reindex-123"
		)

		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			t.Fatalf("Error creating the client: %s\n", err)
		}

		// Prepare indices
		//
		es.Indices.Delete([]string{"test", "reindexed"}, es.Indices.Delete.WithIgnoreUnavailable(true))

		// Index data
		//
		for j := 1; j <= 1000; j++ {
			meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d" } }%s`, j, "\n"))
			data := []byte(`{"content":"` + strings.Repeat("ABC", 100) + `"}`)
			data = append(data, "\n"...)

			buf.Grow(len(meta) + len(data))
			buf.Write(meta)
			buf.Write(data)
		}
		res, err = es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithIndex("test"), es.Bulk.WithRefresh("true"))
		if err != nil {
			t.Fatalf("Failed to index data: %s", err)
		}
		res.Body.Close()
		if res.IsError() {
			t.Fatalf("Failed to index data: %s", res.Status())
		}

		// Launch reindexing task with wait_for_completion=false
		//
		res, err = es.Reindex(
			strings.NewReader(`{"source":{"index":"test"}, "dest": {"index":"reindexed"}}`),
			es.Reindex.WithWaitForCompletion(false),
			es.Reindex.WithRequestsPerSecond(1),
			es.Reindex.WithOpaqueID(requestID))
		if err != nil {
			t.Fatalf("Failed to reindex: %s", err)
		}
		if res.IsError() {
			t.Fatalf("Failed to reindex: %s", res.Status())
		}
		time.Sleep(10 * time.Millisecond)

		res, err = es.Tasks.List(es.Tasks.List.WithPretty())
		if err != nil {
			t.Fatalf("ERROR: %s", err)
		}
		res.Body.Close()
		if res.IsError() {
			t.Fatalf("Failed to get tasks: %s", res.Status())
		}

		// Get the list of tasks
		//
		res, err = es.Tasks.List(es.Tasks.List.WithPretty())
		if err != nil {
			t.Fatalf("ERROR: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			t.Fatalf("Failed to get tasks: %s", res.Status())
		}

		type task struct {
			Node        string
			ID          int
			Action      string
			RunningTime time.Duration `json:"running_time_in_nanos"`
			Cancellable bool
			Headers     map[string]interface{}
		}

		type node struct {
			Tasks map[string]task
		}

		var nodes map[string]map[string]node
		if err := json.NewDecoder(res.Body).Decode(&nodes); err != nil {
			t.Fatalf("Failed to decode response: %s", err)
		}

		var hasReindexTask bool

		for _, n := range nodes["nodes"] {
			for taskID, task := range n.Tasks {
				if task.Headers["X-Opaque-Id"] == requestID {
					if strings.Contains(task.Action, "reindex") {
						hasReindexTask = true
					}
					fmt.Printf("* %s, %s | %s (%s)\n", requestID, taskID, task.Action, task.RunningTime)
				}
			}
		}

		if !hasReindexTask {
			t.Errorf("Expected reindex task in %+v", nodes["nodes"])
		}

		for _, n := range nodes["nodes"] {
			for taskID, task := range n.Tasks {
				if task.Headers["X-Opaque-Id"] == requestID {
					if task.Cancellable {
						fmt.Printf("=> Closing task %s\n", taskID)
						res, err = es.Tasks.Cancel(es.Tasks.Cancel.WithTaskID(taskID))
						if err != nil {
							t.Fatalf("ERROR: %s", err)
						}
						res.Body.Close()
						if res.IsError() {
							t.Fatalf("Failed to cancel task: %s", res)
						}
					}
				}
			}
		}
	})
}
