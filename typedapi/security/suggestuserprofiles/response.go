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

package suggestuserprofiles

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package suggestuserprofiles
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/suggest_user_profiles/Response.ts#L29-L44
type Response struct {

	// Profiles A list of profile documents, ordered by relevance, that match the search
	// criteria.
	Profiles []types.UserProfile `json:"profiles"`
	// Took The number of milliseconds it took Elasticsearch to run the request.
	Took int64 `json:"took"`
	// Total Metadata about the number of matching profiles.
	Total types.TotalUserProfiles `json:"total"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
