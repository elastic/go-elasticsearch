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

// EqlFeatures type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L94-L102
type EqlFeatures struct {
	Event     uint                 `json:"event"`
	Join      uint                 `json:"join"`
	Joins     EqlFeaturesJoin      `json:"joins"`
	Keys      EqlFeaturesKeys      `json:"keys"`
	Pipes     EqlFeaturesPipes     `json:"pipes"`
	Sequence  uint                 `json:"sequence"`
	Sequences EqlFeaturesSequences `json:"sequences"`
}

// EqlFeaturesBuilder holds EqlFeatures struct and provides a builder API.
type EqlFeaturesBuilder struct {
	v *EqlFeatures
}

// NewEqlFeatures provides a builder for the EqlFeatures struct.
func NewEqlFeaturesBuilder() *EqlFeaturesBuilder {
	r := EqlFeaturesBuilder{
		&EqlFeatures{},
	}

	return &r
}

// Build finalize the chain and returns the EqlFeatures struct
func (rb *EqlFeaturesBuilder) Build() EqlFeatures {
	return *rb.v
}

func (rb *EqlFeaturesBuilder) Event(event uint) *EqlFeaturesBuilder {
	rb.v.Event = event
	return rb
}

func (rb *EqlFeaturesBuilder) Join(join uint) *EqlFeaturesBuilder {
	rb.v.Join = join
	return rb
}

func (rb *EqlFeaturesBuilder) Joins(joins *EqlFeaturesJoinBuilder) *EqlFeaturesBuilder {
	v := joins.Build()
	rb.v.Joins = v
	return rb
}

func (rb *EqlFeaturesBuilder) Keys(keys *EqlFeaturesKeysBuilder) *EqlFeaturesBuilder {
	v := keys.Build()
	rb.v.Keys = v
	return rb
}

func (rb *EqlFeaturesBuilder) Pipes(pipes *EqlFeaturesPipesBuilder) *EqlFeaturesBuilder {
	v := pipes.Build()
	rb.v.Pipes = v
	return rb
}

func (rb *EqlFeaturesBuilder) Sequence(sequence uint) *EqlFeaturesBuilder {
	rb.v.Sequence = sequence
	return rb
}

func (rb *EqlFeaturesBuilder) Sequences(sequences *EqlFeaturesSequencesBuilder) *EqlFeaturesBuilder {
	v := sequences.Build()
	rb.v.Sequences = v
	return rb
}
