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

// Deletes a model alias that refers to the trained model
package deletetrainedmodelalias

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
	modelaliasMask = iota + 1

	modelidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteTrainedModelAlias struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	modelalias string
	modelid    string
}

// NewDeleteTrainedModelAlias type alias for index.
type NewDeleteTrainedModelAlias func(modelid, modelalias string) *DeleteTrainedModelAlias

// NewDeleteTrainedModelAliasFunc returns a new instance of DeleteTrainedModelAlias with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteTrainedModelAliasFunc(tp elastictransport.Interface) NewDeleteTrainedModelAlias {
	return func(modelid, modelalias string) *DeleteTrainedModelAlias {
		n := New(tp)

		n._modelalias(modelalias)

		n._modelid(modelid)

		return n
	}
}

// Deletes a model alias that refers to the trained model
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-trained-models-aliases.html
func New(tp elastictransport.Interface) *DeleteTrainedModelAlias {
	r := &DeleteTrainedModelAlias{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *DeleteTrainedModelAlias) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == modelidMask|modelaliasMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")

		path.WriteString(r.modelid)
		path.WriteString("/")
		path.WriteString("model_aliases")
		path.WriteString("/")

		path.WriteString(r.modelalias)

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
func (r DeleteTrainedModelAlias) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the DeleteTrainedModelAlias query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a deletetrainedmodelalias.Response
func (r DeleteTrainedModelAlias) Do(ctx context.Context) (*Response, error) {

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
func (r DeleteTrainedModelAlias) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the DeleteTrainedModelAlias headers map.
func (r *DeleteTrainedModelAlias) Header(key, value string) *DeleteTrainedModelAlias {
	r.headers.Set(key, value)

	return r
}

// ModelAlias The model alias to delete.
// API Name: modelalias
func (r *DeleteTrainedModelAlias) _modelalias(modelalias string) *DeleteTrainedModelAlias {
	r.paramSet |= modelaliasMask
	r.modelalias = modelalias

	return r
}

// ModelId The trained model ID to which the model alias refers.
// API Name: modelid
func (r *DeleteTrainedModelAlias) _modelid(modelid string) *DeleteTrainedModelAlias {
	r.paramSet |= modelidMask
	r.modelid = modelid

	return r
}
