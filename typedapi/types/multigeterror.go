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

// MultiGetError type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/mget/types.ts#L62-L66
type MultiGetError struct {
	Error  ErrorCause `json:"error"`
	Id_    Id         `json:"_id"`
	Index_ IndexName  `json:"_index"`
}

// MultiGetErrorBuilder holds MultiGetError struct and provides a builder API.
type MultiGetErrorBuilder struct {
	v *MultiGetError
}

// NewMultiGetError provides a builder for the MultiGetError struct.
func NewMultiGetErrorBuilder() *MultiGetErrorBuilder {
	r := MultiGetErrorBuilder{
		&MultiGetError{},
	}

	return &r
}

// Build finalize the chain and returns the MultiGetError struct
func (rb *MultiGetErrorBuilder) Build() MultiGetError {
	return *rb.v
}

func (rb *MultiGetErrorBuilder) Error(error *ErrorCauseBuilder) *MultiGetErrorBuilder {
	v := error.Build()
	rb.v.Error = v
	return rb
}

func (rb *MultiGetErrorBuilder) Id_(id_ Id) *MultiGetErrorBuilder {
	rb.v.Id_ = id_
	return rb
}

func (rb *MultiGetErrorBuilder) Index_(index_ IndexName) *MultiGetErrorBuilder {
	rb.v.Index_ = index_
	return rb
}
