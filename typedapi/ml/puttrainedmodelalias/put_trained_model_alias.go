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

// Create or update a trained model alias.
// A trained model alias is a logical name used to reference a single trained
// model.
// You can use aliases instead of trained model identifiers to make it easier to
// reference your models. For example, you can use aliases in inference
// aggregations and processors.
// An alias must be unique and refer to only a single trained model. However,
// you can have multiple aliases for each trained model.
// If you use this API to update an alias such that it references a different
// trained model ID and the model uses a different type of data frame analytics,
// an error occurs. For example, this situation occurs if you have a trained
// model for regression analysis and a trained model for classification
// analysis; you cannot reassign an alias from one type of trained model to
// another.
// If you use this API to update an alias and there are very few input fields in
// common between the old and new trained models for the model alias, the API
// returns a warning.
package puttrainedmodelalias

import (
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

	raw io.Reader

	paramSet int

	modelalias string
	modelid    string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutTrainedModelAlias type alias for index.
type NewPutTrainedModelAlias func(modelid, modelalias string) *PutTrainedModelAlias

// NewPutTrainedModelAliasFunc returns a new instance of PutTrainedModelAlias with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutTrainedModelAliasFunc(tp elastictransport.Interface) NewPutTrainedModelAlias {
	return func(modelid, modelalias string) *PutTrainedModelAlias {
		n := New(tp)

		n._modelalias(modelalias)

		n._modelid(modelid)

		return n
	}
}

// Create or update a trained model alias.
// A trained model alias is a logical name used to reference a single trained
// model.
// You can use aliases instead of trained model identifiers to make it easier to
// reference your models. For example, you can use aliases in inference
// aggregations and processors.
// An alias must be unique and refer to only a single trained model. However,
// you can have multiple aliases for each trained model.
// If you use this API to update an alias such that it references a different
// trained model ID and the model uses a different type of data frame analytics,
// an error occurs. For example, this situation occurs if you have a trained
// model for regression analysis and a trained model for classification
// analysis; you cannot reassign an alias from one type of trained model to
// another.
// If you use this API to update an alias and there are very few input fields in
// common between the old and new trained models for the model alias, the API
// returns a warning.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models-aliases.html
func New(tp elastictransport.Interface) *PutTrainedModelAlias {
	r := &PutTrainedModelAlias{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
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

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "modelid", r.modelid)
		}
		path.WriteString(r.modelid)
		path.WriteString("/")
		path.WriteString("model_aliases")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "modelalias", r.modelalias)
		}
		path.WriteString(r.modelalias)

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
func (r PutTrainedModelAlias) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.put_trained_model_alias")
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
		instrument.BeforeRequest(req, "ml.put_trained_model_alias")
		if reader := instrument.RecordRequestBody(ctx, "ml.put_trained_model_alias", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.put_trained_model_alias")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutTrainedModelAlias query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a puttrainedmodelalias.Response
func (r PutTrainedModelAlias) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.put_trained_model_alias")
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r PutTrainedModelAlias) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.put_trained_model_alias")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the PutTrainedModelAlias query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
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
func (r *PutTrainedModelAlias) _modelalias(modelalias string) *PutTrainedModelAlias {
	r.paramSet |= modelaliasMask
	r.modelalias = modelalias

	return r
}

// ModelId The identifier for the trained model that the alias refers to.
// API Name: modelid
func (r *PutTrainedModelAlias) _modelid(modelid string) *PutTrainedModelAlias {
	r.paramSet |= modelidMask
	r.modelid = modelid

	return r
}

// Reassign Specifies whether the alias gets reassigned to the specified trained
// model if it is already assigned to a different model. If the alias is
// already assigned and this parameter is false, the API returns an error.
// API name: reassign
func (r *PutTrainedModelAlias) Reassign(reassign bool) *PutTrainedModelAlias {
	r.values.Set("reassign", strconv.FormatBool(reassign))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutTrainedModelAlias) ErrorTrace(errortrace bool) *PutTrainedModelAlias {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutTrainedModelAlias) FilterPath(filterpaths ...string) *PutTrainedModelAlias {
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
func (r *PutTrainedModelAlias) Human(human bool) *PutTrainedModelAlias {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutTrainedModelAlias) Pretty(pretty bool) *PutTrainedModelAlias {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
