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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// IlmPolicyStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/52c473efb1fb5320a5bac12572d0b285882862fb/specification/xpack/usage/types.ts#L167-L170
type IlmPolicyStatistics struct {
	IndicesManaged int         `json:"indices_managed"`
	Phases         UsagePhases `json:"phases"`
}

func (s *IlmPolicyStatistics) UnmarshalJSON(data []byte) error {

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

		case "indices_managed":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndicesManaged", err)
				}
				s.IndicesManaged = value
			case float64:
				f := int(v)
				s.IndicesManaged = f
			}

		case "phases":
			if err := dec.Decode(&s.Phases); err != nil {
				return fmt.Errorf("%s | %w", "Phases", err)
			}

		}
	}
	return nil
}

// NewIlmPolicyStatistics returns a IlmPolicyStatistics.
func NewIlmPolicyStatistics() *IlmPolicyStatistics {
	r := &IlmPolicyStatistics{}

	return r
}

// false
