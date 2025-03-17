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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package termsenum

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package termsenum
//
// https://github.com/elastic/elasticsearch-specification/blob/c75a0abec670d027d13eb8d6f23374f86621c76b/specification/_global/terms_enum/TermsEnumResponse.ts#L22-L32
type Response struct {

	// Complete If `false`, the returned terms set may be incomplete and should be treated as
	// approximate.
	// This can occur due to a few reasons, such as a request timeout or a node
	// error.
	Complete bool                  `json:"complete"`
	Shards_  types.ShardStatistics `json:"_shards"`
	Terms    []string              `json:"terms"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
