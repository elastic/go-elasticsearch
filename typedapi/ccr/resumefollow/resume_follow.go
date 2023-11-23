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

// Resumes a follower index that has been paused
package resumefollow

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
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ResumeFollow struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	index string
}

// NewResumeFollow type alias for index.
type NewResumeFollow func(index string) *ResumeFollow

// NewResumeFollowFunc returns a new instance of ResumeFollow with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewResumeFollowFunc(tp elastictransport.Interface) NewResumeFollow {
	return func(index string) *ResumeFollow {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Resumes a follower index that has been paused
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-resume-follow.html
func New(tp elastictransport.Interface) *ResumeFollow {
	r := &ResumeFollow{
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
func (r *ResumeFollow) Raw(raw io.Reader) *ResumeFollow {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *ResumeFollow) Request(req *Request) *ResumeFollow {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ResumeFollow) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for ResumeFollow: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_ccr")
		path.WriteString("/")
		path.WriteString("resume_follow")

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
func (r ResumeFollow) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the ResumeFollow query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a resumefollow.Response
func (r ResumeFollow) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the ResumeFollow headers map.
func (r *ResumeFollow) Header(key, value string) *ResumeFollow {
	r.headers.Set(key, value)

	return r
}

// Index The name of the follow index to resume following.
// API Name: index
func (r *ResumeFollow) _index(index string) *ResumeFollow {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// API name: max_outstanding_read_requests
func (r *ResumeFollow) MaxOutstandingReadRequests(maxoutstandingreadrequests int64) *ResumeFollow {

	r.req.MaxOutstandingReadRequests = &maxoutstandingreadrequests

	return r
}

// API name: max_outstanding_write_requests
func (r *ResumeFollow) MaxOutstandingWriteRequests(maxoutstandingwriterequests int64) *ResumeFollow {

	r.req.MaxOutstandingWriteRequests = &maxoutstandingwriterequests

	return r
}

// API name: max_read_request_operation_count
func (r *ResumeFollow) MaxReadRequestOperationCount(maxreadrequestoperationcount int64) *ResumeFollow {

	r.req.MaxReadRequestOperationCount = &maxreadrequestoperationcount

	return r
}

// API name: max_read_request_size
func (r *ResumeFollow) MaxReadRequestSize(maxreadrequestsize string) *ResumeFollow {

	r.req.MaxReadRequestSize = &maxreadrequestsize

	return r
}

// API name: max_retry_delay
func (r *ResumeFollow) MaxRetryDelay(duration types.Duration) *ResumeFollow {
	r.req.MaxRetryDelay = duration

	return r
}

// API name: max_write_buffer_count
func (r *ResumeFollow) MaxWriteBufferCount(maxwritebuffercount int64) *ResumeFollow {

	r.req.MaxWriteBufferCount = &maxwritebuffercount

	return r
}

// API name: max_write_buffer_size
func (r *ResumeFollow) MaxWriteBufferSize(maxwritebuffersize string) *ResumeFollow {

	r.req.MaxWriteBufferSize = &maxwritebuffersize

	return r
}

// API name: max_write_request_operation_count
func (r *ResumeFollow) MaxWriteRequestOperationCount(maxwriterequestoperationcount int64) *ResumeFollow {

	r.req.MaxWriteRequestOperationCount = &maxwriterequestoperationcount

	return r
}

// API name: max_write_request_size
func (r *ResumeFollow) MaxWriteRequestSize(maxwriterequestsize string) *ResumeFollow {

	r.req.MaxWriteRequestSize = &maxwriterequestsize

	return r
}

// API name: read_poll_timeout
func (r *ResumeFollow) ReadPollTimeout(duration types.Duration) *ResumeFollow {
	r.req.ReadPollTimeout = duration

	return r
}
