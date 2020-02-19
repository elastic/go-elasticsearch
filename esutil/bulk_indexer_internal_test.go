// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package esutil

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

var defaultRoundTripFunc = func(*http.Request) (*http.Response, error) {
	return &http.Response{Body: ioutil.NopCloser(strings.NewReader(`{}`))}, nil
}

type mockTransport struct {
	RoundTripFunc func(*http.Request) (*http.Response, error)
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.RoundTripFunc == nil {
		return defaultRoundTripFunc(req)
	}
	return t.RoundTripFunc(req)
}

func TestBulkIndexer(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		var wg sync.WaitGroup

		es, _ := elasticsearch.NewClient(elasticsearch.Config{Transport: &mockTransport{}})
		bi, _ := NewBulkIndexer(BulkIndexerConfig{NumWorkers: 2, FlushBytes: 50, Client: es})
		numItems := 3

		for i := 1; i <= numItems; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				err := bi.Add(context.Background(), BulkIndexerItem{
					Action:     "index",
					DocumentID: strconv.Itoa(i),
					Body:       strings.NewReader(fmt.Sprintf(`{"title":"foo-%d"}`, i)),
				})
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}(i)
		}
		wg.Wait()

		if err := bi.Close(context.Background()); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		stats := bi.Stats()

		if stats.NumAdded != uint(numItems) {
			t.Errorf("Unexpected NumAdded: %d", stats.NumAdded)
		}

		if stats.NumFailed != 0 {
			t.Errorf("Unexpected NumFailed: %d", stats.NumFailed)
		}

		fmt.Println("NumAdded:", stats.NumAdded, "NumFailed:", stats.NumFailed)
	})

	t.Run("Add() Timeout", func(t *testing.T) {
		es, _ := elasticsearch.NewClient(elasticsearch.Config{Transport: &mockTransport{}})
		bi, _ := NewBulkIndexer(BulkIndexerConfig{NumWorkers: 1, Client: es})
		ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
		defer cancel()

		var errs []error
		for i := 0; i < 10; i++ {
			errs = append(errs, bi.Add(ctx, BulkIndexerItem{Action: "delete", DocumentID: "timeout"}))
		}
		if err := bi.Close(context.Background()); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		var gotError bool
		for _, err := range errs {
			if err != nil && err.Error() == "context deadline exceeded" {
				gotError = true
			}
		}
		if !gotError {
			t.Errorf("Expected timeout error, but none in: %s", errs)
		}
	})

	t.Run("Callbacks", func(t *testing.T) {
		var (
			countSuccessful uint64
			countFailed     uint64
			failedIDs       []string

			numItems       = 4
			numFailed      = 2
			bodyContent, _ = ioutil.ReadFile("testdata/bulk_response_1.json")
		)

		es, _ := elasticsearch.NewClient(elasticsearch.Config{Transport: &mockTransport{
			RoundTripFunc: func(*http.Request) (*http.Response, error) {
				return &http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(bodyContent))}, nil
			},
		}})
		bi, _ := NewBulkIndexer(BulkIndexerConfig{NumWorkers: 1, Client: es})

		successFunc := func(item BulkIndexerItem, res BulkIndexerResponseItem) {
			atomic.AddUint64(&countSuccessful, 1)
		}
		failureFunc := func(item BulkIndexerItem, res BulkIndexerResponseItem, err error) {
			atomic.AddUint64(&countFailed, 1)
			failedIDs = append(failedIDs, item.DocumentID)
		}

		if err := bi.Add(context.Background(), BulkIndexerItem{
			Action:     "index",
			DocumentID: "1",
			Body:       strings.NewReader(`{"title":"foo"}`),
			OnSuccess:  successFunc,
			OnFailure:  failureFunc,
		}); err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if err := bi.Add(context.Background(), BulkIndexerItem{
			Action:     "create",
			DocumentID: "1",
			Body:       strings.NewReader(`{"title":"bar"}`),
			OnSuccess:  successFunc,
			OnFailure:  failureFunc,
		}); err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if err := bi.Add(context.Background(), BulkIndexerItem{
			Action:     "delete",
			DocumentID: "2",
			OnSuccess:  successFunc,
			OnFailure:  failureFunc,
		}); err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if err := bi.Add(context.Background(), BulkIndexerItem{
			Action:     "update",
			DocumentID: "3",
			Body:       strings.NewReader(`{"doc":{"title":"qux"}}`),
			OnSuccess:  successFunc,
			OnFailure:  failureFunc,
		}); err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if err := bi.Close(context.Background()); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		stats := bi.Stats()

		if stats.NumAdded != uint(numItems) {
			t.Errorf("Unexpected NumAdded: %d", stats.NumAdded)
		}

		// Two failures are expected:
		//
		// * Operation #2: document can't be created, because a document with the same ID already exists.
		// * Operation #3: document can't be deleted, because it doesn't exist.

		if stats.NumFailed != uint(numFailed) {
			t.Errorf("Unexpected NumFailed: %d", stats.NumFailed)
		}

		if countSuccessful != uint64(numItems-numFailed) {
			t.Errorf("Unexpected countSuccessful: %d", countSuccessful)
		}

		if countFailed != uint64(numFailed) {
			t.Errorf("Unexpected countFailed: %d", countFailed)
		}

		if !reflect.DeepEqual(failedIDs, []string{"1", "2"}) {
			t.Errorf("Unexpected failedIDs: %#v", failedIDs)
		}
	})
}
