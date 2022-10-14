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
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


package types

// RolloverConditions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/rollover/types.ts#L24-L37
type RolloverConditions struct {
	MaxAge              *Duration                `json:"max_age,omitempty"`
	MaxAgeMillis        *DurationValueUnitMillis `json:"max_age_millis,omitempty"`
	MaxDocs             *int64                   `json:"max_docs,omitempty"`
	MaxPrimaryShardDocs *int64                   `json:"max_primary_shard_docs,omitempty"`
	MaxPrimaryShardSize *ByteSize                `json:"max_primary_shard_size,omitempty"`
	MaxSize             *ByteSize                `json:"max_size,omitempty"`
	MaxSizeBytes        *ByteSize                `json:"max_size_bytes,omitempty"`
	MinAge              *Duration                `json:"min_age,omitempty"`
	MinDocs             *int64                   `json:"min_docs,omitempty"`
	MinPrimaryShardDocs *int64                   `json:"min_primary_shard_docs,omitempty"`
	MinPrimaryShardSize *ByteSize                `json:"min_primary_shard_size,omitempty"`
	MinSize             *ByteSize                `json:"min_size,omitempty"`
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

func (rb *RolloverConditionsBuilder) MaxPrimaryShardDocs(maxprimarysharddocs int64) *RolloverConditionsBuilder {
	rb.v.MaxPrimaryShardDocs = &maxprimarysharddocs
	return rb
}

func (rb *RolloverConditionsBuilder) MaxPrimaryShardSize(maxprimaryshardsize *ByteSizeBuilder) *RolloverConditionsBuilder {
	v := maxprimaryshardsize.Build()
	rb.v.MaxPrimaryShardSize = &v
	return rb
}

func (rb *RolloverConditionsBuilder) MaxSize(maxsize *ByteSizeBuilder) *RolloverConditionsBuilder {
	v := maxsize.Build()
	rb.v.MaxSize = &v
	return rb
}

func (rb *RolloverConditionsBuilder) MaxSizeBytes(maxsizebytes *ByteSizeBuilder) *RolloverConditionsBuilder {
	v := maxsizebytes.Build()
	rb.v.MaxSizeBytes = &v
	return rb
}

func (rb *RolloverConditionsBuilder) MinAge(minage *DurationBuilder) *RolloverConditionsBuilder {
	v := minage.Build()
	rb.v.MinAge = &v
	return rb
}

func (rb *RolloverConditionsBuilder) MinDocs(mindocs int64) *RolloverConditionsBuilder {
	rb.v.MinDocs = &mindocs
	return rb
}

func (rb *RolloverConditionsBuilder) MinPrimaryShardDocs(minprimarysharddocs int64) *RolloverConditionsBuilder {
	rb.v.MinPrimaryShardDocs = &minprimarysharddocs
	return rb
}

func (rb *RolloverConditionsBuilder) MinPrimaryShardSize(minprimaryshardsize *ByteSizeBuilder) *RolloverConditionsBuilder {
	v := minprimaryshardsize.Build()
	rb.v.MinPrimaryShardSize = &v
	return rb
}

func (rb *RolloverConditionsBuilder) MinSize(minsize *ByteSizeBuilder) *RolloverConditionsBuilder {
	v := minsize.Build()
	rb.v.MinSize = &v
	return rb
}
