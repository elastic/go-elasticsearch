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

// Forces the execution of a stored watch.
package executewatch

import (
	gobytes "bytes"
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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actionexecutionmode"
)

const (
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ExecuteWatch struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	id string
}

// NewExecuteWatch type alias for index.
type NewExecuteWatch func() *ExecuteWatch

// NewExecuteWatchFunc returns a new instance of ExecuteWatch with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExecuteWatchFunc(tp elastictransport.Interface) NewExecuteWatch {
	return func() *ExecuteWatch {
		n := New(tp)

		return n
	}
}

// Forces the execution of a stored watch.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-execute-watch.html
func New(tp elastictransport.Interface) *ExecuteWatch {
	r := &ExecuteWatch{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *ExecuteWatch) Raw(raw io.Reader) *ExecuteWatch {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *ExecuteWatch) Request(req *Request) *ExecuteWatch {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ExecuteWatch) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for ExecuteWatch: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_watcher")
		path.WriteString("/")
		path.WriteString("watch")
		path.WriteString("/")

		path.WriteString(r.id)
		path.WriteString("/")
		path.WriteString("_execute")

		method = http.MethodPut
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_watcher")
		path.WriteString("/")
		path.WriteString("watch")
		path.WriteString("/")
		path.WriteString("_execute")

		method = http.MethodPut
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

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r ExecuteWatch) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the ExecuteWatch query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a executewatch.Response
func (r ExecuteWatch) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the ExecuteWatch headers map.
func (r *ExecuteWatch) Header(key, value string) *ExecuteWatch {
	r.headers.Set(key, value)

	return r
}

// Id Identifier for the watch.
// API Name: id
func (r *ExecuteWatch) Id(id string) *ExecuteWatch {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Debug Defines whether the watch runs in debug mode.
// API name: debug
func (r *ExecuteWatch) Debug(debug bool) *ExecuteWatch {
	r.values.Set("debug", strconv.FormatBool(debug))

	return r
}

// ActionModes Determines how to handle the watch actions as part of the watch execution.
// API name: action_modes
func (r *ExecuteWatch) ActionModes(actionmodes map[string]actionexecutionmode.ActionExecutionMode) *ExecuteWatch {

	r.req.ActionModes = actionmodes

	return r
}

// AlternativeInput When present, the watch uses this object as a payload instead of executing
// its own input.
// API name: alternative_input
func (r *ExecuteWatch) AlternativeInput(alternativeinput map[string]json.RawMessage) *ExecuteWatch {

	r.req.AlternativeInput = alternativeinput

	return r
}

// IgnoreCondition When set to `true`, the watch execution uses the always condition. This can
// also be specified as an HTTP parameter.
// API name: ignore_condition
func (r *ExecuteWatch) IgnoreCondition(ignorecondition bool) *ExecuteWatch {
	r.req.IgnoreCondition = &ignorecondition

	return r
}

// RecordExecution When set to `true`, the watch record representing the watch execution result
// is persisted to the `.watcher-history` index for the current time. In
// addition, the status of the watch is updated, possibly throttling subsequent
// executions. This can also be specified as an HTTP parameter.
// API name: record_execution
func (r *ExecuteWatch) RecordExecution(recordexecution bool) *ExecuteWatch {
	r.req.RecordExecution = &recordexecution

	return r
}

// API name: simulated_actions
func (r *ExecuteWatch) SimulatedActions(simulatedactions *types.SimulatedActions) *ExecuteWatch {

	r.req.SimulatedActions = simulatedactions

	return r
}

// TriggerData This structure is parsed as the data of the trigger event that will be used
// during the watch execution
// API name: trigger_data
func (r *ExecuteWatch) TriggerData(triggerdata *types.ScheduleTriggerEvent) *ExecuteWatch {

	r.req.TriggerData = triggerdata

	return r
}

// Watch When present, this watch is used instead of the one specified in the request.
// This watch is not persisted to the index and record_execution cannot be set.
// API name: watch
func (r *ExecuteWatch) Watch(watch *types.Watch) *ExecuteWatch {

	r.req.Watch = watch

	return r
}
