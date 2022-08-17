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

// IntervalsAllOf type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L49-L56
type IntervalsAllOf struct {
	Filter    *IntervalsFilter     `json:"filter,omitempty"`
	Intervals []IntervalsContainer `json:"intervals"`
	MaxGaps   *int                 `json:"max_gaps,omitempty"`
	Ordered   *bool                `json:"ordered,omitempty"`
}

// IntervalsAllOfBuilder holds IntervalsAllOf struct and provides a builder API.
type IntervalsAllOfBuilder struct {
	v *IntervalsAllOf
}

// NewIntervalsAllOf provides a builder for the IntervalsAllOf struct.
func NewIntervalsAllOfBuilder() *IntervalsAllOfBuilder {
	r := IntervalsAllOfBuilder{
		&IntervalsAllOf{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsAllOf struct
func (rb *IntervalsAllOfBuilder) Build() IntervalsAllOf {
	return *rb.v
}

func (rb *IntervalsAllOfBuilder) Filter(filter *IntervalsFilterBuilder) *IntervalsAllOfBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *IntervalsAllOfBuilder) Intervals(intervals []IntervalsContainerBuilder) *IntervalsAllOfBuilder {
	tmp := make([]IntervalsContainer, len(intervals))
	for _, value := range intervals {
		tmp = append(tmp, value.Build())
	}
	rb.v.Intervals = tmp
	return rb
}

func (rb *IntervalsAllOfBuilder) MaxGaps(maxgaps int) *IntervalsAllOfBuilder {
	rb.v.MaxGaps = &maxgaps
	return rb
}

func (rb *IntervalsAllOfBuilder) Ordered(ordered bool) *IntervalsAllOfBuilder {
	rb.v.Ordered = &ordered
	return rb
}
