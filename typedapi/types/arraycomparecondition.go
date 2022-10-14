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
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


package types

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conditionop"
)

// ArrayCompareCondition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/watcher/_types/Conditions.ts#L32-L36
type ArrayCompareCondition struct {
	ArrayCompareCondition map[conditionop.ConditionOp]ArrayCompareOpParams `json:"-"`
	Path                  string                                           `json:"path"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s ArrayCompareCondition) MarshalJSON() ([]byte, error) {
	type opt ArrayCompareCondition
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.ArrayCompareCondition {
		tmp[fmt.Sprintf("%s", key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ArrayCompareConditionBuilder holds ArrayCompareCondition struct and provides a builder API.
type ArrayCompareConditionBuilder struct {
	v *ArrayCompareCondition
}

// NewArrayCompareCondition provides a builder for the ArrayCompareCondition struct.
func NewArrayCompareConditionBuilder() *ArrayCompareConditionBuilder {
	r := ArrayCompareConditionBuilder{
		&ArrayCompareCondition{
			ArrayCompareCondition: make(map[conditionop.ConditionOp]ArrayCompareOpParams, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ArrayCompareCondition struct
func (rb *ArrayCompareConditionBuilder) Build() ArrayCompareCondition {
	return *rb.v
}

func (rb *ArrayCompareConditionBuilder) ArrayCompareCondition(values map[conditionop.ConditionOp]*ArrayCompareOpParamsBuilder) *ArrayCompareConditionBuilder {
	tmp := make(map[conditionop.ConditionOp]ArrayCompareOpParams, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.ArrayCompareCondition = tmp
	return rb
}

func (rb *ArrayCompareConditionBuilder) Path(path string) *ArrayCompareConditionBuilder {
	rb.v.Path = path
	return rb
}
