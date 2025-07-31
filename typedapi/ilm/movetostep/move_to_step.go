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

// Move to a lifecycle step.
// Manually move an index into a specific step in the lifecycle policy and run
// that step.
//
// WARNING: This operation can result in the loss of data. Manually moving an
// index into a specific step runs that step even if it has already been
// performed. This is a potentially destructive action and this should be
// considered an expert level API.
//
// You must specify both the current step and the step to be executed in the
// body of the request.
// The request will fail if the current step does not match the step currently
// running for the index
// This is to prevent the index from being moved from an unexpected step into
// the next step.
//
// When specifying the target (`next_step`) to which the index will be moved,
// either the name or both the action and name fields are optional.
// If only the phase is specified, the index will move to the first step of the
// first action in the target phase.
// If the phase and action are specified, the index will move to the first step
// of the specified action in the specified phase.
// Only actions specified in the ILM policy are considered valid.
// An index cannot move to a step that is not part of its policy.
package movetostep

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
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MoveToStep struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewMoveToStep type alias for index.
type NewMoveToStep func(index string) *MoveToStep

// NewMoveToStepFunc returns a new instance of MoveToStep with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMoveToStepFunc(tp elastictransport.Interface) NewMoveToStep {
	return func(index string) *MoveToStep {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Move to a lifecycle step.
// Manually move an index into a specific step in the lifecycle policy and run
// that step.
//
// WARNING: This operation can result in the loss of data. Manually moving an
// index into a specific step runs that step even if it has already been
// performed. This is a potentially destructive action and this should be
// considered an expert level API.
//
// You must specify both the current step and the step to be executed in the
// body of the request.
// The request will fail if the current step does not match the step currently
// running for the index
// This is to prevent the index from being moved from an unexpected step into
// the next step.
//
// When specifying the target (`next_step`) to which the index will be moved,
// either the name or both the action and name fields are optional.
// If only the phase is specified, the index will move to the first step of the
// first action in the target phase.
// If the phase and action are specified, the index will move to the first step
// of the specified action in the specified phase.
// Only actions specified in the ILM policy are considered valid.
// An index cannot move to a step that is not part of its policy.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-move-to-step.html
func New(tp elastictransport.Interface) *MoveToStep {
	r := &MoveToStep{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *MoveToStep) Raw(raw io.Reader) *MoveToStep {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *MoveToStep) Request(req *Request) *MoveToStep {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *MoveToStep) HttpRequest(ctx context.Context) (*http.Request, error) {
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for MoveToStep: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")
		path.WriteString("_ilm")
		path.WriteString("/")
		path.WriteString("move")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)

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

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
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
func (r MoveToStep) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ilm.move_to_step")
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
		instrument.BeforeRequest(req, "ilm.move_to_step")
		if reader := instrument.RecordRequestBody(ctx, "ilm.move_to_step", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ilm.move_to_step")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the MoveToStep query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a movetostep.Response
func (r MoveToStep) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ilm.move_to_step")
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

// Header set a key, value pair in the MoveToStep headers map.
func (r *MoveToStep) Header(key, value string) *MoveToStep {
	r.headers.Set(key, value)

	return r
}

// Index The name of the index whose lifecycle step is to change
// API Name: index
func (r *MoveToStep) _index(index string) *MoveToStep {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *MoveToStep) ErrorTrace(errortrace bool) *MoveToStep {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *MoveToStep) FilterPath(filterpaths ...string) *MoveToStep {
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
func (r *MoveToStep) Human(human bool) *MoveToStep {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *MoveToStep) Pretty(pretty bool) *MoveToStep {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// CurrentStep The step that the index is expected to be in.
// API name: current_step
func (r *MoveToStep) CurrentStep(currentstep *types.StepKey) *MoveToStep {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.CurrentStep = *currentstep

	return r
}

// NextStep The step that you want to run.
// API name: next_step
func (r *MoveToStep) NextStep(nextstep *types.StepKey) *MoveToStep {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.NextStep = *nextstep

	return r
}
