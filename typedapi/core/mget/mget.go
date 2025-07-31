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

// Get multiple documents.
//
// Get multiple JSON documents by ID from one or more indices.
// If you specify an index in the request URI, you only need to specify the
// document IDs in the request body.
// To ensure fast responses, this multi get (mget) API responds with partial
// results if one or more shards fail.
//
// **Filter source fields**
//
// By default, the `_source` field is returned for every document (if stored).
// Use the `_source` and `_source_include` or `source_exclude` attributes to
// filter what fields are returned for a particular document.
// You can include the `_source`, `_source_includes`, and `_source_excludes`
// query parameters in the request URI to specify the defaults to use when there
// are no per-document instructions.
//
// **Get stored fields**
//
// Use the `stored_fields` attribute to specify the set of stored fields you
// want to retrieve.
// Any requested fields that are not stored are ignored.
// You can include the `stored_fields` query parameter in the request URI to
// specify the defaults to use when there are no per-document instructions.
package mget

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
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Mget struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewMget type alias for index.
type NewMget func() *Mget

// NewMgetFunc returns a new instance of Mget with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMgetFunc(tp elastictransport.Interface) NewMget {
	return func() *Mget {
		n := New(tp)

		return n
	}
}

// Get multiple documents.
//
// Get multiple JSON documents by ID from one or more indices.
// If you specify an index in the request URI, you only need to specify the
// document IDs in the request body.
// To ensure fast responses, this multi get (mget) API responds with partial
// results if one or more shards fail.
//
// **Filter source fields**
//
// By default, the `_source` field is returned for every document (if stored).
// Use the `_source` and `_source_include` or `source_exclude` attributes to
// filter what fields are returned for a particular document.
// You can include the `_source`, `_source_includes`, and `_source_excludes`
// query parameters in the request URI to specify the defaults to use when there
// are no per-document instructions.
//
// **Get stored fields**
//
// Use the `stored_fields` attribute to specify the set of stored fields you
// want to retrieve.
// Any requested fields that are not stored are ignored.
// You can include the `stored_fields` query parameter in the request URI to
// specify the defaults to use when there are no per-document instructions.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-get.html
func New(tp elastictransport.Interface) *Mget {
	r := &Mget{
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
func (r *Mget) Raw(raw io.Reader) *Mget {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Mget) Request(req *Request) *Mget {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Mget) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Mget: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_mget")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_mget")

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
func (r Mget) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "mget")
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
		instrument.BeforeRequest(req, "mget")
		if reader := instrument.RecordRequestBody(ctx, "mget", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "mget")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Mget query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a mget.Response
func (r Mget) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "mget")
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

// Header set a key, value pair in the Mget headers map.
func (r *Mget) Header(key, value string) *Mget {
	r.headers.Set(key, value)

	return r
}

// Index Name of the index to retrieve documents from when `ids` are specified, or
// when a document in the `docs` array does not specify an index.
// API Name: index
func (r *Mget) Index(index string) *Mget {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// ForceSyntheticSource Should this request force synthetic _source?
// Use this to test if the mapping supports synthetic _source and to get a sense
// of the worst case performance.
// Fetches with this enabled will be slower the enabling synthetic source
// natively in the index.
// API name: force_synthetic_source
func (r *Mget) ForceSyntheticSource(forcesyntheticsource bool) *Mget {
	r.values.Set("force_synthetic_source", strconv.FormatBool(forcesyntheticsource))

	return r
}

// Preference Specifies the node or shard the operation should be performed on. Random by
// default.
// API name: preference
func (r *Mget) Preference(preference string) *Mget {
	r.values.Set("preference", preference)

	return r
}

// Realtime If `true`, the request is real-time as opposed to near-real-time.
// API name: realtime
func (r *Mget) Realtime(realtime bool) *Mget {
	r.values.Set("realtime", strconv.FormatBool(realtime))

	return r
}

// Refresh If `true`, the request refreshes relevant shards before retrieving documents.
// API name: refresh
func (r *Mget) Refresh(refresh bool) *Mget {
	r.values.Set("refresh", strconv.FormatBool(refresh))

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *Mget) Routing(routing string) *Mget {
	r.values.Set("routing", routing)

	return r
}

// Source_ True or false to return the `_source` field or not, or a list of fields to
// return.
// API name: _source
func (r *Mget) Source_(sourceconfigparam string) *Mget {
	r.values.Set("_source", sourceconfigparam)

	return r
}

// SourceExcludes_ A comma-separated list of source fields to exclude from the response.
// You can also use this parameter to exclude fields from the subset specified
// in `_source_includes` query parameter.
// API name: _source_excludes
func (r *Mget) SourceExcludes_(fields ...string) *Mget {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A comma-separated list of source fields to include in the response.
// If this parameter is specified, only these source fields are returned. You
// can exclude fields from this subset using the `_source_excludes` query
// parameter.
// If the `_source` parameter is `false`, this parameter is ignored.
// API name: _source_includes
func (r *Mget) SourceIncludes_(fields ...string) *Mget {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// StoredFields If `true`, retrieves the document fields stored in the index rather than the
// document `_source`.
// API name: stored_fields
func (r *Mget) StoredFields(fields ...string) *Mget {
	r.values.Set("stored_fields", strings.Join(fields, ","))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Mget) ErrorTrace(errortrace bool) *Mget {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Mget) FilterPath(filterpaths ...string) *Mget {
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
func (r *Mget) Human(human bool) *Mget {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Mget) Pretty(pretty bool) *Mget {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Docs The documents you want to retrieve. Required if no index is specified in the
// request URI.
// API name: docs
func (r *Mget) Docs(docs ...types.MgetOperation) *Mget {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Docs = docs

	return r
}

// Ids The IDs of the documents you want to retrieve. Allowed when the index is
// specified in the request URI.
// API name: ids
func (r *Mget) Ids(ids ...string) *Mget {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Ids = ids

	return r
}
