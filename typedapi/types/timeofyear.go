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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/month"
)

// TimeOfYear type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/watcher/_types/Schedule.ts#L125-L129
type TimeOfYear struct {
	At  []string      `json:"at"`
	Int []month.Month `json:"int"`
	On  []int         `json:"on"`
}

// TimeOfYearBuilder holds TimeOfYear struct and provides a builder API.
type TimeOfYearBuilder struct {
	v *TimeOfYear
}

// NewTimeOfYear provides a builder for the TimeOfYear struct.
func NewTimeOfYearBuilder() *TimeOfYearBuilder {
	r := TimeOfYearBuilder{
		&TimeOfYear{},
	}

	return &r
}

// Build finalize the chain and returns the TimeOfYear struct
func (rb *TimeOfYearBuilder) Build() TimeOfYear {
	return *rb.v
}

func (rb *TimeOfYearBuilder) At(at ...string) *TimeOfYearBuilder {
	rb.v.At = at
	return rb
}

func (rb *TimeOfYearBuilder) Int(int ...month.Month) *TimeOfYearBuilder {
	rb.v.Int = int
	return rb
}

func (rb *TimeOfYearBuilder) On(on ...int) *TimeOfYearBuilder {
	rb.v.On = on
	return rb
}
