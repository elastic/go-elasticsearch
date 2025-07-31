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

// Preview a transform.
// Generates a preview of the results that you will get when you create a
// transform with the same configuration.
//
// It returns a maximum of 100 results. The calculations are based on all the
// current data in the source index. It also
// generates a list of mappings and settings for the destination index. These
// values are determined based on the field
// types of the source index and the transform aggregations.
package previewtransform

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
	transformidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PreviewTransform struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	transformid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPreviewTransform type alias for index.
type NewPreviewTransform func() *PreviewTransform

// NewPreviewTransformFunc returns a new instance of PreviewTransform with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPreviewTransformFunc(tp elastictransport.Interface) NewPreviewTransform {
	return func() *PreviewTransform {
		n := New(tp)

		return n
	}
}

// Preview a transform.
// Generates a preview of the results that you will get when you create a
// transform with the same configuration.
//
// It returns a maximum of 100 results. The calculations are based on all the
// current data in the source index. It also
// generates a list of mappings and settings for the destination index. These
// values are determined based on the field
// types of the source index and the transform aggregations.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/preview-transform.html
func New(tp elastictransport.Interface) *PreviewTransform {
	r := &PreviewTransform{
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
func (r *PreviewTransform) Raw(raw io.Reader) *PreviewTransform {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PreviewTransform) Request(req *Request) *PreviewTransform {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PreviewTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PreviewTransform: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == transformidMask:
		path.WriteString("/")
		path.WriteString("_transform")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "transformid", r.transformid)
		}
		path.WriteString(r.transformid)
		path.WriteString("/")
		path.WriteString("_preview")

		method = http.MethodPost
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_transform")
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
func (r PreviewTransform) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "transform.preview_transform")
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
		instrument.BeforeRequest(req, "transform.preview_transform")
		if reader := instrument.RecordRequestBody(ctx, "transform.preview_transform", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "transform.preview_transform")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PreviewTransform query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a previewtransform.Response
func (r PreviewTransform) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "transform.preview_transform")
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

// Header set a key, value pair in the PreviewTransform headers map.
func (r *PreviewTransform) Header(key, value string) *PreviewTransform {
	r.headers.Set(key, value)

	return r
}

// TransformId Identifier for the transform to preview. If you specify this path parameter,
// you cannot provide transform
// configuration details in the request body.
// API Name: transformid
func (r *PreviewTransform) TransformId(transformid string) *PreviewTransform {
	r.paramSet |= transformidMask
	r.transformid = transformid

	return r
}

// Timeout Period to wait for a response. If no response is received before the
// timeout expires, the request fails and returns an error.
// API name: timeout
func (r *PreviewTransform) Timeout(duration string) *PreviewTransform {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PreviewTransform) ErrorTrace(errortrace bool) *PreviewTransform {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PreviewTransform) FilterPath(filterpaths ...string) *PreviewTransform {
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
func (r *PreviewTransform) Human(human bool) *PreviewTransform {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PreviewTransform) Pretty(pretty bool) *PreviewTransform {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Description Free text description of the transform.
// API name: description
func (r *PreviewTransform) Description(description string) *PreviewTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Description = &description

	return r
}

// Dest The destination for the transform.
// API name: dest
func (r *PreviewTransform) Dest(dest *types.TransformDestination) *PreviewTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Dest = dest

	return r
}

// Frequency The interval between checks for changes in the source indices when the
// transform is running continuously. Also determines the retry interval in
// the event of transient failures while the transform is searching or
// indexing. The minimum value is 1s and the maximum is 1h.
// API name: frequency
func (r *PreviewTransform) Frequency(duration types.Duration) *PreviewTransform {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Frequency = duration

	return r
}

// Latest The latest method transforms the data by finding the latest document for
// each unique key.
// API name: latest
func (r *PreviewTransform) Latest(latest *types.Latest) *PreviewTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Latest = latest

	return r
}

// Pivot The pivot method transforms the data by aggregating and grouping it.
// These objects define the group by fields and the aggregation to reduce
// the data.
// API name: pivot
func (r *PreviewTransform) Pivot(pivot *types.Pivot) *PreviewTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Pivot = pivot

	return r
}

// RetentionPolicy Defines a retention policy for the transform. Data that meets the defined
// criteria is deleted from the destination index.
// API name: retention_policy
func (r *PreviewTransform) RetentionPolicy(retentionpolicy *types.RetentionPolicyContainer) *PreviewTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RetentionPolicy = retentionpolicy

	return r
}

// Settings Defines optional transform settings.
// API name: settings
func (r *PreviewTransform) Settings(settings *types.Settings) *PreviewTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Settings = settings

	return r
}

// Source The source of the data for the transform.
// API name: source
func (r *PreviewTransform) Source(source *types.TransformSource) *PreviewTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Source = source

	return r
}

// Sync Defines the properties transforms require to run continuously.
// API name: sync
func (r *PreviewTransform) Sync(sync *types.SyncContainer) *PreviewTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Sync = sync

	return r
}
