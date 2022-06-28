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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// LongRareTermsBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/Aggregate.ts#L425-L428
type LongRareTermsBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
	Key          int64                       `json:"key"`
	KeyAsString  *string                     `json:"key_as_string,omitempty"`
}

// LongRareTermsBucketBuilder holds LongRareTermsBucket struct and provides a builder API.
type LongRareTermsBucketBuilder struct {
	v *LongRareTermsBucket
}

// NewLongRareTermsBucket provides a builder for the LongRareTermsBucket struct.
func NewLongRareTermsBucketBuilder() *LongRareTermsBucketBuilder {
	r := LongRareTermsBucketBuilder{
		&LongRareTermsBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the LongRareTermsBucket struct
func (rb *LongRareTermsBucketBuilder) Build() LongRareTermsBucket {
	return *rb.v
}

func (rb *LongRareTermsBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *LongRareTermsBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *LongRareTermsBucketBuilder) DocCount(doccount int64) *LongRareTermsBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *LongRareTermsBucketBuilder) Key(key int64) *LongRareTermsBucketBuilder {
	rb.v.Key = key
	return rb
}

func (rb *LongRareTermsBucketBuilder) KeyAsString(keyasstring string) *LongRareTermsBucketBuilder {
	rb.v.KeyAsString = &keyasstring
	return rb
}
