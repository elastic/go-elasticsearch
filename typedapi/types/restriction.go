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
// https://github.com/elastic/elasticsearch-specification/tree/f6a370d0fba975752c644fc730f7c45610e28f36

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/restrictionworkflow"
)

// Restriction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/f6a370d0fba975752c644fc730f7c45610e28f36/specification/security/_types/RoleDescriptor.ts#L135-L141
type Restriction struct {
	// Workflows A list of workflows to which the API key is restricted.
	// NOTE: In order to use a role restriction, an API key must be created with a
	// single role descriptor.
	Workflows []restrictionworkflow.RestrictionWorkflow `json:"workflows"`
}

// NewRestriction returns a Restriction.
func NewRestriction() *Restriction {
	r := &Restriction{}

	return r
}

// true

type RestrictionVariant interface {
	RestrictionCaster() *Restriction
}

func (s *Restriction) RestrictionCaster() *Restriction {
	return s
}
