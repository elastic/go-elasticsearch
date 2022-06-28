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

// TranslogStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/indices/recovery/types.ts#L94-L101
type TranslogStatus struct {
	Percent           Percentage  `json:"percent"`
	Recovered         int64       `json:"recovered"`
	Total             int64       `json:"total"`
	TotalOnStart      int64       `json:"total_on_start"`
	TotalTime         *string     `json:"total_time,omitempty"`
	TotalTimeInMillis EpochMillis `json:"total_time_in_millis"`
}

// TranslogStatusBuilder holds TranslogStatus struct and provides a builder API.
type TranslogStatusBuilder struct {
	v *TranslogStatus
}

// NewTranslogStatus provides a builder for the TranslogStatus struct.
func NewTranslogStatusBuilder() *TranslogStatusBuilder {
	r := TranslogStatusBuilder{
		&TranslogStatus{},
	}

	return &r
}

// Build finalize the chain and returns the TranslogStatus struct
func (rb *TranslogStatusBuilder) Build() TranslogStatus {
	return *rb.v
}

func (rb *TranslogStatusBuilder) Percent(percent *PercentageBuilder) *TranslogStatusBuilder {
	v := percent.Build()
	rb.v.Percent = v
	return rb
}

func (rb *TranslogStatusBuilder) Recovered(recovered int64) *TranslogStatusBuilder {
	rb.v.Recovered = recovered
	return rb
}

func (rb *TranslogStatusBuilder) Total(total int64) *TranslogStatusBuilder {
	rb.v.Total = total
	return rb
}

func (rb *TranslogStatusBuilder) TotalOnStart(totalonstart int64) *TranslogStatusBuilder {
	rb.v.TotalOnStart = totalonstart
	return rb
}

func (rb *TranslogStatusBuilder) TotalTime(totaltime string) *TranslogStatusBuilder {
	rb.v.TotalTime = &totaltime
	return rb
}

func (rb *TranslogStatusBuilder) TotalTimeInMillis(totaltimeinmillis *EpochMillisBuilder) *TranslogStatusBuilder {
	v := totaltimeinmillis.Build()
	rb.v.TotalTimeInMillis = v
	return rb
}
