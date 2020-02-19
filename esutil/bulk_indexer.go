// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package esutil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

// BulkIndexer represents a parallel, asynchronous, efficient indexer for Elasticsearch.
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

// BulkIndexerConfig represents configuration of the indexer.
//
type BulkIndexerConfig struct {
	NumWorkers int // The number of workers. Defaults to runtime.NumCPU().
	FlushBytes int // The flush threshold in bytes. Defaults to 5MB.

	Client  *elasticsearch.Client   // The Elasticsearch client.
	Decoder BulkResponseJSONDecoder // A custom JSON decoder.

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

// BulkIndexerStats represents the indexer statistics.
//
type BulkIndexerStats struct {
	NumAdded   uint
	NumFlushed uint
	NumFailed  uint
	NumIndexed uint
	NumCreated uint
	NumUpdated uint
	NumDeleted uint
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

	OnSuccess func(item BulkIndexerItem, res BulkIndexerResponseItem)            // Per item
	OnFailure func(item BulkIndexerItem, res BulkIndexerResponseItem, err error) // Per item
}

// BulkIndexerResponse represents the Elasticsearch response.
//
type BulkIndexerResponse struct {
	Took      int                                  `json:"took"`
	HasErrors bool                                 `json:"errors"`
	Items     []map[string]BulkIndexerResponseItem `json:"items,omitempty"`
}

// BulkIndexerResponseItem represents the Elasticsearch response item.
//
type BulkIndexerResponseItem struct {
	Index      string `json:"_index"`
	DocumentID string `json:"_id"`
	Version    int64  `json:"_version,omitempty"`
	Result     string `json:"result,omitempty"`
	Status     int    `json:"status,omitempty"`
	SeqNo      int64  `json:"_seq_no,omitempty"`
	PrimTerm   int64  `json:"_primary_term,omitempty"`

	Shards struct {
		Total      int `json:"total,omitempty"`
		Successful int `json:"successful,omitempty"`
		Failed     int `json:"failed,omitempty"`
	} `json:"_shards,omitempty"`

	Error struct {
		Type   string `json:"type"`
		Reason string `json:"reason"`
		Cause  struct {
			Type   string `json:"type"`
			Reason string `json:"reason"`
		} `json:"caused_by"`
	} `json:"error,omitempty"`
}

// BulkResponseJSONDecoder defines the interface for custom JSON decoders.
//
type BulkResponseJSONDecoder interface {
	UnmarshalFromReader(io.Reader, *BulkIndexerResponse) error
}

type bulkIndexer struct {
	wg      sync.WaitGroup
	queue   chan BulkIndexerItem
	workers []*worker
	stats   *bulkIndexerStats

	config BulkIndexerConfig
}

type bulkIndexerStats struct {
	numAdded   uint64
	numFlushed uint64
	numFailed  uint64
	numIndexed uint64
	numCreated uint64
	numUpdated uint64
	numDeleted uint64
}

// NewBulkIndexer creates a new bulk indexer.
//
func NewBulkIndexer(cfg BulkIndexerConfig) (BulkIndexer, error) {
	if cfg.Client == nil {
		cfg.Client, _ = elasticsearch.NewDefaultClient()
	}

	if cfg.Decoder == nil {
		cfg.Decoder = defaultJSONDecoder{}
	}

	if cfg.NumWorkers == 0 {
		cfg.NumWorkers = runtime.NumCPU()
	}

	if cfg.FlushBytes == 0 {
		cfg.FlushBytes = 5e+6
	}

	bi := bulkIndexer{
		config: cfg,
		stats:  &bulkIndexerStats{},
	}

	bi.init()

	return &bi, nil
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
				// TODO(karmi): Wrap error
				return fmt.Errorf("%s", err)
			}
		}
	}

	return nil
}

// Stats returns indexer statistics.
//
func (bi *bulkIndexer) Stats() BulkIndexerStats {
	return BulkIndexerStats{
		NumAdded:   uint(atomic.LoadUint64(&bi.stats.numAdded)),
		NumFlushed: uint(atomic.LoadUint64(&bi.stats.numFlushed)),
		NumFailed:  uint(atomic.LoadUint64(&bi.stats.numFailed)),
		NumIndexed: uint(atomic.LoadUint64(&bi.stats.numIndexed)),
		NumCreated: uint(atomic.LoadUint64(&bi.stats.numCreated)),
		NumUpdated: uint(atomic.LoadUint64(&bi.stats.numUpdated)),
		NumDeleted: uint(atomic.LoadUint64(&bi.stats.numDeleted)),
	}
}

