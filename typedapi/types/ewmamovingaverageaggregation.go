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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

// EwmaMovingAverageAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/pipeline.ts#L278-L281
type EwmaMovingAverageAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`
	// Format `DecimalFormat` pattern for the output value.
	// If specified, the formatted value is returned in the aggregation’s
	// `value_as_string` property.
	Format *string `json:"format,omitempty"`
	// GapPolicy Policy to apply when gaps are found in the data.
	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Minimize  *bool                `json:"minimize,omitempty"`
	Model     string               `json:"model,omitempty"`
	Predict   *int                 `json:"predict,omitempty"`
	Settings  EwmaModelSettings    `json:"settings"`
	Window    *int                 `json:"window,omitempty"`
}

func (s *EwmaMovingAverageAggregation) UnmarshalJSON(data []byte) error {

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
				return fmt.Errorf("%s | %w", "BucketsPath", err)
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = &o

		case "gap_policy":
			if err := dec.Decode(&s.GapPolicy); err != nil {
				return fmt.Errorf("%s | %w", "GapPolicy", err)
			}

		case "minimize":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Minimize", err)
				}
				s.Minimize = &value
			case bool:
				s.Minimize = &v
			}

		case "model":
			if err := dec.Decode(&s.Model); err != nil {
				return fmt.Errorf("%s | %w", "Model", err)
			}

		case "predict":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Predict", err)
				}
				s.Predict = &value
			case float64:
				f := int(v)
				s.Predict = &f
			}

		case "settings":
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		case "window":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Window", err)
				}
				s.Window = &value
			case float64:
				f := int(v)
				s.Window = &f
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s EwmaMovingAverageAggregation) MarshalJSON() ([]byte, error) {
	type innerEwmaMovingAverageAggregation EwmaMovingAverageAggregation
	tmp := innerEwmaMovingAverageAggregation{
		BucketsPath: s.BucketsPath,
		Format:      s.Format,
		GapPolicy:   s.GapPolicy,
		Minimize:    s.Minimize,
		Model:       s.Model,
		Predict:     s.Predict,
		Settings:    s.Settings,
		Window:      s.Window,
	}

	tmp.Model = "ewma"

	return json.Marshal(tmp)
}

// NewEwmaMovingAverageAggregation returns a EwmaMovingAverageAggregation.
func NewEwmaMovingAverageAggregation() *EwmaMovingAverageAggregation {
	r := &EwmaMovingAverageAggregation{}

	return r
}

type EwmaMovingAverageAggregationVariant interface {
	EwmaMovingAverageAggregationCaster() *EwmaMovingAverageAggregation
}

func (s *EwmaMovingAverageAggregation) EwmaMovingAverageAggregationCaster() *EwmaMovingAverageAggregation {
	return s
}

func (s *EwmaMovingAverageAggregation) MovingAverageAggregationCaster() *MovingAverageAggregation {
	o := MovingAverageAggregation(s)
	return &o
}
