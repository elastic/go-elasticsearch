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

// AdaptiveSelection type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L163-L171
type AdaptiveSelection struct {
	AvgQueueSize      *int64  `json:"avg_queue_size,omitempty"`
	AvgResponseTime   *int64  `json:"avg_response_time,omitempty"`
	AvgResponseTimeNs *int64  `json:"avg_response_time_ns,omitempty"`
	AvgServiceTime    *string `json:"avg_service_time,omitempty"`
	AvgServiceTimeNs  *int64  `json:"avg_service_time_ns,omitempty"`
	OutgoingSearches  *int64  `json:"outgoing_searches,omitempty"`
	Rank              *string `json:"rank,omitempty"`
}

// AdaptiveSelectionBuilder holds AdaptiveSelection struct and provides a builder API.
type AdaptiveSelectionBuilder struct {
	v *AdaptiveSelection
}

// NewAdaptiveSelection provides a builder for the AdaptiveSelection struct.
func NewAdaptiveSelectionBuilder() *AdaptiveSelectionBuilder {
	r := AdaptiveSelectionBuilder{
		&AdaptiveSelection{},
	}

	return &r
}

// Build finalize the chain and returns the AdaptiveSelection struct
func (rb *AdaptiveSelectionBuilder) Build() AdaptiveSelection {
	return *rb.v
}

func (rb *AdaptiveSelectionBuilder) AvgQueueSize(avgqueuesize int64) *AdaptiveSelectionBuilder {
	rb.v.AvgQueueSize = &avgqueuesize
	return rb
}

func (rb *AdaptiveSelectionBuilder) AvgResponseTime(avgresponsetime int64) *AdaptiveSelectionBuilder {
	rb.v.AvgResponseTime = &avgresponsetime
	return rb
}

func (rb *AdaptiveSelectionBuilder) AvgResponseTimeNs(avgresponsetimens int64) *AdaptiveSelectionBuilder {
	rb.v.AvgResponseTimeNs = &avgresponsetimens
	return rb
}

func (rb *AdaptiveSelectionBuilder) AvgServiceTime(avgservicetime string) *AdaptiveSelectionBuilder {
	rb.v.AvgServiceTime = &avgservicetime
	return rb
}

func (rb *AdaptiveSelectionBuilder) AvgServiceTimeNs(avgservicetimens int64) *AdaptiveSelectionBuilder {
	rb.v.AvgServiceTimeNs = &avgservicetimens
	return rb
}

func (rb *AdaptiveSelectionBuilder) OutgoingSearches(outgoingsearches int64) *AdaptiveSelectionBuilder {
	rb.v.OutgoingSearches = &outgoingsearches
	return rb
}

func (rb *AdaptiveSelectionBuilder) Rank(rank string) *AdaptiveSelectionBuilder {
	rb.v.Rank = &rank
	return rb
}
