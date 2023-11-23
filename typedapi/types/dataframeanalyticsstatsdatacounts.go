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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// DataframeAnalyticsStatsDataCounts type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/DataframeAnalytics.ts#L364-L371
type DataframeAnalyticsStatsDataCounts struct {
	// SkippedDocsCount The number of documents that are skipped during the analysis because they
	// contained values that are not supported by the analysis. For example, outlier
	// detection does not support missing fields so it skips documents with missing
	// fields. Likewise, all types of analysis skip documents that contain arrays
	// with more than one element.
	SkippedDocsCount int `json:"skipped_docs_count"`
	// TestDocsCount The number of documents that are not used for training the model and can be
	// used for testing.
	TestDocsCount int `json:"test_docs_count"`
	// TrainingDocsCount The number of documents that are used for training the model.
	TrainingDocsCount int `json:"training_docs_count"`
}

func (s *DataframeAnalyticsStatsDataCounts) UnmarshalJSON(data []byte) error {

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

		case "skipped_docs_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.SkippedDocsCount = value
			case float64:
				f := int(v)
				s.SkippedDocsCount = f
			}

		case "test_docs_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TestDocsCount = value
			case float64:
				f := int(v)
				s.TestDocsCount = f
			}

		case "training_docs_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TrainingDocsCount = value
			case float64:
				f := int(v)
				s.TrainingDocsCount = f
			}

		}
	}
	return nil
}

// NewDataframeAnalyticsStatsDataCounts returns a DataframeAnalyticsStatsDataCounts.
func NewDataframeAnalyticsStatsDataCounts() *DataframeAnalyticsStatsDataCounts {
	r := &DataframeAnalyticsStatsDataCounts{}

	return r
}
