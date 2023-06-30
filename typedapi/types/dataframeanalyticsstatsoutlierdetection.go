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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// DataframeAnalyticsStatsOutlierDetection type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/DataframeAnalytics.ts#L389-L393
type DataframeAnalyticsStatsOutlierDetection struct {
	Parameters  OutlierDetectionParameters `json:"parameters"`
	Timestamp   int64                      `json:"timestamp"`
	TimingStats TimingStats                `json:"timing_stats"`
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
				return err
			}

		case "timestamp":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return err
			}

		case "timing_stats":
			if err := dec.Decode(&s.TimingStats); err != nil {
				return err
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
