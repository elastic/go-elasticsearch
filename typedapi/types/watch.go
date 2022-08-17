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

// Watch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Watch.ts#L37-L47
type Watch struct {
	Actions                map[IndexName]Action     `json:"actions"`
	Condition              ConditionContainer       `json:"condition"`
	Input                  InputContainer           `json:"input"`
	Metadata               *Metadata                `json:"metadata,omitempty"`
	Status                 *WatchStatus             `json:"status,omitempty"`
	ThrottlePeriod         *Duration                `json:"throttle_period,omitempty"`
	ThrottlePeriodInMillis *DurationValueUnitMillis `json:"throttle_period_in_millis,omitempty"`
	Transform              *TransformContainer      `json:"transform,omitempty"`
	Trigger                TriggerContainer         `json:"trigger"`
}

// WatchBuilder holds Watch struct and provides a builder API.
type WatchBuilder struct {
	v *Watch
}

// NewWatch provides a builder for the Watch struct.
func NewWatchBuilder() *WatchBuilder {
	r := WatchBuilder{
		&Watch{
			Actions: make(map[IndexName]Action, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Watch struct
func (rb *WatchBuilder) Build() Watch {
	return *rb.v
}

func (rb *WatchBuilder) Actions(values map[IndexName]*ActionBuilder) *WatchBuilder {
	tmp := make(map[IndexName]Action, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Actions = tmp
	return rb
}

func (rb *WatchBuilder) Condition(condition *ConditionContainerBuilder) *WatchBuilder {
	v := condition.Build()
	rb.v.Condition = v
	return rb
}

func (rb *WatchBuilder) Input(input *InputContainerBuilder) *WatchBuilder {
	v := input.Build()
	rb.v.Input = v
	return rb
}

func (rb *WatchBuilder) Metadata(metadata *MetadataBuilder) *WatchBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

func (rb *WatchBuilder) Status(status *WatchStatusBuilder) *WatchBuilder {
	v := status.Build()
	rb.v.Status = &v
	return rb
}

func (rb *WatchBuilder) ThrottlePeriod(throttleperiod *DurationBuilder) *WatchBuilder {
	v := throttleperiod.Build()
	rb.v.ThrottlePeriod = &v
	return rb
}

func (rb *WatchBuilder) ThrottlePeriodInMillis(throttleperiodinmillis *DurationValueUnitMillisBuilder) *WatchBuilder {
	v := throttleperiodinmillis.Build()
	rb.v.ThrottlePeriodInMillis = &v
	return rb
}

func (rb *WatchBuilder) Transform(transform *TransformContainerBuilder) *WatchBuilder {
	v := transform.Build()
	rb.v.Transform = &v
	return rb
}

func (rb *WatchBuilder) Trigger(trigger *TriggerContainerBuilder) *WatchBuilder {
	v := trigger.Build()
	rb.v.Trigger = v
	return rb
}
