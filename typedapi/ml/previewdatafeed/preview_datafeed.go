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

// Preview a datafeed.
// This API returns the first "page" of search results from a datafeed.
// You can preview an existing datafeed or provide configuration details for a
// datafeed
// and anomaly detection job in the API. The preview shows the structure of the
// data
// that will be passed to the anomaly detection engine.
// IMPORTANT: When Elasticsearch security features are enabled, the preview uses
// the credentials of the user that
// called the API. However, when the datafeed starts it uses the roles of the
// last user that created or updated the
// datafeed. To get a preview that accurately reflects the behavior of the
// datafeed, use the appropriate credentials.
// You can also use secondary authorization headers to supply the credentials.
package previewdatafeed

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
	datafeedidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PreviewDatafeed struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	datafeedid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPreviewDatafeed type alias for index.
type NewPreviewDatafeed func() *PreviewDatafeed

// NewPreviewDatafeedFunc returns a new instance of PreviewDatafeed with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPreviewDatafeedFunc(tp elastictransport.Interface) NewPreviewDatafeed {
	return func() *PreviewDatafeed {
		n := New(tp)

		return n
	}
}

// Preview a datafeed.
// This API returns the first "page" of search results from a datafeed.
// You can preview an existing datafeed or provide configuration details for a
// datafeed
// and anomaly detection job in the API. The preview shows the structure of the
// data
// that will be passed to the anomaly detection engine.
// IMPORTANT: When Elasticsearch security features are enabled, the preview uses
// the credentials of the user that
// called the API. However, when the datafeed starts it uses the roles of the
// last user that created or updated the
// datafeed. To get a preview that accurately reflects the behavior of the
// datafeed, use the appropriate credentials.
// You can also use secondary authorization headers to supply the credentials.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-preview-datafeed.html
func New(tp elastictransport.Interface) *PreviewDatafeed {
	r := &PreviewDatafeed{
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
func (r *PreviewDatafeed) Raw(raw io.Reader) *PreviewDatafeed {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PreviewDatafeed) Request(req *Request) *PreviewDatafeed {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PreviewDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PreviewDatafeed: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

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
		path.WriteString("/")
		path.WriteString("_preview")

		method = http.MethodPost
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("datafeeds")
		path.WriteString("/")
		path.WriteString("_preview")

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
func (r PreviewDatafeed) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.preview_datafeed")
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
		instrument.BeforeRequest(req, "ml.preview_datafeed")
		if reader := instrument.RecordRequestBody(ctx, "ml.preview_datafeed", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.preview_datafeed")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PreviewDatafeed query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a previewdatafeed.Response
func (r PreviewDatafeed) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.preview_datafeed")
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

// Header set a key, value pair in the PreviewDatafeed headers map.
func (r *PreviewDatafeed) Header(key, value string) *PreviewDatafeed {
	r.headers.Set(key, value)

	return r
}

// DatafeedId A numerical character string that uniquely identifies the datafeed. This
// identifier can contain lowercase
// alphanumeric characters (a-z and 0-9), hyphens, and underscores. It must
// start and end with alphanumeric
// characters. NOTE: If you use this path parameter, you cannot provide datafeed
// or anomaly detection job
// configuration details in the request body.
// API Name: datafeedid
func (r *PreviewDatafeed) DatafeedId(datafeedid string) *PreviewDatafeed {
	r.paramSet |= datafeedidMask
	r.datafeedid = datafeedid

	return r
}

// Start The start time from where the datafeed preview should begin
// API name: start
func (r *PreviewDatafeed) Start(datetime string) *PreviewDatafeed {
	r.values.Set("start", datetime)

	return r
}

// End The end time when the datafeed preview should stop
// API name: end
func (r *PreviewDatafeed) End(datetime string) *PreviewDatafeed {
	r.values.Set("end", datetime)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PreviewDatafeed) ErrorTrace(errortrace bool) *PreviewDatafeed {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PreviewDatafeed) FilterPath(filterpaths ...string) *PreviewDatafeed {
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
func (r *PreviewDatafeed) Human(human bool) *PreviewDatafeed {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PreviewDatafeed) Pretty(pretty bool) *PreviewDatafeed {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// DatafeedConfig The datafeed definition to preview.
// API name: datafeed_config
func (r *PreviewDatafeed) DatafeedConfig(datafeedconfig *types.DatafeedConfig) *PreviewDatafeed {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.DatafeedConfig = datafeedconfig

	return r
}

// JobConfig The configuration details for the anomaly detection job that is associated
// with the datafeed. If the
// `datafeed_config` object does not include a `job_id` that references an
// existing anomaly detection job, you must
// supply this `job_config` object. If you include both a `job_id` and a
// `job_config`, the latter information is
// used. You cannot specify a `job_config` object unless you also supply a
// `datafeed_config` object.
// API name: job_config
func (r *PreviewDatafeed) JobConfig(jobconfig *types.JobConfig) *PreviewDatafeed {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.JobConfig = jobconfig

	return r
}
