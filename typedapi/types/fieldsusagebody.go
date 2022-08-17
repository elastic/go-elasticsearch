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
	"encoding/json"
)

// FieldsUsageBody type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/field_usage_stats/IndicesFieldUsageStatsResponse.ts#L32-L36
type FieldsUsageBody struct {
	FieldsUsageBody map[IndexName]UsageStatsIndex `json:"-"`
	Shards_         ShardStatistics               `json:"_shards"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s FieldsUsageBody) MarshalJSON() ([]byte, error) {
	type opt FieldsUsageBody
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
	for key, value := range s.FieldsUsageBody {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// FieldsUsageBodyBuilder holds FieldsUsageBody struct and provides a builder API.
type FieldsUsageBodyBuilder struct {
	v *FieldsUsageBody
}

// NewFieldsUsageBody provides a builder for the FieldsUsageBody struct.
func NewFieldsUsageBodyBuilder() *FieldsUsageBodyBuilder {
	r := FieldsUsageBodyBuilder{
		&FieldsUsageBody{
			FieldsUsageBody: make(map[IndexName]UsageStatsIndex, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the FieldsUsageBody struct
func (rb *FieldsUsageBodyBuilder) Build() FieldsUsageBody {
	return *rb.v
}

func (rb *FieldsUsageBodyBuilder) FieldsUsageBody(values map[IndexName]*UsageStatsIndexBuilder) *FieldsUsageBodyBuilder {
	tmp := make(map[IndexName]UsageStatsIndex, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.FieldsUsageBody = tmp
	return rb
}

func (rb *FieldsUsageBodyBuilder) Shards_(shards_ *ShardStatisticsBuilder) *FieldsUsageBodyBuilder {
	v := shards_.Build()
	rb.v.Shards_ = v
	return rb
}
