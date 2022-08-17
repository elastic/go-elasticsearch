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

// DataframeAnalysisFeatureProcessorMultiEncoding type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L269-L272
type DataframeAnalysisFeatureProcessorMultiEncoding struct {
	// Processors The ordered array of custom processors to execute. Must be more than 1.
	Processors []int `json:"processors"`
}

// DataframeAnalysisFeatureProcessorMultiEncodingBuilder holds DataframeAnalysisFeatureProcessorMultiEncoding struct and provides a builder API.
type DataframeAnalysisFeatureProcessorMultiEncodingBuilder struct {
	v *DataframeAnalysisFeatureProcessorMultiEncoding
}

// NewDataframeAnalysisFeatureProcessorMultiEncoding provides a builder for the DataframeAnalysisFeatureProcessorMultiEncoding struct.
func NewDataframeAnalysisFeatureProcessorMultiEncodingBuilder() *DataframeAnalysisFeatureProcessorMultiEncodingBuilder {
	r := DataframeAnalysisFeatureProcessorMultiEncodingBuilder{
		&DataframeAnalysisFeatureProcessorMultiEncoding{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalysisFeatureProcessorMultiEncoding struct
func (rb *DataframeAnalysisFeatureProcessorMultiEncodingBuilder) Build() DataframeAnalysisFeatureProcessorMultiEncoding {
	return *rb.v
}

// Processors The ordered array of custom processors to execute. Must be more than 1.

func (rb *DataframeAnalysisFeatureProcessorMultiEncodingBuilder) Processors(processors ...int) *DataframeAnalysisFeatureProcessorMultiEncodingBuilder {
	rb.v.Processors = processors
	return rb
}
