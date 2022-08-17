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

// ExploreControls type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/graph/_types/ExploreControls.ts#L24-L29
type ExploreControls struct {
	SampleDiversity *SampleDiversity `json:"sample_diversity,omitempty"`
	SampleSize      *int             `json:"sample_size,omitempty"`
	Timeout         *Duration        `json:"timeout,omitempty"`
	UseSignificance bool             `json:"use_significance"`
}

// ExploreControlsBuilder holds ExploreControls struct and provides a builder API.
type ExploreControlsBuilder struct {
	v *ExploreControls
}

// NewExploreControls provides a builder for the ExploreControls struct.
func NewExploreControlsBuilder() *ExploreControlsBuilder {
	r := ExploreControlsBuilder{
		&ExploreControls{},
	}

	return &r
}

// Build finalize the chain and returns the ExploreControls struct
func (rb *ExploreControlsBuilder) Build() ExploreControls {
	return *rb.v
}

func (rb *ExploreControlsBuilder) SampleDiversity(samplediversity *SampleDiversityBuilder) *ExploreControlsBuilder {
	v := samplediversity.Build()
	rb.v.SampleDiversity = &v
	return rb
}

func (rb *ExploreControlsBuilder) SampleSize(samplesize int) *ExploreControlsBuilder {
	rb.v.SampleSize = &samplesize
	return rb
}

func (rb *ExploreControlsBuilder) Timeout(timeout *DurationBuilder) *ExploreControlsBuilder {
	v := timeout.Build()
	rb.v.Timeout = &v
	return rb
}

func (rb *ExploreControlsBuilder) UseSignificance(usesignificance bool) *ExploreControlsBuilder {
	rb.v.UseSignificance = usesignificance
	return rb
}
