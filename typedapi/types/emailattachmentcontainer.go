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

// EmailAttachmentContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L211-L216
type EmailAttachmentContainer struct {
	Data      *DataEmailAttachment      `json:"data,omitempty"`
	Http      *HttpEmailAttachment      `json:"http,omitempty"`
	Reporting *ReportingEmailAttachment `json:"reporting,omitempty"`
}

// EmailAttachmentContainerBuilder holds EmailAttachmentContainer struct and provides a builder API.
type EmailAttachmentContainerBuilder struct {
	v *EmailAttachmentContainer
}

// NewEmailAttachmentContainer provides a builder for the EmailAttachmentContainer struct.
func NewEmailAttachmentContainerBuilder() *EmailAttachmentContainerBuilder {
	r := EmailAttachmentContainerBuilder{
		&EmailAttachmentContainer{},
	}

	return &r
}

// Build finalize the chain and returns the EmailAttachmentContainer struct
func (rb *EmailAttachmentContainerBuilder) Build() EmailAttachmentContainer {
	return *rb.v
}

func (rb *EmailAttachmentContainerBuilder) Data(data *DataEmailAttachmentBuilder) *EmailAttachmentContainerBuilder {
	v := data.Build()
	rb.v.Data = &v
	return rb
}

func (rb *EmailAttachmentContainerBuilder) Http(http *HttpEmailAttachmentBuilder) *EmailAttachmentContainerBuilder {
	v := http.Build()
	rb.v.Http = &v
	return rb
}

func (rb *EmailAttachmentContainerBuilder) Reporting(reporting *ReportingEmailAttachmentBuilder) *EmailAttachmentContainerBuilder {
	v := reporting.Build()
	rb.v.Reporting = &v
	return rb
}
