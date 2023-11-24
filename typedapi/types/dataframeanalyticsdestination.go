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
)

// DataframeAnalyticsDestination type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/DataframeAnalytics.ts#L77-L82
type DataframeAnalyticsDestination struct {
	// Index Defines the destination index to store the results of the data frame
	// analytics job.
	Index string `json:"index"`
	// ResultsField Defines the name of the field in which to store the results of the analysis.
	// Defaults to `ml`.
	ResultsField *string `json:"results_field,omitempty"`
}

func (s *DataframeAnalyticsDestination) UnmarshalJSON(data []byte) error {

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

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "results_field":
			if err := dec.Decode(&s.ResultsField); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataframeAnalyticsDestination returns a DataframeAnalyticsDestination.
func NewDataframeAnalyticsDestination() *DataframeAnalyticsDestination {
	r := &DataframeAnalyticsDestination{}

	return r
}
