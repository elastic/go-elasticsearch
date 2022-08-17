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

// MatrixStatsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L708-L712
type MatrixStatsAggregate struct {
	DocCount int64               `json:"doc_count"`
	Fields   []MatrixStatsFields `json:"fields"`
	Meta     *Metadata           `json:"meta,omitempty"`
}

// MatrixStatsAggregateBuilder holds MatrixStatsAggregate struct and provides a builder API.
type MatrixStatsAggregateBuilder struct {
	v *MatrixStatsAggregate
}

// NewMatrixStatsAggregate provides a builder for the MatrixStatsAggregate struct.
func NewMatrixStatsAggregateBuilder() *MatrixStatsAggregateBuilder {
	r := MatrixStatsAggregateBuilder{
		&MatrixStatsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the MatrixStatsAggregate struct
func (rb *MatrixStatsAggregateBuilder) Build() MatrixStatsAggregate {
	return *rb.v
}

func (rb *MatrixStatsAggregateBuilder) DocCount(doccount int64) *MatrixStatsAggregateBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *MatrixStatsAggregateBuilder) Fields(fields []MatrixStatsFieldsBuilder) *MatrixStatsAggregateBuilder {
	tmp := make([]MatrixStatsFields, len(fields))
	for _, value := range fields {
		tmp = append(tmp, value.Build())
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *MatrixStatsAggregateBuilder) Meta(meta *MetadataBuilder) *MatrixStatsAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
