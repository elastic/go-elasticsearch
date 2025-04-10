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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _dataframeAnalysisFeatureProcessorTargetMeanEncoding struct {
	v *types.DataframeAnalysisFeatureProcessorTargetMeanEncoding
}

// The configuration information necessary to perform target mean encoding.
func NewDataframeAnalysisFeatureProcessorTargetMeanEncoding(defaultvalue int) *_dataframeAnalysisFeatureProcessorTargetMeanEncoding {

	tmp := &_dataframeAnalysisFeatureProcessorTargetMeanEncoding{v: types.NewDataframeAnalysisFeatureProcessorTargetMeanEncoding()}

	tmp.DefaultValue(defaultvalue)

	return tmp

}

func (s *_dataframeAnalysisFeatureProcessorTargetMeanEncoding) DefaultValue(defaultvalue int) *_dataframeAnalysisFeatureProcessorTargetMeanEncoding {

	s.v.DefaultValue = defaultvalue

	return s
}

func (s *_dataframeAnalysisFeatureProcessorTargetMeanEncoding) FeatureName(name string) *_dataframeAnalysisFeatureProcessorTargetMeanEncoding {

	s.v.FeatureName = name

	return s
}

func (s *_dataframeAnalysisFeatureProcessorTargetMeanEncoding) Field(field string) *_dataframeAnalysisFeatureProcessorTargetMeanEncoding {

	s.v.Field = field

	return s
}

func (s *_dataframeAnalysisFeatureProcessorTargetMeanEncoding) TargetMap(targetmap map[string]json.RawMessage) *_dataframeAnalysisFeatureProcessorTargetMeanEncoding {

	s.v.TargetMap = targetmap
	return s
}

func (s *_dataframeAnalysisFeatureProcessorTargetMeanEncoding) AddTargetMap(key string, value json.RawMessage) *_dataframeAnalysisFeatureProcessorTargetMeanEncoding {

	var tmp map[string]json.RawMessage
	if s.v.TargetMap == nil {
		s.v.TargetMap = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.TargetMap
	}

	tmp[key] = value

	s.v.TargetMap = tmp
	return s
}

func (s *_dataframeAnalysisFeatureProcessorTargetMeanEncoding) DataframeAnalysisFeatureProcessorCaster() *types.DataframeAnalysisFeatureProcessor {
	container := types.NewDataframeAnalysisFeatureProcessor()

	container.TargetMeanEncoding = s.v

	return container
}

func (s *_dataframeAnalysisFeatureProcessorTargetMeanEncoding) DataframeAnalysisFeatureProcessorTargetMeanEncodingCaster() *types.DataframeAnalysisFeatureProcessorTargetMeanEncoding {
	return s.v
}
