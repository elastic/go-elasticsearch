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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

// JobUsage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/xpack/usage/types.ts#L346-L352
type JobUsage struct {
	Count     int              `json:"count"`
	CreatedBy map[string]int64 `json:"created_by"`
	Detectors JobStatistics    `json:"detectors"`
	Forecasts MlJobForecasts   `json:"forecasts"`
	ModelSize JobStatistics    `json:"model_size"`
}

// JobUsageBuilder holds JobUsage struct and provides a builder API.
type JobUsageBuilder struct {
	v *JobUsage
}

// NewJobUsage provides a builder for the JobUsage struct.
func NewJobUsageBuilder() *JobUsageBuilder {
	r := JobUsageBuilder{
		&JobUsage{
			CreatedBy: make(map[string]int64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the JobUsage struct
func (rb *JobUsageBuilder) Build() JobUsage {
	return *rb.v
}

func (rb *JobUsageBuilder) Count(count int) *JobUsageBuilder {
	rb.v.Count = count
	return rb
}

func (rb *JobUsageBuilder) CreatedBy(value map[string]int64) *JobUsageBuilder {
	rb.v.CreatedBy = value
	return rb
}

func (rb *JobUsageBuilder) Detectors(detectors *JobStatisticsBuilder) *JobUsageBuilder {
	v := detectors.Build()
	rb.v.Detectors = v
	return rb
}

func (rb *JobUsageBuilder) Forecasts(forecasts *MlJobForecastsBuilder) *JobUsageBuilder {
	v := forecasts.Build()
	rb.v.Forecasts = v
	return rb
}

func (rb *JobUsageBuilder) ModelSize(modelsize *JobStatisticsBuilder) *JobUsageBuilder {
	v := modelsize.Build()
	rb.v.ModelSize = v
	return rb
}
