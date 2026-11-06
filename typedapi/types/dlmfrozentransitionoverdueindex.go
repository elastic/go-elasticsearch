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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dlmfrozentransitionstate"
)

// DlmFrozenTransitionOverdueIndex type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/health_report/types.ts#L292-L297
type DlmFrozenTransitionOverdueIndex struct {
	// IndexName The name of the overdue index.
	IndexName string `json:"index_name"`
	// TransitionState The transition state of the overdue index.
	TransitionState dlmfrozentransitionstate.DlmFrozenTransitionState `json:"transition_state"`
}

func (s *DlmFrozenTransitionOverdueIndex) UnmarshalJSON(data []byte) error {

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

		case "index_name":
			if err := dec.Decode(&s.IndexName); err != nil {
				return fmt.Errorf("%s | %w", "IndexName", err)
			}

		case "transition_state":
			if err := dec.Decode(&s.TransitionState); err != nil {
				return fmt.Errorf("%s | %w", "TransitionState", err)
			}

		}
	}
	return nil
}

// NewDlmFrozenTransitionOverdueIndex returns a DlmFrozenTransitionOverdueIndex.
func NewDlmFrozenTransitionOverdueIndex() *DlmFrozenTransitionOverdueIndex {
	r := &DlmFrozenTransitionOverdueIndex{}

	return r
}
