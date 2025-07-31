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

// Get the cluster state.
// Get comprehensive information about the state of the cluster.
//
// The cluster state is an internal data structure which keeps track of a
// variety of information needed by every node, including the identity and
// attributes of the other nodes in the cluster; cluster-wide settings; index
// metadata, including the mapping and settings for each index; the location and
// status of every shard copy in the cluster.
//
// The elected master node ensures that every node in the cluster has a copy of
// the same cluster state.
// This API lets you retrieve a representation of this internal state for
// debugging or diagnostic purposes.
// You may need to consult the Elasticsearch source code to determine the
// precise meaning of the response.
//
// By default the API will route requests to the elected master node since this
// node is the authoritative source of cluster states.
// You can also retrieve the cluster state held on the node handling the API
// request by adding the `?local=true` query parameter.
//
// Elasticsearch may need to expend significant effort to compute a response to
// this API in larger clusters, and the response may comprise a very large
// quantity of data.
// If you use this API repeatedly, your cluster may become unstable.
//
// WARNING: The response is a representation of an internal data structure.
// Its format is not subject to the same compatibility guarantees as other more
// stable APIs and may change from version to version.
// Do not query this API using external monitoring tools.
// Instead, obtain the information you require using other more stable cluster
// APIs.
package state

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
	metricMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type State struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	metric string
	index  string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewState type alias for index.
type NewState func() *State

// NewStateFunc returns a new instance of State with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStateFunc(tp elastictransport.Interface) NewState {
	return func() *State {
		n := New(tp)

		return n
	}
}

// Get the cluster state.
// Get comprehensive information about the state of the cluster.
//
// The cluster state is an internal data structure which keeps track of a
// variety of information needed by every node, including the identity and
// attributes of the other nodes in the cluster; cluster-wide settings; index
// metadata, including the mapping and settings for each index; the location and
// status of every shard copy in the cluster.
//
// The elected master node ensures that every node in the cluster has a copy of
// the same cluster state.
// This API lets you retrieve a representation of this internal state for
// debugging or diagnostic purposes.
// You may need to consult the Elasticsearch source code to determine the
// precise meaning of the response.
//
// By default the API will route requests to the elected master node since this
// node is the authoritative source of cluster states.
// You can also retrieve the cluster state held on the node handling the API
// request by adding the `?local=true` query parameter.
//
// Elasticsearch may need to expend significant effort to compute a response to
// this API in larger clusters, and the response may comprise a very large
// quantity of data.
// If you use this API repeatedly, your cluster may become unstable.
//
// WARNING: The response is a representation of an internal data structure.
// Its format is not subject to the same compatibility guarantees as other more
// stable APIs and may change from version to version.
// Do not query this API using external monitoring tools.
// Instead, obtain the information you require using other more stable cluster
// APIs.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-state.html
func New(tp elastictransport.Interface) *State {
	r := &State{
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
func (r *State) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cluster")
		path.WriteString("/")
		path.WriteString("state")

		method = http.MethodGet
	case r.paramSet == metricMask:
		path.WriteString("/")
		path.WriteString("_cluster")
		path.WriteString("/")
		path.WriteString("state")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "metric", r.metric)
		}
		path.WriteString(r.metric)

		method = http.MethodGet
	case r.paramSet == metricMask|indexMask:
		path.WriteString("/")
		path.WriteString("_cluster")
		path.WriteString("/")
		path.WriteString("state")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "metric", r.metric)
		}
		path.WriteString(r.metric)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)

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
func (r State) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cluster.state")
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
		instrument.BeforeRequest(req, "cluster.state")
		if reader := instrument.RecordRequestBody(ctx, "cluster.state", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.state")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the State query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a state.Response
func (r State) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.state")
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
func (r State) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.state")
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
		err := fmt.Errorf("an error happened during the State query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the State headers map.
func (r *State) Header(key, value string) *State {
	r.headers.Set(key, value)

	return r
}

// Metric Limit the information returned to the specified metrics
// API Name: metric
func (r *State) Metric(metric string) *State {
	r.paramSet |= metricMask
	r.metric = metric

	return r
}

// Index A comma-separated list of index names; use `_all` or empty string to perform
// the operation on all indices
// API Name: index
func (r *State) Index(index string) *State {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *State) AllowNoIndices(allownoindices bool) *State {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *State) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *State {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// FlatSettings Return settings in flat format (default: false)
// API name: flat_settings
func (r *State) FlatSettings(flatsettings bool) *State {
	r.values.Set("flat_settings", strconv.FormatBool(flatsettings))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *State) IgnoreUnavailable(ignoreunavailable bool) *State {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Local Return local information, do not retrieve the state from master node
// (default: false)
// API name: local
func (r *State) Local(local bool) *State {
	r.values.Set("local", strconv.FormatBool(local))

	return r
}

// MasterTimeout Specify timeout for connection to master
// API name: master_timeout
func (r *State) MasterTimeout(duration string) *State {
	r.values.Set("master_timeout", duration)

	return r
}

// WaitForMetadataVersion Wait for the metadata version to be equal or greater than the specified
// metadata version
// API name: wait_for_metadata_version
func (r *State) WaitForMetadataVersion(versionnumber string) *State {
	r.values.Set("wait_for_metadata_version", versionnumber)

	return r
}

// WaitForTimeout The maximum time to wait for wait_for_metadata_version before timing out
// API name: wait_for_timeout
func (r *State) WaitForTimeout(duration string) *State {
	r.values.Set("wait_for_timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *State) ErrorTrace(errortrace bool) *State {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *State) FilterPath(filterpaths ...string) *State {
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
func (r *State) Human(human bool) *State {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *State) Pretty(pretty bool) *State {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
