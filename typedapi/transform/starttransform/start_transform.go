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

// Start a transform.
//
// When you start a transform, it creates the destination index if it does not
// already exist. The `number_of_shards` is
// set to `1` and the `auto_expand_replicas` is set to `0-1`. If it is a pivot
// transform, it deduces the mapping
// definitions for the destination index from the source indices and the
// transform aggregations. If fields in the
// destination index are derived from scripts (as in the case of
// `scripted_metric` or `bucket_script` aggregations),
// the transform uses dynamic mappings unless an index template exists. If it is
// a latest transform, it does not deduce
// mapping definitions; it uses dynamic mappings. To use explicit mappings,
// create the destination index before you
// start the transform. Alternatively, you can create an index template, though
// it does not affect the deduced mappings
// in a pivot transform.
//
// When the transform starts, a series of validations occur to ensure its
// success. If you deferred validation when you
// created the transform, they occur when you start the transform—​with the
// exception of privilege checks. When
// Elasticsearch security features are enabled, the transform remembers which
// roles the user that created it had at the
// time of creation and uses those same roles. If those roles do not have the
// required privileges on the source and
// destination indices, the transform fails when it attempts unauthorized
// operations.
package starttransform

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
	transformidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type StartTransform struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	transformid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewStartTransform type alias for index.
type NewStartTransform func(transformid string) *StartTransform

// NewStartTransformFunc returns a new instance of StartTransform with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStartTransformFunc(tp elastictransport.Interface) NewStartTransform {
	return func(transformid string) *StartTransform {
		n := New(tp)

		n._transformid(transformid)

		return n
	}
}

// Start a transform.
//
// When you start a transform, it creates the destination index if it does not
// already exist. The `number_of_shards` is
// set to `1` and the `auto_expand_replicas` is set to `0-1`. If it is a pivot
// transform, it deduces the mapping
// definitions for the destination index from the source indices and the
// transform aggregations. If fields in the
// destination index are derived from scripts (as in the case of
// `scripted_metric` or `bucket_script` aggregations),
// the transform uses dynamic mappings unless an index template exists. If it is
// a latest transform, it does not deduce
// mapping definitions; it uses dynamic mappings. To use explicit mappings,
// create the destination index before you
// start the transform. Alternatively, you can create an index template, though
// it does not affect the deduced mappings
// in a pivot transform.
//
// When the transform starts, a series of validations occur to ensure its
// success. If you deferred validation when you
// created the transform, they occur when you start the transform—​with the
// exception of privilege checks. When
// Elasticsearch security features are enabled, the transform remembers which
// roles the user that created it had at the
// time of creation and uses those same roles. If those roles do not have the
// required privileges on the source and
// destination indices, the transform fails when it attempts unauthorized
// operations.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-transform.html
func New(tp elastictransport.Interface) *StartTransform {
	r := &StartTransform{
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
func (r *StartTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

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
		path.WriteString("_start")

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r StartTransform) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "transform.start_transform")
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
		instrument.BeforeRequest(req, "transform.start_transform")
		if reader := instrument.RecordRequestBody(ctx, "transform.start_transform", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "transform.start_transform")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the StartTransform query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a starttransform.Response
func (r StartTransform) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "transform.start_transform")
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
func (r StartTransform) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "transform.start_transform")
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
		err := fmt.Errorf("an error happened during the StartTransform query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the StartTransform headers map.
func (r *StartTransform) Header(key, value string) *StartTransform {
	r.headers.Set(key, value)

	return r
}

// TransformId Identifier for the transform.
// API Name: transformid
func (r *StartTransform) _transformid(transformid string) *StartTransform {
	r.paramSet |= transformidMask
	r.transformid = transformid

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and returns an error.
// API name: timeout
func (r *StartTransform) Timeout(duration string) *StartTransform {
	r.values.Set("timeout", duration)

	return r
}

// From Restricts the set of transformed entities to those changed after this time.
// Relative times like now-30d are supported. Only applicable for continuous
// transforms.
// API name: from
func (r *StartTransform) From(from string) *StartTransform {
	r.values.Set("from", from)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *StartTransform) ErrorTrace(errortrace bool) *StartTransform {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *StartTransform) FilterPath(filterpaths ...string) *StartTransform {
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
func (r *StartTransform) Human(human bool) *StartTransform {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *StartTransform) Pretty(pretty bool) *StartTransform {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
