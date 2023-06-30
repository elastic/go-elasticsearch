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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package explainlifecycle

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package explainlifecycle
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ilm/explain_lifecycle/ExplainLifecycleResponse.ts#L24-L28

type Response struct {
	Indices map[string]types.LifecycleExplain `json:"indices"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Indices: make(map[string]types.LifecycleExplain, 0),
	}
	return r
}

func (s *Response) UnmarshalJSON(data []byte) error {
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

		case "indices":
			if s.Indices == nil {
				s.Indices = make(map[string]types.LifecycleExplain, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]interface{})
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

				switch kind["managed"] {
				case true:
					oo := types.NewLifecycleExplainManaged()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Indices[key] = oo
				case false:
					oo := types.NewLifecycleExplainUnmanaged()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Indices[key] = oo
				default:
					if err := localDec.Decode(&s.Indices); err != nil {
						return err
					}
				}
			}

		}
	}
	return nil
}
