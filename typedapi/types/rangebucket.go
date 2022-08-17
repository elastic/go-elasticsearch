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

// RangeBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L514-L521
type RangeBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"-"`
	DocCount     int64                       `json:"doc_count"`
	From         *float64                    `json:"from,omitempty"`
	FromAsString *string                     `json:"from_as_string,omitempty"`
	// Key The bucket key. Present if the aggregation is _not_ keyed
	Key        *string  `json:"key,omitempty"`
	To         *float64 `json:"to,omitempty"`
	ToAsString *string  `json:"to_as_string,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s RangeBucket) MarshalJSON() ([]byte, error) {
	type opt RangeBucket
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

// RangeBucketBuilder holds RangeBucket struct and provides a builder API.
type RangeBucketBuilder struct {
	v *RangeBucket
}

// NewRangeBucket provides a builder for the RangeBucket struct.
func NewRangeBucketBuilder() *RangeBucketBuilder {
	r := RangeBucketBuilder{
		&RangeBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the RangeBucket struct
func (rb *RangeBucketBuilder) Build() RangeBucket {
	return *rb.v
}

func (rb *RangeBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *RangeBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *RangeBucketBuilder) DocCount(doccount int64) *RangeBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *RangeBucketBuilder) From(from float64) *RangeBucketBuilder {
	rb.v.From = &from
	return rb
}

func (rb *RangeBucketBuilder) FromAsString(fromasstring string) *RangeBucketBuilder {
	rb.v.FromAsString = &fromasstring
	return rb
}

// Key The bucket key. Present if the aggregation is _not_ keyed

func (rb *RangeBucketBuilder) Key(key string) *RangeBucketBuilder {
	rb.v.Key = &key
	return rb
}

func (rb *RangeBucketBuilder) To(to float64) *RangeBucketBuilder {
	rb.v.To = &to
	return rb
}

func (rb *RangeBucketBuilder) ToAsString(toasstring string) *RangeBucketBuilder {
	rb.v.ToAsString = &toasstring
	return rb
}
