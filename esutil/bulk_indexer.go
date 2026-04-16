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
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
)

// ErrIndexerClosed is returned when Flush is called on a closed indexer.
var ErrIndexerClosed = errors.New("bulk indexer is closed")

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

	// Flush drains all currently queued items, flushes all worker buffers to
	// Elasticsearch, and waits for the flushes to complete. The indexer remains
	// usable after Flush returns; new items may be added with Add.
	//
	// Per-item results are delivered through the OnSuccess and OnFailure
	// callbacks on each BulkIndexerItem, exactly as with automatic flushes.
	// Flush itself only returns an error for transport-level failures or
	// context cancellation.
	//
	// Items added concurrently with Flush may or may not be included in the
	// flush; only items already enqueued at the time each worker receives the
	// flush signal are guaranteed to be sent.
	//
	// Concurrent calls to Flush are serialised. Flush returns ErrIndexerClosed
	// if the indexer has been closed.
	//
	// It is safe for concurrent use with Add.
	Flush(context.Context) error

	// Stats returns indexer statistics.
	Stats() BulkIndexerStats
}

// BulkIndexerConfig represents configuration of the indexer.
type BulkIndexerConfig struct {
	NumWorkers          int           // The number of workers. Defaults to runtime.NumCPU().
	FlushBytes          int           // The flush threshold in bytes. Defaults to 5MB.
	FlushInterval       time.Duration // The flush threshold as duration. Defaults to 30sec.
	FlushJitter         time.Duration // Max random jitter added to FlushInterval per worker, to avoid lockstep flushes. Defaults to 0 (disabled).
	QueueSizeMultiplier int           // The capacity of each worker's item queue. Total capacity is NumWorkers * QueueSizeMultiplier. Defaults to 1.

	Client      esapi.Transport         // The Elasticsearch client (required).
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
	FlushedMs    uint64
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
	ctx             context.Context // per-item context, populated from Add(ctx, item)
	meta            bytes.Buffer    // Item metadata header
	payloadLength   int             // Item payload total length metadata+newline+body length
	flushDone       chan struct{}   // non-nil marks a flush sentinel; worker closes it when flush completes

	OnSuccess func(context.Context, BulkIndexerItem, BulkIndexerResponseItem)        // Per item
	OnFailure func(context.Context, BulkIndexerItem, BulkIndexerResponseItem, error) // Per item
}

// marshallMeta format as JSON the item metadata.
//
//nolint:gocyclo
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
	workers []*worker
	stats   *bulkIndexerStats

	config     BulkIndexerConfig
	addCounter atomic.Uint64 // round-robin dispatch counter for Add()
	flushMu    sync.Mutex    // serialises Flush() and Close()
	closed     atomic.Bool   // set by Close()
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
	flushedMs    uint64
}

