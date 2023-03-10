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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package invalidateapikey

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package invalidateapikey
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/security/invalidate_api_key/SecurityInvalidateApiKeyResponse.ts#L23-L30

type Response struct {
	ErrorCount                   int                `json:"error_count"`
	ErrorDetails                 []types.ErrorCause `json:"error_details,omitempty"`
	InvalidatedApiKeys           []string           `json:"invalidated_api_keys"`
	PreviouslyInvalidatedApiKeys []string           `json:"previously_invalidated_api_keys"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
