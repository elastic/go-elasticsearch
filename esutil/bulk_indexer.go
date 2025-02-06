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
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// BulkIndexer represents a parallel, asynchronous, efficient indexer for Elasticsearch.
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
	RequireAlias        bool
	Source              []string
	SourceExcludes      []string
	SourceIncludes      []string
	Timeout             time.Duration
	WaitForActiveShards string
}

// BulkIndexerStats represents the indexer statistics.
type BulkIndexerStats struct {
	NumAdded     uint64
	NumFlushed   uint64
	NumFailed    uint64
	NumIndexed   uint64
	NumCreated   uint64
	NumUpdated   uint64
	NumDeleted   uint64
	NumRequests  uint64
	FlushedBytes uint64
}

// BulkIndexerItem represents an indexer item.
type BulkIndexerItem struct {
	Index           string
	Action          string
	DocumentID      string
	Routing         string
	RequireAlias    bool
	Version         *int64
	VersionType     string
	Body            io.ReadSeeker
	RetryOnConflict *int
	IfSeqNo         *int64
	IfPrimaryTerm   *int64
	meta            bytes.Buffer // Item metadata header
	payloadLength   int          // Item payload total length metadata+newline+body length

	OnSuccess func(context.Context, BulkIndexerItem, BulkIndexerResponseItem)        // Per item
	OnFailure func(context.Context, BulkIndexerItem, BulkIndexerResponseItem, error) // Per item
}

// marshallMeta format as JSON the item metadata.
func (item *BulkIndexerItem) marshallMeta() {
	// Pre-allocate a buffer large enough for most use cases.
	// 'aux = aux[:0]' resets the length without changing the capacity.
	aux := make([]byte, 0, 256)

	item.meta.WriteRune('{')
	item.meta.Write(strconv.AppendQuote(aux, item.Action))
	aux = aux[:0]
	item.meta.WriteRune(':')
	item.meta.WriteRune('{')
	if item.DocumentID != "" {
		item.meta.WriteString(`"_id":`)
		item.meta.Write(strconv.AppendQuote(aux, item.DocumentID))
		aux = aux[:0]
	}

	if item.DocumentID != "" && item.Version != nil {
		item.meta.WriteRune(',')
		item.meta.WriteString(`"version":`)
		item.meta.WriteString(strconv.FormatInt(*item.Version, 10))
	}

	if item.DocumentID != "" && item.VersionType != "" {
		item.meta.WriteRune(',')
		item.meta.WriteString(`"version_type":`)
		item.meta.Write(strconv.AppendQuote(aux, item.VersionType))
		aux = aux[:0]
	}

	if item.Routing != "" {
		if item.DocumentID != "" {
			item.meta.WriteRune(',')
		}
		item.meta.WriteString(`"routing":`)
		item.meta.Write(strconv.AppendQuote(aux, item.Routing))
		aux = aux[:0]
	}
	if item.Index != "" {
		if item.DocumentID != "" || item.Routing != "" {
			item.meta.WriteRune(',')
		}
		item.meta.WriteString(`"_index":`)
		item.meta.Write(strconv.AppendQuote(aux, item.Index))
		aux = aux[:0]
	}
	if item.RetryOnConflict != nil && item.Action == "update" {
		if item.DocumentID != "" || item.Routing != "" || item.Index != "" {
			item.meta.WriteString(",")
		}
		item.meta.WriteString(`"retry_on_conflict":`)
		item.meta.Write(strconv.AppendInt(aux, int64(*item.RetryOnConflict), 10))
		aux = aux[:0]
	}
	if item.RequireAlias {
		if item.DocumentID != "" || item.Routing != "" || item.Index != "" || item.RetryOnConflict != nil {
			item.meta.WriteString(",")
		}
		item.meta.WriteString(`"require_alias":`)
		item.meta.Write(strconv.AppendBool(aux, item.RequireAlias))
		aux = aux[:0]
	}

	if item.DocumentID != "" && item.IfSeqNo != nil && item.IfPrimaryTerm != nil {
		item.meta.WriteRune(',')
		item.meta.WriteString(`"if_seq_no":`)
		item.meta.WriteString(strconv.FormatInt(*item.IfSeqNo, 10))
		item.meta.WriteRune(',')
		item.meta.WriteString(`"if_primary_term":`)
		item.meta.WriteString(strconv.FormatInt(*item.IfPrimaryTerm, 10))
	}

	item.meta.WriteRune('}')
	item.meta.WriteRune('}')
	item.meta.WriteRune('\n')
}

