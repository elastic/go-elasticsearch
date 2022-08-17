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


package updatetransform

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package updatetransform
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/update_transform/UpdateTransformRequest.ts#L31-L105
type Request struct {

	// Description Free text description of the transform.
	Description *string `json:"description,omitempty"`

	// Dest The destination for the transform.
	Dest *types.Destination `json:"dest,omitempty"`

	// Frequency The interval between checks for changes in the source indices when the
	// transform is running continuously. Also determines the retry interval in
	// the event of transient failures while the transform is searching or
	// indexing. The minimum value is 1s and the maximum is 1h.
	Frequency *types.Duration `json:"frequency,omitempty"`

	// Meta_ Defines optional transform metadata.
	Meta_ *types.Metadata `json:"_meta,omitempty"`

	// RetentionPolicy Defines a retention policy for the transform. Data that meets the defined
	// criteria is deleted from the destination index.
	RetentionPolicy types.RetentionPolicyContainer `json:"retention_policy,omitempty"`

	// Settings Defines optional transform settings.
	Settings *types.Settings `json:"settings,omitempty"`

	// Source The source of the data for the transform.
	Source *types.Source `json:"source,omitempty"`

	// Sync Defines the properties transforms require to run continuously.
	Sync *types.SyncContainer `json:"sync,omitempty"`
}

// RequestBuilder is the builder API for the updatetransform.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Updatetransform request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Description(description string) *RequestBuilder {
	rb.v.Description = &description
	return rb
}

func (rb *RequestBuilder) Dest(dest *types.DestinationBuilder) *RequestBuilder {
	v := dest.Build()
	rb.v.Dest = &v
	return rb
}

func (rb *RequestBuilder) Frequency(frequency *types.DurationBuilder) *RequestBuilder {
	v := frequency.Build()
	rb.v.Frequency = &v
	return rb
}

func (rb *RequestBuilder) Meta_(meta_ *types.MetadataBuilder) *RequestBuilder {
	v := meta_.Build()
	rb.v.Meta_ = &v
	return rb
}

func (rb *RequestBuilder) RetentionPolicy(retentionpolicy types.RetentionPolicyContainer) *RequestBuilder {
	rb.v.RetentionPolicy = retentionpolicy
	return rb
}

func (rb *RequestBuilder) Settings(settings *types.SettingsBuilder) *RequestBuilder {
	v := settings.Build()
	rb.v.Settings = &v
	return rb
}

func (rb *RequestBuilder) Source(source *types.SourceBuilder) *RequestBuilder {
	v := source.Build()
	rb.v.Source = &v
	return rb
}

func (rb *RequestBuilder) Sync(sync *types.SyncContainerBuilder) *RequestBuilder {
	v := sync.Build()
	rb.v.Sync = &v
	return rb
}
