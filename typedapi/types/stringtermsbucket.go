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

// StringTermsBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/Aggregate.ts#L382-L384
type StringTermsBucket struct {
	Aggregations  map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount      int64                       `json:"doc_count"`
	DocCountError *int64                      `json:"doc_count_error,omitempty"`
	Key           string                      `json:"key"`
}

// StringTermsBucketBuilder holds StringTermsBucket struct and provides a builder API.
type StringTermsBucketBuilder struct {
	v *StringTermsBucket
}

// NewStringTermsBucket provides a builder for the StringTermsBucket struct.
func NewStringTermsBucketBuilder() *StringTermsBucketBuilder {
	r := StringTermsBucketBuilder{
		&StringTermsBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the StringTermsBucket struct
func (rb *StringTermsBucketBuilder) Build() StringTermsBucket {
	return *rb.v
}

func (rb *StringTermsBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *StringTermsBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *StringTermsBucketBuilder) DocCount(doccount int64) *StringTermsBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *StringTermsBucketBuilder) DocCountError(doccounterror int64) *StringTermsBucketBuilder {
	rb.v.DocCountError = &doccounterror
	return rb
}

func (rb *StringTermsBucketBuilder) Key(key string) *StringTermsBucketBuilder {
	rb.v.Key = key
	return rb
}
