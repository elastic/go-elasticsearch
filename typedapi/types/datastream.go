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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
)

// DataStream type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/DataStream.ts#L31-L46
type DataStream struct {
	AllowCustomRouting *bool                     `json:"allow_custom_routing,omitempty"`
	Generation         int                       `json:"generation"`
	Hidden             bool                      `json:"hidden"`
	IlmPolicy          *Name                     `json:"ilm_policy,omitempty"`
	Indices            []DataStreamIndex         `json:"indices"`
	Meta_              *Metadata                 `json:"_meta,omitempty"`
	Name               DataStreamName            `json:"name"`
	Replicated         *bool                     `json:"replicated,omitempty"`
	Status             healthstatus.HealthStatus `json:"status"`
	System             *bool                     `json:"system,omitempty"`
	Template           Name                      `json:"template"`
	TimestampField     DataStreamTimestampField  `json:"timestamp_field"`
}

// DataStreamBuilder holds DataStream struct and provides a builder API.
type DataStreamBuilder struct {
	v *DataStream
}

// NewDataStream provides a builder for the DataStream struct.
func NewDataStreamBuilder() *DataStreamBuilder {
	r := DataStreamBuilder{
		&DataStream{},
	}

	return &r
}

// Build finalize the chain and returns the DataStream struct
func (rb *DataStreamBuilder) Build() DataStream {
	return *rb.v
}

func (rb *DataStreamBuilder) AllowCustomRouting(allowcustomrouting bool) *DataStreamBuilder {
	rb.v.AllowCustomRouting = &allowcustomrouting
	return rb
}

func (rb *DataStreamBuilder) Generation(generation int) *DataStreamBuilder {
	rb.v.Generation = generation
	return rb
}

func (rb *DataStreamBuilder) Hidden(hidden bool) *DataStreamBuilder {
	rb.v.Hidden = hidden
	return rb
}

func (rb *DataStreamBuilder) IlmPolicy(ilmpolicy Name) *DataStreamBuilder {
	rb.v.IlmPolicy = &ilmpolicy
	return rb
}

func (rb *DataStreamBuilder) Indices(indices []DataStreamIndexBuilder) *DataStreamBuilder {
	tmp := make([]DataStreamIndex, len(indices))
	for _, value := range indices {
		tmp = append(tmp, value.Build())
	}
	rb.v.Indices = tmp
	return rb
}

func (rb *DataStreamBuilder) Meta_(meta_ *MetadataBuilder) *DataStreamBuilder {
	v := meta_.Build()
	rb.v.Meta_ = &v
	return rb
}

func (rb *DataStreamBuilder) Name(name DataStreamName) *DataStreamBuilder {
	rb.v.Name = name
	return rb
}

func (rb *DataStreamBuilder) Replicated(replicated bool) *DataStreamBuilder {
	rb.v.Replicated = &replicated
	return rb
}

func (rb *DataStreamBuilder) Status(status healthstatus.HealthStatus) *DataStreamBuilder {
	rb.v.Status = status
	return rb
}

func (rb *DataStreamBuilder) System(system bool) *DataStreamBuilder {
	rb.v.System = &system
	return rb
}

func (rb *DataStreamBuilder) Template(template Name) *DataStreamBuilder {
	rb.v.Template = template
	return rb
}

func (rb *DataStreamBuilder) TimestampField(timestampfield *DataStreamTimestampFieldBuilder) *DataStreamBuilder {
	v := timestampfield.Build()
	rb.v.TimestampField = v
	return rb
}
