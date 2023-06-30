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

// Creates a new model alias (or reassigns an existing one) to refer to the
// trained model
package puttrainedmodelalias

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
	modelaliasMask = iota + 1

	modelidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutTrainedModelAlias struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	modelalias string
	modelid    string
}

// NewPutTrainedModelAlias type alias for index.
type NewPutTrainedModelAlias func(modelid, modelalias string) *PutTrainedModelAlias

// NewPutTrainedModelAliasFunc returns a new instance of PutTrainedModelAlias with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutTrainedModelAliasFunc(tp elastictransport.Interface) NewPutTrainedModelAlias {
	return func(modelid, modelalias string) *PutTrainedModelAlias {
		n := New(tp)

		n.ModelAlias(modelalias)

		n.ModelId(modelid)

		return n
	}
}

// Creates a new model alias (or reassigns an existing one) to refer to the
// trained model
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models-aliases.html
func New(tp elastictransport.Interface) *PutTrainedModelAlias {
	r := &PutTrainedModelAlias{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutTrainedModelAlias) HttpRequest(ctx context.Context) (*http.Request, error) {
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

		method = http.MethodPut
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
func (r PutTrainedModelAlias) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutTrainedModelAlias query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a puttrainedmodelalias.Response
func (r PutTrainedModelAlias) Do(ctx context.Context) (*Response, error) {

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
func (r PutTrainedModelAlias) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the PutTrainedModelAlias headers map.
func (r *PutTrainedModelAlias) Header(key, value string) *PutTrainedModelAlias {
	r.headers.Set(key, value)

	return r
}

// ModelAlias The alias to create or update. This value cannot end in numbers.
// API Name: modelalias
func (r *PutTrainedModelAlias) ModelAlias(v string) *PutTrainedModelAlias {
	r.paramSet |= modelaliasMask
	r.modelalias = v

	return r
}

// ModelId The identifier for the trained model that the alias refers to.
// API Name: modelid
func (r *PutTrainedModelAlias) ModelId(v string) *PutTrainedModelAlias {
	r.paramSet |= modelidMask
	r.modelid = v

	return r
}

// Reassign Specifies whether the alias gets reassigned to the specified trained
// model if it is already assigned to a different model. If the alias is
// already assigned and this parameter is false, the API returns an error.
// API name: reassign
func (r *PutTrainedModelAlias) Reassign(b bool) *PutTrainedModelAlias {
	r.values.Set("reassign", strconv.FormatBool(b))

	return r
}
