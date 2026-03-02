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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeAnalysisFeatureProcessorMultiEncoding struct {
	v *types.DataframeAnalysisFeatureProcessorMultiEncoding
}

// The configuration information necessary to perform multi encoding. It allows
// multiple processors to be changed together. This way the output of a
// processor can then be passed to another as an input.
func NewDataframeAnalysisFeatureProcessorMultiEncoding() *_dataframeAnalysisFeatureProcessorMultiEncoding {

	return &_dataframeAnalysisFeatureProcessorMultiEncoding{v: types.NewDataframeAnalysisFeatureProcessorMultiEncoding()}

}

func (s *_dataframeAnalysisFeatureProcessorMultiEncoding) Processors(processors ...int) *_dataframeAnalysisFeatureProcessorMultiEncoding {

	for _, v := range processors {

		s.v.Processors = append(s.v.Processors, v)

	}
	return s
}

func (s *_dataframeAnalysisFeatureProcessorMultiEncoding) DataframeAnalysisFeatureProcessorCaster() *types.DataframeAnalysisFeatureProcessor {
	container := types.NewDataframeAnalysisFeatureProcessor()

	container.MultiEncoding = s.v

	return container
}

func (s *_dataframeAnalysisFeatureProcessorMultiEncoding) DataframeAnalysisFeatureProcessorMultiEncodingCaster() *types.DataframeAnalysisFeatureProcessorMultiEncoding {
	return s.v
}
