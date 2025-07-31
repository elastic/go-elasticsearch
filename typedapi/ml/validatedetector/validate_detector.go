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

// Validate an anomaly detection job.
package validatedetector

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/excludefrequent"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ValidateDetector struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewValidateDetector type alias for index.
type NewValidateDetector func() *ValidateDetector

// NewValidateDetectorFunc returns a new instance of ValidateDetector with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewValidateDetectorFunc(tp elastictransport.Interface) NewValidateDetector {
	return func() *ValidateDetector {
		n := New(tp)

		return n
	}
}

// Validate an anomaly detection job.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/v8
func New(tp elastictransport.Interface) *ValidateDetector {
	r := &ValidateDetector{
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
func (r *ValidateDetector) Raw(raw io.Reader) *ValidateDetector {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *ValidateDetector) Request(req *Request) *ValidateDetector {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ValidateDetector) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for ValidateDetector: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")
		path.WriteString("_validate")
		path.WriteString("/")
		path.WriteString("detector")

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
func (r ValidateDetector) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.validate_detector")
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
		instrument.BeforeRequest(req, "ml.validate_detector")
		if reader := instrument.RecordRequestBody(ctx, "ml.validate_detector", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.validate_detector")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ValidateDetector query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a validatedetector.Response
func (r ValidateDetector) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.validate_detector")
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

// Header set a key, value pair in the ValidateDetector headers map.
func (r *ValidateDetector) Header(key, value string) *ValidateDetector {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ValidateDetector) ErrorTrace(errortrace bool) *ValidateDetector {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ValidateDetector) FilterPath(filterpaths ...string) *ValidateDetector {
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
func (r *ValidateDetector) Human(human bool) *ValidateDetector {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ValidateDetector) Pretty(pretty bool) *ValidateDetector {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// ByFieldName The field used to split the data. In particular, this property is used for
// analyzing the splits with respect to their own history. It is used for
// finding unusual values in the context of the split.
// API name: by_field_name
func (r *ValidateDetector) ByFieldName(field string) *ValidateDetector {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ByFieldName = &field

	return r
}

// CustomRules Custom rules enable you to customize the way detectors operate. For example,
// a rule may dictate conditions under which results should be skipped. Kibana
// refers to custom rules as job rules.
// API name: custom_rules
func (r *ValidateDetector) CustomRules(customrules ...types.DetectionRule) *ValidateDetector {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.CustomRules = customrules

	return r
}

// DetectorDescription A description of the detector.
// API name: detector_description
func (r *ValidateDetector) DetectorDescription(detectordescription string) *ValidateDetector {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.DetectorDescription = &detectordescription

	return r
}

// DetectorIndex A unique identifier for the detector. This identifier is based on the order
// of the detectors in the `analysis_config`, starting at zero. If you specify a
// value for this property, it is ignored.
// API name: detector_index
func (r *ValidateDetector) DetectorIndex(detectorindex int) *ValidateDetector {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.DetectorIndex = &detectorindex

	return r
}

// ExcludeFrequent If set, frequent entities are excluded from influencing the anomaly results.
// Entities can be considered frequent over time or frequent in a population. If
// you are working with both over and by fields, you can set `exclude_frequent`
// to `all` for both fields, or to `by` or `over` for those specific fields.
// API name: exclude_frequent
func (r *ValidateDetector) ExcludeFrequent(excludefrequent excludefrequent.ExcludeFrequent) *ValidateDetector {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ExcludeFrequent = &excludefrequent

	return r
}

// FieldName The field that the detector uses in the function. If you use an event rate
// function such as count or rare, do not specify this field. The `field_name`
// cannot contain double quotes or backslashes.
// API name: field_name
func (r *ValidateDetector) FieldName(field string) *ValidateDetector {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.FieldName = &field

	return r
}

// Function The analysis function that is used. For example, `count`, `rare`, `mean`,
// `min`, `max`, or `sum`.
// API name: function
func (r *ValidateDetector) Function(function string) *ValidateDetector {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Function = &function

	return r
}

// OverFieldName The field used to split the data. In particular, this property is used for
// analyzing the splits with respect to the history of all splits. It is used
// for finding unusual values in the population of all splits.
// API name: over_field_name
func (r *ValidateDetector) OverFieldName(field string) *ValidateDetector {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.OverFieldName = &field

	return r
}

// PartitionFieldName The field used to segment the analysis. When you use this property, you have
// completely independent baselines for each value of this field.
// API name: partition_field_name
func (r *ValidateDetector) PartitionFieldName(field string) *ValidateDetector {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.PartitionFieldName = &field

	return r
}

// UseNull Defines whether a new series is used as the null series when there is no
// value for the by or partition fields.
// API name: use_null
func (r *ValidateDetector) UseNull(usenull bool) *ValidateDetector {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.UseNull = &usenull

	return r
}
