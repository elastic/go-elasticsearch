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

// DoubleTermsBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L405-L408
type DoubleTermsBucket struct {
	Aggregations  map[AggregateName]Aggregate `json:"-"`
	DocCount      int64                       `json:"doc_count"`
	DocCountError *int64                      `json:"doc_count_error,omitempty"`
	Key           float64                     `json:"key"`
	KeyAsString   *string                     `json:"key_as_string,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s DoubleTermsBucket) MarshalJSON() ([]byte, error) {
	type opt DoubleTermsBucket
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

// DoubleTermsBucketBuilder holds DoubleTermsBucket struct and provides a builder API.
type DoubleTermsBucketBuilder struct {
	v *DoubleTermsBucket
}

// NewDoubleTermsBucket provides a builder for the DoubleTermsBucket struct.
func NewDoubleTermsBucketBuilder() *DoubleTermsBucketBuilder {
	r := DoubleTermsBucketBuilder{
		&DoubleTermsBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DoubleTermsBucket struct
func (rb *DoubleTermsBucketBuilder) Build() DoubleTermsBucket {
	return *rb.v
}

func (rb *DoubleTermsBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *DoubleTermsBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *DoubleTermsBucketBuilder) DocCount(doccount int64) *DoubleTermsBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *DoubleTermsBucketBuilder) DocCountError(doccounterror int64) *DoubleTermsBucketBuilder {
	rb.v.DocCountError = &doccounterror
	return rb
}

func (rb *DoubleTermsBucketBuilder) Key(key float64) *DoubleTermsBucketBuilder {
	rb.v.Key = key
	return rb
}

func (rb *DoubleTermsBucketBuilder) KeyAsString(keyasstring string) *DoubleTermsBucketBuilder {
	rb.v.KeyAsString = &keyasstring
	return rb
}
