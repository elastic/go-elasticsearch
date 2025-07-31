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

// Create a transform.
// Creates a transform.
//
// A transform copies data from source indices, transforms it, and persists it
// into an entity-centric destination index. You can also think of the
// destination index as a two-dimensional tabular data structure (known as
// a data frame). The ID for each document in the data frame is generated from a
// hash of the entity, so there is a
// unique row per entity.
//
// You must choose either the latest or pivot method for your transform; you
// cannot use both in a single transform. If
// you choose to use the pivot method for your transform, the entities are
// defined by the set of `group_by` fields in
// the pivot object. If you choose to use the latest method, the entities are
// defined by the `unique_key` field values
// in the latest object.
//
// You must have `create_index`, `index`, and `read` privileges on the
// destination index and `read` and
// `view_index_metadata` privileges on the source indices. When Elasticsearch
// security features are enabled, the
// transform remembers which roles the user that created it had at the time of
// creation and uses those same roles. If
// those roles do not have the required privileges on the source and destination
// indices, the transform fails when it
// attempts unauthorized operations.
//
// NOTE: You must use Kibana or this API to create a transform. Do not add a
// transform directly into any
// `.transform-internal*` indices using the Elasticsearch index API. If
// Elasticsearch security features are enabled, do
// not give users any privileges on `.transform-internal*` indices. If you used
// transforms prior to 7.5, also do not
// give users any privileges on `.data-frame-internal*` indices.
package puttransform

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

type PutTransform struct {
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

// NewPutTransform type alias for index.
type NewPutTransform func(transformid string) *PutTransform

// NewPutTransformFunc returns a new instance of PutTransform with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutTransformFunc(tp elastictransport.Interface) NewPutTransform {
	return func(transformid string) *PutTransform {
		n := New(tp)

		n._transformid(transformid)

		return n
	}
}

// Create a transform.
// Creates a transform.
//
// A transform copies data from source indices, transforms it, and persists it
// into an entity-centric destination index. You can also think of the
// destination index as a two-dimensional tabular data structure (known as
// a data frame). The ID for each document in the data frame is generated from a
// hash of the entity, so there is a
// unique row per entity.
//
// You must choose either the latest or pivot method for your transform; you
// cannot use both in a single transform. If
// you choose to use the pivot method for your transform, the entities are
// defined by the set of `group_by` fields in
// the pivot object. If you choose to use the latest method, the entities are
// defined by the `unique_key` field values
// in the latest object.
//
// You must have `create_index`, `index`, and `read` privileges on the
// destination index and `read` and
// `view_index_metadata` privileges on the source indices. When Elasticsearch
// security features are enabled, the
// transform remembers which roles the user that created it had at the time of
// creation and uses those same roles. If
// those roles do not have the required privileges on the source and destination
// indices, the transform fails when it
// attempts unauthorized operations.
//
// NOTE: You must use Kibana or this API to create a transform. Do not add a
// transform directly into any
// `.transform-internal*` indices using the Elasticsearch index API. If
// Elasticsearch security features are enabled, do
// not give users any privileges on `.transform-internal*` indices. If you used
// transforms prior to 7.5, also do not
// give users any privileges on `.data-frame-internal*` indices.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-transform.html
func New(tp elastictransport.Interface) *PutTransform {
	r := &PutTransform{
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
func (r *PutTransform) Raw(raw io.Reader) *PutTransform {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutTransform) Request(req *Request) *PutTransform {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutTransform: %w", err)
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
func (r PutTransform) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "transform.put_transform")
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
		instrument.BeforeRequest(req, "transform.put_transform")
		if reader := instrument.RecordRequestBody(ctx, "transform.put_transform", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "transform.put_transform")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutTransform query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a puttransform.Response
func (r PutTransform) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "transform.put_transform")
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

// Header set a key, value pair in the PutTransform headers map.
func (r *PutTransform) Header(key, value string) *PutTransform {
	r.headers.Set(key, value)

	return r
}

// TransformId Identifier for the transform. This identifier can contain lowercase
// alphanumeric characters (a-z and 0-9),
// hyphens, and underscores. It has a 64 character limit and must start and end
// with alphanumeric characters.
// API Name: transformid
func (r *PutTransform) _transformid(transformid string) *PutTransform {
	r.paramSet |= transformidMask
	r.transformid = transformid

	return r
}

// DeferValidation When the transform is created, a series of validations occur to ensure its
// success. For example, there is a
// check for the existence of the source indices and a check that the
// destination index is not part of the source
// index pattern. You can use this parameter to skip the checks, for example
// when the source index does not exist
// until after the transform is created. The validations are always run when you
// start the transform, however, with
// the exception of privilege checks.
// API name: defer_validation
func (r *PutTransform) DeferValidation(defervalidation bool) *PutTransform {
	r.values.Set("defer_validation", strconv.FormatBool(defervalidation))

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and returns an error.
// API name: timeout
func (r *PutTransform) Timeout(duration string) *PutTransform {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutTransform) ErrorTrace(errortrace bool) *PutTransform {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutTransform) FilterPath(filterpaths ...string) *PutTransform {
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
func (r *PutTransform) Human(human bool) *PutTransform {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutTransform) Pretty(pretty bool) *PutTransform {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Description Free text description of the transform.
// API name: description
func (r *PutTransform) Description(description string) *PutTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Description = &description

	return r
}

// Dest The destination for the transform.
// API name: dest
func (r *PutTransform) Dest(dest *types.TransformDestination) *PutTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Dest = *dest

	return r
}

// Frequency The interval between checks for changes in the source indices when the
// transform is running continuously. Also
// determines the retry interval in the event of transient failures while the
// transform is searching or indexing.
// The minimum value is `1s` and the maximum is `1h`.
// API name: frequency
func (r *PutTransform) Frequency(duration types.Duration) *PutTransform {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Frequency = duration

	return r
}

// Latest The latest method transforms the data by finding the latest document for each
// unique key.
// API name: latest
func (r *PutTransform) Latest(latest *types.Latest) *PutTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Latest = latest

	return r
}

// Meta_ Defines optional transform metadata.
// API name: _meta
func (r *PutTransform) Meta_(metadata types.Metadata) *PutTransform {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Meta_ = metadata

	return r
}

// Pivot The pivot method transforms the data by aggregating and grouping it. These
// objects define the group by fields
// and the aggregation to reduce the data.
// API name: pivot
func (r *PutTransform) Pivot(pivot *types.Pivot) *PutTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Pivot = pivot

	return r
}

// RetentionPolicy Defines a retention policy for the transform. Data that meets the defined
// criteria is deleted from the
// destination index.
// API name: retention_policy
func (r *PutTransform) RetentionPolicy(retentionpolicy *types.RetentionPolicyContainer) *PutTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RetentionPolicy = retentionpolicy

	return r
}

// Settings Defines optional transform settings.
// API name: settings
func (r *PutTransform) Settings(settings *types.Settings) *PutTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Settings = settings

	return r
}

// Source The source of the data for the transform.
// API name: source
func (r *PutTransform) Source(source *types.TransformSource) *PutTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Source = *source

	return r
}

// Sync Defines the properties transforms require to run continuously.
// API name: sync
func (r *PutTransform) Sync(sync *types.SyncContainer) *PutTransform {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Sync = sync

	return r
}
