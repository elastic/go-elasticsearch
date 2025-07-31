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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ApplicationPrivilegesCheck type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/has_privileges/types.ts#L24-L32
type ApplicationPrivilegesCheck struct {
	// Application The name of the application.
	Application string `json:"application"`
	// Privileges A list of the privileges that you want to check for the specified resources.
	// It may be either application privilege names or the names of actions that are
	// granted by those privileges
	Privileges []string `json:"privileges"`
	// Resources A list of resource names against which the privileges should be checked.
	Resources []string `json:"resources"`
}

func (s *ApplicationPrivilegesCheck) UnmarshalJSON(data []byte) error {

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
				return fmt.Errorf("%s | %w", "Application", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Application = o

		case "privileges":
			if err := dec.Decode(&s.Privileges); err != nil {
				return fmt.Errorf("%s | %w", "Privileges", err)
			}

		case "resources":
			if err := dec.Decode(&s.Resources); err != nil {
				return fmt.Errorf("%s | %w", "Resources", err)
			}

		}
	}
	return nil
}

// NewApplicationPrivilegesCheck returns a ApplicationPrivilegesCheck.
func NewApplicationPrivilegesCheck() *ApplicationPrivilegesCheck {
	r := &ApplicationPrivilegesCheck{}

	return r
}
