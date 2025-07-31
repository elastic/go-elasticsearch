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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// DataframeAnalyticsStatsOutlierDetection type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/DataframeAnalytics.ts#L405-L418
type DataframeAnalyticsStatsOutlierDetection struct {
	// Parameters The list of job parameters specified by the user or determined by algorithmic
	// heuristics.
	Parameters OutlierDetectionParameters `json:"parameters"`
	// Timestamp The timestamp when the statistics were reported in milliseconds since the
	// epoch.
	Timestamp int64 `json:"timestamp"`
	// TimingStats An object containing time statistics about the data frame analytics job.
	TimingStats TimingStats `json:"timing_stats"`
}

func (s *DataframeAnalyticsStatsOutlierDetection) UnmarshalJSON(data []byte) error {

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

		case "parameters":
			if err := dec.Decode(&s.Parameters); err != nil {
				return fmt.Errorf("%s | %w", "Parameters", err)
			}

		case "timestamp":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return fmt.Errorf("%s | %w", "Timestamp", err)
			}

		case "timing_stats":
			if err := dec.Decode(&s.TimingStats); err != nil {
				return fmt.Errorf("%s | %w", "TimingStats", err)
			}

		}
	}
	return nil
}

// NewDataframeAnalyticsStatsOutlierDetection returns a DataframeAnalyticsStatsOutlierDetection.
func NewDataframeAnalyticsStatsOutlierDetection() *DataframeAnalyticsStatsOutlierDetection {
	r := &DataframeAnalyticsStatsOutlierDetection{}

	return r
}
