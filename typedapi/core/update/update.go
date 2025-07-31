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

// Update a document.
//
// Update a document by running a script or passing a partial document.
//
// If the Elasticsearch security features are enabled, you must have the `index`
// or `write` index privilege for the target index or index alias.
//
// The script can update, delete, or skip modifying the document.
// The API also supports passing a partial document, which is merged into the
// existing document.
// To fully replace an existing document, use the index API.
// This operation:
//
// * Gets the document (collocated with the shard) from the index.
// * Runs the specified script.
// * Indexes the result.
//
// The document must still be reindexed, but using this API removes some network
// roundtrips and reduces chances of version conflicts between the GET and the
// index operation.
//
// The `_source` field must be enabled to use this API.
// In addition to `_source`, you can access the following variables through the
// `ctx` map: `_index`, `_type`, `_id`, `_version`, `_routing`, and `_now` (the
// current timestamp).
package update

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Update struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id    string
	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewUpdate type alias for index.
type NewUpdate func(index, id string) *Update

// NewUpdateFunc returns a new instance of Update with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateFunc(tp elastictransport.Interface) NewUpdate {
	return func(index, id string) *Update {
		n := New(tp)

		n._id(id)

		n._index(index)

		return n
	}
}

// Update a document.
//
// Update a document by running a script or passing a partial document.
//
// If the Elasticsearch security features are enabled, you must have the `index`
// or `write` index privilege for the target index or index alias.
//
// The script can update, delete, or skip modifying the document.
// The API also supports passing a partial document, which is merged into the
// existing document.
// To fully replace an existing document, use the index API.
// This operation:
//
// * Gets the document (collocated with the shard) from the index.
// * Runs the specified script.
// * Indexes the result.
//
// The document must still be reindexed, but using this API removes some network
// roundtrips and reduces chances of version conflicts between the GET and the
// index operation.
//
// The `_source` field must be enabled to use this API.
// In addition to `_source`, you can access the following variables through the
// `ctx` map: `_index`, `_type`, `_id`, `_version`, `_routing`, and `_now` (the
// current timestamp).
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update.html
func New(tp elastictransport.Interface) *Update {
	r := &Update{
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
func (r *Update) Raw(raw io.Reader) *Update {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Update) Request(req *Request) *Update {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Update) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Update: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_update")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

		method = http.MethodPost
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
func (r Update) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "update")
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
		instrument.BeforeRequest(req, "update")
		if reader := instrument.RecordRequestBody(ctx, "update", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "update")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Update query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a update.Response
func (r Update) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "update")
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

// Header set a key, value pair in the Update headers map.
func (r *Update) Header(key, value string) *Update {
	r.headers.Set(key, value)

	return r
}

