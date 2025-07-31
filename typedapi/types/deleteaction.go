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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DeleteAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ilm/_types/Phase.ts#L149-L151
type DeleteAction struct {
	DeleteSearchableSnapshot *bool `json:"delete_searchable_snapshot,omitempty"`
}

func (s *DeleteAction) UnmarshalJSON(data []byte) error {

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

		case "delete_searchable_snapshot":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DeleteSearchableSnapshot", err)
				}
				s.DeleteSearchableSnapshot = &value
			case bool:
				s.DeleteSearchableSnapshot = &v
			}

		}
	}
	return nil
}

// NewDeleteAction returns a DeleteAction.
func NewDeleteAction() *DeleteAction {
	r := &DeleteAction{}

	return r
}

type DeleteActionVariant interface {
	DeleteActionCaster() *DeleteAction
}

func (s *DeleteAction) DeleteActionCaster() *DeleteAction {
	return s
}
