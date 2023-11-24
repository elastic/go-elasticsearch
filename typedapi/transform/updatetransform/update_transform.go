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

// Updates certain properties of a transform.
package updatetransform

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

type UpdateTransform struct {
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

// NewUpdateTransform type alias for index.
type NewUpdateTransform func(transformid string) *UpdateTransform

// NewUpdateTransformFunc returns a new instance of UpdateTransform with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateTransformFunc(tp elastictransport.Interface) NewUpdateTransform {
	return func(transformid string) *UpdateTransform {
		n := New(tp)

		n._transformid(transformid)

		return n
	}
}

// Updates certain properties of a transform.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-transform.html
func New(tp elastictransport.Interface) *UpdateTransform {
	r := &UpdateTransform{
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
func (r *UpdateTransform) Raw(raw io.Reader) *UpdateTransform {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *UpdateTransform) Request(req *Request) *UpdateTransform {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *UpdateTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for UpdateTransform: %w", err)
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
		path.WriteString("/")
		path.WriteString("_update")

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
func (r UpdateTransform) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the UpdateTransform query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a updatetransform.Response
func (r UpdateTransform) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the UpdateTransform headers map.
func (r *UpdateTransform) Header(key, value string) *UpdateTransform {
	r.headers.Set(key, value)

	return r
}

// TransformId Identifier for the transform.
// API Name: transformid
func (r *UpdateTransform) _transformid(transformid string) *UpdateTransform {
	r.paramSet |= transformidMask
	r.transformid = transformid

	return r
}

// DeferValidation When true, deferrable validations are not run. This behavior may be
// desired if the source index does not exist until after the transform is
// created.
// API name: defer_validation
func (r *UpdateTransform) DeferValidation(defervalidation bool) *UpdateTransform {
	r.values.Set("defer_validation", strconv.FormatBool(defervalidation))

	return r
}

// Timeout Period to wait for a response. If no response is received before the
// timeout expires, the request fails and returns an error.
// API name: timeout
func (r *UpdateTransform) Timeout(duration string) *UpdateTransform {
	r.values.Set("timeout", duration)

	return r
}

// Description Free text description of the transform.
// API name: description
func (r *UpdateTransform) Description(description string) *UpdateTransform {

	r.req.Description = &description

	return r
}

// Dest The destination for the transform.
// API name: dest
func (r *UpdateTransform) Dest(dest *types.TransformDestination) *UpdateTransform {

	r.req.Dest = dest

	return r
}

// Frequency The interval between checks for changes in the source indices when the
// transform is running continuously. Also determines the retry interval in
// the event of transient failures while the transform is searching or
// indexing. The minimum value is 1s and the maximum is 1h.
// API name: frequency
func (r *UpdateTransform) Frequency(duration types.Duration) *UpdateTransform {
	r.req.Frequency = duration

	return r
}

// Meta_ Defines optional transform metadata.
// API name: _meta
func (r *UpdateTransform) Meta_(metadata types.Metadata) *UpdateTransform {
	r.req.Meta_ = metadata

	return r
}

// RetentionPolicy Defines a retention policy for the transform. Data that meets the defined
// criteria is deleted from the destination index.
// API name: retention_policy
func (r *UpdateTransform) RetentionPolicy(retentionpolicy types.RetentionPolicyContainer) *UpdateTransform {
	r.req.RetentionPolicy = retentionpolicy

	return r
}

// Settings Defines optional transform settings.
// API name: settings
func (r *UpdateTransform) Settings(settings *types.Settings) *UpdateTransform {

	r.req.Settings = settings

	return r
}

// Source The source of the data for the transform.
// API name: source
func (r *UpdateTransform) Source(source *types.TransformSource) *UpdateTransform {

	r.req.Source = source

	return r
}

// Sync Defines the properties transforms require to run continuously.
// API name: sync
func (r *UpdateTransform) Sync(sync *types.SyncContainer) *UpdateTransform {

	r.req.Sync = sync

	return r
}
