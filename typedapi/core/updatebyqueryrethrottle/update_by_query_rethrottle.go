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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Changes the number of requests per second for a particular Update By Query
// operation.
package updatebyqueryrethrottle

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
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	taskidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateByQueryRethrottle struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	taskid string
}

// NewUpdateByQueryRethrottle type alias for index.
type NewUpdateByQueryRethrottle func(taskid string) *UpdateByQueryRethrottle

// NewUpdateByQueryRethrottleFunc returns a new instance of UpdateByQueryRethrottle with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateByQueryRethrottleFunc(tp elastictransport.Interface) NewUpdateByQueryRethrottle {
	return func(taskid string) *UpdateByQueryRethrottle {
		n := New(tp)

		n.TaskId(taskid)

		return n
	}
}

// Changes the number of requests per second for a particular Update By Query
// operation.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
func New(tp elastictransport.Interface) *UpdateByQueryRethrottle {
	r := &UpdateByQueryRethrottle{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *UpdateByQueryRethrottle) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == taskidMask:
		path.WriteString("/")
		path.WriteString("_update_by_query")
		path.WriteString("/")

		path.WriteString(r.taskid)
		path.WriteString("/")
		path.WriteString("_rethrottle")

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
func (r UpdateByQueryRethrottle) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the UpdateByQueryRethrottle query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a updatebyqueryrethrottle.Response
func (r UpdateByQueryRethrottle) Do(ctx context.Context) (*Response, error) {

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

	return nil, errorResponse
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r UpdateByQueryRethrottle) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the UpdateByQueryRethrottle headers map.
func (r *UpdateByQueryRethrottle) Header(key, value string) *UpdateByQueryRethrottle {
	r.headers.Set(key, value)

	return r
}

// TaskId The task id to rethrottle
// API Name: taskid
func (r *UpdateByQueryRethrottle) TaskId(v string) *UpdateByQueryRethrottle {
	r.paramSet |= taskidMask
	r.taskid = v

	return r
}

// RequestsPerSecond The throttle to set on this request in floating sub-requests per second. -1
// means set no throttle.
// API name: requests_per_second
func (r *UpdateByQueryRethrottle) RequestsPerSecond(v string) *UpdateByQueryRethrottle {
	r.values.Set("requests_per_second", v)

	return r
}
