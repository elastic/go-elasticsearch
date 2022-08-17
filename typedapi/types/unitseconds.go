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

// UnitSeconds type alias.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Time.ts#L29-L30
type UnitSeconds int64

// UnitSecondsBuilder holds UnitSeconds struct and provides a builder API.
type UnitSecondsBuilder struct {
	v UnitSeconds
}

// NewUnitSeconds provides a builder for the UnitSeconds struct.
func NewUnitSecondsBuilder() *UnitSecondsBuilder {
	return &UnitSecondsBuilder{}
}

// Build finalize the chain and returns the UnitSeconds struct
func (b *UnitSecondsBuilder) Build() UnitSeconds {
	return b.v
}

func (b *UnitSecondsBuilder) UnitSeconds(value UnitSeconds) *UnitSecondsBuilder {
	b.v = value
	return b
}
