package esapi_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

// TODO(karmi): Refactor into a shared mock/testing package

var (
	defaultResponse    = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader("MOCK"))}
	defaultRoundTripFn = func(*http.Request) (*http.Response, error) { return defaultResponse, nil }
	errorRoundTripFn   = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 400,
			Body: ioutil.NopCloser(strings.NewReader(`
					{ "error" : {
					    "root_cause" : [
					      {
					        "type" : "parsing_exception",
					        "reason" : "no [query] registered for [foo]",
					        "line" : 1,
					        "col" : 22
					      }
					    ],
					    "type" : "parsing_exception",
					    "reason" : "no [query] registered for [foo]",
					    "line" : 1,
					    "col" : 22
					  },
					  "status" : 400
					}`)),
		}, nil
	}
)

type FakeTransport struct {
	RoundTripFn func(*http.Request) (*http.Response, error)
}

func (t *FakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.RoundTripFn(req)
}

func newFakeClient(b *testing.B) *elasticsearch.Client {
	cfg := elasticsearch.Config{Transport: &FakeTransport{RoundTripFn: defaultRoundTripFn}}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		b.Fatalf("Unexpected error when creating a client: %s", err)
	}

	return es
}

func newFakeClientWithError(b *testing.B) *elasticsearch.Client {
	cfg := elasticsearch.Config{Transport: &FakeTransport{RoundTripFn: errorRoundTripFn}}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		b.Fatalf("Unexpected error when creating a client: %s", err)
	}

	return es
}

func BenchmarkAPI(b *testing.B) {
	var es = newFakeClient(b)
	var eserr = newFakeClientWithError(b)

	b.Run("client.Info()                      ", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if _, err := es.Info(); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("client.Info() WithContext          ", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if _, err := es.Info(es.Info.WithContext(context.Background())); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("InfoRequest{}.Do()                 ", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			req := esapi.InfoRequest{}
			if _, err := req.Do(context.Background(), es); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("client.Cluster.Health()            ", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if _, err := es.Cluster.Health(); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("client.Cluster.Health() With()     ", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := es.Cluster.Health(
				es.Cluster.Health.WithContext(context.Background()),
				es.Cluster.Health.WithLevel("indices"),
				es.Cluster.Health.WithPretty(),
			)

			if err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("ClusterHealthRequest{}.Do()        ", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			req := esapi.ClusterHealthRequest{}
			if _, err := req.Do(context.Background(), es); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("ClusterHealthRequest{...}.Do()     ", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			req := esapi.ClusterHealthRequest{Level: "indices", Pretty: true}
			if _, err := req.Do(context.Background(), es); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("client.Index() With()             ", func(b *testing.B) {
		indx := "test"
		body := strings.NewReader(`{"title" : "Test"}`)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := es.Index(
				indx,
				body,
				es.Index.WithDocumentID("1"),
				es.Index.WithRefresh("true"),
				es.Index.WithContext(context.Background()),
				es.Index.WithRefresh("true"),
				es.Index.WithPretty(),
				es.Index.WithTimeout(100),
			)

			if err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("IndexRequest{...}.Do()             ", func(b *testing.B) {
		b.ResetTimer()
		var body strings.Builder
		for i := 0; i < b.N; i++ {
			docID := strconv.FormatInt(int64(i), 10)
			body.Reset()
			body.WriteString(`{"foo" : "bar `)
			body.WriteString(docID)
			body.WriteString(`	" }`)

			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: docID,
				Body:       strings.NewReader(body.String()),
				Refresh:    "true",
				Pretty:     true,
				Timeout:    100,
			}

			if _, err := req.Do(context.Background(), es); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("IndexRequest{...}.Do() reused      ", func(b *testing.B) {
		b.ResetTimer()
		var body strings.Builder

		req := esapi.IndexRequest{}
		req.Index = "test"
		req.Refresh = "true"
		req.Pretty = true
		req.Timeout = 100

		for i := 0; i < b.N; i++ {
			docID := strconv.FormatInt(int64(i), 10)
			body.Reset()
			body.WriteString(`{"foo" : "bar `)
			body.WriteString(docID)
			body.WriteString(`	" }`)
			req.DocumentID = docID
			req.Body = strings.NewReader(body.String())

			if _, err := req.Do(context.Background(), es); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("client.Search()                   ", func(b *testing.B) {
		body := strings.NewReader(`{"query" : { "match_all" : {} } }`)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := es.Search(es.Search.WithContext(context.Background()), es.Search.WithBody(body))

			if err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("client.Search() with error        ", func(b *testing.B) {
		body := strings.NewReader(`{}`)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := eserr.Search(eserr.Search.WithContext(context.Background()), eserr.Search.WithBody(body))

			if err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("SearchRequest{...}.Do()            ", func(b *testing.B) {
		body := strings.NewReader(`{"query" : { "match_all" : {} } }`)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			req := esapi.SearchRequest{Body: body}
			if _, err := req.Do(context.Background(), es); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("SearchRequest{...}.Do() with error ", func(b *testing.B) {
		body := strings.NewReader(`{}`)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			req := esapi.SearchRequest{Body: body}
			if _, err := req.Do(context.Background(), eserr); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})
}
