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

// ClusterStateUpdate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L120-L136
type ClusterStateUpdate struct {
	CommitTime                    *Duration                `json:"commit_time,omitempty"`
	CommitTimeMillis              *DurationValueUnitMillis `json:"commit_time_millis,omitempty"`
	CompletionTime                *Duration                `json:"completion_time,omitempty"`
	CompletionTimeMillis          *DurationValueUnitMillis `json:"completion_time_millis,omitempty"`
	ComputationTime               *Duration                `json:"computation_time,omitempty"`
	ComputationTimeMillis         *DurationValueUnitMillis `json:"computation_time_millis,omitempty"`
	ContextConstructionTime       *Duration                `json:"context_construction_time,omitempty"`
	ContextConstructionTimeMillis *DurationValueUnitMillis `json:"context_construction_time_millis,omitempty"`
	Count                         int64                    `json:"count"`
	MasterApplyTime               *Duration                `json:"master_apply_time,omitempty"`
	MasterApplyTimeMillis         *DurationValueUnitMillis `json:"master_apply_time_millis,omitempty"`
	NotificationTime              *Duration                `json:"notification_time,omitempty"`
	NotificationTimeMillis        *DurationValueUnitMillis `json:"notification_time_millis,omitempty"`
	PublicationTime               *Duration                `json:"publication_time,omitempty"`
	PublicationTimeMillis         *DurationValueUnitMillis `json:"publication_time_millis,omitempty"`
}

// ClusterStateUpdateBuilder holds ClusterStateUpdate struct and provides a builder API.
type ClusterStateUpdateBuilder struct {
	v *ClusterStateUpdate
}

// NewClusterStateUpdate provides a builder for the ClusterStateUpdate struct.
func NewClusterStateUpdateBuilder() *ClusterStateUpdateBuilder {
	r := ClusterStateUpdateBuilder{
		&ClusterStateUpdate{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterStateUpdate struct
func (rb *ClusterStateUpdateBuilder) Build() ClusterStateUpdate {
	return *rb.v
}

func (rb *ClusterStateUpdateBuilder) CommitTime(committime *DurationBuilder) *ClusterStateUpdateBuilder {
	v := committime.Build()
	rb.v.CommitTime = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) CommitTimeMillis(committimemillis *DurationValueUnitMillisBuilder) *ClusterStateUpdateBuilder {
	v := committimemillis.Build()
	rb.v.CommitTimeMillis = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) CompletionTime(completiontime *DurationBuilder) *ClusterStateUpdateBuilder {
	v := completiontime.Build()
	rb.v.CompletionTime = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) CompletionTimeMillis(completiontimemillis *DurationValueUnitMillisBuilder) *ClusterStateUpdateBuilder {
	v := completiontimemillis.Build()
	rb.v.CompletionTimeMillis = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) ComputationTime(computationtime *DurationBuilder) *ClusterStateUpdateBuilder {
	v := computationtime.Build()
	rb.v.ComputationTime = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) ComputationTimeMillis(computationtimemillis *DurationValueUnitMillisBuilder) *ClusterStateUpdateBuilder {
	v := computationtimemillis.Build()
	rb.v.ComputationTimeMillis = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) ContextConstructionTime(contextconstructiontime *DurationBuilder) *ClusterStateUpdateBuilder {
	v := contextconstructiontime.Build()
	rb.v.ContextConstructionTime = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) ContextConstructionTimeMillis(contextconstructiontimemillis *DurationValueUnitMillisBuilder) *ClusterStateUpdateBuilder {
	v := contextconstructiontimemillis.Build()
	rb.v.ContextConstructionTimeMillis = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) Count(count int64) *ClusterStateUpdateBuilder {
	rb.v.Count = count
	return rb
}

func (rb *ClusterStateUpdateBuilder) MasterApplyTime(masterapplytime *DurationBuilder) *ClusterStateUpdateBuilder {
	v := masterapplytime.Build()
	rb.v.MasterApplyTime = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) MasterApplyTimeMillis(masterapplytimemillis *DurationValueUnitMillisBuilder) *ClusterStateUpdateBuilder {
	v := masterapplytimemillis.Build()
	rb.v.MasterApplyTimeMillis = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) NotificationTime(notificationtime *DurationBuilder) *ClusterStateUpdateBuilder {
	v := notificationtime.Build()
	rb.v.NotificationTime = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) NotificationTimeMillis(notificationtimemillis *DurationValueUnitMillisBuilder) *ClusterStateUpdateBuilder {
	v := notificationtimemillis.Build()
	rb.v.NotificationTimeMillis = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) PublicationTime(publicationtime *DurationBuilder) *ClusterStateUpdateBuilder {
	v := publicationtime.Build()
	rb.v.PublicationTime = &v
	return rb
}

func (rb *ClusterStateUpdateBuilder) PublicationTimeMillis(publicationtimemillis *DurationValueUnitMillisBuilder) *ClusterStateUpdateBuilder {
	v := publicationtimemillis.Build()
	rb.v.PublicationTimeMillis = &v
	return rb
}
