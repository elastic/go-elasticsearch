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
// https://github.com/elastic/elasticsearch-specification/tree/656080dee792f93a849cd7fbf566f528f858a5ea

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conditionop"
)

// ArrayCompareCondition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/656080dee792f93a849cd7fbf566f528f858a5ea/specification/watcher/_types/Conditions.ts#L32-L36
type ArrayCompareCondition struct {
	ArrayCompareCondition map[conditionop.ConditionOp]ArrayCompareOpParams `json:"ArrayCompareCondition,omitempty"`
	Path                  string                                           `json:"path"`
}

func (s *ArrayCompareCondition) UnmarshalJSON(data []byte) error {

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

		case "ArrayCompareCondition":
			if s.ArrayCompareCondition == nil {
				s.ArrayCompareCondition = make(map[conditionop.ConditionOp]ArrayCompareOpParams, 0)
			}
			if err := dec.Decode(&s.ArrayCompareCondition); err != nil {
				return err
			}

		case "path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Path = o

		default:

		}
	}
	return nil
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
	delete(tmp, "ArrayCompareCondition")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewArrayCompareCondition returns a ArrayCompareCondition.
func NewArrayCompareCondition() *ArrayCompareCondition {
	r := &ArrayCompareCondition{
		ArrayCompareCondition: make(map[conditionop.ConditionOp]ArrayCompareOpParams, 0),
	}

	return r
}
