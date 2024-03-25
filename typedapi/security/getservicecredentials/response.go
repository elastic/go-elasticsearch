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
// https://github.com/elastic/elasticsearch-specification/tree/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757

package getservicecredentials

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package getservicecredentials
//
// https://github.com/elastic/elasticsearch-specification/blob/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757/specification/security/get_service_credentials/GetServiceCredentialsResponse.ts#L25-L33
type Response struct {
	Count int `json:"count"`
	// NodesCredentials Contains service account credentials collected from all nodes of the cluster
	NodesCredentials types.NodesCredentials    `json:"nodes_credentials"`
	ServiceAccount   string                    `json:"service_account"`
	Tokens           map[string]types.Metadata `json:"tokens"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Tokens: make(map[string]types.Metadata, 0),
	}
	return r
}
