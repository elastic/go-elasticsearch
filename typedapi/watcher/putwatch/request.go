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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package putwatch

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putwatch
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/put_watch/WatcherPutWatchRequest.ts#L30-L54
type Request struct {
	Actions map[string]types.Action `json:"actions,omitempty"`

	Condition *types.ConditionContainer `json:"condition,omitempty"`

	Input *types.InputContainer `json:"input,omitempty"`

	Metadata *types.Metadata `json:"metadata,omitempty"`

	ThrottlePeriod *string `json:"throttle_period,omitempty"`

	Transform *types.TransformContainer `json:"transform,omitempty"`

	Trigger *types.TriggerContainer `json:"trigger,omitempty"`
}

// RequestBuilder is the builder API for the putwatch.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			Actions: make(map[string]types.Action, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putwatch request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Actions(values map[string]*types.ActionBuilder) *RequestBuilder {
	tmp := make(map[string]types.Action, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Actions = tmp
	return rb
}

func (rb *RequestBuilder) Condition(condition *types.ConditionContainerBuilder) *RequestBuilder {
	v := condition.Build()
	rb.v.Condition = &v
	return rb
}

func (rb *RequestBuilder) Input(input *types.InputContainerBuilder) *RequestBuilder {
	v := input.Build()
	rb.v.Input = &v
	return rb
}

func (rb *RequestBuilder) Metadata(metadata *types.MetadataBuilder) *RequestBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

func (rb *RequestBuilder) ThrottlePeriod(throttleperiod string) *RequestBuilder {
	rb.v.ThrottlePeriod = &throttleperiod
	return rb
}

func (rb *RequestBuilder) Transform(transform *types.TransformContainerBuilder) *RequestBuilder {
	v := transform.Build()
	rb.v.Transform = &v
	return rb
}

func (rb *RequestBuilder) Trigger(trigger *types.TriggerContainerBuilder) *RequestBuilder {
	v := trigger.Build()
	rb.v.Trigger = &v
	return rb
}
