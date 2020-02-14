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
	// Add adds an item to the indexer. It is asynchronous, and returns an error
	// only when the item cannot be added. Use the OnSuccess and OnFailure
	// callbacks to get the operation result per item.
	Add(item BulkIndexerItem) error

	// Wait blocks until all added items are indexed.
	Wait(ctx context.Context) error

	NumAdded() uint   // NumAdded returns the number of items added to the indexer.
	NumFailed() uint  // NumFailed returns the number of failed items.
	NumCreated() uint // NumCreated returns the number of succesfull "create" operations.
	NumUpdated() uint // NumUpdated returns the number of succesfull "update" operations.
	NumDeleted() uint // NumDeleted returns the number of succesfull "delete" operations.
}

// bulkIndexer is a parallel, asynchronous indexer.
//
type bulkIndexer struct {
	mu    sync.Mutex
	wg    sync.WaitGroup
	queue chan BulkIndexerItem

	client *elasticsearch.Client

	numWorkers          int
	flushThresholdItems uint
	flushThresholdBytes uint

	numAdded   uint
	numCreated uint
	numUpdated uint
	numDeleted uint
	numFailed  uint

	// ?: Cancellation, timeouts? context.WithCancel() / context.WithTimeout()?
}

// BulkIndexerConfig represents configuration of the bulk indexer.
//
type BulkIndexerConfig struct {
	Client *elasticsearch.Client // The Elasticsearch client.

	NumWorkers          int  // The number of workers. Defaults to runtime.NumCPU().
	FlushThresholdItems uint // The flush threshold for items. Defaults to 1000.
	FlushThresholdBytes uint // The flush threshold in bytes. Defaults 5MB.
}

// NewBulkIndexer creates a new bulk indexer.
//
func NewBulkIndexer(cfg BulkIndexerConfig) (BulkIndexer, error) {
	if cfg.NumWorkers == 0 {
		cfg.NumWorkers = runtime.NumCPU()
	}

	if cfg.FlushThresholdItems == 0 {
		cfg.FlushThresholdItems = 1000
	}

	if cfg.FlushThresholdBytes == 0 {
		cfg.FlushThresholdBytes = 5e+6
	}

	bi := bulkIndexer{
		numWorkers:          cfg.NumWorkers,
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

	OnSuccess func(item BulkIndexerItem, req interface{}, res interface{})            // Per item
	OnFailure func(item BulkIndexerItem, req interface{}, res interface{}, err error) // Per item
}

func (bi *bulkIndexer) Add(item BulkIndexerItem) error {
	bi.mu.Lock()
	defer bi.mu.Unlock()

	bi.numAdded++
	bi.queue <- item

	return nil
}

func (bi *bulkIndexer) Wait(ctx context.Context) error {
	close(bi.queue)
	bi.wg.Wait()
	return nil
}

func (bi *bulkIndexer) NumAdded() uint   { return bi.numAdded }
func (bi *bulkIndexer) NumFailed() uint  { return bi.numFailed }
func (bi *bulkIndexer) NumCreated() uint { return bi.numCreated }
func (bi *bulkIndexer) NumUpdated() uint { return bi.numUpdated }
func (bi *bulkIndexer) NumDeleted() uint { return bi.numDeleted }

// init initializes the bulk indexer.
//
func (bi *bulkIndexer) init() {
	bi.queue = make(chan BulkIndexerItem, bi.numWorkers)

	for i := 0; i < bi.numWorkers; i++ {
		w := worker{ch: bi.queue, wg: &bi.wg, buf: bytes.NewBuffer(make([]byte, 0, bi.flushThresholdBytes))}
		w.run()
	}
	bi.wg.Add(bi.numWorkers)
}

// worker represents an indexer worker.
//
type worker struct {
	ch  <-chan BulkIndexerItem
	wg  *sync.WaitGroup
	buf *bytes.Buffer
}

// run launches the worker goroutine.
//
func (w *worker) run() {
	go func() {
		fmt.Println("Started worker")
		defer w.wg.Done()

		for item := range w.ch {
			fmt.Println("Got item", item)
		}
	}()
}
