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

package flushjob

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package flushjob
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/flush_job/MlFlushJobRequest.ts#L24-L107
type Request struct {

	// AdvanceTime Refer to the description for the `advance_time` query parameter.
	AdvanceTime types.DateTime `json:"advance_time,omitempty"`
	// CalcInterim Refer to the description for the `calc_interim` query parameter.
	CalcInterim *bool `json:"calc_interim,omitempty"`
	// End Refer to the description for the `end` query parameter.
	End types.DateTime `json:"end,omitempty"`
	// SkipTime Refer to the description for the `skip_time` query parameter.
	SkipTime types.DateTime `json:"skip_time,omitempty"`
	// Start Refer to the description for the `start` query parameter.
	Start types.DateTime `json:"start,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Flushjob request: %w", err)
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

		case "advance_time":
			if err := dec.Decode(&s.AdvanceTime); err != nil {
				return fmt.Errorf("%s | %w", "AdvanceTime", err)
			}

		case "calc_interim":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CalcInterim", err)
				}
				s.CalcInterim = &value
			case bool:
				s.CalcInterim = &v
			}

		case "end":
			if err := dec.Decode(&s.End); err != nil {
				return fmt.Errorf("%s | %w", "End", err)
			}

		case "skip_time":
			if err := dec.Decode(&s.SkipTime); err != nil {
				return fmt.Errorf("%s | %w", "SkipTime", err)
			}

		case "start":
			if err := dec.Decode(&s.Start); err != nil {
				return fmt.Errorf("%s | %w", "Start", err)
			}

		}
	}
	return nil
}
