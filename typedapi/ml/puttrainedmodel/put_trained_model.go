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

// Creates an inference trained model.
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	modelid string
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

// Creates an inference trained model.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models.html
func New(tp elastictransport.Interface) *PutTrainedModel {
	r := &PutTrainedModel{
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutTrainedModel: %w", err)
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
func (r PutTrainedModel) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutTrainedModel query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a puttrainedmodel.Response
func (r PutTrainedModel) Do(ctx context.Context) (*Response, error) {

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

// DeferDefinitionDecompression If set to `true` and a `compressed_definition` is provided, the request
// defers definition decompression and skips relevant validations.
// API name: defer_definition_decompression
func (r *PutTrainedModel) DeferDefinitionDecompression(deferdefinitiondecompression bool) *PutTrainedModel {
	r.values.Set("defer_definition_decompression", strconv.FormatBool(deferdefinitiondecompression))

	return r
}

// CompressedDefinition The compressed (GZipped and Base64 encoded) inference definition of the
// model. If compressed_definition is specified, then definition cannot be
// specified.
// API name: compressed_definition
func (r *PutTrainedModel) CompressedDefinition(compresseddefinition string) *PutTrainedModel {

	r.req.CompressedDefinition = &compresseddefinition

	return r
}

// Definition The inference definition for the model. If definition is specified, then
// compressed_definition cannot be specified.
// API name: definition
func (r *PutTrainedModel) Definition(definition *types.Definition) *PutTrainedModel {

	r.req.Definition = definition

	return r
}

// Description A human-readable description of the inference trained model.
// API name: description
func (r *PutTrainedModel) Description(description string) *PutTrainedModel {

	r.req.Description = &description

	return r
}

// InferenceConfig The default configuration for inference. This can be either a regression
// or classification configuration. It must match the underlying
// definition.trained_model's target_type. For pre-packaged models such as
// ELSER the config is not required.
// API name: inference_config
func (r *PutTrainedModel) InferenceConfig(inferenceconfig *types.InferenceConfigCreateContainer) *PutTrainedModel {

	r.req.InferenceConfig = inferenceconfig

	return r
}

// Input The input field names for the model definition.
// API name: input
func (r *PutTrainedModel) Input(input *types.Input) *PutTrainedModel {

	r.req.Input = input

	return r
}

// Metadata An object map that contains metadata about the model.
// API name: metadata
//
// metadata should be a json.RawMessage or a structure
// if a structure is provided, the client will defer a json serialization
// prior to sending the payload to Elasticsearch.
func (r *PutTrainedModel) Metadata(metadata interface{}) *PutTrainedModel {
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

	r.req.ModelSizeBytes = &modelsizebytes

	return r
}

// ModelType The model type.
// API name: model_type
func (r *PutTrainedModel) ModelType(modeltype trainedmodeltype.TrainedModelType) *PutTrainedModel {
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

	r.req.PlatformArchitecture = &platformarchitecture

	return r
}

// Tags An array of tags to organize the model.
// API name: tags
func (r *PutTrainedModel) Tags(tags ...string) *PutTrainedModel {
	r.req.Tags = tags

	return r
}
