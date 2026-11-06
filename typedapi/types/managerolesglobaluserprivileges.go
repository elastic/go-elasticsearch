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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

// ManageRolesGlobalUserPrivileges type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/security/_types/Privileges.ts#L459-L465
type ManageRolesGlobalUserPrivileges struct {
	// Manage The index-scoped role management privileges. Absent when the `role` section
	// is present but grants no privileges.
	Manage *ManageRolesPrivileges `json:"manage,omitempty"`
}

// NewManageRolesGlobalUserPrivileges returns a ManageRolesGlobalUserPrivileges.
func NewManageRolesGlobalUserPrivileges() *ManageRolesGlobalUserPrivileges {
	r := &ManageRolesGlobalUserPrivileges{}

	return r
}

type ManageRolesGlobalUserPrivilegesVariant interface {
	ManageRolesGlobalUserPrivilegesCaster() *ManageRolesGlobalUserPrivileges
}

func (s *ManageRolesGlobalUserPrivileges) ManageRolesGlobalUserPrivilegesCaster() *ManageRolesGlobalUserPrivileges {
	return s
}
