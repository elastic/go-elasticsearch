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

// NodeInfoIngestInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L123-L125
type NodeInfoIngestInfo struct {
	Downloader NodeInfoIngestDownloader `json:"downloader"`
}

// NodeInfoIngestInfoBuilder holds NodeInfoIngestInfo struct and provides a builder API.
type NodeInfoIngestInfoBuilder struct {
	v *NodeInfoIngestInfo
}

// NewNodeInfoIngestInfo provides a builder for the NodeInfoIngestInfo struct.
func NewNodeInfoIngestInfoBuilder() *NodeInfoIngestInfoBuilder {
	r := NodeInfoIngestInfoBuilder{
		&NodeInfoIngestInfo{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoIngestInfo struct
func (rb *NodeInfoIngestInfoBuilder) Build() NodeInfoIngestInfo {
	return *rb.v
}

func (rb *NodeInfoIngestInfoBuilder) Downloader(downloader *NodeInfoIngestDownloaderBuilder) *NodeInfoIngestInfoBuilder {
	v := downloader.Build()
	rb.v.Downloader = v
	return rb
}
