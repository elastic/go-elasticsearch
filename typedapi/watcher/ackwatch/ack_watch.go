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

// Acknowledge a watch.
// Acknowledging a watch enables you to manually throttle the execution of the
// watch's actions.
//
// The acknowledgement state of an action is stored in the
// `status.actions.<id>.ack.state` structure.
//
// IMPORTANT: If the specified watch is currently being executed, this API will
// return an error
// The reason for this behavior is to prevent overwriting the watch status from
// a watch execution.
//
// Acknowledging an action throttles further executions of that action until its
// `ack.state` is reset to `awaits_successful_execution`.
// This happens when the condition of the watch is not met (the condition
// evaluates to false).
package ackwatch

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
	watchidMask = iota + 1

	actionidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type AckWatch struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	watchid  string
	actionid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewAckWatch type alias for index.
type NewAckWatch func(watchid string) *AckWatch

// NewAckWatchFunc returns a new instance of AckWatch with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewAckWatchFunc(tp elastictransport.Interface) NewAckWatch {
	return func(watchid string) *AckWatch {
		n := New(tp)

		n._watchid(watchid)

		return n
	}
}

// Acknowledge a watch.
// Acknowledging a watch enables you to manually throttle the execution of the
// watch's actions.
//
// The acknowledgement state of an action is stored in the
// `status.actions.<id>.ack.state` structure.
//
// IMPORTANT: If the specified watch is currently being executed, this API will
// return an error
// The reason for this behavior is to prevent overwriting the watch status from
// a watch execution.
//
// Acknowledging an action throttles further executions of that action until its
// `ack.state` is reset to `awaits_successful_execution`.
// This happens when the condition of the watch is not met (the condition
// evaluates to false).
//
// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-ack-watch
func New(tp elastictransport.Interface) *AckWatch {
	r := &AckWatch{
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
func (r *AckWatch) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == watchidMask:
		path.WriteString("/")
		path.WriteString("_watcher")
		path.WriteString("/")
		path.WriteString("watch")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "watchid", r.watchid)
		}
		path.WriteString(r.watchid)
		path.WriteString("/")
		path.WriteString("_ack")

		method = http.MethodPut
	case r.paramSet == watchidMask|actionidMask:
		path.WriteString("/")
		path.WriteString("_watcher")
		path.WriteString("/")
		path.WriteString("watch")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "watchid", r.watchid)
		}
		path.WriteString(r.watchid)
		path.WriteString("/")
		path.WriteString("_ack")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "actionid", r.actionid)
		}
		path.WriteString(r.actionid)

		method = http.MethodPut
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
func (r AckWatch) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "watcher.ack_watch")
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
		instrument.BeforeRequest(req, "watcher.ack_watch")
		if reader := instrument.RecordRequestBody(ctx, "watcher.ack_watch", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "watcher.ack_watch")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the AckWatch query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a ackwatch.Response
func (r AckWatch) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "watcher.ack_watch")
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
func (r AckWatch) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "watcher.ack_watch")
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
		err := fmt.Errorf("an error happened during the AckWatch query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the AckWatch headers map.
func (r *AckWatch) Header(key, value string) *AckWatch {
	r.headers.Set(key, value)

	return r
}

// WatchId The watch identifier.
// API Name: watchid
func (r *AckWatch) _watchid(watchid string) *AckWatch {
	r.paramSet |= watchidMask
	r.watchid = watchid

	return r
}

// ActionId A comma-separated list of the action identifiers to acknowledge.
// If you omit this parameter, all of the actions of the watch are acknowledged.
// API Name: actionid
func (r *AckWatch) ActionId(actionid string) *AckWatch {
	r.paramSet |= actionidMask
	r.actionid = actionid

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *AckWatch) ErrorTrace(errortrace bool) *AckWatch {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *AckWatch) FilterPath(filterpaths ...string) *AckWatch {
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
func (r *AckWatch) Human(human bool) *AckWatch {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *AckWatch) Pretty(pretty bool) *AckWatch {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
