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

// Retrieves anomaly detection job results for one or more categories.
package getcategories

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

	categoryidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetCategories struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int

	jobid      string
	categoryid string
}

// NewGetCategories type alias for index.
type NewGetCategories func(jobid string) *GetCategories

// NewGetCategoriesFunc returns a new instance of GetCategories with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetCategoriesFunc(tp elastictransport.Interface) NewGetCategories {
	return func(jobid string) *GetCategories {
		n := New(tp)

		n.JobId(jobid)

		return n
	}
}

// Retrieves anomaly detection job results for one or more categories.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-category.html
func New(tp elastictransport.Interface) *GetCategories {
	r := &GetCategories{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *GetCategories) Raw(raw io.Reader) *GetCategories {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *GetCategories) Request(req *Request) *GetCategories {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetCategories) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for GetCategories: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == jobidMask|categoryidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")

		path.WriteString(r.jobid)
		path.WriteString("/")
		path.WriteString("results")
		path.WriteString("/")
		path.WriteString("categories")
		path.WriteString("/")

		path.WriteString(r.categoryid)

		method = http.MethodPost
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
		path.WriteString("categories")

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
func (r GetCategories) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetCategories query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getcategories.Response
func (r GetCategories) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the GetCategories headers map.
func (r *GetCategories) Header(key, value string) *GetCategories {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job.
// API Name: jobid
func (r *GetCategories) JobId(v string) *GetCategories {
	r.paramSet |= jobidMask
	r.jobid = v

	return r
}

// CategoryId Identifier for the category, which is unique in the job. If you specify
// neither the category ID nor the partition_field_value, the API returns
// information about all categories. If you specify only the
// partition_field_value, it returns information about all categories for
// the specified partition.
// API Name: categoryid
func (r *GetCategories) CategoryId(v string) *GetCategories {
	r.paramSet |= categoryidMask
	r.categoryid = v

	return r
}

// From Skips the specified number of categories.
// API name: from
func (r *GetCategories) From(i int) *GetCategories {
	r.values.Set("from", strconv.Itoa(i))

	return r
}

// PartitionFieldValue Only return categories for the specified partition.
// API name: partition_field_value
func (r *GetCategories) PartitionFieldValue(v string) *GetCategories {
	r.values.Set("partition_field_value", v)

	return r
}

// Size Specifies the maximum number of categories to obtain.
// API name: size
func (r *GetCategories) Size(i int) *GetCategories {
	r.values.Set("size", strconv.Itoa(i))

	return r
}
