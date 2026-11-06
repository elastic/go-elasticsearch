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

package startrecovery

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/recoverymode"
)

// Request holds the request body struct for the package startrecovery
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/data_recovery/start_recovery/DataRecoveryStartRecoveryRequest.ts#L25-L76
type Request struct {
	// Metadata Custom metadata for the recovery.
	Metadata types.Metadata `json:"metadata,omitempty"`
	// Mode How to recover the sources: `replace` recovers them in place, over the data
	// that currently exists, and `rename` recovers them under new names.
	Mode recoverymode.RecoveryMode `json:"mode"`
	// RecoveryPointEndTime The exact end time of the recovery point to recover from, as a date or epoch
	// milliseconds. An `end_time` returned when listing recovery points can be
	// passed back as it is.
	RecoveryPointEndTime types.DateTime `json:"recovery_point_end_time"`
	// RenameStrategy How to derive the new name of each recovered source from its name. Required
	// when `mode` is `rename`, and not allowed when `mode` is `replace`.
	RenameStrategy *types.RenameStrategy `json:"rename_strategy,omitempty"`
	// Sources The sources to recover. Each value is an index or data stream expression.
	// Wildcards and exclusions prefixed with `-` are supported. The list must not
	// be empty and must not contain blank values.
	Sources []string `json:"sources"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Startrecovery request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "mode":
			if err := dec.Decode(&s.Mode); err != nil {
				return fmt.Errorf("%s | %w", "Mode", err)
			}

		case "recovery_point_end_time":
			if err := dec.Decode(&s.RecoveryPointEndTime); err != nil {
				return fmt.Errorf("%s | %w", "RecoveryPointEndTime", err)
			}

		case "rename_strategy":
			if err := dec.Decode(&s.RenameStrategy); err != nil {
				return fmt.Errorf("%s | %w", "RenameStrategy", err)
			}

		case "sources":
			if err := dec.Decode(&s.Sources); err != nil {
				return fmt.Errorf("%s | %w", "Sources", err)
			}

		}
	}
	return nil
}
