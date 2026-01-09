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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _rRFRetrieverEntry struct {
	v types.RRFRetrieverEntry
}

func NewRRFRetrieverEntry() *_rRFRetrieverEntry {
	return &_rRFRetrieverEntry{v: nil}
}

func (u *_rRFRetrieverEntry) RetrieverContainer(retrievercontainer types.RetrieverContainerVariant) *_rRFRetrieverEntry {

	u.v = retrievercontainer.RetrieverContainerCaster()

	return u
}

// Interface implementation for RetrieverContainer in RRFRetrieverEntry union
func (u *_retrieverContainer) RRFRetrieverEntryCaster() *types.RRFRetrieverEntry {
	t := types.RRFRetrieverEntry(u.v)
	return &t
}

func (u *_rRFRetrieverEntry) RRFRetrieverComponent(rrfretrievercomponent types.RRFRetrieverComponentVariant) *_rRFRetrieverEntry {

	u.v = rrfretrievercomponent.RRFRetrieverComponentCaster()

	return u
}

// Interface implementation for RRFRetrieverComponent in RRFRetrieverEntry union
func (u *_rRFRetrieverComponent) RRFRetrieverEntryCaster() *types.RRFRetrieverEntry {
	t := types.RRFRetrieverEntry(u.v)
	return &t
}

func (u *_rRFRetrieverEntry) RRFRetrieverEntryCaster() *types.RRFRetrieverEntry {
	return &u.v
}
