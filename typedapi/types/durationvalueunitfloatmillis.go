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

// DurationValueUnitFloatMillis type alias.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Time.ts#L67-L67
type DurationValueUnitFloatMillis UnitFloatMillis

// DurationValueUnitFloatMillisBuilder holds DurationValueUnitFloatMillis struct and provides a builder API.
type DurationValueUnitFloatMillisBuilder struct {
	v DurationValueUnitFloatMillis
}

// NewDurationValueUnitFloatMillis provides a builder for the DurationValueUnitFloatMillis struct.
func NewDurationValueUnitFloatMillisBuilder() *DurationValueUnitFloatMillisBuilder {
	return &DurationValueUnitFloatMillisBuilder{}
}

// Build finalize the chain and returns the DurationValueUnitFloatMillis struct
func (b *DurationValueUnitFloatMillisBuilder) Build() DurationValueUnitFloatMillis {
	return b.v
}

func (b *DurationValueUnitFloatMillisBuilder) DurationValueUnitFloatMillis(value DurationValueUnitFloatMillis) *DurationValueUnitFloatMillisBuilder {
	b.v = value
	return b
}
