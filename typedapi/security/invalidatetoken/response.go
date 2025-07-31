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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package invalidatetoken

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package invalidatetoken
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/invalidate_token/SecurityInvalidateTokenResponse.ts#L23-L46
type Response struct {

	// ErrorCount The number of errors that were encountered when invalidating the tokens.
	ErrorCount int64 `json:"error_count"`
	// ErrorDetails Details about the errors.
	// This field is not present in the response when `error_count` is `0`.
	ErrorDetails []types.ErrorCause `json:"error_details,omitempty"`
	// InvalidatedTokens The number of the tokens that were invalidated as part of this request.
	InvalidatedTokens int64 `json:"invalidated_tokens"`
	// PreviouslyInvalidatedTokens The number of tokens that were already invalidated.
	PreviouslyInvalidatedTokens int64 `json:"previously_invalidated_tokens"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
