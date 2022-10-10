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
// https://github.com/elastic/elasticsearch-specification/tree/e0ea3dc890d394d682096cc862b3bd879d9422e9


package types

// ErrorResponseBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e0ea3dc890d394d682096cc862b3bd879d9422e9/specification/_types/Base.ts#L66-L75
type ErrorResponseBase struct {
	Error  ErrorCause `json:"error"`
	Status int        `json:"status"`
}

// ErrorResponseBaseBuilder holds ErrorResponseBase struct and provides a builder API.
type ErrorResponseBaseBuilder struct {
	v *ErrorResponseBase
}

// NewErrorResponseBase provides a builder for the ErrorResponseBase struct.
func NewErrorResponseBaseBuilder() *ErrorResponseBaseBuilder {
	r := ErrorResponseBaseBuilder{
		&ErrorResponseBase{},
	}

	return &r
}

// Build finalize the chain and returns the ErrorResponseBase struct
func (rb *ErrorResponseBaseBuilder) Build() ErrorResponseBase {
	return *rb.v
}

func (rb *ErrorResponseBaseBuilder) Error(error *ErrorCauseBuilder) *ErrorResponseBaseBuilder {
	v := error.Build()
	rb.v.Error = v
	return rb
}

func (rb *ErrorResponseBaseBuilder) Status(status int) *ErrorResponseBaseBuilder {
	rb.v.Status = status
	return rb
}