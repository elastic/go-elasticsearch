package esutil

import (
	"bytes"
	"io"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

// ?: `BulkProcessor` to mirror Java?
//
type BulkIndexer struct {
	Client *elasticsearch.Client

	BatchSuccessCallback func(req interface{}, res interface{})            // Per batch
	BatchErrorCallback   func(req interface{}, res interface{}, err error) // Per batch

	FlushInterval       time.Duration // Default 30sec
	FlushThresholdItems uint          // Default 1000
	FlushThresholdBytes uint          // Default 5MB

	numPublished uint
	numCreated   uint
	numUpdated   uint
	numDeleted   uint
	numFailed    uint

	buffer *bytes.Buffer // sync.Pool.Get(), make(), default capacity = FlushThresholdBytes (5MB)
	// ?: Or rather a queue object, and keep buffer local during execution?
	queue *struct { // BulkIndexerQueue
		mu    sync.Mutex
		queue []*BulkIndexerItem
	}

	// ?: Cancellation, timeouts? context.WithCancel() / context.WithTimeout()?

	// ?: ConcurrentRequests (=> semaphore)
}

type BulkIndexerItem struct {
	Action string
	Body   io.Reader

	Meta interface{} // ?: Use (context.Context).Value instead?

	SuccessCallback func(req interface{}, res interface{})            // Per item (Needed?)
	ErrorCallback   func(req interface{}, res interface{}, err error) // Per item (Needed?)
}

// ?: Variadic arguments? More flexible.
// ?: Add(interface{}) and then type-switch over eg. `esapi.IndexRequest{}`?
//
// !: Must be threadsafe
//
func (bi BulkIndexer) Add(item *BulkIndexerItem) {
	// Asynchronous, use callback for getting operation result
	// Add item to the indexer buffer and return
}

// ?: Should be exposed? Java `BulkProcessor` does. Blocking operation.
func (bi BulkIndexer) Flush() error { return nil }

// ?: Wait for operations to finish, similar to Java's `awaitClose()`. Blocking operation.
func (bi BulkIndexer) Wait() error { return nil }

func (bi BulkIndexer) NumPublished() uint { return bi.numPublished }
func (bi BulkIndexer) NumFailed() uint    { return bi.numFailed }
