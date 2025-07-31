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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Create a trained model vocabulary.
// This API is supported only for natural language processing (NLP) models.
// The vocabulary is stored in the index as described in
// `inference_config.*.vocabulary` of the trained model definition.
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

type PutTrainedModelVocabulary struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	modelid string

	spanStarted bool

	instrument elastictransport.Instrumentation
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

// Create a trained model vocabulary.
// This API is supported only for natural language processing (NLP) models.
// The vocabulary is stored in the index as described in
// `inference_config.*.vocabulary` of the trained model definition.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-vocabulary.html
func New(tp elastictransport.Interface) *PutTrainedModelVocabulary {
	r := &PutTrainedModelVocabulary{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutTrainedModelVocabulary: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == modelidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "modelid", r.modelid)
		}
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
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
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
func (r PutTrainedModelVocabulary) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.put_trained_model_vocabulary")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "ml.put_trained_model_vocabulary")
		if reader := instrument.RecordRequestBody(ctx, "ml.put_trained_model_vocabulary", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.put_trained_model_vocabulary")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutTrainedModelVocabulary query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a puttrainedmodelvocabulary.Response
func (r PutTrainedModelVocabulary) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.put_trained_model_vocabulary")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
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

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutTrainedModelVocabulary) ErrorTrace(errortrace bool) *PutTrainedModelVocabulary {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutTrainedModelVocabulary) FilterPath(filterpaths ...string) *PutTrainedModelVocabulary {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *PutTrainedModelVocabulary) Human(human bool) *PutTrainedModelVocabulary {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutTrainedModelVocabulary) Pretty(pretty bool) *PutTrainedModelVocabulary {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Merges The optional model merges if required by the tokenizer.
// API name: merges
func (r *PutTrainedModelVocabulary) Merges(merges ...string) *PutTrainedModelVocabulary {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Merges = merges

	return r
}

// Scores The optional vocabulary value scores if required by the tokenizer.
// API name: scores
func (r *PutTrainedModelVocabulary) Scores(scores ...types.Float64) *PutTrainedModelVocabulary {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Scores = scores

	return r
}

// Vocabulary The model vocabulary, which must not be empty.
// API name: vocabulary
func (r *PutTrainedModelVocabulary) Vocabulary(vocabularies ...string) *PutTrainedModelVocabulary {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Vocabulary = vocabularies

	return r
}
