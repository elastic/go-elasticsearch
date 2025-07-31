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
)

// TranslogRetention type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/_types/IndexSettings.ts#L396-L415
type TranslogRetention struct {
	// Age This controls the maximum duration for which translog files are kept by each
	// shard. Keeping more
	// translog files increases the chance of performing an operation based sync
	// when recovering replicas. If
	// the translog files are not sufficient, replica recovery will fall back to a
	// file based sync. This setting
	// is ignored, and should not be set, if soft deletes are enabled. Soft deletes
	// are enabled by default in
	// indices created in Elasticsearch versions 7.0.0 and later.
	Age Duration `json:"age,omitempty"`
	// Size This controls the total size of translog files to keep for each shard.
	// Keeping more translog files increases
	// the chance of performing an operation based sync when recovering a replica.
	// If the translog files are not
	// sufficient, replica recovery will fall back to a file based sync. This
	// setting is ignored, and should not be
	// set, if soft deletes are enabled. Soft deletes are enabled by default in
	// indices created in Elasticsearch
	// versions 7.0.0 and later.
	Size ByteSize `json:"size,omitempty"`
}

func (s *TranslogRetention) UnmarshalJSON(data []byte) error {

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

		case "age":
			if err := dec.Decode(&s.Age); err != nil {
				return fmt.Errorf("%s | %w", "Age", err)
			}

		case "size":
			if err := dec.Decode(&s.Size); err != nil {
				return fmt.Errorf("%s | %w", "Size", err)
			}

		}
	}
	return nil
}

// NewTranslogRetention returns a TranslogRetention.
func NewTranslogRetention() *TranslogRetention {
	r := &TranslogRetention{}

	return r
}
