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

// VerifyIndex type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/recovery/types.ts#L111-L116
type VerifyIndex struct {
	CheckIndexTime         *Duration               `json:"check_index_time,omitempty"`
	CheckIndexTimeInMillis DurationValueUnitMillis `json:"check_index_time_in_millis"`
	TotalTime              *Duration               `json:"total_time,omitempty"`
	TotalTimeInMillis      DurationValueUnitMillis `json:"total_time_in_millis"`
}

// VerifyIndexBuilder holds VerifyIndex struct and provides a builder API.
type VerifyIndexBuilder struct {
	v *VerifyIndex
}

// NewVerifyIndex provides a builder for the VerifyIndex struct.
func NewVerifyIndexBuilder() *VerifyIndexBuilder {
	r := VerifyIndexBuilder{
		&VerifyIndex{},
	}

	return &r
}

// Build finalize the chain and returns the VerifyIndex struct
func (rb *VerifyIndexBuilder) Build() VerifyIndex {
	return *rb.v
}

func (rb *VerifyIndexBuilder) CheckIndexTime(checkindextime *DurationBuilder) *VerifyIndexBuilder {
	v := checkindextime.Build()
	rb.v.CheckIndexTime = &v
	return rb
}

func (rb *VerifyIndexBuilder) CheckIndexTimeInMillis(checkindextimeinmillis *DurationValueUnitMillisBuilder) *VerifyIndexBuilder {
	v := checkindextimeinmillis.Build()
	rb.v.CheckIndexTimeInMillis = v
	return rb
}

func (rb *VerifyIndexBuilder) TotalTime(totaltime *DurationBuilder) *VerifyIndexBuilder {
	v := totaltime.Build()
	rb.v.TotalTime = &v
	return rb
}

func (rb *VerifyIndexBuilder) TotalTimeInMillis(totaltimeinmillis *DurationValueUnitMillisBuilder) *VerifyIndexBuilder {
	v := totaltimeinmillis.Build()
	rb.v.TotalTimeInMillis = v
	return rb
}
