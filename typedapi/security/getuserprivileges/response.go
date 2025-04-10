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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package getuserprivileges

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package getuserprivileges
//
// https://github.com/elastic/elasticsearch-specification/blob/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9/specification/security/get_user_privileges/SecurityGetUserPrivilegesResponse.ts#L28-L38
type Response struct {
	Applications  []types.ApplicationPrivileges       `json:"applications"`
	Cluster       []string                            `json:"cluster"`
	Global        []types.GlobalPrivilege             `json:"global"`
	Indices       []types.UserIndicesPrivileges       `json:"indices"`
	RemoteCluster []types.RemoteClusterPrivileges     `json:"remote_cluster,omitempty"`
	RemoteIndices []types.RemoteUserIndicesPrivileges `json:"remote_indices,omitempty"`
	RunAs         []string                            `json:"run_as"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
