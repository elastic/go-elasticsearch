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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// LifecycleExplainManaged type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/ilm/explain_lifecycle/types.ts#L26-L45
type LifecycleExplainManaged struct {
	Action                  Name                           `json:"action"`
	ActionTimeMillis        EpochMillis                    `json:"action_time_millis"`
	Age                     Time                           `json:"age"`
	FailedStep              *Name                          `json:"failed_step,omitempty"`
	FailedStepRetryCount    *int                           `json:"failed_step_retry_count,omitempty"`
	Index                   IndexName                      `json:"index"`
	IndexCreationDateMillis *EpochMillis                   `json:"index_creation_date_millis,omitempty"`
	IsAutoRetryableError    *bool                          `json:"is_auto_retryable_error,omitempty"`
	LifecycleDateMillis     EpochMillis                    `json:"lifecycle_date_millis"`
	Managed                 string                         `json:"managed,omitempty"`
	Phase                   Name                           `json:"phase"`
	PhaseExecution          LifecycleExplainPhaseExecution `json:"phase_execution"`
	PhaseTimeMillis         EpochMillis                    `json:"phase_time_millis"`
	Policy                  Name                           `json:"policy"`
	Step                    Name                           `json:"step"`
	StepInfo                map[string]interface{}         `json:"step_info,omitempty"`
	StepTimeMillis          EpochMillis                    `json:"step_time_millis"`
	TimeSinceIndexCreation  *Time                          `json:"time_since_index_creation,omitempty"`
}

// LifecycleExplainManagedBuilder holds LifecycleExplainManaged struct and provides a builder API.
type LifecycleExplainManagedBuilder struct {
	v *LifecycleExplainManaged
}

// NewLifecycleExplainManaged provides a builder for the LifecycleExplainManaged struct.
func NewLifecycleExplainManagedBuilder() *LifecycleExplainManagedBuilder {
	r := LifecycleExplainManagedBuilder{
		&LifecycleExplainManaged{
			StepInfo: make(map[string]interface{}, 0),
		},
	}

	r.v.Managed = "true"

	return &r
}

// Build finalize the chain and returns the LifecycleExplainManaged struct
func (rb *LifecycleExplainManagedBuilder) Build() LifecycleExplainManaged {
	return *rb.v
}

func (rb *LifecycleExplainManagedBuilder) Action(action Name) *LifecycleExplainManagedBuilder {
	rb.v.Action = action
	return rb
}

func (rb *LifecycleExplainManagedBuilder) ActionTimeMillis(actiontimemillis *EpochMillisBuilder) *LifecycleExplainManagedBuilder {
	v := actiontimemillis.Build()
	rb.v.ActionTimeMillis = v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) Age(age *TimeBuilder) *LifecycleExplainManagedBuilder {
	v := age.Build()
	rb.v.Age = v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) FailedStep(failedstep Name) *LifecycleExplainManagedBuilder {
	rb.v.FailedStep = &failedstep
	return rb
}

func (rb *LifecycleExplainManagedBuilder) FailedStepRetryCount(failedstepretrycount int) *LifecycleExplainManagedBuilder {
	rb.v.FailedStepRetryCount = &failedstepretrycount
	return rb
}

func (rb *LifecycleExplainManagedBuilder) Index(index IndexName) *LifecycleExplainManagedBuilder {
	rb.v.Index = index
	return rb
}

func (rb *LifecycleExplainManagedBuilder) IndexCreationDateMillis(indexcreationdatemillis *EpochMillisBuilder) *LifecycleExplainManagedBuilder {
	v := indexcreationdatemillis.Build()
	rb.v.IndexCreationDateMillis = &v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) IsAutoRetryableError(isautoretryableerror bool) *LifecycleExplainManagedBuilder {
	rb.v.IsAutoRetryableError = &isautoretryableerror
	return rb
}

func (rb *LifecycleExplainManagedBuilder) LifecycleDateMillis(lifecycledatemillis *EpochMillisBuilder) *LifecycleExplainManagedBuilder {
	v := lifecycledatemillis.Build()
	rb.v.LifecycleDateMillis = v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) Phase(phase Name) *LifecycleExplainManagedBuilder {
	rb.v.Phase = phase
	return rb
}

func (rb *LifecycleExplainManagedBuilder) PhaseExecution(phaseexecution *LifecycleExplainPhaseExecutionBuilder) *LifecycleExplainManagedBuilder {
	v := phaseexecution.Build()
	rb.v.PhaseExecution = v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) PhaseTimeMillis(phasetimemillis *EpochMillisBuilder) *LifecycleExplainManagedBuilder {
	v := phasetimemillis.Build()
	rb.v.PhaseTimeMillis = v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) Policy(policy Name) *LifecycleExplainManagedBuilder {
	rb.v.Policy = policy
	return rb
}

func (rb *LifecycleExplainManagedBuilder) Step(step Name) *LifecycleExplainManagedBuilder {
	rb.v.Step = step
	return rb
}

func (rb *LifecycleExplainManagedBuilder) StepInfo(value map[string]interface{}) *LifecycleExplainManagedBuilder {
	rb.v.StepInfo = value
	return rb
}

func (rb *LifecycleExplainManagedBuilder) StepTimeMillis(steptimemillis *EpochMillisBuilder) *LifecycleExplainManagedBuilder {
	v := steptimemillis.Build()
	rb.v.StepTimeMillis = v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) TimeSinceIndexCreation(timesinceindexcreation *TimeBuilder) *LifecycleExplainManagedBuilder {
	v := timesinceindexcreation.Build()
	rb.v.TimeSinceIndexCreation = &v
	return rb
}
