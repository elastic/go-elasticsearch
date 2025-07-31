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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Set the connector sync job stats.
// Stats include: `deleted_document_count`, `indexed_document_count`,
// `indexed_document_volume`, and `total_document_count`.
// You can also update `last_seen`.
// This API is mainly used by the connector service for updating sync job
// information.
//
// To sync data using self-managed connectors, you need to deploy the Elastic
// connector service on your own infrastructure.
// This service runs automatically on Elastic Cloud for Elastic managed
// connectors.
package syncjobupdatestats

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	connectorsyncjobidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SyncJobUpdateStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	connectorsyncjobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewSyncJobUpdateStats type alias for index.
type NewSyncJobUpdateStats func(connectorsyncjobid string) *SyncJobUpdateStats

// NewSyncJobUpdateStatsFunc returns a new instance of SyncJobUpdateStats with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSyncJobUpdateStatsFunc(tp elastictransport.Interface) NewSyncJobUpdateStats {
	return func(connectorsyncjobid string) *SyncJobUpdateStats {
		n := New(tp)

		n._connectorsyncjobid(connectorsyncjobid)

		return n
	}
}

// Set the connector sync job stats.
// Stats include: `deleted_document_count`, `indexed_document_count`,
// `indexed_document_volume`, and `total_document_count`.
// You can also update `last_seen`.
// This API is mainly used by the connector service for updating sync job
// information.
//
// To sync data using self-managed connectors, you need to deploy the Elastic
// connector service on your own infrastructure.
// This service runs automatically on Elastic Cloud for Elastic managed
// connectors.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/set-connector-sync-job-stats-api.html
func New(tp elastictransport.Interface) *SyncJobUpdateStats {
	r := &SyncJobUpdateStats{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *SyncJobUpdateStats) Raw(raw io.Reader) *SyncJobUpdateStats {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *SyncJobUpdateStats) Request(req *Request) *SyncJobUpdateStats {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *SyncJobUpdateStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for SyncJobUpdateStats: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == connectorsyncjobidMask:
		path.WriteString("/")
		path.WriteString("_connector")
		path.WriteString("/")
		path.WriteString("_sync_job")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "connectorsyncjobid", r.connectorsyncjobid)
		}
		path.WriteString(r.connectorsyncjobid)
		path.WriteString("/")
		path.WriteString("_stats")

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r SyncJobUpdateStats) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "connector.sync_job_update_stats")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "connector.sync_job_update_stats")
		if reader := instrument.RecordRequestBody(ctx, "connector.sync_job_update_stats", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "connector.sync_job_update_stats")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the SyncJobUpdateStats query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a syncjobupdatestats.Response
func (r SyncJobUpdateStats) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "connector.sync_job_update_stats")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// Header set a key, value pair in the SyncJobUpdateStats headers map.
func (r *SyncJobUpdateStats) Header(key, value string) *SyncJobUpdateStats {
	r.headers.Set(key, value)

	return r
}

// ConnectorSyncJobId The unique identifier of the connector sync job.
// API Name: connectorsyncjobid
func (r *SyncJobUpdateStats) _connectorsyncjobid(connectorsyncjobid string) *SyncJobUpdateStats {
	r.paramSet |= connectorsyncjobidMask
	r.connectorsyncjobid = connectorsyncjobid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *SyncJobUpdateStats) ErrorTrace(errortrace bool) *SyncJobUpdateStats {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *SyncJobUpdateStats) FilterPath(filterpaths ...string) *SyncJobUpdateStats {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *SyncJobUpdateStats) Human(human bool) *SyncJobUpdateStats {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *SyncJobUpdateStats) Pretty(pretty bool) *SyncJobUpdateStats {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// DeletedDocumentCount The number of documents the sync job deleted.
// API name: deleted_document_count
func (r *SyncJobUpdateStats) DeletedDocumentCount(deleteddocumentcount int64) *SyncJobUpdateStats {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.DeletedDocumentCount = deleteddocumentcount

	return r
}

// IndexedDocumentCount The number of documents the sync job indexed.
// API name: indexed_document_count
func (r *SyncJobUpdateStats) IndexedDocumentCount(indexeddocumentcount int64) *SyncJobUpdateStats {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexedDocumentCount = indexeddocumentcount

	return r
}

// IndexedDocumentVolume The total size of the data (in MiB) the sync job indexed.
// API name: indexed_document_volume
func (r *SyncJobUpdateStats) IndexedDocumentVolume(indexeddocumentvolume int64) *SyncJobUpdateStats {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexedDocumentVolume = indexeddocumentvolume

	return r
}

// LastSeen The timestamp to use in the `last_seen` property for the connector sync job.
// API name: last_seen
func (r *SyncJobUpdateStats) LastSeen(duration types.Duration) *SyncJobUpdateStats {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LastSeen = duration

	return r
}

// Metadata The connector-specific metadata.
// API name: metadata
func (r *SyncJobUpdateStats) Metadata(metadata types.Metadata) *SyncJobUpdateStats {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Metadata = metadata

	return r
}

// TotalDocumentCount The total number of documents in the target index after the sync job
// finished.
// API name: total_document_count
func (r *SyncJobUpdateStats) TotalDocumentCount(totaldocumentcount int) *SyncJobUpdateStats {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.TotalDocumentCount = &totaldocumentcount

	return r
}
