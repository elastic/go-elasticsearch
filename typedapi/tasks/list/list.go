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

// Get all tasks.
// Get information about the tasks currently running on one or more nodes in the
// cluster.
//
// WARNING: The task management API is new and should still be considered a beta
// feature.
// The API may change in ways that are not backwards compatible.
//
// **Identifying running tasks**
//
// The `X-Opaque-Id header`, when provided on the HTTP request header, is going
// to be returned as a header in the response as well as in the headers field
// for in the task information.
// This enables you to track certain calls or associate certain tasks with the
// client that started them.
// For example:
//
// ```
// curl -i -H "X-Opaque-Id: 123456"
// "http://localhost:9200/_tasks?group_by=parents"
// ```
//
// The API returns the following result:
//
// ```
// HTTP/1.1 200 OK
// X-Opaque-Id: 123456
// content-type: application/json; charset=UTF-8
// content-length: 831
//
//	{
//	  "tasks" : {
//	    "u5lcZHqcQhu-rUoFaqDphA:45" : {
//	      "node" : "u5lcZHqcQhu-rUoFaqDphA",
//	      "id" : 45,
//	      "type" : "transport",
//	      "action" : "cluster:monitor/tasks/lists",
//	      "start_time_in_millis" : 1513823752749,
//	      "running_time_in_nanos" : 293139,
//	      "cancellable" : false,
//	      "headers" : {
//	        "X-Opaque-Id" : "123456"
//	      },
//	      "children" : [
//	        {
//	          "node" : "u5lcZHqcQhu-rUoFaqDphA",
//	          "id" : 46,
//	          "type" : "direct",
//	          "action" : "cluster:monitor/tasks/lists[n]",
//	          "start_time_in_millis" : 1513823752750,
//	          "running_time_in_nanos" : 92133,
//	          "cancellable" : false,
//	          "parent_task_id" : "u5lcZHqcQhu-rUoFaqDphA:45",
//	          "headers" : {
//	            "X-Opaque-Id" : "123456"
//	          }
//	        }
//	      ]
//	    }
//	  }
//	 }
//
// ```
// In this example, `X-Opaque-Id: 123456` is the ID as a part of the response
// header.
// The `X-Opaque-Id` in the task `headers` is the ID for the task that was
// initiated by the REST request.
// The `X-Opaque-Id` in the children `headers` is the child task of the task
// that was initiated by the REST request.
package list

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/groupby"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type List struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewList type alias for index.
type NewList func() *List

// NewListFunc returns a new instance of List with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewListFunc(tp elastictransport.Interface) NewList {
	return func() *List {
		n := New(tp)

		return n
	}
}

// Get all tasks.
// Get information about the tasks currently running on one or more nodes in the
// cluster.
//
// WARNING: The task management API is new and should still be considered a beta
// feature.
// The API may change in ways that are not backwards compatible.
//
// **Identifying running tasks**
//
// The `X-Opaque-Id header`, when provided on the HTTP request header, is going
// to be returned as a header in the response as well as in the headers field
// for in the task information.
// This enables you to track certain calls or associate certain tasks with the
// client that started them.
// For example:
//
// ```
// curl -i -H "X-Opaque-Id: 123456"
// "http://localhost:9200/_tasks?group_by=parents"
// ```
//
// The API returns the following result:
//
// ```
// HTTP/1.1 200 OK
// X-Opaque-Id: 123456
// content-type: application/json; charset=UTF-8
// content-length: 831
//
//	{
//	  "tasks" : {
//	    "u5lcZHqcQhu-rUoFaqDphA:45" : {
//	      "node" : "u5lcZHqcQhu-rUoFaqDphA",
//	      "id" : 45,
//	      "type" : "transport",
//	      "action" : "cluster:monitor/tasks/lists",
//	      "start_time_in_millis" : 1513823752749,
//	      "running_time_in_nanos" : 293139,
//	      "cancellable" : false,
//	      "headers" : {
//	        "X-Opaque-Id" : "123456"
//	      },
//	      "children" : [
//	        {
//	          "node" : "u5lcZHqcQhu-rUoFaqDphA",
//	          "id" : 46,
//	          "type" : "direct",
//	          "action" : "cluster:monitor/tasks/lists[n]",
//	          "start_time_in_millis" : 1513823752750,
//	          "running_time_in_nanos" : 92133,
//	          "cancellable" : false,
//	          "parent_task_id" : "u5lcZHqcQhu-rUoFaqDphA:45",
//	          "headers" : {
//	            "X-Opaque-Id" : "123456"
//	          }
//	        }
//	      ]
//	    }
//	  }
//	 }
//
// ```
// In this example, `X-Opaque-Id: 123456` is the ID as a part of the response
// header.
// The `X-Opaque-Id` in the task `headers` is the ID for the task that was
// initiated by the REST request.
// The `X-Opaque-Id` in the children `headers` is the child task of the task
// that was initiated by the REST request.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
func New(tp elastictransport.Interface) *List {
	r := &List{
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
func (r *List) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_tasks")

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
func (r List) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "tasks.list")
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
		instrument.BeforeRequest(req, "tasks.list")
		if reader := instrument.RecordRequestBody(ctx, "tasks.list", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "tasks.list")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the List query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a list.Response
func (r List) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "tasks.list")
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
func (r List) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "tasks.list")
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
		err := fmt.Errorf("an error happened during the List query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the List headers map.
func (r *List) Header(key, value string) *List {
	r.headers.Set(key, value)

	return r
}

// Actions A comma-separated list or wildcard expression of actions used to limit the
// request.
// For example, you can use `cluser:*` to retrieve all cluster-related tasks.
// API name: actions
func (r *List) Actions(actions ...string) *List {
	tmp := []string{}
	for _, item := range actions {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("actions", strings.Join(tmp, ","))

	return r
}

// Detailed If `true`, the response includes detailed information about the running
// tasks.
// This information is useful to distinguish tasks from each other but is more
// costly to run.
// API name: detailed
func (r *List) Detailed(detailed bool) *List {
	r.values.Set("detailed", strconv.FormatBool(detailed))

	return r
}

// GroupBy A key that is used to group tasks in the response.
// The task lists can be grouped either by nodes or by parent tasks.
// API name: group_by
func (r *List) GroupBy(groupby groupby.GroupBy) *List {
	r.values.Set("group_by", groupby.String())

	return r
}

// Nodes A comma-separated list of node IDs or names that is used to limit the
// returned information.
// API name: nodes
func (r *List) Nodes(nodeids ...string) *List {
	r.values.Set("nodes", strings.Join(nodeids, ","))

	return r
}

// ParentTaskId A parent task identifier that is used to limit returned information.
// To return all tasks, omit this parameter or use a value of `-1`.
// If the parent task is not found, the API does not return a 404 response code.
// API name: parent_task_id
func (r *List) ParentTaskId(id string) *List {
	r.values.Set("parent_task_id", id)

	return r
}

// Timeout The period to wait for each node to respond.
// If a node does not respond before its timeout expires, the response does not
// include its information.
// However, timed out nodes are included in the `node_failures` property.
// API name: timeout
func (r *List) Timeout(duration string) *List {
	r.values.Set("timeout", duration)

	return r
}

// WaitForCompletion If `true`, the request blocks until the operation is complete.
// API name: wait_for_completion
func (r *List) WaitForCompletion(waitforcompletion bool) *List {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *List) ErrorTrace(errortrace bool) *List {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *List) FilterPath(filterpaths ...string) *List {
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
func (r *List) Human(human bool) *List {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *List) Pretty(pretty bool) *List {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
