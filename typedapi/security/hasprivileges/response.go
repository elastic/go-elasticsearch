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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package hasprivileges

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package hasprivileges
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/security/has_privileges/SecurityHasPrivilegesResponse.ts#L24-L32

type Response struct {
	Application     types.ApplicationsPrivileges `json:"application"`
	Cluster         map[string]bool              `json:"cluster"`
	HasAllRequested bool                         `json:"has_all_requested"`
	Index           map[string]types.Privileges  `json:"index"`
	Username        string                       `json:"username"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Cluster: make(map[string]bool, 0),
		Index:   make(map[string]types.Privileges, 0),
	}
	return r
}
