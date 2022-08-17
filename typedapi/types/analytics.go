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

// Analytics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L313-L315
type Analytics struct {
	Available bool                `json:"available"`
	Enabled   bool                `json:"enabled"`
	Stats     AnalyticsStatistics `json:"stats"`
}

// AnalyticsBuilder holds Analytics struct and provides a builder API.
type AnalyticsBuilder struct {
	v *Analytics
}

// NewAnalytics provides a builder for the Analytics struct.
func NewAnalyticsBuilder() *AnalyticsBuilder {
	r := AnalyticsBuilder{
		&Analytics{},
	}

	return &r
}

// Build finalize the chain and returns the Analytics struct
func (rb *AnalyticsBuilder) Build() Analytics {
	return *rb.v
}

func (rb *AnalyticsBuilder) Available(available bool) *AnalyticsBuilder {
	rb.v.Available = available
	return rb
}

func (rb *AnalyticsBuilder) Enabled(enabled bool) *AnalyticsBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *AnalyticsBuilder) Stats(stats *AnalyticsStatisticsBuilder) *AnalyticsBuilder {
	v := stats.Build()
	rb.v.Stats = v
	return rb
}
