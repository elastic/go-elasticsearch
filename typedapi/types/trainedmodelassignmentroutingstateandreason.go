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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/routingstate"
)

// TrainedModelAssignmentRoutingStateAndReason type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/TrainedModel.ts#L431-L441
type TrainedModelAssignmentRoutingStateAndReason struct {
	// Reason The reason for the current state. It is usually populated only when the
	// `routing_state` is `failed`.
	Reason *string `json:"reason,omitempty"`
	// RoutingState The current routing state.
	RoutingState routingstate.RoutingState `json:"routing_state"`
}

func (s *TrainedModelAssignmentRoutingStateAndReason) UnmarshalJSON(data []byte) error {

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

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Reason", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = &o

		case "routing_state":
			if err := dec.Decode(&s.RoutingState); err != nil {
				return fmt.Errorf("%s | %w", "RoutingState", err)
			}

		}
	}
	return nil
}

// NewTrainedModelAssignmentRoutingStateAndReason returns a TrainedModelAssignmentRoutingStateAndReason.
func NewTrainedModelAssignmentRoutingStateAndReason() *TrainedModelAssignmentRoutingStateAndReason {
	r := &TrainedModelAssignmentRoutingStateAndReason{}

	return r
}
