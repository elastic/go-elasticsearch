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
// https://github.com/elastic/elasticsearch-specification/tree/ec3159eb31c62611202a4fb157ea88fa6ff78e1a


package types

// FiltersAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/_types/aggregations/bucket.ts#L169-L174
type FiltersAggregation struct {
	Filters        *BucketsQuery          `json:"filters,omitempty"`
	Keyed          *bool                  `json:"keyed,omitempty"`
	Meta           map[string]interface{} `json:"meta,omitempty"`
	Name           *string                `json:"name,omitempty"`
	OtherBucket    *bool                  `json:"other_bucket,omitempty"`
	OtherBucketKey *string                `json:"other_bucket_key,omitempty"`
}

// NewFiltersAggregation returns a FiltersAggregation.
func NewFiltersAggregation() *FiltersAggregation {
	r := &FiltersAggregation{}

	return r
}
