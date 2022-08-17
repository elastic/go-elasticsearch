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

// Counter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L34-L37
type Counter struct {
	Active int64 `json:"active"`
	Total  int64 `json:"total"`
}

// CounterBuilder holds Counter struct and provides a builder API.
type CounterBuilder struct {
	v *Counter
}

// NewCounter provides a builder for the Counter struct.
func NewCounterBuilder() *CounterBuilder {
	r := CounterBuilder{
		&Counter{},
	}

	return &r
}

// Build finalize the chain and returns the Counter struct
func (rb *CounterBuilder) Build() Counter {
	return *rb.v
}

func (rb *CounterBuilder) Active(active int64) *CounterBuilder {
	rb.v.Active = active
	return rb
}

func (rb *CounterBuilder) Total(total int64) *CounterBuilder {
	rb.v.Total = total
	return rb
}
