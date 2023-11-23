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

// Creates a repository.
package createrepository

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
	repositoryMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CreateRepository struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	repository string
}

// NewCreateRepository type alias for index.
type NewCreateRepository func(repository string) *CreateRepository

// NewCreateRepositoryFunc returns a new instance of CreateRepository with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCreateRepositoryFunc(tp elastictransport.Interface) NewCreateRepository {
	return func(repository string) *CreateRepository {
		n := New(tp)

		n._repository(repository)

		return n
	}
}

// Creates a repository.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
func New(tp elastictransport.Interface) *CreateRepository {
	r := &CreateRepository{
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
func (r *CreateRepository) Raw(raw io.Reader) *CreateRepository {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *CreateRepository) Request(req *Request) *CreateRepository {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *CreateRepository) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for CreateRepository: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == repositoryMask:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")

		path.WriteString(r.repository)

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
func (r CreateRepository) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the CreateRepository query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a createrepository.Response
func (r CreateRepository) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the CreateRepository headers map.
func (r *CreateRepository) Header(key, value string) *CreateRepository {
	r.headers.Set(key, value)

	return r
}

// Repository A repository name
// API Name: repository
func (r *CreateRepository) _repository(repository string) *CreateRepository {
	r.paramSet |= repositoryMask
	r.repository = repository

	return r
}

// MasterTimeout Explicit operation timeout for connection to master node
// API name: master_timeout
func (r *CreateRepository) MasterTimeout(duration string) *CreateRepository {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *CreateRepository) Timeout(duration string) *CreateRepository {
	r.values.Set("timeout", duration)

	return r
}

// Verify Whether to verify the repository after creation
// API name: verify
func (r *CreateRepository) Verify(verify bool) *CreateRepository {
	r.values.Set("verify", strconv.FormatBool(verify))

	return r
}

// API name: settings
func (r *CreateRepository) Settings(settings *types.RepositorySettings) *CreateRepository {

	r.req.Settings = *settings

	return r
}

// API name: type
func (r *CreateRepository) Type(type_ string) *CreateRepository {

	r.req.Type = type_

	return r
}
