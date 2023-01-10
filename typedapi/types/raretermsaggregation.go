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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// RareTermsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/aggregations/bucket.ts#L304-L312
type RareTermsAggregation struct {
	Exclude     []string               `json:"exclude,omitempty"`
	Field       *string                `json:"field,omitempty"`
	Include     *TermsInclude          `json:"include,omitempty"`
	MaxDocCount *int64                 `json:"max_doc_count,omitempty"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
	Missing     *Missing               `json:"missing,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Precision   *float64               `json:"precision,omitempty"`
	ValueType   *string                `json:"value_type,omitempty"`
}

// NewRareTermsAggregation returns a RareTermsAggregation.
func NewRareTermsAggregation() *RareTermsAggregation {
	r := &RareTermsAggregation{}

	return r
}
