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

// VariableWidthHistogramBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L353-L360
type VariableWidthHistogramBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"-"`
	DocCount     int64                       `json:"doc_count"`
	Key          float64                     `json:"key"`
	KeyAsString  *string                     `json:"key_as_string,omitempty"`
	Max          float64                     `json:"max"`
	MaxAsString  *string                     `json:"max_as_string,omitempty"`
	Min          float64                     `json:"min"`
	MinAsString  *string                     `json:"min_as_string,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s VariableWidthHistogramBucket) MarshalJSON() ([]byte, error) {
	type opt VariableWidthHistogramBucket
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

// VariableWidthHistogramBucketBuilder holds VariableWidthHistogramBucket struct and provides a builder API.
type VariableWidthHistogramBucketBuilder struct {
	v *VariableWidthHistogramBucket
}

// NewVariableWidthHistogramBucket provides a builder for the VariableWidthHistogramBucket struct.
func NewVariableWidthHistogramBucketBuilder() *VariableWidthHistogramBucketBuilder {
	r := VariableWidthHistogramBucketBuilder{
		&VariableWidthHistogramBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the VariableWidthHistogramBucket struct
func (rb *VariableWidthHistogramBucketBuilder) Build() VariableWidthHistogramBucket {
	return *rb.v
}

func (rb *VariableWidthHistogramBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *VariableWidthHistogramBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *VariableWidthHistogramBucketBuilder) DocCount(doccount int64) *VariableWidthHistogramBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *VariableWidthHistogramBucketBuilder) Key(key float64) *VariableWidthHistogramBucketBuilder {
	rb.v.Key = key
	return rb
}

func (rb *VariableWidthHistogramBucketBuilder) KeyAsString(keyasstring string) *VariableWidthHistogramBucketBuilder {
	rb.v.KeyAsString = &keyasstring
	return rb
}

func (rb *VariableWidthHistogramBucketBuilder) Max(max float64) *VariableWidthHistogramBucketBuilder {
	rb.v.Max = max
	return rb
}

func (rb *VariableWidthHistogramBucketBuilder) MaxAsString(maxasstring string) *VariableWidthHistogramBucketBuilder {
	rb.v.MaxAsString = &maxasstring
	return rb
}

func (rb *VariableWidthHistogramBucketBuilder) Min(min float64) *VariableWidthHistogramBucketBuilder {
	rb.v.Min = min
	return rb
}

func (rb *VariableWidthHistogramBucketBuilder) MinAsString(minasstring string) *VariableWidthHistogramBucketBuilder {
	rb.v.MinAsString = &minasstring
	return rb
}
