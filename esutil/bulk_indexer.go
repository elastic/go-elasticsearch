package esutil

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

// BulkIndexer is a parallel, asynchronous indexer.
//
type BulkIndexer struct {
	sync.Mutex

	Client *elasticsearch.Client

	FlushInterval       time.Duration // Default 30sec
	FlushThresholdItems uint          // Default 1000
	FlushThresholdBytes uint          // Default 5MB

	NumWorkers int // Semaphore, default to runtime.NumCPU

	numAdded   uint
	numCreated uint
	numUpdated uint
	numDeleted uint
	numFailed  uint

	workers workerPool

	// ?: Cancellation, timeouts? context.WithCancel() / context.WithTimeout()?
}

// NewBulkIndexer creates a new bulk indexer.
//
func NewBulkIndexer() *BulkIndexer {
	bi := BulkIndexer{}

	bufPool := sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, bi.FlushThresholdBytes))
		},
	}

	for i := 0; i < bi.NumWorkers; i++ {
		bi.workers = append(bi.workers, worker{bufPool: bufPool})
	}

	return &bi
}

// workerPool represents the pool of workers.
//
type workerPool []worker

// worker represents an indexer worker.
//
type worker struct {
	bufPool sync.Pool // (sync.Pool).Get().(*bytes.Buffer)
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
	bi.Lock()
	defer bi.Unlock()

	bi.numAdded++

	return nil
}

// Flush forcefully flushes all added items.
//
func (bi *BulkIndexer) Flush(ctx context.Context) error { return nil }

// Wait blocks until all added items are flushed.
//
func (bi *BulkIndexer) Wait(ctx context.Context) error { return nil }

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