// computeLength calculate the size of the body and the metadata.
func (item *BulkIndexerItem) computeLength() error {
	if item.Body != nil {
		n, err := item.Body.Seek(0, io.SeekEnd)
		if err != nil {
			return err
		}
		item.payloadLength += int(n)
		_, err = item.Body.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}
	}
	item.payloadLength += len(item.meta.Bytes())
	// Add one byte to account for newline at the end of payload.
	item.payloadLength++

	return nil
}

// BulkIndexerResponse represents the Elasticsearch response.
type BulkIndexerResponse struct {
	Took      int                                  `json:"took"`
	HasErrors bool                                 `json:"errors"`
	Items     []map[string]BulkIndexerResponseItem `json:"items,omitempty"`
}

// BulkIndexerResponseItem represents the Elasticsearch response item.
type BulkIndexerResponseItem struct {
	Index        string `json:"_index"`
	DocumentID   string `json:"_id"`
	Version      int64  `json:"_version"`
	Result       string `json:"result"`
	Status       int    `json:"status"`
	SeqNo        int64  `json:"_seq_no"`
	PrimTerm     int64  `json:"_primary_term"`
	FailureStore string `json:"failure_store,omitempty"`

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
type BulkResponseJSONDecoder interface {
	UnmarshalFromReader(io.Reader, *BulkIndexerResponse) error
}

// BulkIndexerDebugLogger defines the interface for a debugging logger.
type BulkIndexerDebugLogger interface {
	Printf(string, ...interface{})
}

type bulkIndexer struct {
	wg      sync.WaitGroup
	queue   chan BulkIndexerItem
	workers []*worker
	stats   *bulkIndexerStats

	config BulkIndexerConfig
}

type bulkIndexerStats struct {
	numAdded     uint64
	numFlushed   uint64
	numFailed    uint64
	numIndexed   uint64
	numCreated   uint64
	numUpdated   uint64
	numDeleted   uint64
	numRequests  uint64
	flushedBytes uint64
}

// NewBulkIndexer creates a new bulk indexer.
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
		stats:  &bulkIndexerStats{},
	}

	bi.init()

	return &bi, nil
}

