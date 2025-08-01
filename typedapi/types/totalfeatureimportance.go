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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// TotalFeatureImportance type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/TrainedModel.ts#L290-L297
type TotalFeatureImportance struct {
	// Classes If the trained model is a classification model, feature importance statistics
	// are gathered per target class value.
	Classes []TotalFeatureImportanceClass `json:"classes"`
	// FeatureName The feature for which this importance was calculated.
	FeatureName string `json:"feature_name"`
	// Importance A collection of feature importance statistics related to the training data
	// set for this particular feature.
	Importance []TotalFeatureImportanceStatistics `json:"importance"`
}

func (s *TotalFeatureImportance) UnmarshalJSON(data []byte) error {

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

		case "classes":
			if err := dec.Decode(&s.Classes); err != nil {
				return fmt.Errorf("%s | %w", "Classes", err)
			}

		case "feature_name":
			if err := dec.Decode(&s.FeatureName); err != nil {
				return fmt.Errorf("%s | %w", "FeatureName", err)
			}

		case "importance":
			if err := dec.Decode(&s.Importance); err != nil {
				return fmt.Errorf("%s | %w", "Importance", err)
			}

		}
	}
	return nil
}

// NewTotalFeatureImportance returns a TotalFeatureImportance.
func NewTotalFeatureImportance() *TotalFeatureImportance {
	r := &TotalFeatureImportance{}

	return r
}
