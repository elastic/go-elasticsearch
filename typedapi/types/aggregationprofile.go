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

// AggregationProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/profile.ts#L75-L82
type AggregationProfile struct {
	Breakdown   AggregationBreakdown     `json:"breakdown"`
	Children    []AggregationProfile     `json:"children,omitempty"`
	Debug       *AggregationProfileDebug `json:"debug,omitempty"`
	Description string                   `json:"description"`
	TimeInNanos DurationValueUnitNanos   `json:"time_in_nanos"`
	Type        string                   `json:"type"`
}

// AggregationProfileBuilder holds AggregationProfile struct and provides a builder API.
type AggregationProfileBuilder struct {
	v *AggregationProfile
}

// NewAggregationProfile provides a builder for the AggregationProfile struct.
func NewAggregationProfileBuilder() *AggregationProfileBuilder {
	r := AggregationProfileBuilder{
		&AggregationProfile{},
	}

	return &r
}

// Build finalize the chain and returns the AggregationProfile struct
func (rb *AggregationProfileBuilder) Build() AggregationProfile {
	return *rb.v
}

func (rb *AggregationProfileBuilder) Breakdown(breakdown *AggregationBreakdownBuilder) *AggregationProfileBuilder {
	v := breakdown.Build()
	rb.v.Breakdown = v
	return rb
}

func (rb *AggregationProfileBuilder) Children(children []AggregationProfileBuilder) *AggregationProfileBuilder {
	tmp := make([]AggregationProfile, len(children))
	for _, value := range children {
		tmp = append(tmp, value.Build())
	}
	rb.v.Children = tmp
	return rb
}

func (rb *AggregationProfileBuilder) Debug(debug *AggregationProfileDebugBuilder) *AggregationProfileBuilder {
	v := debug.Build()
	rb.v.Debug = &v
	return rb
}

func (rb *AggregationProfileBuilder) Description(description string) *AggregationProfileBuilder {
	rb.v.Description = description
	return rb
}

func (rb *AggregationProfileBuilder) TimeInNanos(timeinnanos *DurationValueUnitNanosBuilder) *AggregationProfileBuilder {
	v := timeinnanos.Build()
	rb.v.TimeInNanos = v
	return rb
}

func (rb *AggregationProfileBuilder) Type_(type_ string) *AggregationProfileBuilder {
	rb.v.Type = type_
	return rb
}
