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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"encoding/json"
	"fmt"
)

// RetentionPolicyContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/transform/_types/Transform.ts#L80-L86
type RetentionPolicyContainer struct {
	AdditionalRetentionPolicyContainerProperty map[string]json.RawMessage `json:"-"`
	// Time Specifies that the transform uses a time field to set the retention policy.
	Time *RetentionPolicy `json:"time,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s RetentionPolicyContainer) MarshalJSON() ([]byte, error) {
	type opt RetentionPolicyContainer
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
	for key, value := range s.AdditionalRetentionPolicyContainerProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalRetentionPolicyContainerProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewRetentionPolicyContainer returns a RetentionPolicyContainer.
func NewRetentionPolicyContainer() *RetentionPolicyContainer {
	r := &RetentionPolicyContainer{
		AdditionalRetentionPolicyContainerProperty: make(map[string]json.RawMessage),
	}

	return r
}
