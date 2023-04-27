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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// DelayedDataCheckConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ml/_types/Datafeed.ts#L119-L130
type DelayedDataCheckConfig struct {
	// CheckWindow The window of time that is searched for late data. This window of time ends
	// with the latest finalized bucket.
	// It defaults to null, which causes an appropriate `check_window` to be
	// calculated when the real-time datafeed runs.
	// In particular, the default `check_window` span calculation is based on the
	// maximum of `2h` or `8 * bucket_span`.
	CheckWindow Duration `json:"check_window,omitempty"`
	// Enabled Specifies whether the datafeed periodically checks for delayed data.
	Enabled bool `json:"enabled"`
}

func (s *DelayedDataCheckConfig) UnmarshalJSON(data []byte) error {

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

		case "check_window":
			if err := dec.Decode(&s.CheckWindow); err != nil {
				return err
			}

		case "enabled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		}
	}
	return nil
}

// NewDelayedDataCheckConfig returns a DelayedDataCheckConfig.
func NewDelayedDataCheckConfig() *DelayedDataCheckConfig {
	r := &DelayedDataCheckConfig{}

	return r
}
