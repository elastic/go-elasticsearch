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


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/acknowledgementoptions"
)

// AcknowledgeState type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Action.ts#L112-L115
type AcknowledgeState struct {
	State     acknowledgementoptions.AcknowledgementOptions `json:"state"`
	Timestamp DateTime                                      `json:"timestamp"`
}

// AcknowledgeStateBuilder holds AcknowledgeState struct and provides a builder API.
type AcknowledgeStateBuilder struct {
	v *AcknowledgeState
}

// NewAcknowledgeState provides a builder for the AcknowledgeState struct.
func NewAcknowledgeStateBuilder() *AcknowledgeStateBuilder {
	r := AcknowledgeStateBuilder{
		&AcknowledgeState{},
	}

	return &r
}

// Build finalize the chain and returns the AcknowledgeState struct
func (rb *AcknowledgeStateBuilder) Build() AcknowledgeState {
	return *rb.v
}

func (rb *AcknowledgeStateBuilder) State(state acknowledgementoptions.AcknowledgementOptions) *AcknowledgeStateBuilder {
	rb.v.State = state
	return rb
}

func (rb *AcknowledgeStateBuilder) Timestamp(timestamp *DateTimeBuilder) *AcknowledgeStateBuilder {
	v := timestamp.Build()
	rb.v.Timestamp = v
	return rb
}
