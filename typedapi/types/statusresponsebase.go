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

// StatusResponseBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/async_search/status/AsyncSearchStatusResponse.ts#L24-L27
type StatusResponseBase struct {
	CompletionStatus       *int                `json:"completion_status,omitempty"`
	ExpirationTime         *DateTime           `json:"expiration_time,omitempty"`
	ExpirationTimeInMillis EpochTimeUnitMillis `json:"expiration_time_in_millis"`
	Id                     *Id                 `json:"id,omitempty"`
	IsPartial              bool                `json:"is_partial"`
	IsRunning              bool                `json:"is_running"`
	Shards_                ShardStatistics     `json:"_shards"`
	StartTime              *DateTime           `json:"start_time,omitempty"`
	StartTimeInMillis      EpochTimeUnitMillis `json:"start_time_in_millis"`
}

// StatusResponseBaseBuilder holds StatusResponseBase struct and provides a builder API.
type StatusResponseBaseBuilder struct {
	v *StatusResponseBase
}

// NewStatusResponseBase provides a builder for the StatusResponseBase struct.
func NewStatusResponseBaseBuilder() *StatusResponseBaseBuilder {
	r := StatusResponseBaseBuilder{
		&StatusResponseBase{},
	}

	return &r
}

// Build finalize the chain and returns the StatusResponseBase struct
func (rb *StatusResponseBaseBuilder) Build() StatusResponseBase {
	return *rb.v
}

func (rb *StatusResponseBaseBuilder) CompletionStatus(completionstatus int) *StatusResponseBaseBuilder {
	rb.v.CompletionStatus = &completionstatus
	return rb
}

func (rb *StatusResponseBaseBuilder) ExpirationTime(expirationtime *DateTimeBuilder) *StatusResponseBaseBuilder {
	v := expirationtime.Build()
	rb.v.ExpirationTime = &v
	return rb
}

func (rb *StatusResponseBaseBuilder) ExpirationTimeInMillis(expirationtimeinmillis *EpochTimeUnitMillisBuilder) *StatusResponseBaseBuilder {
	v := expirationtimeinmillis.Build()
	rb.v.ExpirationTimeInMillis = v
	return rb
}

func (rb *StatusResponseBaseBuilder) Id(id Id) *StatusResponseBaseBuilder {
	rb.v.Id = &id
	return rb
}

func (rb *StatusResponseBaseBuilder) IsPartial(ispartial bool) *StatusResponseBaseBuilder {
	rb.v.IsPartial = ispartial
	return rb
}

func (rb *StatusResponseBaseBuilder) IsRunning(isrunning bool) *StatusResponseBaseBuilder {
	rb.v.IsRunning = isrunning
	return rb
}

func (rb *StatusResponseBaseBuilder) Shards_(shards_ *ShardStatisticsBuilder) *StatusResponseBaseBuilder {
	v := shards_.Build()
	rb.v.Shards_ = v
	return rb
}

func (rb *StatusResponseBaseBuilder) StartTime(starttime *DateTimeBuilder) *StatusResponseBaseBuilder {
	v := starttime.Build()
	rb.v.StartTime = &v
	return rb
}

func (rb *StatusResponseBaseBuilder) StartTimeInMillis(starttimeinmillis *EpochTimeUnitMillisBuilder) *StatusResponseBaseBuilder {
	v := starttimeinmillis.Build()
	rb.v.StartTimeInMillis = v
	return rb
}
