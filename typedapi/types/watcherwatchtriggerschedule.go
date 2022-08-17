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

// WatcherWatchTriggerSchedule type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L455-L458
type WatcherWatchTriggerSchedule struct {
	Active int64   `json:"active"`
	All_   Counter `json:"_all"`
	Cron   Counter `json:"cron"`
	Total  int64   `json:"total"`
}

// WatcherWatchTriggerScheduleBuilder holds WatcherWatchTriggerSchedule struct and provides a builder API.
type WatcherWatchTriggerScheduleBuilder struct {
	v *WatcherWatchTriggerSchedule
}

// NewWatcherWatchTriggerSchedule provides a builder for the WatcherWatchTriggerSchedule struct.
func NewWatcherWatchTriggerScheduleBuilder() *WatcherWatchTriggerScheduleBuilder {
	r := WatcherWatchTriggerScheduleBuilder{
		&WatcherWatchTriggerSchedule{},
	}

	return &r
}

// Build finalize the chain and returns the WatcherWatchTriggerSchedule struct
func (rb *WatcherWatchTriggerScheduleBuilder) Build() WatcherWatchTriggerSchedule {
	return *rb.v
}

func (rb *WatcherWatchTriggerScheduleBuilder) Active(active int64) *WatcherWatchTriggerScheduleBuilder {
	rb.v.Active = active
	return rb
}

func (rb *WatcherWatchTriggerScheduleBuilder) All_(all_ *CounterBuilder) *WatcherWatchTriggerScheduleBuilder {
	v := all_.Build()
	rb.v.All_ = v
	return rb
}

func (rb *WatcherWatchTriggerScheduleBuilder) Cron(cron *CounterBuilder) *WatcherWatchTriggerScheduleBuilder {
	v := cron.Build()
	rb.v.Cron = v
	return rb
}

func (rb *WatcherWatchTriggerScheduleBuilder) Total(total int64) *WatcherWatchTriggerScheduleBuilder {
	rb.v.Total = total
	return rb
}
