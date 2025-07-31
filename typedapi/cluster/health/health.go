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

// Get the cluster health status.
//
// You can also use the API to get the health status of only specified data
// streams and indices.
// For data streams, the API retrieves the health status of the stream’s backing
// indices.
//
// The cluster health status is: green, yellow or red.
// On the shard level, a red status indicates that the specific shard is not
// allocated in the cluster. Yellow means that the primary shard is allocated
// but replicas are not. Green means that all shards are allocated.
// The index level status is controlled by the worst shard status.
//
// One of the main benefits of the API is the ability to wait until the cluster
// reaches a certain high watermark health level.
// The cluster status is controlled by the worst index status.
package health

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/level"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/waitforevents"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Health struct {
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

// NewHealth type alias for index.
type NewHealth func() *Health

// NewHealthFunc returns a new instance of Health with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewHealthFunc(tp elastictransport.Interface) NewHealth {
	return func() *Health {
		n := New(tp)

		return n
	}
}

// Get the cluster health status.
//
// You can also use the API to get the health status of only specified data
// streams and indices.
// For data streams, the API retrieves the health status of the stream’s backing
// indices.
//
// The cluster health status is: green, yellow or red.
// On the shard level, a red status indicates that the specific shard is not
// allocated in the cluster. Yellow means that the primary shard is allocated
// but replicas are not. Green means that all shards are allocated.
// The index level status is controlled by the worst shard status.
//
// One of the main benefits of the API is the ability to wait until the cluster
// reaches a certain high watermark health level.
// The cluster status is controlled by the worst index status.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html
func New(tp elastictransport.Interface) *Health {
	r := &Health{
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
func (r *Health) HttpRequest(ctx context.Context) (*http.Request, error) {
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
		path.WriteString("health")

		method = http.MethodGet
	case r.paramSet == indexMask:
		path.WriteString("/")
		path.WriteString("_cluster")
		path.WriteString("/")
		path.WriteString("health")
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
func (r Health) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cluster.health")
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
		instrument.BeforeRequest(req, "cluster.health")
		if reader := instrument.RecordRequestBody(ctx, "cluster.health", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.health")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Health query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a health.Response
func (r Health) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.health")
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

	if res.StatusCode < 299 || slices.Contains([]int{408}, res.StatusCode) {

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
func (r Health) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.health")
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
		err := fmt.Errorf("an error happened during the Health query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Health headers map.
func (r *Health) Header(key, value string) *Health {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and index aliases used to
// limit the request. Wildcard expressions (`*`) are supported. To target all
// data streams and indices in a cluster, omit this parameter or use _all or
// `*`.
// API Name: index
func (r *Health) Index(index string) *Health {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *Health) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Health {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// Level Can be one of cluster, indices or shards. Controls the details level of the
// health information returned.
// API name: level
func (r *Health) Level(level level.Level) *Health {
	r.values.Set("level", level.String())

	return r
}

// Local If true, the request retrieves information from the local node only. Defaults
// to false, which means information is retrieved from the master node.
// API name: local
func (r *Health) Local(local bool) *Health {
	r.values.Set("local", strconv.FormatBool(local))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *Health) MasterTimeout(duration string) *Health {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and returns an error.
// API name: timeout
func (r *Health) Timeout(duration string) *Health {
	r.values.Set("timeout", duration)

	return r
}

// WaitForActiveShards A number controlling to how many active shards to wait for, all to wait for
// all shards in the cluster to be active, or 0 to not wait.
// API name: wait_for_active_shards
func (r *Health) WaitForActiveShards(waitforactiveshards string) *Health {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// WaitForEvents Can be one of immediate, urgent, high, normal, low, languid. Wait until all
// currently queued events with the given priority are processed.
// API name: wait_for_events
func (r *Health) WaitForEvents(waitforevents waitforevents.WaitForEvents) *Health {
	r.values.Set("wait_for_events", waitforevents.String())

	return r
}

// WaitForNodes The request waits until the specified number N of nodes is available. It also
// accepts >=N, <=N, >N and <N. Alternatively, it is possible to use ge(N),
// le(N), gt(N) and lt(N) notation.
// API name: wait_for_nodes
func (r *Health) WaitForNodes(waitfornodes string) *Health {
	r.values.Set("wait_for_nodes", waitfornodes)

	return r
}

// WaitForNoInitializingShards A boolean value which controls whether to wait (until the timeout provided)
// for the cluster to have no shard initializations. Defaults to false, which
// means it will not wait for initializing shards.
// API name: wait_for_no_initializing_shards
func (r *Health) WaitForNoInitializingShards(waitfornoinitializingshards bool) *Health {
	r.values.Set("wait_for_no_initializing_shards", strconv.FormatBool(waitfornoinitializingshards))

	return r
}

// WaitForNoRelocatingShards A boolean value which controls whether to wait (until the timeout provided)
// for the cluster to have no shard relocations. Defaults to false, which means
// it will not wait for relocating shards.
// API name: wait_for_no_relocating_shards
func (r *Health) WaitForNoRelocatingShards(waitfornorelocatingshards bool) *Health {
	r.values.Set("wait_for_no_relocating_shards", strconv.FormatBool(waitfornorelocatingshards))

	return r
}

// WaitForStatus One of green, yellow or red. Will wait (until the timeout provided) until the
// status of the cluster changes to the one provided or better, i.e. green >
// yellow > red. By default, will not wait for any status.
// API name: wait_for_status
func (r *Health) WaitForStatus(waitforstatus healthstatus.HealthStatus) *Health {
	r.values.Set("wait_for_status", waitforstatus.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Health) ErrorTrace(errortrace bool) *Health {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Health) FilterPath(filterpaths ...string) *Health {
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
func (r *Health) Human(human bool) *Health {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Health) Pretty(pretty bool) *Health {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
