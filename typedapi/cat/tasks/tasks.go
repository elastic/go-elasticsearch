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

// Returns information about the tasks currently executing on one or more nodes
// in the cluster.
package tasks

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
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Tasks struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int
}

// NewTasks type alias for index.
type NewTasks func() *Tasks

// NewTasksFunc returns a new instance of Tasks with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewTasksFunc(tp elastictransport.Interface) NewTasks {
	return func() *Tasks {
		n := New(tp)

		return n
	}
}

// Returns information about the tasks currently executing on one or more nodes
// in the cluster.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/tasks.html
func New(tp elastictransport.Interface) *Tasks {
	r := &Tasks{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Tasks) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("tasks")

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
func (r Tasks) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Tasks query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a tasks.Response
func (r Tasks) Do(ctx context.Context) (Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(&response)
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
func (r Tasks) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the Tasks headers map.
func (r *Tasks) Header(key, value string) *Tasks {
	r.headers.Set(key, value)

	return r
}

// Actions The task action names, which are used to limit the response.
// API name: actions
func (r *Tasks) Actions(actions ...string) *Tasks {
	tmp := []string{}
	for _, item := range actions {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("actions", strings.Join(tmp, ","))

	return r
}

// Detailed If `true`, the response includes detailed information about shard recoveries.
// API name: detailed
func (r *Tasks) Detailed(detailed bool) *Tasks {
	r.values.Set("detailed", strconv.FormatBool(detailed))

	return r
}

// NodeId Unique node identifiers, which are used to limit the response.
// API name: node_id
func (r *Tasks) NodeId(nodeids ...string) *Tasks {
	tmp := []string{}
	for _, item := range nodeids {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("node_id", strings.Join(tmp, ","))

	return r
}

// ParentTaskId The parent task identifier, which is used to limit the response.
// API name: parent_task_id
func (r *Tasks) ParentTaskId(parenttaskid string) *Tasks {
	r.values.Set("parent_task_id", parenttaskid)

	return r
}
