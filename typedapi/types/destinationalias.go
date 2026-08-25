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
// https://github.com/elastic/elasticsearch-specification/tree/abf9c2c6bb21328339daa197aae15af2ecbc46f0

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DestinationAlias type.
//
// https://github.com/elastic/elasticsearch-specification/blob/abf9c2c6bb21328339daa197aae15af2ecbc46f0/specification/transform/_types/Transform.ts#L64-L76
type DestinationAlias struct {
	// Alias The name of the alias.
	Alias string `json:"alias"`
	// MoveOnCreation Whether the destination index should be the only index in this alias. If
	// `true`, all the other indices will be removed from this alias before adding
	// the destination index to this alias. This does not delete the removed
	// indices; it only removes them from the alias.
	MoveOnCreation *bool `json:"move_on_creation,omitempty"`
}

func (s *DestinationAlias) UnmarshalJSON(data []byte) error {

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

		case "alias":
			if err := dec.Decode(&s.Alias); err != nil {
				return fmt.Errorf("%s | %w", "Alias", err)
			}

		case "move_on_creation":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MoveOnCreation", err)
				}
				s.MoveOnCreation = &value
			case bool:
				s.MoveOnCreation = &v
			}

		}
	}
	return nil
}

// NewDestinationAlias returns a DestinationAlias.
func NewDestinationAlias() *DestinationAlias {
	r := &DestinationAlias{}

	return r
}

type DestinationAliasVariant interface {
	DestinationAliasCaster() *DestinationAlias
}

func (s *DestinationAlias) DestinationAliasCaster() *DestinationAlias {
	return s
}
