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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dataframestate"
)

// DataframeAnalytics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L324-L341
type DataframeAnalytics struct {
	// AnalysisStats An object containing information about the analysis job.
	AnalysisStats *DataframeAnalyticsStatsContainer `json:"analysis_stats,omitempty"`
	// AssignmentExplanation For running jobs only, contains messages relating to the selection of a node
	// to run the job.
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`
	// DataCounts An object that provides counts for the quantity of documents skipped, used in
	// training, or available for testing.
	DataCounts DataframeAnalyticsStatsDataCounts `json:"data_counts"`
	// Id The unique identifier of the data frame analytics job.
	Id Id `json:"id"`
	// MemoryUsage An object describing memory usage of the analytics. It is present only after
	// the job is started and memory usage is reported.
	MemoryUsage DataframeAnalyticsStatsMemoryUsage `json:"memory_usage"`
	// Node Contains properties for the node that runs the job. This information is
	// available only for running jobs.
	Node *NodeAttributes `json:"node,omitempty"`
	// Progress The progress report of the data frame analytics job by phase.
	Progress []DataframeAnalyticsStatsProgress `json:"progress"`
	// State The status of the data frame analytics job, which can be one of the following
	// values: failed, started, starting, stopping, stopped.
	State dataframestate.DataframeState `json:"state"`
}

// DataframeAnalyticsBuilder holds DataframeAnalytics struct and provides a builder API.
type DataframeAnalyticsBuilder struct {
	v *DataframeAnalytics
}

// NewDataframeAnalytics provides a builder for the DataframeAnalytics struct.
func NewDataframeAnalyticsBuilder() *DataframeAnalyticsBuilder {
	r := DataframeAnalyticsBuilder{
		&DataframeAnalytics{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalytics struct
func (rb *DataframeAnalyticsBuilder) Build() DataframeAnalytics {
	return *rb.v
}

// AnalysisStats An object containing information about the analysis job.

func (rb *DataframeAnalyticsBuilder) AnalysisStats(analysisstats *DataframeAnalyticsStatsContainerBuilder) *DataframeAnalyticsBuilder {
	v := analysisstats.Build()
	rb.v.AnalysisStats = &v
	return rb
}

// AssignmentExplanation For running jobs only, contains messages relating to the selection of a node
// to run the job.

func (rb *DataframeAnalyticsBuilder) AssignmentExplanation(assignmentexplanation string) *DataframeAnalyticsBuilder {
	rb.v.AssignmentExplanation = &assignmentexplanation
	return rb
}

// DataCounts An object that provides counts for the quantity of documents skipped, used in
// training, or available for testing.

func (rb *DataframeAnalyticsBuilder) DataCounts(datacounts *DataframeAnalyticsStatsDataCountsBuilder) *DataframeAnalyticsBuilder {
	v := datacounts.Build()
	rb.v.DataCounts = v
	return rb
}

// Id The unique identifier of the data frame analytics job.

func (rb *DataframeAnalyticsBuilder) Id(id Id) *DataframeAnalyticsBuilder {
	rb.v.Id = id
	return rb
}

// MemoryUsage An object describing memory usage of the analytics. It is present only after
// the job is started and memory usage is reported.

func (rb *DataframeAnalyticsBuilder) MemoryUsage(memoryusage *DataframeAnalyticsStatsMemoryUsageBuilder) *DataframeAnalyticsBuilder {
	v := memoryusage.Build()
	rb.v.MemoryUsage = v
	return rb
}

// Node Contains properties for the node that runs the job. This information is
// available only for running jobs.

func (rb *DataframeAnalyticsBuilder) Node(node *NodeAttributesBuilder) *DataframeAnalyticsBuilder {
	v := node.Build()
	rb.v.Node = &v
	return rb
}

// Progress The progress report of the data frame analytics job by phase.

func (rb *DataframeAnalyticsBuilder) Progress(progress []DataframeAnalyticsStatsProgressBuilder) *DataframeAnalyticsBuilder {
	tmp := make([]DataframeAnalyticsStatsProgress, len(progress))
	for _, value := range progress {
		tmp = append(tmp, value.Build())
	}
	rb.v.Progress = tmp
	return rb
}

// State The status of the data frame analytics job, which can be one of the following
// values: failed, started, starting, stopping, stopped.

func (rb *DataframeAnalyticsBuilder) State(state dataframestate.DataframeState) *DataframeAnalyticsBuilder {
	rb.v.State = state
	return rb
}
