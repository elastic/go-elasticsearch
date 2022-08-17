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
// https://github.com/elastic/elasticsearch-specification/tree/e0ea3dc890d394d682096cc862b3bd879d9422e9


package types

import (
	"encoding/json"
)

// FilterAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e0ea3dc890d394d682096cc862b3bd879d9422e9/specification/_types/aggregations/Aggregate.ts#L482-L483
type FilterAggregate struct {
	Aggregations map[AggregateName]Aggregate `json:"-"`
	DocCount     int64                       `json:"doc_count"`
	Meta         *Metadata                   `json:"meta,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s FilterAggregate) MarshalJSON() ([]byte, error) {
	type opt FilterAggregate
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.Aggregations {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// FilterAggregateBuilder holds FilterAggregate struct and provides a builder API.
type FilterAggregateBuilder struct {
	v *FilterAggregate
}

// NewFilterAggregate provides a builder for the FilterAggregate struct.
func NewFilterAggregateBuilder() *FilterAggregateBuilder {
	r := FilterAggregateBuilder{
		&FilterAggregate{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the FilterAggregate struct
func (rb *FilterAggregateBuilder) Build() FilterAggregate {
	return *rb.v
}

func (rb *FilterAggregateBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *FilterAggregateBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *FilterAggregateBuilder) DocCount(doccount int64) *FilterAggregateBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *FilterAggregateBuilder) Meta(meta *MetadataBuilder) *FilterAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
