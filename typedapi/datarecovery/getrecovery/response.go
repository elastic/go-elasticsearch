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

package getrecovery

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/recoverymode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/recoveryrunstatus"
)

// Response holds the response body struct for the package getrecovery
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/data_recovery/get_recovery/DataRecoveryGetRecoveryResponse.ts#L22-L25
type Response struct {
	// CompletedAt The time when the recovery finished. It is omitted while the recovery is
	// running.
	CompletedAt types.DateTime `json:"completed_at,omitempty"`
	// DestinationMode Whether the sources are recovered in place (`replace`) or under new names
	// (`rename`).
	DestinationMode recoverymode.RecoveryMode `json:"destination_mode"`
	// Failures The sources that could not be recovered. The list is capped, so it can be
	// shorter than `total_failure_count`.
	Failures []types.RecoveryFailure `json:"failures"`
	// Metadata The custom metadata that was provided when the recovery was started.
	Metadata types.Metadata `json:"metadata"`
	// Progress The progress of the recovery.
	Progress types.RecoveryProgress `json:"progress"`
	// RecoveryId The identifier of the recovery.
	RecoveryId string `json:"recovery_id"`
	// RecoveryPoint The recovery point that the sources are recovered from.
	RecoveryPoint types.Point `json:"recovery_point"`
	// Sources The sources that the requested source expressions resolved to, with the
	// destinations that they are recovered to.
	Sources []types.ResolvedRecoverySource `json:"sources"`
	// StartedAt The time when the recovery started.
	StartedAt types.DateTime `json:"started_at"`
	// Status The status of the recovery.
	Status recoveryrunstatus.RecoveryRunStatus `json:"status"`
	// TotalFailureCount The total number of sources that could not be recovered, including any that
	// are not listed in `failures`.
	TotalFailureCount int `json:"total_failure_count"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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

		case "completed_at":
			if err := dec.Decode(&s.CompletedAt); err != nil {
				return fmt.Errorf("%s | %w", "CompletedAt", err)
			}

		case "destination_mode":
			if err := dec.Decode(&s.DestinationMode); err != nil {
				return fmt.Errorf("%s | %w", "DestinationMode", err)
			}

		case "failures":
			if err := dec.Decode(&s.Failures); err != nil {
				return fmt.Errorf("%s | %w", "Failures", err)
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "progress":
			if err := dec.Decode(&s.Progress); err != nil {
				return fmt.Errorf("%s | %w", "Progress", err)
			}

		case "recovery_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RecoveryId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RecoveryId = o

		case "recovery_point":
			if err := dec.Decode(&s.RecoveryPoint); err != nil {
				return fmt.Errorf("%s | %w", "RecoveryPoint", err)
			}

		case "sources":
			if err := dec.Decode(&s.Sources); err != nil {
				return fmt.Errorf("%s | %w", "Sources", err)
			}

		case "started_at":
			if err := dec.Decode(&s.StartedAt); err != nil {
				return fmt.Errorf("%s | %w", "StartedAt", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "total_failure_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalFailureCount", err)
				}
				s.TotalFailureCount = value
			case float64:
				f := int(v)
				s.TotalFailureCount = f
			}

		}
	}
	return nil
}
