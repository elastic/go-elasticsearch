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

// ClusterAppliedStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L84-L86
type ClusterAppliedStats struct {
	Recordings []Recording `json:"recordings,omitempty"`
}

// ClusterAppliedStatsBuilder holds ClusterAppliedStats struct and provides a builder API.
type ClusterAppliedStatsBuilder struct {
	v *ClusterAppliedStats
}

// NewClusterAppliedStats provides a builder for the ClusterAppliedStats struct.
func NewClusterAppliedStatsBuilder() *ClusterAppliedStatsBuilder {
	r := ClusterAppliedStatsBuilder{
		&ClusterAppliedStats{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterAppliedStats struct
func (rb *ClusterAppliedStatsBuilder) Build() ClusterAppliedStats {
	return *rb.v
}

func (rb *ClusterAppliedStatsBuilder) Recordings(recordings []RecordingBuilder) *ClusterAppliedStatsBuilder {
	tmp := make([]Recording, len(recordings))
	for _, value := range recordings {
		tmp = append(tmp, value.Build())
	}
	rb.v.Recordings = tmp
	return rb
}
