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

// DataframeAnalyticsMemoryEstimation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/DataframeAnalytics.ts#L70-L75
type DataframeAnalyticsMemoryEstimation struct {
	// ExpectedMemoryWithDisk Estimated memory usage under the assumption that overflowing to disk is
	// allowed during data frame analytics. expected_memory_with_disk is usually
	// smaller than expected_memory_without_disk as using disk allows to limit the
	// main memory needed to perform data frame analytics.
	ExpectedMemoryWithDisk string `json:"expected_memory_with_disk"`
	// ExpectedMemoryWithoutDisk Estimated memory usage under the assumption that the whole data frame
	// analytics should happen in memory (i.e. without overflowing to disk).
	ExpectedMemoryWithoutDisk string `json:"expected_memory_without_disk"`
}

func (s *DataframeAnalyticsMemoryEstimation) UnmarshalJSON(data []byte) error {

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

		case "expected_memory_with_disk":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ExpectedMemoryWithDisk = o

		case "expected_memory_without_disk":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ExpectedMemoryWithoutDisk = o

		}
	}
	return nil
}

// NewDataframeAnalyticsMemoryEstimation returns a DataframeAnalyticsMemoryEstimation.
func NewDataframeAnalyticsMemoryEstimation() *DataframeAnalyticsMemoryEstimation {
	r := &DataframeAnalyticsMemoryEstimation{}

	return r
}
