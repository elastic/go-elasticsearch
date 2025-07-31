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

package getbuiltinprivileges

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/remoteclusterprivilege"
)

// Response holds the response body struct for the package getbuiltinprivileges
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/security/get_builtin_privileges/SecurityGetBuiltinPrivilegesResponse.ts#L26-L42
type Response struct {

	// Cluster The list of cluster privileges that are understood by this version of
	// Elasticsearch.
	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster"`
	// Index The list of index privileges that are understood by this version of
	// Elasticsearch.
	Index []string `json:"index"`
	// RemoteCluster The list of remote_cluster privileges that are understood by this version of
	// Elasticsearch.
	RemoteCluster []remoteclusterprivilege.RemoteClusterPrivilege `json:"remote_cluster"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
