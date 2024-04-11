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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Ensemble type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/ml/put_trained_model/types.ts#L93-L99
type Ensemble struct {
	AggregateOutput      *AggregateOutput `json:"aggregate_output,omitempty"`
	ClassificationLabels []string         `json:"classification_labels,omitempty"`
	FeatureNames         []string         `json:"feature_names,omitempty"`
	TargetType           *string          `json:"target_type,omitempty"`
	TrainedModels        []TrainedModel   `json:"trained_models"`
}

func (s *Ensemble) UnmarshalJSON(data []byte) error {

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

		case "aggregate_output":
			if err := dec.Decode(&s.AggregateOutput); err != nil {
				return fmt.Errorf("%s | %w", "AggregateOutput", err)
			}

		case "classification_labels":
			if err := dec.Decode(&s.ClassificationLabels); err != nil {
				return fmt.Errorf("%s | %w", "ClassificationLabels", err)
			}

		case "feature_names":
			if err := dec.Decode(&s.FeatureNames); err != nil {
				return fmt.Errorf("%s | %w", "FeatureNames", err)
			}

		case "target_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TargetType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TargetType = &o

		case "trained_models":
			if err := dec.Decode(&s.TrainedModels); err != nil {
				return fmt.Errorf("%s | %w", "TrainedModels", err)
			}

		}
	}
	return nil
}

// NewEnsemble returns a Ensemble.
func NewEnsemble() *Ensemble {
	r := &Ensemble{}

	return r
}
