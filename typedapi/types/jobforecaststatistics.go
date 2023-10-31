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

// JobForecastStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/Job.ts#L343-L350
type JobForecastStatistics struct {
	ForecastedJobs   int              `json:"forecasted_jobs"`
	MemoryBytes      *JobStatistics   `json:"memory_bytes,omitempty"`
	ProcessingTimeMs *JobStatistics   `json:"processing_time_ms,omitempty"`
	Records          *JobStatistics   `json:"records,omitempty"`
	Status           map[string]int64 `json:"status,omitempty"`
	Total            int64            `json:"total"`
}

func (s *JobForecastStatistics) UnmarshalJSON(data []byte) error {

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

		case "forecasted_jobs":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ForecastedJobs = value
			case float64:
				f := int(v)
				s.ForecastedJobs = f
			}

		case "memory_bytes":
			if err := dec.Decode(&s.MemoryBytes); err != nil {
				return err
			}

		case "processing_time_ms":
			if err := dec.Decode(&s.ProcessingTimeMs); err != nil {
				return err
			}

		case "records":
			if err := dec.Decode(&s.Records); err != nil {
				return err
			}

		case "status":
			if s.Status == nil {
				s.Status = make(map[string]int64, 0)
			}
			if err := dec.Decode(&s.Status); err != nil {
				return err
			}

		case "total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Total = value
			case float64:
				f := int64(v)
				s.Total = f
			}

		}
	}
	return nil
}

// NewJobForecastStatistics returns a JobForecastStatistics.
func NewJobForecastStatistics() *JobForecastStatistics {
	r := &JobForecastStatistics{
		Status: make(map[string]int64, 0),
	}

	return r
}
