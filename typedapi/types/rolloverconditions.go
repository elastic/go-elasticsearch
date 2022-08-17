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

// RolloverConditions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/rollover/types.ts#L24-L32
type RolloverConditions struct {
	MaxAge                   *Duration                `json:"max_age,omitempty"`
	MaxAgeMillis             *DurationValueUnitMillis `json:"max_age_millis,omitempty"`
	MaxDocs                  *int64                   `json:"max_docs,omitempty"`
	MaxPrimaryShardSize      *ByteSize                `json:"max_primary_shard_size,omitempty"`
	MaxPrimaryShardSizeBytes *ByteSize                `json:"max_primary_shard_size_bytes,omitempty"`
	MaxSize                  *string                  `json:"max_size,omitempty"`
	MaxSizeBytes             *ByteSize                `json:"max_size_bytes,omitempty"`
}

// RolloverConditionsBuilder holds RolloverConditions struct and provides a builder API.
type RolloverConditionsBuilder struct {
	v *RolloverConditions
}

// NewRolloverConditions provides a builder for the RolloverConditions struct.
func NewRolloverConditionsBuilder() *RolloverConditionsBuilder {
	r := RolloverConditionsBuilder{
		&RolloverConditions{},
	}

	return &r
}

// Build finalize the chain and returns the RolloverConditions struct
func (rb *RolloverConditionsBuilder) Build() RolloverConditions {
	return *rb.v
}

func (rb *RolloverConditionsBuilder) MaxAge(maxage *DurationBuilder) *RolloverConditionsBuilder {
	v := maxage.Build()
	rb.v.MaxAge = &v
	return rb
}

func (rb *RolloverConditionsBuilder) MaxAgeMillis(maxagemillis *DurationValueUnitMillisBuilder) *RolloverConditionsBuilder {
	v := maxagemillis.Build()
	rb.v.MaxAgeMillis = &v
	return rb
}

func (rb *RolloverConditionsBuilder) MaxDocs(maxdocs int64) *RolloverConditionsBuilder {
	rb.v.MaxDocs = &maxdocs
	return rb
}

func (rb *RolloverConditionsBuilder) MaxPrimaryShardSize(maxprimaryshardsize *ByteSizeBuilder) *RolloverConditionsBuilder {
	v := maxprimaryshardsize.Build()
	rb.v.MaxPrimaryShardSize = &v
	return rb
}

func (rb *RolloverConditionsBuilder) MaxPrimaryShardSizeBytes(maxprimaryshardsizebytes *ByteSizeBuilder) *RolloverConditionsBuilder {
	v := maxprimaryshardsizebytes.Build()
	rb.v.MaxPrimaryShardSizeBytes = &v
	return rb
}

func (rb *RolloverConditionsBuilder) MaxSize(maxsize string) *RolloverConditionsBuilder {
	rb.v.MaxSize = &maxsize
	return rb
}

func (rb *RolloverConditionsBuilder) MaxSizeBytes(maxsizebytes *ByteSizeBuilder) *RolloverConditionsBuilder {
	v := maxsizebytes.Build()
	rb.v.MaxSizeBytes = &v
	return rb
}
