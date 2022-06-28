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

// Row type alias.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/sql/types.ts#L28-L28
type Row []interface{}

// RowBuilder holds Row struct and provides a builder API.
type RowBuilder struct {
	v Row
}

// NewRow provides a builder for the Row struct.
func NewRowBuilder() *RowBuilder {
	return &RowBuilder{}
}

// Build finalize the chain and returns the Row struct
func (b *RowBuilder) Build() Row {
	return b.v
}

func (b *RowBuilder) Row(value Row) *RowBuilder {
	b.v = value
	return b
}
