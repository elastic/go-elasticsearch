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
)

// DataStreamOptionsTemplate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/DataStreamOptions.ts#L36-L41
type DataStreamOptionsTemplate struct {
	FailureStore *DataStreamFailureStoreTemplate `json:"failure_store,omitempty"`
}

func (s *DataStreamOptionsTemplate) UnmarshalJSON(data []byte) error {

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

		case "failure_store":
			if err := dec.Decode(&s.FailureStore); err != nil {
				return fmt.Errorf("%s | %w", "FailureStore", err)
			}

		}
	}
	return nil
}

// NewDataStreamOptionsTemplate returns a DataStreamOptionsTemplate.
func NewDataStreamOptionsTemplate() *DataStreamOptionsTemplate {
	r := &DataStreamOptionsTemplate{}

	return r
}

type DataStreamOptionsTemplateVariant interface {
	DataStreamOptionsTemplateCaster() *DataStreamOptionsTemplate
}

func (s *DataStreamOptionsTemplate) DataStreamOptionsTemplateCaster() *DataStreamOptionsTemplate {
	return s
}
