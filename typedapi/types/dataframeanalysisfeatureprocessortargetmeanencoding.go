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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// DataframeAnalysisFeatureProcessorTargetMeanEncoding type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/DataframeAnalytics.ts#L295-L304
type DataframeAnalysisFeatureProcessorTargetMeanEncoding struct {
	// DefaultValue The default value if field value is not found in the target_map.
	DefaultValue int `json:"default_value"`
	// FeatureName The resulting feature name.
	FeatureName string `json:"feature_name"`
	// Field The name of the field to encode.
	Field string `json:"field"`
	// TargetMap The field value to target mean transition map.
	TargetMap map[string]json.RawMessage `json:"target_map"`
}

func (s *DataframeAnalysisFeatureProcessorTargetMeanEncoding) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "default_value":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.DefaultValue = value
			case float64:
				f := int(v)
				s.DefaultValue = f
			}

		case "feature_name":
			if err := dec.Decode(&s.FeatureName); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "target_map":
			if s.TargetMap == nil {
				s.TargetMap = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.TargetMap); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataframeAnalysisFeatureProcessorTargetMeanEncoding returns a DataframeAnalysisFeatureProcessorTargetMeanEncoding.
func NewDataframeAnalysisFeatureProcessorTargetMeanEncoding() *DataframeAnalysisFeatureProcessorTargetMeanEncoding {
	r := &DataframeAnalysisFeatureProcessorTargetMeanEncoding{
		TargetMap: make(map[string]json.RawMessage, 0),
	}

	return r
}
