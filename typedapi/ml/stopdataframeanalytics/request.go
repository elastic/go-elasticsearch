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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package stopdataframeanalytics

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Request holds the request body struct for the package stopdataframeanalytics
//
// https://github.com/elastic/elasticsearch-specification/blob/e196f9953fa743572ee46884835f1934bce9a16b/specification/ml/stop_data_frame_analytics/MlStopDataFrameAnalyticsRequest.ts#L24-L113
type Request struct {
	// AllowNoMatch Specifies what to do when the request:
	//
	// 1. Contains wildcard expressions and there are no data frame analytics
	// jobs that match.
	// 2. Contains the _all string or no identifiers and there are no matches.
	// 3. Contains wildcard expressions and there are only partial matches.
	//
	// The default value is true, which returns an empty data_frame_analytics
	// array when there are no matches and the subset of results when there are
	// partial matches. If this parameter is false, the request returns a 404
	// status code when there are no matches or only partial matches.
	AllowNoMatch *bool `json:"allow_no_match,omitempty"`
	// Force If true, the data frame analytics job is stopped forcefully.
	Force *bool `json:"force,omitempty"`
	// Id If provided, must be the same identifier as in the path.
	Id *string `json:"id,omitempty"`
	// Timeout Controls the amount of time to wait until the data frame analytics job
	// stops. Defaults to 20 seconds.
	Timeout types.Duration `json:"timeout,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Stopdataframeanalytics request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "allow_no_match":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowNoMatch", err)
				}
				s.AllowNoMatch = &value
			case bool:
				s.AllowNoMatch = &v
			}

		case "force":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Force", err)
				}
				s.Force = &value
			case bool:
				s.Force = &v
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "timeout":
			if err := dec.Decode(&s.Timeout); err != nil {
				return fmt.Errorf("%s | %w", "Timeout", err)
			}

		}
	}
	return nil
}
