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

// Retrieves a synonym rule from a synonym set
package getsynonymrule

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
	setidMask = iota + 1

	ruleidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetSynonymRule struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	setid  string
	ruleid string
}

// NewGetSynonymRule type alias for index.
type NewGetSynonymRule func(setid, ruleid string) *GetSynonymRule

// NewGetSynonymRuleFunc returns a new instance of GetSynonymRule with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetSynonymRuleFunc(tp elastictransport.Interface) NewGetSynonymRule {
	return func(setid, ruleid string) *GetSynonymRule {
		n := New(tp)

		n._setid(setid)

		n._ruleid(ruleid)

		return n
	}
}

// Retrieves a synonym rule from a synonym set
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/get-synonym-rule.html
func New(tp elastictransport.Interface) *GetSynonymRule {
	r := &GetSynonymRule{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetSynonymRule) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == setidMask|ruleidMask:
		path.WriteString("/")
		path.WriteString("_synonyms")
		path.WriteString("/")

		path.WriteString(r.setid)
		path.WriteString("/")

		path.WriteString(r.ruleid)

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
func (r GetSynonymRule) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetSynonymRule query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getsynonymrule.Response
func (r GetSynonymRule) Do(ctx context.Context) (*Response, error) {

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
func (r GetSynonymRule) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GetSynonymRule headers map.
func (r *GetSynonymRule) Header(key, value string) *GetSynonymRule {
	r.headers.Set(key, value)

	return r
}

// SetId The id of the synonym set to retrieve the synonym rule from
// API Name: setid
func (r *GetSynonymRule) _setid(setid string) *GetSynonymRule {
	r.paramSet |= setidMask
	r.setid = setid

	return r
}

// RuleId The id of the synonym rule to retrieve
// API Name: ruleid
func (r *GetSynonymRule) _ruleid(ruleid string) *GetSynonymRule {
	r.paramSet |= ruleidMask
	r.ruleid = ruleid

	return r
}
