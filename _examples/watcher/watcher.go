// +build ignore

package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/tidwall/gjson"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/estransport"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

func main() {
	log.SetFlags(0)
	rand.Seed(time.Now().UnixNano())

	var (
		res *esapi.Response
		err error
	)

	// Create the Elasticsearch client
	//
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"https://elastic:elastic@localhost:9200"},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Logger: &estransport.ColorLogger{
			Output: os.Stdout,
			// EnableRequestBody:  true,
			// EnableResponseBody: true,
		},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Start the webhook server
	started := make(chan bool)
	go startServer(started)
	<-started

	// _ = func() {

	// Delete the Watcher and test indices
	//
	es.Indices.Delete([]string{"test-errors", "alerts", ".watcher-history-*"}, es.Indices.Delete.WithIgnoreUnavailable(true))

	// Register a new watch
	//
	res, _ = es.Watcher.PutWatch(
		"error-500",
		es.Watcher.PutWatch.WithBody(
			strings.NewReader(`
{
  "metadata": {
    "tags": [
      "errors"
    ]
  },
  "trigger": {
    "schedule": {
      "interval": "10s"
    }
  },
  "condition": {
    "compare": {
      "ctx.payload.hits.total": {
        "gt": 3
      }
    }
  },
  "throttle_period": "30s",
  "input": {
    "search": {
      "request": {
        "indices": [
          "test-errors"
        ],
        "body": {
          "query": {
            "bool": {
              "must": [
                {
                  "match": {
                    "status": 500
                  }
                },
                {
                  "range": {
                    "timestamp": {
                      "from": "{{ctx.trigger.scheduled_time}}||-5m",
                      "to": "{{ctx.trigger.triggered_time}}"
                    }
                  }
                }
              ]
            }
          },
          "aggregations": {
            "hosts": {
              "terms": {
                "field": "host"
              }
            }
          }
        }
      }
    }
  },
  "actions": {
    "index_payload": {
      "transform": {
        "script": {
          "lang": "painless",
          "source": "[ 'watch_id': ctx.watch_id, 'payload': ctx.payload ]"
        }
      },
      "index": {
        "index": "alerts"
      }
    },
    "ping_webhook": {
      "webhook": {
        "method": "post",
        "url": "http://localhost:4567",
        "body": "{\"watch_id\" : \"{{ctx.watch_id}}\", \"payload\" : \"{{ctx.payload}}\"}"
      }
    },
    "send_email": {
      "transform": {
        "script": {
          "source": "[ 'total': ctx.payload.hits.total, 'hosts': ctx.payload.aggregations.hosts.buckets.collect(bucket -> [ 'host': bucket.key, 'errors': bucket.doc_count ]), 'errors': ctx.payload.hits.hits.collect(d -> d._source) ]"
        }
      },
      "email": {
        "to": "alert@example.com",
        "subject": "[ALERT] {{ctx.watch_id}}",
        "body": "Received {{ctx.payload.total}} errors in the last 5 minutes.\nHosts:\n{{#ctx.payload.hosts}}- {{host}} ({{errors}} errors)\n{{/ctx.payload.hosts}}\nA file with complete data is attached to this message.\n\n",
        "attachments": {
          "data.yml": {
            "data": {
              "format": "yaml"
            }
          }
        }
      }
    }
  }
}`),
		),
	)
	if res.IsError() {
		log.Fatalf("Error storing the watch: %s", res)
	}

	// Create the index and store example documents
	//
	es.Indices.Create(
		"test-errors",
		es.Indices.Create.WithBody(strings.NewReader(`{ "mappings": { "properties": { "host": { "type": "keyword" } } }}`)),
	)
	for i := 1; i <= 15; i++ {
		es.Index(
			"test-errors",
			esutil.NewJSONReader(map[string]interface{}{
				"timestamp": time.Now().UTC().Format(time.RFC3339),
				"status":    fmt.Sprintf(`"%d00"`, randNum(4, 5)),
				"host":      fmt.Sprintf(`"10.0.0.%d"`, randNum(1, 3)),
			}),
		)
	}

	fmt.Print("\nWaiting 30 seconds...")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for countdown := 30; countdown > 0; countdown-- {
		select {
		case <-ticker.C:
			fmt.Print(".")
		}
	}
	fmt.Print("\n\n")
	// }

	res, err = es.Search(
		es.Search.WithIndex(".watcher-history-*"),
		es.Search.WithQuery("watch_id:error-500"),
		es.Search.WithSort("trigger_event.triggered_time:asc"),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	if res.IsError() {
		log.Fatalf("Error response: %s", res)
	}

	gjson.Get(read(res.Body), "hits.hits").ForEach(func(key, value gjson.Result) bool {
		values := gjson.GetMany(value.String(), "_source.watch_id", "_source.state", "_source.result.execution_time")
		id, state := values[0], strings.ToUpper(values[1].String())
		timestamp, _ := time.Parse(time.RFC3339, values[2].String())
		fmt.Printf("[%s] %10s at %s", id, state, timestamp.Format(time.Kitchen))
		if state == "EXECUTED" {
			fmt.Print(" | ")
			actions := value.Get("_source.result.actions").Array()
			for i, v := range actions {
				fmt.Printf("%s:%s", v.Get("id"), v.Get("status"))
				if i < len(actions)-1 {
					fmt.Print(", ")
				}
			}
		}
		fmt.Print("\n")
		return true
	})
}

func randNum(min, max int) int {
	return min + rand.Intn(max+1-min)
}

func read(r io.Reader) string {
	var b bytes.Buffer
	b.ReadFrom(r)
	return b.String()
}

func startServer(started chan<- bool) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("WEBHOOK: %s", r)
		fmt.Fprint(w, "HELLO")
	})
	go http.ListenAndServe("localhost:4567", nil)
	started <- true
}
