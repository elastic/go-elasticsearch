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

// Package indexingjobstate
package indexingjobstate

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/rollup/get_jobs/types.ts#L66-L72
type IndexingJobState struct {
	Name string
}

var (
	Started = IndexingJobState{"started"}

	Indexing = IndexingJobState{"indexing"}

	Stopping = IndexingJobState{"stopping"}

	Stopped = IndexingJobState{"stopped"}

	Aborting = IndexingJobState{"aborting"}
)

func (i IndexingJobState) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *IndexingJobState) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "started":
		*i = Started
	case "indexing":
		*i = Indexing
	case "stopping":
		*i = Stopping
	case "stopped":
		*i = Stopped
	case "aborting":
		*i = Aborting
	default:
		*i = IndexingJobState{string(text)}
	}

	return nil
}

func (i IndexingJobState) String() string {
	return i.Name
}