// Add adds an item to the indexer.
//
// Adding an item after a call to Close() will panic.
func (bi *bulkIndexer) Add(ctx context.Context, item BulkIndexerItem) error {
	atomic.AddUint64(&bi.stats.numAdded, 1)

	// Serialize metadata to JSON
	item.marshallMeta()
	// Compute length for body & metadata
	if err := item.computeLength(); err != nil {
		return err
	}

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
// which triggers the workers to flush and stop.
func (bi *bulkIndexer) Close(ctx context.Context) error {
	close(bi.queue)

	select {
	case <-ctx.Done():
		if bi.config.OnError != nil {
			bi.config.OnError(ctx, ctx.Err())
		}
		return ctx.Err()
	default:
		bi.wg.Wait()
	}

	return nil
}

// Stats returns indexer statistics.
func (bi *bulkIndexer) Stats() BulkIndexerStats {
	return BulkIndexerStats{
		NumAdded:     atomic.LoadUint64(&bi.stats.numAdded),
		NumFlushed:   atomic.LoadUint64(&bi.stats.numFlushed),
		NumFailed:    atomic.LoadUint64(&bi.stats.numFailed),
		NumIndexed:   atomic.LoadUint64(&bi.stats.numIndexed),
		NumCreated:   atomic.LoadUint64(&bi.stats.numCreated),
		NumUpdated:   atomic.LoadUint64(&bi.stats.numUpdated),
		NumDeleted:   atomic.LoadUint64(&bi.stats.numDeleted),
		NumRequests:  atomic.LoadUint64(&bi.stats.numRequests),
		FlushedBytes: atomic.LoadUint64(&bi.stats.flushedBytes),
	}
}

// init initializes the bulk indexer.
func (bi *bulkIndexer) init() {
	bi.queue = make(chan BulkIndexerItem, bi.config.NumWorkers)

	for i := 1; i <= bi.config.NumWorkers; i++ {
		bi.wg.Add(1)
		w := worker{
			id:     i,
			ch:     bi.queue,
			bi:     bi,
			buf:    bytes.NewBuffer(make([]byte, 0, bi.config.FlushBytes)),
			ticker: time.NewTicker(bi.config.FlushInterval),
		}
		w.run()
		bi.workers = append(bi.workers, &w)
	}
}

// worker represents an indexer worker.
type worker struct {
	id     int
	ch     <-chan BulkIndexerItem
	bi     *bulkIndexer
	buf    *bytes.Buffer
	items  []BulkIndexerItem
	ticker *time.Ticker
}

// run launches the worker in a goroutine.
func (w *worker) run() {
	go func() {
		ctx := context.Background()

		if w.bi.config.DebugLogger != nil {
			w.bi.config.DebugLogger.Printf("[worker-%03d] Started\n", w.id)
		}
		defer func() {
			w.flush(ctx)
			w.ticker.Stop()
			w.bi.wg.Done()
		}()

		for {
			select {
			case <-w.ticker.C:
				if w.bi.config.DebugLogger != nil {
					w.bi.config.DebugLogger.Printf("[worker-%03d] Auto-flushing after %s\n",
						w.id, w.bi.config.FlushInterval)
				}
				w.flush(ctx)
			case item, ok := <-w.ch:
				if !ok {
					return
				}

				if w.bi.config.DebugLogger != nil {
					w.bi.config.DebugLogger.Printf("[worker-%03d] Received item [%s:%s]\n", w.id, item.Action, item.DocumentID)
				}

				oversizePayload := w.bi.config.FlushBytes <= item.payloadLength
				if !oversizePayload && w.buf.Len() > 0 && w.buf.Len()+item.payloadLength >= w.bi.config.FlushBytes {
					if !w.flush(ctx) {
						continue
					}
				}

				if err := w.writeMeta(&item); err != nil {
					if item.OnFailure != nil {
						item.OnFailure(ctx, item, BulkIndexerResponseItem{}, err)
					}
					atomic.AddUint64(&w.bi.stats.numFailed, 1)
					continue
				}

				if err := w.writeBody(&item); err != nil {
					if item.OnFailure != nil {
						item.OnFailure(ctx, item, BulkIndexerResponseItem{}, err)
					}
					atomic.AddUint64(&w.bi.stats.numFailed, 1)
					continue
				}

				w.items = append(w.items, item)
				// Should the item payload exceed the configured FlushBytes flush happens instantly.
				if oversizePayload {
					if w.bi.config.DebugLogger != nil {
						w.bi.config.DebugLogger.Printf("[worker-%03d] Oversize Payload in item [%s:%s]\n", w.id, item.Action, item.DocumentID)
					}
					w.flush(ctx)
				}
			}
		}
	}()
}

// writeMeta writes the item metadata to the buffer.
func (w *worker) writeMeta(item *BulkIndexerItem) error {
	if _, err := w.buf.Write(item.meta.Bytes()); err != nil {
		return err
	}
	return nil
}

// writeBody writes the item body to the buffer.
func (w *worker) writeBody(item *BulkIndexerItem) error {
	if item.Body != nil {
		if _, err := w.buf.ReadFrom(item.Body); err != nil {
			if w.bi.config.OnError != nil {
				w.bi.config.OnError(context.Background(), err)
			}
			return err
		}
		item.Body.Seek(0, io.SeekStart)
		w.buf.WriteRune('\n')
	}
	return nil
}

// flush writes out the worker buffer and handles errors.
// It also restarts the ticker.
// Returns true to indicate success.
func (w *worker) flush(ctx context.Context) bool {
	ok := true
	if err := w.flushBuffer(ctx); err != nil {
		if w.bi.config.OnError != nil {
			w.bi.config.OnError(ctx, err)
		}
		ok = false
	}
	w.ticker.Reset(w.bi.config.FlushInterval)
	return ok
}

// flushBuffer writes out the worker buffer.
func (w *worker) flushBuffer(ctx context.Context) error {
	if w.bi.config.OnFlushStart != nil {
		ctx = w.bi.config.OnFlushStart(ctx)
	}

	if w.bi.config.OnFlushEnd != nil {
		defer func() { w.bi.config.OnFlushEnd(ctx) }()
	}

	bufLen := w.buf.Len()

	if bufLen < 1 {
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
		w.items = nil
		if w.buf.Cap() > w.bi.config.FlushBytes {
			w.buf = bytes.NewBuffer(make([]byte, 0, w.bi.config.FlushBytes))
		} else {
			w.buf.Reset()
		}
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
		Header:     w.bi.config.Header.Clone(),
	}
	if w.bi.config.RequireAlias {
		req.RequireAlias = &w.bi.config.RequireAlias
	}

	// Add Header and MetaHeader to config if not already set
	if req.Header == nil {
		req.Header = http.Header{}
	}
	req.Header.Set(elasticsearch.HeaderClientMeta, "h=bp")

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
			w.bi.config.OnError(ctx, fmt.Errorf("flush: %s", res.String()))
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

	atomic.AddUint64(&w.bi.stats.flushedBytes, uint64(bufLen))

	return err
}

type defaultJSONDecoder struct{}

func (d defaultJSONDecoder) UnmarshalFromReader(r io.Reader, blk *BulkIndexerResponse) error {
	return json.NewDecoder(r).Decode(blk)
}
