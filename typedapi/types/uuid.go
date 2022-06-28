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

// Uuid type alias.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/common.ts#L99-L99
type Uuid string

// UuidBuilder holds Uuid struct and provides a builder API.
type UuidBuilder struct {
	v Uuid
}

// NewUuid provides a builder for the Uuid struct.
func NewUuidBuilder() *UuidBuilder {
	return &UuidBuilder{}
}

// Build finalize the chain and returns the Uuid struct
func (b *UuidBuilder) Build() Uuid {
	return b.v
}

func (b *UuidBuilder) Uuid(value Uuid) *UuidBuilder {
	b.v = value
	return b
}
