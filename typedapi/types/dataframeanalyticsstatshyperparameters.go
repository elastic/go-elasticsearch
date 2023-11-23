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

// DataframeAnalyticsStatsHyperparameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/DataframeAnalytics.ts#L383-L402
type DataframeAnalyticsStatsHyperparameters struct {
	// Hyperparameters An object containing the parameters of the classification analysis job.
	Hyperparameters Hyperparameters `json:"hyperparameters"`
	// Iteration The number of iterations on the analysis.
	Iteration int `json:"iteration"`
	// Timestamp The timestamp when the statistics were reported in milliseconds since the
	// epoch.
	Timestamp int64 `json:"timestamp"`
	// TimingStats An object containing time statistics about the data frame analytics job.
	TimingStats TimingStats `json:"timing_stats"`
	// ValidationLoss An object containing information about validation loss.
	ValidationLoss ValidationLoss `json:"validation_loss"`
}

func (s *DataframeAnalyticsStatsHyperparameters) UnmarshalJSON(data []byte) error {

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

		case "hyperparameters":
			if err := dec.Decode(&s.Hyperparameters); err != nil {
				return err
			}

		case "iteration":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Iteration = value
			case float64:
				f := int(v)
				s.Iteration = f
			}

		case "timestamp":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return err
			}

		case "timing_stats":
			if err := dec.Decode(&s.TimingStats); err != nil {
				return err
			}

		case "validation_loss":
			if err := dec.Decode(&s.ValidationLoss); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataframeAnalyticsStatsHyperparameters returns a DataframeAnalyticsStatsHyperparameters.
func NewDataframeAnalyticsStatsHyperparameters() *DataframeAnalyticsStatsHyperparameters {
	r := &DataframeAnalyticsStatsHyperparameters{}

	return r
}
