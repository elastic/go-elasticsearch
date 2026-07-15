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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// GlobalRetention type.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/indices/get_data_lifecycle/_types/response.ts#L22-L25
type GlobalRetention struct {
	DefaultRetention Duration `json:"default_retention,omitempty"`
	MaxRetention     Duration `json:"max_retention,omitempty"`
}

func (s *GlobalRetention) UnmarshalJSON(data []byte) error {

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

		case "default_retention":
			if err := dec.Decode(&s.DefaultRetention); err != nil {
				return fmt.Errorf("%s | %w", "DefaultRetention", err)
			}

		case "max_retention":
			if err := dec.Decode(&s.MaxRetention); err != nil {
				return fmt.Errorf("%s | %w", "MaxRetention", err)
			}

		}
	}
	return nil
}

// NewGlobalRetention returns a GlobalRetention.
func NewGlobalRetention() *GlobalRetention {
	r := &GlobalRetention{}

	return r
}
