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

// FieldStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/termvectors/types.ts#L28-L32
type FieldStatistics struct {
	DocCount   int   `json:"doc_count"`
	SumDocFreq int64 `json:"sum_doc_freq"`
	SumTtf     int64 `json:"sum_ttf"`
}

// FieldStatisticsBuilder holds FieldStatistics struct and provides a builder API.
type FieldStatisticsBuilder struct {
	v *FieldStatistics
}

// NewFieldStatistics provides a builder for the FieldStatistics struct.
func NewFieldStatisticsBuilder() *FieldStatisticsBuilder {
	r := FieldStatisticsBuilder{
		&FieldStatistics{},
	}

	return &r
}

// Build finalize the chain and returns the FieldStatistics struct
func (rb *FieldStatisticsBuilder) Build() FieldStatistics {
	return *rb.v
}

func (rb *FieldStatisticsBuilder) DocCount(doccount int) *FieldStatisticsBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *FieldStatisticsBuilder) SumDocFreq(sumdocfreq int64) *FieldStatisticsBuilder {
	rb.v.SumDocFreq = sumdocfreq
	return rb
}

func (rb *FieldStatisticsBuilder) SumTtf(sumttf int64) *FieldStatisticsBuilder {
	rb.v.SumTtf = sumttf
	return rb
}
