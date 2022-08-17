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

// NodeStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Node.ts#L28-L39
type NodeStatistics struct {
	// Failed Number of nodes that rejected the request or failed to respond. If this value
	// is not 0, a reason for the rejection or failure is included in the response.
	Failed   int          `json:"failed"`
	Failures []ErrorCause `json:"failures,omitempty"`
	// Successful Number of nodes that responded successfully to the request.
	Successful int `json:"successful"`
	// Total Total number of nodes selected by the request.
	Total int `json:"total"`
}

// NodeStatisticsBuilder holds NodeStatistics struct and provides a builder API.
type NodeStatisticsBuilder struct {
	v *NodeStatistics
}

// NewNodeStatistics provides a builder for the NodeStatistics struct.
func NewNodeStatisticsBuilder() *NodeStatisticsBuilder {
	r := NodeStatisticsBuilder{
		&NodeStatistics{},
	}

	return &r
}

// Build finalize the chain and returns the NodeStatistics struct
func (rb *NodeStatisticsBuilder) Build() NodeStatistics {
	return *rb.v
}

// Failed Number of nodes that rejected the request or failed to respond. If this value
// is not 0, a reason for the rejection or failure is included in the response.

func (rb *NodeStatisticsBuilder) Failed(failed int) *NodeStatisticsBuilder {
	rb.v.Failed = failed
	return rb
}

func (rb *NodeStatisticsBuilder) Failures(failures []ErrorCauseBuilder) *NodeStatisticsBuilder {
	tmp := make([]ErrorCause, len(failures))
	for _, value := range failures {
		tmp = append(tmp, value.Build())
	}
	rb.v.Failures = tmp
	return rb
}

// Successful Number of nodes that responded successfully to the request.

func (rb *NodeStatisticsBuilder) Successful(successful int) *NodeStatisticsBuilder {
	rb.v.Successful = successful
	return rb
}

// Total Total number of nodes selected by the request.

func (rb *NodeStatisticsBuilder) Total(total int) *NodeStatisticsBuilder {
	rb.v.Total = total
	return rb
}
