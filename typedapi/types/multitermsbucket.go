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

// MultiTermsBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L452-L456
type MultiTermsBucket struct {
	Aggregations            map[AggregateName]Aggregate `json:"-"`
	DocCount                int64                       `json:"doc_count"`
	DocCountErrorUpperBound *int64                      `json:"doc_count_error_upper_bound,omitempty"`
	Key                     []string                    `json:"key"`
	KeyAsString             *string                     `json:"key_as_string,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s MultiTermsBucket) MarshalJSON() ([]byte, error) {
	type opt MultiTermsBucket
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

// MultiTermsBucketBuilder holds MultiTermsBucket struct and provides a builder API.
type MultiTermsBucketBuilder struct {
	v *MultiTermsBucket
}

// NewMultiTermsBucket provides a builder for the MultiTermsBucket struct.
func NewMultiTermsBucketBuilder() *MultiTermsBucketBuilder {
	r := MultiTermsBucketBuilder{
		&MultiTermsBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the MultiTermsBucket struct
func (rb *MultiTermsBucketBuilder) Build() MultiTermsBucket {
	return *rb.v
}

func (rb *MultiTermsBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *MultiTermsBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *MultiTermsBucketBuilder) DocCount(doccount int64) *MultiTermsBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *MultiTermsBucketBuilder) DocCountErrorUpperBound(doccounterrorupperbound int64) *MultiTermsBucketBuilder {
	rb.v.DocCountErrorUpperBound = &doccounterrorupperbound
	return rb
}

func (rb *MultiTermsBucketBuilder) Key(key ...string) *MultiTermsBucketBuilder {
	rb.v.Key = key
	return rb
}

func (rb *MultiTermsBucketBuilder) KeyAsString(keyasstring string) *MultiTermsBucketBuilder {
	rb.v.KeyAsString = &keyasstring
	return rb
}
