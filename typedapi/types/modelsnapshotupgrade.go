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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/snapshotupgradestate"
)

// ModelSnapshotUpgrade type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Model.ts#L48-L54
type ModelSnapshotUpgrade struct {
	AssignmentExplanation string                                    `json:"assignment_explanation"`
	JobId                 Id                                        `json:"job_id"`
	Node                  DiscoveryNode                             `json:"node"`
	SnapshotId            Id                                        `json:"snapshot_id"`
	State                 snapshotupgradestate.SnapshotUpgradeState `json:"state"`
}

// ModelSnapshotUpgradeBuilder holds ModelSnapshotUpgrade struct and provides a builder API.
type ModelSnapshotUpgradeBuilder struct {
	v *ModelSnapshotUpgrade
}

// NewModelSnapshotUpgrade provides a builder for the ModelSnapshotUpgrade struct.
func NewModelSnapshotUpgradeBuilder() *ModelSnapshotUpgradeBuilder {
	r := ModelSnapshotUpgradeBuilder{
		&ModelSnapshotUpgrade{},
	}

	return &r
}

// Build finalize the chain and returns the ModelSnapshotUpgrade struct
func (rb *ModelSnapshotUpgradeBuilder) Build() ModelSnapshotUpgrade {
	return *rb.v
}

func (rb *ModelSnapshotUpgradeBuilder) AssignmentExplanation(assignmentexplanation string) *ModelSnapshotUpgradeBuilder {
	rb.v.AssignmentExplanation = assignmentexplanation
	return rb
}

func (rb *ModelSnapshotUpgradeBuilder) JobId(jobid Id) *ModelSnapshotUpgradeBuilder {
	rb.v.JobId = jobid
	return rb
}

func (rb *ModelSnapshotUpgradeBuilder) Node(node *DiscoveryNodeBuilder) *ModelSnapshotUpgradeBuilder {
	v := node.Build()
	rb.v.Node = v
	return rb
}

func (rb *ModelSnapshotUpgradeBuilder) SnapshotId(snapshotid Id) *ModelSnapshotUpgradeBuilder {
	rb.v.SnapshotId = snapshotid
	return rb
}

func (rb *ModelSnapshotUpgradeBuilder) State(state snapshotupgradestate.SnapshotUpgradeState) *ModelSnapshotUpgradeBuilder {
	rb.v.State = state
	return rb
}
