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

// GeoIpProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L105-L112
type GeoIpProcessor struct {
	DatabaseFile  string               `json:"database_file"`
	Field         Field                `json:"field"`
	FirstOnly     bool                 `json:"first_only"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing bool                 `json:"ignore_missing"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Properties    []string             `json:"properties"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   Field                `json:"target_field"`
}

// GeoIpProcessorBuilder holds GeoIpProcessor struct and provides a builder API.
type GeoIpProcessorBuilder struct {
	v *GeoIpProcessor
}

// NewGeoIpProcessor provides a builder for the GeoIpProcessor struct.
func NewGeoIpProcessorBuilder() *GeoIpProcessorBuilder {
	r := GeoIpProcessorBuilder{
		&GeoIpProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the GeoIpProcessor struct
func (rb *GeoIpProcessorBuilder) Build() GeoIpProcessor {
	return *rb.v
}

func (rb *GeoIpProcessorBuilder) DatabaseFile(databasefile string) *GeoIpProcessorBuilder {
	rb.v.DatabaseFile = databasefile
	return rb
}

func (rb *GeoIpProcessorBuilder) Field(field Field) *GeoIpProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *GeoIpProcessorBuilder) FirstOnly(firstonly bool) *GeoIpProcessorBuilder {
	rb.v.FirstOnly = firstonly
	return rb
}

func (rb *GeoIpProcessorBuilder) If_(if_ string) *GeoIpProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *GeoIpProcessorBuilder) IgnoreFailure(ignorefailure bool) *GeoIpProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *GeoIpProcessorBuilder) IgnoreMissing(ignoremissing bool) *GeoIpProcessorBuilder {
	rb.v.IgnoreMissing = ignoremissing
	return rb
}

func (rb *GeoIpProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *GeoIpProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *GeoIpProcessorBuilder) Properties(properties ...string) *GeoIpProcessorBuilder {
	rb.v.Properties = properties
	return rb
}

func (rb *GeoIpProcessorBuilder) Tag(tag string) *GeoIpProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *GeoIpProcessorBuilder) TargetField(targetfield Field) *GeoIpProcessorBuilder {
	rb.v.TargetField = targetfield
	return rb
}
