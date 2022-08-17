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

// DataframeAnalyticsStatsHyperparameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L380-L387
type DataframeAnalyticsStatsHyperparameters struct {
	Hyperparameters Hyperparameters `json:"hyperparameters"`
	// Iteration The number of iterations on the analysis.
	Iteration      int                 `json:"iteration"`
	Timestamp      EpochTimeUnitMillis `json:"timestamp"`
	TimingStats    TimingStats         `json:"timing_stats"`
	ValidationLoss ValidationLoss      `json:"validation_loss"`
}

// DataframeAnalyticsStatsHyperparametersBuilder holds DataframeAnalyticsStatsHyperparameters struct and provides a builder API.
type DataframeAnalyticsStatsHyperparametersBuilder struct {
	v *DataframeAnalyticsStatsHyperparameters
}

// NewDataframeAnalyticsStatsHyperparameters provides a builder for the DataframeAnalyticsStatsHyperparameters struct.
func NewDataframeAnalyticsStatsHyperparametersBuilder() *DataframeAnalyticsStatsHyperparametersBuilder {
	r := DataframeAnalyticsStatsHyperparametersBuilder{
		&DataframeAnalyticsStatsHyperparameters{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsStatsHyperparameters struct
func (rb *DataframeAnalyticsStatsHyperparametersBuilder) Build() DataframeAnalyticsStatsHyperparameters {
	return *rb.v
}

func (rb *DataframeAnalyticsStatsHyperparametersBuilder) Hyperparameters(hyperparameters *HyperparametersBuilder) *DataframeAnalyticsStatsHyperparametersBuilder {
	v := hyperparameters.Build()
	rb.v.Hyperparameters = v
	return rb
}

// Iteration The number of iterations on the analysis.

func (rb *DataframeAnalyticsStatsHyperparametersBuilder) Iteration(iteration int) *DataframeAnalyticsStatsHyperparametersBuilder {
	rb.v.Iteration = iteration
	return rb
}

func (rb *DataframeAnalyticsStatsHyperparametersBuilder) Timestamp(timestamp *EpochTimeUnitMillisBuilder) *DataframeAnalyticsStatsHyperparametersBuilder {
	v := timestamp.Build()
	rb.v.Timestamp = v
	return rb
}

func (rb *DataframeAnalyticsStatsHyperparametersBuilder) TimingStats(timingstats *TimingStatsBuilder) *DataframeAnalyticsStatsHyperparametersBuilder {
	v := timingstats.Build()
	rb.v.TimingStats = v
	return rb
}

func (rb *DataframeAnalyticsStatsHyperparametersBuilder) ValidationLoss(validationloss *ValidationLossBuilder) *DataframeAnalyticsStatsHyperparametersBuilder {
	v := validationloss.Build()
	rb.v.ValidationLoss = v
	return rb
}
