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

// Breaker type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/nodes/_types/Stats.ts#L173-L180
type Breaker struct {
	EstimatedSize        *string  `json:"estimated_size,omitempty"`
	EstimatedSizeInBytes *int64   `json:"estimated_size_in_bytes,omitempty"`
	LimitSize            *string  `json:"limit_size,omitempty"`
	LimitSizeInBytes     *int64   `json:"limit_size_in_bytes,omitempty"`
	Overhead             *float32 `json:"overhead,omitempty"`
	Tripped              *float32 `json:"tripped,omitempty"`
}

// NewBreaker returns a Breaker.
func NewBreaker() *Breaker {
	r := &Breaker{}

	return r
}
