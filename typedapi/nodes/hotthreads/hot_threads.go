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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Returns information about hot threads on each node in the cluster.
package hotthreads

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/threadtype"
)

const (
	nodeidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type HotThreads struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	nodeid string
}

// NewHotThreads type alias for index.
type NewHotThreads func() *HotThreads

// NewHotThreadsFunc returns a new instance of HotThreads with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewHotThreadsFunc(tp elastictransport.Interface) NewHotThreads {
	return func() *HotThreads {
		n := New(tp)

		return n
	}
}

// Returns information about hot threads on each node in the cluster.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/cluster-nodes-hot-threads.html
func New(tp elastictransport.Interface) *HotThreads {
	r := &HotThreads{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *HotThreads) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")
		path.WriteString("hot_threads")

		method = http.MethodGet
	case r.paramSet == nodeidMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")

		path.WriteString(r.nodeid)
		path.WriteString("/")
		path.WriteString("hot_threads")

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
		req.Header.Set("Accept", "text/plain")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r HotThreads) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the HotThreads query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a hotthreads.Response
func (r HotThreads) Do(ctx context.Context) (*Response, error) {

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
func (r HotThreads) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the HotThreads headers map.
func (r *HotThreads) Header(key, value string) *HotThreads {
	r.headers.Set(key, value)

	return r
}

// NodeId List of node IDs or names used to limit returned information.
// API Name: nodeid
func (r *HotThreads) NodeId(v string) *HotThreads {
	r.paramSet |= nodeidMask
	r.nodeid = v

	return r
}

// IgnoreIdleThreads If true, known idle threads (e.g. waiting in a socket select, or to get
// a task from an empty queue) are filtered out.
// API name: ignore_idle_threads
func (r *HotThreads) IgnoreIdleThreads(b bool) *HotThreads {
	r.values.Set("ignore_idle_threads", strconv.FormatBool(b))

	return r
}

// Interval The interval to do the second sampling of threads.
// API name: interval
func (r *HotThreads) Interval(v string) *HotThreads {
	r.values.Set("interval", v)

	return r
}

// Snapshots Number of samples of thread stacktrace.
// API name: snapshots
func (r *HotThreads) Snapshots(v string) *HotThreads {
	r.values.Set("snapshots", v)

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response
// is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *HotThreads) MasterTimeout(v string) *HotThreads {
	r.values.Set("master_timeout", v)

	return r
}

// Threads Specifies the number of hot threads to provide information for.
// API name: threads
func (r *HotThreads) Threads(v string) *HotThreads {
	r.values.Set("threads", v)

	return r
}

// Timeout Period to wait for a response. If no response is received
// before the timeout expires, the request fails and returns an error.
// API name: timeout
func (r *HotThreads) Timeout(v string) *HotThreads {
	r.values.Set("timeout", v)

	return r
}

// Type The type to sample.
// API name: type
func (r *HotThreads) Type(enum threadtype.ThreadType) *HotThreads {
	r.values.Set("type", enum.String())

	return r
}

// Sort The sort order for 'cpu' type (default: total)
// API name: sort
func (r *HotThreads) Sort(enum threadtype.ThreadType) *HotThreads {
	r.values.Set("sort", enum.String())

	return r
}
