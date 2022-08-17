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

// ElasticsearchVersionInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Base.ts#L54-L64
type ElasticsearchVersionInfo struct {
	BuildDate                        DateTime      `json:"build_date"`
	BuildFlavor                      string        `json:"build_flavor"`
	BuildHash                        string        `json:"build_hash"`
	BuildSnapshot                    bool          `json:"build_snapshot"`
	BuildType                        string        `json:"build_type"`
	Int                              string        `json:"number"`
	LuceneVersion                    VersionString `json:"lucene_version"`
	MinimumIndexCompatibilityVersion VersionString `json:"minimum_index_compatibility_version"`
	MinimumWireCompatibilityVersion  VersionString `json:"minimum_wire_compatibility_version"`
}

// ElasticsearchVersionInfoBuilder holds ElasticsearchVersionInfo struct and provides a builder API.
type ElasticsearchVersionInfoBuilder struct {
	v *ElasticsearchVersionInfo
}

// NewElasticsearchVersionInfo provides a builder for the ElasticsearchVersionInfo struct.
func NewElasticsearchVersionInfoBuilder() *ElasticsearchVersionInfoBuilder {
	r := ElasticsearchVersionInfoBuilder{
		&ElasticsearchVersionInfo{},
	}

	return &r
}

// Build finalize the chain and returns the ElasticsearchVersionInfo struct
func (rb *ElasticsearchVersionInfoBuilder) Build() ElasticsearchVersionInfo {
	return *rb.v
}

func (rb *ElasticsearchVersionInfoBuilder) BuildDate(builddate *DateTimeBuilder) *ElasticsearchVersionInfoBuilder {
	v := builddate.Build()
	rb.v.BuildDate = v
	return rb
}

func (rb *ElasticsearchVersionInfoBuilder) BuildFlavor(buildflavor string) *ElasticsearchVersionInfoBuilder {
	rb.v.BuildFlavor = buildflavor
	return rb
}

func (rb *ElasticsearchVersionInfoBuilder) BuildHash(buildhash string) *ElasticsearchVersionInfoBuilder {
	rb.v.BuildHash = buildhash
	return rb
}

func (rb *ElasticsearchVersionInfoBuilder) BuildSnapshot(buildsnapshot bool) *ElasticsearchVersionInfoBuilder {
	rb.v.BuildSnapshot = buildsnapshot
	return rb
}

func (rb *ElasticsearchVersionInfoBuilder) BuildType(buildtype string) *ElasticsearchVersionInfoBuilder {
	rb.v.BuildType = buildtype
	return rb
}

func (rb *ElasticsearchVersionInfoBuilder) Int(int string) *ElasticsearchVersionInfoBuilder {
	rb.v.Int = int
	return rb
}

func (rb *ElasticsearchVersionInfoBuilder) LuceneVersion(luceneversion VersionString) *ElasticsearchVersionInfoBuilder {
	rb.v.LuceneVersion = luceneversion
	return rb
}

func (rb *ElasticsearchVersionInfoBuilder) MinimumIndexCompatibilityVersion(minimumindexcompatibilityversion VersionString) *ElasticsearchVersionInfoBuilder {
	rb.v.MinimumIndexCompatibilityVersion = minimumindexcompatibilityversion
	return rb
}

func (rb *ElasticsearchVersionInfoBuilder) MinimumWireCompatibilityVersion(minimumwirecompatibilityversion VersionString) *ElasticsearchVersionInfoBuilder {
	rb.v.MinimumWireCompatibilityVersion = minimumwirecompatibilityversion
	return rb
}
