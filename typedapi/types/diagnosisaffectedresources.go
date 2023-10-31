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
)

// DiagnosisAffectedResources type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/health_report/types.ts#L57-L63
type DiagnosisAffectedResources struct {
	FeatureStates        []string        `json:"feature_states,omitempty"`
	Indices              []string        `json:"indices,omitempty"`
	Nodes                []IndicatorNode `json:"nodes,omitempty"`
	SlmPolicies          []string        `json:"slm_policies,omitempty"`
	SnapshotRepositories []string        `json:"snapshot_repositories,omitempty"`
}

func (s *DiagnosisAffectedResources) UnmarshalJSON(data []byte) error {

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

		case "feature_states":
			if err := dec.Decode(&s.FeatureStates); err != nil {
				return err
			}

		case "indices":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Indices = append(s.Indices, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Indices); err != nil {
					return err
				}
			}

		case "nodes":
			if err := dec.Decode(&s.Nodes); err != nil {
				return err
			}

		case "slm_policies":
			if err := dec.Decode(&s.SlmPolicies); err != nil {
				return err
			}

		case "snapshot_repositories":
			if err := dec.Decode(&s.SnapshotRepositories); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDiagnosisAffectedResources returns a DiagnosisAffectedResources.
func NewDiagnosisAffectedResources() *DiagnosisAffectedResources {
	r := &DiagnosisAffectedResources{}

	return r
}
