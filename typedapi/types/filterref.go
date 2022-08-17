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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/filtertype"
)

// FilterRef type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Filter.ts#L31-L41
type FilterRef struct {
	// FilterId The identifier for the filter.
	FilterId Id `json:"filter_id"`
	// FilterType If set to `include`, the rule applies for values in the filter. If set to
	// `exclude`, the rule applies for values not in the filter.
	FilterType *filtertype.FilterType `json:"filter_type,omitempty"`
}

// FilterRefBuilder holds FilterRef struct and provides a builder API.
type FilterRefBuilder struct {
	v *FilterRef
}

// NewFilterRef provides a builder for the FilterRef struct.
func NewFilterRefBuilder() *FilterRefBuilder {
	r := FilterRefBuilder{
		&FilterRef{},
	}

	return &r
}

// Build finalize the chain and returns the FilterRef struct
func (rb *FilterRefBuilder) Build() FilterRef {
	return *rb.v
}

// FilterId The identifier for the filter.

func (rb *FilterRefBuilder) FilterId(filterid Id) *FilterRefBuilder {
	rb.v.FilterId = filterid
	return rb
}

// FilterType If set to `include`, the rule applies for values in the filter. If set to
// `exclude`, the rule applies for values not in the filter.

func (rb *FilterRefBuilder) FilterType(filtertype filtertype.FilterType) *FilterRefBuilder {
	rb.v.FilterType = &filtertype
	return rb
}
