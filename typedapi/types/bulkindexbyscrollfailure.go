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

// BulkIndexByScrollFailure type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Errors.ts#L58-L64
type BulkIndexByScrollFailure struct {
	Cause  ErrorCause `json:"cause"`
	Id     Id         `json:"id"`
	Index  IndexName  `json:"index"`
	Status int        `json:"status"`
	Type   string     `json:"type"`
}

// BulkIndexByScrollFailureBuilder holds BulkIndexByScrollFailure struct and provides a builder API.
type BulkIndexByScrollFailureBuilder struct {
	v *BulkIndexByScrollFailure
}

// NewBulkIndexByScrollFailure provides a builder for the BulkIndexByScrollFailure struct.
func NewBulkIndexByScrollFailureBuilder() *BulkIndexByScrollFailureBuilder {
	r := BulkIndexByScrollFailureBuilder{
		&BulkIndexByScrollFailure{},
	}

	return &r
}

// Build finalize the chain and returns the BulkIndexByScrollFailure struct
func (rb *BulkIndexByScrollFailureBuilder) Build() BulkIndexByScrollFailure {
	return *rb.v
}

func (rb *BulkIndexByScrollFailureBuilder) Cause(cause *ErrorCauseBuilder) *BulkIndexByScrollFailureBuilder {
	v := cause.Build()
	rb.v.Cause = v
	return rb
}

func (rb *BulkIndexByScrollFailureBuilder) Id(id Id) *BulkIndexByScrollFailureBuilder {
	rb.v.Id = id
	return rb
}

func (rb *BulkIndexByScrollFailureBuilder) Index(index IndexName) *BulkIndexByScrollFailureBuilder {
	rb.v.Index = index
	return rb
}

func (rb *BulkIndexByScrollFailureBuilder) Status(status int) *BulkIndexByScrollFailureBuilder {
	rb.v.Status = status
	return rb
}

func (rb *BulkIndexByScrollFailureBuilder) Type_(type_ string) *BulkIndexByScrollFailureBuilder {
	rb.v.Type = type_
	return rb
}
