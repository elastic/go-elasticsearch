// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package esutil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/estransport"
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
	NumWorkers    int           // The number of workers. Defaults to runtime.NumCPU().
	FlushBytes    int           // The flush threshold in bytes. Defaults to 5MB.
	FlushInterval time.Duration // The flush threshold as duration. Defaults to 30sec.

	Client      *elasticsearch.Client   // The Elasticsearch client.
	Decoder     BulkResponseJSONDecoder // A custom JSON decoder.
	DebugLogger BulkIndexerDebugLogger  // An optional logger for debugging.

	OnError      func(context.Context, error)          // Called for indexer errors.
	OnFlushStart func(context.Context) context.Context // Called when the flush starts.
	OnFlushEnd   func(context.Context)                 // Called when the flush ends.

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
	NumAdded    uint64
	NumFlushed  uint64
	NumFailed   uint64
	NumIndexed  uint64
	NumCreated  uint64
	NumUpdated  uint64
	NumDeleted  uint64
	NumRequests uint64
}

// BulkIndexerItem represents an indexer item.
//
type BulkIndexerItem struct {
	Index           string
	Action          string
	DocumentID      string
	Routing         string
	Version         *int64
	VersionType     string
	Body            io.Reader
	RetryOnConflict *int

	OnSuccess func(context.Context, BulkIndexerItem, BulkIndexerResponseItem)        // Per item
	OnFailure func(context.Context, BulkIndexerItem, BulkIndexerResponseItem, error) // Per item
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
	Version    int64  `json:"_version"`
	Result     string `json:"result"`
	Status     int    `json:"status"`
	SeqNo      int64  `json:"_seq_no"`
	PrimTerm   int64  `json:"_primary_term"`

	Shards struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`

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

// BulkIndexerDebugLogger defines the interface for a debugging logger.
//
type BulkIndexerDebugLogger interface {
	Printf(string, ...interface{})
}

type bulkIndexer struct {
	wg      sync.WaitGroup
	queue   chan BulkIndexerItem
	workers []*worker
	ticker  *time.Ticker
	done    chan bool
	stats   *bulkIndexerStats

	config BulkIndexerConfig
}

type bulkIndexerStats struct {
	numAdded    uint64
	numFlushed  uint64
	numFailed   uint64
	numIndexed  uint64
	numCreated  uint64
	numUpdated  uint64
	numDeleted  uint64
	numRequests uint64
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

	if cfg.FlushInterval == 0 {
		cfg.FlushInterval = 30 * time.Second
	}

	bi := bulkIndexer{
		config: cfg,
		done:   make(chan bool),
		stats:  &bulkIndexerStats{},
	}

	bi.init()

	return &bi, nil
}

// Add adds an item to the indexer.
//
// Adding an item after a call to Close() will panic.
//
func (bi *bulkIndexer) Add(ctx context.Context, item BulkIndexerItem) error {
	atomic.AddUint64(&bi.stats.numAdded, 1)

	select {
	case <-ctx.Done():
		if bi.config.OnError != nil {
			bi.config.OnError(ctx, ctx.Err())
		}
		return ctx.Err()
	case bi.queue <- item:
	}

	return nil
}

// Close stops the periodic flush, closes the indexer queue channel,
// notifies the done channel and calls flush on all writers.
//
func (bi *bulkIndexer) Close(ctx context.Context) error {
	bi.ticker.Stop()
	close(bi.queue)
	bi.done <- true

	select {
	case <-ctx.Done():
		if bi.config.OnError != nil {
			bi.config.OnError(ctx, ctx.Err())
		}
		return ctx.Err()
	default:
		bi.wg.Wait()
	}

	for _, w := range bi.workers {
		w.mu.Lock()
		if w.buf.Len() > 0 {
			if err := w.flush(ctx); err != nil {
				w.mu.Unlock()
				if bi.config.OnError != nil {
					bi.config.OnError(ctx, err)
				}
				continue
			}
		}
		w.mu.Unlock()
	}
	return nil
}

// Stats returns indexer statistics.
//
func (bi *bulkIndexer) Stats() BulkIndexerStats {
	return BulkIndexerStats{
		NumAdded:    atomic.LoadUint64(&bi.stats.numAdded),
		NumFlushed:  atomic.LoadUint64(&bi.stats.numFlushed),
		NumFailed:   atomic.LoadUint64(&bi.stats.numFailed),
		NumIndexed:  atomic.LoadUint64(&bi.stats.numIndexed),
		NumCreated:  atomic.LoadUint64(&bi.stats.numCreated),
		NumUpdated:  atomic.LoadUint64(&bi.stats.numUpdated),
		NumDeleted:  atomic.LoadUint64(&bi.stats.numDeleted),
		NumRequests: atomic.LoadUint64(&bi.stats.numRequests),
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

	bi.ticker = time.NewTicker(bi.config.FlushInterval)
	go func() {
		ctx := context.Background()
		for {
			select {
			case <-bi.done:
				return
			case <-bi.ticker.C:
				if bi.config.DebugLogger != nil {
					bi.config.DebugLogger.Printf("[indexer] Auto-flushing workers after %s\n", bi.config.FlushInterval)
				}
				for _, w := range bi.workers {
					w.mu.Lock()
					if w.buf.Len() > 0 {
						if err := w.flush(ctx); err != nil {
							w.mu.Unlock()
							if bi.config.OnError != nil {
								bi.config.OnError(ctx, err)
							}
							continue
						}
					}
					w.mu.Unlock()
				}
			}
		}
	}()
}

// worker represents an indexer worker.
//
type worker struct {
	id    int
	ch    <-chan BulkIndexerItem
	mu    sync.Mutex
	bi    *bulkIndexer
	buf   *bytes.Buffer
	aux   []byte
	items []BulkIndexerItem
}

// run launches the worker in a goroutine.
//
func (w *worker) run() {
	go func() {
		ctx := context.Background()

		if w.bi.config.DebugLogger != nil {
			w.bi.config.DebugLogger.Printf("[worker-%03d] Started\n", w.id)
		}
		defer w.bi.wg.Done()

		for item := range w.ch {
			w.mu.Lock()

			if w.bi.config.DebugLogger != nil {
				w.bi.config.DebugLogger.Printf("[worker-%03d] Received item [%s:%s]\n", w.id, item.Action, item.DocumentID)
			}

			if err := w.writeMeta(item); err != nil {
				if item.OnFailure != nil {
					item.OnFailure(ctx, item, BulkIndexerResponseItem{}, err)
				}
				atomic.AddUint64(&w.bi.stats.numFailed, 1)
				w.mu.Unlock()
				continue
			}

			if err := w.writeBody(&item); err != nil {
				if item.OnFailure != nil {
					item.OnFailure(ctx, item, BulkIndexerResponseItem{}, err)
				}
				atomic.AddUint64(&w.bi.stats.numFailed, 1)
				w.mu.Unlock()
				continue
			}

			w.items = append(w.items, item)
			if w.buf.Len() >= w.bi.config.FlushBytes {
				if err := w.flush(ctx); err != nil {
					w.mu.Unlock()
					if w.bi.config.OnError != nil {
						w.bi.config.OnError(ctx, err)
					}
					continue
				}
			}
			w.mu.Unlock()
		}
	}()
}

// writeMeta formats and writes the item metadata to the buffer; it must be called under a lock.
//
func (w *worker) writeMeta(item BulkIndexerItem) error {
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

	if item.DocumentID != "" && item.Version != nil {
		w.buf.WriteRune(',')
		w.buf.WriteString(`"version":`)
		w.buf.WriteString(strconv.FormatInt(*item.Version, 10))
	}

	if item.DocumentID != "" && item.VersionType != "" {
		w.buf.WriteRune(',')
		w.buf.WriteString(`"version_type":`)
		w.aux = strconv.AppendQuote(w.aux, item.VersionType)
		w.buf.Write(w.aux)
		w.aux = w.aux[:0]
	}

	if item.Routing != "" {
		if item.DocumentID != "" {
			w.buf.WriteRune(',')
		}
		w.buf.WriteString(`"_routing":`)
		w.aux = strconv.AppendQuote(w.aux, item.Routing)
		w.buf.Write(w.aux)
		w.aux = w.aux[:0]
	}
	if item.Index != "" {
		if item.DocumentID != "" || item.Routing != "" {
			w.buf.WriteRune(',')
		}
		w.buf.WriteString(`"_index":`)
		w.aux = strconv.AppendQuote(w.aux, item.Index)
		w.buf.Write(w.aux)
		w.aux = w.aux[:0]
	}
	w.buf.WriteRune('}')
	w.buf.WriteRune('}')
	w.buf.WriteRune('\n')
	return nil
}

// writeBody writes the item body to the buffer; it must be called under a lock.
//
func (w *worker) writeBody(item *BulkIndexerItem) error {
	if item.Body != nil {

		var getBody func() io.Reader

		if item.OnSuccess != nil || item.OnFailure != nil {
			var buf bytes.Buffer
			buf.ReadFrom(item.Body)
			getBody = func() io.Reader {
				r := buf
				return ioutil.NopCloser(&r)
			}
			item.Body = getBody()
		}

		if _, err := w.buf.ReadFrom(item.Body); err != nil {
			if w.bi.config.OnError != nil {
				w.bi.config.OnError(context.Background(), err)
			}
			return err
		}
		w.buf.WriteRune('\n')

		if getBody != nil && (item.OnSuccess != nil || item.OnFailure != nil) {
			item.Body = getBody()
		}
	}
	return nil
}

// flush writes out the worker buffer; it must be called under a lock.
//
func (w *worker) flush(ctx context.Context) error {
	if w.bi.config.OnFlushStart != nil {
		ctx = w.bi.config.OnFlushStart(ctx)
	}

	if w.bi.config.OnFlushEnd != nil {
		defer func() { w.bi.config.OnFlushEnd(ctx) }()
	}

	if w.buf.Len() < 1 {
		if w.bi.config.DebugLogger != nil {
			w.bi.config.DebugLogger.Printf("[worker-%03d] Flush: Buffer empty\n", w.id)
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
	}()

	if w.bi.config.DebugLogger != nil {
		w.bi.config.DebugLogger.Printf("[worker-%03d] Flush: %s\n", w.id, w.buf.String())
	}

	atomic.AddUint64(&w.bi.stats.numRequests, 1)
	req := esapi.BulkRequest{
		Index: w.bi.config.Index,
		Body:  w.buf,

		Pipeline:            w.bi.config.Pipeline,
		Refresh:             w.bi.config.Refresh,
		Routing:             w.bi.config.Routing,
		Source:              w.bi.config.Source,
		SourceExcludes:      w.bi.config.SourceExcludes,
		SourceIncludes:      w.bi.config.SourceIncludes,
		Timeout:             w.bi.config.Timeout,
		WaitForActiveShards: w.bi.config.WaitForActiveShards,

		Pretty:     w.bi.config.Pretty,
		Human:      w.bi.config.Human,
		ErrorTrace: w.bi.config.ErrorTrace,
		FilterPath: w.bi.config.FilterPath,
		Header:     w.bi.config.Header,
	}

	// Add Header and MetaHeader to config if not already set
	if req.Header == nil {
		req.Header = http.Header{}
	}
	req.Header.Set(estransport.HeaderClientMeta, "h=bp")

	res, err := req.Do(ctx, w.bi.config.Client)
	if err != nil {
		atomic.AddUint64(&w.bi.stats.numFailed, uint64(len(w.items)))
		if w.bi.config.OnError != nil {
			w.bi.config.OnError(ctx, fmt.Errorf("flush: %s", err))
		}
		return fmt.Errorf("flush: %s", err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	if res.IsError() {
		atomic.AddUint64(&w.bi.stats.numFailed, uint64(len(w.items)))
		// TODO(karmi): Wrap error (include response struct)
		if w.bi.config.OnError != nil {
			w.bi.config.OnError(ctx, fmt.Errorf("flush: %s", err))
		}
		return fmt.Errorf("flush: %s", res.String())
	}

	if err := w.bi.config.Decoder.UnmarshalFromReader(res.Body, &blk); err != nil {
		// TODO(karmi): Wrap error (include response struct)
		if w.bi.config.OnError != nil {
			w.bi.config.OnError(ctx, fmt.Errorf("flush: %s", err))
		}
		return fmt.Errorf("flush: error parsing response body: %s", err)
	}

	for i, blkItem := range blk.Items {
		var (
			item BulkIndexerItem
			info BulkIndexerResponseItem
			op   string
		)

		item = w.items[i]
		// The Elasticsearch bulk response contains an array of maps like this:
		//   [ { "index": { ... } }, { "create": { ... } }, ... ]
		// We range over the map, to set the first key and value as "op" and "info".
		for k, v := range blkItem {
			op = k
			info = v
		}
		if info.Error.Type != "" || info.Status > 201 {
			atomic.AddUint64(&w.bi.stats.numFailed, 1)
			if item.OnFailure != nil {
				item.OnFailure(ctx, item, info, nil)
			}
		} else {
			atomic.AddUint64(&w.bi.stats.numFlushed, 1)

			switch op {
			case "index":
				atomic.AddUint64(&w.bi.stats.numIndexed, 1)
			case "create":
				atomic.AddUint64(&w.bi.stats.numCreated, 1)
			case "delete":
				atomic.AddUint64(&w.bi.stats.numDeleted, 1)
			case "update":
				atomic.AddUint64(&w.bi.stats.numUpdated, 1)
			}

			if item.OnSuccess != nil {
				item.OnSuccess(ctx, item, info)
			}
		}
	}

	return err
}

type defaultJSONDecoder struct{}

func (d defaultJSONDecoder) UnmarshalFromReader(r io.Reader, blk *BulkIndexerResponse) error {
	return json.NewDecoder(r).Decode(blk)
}
