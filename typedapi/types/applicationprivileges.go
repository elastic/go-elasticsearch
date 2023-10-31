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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ApplicationPrivileges type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/security/_types/Privileges.ts#L26-L39
type ApplicationPrivileges struct {
	// Application The name of the application to which this entry applies.
	Application string `json:"application"`
	// Privileges A list of strings, where each element is the name of an application privilege
	// or action.
	Privileges []string `json:"privileges"`
	// Resources A list resources to which the privileges are applied.
	Resources []string `json:"resources"`
}

func (s *ApplicationPrivileges) UnmarshalJSON(data []byte) error {

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

		case "application":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Application = o

		case "privileges":
			if err := dec.Decode(&s.Privileges); err != nil {
				return err
			}

		case "resources":
			if err := dec.Decode(&s.Resources); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewApplicationPrivileges returns a ApplicationPrivileges.
func NewApplicationPrivileges() *ApplicationPrivileges {
	r := &ApplicationPrivileges{}

	return r
}
