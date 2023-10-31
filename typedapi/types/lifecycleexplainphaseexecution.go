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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// LifecycleExplainPhaseExecution type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ilm/explain_lifecycle/types.ts#L64-L68
type LifecycleExplainPhaseExecution struct {
	ModifiedDateInMillis int64  `json:"modified_date_in_millis"`
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
				return err
			}

		case "policy":
			if err := dec.Decode(&s.Policy); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
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
