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
// https://github.com/elastic/elasticsearch-specification/tree/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6

package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/datasourceprivilege"
)

// DataSourcePrivileges type.
//
// https://github.com/elastic/elasticsearch-specification/blob/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6/specification/security/_types/Privileges.ts#L435-L444
type DataSourcePrivileges struct {
	// Names A list of data source names or wildcard patterns to which the permissions in
	// this entry apply.
	Names []string `json:"names"`
	// Privileges The data source privileges that owners of the role have for the specified
	// data sources.
	Privileges []datasourceprivilege.DataSourcePrivilege `json:"privileges"`
}

// NewDataSourcePrivileges returns a DataSourcePrivileges.
func NewDataSourcePrivileges() *DataSourcePrivileges {
	r := &DataSourcePrivileges{}

	return r
}

type DataSourcePrivilegesVariant interface {
	DataSourcePrivilegesCaster() *DataSourcePrivileges
}

func (s *DataSourcePrivileges) DataSourcePrivilegesCaster() *DataSourcePrivileges {
	return s
}
