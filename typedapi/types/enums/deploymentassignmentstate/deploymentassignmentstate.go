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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package deploymentassignmentstate
package deploymentassignmentstate

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/ml/_types/TrainedModel.ts#L347-L364
type DeploymentAssignmentState struct {
	Name string
}

var (

	// Started The deployment is usable; at least one node has the model allocated.
	Started = DeploymentAssignmentState{"started"}

	// Starting The deployment has recently started but is not yet usable; the model is not
	// allocated on any nodes.
	Starting = DeploymentAssignmentState{"starting"}

	// Stopping The deployment is preparing to stop and deallocate the model from the
	// relevant nodes.
	Stopping = DeploymentAssignmentState{"stopping"}

	// Failed The deployment is on a failed state and must be re-deployed.
	Failed = DeploymentAssignmentState{"failed"}
)

func (d DeploymentAssignmentState) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DeploymentAssignmentState) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "started":
		*d = Started
	case "starting":
		*d = Starting
	case "stopping":
		*d = Stopping
	case "failed":
		*d = Failed
	default:
		*d = DeploymentAssignmentState{string(text)}
	}

	return nil
}

func (d DeploymentAssignmentState) String() string {
	return d.Name
}
