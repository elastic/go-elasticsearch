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
//
// Code generated from specification version 8.19.0: DO NOT EDIT

package esapi

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func newMLPutTrainedModelDefinitionPartFunc(t Transport) MLPutTrainedModelDefinitionPart {
	return func(body io.Reader, model_id string, part *int, o ...func(*MLPutTrainedModelDefinitionPartRequest)) (*Response, error) {
		var r = MLPutTrainedModelDefinitionPartRequest{Body: body, ModelID: model_id, Part: part}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.Instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// MLPutTrainedModelDefinitionPart - Creates part of a trained model definition
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-definition-part.html.
type MLPutTrainedModelDefinitionPart func(body io.Reader, model_id string, part *int, o ...func(*MLPutTrainedModelDefinitionPartRequest)) (*Response, error)

// MLPutTrainedModelDefinitionPartRequest configures the ML Put Trained Model Definition Part API request.
type MLPutTrainedModelDefinitionPartRequest struct {
	Body io.Reader

	ModelID string
	Part    *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r MLPutTrainedModelDefinitionPartRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.put_trained_model_definition_part")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	if r.Part == nil {
		return nil, errors.New("part is required and cannot be nil")
	}

	path.Grow(7 + 1 + len("_ml") + 1 + len("trained_models") + 1 + len(r.ModelID) + 1 + len("definition") + 1 + len(strconv.Itoa(*r.Part)))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("trained_models")
	path.WriteString("/")
	path.WriteString(r.ModelID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "model_id", r.ModelID)
	}
	path.WriteString("/")
	path.WriteString("definition")
	path.WriteString("/")
	path.WriteString(strconv.Itoa(*r.Part))
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "part", strconv.Itoa(*r.Part))
	}

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "ml.put_trained_model_definition_part")
		if reader := instrument.RecordRequestBody(ctx, "ml.put_trained_model_definition_part", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.put_trained_model_definition_part")
	}
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
func (f MLPutTrainedModelDefinitionPart) WithContext(v context.Context) func(*MLPutTrainedModelDefinitionPartRequest) {
	return func(r *MLPutTrainedModelDefinitionPartRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLPutTrainedModelDefinitionPart) WithPretty() func(*MLPutTrainedModelDefinitionPartRequest) {
	return func(r *MLPutTrainedModelDefinitionPartRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLPutTrainedModelDefinitionPart) WithHuman() func(*MLPutTrainedModelDefinitionPartRequest) {
	return func(r *MLPutTrainedModelDefinitionPartRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLPutTrainedModelDefinitionPart) WithErrorTrace() func(*MLPutTrainedModelDefinitionPartRequest) {
	return func(r *MLPutTrainedModelDefinitionPartRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLPutTrainedModelDefinitionPart) WithFilterPath(v ...string) func(*MLPutTrainedModelDefinitionPartRequest) {
	return func(r *MLPutTrainedModelDefinitionPartRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLPutTrainedModelDefinitionPart) WithHeader(h map[string]string) func(*MLPutTrainedModelDefinitionPartRequest) {
	return func(r *MLPutTrainedModelDefinitionPartRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLPutTrainedModelDefinitionPart) WithOpaqueID(s string) func(*MLPutTrainedModelDefinitionPartRequest) {
	return func(r *MLPutTrainedModelDefinitionPartRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
