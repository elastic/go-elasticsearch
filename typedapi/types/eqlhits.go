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

// EqlHits type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/eql/_types/EqlHits.ts#L25-L39
type EqlHits struct {
	// Events Contains events matching the query. Each object represents a matching event.
	Events []HitsEvent `json:"events,omitempty"`
	// Sequences Contains event sequences matching the query. Each object represents a
	// matching sequence. This parameter is only returned for EQL queries containing
	// a sequence.
	Sequences []HitsSequence `json:"sequences,omitempty"`
	// Total Metadata about the number of matching events or sequences.
	Total *TotalHits `json:"total,omitempty"`
}

// EqlHitsBuilder holds EqlHits struct and provides a builder API.
type EqlHitsBuilder struct {
	v *EqlHits
}

// NewEqlHits provides a builder for the EqlHits struct.
func NewEqlHitsBuilder() *EqlHitsBuilder {
	r := EqlHitsBuilder{
		&EqlHits{},
	}

	return &r
}

// Build finalize the chain and returns the EqlHits struct
func (rb *EqlHitsBuilder) Build() EqlHits {
	return *rb.v
}

// Events Contains events matching the query. Each object represents a matching event.

func (rb *EqlHitsBuilder) Events(events []HitsEventBuilder) *EqlHitsBuilder {
	tmp := make([]HitsEvent, len(events))
	for _, value := range events {
		tmp = append(tmp, value.Build())
	}
	rb.v.Events = tmp
	return rb
}

// Sequences Contains event sequences matching the query. Each object represents a
// matching sequence. This parameter is only returned for EQL queries containing
// a sequence.

func (rb *EqlHitsBuilder) Sequences(sequences []HitsSequenceBuilder) *EqlHitsBuilder {
	tmp := make([]HitsSequence, len(sequences))
	for _, value := range sequences {
		tmp = append(tmp, value.Build())
	}
	rb.v.Sequences = tmp
	return rb
}

// Total Metadata about the number of matching events or sequences.

func (rb *EqlHitsBuilder) Total(total *TotalHitsBuilder) *EqlHitsBuilder {
	v := total.Build()
	rb.v.Total = &v
	return rb
}
