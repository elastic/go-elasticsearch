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

// AsyncSearchDocumentResponseBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/async_search/_types/AsyncSearchResponseBase.ts#L33-L37
type AsyncSearchDocumentResponseBase struct {
	ExpirationTime         *DateTime           `json:"expiration_time,omitempty"`
	ExpirationTimeInMillis EpochTimeUnitMillis `json:"expiration_time_in_millis"`
	Id                     *Id                 `json:"id,omitempty"`
	IsPartial              bool                `json:"is_partial"`
	IsRunning              bool                `json:"is_running"`
	Response               AsyncSearch         `json:"response"`
	StartTime              *DateTime           `json:"start_time,omitempty"`
	StartTimeInMillis      EpochTimeUnitMillis `json:"start_time_in_millis"`
}

// AsyncSearchDocumentResponseBaseBuilder holds AsyncSearchDocumentResponseBase struct and provides a builder API.
type AsyncSearchDocumentResponseBaseBuilder struct {
	v *AsyncSearchDocumentResponseBase
}

// NewAsyncSearchDocumentResponseBase provides a builder for the AsyncSearchDocumentResponseBase struct.
func NewAsyncSearchDocumentResponseBaseBuilder() *AsyncSearchDocumentResponseBaseBuilder {
	r := AsyncSearchDocumentResponseBaseBuilder{
		&AsyncSearchDocumentResponseBase{},
	}

	return &r
}

// Build finalize the chain and returns the AsyncSearchDocumentResponseBase struct
func (rb *AsyncSearchDocumentResponseBaseBuilder) Build() AsyncSearchDocumentResponseBase {
	return *rb.v
}

func (rb *AsyncSearchDocumentResponseBaseBuilder) ExpirationTime(expirationtime *DateTimeBuilder) *AsyncSearchDocumentResponseBaseBuilder {
	v := expirationtime.Build()
	rb.v.ExpirationTime = &v
	return rb
}

func (rb *AsyncSearchDocumentResponseBaseBuilder) ExpirationTimeInMillis(expirationtimeinmillis *EpochTimeUnitMillisBuilder) *AsyncSearchDocumentResponseBaseBuilder {
	v := expirationtimeinmillis.Build()
	rb.v.ExpirationTimeInMillis = v
	return rb
}

func (rb *AsyncSearchDocumentResponseBaseBuilder) Id(id Id) *AsyncSearchDocumentResponseBaseBuilder {
	rb.v.Id = &id
	return rb
}

func (rb *AsyncSearchDocumentResponseBaseBuilder) IsPartial(ispartial bool) *AsyncSearchDocumentResponseBaseBuilder {
	rb.v.IsPartial = ispartial
	return rb
}

func (rb *AsyncSearchDocumentResponseBaseBuilder) IsRunning(isrunning bool) *AsyncSearchDocumentResponseBaseBuilder {
	rb.v.IsRunning = isrunning
	return rb
}

func (rb *AsyncSearchDocumentResponseBaseBuilder) Response(response *AsyncSearchBuilder) *AsyncSearchDocumentResponseBaseBuilder {
	v := response.Build()
	rb.v.Response = v
	return rb
}

func (rb *AsyncSearchDocumentResponseBaseBuilder) StartTime(starttime *DateTimeBuilder) *AsyncSearchDocumentResponseBaseBuilder {
	v := starttime.Build()
	rb.v.StartTime = &v
	return rb
}

func (rb *AsyncSearchDocumentResponseBaseBuilder) StartTimeInMillis(starttimeinmillis *EpochTimeUnitMillisBuilder) *AsyncSearchDocumentResponseBaseBuilder {
	v := starttimeinmillis.Build()
	rb.v.StartTimeInMillis = v
	return rb
}
