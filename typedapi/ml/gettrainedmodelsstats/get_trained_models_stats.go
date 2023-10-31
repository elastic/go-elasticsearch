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

// Retrieves usage information for trained inference models.
package gettrainedmodelsstats

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
)

const (
	modelidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetTrainedModelsStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	modelid string
}

// NewGetTrainedModelsStats type alias for index.
type NewGetTrainedModelsStats func() *GetTrainedModelsStats

// NewGetTrainedModelsStatsFunc returns a new instance of GetTrainedModelsStats with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetTrainedModelsStatsFunc(tp elastictransport.Interface) NewGetTrainedModelsStats {
	return func() *GetTrainedModelsStats {
		n := New(tp)

		return n
	}
}

// Retrieves usage information for trained inference models.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trained-models-stats.html
func New(tp elastictransport.Interface) *GetTrainedModelsStats {
	r := &GetTrainedModelsStats{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetTrainedModelsStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == modelidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")

		path.WriteString(r.modelid)
		path.WriteString("/")
		path.WriteString("_stats")

		method = http.MethodGet
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")
		path.WriteString("_stats")

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
func (r GetTrainedModelsStats) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetTrainedModelsStats query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a gettrainedmodelsstats.Response
func (r GetTrainedModelsStats) Do(ctx context.Context) (*Response, error) {

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
func (r GetTrainedModelsStats) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GetTrainedModelsStats headers map.
func (r *GetTrainedModelsStats) Header(key, value string) *GetTrainedModelsStats {
	r.headers.Set(key, value)

	return r
}

// ModelId The unique identifier of the trained model or a model alias. It can be a
// comma-separated list or a wildcard expression.
// API Name: modelid
func (r *GetTrainedModelsStats) ModelId(modelid string) *GetTrainedModelsStats {
	r.paramSet |= modelidMask
	r.modelid = modelid

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
// - Contains wildcard expressions and there are no models that match.
// - Contains the _all string or no identifiers and there are no matches.
// - Contains wildcard expressions and there are only partial matches.
//
// If true, it returns an empty array when there are no matches and the
// subset of results when there are partial matches.
// API name: allow_no_match
func (r *GetTrainedModelsStats) AllowNoMatch(allownomatch bool) *GetTrainedModelsStats {
	r.values.Set("allow_no_match", strconv.FormatBool(allownomatch))

	return r
}

// From Skips the specified number of models.
// API name: from
func (r *GetTrainedModelsStats) From(from int) *GetTrainedModelsStats {
	r.values.Set("from", strconv.Itoa(from))

	return r
}

// Size Specifies the maximum number of models to obtain.
// API name: size
func (r *GetTrainedModelsStats) Size(size int) *GetTrainedModelsStats {
	r.values.Set("size", strconv.Itoa(size))

	return r
}
