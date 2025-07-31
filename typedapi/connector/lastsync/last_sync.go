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

// Update the connector last sync stats.
//
// Update the fields related to the last sync of a connector.
// This action is used for analytics and monitoring.
package lastsync

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/syncstatus"
)

const (
	connectoridMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type LastSync struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	connectorid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewLastSync type alias for index.
type NewLastSync func(connectorid string) *LastSync

// NewLastSyncFunc returns a new instance of LastSync with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewLastSyncFunc(tp elastictransport.Interface) NewLastSync {
	return func(connectorid string) *LastSync {
		n := New(tp)

		n._connectorid(connectorid)

		return n
	}
}

// Update the connector last sync stats.
//
// Update the fields related to the last sync of a connector.
// This action is used for analytics and monitoring.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-last-sync-api.html
func New(tp elastictransport.Interface) *LastSync {
	r := &LastSync{
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
func (r *LastSync) Raw(raw io.Reader) *LastSync {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *LastSync) Request(req *Request) *LastSync {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *LastSync) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for LastSync: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == connectoridMask:
		path.WriteString("/")
		path.WriteString("_connector")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "connectorid", r.connectorid)
		}
		path.WriteString(r.connectorid)
		path.WriteString("/")
		path.WriteString("_last_sync")

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
func (r LastSync) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "connector.last_sync")
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
		instrument.BeforeRequest(req, "connector.last_sync")
		if reader := instrument.RecordRequestBody(ctx, "connector.last_sync", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "connector.last_sync")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the LastSync query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a lastsync.Response
func (r LastSync) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "connector.last_sync")
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

// Header set a key, value pair in the LastSync headers map.
func (r *LastSync) Header(key, value string) *LastSync {
	r.headers.Set(key, value)

	return r
}

// ConnectorId The unique identifier of the connector to be updated
// API Name: connectorid
func (r *LastSync) _connectorid(connectorid string) *LastSync {
	r.paramSet |= connectoridMask
	r.connectorid = connectorid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *LastSync) ErrorTrace(errortrace bool) *LastSync {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *LastSync) FilterPath(filterpaths ...string) *LastSync {
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
func (r *LastSync) Human(human bool) *LastSync {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *LastSync) Pretty(pretty bool) *LastSync {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// API name: last_access_control_sync_error
func (r *LastSync) LastAccessControlSyncError(lastaccesscontrolsyncerror string) *LastSync {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.LastAccessControlSyncError = &lastaccesscontrolsyncerror

	return r
}

// API name: last_access_control_sync_scheduled_at
func (r *LastSync) LastAccessControlSyncScheduledAt(datetime types.DateTime) *LastSync {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LastAccessControlSyncScheduledAt = datetime

	return r
}

// API name: last_access_control_sync_status
func (r *LastSync) LastAccessControlSyncStatus(lastaccesscontrolsyncstatus syncstatus.SyncStatus) *LastSync {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LastAccessControlSyncStatus = &lastaccesscontrolsyncstatus

	return r
}

// API name: last_deleted_document_count
func (r *LastSync) LastDeletedDocumentCount(lastdeleteddocumentcount int64) *LastSync {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.LastDeletedDocumentCount = &lastdeleteddocumentcount

	return r
}

// API name: last_incremental_sync_scheduled_at
func (r *LastSync) LastIncrementalSyncScheduledAt(datetime types.DateTime) *LastSync {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LastIncrementalSyncScheduledAt = datetime

	return r
}

// API name: last_indexed_document_count
func (r *LastSync) LastIndexedDocumentCount(lastindexeddocumentcount int64) *LastSync {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.LastIndexedDocumentCount = &lastindexeddocumentcount

	return r
}

// API name: last_seen
func (r *LastSync) LastSeen(datetime types.DateTime) *LastSync {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LastSeen = datetime

	return r
}

// API name: last_sync_error
func (r *LastSync) LastSyncError(lastsyncerror string) *LastSync {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.LastSyncError = &lastsyncerror

	return r
}

// API name: last_sync_scheduled_at
func (r *LastSync) LastSyncScheduledAt(datetime types.DateTime) *LastSync {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LastSyncScheduledAt = datetime

	return r
}

// API name: last_sync_status
func (r *LastSync) LastSyncStatus(lastsyncstatus syncstatus.SyncStatus) *LastSync {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LastSyncStatus = &lastsyncstatus

	return r
}

// API name: last_synced
func (r *LastSync) LastSynced(datetime types.DateTime) *LastSync {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LastSynced = datetime

	return r
}

// API name: sync_cursor
//
// synccursor should be a json.RawMessage or a structure
// if a structure is provided, the client will defer a json serialization
// prior to sending the payload to Elasticsearch.
func (r *LastSync) SyncCursor(synccursor any) *LastSync {
	if r.req == nil {
		r.req = NewRequest()
	}
	switch casted := synccursor.(type) {
	case json.RawMessage:
		r.req.SyncCursor = casted
	default:
		r.deferred = append(r.deferred, func(request *Request) error {
			data, err := json.Marshal(synccursor)
			if err != nil {
				return err
			}
			r.req.SyncCursor = data
			return nil
		})
	}

	return r
}
