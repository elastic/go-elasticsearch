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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// LifecycleExplainManaged type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ilm/explain_lifecycle/types.ts#L26-L52
type LifecycleExplainManaged struct {
	Action                  *string                         `json:"action,omitempty"`
	ActionTime              DateTime                        `json:"action_time,omitempty"`
	ActionTimeMillis        *int64                          `json:"action_time_millis,omitempty"`
	Age                     Duration                        `json:"age,omitempty"`
	FailedStep              *string                         `json:"failed_step,omitempty"`
	FailedStepRetryCount    *int                            `json:"failed_step_retry_count,omitempty"`
	Index                   *string                         `json:"index,omitempty"`
	IndexCreationDate       DateTime                        `json:"index_creation_date,omitempty"`
	IndexCreationDateMillis *int64                          `json:"index_creation_date_millis,omitempty"`
	IsAutoRetryableError    *bool                           `json:"is_auto_retryable_error,omitempty"`
	LifecycleDate           DateTime                        `json:"lifecycle_date,omitempty"`
	LifecycleDateMillis     *int64                          `json:"lifecycle_date_millis,omitempty"`
	Managed                 bool                            `json:"managed,omitempty"`
	Phase                   string                          `json:"phase"`
	PhaseExecution          *LifecycleExplainPhaseExecution `json:"phase_execution,omitempty"`
	PhaseTime               DateTime                        `json:"phase_time,omitempty"`
	PhaseTimeMillis         *int64                          `json:"phase_time_millis,omitempty"`
	Policy                  string                          `json:"policy"`
	Step                    *string                         `json:"step,omitempty"`
	StepInfo                map[string]json.RawMessage      `json:"step_info,omitempty"`
	StepTime                DateTime                        `json:"step_time,omitempty"`
	StepTimeMillis          *int64                          `json:"step_time_millis,omitempty"`
	TimeSinceIndexCreation  Duration                        `json:"time_since_index_creation,omitempty"`
}

func (s *LifecycleExplainManaged) UnmarshalJSON(data []byte) error {

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

		case "action":
			if err := dec.Decode(&s.Action); err != nil {
				return err
			}

		case "action_time":
			if err := dec.Decode(&s.ActionTime); err != nil {
				return err
			}

		case "action_time_millis":
			if err := dec.Decode(&s.ActionTimeMillis); err != nil {
				return err
			}

		case "age":
			if err := dec.Decode(&s.Age); err != nil {
				return err
			}

		case "failed_step":
			if err := dec.Decode(&s.FailedStep); err != nil {
				return err
			}

		case "failed_step_retry_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FailedStepRetryCount = &value
			case float64:
				f := int(v)
				s.FailedStepRetryCount = &f
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "index_creation_date":
			if err := dec.Decode(&s.IndexCreationDate); err != nil {
				return err
			}

		case "index_creation_date_millis":
			if err := dec.Decode(&s.IndexCreationDateMillis); err != nil {
				return err
			}

		case "is_auto_retryable_error":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IsAutoRetryableError = &value
			case bool:
				s.IsAutoRetryableError = &v
			}

		case "lifecycle_date":
			if err := dec.Decode(&s.LifecycleDate); err != nil {
				return err
			}

		case "lifecycle_date_millis":
			if err := dec.Decode(&s.LifecycleDateMillis); err != nil {
				return err
			}

		case "managed":
			if err := dec.Decode(&s.Managed); err != nil {
				return err
			}

		case "phase":
			if err := dec.Decode(&s.Phase); err != nil {
				return err
			}

		case "phase_execution":
			if err := dec.Decode(&s.PhaseExecution); err != nil {
				return err
			}

		case "phase_time":
			if err := dec.Decode(&s.PhaseTime); err != nil {
				return err
			}

		case "phase_time_millis":
			if err := dec.Decode(&s.PhaseTimeMillis); err != nil {
				return err
			}

		case "policy":
			if err := dec.Decode(&s.Policy); err != nil {
				return err
			}

		case "step":
			if err := dec.Decode(&s.Step); err != nil {
				return err
			}

		case "step_info":
			if s.StepInfo == nil {
				s.StepInfo = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.StepInfo); err != nil {
				return err
			}

		case "step_time":
			if err := dec.Decode(&s.StepTime); err != nil {
				return err
			}

		case "step_time_millis":
			if err := dec.Decode(&s.StepTimeMillis); err != nil {
				return err
			}

		case "time_since_index_creation":
			if err := dec.Decode(&s.TimeSinceIndexCreation); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewLifecycleExplainManaged returns a LifecycleExplainManaged.
func NewLifecycleExplainManaged() *LifecycleExplainManaged {
	r := &LifecycleExplainManaged{
		StepInfo: make(map[string]json.RawMessage, 0),
	}

	r.Managed = true

	return r
}
