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

// MultiBucketBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/Aggregate.ts#L318-L325
type MultiBucketBase struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
}

// MultiBucketBaseBuilder holds MultiBucketBase struct and provides a builder API.
type MultiBucketBaseBuilder struct {
	v *MultiBucketBase
}

// NewMultiBucketBase provides a builder for the MultiBucketBase struct.
func NewMultiBucketBaseBuilder() *MultiBucketBaseBuilder {
	r := MultiBucketBaseBuilder{
		&MultiBucketBase{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the MultiBucketBase struct
func (rb *MultiBucketBaseBuilder) Build() MultiBucketBase {
	return *rb.v
}

func (rb *MultiBucketBaseBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *MultiBucketBaseBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *MultiBucketBaseBuilder) DocCount(doccount int64) *MultiBucketBaseBuilder {
	rb.v.DocCount = doccount
	return rb
}
