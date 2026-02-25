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

// Package jobstate
package jobstate

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/ml/_types/Job.ts#L36-L52
type JobState struct {
	Name string
}

var (

	// Closing The job close action is in progress and has not yet completed. A closing job
	// cannot accept further data.
	Closing = JobState{"closing"}

	// Closed The job finished successfully with its model state persisted. The job must be
	// opened before it can accept further data.
	Closed = JobState{"closed"}

	// Opened The job is available to receive and process data.
	Opened = JobState{"opened"}

	// Failed The job did not finish successfully due to an error. This situation can occur
	// due to invalid input data, a fatal error occurring during the analysis, or an
	// external interaction such as the process being killed by the Linux out of
	// memory (OOM) killer. If the job had irrevocably failed, it must be force
	// closed and then deleted. If the datafeed can be corrected, the job can be
	// closed and then re-opened.
	Failed = JobState{"failed"}

	// Opening The job open action is in progress and has not yet completed.
	Opening = JobState{"opening"}
)

func (j JobState) MarshalText() (text []byte, err error) {
	return []byte(j.String()), nil
}

func (j *JobState) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "closing":
		*j = Closing
	case "closed":
		*j = Closed
	case "opened":
		*j = Opened
	case "failed":
		*j = Failed
	case "opening":
		*j = Opening
	default:
		*j = JobState{string(text)}
	}

	return nil
}

func (j JobState) String() string {
	return j.Name
}
