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

// DateIndexNameProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L164-L177
type DateIndexNameProcessor struct {
	DateFormats []string `json:"date_formats"`
	// DateRounding How to round the date when formatting the date into the index name. Valid
	// values are:
	// `y` (year), `M` (month), `w` (week), `d` (day), `h` (hour), `m` (minute) and
	// `s` (second).
	// Supports template snippets.
	DateRounding    string               `json:"date_rounding"`
	Field           Field                `json:"field"`
	If              *string              `json:"if,omitempty"`
	IgnoreFailure   *bool                `json:"ignore_failure,omitempty"`
	IndexNameFormat string               `json:"index_name_format"`
	IndexNamePrefix string               `json:"index_name_prefix"`
	Locale          string               `json:"locale"`
	OnFailure       []ProcessorContainer `json:"on_failure,omitempty"`
	Tag             *string              `json:"tag,omitempty"`
	Timezone        string               `json:"timezone"`
}

// DateIndexNameProcessorBuilder holds DateIndexNameProcessor struct and provides a builder API.
type DateIndexNameProcessorBuilder struct {
	v *DateIndexNameProcessor
}

// NewDateIndexNameProcessor provides a builder for the DateIndexNameProcessor struct.
func NewDateIndexNameProcessorBuilder() *DateIndexNameProcessorBuilder {
	r := DateIndexNameProcessorBuilder{
		&DateIndexNameProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the DateIndexNameProcessor struct
func (rb *DateIndexNameProcessorBuilder) Build() DateIndexNameProcessor {
	return *rb.v
}

func (rb *DateIndexNameProcessorBuilder) DateFormats(date_formats ...string) *DateIndexNameProcessorBuilder {
	rb.v.DateFormats = date_formats
	return rb
}

// DateRounding How to round the date when formatting the date into the index name. Valid
// values are:
// `y` (year), `M` (month), `w` (week), `d` (day), `h` (hour), `m` (minute) and
// `s` (second).
// Supports template snippets.

func (rb *DateIndexNameProcessorBuilder) DateRounding(daterounding string) *DateIndexNameProcessorBuilder {
	rb.v.DateRounding = daterounding
	return rb
}

func (rb *DateIndexNameProcessorBuilder) Field(field Field) *DateIndexNameProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *DateIndexNameProcessorBuilder) If_(if_ string) *DateIndexNameProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *DateIndexNameProcessorBuilder) IgnoreFailure(ignorefailure bool) *DateIndexNameProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *DateIndexNameProcessorBuilder) IndexNameFormat(indexnameformat string) *DateIndexNameProcessorBuilder {
	rb.v.IndexNameFormat = indexnameformat
	return rb
}

func (rb *DateIndexNameProcessorBuilder) IndexNamePrefix(indexnameprefix string) *DateIndexNameProcessorBuilder {
	rb.v.IndexNamePrefix = indexnameprefix
	return rb
}

func (rb *DateIndexNameProcessorBuilder) Locale(locale string) *DateIndexNameProcessorBuilder {
	rb.v.Locale = locale
	return rb
}

func (rb *DateIndexNameProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *DateIndexNameProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *DateIndexNameProcessorBuilder) Tag(tag string) *DateIndexNameProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *DateIndexNameProcessorBuilder) Timezone(timezone string) *DateIndexNameProcessorBuilder {
	rb.v.Timezone = timezone
	return rb
}
