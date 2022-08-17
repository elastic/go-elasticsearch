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

// DataframeAnalysisFeatureProcessorOneHotEncoding type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeAnalytics.ts#L288-L293
type DataframeAnalysisFeatureProcessorOneHotEncoding struct {
	// Field The name of the field to encode.
	Field Field `json:"field"`
	// HotMap The one hot map mapping the field value with the column name.
	HotMap string `json:"hot_map"`
}

// DataframeAnalysisFeatureProcessorOneHotEncodingBuilder holds DataframeAnalysisFeatureProcessorOneHotEncoding struct and provides a builder API.
type DataframeAnalysisFeatureProcessorOneHotEncodingBuilder struct {
	v *DataframeAnalysisFeatureProcessorOneHotEncoding
}

// NewDataframeAnalysisFeatureProcessorOneHotEncoding provides a builder for the DataframeAnalysisFeatureProcessorOneHotEncoding struct.
func NewDataframeAnalysisFeatureProcessorOneHotEncodingBuilder() *DataframeAnalysisFeatureProcessorOneHotEncodingBuilder {
	r := DataframeAnalysisFeatureProcessorOneHotEncodingBuilder{
		&DataframeAnalysisFeatureProcessorOneHotEncoding{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalysisFeatureProcessorOneHotEncoding struct
func (rb *DataframeAnalysisFeatureProcessorOneHotEncodingBuilder) Build() DataframeAnalysisFeatureProcessorOneHotEncoding {
	return *rb.v
}

// Field The name of the field to encode.

func (rb *DataframeAnalysisFeatureProcessorOneHotEncodingBuilder) Field(field Field) *DataframeAnalysisFeatureProcessorOneHotEncodingBuilder {
	rb.v.Field = field
	return rb
}

// HotMap The one hot map mapping the field value with the column name.

func (rb *DataframeAnalysisFeatureProcessorOneHotEncodingBuilder) HotMap(hotmap string) *DataframeAnalysisFeatureProcessorOneHotEncodingBuilder {
	rb.v.HotMap = hotmap
	return rb
}
