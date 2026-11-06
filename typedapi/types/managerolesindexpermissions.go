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

// ManageRolesIndexPermissions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/security/_types/Privileges.ts#L471-L480
type ManageRolesIndexPermissions struct {
	// Names A list of indices (or index name patterns) that the owners of the role can
	// manage roles for.
	Names []string `json:"names"`
	// Privileges The index privileges that the managed roles are allowed to grant on the
	// specified indices.
	Privileges []string `json:"privileges"`
}

// NewManageRolesIndexPermissions returns a ManageRolesIndexPermissions.
func NewManageRolesIndexPermissions() *ManageRolesIndexPermissions {
	r := &ManageRolesIndexPermissions{}

	return r
}

type ManageRolesIndexPermissionsVariant interface {
	ManageRolesIndexPermissionsCaster() *ManageRolesIndexPermissions
}

func (s *ManageRolesIndexPermissions) ManageRolesIndexPermissionsCaster() *ManageRolesIndexPermissions {
	return s
}
