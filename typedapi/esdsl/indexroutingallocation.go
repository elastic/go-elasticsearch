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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexroutingallocationoptions"
)

type _indexRoutingAllocation struct {
	v *types.IndexRoutingAllocation
}

func NewIndexRoutingAllocation() *_indexRoutingAllocation {

	return &_indexRoutingAllocation{v: types.NewIndexRoutingAllocation()}

}

func (s *_indexRoutingAllocation) Disk(disk types.IndexRoutingAllocationDiskVariant) *_indexRoutingAllocation {

	s.v.Disk = disk.IndexRoutingAllocationDiskCaster()

	return s
}

func (s *_indexRoutingAllocation) Enable(enable indexroutingallocationoptions.IndexRoutingAllocationOptions) *_indexRoutingAllocation {

	s.v.Enable = &enable
	return s
}

func (s *_indexRoutingAllocation) Include(include types.IndexRoutingAllocationIncludeVariant) *_indexRoutingAllocation {

	s.v.Include = include.IndexRoutingAllocationIncludeCaster()

	return s
}

func (s *_indexRoutingAllocation) InitialRecovery(initialrecovery types.IndexRoutingAllocationInitialRecoveryVariant) *_indexRoutingAllocation {

	s.v.InitialRecovery = initialrecovery.IndexRoutingAllocationInitialRecoveryCaster()

	return s
}

func (s *_indexRoutingAllocation) IndexRoutingAllocationCaster() *types.IndexRoutingAllocation {
	return s.v
}
