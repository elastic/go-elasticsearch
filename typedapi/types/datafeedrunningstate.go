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

// DatafeedRunningState type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/Datafeed.ts#L198-L212
type DatafeedRunningState struct {
	// RealTimeConfigured Indicates if the datafeed is "real-time"; meaning that the datafeed has no
	// configured `end` time.
	RealTimeConfigured bool `json:"real_time_configured"`
	// RealTimeRunning Indicates whether the datafeed has finished running on the available past
	// data.
	// For datafeeds without a configured `end` time, this means that the datafeed
	// is now running on "real-time" data.
	RealTimeRunning bool `json:"real_time_running"`
	// SearchInterval Provides the latest time interval the datafeed has searched.
	SearchInterval *RunningStateSearchInterval `json:"search_interval,omitempty"`
}

func (s *DatafeedRunningState) UnmarshalJSON(data []byte) error {

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

		case "real_time_configured":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.RealTimeConfigured = value
			case bool:
				s.RealTimeConfigured = v
			}

		case "real_time_running":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.RealTimeRunning = value
			case bool:
				s.RealTimeRunning = v
			}

		case "search_interval":
			if err := dec.Decode(&s.SearchInterval); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDatafeedRunningState returns a DatafeedRunningState.
func NewDatafeedRunningState() *DatafeedRunningState {
	r := &DatafeedRunningState{}

	return r
}
