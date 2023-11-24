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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Instantiates a transform.
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	transformid string
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

// Instantiates a transform.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-transform.html
func New(tp elastictransport.Interface) *PutTransform {
	r := &PutTransform{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutTransform: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == transformidMask:
		path.WriteString("/")
		path.WriteString("_transform")
		path.WriteString("/")

		path.WriteString(r.transformid)

		method = http.MethodPut
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
func (r PutTransform) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutTransform query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a puttransform.Response
func (r PutTransform) Do(ctx context.Context) (*Response, error) {

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

// Description Free text description of the transform.
// API name: description
func (r *PutTransform) Description(description string) *PutTransform {

	r.req.Description = &description

	return r
}

// Dest The destination for the transform.
// API name: dest
func (r *PutTransform) Dest(dest *types.TransformDestination) *PutTransform {

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
	r.req.Frequency = duration

	return r
}

// Latest The latest method transforms the data by finding the latest document for each
// unique key.
// API name: latest
func (r *PutTransform) Latest(latest *types.Latest) *PutTransform {

	r.req.Latest = latest

	return r
}

// Meta_ Defines optional transform metadata.
// API name: _meta
func (r *PutTransform) Meta_(metadata types.Metadata) *PutTransform {
	r.req.Meta_ = metadata

	return r
}

// Pivot The pivot method transforms the data by aggregating and grouping it. These
// objects define the group by fields
// and the aggregation to reduce the data.
// API name: pivot
func (r *PutTransform) Pivot(pivot *types.Pivot) *PutTransform {

	r.req.Pivot = pivot

	return r
}

// RetentionPolicy Defines a retention policy for the transform. Data that meets the defined
// criteria is deleted from the
// destination index.
// API name: retention_policy
func (r *PutTransform) RetentionPolicy(retentionpolicy *types.RetentionPolicyContainer) *PutTransform {

	r.req.RetentionPolicy = retentionpolicy

	return r
}

// Settings Defines optional transform settings.
// API name: settings
func (r *PutTransform) Settings(settings *types.Settings) *PutTransform {

	r.req.Settings = settings

	return r
}

// Source The source of the data for the transform.
// API name: source
func (r *PutTransform) Source(source *types.TransformSource) *PutTransform {

	r.req.Source = *source

	return r
}

// Sync Defines the properties transforms require to run continuously.
// API name: sync
func (r *PutTransform) Sync(sync *types.SyncContainer) *PutTransform {

	r.req.Sync = sync

	return r
}