// Id A unique identifier for the document to be updated.
// API Name: id
func (r *Update) _id(id string) *Update {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index The name of the target index.
// By default, the index is created automatically if it doesn't exist.
// API Name: index
func (r *Update) _index(index string) *Update {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// IfPrimaryTerm Only perform the operation if the document has this primary term.
// API name: if_primary_term
func (r *Update) IfPrimaryTerm(ifprimaryterm string) *Update {
	r.values.Set("if_primary_term", ifprimaryterm)

	return r
}

// IfSeqNo Only perform the operation if the document has this sequence number.
// API name: if_seq_no
func (r *Update) IfSeqNo(sequencenumber string) *Update {
	r.values.Set("if_seq_no", sequencenumber)

	return r
}

// IncludeSourceOnError True or false if to include the document source in the error message in case
// of parsing errors.
// API name: include_source_on_error
func (r *Update) IncludeSourceOnError(includesourceonerror bool) *Update {
	r.values.Set("include_source_on_error", strconv.FormatBool(includesourceonerror))

	return r
}

// Lang The script language.
// API name: lang
func (r *Update) Lang(lang string) *Update {
	r.values.Set("lang", lang)

	return r
}

// Refresh If 'true', Elasticsearch refreshes the affected shards to make this operation
// visible to search.
// If 'wait_for', it waits for a refresh to make this operation visible to
// search.
// If 'false', it does nothing with refreshes.
// API name: refresh
func (r *Update) Refresh(refresh refresh.Refresh) *Update {
	r.values.Set("refresh", refresh.String())

	return r
}

// RequireAlias If `true`, the destination must be an index alias.
// API name: require_alias
func (r *Update) RequireAlias(requirealias bool) *Update {
	r.values.Set("require_alias", strconv.FormatBool(requirealias))

	return r
}

// RetryOnConflict The number of times the operation should be retried when a conflict occurs.
// API name: retry_on_conflict
func (r *Update) RetryOnConflict(retryonconflict int) *Update {
	r.values.Set("retry_on_conflict", strconv.Itoa(retryonconflict))

	return r
}

// Routing A custom value used to route operations to a specific shard.
// API name: routing
func (r *Update) Routing(routing string) *Update {
	r.values.Set("routing", routing)

	return r
}

// Timeout The period to wait for the following operations: dynamic mapping updates and
// waiting for active shards.
// Elasticsearch waits for at least the timeout period before failing.
// The actual wait time could be longer, particularly when multiple waits occur.
// API name: timeout
func (r *Update) Timeout(duration string) *Update {
	r.values.Set("timeout", duration)

	return r
}

// WaitForActiveShards The number of copies of each shard that must be active before proceeding with
// the operation.
// Set to 'all' or any positive integer up to the total number of shards in the
// index (`number_of_replicas`+1).
// The default value of `1` means it waits for each primary shard to be active.
// API name: wait_for_active_shards
func (r *Update) WaitForActiveShards(waitforactiveshards string) *Update {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// SourceExcludes_ The source fields you want to exclude.
// API name: _source_excludes
func (r *Update) SourceExcludes_(fields ...string) *Update {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ The source fields you want to retrieve.
// API name: _source_includes
func (r *Update) SourceIncludes_(fields ...string) *Update {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Update) ErrorTrace(errortrace bool) *Update {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Update) FilterPath(filterpaths ...string) *Update {
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
func (r *Update) Human(human bool) *Update {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Update) Pretty(pretty bool) *Update {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// DetectNoop If `true`, the `result` in the response is set to `noop` (no operation) when
// there are no changes to the document.
// API name: detect_noop
func (r *Update) DetectNoop(detectnoop bool) *Update {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.DetectNoop = &detectnoop

	return r
}

// Doc A partial update to an existing document.
// If both `doc` and `script` are specified, `doc` is ignored.
// API name: doc
//
// doc should be a json.RawMessage or a structure
// if a structure is provided, the client will defer a json serialization
// prior to sending the payload to Elasticsearch.
func (r *Update) Doc(doc any) *Update {
	if r.req == nil {
		r.req = NewRequest()
	}
	switch casted := doc.(type) {
	case json.RawMessage:
		r.req.Doc = casted
	default:
		r.deferred = append(r.deferred, func(request *Request) error {
			data, err := json.Marshal(doc)
			if err != nil {
				return err
			}
			r.req.Doc = data
			return nil
		})
	}

	return r
}

// DocAsUpsert If `true`, use the contents of 'doc' as the value of 'upsert'.
// NOTE: Using ingest pipelines with `doc_as_upsert` is not supported.
// API name: doc_as_upsert
func (r *Update) DocAsUpsert(docasupsert bool) *Update {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.DocAsUpsert = &docasupsert

	return r
}

// Script The script to run to update the document.
// API name: script
func (r *Update) Script(script *types.Script) *Update {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Script = script

	return r
}

// ScriptedUpsert If `true`, run the script whether or not the document exists.
// API name: scripted_upsert
func (r *Update) ScriptedUpsert(scriptedupsert bool) *Update {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ScriptedUpsert = &scriptedupsert

	return r
}

// Source_ If `false`, turn off source retrieval.
// You can also specify a comma-separated list of the fields you want to
// retrieve.
// API name: _source
func (r *Update) Source_(sourceconfig types.SourceConfig) *Update {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Source_ = sourceconfig

	return r
}

// Upsert If the document does not already exist, the contents of 'upsert' are inserted
// as a new document.
// If the document exists, the 'script' is run.
// API name: upsert
//
// upsert should be a json.RawMessage or a structure
// if a structure is provided, the client will defer a json serialization
// prior to sending the payload to Elasticsearch.
func (r *Update) Upsert(upsert any) *Update {
	if r.req == nil {
		r.req = NewRequest()
	}
	switch casted := upsert.(type) {
	case json.RawMessage:
		r.req.Upsert = casted
	default:
		r.deferred = append(r.deferred, func(request *Request) error {
			data, err := json.Marshal(upsert)
			if err != nil {
				return err
			}
			r.req.Upsert = data
			return nil
		})
	}

	return r
}
