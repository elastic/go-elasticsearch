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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Package deploymentstate
package deploymentstate

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/TrainedModel.ts#L273-L286
type DeploymentState struct {
	Name string
}

var (
	Started = DeploymentState{"started"}

	Starting = DeploymentState{"starting"}

	Stopping = DeploymentState{"stopping"}
)

func (d DeploymentState) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DeploymentState) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "started":
		*d = Started
	case "starting":
		*d = Starting
	case "stopping":
		*d = Stopping
	default:
		*d = DeploymentState{string(text)}
	}

	return nil
}

func (d DeploymentState) String() string {
	return d.Name
}
