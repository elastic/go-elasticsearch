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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Package routingstate
package routingstate

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/TrainedModel.ts#L335-L356
type RoutingState struct {
	Name string
}

var (
	Failed = RoutingState{"failed"}

	Started = RoutingState{"started"}

	Starting = RoutingState{"starting"}

	Stopped = RoutingState{"stopped"}

	Stopping = RoutingState{"stopping"}
)

func (r RoutingState) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RoutingState) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "failed":
		*r = Failed
	case "started":
		*r = Started
	case "starting":
		*r = Starting
	case "stopped":
		*r = Stopped
	case "stopping":
		*r = Stopping
	default:
		*r = RoutingState{string(text)}
	}

	return nil
}

func (r RoutingState) String() string {
	return r.Name
}
