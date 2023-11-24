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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Creates a trained model vocabulary
package puttrainedmodelvocabulary

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	modelidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutTrainedModelVocabulary struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	modelid string
}

// NewPutTrainedModelVocabulary type alias for index.
type NewPutTrainedModelVocabulary func(modelid string) *PutTrainedModelVocabulary

// NewPutTrainedModelVocabularyFunc returns a new instance of PutTrainedModelVocabulary with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutTrainedModelVocabularyFunc(tp elastictransport.Interface) NewPutTrainedModelVocabulary {
	return func(modelid string) *PutTrainedModelVocabulary {
		n := New(tp)

		n._modelid(modelid)

		return n
	}
}

// Creates a trained model vocabulary
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-vocabulary.html
func New(tp elastictransport.Interface) *PutTrainedModelVocabulary {
	r := &PutTrainedModelVocabulary{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutTrainedModelVocabulary) Raw(raw io.Reader) *PutTrainedModelVocabulary {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutTrainedModelVocabulary) Request(req *Request) *PutTrainedModelVocabulary {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutTrainedModelVocabulary) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutTrainedModelVocabulary: %w", err)
		}

		r.buf.Write(data)

	}

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
		path.WriteString("vocabulary")

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
func (r PutTrainedModelVocabulary) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutTrainedModelVocabulary query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a puttrainedmodelvocabulary.Response
func (r PutTrainedModelVocabulary) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutTrainedModelVocabulary headers map.
func (r *PutTrainedModelVocabulary) Header(key, value string) *PutTrainedModelVocabulary {
	r.headers.Set(key, value)

	return r
}

// ModelId The unique identifier of the trained model.
// API Name: modelid
func (r *PutTrainedModelVocabulary) _modelid(modelid string) *PutTrainedModelVocabulary {
	r.paramSet |= modelidMask
	r.modelid = modelid

	return r
}

// Merges The optional model merges if required by the tokenizer.
// API name: merges
func (r *PutTrainedModelVocabulary) Merges(merges ...string) *PutTrainedModelVocabulary {
	r.req.Merges = merges

	return r
}

// Scores The optional vocabulary value scores if required by the tokenizer.
// API name: scores
func (r *PutTrainedModelVocabulary) Scores(scores ...types.Float64) *PutTrainedModelVocabulary {
	r.req.Scores = scores

	return r
}

// Vocabulary The model vocabulary, which must not be empty.
// API name: vocabulary
func (r *PutTrainedModelVocabulary) Vocabulary(vocabularies ...string) *PutTrainedModelVocabulary {
	r.req.Vocabulary = vocabularies

	return r
}
