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

// Processor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L156-L161
type Processor struct {
	Count        *int64                   `json:"count,omitempty"`
	Current      *int64                   `json:"current,omitempty"`
	Failed       *int64                   `json:"failed,omitempty"`
	TimeInMillis *DurationValueUnitMillis `json:"time_in_millis,omitempty"`
}

// ProcessorBuilder holds Processor struct and provides a builder API.
type ProcessorBuilder struct {
	v *Processor
}

// NewProcessor provides a builder for the Processor struct.
func NewProcessorBuilder() *ProcessorBuilder {
	r := ProcessorBuilder{
		&Processor{},
	}

	return &r
}

// Build finalize the chain and returns the Processor struct
func (rb *ProcessorBuilder) Build() Processor {
	return *rb.v
}

func (rb *ProcessorBuilder) Count(count int64) *ProcessorBuilder {
	rb.v.Count = &count
	return rb
}

func (rb *ProcessorBuilder) Current(current int64) *ProcessorBuilder {
	rb.v.Current = &current
	return rb
}

func (rb *ProcessorBuilder) Failed(failed int64) *ProcessorBuilder {
	rb.v.Failed = &failed
	return rb
}

func (rb *ProcessorBuilder) TimeInMillis(timeinmillis *DurationValueUnitMillisBuilder) *ProcessorBuilder {
	v := timeinmillis.Build()
	rb.v.TimeInMillis = &v
	return rb
}
