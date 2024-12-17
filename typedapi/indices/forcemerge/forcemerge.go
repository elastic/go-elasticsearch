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
// https://github.com/elastic/elasticsearch-specification/tree/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64

// Force a merge.
// Perform the force merge operation on the shards of one or more indices.
// For data streams, the API forces a merge on the shards of the stream's
// backing indices.
//
// Merging reduces the number of segments in each shard by merging some of them
// together and also frees up the space used by deleted documents.
// Merging normally happens automatically, but sometimes it is useful to trigger
// a merge manually.
//
// WARNING: We recommend force merging only a read-only index (meaning the index
// is no longer receiving writes).
// When documents are updated or deleted, the old version is not immediately
// removed but instead soft-deleted and marked with a "tombstone".
// These soft-deleted documents are automatically cleaned up during regular
// segment merges.
// But force merge can cause very large (greater than 5 GB) segments to be
// produced, which are not eligible for regular merges.
// So the number of soft-deleted documents can then grow rapidly, resulting in
// higher disk usage and worse search performance.
// If you regularly force merge an index receiving writes, this can also make
// snapshots more expensive, since the new documents can't be backed up
// incrementally.
package forcemerge

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

type Forcemerge struct {
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

// NewForcemerge type alias for index.
type NewForcemerge func() *Forcemerge

// NewForcemergeFunc returns a new instance of Forcemerge with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewForcemergeFunc(tp elastictransport.Interface) NewForcemerge {
	return func() *Forcemerge {
		n := New(tp)

		return n
	}
}

// Force a merge.
// Perform the force merge operation on the shards of one or more indices.
// For data streams, the API forces a merge on the shards of the stream's
// backing indices.
//
// Merging reduces the number of segments in each shard by merging some of them
// together and also frees up the space used by deleted documents.
// Merging normally happens automatically, but sometimes it is useful to trigger
// a merge manually.
//
// WARNING: We recommend force merging only a read-only index (meaning the index
// is no longer receiving writes).
// When documents are updated or deleted, the old version is not immediately
// removed but instead soft-deleted and marked with a "tombstone".
// These soft-deleted documents are automatically cleaned up during regular
// segment merges.
// But force merge can cause very large (greater than 5 GB) segments to be
// produced, which are not eligible for regular merges.
// So the number of soft-deleted documents can then grow rapidly, resulting in
// higher disk usage and worse search performance.
// If you regularly force merge an index receiving writes, this can also make
// snapshots more expensive, since the new documents can't be backed up
// incrementally.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-forcemerge.html
func New(tp elastictransport.Interface) *Forcemerge {
	r := &Forcemerge{
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
func (r *Forcemerge) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_forcemerge")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_forcemerge")

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
func (r Forcemerge) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.forcemerge")
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
		instrument.BeforeRequest(req, "indices.forcemerge")
		if reader := instrument.RecordRequestBody(ctx, "indices.forcemerge", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.forcemerge")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Forcemerge query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a forcemerge.Response
func (r Forcemerge) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.forcemerge")
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
func (r Forcemerge) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.forcemerge")
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
		err := fmt.Errorf("an error happened during the Forcemerge query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Forcemerge headers map.
func (r *Forcemerge) Header(key, value string) *Forcemerge {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names; use `_all` or empty string to perform
// the operation on all indices
// API Name: index
func (r *Forcemerge) Index(index string) *Forcemerge {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *Forcemerge) AllowNoIndices(allownoindices bool) *Forcemerge {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *Forcemerge) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Forcemerge {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// Flush Specify whether the index should be flushed after performing the operation
// (default: true)
// API name: flush
func (r *Forcemerge) Flush(flush bool) *Forcemerge {
	r.values.Set("flush", strconv.FormatBool(flush))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *Forcemerge) IgnoreUnavailable(ignoreunavailable bool) *Forcemerge {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// MaxNumSegments The number of segments the index should be merged into (default: dynamic)
// API name: max_num_segments
func (r *Forcemerge) MaxNumSegments(maxnumsegments string) *Forcemerge {
	r.values.Set("max_num_segments", maxnumsegments)

	return r
}

// OnlyExpungeDeletes Specify whether the operation should only expunge deleted documents
// API name: only_expunge_deletes
func (r *Forcemerge) OnlyExpungeDeletes(onlyexpungedeletes bool) *Forcemerge {
	r.values.Set("only_expunge_deletes", strconv.FormatBool(onlyexpungedeletes))

	return r
}

// WaitForCompletion Should the request wait until the force merge is completed.
// API name: wait_for_completion
func (r *Forcemerge) WaitForCompletion(waitforcompletion bool) *Forcemerge {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Forcemerge) ErrorTrace(errortrace bool) *Forcemerge {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Forcemerge) FilterPath(filterpaths ...string) *Forcemerge {
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
func (r *Forcemerge) Human(human bool) *Forcemerge {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Forcemerge) Pretty(pretty bool) *Forcemerge {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
