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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/lifecycleoperationmode"
)

// SlmIndicatorDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_global/health_report/types.ts#L160-L164
type SlmIndicatorDetails struct {
	Policies          int64                                         `json:"policies"`
	SlmStatus         lifecycleoperationmode.LifecycleOperationMode `json:"slm_status"`
	UnhealthyPolicies SlmIndicatorUnhealthyPolicies                 `json:"unhealthy_policies"`
}

func (s *SlmIndicatorDetails) UnmarshalJSON(data []byte) error {

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

		case "policies":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Policies", err)
				}
				s.Policies = value
			case float64:
				f := int64(v)
				s.Policies = f
			}

		case "slm_status":
			if err := dec.Decode(&s.SlmStatus); err != nil {
				return fmt.Errorf("%s | %w", "SlmStatus", err)
			}

		case "unhealthy_policies":
			if err := dec.Decode(&s.UnhealthyPolicies); err != nil {
				return fmt.Errorf("%s | %w", "UnhealthyPolicies", err)
			}

		}
	}
	return nil
}

// NewSlmIndicatorDetails returns a SlmIndicatorDetails.
func NewSlmIndicatorDetails() *SlmIndicatorDetails {
	r := &SlmIndicatorDetails{}

	return r
}
