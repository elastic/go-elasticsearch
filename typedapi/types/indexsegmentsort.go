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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/segmentsortmissing"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/segmentsortmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/segmentsortorder"
)

// IndexSegmentSort type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSegmentSort.ts#L22-L27
type IndexSegmentSort struct {
	Field   *Fields                                 `json:"field,omitempty"`
	Missing []segmentsortmissing.SegmentSortMissing `json:"missing,omitempty"`
	Mode    []segmentsortmode.SegmentSortMode       `json:"mode,omitempty"`
	Order   []segmentsortorder.SegmentSortOrder     `json:"order,omitempty"`
}

// IndexSegmentSortBuilder holds IndexSegmentSort struct and provides a builder API.
type IndexSegmentSortBuilder struct {
	v *IndexSegmentSort
}

// NewIndexSegmentSort provides a builder for the IndexSegmentSort struct.
func NewIndexSegmentSortBuilder() *IndexSegmentSortBuilder {
	r := IndexSegmentSortBuilder{
		&IndexSegmentSort{},
	}

	return &r
}

// Build finalize the chain and returns the IndexSegmentSort struct
func (rb *IndexSegmentSortBuilder) Build() IndexSegmentSort {
	return *rb.v
}

func (rb *IndexSegmentSortBuilder) Field(field *FieldsBuilder) *IndexSegmentSortBuilder {
	v := field.Build()
	rb.v.Field = &v
	return rb
}

func (rb *IndexSegmentSortBuilder) Missing(arg []segmentsortmissing.SegmentSortMissing) *IndexSegmentSortBuilder {
	rb.v.Missing = arg
	return rb
}

func (rb *IndexSegmentSortBuilder) Mode(arg []segmentsortmode.SegmentSortMode) *IndexSegmentSortBuilder {
	rb.v.Mode = arg
	return rb
}

func (rb *IndexSegmentSortBuilder) Order(arg []segmentsortorder.SegmentSortOrder) *IndexSegmentSortBuilder {
	rb.v.Order = arg
	return rb
}
