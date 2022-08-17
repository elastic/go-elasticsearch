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

// Settings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/_types/Transform.ts#L98-L133
type Settings struct {
	// AlignCheckpoints Specifies whether the transform checkpoint ranges should be optimized for
	// performance. Such optimization can align
	// checkpoint ranges with the date histogram interval when date histogram is
	// specified as a group source in the
	// transform config. As a result, less document updates in the destination index
	// will be performed thus improving
	// overall performance.
	AlignCheckpoints *bool `json:"align_checkpoints,omitempty"`
	// DatesAsEpochMillis Defines if dates in the ouput should be written as ISO formatted string or as
	// millis since epoch. epoch_millis was
	// the default for transforms created before version 7.11. For compatible output
	// set this value to `true`.
	DatesAsEpochMillis *bool `json:"dates_as_epoch_millis,omitempty"`
	// DeduceMappings Specifies whether the transform should deduce the destination index mappings
	// from the transform configuration.
	DeduceMappings *bool `json:"deduce_mappings,omitempty"`
	// DocsPerSecond Specifies a limit on the number of input documents per second. This setting
	// throttles the transform by adding a
	// wait time between search requests. The default value is null, which disables
	// throttling.
	DocsPerSecond *float32 `json:"docs_per_second,omitempty"`
	// MaxPageSearchSize Defines the initial page size to use for the composite aggregation for each
	// checkpoint. If circuit breaker
	// exceptions occur, the page size is dynamically adjusted to a lower value. The
	// minimum value is `10` and the
	// maximum is `65,536`.
	MaxPageSearchSize *int `json:"max_page_search_size,omitempty"`
}

// SettingsBuilder holds Settings struct and provides a builder API.
type SettingsBuilder struct {
	v *Settings
}

// NewSettings provides a builder for the Settings struct.
func NewSettingsBuilder() *SettingsBuilder {
	r := SettingsBuilder{
		&Settings{},
	}

	return &r
}

// Build finalize the chain and returns the Settings struct
func (rb *SettingsBuilder) Build() Settings {
	return *rb.v
}

// AlignCheckpoints Specifies whether the transform checkpoint ranges should be optimized for
// performance. Such optimization can align
// checkpoint ranges with the date histogram interval when date histogram is
// specified as a group source in the
// transform config. As a result, less document updates in the destination index
// will be performed thus improving
// overall performance.

func (rb *SettingsBuilder) AlignCheckpoints(aligncheckpoints bool) *SettingsBuilder {
	rb.v.AlignCheckpoints = &aligncheckpoints
	return rb
}

// DatesAsEpochMillis Defines if dates in the ouput should be written as ISO formatted string or as
// millis since epoch. epoch_millis was
// the default for transforms created before version 7.11. For compatible output
// set this value to `true`.

func (rb *SettingsBuilder) DatesAsEpochMillis(datesasepochmillis bool) *SettingsBuilder {
	rb.v.DatesAsEpochMillis = &datesasepochmillis
	return rb
}

// DeduceMappings Specifies whether the transform should deduce the destination index mappings
// from the transform configuration.

func (rb *SettingsBuilder) DeduceMappings(deducemappings bool) *SettingsBuilder {
	rb.v.DeduceMappings = &deducemappings
	return rb
}

// DocsPerSecond Specifies a limit on the number of input documents per second. This setting
// throttles the transform by adding a
// wait time between search requests. The default value is null, which disables
// throttling.

func (rb *SettingsBuilder) DocsPerSecond(docspersecond float32) *SettingsBuilder {
	rb.v.DocsPerSecond = &docspersecond
	return rb
}

// MaxPageSearchSize Defines the initial page size to use for the composite aggregation for each
// checkpoint. If circuit breaker
// exceptions occur, the page size is dynamically adjusted to a lower value. The
// minimum value is `10` and the
// maximum is `65,536`.

func (rb *SettingsBuilder) MaxPageSearchSize(maxpagesearchsize int) *SettingsBuilder {
	rb.v.MaxPageSearchSize = &maxpagesearchsize
	return rb
}
