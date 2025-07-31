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

// Create or update a watch.
// When a watch is registered, a new document that represents the watch is added
// to the `.watches` index and its trigger is immediately registered with the
// relevant trigger engine.
// Typically for the `schedule` trigger, the scheduler is the trigger engine.
//
// IMPORTANT: You must use Kibana or this API to create a watch.
// Do not add a watch directly to the `.watches` index by using the
// Elasticsearch index API.
// If Elasticsearch security features are enabled, do not give users write
// privileges on the `.watches` index.
//
// When you add a watch you can also define its initial active state by setting
// the *active* parameter.
//
// When Elasticsearch security features are enabled, your watch can index or
// search only on indices for which the user that stored the watch has
// privileges.
// If the user is able to read index `a`, but not index `b`, the same will apply
// when the watch runs.
package putwatch

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
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutWatch struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutWatch type alias for index.
type NewPutWatch func(id string) *PutWatch

// NewPutWatchFunc returns a new instance of PutWatch with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutWatchFunc(tp elastictransport.Interface) NewPutWatch {
	return func(id string) *PutWatch {
		n := New(tp)

		n._id(id)

		return n
	}
}

// Create or update a watch.
// When a watch is registered, a new document that represents the watch is added
// to the `.watches` index and its trigger is immediately registered with the
// relevant trigger engine.
// Typically for the `schedule` trigger, the scheduler is the trigger engine.
//
// IMPORTANT: You must use Kibana or this API to create a watch.
// Do not add a watch directly to the `.watches` index by using the
// Elasticsearch index API.
// If Elasticsearch security features are enabled, do not give users write
// privileges on the `.watches` index.
//
// When you add a watch you can also define its initial active state by setting
// the *active* parameter.
//
// When Elasticsearch security features are enabled, your watch can index or
// search only on indices for which the user that stored the watch has
// privileges.
// If the user is able to read index `a`, but not index `b`, the same will apply
// when the watch runs.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-watcher-put-watch
func New(tp elastictransport.Interface) *PutWatch {
	r := &PutWatch{
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
func (r *PutWatch) Raw(raw io.Reader) *PutWatch {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutWatch) Request(req *Request) *PutWatch {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutWatch) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutWatch: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_watcher")
		path.WriteString("/")
		path.WriteString("watch")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

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
func (r PutWatch) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "watcher.put_watch")
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
		instrument.BeforeRequest(req, "watcher.put_watch")
		if reader := instrument.RecordRequestBody(ctx, "watcher.put_watch", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "watcher.put_watch")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutWatch query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putwatch.Response
func (r PutWatch) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "watcher.put_watch")
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

// Header set a key, value pair in the PutWatch headers map.
func (r *PutWatch) Header(key, value string) *PutWatch {
	r.headers.Set(key, value)

	return r
}

// Id The identifier for the watch.
// API Name: id
func (r *PutWatch) _id(id string) *PutWatch {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Active The initial state of the watch.
// The default value is `true`, which means the watch is active by default.
// API name: active
func (r *PutWatch) Active(active bool) *PutWatch {
	r.values.Set("active", strconv.FormatBool(active))

	return r
}

// IfPrimaryTerm only update the watch if the last operation that has changed the watch has
// the specified primary term
// API name: if_primary_term
func (r *PutWatch) IfPrimaryTerm(ifprimaryterm string) *PutWatch {
	r.values.Set("if_primary_term", ifprimaryterm)

	return r
}

// IfSeqNo only update the watch if the last operation that has changed the watch has
// the specified sequence number
// API name: if_seq_no
func (r *PutWatch) IfSeqNo(sequencenumber string) *PutWatch {
	r.values.Set("if_seq_no", sequencenumber)

	return r
}

// Version Explicit version number for concurrency control
// API name: version
func (r *PutWatch) Version(versionnumber string) *PutWatch {
	r.values.Set("version", versionnumber)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutWatch) ErrorTrace(errortrace bool) *PutWatch {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutWatch) FilterPath(filterpaths ...string) *PutWatch {
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
func (r *PutWatch) Human(human bool) *PutWatch {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutWatch) Pretty(pretty bool) *PutWatch {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Actions The list of actions that will be run if the condition matches.
// API name: actions
func (r *PutWatch) Actions(actions map[string]types.WatcherAction) *PutWatch {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Actions = actions

	return r
}

// Condition The condition that defines if the actions should be run.
// API name: condition
func (r *PutWatch) Condition(condition *types.WatcherCondition) *PutWatch {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Condition = condition

	return r
}

// Input The input that defines the input that loads the data for the watch.
// API name: input
func (r *PutWatch) Input(input *types.WatcherInput) *PutWatch {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Input = input

	return r
}

// Metadata Metadata JSON that will be copied into the history entries.
// API name: metadata
func (r *PutWatch) Metadata(metadata types.Metadata) *PutWatch {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Metadata = metadata

	return r
}

// ThrottlePeriod The minimum time between actions being run.
// The default is 5 seconds.
// This default can be changed in the config file with the setting
// `xpack.watcher.throttle.period.default_period`.
// If both this value and the `throttle_period_in_millis` parameter are
// specified, Watcher uses the last parameter included in the request.
// API name: throttle_period
func (r *PutWatch) ThrottlePeriod(duration types.Duration) *PutWatch {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ThrottlePeriod = duration

	return r
}

// ThrottlePeriodInMillis Minimum time in milliseconds between actions being run. Defaults to 5000. If
// both this value and the throttle_period parameter are specified, Watcher uses
// the last parameter included in the request.
// API name: throttle_period_in_millis
func (r *PutWatch) ThrottlePeriodInMillis(durationvalueunitmillis int64) *PutWatch {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ThrottlePeriodInMillis = &durationvalueunitmillis

	return r
}

// Transform The transform that processes the watch payload to prepare it for the watch
// actions.
// API name: transform
func (r *PutWatch) Transform(transform *types.TransformContainer) *PutWatch {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Transform = transform

	return r
}

// Trigger The trigger that defines when the watch should run.
// API name: trigger
func (r *PutWatch) Trigger(trigger *types.TriggerContainer) *PutWatch {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Trigger = trigger

	return r
}
