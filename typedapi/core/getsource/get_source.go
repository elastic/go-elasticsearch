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

// Get a document's source.
//
// Get the source of a document.
// For example:
//
// ```
// GET my-index-000001/_source/1
// ```
//
// You can use the source filtering parameters to control which parts of the
// `_source` are returned:
//
// ```
// GET
// my-index-000001/_source/1/?_source_includes=*.id&_source_excludes=entities
// ```
package getsource

import (
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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetSource struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id    string
	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewGetSource type alias for index.
type NewGetSource func(index, id string) *GetSource

// NewGetSourceFunc returns a new instance of GetSource with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetSourceFunc(tp elastictransport.Interface) NewGetSource {
	return func(index, id string) *GetSource {
		n := New(tp)

		n._id(id)

		n._index(index)

		return n
	}
}

// Get a document's source.
//
// Get the source of a document.
// For example:
//
// ```
// GET my-index-000001/_source/1
// ```
//
// You can use the source filtering parameters to control which parts of the
// `_source` are returned:
//
// ```
// GET
// my-index-000001/_source/1/?_source_includes=*.id&_source_excludes=entities
// ```
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
func New(tp elastictransport.Interface) *GetSource {
	r := &GetSource{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetSource) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_source")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

		method = http.MethodGet
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r GetSource) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "get_source")
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
		instrument.BeforeRequest(req, "get_source")
		if reader := instrument.RecordRequestBody(ctx, "get_source", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "get_source")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetSource query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getsource.Response
func (r GetSource) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "get_source")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := new(Response)

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return *response, nil
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetSource) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "get_source")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the GetSource query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the GetSource headers map.
func (r *GetSource) Header(key, value string) *GetSource {
	r.headers.Set(key, value)

	return r
}

// Id A unique document identifier.
// API Name: id
func (r *GetSource) _id(id string) *GetSource {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Index The name of the index that contains the document.
// API Name: index
func (r *GetSource) _index(index string) *GetSource {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Preference The node or shard the operation should be performed on.
// By default, the operation is randomized between the shard replicas.
// API name: preference
func (r *GetSource) Preference(preference string) *GetSource {
	r.values.Set("preference", preference)

	return r
}

// Realtime If `true`, the request is real-time as opposed to near-real-time.
// API name: realtime
func (r *GetSource) Realtime(realtime bool) *GetSource {
	r.values.Set("realtime", strconv.FormatBool(realtime))

	return r
}

// Refresh If `true`, the request refreshes the relevant shards before retrieving the
// document.
// Setting it to `true` should be done after careful thought and verification
// that this does not cause a heavy load on the system (and slow down indexing).
// API name: refresh
func (r *GetSource) Refresh(refresh bool) *GetSource {
	r.values.Set("refresh", strconv.FormatBool(refresh))

	return r
}

// Routing A custom value used to route operations to a specific shard.
// API name: routing
func (r *GetSource) Routing(routing string) *GetSource {
	r.values.Set("routing", routing)

	return r
}

// Source_ Indicates whether to return the `_source` field (`true` or `false`) or lists
// the fields to return.
// API name: _source
func (r *GetSource) Source_(sourceconfigparam string) *GetSource {
	r.values.Set("_source", sourceconfigparam)

	return r
}

// SourceExcludes_ A comma-separated list of source fields to exclude in the response.
// API name: _source_excludes
func (r *GetSource) SourceExcludes_(fields ...string) *GetSource {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A comma-separated list of source fields to include in the response.
// API name: _source_includes
func (r *GetSource) SourceIncludes_(fields ...string) *GetSource {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// Version The version number for concurrency control.
// It must match the current version of the document for the request to succeed.
// API name: version
func (r *GetSource) Version(versionnumber string) *GetSource {
	r.values.Set("version", versionnumber)

	return r
}

// VersionType The version type.
// API name: version_type
func (r *GetSource) VersionType(versiontype versiontype.VersionType) *GetSource {
	r.values.Set("version_type", versiontype.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetSource) ErrorTrace(errortrace bool) *GetSource {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetSource) FilterPath(filterpaths ...string) *GetSource {
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
func (r *GetSource) Human(human bool) *GetSource {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GetSource) Pretty(pretty bool) *GetSource {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