// init initializes the bulk indexer.
//
func (bi *bulkIndexer) init() {
	bi.queue = make(chan BulkIndexerItem, bi.config.NumWorkers)

	for i := 1; i <= bi.config.NumWorkers; i++ {
		w := worker{
			id:  i,
			ch:  bi.queue,
			bi:  bi,
			buf: bytes.NewBuffer(make([]byte, 0, bi.config.FlushBytes)),
			aux: make([]byte, 0, 512)}
		w.run()
		bi.workers = append(bi.workers, &w)
	}
	bi.wg.Add(bi.config.NumWorkers)
}

// worker represents an indexer worker.
//
type worker struct {
	id    int
	ch    <-chan BulkIndexerItem
	bi    *bulkIndexer
	buf   *bytes.Buffer
	aux   []byte
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
		defer w.bi.wg.Done()

		for item := range w.ch {
			if os.Getenv("DEBUG") != "" {
				fmt.Printf(">>> [worker-%03d] Got %s:%s\n", w.id, item.Action, item.DocumentID)
			}

			if err := w.writeMeta(item); err != nil {
				if item.OnFailure != nil {
					item.OnFailure(item, BulkIndexerResponseItem{}, err)
				}
				atomic.AddUint64(&w.bi.stats.numFailed, 1)
				continue
			}

			if err := w.writeBody(item); err != nil {
				if item.OnFailure != nil {
					item.OnFailure(item, BulkIndexerResponseItem{}, err)
				}
				atomic.AddUint64(&w.bi.stats.numFailed, 1)
				continue
			}

			w.items = append(w.items, item)
			if w.buf.Len() >= w.bi.config.FlushBytes {
				w.flush()
			}
		}
	}()
}

// writeMeta formats and writes the item metadata to the buffer.
//
func (w *worker) writeMeta(item BulkIndexerItem) error {
	// TODO(karmi): Handle errors
	w.buf.WriteRune('{')
	w.aux = strconv.AppendQuote(w.aux, item.Action)
	w.buf.Write(w.aux)
	w.aux = w.aux[:0]
	w.buf.WriteRune(':')
	w.buf.WriteRune('{')
	if item.DocumentID != "" {
		w.buf.WriteString(`"_id":`)
		w.aux = strconv.AppendQuote(w.aux, item.DocumentID)
		w.buf.Write(w.aux)
		w.aux = w.aux[:0]
	}
	w.buf.WriteRune('}')
	w.buf.WriteRune('}')
	w.buf.WriteRune('\n')
	return nil
}

// writeBody writes the item body to the buffer.
//
func (w *worker) writeBody(item BulkIndexerItem) error {
	// TODO(karmi): Handle errors
	if item.Body != nil {
		if _, err := w.buf.ReadFrom(item.Body); err != nil {
			return err
		}
		w.buf.WriteRune('\n')
	}
	return nil
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

	var (
		err error
		blk BulkIndexerResponse
	)

	defer func() {
		w.items = w.items[:0]
		w.buf.Reset()
		w.isFlushing = false
	}()

	w.isFlushing = true

	if os.Getenv("DEBUG") != "" {
		fmt.Printf(">>> [worker-%03d] FLUSH BUFFER: %s\n", w.id, w.buf.String())
	}

	res, err := w.bi.config.Client.Bulk(w.buf, w.bi.config.Client.Bulk.WithIndex(w.bi.config.Index))
	if err != nil {
		atomic.AddUint64(&w.bi.stats.numFailed, uint64(len(w.items)))
		// TODO(karmi): Wrap error
		return fmt.Errorf("flush: %s", err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	if res.IsError() {
		atomic.AddUint64(&w.bi.stats.numFailed, uint64(len(w.items)))
		// TODO(karmi): Wrap error (include response struct)
		return fmt.Errorf("flush: %s", res.String())
	}

	if err := w.bi.config.Decoder.UnmarshalFromReader(res.Body, &blk); err != nil {
		// TODO(karmi): Wrap error (include response struct)
		return fmt.Errorf("flush: error parsing response body: %s", err)
	}

	for i, blkItem := range blk.Items {
		var (
			item BulkIndexerItem
			info BulkIndexerResponseItem
		)

		item = w.items[i]
		for _, v := range blkItem {
			info = v
		}
		if info.Error.Type != "" || info.Status > 201 {
			atomic.AddUint64(&w.bi.stats.numFailed, 1)
			if item.OnFailure != nil {
				item.OnFailure(item, info, nil)
			}
		} else {
			atomic.AddUint64(&w.bi.stats.numFlushed, 1)
			if item.OnSuccess != nil {
				item.OnSuccess(item, info)
			}
		}
	}

	return err
}

type defaultJSONDecoder struct{}

func (d defaultJSONDecoder) UnmarshalFromReader(r io.Reader, blk *BulkIndexerResponse) error {
	return json.NewDecoder(r).Decode(blk)
}
