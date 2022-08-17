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

// SampleDiversity type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/graph/_types/ExploreControls.ts#L31-L34
type SampleDiversity struct {
	Field           Field `json:"field"`
	MaxDocsPerValue int   `json:"max_docs_per_value"`
}

// SampleDiversityBuilder holds SampleDiversity struct and provides a builder API.
type SampleDiversityBuilder struct {
	v *SampleDiversity
}

// NewSampleDiversity provides a builder for the SampleDiversity struct.
func NewSampleDiversityBuilder() *SampleDiversityBuilder {
	r := SampleDiversityBuilder{
		&SampleDiversity{},
	}

	return &r
}

// Build finalize the chain and returns the SampleDiversity struct
func (rb *SampleDiversityBuilder) Build() SampleDiversity {
	return *rb.v
}

func (rb *SampleDiversityBuilder) Field(field Field) *SampleDiversityBuilder {
	rb.v.Field = field
	return rb
}

func (rb *SampleDiversityBuilder) MaxDocsPerValue(maxdocspervalue int) *SampleDiversityBuilder {
	rb.v.MaxDocsPerValue = maxdocspervalue
	return rb
}
