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

// Creates or updates a snapshot lifecycle policy.
package putlifecycle

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
	policyidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutLifecycle struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	policyid string
}

// NewPutLifecycle type alias for index.
type NewPutLifecycle func(policyid string) *PutLifecycle

// NewPutLifecycleFunc returns a new instance of PutLifecycle with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutLifecycleFunc(tp elastictransport.Interface) NewPutLifecycle {
	return func(policyid string) *PutLifecycle {
		n := New(tp)

		n._policyid(policyid)

		return n
	}
}

// Creates or updates a snapshot lifecycle policy.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-put-policy.html
func New(tp elastictransport.Interface) *PutLifecycle {
	r := &PutLifecycle{
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
func (r *PutLifecycle) Raw(raw io.Reader) *PutLifecycle {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutLifecycle) Request(req *Request) *PutLifecycle {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutLifecycle: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == policyidMask:
		path.WriteString("/")
		path.WriteString("_slm")
		path.WriteString("/")
		path.WriteString("policy")
		path.WriteString("/")

		path.WriteString(r.policyid)

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
func (r PutLifecycle) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutLifecycle query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putlifecycle.Response
func (r PutLifecycle) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutLifecycle headers map.
func (r *PutLifecycle) Header(key, value string) *PutLifecycle {
	r.headers.Set(key, value)

	return r
}

// PolicyId ID for the snapshot lifecycle policy you want to create or update.
// API Name: policyid
func (r *PutLifecycle) _policyid(policyid string) *PutLifecycle {
	r.paramSet |= policyidMask
	r.policyid = policyid

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *PutLifecycle) MasterTimeout(duration string) *PutLifecycle {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and returns an error.
// API name: timeout
func (r *PutLifecycle) Timeout(duration string) *PutLifecycle {
	r.values.Set("timeout", duration)

	return r
}

// Config Configuration for each snapshot created by the policy.
// API name: config
func (r *PutLifecycle) Config(config *types.Configuration) *PutLifecycle {

	r.req.Config = config

	return r
}

// Name Name automatically assigned to each snapshot created by the policy. Date math
// is supported. To prevent conflicting snapshot names, a UUID is automatically
// appended to each snapshot name.
// API name: name
func (r *PutLifecycle) Name(name string) *PutLifecycle {
	r.req.Name = &name

	return r
}

// Repository Repository used to store snapshots created by this policy. This repository
// must exist prior to the policy’s creation. You can create a repository using
// the snapshot repository API.
// API name: repository
func (r *PutLifecycle) Repository(repository string) *PutLifecycle {

	r.req.Repository = &repository

	return r
}

// Retention Retention rules used to retain and delete snapshots created by the policy.
// API name: retention
func (r *PutLifecycle) Retention(retention *types.Retention) *PutLifecycle {

	r.req.Retention = retention

	return r
}

// Schedule Periodic or absolute schedule at which the policy creates snapshots. SLM
// applies schedule changes immediately.
// API name: schedule
func (r *PutLifecycle) Schedule(cronexpression string) *PutLifecycle {
	r.req.Schedule = &cronexpression

	return r
}
