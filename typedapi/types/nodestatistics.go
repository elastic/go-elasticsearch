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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// NodeStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/Node.ts#L28-L39
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

// NewNodeStatistics returns a NodeStatistics.
func NewNodeStatistics() *NodeStatistics {
	r := &NodeStatistics{}

	return r
}
