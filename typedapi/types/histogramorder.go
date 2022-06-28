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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// HistogramOrder type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/bucket.ts#L245-L248
type HistogramOrder struct {
	Count_ *sortorder.SortOrder `json:"_count,omitempty"`
	Key_   *sortorder.SortOrder `json:"_key,omitempty"`
}

// HistogramOrderBuilder holds HistogramOrder struct and provides a builder API.
type HistogramOrderBuilder struct {
	v *HistogramOrder
}

// NewHistogramOrder provides a builder for the HistogramOrder struct.
func NewHistogramOrderBuilder() *HistogramOrderBuilder {
	r := HistogramOrderBuilder{
		&HistogramOrder{},
	}

	return &r
}

// Build finalize the chain and returns the HistogramOrder struct
func (rb *HistogramOrderBuilder) Build() HistogramOrder {
	return *rb.v
}

func (rb *HistogramOrderBuilder) Count_(count_ sortorder.SortOrder) *HistogramOrderBuilder {
	rb.v.Count_ = &count_
	return rb
}

func (rb *HistogramOrderBuilder) Key_(key_ sortorder.SortOrder) *HistogramOrderBuilder {
	rb.v.Key_ = &key_
	return rb
}
