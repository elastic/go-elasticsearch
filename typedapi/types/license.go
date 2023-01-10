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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/licensetype"
)

// License type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/license/_types/License.ts#L42-L53
type License struct {
	ExpiryDateInMillis int64                   `json:"expiry_date_in_millis"`
	IssueDateInMillis  int64                   `json:"issue_date_in_millis"`
	IssuedTo           string                  `json:"issued_to"`
	Issuer             string                  `json:"issuer"`
	MaxNodes           int64                   `json:"max_nodes,omitempty"`
	MaxResourceUnits   *int64                  `json:"max_resource_units,omitempty"`
	Signature          string                  `json:"signature"`
	StartDateInMillis  *int64                  `json:"start_date_in_millis,omitempty"`
	Type               licensetype.LicenseType `json:"type"`
	Uid                string                  `json:"uid"`
}

// NewLicense returns a License.
func NewLicense() *License {
	r := &License{}

	return r
}
