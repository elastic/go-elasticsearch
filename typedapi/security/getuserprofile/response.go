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

package getuserprofile

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package getuserprofile
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/get_user_profile/Response.ts#L23-L33
type Response struct {
	Errors *types.GetUserProfileErrors `json:"errors,omitempty"`
	// Profiles A successful call returns the JSON representation of the user profile and its
	// internal versioning numbers.
	// The API returns an empty object if no profile document is found for the
	// provided `uid`.
	// The content of the data field is not returned by default to avoid
	// deserializing a potential large payload.
	Profiles []types.UserProfileWithMetadata `json:"profiles"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