// NewBulkIndexer creates a new bulk indexer.
//
// Returns an error when cfg.Client is nil.
func NewBulkIndexer(cfg BulkIndexerConfig) (BulkIndexer, error) {
	if cfg.Client == nil {
		return nil, fmt.Errorf("BulkIndexerConfig.Client is required")
	}

	if cfg.Decoder == nil {
		cfg.Decoder = defaultJSONDecoder{}
	}

	if cfg.NumWorkers <= 0 {
		cfg.NumWorkers = runtime.NumCPU()
	}

	if cfg.FlushBytes <= 0 {
		cfg.FlushBytes = 5e+6
	}

	if cfg.FlushInterval <= 0 {
		cfg.FlushInterval = 30 * time.Second
	}

	if cfg.QueueSizeMultiplier <= 0 {
		cfg.QueueSizeMultiplier = 1
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

	if item.ctx == nil {
		item.ctx = ctx
	}

	// Serialize metadata to JSON
	item.marshallMeta()
	// Compute length for body & metadata
	if err := item.computeLength(); err != nil {
		return err
	}

	n := bi.addCounter.Add(1) - 1
	numWorkers := uint64(len(bi.workers))
	primary := bi.workers[n%numWorkers]

	// Avoid the O(N) scan when the hashed worker can take the item immediately.
	select {
	case primary.ch <- item:
		return nil
	default:
	}

	// Spread load: prefer making progress on any idle worker over blocking
	// on one busy worker. Start after primary to avoid retrying it and to
	// distribute scan pressure evenly under sustained back-pressure.
	for i := uint64(1); i < numWorkers; i++ {
		w := bi.workers[(n+i)%numWorkers]
		select {
		case w.ch <- item:
			return nil
		default:
		}
	}

	// Back-pressure: all queues full; block until primary drains or the
	// context is cancelled.
	select {
	case <-ctx.Done():
		if bi.config.OnError != nil {
			bi.config.OnError(ctx, ctx.Err())
		}
		return ctx.Err()
	case primary.ch <- item:
	}

	return nil
}

// Close stops the periodic flush, closes the indexer queue channel,
// which triggers the workers to flush and stop.
// Note: it is the user's responsibility to call Close on the elasticsearch Client passed in to the BulkIndexerConfig.
func (bi *bulkIndexer) Close(ctx context.Context) error {
	bi.flushMu.Lock()
	bi.closed.Store(true)
	for _, w := range bi.workers {
		close(w.ch)
	}
	bi.flushMu.Unlock()

	if err := ctx.Err(); err != nil {
		if bi.config.OnError != nil {
			bi.config.OnError(ctx, err)
		}
		return err
	}

	done := make(chan struct{})
	go func() {
		bi.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		if bi.config.OnError != nil {
			bi.config.OnError(ctx, ctx.Err())
		}
		return ctx.Err()
	}
}

// Flush drains all currently queued items, flushes all worker buffers to
// Elasticsearch, and waits for the flushes to complete. The indexer remains
// usable after Flush returns.
func (bi *bulkIndexer) Flush(ctx context.Context) error {
	if bi.closed.Load() {
		return ErrIndexerClosed
	}

	bi.flushMu.Lock()
	defer bi.flushMu.Unlock()

	if bi.closed.Load() {
		return ErrIndexerClosed
	}

	if err := ctx.Err(); err != nil {
		return err
	}

	numWorkers := len(bi.workers)
	dones := make([]chan struct{}, numWorkers)
	for i := range dones {
		dones[i] = make(chan struct{})
	}

	for i, w := range bi.workers {
		sentinel := BulkIndexerItem{flushDone: dones[i]}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case w.ch <- sentinel:
		}
	}

	for i := 0; i < numWorkers; i++ {
		select {
		case <-dones[i]:
		case <-ctx.Done():
			if bi.config.OnError != nil {
				bi.config.OnError(ctx, ctx.Err())
			}
			return ctx.Err()
		}
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
		FlushedMs:    atomic.LoadUint64(&bi.stats.flushedMs),
	}
}

// init initializes the bulk indexer.
func (bi *bulkIndexer) init() {
	for i := 1; i <= bi.config.NumWorkers; i++ {
		bi.wg.Add(1)
		w := worker{
			id:     i,
			ch:     make(chan BulkIndexerItem, bi.config.QueueSizeMultiplier),
			bi:     bi,
			buf:    bytes.NewBuffer(make([]byte, 0, bi.config.FlushBytes)),
			ticker: time.NewTicker(bi.nextFlushInterval()),
		}
		w.run()
		bi.workers = append(bi.workers, &w)
	}
}

// nextFlushInterval returns FlushInterval plus a random jitter in
// [0, FlushJitter) when FlushJitter > 0. Callers are independent worker
// goroutines; math/rand/v2's top-level functions are goroutine-safe.
func (bi *bulkIndexer) nextFlushInterval() time.Duration {
	if bi.config.FlushJitter <= 0 {
		return bi.config.FlushInterval
	}
	return bi.config.FlushInterval + rand.N(bi.config.FlushJitter) //nolint:gosec // G404: non-crypto randomness is intentional for flush-interval jitter
}

