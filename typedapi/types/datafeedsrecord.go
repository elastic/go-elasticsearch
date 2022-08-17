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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/datafeedstate"
)

// DatafeedsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/ml_datafeeds/types.ts#L22-L83
type DatafeedsRecord struct {
	// AssignmentExplanation why the datafeed is or is not assigned to a node
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`
	// BucketsCount bucket count
	BucketsCount *string `json:"buckets.count,omitempty"`
	// Id the datafeed_id
	Id *string `json:"id,omitempty"`
	// NodeAddress network address of the assigned node
	NodeAddress *string `json:"node.address,omitempty"`
	// NodeEphemeralId ephemeral id of the assigned node
	NodeEphemeralId *string `json:"node.ephemeral_id,omitempty"`
	// NodeId id of the assigned node
	NodeId *string `json:"node.id,omitempty"`
	// NodeName name of the assigned node
	NodeName *string `json:"node.name,omitempty"`
	// SearchBucketAvg the average search time per bucket (millisecond)
	SearchBucketAvg *string `json:"search.bucket_avg,omitempty"`
	// SearchCount number of searches ran by the datafeed
	SearchCount *string `json:"search.count,omitempty"`
	// SearchExpAvgHour the exponential average search time per hour (millisecond)
	SearchExpAvgHour *string `json:"search.exp_avg_hour,omitempty"`
	// SearchTime the total search time
	SearchTime *string `json:"search.time,omitempty"`
	// State the datafeed state
	State *datafeedstate.DatafeedState `json:"state,omitempty"`
}

// DatafeedsRecordBuilder holds DatafeedsRecord struct and provides a builder API.
type DatafeedsRecordBuilder struct {
	v *DatafeedsRecord
}

// NewDatafeedsRecord provides a builder for the DatafeedsRecord struct.
func NewDatafeedsRecordBuilder() *DatafeedsRecordBuilder {
	r := DatafeedsRecordBuilder{
		&DatafeedsRecord{},
	}

	return &r
}

// Build finalize the chain and returns the DatafeedsRecord struct
func (rb *DatafeedsRecordBuilder) Build() DatafeedsRecord {
	return *rb.v
}

// AssignmentExplanation why the datafeed is or is not assigned to a node

func (rb *DatafeedsRecordBuilder) AssignmentExplanation(assignmentexplanation string) *DatafeedsRecordBuilder {
	rb.v.AssignmentExplanation = &assignmentexplanation
	return rb
}

// BucketsCount bucket count

func (rb *DatafeedsRecordBuilder) BucketsCount(bucketscount string) *DatafeedsRecordBuilder {
	rb.v.BucketsCount = &bucketscount
	return rb
}

// Id the datafeed_id

func (rb *DatafeedsRecordBuilder) Id(id string) *DatafeedsRecordBuilder {
	rb.v.Id = &id
	return rb
}

// NodeAddress network address of the assigned node

func (rb *DatafeedsRecordBuilder) NodeAddress(nodeaddress string) *DatafeedsRecordBuilder {
	rb.v.NodeAddress = &nodeaddress
	return rb
}

// NodeEphemeralId ephemeral id of the assigned node

func (rb *DatafeedsRecordBuilder) NodeEphemeralId(nodeephemeralid string) *DatafeedsRecordBuilder {
	rb.v.NodeEphemeralId = &nodeephemeralid
	return rb
}

// NodeId id of the assigned node

func (rb *DatafeedsRecordBuilder) NodeId(nodeid string) *DatafeedsRecordBuilder {
	rb.v.NodeId = &nodeid
	return rb
}

// NodeName name of the assigned node

func (rb *DatafeedsRecordBuilder) NodeName(nodename string) *DatafeedsRecordBuilder {
	rb.v.NodeName = &nodename
	return rb
}

// SearchBucketAvg the average search time per bucket (millisecond)

func (rb *DatafeedsRecordBuilder) SearchBucketAvg(searchbucketavg string) *DatafeedsRecordBuilder {
	rb.v.SearchBucketAvg = &searchbucketavg
	return rb
}

// SearchCount number of searches ran by the datafeed

func (rb *DatafeedsRecordBuilder) SearchCount(searchcount string) *DatafeedsRecordBuilder {
	rb.v.SearchCount = &searchcount
	return rb
}

// SearchExpAvgHour the exponential average search time per hour (millisecond)

func (rb *DatafeedsRecordBuilder) SearchExpAvgHour(searchexpavghour string) *DatafeedsRecordBuilder {
	rb.v.SearchExpAvgHour = &searchexpavghour
	return rb
}

// SearchTime the total search time

func (rb *DatafeedsRecordBuilder) SearchTime(searchtime string) *DatafeedsRecordBuilder {
	rb.v.SearchTime = &searchtime
	return rb
}

// State the datafeed state

func (rb *DatafeedsRecordBuilder) State(state datafeedstate.DatafeedState) *DatafeedsRecordBuilder {
	rb.v.State = &state
	return rb
}
