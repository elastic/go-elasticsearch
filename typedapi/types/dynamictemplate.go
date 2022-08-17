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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/matchtype"
)

// DynamicTemplate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/dynamic-template.ts#L22-L30
type DynamicTemplate struct {
	Mapping          *Property            `json:"mapping,omitempty"`
	Match            *string              `json:"match,omitempty"`
	MatchMappingType *string              `json:"match_mapping_type,omitempty"`
	MatchPattern     *matchtype.MatchType `json:"match_pattern,omitempty"`
	PathMatch        *string              `json:"path_match,omitempty"`
	PathUnmatch      *string              `json:"path_unmatch,omitempty"`
	Unmatch          *string              `json:"unmatch,omitempty"`
}

// DynamicTemplateBuilder holds DynamicTemplate struct and provides a builder API.
type DynamicTemplateBuilder struct {
	v *DynamicTemplate
}

// NewDynamicTemplate provides a builder for the DynamicTemplate struct.
func NewDynamicTemplateBuilder() *DynamicTemplateBuilder {
	r := DynamicTemplateBuilder{
		&DynamicTemplate{},
	}

	return &r
}

// Build finalize the chain and returns the DynamicTemplate struct
func (rb *DynamicTemplateBuilder) Build() DynamicTemplate {
	return *rb.v
}

func (rb *DynamicTemplateBuilder) Mapping(mapping *PropertyBuilder) *DynamicTemplateBuilder {
	v := mapping.Build()
	rb.v.Mapping = &v
	return rb
}

func (rb *DynamicTemplateBuilder) Match(match string) *DynamicTemplateBuilder {
	rb.v.Match = &match
	return rb
}

func (rb *DynamicTemplateBuilder) MatchMappingType(matchmappingtype string) *DynamicTemplateBuilder {
	rb.v.MatchMappingType = &matchmappingtype
	return rb
}

func (rb *DynamicTemplateBuilder) MatchPattern(matchpattern matchtype.MatchType) *DynamicTemplateBuilder {
	rb.v.MatchPattern = &matchpattern
	return rb
}

func (rb *DynamicTemplateBuilder) PathMatch(pathmatch string) *DynamicTemplateBuilder {
	rb.v.PathMatch = &pathmatch
	return rb
}

func (rb *DynamicTemplateBuilder) PathUnmatch(pathunmatch string) *DynamicTemplateBuilder {
	rb.v.PathUnmatch = &pathunmatch
	return rb
}

func (rb *DynamicTemplateBuilder) Unmatch(unmatch string) *DynamicTemplateBuilder {
	rb.v.Unmatch = &unmatch
	return rb
}
