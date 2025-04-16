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
// https://github.com/elastic/elasticsearch-specification/tree/f6a370d0fba975752c644fc730f7c45610e28f36

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// IndexSettingsLifecycleStep type.
//
// https://github.com/elastic/elasticsearch-specification/blob/f6a370d0fba975752c644fc730f7c45610e28f36/specification/indices/_types/IndexSettings.ts#L325-L331
type IndexSettingsLifecycleStep struct {
	// WaitTimeThreshold Time to wait for the cluster to resolve allocation issues during an ILM
	// shrink action. Must be greater than 1h (1 hour).
	// See Shard allocation for shrink.
	WaitTimeThreshold Duration `json:"wait_time_threshold,omitempty"`
}

func (s *IndexSettingsLifecycleStep) UnmarshalJSON(data []byte) error {

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

		case "wait_time_threshold":
			if err := dec.Decode(&s.WaitTimeThreshold); err != nil {
				return fmt.Errorf("%s | %w", "WaitTimeThreshold", err)
			}

		}
	}
	return nil
}

// NewIndexSettingsLifecycleStep returns a IndexSettingsLifecycleStep.
func NewIndexSettingsLifecycleStep() *IndexSettingsLifecycleStep {
	r := &IndexSettingsLifecycleStep{}

	return r
}

// true

type IndexSettingsLifecycleStepVariant interface {
	IndexSettingsLifecycleStepCaster() *IndexSettingsLifecycleStep
}

func (s *IndexSettingsLifecycleStep) IndexSettingsLifecycleStepCaster() *IndexSettingsLifecycleStep {
	return s
}
