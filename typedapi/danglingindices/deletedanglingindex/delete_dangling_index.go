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


// Deletes the specified dangling index
package deletedanglingindex

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
)

const (
	indexuuidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteDanglingIndex struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	indexuuid string
}

// NewDeleteDanglingIndex type alias for index.
type NewDeleteDanglingIndex func(indexuuid string) *DeleteDanglingIndex

// NewDeleteDanglingIndexFunc returns a new instance of DeleteDanglingIndex with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteDanglingIndexFunc(tp elastictransport.Interface) NewDeleteDanglingIndex {
	return func(indexuuid string) *DeleteDanglingIndex {
		n := New(tp)

		n.IndexUuid(indexuuid)

		return n
	}
}

// Deletes the specified dangling index
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/modules-gateway-dangling-indices.html
func New(tp elastictransport.Interface) *DeleteDanglingIndex {
	r := &DeleteDanglingIndex{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *DeleteDanglingIndex) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexuuidMask:
		path.WriteString("/")
		path.WriteString("_dangling")
		path.WriteString("/")

		path.WriteString(r.indexuuid)

		method = http.MethodDelete
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
func (r DeleteDanglingIndex) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the DeleteDanglingIndex query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r DeleteDanglingIndex) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the DeleteDanglingIndex headers map.
func (r *DeleteDanglingIndex) Header(key, value string) *DeleteDanglingIndex {
	r.headers.Set(key, value)

	return r
}

// IndexUuid The UUID of the dangling index
// API Name: indexuuid
func (r *DeleteDanglingIndex) IndexUuid(v string) *DeleteDanglingIndex {
	r.paramSet |= indexuuidMask
	r.indexuuid = v

	return r
}

// AcceptDataLoss Must be set to true in order to delete the dangling index
// API name: accept_data_loss
func (r *DeleteDanglingIndex) AcceptDataLoss(b bool) *DeleteDanglingIndex {
	r.values.Set("accept_data_loss", strconv.FormatBool(b))

	return r
}

// MasterTimeout Specify timeout for connection to master
// API name: master_timeout
func (r *DeleteDanglingIndex) MasterTimeout(value string) *DeleteDanglingIndex {
	r.values.Set("master_timeout", value)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *DeleteDanglingIndex) Timeout(value string) *DeleteDanglingIndex {
	r.values.Set("timeout", value)

	return r
}
