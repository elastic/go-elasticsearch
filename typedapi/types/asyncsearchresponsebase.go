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

// AsyncSearchResponseBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/async_search/_types/AsyncSearchResponseBase.ts#L24-L30
type AsyncSearchResponseBase struct {
	ExpirationTimeInMillis EpochMillis `json:"expiration_time_in_millis"`
	Id                     *Id         `json:"id,omitempty"`
	IsPartial              bool        `json:"is_partial"`
	IsRunning              bool        `json:"is_running"`
	StartTimeInMillis      EpochMillis `json:"start_time_in_millis"`
}

// AsyncSearchResponseBaseBuilder holds AsyncSearchResponseBase struct and provides a builder API.
type AsyncSearchResponseBaseBuilder struct {
	v *AsyncSearchResponseBase
}

// NewAsyncSearchResponseBase provides a builder for the AsyncSearchResponseBase struct.
func NewAsyncSearchResponseBaseBuilder() *AsyncSearchResponseBaseBuilder {
	r := AsyncSearchResponseBaseBuilder{
		&AsyncSearchResponseBase{},
	}

	return &r
}

// Build finalize the chain and returns the AsyncSearchResponseBase struct
func (rb *AsyncSearchResponseBaseBuilder) Build() AsyncSearchResponseBase {
	return *rb.v
}

func (rb *AsyncSearchResponseBaseBuilder) ExpirationTimeInMillis(expirationtimeinmillis *EpochMillisBuilder) *AsyncSearchResponseBaseBuilder {
	v := expirationtimeinmillis.Build()
	rb.v.ExpirationTimeInMillis = v
	return rb
}

func (rb *AsyncSearchResponseBaseBuilder) Id(id Id) *AsyncSearchResponseBaseBuilder {
	rb.v.Id = &id
	return rb
}

func (rb *AsyncSearchResponseBaseBuilder) IsPartial(ispartial bool) *AsyncSearchResponseBaseBuilder {
	rb.v.IsPartial = ispartial
	return rb
}

func (rb *AsyncSearchResponseBaseBuilder) IsRunning(isrunning bool) *AsyncSearchResponseBaseBuilder {
	rb.v.IsRunning = isrunning
	return rb
}

func (rb *AsyncSearchResponseBaseBuilder) StartTimeInMillis(starttimeinmillis *EpochMillisBuilder) *AsyncSearchResponseBaseBuilder {
	v := starttimeinmillis.Build()
	rb.v.StartTimeInMillis = v
	return rb
}
