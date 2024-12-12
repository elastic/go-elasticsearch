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
// https://github.com/elastic/elasticsearch-specification/tree/1ed5f4795fc7c4d9875601f883b8d5fb9023c526

package getfeatureupgradestatus

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/migrationstatus"
)

// Response holds the response body struct for the package getfeatureupgradestatus
//
// https://github.com/elastic/elasticsearch-specification/blob/1ed5f4795fc7c4d9875601f883b8d5fb9023c526/specification/migration/get_feature_upgrade_status/GetFeatureUpgradeStatusResponse.ts#L23-L28
type Response struct {
	Features        []types.GetMigrationFeature     `json:"features"`
	MigrationStatus migrationstatus.MigrationStatus `json:"migration_status"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
