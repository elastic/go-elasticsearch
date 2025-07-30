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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeAnalysisFeatureProcessorOneHotEncoding struct {
	v *types.DataframeAnalysisFeatureProcessorOneHotEncoding
}

// The configuration information necessary to perform one hot encoding.
func NewDataframeAnalysisFeatureProcessorOneHotEncoding(hotmap string) *_dataframeAnalysisFeatureProcessorOneHotEncoding {

	tmp := &_dataframeAnalysisFeatureProcessorOneHotEncoding{v: types.NewDataframeAnalysisFeatureProcessorOneHotEncoding()}

	tmp.HotMap(hotmap)

	return tmp

}

func (s *_dataframeAnalysisFeatureProcessorOneHotEncoding) Field(field string) *_dataframeAnalysisFeatureProcessorOneHotEncoding {

	s.v.Field = field

	return s
}

func (s *_dataframeAnalysisFeatureProcessorOneHotEncoding) HotMap(hotmap string) *_dataframeAnalysisFeatureProcessorOneHotEncoding {

	s.v.HotMap = hotmap

	return s
}

func (s *_dataframeAnalysisFeatureProcessorOneHotEncoding) DataframeAnalysisFeatureProcessorCaster() *types.DataframeAnalysisFeatureProcessor {
	container := types.NewDataframeAnalysisFeatureProcessor()

	container.OneHotEncoding = s.v

	return container
}

func (s *_dataframeAnalysisFeatureProcessorOneHotEncoding) DataframeAnalysisFeatureProcessorOneHotEncodingCaster() *types.DataframeAnalysisFeatureProcessorOneHotEncoding {
	return s.v
}
