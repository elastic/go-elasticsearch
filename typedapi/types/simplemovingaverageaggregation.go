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
// https://github.com/elastic/elasticsearch-specification/tree/1ad7fe36297b3a8e187b2259dedaf68a47bc236e

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gappolicy"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// SimpleMovingAverageAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1ad7fe36297b3a8e187b2259dedaf68a47bc236e/specification/_types/aggregations/pipeline.ts#L207-L210
type SimpleMovingAverageAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath BucketsPath                `json:"buckets_path,omitempty"`
	Format      *string                    `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy       `json:"gap_policy,omitempty"`
	Meta        map[string]json.RawMessage `json:"meta,omitempty"`
	Minimize    *bool                      `json:"minimize,omitempty"`
	Model       string                     `json:"model,omitempty"`
	Name        *string                    `json:"name,omitempty"`
	Predict     *int                       `json:"predict,omitempty"`
	Settings    EmptyObject                `json:"settings"`
	Window      *int                       `json:"window,omitempty"`
}

func (s *SimpleMovingAverageAggregation) UnmarshalJSON(data []byte) error {
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

		case "buckets_path":
			if err := dec.Decode(&s.BucketsPath); err != nil {
				return err
			}

		case "format":
			if err := dec.Decode(&s.Format); err != nil {
				return err
			}

		case "gap_policy":
			if err := dec.Decode(&s.GapPolicy); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "minimize":
			if err := dec.Decode(&s.Minimize); err != nil {
				return err
			}

		case "model":
			if err := dec.Decode(&s.Model); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "predict":
			if err := dec.Decode(&s.Predict); err != nil {
				return err
			}

		case "settings":
			if err := dec.Decode(&s.Settings); err != nil {
				return err
			}

		case "window":
			if err := dec.Decode(&s.Window); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSimpleMovingAverageAggregation returns a SimpleMovingAverageAggregation.
func NewSimpleMovingAverageAggregation() *SimpleMovingAverageAggregation {
	r := &SimpleMovingAverageAggregation{}

	r.Model = "simple"

	return r
}
