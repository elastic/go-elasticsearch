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

// HitsEvent type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/eql/_types/EqlHits.ts#L41-L49
type HitsEvent struct {
	Fields map[Field][]interface{} `json:"fields,omitempty"`
	// Id_ Unique identifier for the event. This ID is only unique within the index.
	Id_ Id `json:"_id"`
	// Index_ Name of the index containing the event.
	Index_ IndexName `json:"_index"`
	// Source_ Original JSON body passed for the event at index time.
	Source_ interface{} `json:"_source,omitempty"`
}

// HitsEventBuilder holds HitsEvent struct and provides a builder API.
type HitsEventBuilder struct {
	v *HitsEvent
}

// NewHitsEvent provides a builder for the HitsEvent struct.
func NewHitsEventBuilder() *HitsEventBuilder {
	r := HitsEventBuilder{
		&HitsEvent{
			Fields: make(map[Field][]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the HitsEvent struct
func (rb *HitsEventBuilder) Build() HitsEvent {
	return *rb.v
}

func (rb *HitsEventBuilder) Fields(value map[Field][]interface{}) *HitsEventBuilder {
	rb.v.Fields = value
	return rb
}

// Id_ Unique identifier for the event. This ID is only unique within the index.

func (rb *HitsEventBuilder) Id_(id_ Id) *HitsEventBuilder {
	rb.v.Id_ = id_
	return rb
}

// Index_ Name of the index containing the event.

func (rb *HitsEventBuilder) Index_(index_ IndexName) *HitsEventBuilder {
	rb.v.Index_ = index_
	return rb
}

// Source_ Original JSON body passed for the event at index time.

func (rb *HitsEventBuilder) Source_(source_ interface{}) *HitsEventBuilder {
	rb.v.Source_ = source_
	return rb
}
