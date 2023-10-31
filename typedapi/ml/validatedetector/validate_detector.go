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

// Validates an anomaly detection detector.
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int
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

// Validates an anomaly detection detector.
//
// https://www.elastic.co/guide/en/machine-learning/current/ml-jobs.html
func New(tp elastictransport.Interface) *ValidateDetector {
	r := &ValidateDetector{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for ValidateDetector: %w", err)
		}

		r.buf.Write(data)

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
func (r ValidateDetector) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the ValidateDetector query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a validatedetector.Response
func (r ValidateDetector) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the ValidateDetector headers map.
func (r *ValidateDetector) Header(key, value string) *ValidateDetector {
	r.headers.Set(key, value)

	return r
}

// ByFieldName The field used to split the data. In particular, this property is used for
// analyzing the splits with respect to their own history. It is used for
// finding unusual values in the context of the split.
// API name: by_field_name
func (r *ValidateDetector) ByFieldName(field string) *ValidateDetector {
	r.req.ByFieldName = &field

	return r
}

// CustomRules Custom rules enable you to customize the way detectors operate. For example,
// a rule may dictate conditions under which results should be skipped. Kibana
// refers to custom rules as job rules.
// API name: custom_rules
func (r *ValidateDetector) CustomRules(customrules ...types.DetectionRule) *ValidateDetector {
	r.req.CustomRules = customrules

	return r
}

// DetectorDescription A description of the detector.
// API name: detector_description
func (r *ValidateDetector) DetectorDescription(detectordescription string) *ValidateDetector {

	r.req.DetectorDescription = &detectordescription

	return r
}

// DetectorIndex A unique identifier for the detector. This identifier is based on the order
// of the detectors in the `analysis_config`, starting at zero. If you specify a
// value for this property, it is ignored.
// API name: detector_index
func (r *ValidateDetector) DetectorIndex(detectorindex int) *ValidateDetector {
	r.req.DetectorIndex = &detectorindex

	return r
}

// ExcludeFrequent If set, frequent entities are excluded from influencing the anomaly results.
// Entities can be considered frequent over time or frequent in a population. If
// you are working with both over and by fields, you can set `exclude_frequent`
// to `all` for both fields, or to `by` or `over` for those specific fields.
// API name: exclude_frequent
func (r *ValidateDetector) ExcludeFrequent(excludefrequent excludefrequent.ExcludeFrequent) *ValidateDetector {
	r.req.ExcludeFrequent = &excludefrequent

	return r
}

// FieldName The field that the detector uses in the function. If you use an event rate
// function such as count or rare, do not specify this field. The `field_name`
// cannot contain double quotes or backslashes.
// API name: field_name
func (r *ValidateDetector) FieldName(field string) *ValidateDetector {
	r.req.FieldName = &field

	return r
}

// Function The analysis function that is used. For example, `count`, `rare`, `mean`,
// `min`, `max`, or `sum`.
// API name: function
func (r *ValidateDetector) Function(function string) *ValidateDetector {

	r.req.Function = &function

	return r
}

// OverFieldName The field used to split the data. In particular, this property is used for
// analyzing the splits with respect to the history of all splits. It is used
// for finding unusual values in the population of all splits.
// API name: over_field_name
func (r *ValidateDetector) OverFieldName(field string) *ValidateDetector {
	r.req.OverFieldName = &field

	return r
}

// PartitionFieldName The field used to segment the analysis. When you use this property, you have
// completely independent baselines for each value of this field.
// API name: partition_field_name
func (r *ValidateDetector) PartitionFieldName(field string) *ValidateDetector {
	r.req.PartitionFieldName = &field

	return r
}

// UseNull Defines whether a new series is used as the null series when there is no
// value for the by or partition fields.
// API name: use_null
func (r *ValidateDetector) UseNull(usenull bool) *ValidateDetector {
	r.req.UseNull = &usenull

	return r
}
