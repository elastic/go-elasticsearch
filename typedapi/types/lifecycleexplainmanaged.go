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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// LifecycleExplainManaged type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ilm/explain_lifecycle/types.ts#L26-L52
type LifecycleExplainManaged struct {
	Action                  *Name                           `json:"action,omitempty"`
	ActionTime              *DateTime                       `json:"action_time,omitempty"`
	ActionTimeMillis        *EpochTimeUnitMillis            `json:"action_time_millis,omitempty"`
	Age                     *Duration                       `json:"age,omitempty"`
	FailedStep              *Name                           `json:"failed_step,omitempty"`
	FailedStepRetryCount    *int                            `json:"failed_step_retry_count,omitempty"`
	Index                   *IndexName                      `json:"index,omitempty"`
	IndexCreationDate       *DateTime                       `json:"index_creation_date,omitempty"`
	IndexCreationDateMillis *EpochTimeUnitMillis            `json:"index_creation_date_millis,omitempty"`
	IsAutoRetryableError    *bool                           `json:"is_auto_retryable_error,omitempty"`
	LifecycleDate           *DateTime                       `json:"lifecycle_date,omitempty"`
	LifecycleDateMillis     *EpochTimeUnitMillis            `json:"lifecycle_date_millis,omitempty"`
	Managed                 string                          `json:"managed,omitempty"`
	Phase                   Name                            `json:"phase"`
	PhaseExecution          *LifecycleExplainPhaseExecution `json:"phase_execution,omitempty"`
	PhaseTime               *DateTime                       `json:"phase_time,omitempty"`
	PhaseTimeMillis         *EpochTimeUnitMillis            `json:"phase_time_millis,omitempty"`
	Policy                  Name                            `json:"policy"`
	Step                    *Name                           `json:"step,omitempty"`
	StepInfo                map[string]interface{}          `json:"step_info,omitempty"`
	StepTime                *DateTime                       `json:"step_time,omitempty"`
	StepTimeMillis          *EpochTimeUnitMillis            `json:"step_time_millis,omitempty"`
	TimeSinceIndexCreation  *Duration                       `json:"time_since_index_creation,omitempty"`
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
	rb.v.Action = &action
	return rb
}

func (rb *LifecycleExplainManagedBuilder) ActionTime(actiontime *DateTimeBuilder) *LifecycleExplainManagedBuilder {
	v := actiontime.Build()
	rb.v.ActionTime = &v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) ActionTimeMillis(actiontimemillis *EpochTimeUnitMillisBuilder) *LifecycleExplainManagedBuilder {
	v := actiontimemillis.Build()
	rb.v.ActionTimeMillis = &v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) Age(age *DurationBuilder) *LifecycleExplainManagedBuilder {
	v := age.Build()
	rb.v.Age = &v
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
	rb.v.Index = &index
	return rb
}

func (rb *LifecycleExplainManagedBuilder) IndexCreationDate(indexcreationdate *DateTimeBuilder) *LifecycleExplainManagedBuilder {
	v := indexcreationdate.Build()
	rb.v.IndexCreationDate = &v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) IndexCreationDateMillis(indexcreationdatemillis *EpochTimeUnitMillisBuilder) *LifecycleExplainManagedBuilder {
	v := indexcreationdatemillis.Build()
	rb.v.IndexCreationDateMillis = &v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) IsAutoRetryableError(isautoretryableerror bool) *LifecycleExplainManagedBuilder {
	rb.v.IsAutoRetryableError = &isautoretryableerror
	return rb
}

func (rb *LifecycleExplainManagedBuilder) LifecycleDate(lifecycledate *DateTimeBuilder) *LifecycleExplainManagedBuilder {
	v := lifecycledate.Build()
	rb.v.LifecycleDate = &v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) LifecycleDateMillis(lifecycledatemillis *EpochTimeUnitMillisBuilder) *LifecycleExplainManagedBuilder {
	v := lifecycledatemillis.Build()
	rb.v.LifecycleDateMillis = &v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) Phase(phase Name) *LifecycleExplainManagedBuilder {
	rb.v.Phase = phase
	return rb
}

func (rb *LifecycleExplainManagedBuilder) PhaseExecution(phaseexecution *LifecycleExplainPhaseExecutionBuilder) *LifecycleExplainManagedBuilder {
	v := phaseexecution.Build()
	rb.v.PhaseExecution = &v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) PhaseTime(phasetime *DateTimeBuilder) *LifecycleExplainManagedBuilder {
	v := phasetime.Build()
	rb.v.PhaseTime = &v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) PhaseTimeMillis(phasetimemillis *EpochTimeUnitMillisBuilder) *LifecycleExplainManagedBuilder {
	v := phasetimemillis.Build()
	rb.v.PhaseTimeMillis = &v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) Policy(policy Name) *LifecycleExplainManagedBuilder {
	rb.v.Policy = policy
	return rb
}

func (rb *LifecycleExplainManagedBuilder) Step(step Name) *LifecycleExplainManagedBuilder {
	rb.v.Step = &step
	return rb
}

func (rb *LifecycleExplainManagedBuilder) StepInfo(value map[string]interface{}) *LifecycleExplainManagedBuilder {
	rb.v.StepInfo = value
	return rb
}

func (rb *LifecycleExplainManagedBuilder) StepTime(steptime *DateTimeBuilder) *LifecycleExplainManagedBuilder {
	v := steptime.Build()
	rb.v.StepTime = &v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) StepTimeMillis(steptimemillis *EpochTimeUnitMillisBuilder) *LifecycleExplainManagedBuilder {
	v := steptimemillis.Build()
	rb.v.StepTimeMillis = &v
	return rb
}

func (rb *LifecycleExplainManagedBuilder) TimeSinceIndexCreation(timesinceindexcreation *DurationBuilder) *LifecycleExplainManagedBuilder {
	v := timesinceindexcreation.Build()
	rb.v.TimeSinceIndexCreation = &v
	return rb
}
