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

// Create a trained model.
// Enable you to supply a trained model that is not created by data frame
// analytics.
package puttrainedmodel

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/trainedmodeltype"
)

const (
	modelidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutTrainedModel struct {
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

// NewPutTrainedModel type alias for index.
type NewPutTrainedModel func(modelid string) *PutTrainedModel

// NewPutTrainedModelFunc returns a new instance of PutTrainedModel with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutTrainedModelFunc(tp elastictransport.Interface) NewPutTrainedModel {
	return func(modelid string) *PutTrainedModel {
		n := New(tp)

		n._modelid(modelid)

		return n
	}
}

// Create a trained model.
// Enable you to supply a trained model that is not created by data frame
// analytics.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models.html
func New(tp elastictransport.Interface) *PutTrainedModel {
	r := &PutTrainedModel{
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
func (r *PutTrainedModel) Raw(raw io.Reader) *PutTrainedModel {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutTrainedModel) Request(req *Request) *PutTrainedModel {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutTrainedModel) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutTrainedModel: %w", err)
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
func (r PutTrainedModel) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.put_trained_model")
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
		instrument.BeforeRequest(req, "ml.put_trained_model")
		if reader := instrument.RecordRequestBody(ctx, "ml.put_trained_model", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.put_trained_model")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutTrainedModel query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a puttrainedmodel.Response
func (r PutTrainedModel) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.put_trained_model")
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

// Header set a key, value pair in the PutTrainedModel headers map.
func (r *PutTrainedModel) Header(key, value string) *PutTrainedModel {
	r.headers.Set(key, value)

	return r
}

// ModelId The unique identifier of the trained model.
// API Name: modelid
func (r *PutTrainedModel) _modelid(modelid string) *PutTrainedModel {
	r.paramSet |= modelidMask
	r.modelid = modelid

	return r
}

// DeferDefinitionDecompression If set to `true` and a `compressed_definition` is provided,
// the request defers definition decompression and skips relevant
// validations.
// API name: defer_definition_decompression
func (r *PutTrainedModel) DeferDefinitionDecompression(deferdefinitiondecompression bool) *PutTrainedModel {
	r.values.Set("defer_definition_decompression", strconv.FormatBool(deferdefinitiondecompression))

	return r
}

// WaitForCompletion Whether to wait for all child operations (e.g. model download)
// to complete.
// API name: wait_for_completion
func (r *PutTrainedModel) WaitForCompletion(waitforcompletion bool) *PutTrainedModel {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutTrainedModel) ErrorTrace(errortrace bool) *PutTrainedModel {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutTrainedModel) FilterPath(filterpaths ...string) *PutTrainedModel {
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
func (r *PutTrainedModel) Human(human bool) *PutTrainedModel {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutTrainedModel) Pretty(pretty bool) *PutTrainedModel {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// CompressedDefinition The compressed (GZipped and Base64 encoded) inference definition of the
// model. If compressed_definition is specified, then definition cannot be
// specified.
// API name: compressed_definition
func (r *PutTrainedModel) CompressedDefinition(compresseddefinition string) *PutTrainedModel {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.CompressedDefinition = &compresseddefinition

	return r
}

// Definition The inference definition for the model. If definition is specified, then
// compressed_definition cannot be specified.
// API name: definition
func (r *PutTrainedModel) Definition(definition *types.Definition) *PutTrainedModel {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Definition = definition

	return r
}

// Description A human-readable description of the inference trained model.
// API name: description
func (r *PutTrainedModel) Description(description string) *PutTrainedModel {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Description = &description

	return r
}

// InferenceConfig The default configuration for inference. This can be either a regression
// or classification configuration. It must match the underlying
// definition.trained_model's target_type. For pre-packaged models such as
// ELSER the config is not required.
// API name: inference_config
func (r *PutTrainedModel) InferenceConfig(inferenceconfig *types.InferenceConfigCreateContainer) *PutTrainedModel {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.InferenceConfig = inferenceconfig

	return r
}

// Input The input field names for the model definition.
// API name: input
func (r *PutTrainedModel) Input(input *types.Input) *PutTrainedModel {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Input = input

	return r
}

// Metadata An object map that contains metadata about the model.
// API name: metadata
//
// metadata should be a json.RawMessage or a structure
// if a structure is provided, the client will defer a json serialization
// prior to sending the payload to Elasticsearch.
func (r *PutTrainedModel) Metadata(metadata any) *PutTrainedModel {
	if r.req == nil {
		r.req = NewRequest()
	}
	switch casted := metadata.(type) {
	case json.RawMessage:
		r.req.Metadata = casted
	default:
		r.deferred = append(r.deferred, func(request *Request) error {
			data, err := json.Marshal(metadata)
			if err != nil {
				return err
			}
			r.req.Metadata = data
			return nil
		})
	}

	return r
}

// ModelSizeBytes The estimated memory usage in bytes to keep the trained model in memory.
// This property is supported only if defer_definition_decompression is true
// or the model definition is not supplied.
// API name: model_size_bytes
func (r *PutTrainedModel) ModelSizeBytes(modelsizebytes int64) *PutTrainedModel {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ModelSizeBytes = &modelsizebytes

	return r
}

// ModelType The model type.
// API name: model_type
func (r *PutTrainedModel) ModelType(modeltype trainedmodeltype.TrainedModelType) *PutTrainedModel {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ModelType = &modeltype

	return r
}

// PlatformArchitecture The platform architecture (if applicable) of the trained mode. If the model
// only works on one platform, because it is heavily optimized for a particular
// processor architecture and OS combination, then this field specifies which.
// The format of the string must match the platform identifiers used by
// Elasticsearch,
// so one of, `linux-x86_64`, `linux-aarch64`, `darwin-x86_64`,
// `darwin-aarch64`,
// or `windows-x86_64`. For portable models (those that work independent of
// processor
// architecture or OS features), leave this field unset.
// API name: platform_architecture
func (r *PutTrainedModel) PlatformArchitecture(platformarchitecture string) *PutTrainedModel {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.PlatformArchitecture = &platformarchitecture

	return r
}

// PrefixStrings Optional prefix strings applied at inference
// API name: prefix_strings
func (r *PutTrainedModel) PrefixStrings(prefixstrings *types.TrainedModelPrefixStrings) *PutTrainedModel {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.PrefixStrings = prefixstrings

	return r
}

// Tags An array of tags to organize the model.
// API name: tags
func (r *PutTrainedModel) Tags(tags ...string) *PutTrainedModel {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Tags = tags

	return r
}
