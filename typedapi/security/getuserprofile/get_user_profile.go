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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Retrieves user profiles for the given unique ID(s).
package getuserprofile

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
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	uidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetUserProfile struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	uid string
}

// NewGetUserProfile type alias for index.
type NewGetUserProfile func(uid string) *GetUserProfile

// NewGetUserProfileFunc returns a new instance of GetUserProfile with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetUserProfileFunc(tp elastictransport.Interface) NewGetUserProfile {
	return func(uid string) *GetUserProfile {
		n := New(tp)

		n._uid(uid)

		return n
	}
}

// Retrieves user profiles for the given unique ID(s).
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user-profile.html
func New(tp elastictransport.Interface) *GetUserProfile {
	r := &GetUserProfile{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetUserProfile) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == uidMask:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("profile")
		path.WriteString("/")

		path.WriteString(r.uid)

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
func (r GetUserProfile) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetUserProfile query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getuserprofile.Response
func (r GetUserProfile) Do(ctx context.Context) (*Response, error) {

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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetUserProfile) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GetUserProfile headers map.
func (r *GetUserProfile) Header(key, value string) *GetUserProfile {
	r.headers.Set(key, value)

	return r
}

// Uid A unique identifier for the user profile.
// API Name: uid
func (r *GetUserProfile) _uid(uids ...string) *GetUserProfile {
	r.paramSet |= uidMask
	r.uid = strings.Join(uids, ",")

	return r
}

// Data List of filters for the `data` field of the profile document.
// To return all content use `data=*`. To return a subset of content
// use `data=<key>` to retrieve content nested under the specified `<key>`.
// By default returns no `data` content.
// API name: data
func (r *GetUserProfile) Data(data ...string) *GetUserProfile {
	tmp := []string{}
	for _, item := range data {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("data", strings.Join(tmp, ","))

	return r
}
