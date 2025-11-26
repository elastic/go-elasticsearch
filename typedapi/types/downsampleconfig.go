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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/samplingmethod"
)

// DownsampleConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/aa1459fbdcaf57c653729142b3b6e9982373bb1c/specification/indices/_types/Downsample.ts#L22-L31
type DownsampleConfig struct {
	// FixedInterval The interval at which to aggregate the original time series index.
	FixedInterval string `json:"fixed_interval"`
	// SamplingMethod The sampling method used to reduce the documents; it can be either
	// `aggregate` or `last_value`. Defaults to `aggregate`.
	SamplingMethod *samplingmethod.SamplingMethod `json:"sampling_method,omitempty"`
}

func (s *DownsampleConfig) UnmarshalJSON(data []byte) error {

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

		case "fixed_interval":
			if err := dec.Decode(&s.FixedInterval); err != nil {
				return fmt.Errorf("%s | %w", "FixedInterval", err)
			}

		case "sampling_method":
			if err := dec.Decode(&s.SamplingMethod); err != nil {
				return fmt.Errorf("%s | %w", "SamplingMethod", err)
			}

		}
	}
	return nil
}

// NewDownsampleConfig returns a DownsampleConfig.
func NewDownsampleConfig() *DownsampleConfig {
	r := &DownsampleConfig{}

	return r
}

type DownsampleConfigVariant interface {
	DownsampleConfigCaster() *DownsampleConfig
}

func (s *DownsampleConfig) DownsampleConfigCaster() *DownsampleConfig {
	return s
}
