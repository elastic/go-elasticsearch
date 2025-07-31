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

// Cancel a task.
//
// WARNING: The task management API is new and should still be considered a beta
// feature.
// The API may change in ways that are not backwards compatible.
//
// A task may continue to run for some time after it has been cancelled because
// it may not be able to safely stop its current activity straight away.
// It is also possible that Elasticsearch must complete its work on other tasks
// before it can process the cancellation.
// The get task information API will continue to list these cancelled tasks
// until they complete.
// The cancelled flag in the response indicates that the cancellation command
// has been processed and the task will stop as soon as possible.
//
// To troubleshoot why a cancelled task does not complete promptly, use the get
// task information API with the `?detailed` parameter to identify the other
// tasks the system is running.
// You can also use the node hot threads API to obtain detailed information
// about the work the system is doing instead of completing the cancelled task.
package cancel

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

	raw io.Reader

	paramSet int

	taskid string

	spanStarted bool

	instrument elastictransport.Instrumentation
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

// Cancel a task.
//
// WARNING: The task management API is new and should still be considered a beta
// feature.
// The API may change in ways that are not backwards compatible.
//
// A task may continue to run for some time after it has been cancelled because
// it may not be able to safely stop its current activity straight away.
// It is also possible that Elasticsearch must complete its work on other tasks
// before it can process the cancellation.
// The get task information API will continue to list these cancelled tasks
// until they complete.
// The cancelled flag in the response indicates that the cancellation command
// has been processed and the task will stop as soon as possible.
//
// To troubleshoot why a cancelled task does not complete promptly, use the get
// task information API with the `?detailed` parameter to identify the other
// tasks the system is running.
// You can also use the node hot threads API to obtain detailed information
// about the work the system is doing instead of completing the cancelled task.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
func New(tp elastictransport.Interface) *Cancel {
	r := &Cancel{
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

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "taskid", r.taskid)
		}
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
func (r Cancel) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "tasks.cancel")
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
		instrument.BeforeRequest(req, "tasks.cancel")
		if reader := instrument.RecordRequestBody(ctx, "tasks.cancel", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "tasks.cancel")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Cancel query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a cancel.Response
func (r Cancel) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "tasks.cancel")
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
func (r Cancel) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "tasks.cancel")
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
		err := fmt.Errorf("an error happened during the Cancel query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Cancel headers map.
func (r *Cancel) Header(key, value string) *Cancel {
	r.headers.Set(key, value)

	return r
}

// TaskId The task identifier.
// API Name: taskid
func (r *Cancel) TaskId(taskid string) *Cancel {
	r.paramSet |= taskidMask
	r.taskid = taskid

	return r
}

// Actions A comma-separated list or wildcard expression of actions that is used to
// limit the request.
// API name: actions
func (r *Cancel) Actions(actions ...string) *Cancel {
	tmp := []string{}
	for _, item := range actions {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("actions", strings.Join(tmp, ","))

	return r
}

// Nodes A comma-separated list of node IDs or names that is used to limit the
// request.
// API name: nodes
func (r *Cancel) Nodes(nodes ...string) *Cancel {
	tmp := []string{}
	for _, item := range nodes {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("nodes", strings.Join(tmp, ","))

	return r
}

// ParentTaskId A parent task ID that is used to limit the tasks.
// API name: parent_task_id
func (r *Cancel) ParentTaskId(parenttaskid string) *Cancel {
	r.values.Set("parent_task_id", parenttaskid)

	return r
}

// WaitForCompletion If true, the request blocks until all found tasks are complete.
// API name: wait_for_completion
func (r *Cancel) WaitForCompletion(waitforcompletion bool) *Cancel {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Cancel) ErrorTrace(errortrace bool) *Cancel {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Cancel) FilterPath(filterpaths ...string) *Cancel {
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
func (r *Cancel) Human(human bool) *Cancel {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Cancel) Pretty(pretty bool) *Cancel {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
