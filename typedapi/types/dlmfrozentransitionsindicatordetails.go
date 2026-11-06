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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dlmfrozentransitionstate"
)

// An index is overdue when it has passed its `frozen_after` age by more than
// the configured threshold and has not yet completed its transition to the
// frozen tier. Indices with a running transition are making progress and are
// not counted as overdue.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/health_report/types.ts#L250-L279
type DlmFrozenTransitionsIndicatorDetails struct {
	// DefaultRepositoryConfigured Indicates whether a default snapshot repository is configured with the
	// `repositories.default_repository` cluster setting.
	DefaultRepositoryConfigured *bool `json:"default_repository_configured,omitempty"`
	// GeneratedAtMillis The time, in milliseconds since the Unix epoch, when the master node
	// generated the health snapshot. Only present when the indicator status is
	// `unknown`, meaning the health snapshot is stale.
	GeneratedAtMillis *int64 `json:"generated_at_millis,omitempty"`
	// OverdueIndicesCount The total number of overdue indices in the cluster.
	OverdueIndicesCount *int `json:"overdue_indices_count,omitempty"`
	// OverdueIndicesCountByState The number of overdue indices in the cluster for each transition state. Only
	// present when there are overdue indices.
	OverdueIndicesCountByState map[dlmfrozentransitionstate.DlmFrozenTransitionState]int `json:"overdue_indices_count_by_state,omitempty"`
	// OverdueIndicesSample A sample of the overdue indices and their transition states, sorted by index
	// name. The sample is limited to 100 indices for each transition state, so it
	// can contain fewer indices than `overdue_indices_count`. Only present when
	// there are overdue indices.
	OverdueIndicesSample []DlmFrozenTransitionOverdueIndex `json:"overdue_indices_sample,omitempty"`
	// ServiceRunning Indicates whether the frozen transition service is running on the current
	// master node.
	ServiceRunning *bool `json:"service_running,omitempty"`
	// TransitionsEnabled Indicates whether frozen-tier transitions are enabled by the
	// `dlm.frozen_transitions.enabled` cluster setting.
	TransitionsEnabled *bool `json:"transitions_enabled,omitempty"`
}

func (s *DlmFrozenTransitionsIndicatorDetails) UnmarshalJSON(data []byte) error {

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

		case "default_repository_configured":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DefaultRepositoryConfigured", err)
				}
				s.DefaultRepositoryConfigured = &value
			case bool:
				s.DefaultRepositoryConfigured = &v
			}

		case "generated_at_millis":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "GeneratedAtMillis", err)
				}
				s.GeneratedAtMillis = &value
			case float64:
				f := int64(v)
				s.GeneratedAtMillis = &f
			}

		case "overdue_indices_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OverdueIndicesCount", err)
				}
				s.OverdueIndicesCount = &value
			case float64:
				f := int(v)
				s.OverdueIndicesCount = &f
			}

		case "overdue_indices_count_by_state":
			if s.OverdueIndicesCountByState == nil {
				s.OverdueIndicesCountByState = make(map[dlmfrozentransitionstate.DlmFrozenTransitionState]int, 0)
			}
			if err := dec.Decode(&s.OverdueIndicesCountByState); err != nil {
				return fmt.Errorf("%s | %w", "OverdueIndicesCountByState", err)
			}

		case "overdue_indices_sample":
			if err := dec.Decode(&s.OverdueIndicesSample); err != nil {
				return fmt.Errorf("%s | %w", "OverdueIndicesSample", err)
			}

		case "service_running":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ServiceRunning", err)
				}
				s.ServiceRunning = &value
			case bool:
				s.ServiceRunning = &v
			}

		case "transitions_enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TransitionsEnabled", err)
				}
				s.TransitionsEnabled = &value
			case bool:
				s.TransitionsEnabled = &v
			}

		}
	}
	return nil
}

// NewDlmFrozenTransitionsIndicatorDetails returns a DlmFrozenTransitionsIndicatorDetails.
func NewDlmFrozenTransitionsIndicatorDetails() *DlmFrozenTransitionsIndicatorDetails {
	r := &DlmFrozenTransitionsIndicatorDetails{
		OverdueIndicesCountByState: make(map[dlmfrozentransitionstate.DlmFrozenTransitionState]int),
	}

	return r
}
