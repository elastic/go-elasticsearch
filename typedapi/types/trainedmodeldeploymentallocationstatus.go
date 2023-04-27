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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/deploymentallocationstate"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// TrainedModelDeploymentAllocationStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ml/_types/TrainedModel.ts#L378-L385
type TrainedModelDeploymentAllocationStatus struct {
	// AllocationCount The current number of nodes where the model is allocated.
	AllocationCount int `json:"allocation_count"`
	// State The detailed allocation state related to the nodes.
	State deploymentallocationstate.DeploymentAllocationState `json:"state"`
	// TargetAllocationCount The desired number of nodes for model allocation.
	TargetAllocationCount int `json:"target_allocation_count"`
}

func (s *TrainedModelDeploymentAllocationStatus) UnmarshalJSON(data []byte) error {

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

		case "allocation_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.AllocationCount = value
			case float64:
				f := int(v)
				s.AllocationCount = f
			}

		case "state":
			if err := dec.Decode(&s.State); err != nil {
				return err
			}

		case "target_allocation_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TargetAllocationCount = value
			case float64:
				f := int(v)
				s.TargetAllocationCount = f
			}

		}
	}
	return nil
}

// NewTrainedModelDeploymentAllocationStatus returns a TrainedModelDeploymentAllocationStatus.
func NewTrainedModelDeploymentAllocationStatus() *TrainedModelDeploymentAllocationStatus {
	r := &TrainedModelDeploymentAllocationStatus{}

	return r
}
