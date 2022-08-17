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


package validate

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package validate
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/validate/MlValidateJobRequest.ts#L27-L45
type Request struct {
	AnalysisConfig *types.AnalysisConfig `json:"analysis_config,omitempty"`

	AnalysisLimits *types.AnalysisLimits `json:"analysis_limits,omitempty"`

	DataDescription *types.DataDescription `json:"data_description,omitempty"`

	Description *string `json:"description,omitempty"`

	JobId *types.Id `json:"job_id,omitempty"`

	ModelPlot *types.ModelPlotConfig `json:"model_plot,omitempty"`

	ModelSnapshotId *types.Id `json:"model_snapshot_id,omitempty"`

	ModelSnapshotRetentionDays *int64 `json:"model_snapshot_retention_days,omitempty"`

	ResultsIndexName *types.IndexName `json:"results_index_name,omitempty"`
}

// RequestBuilder is the builder API for the validate.Request
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
		return nil, fmt.Errorf("could not deserialise json into Validate request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) AnalysisConfig(analysisconfig *types.AnalysisConfigBuilder) *RequestBuilder {
	v := analysisconfig.Build()
	rb.v.AnalysisConfig = &v
	return rb
}

func (rb *RequestBuilder) AnalysisLimits(analysislimits *types.AnalysisLimitsBuilder) *RequestBuilder {
	v := analysislimits.Build()
	rb.v.AnalysisLimits = &v
	return rb
}

func (rb *RequestBuilder) DataDescription(datadescription *types.DataDescriptionBuilder) *RequestBuilder {
	v := datadescription.Build()
	rb.v.DataDescription = &v
	return rb
}

func (rb *RequestBuilder) Description(description string) *RequestBuilder {
	rb.v.Description = &description
	return rb
}

func (rb *RequestBuilder) JobId(jobid types.Id) *RequestBuilder {
	rb.v.JobId = &jobid
	return rb
}

func (rb *RequestBuilder) ModelPlot(modelplot *types.ModelPlotConfigBuilder) *RequestBuilder {
	v := modelplot.Build()
	rb.v.ModelPlot = &v
	return rb
}

func (rb *RequestBuilder) ModelSnapshotId(modelsnapshotid types.Id) *RequestBuilder {
	rb.v.ModelSnapshotId = &modelsnapshotid
	return rb
}

func (rb *RequestBuilder) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *RequestBuilder {
	rb.v.ModelSnapshotRetentionDays = &modelsnapshotretentiondays
	return rb
}

func (rb *RequestBuilder) ResultsIndexName(resultsindexname types.IndexName) *RequestBuilder {
	rb.v.ResultsIndexName = &resultsindexname
	return rb
}
