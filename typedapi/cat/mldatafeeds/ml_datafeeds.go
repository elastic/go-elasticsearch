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


// Gets configuration and usage information about datafeeds.
package mldatafeeds

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeunit"
)

const (
	datafeedidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MlDatafeeds struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	datafeedid string
}

// NewMlDatafeeds type alias for index.
type NewMlDatafeeds func() *MlDatafeeds

// NewMlDatafeedsFunc returns a new instance of MlDatafeeds with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMlDatafeedsFunc(tp elastictransport.Interface) NewMlDatafeeds {
	return func() *MlDatafeeds {
		n := New(tp)

		return n
	}
}

// Gets configuration and usage information about datafeeds.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/cat-datafeeds.html
func New(tp elastictransport.Interface) *MlDatafeeds {
	r := &MlDatafeeds{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *MlDatafeeds) HttpRequest(ctx context.Context) (*http.Request, error) {
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
		path.WriteString("datafeeds")

		method = http.MethodGet
	case r.paramSet == datafeedidMask:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("ml")
		path.WriteString("/")
		path.WriteString("datafeeds")
		path.WriteString("/")

		path.WriteString(r.datafeedid)

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

// Do runs the http.Request through the provided transport.
func (r MlDatafeeds) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the MlDatafeeds query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r MlDatafeeds) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Do(ctx)

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

// Header set a key, value pair in the MlDatafeeds headers map.
func (r *MlDatafeeds) Header(key, value string) *MlDatafeeds {
	r.headers.Set(key, value)

	return r
}

// DatafeedId A numerical character string that uniquely identifies the datafeed.
// API Name: datafeedid
func (r *MlDatafeeds) DatafeedId(v string) *MlDatafeeds {
	r.paramSet |= datafeedidMask
	r.datafeedid = v

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
// * Contains wildcard expressions and there are no datafeeds that match.
// * Contains the `_all` string or no identifiers and there are no matches.
// * Contains wildcard expressions and there are only partial matches.
//
// If `true`, the API returns an empty datafeeds array when there are no matches
// and the subset of results when
// there are partial matches. If `false`, the API returns a 404 status code when
// there are no matches or only
// partial matches.
// API name: allow_no_match
func (r *MlDatafeeds) AllowNoMatch(b bool) *MlDatafeeds {
	r.values.Set("allow_no_match", strconv.FormatBool(b))

	return r
}

// H Comma-separated list of column names to display.
// API name: h
func (r *MlDatafeeds) H(value string) *MlDatafeeds {
	r.values.Set("h", value)

	return r
}

// S Comma-separated list of column names or column aliases used to sort the
// response.
// API name: s
func (r *MlDatafeeds) S(value string) *MlDatafeeds {
	r.values.Set("s", value)

	return r
}

// Time The unit used to display time values.
// API name: time
func (r *MlDatafeeds) Time(enum timeunit.TimeUnit) *MlDatafeeds {
	r.values.Set("time", enum.String())

	return r
}
