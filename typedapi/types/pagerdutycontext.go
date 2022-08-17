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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/pagerdutycontexttype"
)

// PagerDutyContext type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L61-L65
type PagerDutyContext struct {
	Href *string                                   `json:"href,omitempty"`
	Src  *string                                   `json:"src,omitempty"`
	Type pagerdutycontexttype.PagerDutyContextType `json:"type"`
}

// PagerDutyContextBuilder holds PagerDutyContext struct and provides a builder API.
type PagerDutyContextBuilder struct {
	v *PagerDutyContext
}

// NewPagerDutyContext provides a builder for the PagerDutyContext struct.
func NewPagerDutyContextBuilder() *PagerDutyContextBuilder {
	r := PagerDutyContextBuilder{
		&PagerDutyContext{},
	}

	return &r
}

// Build finalize the chain and returns the PagerDutyContext struct
func (rb *PagerDutyContextBuilder) Build() PagerDutyContext {
	return *rb.v
}

func (rb *PagerDutyContextBuilder) Href(href string) *PagerDutyContextBuilder {
	rb.v.Href = &href
	return rb
}

func (rb *PagerDutyContextBuilder) Src(src string) *PagerDutyContextBuilder {
	rb.v.Src = &src
	return rb
}

func (rb *PagerDutyContextBuilder) Type_(type_ pagerdutycontexttype.PagerDutyContextType) *PagerDutyContextBuilder {
	rb.v.Type = type_
	return rb
}
