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

// UsageStatsShards type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/field_usage_stats/IndicesFieldUsageStatsResponse.ts#L42-L47
type UsageStatsShards struct {
	Routing                 ShardRouting        `json:"routing"`
	Stats                   ShardsStats         `json:"stats"`
	TrackingId              string              `json:"tracking_id"`
	TrackingStartedAtMillis EpochTimeUnitMillis `json:"tracking_started_at_millis"`
}

// UsageStatsShardsBuilder holds UsageStatsShards struct and provides a builder API.
type UsageStatsShardsBuilder struct {
	v *UsageStatsShards
}

// NewUsageStatsShards provides a builder for the UsageStatsShards struct.
func NewUsageStatsShardsBuilder() *UsageStatsShardsBuilder {
	r := UsageStatsShardsBuilder{
		&UsageStatsShards{},
	}

	return &r
}

// Build finalize the chain and returns the UsageStatsShards struct
func (rb *UsageStatsShardsBuilder) Build() UsageStatsShards {
	return *rb.v
}

func (rb *UsageStatsShardsBuilder) Routing(routing *ShardRoutingBuilder) *UsageStatsShardsBuilder {
	v := routing.Build()
	rb.v.Routing = v
	return rb
}

func (rb *UsageStatsShardsBuilder) Stats(stats *ShardsStatsBuilder) *UsageStatsShardsBuilder {
	v := stats.Build()
	rb.v.Stats = v
	return rb
}

func (rb *UsageStatsShardsBuilder) TrackingId(trackingid string) *UsageStatsShardsBuilder {
	rb.v.TrackingId = trackingid
	return rb
}

func (rb *UsageStatsShardsBuilder) TrackingStartedAtMillis(trackingstartedatmillis *EpochTimeUnitMillisBuilder) *UsageStatsShardsBuilder {
	v := trackingstartedatmillis.Build()
	rb.v.TrackingStartedAtMillis = v
	return rb
}
