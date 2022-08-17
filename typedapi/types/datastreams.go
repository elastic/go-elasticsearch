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

// DataStreams type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L76-L79
type DataStreams struct {
	Available    bool  `json:"available"`
	DataStreams  int64 `json:"data_streams"`
	Enabled      bool  `json:"enabled"`
	IndicesCount int64 `json:"indices_count"`
}

// DataStreamsBuilder holds DataStreams struct and provides a builder API.
type DataStreamsBuilder struct {
	v *DataStreams
}

// NewDataStreams provides a builder for the DataStreams struct.
func NewDataStreamsBuilder() *DataStreamsBuilder {
	r := DataStreamsBuilder{
		&DataStreams{},
	}

	return &r
}

// Build finalize the chain and returns the DataStreams struct
func (rb *DataStreamsBuilder) Build() DataStreams {
	return *rb.v
}

func (rb *DataStreamsBuilder) Available(available bool) *DataStreamsBuilder {
	rb.v.Available = available
	return rb
}

func (rb *DataStreamsBuilder) DataStreams(datastreams int64) *DataStreamsBuilder {
	rb.v.DataStreams = datastreams
	return rb
}

func (rb *DataStreamsBuilder) Enabled(enabled bool) *DataStreamsBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *DataStreamsBuilder) IndicesCount(indicescount int64) *DataStreamsBuilder {
	rb.v.IndicesCount = indicescount
	return rb
}
