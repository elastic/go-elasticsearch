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

// Get datafeeds configuration info.
// You can get information for multiple datafeeds in a single API request by
// using a comma-separated list of datafeeds or a wildcard expression. You can
// get information for all datafeeds by using `_all`, by specifying `*` as the
// `<feed_id>`, or by omitting the `<feed_id>`.
// This API returns a maximum of 10,000 datafeeds.
package getdatafeeds

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
	datafeedidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetDatafeeds struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	datafeedid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewGetDatafeeds type alias for index.
type NewGetDatafeeds func() *GetDatafeeds

// NewGetDatafeedsFunc returns a new instance of GetDatafeeds with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetDatafeedsFunc(tp elastictransport.Interface) NewGetDatafeeds {
	return func() *GetDatafeeds {
		n := New(tp)

		return n
	}
}

// Get datafeeds configuration info.
// You can get information for multiple datafeeds in a single API request by
// using a comma-separated list of datafeeds or a wildcard expression. You can
// get information for all datafeeds by using `_all`, by specifying `*` as the
// `<feed_id>`, or by omitting the `<feed_id>`.
// This API returns a maximum of 10,000 datafeeds.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed.html
func New(tp elastictransport.Interface) *GetDatafeeds {
	r := &GetDatafeeds{
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
func (r *GetDatafeeds) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == datafeedidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("datafeeds")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "datafeedid", r.datafeedid)
		}
		path.WriteString(r.datafeedid)

		method = http.MethodGet
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("datafeeds")

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
func (r GetDatafeeds) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.get_datafeeds")
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
		instrument.BeforeRequest(req, "ml.get_datafeeds")
		if reader := instrument.RecordRequestBody(ctx, "ml.get_datafeeds", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.get_datafeeds")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetDatafeeds query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getdatafeeds.Response
func (r GetDatafeeds) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_datafeeds")
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
func (r GetDatafeeds) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_datafeeds")
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
		err := fmt.Errorf("an error happened during the GetDatafeeds query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the GetDatafeeds headers map.
func (r *GetDatafeeds) Header(key, value string) *GetDatafeeds {
	r.headers.Set(key, value)

	return r
}

// DatafeedId Identifier for the datafeed. It can be a datafeed identifier or a
// wildcard expression. If you do not specify one of these options, the API
// returns information about all datafeeds.
// API Name: datafeedid
func (r *GetDatafeeds) DatafeedId(datafeedid string) *GetDatafeeds {
	r.paramSet |= datafeedidMask
	r.datafeedid = datafeedid

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
// 1. Contains wildcard expressions and there are no datafeeds that match.
// 2. Contains the `_all` string or no identifiers and there are no matches.
// 3. Contains wildcard expressions and there are only partial matches.
//
// The default value is `true`, which returns an empty `datafeeds` array
// when there are no matches and the subset of results when there are
// partial matches. If this parameter is `false`, the request returns a
// `404` status code when there are no matches or only partial matches.
// API name: allow_no_match
func (r *GetDatafeeds) AllowNoMatch(allownomatch bool) *GetDatafeeds {
	r.values.Set("allow_no_match", strconv.FormatBool(allownomatch))

	return r
}

// ExcludeGenerated Indicates if certain fields should be removed from the configuration on
// retrieval. This allows the configuration to be in an acceptable format to
// be retrieved and then added to another cluster.
// API name: exclude_generated
func (r *GetDatafeeds) ExcludeGenerated(excludegenerated bool) *GetDatafeeds {
	r.values.Set("exclude_generated", strconv.FormatBool(excludegenerated))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetDatafeeds) ErrorTrace(errortrace bool) *GetDatafeeds {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetDatafeeds) FilterPath(filterpaths ...string) *GetDatafeeds {
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
func (r *GetDatafeeds) Human(human bool) *GetDatafeeds {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GetDatafeeds) Pretty(pretty bool) *GetDatafeeds {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
