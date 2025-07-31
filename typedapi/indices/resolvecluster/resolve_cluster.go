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

// Resolve the cluster.
//
// Resolve the specified index expressions to return information about each
// cluster, including the local "querying" cluster, if included.
// If no index expression is provided, the API will return information about all
// the remote clusters that are configured on the querying cluster.
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
// remote cluster specified in the index expression. Note that this endpoint
// actively attempts to contact the remote clusters, unlike the `remote/info`
// endpoint.
// * Whether each remote cluster is configured with `skip_unavailable` as `true`
// or `false`.
// * Whether there are any indices, aliases, or data streams on that cluster
// that match the index expression.
// * Whether the search is likely to have errors returned when you do the
// cross-cluster search (including any authorization errors if you do not have
// permission to query the index).
// * Cluster version information, including the Elasticsearch server version.
//
// For example, `GET /_resolve/cluster/my-index-*,cluster*:my-index-*` returns
// information about the local cluster and all remotely configured clusters that
// start with the alias `cluster*`.
// Each cluster returns information about whether it has any indices, aliases or
// data streams that match `my-index-*`.
//
// ## Note on backwards compatibility
// The ability to query without an index expression was added in version 8.18,
// so when
// querying remote clusters older than that, the local cluster will send the
// index
// expression `dummy*` to those remote clusters. Thus, if an errors occur, you
// may see a reference
// to that index expression even though you didn't request it. If it causes a
// problem, you can
// instead include an index expression like `*:*` to bypass the issue.
//
// ## Advantages of using this endpoint before a cross-cluster search
//
// You may want to exclude a cluster or index from a search when:
//
// * A remote cluster is not currently connected and is configured with
// `skip_unavailable=false`. Running a cross-cluster search under those
// conditions will cause the entire search to fail.
// * A cluster has no matching indices, aliases or data streams for the index
// expression (or your user does not have permissions to search them). For
// example, suppose your index expression is `logs*,remote1:logs*` and the
// remote1 cluster has no indices, aliases or data streams that match `logs*`.
// In that case, that cluster will return no results from that cluster if you
// include it in a cross-cluster search.
// * The index expression (combined with any query parameters you specify) will
// likely cause an exception to be thrown when you do the search. In these
// cases, the "error" field in the `_resolve/cluster` response will be present.
// (This is also where security/permission errors will be shown.)
// * A remote cluster is an older version that does not support the feature you
// want to use in your search.
//
// ## Test availability of remote clusters
//
// The `remote/info` endpoint is commonly used to test whether the "local"
// cluster (the cluster being queried) is connected to its remote clusters, but
// it does not necessarily reflect whether the remote cluster is available or
// not.
// The remote cluster may be available, while the local cluster is not currently
// connected to it.
//
// You can use the `_resolve/cluster` API to attempt to reconnect to remote
// clusters.
// For example with `GET _resolve/cluster` or `GET _resolve/cluster/*:*`.
// The `connected` field in the response will indicate whether it was
// successful.
// If a connection was (re-)established, this will also cause the `remote/info`
// endpoint to now indicate a connected status.
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
type NewResolveCluster func() *ResolveCluster

// NewResolveClusterFunc returns a new instance of ResolveCluster with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewResolveClusterFunc(tp elastictransport.Interface) NewResolveCluster {
	return func() *ResolveCluster {
		n := New(tp)

		return n
	}
}

