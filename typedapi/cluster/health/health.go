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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Returns basic information about the health of the cluster.
package health

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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

	buf *gobytes.Buffer

	paramSet int

	index string
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

// Returns basic information about the health of the cluster.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/cluster-health.html
func New(tp elastictransport.Interface) *Health {
	r := &Health{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
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

		path.WriteString(r.index)

		method = http.MethodGet
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
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
func (r Health) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Health query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a health.Response
func (r Health) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	if res.StatusCode == 408 {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		errorResponse := types.NewElasticsearchError()
		err = json.NewDecoder(gobytes.NewReader(data)).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}

		if errorResponse.Status == 0 {
			err = json.NewDecoder(gobytes.NewReader(data)).Decode(&response)
			if err != nil {
				return nil, err
			}

			return response, nil
		}

		return nil, errorResponse
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r Health) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the Health headers map.
func (r *Health) Header(key, value string) *Health {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and index aliases used to
// limit the request. Wildcard expressions (*) are supported. To target all data
// streams and indices in a cluster, omit this parameter or use _all or *.
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
