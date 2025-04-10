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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package updateapikey

// Response holds the response body struct for the package updateapikey
//
// https://github.com/elastic/elasticsearch-specification/blob/beeb1dc688bcc058488dcc45d9cbd2cd364e9943/specification/security/update_api_key/Response.ts#L20-L28
type Response struct {

	// Updated If `true`, the API key was updated.
	// If `false`, the API key didn't change because no change was detected.
	Updated bool `json:"updated"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
