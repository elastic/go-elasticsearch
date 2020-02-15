// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package esutil

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

// BulkIndexer represents a parallel, asynchronous, bulk indexer.
//
type BulkIndexer interface {
	// Add adds an item to the indexer. It returns an error when the item cannot be added.
	// Use the OnSuccess and OnFailure callbacks to get the operation result for the item.
	//
	// You must call the Close() method after you're done adding items.
	//
	// It is safe for concurrent use. When it's called from goroutines,
	// they must finish before the call to Close, eg. using sync.WaitGroup.
	Add(context.Context, BulkIndexerItem) error

	// Close waits until all added items are flushed and closes the indexer.
	Close(context.Context) error

	// Stats returns indexer statistics.
	Stats() BulkIndexerStats
}

// BulkIndexerStats represents the indexer statistics.
//
type BulkIndexerStats struct {
	NumAdded   uint
	NumFailed  uint
	NumIndexed uint
	NumCreated uint
	NumUpdated uint
	NumDeleted uint
}

type bulkIndexer struct {
	wg      sync.WaitGroup
	queue   chan BulkIndexerItem
	workers []*worker
	stats   *bulkIndexerStats

	client *elasticsearch.Client

	numWorkers int
	flushBytes int
}

type bulkIndexerStats struct {
	numAdded   uint64
	numFailed  uint64
	numIndexed uint64
	numCreated uint64
	numUpdated uint64
	numDeleted uint64
}

// BulkIndexerConfig represents configuration of the indexer.
//
type BulkIndexerConfig struct {
	NumWorkers int // The number of workers. Defaults to runtime.NumCPU().
	FlushBytes int // The flush threshold in bytes. Defaults to 5MB.

	Client *elasticsearch.Client // The Elasticsearch client.

	// Parameters of the Bulk API.
	Index               string
	ErrorTrace          bool
	FilterPath          []string
	Header              http.Header
	Human               bool
	Pipeline            string
	Pretty              bool
	Refresh             string
	Routing             string
	Source              []string
	SourceExcludes      []string
	SourceIncludes      []string
	Timeout             time.Duration
	WaitForActiveShards string
}

// NewBulkIndexer creates a new bulk indexer.
//
func NewBulkIndexer(cfg BulkIndexerConfig) (BulkIndexer, error) {
	if cfg.NumWorkers == 0 {
		cfg.NumWorkers = runtime.NumCPU()
	}

	if cfg.FlushBytes == 0 {
		cfg.FlushBytes = 5e+6
	}

	bi := bulkIndexer{
		numWorkers: cfg.NumWorkers,
		flushBytes: cfg.FlushBytes,

		stats: &bulkIndexerStats{},
	}

	bi.init()

	return &bi, nil
}

// BulkIndexerItem represents an indexer item.
//
type BulkIndexerItem struct {
	Index           string
	Action          string
	DocumentID      string
	Body            io.Reader
	RetryOnConflict *int

	Metadata interface{} // To use eg. in OnSuccess/OnFailure callbacks

	OnSuccess func(item BulkIndexerItem, req interface{}, res interface{})            // Per item
	OnFailure func(item BulkIndexerItem, req interface{}, res interface{}, err error) // Per item
}

// Add adds an item to the indexer.
//
func (bi *bulkIndexer) Add(ctx context.Context, item BulkIndexerItem) error {
	atomic.AddUint64(&bi.stats.numAdded, 1)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case bi.queue <- item:
	}

	return nil
}

// Close calls flush on writers and closes the indexer queue channel.
//
func (bi *bulkIndexer) Close(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		close(bi.queue)
		bi.wg.Wait()
	}

	if os.Getenv("DEBUG") != "" {
		fmt.Println("==> Flushing remaining buffers...")
	}
	for _, w := range bi.workers {
		if w.buf.Len() > 0 {
			if err := w.flush(); err != nil {
				return fmt.Errorf("flush: %s", err)
			}
		}
	}

	return nil
}

// Stats returns indexer statistics.
//
func (bi *bulkIndexer) Stats() BulkIndexerStats {
	return BulkIndexerStats{
		NumAdded:   uint(bi.stats.numAdded),
		NumFailed:  uint(bi.stats.numFailed),
		NumIndexed: uint(bi.stats.numIndexed),
		NumCreated: uint(bi.stats.numCreated),
		NumUpdated: uint(bi.stats.numUpdated),
		NumDeleted: uint(bi.stats.numDeleted),
	}
}

// init initializes the bulk indexer.
//
func (bi *bulkIndexer) init() {
	bi.queue = make(chan BulkIndexerItem, bi.numWorkers)

	for i := 1; i <= bi.numWorkers; i++ {
		w := worker{
			id:  i,
			ch:  bi.queue,
			wg:  &bi.wg,
			buf: bytes.NewBuffer(make([]byte, 0, bi.flushBytes)),
			byt: bi.flushBytes}
		w.run()
		bi.workers = append(bi.workers, &w)
	}
	bi.wg.Add(bi.numWorkers)
}

// worker represents an indexer worker.
//
type worker struct {
	id    int
	ch    <-chan BulkIndexerItem
	wg    *sync.WaitGroup
	buf   *bytes.Buffer
	byt   int
	items []BulkIndexerItem

	isFlushing bool
}

// run launches the worker in a goroutine.
//
func (w *worker) run() {
	go func() {
		if os.Getenv("DEBUG") != "" {
			fmt.Printf("--> [worker-%03d] Started\n", w.id)
		}
		defer w.wg.Done()

		for item := range w.ch {
			if os.Getenv("DEBUG") != "" {
				fmt.Printf(">>> [worker-%03d] Got %s:%s\n", w.id, item.Action, item.DocumentID)
			}
			w.items = append(w.items, item)

			w.buf.WriteString(item.Action)
			w.buf.WriteRune('/')
			if item.DocumentID != "" {
				w.buf.WriteString(item.DocumentID)
			}
			w.buf.WriteRune(':')

			if item.Body != nil {
				if _, err := w.buf.ReadFrom(item.Body); err != nil {
					if item.OnFailure != nil {
						item.OnFailure(item, nil, nil, err)
					}
					continue
				}
				w.buf.WriteRune(',')
			}

			if w.buf.Len() >= w.byt {
				w.flush()
			}
		}
	}()
}

// flush writes out the worker buffer.
//
func (w *worker) flush() error {
	if w.isFlushing {
		if os.Getenv("DEBUG") != "" {
			fmt.Printf(">>> [worker-%03d] Already flushing\n", w.id)
		}
		return nil
	}
	w.isFlushing = true

	// Simulate copy operation
	var b bytes.Buffer
	b.Grow(w.buf.Len())
	io.Copy(&b, w.buf)
	// -----------------------
	if os.Getenv("DEBUG") != "" {
		fmt.Printf(">>> [worker-%03d] FLUSH BUFFER: %s\n", w.id, b.String())
	}

	var err error
	if err == nil {
		for _, item := range w.items {
			if item.OnSuccess != nil {
				item.OnSuccess(item, nil, nil)
			}
		}
	} else {
		for _, item := range w.items {
			if item.OnFailure != nil {
				item.OnFailure(item, nil, nil, fmt.Errorf("flush: %s", err))
			}
		}
	}

	w.items = w.items[:0]
	w.buf.Reset()
	w.isFlushing = false

	return err
}
