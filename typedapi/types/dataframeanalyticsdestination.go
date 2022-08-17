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

// DataframeAnalyticsDestination type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L77-L82
type DataframeAnalyticsDestination struct {
	// Index Defines the destination index to store the results of the data frame
	// analytics job.
	Index IndexName `json:"index"`
	// ResultsField Defines the name of the field in which to store the results of the analysis.
	// Defaults to `ml`.
	ResultsField *Field `json:"results_field,omitempty"`
}

// DataframeAnalyticsDestinationBuilder holds DataframeAnalyticsDestination struct and provides a builder API.
type DataframeAnalyticsDestinationBuilder struct {
	v *DataframeAnalyticsDestination
}

// NewDataframeAnalyticsDestination provides a builder for the DataframeAnalyticsDestination struct.
func NewDataframeAnalyticsDestinationBuilder() *DataframeAnalyticsDestinationBuilder {
	r := DataframeAnalyticsDestinationBuilder{
		&DataframeAnalyticsDestination{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsDestination struct
func (rb *DataframeAnalyticsDestinationBuilder) Build() DataframeAnalyticsDestination {
	return *rb.v
}

// Index Defines the destination index to store the results of the data frame
// analytics job.

func (rb *DataframeAnalyticsDestinationBuilder) Index(index IndexName) *DataframeAnalyticsDestinationBuilder {
	rb.v.Index = index
	return rb
}

// ResultsField Defines the name of the field in which to store the results of the analysis.
// Defaults to `ml`.

func (rb *DataframeAnalyticsDestinationBuilder) ResultsField(resultsfield Field) *DataframeAnalyticsDestinationBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}
