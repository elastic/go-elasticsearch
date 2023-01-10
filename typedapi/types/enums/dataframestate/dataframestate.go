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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Package dataframestate
package dataframestate

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/Dataframe.ts#L20-L26
type DataframeState struct {
	Name string
}

var (
	Started = DataframeState{"started"}

	Stopped = DataframeState{"stopped"}

	Starting = DataframeState{"starting"}

	Stopping = DataframeState{"stopping"}

	Failed = DataframeState{"failed"}
)

func (d DataframeState) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DataframeState) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "started":
		*d = Started
	case "stopped":
		*d = Stopped
	case "starting":
		*d = Starting
	case "stopping":
		*d = Stopping
	case "failed":
		*d = Failed
	default:
		*d = DataframeState{string(text)}
	}

	return nil
}

func (d DataframeState) String() string {
	return d.Name
}
