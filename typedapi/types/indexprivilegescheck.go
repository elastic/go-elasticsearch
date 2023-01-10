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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexprivilege"
)

// IndexPrivilegesCheck type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/security/has_privileges/types.ts#L33-L44
type IndexPrivilegesCheck struct {
	// AllowRestrictedIndices This needs to be set to true (default is false) if using wildcards or regexps
	// for patterns that cover restricted indices.
	// Implicitly, restricted indices do not match index patterns because restricted
	// indices usually have limited privileges and including them in pattern tests
	// would render most such tests false.
	// If restricted indices are explicitly included in the names list, privileges
	// will be checked against them regardless of the value of
	// allow_restricted_indices.
	AllowRestrictedIndices *bool `json:"allow_restricted_indices,omitempty"`
	// Names A list of indices.
	Names []string `json:"names"`
	// Privileges A list of the privileges that you want to check for the specified indices.
	Privileges []indexprivilege.IndexPrivilege `json:"privileges"`
}

// NewIndexPrivilegesCheck returns a IndexPrivilegesCheck.
func NewIndexPrivilegesCheck() *IndexPrivilegesCheck {
	r := &IndexPrivilegesCheck{}

	return r
}
