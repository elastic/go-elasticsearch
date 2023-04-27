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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

// Retrieves configuration information for a trained inference model.
package gettrainedmodels

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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/include"
)

const (
	modelidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetTrainedModels struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	modelid string
}

// NewGetTrainedModels type alias for index.
type NewGetTrainedModels func() *GetTrainedModels

// NewGetTrainedModelsFunc returns a new instance of GetTrainedModels with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetTrainedModelsFunc(tp elastictransport.Interface) NewGetTrainedModels {
	return func() *GetTrainedModels {
		n := New(tp)

		return n
	}
}

// Retrieves configuration information for a trained inference model.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trained-models.html
func New(tp elastictransport.Interface) *GetTrainedModels {
	r := &GetTrainedModels{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetTrainedModels) HttpRequest(ctx context.Context) (*http.Request, error) {
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

		method = http.MethodGet
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")

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
func (r GetTrainedModels) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetTrainedModels query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a gettrainedmodels.Response
func (r GetTrainedModels) Do(ctx context.Context) (*Response, error) {

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

	return nil, errorResponse
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetTrainedModels) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GetTrainedModels headers map.
func (r *GetTrainedModels) Header(key, value string) *GetTrainedModels {
	r.headers.Set(key, value)

	return r
}

// ModelId The unique identifier of the trained model.
// API Name: modelid
func (r *GetTrainedModels) ModelId(v string) *GetTrainedModels {
	r.paramSet |= modelidMask
	r.modelid = v

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
func (r *GetTrainedModels) AllowNoMatch(b bool) *GetTrainedModels {
	r.values.Set("allow_no_match", strconv.FormatBool(b))

	return r
}

// DecompressDefinition Specifies whether the included model definition should be returned as a
// JSON map (true) or in a custom compressed format (false).
// API name: decompress_definition
func (r *GetTrainedModels) DecompressDefinition(b bool) *GetTrainedModels {
	r.values.Set("decompress_definition", strconv.FormatBool(b))

	return r
}

// ExcludeGenerated Indicates if certain fields should be removed from the configuration on
// retrieval. This allows the configuration to be in an acceptable format to
// be retrieved and then added to another cluster.
// API name: exclude_generated
func (r *GetTrainedModels) ExcludeGenerated(b bool) *GetTrainedModels {
	r.values.Set("exclude_generated", strconv.FormatBool(b))

	return r
}

// From Skips the specified number of models.
// API name: from
func (r *GetTrainedModels) From(i int) *GetTrainedModels {
	r.values.Set("from", strconv.Itoa(i))

	return r
}

// Include A comma delimited string of optional fields to include in the response
// body.
// API name: include
func (r *GetTrainedModels) Include(enum include.Include) *GetTrainedModels {
	r.values.Set("include", enum.String())

	return r
}

// Size Specifies the maximum number of models to obtain.
// API name: size
func (r *GetTrainedModels) Size(i int) *GetTrainedModels {
	r.values.Set("size", strconv.Itoa(i))

	return r
}

// Tags A comma delimited string of tags. A trained model can have many tags, or
// none. When supplied, only trained models that contain all the supplied
// tags are returned.
// API name: tags
func (r *GetTrainedModels) Tags(v string) *GetTrainedModels {
	r.values.Set("tags", v)

	return r
}
