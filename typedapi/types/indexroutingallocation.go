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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexroutingallocationoptions"
)

// IndexRoutingAllocation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexRouting.ts#L27-L32
type IndexRoutingAllocation struct {
	Disk            *IndexRoutingAllocationDisk                                  `json:"disk,omitempty"`
	Enable          *indexroutingallocationoptions.IndexRoutingAllocationOptions `json:"enable,omitempty"`
	Include         *IndexRoutingAllocationInclude                               `json:"include,omitempty"`
	InitialRecovery *IndexRoutingAllocationInitialRecovery                       `json:"initial_recovery,omitempty"`
}

// IndexRoutingAllocationBuilder holds IndexRoutingAllocation struct and provides a builder API.
type IndexRoutingAllocationBuilder struct {
	v *IndexRoutingAllocation
}

// NewIndexRoutingAllocation provides a builder for the IndexRoutingAllocation struct.
func NewIndexRoutingAllocationBuilder() *IndexRoutingAllocationBuilder {
	r := IndexRoutingAllocationBuilder{
		&IndexRoutingAllocation{},
	}

	return &r
}

// Build finalize the chain and returns the IndexRoutingAllocation struct
func (rb *IndexRoutingAllocationBuilder) Build() IndexRoutingAllocation {
	return *rb.v
}

func (rb *IndexRoutingAllocationBuilder) Disk(disk *IndexRoutingAllocationDiskBuilder) *IndexRoutingAllocationBuilder {
	v := disk.Build()
	rb.v.Disk = &v
	return rb
}

func (rb *IndexRoutingAllocationBuilder) Enable(enable indexroutingallocationoptions.IndexRoutingAllocationOptions) *IndexRoutingAllocationBuilder {
	rb.v.Enable = &enable
	return rb
}

func (rb *IndexRoutingAllocationBuilder) Include(include *IndexRoutingAllocationIncludeBuilder) *IndexRoutingAllocationBuilder {
	v := include.Build()
	rb.v.Include = &v
	return rb
}

func (rb *IndexRoutingAllocationBuilder) InitialRecovery(initialrecovery *IndexRoutingAllocationInitialRecoveryBuilder) *IndexRoutingAllocationBuilder {
	v := initialrecovery.Build()
	rb.v.InitialRecovery = &v
	return rb
}
