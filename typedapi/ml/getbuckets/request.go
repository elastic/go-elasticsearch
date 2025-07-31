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

package getbuckets

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package getbuckets
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/get_buckets/MlGetBucketsRequest.ts#L26-L145
type Request struct {

	// AnomalyScore Refer to the description for the `anomaly_score` query parameter.
	AnomalyScore *types.Float64 `json:"anomaly_score,omitempty"`
	// Desc Refer to the description for the `desc` query parameter.
	Desc *bool `json:"desc,omitempty"`
	// End Refer to the description for the `end` query parameter.
	End types.DateTime `json:"end,omitempty"`
	// ExcludeInterim Refer to the description for the `exclude_interim` query parameter.
	ExcludeInterim *bool `json:"exclude_interim,omitempty"`
	// Expand Refer to the description for the `expand` query parameter.
	Expand *bool       `json:"expand,omitempty"`
	Page   *types.Page `json:"page,omitempty"`
	// Sort Refer to the desription for the `sort` query parameter.
	Sort *string `json:"sort,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Getbuckets request: %w", err)
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

		case "anomaly_score":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "AnomalyScore", err)
				}
				f := types.Float64(value)
				s.AnomalyScore = &f
			case float64:
				f := types.Float64(v)
				s.AnomalyScore = &f
			}

		case "desc":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Desc", err)
				}
				s.Desc = &value
			case bool:
				s.Desc = &v
			}

		case "end":
			if err := dec.Decode(&s.End); err != nil {
				return fmt.Errorf("%s | %w", "End", err)
			}

		case "exclude_interim":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ExcludeInterim", err)
				}
				s.ExcludeInterim = &value
			case bool:
				s.ExcludeInterim = &v
			}

		case "expand":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Expand", err)
				}
				s.Expand = &value
			case bool:
				s.Expand = &v
			}

		case "page":
			if err := dec.Decode(&s.Page); err != nil {
				return fmt.Errorf("%s | %w", "Page", err)
			}

		case "sort":
			if err := dec.Decode(&s.Sort); err != nil {
				return fmt.Errorf("%s | %w", "Sort", err)
			}

		case "start":
			if err := dec.Decode(&s.Start); err != nil {
				return fmt.Errorf("%s | %w", "Start", err)
			}

		}
	}
	return nil
}
