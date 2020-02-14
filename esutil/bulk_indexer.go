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

// BulkIndexer is a parallel, asynchronous indexer.
//
type BulkIndexer struct {
	mu    sync.Mutex
	wg    sync.WaitGroup
	queue chan BulkIndexerItem

	client *elasticsearch.Client

	numWorkers          int
	flushInterval       time.Duration
	flushThresholdItems uint
	flushThresholdBytes uint

	numAdded   uint
	numCreated uint
	numUpdated uint
	numDeleted uint
	numFailed  uint

	// ?: Cancellation, timeouts? context.WithCancel() / context.WithTimeout()?
}

type BulkIndexerConfig struct {
	Client *elasticsearch.Client // The Elasticsearch client.

	NumWorkers          int           // The number of workers. Defaults to runtime.NumCPU().
	FlushInterval       time.Duration // The flush interval. Defaults to 30s.
	FlushThresholdItems uint          // The flush threshold for items. Defaults to 1000.
	FlushThresholdBytes uint          // The flush threshold in bytes. Defaults 5MB.
}

// NewBulkIndexer creates a new bulk indexer.
//
func NewBulkIndexer(cfg BulkIndexerConfig) (*BulkIndexer, error) {
	if cfg.NumWorkers == 0 {
		cfg.NumWorkers = runtime.NumCPU()
	}

	if cfg.FlushInterval == 0 {
		cfg.FlushInterval = 30 * time.Second
	}

	if cfg.FlushThresholdItems == 0 {
		cfg.FlushThresholdItems = 1000
	}

	if cfg.FlushThresholdBytes == 0 {
		cfg.FlushThresholdBytes = 5e+6
	}

	bi := BulkIndexer{
		numWorkers:          cfg.NumWorkers,
		flushInterval:       cfg.FlushInterval,
		flushThresholdItems: cfg.FlushThresholdItems,
		flushThresholdBytes: cfg.FlushThresholdBytes,
	}

	bi.init()

	return &bi, nil
}

// BulkIndexerItem represents an indexer item.
//
type BulkIndexerItem struct {
	Index  string
	Action string
	Body   io.Reader

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

	Meta interface{} // To use eg. in OnSuccess/OnFailure callbacks

	OnSuccess func(req interface{}, res interface{})            // Per item
	OnFailure func(req interface{}, res interface{}, err error) // Per item
}

// Add adds an item to the indexer. It is asynchronous, and returns an error
// only when the item cannot be added. Use the OnSuccess and OnFailure
// callbacks to get the operation result.
//
func (bi *BulkIndexer) Add(item BulkIndexerItem) error {
	bi.mu.Lock()
	defer bi.mu.Unlock()

	bi.numAdded++
	bi.queue <- item

	return nil
}

// Flush forcefully flushes all added items.
//
func (bi *BulkIndexer) Flush(ctx context.Context) error { return nil }

// Wait blocks until all added items are flushed.
//
func (bi *BulkIndexer) Wait(ctx context.Context) error {
	close(bi.queue)
	bi.wg.Wait()
	return nil
}

// NumAdded returns the number of items added to the indexer.
func (bi *BulkIndexer) NumAdded() uint { return bi.numAdded }

// NumFailed returns the number of failed items.
func (bi *BulkIndexer) NumFailed() uint { return bi.numFailed }

// NumCreated returns the number of succesfull "create" operations.
func (bi *BulkIndexer) NumCreated() uint { return bi.numCreated }

// NumUpdated returns the number of succesfull "update" operations.
func (bi *BulkIndexer) NumUpdated() uint { return bi.numUpdated }

// NumDeleted returns the number of succesfull "delete" operations.
func (bi *BulkIndexer) NumDeleted() uint { return bi.numDeleted }

// worker represents an indexer worker.
//
type worker struct {
	ch  <-chan BulkIndexerItem
	wg  *sync.WaitGroup
	buf *bytes.Buffer
}

func (bi *BulkIndexer) init() {
	bi.queue = make(chan BulkIndexerItem, bi.numWorkers)

	for i := 0; i < bi.numWorkers; i++ {
		w := worker{ch: bi.queue, wg: &bi.wg, buf: bytes.NewBuffer(make([]byte, 0, bi.flushThresholdBytes))}
		w.run()
	}
	bi.wg.Add(bi.numWorkers)
}

func (w *worker) run() {
	go func() {
		fmt.Println("Started worker")
		defer w.wg.Done()

		for item := range w.ch {
			fmt.Println("Got item", item)
		}
	}()
}