// worker represents an indexer worker.
type worker struct {
	id     int
	ch     chan BulkIndexerItem
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

				if item.flushDone != nil {
					if w.bi.config.DebugLogger != nil {
						w.bi.config.DebugLogger.Printf("[worker-%03d] Flush sentinel received\n", w.id)
					}
					w.flush(ctx)
					close(item.flushDone)
					continue
				}

				if w.bi.config.DebugLogger != nil {
					w.bi.config.DebugLogger.Printf("[worker-%03d] Received item [%s:%s]\n", w.id, item.Action, item.DocumentID)
				}

				oversizePayload := w.bi.config.FlushBytes <= item.payloadLength
				if !oversizePayload && w.buf.Len() > 0 && w.buf.Len()+item.payloadLength >= w.bi.config.FlushBytes {
					w.flush(ctx)
				}

				if err := w.writeMeta(&item); err != nil {
					if item.OnFailure != nil {
						item.OnFailure(item.ctx, item, BulkIndexerResponseItem{}, err)
					}
					atomic.AddUint64(&w.bi.stats.numFailed, 1)
					continue
				}

				if err := w.writeBody(&item); err != nil {
					if item.OnFailure != nil {
						item.OnFailure(item.ctx, item, BulkIndexerResponseItem{}, err)
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
		if _, err := item.Body.Seek(0, io.SeekStart); err != nil {
			if w.bi.config.OnError != nil {
				w.bi.config.OnError(context.Background(), err)
			}
			return err
		}
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
		ok = false
	}
	w.ticker.Reset(w.bi.nextFlushInterval())
	return ok
}

// flushBuffer writes out the worker buffer.
//
//nolint:gocyclo
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
		Routing:             strings.Split(w.bi.config.Routing, ","),
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
	if transport, ok := w.bi.config.Client.(elastictransport.Instrumented); ok {
		req.Instrument = transport.InstrumentationEnabled()
	}

	// Add Header and MetaHeader to config if not already set
	if req.Header == nil {
		req.Header = http.Header{}
	}
	req.Header.Set(elasticsearch.HeaderClientMeta, "h=bp")

	start := time.Now()
	res, err := req.Do(ctx, w.bi.config.Client)
	if err != nil {
		atomic.AddUint64(&w.bi.stats.numFailed, uint64(len(w.items)))
		flushErr := fmt.Errorf("flush: %w", err)
		w.handleError(ctx, flushErr)
		return flushErr
	}

	if res.Body != nil {
		defer func() { _ = res.Body.Close() }()
	}
	if res.IsError() {
		atomic.AddUint64(&w.bi.stats.numFailed, uint64(len(w.items)))
		flushErr := fmt.Errorf("flush: %s", res.String())
		w.handleError(ctx, flushErr)
		return flushErr
	}

	if unmarshalErr := w.bi.config.Decoder.UnmarshalFromReader(res.Body, &blk); unmarshalErr != nil {
		// TODO(karmi): Wrap error (include response struct)
		w.handleError(ctx, fmt.Errorf("flush: %w", unmarshalErr))
		return fmt.Errorf("flush: error parsing response body: %w", unmarshalErr)
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
				item.OnFailure(item.ctx, item, info, nil)
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
				item.OnSuccess(item.ctx, item, info)
			}
		}
	}

	atomic.AddUint64(&w.bi.stats.flushedBytes, uint64(bufLen))
	if elapsed := time.Since(start).Milliseconds(); elapsed > 0 {
		atomic.AddUint64(&w.bi.stats.flushedMs, uint64(elapsed)) //nolint:gosec // elapsed is guaranteed positive
	}

	return err
}

func (w *worker) notifyItemsOnError(err error) {
	for _, item := range w.items {
		if item.OnFailure != nil {
			item.OnFailure(item.ctx, item, BulkIndexerResponseItem{}, err)
		}
	}
}

func (w *worker) handleError(ctx context.Context, err error) {
	if w.bi.config.OnError != nil {
		w.bi.config.OnError(ctx, err)
	}
	w.notifyItemsOnError(err)
}

type defaultJSONDecoder struct{}

func (d defaultJSONDecoder) UnmarshalFromReader(r io.Reader, blk *BulkIndexerResponse) error {
	return json.NewDecoder(r).Decode(blk)
}
