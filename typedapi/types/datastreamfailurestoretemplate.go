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

// DataStreamFailureStoreTemplate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/DataStreamFailureStore.ts#L39-L54
type DataStreamFailureStoreTemplate struct {
	// Enabled If defined, it turns the failure store on/off (`true`/`false`) for this data
	// stream. A data stream failure store
	// that's disabled (enabled: `false`) will redirect no new failed indices to the
	// failure store; however, it will
	// not remove any existing data from the failure store.
	Enabled *bool `json:"enabled,omitempty"`
	// Lifecycle If defined, it specifies the lifecycle configuration for the failure store of
	// this data stream.
	Lifecycle *FailureStoreLifecycleTemplate `json:"lifecycle,omitempty"`
}

func (s *DataStreamFailureStoreTemplate) UnmarshalJSON(data []byte) error {

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

		case "enabled":
			if err := dec.Decode(&s.Enabled); err != nil {
				return fmt.Errorf("%s | %w", "Enabled", err)
			}

		case "lifecycle":
			if err := dec.Decode(&s.Lifecycle); err != nil {
				return fmt.Errorf("%s | %w", "Lifecycle", err)
			}

		}
	}
	return nil
}

// NewDataStreamFailureStoreTemplate returns a DataStreamFailureStoreTemplate.
func NewDataStreamFailureStoreTemplate() *DataStreamFailureStoreTemplate {
	r := &DataStreamFailureStoreTemplate{}

	return r
}

type DataStreamFailureStoreTemplateVariant interface {
	DataStreamFailureStoreTemplateCaster() *DataStreamFailureStoreTemplate
}

func (s *DataStreamFailureStoreTemplate) DataStreamFailureStoreTemplateCaster() *DataStreamFailureStoreTemplate {
	return s
}
