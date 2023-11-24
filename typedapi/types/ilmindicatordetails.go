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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/lifecycleoperationmode"
)

// IlmIndicatorDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_global/health_report/types.ts#L150-L153
type IlmIndicatorDetails struct {
	IlmStatus lifecycleoperationmode.LifecycleOperationMode `json:"ilm_status"`
	Policies  int64                                         `json:"policies"`
}

func (s *IlmIndicatorDetails) UnmarshalJSON(data []byte) error {

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

		case "ilm_status":
			if err := dec.Decode(&s.IlmStatus); err != nil {
				return err
			}

		case "policies":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Policies = value
			case float64:
				f := int64(v)
				s.Policies = f
			}

		}
	}
	return nil
}

// NewIlmIndicatorDetails returns a IlmIndicatorDetails.
func NewIlmIndicatorDetails() *IlmIndicatorDetails {
	r := &IlmIndicatorDetails{}

	return r
}
