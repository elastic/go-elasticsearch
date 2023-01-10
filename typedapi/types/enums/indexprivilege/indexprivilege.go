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


// Package indexprivilege
package indexprivilege

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/security/_types/Privileges.ts#L165-L185
type IndexPrivilege struct {
	Name string
}

var (
	None = IndexPrivilege{"none"}

	All = IndexPrivilege{"all"}

	Autoconfigure = IndexPrivilege{"auto_configure"}

	Create = IndexPrivilege{"create"}

	Createdoc = IndexPrivilege{"create_doc"}

	Createindex = IndexPrivilege{"create_index"}

	Delete = IndexPrivilege{"delete"}

	Deleteindex = IndexPrivilege{"delete_index"}

	Index = IndexPrivilege{"index"}

	Maintenance = IndexPrivilege{"maintenance"}

	Manage = IndexPrivilege{"manage"}

	Managefollowindex = IndexPrivilege{"manage_follow_index"}

	Manageilm = IndexPrivilege{"manage_ilm"}

	Manageleaderindex = IndexPrivilege{"manage_leader_index"}

	Monitor = IndexPrivilege{"monitor"}

	Read = IndexPrivilege{"read"}

	Readcrosscluster = IndexPrivilege{"read_cross_cluster"}

	Viewindexmetadata = IndexPrivilege{"view_index_metadata"}

	Write = IndexPrivilege{"write"}
)

func (i IndexPrivilege) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *IndexPrivilege) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "none":
		*i = None
	case "all":
		*i = All
	case "auto_configure":
		*i = Autoconfigure
	case "create":
		*i = Create
	case "create_doc":
		*i = Createdoc
	case "create_index":
		*i = Createindex
	case "delete":
		*i = Delete
	case "delete_index":
		*i = Deleteindex
	case "index":
		*i = Index
	case "maintenance":
		*i = Maintenance
	case "manage":
		*i = Manage
	case "manage_follow_index":
		*i = Managefollowindex
	case "manage_ilm":
		*i = Manageilm
	case "manage_leader_index":
		*i = Manageleaderindex
	case "monitor":
		*i = Monitor
	case "read":
		*i = Read
	case "read_cross_cluster":
		*i = Readcrosscluster
	case "view_index_metadata":
		*i = Viewindexmetadata
	case "write":
		*i = Write
	default:
		*i = IndexPrivilege{string(text)}
	}

	return nil
}

func (i IndexPrivilege) String() string {
	return i.Name
}
