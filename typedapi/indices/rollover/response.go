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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package rollover

// Response holds the response body struct for the package rollover
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/indices/rollover/IndicesRolloverResponse.ts#L22-L32

type Response struct {
	Acknowledged       bool            `json:"acknowledged"`
	Conditions         map[string]bool `json:"conditions"`
	DryRun             bool            `json:"dry_run"`
	NewIndex           string          `json:"new_index"`
	OldIndex           string          `json:"old_index"`
	RolledOver         bool            `json:"rolled_over"`
	ShardsAcknowledged bool            `json:"shards_acknowledged"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Conditions: make(map[string]bool, 0),
	}
	return r
}
