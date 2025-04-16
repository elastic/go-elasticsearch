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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// LifecycleExplainPhaseExecution type.
//
// https://github.com/elastic/elasticsearch-specification/blob/52c473efb1fb5320a5bac12572d0b285882862fb/specification/ilm/explain_lifecycle/types.ts#L69-L74
type LifecycleExplainPhaseExecution struct {
	ModifiedDateInMillis int64  `json:"modified_date_in_millis"`
	PhaseDefinition      *Phase `json:"phase_definition,omitempty"`
	Policy               string `json:"policy"`
	Version              int64  `json:"version"`
}

func (s *LifecycleExplainPhaseExecution) UnmarshalJSON(data []byte) error {

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

		case "modified_date_in_millis":
			if err := dec.Decode(&s.ModifiedDateInMillis); err != nil {
				return fmt.Errorf("%s | %w", "ModifiedDateInMillis", err)
			}

		case "phase_definition":
			if err := dec.Decode(&s.PhaseDefinition); err != nil {
				return fmt.Errorf("%s | %w", "PhaseDefinition", err)
			}

		case "policy":
			if err := dec.Decode(&s.Policy); err != nil {
				return fmt.Errorf("%s | %w", "Policy", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// NewLifecycleExplainPhaseExecution returns a LifecycleExplainPhaseExecution.
func NewLifecycleExplainPhaseExecution() *LifecycleExplainPhaseExecution {
	r := &LifecycleExplainPhaseExecution{}

	return r
}

// false
