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
	"strconv"
)

// ScriptRescoreProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/search/_types/profile.ts#L209-L213
type ScriptRescoreProfile struct {
	// Breakdown The `script` rescorer does not report any steps.
	Breakdown EmptyObject `json:"breakdown"`
	// Debug Extra details about the rescorer.
	//
	// Every rescorer reports `window_size`, the number of top documents it was
	// configured to rescore. `timed_out` is set to `true` if the search timed out
	// while the rescorer was running. Rescorers can add their own keys.
	Debug map[string]json.RawMessage `json:"debug,omitempty"`
	// Description What the rescorer runs.
	Description string `json:"description"`
	// Time The time spent in the rescorer, in human-readable units. Only returned when
	// `human` is `true`.
	Time Duration `json:"time,omitempty"`
	// TimeInNanos The total time spent in the rescorer, including the time spent running its
	// queries.
	TimeInNanos int64  `json:"time_in_nanos"`
	Type        string `json:"type,omitempty"`
}

func (s *ScriptRescoreProfile) UnmarshalJSON(data []byte) error {

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

		case "breakdown":
			if err := dec.Decode(&s.Breakdown); err != nil {
				return fmt.Errorf("%s | %w", "Breakdown", err)
			}

		case "debug":
			if s.Debug == nil {
				s.Debug = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Debug); err != nil {
				return fmt.Errorf("%s | %w", "Debug", err)
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = o

		case "time":
			if err := dec.Decode(&s.Time); err != nil {
				return fmt.Errorf("%s | %w", "Time", err)
			}

		case "time_in_nanos":
			if err := dec.Decode(&s.TimeInNanos); err != nil {
				return fmt.Errorf("%s | %w", "TimeInNanos", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s ScriptRescoreProfile) MarshalJSON() ([]byte, error) {
	type innerScriptRescoreProfile ScriptRescoreProfile
	tmp := innerScriptRescoreProfile{
		Breakdown:   s.Breakdown,
		Debug:       s.Debug,
		Description: s.Description,
		Time:        s.Time,
		TimeInNanos: s.TimeInNanos,
		Type:        s.Type,
	}

	tmp.Type = "script"

	return json.Marshal(tmp)
}

// NewScriptRescoreProfile returns a ScriptRescoreProfile.
func NewScriptRescoreProfile() *ScriptRescoreProfile {
	r := &ScriptRescoreProfile{
		Debug: make(map[string]json.RawMessage),
	}

	return r
}
