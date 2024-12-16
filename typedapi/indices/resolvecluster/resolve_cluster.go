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

// Resolve the cluster.
// Resolve the specified index expressions to return information about each
// cluster, including the local cluster, if included.
// Multiple patterns and remote clusters are supported.
//
// This endpoint is useful before doing a cross-cluster search in order to
// determine which remote clusters should be included in a search.
//
// You use the same index expression with this endpoint as you would for
// cross-cluster search.
// Index and cluster exclusions are also supported with this endpoint.
//
// For each cluster in the index expression, information is returned about:
//
// * Whether the querying ("local") cluster is currently connected to each
// remote cluster in the index expression scope.
// * Whether each remote cluster is configured with `skip_unavailable` as `true`
// or `false`.
// * Whether there are any indices, aliases, or data streams on that cluster
// that match the index expression.
// * Whether the search is likely to have errors returned when you do the
// cross-cluster search (including any authorization errors if you do not have
// permission to query the index).
// * Cluster version information, including the Elasticsearch server version.
package resolvecluster

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
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ResolveCluster struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewResolveCluster type alias for index.
type NewResolveCluster func(name string) *ResolveCluster

// NewResolveClusterFunc returns a new instance of ResolveCluster with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewResolveClusterFunc(tp elastictransport.Interface) NewResolveCluster {
	return func(name string) *ResolveCluster {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Resolve the cluster.
// Resolve the specified index expressions to return information about each
// cluster, including the local cluster, if included.
// Multiple patterns and remote clusters are supported.
//
// This endpoint is useful before doing a cross-cluster search in order to
// determine which remote clusters should be included in a search.
//
// You use the same index expression with this endpoint as you would for
// cross-cluster search.
// Index and cluster exclusions are also supported with this endpoint.
//
// For each cluster in the index expression, information is returned about:
//
// * Whether the querying ("local") cluster is currently connected to each
// remote cluster in the index expression scope.
// * Whether each remote cluster is configured with `skip_unavailable` as `true`
// or `false`.
// * Whether there are any indices, aliases, or data streams on that cluster
// that match the index expression.
// * Whether the search is likely to have errors returned when you do the
// cross-cluster search (including any authorization errors if you do not have
// permission to query the index).
// * Cluster version information, including the Elasticsearch server version.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-resolve-cluster-api.html
func New(tp elastictransport.Interface) *ResolveCluster {
	r := &ResolveCluster{
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
func (r *ResolveCluster) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_resolve")
		path.WriteString("/")
		path.WriteString("cluster")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)

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
func (r ResolveCluster) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.resolve_cluster")
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
		instrument.BeforeRequest(req, "indices.resolve_cluster")
		if reader := instrument.RecordRequestBody(ctx, "indices.resolve_cluster", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.resolve_cluster")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ResolveCluster query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a resolvecluster.Response
func (r ResolveCluster) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.resolve_cluster")
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
		err = json.NewDecoder(res.Body).Decode(&response)
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
func (r ResolveCluster) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.resolve_cluster")
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
		err := fmt.Errorf("an error happened during the ResolveCluster query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the ResolveCluster headers map.
func (r *ResolveCluster) Header(key, value string) *ResolveCluster {
	r.headers.Set(key, value)

	return r
}

// Name Comma-separated name(s) or index pattern(s) of the indices, aliases, and data
// streams to resolve.
// Resources on remote clusters can be specified using the `<cluster>`:`<name>`
// syntax.
// API Name: name
func (r *ResolveCluster) _name(name string) *ResolveCluster {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// AllowNoIndices If false, the request returns an error if any wildcard expression, index
// alias, or _all value targets only missing
// or closed indices. This behavior applies even if the request targets other
// open indices. For example, a request
// targeting foo*,bar* returns an error if an index starts with foo but no index
// starts with bar.
// API name: allow_no_indices
func (r *ResolveCluster) AllowNoIndices(allownoindices bool) *ResolveCluster {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`.
// Valid values are: `all`, `open`, `closed`, `hidden`, `none`.
// API name: expand_wildcards
func (r *ResolveCluster) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ResolveCluster {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreThrottled If true, concrete, expanded or aliased indices are ignored when frozen.
// Defaults to false.
// API name: ignore_throttled
func (r *ResolveCluster) IgnoreThrottled(ignorethrottled bool) *ResolveCluster {
	r.values.Set("ignore_throttled", strconv.FormatBool(ignorethrottled))

	return r
}

// IgnoreUnavailable If false, the request returns an error if it targets a missing or closed
// index. Defaults to false.
// API name: ignore_unavailable
func (r *ResolveCluster) IgnoreUnavailable(ignoreunavailable bool) *ResolveCluster {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ResolveCluster) ErrorTrace(errortrace bool) *ResolveCluster {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ResolveCluster) FilterPath(filterpaths ...string) *ResolveCluster {
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
func (r *ResolveCluster) Human(human bool) *ResolveCluster {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ResolveCluster) Pretty(pretty bool) *ResolveCluster {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
