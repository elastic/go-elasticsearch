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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// RankFeatureProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/core.ts#L170-L173
type RankFeatureProperty struct {
	Dynamic             *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields              map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove         *int                           `json:"ignore_above,omitempty"`
	LocalMetadata       *Metadata                      `json:"local_metadata,omitempty"`
	Meta                map[string]string              `json:"meta,omitempty"`
	PositiveScoreImpact *bool                          `json:"positive_score_impact,omitempty"`
	Properties          map[PropertyName]Property      `json:"properties,omitempty"`
	Type                string                         `json:"type,omitempty"`
}

// RankFeaturePropertyBuilder holds RankFeatureProperty struct and provides a builder API.
type RankFeaturePropertyBuilder struct {
	v *RankFeatureProperty
}

// NewRankFeatureProperty provides a builder for the RankFeatureProperty struct.
func NewRankFeaturePropertyBuilder() *RankFeaturePropertyBuilder {
	r := RankFeaturePropertyBuilder{
		&RankFeatureProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "rank_feature"

	return &r
}

// Build finalize the chain and returns the RankFeatureProperty struct
func (rb *RankFeaturePropertyBuilder) Build() RankFeatureProperty {
	return *rb.v
}

func (rb *RankFeaturePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *RankFeaturePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *RankFeaturePropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *RankFeaturePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *RankFeaturePropertyBuilder) IgnoreAbove(ignoreabove int) *RankFeaturePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *RankFeaturePropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *RankFeaturePropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *RankFeaturePropertyBuilder) Meta(value map[string]string) *RankFeaturePropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *RankFeaturePropertyBuilder) PositiveScoreImpact(positivescoreimpact bool) *RankFeaturePropertyBuilder {
	rb.v.PositiveScoreImpact = &positivescoreimpact
	return rb
}

func (rb *RankFeaturePropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *RankFeaturePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}
