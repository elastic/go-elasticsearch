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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Retrieves anomaly records for an anomaly detection job.
package getrecords

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
	jobidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetRecords struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int

	jobid string
}

// NewGetRecords type alias for index.
type NewGetRecords func(jobid string) *GetRecords

// NewGetRecordsFunc returns a new instance of GetRecords with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetRecordsFunc(tp elastictransport.Interface) NewGetRecords {
	return func(jobid string) *GetRecords {
		n := New(tp)

		n.JobId(jobid)

		return n
	}
}

// Retrieves anomaly records for an anomaly detection job.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-record.html
func New(tp elastictransport.Interface) *GetRecords {
	r := &GetRecords{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *GetRecords) Raw(raw io.Reader) *GetRecords {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *GetRecords) Request(req *Request) *GetRecords {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetRecords) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for GetRecords: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == jobidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")

		path.WriteString(r.jobid)
		path.WriteString("/")
		path.WriteString("results")
		path.WriteString("/")
		path.WriteString("records")

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
func (r GetRecords) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetRecords query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getrecords.Response
func (r GetRecords) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the GetRecords headers map.
func (r *GetRecords) Header(key, value string) *GetRecords {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job.
// API Name: jobid
func (r *GetRecords) JobId(v string) *GetRecords {
	r.paramSet |= jobidMask
	r.jobid = v

	return r
}

// Desc If true, the results are sorted in descending order.
// API name: desc
func (r *GetRecords) Desc(b bool) *GetRecords {
	r.values.Set("desc", strconv.FormatBool(b))

	return r
}

// End Returns records with timestamps earlier than this time. The default value
// means results are not limited to specific timestamps.
// API name: end
func (r *GetRecords) End(v string) *GetRecords {
	r.values.Set("end", v)

	return r
}

// ExcludeInterim If `true`, the output excludes interim results.
// API name: exclude_interim
func (r *GetRecords) ExcludeInterim(b bool) *GetRecords {
	r.values.Set("exclude_interim", strconv.FormatBool(b))

	return r
}

// From Skips the specified number of records.
// API name: from
func (r *GetRecords) From(i int) *GetRecords {
	r.values.Set("from", strconv.Itoa(i))

	return r
}

// RecordScore Returns records with anomaly scores greater or equal than this value.
// API name: record_score
func (r *GetRecords) RecordScore(v string) *GetRecords {
	r.values.Set("record_score", v)

	return r
}

// Size Specifies the maximum number of records to obtain.
// API name: size
func (r *GetRecords) Size(i int) *GetRecords {
	r.values.Set("size", strconv.Itoa(i))

	return r
}

// Sort Specifies the sort field for the requested records.
// API name: sort
func (r *GetRecords) Sort(v string) *GetRecords {
	r.values.Set("sort", v)

	return r
}

// Start Returns records with timestamps after this time. The default value means
// results are not limited to specific timestamps.
// API name: start
func (r *GetRecords) Start(v string) *GetRecords {
	r.values.Set("start", v)

	return r
}
