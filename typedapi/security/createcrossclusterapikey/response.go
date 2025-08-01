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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package createcrossclusterapikey

// Response holds the response body struct for the package createcrossclusterapikey
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/security/create_cross_cluster_api_key/CreateCrossClusterApiKeyResponse.ts#L23-L48
type Response struct {

	// ApiKey Generated API key.
	ApiKey string `json:"api_key"`
	// Encoded API key credentials which is the base64-encoding of
	// the UTF-8 representation of `id` and `api_key` joined
	// by a colon (`:`).
	Encoded string `json:"encoded"`
	// Expiration Expiration in milliseconds for the API key.
	Expiration *int64 `json:"expiration,omitempty"`
	// Id Unique ID for this API key.
	Id string `json:"id"`
	// Name Specifies the name for this API key.
	Name string `json:"name"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
