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
// https://github.com/elastic/elasticsearch-specification/tree/3ea9ce260df22d3244bff5bace485dd97ff4046d

package types

import (
	"encoding/json"
	"fmt"
)

// TriggerContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3ea9ce260df22d3244bff5bace485dd97ff4046d/specification/watcher/_types/Trigger.ts#L23-L28
type TriggerContainer struct {
	AdditionalTriggerContainerProperty map[string]json.RawMessage `json:"-"`
	Schedule                           *ScheduleContainer         `json:"schedule,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s TriggerContainer) MarshalJSON() ([]byte, error) {
	type opt TriggerContainer
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalTriggerContainerProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalTriggerContainerProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewTriggerContainer returns a TriggerContainer.
func NewTriggerContainer() *TriggerContainer {
	r := &TriggerContainer{
		AdditionalTriggerContainerProperty: make(map[string]json.RawMessage),
	}

	return r
}

// true

type TriggerContainerVariant interface {
	TriggerContainerCaster() *TriggerContainer
}

func (s *TriggerContainer) TriggerContainerCaster() *TriggerContainer {
	return s
}
