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


// Creates a rollup job.
package putjob

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
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutJob struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	id string
}

// NewPutJob type alias for index.
type NewPutJob func(id string) *PutJob

// NewPutJobFunc returns a new instance of PutJob with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutJobFunc(tp elastictransport.Interface) NewPutJob {
	return func(id string) *PutJob {
		n := New(tp)

		n.Id(id)

		return n
	}
}

// Creates a rollup job.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/rollup-put-job.html
func New(tp elastictransport.Interface) *PutJob {
	r := &PutJob{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutJob) Raw(raw json.RawMessage) *PutJob {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutJob) Request(req *Request) *PutJob {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutJob: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_rollup")
		path.WriteString("/")
		path.WriteString("job")
		path.WriteString("/")

		path.WriteString(r.id)

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
func (r PutJob) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutJob query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the PutJob headers map.
func (r *PutJob) Header(key, value string) *PutJob {
	r.headers.Set(key, value)

	return r
}

// Id Identifier for the rollup job. This can be any alphanumeric string and
// uniquely identifies the
// data that is associated with the rollup job. The ID is persistent; it is
// stored with the rolled
// up data. If you create a job, let it run for a while, then delete the job,
// the data that the job
// rolled up is still be associated with this job ID. You cannot create a new
// job with the same ID
// since that could lead to problems with mismatched job configurations.
// API Name: id
func (r *PutJob) Id(v string) *PutJob {
	r.paramSet |= idMask
	r.id = v

	return r
}
