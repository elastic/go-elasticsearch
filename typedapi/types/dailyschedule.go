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

// DailySchedule type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Schedule.ts#L33-L35
type DailySchedule struct {
	At []ScheduleTimeOfDay `json:"at"`
}

// DailyScheduleBuilder holds DailySchedule struct and provides a builder API.
type DailyScheduleBuilder struct {
	v *DailySchedule
}

// NewDailySchedule provides a builder for the DailySchedule struct.
func NewDailyScheduleBuilder() *DailyScheduleBuilder {
	r := DailyScheduleBuilder{
		&DailySchedule{},
	}

	return &r
}

// Build finalize the chain and returns the DailySchedule struct
func (rb *DailyScheduleBuilder) Build() DailySchedule {
	return *rb.v
}

func (rb *DailyScheduleBuilder) At(at ...ScheduleTimeOfDay) *DailyScheduleBuilder {
	rb.v.At = at
	return rb
}
