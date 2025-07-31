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
)

// SlowlogTresholdLevels type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/IndexSettings.ts#L540-L545
type SlowlogTresholdLevels struct {
	Debug Duration `json:"debug,omitempty"`
	Info  Duration `json:"info,omitempty"`
	Trace Duration `json:"trace,omitempty"`
	Warn  Duration `json:"warn,omitempty"`
}

func (s *SlowlogTresholdLevels) UnmarshalJSON(data []byte) error {

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

		case "debug":
			if err := dec.Decode(&s.Debug); err != nil {
				return fmt.Errorf("%s | %w", "Debug", err)
			}

		case "info":
			if err := dec.Decode(&s.Info); err != nil {
				return fmt.Errorf("%s | %w", "Info", err)
			}

		case "trace":
			if err := dec.Decode(&s.Trace); err != nil {
				return fmt.Errorf("%s | %w", "Trace", err)
			}

		case "warn":
			if err := dec.Decode(&s.Warn); err != nil {
				return fmt.Errorf("%s | %w", "Warn", err)
			}

		}
	}
	return nil
}

// NewSlowlogTresholdLevels returns a SlowlogTresholdLevels.
func NewSlowlogTresholdLevels() *SlowlogTresholdLevels {
	r := &SlowlogTresholdLevels{}

	return r
}

type SlowlogTresholdLevelsVariant interface {
	SlowlogTresholdLevelsCaster() *SlowlogTresholdLevels
}

func (s *SlowlogTresholdLevels) SlowlogTresholdLevelsCaster() *SlowlogTresholdLevels {
	return s
}
