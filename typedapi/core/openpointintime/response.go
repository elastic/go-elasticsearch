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
// https://github.com/elastic/elasticsearch-specification/tree/0f6f3696eb685db8b944feefb6a209ad7e385b9c

package openpointintime

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package openpointintime
//
// https://github.com/elastic/elasticsearch-specification/blob/0f6f3696eb685db8b944feefb6a209ad7e385b9c/specification/_global/open_point_in_time/OpenPointInTimeResponse.ts#L23-L29
type Response struct {
	Id string `json:"id"`
	// Shards_ Shards used to create the PIT
	Shards_ types.ShardStatistics `json:"_shards"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
