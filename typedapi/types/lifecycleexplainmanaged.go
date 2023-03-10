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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"encoding/json"
)

// LifecycleExplainManaged type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ilm/explain_lifecycle/types.ts#L26-L52
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

// NewLifecycleExplainManaged returns a LifecycleExplainManaged.
func NewLifecycleExplainManaged() *LifecycleExplainManaged {
	r := &LifecycleExplainManaged{
		StepInfo: make(map[string]json.RawMessage, 0),
	}

	r.Managed = true

	return r
}
