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

// SignificantStringTermsBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/Aggregate.ts#L576-L578
type SignificantStringTermsBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	BgCount      int64                       `json:"bg_count"`
	DocCount     int64                       `json:"doc_count"`
	Key          string                      `json:"key"`
	Score        float64                     `json:"score"`
}

// SignificantStringTermsBucketBuilder holds SignificantStringTermsBucket struct and provides a builder API.
type SignificantStringTermsBucketBuilder struct {
	v *SignificantStringTermsBucket
}

// NewSignificantStringTermsBucket provides a builder for the SignificantStringTermsBucket struct.
func NewSignificantStringTermsBucketBuilder() *SignificantStringTermsBucketBuilder {
	r := SignificantStringTermsBucketBuilder{
		&SignificantStringTermsBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SignificantStringTermsBucket struct
func (rb *SignificantStringTermsBucketBuilder) Build() SignificantStringTermsBucket {
	return *rb.v
}

func (rb *SignificantStringTermsBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *SignificantStringTermsBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *SignificantStringTermsBucketBuilder) BgCount(bgcount int64) *SignificantStringTermsBucketBuilder {
	rb.v.BgCount = bgcount
	return rb
}

func (rb *SignificantStringTermsBucketBuilder) DocCount(doccount int64) *SignificantStringTermsBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *SignificantStringTermsBucketBuilder) Key(key string) *SignificantStringTermsBucketBuilder {
	rb.v.Key = key
	return rb
}

func (rb *SignificantStringTermsBucketBuilder) Score(score float64) *SignificantStringTermsBucketBuilder {
	rb.v.Score = score
	return rb
}