// Resolve the cluster.
//
// Resolve the specified index expressions to return information about each
// cluster, including the local "querying" cluster, if included.
// If no index expression is provided, the API will return information about all
// the remote clusters that are configured on the querying cluster.
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
// remote cluster specified in the index expression. Note that this endpoint
// actively attempts to contact the remote clusters, unlike the `remote/info`
// endpoint.
// * Whether each remote cluster is configured with `skip_unavailable` as `true`
// or `false`.
// * Whether there are any indices, aliases, or data streams on that cluster
// that match the index expression.
// * Whether the search is likely to have errors returned when you do the
// cross-cluster search (including any authorization errors if you do not have
// permission to query the index).
// * Cluster version information, including the Elasticsearch server version.
//
// For example, `GET /_resolve/cluster/my-index-*,cluster*:my-index-*` returns
// information about the local cluster and all remotely configured clusters that
// start with the alias `cluster*`.
// Each cluster returns information about whether it has any indices, aliases or
// data streams that match `my-index-*`.
//
// ## Note on backwards compatibility
// The ability to query without an index expression was added in version 8.18,
// so when
// querying remote clusters older than that, the local cluster will send the
// index
// expression `dummy*` to those remote clusters. Thus, if an errors occur, you
// may see a reference
// to that index expression even though you didn't request it. If it causes a
// problem, you can
// instead include an index expression like `*:*` to bypass the issue.
//
// ## Advantages of using this endpoint before a cross-cluster search
//
// You may want to exclude a cluster or index from a search when:
//
// * A remote cluster is not currently connected and is configured with
// `skip_unavailable=false`. Running a cross-cluster search under those
// conditions will cause the entire search to fail.
// * A cluster has no matching indices, aliases or data streams for the index
// expression (or your user does not have permissions to search them). For
// example, suppose your index expression is `logs*,remote1:logs*` and the
// remote1 cluster has no indices, aliases or data streams that match `logs*`.
// In that case, that cluster will return no results from that cluster if you
// include it in a cross-cluster search.
// * The index expression (combined with any query parameters you specify) will
// likely cause an exception to be thrown when you do the search. In these
// cases, the "error" field in the `_resolve/cluster` response will be present.
// (This is also where security/permission errors will be shown.)
// * A remote cluster is an older version that does not support the feature you
// want to use in your search.
//
// ## Test availability of remote clusters
//
// The `remote/info` endpoint is commonly used to test whether the "local"
// cluster (the cluster being queried) is connected to its remote clusters, but
// it does not necessarily reflect whether the remote cluster is available or
// not.
// The remote cluster may be available, while the local cluster is not currently
// connected to it.
//
// You can use the `_resolve/cluster` API to attempt to reconnect to remote
// clusters.
// For example with `GET _resolve/cluster` or `GET _resolve/cluster/*:*`.
// The `connected` field in the response will indicate whether it was
// successful.
// If a connection was (re-)established, this will also cause the `remote/info`
// endpoint to now indicate a connected status.
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
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_resolve")
		path.WriteString("/")
		path.WriteString("cluster")

		method = http.MethodGet
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

// Name A comma-separated list of names or index patterns for the indices, aliases,
// and data streams to resolve.
// Resources on remote clusters can be specified using the `<cluster>`:`<name>`
// syntax.
// Index and cluster exclusions (e.g., `-cluster1:*`) are also supported.
// If no index expression is specified, information about all remote clusters
// configured on the local cluster
// is returned without doing any index matching
// API Name: name
func (r *ResolveCluster) Name(name string) *ResolveCluster {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// AllowNoIndices If false, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing
// or closed indices. This behavior applies even if the request targets other
// open indices. For example, a request
// targeting `foo*,bar*` returns an error if an index starts with `foo` but no
// index starts with `bar`.
// NOTE: This option is only supported when specifying an index expression. You
// will get an error if you specify index
// options to the `_resolve/cluster` API endpoint that takes no index
// expression.
// API name: allow_no_indices
func (r *ResolveCluster) AllowNoIndices(allownoindices bool) *ResolveCluster {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`.
// NOTE: This option is only supported when specifying an index expression. You
// will get an error if you specify index
// options to the `_resolve/cluster` API endpoint that takes no index
// expression.
// API name: expand_wildcards
func (r *ResolveCluster) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ResolveCluster {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreThrottled If true, concrete, expanded, or aliased indices are ignored when frozen.
// NOTE: This option is only supported when specifying an index expression. You
// will get an error if you specify index
// options to the `_resolve/cluster` API endpoint that takes no index
// expression.
// API name: ignore_throttled
func (r *ResolveCluster) IgnoreThrottled(ignorethrottled bool) *ResolveCluster {
	r.values.Set("ignore_throttled", strconv.FormatBool(ignorethrottled))

	return r
}

// IgnoreUnavailable If false, the request returns an error if it targets a missing or closed
// index.
// NOTE: This option is only supported when specifying an index expression. You
// will get an error if you specify index
// options to the `_resolve/cluster` API endpoint that takes no index
// expression.
// API name: ignore_unavailable
func (r *ResolveCluster) IgnoreUnavailable(ignoreunavailable bool) *ResolveCluster {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Timeout The maximum time to wait for remote clusters to respond.
// If a remote cluster does not respond within this timeout period, the API
// response
// will show the cluster as not connected and include an error message that the
// request timed out.
//
// The default timeout is unset and the query can take
// as long as the networking layer is configured to wait for remote clusters
// that are
// not responding (typically 30 seconds).
// API name: timeout
func (r *ResolveCluster) Timeout(duration string) *ResolveCluster {
	r.values.Set("timeout", duration)

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
