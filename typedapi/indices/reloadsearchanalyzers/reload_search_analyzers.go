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

// Reload search analyzers.
// Reload an index's search analyzers and their resources.
// For data streams, the API reloads search analyzers and resources for the
// stream's backing indices.
//
// IMPORTANT: After reloading the search analyzers you should clear the request
// cache to make sure it doesn't contain responses derived from the previous
// versions of the analyzer.
//
// You can use the reload search analyzers API to pick up changes to synonym
// files used in the `synonym_graph` or `synonym` token filter of a search
// analyzer.
// To be eligible, the token filter must have an `updateable` flag of `true` and
// only be used in search analyzers.
//
// NOTE: This API does not perform a reload for each shard of an index.
// Instead, it performs a reload for each node containing index shards.
// As a result, the total shard count returned by the API can differ from the
// number of index shards.
// Because reloading affects every node with an index shard, it is important to
// update the synonym file on every data node in the cluster--including nodes
// that don't contain a shard replica--before using this API.
// This ensures the synonym file is updated everywhere in the cluster in case
// shards are relocated in the future.
package reloadsearchanalyzers

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ReloadSearchAnalyzers struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewReloadSearchAnalyzers type alias for index.
type NewReloadSearchAnalyzers func(index string) *ReloadSearchAnalyzers

// NewReloadSearchAnalyzersFunc returns a new instance of ReloadSearchAnalyzers with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewReloadSearchAnalyzersFunc(tp elastictransport.Interface) NewReloadSearchAnalyzers {
	return func(index string) *ReloadSearchAnalyzers {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Reload search analyzers.
// Reload an index's search analyzers and their resources.
// For data streams, the API reloads search analyzers and resources for the
// stream's backing indices.
//
// IMPORTANT: After reloading the search analyzers you should clear the request
// cache to make sure it doesn't contain responses derived from the previous
// versions of the analyzer.
//
// You can use the reload search analyzers API to pick up changes to synonym
// files used in the `synonym_graph` or `synonym` token filter of a search
// analyzer.
// To be eligible, the token filter must have an `updateable` flag of `true` and
// only be used in search analyzers.
//
// NOTE: This API does not perform a reload for each shard of an index.
// Instead, it performs a reload for each node containing index shards.
// As a result, the total shard count returned by the API can differ from the
// number of index shards.
// Because reloading affects every node with an index shard, it is important to
// update the synonym file on every data node in the cluster--including nodes
// that don't contain a shard replica--before using this API.
// This ensures the synonym file is updated everywhere in the cluster in case
// shards are relocated in the future.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-reload-analyzers.html
func New(tp elastictransport.Interface) *ReloadSearchAnalyzers {
	r := &ReloadSearchAnalyzers{
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
func (r *ReloadSearchAnalyzers) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_reload_search_analyzers")

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r ReloadSearchAnalyzers) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.reload_search_analyzers")
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
		instrument.BeforeRequest(req, "indices.reload_search_analyzers")
		if reader := instrument.RecordRequestBody(ctx, "indices.reload_search_analyzers", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.reload_search_analyzers")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ReloadSearchAnalyzers query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a reloadsearchanalyzers.Response
func (r ReloadSearchAnalyzers) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.reload_search_analyzers")
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r ReloadSearchAnalyzers) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.reload_search_analyzers")
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
		err := fmt.Errorf("an error happened during the ReloadSearchAnalyzers query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the ReloadSearchAnalyzers headers map.
func (r *ReloadSearchAnalyzers) Header(key, value string) *ReloadSearchAnalyzers {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names to reload analyzers for
// API Name: index
func (r *ReloadSearchAnalyzers) _index(index string) *ReloadSearchAnalyzers {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *ReloadSearchAnalyzers) AllowNoIndices(allownoindices bool) *ReloadSearchAnalyzers {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *ReloadSearchAnalyzers) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ReloadSearchAnalyzers {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *ReloadSearchAnalyzers) IgnoreUnavailable(ignoreunavailable bool) *ReloadSearchAnalyzers {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Resource Changed resource to reload analyzers from if applicable
// API name: resource
func (r *ReloadSearchAnalyzers) Resource(resource string) *ReloadSearchAnalyzers {
	r.values.Set("resource", resource)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ReloadSearchAnalyzers) ErrorTrace(errortrace bool) *ReloadSearchAnalyzers {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ReloadSearchAnalyzers) FilterPath(filterpaths ...string) *ReloadSearchAnalyzers {
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
func (r *ReloadSearchAnalyzers) Human(human bool) *ReloadSearchAnalyzers {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ReloadSearchAnalyzers) Pretty(pretty bool) *ReloadSearchAnalyzers {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
