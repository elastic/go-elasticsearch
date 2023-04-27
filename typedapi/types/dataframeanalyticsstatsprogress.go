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
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// DataframeAnalyticsStatsProgress type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ml/_types/DataframeAnalytics.ts#L343-L348
type DataframeAnalyticsStatsProgress struct {
	// Phase Defines the phase of the data frame analytics job.
	Phase string `json:"phase"`
	// ProgressPercent The progress that the data frame analytics job has made expressed in
	// percentage.
	ProgressPercent int `json:"progress_percent"`
}

func (s *DataframeAnalyticsStatsProgress) UnmarshalJSON(data []byte) error {

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

		case "phase":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Phase = o

		case "progress_percent":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ProgressPercent = value
			case float64:
				f := int(v)
				s.ProgressPercent = f
			}

		}
	}
	return nil
}

// NewDataframeAnalyticsStatsProgress returns a DataframeAnalyticsStatsProgress.
func NewDataframeAnalyticsStatsProgress() *DataframeAnalyticsStatsProgress {
	r := &DataframeAnalyticsStatsProgress{}

	return r
}
