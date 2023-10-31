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

// TrainedModelTree type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/put_trained_model/types.ts#L74-L79
type TrainedModelTree struct {
	ClassificationLabels []string               `json:"classification_labels,omitempty"`
	FeatureNames         []string               `json:"feature_names"`
	TargetType           *string                `json:"target_type,omitempty"`
	TreeStructure        []TrainedModelTreeNode `json:"tree_structure"`
}

func (s *TrainedModelTree) UnmarshalJSON(data []byte) error {

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

		case "classification_labels":
			if err := dec.Decode(&s.ClassificationLabels); err != nil {
				return err
			}

		case "feature_names":
			if err := dec.Decode(&s.FeatureNames); err != nil {
				return err
			}

		case "target_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TargetType = &o

		case "tree_structure":
			if err := dec.Decode(&s.TreeStructure); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewTrainedModelTree returns a TrainedModelTree.
func NewTrainedModelTree() *TrainedModelTree {
	r := &TrainedModelTree{}

	return r
}
