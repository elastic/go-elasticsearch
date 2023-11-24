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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Cancels a task, if it can be cancelled through an API.
package cancel

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

const (
	taskidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Cancel struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	taskid string
}

// NewCancel type alias for index.
type NewCancel func() *Cancel

// NewCancelFunc returns a new instance of Cancel with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCancelFunc(tp elastictransport.Interface) NewCancel {
	return func() *Cancel {
		n := New(tp)

		return n
	}
}

// Cancels a task, if it can be cancelled through an API.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
func New(tp elastictransport.Interface) *Cancel {
	r := &Cancel{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Cancel) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_tasks")
		path.WriteString("/")
		path.WriteString("_cancel")

		method = http.MethodPost
	case r.paramSet == taskidMask:
		path.WriteString("/")
		path.WriteString("_tasks")
		path.WriteString("/")

		path.WriteString(r.taskid)
		path.WriteString("/")
		path.WriteString("_cancel")

		method = http.MethodPost
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
func (r Cancel) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Cancel query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a cancel.Response
func (r Cancel) Do(ctx context.Context) (*Response, error) {

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
func (r Cancel) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the Cancel headers map.
func (r *Cancel) Header(key, value string) *Cancel {
	r.headers.Set(key, value)

	return r
}

// TaskId ID of the task.
// API Name: taskid
func (r *Cancel) TaskId(taskid string) *Cancel {
	r.paramSet |= taskidMask
	r.taskid = taskid

	return r
}

// Actions Comma-separated list or wildcard expression of actions used to limit the
// request.
// API name: actions
func (r *Cancel) Actions(actions ...string) *Cancel {
	tmp := []string{}
	for _, item := range actions {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("actions", strings.Join(tmp, ","))

	return r
}

// Nodes Comma-separated list of node IDs or names used to limit the request.
// API name: nodes
func (r *Cancel) Nodes(nodes ...string) *Cancel {
	tmp := []string{}
	for _, item := range nodes {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("nodes", strings.Join(tmp, ","))

	return r
}

// ParentTaskId Parent task ID used to limit the tasks.
// API name: parent_task_id
func (r *Cancel) ParentTaskId(parenttaskid string) *Cancel {
	r.values.Set("parent_task_id", parenttaskid)

	return r
}

// WaitForCompletion Should the request block until the cancellation of the task and its
// descendant tasks is completed. Defaults to false
// API name: wait_for_completion
func (r *Cancel) WaitForCompletion(waitforcompletion bool) *Cancel {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}
