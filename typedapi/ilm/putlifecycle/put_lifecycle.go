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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Creates a lifecycle policy
package putlifecycle

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	policyMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutLifecycle struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	policy string
}

// NewPutLifecycle type alias for index.
type NewPutLifecycle func(policy string) *PutLifecycle

// NewPutLifecycleFunc returns a new instance of PutLifecycle with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutLifecycleFunc(tp elastictransport.Interface) NewPutLifecycle {
	return func(policy string) *PutLifecycle {
		n := New(tp)

		n.Policy(policy)

		return n
	}
}

// Creates a lifecycle policy
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-put-lifecycle.html
func New(tp elastictransport.Interface) *PutLifecycle {
	r := &PutLifecycle{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutLifecycle) Raw(raw json.RawMessage) *PutLifecycle {
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

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutLifecycle: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == policyMask:
		path.WriteString("/")
		path.WriteString("_ilm")
		path.WriteString("/")
		path.WriteString("policy")
		path.WriteString("/")

		path.WriteString(r.policy)

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

// Do runs the http.Request through the provided transport.
func (r PutLifecycle) Do(ctx context.Context) (*http.Response, error) {
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

// Header set a key, value pair in the PutLifecycle headers map.
func (r *PutLifecycle) Header(key, value string) *PutLifecycle {
	r.headers.Set(key, value)

	return r
}

// Policy Identifier for the policy.
// API Name: policy
func (r *PutLifecycle) Policy(v string) *PutLifecycle {
	r.paramSet |= policyMask
	r.policy = v

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *PutLifecycle) MasterTimeout(value string) *PutLifecycle {
	r.values.Set("master_timeout", value)

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and returns an error.
// API name: timeout
func (r *PutLifecycle) Timeout(value string) *PutLifecycle {
	r.values.Set("timeout", value)

	return r
}
