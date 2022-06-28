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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// ClusterStateUpdate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/nodes/_types/Stats.ts#L119-L135
type ClusterStateUpdate struct {
	CommitTime                    *string `json:"commit_time,omitempty"`
	CommitTimeMillis              *int64  `json:"commit_time_millis,omitempty"`
	CompletionTime                *string `json:"completion_time,omitempty"`
	CompletionTimeMillis          *int64  `json:"completion_time_millis,omitempty"`
	ComputationTime               *string `json:"computation_time,omitempty"`
	ComputationTimeMillis         *int64  `json:"computation_time_millis,omitempty"`
	ContextConstructionTime       *string `json:"context_construction_time,omitempty"`
	ContextConstructionTimeMillis *int64  `json:"context_construction_time_millis,omitempty"`
	Count                         *int64  `json:"count,omitempty"`
	MasterApplyTime               *string `json:"master_apply_time,omitempty"`
	MasterApplyTimeMillis         *int64  `json:"master_apply_time_millis,omitempty"`
	NotificationTime              *string `json:"notification_time,omitempty"`
	NotificationTimeMillis        *int64  `json:"notification_time_millis,omitempty"`
	PublicationTime               *string `json:"publication_time,omitempty"`
	PublicationTimeMillis         *int64  `json:"publication_time_millis,omitempty"`
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

func (rb *ClusterStateUpdateBuilder) CommitTime(committime string) *ClusterStateUpdateBuilder {
	rb.v.CommitTime = &committime
	return rb
}

func (rb *ClusterStateUpdateBuilder) CommitTimeMillis(committimemillis int64) *ClusterStateUpdateBuilder {
	rb.v.CommitTimeMillis = &committimemillis
	return rb
}

func (rb *ClusterStateUpdateBuilder) CompletionTime(completiontime string) *ClusterStateUpdateBuilder {
	rb.v.CompletionTime = &completiontime
	return rb
}

func (rb *ClusterStateUpdateBuilder) CompletionTimeMillis(completiontimemillis int64) *ClusterStateUpdateBuilder {
	rb.v.CompletionTimeMillis = &completiontimemillis
	return rb
}

func (rb *ClusterStateUpdateBuilder) ComputationTime(computationtime string) *ClusterStateUpdateBuilder {
	rb.v.ComputationTime = &computationtime
	return rb
}

func (rb *ClusterStateUpdateBuilder) ComputationTimeMillis(computationtimemillis int64) *ClusterStateUpdateBuilder {
	rb.v.ComputationTimeMillis = &computationtimemillis
	return rb
}

func (rb *ClusterStateUpdateBuilder) ContextConstructionTime(contextconstructiontime string) *ClusterStateUpdateBuilder {
	rb.v.ContextConstructionTime = &contextconstructiontime
	return rb
}

func (rb *ClusterStateUpdateBuilder) ContextConstructionTimeMillis(contextconstructiontimemillis int64) *ClusterStateUpdateBuilder {
	rb.v.ContextConstructionTimeMillis = &contextconstructiontimemillis
	return rb
}

func (rb *ClusterStateUpdateBuilder) Count(count int64) *ClusterStateUpdateBuilder {
	rb.v.Count = &count
	return rb
}

func (rb *ClusterStateUpdateBuilder) MasterApplyTime(masterapplytime string) *ClusterStateUpdateBuilder {
	rb.v.MasterApplyTime = &masterapplytime
	return rb
}

func (rb *ClusterStateUpdateBuilder) MasterApplyTimeMillis(masterapplytimemillis int64) *ClusterStateUpdateBuilder {
	rb.v.MasterApplyTimeMillis = &masterapplytimemillis
	return rb
}

func (rb *ClusterStateUpdateBuilder) NotificationTime(notificationtime string) *ClusterStateUpdateBuilder {
	rb.v.NotificationTime = &notificationtime
	return rb
}

func (rb *ClusterStateUpdateBuilder) NotificationTimeMillis(notificationtimemillis int64) *ClusterStateUpdateBuilder {
	rb.v.NotificationTimeMillis = &notificationtimemillis
	return rb
}

func (rb *ClusterStateUpdateBuilder) PublicationTime(publicationtime string) *ClusterStateUpdateBuilder {
	rb.v.PublicationTime = &publicationtime
	return rb
}

func (rb *ClusterStateUpdateBuilder) PublicationTimeMillis(publicationtimemillis int64) *ClusterStateUpdateBuilder {
	rb.v.PublicationTimeMillis = &publicationtimemillis
	return rb
}
