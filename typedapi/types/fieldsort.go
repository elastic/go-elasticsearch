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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldsortnumerictype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// FieldSort type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/sort.ts#L43-L52
type FieldSort struct {
	Format       *string                                    `json:"format,omitempty"`
	Missing      *Missing                                   `json:"missing,omitempty"`
	Mode         *sortmode.SortMode                         `json:"mode,omitempty"`
	Nested       *NestedSortValue                           `json:"nested,omitempty"`
	NumericType  *fieldsortnumerictype.FieldSortNumericType `json:"numeric_type,omitempty"`
	Order        *sortorder.SortOrder                       `json:"order,omitempty"`
	UnmappedType *fieldtype.FieldType                       `json:"unmapped_type,omitempty"`
}

// FieldSortBuilder holds FieldSort struct and provides a builder API.
type FieldSortBuilder struct {
	v *FieldSort
}

// NewFieldSort provides a builder for the FieldSort struct.
func NewFieldSortBuilder() *FieldSortBuilder {
	r := FieldSortBuilder{
		&FieldSort{},
	}

	return &r
}

// Build finalize the chain and returns the FieldSort struct
func (rb *FieldSortBuilder) Build() FieldSort {
	return *rb.v
}

func (rb *FieldSortBuilder) Format(format string) *FieldSortBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *FieldSortBuilder) Missing(missing *MissingBuilder) *FieldSortBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *FieldSortBuilder) Mode(mode sortmode.SortMode) *FieldSortBuilder {
	rb.v.Mode = &mode
	return rb
}

func (rb *FieldSortBuilder) Nested(nested *NestedSortValueBuilder) *FieldSortBuilder {
	v := nested.Build()
	rb.v.Nested = &v
	return rb
}

func (rb *FieldSortBuilder) NumericType(numerictype fieldsortnumerictype.FieldSortNumericType) *FieldSortBuilder {
	rb.v.NumericType = &numerictype
	return rb
}

func (rb *FieldSortBuilder) Order(order sortorder.SortOrder) *FieldSortBuilder {
	rb.v.Order = &order
	return rb
}

func (rb *FieldSortBuilder) UnmappedType(unmappedtype fieldtype.FieldType) *FieldSortBuilder {
	rb.v.UnmappedType = &unmappedtype
	return rb
}
