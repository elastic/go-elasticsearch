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

// SearchProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/profile.ts#L124-L128
type SearchProfile struct {
	Collector   []Collector    `json:"collector"`
	Query       []QueryProfile `json:"query"`
	RewriteTime int64          `json:"rewrite_time"`
}

// SearchProfileBuilder holds SearchProfile struct and provides a builder API.
type SearchProfileBuilder struct {
	v *SearchProfile
}

// NewSearchProfile provides a builder for the SearchProfile struct.
func NewSearchProfileBuilder() *SearchProfileBuilder {
	r := SearchProfileBuilder{
		&SearchProfile{},
	}

	return &r
}

// Build finalize the chain and returns the SearchProfile struct
func (rb *SearchProfileBuilder) Build() SearchProfile {
	return *rb.v
}

func (rb *SearchProfileBuilder) Collector(collector []CollectorBuilder) *SearchProfileBuilder {
	tmp := make([]Collector, len(collector))
	for _, value := range collector {
		tmp = append(tmp, value.Build())
	}
	rb.v.Collector = tmp
	return rb
}

func (rb *SearchProfileBuilder) Query(query []QueryProfileBuilder) *SearchProfileBuilder {
	tmp := make([]QueryProfile, len(query))
	for _, value := range query {
		tmp = append(tmp, value.Build())
	}
	rb.v.Query = tmp
	return rb
}

func (rb *SearchProfileBuilder) RewriteTime(rewritetime int64) *SearchProfileBuilder {
	rb.v.RewriteTime = rewritetime
	return rb
}
