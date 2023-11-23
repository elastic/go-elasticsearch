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

// Stops one or more transforms.
package stoptransform

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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

type StopTransform struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	transformid string
}

// NewStopTransform type alias for index.
type NewStopTransform func(transformid string) *StopTransform

// NewStopTransformFunc returns a new instance of StopTransform with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStopTransformFunc(tp elastictransport.Interface) NewStopTransform {
	return func(transformid string) *StopTransform {
		n := New(tp)

		n._transformid(transformid)

		return n
	}
}

// Stops one or more transforms.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-transform.html
func New(tp elastictransport.Interface) *StopTransform {
	r := &StopTransform{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *StopTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
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

		path.WriteString(r.transformid)
		path.WriteString("/")
		path.WriteString("_stop")

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r StopTransform) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the StopTransform query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a stoptransform.Response
func (r StopTransform) Do(ctx context.Context) (*Response, error) {

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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r StopTransform) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the StopTransform headers map.
func (r *StopTransform) Header(key, value string) *StopTransform {
	r.headers.Set(key, value)

	return r
}

// TransformId Identifier for the transform. To stop multiple transforms, use a
// comma-separated list or a wildcard expression.
// To stop all transforms, use `_all` or `*` as the identifier.
// API Name: transformid
func (r *StopTransform) _transformid(transformid string) *StopTransform {
	r.paramSet |= transformidMask
	r.transformid = transformid

	return r
}

// AllowNoMatch Specifies what to do when the request: contains wildcard expressions and
// there are no transforms that match;
// contains the `_all` string or no identifiers and there are no matches;
// contains wildcard expressions and there
// are only partial matches.
//
// If it is true, the API returns a successful acknowledgement message when
// there are no matches. When there are
// only partial matches, the API stops the appropriate transforms.
//
// If it is false, the request returns a 404 status code when there are no
// matches or only partial matches.
// API name: allow_no_match
func (r *StopTransform) AllowNoMatch(allownomatch bool) *StopTransform {
	r.values.Set("allow_no_match", strconv.FormatBool(allownomatch))

	return r
}

// Force If it is true, the API forcefully stops the transforms.
// API name: force
func (r *StopTransform) Force(force bool) *StopTransform {
	r.values.Set("force", strconv.FormatBool(force))

	return r
}

// Timeout Period to wait for a response when `wait_for_completion` is `true`. If no
// response is received before the
// timeout expires, the request returns a timeout exception. However, the
// request continues processing and
// eventually moves the transform to a STOPPED state.
// API name: timeout
func (r *StopTransform) Timeout(duration string) *StopTransform {
	r.values.Set("timeout", duration)

	return r
}

// WaitForCheckpoint If it is true, the transform does not completely stop until the current
// checkpoint is completed. If it is false,
// the transform stops as soon as possible.
// API name: wait_for_checkpoint
func (r *StopTransform) WaitForCheckpoint(waitforcheckpoint bool) *StopTransform {
	r.values.Set("wait_for_checkpoint", strconv.FormatBool(waitforcheckpoint))

	return r
}

// WaitForCompletion If it is true, the API blocks until the indexer state completely stops. If it
// is false, the API returns
// immediately and the indexer is stopped asynchronously in the background.
// API name: wait_for_completion
func (r *StopTransform) WaitForCompletion(waitforcompletion bool) *StopTransform {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}
