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

// ClusterIngest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L144-L147
type ClusterIngest struct {
	NumberOfPipelines int                         `json:"number_of_pipelines"`
	ProcessorStats    map[string]ClusterProcessor `json:"processor_stats"`
}

// ClusterIngestBuilder holds ClusterIngest struct and provides a builder API.
type ClusterIngestBuilder struct {
	v *ClusterIngest
}

// NewClusterIngest provides a builder for the ClusterIngest struct.
func NewClusterIngestBuilder() *ClusterIngestBuilder {
	r := ClusterIngestBuilder{
		&ClusterIngest{
			ProcessorStats: make(map[string]ClusterProcessor, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ClusterIngest struct
func (rb *ClusterIngestBuilder) Build() ClusterIngest {
	return *rb.v
}

func (rb *ClusterIngestBuilder) NumberOfPipelines(numberofpipelines int) *ClusterIngestBuilder {
	rb.v.NumberOfPipelines = numberofpipelines
	return rb
}

func (rb *ClusterIngestBuilder) ProcessorStats(values map[string]*ClusterProcessorBuilder) *ClusterIngestBuilder {
	tmp := make(map[string]ClusterProcessor, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.ProcessorStats = tmp
	return rb
}
