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

// AutoFollowPattern type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ccr/get_auto_follow_pattern/types.ts#L23-L26
type AutoFollowPattern struct {
	Name    Name                     `json:"name"`
	Pattern AutoFollowPatternSummary `json:"pattern"`
}

// AutoFollowPatternBuilder holds AutoFollowPattern struct and provides a builder API.
type AutoFollowPatternBuilder struct {
	v *AutoFollowPattern
}

// NewAutoFollowPattern provides a builder for the AutoFollowPattern struct.
func NewAutoFollowPatternBuilder() *AutoFollowPatternBuilder {
	r := AutoFollowPatternBuilder{
		&AutoFollowPattern{},
	}

	return &r
}

// Build finalize the chain and returns the AutoFollowPattern struct
func (rb *AutoFollowPatternBuilder) Build() AutoFollowPattern {
	return *rb.v
}

func (rb *AutoFollowPatternBuilder) Name(name Name) *AutoFollowPatternBuilder {
	rb.v.Name = name
	return rb
}

func (rb *AutoFollowPatternBuilder) Pattern(pattern *AutoFollowPatternSummaryBuilder) *AutoFollowPatternBuilder {
	v := pattern.Build()
	rb.v.Pattern = v
	return rb
}
