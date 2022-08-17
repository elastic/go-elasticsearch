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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/unassignedinformationreason"
)

// UnassignedInformation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/allocation_explain/types.ts#L117-L125
type UnassignedInformation struct {
	AllocationStatus         *string                                                 `json:"allocation_status,omitempty"`
	At                       DateTime                                                `json:"at"`
	Delayed                  *bool                                                   `json:"delayed,omitempty"`
	Details                  *string                                                 `json:"details,omitempty"`
	FailedAllocationAttempts *int                                                    `json:"failed_allocation_attempts,omitempty"`
	LastAllocationStatus     *string                                                 `json:"last_allocation_status,omitempty"`
	Reason                   unassignedinformationreason.UnassignedInformationReason `json:"reason"`
}

// UnassignedInformationBuilder holds UnassignedInformation struct and provides a builder API.
type UnassignedInformationBuilder struct {
	v *UnassignedInformation
}

// NewUnassignedInformation provides a builder for the UnassignedInformation struct.
func NewUnassignedInformationBuilder() *UnassignedInformationBuilder {
	r := UnassignedInformationBuilder{
		&UnassignedInformation{},
	}

	return &r
}

// Build finalize the chain and returns the UnassignedInformation struct
func (rb *UnassignedInformationBuilder) Build() UnassignedInformation {
	return *rb.v
}

func (rb *UnassignedInformationBuilder) AllocationStatus(allocationstatus string) *UnassignedInformationBuilder {
	rb.v.AllocationStatus = &allocationstatus
	return rb
}

func (rb *UnassignedInformationBuilder) At(at *DateTimeBuilder) *UnassignedInformationBuilder {
	v := at.Build()
	rb.v.At = v
	return rb
}

func (rb *UnassignedInformationBuilder) Delayed(delayed bool) *UnassignedInformationBuilder {
	rb.v.Delayed = &delayed
	return rb
}

func (rb *UnassignedInformationBuilder) Details(details string) *UnassignedInformationBuilder {
	rb.v.Details = &details
	return rb
}

func (rb *UnassignedInformationBuilder) FailedAllocationAttempts(failedallocationattempts int) *UnassignedInformationBuilder {
	rb.v.FailedAllocationAttempts = &failedallocationattempts
	return rb
}

func (rb *UnassignedInformationBuilder) LastAllocationStatus(lastallocationstatus string) *UnassignedInformationBuilder {
	rb.v.LastAllocationStatus = &lastallocationstatus
	return rb
}

func (rb *UnassignedInformationBuilder) Reason(reason unassignedinformationreason.UnassignedInformationReason) *UnassignedInformationBuilder {
	rb.v.Reason = reason
	return rb
}
