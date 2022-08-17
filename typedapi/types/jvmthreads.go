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

// JvmThreads type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L346-L349
type JvmThreads struct {
	Count     *int64 `json:"count,omitempty"`
	PeakCount *int64 `json:"peak_count,omitempty"`
}

// JvmThreadsBuilder holds JvmThreads struct and provides a builder API.
type JvmThreadsBuilder struct {
	v *JvmThreads
}

// NewJvmThreads provides a builder for the JvmThreads struct.
func NewJvmThreadsBuilder() *JvmThreadsBuilder {
	r := JvmThreadsBuilder{
		&JvmThreads{},
	}

	return &r
}

// Build finalize the chain and returns the JvmThreads struct
func (rb *JvmThreadsBuilder) Build() JvmThreads {
	return *rb.v
}

func (rb *JvmThreadsBuilder) Count(count int64) *JvmThreadsBuilder {
	rb.v.Count = &count
	return rb
}

func (rb *JvmThreadsBuilder) PeakCount(peakcount int64) *JvmThreadsBuilder {
	rb.v.PeakCount = &peakcount
	return rb
}
