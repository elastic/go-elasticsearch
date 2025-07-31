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

// Get trained models.
//
// Get configuration and usage information about inference trained models.
//
// IMPORTANT: CAT APIs are only intended for human consumption using the Kibana
// console or command line. They are not intended for use by applications. For
// application consumption, use the get trained models statistics API.
package mltrainedmodels

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/cattrainedmodelscolumn"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeunit"
)

const (
	modelidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MlTrainedModels struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	modelid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewMlTrainedModels type alias for index.
type NewMlTrainedModels func() *MlTrainedModels

// NewMlTrainedModelsFunc returns a new instance of MlTrainedModels with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMlTrainedModelsFunc(tp elastictransport.Interface) NewMlTrainedModels {
	return func() *MlTrainedModels {
		n := New(tp)

		return n
	}
}

// Get trained models.
//
// Get configuration and usage information about inference trained models.
//
// IMPORTANT: CAT APIs are only intended for human consumption using the Kibana
// console or command line. They are not intended for use by applications. For
// application consumption, use the get trained models statistics API.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-trained-model.html
func New(tp elastictransport.Interface) *MlTrainedModels {
	r := &MlTrainedModels{
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
func (r *MlTrainedModels) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("ml")
		path.WriteString("/")
		path.WriteString("trained_models")

		method = http.MethodGet
	case r.paramSet == modelidMask:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "modelid", r.modelid)
		}
		path.WriteString(r.modelid)

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
func (r MlTrainedModels) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cat.ml_trained_models")
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
		instrument.BeforeRequest(req, "cat.ml_trained_models")
		if reader := instrument.RecordRequestBody(ctx, "cat.ml_trained_models", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cat.ml_trained_models")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the MlTrainedModels query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a mltrainedmodels.Response
func (r MlTrainedModels) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.ml_trained_models")
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
		err = json.NewDecoder(res.Body).Decode(&response)
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
func (r MlTrainedModels) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.ml_trained_models")
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
		err := fmt.Errorf("an error happened during the MlTrainedModels query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the MlTrainedModels headers map.
func (r *MlTrainedModels) Header(key, value string) *MlTrainedModels {
	r.headers.Set(key, value)

	return r
}

// ModelId A unique identifier for the trained model.
// API Name: modelid
func (r *MlTrainedModels) ModelId(modelid string) *MlTrainedModels {
	r.paramSet |= modelidMask
	r.modelid = modelid

	return r
}

// AllowNoMatch Specifies what to do when the request: contains wildcard expressions and
// there are no models that match; contains the `_all` string or no identifiers
// and there are no matches; contains wildcard expressions and there are only
// partial matches.
// If `true`, the API returns an empty array when there are no matches and the
// subset of results when there are partial matches.
// If `false`, the API returns a 404 status code when there are no matches or
// only partial matches.
// API name: allow_no_match
func (r *MlTrainedModels) AllowNoMatch(allownomatch bool) *MlTrainedModels {
	r.values.Set("allow_no_match", strconv.FormatBool(allownomatch))

	return r
}

// Bytes The unit used to display byte values.
// API name: bytes
func (r *MlTrainedModels) Bytes(bytes bytes.Bytes) *MlTrainedModels {
	r.values.Set("bytes", bytes.String())

	return r
}

// H A comma-separated list of column names to display.
// API name: h
func (r *MlTrainedModels) H(cattrainedmodelscolumns ...cattrainedmodelscolumn.CatTrainedModelsColumn) *MlTrainedModels {
	tmp := []string{}
	for _, item := range cattrainedmodelscolumns {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// S A comma-separated list of column names or aliases used to sort the response.
// API name: s
func (r *MlTrainedModels) S(cattrainedmodelscolumns ...cattrainedmodelscolumn.CatTrainedModelsColumn) *MlTrainedModels {
	tmp := []string{}
	for _, item := range cattrainedmodelscolumns {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// From Skips the specified number of transforms.
// API name: from
func (r *MlTrainedModels) From(from int) *MlTrainedModels {
	r.values.Set("from", strconv.Itoa(from))

	return r
}

// Size The maximum number of transforms to display.
// API name: size
func (r *MlTrainedModels) Size(size int) *MlTrainedModels {
	r.values.Set("size", strconv.Itoa(size))

	return r
}

// Time Unit used to display time values.
// API name: time
func (r *MlTrainedModels) Time(time timeunit.TimeUnit) *MlTrainedModels {
	r.values.Set("time", time.String())

	return r
}

// Format Specifies the format to return the columnar data in, can be set to
// `text`, `json`, `cbor`, `yaml`, or `smile`.
// API name: format
func (r *MlTrainedModels) Format(format string) *MlTrainedModels {
	r.values.Set("format", format)

	return r
}

// Help When set to `true` will output available columns. This option
// can't be combined with any other query string option.
// API name: help
func (r *MlTrainedModels) Help(help bool) *MlTrainedModels {
	r.values.Set("help", strconv.FormatBool(help))

	return r
}

// V When set to `true` will enable verbose output.
// API name: v
func (r *MlTrainedModels) V(v bool) *MlTrainedModels {
	r.values.Set("v", strconv.FormatBool(v))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *MlTrainedModels) ErrorTrace(errortrace bool) *MlTrainedModels {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *MlTrainedModels) FilterPath(filterpaths ...string) *MlTrainedModels {
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
func (r *MlTrainedModels) Human(human bool) *MlTrainedModels {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *MlTrainedModels) Pretty(pretty bool) *MlTrainedModels {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
