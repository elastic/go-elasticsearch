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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dataattachmentformat"
)

// DataEmailAttachment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L234-L236
type DataEmailAttachment struct {
	Format *dataattachmentformat.DataAttachmentFormat `json:"format,omitempty"`
}

// DataEmailAttachmentBuilder holds DataEmailAttachment struct and provides a builder API.
type DataEmailAttachmentBuilder struct {
	v *DataEmailAttachment
}

// NewDataEmailAttachment provides a builder for the DataEmailAttachment struct.
func NewDataEmailAttachmentBuilder() *DataEmailAttachmentBuilder {
	r := DataEmailAttachmentBuilder{
		&DataEmailAttachment{},
	}

	return &r
}

// Build finalize the chain and returns the DataEmailAttachment struct
func (rb *DataEmailAttachmentBuilder) Build() DataEmailAttachment {
	return *rb.v
}

func (rb *DataEmailAttachmentBuilder) Format(format dataattachmentformat.DataAttachmentFormat) *DataEmailAttachmentBuilder {
	rb.v.Format = &format
	return rb
}
