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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DataframeAnalyticsStatsProgress type.
//
// https://github.com/elastic/elasticsearch-specification/blob/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9/specification/ml/_types/DataframeAnalytics.ts#L347-L352
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
				return fmt.Errorf("%s | %w", "Phase", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Phase = o

		case "progress_percent":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ProgressPercent", err)
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

// false
