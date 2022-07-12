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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// PhraseSuggestCollate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_global/search/_types/suggester.ts#L175-L179
type PhraseSuggestCollate struct {
	Params map[string]interface{}    `json:"params,omitempty"`
	Prune  *bool                     `json:"prune,omitempty"`
	Query  PhraseSuggestCollateQuery `json:"query"`
}

// PhraseSuggestCollateBuilder holds PhraseSuggestCollate struct and provides a builder API.
type PhraseSuggestCollateBuilder struct {
	v *PhraseSuggestCollate
}

// NewPhraseSuggestCollate provides a builder for the PhraseSuggestCollate struct.
func NewPhraseSuggestCollateBuilder() *PhraseSuggestCollateBuilder {
	r := PhraseSuggestCollateBuilder{
		&PhraseSuggestCollate{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the PhraseSuggestCollate struct
func (rb *PhraseSuggestCollateBuilder) Build() PhraseSuggestCollate {
	return *rb.v
}

func (rb *PhraseSuggestCollateBuilder) Params(value map[string]interface{}) *PhraseSuggestCollateBuilder {
	rb.v.Params = value
	return rb
}

func (rb *PhraseSuggestCollateBuilder) Prune(prune bool) *PhraseSuggestCollateBuilder {
	rb.v.Prune = &prune
	return rb
}

func (rb *PhraseSuggestCollateBuilder) Query(query *PhraseSuggestCollateQueryBuilder) *PhraseSuggestCollateBuilder {
	v := query.Build()
	rb.v.Query = v
	return rb
}
