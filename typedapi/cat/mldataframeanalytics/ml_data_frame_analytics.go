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

// Gets configuration and usage information about data frame analytics jobs.
package mldataframeanalytics

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/bytes"
)

const (
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MlDataFrameAnalytics struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	id string
}

// NewMlDataFrameAnalytics type alias for index.
type NewMlDataFrameAnalytics func() *MlDataFrameAnalytics

// NewMlDataFrameAnalyticsFunc returns a new instance of MlDataFrameAnalytics with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMlDataFrameAnalyticsFunc(tp elastictransport.Interface) NewMlDataFrameAnalytics {
	return func() *MlDataFrameAnalytics {
		n := New(tp)

		return n
	}
}

// Gets configuration and usage information about data frame analytics jobs.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/cat-dfanalytics.html
func New(tp elastictransport.Interface) *MlDataFrameAnalytics {
	r := &MlDataFrameAnalytics{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *MlDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("ml")
		path.WriteString("/")
		path.WriteString("data_frame")
		path.WriteString("/")
		path.WriteString("analytics")

		method = http.MethodGet
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("ml")
		path.WriteString("/")
		path.WriteString("data_frame")
		path.WriteString("/")
		path.WriteString("analytics")
		path.WriteString("/")

		path.WriteString(r.id)

		method = http.MethodGet
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
func (r MlDataFrameAnalytics) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the MlDataFrameAnalytics query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a mldataframeanalytics.Response
func (r MlDataFrameAnalytics) Do(ctx context.Context) (Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(&response)
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
func (r MlDataFrameAnalytics) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the MlDataFrameAnalytics headers map.
func (r *MlDataFrameAnalytics) Header(key, value string) *MlDataFrameAnalytics {
	r.headers.Set(key, value)

	return r
}

// Id The ID of the data frame analytics to fetch
// API Name: id
func (r *MlDataFrameAnalytics) Id(v string) *MlDataFrameAnalytics {
	r.paramSet |= idMask
	r.id = v

	return r
}

// AllowNoMatch Whether to ignore if a wildcard expression matches no configs. (This includes
// `_all` string or when no configs have been specified)
// API name: allow_no_match
func (r *MlDataFrameAnalytics) AllowNoMatch(b bool) *MlDataFrameAnalytics {
	r.values.Set("allow_no_match", strconv.FormatBool(b))

	return r
}

// Bytes The unit in which to display byte values
// API name: bytes
func (r *MlDataFrameAnalytics) Bytes(enum bytes.Bytes) *MlDataFrameAnalytics {
	r.values.Set("bytes", enum.String())

	return r
}

// H Comma-separated list of column names to display.
// API name: h
func (r *MlDataFrameAnalytics) H(v string) *MlDataFrameAnalytics {
	r.values.Set("h", v)

	return r
}

// S Comma-separated list of column names or column aliases used to sort the
// response.
// API name: s
func (r *MlDataFrameAnalytics) S(v string) *MlDataFrameAnalytics {
	r.values.Set("s", v)

	return r
}

// Time Unit used to display time values.
// API name: time
func (r *MlDataFrameAnalytics) Time(v string) *MlDataFrameAnalytics {
	r.values.Set("time", v)

	return r
}
