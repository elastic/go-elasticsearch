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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/holtwinterstype"
)

// HoltWintersModelSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/aggregations/pipeline.ts#L235-L242
type HoltWintersModelSettings struct {
	Alpha  *float32                         `json:"alpha,omitempty"`
	Beta   *float32                         `json:"beta,omitempty"`
	Gamma  *float32                         `json:"gamma,omitempty"`
	Pad    *bool                            `json:"pad,omitempty"`
	Period *int                             `json:"period,omitempty"`
	Type   *holtwinterstype.HoltWintersType `json:"type,omitempty"`
}

// HoltWintersModelSettingsBuilder holds HoltWintersModelSettings struct and provides a builder API.
type HoltWintersModelSettingsBuilder struct {
	v *HoltWintersModelSettings
}

// NewHoltWintersModelSettings provides a builder for the HoltWintersModelSettings struct.
func NewHoltWintersModelSettingsBuilder() *HoltWintersModelSettingsBuilder {
	r := HoltWintersModelSettingsBuilder{
		&HoltWintersModelSettings{},
	}

	return &r
}

// Build finalize the chain and returns the HoltWintersModelSettings struct
func (rb *HoltWintersModelSettingsBuilder) Build() HoltWintersModelSettings {
	return *rb.v
}

func (rb *HoltWintersModelSettingsBuilder) Alpha(alpha float32) *HoltWintersModelSettingsBuilder {
	rb.v.Alpha = &alpha
	return rb
}

func (rb *HoltWintersModelSettingsBuilder) Beta(beta float32) *HoltWintersModelSettingsBuilder {
	rb.v.Beta = &beta
	return rb
}

func (rb *HoltWintersModelSettingsBuilder) Gamma(gamma float32) *HoltWintersModelSettingsBuilder {
	rb.v.Gamma = &gamma
	return rb
}

func (rb *HoltWintersModelSettingsBuilder) Pad(pad bool) *HoltWintersModelSettingsBuilder {
	rb.v.Pad = &pad
	return rb
}

func (rb *HoltWintersModelSettingsBuilder) Period(period int) *HoltWintersModelSettingsBuilder {
	rb.v.Period = &period
	return rb
}

func (rb *HoltWintersModelSettingsBuilder) Type_(type_ holtwinterstype.HoltWintersType) *HoltWintersModelSettingsBuilder {
	rb.v.Type = &type_
	return rb
}
