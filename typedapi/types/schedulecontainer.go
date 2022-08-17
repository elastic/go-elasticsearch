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

// ScheduleContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Schedule.ts#L85-L96
type ScheduleContainer struct {
	Cron     *CronExpression `json:"cron,omitempty"`
	Daily    *DailySchedule  `json:"daily,omitempty"`
	Hourly   *HourlySchedule `json:"hourly,omitempty"`
	Interval *Duration       `json:"interval,omitempty"`
	Monthly  []TimeOfMonth   `json:"monthly,omitempty"`
	Weekly   []TimeOfWeek    `json:"weekly,omitempty"`
	Yearly   []TimeOfYear    `json:"yearly,omitempty"`
}

// ScheduleContainerBuilder holds ScheduleContainer struct and provides a builder API.
type ScheduleContainerBuilder struct {
	v *ScheduleContainer
}

// NewScheduleContainer provides a builder for the ScheduleContainer struct.
func NewScheduleContainerBuilder() *ScheduleContainerBuilder {
	r := ScheduleContainerBuilder{
		&ScheduleContainer{},
	}

	return &r
}

// Build finalize the chain and returns the ScheduleContainer struct
func (rb *ScheduleContainerBuilder) Build() ScheduleContainer {
	return *rb.v
}

func (rb *ScheduleContainerBuilder) Cron(cron CronExpression) *ScheduleContainerBuilder {
	rb.v.Cron = &cron
	return rb
}

func (rb *ScheduleContainerBuilder) Daily(daily *DailyScheduleBuilder) *ScheduleContainerBuilder {
	v := daily.Build()
	rb.v.Daily = &v
	return rb
}

func (rb *ScheduleContainerBuilder) Hourly(hourly *HourlyScheduleBuilder) *ScheduleContainerBuilder {
	v := hourly.Build()
	rb.v.Hourly = &v
	return rb
}

func (rb *ScheduleContainerBuilder) Interval(interval *DurationBuilder) *ScheduleContainerBuilder {
	v := interval.Build()
	rb.v.Interval = &v
	return rb
}

func (rb *ScheduleContainerBuilder) Monthly(arg []TimeOfMonth) *ScheduleContainerBuilder {
	rb.v.Monthly = arg
	return rb
}

func (rb *ScheduleContainerBuilder) Weekly(arg []TimeOfWeek) *ScheduleContainerBuilder {
	rb.v.Weekly = arg
	return rb
}

func (rb *ScheduleContainerBuilder) Yearly(arg []TimeOfYear) *ScheduleContainerBuilder {
	rb.v.Yearly = arg
	return rb
}
