package esutil

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

// BulkIndexer represents a parallel, asynchronous indexer.
//
type BulkIndexer interface {
	// Add adds an item to the indexer. It returns an error when the item cannot be added.
	// Use the OnSuccess and OnFailure callbacks to get the operation result for the item.
	// You must call the Wait() method to allow the indexer finish the indexing.
	// It is safe for concurrent use, but all goroutines must finish before the call to Wait().
	Add(context.Context, BulkIndexerItem) error

	// Wait blocks until all added items are indexed.
	Wait(context.Context) error

	NumAdded() int   // NumAdded returns the number of items added to the indexer.
	NumFailed() int  // NumFailed returns the number of failed items.
	NumIndexed() int // NumCreated returns the number of succesfull "index" operations.
	NumCreated() int // NumCreated returns the number of succesfull "create" operations.
	NumUpdated() int // NumUpdated returns the number of succesfull "update" operations.
	NumDeleted() int // NumDeleted returns the number of succesfull "delete" operations.
}

type bulkIndexer struct {
	mu      sync.Mutex
	wg      sync.WaitGroup
	queue   chan BulkIndexerItem
	workers []*worker

	client *elasticsearch.Client

	numWorkers          int
	flushThresholdBytes int

	numAdded   int
	numFailed  int
	numIndexed int
	numCreated int
	numUpdated int
	numDeleted int
}

// BulkIndexerConfig represents configuration of the bulk indexer.
//
type BulkIndexerConfig struct {
	Client *elasticsearch.Client // The Elasticsearch client.

	NumWorkers          int // The number of workers. Defaults to runtime.NumCPU().
	FlushThresholdBytes int // The flush threshold in bytes. Defaults to 5MB.

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

	if cfg.FlushThresholdBytes == 0 {
		cfg.FlushThresholdBytes = 5e+6
	}

	bi := bulkIndexer{
		numWorkers:          cfg.NumWorkers,
		flushThresholdBytes: cfg.FlushThresholdBytes,
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

	Meta interface{} // To use eg. in OnSuccess/OnFailure callbacks

	OnSuccess func(item BulkIndexerItem, req interface{}, res interface{})            // Per item
	OnFailure func(item BulkIndexerItem, req interface{}, res interface{}, err error) // Per item
}

func (bi *bulkIndexer) Add(ctx context.Context, item BulkIndexerItem) error {
	bi.mu.Lock()
	defer bi.mu.Unlock()

	bi.numAdded++

	select {
	case <-ctx.Done():
		return ctx.Err()
	case bi.queue <- item:
	}

	return nil
}

func (bi *bulkIndexer) Wait(ctx context.Context) error {
	close(bi.queue)
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		bi.wg.Wait()
	}

	fmt.Println("==> Flushing remaining buffers...")
	for _, w := range bi.workers {
		if w.buf.Len() > 0 {
			if err := w.flush(); err != nil {
				return fmt.Errorf("flush: %s", err)
			}
		}
	}

	return nil
}

func (bi *bulkIndexer) NumAdded() int   { bi.mu.Lock(); defer bi.mu.Unlock(); return bi.numAdded }
func (bi *bulkIndexer) NumFailed() int  { bi.mu.Lock(); defer bi.mu.Unlock(); return bi.numFailed }
func (bi *bulkIndexer) NumIndexed() int { bi.mu.Lock(); defer bi.mu.Unlock(); return bi.numIndexed }
func (bi *bulkIndexer) NumCreated() int { bi.mu.Lock(); defer bi.mu.Unlock(); return bi.numCreated }
func (bi *bulkIndexer) NumUpdated() int { bi.mu.Lock(); defer bi.mu.Unlock(); return bi.numUpdated }
func (bi *bulkIndexer) NumDeleted() int { bi.mu.Lock(); defer bi.mu.Unlock(); return bi.numDeleted }

// init initializes the bulk indexer.
//
func (bi *bulkIndexer) init() {
	bi.queue = make(chan BulkIndexerItem, bi.numWorkers)

	for i := 1; i <= bi.numWorkers; i++ {
		w := worker{
			id:  i,
			ch:  bi.queue,
			wg:  &bi.wg,
			buf: bytes.NewBuffer(make([]byte, 0, bi.flushThresholdBytes)),
			byt: bi.flushThresholdBytes}
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

// run launches the worker goroutine.
//
func (w *worker) run() {
	go func() {
		fmt.Printf("--> [worker-%03d] Started\n", w.id)
		defer w.wg.Done()

		for item := range w.ch {
			fmt.Printf(">>> [worker-%03d] Got %s:%s\n", w.id, item.Action, item.DocumentID)
			w.items = append(w.items, item)

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
				w.items = w.items[:0]
			}
		}
	}()
}

// flush writes out the buffer.
func (w *worker) flush() error {
	if w.isFlushing {
		fmt.Printf(">>> [worker-%03d] Already flushing\n", w.id)
		return nil
	}
	w.isFlushing = true

	fmt.Printf(">>> [worker-%03d] FLUSH BUFFER: %s\n", w.id, w.buf.String())

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

	w.buf.Reset()
	w.isFlushing = false

	return err
}
