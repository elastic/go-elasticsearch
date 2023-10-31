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

// MlInferenceDeployments type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/xpack/usage/types.ts#L227-L232
type MlInferenceDeployments struct {
	Count           int                          `json:"count"`
	InferenceCounts JobStatistics                `json:"inference_counts"`
	ModelSizesBytes JobStatistics                `json:"model_sizes_bytes"`
	TimeMs          MlInferenceDeploymentsTimeMs `json:"time_ms"`
}

func (s *MlInferenceDeployments) UnmarshalJSON(data []byte) error {

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

		case "count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "inference_counts":
			if err := dec.Decode(&s.InferenceCounts); err != nil {
				return err
			}

		case "model_sizes_bytes":
			if err := dec.Decode(&s.ModelSizesBytes); err != nil {
				return err
			}

		case "time_ms":
			if err := dec.Decode(&s.TimeMs); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewMlInferenceDeployments returns a MlInferenceDeployments.
func NewMlInferenceDeployments() *MlInferenceDeployments {
	r := &MlInferenceDeployments{}

	return r
}
