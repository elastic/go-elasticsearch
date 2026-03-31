package elasticsearch

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

type readCloser struct {
	io.Reader
	closed bool
}

func (rc *readCloser) Close() error {
	rc.closed = true
	return nil
}

func TestAutoDrainingReader(t *testing.T) {
	content := strings.NewReader("hello elasticsearch")
	wrapped := &autoDrainingReader{ReadCloser: &readCloser{Reader: content}}
	buf := make([]byte, 5)
	n, err := wrapped.Read(buf)
	if err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if n != 5 {
		t.Errorf("expected 5 bytes, got %d", n)
	}
	if string(buf[:n]) != "hello" {
		t.Errorf("expected 'hello', got '%s'", string(buf[:n]))
	}
	if err := wrapped.Close(); err != nil {
		t.Fatalf("Close error: %v", err)
	}
}

func TestAutoDrainingReaderDrainsOnClose(t *testing.T) {
	content := bytes.Buffer{}
	content.Grow(100)
	content.Write(make([]byte, 100))
	wrapped := &autoDrainingReader{
		ReadCloser: &readCloser{Reader: &content},
	}
	n, err := wrapped.Read(make([]byte, 10))
	if err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if n != 10 {
		t.Errorf("expected 10 bytes, got %d", n)
	}
	if err := wrapped.Close(); err != nil {
		t.Fatalf("Close error: %v", err)
	}
	if content.Len() != 0 {
		t.Errorf("expected content to be drained, remaining: %d bytes", content.Len())
	}
}

func TestAutoDrainingReadWriter(t *testing.T) {
	content := strings.NewReader("hello elasticsearch")
	rwc := &readWriteCloser{
		Reader: content,
		Closer: &readCloser{Reader: content},
	}
	wrapped := &autoDrainingReadWriter{ReadWriteCloser: rwc}
	buf := make([]byte, 5)
	n, err := wrapped.Read(buf)
	if err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if n != 5 {
		t.Errorf("expected 5 bytes, got %d", n)
	}
	if err := wrapped.Close(); err != nil {
		t.Fatalf("Close error: %v", err)
	}
}

func TestMakeAutoDrainingReadCloser(t *testing.T) {
	content := strings.NewReader("elasticsearch")
	rc := &readCloser{Reader: content}
	wrapped := makeAutoDraining(rc)
	if _, ok := wrapped.(*autoDrainingReader); !ok {
		t.Errorf("expected *autoDrainingReader, got %T", wrapped)
	}
	if err := wrapped.Close(); err != nil {
		t.Fatalf("Close error: %v", err)
	}
}

func TestMakeAutoDrainingReadWriteCloser(t *testing.T) {
	content := strings.NewReader("elasticsearch")
	rc := &readCloser{Reader: content}
	wrapped := makeAutoDraining(rc)
	if _, ok := wrapped.(*autoDrainingReader); !ok {
		t.Errorf("expected *autoDrainingReader, got %T", wrapped)
	}
	if err := wrapped.Close(); err != nil {
		t.Fatalf("Close error: %v", err)
	}
}

type readWriteCloser struct {
	io.Reader
	io.Writer
	io.Closer
}

func TestAutoDrainingReadWriterWithReadWriteCloser(t *testing.T) {
	content := strings.NewReader("hello elasticsearch")
	rwc := &readWriteCloser{
		Reader: content,
		Closer: &readCloser{Reader: content},
	}
	wrapped := makeAutoDraining(rwc)
	if _, ok := wrapped.(*autoDrainingReadWriter); !ok {
		t.Errorf("expected *autoDrainingReadWriter, got %T", wrapped)
	}
	if err := wrapped.Close(); err != nil {
		t.Fatalf("Close error: %v", err)
	}
}

func TestAutoDrainBodyEnablesConnectionReuse(t *testing.T) {
	var connectionCount int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&connectionCount, 1)
		w.Header().Set("X-Elastic-Product", "Elasticsearch")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"cluster_name":"test"}`))
	}))
	defer ts.Close()
	es, err := NewClient(Config{
		Addresses:     []string{ts.URL},
		AutoDrainBody: true,
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	for range 5 {
		func() {
			res, err := es.API.Info()
			if err != nil {
				t.Fatalf("Request failed: %v", err)
			}
			defer res.Body.Close()
			if res.IsError() {
				t.Errorf("Request returned error: %s", res.String())
			}
		}()
	}
	if connectionCount != 5 {
		t.Errorf("Expected 5 requests, got %d", connectionCount)
	}
}

func TestAutoDrainBodyWithoutDrain(t *testing.T) {
	var connectionCount int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&connectionCount, 1)
		w.Header().Set("X-Elastic-Product", "Elasticsearch")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"cluster_name":"test"}`))
	}))
	defer ts.Close()
	es, err := NewClient(Config{
		Addresses:     []string{ts.URL},
		AutoDrainBody: false,
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	for range 5 {
		func() {
			res, err := es.API.Info()
			if err != nil {
				t.Fatalf("Request failed: %v", err)
			}
			defer res.Body.Close()
			if res.IsError() {
				t.Errorf("Request returned error: %s", res.String())
			}
		}()
	}
	if connectionCount != 5 {
		t.Errorf("Expected 5 requests, got %d", connectionCount)
	}
}

func TestAutoDrainBodyWithPartialRead(t *testing.T) {
	var drainCalled atomic.Bool
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Elastic-Product", "Elasticsearch")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"cluster_name":"test","test_data":"hello elasticsearch this is a longer response"}`))
	}))
	defer ts.Close()
	es, err := NewClient(Config{
		Addresses:     []string{ts.URL},
		AutoDrainBody: true,
	})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	for range 3 {
		res, err := es.API.Info()
		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}
		buf := make([]byte, 10)
		n, err := res.Body.Read(buf)
		if err != nil && err != io.EOF {
			t.Fatalf("Read error: %v", err)
		}
		if n == 0 {
			t.Error("Expected to read some data")
		}
		if err := res.Body.Close(); err != nil {
			t.Fatalf("Close error: %v", err)
		}
		drainCalled.Store(true)
	}
	if !drainCalled.Load() {
		t.Error("Expected drain to be called")
	}
}
