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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// DocStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/Stats.ts#L63-L66
type DocStats struct {
	Count   int64  `json:"count"`
	Deleted *int64 `json:"deleted,omitempty"`
}

// DocStatsBuilder holds DocStats struct and provides a builder API.
type DocStatsBuilder struct {
	v *DocStats
}

// NewDocStats provides a builder for the DocStats struct.
func NewDocStatsBuilder() *DocStatsBuilder {
	r := DocStatsBuilder{
		&DocStats{},
	}

	return &r
}

// Build finalize the chain and returns the DocStats struct
func (rb *DocStatsBuilder) Build() DocStats {
	return *rb.v
}

func (rb *DocStatsBuilder) Count(count int64) *DocStatsBuilder {
	rb.v.Count = count
	return rb
}

func (rb *DocStatsBuilder) Deleted(deleted int64) *DocStatsBuilder {
	rb.v.Deleted = &deleted
	return rb
}
