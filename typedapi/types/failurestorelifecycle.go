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

// FailureStoreLifecycle type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/_types/DataStreamFailureStore.ts#L56-L72
type FailureStoreLifecycle struct {
	// DataRetention If defined, every document added to this data stream will be stored at least
	// for this time frame.
	// Any time after this duration the document could be deleted.
	// When empty, every document in this data stream will be stored indefinitely.
	DataRetention Duration `json:"data_retention,omitempty"`
	// Enabled If defined, it turns data stream lifecycle on/off (`true`/`false`) for this
	// data stream. A data stream lifecycle
	// that's disabled (enabled: `false`) will have no effect on the data stream.
	Enabled *bool `json:"enabled,omitempty"`
}

func (s *FailureStoreLifecycle) UnmarshalJSON(data []byte) error {

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

		case "data_retention":
			if err := dec.Decode(&s.DataRetention); err != nil {
				return fmt.Errorf("%s | %w", "DataRetention", err)
			}

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = &value
			case bool:
				s.Enabled = &v
			}

		}
	}
	return nil
}

// NewFailureStoreLifecycle returns a FailureStoreLifecycle.
func NewFailureStoreLifecycle() *FailureStoreLifecycle {
	r := &FailureStoreLifecycle{}

	return r
}

type FailureStoreLifecycleVariant interface {
	FailureStoreLifecycleCaster() *FailureStoreLifecycle
}

func (s *FailureStoreLifecycle) FailureStoreLifecycleCaster() *FailureStoreLifecycle {
	return s
}
